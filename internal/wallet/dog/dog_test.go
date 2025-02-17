package dog

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/assert"
)

func TestDogAddress(t *testing.T) {
	type Args struct {
		PK              string
		isCompress      bool
		expectedAddress string
	}

	args := []Args{
		{
			PK:              "ccea9c5a20e2b78c2e0fbdd8ae2d2b67e6b1894ccb7a55fc1de08bd53994ea64",
			isCompress:      true,
			expectedAddress: "deamnevpm5hiiakpfdh768ufiuogdcyiat",
		},
		{
			PK:              "2a70638d5ed9ae0c69a3536d0cc94658c9f9a29cc513c6698f1f045670f49456",
			isCompress:      true,
			expectedAddress: "djb7z3mhg9gtgyatpnwttnys5idbxadkdy",
		},
	}

	for _, arg := range args {
		pkData, err := hex.DecodeString(arg.PK)
		assert.NoError(t, err)

		pk, pubkey := btcec.PrivKeyFromBytes(pkData)
		assert.NotNil(t, pk)
		assert.NotNil(t, pubkey)
		t.Logf("SerializeUncompressed pubkey: %x", pubkey.SerializeUncompressed())
		t.Logf("SerializeCompressed pubkey: %x", pubkey.SerializeCompressed())

		// point := crypto.NewECPointNoCurveCheck(tss.S256(), pubkey.X(), pubkey.Y())
		address, err := DogAddress(pubkey.SerializeCompressed())
		assert.NoError(t, err)
		t.Log("address:", address)
		// 1asfqpzbtfpsba9dwdhyynk4qm5nogntzl
		// dtkijpt6pneorlthj5muec8ojvbxcsmxfo
		assert.Equal(t, strings.ToLower(arg.expectedAddress), strings.ToLower(address))
	}
}
