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
	"github.com/davecgh/go-spew/spew"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
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
	client := rpc.New(rpc.DevNet_RPC)
	pubkey := solana.MustPublicKeyFromBase58("2cz1TgTjQSdmGSjUiL9Z1QupEAUD3S46AX4KB4Uefr59")
	to := solana.MustPublicKeyFromBase58("xJW6cfu7atWKje4bzcxr83Aqpb8aKt3w271pv3jKTPN")
	t.Log(pubkey.String())
	out, err := client.GetBalance(
		context.Background(),
		pubkey,
		rpc.CommitmentFinalized,
	)
	assert.Nil(t, err)
	spew.Dump(out)
	spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports

	lamportsOnAccount := new(big.Float).SetUint64(out.Value)
	// Convert lamports to sol:
	solBalance := new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))
	// WARNING: this is not a precise conversion.
	t.Log("◎", solBalance.Text('f', 10)) // 500000000

	c := NewDevSolClient()
	unSignTx, err := c.BuildSolTransfer(pubkey, to, 500000)
	assert.Nil(t, err)

	sign, err := unSignTx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		p := pk.SerializeSecret()
		pp := (*solana.PrivateKey)(&p)
		t.Logf("pk: %s, pubkey: %s, key: %s", pp.String(), pp.PublicKey().String(), key.String())
		return pp
	})
	assert.Nil(t, err)
	t.Log(sign)

	signTx := unSignTx.BuildSolTransaction(sign[0])

	err = c.SendTransaction(context.Background(), (*solana.Transaction)(signTx))

	assert.Nil(t, err)
	out, err = client.GetBalance(
		context.Background(),
		to,
		rpc.CommitmentFinalized,
	)
	assert.Nil(t, err)
	spew.Dump(out)
	spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports
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
