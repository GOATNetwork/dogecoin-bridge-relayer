package solana

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/davecgh/go-spew/spew"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolanaAddress(t *testing.T) {
	type Args struct {
		PK              string
		ExpectedAddress string
	}

	args := []Args{
		{
			PK:              "31ry6kknaZvLyi4yf1ZCVvmoMXtvg8FkpZFtKqJtV6xmmkgii3cKuuBXS8FWgYxW6R5eUDgeDaZuxo2mEVaEriwr",
			ExpectedAddress: "12bLXhcf6gzHHp6pzWyGHUE2NUpXVpbLva8ikTCd6AtG",
		},
		{
			PK:              "47pXBoR4PA6K6VgQY8VCfBbtukffQhfedXYdqZ27TPB6MQdEhF74znuRrvbe3VjP2pFqByhgxZNbeGZ3v7H1emNe",
			ExpectedAddress: "1LPWKkn3KYD2kFp3dZquTBb7hFabKM9KEFAQ2wi76mt",
		},
		{
			PK:              "5ZnCSBuoktAiv1titQWUzHd9iqvy9sD8vQNMrHxZMR8KMzjwkM3GQyX7qfoZJ6cYU1HLEX6bT25B2rtRhKiM8MVc",
			ExpectedAddress: "12nV3gkyFCAYmScaAsBgU1LyyEFQer3f9x2FRvirDSJ",
		},
	}

	for _, arg := range args {
		data := base58.Decode(arg.PK)
		pk, pubKey := edwards.PrivKeyFromBytes(data)
		assert.NotNil(t, pk)
		assert.NotNil(t, pubKey)

		assert.Equal(t, strings.ToLower(arg.ExpectedAddress), strings.ToLower(base58.Encode(pubKey.Serialize())))

		point := crypto.NewECPointNoCurveCheck(tss.Edwards(), pubKey.X, pubKey.Y)
		solanaAddress := SolanaAddress(point)
		assert.Equal(t, strings.ToLower(arg.ExpectedAddress), strings.ToLower(solanaAddress))
	}
}

func TestGenerateSolAddress(t *testing.T) {
	utils.SkipCI(t)

	data := base58.Decode("5ZnCSBuoktAiv1titQWUzHd9iqvy9sD8vQNMrHxZMR8KMzjwkM3GQyX7qfoZJ6cYU1HLEX6bT25B2rtRhKiM8MVc")
	pk, pubKey := edwards.PrivKeyFromBytes(data)
	assert.NotNil(t, pk)
	assert.NotNil(t, pubKey)
	point := crypto.NewECPointNoCurveCheck(tss.Edwards(), pubKey.X, pubKey.Y)
	address := address.GenerateAddressByPath(point, types.CoinTypeSOL, 1, 1)
	t.Log("address", address)
	assert.Equal(t, strings.ToLower("jxK4DrMrDevCn7UXGhiJPjT36e4XP12cJLFDvP9uvxX"), strings.ToLower(address))

	client := rpc.New(rpc.DevNet_RPC)
	pubkey := solana.MustPublicKeyFromBase58(address)

	out, err := client.GetBalance(
		context.Background(),
		pubkey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}

	spew.Dump(out)
	spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports

	lamportsOnAccount := new(big.Float).SetUint64(out.Value)
	// Convert lamports to sol:
	solBalance := new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))

	// WARNING: this is not a precise conversion.
	fmt.Println("◎", solBalance.Text('f', 10))
}

