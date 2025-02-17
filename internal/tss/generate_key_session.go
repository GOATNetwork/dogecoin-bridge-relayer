package tss

import (
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
)

var _ Session[any] = &GenerateKeySession[any, any, any]{}

type GenerateKeySession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func (g *GenerateKeySession[T, M, D]) Post(data *helper.ReceivedPartyState) {
	if data.IsBroadcast || g.Included(data.ToPartyIds) {
		g.sessionTransport.Post(data)
	}
}

func (m *Scheduler) NewGenerateKeySession(
	ec crypto.CurveType,
	proposalID ProposalID, // msg id
	sessionID party.SessionID,
	msg *Proposal,
) party.SessionID {
	allPartners := m.Participants()
	s := newSession[ProposalID, *Proposal, *LocalPartySaveData](
		m.ctx,
		ec,
		m.p2p,
		m,
		sessionID,
		"",
		m.Proposer(),
		proposalID,
		msg,
		types.GenKeySessionType,
		allPartners,
	)
	party, partyIdMap, endCh, errCh := RunKeyGen(s.ctx, m.isProd, ec, m.LocalSubmitter(), allPartners, s) // todo
	s.party = party
	s.partyIdMap = partyIdMap
	s.errCH = errCh
	s.endCh = endCh
	s.inToOut = m.senateInToOut
	s.Run()
	session := &GenerateKeySession[ProposalID, *Proposal, *LocalPartySaveData]{sessionTransport: s}
	m.AddSession(session)

	return session.SessionID()
}

func (m *Scheduler) JoinGenKeySession(msg SessionMessage[ProposalID, Proposal]) error {
	// check groupID
	if msg.GroupID != m.Participants().GroupID() {
		return fmt.Errorf("GenKeySessionType: %w", ErrGroupIdWrong)
	}
	// check msg
	unSignMsg := m.GenKeyProposal()
	if unSignMsg.String() != msg.Proposal.String() {
		return fmt.Errorf("GenKeyUnSignMsg: %w", ErrTaskSignatureMsgWrong)
	}

	ec := m.curveTypeBySenateSession(msg.SessionID)

	_ = m.NewGenerateKeySession(
		ec,
		msg.ProposalID,
		msg.SessionID,
		&msg.Proposal,
	)

	m.OpenSession(msg)

	return nil
}
