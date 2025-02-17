package tss

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	ecdsaResharing "github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	ecdsaSigning "github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	eddsaResharing "github.com/bnb-chain/tss-lib/v2/eddsa/resharing"
	eddsaSigning "github.com/bnb-chain/tss-lib/v2/eddsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

func Partners() types.Participants {
	// todo online contact get address list
	return lo.Map(
		config.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return ethCrypto.PubkeyToAddress(*pubKey) },
	)
}

func PartyKey(ec crypto.CurveType, participants types.Participants, address common.Address) *big.Int {
	key := new(big.Int).Add(address.Big(), big.NewInt(int64(ec)))
	return key.Add(key, participants.GroupID().Big())
}

func createPartyIDsByGroupWithAlias(ec crypto.CurveType, addressList types.Participants, aliasName string) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(addressList))

	groupID := addressList.GroupID()

	for i, address := range addressList {
		key := PartyKey(ec, addressList, address)
		tssAllPartyIDs[i] = tss.NewPartyID(
			strings.ToLower(key.Text(16)),
			fmt.Sprintf("address: %v, groupID:%v, ec: %s, aliasName: %v", address, groupID, ec.CurveName(), aliasName),
			key,
		)
	}

	return tss.SortPartyIDs(tssAllPartyIDs)
}

func createPartyIDsByGroup(ec crypto.CurveType, addressList types.Participants) tss.SortedPartyIDs {
	return createPartyIDsByGroupWithAlias(ec, addressList, "")
}

func createPartyIDsByAddress(addressList types.Participants) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(addressList))

	for i, address := range addressList {
		key := address.Big()
		tssAllPartyIDs[i] = tss.NewPartyID(
			strings.ToLower(address.String()),
			strings.ToLower(address.String()),
			key,
		)
	}

	return tss.SortPartyIDs(tssAllPartyIDs)
}

func createOldPartyIDsByAddress(addressList types.Participants) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(addressList))

	for i, address := range addressList {
		key := address.Big()
		key = new(big.Int).Add(key, big.NewInt(1)) // key + 1
		tssAllPartyIDs[i] = tss.NewPartyID(
			strings.ToLower(key.Text(16)),
			strings.ToLower(key.Text(16)),
			key,
		)
	}

	return tss.SortPartyIDs(tssAllPartyIDs)
}

func CompareStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func AddressIndex(addressList types.Participants, tssAddress common.Address) int {
	for i, address := range addressList {
		if address == tssAddress {
			return i // Return the index if a match is found
		}
	}

	return -1 // Return -1 if not found
}

func serializeMessageToBeSigned(nonce uint64, data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	bufBytes := buf.Bytes()
	length := len(bufBytes)
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))

	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)

	return append(append(nonceBytes, lengthBytes...), data...), nil
}

// RunKeyGen starts the local keygen party and handles incoming and outgoing
// messages to other parties.
func RunKeyGen(
	ctx context.Context,
	isProd bool,
	ty crypto.CurveType,
	localSubmitter common.Address, // current submitter
	allPartners types.Participants,
	transport helper.Transporter,
) (tss.Party, map[string]*tss.PartyID, chan *LocalPartySaveData, chan *tss.Error) {
	// outgoing messages to other peers
	outCh := make(chan tss.Message, 256)
	// error if keygen fails, contains culprits to blame
	errCh := make(chan *tss.Error, 10)

	var party tss.Party

	log.Debug("creating new local party")

	outEndCh := make(chan *LocalPartySaveData)
	// output data when keygen finished
	ecdsaEndCh := make(chan *ecdsaKeygen.LocalPartySaveData)
	eddsaEndCh := make(chan *eddsaKeygen.LocalPartySaveData)

	params, partyIdMap := NewParam(ty, localSubmitter, allPartners)

	switch ty {
	case crypto.ECDSA:
		// prod
		if isProd {
			preParams, err := ecdsaKeygen.GeneratePreParams(2 * time.Minute)
			utils.Assert(err)

			party = ecdsaKeygen.NewLocalParty(params, outCh, ecdsaEndCh, *preParams)
		} else {
			// dev
			party = ecdsaKeygen.NewLocalParty(params, outCh, ecdsaEndCh)
		}

	case crypto.EDDSA:
		party = eddsaKeygen.NewLocalParty(params, outCh, eddsaEndCh)
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

	log.Debug("local party created", "partyID", party.PartyID())

	helper.RunParty(ctx, party, errCh, outCh, transport, false)

	return party, partyIdMap, outEndCh, errCh
}

// RunReshare starts the local reshare party and handles incoming and outgoing
// messages to other parties.
func RunReshare(
	ctx context.Context,
	params *tss.ReSharingParameters,
	key LocalPartySaveData,
	transport helper.Transporter,
) (tss.Party, chan *LocalPartySaveData, chan *tss.Error) {
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

	helper.RunParty(ctx, party, errCh, outCh, transport, true)

	return party, outEndCh, errCh
}

func RunParty(
	ctx context.Context,
	msg *big.Int,
	params *tss.Parameters,
	key LocalPartySaveData,
	transport helper.Transporter,
	keyDerivationDelta *big.Int,
) (tss.Party, chan *tsscommon.SignatureData, chan *tss.Error) {
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

	helper.RunParty(ctx, party, errCh, outCh, transport, false)

	return party, endCh, errCh
}

func secp256k1Signature(data *tsscommon.SignatureData) []byte {
	first := data.SignatureRecovery[0]
	if first < 27 {
		first += 27
	}

	return append(data.Signature, first)
}

// ConvertP2PMsgData converts the message data to the corresponding struct.
func ConvertP2PMsgData(msg p2p.Message[json.RawMessage]) any {
	switch msg.DataType {
	case types.GenKeySessionType, types.ReShareGroupSessionType,
		types.SignTaskSessionType, // only used test
		types.SignOperationSessionType, types.SignTxSessionType:
		return types.UnmarshalJson[SessionMessage[ProposalID, Proposal]](msg.Data)
	}

	panic("error p2p data type")
}
