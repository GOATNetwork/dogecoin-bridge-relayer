package utils

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func signHash(data []byte) common.Hash {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)

	return crypto.Keccak256Hash([]byte(msg))
}

func TestSign(t *testing.T) {
	hash := common.HexToHash("0x681b436b243e0e795fb16380f257b2ace153dc8917e29346bb4580e41c601e81")
	hash1 := PersonalMsgHash(hash)
	t.Log(hash1)

	fullMessage := "\x19Ethereum Signed Message:\n32"
	msg := append([]byte(fullMessage), hash.Bytes()...)
	hash2 := crypto.Keccak256Hash(msg)
	t.Log(hash2)

	hash3 := signHash(hash.Bytes())
	t.Log(hash3)

	address := common.HexToAddress("00000000000000000000000004d9389cf937b1e6f2258d842e7237e955d6ab04")
	t.Log(address.Hex())

	address = common.HexToAddress("000000000000000000000000c9a4b85549a239b0259e52baa216b3611c34167c")
	t.Log(address.Hex())

	pk, err := crypto.HexToECDSA("76cbb08e5321cec5f584b2b40b4666d9bbbee59eb3022e80d804e8310b17a105")
	assert.Nil(t, err)

	t.Log(crypto.PubkeyToAddress(pk.PublicKey))

	signature, err := PersonalSign(hash, pk)
	assert.Nil(t, err)
	t.Log(signature)
}
