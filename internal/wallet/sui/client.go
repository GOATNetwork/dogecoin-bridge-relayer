package sui

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	sutils "github.com/block-vision/sui-go-sdk/utils"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/blake2b"
)

type RawSignTx struct{}

type UnSignTx models.TxnMetaData

type SignedTx models.SignedTransactionSerializedSig

func (t *UnSignTx) Blake2bHash() []byte {
	return Blake2bHash(t.TxBytes)
}

func (t *UnSignTx) TxDigest() string {
	digest, err := sutils.GetTxDigest(t.TxBytes)
	utils.Assert(err)

	return digest
}

func (t *UnSignTx) SerializedSigWith(signature, publicKey []byte) *models.SignedTransactionSerializedSig {
	return &models.SignedTransactionSerializedSig{
		TxBytes:   t.TxBytes,
		Signature: models.ToSerializedSignature(signature, publicKey),
	}
}

// SuiClient https://docs.sui.io/sui-api-ref#unsafe_paysui
type SuiClient struct {
	client sui.ISuiAPI
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *SuiClient) SetUrl(url string) *SuiClient {
	c.client = sui.NewSuiClient(url)
	return c
}

func NewSuiClient(ctx context.Context) *SuiClient {
	ctx, cancel := context.WithCancel(ctx)

	return &SuiClient{
		client: sui.NewSuiClient(constant.SuiMainnetEndpoint),
		ctx:    ctx,
		cancel: cancel,
	}
}

func NewDevClient() *SuiClient {
	c := NewSuiClient(context.Background())
	c.SetUrl("https://sui-devnet-endpoint.blockvision.org")

	return c
}

func (c *SuiClient) GetBalance(address string, coinAddress string, coinName string) (*models.CoinBalanceResponse, error) {
	res, err := c.client.SuiXGetBalance(c.ctx, models.SuiXGetBalanceRequest{
		Owner:    address,
		CoinType: CoinType(coinAddress, coinName), //"0x2::sui::SUI",
	})
	if err != nil {
		return nil, fmt.Errorf("sui get balance: %w", err)
	}

	return &res, nil
}

