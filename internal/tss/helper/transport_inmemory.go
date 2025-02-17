package helper

import (
	"context"
	"fmt"
	"testing"

	"github.com/bnb-chain/tss-lib/v2/tss"
	log "github.com/sirupsen/logrus"
)

type MemoryTransporter struct {
	PartyID *tss.PartyID
	// incoming messages from other parties.
	recvChan chan *ReceivedPartyState
	// outgoing messages to other parties
	sendChan map[*tss.PartyID]chan *ReceivedPartyState
	// old/new committee only for resigning
	oldCommittee map[*tss.PartyID]chan *ReceivedPartyState
	newCommittee map[*tss.PartyID]chan *ReceivedPartyState
}

func (mt *MemoryTransporter) Post(state *ReceivedPartyState) {
	mt.recvChan <- state
}

var _ Transporter = (*MemoryTransporter)(nil)

func NewMemoryTransporter(partyID *tss.PartyID) *MemoryTransporter {
	ts := &MemoryTransporter{
		PartyID:      partyID,
		recvChan:     make(chan *ReceivedPartyState, 1),
		sendChan:     make(map[*tss.PartyID]chan *ReceivedPartyState),
		oldCommittee: make(map[*tss.PartyID]chan *ReceivedPartyState),
		newCommittee: make(map[*tss.PartyID]chan *ReceivedPartyState),
	}

	return ts
}

func (mt *MemoryTransporter) Release() {
	close(mt.recvChan)
}

func (mt *MemoryTransporter) Send(_ context.Context, data []byte, routing *tss.MessageRouting, isResharing bool) error {
	if isResharing {
		return mt.sendReSharing(data, routing)
	}

	return mt.sendKeygenOrSigning(data, routing)
}

func (mt *MemoryTransporter) sendReSharing(data []byte, routing *tss.MessageRouting) error {
	log.Debug(
		"sending resharing message",
		"to", routing.To,
		"isBroadcast", routing.IsBroadcast,
		"IsToOldCommittee", routing.IsToOldCommittee,
		"IsToOldAndNewCommittees", routing.IsToOldAndNewCommittees,
	)

	dest := routing.To
	if dest == nil {
		return fmt.Errorf("resharing should not have a msg with nil destination")
	}

	if routing.IsToOldCommittee || routing.IsToOldAndNewCommittees {
		log.Debug("sending message to old committee")

		for _, partyID := range dest[:len(mt.oldCommittee)] {
			// Skip sending back to sender
			if partyID == routing.From {
				continue
			}

			ch := mt.oldCommittee[partyID]

			go func(partyID *tss.PartyID, ch chan *ReceivedPartyState) {
				log.Debug("sending message to party", "partyID", partyID, "len(ch)", len(ch))
				ch <- DataRoutingToMessage(data, routing)
				log.Debug("sent message to party", "partyID", partyID, "len(ch)", len(ch))
			}(partyID, ch)
		}
	}

	if !routing.IsToOldCommittee || routing.IsToOldAndNewCommittees {
		log.Debug("sending message to new committee")

		for _, partyID := range dest {
			// Skip sending back to sender
			if partyID == routing.From {
				continue
			}

			ch := mt.newCommittee[partyID]

			go func(partyID *tss.PartyID, ch chan *ReceivedPartyState) {
				log.Debug("sending message to party", "partyID", partyID, "len(ch)", len(ch))
				ch <- DataRoutingToMessage(data, routing)
				log.Debug("sent message to party", "partyID", partyID, "len(ch)", len(ch))
			}(partyID, ch)
		}
	}

	return nil
}

