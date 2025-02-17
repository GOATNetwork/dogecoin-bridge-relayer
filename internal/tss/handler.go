package tss

import (
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ErrTaskNotFound          = gorm.ErrRecordNotFound
	ErrTaskCompleted         = fmt.Errorf("task already completed")
	ErrTaskOrderInconsistent = fmt.Errorf("order of the task is inconsistent")
	ErrTaskIdWrong           = fmt.Errorf("taskId is wrong")
	ErrTaskSignatureMsgWrong = fmt.Errorf("task signature msg is wrong")
	ErrGroupIdWrong          = fmt.Errorf("groupId is wrong")
	ErrSessionIdWrong        = fmt.Errorf("sessionId is wrong")
	ErrProposerWrong         = fmt.Errorf("task proposer is wrong")
)

func (m *Scheduler) Validate(msg SessionMessage[ProposalID, Proposal]) error {
	if m.IsDiscussed(msg.SeqId) {
		return fmt.Errorf("taskID:%v, %w", msg.ProposalID, ErrTaskCompleted)
	}

	if msg.Proposer != m.Proposer() {
		return fmt.Errorf("proposer:(%v, %v), %w", msg.Proposer, m.Proposer(), ErrProposerWrong)
	}

	return nil
}

func (m *Scheduler) GenKeyProposal() Proposal {
	return *SenateProposal
}

func (m *Scheduler) ReShareGroupProposal() Proposal {
	return *SenateProposal
}

func (m *Scheduler) isSenateSession(sessionID party.SessionID) bool {
	return sessionID == SenateSessionIDOfECDSA || sessionID == SenateSessionIDOfEDDSA
}

func (m *Scheduler) curveTypeBySenateSession(sessionID party.SessionID) crypto.CurveType {
	switch sessionID {
	case SenateSessionIDOfEDDSA:
		return crypto.EDDSA
	case SenateSessionIDOfECDSA:
		return crypto.ECDSA
	default:
		panic("unimplemented")
	}
}

func (m *Scheduler) OpenSession(msg SessionMessage[ProposalID, Proposal]) bool {
	session := m.GetSession(msg.SessionID)
	if session != nil {
		if !session.Equal(msg.FromPartyId) { // not from self
			from := session.PartyID(msg.FromPartyId)
			if from != nil {
				session.Post(msg.State(from))
			} else {
				if session.Included(msg.ToPartyIds) {
					log.Errorf("session is nil, but included: %v", msg.SessionID)
				}

				if !m.isSenateSession(msg.SessionID) {
					// panic
					panic(fmt.Errorf("session from not is exist:%v basePath: %v", msg.FromPartyId, m.partyData.basePath))
				}

				log.Errorf("session from not is exist:%v basePath: %v", msg.FromPartyId, m.partyData.basePath)
			}
		}

		return true
	}

	return false
}

// processReceivedProposal handler received msg from other node.
func (m *Scheduler) processReceivedProposal(msg SessionMessage[ProposalID, Proposal]) error {
	log.Debugf("process received proposal id: %v, basePath: %v", msg.SeqId, m.partyData.basePath)

	// todo
	//err := m.Validate(msg)
	//if err != nil {
	//	return err
	//}

	ok := m.OpenSession(msg)
	if ok {
		log.Debugf("open session success, sessionID: %v", msg.SessionID)
		return nil
	}

	log.Debugf("open session fail: session id: %v, msg type: %v,", msg.SessionID, msg.Type)

	var err error
	// build new session
	switch msg.Type {
	case types.GenKeySessionType:
		err = m.JoinGenKeySession(msg)
	case types.ReShareGroupSessionType:
		err = m.JoinReShareGroupSession(msg)
	case types.SignOperationSessionType, types.SignTxSessionType:
		err = m.JoinSignSession(msg)
	case types.SignTaskSessionType: // only used test
		task, errTask := m.GetTask(msg.SeqId)
		if errTask != nil {
			return errTask
		}

		err = m.joinSignTaskSession(msg, task)
	default:
		err = fmt.Errorf("unknown msg type: %v, msg: %v", msg.Type, msg)
	}

	if err != nil {
		return err
	}

	ok = m.OpenSession(msg)
	if !ok {
		log.Debug("session not is exist")
	}

	return nil
}