type Recipient struct {
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

func (c *SuiClient) BuildPaySuiTx(coinType string, from string, recipients []Recipient) (*UnSignTx, error) {
	balanceTotal := decimal.Zero
	amountTotal := decimal.Zero
	amountList := make([]string, 0, len(recipients))
	recipientList := make([]string, 0, len(recipients))
	suiObjectIdList := make([]string, 0, len(recipients))

	for _, v := range recipients {
		recipientList = append(recipientList, v.Recipient)
		amountList = append(amountList, v.Amount)
		amountTotal = amountTotal.Add(decimal.RequireFromString(v.Amount))
	}

	res, err := c.client.SuiXGetCoins(c.ctx, models.SuiXGetCoinsRequest{
		Owner:    from,
		CoinType: coinType,
		Limit:    50,
	})
	if err != nil {
		return nil, fmt.Errorf("call SuiXGetCoins error: %w", err)
	}

	is := false

	for _, coins := range res.Data {
		balance := decimal.RequireFromString(coins.Balance)
		balanceTotal = balanceTotal.Add(balance)

		suiObjectIdList = append(suiObjectIdList, coins.CoinObjectId)

		if balanceTotal.Cmp(amountTotal) >= 0 {
			is = true
			break
		}
	}

	if !is {
		return nil, fmt.Errorf("not sufficient funds")
	}

	txIn, err := c.client.PaySui(c.ctx, models.PaySuiRequest{
		Signer:      from,
		SuiObjectId: suiObjectIdList,
		Recipient:   recipientList,
		Amount:      amountList,
		GasBudget:   "1000000", // todo
	})
	if err != nil {
		return nil, fmt.Errorf("transfer sui: %w", err)
	}

	return (*UnSignTx)(&txIn), nil
}

func (c *SuiClient) BuildTransferTx(coinType, from, to string, amount uint64) (*UnSignTx, error) {
	res, err := c.client.SuiXGetCoins(c.ctx, models.SuiXGetCoinsRequest{
		Owner:    from,
		CoinType: coinType,
		Limit:    50,
	})
	if err != nil {
		return nil, fmt.Errorf("call SuiXGetCoins error: %w", err)
	}

	suiObjectId := ""

	for _, coins := range res.Data {
		balance := decimal.RequireFromString(coins.Balance)
		if balance.Cmp(decimal.NewFromUint64(amount)) >= 0 {
			suiObjectId = coins.CoinObjectId
			break
		}
	}

	if suiObjectId == "" {
		return nil, fmt.Errorf("not sufficient funds")
	}

	txIn, err := c.client.TransferSui(c.ctx, models.TransferSuiRequest{
		Signer:      from,
		SuiObjectId: suiObjectId,
		Recipient:   to,
		Amount:      fmt.Sprintf("%d", amount),
		GasBudget:   "1000000", // todo https://docs.sui.io/concepts/tokenomics/gas-in-sui
	})
	if err != nil {
		return nil, fmt.Errorf("transfer sui: %w", err)
	}

	return (*UnSignTx)(&txIn), nil
}

func (c *SuiClient) BuildCollectFoundTx(coinType string, from, to string) (*UnSignTx, error) {
	var suiObjectId []string

	for {
		res, err := c.client.SuiXGetCoins(c.ctx, models.SuiXGetCoinsRequest{
			Owner:    from,
			CoinType: coinType,
			Limit:    50,
		})
		if err != nil {
			return nil, fmt.Errorf("call SuiXGetCoins error: %w", err)
		}

		for _, coins := range res.Data {
			suiObjectId = append(suiObjectId, coins.CoinObjectId)
		}

		if !res.HasNextPage {
			break
		}
	}

	txIn, err := c.client.PayAllSui(c.ctx, models.PayAllSuiRequest{
		Signer:      from,
		SuiObjectId: suiObjectId,
		Recipient:   to,
		GasBudget:   "1000000", // todo
	})
	if err != nil {
		return nil, fmt.Errorf("transfer sui: %w", err)
	}

	return (*UnSignTx)(&txIn), nil
}

func (c *SuiClient) SendTx(tx *SignedTx) (string, error) {
	res, err := c.client.SuiExecuteTransactionBlock(c.ctx, models.SuiExecuteTransactionBlockRequest{
		TxBytes:   tx.TxBytes,
		Signature: []string{tx.Signature},
		Options: models.SuiTransactionBlockOptions{
			ShowInput:          true,
			ShowRawInput:       true,
			ShowEffects:        true,
			ShowEvents:         true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
		},
		RequestType: "WaitForLocalExecution",
	})
	if err != nil {
		return "", fmt.Errorf("send transaction block: %w", err)
	}

	return res.Digest, nil
}

func (c *SuiClient) WaitTxSuccess(digest string) error {
	result, err := c.client.SuiGetTransactionBlock(c.ctx, models.SuiGetTransactionBlockRequest{
		Digest: digest,
		Options: models.SuiTransactionBlockOptions{
			ShowInput:          true,
			ShowRawInput:       true,
			ShowEffects:        true,
			ShowEvents:         true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
		},
	})
	if err != nil {
		return fmt.Errorf("get transaction block: %w", err)
	}

	log.Info("sui SuiGetTransactionBlock: ", utils.FormatJSON(result))

	if result.Effects.Status.Error != "" {
		return errors.New(result.Effects.Status.Error)
	}

	return nil
}

// TryExecuteTx todo: bug.
func (c *SuiClient) TryExecuteTx(tx *UnSignTx) (string, error) {
	res, err := c.client.SuiDryRunTransactionBlock(c.ctx, models.SuiDryRunTransactionBlockRequest{TxBytes: tx.TxBytes})
	if err != nil {
		return "", fmt.Errorf("try execute transaction block: %w", err)
	}

	fmt.Println(utils.FormatJSON(res))

	return "", nil
}

func Blake2bHash(data string) []byte {
	txBytes, _ := base64.StdEncoding.DecodeString(data)
	message := models.MessageWithIntent(txBytes)
	digest := blake2b.Sum256(message)

	return digest[:]
}

func CoinType(coinAddress, coinName string) string {
	return fmt.Sprintf("%s::%s::%s", coinAddress, strings.ToLower(coinName), strings.ToUpper(coinName))
}