func (mt *MemoryTransporter) sendKeygenOrSigning(data []byte, routing *tss.MessageRouting) error {
	log.Debug(
		"sending message",
		"to", routing.To,
		"isBroadcast", routing.IsBroadcast,
	)

	if routing.IsBroadcast && len(routing.To) == 0 {
		log.Debug("broadcast message to all peers")

		for partyID, ch := range mt.sendChan {
			// Skip send back to sender
			if partyID == routing.From {
				continue
			}

			go func(partyID *tss.PartyID, ch chan *ReceivedPartyState) {
				log.Debug("sending message to party", "partyID", partyID, "len(ch)", len(ch))
				ch <- DataRoutingToMessage(data, routing)
				log.Debug("sent message to party", "partyID", partyID, "len(ch)", len(ch))
			}(partyID, ch)
		}

		log.Debug("done broadcast")

		return nil
	}

	for _, partyID := range routing.To {
		if partyID == routing.From {
			continue
		}

		ch, ok := mt.sendChan[partyID]
		if !ok {
			return fmt.Errorf("party %s not found", partyID)
		}

		go func(partyID *tss.PartyID, ch chan *ReceivedPartyState) {
			log.Debug("sending message to party", "partyID", partyID, "len(ch)", len(ch))
			ch <- DataRoutingToMessage(data, routing)
			log.Debug("sent message to party", "partyID", partyID, "len(ch)", len(ch))
		}(partyID, ch)
	}

	return nil
}

func (mt *MemoryTransporter) AddTarget(partyID *tss.PartyID, ch chan *ReceivedPartyState) {
	mt.sendChan[partyID] = ch
}

// GetReceiver returns a channel for other peer to send messages to.
func (mt *MemoryTransporter) GetReceiver() chan *ReceivedPartyState {
	return mt.recvChan
}

// GetReceiveChannel returns a channel for the current peer to receive messages from
// other peers.
func (mt *MemoryTransporter) GetReceiveChannel(_ string) chan *ReceivedPartyState {
	return mt.recvChan
}

func (mt *MemoryTransporter) AddOldCommitteeTarget(partyID *tss.PartyID, ch chan *ReceivedPartyState) {
	mt.oldCommittee[partyID] = ch
}

func (mt *MemoryTransporter) AddNewCommitteeTarget(partyID *tss.PartyID, ch chan *ReceivedPartyState) {
	mt.newCommittee[partyID] = ch
}

func DataRoutingToMessage(data []byte, routing *tss.MessageRouting) *ReceivedPartyState {
	return &ReceivedPartyState{
		WireBytes:   data,
		From:        routing.From,
		IsBroadcast: routing.IsBroadcast,
	}
}

func CreateAndConnectTransports(
	t *testing.T,
	partyIDs []*tss.PartyID,
) []*MemoryTransporter {
	// Create transport between peers
	transports := make([]*MemoryTransporter, 0)
	for _, partyID := range partyIDs {
		transports = append(transports, NewMemoryTransporter(partyID))
	}

	t.Logf("transports: %+v", transports)

	// Add transport receivers to each other
	for _, transport := range transports {
		for _, otherTransport := range transports {
			transport.AddTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}
	}

	t.Logf("transports connected: %+v", transports)

	return transports
}

func CreateAndConnectReSharingTransports(
	t *testing.T,
	oldCommittee []*tss.PartyID,
	newCommittee []*tss.PartyID,
) ([]*MemoryTransporter, []*MemoryTransporter) {
	// Create transport between peers
	oldTransports := make([]*MemoryTransporter, 0, len(oldCommittee))

	for _, partyID := range oldCommittee {
		mt := NewMemoryTransporter(partyID)
		oldTransports = append(oldTransports, mt)
	}

	newTransports := make([]*MemoryTransporter, 0, len(newCommittee))

	for _, partyID := range newCommittee {
		mt := NewMemoryTransporter(partyID)
		newTransports = append(newTransports, mt)
	}

	t.Logf("old transports: %+v", oldTransports)
	t.Logf("new transports: %+v", newTransports)

	// Add old transport receivers to each other
	for _, transport := range oldTransports {
		for _, otherTransport := range oldTransports {
			transport.AddOldCommitteeTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}

		for _, otherTransport := range newTransports {
			transport.AddNewCommitteeTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}
	}

	t.Logf("old transports connected")

	// Add new transport receivers to each other
	for _, transport := range newTransports {
		for _, otherTransport := range oldTransports {
			transport.AddOldCommitteeTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}

		for _, otherTransport := range newTransports {
			transport.AddNewCommitteeTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}
	}

	t.Logf("new transports connected")

	return oldTransports, newTransports
}
