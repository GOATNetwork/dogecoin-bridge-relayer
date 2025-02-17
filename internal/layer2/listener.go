package layer2

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync/atomic"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Layer2Listener struct {
	p2p                   *p2p.Service
	db                    *db.DatabaseManager
	eventBus              eventbus.Bus
	ethClient             *ethclient.Client
	chainID               atomic.Int64
	contractAddress       []common.Address
	addressBind           map[common.Address]func(types.Log) error
	contractVotingManager *contracts.VotingManagerContract
	participantManager    *contracts.ParticipantManagerContract
	taskManager           *contracts.TaskManagerContract
	accountManager        *contracts.AccountManagerContract
	isSyncing             *atomic.Bool
	chainInfo             *config.Chain
}

func NewLayer2Listener(p *p2p.Service, eventBus eventbus.Bus, db *db.DatabaseManager, chainInfo *config.Chain) *Layer2Listener {
	ethClient, err := DialEthClient(chainInfo)
	if err != nil {
		log.Fatalf("Error creating Layer2 EVM RPC client: %v", err)
	}

	isSyncing := &atomic.Bool{}
	isSyncing.Store(true)
	self := &Layer2Listener{
		p2p:       p,
		db:        db,
		eventBus:  eventBus,
		ethClient: ethClient,
		chainID:   atomic.Int64{},
		isSyncing: isSyncing,
		chainInfo: chainInfo,
	}

	var (
		VotingAddress      = common.HexToAddress(config.AppConfig.Contract.Voter)
		AccountAddress     = common.HexToAddress(config.AppConfig.Contract.Account)
		TaskAddress        = common.HexToAddress(config.AppConfig.Contract.TaskManager)
		ParticipantAddress = common.HexToAddress(config.AppConfig.Contract.Participant)
		DepositAddress     = common.HexToAddress(config.AppConfig.Contract.Deposit)
		AssetAddress       = common.HexToAddress(config.AppConfig.Contract.AssetHandler)
	)

	self.addressBind = map[common.Address]func(types.Log) error{
		VotingAddress:      self.processVotingLog,
		AccountAddress:     self.processAccountLog,
		ParticipantAddress: self.processParticipantLog,
		TaskAddress:        self.processTaskLog,
		DepositAddress:     self.processDepositLog,
		AssetAddress:       self.processAssetLog,
	}
	self.contractAddress = lo.MapToSlice(
		self.addressBind,
		func(item common.Address, _ func(log2 types.Log) error) common.Address { return item },
	)

	var errs []error

	chainId := self.ChainID(context.Background())

	if chainId.Int64() != config.MasterChainId.Int64() {
		err = fmt.Errorf("ChainID mismatch: expected %d, got %d", config.MasterChainId.Int64(), chainId.Int64())
		errs = append(errs, err)
	}

	errs = append(errs, err)
	contractVotingManager, err := contracts.NewVotingManagerContract(VotingAddress, ethClient)
	errs = append(errs, err)
	participantManager, err := contracts.NewParticipantManagerContract(ParticipantAddress, ethClient)
	errs = append(errs, err)
	taskManager, err := contracts.NewTaskManagerContract(TaskAddress, ethClient)
	errs = append(errs, err)
	accountManager, err := contracts.NewAccountManagerContract(TaskAddress, ethClient)
	errs = append(errs, err)

	utils.Assert(errors.Join(errs...))

	self.taskManager = taskManager
	self.contractVotingManager = contractVotingManager
	self.participantManager = participantManager
	self.accountManager = accountManager

	return self
}

