package tss

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	ecdsaSigning "github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	eddsaSigning "github.com/bnb-chain/tss-lib/v2/eddsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
	log "github.com/sirupsen/logrus"
)

var _ Session[any] = &SignSession[any, any, any]{}

type SignSession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func (s *SignSession[T, M, D]) Post(data *helper.ReceivedPartyState) {
	if data.IsBroadcast || s.Included(data.ToPartyIds) {
		s.sessionTransport.Post(data)
	}
}

func createSignParty(
	msg *big.Int,
	params *tss.Parameters,
	key LocalPartySaveData,
	keyDerivationDelta *big.Int,
) (tss.Party, chan *tsscommon.SignatureData, chan tss.Message, chan *tss.Error) {
	// outgoing messages to other peers - not one to not deadlock when a party
	// round is waiting for outgoing messages channel to clear
	outCh := make(chan tss.Message, 100000)
	// output signature when finished
	endCh := make(chan *tsscommon.SignatureData, 256)
	// error if signing fails, contains culprits to blame
	errCh := make(chan *tss.Error, 256)

	log.Debug("creating new local party")

	var party tss.Party

	switch key.CurveType() {
	case crypto.ECDSA:
		if keyDerivationDelta != nil {
			party = ecdsaSigning.NewLocalPartyWithKDD(msg, params, *key.ECDSAData(), keyDerivationDelta, outCh, endCh)
		} else {
			party = ecdsaSigning.NewLocalParty(msg, params, *key.ECDSAData(), outCh, endCh)
		}
	case crypto.EDDSA:
		if keyDerivationDelta != nil {
			party = eddsaSigning.NewLocalPartyWithKDD(msg, params, *key.EDDSAData(), keyDerivationDelta, outCh, endCh)
		} else {
			party = eddsaSigning.NewLocalParty(msg, params, *key.EDDSAData(), outCh, endCh)
		}

	default:
		panic("implement me")
	}

	log.Debug("local signing party created", "partyID", party.PartyID())

	// types.RunParty(ctx, party, errCh, outCh, transport, false)

	return party, endCh, outCh, errCh
}

func RandSessionID() party.SessionID {
	b := make([]byte, 32)
	_, _ = rand.Read(b)

	return common.BytesToHash(b[:])
}

func (m *Scheduler) JoinSignSession(msg SessionMessage[ProposalID, Proposal]) error {
	err := m.verifySign(&msg)
	if err != nil {
		return err
	}

	signerCtx := m.GetSigner(msg.Signer)
	if signerCtx == nil {
		return fmt.Errorf("signer not found in context")
	}

	log.Warn("sign proposal")
	m.NewSignSessionWitKey(
		msg.ChainId,
		msg.SessionID,
		msg.SeqId,
		msg.ProposalID,
		&msg.Proposal,
		msg.Type,
		msg.Data,
		signerCtx,
	)

	return nil
}

func (m *Scheduler) NewSignOperationSession(
	sessionID party.SessionID,
	seqId uint64,
	proposalID ProposalID,
	msg *Proposal,
	data []byte,
) party.SessionID {
	return m.NewSignSessionWitKey(
		types.ChainIdEthereum,
		sessionID,
		seqId,
		proposalID,
		msg,
		types.SignOperationSessionType,
		data,
		m.tssSigner(),
	)
}

func (m *Scheduler) NewSignSession(
	sessionID party.SessionID,
	seqId uint64,
	proposalID ProposalID,
	msg *Proposal,
	signer *SignerContext,
) party.SessionID {
	return m.NewSignSessionWitKey(types.ChainIdEthereum, sessionID, seqId, proposalID, msg, types.SignTaskSessionType, nil, signer)
}

func (m *Scheduler) NewTxSignSession(
	sessionID party.SessionID,
	seqId uint64,
	proposalID ProposalID,
	msg *Proposal,
	signer *SignerContext,
) party.SessionID {
	return m.NewSignSessionWitKey(types.ChainIdEthereum, sessionID, seqId, proposalID, msg, types.SignTxSessionType, nil, signer)
}

func (m *Scheduler) NewSignSessionWitKey(
	chainId uint64,
	sessionID party.SessionID,
	seqId uint64,
	proposalID ProposalID,
	msg *Proposal,
	sessionType string,
	data []byte,
	signer *SignerContext,
) party.SessionID {
	allPartners := m.Participants()
	params, partyIdMap := NewParam(signer.CurveType(), m.LocalSubmitter(), allPartners)
	ctx, cancel := context.WithTimeout(m.ctx, m.signTimeout)
	_ = cancel
	innerSession := newSession[ProposalID, *Proposal, *tsscommon.SignatureData](
		ctx,
		signer.CurveType(),
		m.p2p,
		m,
		sessionID,
		signer.Address(),
		m.Proposer(),
		proposalID,
		msg,
		sessionType,
		allPartners,
	)
	innerSession.session.Data = data
	innerSession.session.SeqId = seqId
	innerSession.session.ChainId = chainId
	innerSession.inToOut = m.sigInToOut
	innerSession.partyIdMap = partyIdMap
	p, endCh, outCh, errCh := createSignParty(msg, params, signer.LocalData(), signer.KeyDerivationDelta())
	innerSession.party = p
	innerSession.endCh = endCh
	innerSession.errCH = errCh
	session := &SignSession[ProposalID, *Proposal, *tsscommon.SignatureData]{sessionTransport: innerSession}
	session.Run()
	m.AddSession(session)

	helper.RunParty(session.ctx, p, errCh, outCh, session, false)

	return session.SessionID()
}
