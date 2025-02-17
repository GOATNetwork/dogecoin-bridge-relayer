package helper

import (
	"context"

	"github.com/bnb-chain/tss-lib/v2/tss"
)

// Transporter is the interface that defines the Send and Receive methods to
// transfer lib-tss messages between parties.
type Transporter interface {
	Send(context.Context, []byte, *tss.MessageRouting, bool) error
	// GetReceiveChannel returns a channel that will be read by the local tss party. This
	// consists of ReceivedPartyState messages received from other parties.
	GetReceiveChannel(partyID string) chan *ReceivedPartyState

	Post(*ReceivedPartyState)

	Release()
}

// ReceivedPartyState is a message received from another party.
type ReceivedPartyState struct {
	WireBytes               []byte
	From                    *tss.PartyID
	ToPartyIds              []string
	IsBroadcast             bool
	IsToOldCommittee        bool `json:"is_to_old_committee"`          // whether the message should be sent to old committee participants rather than the new committee
	IsToOldAndNewCommittees bool `json:"is_to_old_and_new_committees"` // whether the message should be sent to both old and new committee participants
}

// NewReceivedPartyState returns a new ReceivedPartyState.
func NewReceivedPartyState(
	wireBytes []byte,
	from *tss.PartyID,
	isBroadcast bool,
) ReceivedPartyState {
	return ReceivedPartyState{
		WireBytes:   wireBytes,
		From:        from,
		IsBroadcast: isBroadcast,
	}
}
