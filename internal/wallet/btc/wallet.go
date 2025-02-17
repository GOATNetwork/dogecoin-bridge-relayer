package btc

import (
	"context"
	"sync"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	log "github.com/sirupsen/logrus"
)

type WalletClient struct {
	*wallet.BaseWallet
	client           *Client
	ctx              context.Context
	cancel           context.CancelFunc
	state            *state.BtcWalletState
	tss              suite.TssService
	txContext        sync.Map // taskID:TxContext
	chainInfo        *config.Chain
	lastScannedBlock uint64
}

func NewWallet(
	bus eventbus.Bus,
	tss suite.TssService,
	stateDB *state.ContractState,
	state *state.BtcWalletState,
	voter layer2.Voter,
	chainInfo *config.Chain,
) *WalletClient {
	ctx, cancel := context.WithCancel(context.Background())

	return &WalletClient{
		BaseWallet: wallet.NewBaseWallet(bus, stateDB, voter, chainInfo.ChainId),
		ctx:        ctx,
		cancel:     cancel,
		state:      state,
		tss:        tss,
		txContext:  sync.Map{},
		client:     NewClient(ctx, time.Minute, &chaincfg.MainNetParams, chainInfo),
		chainInfo:  chainInfo,
	}
}

func (w *WalletClient) Start(context.Context) {
	log.Info("btc wallet client is starting...")

	go w.ScanBlocks()
	w.tss.RegisterTssClient(w)
	w.receiveL2TaskLoop()
}

func (w *WalletClient) Stop(context.Context) {
	log.Info("btc wallet client is stopping...")
	w.cancel()
}
