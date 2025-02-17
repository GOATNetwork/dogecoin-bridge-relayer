package solana

import (
	"context"
	"fmt"
	"time"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/compute_budget"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/program/token"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

// github.com/blocto/solana-go-sdk.
type (
	RawSignTx types.Message
	SignedTx  types.Transaction
)

type UnSignTx struct {
	*types.Transaction
}

// RawData signature data.
func (t *UnSignTx) RawData() ([]byte, error) {
	return t.Message.Serialize()
}

func (t *UnSignTx) BuildTxWithSignature(signature types.Signature) *SignedTx {
	t.Signatures = append(t.Signatures, signature)
	return (*SignedTx)(t.Transaction)
}

type SolClient struct {
	client *client.Client
}

func NewSolClient() *SolClient {
	return &SolClient{client: client.NewClient(rpc.MainnetRPCEndpoint)}
}

func NewDevSolClient() *SolClient {
	return &SolClient{client: client.NewClient(rpc.DevnetRPCEndpoint)}
}

// BuildTokenTransfer
// https://blog.csdn.net/qq_44543317/article/details/136681475
// https://solana.com/zh/docs/core/tokens#%E8%BD%AC%E7%A7%BB%E4%BB%A3%E5%B8%81
func (c *SolClient) BuildTokenTransfer(s, f, t string, amount uint64, Decimals uint8) (*UnSignTx, error) {
	splToken, from, to := common.PublicKeyFromString(s), common.PublicKeyFromString(f), common.PublicKeyFromString(t)

	ataFrom, _, err := common.FindAssociatedTokenAddress(from, splToken)
	if err != nil {
		return nil, err
	}

	ataTo, _, err := common.FindAssociatedTokenAddress(to, splToken)
	if err != nil {
		return nil, err
	}

	res, err := c.client.GetLatestBlockhash(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get recent blockhash, err: %w", err)
	}

	tx := &types.Transaction{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        from,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				token.TransferChecked(token.TransferCheckedParam{
					From:    ataFrom,
					To:      ataTo,
					Mint:    splToken,
					Signers: []common.PublicKey{from},
					// Auth     common.PublicKey
					Amount:   amount,
					Decimals: Decimals,
				}),
			},
		}),
	}

	return &UnSignTx{tx}, nil
}

func (c *SolClient) BuildSolTransferWithAddress(from, to string, amount uint64) (*UnSignTx, error) {
	return c.BuildSolTransfer(common.PublicKeyFromString(from), common.PublicKeyFromString(to), amount)
}

// BuildSolTransfer https://docs.anza.xyz/runtime/programs/#system-program
func (c *SolClient) BuildSolTransfer(from, to common.PublicKey, amount uint64) (*UnSignTx, error) {
	// to fetch recent blockhash
	// todo verify blockHash
	res, err := c.client.GetLatestBlockhash(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get recent blockhash, err: %w", err)
	}
	// default SetComputeUnitLimit = 1400000
	// default SetComputeUnitPrice = 0
	// create a transfer tx
	tx := &types.Transaction{
		Signatures: nil,
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        from,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				compute_budget.SetComputeUnitPrice(compute_budget.SetComputeUnitPriceParam{
					MicroLamports: c.getPrioritizationFees(from),
				}),
				system.Transfer(system.TransferParam{
					From:   from,
					To:     to,
					Amount: amount,
				}),
			},
		}),
	}

	return &UnSignTx{tx}, nil
}

// https://solana.com/zh/developers/guides/advanced/how-to-use-priority-fees
// https://solana.com/zh/docs/rpc/http/getrecentprioritizationfees
func (c *SolClient) getPrioritizationFees(from common.PublicKey) uint64 {
	out, _ := c.client.GetRecentPrioritizationFees(context.Background(), []common.PublicKey{from})
	if len(out) > 0 {
		out = lo.Filter(out, func(item rpc.PrioritizationFee, index int) bool { return item.PrioritizationFee > 0 })

		minVal := lo.MinBy(out, func(a rpc.PrioritizationFee, b rpc.PrioritizationFee) bool {
			return a.PrioritizationFee < b.PrioritizationFee
		})
		if minVal.PrioritizationFee > 0 {
			return minVal.PrioritizationFee
		}
	}

	return 100
}

// SendTransaction https://solana.com/zh/docs/rpc/http/sendtransaction
func (c *SolClient) SendTransaction(ctx context.Context, tx *types.Transaction) (string, error) {
	sig, err := c.client.SendTransaction(ctx, *tx)
	if err != nil {
		return sig, fmt.Errorf("send transaction: %w", err)
	}

	return sig, nil
}

func (c *SolClient) SyncSendTransaction(ctx context.Context, tx *types.Transaction) (string, error) {
	sig, err := c.SendTransaction(ctx, tx)
	if err != nil {
		return sig, err
	}

	return sig, c.WaitTxSuccess(ctx, sig)
}

func (c *SolClient) WaitTxSuccess(ctx context.Context, sig string) error {
	begin := time.Now()
	defer func() {
		log.Infof("waitTxSuccess, duration_ms: %v", time.Since(begin).Milliseconds())
	}()

	var err error

	count := 3
	for count > 0 {
		// status, err := c.client.GetSignatureStatus(ctx, sig)
		status, err := c.client.GetSignatureStatusWithConfig(ctx, sig, client.GetSignatureStatusesConfig{SearchTransactionHistory: true})
		if err != nil {
			return fmt.Errorf("get confirmed transaction: %w", err)
		}

		log.Debugf("get confirmed transaction status: %v", utils.FormatJSON(status))

		if status == nil {
			time.Sleep(time.Second)

			count--

			continue
		}

		if status.Err != nil {
			return fmt.Errorf("get confirmed transaction: %v", status.Err)
		}

		if *status.ConfirmationStatus == rpc.CommitmentProcessed {
			time.Sleep(time.Second)

			count--
		} else {
			return nil
		}
	}

	return err
}

// EstimateFee https://solana.com/zh/docs/rpc/http/simulatetransaction
// https://solana.com/zh/docs/core/fees#%E8%AE%A1%E7%AE%97%E5%8D%95%E5%85%83%E4%BB%B7%E6%A0%BC
// https://solana.com/zh/docs/core/fees#compute-unit-limit
func (c *SolClient) EstimateFee(ctx context.Context, message types.Message) (uint64, error) {
	// Get the fee the network will charge for a particular Message.
	//
	// **NEW**: This method is only available in solana-core v1.9 or newer. Please use
	// `getFees` for solana-core v1.8.
	_, err := c.client.GetFeeForMessage(ctx, message)
	if err != nil {
		return 0, err
	}

	panic("not implemented")
}

func (c *SolClient) GetBalanceOfSol(ctx context.Context, account common.PublicKey) (uint64, error) {
	out, err := c.client.GetBalance(
		ctx,
		account.String(),
	)
	if err != nil {
		return 0, fmt.Errorf("get balance: %w", err)
	}

	return out, nil
}

func (c *SolClient) GetBalanceOfToken(ctx context.Context, account common.PublicKey) (uint64, error) {
	out, err := c.client.GetTokenAccountBalance(ctx, account.String())
	if err != nil {
		return 0, fmt.Errorf("get balance: %w", err)
	}

	return out.Amount, nil
}
