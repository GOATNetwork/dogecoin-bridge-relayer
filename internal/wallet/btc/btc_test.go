package btc

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/assert"
)

func TestBTCAddress(t *testing.T) {
	type Args struct {
		PK              string
		isCompress      bool
		expectedAddress string
	}

	args := []Args{
		{
			PK:              "ccea9c5a20e2b78c2e0fbdd8ae2d2b67e6b1894ccb7a55fc1de08bd53994ea64",
			isCompress:      false,
			expectedAddress: "14xfJr1DArtYR156XBs28FoYk6sQqirT2s",
		},
		{
			PK:              "2a70638d5ed9ae0c69a3536d0cc94658c9f9a29cc513c6698f1f045670f49456",
			isCompress:      false,
			expectedAddress: "17RTxvdA3pwt1c3SK8x1aGGFVsQCoLLmCR",
		},
		//{
		//	PK:              "2a70638d5ed9ae0c69a3536d0cc94658c9f9a29cc513c6698f1f045670f49456",
		//	isCompress:      true,
		//	expectedAddress: "1KnpLgu82MRN96SFabufNjFDf5rkXKc9j1",
		//},
		{
			PK:              "7d52d4847d04d6afaae94370f885f9c04e962eeba777b960092fda9dbc363cea",
			isCompress:      false,
			expectedAddress: "16zoEAS4SbWz4NsDnZ76qg4Avuomy2pQbi",
		},
		//{
		//	PK:              "7d52d4847d04d6afaae94370f885f9c04e962eeba777b960092fda9dbc363cea",
		//	isCompress:      true,
		//	expectedAddress: "1H5j9kT2MpMKnJ7qV92FyGTfR9YmusEKPX",
		//},
	}

	for _, arg := range args {
		pkData, err := hex.DecodeString(arg.PK)
		assert.NoError(t, err)

		pk, pubkey := btcec.PrivKeyFromBytes(pkData)
		assert.NotNil(t, pk)
		assert.NotNil(t, pubkey)
		t.Logf("SerializeUncompressed pubkey: %x", pubkey.SerializeUncompressed())
		t.Logf("SerializeCompressed pubkey: %x", pubkey.SerializeCompressed())

		point := crypto.NewECPointNoCurveCheck(tss.S256(), pubkey.X(), pubkey.Y())

		var address string

		if arg.isCompress {
			address, err = GenerateCompressedBTCAddress(point)
		} else {
			address, err = GenerateUnCompressedBTCAddress(point)
		}

		assert.NoError(t, err)
		t.Log("address:", address)
		assert.Equal(t, strings.ToLower(arg.expectedAddress), strings.ToLower(address))

		address = GenerateP2WPKHBTCAddress(point)
		t.Log("GenerateP2WPKHBTCAddress address:", address, "len", len([]byte(address)))
	}
}
