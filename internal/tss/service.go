package tss

import (
	"context"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	scheduler *Scheduler
}

func (s Service) TssService() suite.TssService {
	return s.scheduler
}

func NewTssService(p p2p.P2PService, dbm *db.DatabaseManager, bus eventbus.Bus, voter layer2.Voter) *Service {
	scheduler := NewScheduler(

		true,
		p,
		bus,
		state.NewContractState(dbm.GetContractDB()),
		voter,
		crypto.PubkeyToAddress(config.SubmitterPrivateKey.PublicKey),
	)

	return &Service{
		scheduler: scheduler,
	}
}

func (t *Service) Start(ctx context.Context) {
	t.scheduler.Start()

	<-ctx.Done()
	log.Info("TSSService is stopping...")
}

func (t *Service) Stop(ctx context.Context) {
	t.scheduler.Stop()
}
