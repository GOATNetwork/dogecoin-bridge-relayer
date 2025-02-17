package btc

import (
	"fmt"

	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

func init() {
	address.RegisterAddress(types.CoinTypeBTC, GenerateP2WPKHBTCAddress)
}

func GenerateCompressedBTCAddress(p *crypto.ECPoint) (string, error) {
	return btcAddress(NewPublicKeyOfBtc(p).SerializeCompressed())
}

func GenerateUnCompressedBTCAddress(p *crypto.ECPoint) (string, error) {
	return btcAddress(NewPublicKeyOfBtc(p).SerializeUncompressed())
}

func btcAddress(serializedPubKey []byte) (string, error) {
	addr, err := btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %w", err)
	}

	return addr.EncodeAddress(), nil
}

func NewPublicKeyOfBtc(p *crypto.ECPoint) *btcec.PublicKey {
	var (
		x = &btcec.FieldVal{}
		y = &btcec.FieldVal{}
	)

	x.SetByteSlice(p.X().Bytes())
	y.SetByteSlice(p.Y().Bytes())

	return btcec.NewPublicKey(x, y)
}

// GenerateP2WPKHBTCAddress P2WPKH(pay to witness public key hash) address.
func GenerateP2WPKHBTCAddress(p *crypto.ECPoint) string {
	address, err := P2WPKHAddress(NewPublicKeyOfBtc(p).SerializeCompressed())
	utils.Assert(err)

	return address
}

func P2WPKHAddress(serializedPubKey []byte) (string, error) {
	addr, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(serializedPubKey), &chaincfg.MainNetParams)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %w", err)
	}

	return addr.EncodeAddress(), nil
}
