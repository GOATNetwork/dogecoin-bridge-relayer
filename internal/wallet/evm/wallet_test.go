package evm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

const (
	ecdsaPublicKey = "0x026ae06bb6b7a4779ef7d2fbcb5da36bec729c54e8b9c235aa75b09a5e22dd427b"
	eddsaPublicKey = "44a3e1108c206006fbcc5d3a5e33dfba38b0f3bca00fe0ccdfc2267e712271a1"
)

func TestMasterPublicKey(t *testing.T) {
	pubKey, err := ethCrypto.DecompressPubkey(hexutil.MustDecode(ecdsaPublicKey))
	assert.Nil(t, err)

	tssAddress := ethCrypto.PubkeyToAddress(*pubKey)
	t.Log(tssAddress.String())
	assert.Equal(t, tssAddress, common.HexToAddress("0xB43EB0e9Ec8040737FFcc144073C72Cf68bC4bab")) // 🙆
}

func TestEcrecover(t *testing.T) {
	data, err := hexutil.Decode("0x61093330325fd881b8cf5a8825dfc477f4707dd8e11b66fdbf665925020ab9df42249c22e83278c4ae65a5eb21bcee01b0f538d760888d771290f1680225f0e001")
	assert.Nil(t, err)
	pubkey, err := ethCrypto.SigToPub(common.HexToHash("0x594dddaa6bb9b5c41b3c4af2551d59785ee3d64d3fd7b66b1d24b64793201d14").Bytes(), data)
	assert.Nil(t, err)

	address := ethCrypto.PubkeyToAddress(*pubkey)

	t.Log(address)
}

func TestTxHash(t *testing.T) {
	data, err := hexutil.Decode("0x61093330325fd881b8cf5a8825dfc477f4707dd8e11b66fdbf665925020ab9df42249c22e83278c4ae65a5eb21bcee01b0f538d760888d771290f1680225f0e001")
	assert.Nil(t, err)

	to := common.HexToAddress("0xB43EB0e9Ec8040737FFcc144073C72Cf68bC4bab")
	baseTx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     1,
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(100),
		Gas:       100,
		To:        &to,
		Value:     big.NewInt(1),
		Data:      data,
	}

	tx := types.NewTx(baseTx)
	hash := tx.Hash()
	t.Log(hash.String())

	signer := types.LatestSignerForChainID(big.NewInt(1))
	signatureHash := signer.Hash(tx)
	t.Log(signatureHash)
}
