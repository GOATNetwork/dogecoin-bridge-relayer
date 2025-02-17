package tss

import (
	"encoding/json"
	"math/big"
	"reflect"
	"testing"

	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestPartyID(t *testing.T) {
	addressList := types.Participants{
		common.HexToAddress("0x5e7BB104d84c7CB9B682AaC2F3d509f5F406809A"),
		common.HexToAddress("0x33128a8fC17869897dcE68Ed026d694621f6FDfD"),
		common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"),
		common.HexToAddress("0x38015D05f4fEC8AFe15D7cc0386a126574e8077B"),
		common.HexToAddress("0x41ff9AA7e16B8B1a8a8dc4f0eFacd93D02d071c9"),
	}

	partKey := PartyKey(crypto.ECDSA, addressList, common.HexToAddress("0x41ff9AA7e16B8B1a8a8dc4f0eFacd93D02d071c9"))
	t.Log(partKey.Text(16))

	ll := createPartyIDsByGroup(crypto.ECDSA, addressList)
	lo.ForEach(ll, func(item *tss.PartyID, index int) {
		t.Log(new(big.Int).SetBytes(item.Key).Text(16))
	})
}

func TestOldPartyID(t *testing.T) {
	addressList := types.Participants{
		common.HexToAddress("0x5e7BB104d84c7CB9B682AaC2F3d509f5F406809A"),
		common.HexToAddress("0x33128a8fC17869897dcE68Ed026d694621f6FDfD"),
		common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"),
		common.HexToAddress("0x38015D05f4fEC8AFe15D7cc0386a126574e8077B"),
		common.HexToAddress("0x41ff9AA7e16B8B1a8a8dc4f0eFacd93D02d071c9"),
	}

	oldPartyIDs := createOldPartyIDsByAddress(addressList)
	t.Log(oldPartyIDs[0].Id)
	t.Log(oldPartyIDs[0].Moniker)
	t.Logf("%x", oldPartyIDs[0].Key)

	assert.Equal(t, oldPartyIDs[0].Id, oldPartyIDs[0].Moniker)
	assert.Equal(t, oldPartyIDs[0].Id, "33128a8fc17869897dce68ed026d694621f6fdfe")
}

func TestJsonParse(t *testing.T) {
	msg := p2p.Message[types.TssMessage]{
		MessageType: 12,
		RequestId:   "test",
		DataType:    "test",
		Data: types.TssMessage{
			FromPartyId:  "rewr",
			ToPartyIds:   []string{"1", "2"},
			IsBroadcast:  false,
			MsgWireBytes: []byte("1234567890"),
		},
	}

	data, err := json.Marshal(msg)
	assert.Nil(t, err)
	t.Log(string(data))

	msgByte := p2p.Message[json.RawMessage]{}
	err = json.Unmarshal(data, &msgByte)
	assert.Nil(t, err)
	t.Log(msg)
	t.Log(msg.Data)
	t.Log(reflect.ValueOf(msg.Data).String())

	dataMsg := types.TssMessage{}
	err = json.Unmarshal(msgByte.Data, &dataMsg)
	assert.Nil(t, err)
	t.Log(utils.FormatJSON(dataMsg))
}
