package helper

import (
	"context"
	"math/big"

	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	log "github.com/sirupsen/logrus"
)

// RunSign starts the local signing party and handles incoming and outgoing
// messages to other parties.
func RunSign(
	ctx context.Context,
	msg *big.Int,
	params *tss.Parameters,
	key keygen.LocalPartySaveData,
	transport Transporter,
) (tss.Party, chan *common.SignatureData, chan *tss.Error) {
	return Run(ctx, msg, params, key, transport, nil)
}

func Run(
	ctx context.Context,
	msg *big.Int,
	params *tss.Parameters,
	key keygen.LocalPartySaveData,
	transport Transporter,
	keyDerivationDelta *big.Int,
) (tss.Party, chan *common.SignatureData, chan *tss.Error) {
	// outgoing messages to other peers - not one to not deadlock when a party
	// round is waiting for outgoing messages channel to clear
	outCh := make(chan tss.Message, params.PartyCount())
	// output signature when finished
	endCh := make(chan *common.SignatureData, 1)
	// error if signing fails, contains culprits to blame
	errCh := make(chan *tss.Error, 1)

	log.Debug("creating new local party")

	var party tss.Party
	if keyDerivationDelta != nil {
		party = signing.NewLocalPartyWithKDD(msg, params, key, keyDerivationDelta, outCh, endCh)
	} else {
		party = signing.NewLocalParty(msg, params, key, outCh, endCh)
	}

	log.Debug("local signing party created", "partyID", party.PartyID())

	RunParty(ctx, party, errCh, outCh, transport, false)

	return party, endCh, errCh
}

// RunSignWithHD starts the local signing party and handles incoming and outgoing
// messages to other parties.
func RunSignWithHD(
	ctx context.Context,
	msg *big.Int,
	params *tss.Parameters,
	key keygen.LocalPartySaveData,
	transport Transporter,
	keyDerivationDelta *big.Int,
) (tss.Party, chan *common.SignatureData, chan *tss.Error) {
	return Run(ctx, msg, params, key, transport, keyDerivationDelta)
}
