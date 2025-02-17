package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss"
	"github.com/nuvosphere/nudex-voter/internal/types"
	btcWallet "github.com/nuvosphere/nudex-voter/internal/wallet/btc"
	_ "github.com/nuvosphere/nudex-voter/internal/wallet/dog"
	"github.com/nuvosphere/nudex-voter/internal/wallet/evm"
	"github.com/nuvosphere/nudex-voter/internal/wallet/solana"
	"github.com/nuvosphere/nudex-voter/internal/wallet/sui"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	DatabaseManager *db.DatabaseManager
	State           *state.State
	modules         []Module
}

func NewApplication() *Application {
	dbm := db.NewDatabaseManager()
	stateDB := state.InitializeState(dbm)
	libP2PService := p2p.NewLibP2PService(stateDB, config.SubmitterPrivateKey)
	layer2Listener := layer2.NewLayer2Listener(libP2PService, stateDB.Bus(), dbm, config.AppConfig.MasterChainInfo())
	// btcListener := btc.NewBTCListener(libP2PService, stateDB, dbm)
	tssService := tss.NewTssService(libP2PService, dbm, stateDB.Bus(), layer2Listener)

	moules := []Module{
		layer2Listener,
		libP2PService,
		// btcListener,
		tssService,
	}

	for _, chain := range config.AppConfig.Chains {
		var module Module

		log.Debugf("moudle chain : %v", chain)

		switch chain.ChainType {
		case types.ChainBitcoin:
			module = btcWallet.NewWallet(
				stateDB.Bus(),
				tssService.TssService(),
				state.NewContractState(dbm.GetContractDB()),
				state.NewBtcWalletState(dbm.GetWalletDB(chain.ChainId), chain.ChainId),
				layer2Listener,
				&chain,
			)
		case types.ChainEthereum:
			module = evm.NewWallet(
				stateDB.Bus(),
				tssService.TssService(),
				layer2Listener,
				state.NewContractState(dbm.GetContractDB()),
				state.NewEvmWalletState(dbm.GetWalletDB(chain.ChainId)),
				&chain,
			)

		case types.CoinTypeSOL:
			module = solana.NewWallet(
				stateDB.Bus(),
				tssService.TssService(),
				state.NewContractState(dbm.GetContractDB()),
				state.NewSolWalletState(dbm.GetWalletDB(chain.ChainId)),
				layer2Listener,
			)

		case types.ChainSui:
			module = sui.NewWallet(
				stateDB.Bus(),
				tssService.TssService(),
				state.NewContractState(dbm.GetContractDB()),
				state.NewSuiWalletState(dbm.GetWalletDB(chain.ChainId)),
				layer2Listener,
			)
		default:
			panic(fmt.Sprintf("unknown chain type: %d", chain.ChainType))
		}

		moules = append(moules, module)
	}

	return &Application{
		DatabaseManager: dbm,
		State:           stateDB,
		modules:         moules,
	}
}

func (app *Application) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go parallel.ForEach(app.modules, func(module Module, _ int) { module.Start(ctx) })

	<-stop
	log.Info("Receiving exit signal...")

	cancel()
	lo.ForEach(app.modules, func(module Module, _ int) { module.Stop(ctx) })

	app.State.Bus().Close()
	log.Info("Server stopped")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// app := NewApplication()
	// app.Run()
	Execute()
}

type Module interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}
