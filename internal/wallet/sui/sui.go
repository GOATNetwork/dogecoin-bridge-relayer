package sui

import (
	"encoding/hex"
	"fmt"

	"github.com/block-vision/sui-go-sdk/common/keypair"
	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"golang.org/x/crypto/blake2b"
)

func init() {
	address.RegisterAddress(types.CoinTypeSUI, SuiEd25519Address)
}

// https://github.com/MystenLabs/sui/blob/f3b5fdd73da64a0df65fb4323471512b0f57ec4d/sdk/typescript/test/unit/cryptography/ed25519-keypair.test.ts
// https://docs-zh.sui-book.com/concepts/cryptography/transaction-auth/keys-addresses
// m / purpose' / coin_type' / account' / change / address_index
// https://github.com/coming-chat/go-sui-sdk/blob/main/account/account.go
// https://github.com/block-vision/sui-go-sdk
// https://github.com/MystenLabs/sui/blob/main/sdk/typescript/src/keypairs/ed25519/ed25519-hd-key.ts

const SuiChainCodeED25519 = "ed25519 seed"

func SuiEd25519Address(p *crypto.ECPoint) string {
	pubkey := edwards.NewPublicKey(p.X(), p.Y())

	return Ed25519PublicKeyToSuiAddress(pubkey.Serialize())
}

// Ed25519PublicKeyToSuiAddress https://github.com/MystenLabs/sui/blob/main/sdk/typescript/src/cryptography/publickey.ts#L112
func Ed25519PublicKeyToSuiAddress(pubKey []byte) string {
	newPubkey := []byte{byte(keypair.Ed25519Flag)}
	newPubkey = append(newPubkey, pubKey...)

	addrBytes := blake2b.Sum256(newPubkey)

	return fmt.Sprintf("0x%s", hex.EncodeToString(addrBytes[:])[:64])
}
