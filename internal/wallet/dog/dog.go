package dog

import (
	"fmt"

	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet/btc"
)

var MainNetParams = chaincfg.MainNetParams

func init() {
	// P2PKH - D
	// P2SH - 9, A
	// P2PKH (Testnet) - n
	// P2SH (Testnet) - 2
	MainNetParams.PubKeyHashAddrID = 0x1e

	address.RegisterAddress(types.CoinTypeDOG, GenerateP2PKHAddress)
}

func DogAddress(serializedPubKey []byte) (string, error) {
	addr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(serializedPubKey), &MainNetParams)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %w", err)
	}

	return addr.EncodeAddress(), nil
}

func GenerateP2PKHAddress(p *crypto.ECPoint) string {
	address, err := DogAddress(btc.NewPublicKeyOfBtc(p).SerializeCompressed())
	utils.Assert(err)

	return address
}
