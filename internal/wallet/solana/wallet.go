package solana

import (
	"context"
	"sync"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	log "github.com/sirupsen/logrus"
)

type WalletClient struct {
	*wallet.BaseWallet
	ctx       context.Context
	cancel    context.CancelFunc
	event     eventbus.Bus
	state     *state.SolWalletState
	tss       suite.TssService
	client    *SolClient
	txContext sync.Map // taskID:TxContext
	chainInfo *config.Chain
}

func NewWallet(
	bus eventbus.Bus,
	tss suite.TssService,
	stateDB *state.ContractState,
	state *state.SolWalletState,
	voter layer2.Voter,
) *WalletClient {
	ctx, cancel := context.WithCancel(context.Background())

	base := wallet.NewBaseWallet(bus, stateDB, voter, types.ChainIdSolana)

	var c *SolClient

	if base.IsProd() {
		c = NewSolClient()
	} else {
		c = NewDevSolClient()
	}

	return &WalletClient{
		BaseWallet: base,
		ctx:        ctx,
		cancel:     cancel,
		state:      state,
		tss:        tss,
		client:     c,
	}
}

func (w *WalletClient) Start(context.Context) {
	log.Info("solana wallet client is starting...")
	w.tss.RegisterTssClient(w)
	w.receiveL2TaskLoop()
}

func (w *WalletClient) Stop(context.Context) {
	w.cancel()
}
