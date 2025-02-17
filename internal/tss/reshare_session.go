package tss

import (
	"context"
	"fmt"
	"strings"

	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	ecdsaResharing "github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	eddsaResharing "github.com/bnb-chain/tss-lib/v2/eddsa/resharing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

var _ Session[any] = &ReShareGroupSession[any, any, any]{}

type JoinReShareGroupSession[T, M, D any] struct {
	*sessionTransport[T, M, D]
	partyIdMap map[string]*tss.PartyID
	isNew      bool
}

func (r *JoinReShareGroupSession[T, M, D]) Post(state *helper.ReceivedPartyState) {
	if len(state.ToPartyIds) == 0 {
		log.Error("did not expect a msg to have a nil destination during resharing")
		return
	}

	if r.Included(state.ToPartyIds) {
		if state.IsToOldAndNewCommittees {
			r.sessionTransport.Post(state)
			return
		}

		if (!state.IsToOldCommittee && r.isNew) || (state.IsToOldCommittee && !r.isNew) {
			r.sessionTransport.Post(state)
		}
	}
}

func (r *JoinReShareGroupSession[T, M, D]) PartyID(id string) *tss.PartyID {
	pid := r.partyIdMap[strings.ToLower(id)]
	if pid != nil {
		return pid
	}

	return r.sessionTransport.PartyID(id)
}

type ReShareGroupSession[T, M, D any] struct {
	oldSession *sessionTransport[T, M, D]
	newSession *sessionTransport[T, M, D]
}

func createReShareParam(
	ctx context.Context,
	params *tss.ReSharingParameters,
	key LocalPartySaveData,
) (tss.Party, chan *LocalPartySaveData, chan tss.Message, chan *tss.Error) {
	// outgoing messages to other peers
	outCh := make(chan tss.Message, 100000)
	// error if reshare fails, contains culprits to blame
	errCh := make(chan *tss.Error, 256)

	log.Debug("creating new local party")

	outEndCh := make(chan *LocalPartySaveData, 100000)
	// output data when keygen finished
	ecdsaEndCh := make(chan *ecdsaKeygen.LocalPartySaveData, 256)
	eddsaEndCh := make(chan *eddsaKeygen.LocalPartySaveData, 256)

	var party tss.Party

	switch key.CurveType() {
	case crypto.ECDSA:
		data := key.ECDSAData()
		party = ecdsaResharing.NewLocalParty(params, *data, outCh, ecdsaEndCh)

	case crypto.EDDSA:
		party = eddsaResharing.NewLocalParty(params, *key.EDDSAData(), outCh, eddsaEndCh)

	default:
		panic("implement me")
	}

	go func() {
		defer close(eddsaEndCh)
		defer close(ecdsaEndCh)
		select {
		case <-ctx.Done():
			return
		case data := <-ecdsaEndCh:
			outEndCh <- BuildECDSALocalPartySaveData().SetData(data)
		case data := <-eddsaEndCh:
			outEndCh <- BuildEDDSALocalPartySaveData().SetData(data)
		}
	}()

	log.Debug("local resharing party created", "partyID", party.PartyID())

	return party, outEndCh, outCh, errCh
}

func runReShareParty(ctx context.Context, transport helper.Transporter, party tss.Party, outCh chan tss.Message, errCh chan *tss.Error) {
	helper.RunParty(ctx, party, errCh, outCh, transport, true)
}

func (m *Scheduler) NewReShareGroupSession(
	ec crypto.CurveType,
	sessionID party.SessionID,
	proposalID ProposalID, // msg id
	msg *Proposal,
	oldPartners types.Participants,
	newPartners types.Participants,
) party.SessionID {
	m.ecCount.Add(1)
	localSubmitter := m.LocalSubmitter()
	signer := "" // todo

	log.Debugf("newPartners:%v, oldPartners: %v", newPartners, oldPartners)
	oldPartyIDs := createPartyIDsByGroupWithAlias(ec, oldPartners, "old: "+m.partyData.basePath)
	oldPeerCtx := tss.NewPeerContext(oldPartyIDs) // todo

	newPartyIDs := createPartyIDsByGroupWithAlias(ec, newPartners, "new: "+m.partyData.basePath)
	newPeerCtx := tss.NewPeerContext(newPartyIDs)
	newPartKey := PartyKey(ec, newPartners, localSubmitter)
	newPartyID := newPartyIDs.FindByKey(newPartKey)

	log.Debugf("newPartyID: %v", newPartyID)

	newPartyIdMap := lo.SliceToMap(newPartyIDs, func(item *tss.PartyID) (string, *tss.PartyID) { return strings.ToLower(item.Id), item })
	newParams := tss.NewReSharingParameters(
		ec.EC(),
		oldPeerCtx,
		newPeerCtx,
		newPartyID,
		oldPartners.Len(),
		oldPartners.Threshold(),
		newPartners.Len(),
		newPartners.Threshold(),
	)
	oldPartyIdMap := lo.SliceToMap(oldPartyIDs, func(item *tss.PartyID) (string, *tss.PartyID) { return strings.ToLower(item.Id), item })
	oldPartKey := PartyKey(ec, oldPartners, localSubmitter)
	oldPartyID := oldPartyIDs.FindByKey(oldPartKey)
	oldParams := tss.NewReSharingParameters(
		ec.EC(),
		oldPeerCtx,
		newPeerCtx,
		oldPartyID,
		oldPartners.Len(),
		oldPartners.Threshold(),
		newPartners.Len(),
		newPartners.Threshold(),
	)

	joinedNew := newPartners.Contains(m.LocalSubmitter())
	joinedOld := oldPartners.Contains(m.LocalSubmitter())
	// 1. joined new group
	if joinedNew && !joinedOld {
		log.Debugf("new join node: oldPartyIDs: %v,newPartyIDs:%v", oldPartyIDs, newPartyIDs)

		newInnerSession := newSession[ProposalID, *Proposal, *LocalPartySaveData](
			m.ctx,
			ec,
			m.p2p,
			m,
			sessionID,
			signer,
			localSubmitter,
			proposalID,
			msg,
			types.ReShareGroupSessionType,
			newPartners,
		)
		newInnerSession.partyIdMap = newPartyIdMap
		newInnerSession.inToOut = m.senateInToOut
		localData := m.partyData.GenerateNewLocalPartySaveData(ec, newPartners)
		party, outEndCh, outCh, errCh := createReShareParam(newInnerSession.ctx, newParams, *localData)
		newInnerSession.party = party
		newInnerSession.endCh = outEndCh
		newInnerSession.errCH = errCh

		joinSession := &JoinReShareGroupSession[ProposalID, *Proposal, *LocalPartySaveData]{
			sessionTransport: newInnerSession,
			partyIdMap:       oldPartyIdMap,
			isNew:            true,
		}
		m.AddSession(joinSession)
		joinSession.Run()

		runReShareParty(newInnerSession.ctx, joinSession, party, outCh, errCh)

		return joinSession.SessionID()
	}

	// 2. joined old group
	if !joinedNew && joinedOld {
		log.Debugf("remove node: oldPartyIDs: %v,newPartyIDs:%v", oldPartyIDs, newPartyIDs)

		oldInnerSession := newSession[ProposalID, *Proposal, *LocalPartySaveData](
			m.ctx,
			ec,
			m.p2p,
			m,
			sessionID,
			signer,
			localSubmitter,
			proposalID,
			msg,
			types.ReShareGroupSessionType,
			newPartners, // todo
		)
		oldInnerSession.partyIdMap = oldPartyIdMap
		oldInnerSession.inToOut = m.senateInToOut

		localData := m.partyData.GetData(ec)
		party, outEndCh, outCh, errCh := createReShareParam(oldInnerSession.ctx, oldParams, *localData)
		oldInnerSession.party = party
		oldInnerSession.endCh = outEndCh
		oldInnerSession.errCH = errCh
		joinSession := &JoinReShareGroupSession[ProposalID, *Proposal, *LocalPartySaveData]{
			sessionTransport: oldInnerSession,
			partyIdMap:       newPartyIdMap,
			isNew:            false,
		}
		m.AddSession(joinSession)
		joinSession.Run()

		runReShareParty(oldInnerSession.ctx, joinSession, party, outCh, errCh)

		return joinSession.SessionID()
	}

	// 3. both joined old and new group
	reShareSession := &ReShareGroupSession[ProposalID, *Proposal, *LocalPartySaveData]{}
	oldInnerSession := newSession[ProposalID, *Proposal, *LocalPartySaveData](
		m.ctx,
		ec,
		m.p2p,
		m,
		sessionID,
		signer,
		localSubmitter,
		proposalID,
		msg,
		types.ReShareGroupSessionType,
		newPartners, // todo
	)
	oldInnerSession.partyIdMap = oldPartyIdMap
	oldInnerSession.inToOut = make(chan<- *SessionResult[ProposalID, *LocalPartySaveData], 100) // todo

	localData := m.partyData.GetData(ec)
	oldParty, oldOutEndCh, oldOutCh, oldErrCh := createReShareParam(oldInnerSession.ctx, oldParams, *localData)
	oldInnerSession.party = oldParty
	oldInnerSession.endCh = oldOutEndCh
	oldInnerSession.errCH = oldErrCh
	reShareSession.oldSession = oldInnerSession

	newInnerSession := newSession[ProposalID, *Proposal, *LocalPartySaveData](
		m.ctx,
		ec,
		m.p2p,
		m,
		sessionID,
		signer,
		localSubmitter,
		proposalID,
		msg,
		types.ReShareGroupSessionType,
		newPartners,
	)
	newInnerSession.partyIdMap = newPartyIdMap
	newInnerSession.inToOut = m.senateInToOut
	localData = m.partyData.GenerateNewLocalPartySaveData(ec, newPartners)
	party, outEndCh, outCh, errCh := createReShareParam(newInnerSession.ctx, newParams, *localData)
	newInnerSession.party = party
	newInnerSession.endCh = outEndCh
	newInnerSession.errCH = errCh

	reShareSession.newSession = newInnerSession
	m.AddSession(reShareSession)
	reShareSession.Run()

	runReShareParty(oldInnerSession.ctx, reShareSession, oldParty, oldOutCh, oldErrCh) // run old party
	runReShareParty(newInnerSession.ctx, reShareSession, party, outCh, errCh)          // run new party

	return reShareSession.SessionID()
}

func FromRouteMsg(bytes []byte, routing *tss.MessageRouting) *helper.ReceivedPartyState {
	data := helper.ReceivedPartyState{
		WireBytes:               bytes,
		From:                    routing.From,
		IsBroadcast:             routing.IsBroadcast,
		IsToOldCommittee:        routing.IsToOldCommittee,
		IsToOldAndNewCommittees: routing.IsToOldAndNewCommittees,
	}

	return &data
}

func (r *ReShareGroupSession[T, M, D]) Send(ctx context.Context, bytes []byte, routing *tss.MessageRouting, b bool) error {
	var err error

	dest := lo.Map(routing.To, func(item *tss.PartyID, index int) string { return strings.ToLower(item.Id) })
	if len(dest) == 0 {
		return fmt.Errorf("did not expect a msg to have a nil destination during resharing")
	}

	log.Debugf("ReShareGroupSession Send from:%v to: %v", routing.From, routing.To)

	if r.oldSession.Equal(routing.From.Id) { // from
		log.Debugf("msg from oldSession: %v", routing.From.Id)
		err = r.oldSession.Send(ctx, bytes, routing, b)

		if (!routing.IsToOldCommittee || routing.IsToOldAndNewCommittees) && r.newSession.Included(dest) {
			log.Debugf("msg from oldSession: %v, to newSession: %v", routing.From, r.newSession.party.PartyID())
			r.newSession.Post(FromRouteMsg(bytes, routing))
		}
	}

	if r.newSession.Equal(routing.From.Id) {
		log.Debugf("msg from newSession: %v", routing.From.Id)
		err = r.newSession.Send(ctx, bytes, routing, b)

		if (routing.IsToOldCommittee || routing.IsToOldAndNewCommittees) && r.oldSession.Included(dest) {
			log.Debugf("msg from newSession: %v, to oldSession: %v", routing.From, r.oldSession.party.PartyID())
			r.oldSession.Post(FromRouteMsg(bytes, routing))
		}
	}

	return err
}

func (r *ReShareGroupSession[T, M, D]) GetReceiveChannel(partyID string) chan *helper.ReceivedPartyState {
	if r.oldSession.Equal(partyID) {
		log.Debugf("GetReceiveChanne: from oldSession: %v", partyID)
		return r.oldSession.GetReceiveChannel(partyID)
	}

	log.Debugf("GetReceiveChanne: from newSession: %v", partyID)

	return r.newSession.GetReceiveChannel(partyID)
}

func (r *ReShareGroupSession[T, M, D]) Post(state *helper.ReceivedPartyState) {
	log.Infof("from: %v,dest: %v", state.From, state.ToPartyIds)

	if len(state.ToPartyIds) == 0 {
		log.Error("did not expect a msg to have a nil destination during resharing")
		return
	}

	if (state.IsToOldCommittee || state.IsToOldAndNewCommittees) && r.oldSession.Included(state.ToPartyIds) {
		log.Debugf("oldSession from: %v, to: %v, dest: %v", state.From.Moniker, r.oldSession.party.PartyID().Moniker, state.ToPartyIds)
		r.oldSession.Post(state)
	}

	if (!state.IsToOldCommittee || state.IsToOldAndNewCommittees) && r.newSession.Included(state.ToPartyIds) {
		log.Debugf("newSession from: %v, to: %v, dest: %v", state.From.Moniker, r.newSession.party.PartyID().Moniker, state.ToPartyIds)
		r.newSession.Post(state)
	}
}

func (r *ReShareGroupSession[T, M, D]) Release() {
	r.oldSession.Release()
	r.newSession.Release()
}

func (r *ReShareGroupSession[T, M, D]) Name() string {
	return types.ReShareGroupSessionType
}

func (r *ReShareGroupSession[T, M, D]) Type() string {
	return types.ReShareGroupSessionType
}

func (r *ReShareGroupSession[T, M, D]) SessionID() party.SessionID {
	return r.newSession.SessionID()
}

func (r *ReShareGroupSession[T, M, D]) GroupID() party.GroupID {
	return r.newSession.GroupID()
}

func (r *ReShareGroupSession[T, M, D]) ProposalID() T {
	return r.newSession.ProposalID()
}

func (r *ReShareGroupSession[T, M, D]) Proposer() common.Address {
	return r.newSession.Proposer()
}

func (r *ReShareGroupSession[T, M, D]) Threshold() int {
	return r.newSession.Threshold()
}

func (r *ReShareGroupSession[T, M, D]) PartyID(id string) *tss.PartyID {
	pid := r.newSession.PartyID(id)
	if pid != nil {
		return pid
	}

	return r.oldSession.PartyID(id)
}

func (r *ReShareGroupSession[T, M, D]) Equal(id string) bool {
	return r.newSession.Equal(id) || r.oldSession.Equal(id)
}

func (r *ReShareGroupSession[T, M, D]) Included(ids []string) bool {
	return r.newSession.Included(ids) || r.oldSession.Included(ids)
}

func (r *ReShareGroupSession[T, M, D]) Run() {
	r.oldSession.Run()
	r.newSession.Run()
}

func (r *ReShareGroupSession[T, M, D]) Participants() types.Participants {
	return r.newSession.Participants()
}

func (r *ReShareGroupSession[T, M, D]) Signer() string {
	return r.newSession.Signer()
}

func (r *ReShareGroupSession[T, M, D]) ChainId() uint64 {
	return r.newSession.ChainId()
}