func TestSoMasterAddress(t *testing.T) {
	const (
		ecdsaPublicKey = "0x026ae06bb6b7a4779ef7d2fbcb5da36bec729c54e8b9c235aa75b09a5e22dd427b"
		eddsaPublicKey = "44a3e1108c206006fbcc5d3a5e33dfba38b0f3bca00fe0ccdfc2267e712271a1"
	)

	data, err := hex.DecodeString(eddsaPublicKey)
	assert.Nil(t, err)
	// pubkey := common.PublicKeyFromBytes(data)
	pubKey, err := edwards.ParsePubKey(data)
	assert.Nil(t, err)
	assert.NotNil(t, pubKey)
	t.Logf("master pubKey: %x", pubKey.SerializeCompressed()) // 44a3e1108c206006fbcc5d3a5e33dfba38b0f3bca00fe0ccdfc2267e712271a1
	assert.Equal(t, data, pubKey.SerializeCompressed())

	x, is := new(big.Int).SetString("24723480748839663281939995149474662025134806861408495687251712051005140403503", 10)
	assert.True(t, is)
	y, is := new(big.Int).SetString("15126215440702806707084235929657214535861454973064246117622798020766234420036", 10)
	assert.True(t, is)

	// point := crypto.NewECPointNoCurveCheck(tss.Edwards(), pubKey.X, pubKey.Y)
	point, err := crypto.NewECPoint(tss.Edwards(), x, y)
	assert.Nil(t, err)

	hotAddress := address.HotAddressOfSolana(point)
	t.Log("hotAddress", hotAddress)
	t.Logf("hotAddress pubkey: %x", base58.Decode(hotAddress))

	t.Logf("master pubkey: %x", base58.Decode("NZpbDcRYh6r3derZ6Hrs83rrXY1CxaBqi8x1tK3XSFgpYix4p9HuG5YxSv6Er6WjaQ1xjKGkBbPpguSUdtPnGHs9"))

	pp, err := ethCrypto.UnmarshalPubkey(base58.Decode("NZpbDcRYh6r3derZ6Hrs83rrXY1CxaBqi8x1tK3XSFgpYix4p9HuG5YxSv6Er6WjaQ1xjKGkBbPpguSUdtPnGHs9"))
	assert.Nil(t, err)
	point, err = crypto.NewECPoint(tss.Edwards(), pp.X, pp.Y)
	assert.Nil(t, err)
	t.Logf("x: %v, y:%v", pp.X.String(), pp.Y.String())

	hotAddress = address.HotAddressOfSolana(point)
	t.Log("hotAddress", hotAddress)
	t.Logf("hotAddress pubkey: %x", base58.Decode(hotAddress))

	assert.Equal(t, strings.ToLower("ATFdx2yY8uAA345ZPyWYcCcr7Avk6ThUoqTG1jSJDebU"), strings.ToLower(hotAddress))
	t.Logf("hotAddress pubkey: %x", base58.Decode("ESy7hzp2VFD9ew7KWXUHFewWvy3WwoG7LkUpzv2cQXek"))

	t.Logf("master pubkey: %x", base58.Decode("NptdNDejbuc7F6uHyvMGuptZNY16afpP1CoeQUi5gW1gy8xKGu75GmQBf2u3cjUNEs2VcWyJktrVPhuRn6bxZjhc"))
	pp, err = ethCrypto.UnmarshalPubkey(base58.Decode("NptdNDejbuc7F6uHyvMGuptZNY16afpP1CoeQUi5gW1gy8xKGu75GmQBf2u3cjUNEs2VcWyJktrVPhuRn6bxZjhc"))
	assert.Nil(t, err)
	point, err = crypto.NewECPoint(tss.Edwards(), pp.X, pp.Y)
	assert.Nil(t, err)
	t.Logf("x: %v, y:%v", pp.X.String(), pp.Y.String())

	hotAddress = address.HotAddressOfSolana(point)
	t.Log("hotAddress", hotAddress)
	t.Logf("hotAddress pubkey: %x", base58.Decode(hotAddress))
}

func TestSolHotAddress(t *testing.T) {
	testCase := []struct {
		masterKey    string
		childAddress string
	}{
		{
			"44a3e1108c206006fbcc5d3a5e33dfba38b0f3bca00fe0ccdfc2267e712271a1",
			"ATFdx2yY8uAA345ZPyWYcCcr7Avk6ThUoqTG1jSJDebU",
		},
		{
			"8b0bb8bda779a4fa7886c48a72a9b88b4ed346cda8c3ca53e735b6483acdba79",
			"ESy7hzp2VFD9ew7KWXUHFewWvy3WwoG7LkUpzv2cQXek",
		},
	}

	for _, s := range testCase {
		data, err := hex.DecodeString(s.masterKey)
		assert.Nil(t, err)
		pubKey, err := edwards.ParsePubKey(data)
		assert.Nil(t, err)
		point, err := crypto.NewECPoint(tss.Edwards(), pubKey.X, pubKey.Y)
		assert.Nil(t, err)

		hotAddress := address.HotAddressOfSolana(point)
		t.Log("hotAddress", hotAddress)
		t.Logf("hotAddress pubkey: %x", base58.Decode(hotAddress))
		assert.Equal(t, strings.ToLower(s.childAddress), strings.ToLower(hotAddress))
	}
}
