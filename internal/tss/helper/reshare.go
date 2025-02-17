package helper

import (
	"context"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	log "github.com/sirupsen/logrus"
)

// RunReshare starts the local reshare party and handles incoming and outgoing
// messages to other parties.
func RunReshare(
	ctx context.Context,
	params *tss.ReSharingParameters,
	key keygen.LocalPartySaveData,
	transport Transporter,
) (tss.Party, chan *keygen.LocalPartySaveData, chan *tss.Error) {
	// outgoing messages to other peers
	outCh := make(chan tss.Message, 100)
	// output reshared key when finished
	endCh := make(chan *keygen.LocalPartySaveData, 1)
	// error if reshare fails, contains culprits to blame
	errCh := make(chan *tss.Error, 1)

	log.Debug("creating new local party")

	party := resharing.NewLocalParty(params, key, outCh, endCh)
	log.Debug("local resharing party created", "partyID", party.PartyID())

	RunParty(ctx, party, errCh, outCh, transport, true)

	return party, endCh, errCh
}
