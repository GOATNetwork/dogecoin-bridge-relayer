package helper

import (
	"context"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	log "github.com/sirupsen/logrus"
)

// RunKeyGen starts the local keygen party and handles incoming and outgoing
// messages to other parties.
func RunKeyGen(
	ctx context.Context,
	preParams *keygen.LocalPreParams,
	params *tss.Parameters,
	transport Transporter,
) (tss.Party, chan *keygen.LocalPartySaveData, chan *tss.Error) {
	// outgoing messages to other peers
	outCh := make(chan tss.Message, 10)
	// error if keygen fails, contains culprits to blame
	errCh := make(chan *tss.Error, 10)
	// output data when keygen finished
	endCh := make(chan *keygen.LocalPartySaveData, 10)

	log.Debug("creating new local party")

	party := keygen.NewLocalParty(params, outCh, endCh, *preParams)
	log.Debug("local party created", "partyID", party.PartyID())

	RunParty(ctx, party, errCh, outCh, transport, false)

	return party, endCh, errCh
}
