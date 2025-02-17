package sui

import (
	"context"
	"encoding/base64"
	"strings"
	"testing"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestSuiAddress(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString("ALoFvhYh7S9eDM+hxB9cx6O1UHkRSfEdOAL9geHj0DME")
	t.Logf("%x", data)
	assert.Nil(t, err)
	pubkey, err := edwards.ParsePubKey(data[1:])
	assert.Nil(t, err)

	address := Ed25519PublicKeyToSuiAddress(pubkey.Serialize())
	t.Log(address)
	assert.Equal(t, "0x9cbec822bd17762d757bf741ef9dcea763cc414d970baffb52c191978fabe266", address)
}

func TestGenerateSuiAddress(t *testing.T) {
	utils.SkipCI(t)

	data := base58.Decode("5ZnCSBuoktAiv1titQWUzHd9iqvy9sD8vQNMrHxZMR8KMzjwkM3GQyX7qfoZJ6cYU1HLEX6bT25B2rtRhKiM8MVc")
	pk, pubKey := edwards.PrivKeyFromBytes(data)
	assert.NotNil(t, pk)
	assert.NotNil(t, pubKey)
	point := crypto.NewECPointNoCurveCheck(tss.Edwards(), pubKey.X, pubKey.Y)
	address := address.GenerateAddressByPath(point, types.CoinTypeSUI, 1, 1)
	t.Log("address", address)
	assert.Equal(t, strings.ToLower("0x5283816ef0fe030955141418c61ac7e362101eb251ca6e9e9d812ca2e803320c"), strings.ToLower(address))

	// faucet:
	// curl --location --request POST 'https://faucet.devnet.sui.io/v1/gas' --header 'Content-Type: application/json' --data-raw '{"FixedAmountRequest": {"recipient": "0x5283816ef0fe030955141418c61ac7e362101eb251ca6e9e9d812ca2e803320c"}}'
	// devnet rpc: https://sui-devnet-endpoint.blockvision.org
	// tx info: https://devnet.suivision.xyz/account/0x5283816ef0fe030955141418c61ac7e362101eb251ca6e9e9d812ca2e803320c
	client := sui.NewSuiClient("https://sui-devnet-endpoint.blockvision.org")
	res, err := client.SuiXGetBalance(context.Background(), models.SuiXGetBalanceRequest{
		Owner:    address,
		CoinType: "0x2::sui::SUI",
	})
	assert.Nil(t, err)
	t.Logf("res: %v", res)

	rsp, err := client.SuiXGetAllBalance(context.Background(), models.SuiXGetAllBalanceRequest{
		Owner: address,
	})

	assert.Nil(t, err)
	t.Logf("res: %v", rsp)

	rsp1, err := client.SuiXGetCoins(context.Background(), models.SuiXGetCoinsRequest{
		Owner:    "0x9099b85cce1e63a584f981390bf3457611df3f7778c0d77de3f16cb57951bcf9",
		CoinType: "0x2::sui::SUI",
	})

	assert.Nil(t, err)
	t.Log(utils.FormatJSON(rsp1))
}
