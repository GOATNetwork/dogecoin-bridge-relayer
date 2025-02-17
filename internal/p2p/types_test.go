package p2p

import (
	"encoding/hex"
	"encoding/json"
	"testing"
	"time"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/assert"
)

func TestMsgDecode(t *testing.T) {
	originMsg := Message[any]{
		MessageType: 10,
		RequestId:   "test",
		DataType:    "test",
		Data: HeartbeatMessage{
			PeerID:    "100",
			Message:   "test",
			Timestamp: time.Now().Unix(),
		},
	}

	data, err := json.Marshal(originMsg)
	assert.Nil(t, err)
	assert.True(t, len(data) > 0)

	rawMsg := Message[json.RawMessage]{}
	err = json.Unmarshal(data, &rawMsg)
	assert.Nil(t, err)

	heartMsg := HeartbeatMessage{}
	err = json.Unmarshal(rawMsg.Data, &heartMsg)
	assert.Nil(t, err)

	msg := Message[any]{
		MessageType: rawMsg.MessageType,
		RequestId:   rawMsg.RequestId,
		DataType:    rawMsg.DataType,
		Data:        heartMsg,
	}
	assert.Equal(t, originMsg, msg)
}

type Account struct {
	PK      string
	PubKey  string
	Address common.Address
}

var accounts = []Account{
	{
		PK:      "76cbb08e5321cec5f584b2b40b4666d9bbbee59eb3022e80d804e8310b17a105",
		PubKey:  "020b537f46c6da81f84824ce1409bab1f9825fb58b57dcafbf4f4b074e90a0c040",
		Address: common.HexToAddress("0x3a818294ca1F3C27d7588b123Ec43F2546fa07f4"),
	},
	{
		PK:      "ffab86884b5f4696c503e8d0cef97f818d122f44017528c24ce3ac580f12b876",
		PubKey:  "02a8fd23c439e9226f422e94911f06788e0019aa1f8efd4f498f75e4f1d5ef7c0a",
		Address: common.HexToAddress("0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"),
	},
	{
		PK:      "5d0ca3f7b4e63f3308a73537001065ee1d6ff3e217115444b148018a1bcbfaf7",
		PubKey:  "02f82403b0337c908478d381f88582e1051c2a9da22a34cd0a1a5b1d10a85b6256",
		Address: common.HexToAddress("0xf6D37CE75dB465DcDb4c7097bEB9c1D46b171037"),
	},
}

func TestGeneratePeerID(t *testing.T) {
	for _, account := range accounts {
		data, err := hex.DecodeString(account.PK)
		assert.Nil(t, err)
		// pk := ethCrypto.ToECDSAUnsafe(data)
		privateKey, err := crypto.UnmarshalSecp256k1PrivateKey(data)
		assert.Nil(t, err)
		id, err := peer.IDFromPrivateKey(privateKey)
		assert.Nil(t, err)
		t.Log("address", account.Address, "id", id.String())

		secp256k1PK := secp256k1.PrivKeyFromBytes(data)
		stdPK := secp256k1PK.ToECDSA()
		address1 := ethCrypto.PubkeyToAddress(stdPK.PublicKey)
		assert.Equal(t, account.Address.String(), address1.String())
	}
}
