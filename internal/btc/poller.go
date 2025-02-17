package btc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SigHashQueue struct {
	Start  uint64
	Count  int
	Status bool
	Id     string
}

func NewSigHashQueue() *SigHashQueue {
	return &SigHashQueue{
		Start:  0,
		Count:  0,
		Status: false,
		Id:     "",
	}
}

type BTCPoller struct {
	db          *gorm.DB
	state       *state.State
	confirmChan chan *types.BtcBlockExt

	sigFailChan    chan interface{}
	sigFinishChan  chan interface{}
	sigTimeoutChan chan interface{}

	sigHashQueue *SigHashQueue
	sigHashMu    sync.Mutex //lint:ignore U1000 Ignore unused
}

func NewBTCPoller(state *state.State, db *gorm.DB) *BTCPoller {
	return &BTCPoller{
		state:       state,
		db:          db,
		confirmChan: make(chan *types.BtcBlockExt, 64),

		sigFailChan:    make(chan interface{}, 10),
		sigFinishChan:  make(chan interface{}, 10),
		sigTimeoutChan: make(chan interface{}, 10),

		sigHashQueue: NewSigHashQueue(),
	}
}

func (p *BTCPoller) Start(ctx context.Context) {
	go p.pollLoop(ctx)
}

func (p *BTCPoller) Stop() {
}

func (p *BTCPoller) pollLoop(ctx context.Context) {
	for {
		select {
		case block := <-p.confirmChan:
			p.handleConfirmedBlock(block)
		case <-ctx.Done():
			log.Info("Stopping the polling of confirmed blocks...")
			return
		}
	}
}

func (p *BTCPoller) GetBlockHashForTx(txHash chainhash.Hash) (*chainhash.Hash, error) {
	var btcTxOutput db.BtcTXOutput

	if err := p.db.Where("tx_hash = ?", txHash.String()).First(&btcTxOutput).Error; err != nil {
		return nil, fmt.Errorf("failed to find the block hash for the transaction: %w", err)
	}

	blockHashBytes := btcTxOutput.PkScript[:32] // Assuming the block hash is the first 32 bytes of PkScript

	blockHash, err := chainhash.NewHash(blockHashBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to create hash from block hash bytes: %w", err)
	}

	return blockHash, nil
}

func (p *BTCPoller) GetBlockHeader(blockHash *chainhash.Hash) (*wire.BlockHeader, error) {
	var blockData db.BtcBlockData
	if err := p.db.Where("block_hash = ?", blockHash.String()).First(&blockData).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve block header from database: %w", err)
	}

	header := wire.BlockHeader{}

	err := header.Deserialize(bytes.NewReader(blockData.Header))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize block header: %w", err)
	}

	return &header, nil
}

func (p *BTCPoller) GetTxHashes(blockHash *chainhash.Hash) ([]chainhash.Hash, error) {
	var txHashes []chainhash.Hash

	var blockData db.BtcBlockData
	if err := p.db.Where("block_hash = ?", blockHash.String()).First(&blockData).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve block data from database: %w", err)
	}

	err := json.Unmarshal([]byte(blockData.TxHashes), &txHashes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal transaction hash list: %w", err)
	}

	return txHashes, nil
}

func (p *BTCPoller) GetBlock(height uint64) (*db.BtcBlockData, error) {
	var blockData db.BtcBlockData
	if err := p.db.Where("block_height = ?", height).First(&blockData).Error; err != nil {
		return nil, fmt.Errorf("error retrieving block from database: %w", err)
	}

	return &blockData, nil
}

func (p *BTCPoller) handleConfirmedBlock(block *types.BtcBlockExt) {
	// Logic for handling confirmed blocks
	panic("todo")
}