// New an eth client.
func DialEthClient(chainInfo *config.Chain) (*ethclient.Client, error) {
	var opts []rpc.ClientOption

	if chainInfo.Rpc.JwtSecret != "" {
		jwtSecret := common.FromHex(strings.TrimSpace(chainInfo.Rpc.JwtSecret))
		if len(jwtSecret) != 32 {
			return nil, errors.New("jwt secret is not a 32 bytes hex string")
		}

		var jwtKey [32]byte

		copy(jwtKey[:], jwtSecret)
		opts = append(opts, rpc.WithHTTPAuth(node.NewJWTAuth(jwtKey)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Dial the Ethereum node with optional JWT authentication
	client, err := rpc.DialOptions(ctx, chainInfo.Rpc.Url, opts...)
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(client), nil
}

func (l *Layer2Listener) Stop(context.Context) {
	log.Infof("Stoped layer2 listener")
}

func (l *Layer2Listener) Start(ctx context.Context) {
	// Get latest sync height
	var syncStatus db.EVMSyncStatus

	relayerDB := l.db.GetContractDB()

	result := relayerDB.First(&syncStatus)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		syncStatus.LastSyncBlock = uint64(l.chainInfo.StartHeight)
		syncStatus.UpdatedAt = time.Now()
		relayerDB.Create(&syncStatus)
	} else if result.Error != nil {
		log.Fatalf("Error querying sync status: %v", result.Error)
	}

	ticker := time.NewTicker(l.chainInfo.ScanInterval * time.Second)
	defer ticker.Stop()

	log.Infof("Layer2Listener: begin scan log: begin height: %v", syncStatus.LastSyncBlock)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		L:
			for {
				isContinue, err := l.scan(ctx, &syncStatus)
				if err != nil {
					log.Errorf("scan : %v", err)
				}
				if !isContinue {
					l.isSyncing.Store(false)
					log.Infof("Layer2Listener: end scan log: end height: %v", syncStatus.LastSyncBlock)
					l.eventBus.Publish(eventbus.EventTask{}, eventbus.EventSynced{})
					break L
				} else {
					l.isSyncing.Store(true)
				}
			}
		}
	}
}

func (l *Layer2Listener) scan(ctx context.Context, syncStatus *db.EVMSyncStatus) (isContinue bool, err error) {
	latestBlock, err := l.ethClient.BlockNumber(ctx)
	if err != nil {
		return false, fmt.Errorf("error getting latest block number: %w", err)
	}

	var targetBlock uint64
	if latestBlock >= uint64(l.chainInfo.Confirmations) {
		targetBlock = latestBlock - uint64(l.chainInfo.Confirmations)
	} else {
		log.Infof("latestBlock is less than L2Confirmations, setting targetBlock to 0")

		targetBlock = 0
	}

	log.Infof("syncStatus.LastSyncBlock: %v, targetBlock: %v", syncStatus.LastSyncBlock, targetBlock)

	if syncStatus.LastSyncBlock < targetBlock {
		fromBlock := syncStatus.LastSyncBlock + 1

		toBlock := fromBlock + uint64(l.chainInfo.MaxBlockRange) - 1
		if toBlock > targetBlock {
			toBlock = targetBlock
		}

		log.WithFields(log.Fields{"fromBlock": fromBlock, "toBlock": toBlock}).Info("Syncing L2 nudex events")

		filterQuery := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: l.contractAddress,
			// Topics:    batch,
		}

		var logs []types.Log

		err = backoff.Retry(
			func() error {
				logs, err = l.ethClient.FilterLogs(context.Background(), filterQuery)
				return err
			},
			backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5),
		)
		if err != nil {
			return false, fmt.Errorf("failed to filter logs: %w", err)
		}

		for _, vLog := range logs {
			l.processLogs(vLog)
		}

		// Save sync status
		syncStatus.LastSyncBlock = toBlock
		syncStatus.UpdatedAt = time.Now()
		l.db.GetContractDB().Save(syncStatus)

		return true, nil
	}

	return false, nil
}

// stop ctx.
//
//lint:ignore U1000 Ignore unused function
func (l *Layer2Listener) stop() {}

func (l *Layer2Listener) ChainID(ctx context.Context) *big.Int {
	if l.chainID.Load() == 0 {
		chainID, err := l.ethClient.ChainID(ctx)
		if err != nil {
			l.chainID.Store(int64(l.chainInfo.ChainId))
		} else {
			l.chainID.Store(chainID.Int64())
		}
	}

	return big.NewInt(l.chainID.Load())
}

func (l *Layer2Listener) IsSyncing() bool {
	return l.isSyncing.Load()
}
