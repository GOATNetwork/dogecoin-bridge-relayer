package state

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
)

type State struct {
	eventBus eventbus.Bus
	dbm      *db.DatabaseManager
}

func (s *State) Bus() eventbus.Bus {
	return s.eventBus
}

// InitializeState initializes the state by reading from the DB.
func InitializeState(dbm *db.DatabaseManager) *State {
	return &State{
		eventBus: eventbus.NewBus(),
		dbm:      dbm,
	}
}
