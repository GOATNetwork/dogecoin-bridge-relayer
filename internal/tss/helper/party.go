package helper

import (
	"context"
	"time"

	"github.com/bnb-chain/tss-lib/v2/tss"
	log "github.com/sirupsen/logrus"
)

// RunParty starts the local party in the background and handles incoming and
// outgoing messages. Does **not** block.
func RunParty(
	ctx context.Context,
	party tss.Party,
	errCh chan<- *tss.Error,
	outCh <-chan tss.Message,
	transport Transporter,
	isReSharing bool,
) {
	// Start party in goroutine
	log.Debug("Starting party")

	if err := party.Start(); err != nil {
		errCh <- err
	}

	// Process outgoing and incoming messages
	go func() {
		log.Debugf("party.PartyID():%v", party.PartyID())
		incomingMsgCh := transport.GetReceiveChannel(party.PartyID().Id)

		log.Debug("Starting out/in message loop")

		ticker := time.NewTicker(10 * time.Second)

		for {
			log.Debug("waiting for next message...", "partyID: ", party.PartyID())
			select {
			case <-ctx.Done():
				log.Infof("party done:%v ", party.PartyID())
				return

			case <-ticker.C:
				log.Debugf("party: %v, waiting for: %v", party.PartyID(), party.WaitingFor())
			case outgoingMsg := <-outCh:
				log.Debug("outgoing message", "GetTo(): ", outgoingMsg.GetTo())

				data, routing, err := outgoingMsg.WireBytes()
				log.Debug(
					"party outgoing msg write bytes",
					"partyID: ", party.PartyID(),
					"routing: ", routing,
				)

				if err != nil {
					errCh <- party.WrapError(err)
					return
				}

				// Prevent blocking goroutine to receive messages, may deadlock
				// if receive channels are full if not in goroutine.
				// outCh => Send
				// send to other parties
				if err := transport.Send(ctx, data, routing, isReSharing); err != nil {
					log.Error(
						"failed to send output message: ",
						"from PartyID: ", party.PartyID(),
						"err: ", err,
					)
					errCh <- party.WrapError(err)

					return
				}

				log.Debug("done sending outgoing message", "partyID", party.PartyID())
			case incomingMsg := <-incomingMsgCh:
				if incomingMsg == nil {
					log.Debug("done incoming message", "partyID", party.PartyID())
					return
				}

				for !party.Running() {
					log.Debug("party not running...", "partyID: ", party.PartyID())
					time.Sleep(100 * time.Millisecond)
				}

				// Receive => party
				// Running in goroutine prevents blocking when channels get
				// filled up. This may deadlock if not run in a goroutine and
				// blocks receiving messages.
				// go func() {
				log.Debug(
					"received message: ",
					"partyID: ", party.PartyID(),
					"from partyID: ", incomingMsg.From,
					"isBroadcast: ", incomingMsg.IsBroadcast,
					"len(bytes): ", len(incomingMsg.WireBytes),
				)

				// The first return value `ok` is false only when there is
				// an error. This should be fine to ignore as we handle err
				// instead.
				ok, err := party.UpdateFromBytes(
					incomingMsg.WireBytes,
					incomingMsg.From,
					incomingMsg.IsBroadcast,
				)
				if err != nil {
					log.Error("failed to update from bytes", "err", err)
					errCh <- party.WrapError(err)

					return
				}

				if !ok {
					log.Error("UpdateFromBytes: fail")
				}

				log.Debugf(
					"updated party %v from bytes from %v",
					party.PartyID(),
					incomingMsg.From,
				)
			}
		}
	}()
}
