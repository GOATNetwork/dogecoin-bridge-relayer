package bip44

import (
	"crypto/ecdsa"
	"fmt"
	"strconv"
	"strings"

	tssCrypto "github.com/bnb-chain/tss-lib/v2/crypto"
)

const (
	// DeterministicWalletsPurpose Purpose 44'.
	DeterministicWalletsPurpose = uint32(0x8000002C)

	// CoinTypeBitcoin Coin type 0'.
	CoinTypeBitcoin = uint32(0x80000000)
	// CoinTypeEther Coin type 60'.
	CoinTypeEther = uint32(0x8000003c)
	// CoinTypeSolana Coin type 501'.
	CoinTypeSolana = uint32(0x800001f5)

	// Bitcoin the derive key path of Bitcoin.
	Bitcoin = DerivePath("m/44'/0'/0'/0/0")
	// Ethereum the derive key path of Ethereum.
	Ethereum = DerivePath("m/44'/60'/0'/0/0")

	Solana = DerivePath("m/44'/501'/0'/0'/0")

	// FirstHardenedChild is the index of the firxt "harded" child key as per the
	// bip32 spec.
	FirstHardenedChild = uint32(0x80000000)

	// PublicKeyCompressedLength is the byte count of a compressed public key.
	PublicKeyCompressedLength = 33
)

// DerivePath is the key path of Hierarchical Deterministic Wallets.
type DerivePath string

// DerivePathParams is the BIP44 params of a derived path.
type DerivePathParams struct {
	// Purpose is the purpose field of BIP44. Must be 44'
	Purpose uint32
	// CoinType is the coin type for BIP44.
	// 0' for Bitcoin, 60' for Ethereum, 501' for Solana
	// See https://github.com/satoshilabs/slips/blob/master/slip-0044.md for detail.
	CoinType                      uint32
	Account, Change, AddressIndex uint32
	// Extra for any possible path
	Extra []uint32
	// Depth for the actual depth of the path
	Depth uint
}

func parseIndex(s string) (uint32, error) {
	index := uint32(0)
	if strings.HasSuffix(s, "'") {
		index |= FirstHardenedChild
		s = s[:len(s)-1]
	}

	idx, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return index | uint32(idx), nil
}

// ToParams converts a string derive path to DerivePathParams.
func (p DerivePath) ToParams() (*DerivePathParams, error) {
	components := strings.Split(string(p), "/")
	n := len(components)

	if n < 2 || components[0] != "m" {
		return nil, fmt.Errorf("invalid derive path: %s. must have at least 2 components and starts with 'm/'", p)
	}

	params := DerivePathParams{
		Depth: uint(n - 1),
	}

	var err error

	params.Purpose, err = parseIndex(components[1])
	if err != nil {
		return nil, err
	}

	if n >= 3 {
		params.CoinType, err = parseIndex(components[2])
		if err != nil {
			return nil, err
		}
	}

	if n >= 4 {
		params.Account, err = parseIndex(components[3])
		if err != nil {
			return nil, err
		}
	}

	if n >= 5 {
		params.Change, err = parseIndex(components[4])
		if err != nil {
			return nil, err
		}
	}

	if n >= 6 {
		params.AddressIndex, err = parseIndex(components[5])
		if err != nil {
			return nil, err
		}
	}

	if n > 6 {
		extras := components[6:]
		for _, extra := range extras {
			idx, err := parseIndex(extra)
			if err != nil {
				return nil, err
			}

			params.Extra = append(params.Extra, idx)
		}
	}

	return &params, nil
}

// Indexes returns child indexes of this derived path.
func (p *DerivePathParams) Indexes() []uint32 {
	indexes := make([]uint32, 0)
	indexes = append(indexes, p.Purpose)

	if p.Depth >= 2 {
		indexes = append(indexes, p.CoinType)
	}

	if p.Depth >= 3 {
		indexes = append(indexes, p.Account)
	}

	if p.Depth >= 4 {
		indexes = append(indexes, p.Change)
	}

	if p.Depth >= 5 {
		indexes = append(indexes, p.AddressIndex)
	}

	if p.Depth > 5 {
		indexes = append(indexes, p.Extra...)
	}

	return indexes
}

// DerivingPubKeyFromPath EDDSA hd support:
// https://github.com/bnb-chain/tss-lib/pull/299/files
// https://github.com/bnb-chain/tss-lib/pull/300
//func DerivingPubKeyFromPath(masterPub ecdsa.PublicKey, path []uint32) (*big.Int, *ckd.ExtendedKey, error) {
//	// Generate key and chaincode
//	// h := hmac.New(sha512.New, []byte("Bitcoin seed"))
//	h := hmac.New(sha512.New, append(masterPub.X.Bytes(), masterPub.Y.Bytes()...)) // todo
//	intermediary := h.Sum(nil)
//	// Split it into our key and chain code
//	// keyBytes := intermediary[:32]
//	chainCode := intermediary[32:] // todo
//	net := &chaincfg.MainNetParams // todo
//	extendedParentPk := &ckd.ExtendedKey{
//		PublicKey:  masterPub,
//		Depth:      0,
//		ChildIndex: 0,
//		ChainCode:  chainCode[:],
//		ParentFP:   []byte{0x00, 0x00, 0x00, 0x00},
//		Version:    net.HDPrivateKeyID[:],
//	}
//
//	ec := btcec.S256()
//
//	return ckd.DeriveChildKeyFromHierarchy(path, extendedParentPk, ec.Params().N, ec)
//}

func Bip44DerivationPath(coinType, account, index uint32) DerivePath {
	// https://learnblockchain.cn/2018/09/28/hdwallet
	// m / purpose' / coin' / account' / change / address_index
	// coin list https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	return DerivePath(fmt.Sprintf("m/44/%d/%d/0/%d", coinType, account, index))
}

func ECPoint(key *ecdsa.PublicKey) *tssCrypto.ECPoint {
	return tssCrypto.NewECPointNoCurveCheck(key.Curve, key.X, key.Y)
}
