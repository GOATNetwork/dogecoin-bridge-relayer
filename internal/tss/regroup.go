package tss

import (
	"fmt"
	"slices"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

func (m *Scheduler) reGroupResultLoop() {
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("reGroup result loop stopping...")
			case sessionResult := <-m.senateInToOut:
				m.SaveSenateSessionResult(sessionResult)
				m.ecCount.Add(-1)

				if m.ecCount.Load() == 0 {
					newGroup := m.newGroup.Swap(nullNewGroup).(*NewGroup)
					m.partners.Store(newGroup.NewParts)
					m.p2p.UpdateParticipants(newGroup.NewParts)
					log.Infof("regroup success!!!: new groupID: %v", newGroup.NewParts.GroupID())
				}
			}
		}
	}()
}

func (m *Scheduler) processReGroupProposal(v *types.ParticipantEvent) {
	newParts := types.Participants{}
	oldParts := m.Participants()

	log.Debugf("ParticipantEvent: %v, address: %v", v.EventName, v.Address)

	switch v.EventName {
	case layer2.ParticipantAdded:
		joinAddress := v.Address[0]
		if !slices.Contains(oldParts, joinAddress) {
			newParts = append(oldParts, joinAddress)
		}
	case layer2.ParticipantRemoved:
		joinAddress := v.Address[0]
		newParts = lo.Filter(oldParts, func(item common.Address, index int) bool { return item != joinAddress })
	case layer2.ParticipantReset:
		newParts = v.Address
	}

	log.Debugf("newParts: %v, oldParts: %v", newParts, oldParts)

	if len(newParts) > 0 && newParts.GroupID() != oldParts.GroupID() {
		m.newGroup.Store(&NewGroup{
			Event:    v,
			NewParts: newParts,
			OldParts: oldParts,
		})

		if m.IsCanProposal() {
			for {
				_, have := lo.Find(newParts, func(item common.Address) bool {
					return !m.p2p.IsOnline(item.String())
				})

				if !have {
					break
				}

				time.Sleep(1 * time.Second)
			}

			_ = m.NewReShareGroupSession(
				crypto.ECDSA,
				SenateSessionIDOfECDSA,
				SenateProposalIDOfECDSA,
				SenateProposal,
				oldParts,
				newParts,
			)
			_ = m.NewReShareGroupSession(
				crypto.EDDSA,
				SenateSessionIDOfEDDSA,
				SenateProposalIDOfEDDSA,
				SenateProposal,
				oldParts,
				newParts,
			)

			log.Info("Leader NewReShareGroupSession stared")
		} else {
			log.Info("Candidate NewReShareGroupSession stared")
		}
	}
}

func (m *Scheduler) isReShareGroup() bool {
	newGroup := m.newGroup.Load().(*NewGroup)
	if newGroup != nullNewGroup {
		return true
	}

	// get latest participants compare local participants
	partners, err := m.voter.Participants()
	utils.Assert(err)

	old := m.Participants()
	if old.GroupID() != partners.GroupID() {
		g := &NewGroup{
			NewParts: partners,
			OldParts: old,
		}
		m.newGroup.Store(g)

		return true
	}

	return false
}

func (m *Scheduler) JoinReShareGroupSession(msg SessionMessage[ProposalID, Proposal]) error {
	// todo How find new part?
	is := m.isReShareGroup()
	if !is {
		return fmt.Errorf("not new group")
	}

	newGroup := m.newGroup.Load().(*NewGroup)
	// check groupID
	if msg.GroupID != newGroup.NewParts.GroupID() {
		return fmt.Errorf("JoinReShareGroupSession: session id: %v, msg.SessionID, %w", msg.SessionID, ErrGroupIdWrong)
	}
	// check msg
	unSignMsg := m.ReShareGroupProposal()
	if unSignMsg.Cmp(&msg.Proposal) != 0 {
		return fmt.Errorf("JoinReShareGroupSession: proposal error, session id: %v", msg.SessionID)
	}

	ec := m.curveTypeBySenateSession(msg.SessionID)

	_ = m.NewReShareGroupSession(
		ec,
		msg.SessionID,
		msg.ProposalID,
		&msg.Proposal,
		newGroup.OldParts,
		newGroup.NewParts,
	)

	m.OpenSession(msg)

	return nil
}
