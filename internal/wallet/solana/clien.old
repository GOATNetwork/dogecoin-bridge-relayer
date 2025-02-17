package solana

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

// https://github.com/gagliardetto/solana-go

type (
	RawSignTx solana.Message
	SignedTx  solana.Transaction
)

type UnSignTx struct {
	*solana.Transaction
}

// RawData signature data
func (t *UnSignTx) RawData() ([]byte, error) {
	return t.Message.MarshalBinary()
}

func (t *UnSignTx) BuildSolTransaction(signature solana.Signature) *SignedTx {
	t.Signatures = append(t.Signatures, signature)
	return (*SignedTx)(t.Transaction)
}

type SolClient struct {
	client *rpc.Client
}

func NewSolClient() *SolClient {
	client := rpc.New(rpc.MainNetBeta_RPC)
	return &SolClient{
		client: client,
	}
}

func NewDevSolClient() *SolClient {
	client := rpc.New(rpc.DevNet_RPC)
	return &SolClient{
		client: client,
	}
}

// BuildTokenTransfer https://solana.com/zh/docs/core/tokens#%E8%BD%AC%E7%A7%BB%E4%BB%A3%E5%B8%81
func (c *SolClient) BuildTokenTransfer(splToken, from, to, t solana.PublicKey, amount uint64) (*UnSignTx, error) {
	ins, err := token.NewTransferInstructionBuilder().
		SetOwnerAccount(splToken). // todo
		SetSourceAccount(from).
		SetDestinationAccount(to).
		SetAmount(amount).
		ValidateAndBuild()
	if err != nil {
		return nil, fmt.Errorf("build token transfer: %w", err)
	}
	return c.buildPrioritizationFees(from, solana.NewTransactionBuilder().AddInstruction(ins))
}

// BuildSolTransfer https://docs.anza.xyz/runtime/programs/#system-program
func (c *SolClient) BuildSolTransfer(from, to solana.PublicKey, amount uint64) (*UnSignTx, error) {
	transfer := system.NewTransferInstruction(amount, from, to)
	ins, err := transfer.ValidateAndBuild()
	if err != nil {
		return nil, fmt.Errorf("failed to build transfer instruction: %w", err)
	}
	return c.buildPrioritizationFees(from, solana.NewTransactionBuilder().AddInstruction(ins))
}

// https://solana.com/zh/developers/guides/advanced/how-to-use-priority-fees
// https://solana.com/zh/docs/rpc/http/getrecentprioritizationfees
func (c *SolClient) buildPrioritizationFees(from solana.PublicKey, builder *solana.TransactionBuilder) (*UnSignTx, error) {
	builder.SetFeePayer(from)
	out, err := c.client.GetRecentPrioritizationFees(context.Background(), solana.PublicKeySlice{from})
	if err != nil {
		return nil, fmt.Errorf("get recent prioritization fees: %w", err)
	}
	if len(out) > 0 {
		out = lo.Filter(out, func(item rpc.PriorizationFeeResult, index int) bool { return item.PrioritizationFee > 0 })
		minVal := lo.MinBy(out, func(a rpc.PriorizationFeeResult, b rpc.PriorizationFeeResult) bool {
			return a.PrioritizationFee < b.PrioritizationFee
		})
		if minVal.PrioritizationFee == 0 {
			minVal.PrioritizationFee = 100 // todo
		}
		prioritizationFee, err := computebudget.NewSetComputeUnitPriceInstruction(minVal.PrioritizationFee).ValidateAndBuild()
		if err != nil {
			return nil, fmt.Errorf("validate and build prioritization fee set: %w", err)
		}
		// default SetComputeUnitLimit = 1400000
		// default SetComputeUnitPrice = 0
		builder.AddInstruction(prioritizationFee)
	}

	tx, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("build transaction: %w", err)
	}
	return &UnSignTx{tx}, nil
}

// SendTransaction https://solana.com/zh/docs/rpc/http/sendtransaction
func (c *SolClient) SendTransaction(ctx context.Context, tx *solana.Transaction) error {
	sig, err := c.client.SendTransaction(ctx, tx)
	if err != nil {
		return fmt.Errorf("send transaction: %w", err)
	}

	return c.waitTxSuccess(ctx, sig)
}

func (c *SolClient) waitTxSuccess(ctx context.Context, sig solana.Signature) error {
	begin := time.Now()
	defer func() {
		log.Info("waitTxSuccess", "duration_ms", time.Since(begin).Milliseconds())
	}()

	var err error
	count := 3
	for count > 0 {
		_, err = c.client.GetConfirmedTransaction(ctx, sig)
		if err != nil && !errors.Is(err, rpc.ErrNotFound) {
			return fmt.Errorf("get confirmed transaction: %w", err)
		}
		time.Sleep(time.Second)
		count--
	}

	return err
}

// EstimateFee https://solana.com/zh/docs/rpc/http/simulatetransaction
// https://solana.com/zh/docs/core/fees#%E8%AE%A1%E7%AE%97%E5%8D%95%E5%85%83%E4%BB%B7%E6%A0%BC
// https://solana.com/zh/docs/core/fees#compute-unit-limit
func (c *SolClient) EstimateFee(ctx context.Context, msg string) (uint64, error) {
	// Get the fee the network will charge for a particular Message.
	//
	// **NEW**: This method is only available in solana-core v1.9 or newer. Please use
	// `getFees` for solana-core v1.8.
	_, err := c.client.GetFeeForMessage(ctx, msg, rpc.CommitmentFinalized)
	if err != nil {
		return 0, err
	}
	panic("not implemented")
}

func (c *SolClient) GetBalance(ctx context.Context, account solana.PublicKey) (uint64, error) {
	out, err := c.client.GetBalance(
		ctx,
		account,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return 0, fmt.Errorf("get balance: %w", err)
	}
	return out.Value, nil
}
