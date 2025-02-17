package sui

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
	state     *state.SuiWalletState
	tss       suite.TssService
	txContext sync.Map // taskID:TxContext
	client    *SuiClient
	chainInfo *config.Chain
}

func NewWallet(
	bus eventbus.Bus,
	tss suite.TssService,
	stateDB *state.ContractState,
	state *state.SuiWalletState,
	voter layer2.Voter,
) *WalletClient {
	ctx, cancel := context.WithCancel(context.Background())
	base := wallet.NewBaseWallet(bus, stateDB, voter, types.ChainIdSui)

	var c *SuiClient

	if base.IsProd() {
		c = NewSuiClient(ctx)
	} else {
		c = NewDevClient()
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
	log.Info("sui wallet client is starting...")
	w.tss.RegisterTssClient(w)
	w.receiveL2TaskLoop()
}

func (w *WalletClient) Stop(context.Context) {
	w.cancel()
}
