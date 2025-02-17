package tss

import (
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
)

type Session[T any] interface {
	helper.Transporter
	types.ChainID
	Releaser
	Name
	Type
	SessionID
	GroupID
	ProposalIdent[T]
	Proposer
	Threshold
	PartyID
	Equal
	Included
	Run
	Participants
	Signer
}

type Signer interface {
	Signer() string
}

type Run interface {
	Run()
}

type Name interface {
	Name() string
}

type Type interface {
	Type() string
}

type SessionID interface {
	SessionID() party.SessionID
}

type GroupID interface {
	GroupID() party.GroupID
}

type ProposalIdent[T any] interface {
	ProposalID() T
}

type Proposer interface {
	Proposer() common.Address
}

type PartyID interface {
	PartyID(id string) *tss.PartyID
}

type Included interface {
	Included(ids []string) bool
}

type Equal interface {
	Equal(id string) bool
}

type Threshold interface {
	Threshold() int
}

type Participants interface {
	Participants() types.Participants
}
