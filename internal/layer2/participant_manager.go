package layer2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

type ParticipantManager interface {
	Participants() (types.Participants, error)
	IsParticipant(participant common.Address) (bool, error)
	GetRandomParticipant(_salt common.Address) (common.Address, error)
}

func (l *Layer2Listener) Participants() (types.Participants, error) {
	return l.participantManager.GetParticipants(nil)
}

func (l *Layer2Listener) IsParticipant(participant common.Address) (bool, error) {
	return l.participantManager.IsParticipant(nil, participant)
}

func (l *Layer2Listener) GetRandomParticipant(_salt common.Address) (common.Address, error) {
	return l.participantManager.GetRandomParticipant(nil, _salt)
}
