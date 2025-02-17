package solana

import (
	"context"
	"math/big"
	"testing"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/system"
	rpco "github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestBigInt(t *testing.T) {
	data := []string{
		"wertyuiofghjklgfhjkulytfyguitfgjhuyfgcjvhklhkgjvjb3465789yughuygujvhbuhygvhbjuhgjvbjkhgjkhgvhghvhuyutffghyugvjhbnnjh",
		`"transaction": {
				"message": {
					"header": {
						"numReadonlySignedAccounts": 0,
						"numReadonlyUnsignedAccounts": 1,
						"numRequiredSignatures": 1
					},
					"accountKeys": [
					"3z9vL1zjN6qyAFHhHQdWYRTFAcy69pJydkZmSFBKHg1R",
					"5snoUseZG8s8CDFHrXY2ZHaCrJYsW457piktDmhyb5Jd",
					"11111111111111111111111111111111"
				],
					"recentBlockhash": "DzfXchZJoLMG3cNftcf2sw7qatkkuwQf4xH15N5wkKAb",
					"instructions": [
				{
					"accounts": [
					0,
					1
				],
					"data": "3Bxs4NN8M2Yn4TLb",
					"programIdIndex": 2,
					"stackHeight": null
				}
				],
					"indexToProgramIds": {}
				},
				"signatures": [
				"5LrcE2f6uvydKRquEJ8xp19heGxSvqsVbcqUeFoiWbXe8JNip7ftPQNTAVPyTK7ijVdpkzmKKaAQR7MWMmujAhXD"
			]
		}`,
		` {
			  "keys": [
				{
				  "pubkey": "3z9vL1zjN6qyAFHhHQdWYRTFAcy69pJydkZmSFBKHg1R",
				  "isSigner": true,
				  "isWritable": true
				},
				{
				  "pubkey": "BpvxsLYKQZTH42jjtWHZpsVSa7s6JVwLKwBptPSHXuZc",
				  "isSigner": false,
				  "isWritable": true
				}
			  ],
			  "programId": "11111111111111111111111111111111",
			  "data": [2,0,0,0,128,150,152,0,0,0,0,0]
			}`,
	}

	for _, d := range data {
		bigInt := new(big.Int)
		bigInt.SetBytes([]byte(d))
		t.Log(bigInt.String())
		t.Logf("%v", bigInt)
		t.Logf("%x", bigInt.Bytes())
	}
}

func TestSolTransfer(t *testing.T) {
	utils.SkipCI(t)

	// https://key.tokenpocket.pro/?locale=zh#/?network=SOL
	// data := base58.Decode("5ZnCSBuoktAiv1titQWUzHd9iqvy9sD8vQNMrHxZMR8KMzjwkM3GQyX7qfoZJ6cYU1HLEX6bT25B2rtRhKiM8MVc")
	data := base58.Decode("3VNyeZkEXWG7ewAiXrnGDNNWsrsUHBzMQmC6Mt7wjfagmEXHtu8Tx4C7pNujguc9yhr2DJbTerN6hdvJRzP7V6Aw")
	pk, pubKey := edwards.PrivKeyFromBytes(data)
	assert.NotNil(t, pk)
	assert.NotNil(t, pubKey)

	c := NewDevSolClient()

	p := common.PublicKeyFromBytes(pubKey.Serialize())
	t.Log("p", p.String())

	pubkey := common.PublicKeyFromString("2cz1TgTjQSdmGSjUiL9Z1QupEAUD3S46AX4KB4Uefr59")
	// pubkey := common.PublicKeyFromString("jxK4DrMrDevCn7UXGhiJPjT36e4XP12cJLFDvP9uvxX")
	to := common.PublicKeyFromString("xJW6cfu7atWKje4bzcxr83Aqpb8aKt3w271pv3jKTPN")

	t.Log(pubkey.String())
	out, err := c.GetBalanceOfSol(context.Background(), pubkey)
	assert.Nil(t, err)
	t.Log("from balance:", out)

	out, err = c.GetBalanceOfSol(context.Background(), to)
	assert.Nil(t, err)
	t.Log("to balance:", out)

	unSignTx, err := c.BuildSolTransfer(pubkey, to, 5000)
	assert.Nil(t, err)

	raw, err := unSignTx.RawData()
	assert.Nil(t, err)

	sin, err := pk.Sign(raw)
	assert.Nil(t, err)

	signedTx := unSignTx.BuildTxWithSignature(sin.Serialize())

	sigData := base58.Encode(sin.Serialize())
	t.Log("sigData", sigData)

	signature, err := c.SyncSendTransaction(context.Background(), (*types.Transaction)(signedTx))
	assert.Nil(t, err)
	t.Log("signature", signature)

	out, err = c.GetBalanceOfSol(context.Background(), to)
	assert.Nil(t, err)
	t.Log("to balance:", out)
}

func TestSolTransfer1(t *testing.T) {
	utils.SkipCI(t)

	data := base58.Decode("3VNyeZkEXWG7ewAiXrnGDNNWsrsUHBzMQmC6Mt7wjfagmEXHtu8Tx4C7pNujguc9yhr2DJbTerN6hdvJRzP7V6Aw")
	// data := base58.Decode("5ZnCSBuoktAiv1titQWUzHd9iqvy9sD8vQNMrHxZMR8KMzjwkM3GQyX7qfoZJ6cYU1HLEX6bT25B2rtRhKiM8MVc")
	pk, pubKey := edwards.PrivKeyFromBytes(data)
	assert.NotNil(t, pk)
	assert.NotNil(t, pubKey)

	pubkey := common.PublicKeyFromString("2cz1TgTjQSdmGSjUiL9Z1QupEAUD3S46AX4KB4Uefr59")
	// pubkey := common.PublicKeyFromString("jxK4DrMrDevCn7UXGhiJPjT36e4XP12cJLFDvP9uvxX")
	to := common.PublicKeyFromString("xJW6cfu7atWKje4bzcxr83Aqpb8aKt3w271pv3jKTPN")

	c := client.NewClient(rpco.DevnetRPCEndpoint)

	// to fetch recent blockhash
	res, err := c.GetLatestBlockhash(context.Background())
	assert.Nil(t, err)

	// create a message
	message := types.NewMessage(types.NewMessageParam{
		FeePayer:        pubkey,
		RecentBlockhash: res.Blockhash, // recent blockhash
		Instructions: []types.Instruction{
			system.Transfer(system.TransferParam{
				From:   pubkey, // from
				To:     to,     // to
				Amount: 1e7,    // 1 SOL
			}),
		},
	})

	// create tx by message + signer
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: message,
		Signers: []types.Account{
			{
				PublicKey:  pubkey,
				PrivateKey: pk.SerializeSecret(),
			},
		},
	})

	assert.Nil(t, err)

	// send tx
	txhash, err := c.SendTransaction(context.Background(), tx)
	assert.Nil(t, err)

	t.Log("txhash:", txhash)
}

func TestBalanceOfSol(t *testing.T) {
	utils.SkipCI(t)

	hotAddress := "ATFdx2yY8uAA345ZPyWYcCcr7Avk6ThUoqTG1jSJDebU"
	c := NewDevSolClient()

	amount, err := c.GetBalanceOfSol(context.Background(), common.PublicKeyFromString(hotAddress))
	assert.Nil(t, err)
	t.Log("amount:", amount)

	kk := common.PublicKeyFromString("ESy7hzp2VFD9ew7KWXUHFewWvy3WwoG7LkUpzv2cQXek")
	t.Logf("kk: %x", kk.Bytes())
}
