package evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const defaultGasLimit = 1000000

type WalletClient struct {
	*wallet.BaseWallet
	chainInfo           *config.Chain
	ctx                 context.Context
	cancel              context.CancelFunc
	client              *ethclient.Client
	submitterPrivateKey *ecdsa.PrivateKey
	submitter           common.Address
	chainID             atomic.Int64
	walletState         *state.EvmWalletState
	tss                 suite.TssService
	currentVoterNonce   *atomic.Uint64
	submitTaskQueue     *pool.Pool[uint64] // submit task
	operationsQueue     *pool.Pool[uint64] // pending batch task
	pendingTx           sync.Map           // txHash: EvmTx
}

func (w *WalletClient) Start(context.Context) {
	log.Info("evm wallet client is starting...")
	w.tss.RegisterTssClient(w)
	w.receiveL2TaskLoop()
	w.receiveSubmitTaskLoop()
	w.removeSubmitTaskTaskLoop()

	// enable master chain submit operation
	if config.AppConfig.IsMaster(w.chainInfo.ChainId) {
		w.loopProcessOperation()
	}

	w.tickerRetrySendTx()
	w.LoadUnCompletedTasks()
}

func (w *WalletClient) Stop(context.Context) {
	w.cancel()
}

func NewWallet(
	bus eventbus.Bus,
	tss suite.TssService,
	voter layer2.Voter,
	stateDB *state.ContractState,
	walletState *state.EvmWalletState,
	chainInfo *config.Chain,
) *WalletClient {
	client, err := ethclient.Dial(chainInfo.Rpc.Url)
	utils.Assert(err)

	currentNonce := &atomic.Uint64{}
	nonce, _ := voter.TssNonce()

	if nonce != nil {
		currentNonce.Store(nonce.Uint64())
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &WalletClient{
		chainInfo:           chainInfo,
		BaseWallet:          wallet.NewBaseWallet(bus, stateDB, voter, chainInfo.ChainId),
		ctx:                 ctx,
		cancel:              cancel,
		client:              client,
		submitterPrivateKey: config.SubmitterPrivateKey,
		submitter:           crypto.PubkeyToAddress(config.SubmitterPrivateKey.PublicKey),
		pendingTx:           sync.Map{},
		chainID:             atomic.Int64{},
		walletState:         walletState,
		tss:                 tss,
		currentVoterNonce:   currentNonce,
		submitTaskQueue:     pool.NewTaskPool[uint64](),
		operationsQueue:     pool.NewTaskPool[uint64](),
	}
}

func (w *WalletClient) signatureHash(tx *types.Transaction) common.Hash {
	signer := types.LatestSignerForChainID(big.NewInt(w.chainID.Load()))
	signatureHash := signer.Hash(tx)

	return signatureHash
}

func (w *WalletClient) BalanceOf(erc20Token, owner common.Address) (*big.Int, error) {
	erc20, err := contracts.NewERC20(erc20Token, w.client)
	if err != nil {
		return nil, err
	}

	return erc20.BalanceOf(nil, owner)
}

func (w *WalletClient) BalanceAt(owner common.Address) (*big.Int, error) {
	return w.client.BalanceAt(w.ctx, owner, nil)
}

func (w *WalletClient) ChainID() *big.Int {
	if w.chainID.Load() == 0 {
		var (
			chainID *big.Int
			err     error
		)

		err = backoff.Retry(
			func() error {
				chainID, err = w.client.ChainID(w.ctx)
				return err
			},
			backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 3),
		)
		if err != nil {
			return big.NewInt(w.chainID.Load())
		}

		w.chainID.Store(chainID.Int64())
	}

	return big.NewInt(w.chainID.Load())
}

func (w *WalletClient) SendSignedTx(ctx *TxContext) error {
	err := w.sendTransaction(ctx)
	if errors.Is(err, ErrIntrinsicGasTooLow) {
		return w.SpeedSendTx(ctx)
	}

	if errors.Is(err, ErrNonceTooLow) || errors.Is(err, ErrReplacement) {
		return w.SendUnSignTx(ctx)
	}

	return err
}

func (w *WalletClient) EstimateGasAPI(msg ethereum.CallMsg) (uint64, error) {
	gasLimit, err := w.client.EstimateGas(w.ctx, msg)
	if err != nil {
		err = errors.Join(ErrEstimateGas, wrapError(err))
		gasLimit = defaultGasLimit

		if errors.Is(err, ErrInsufficientFunds) {
			err = errors.Join(fmt.Errorf("insufficient funds for sender: %v", msg.From), err)
		}
	}

	return gasLimit, err
}

func (w *WalletClient) EstimateGas(account, contractAddress common.Address, data []byte) (uint64, error) {
	// Estimate GasPrice
	// gasPrice := opt.GasPrice
	price, err := w.client.SuggestGasPrice(w.ctx)
	if err != nil {
		return 0, fmt.Errorf("SuggestGasPrice error: %w, account address: %v", err, account)
	}

	msg := ethereum.CallMsg{
		From:     account,
		To:       &contractAddress,
		Data:     data,
		GasPrice: price,
	}

	return w.EstimateGasAPI(msg)
}

func (w *WalletClient) BuildUnSignTx(
	account, contractAddress common.Address,
	value *big.Int, calldata []byte,
	ty int,
	seqID uint64,
) (*types.Transaction, error) {
	head, err := w.client.HeaderByNumber(w.ctx, nil)
	if err != nil {
		return nil, wrapError(err)
	}

	// EIP1559
	// https://learnblockchain.cn/article/8593
	// https://blog.csdn.net/vigor2323/article/details/122817104
	// https://metamask.io/1559/
	// https://www.blocknative.com/blog/eip-1559-fees
	// 1、Transaction Fee=GasUsed*GasPrice=GasUsed*(BaseFee+MaxPriorityFee)；
	// 2、MaxFee = MaxGasPrice = (2*BaseFee)+MaxPriorityFee
	// 3、Burnt=BaseFee*GasUsed
	// 4、TxSavingsFees=MaxFee*GasUsed−(BaseFee+MaxPriorityFee)*GasUsed = (MaxFee-(BaseFee+MaxPriorityFee))*GasUsed
	// 5、Tip(minter get fee) = Min(Max fee - Base fee, Max priority fee)
	// 6、usedGasPrice = min(MaxPriorityFeePerGas + basefee, MaxFeePerGas)
	gasTipCap, err := w.client.SuggestGasTipCap(w.ctx)
	if err != nil {
		return nil, wrapError(err)
	}

	gasFeeCap := new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
	)
	if gasFeeCap.Cmp(gasTipCap) < 0 {
		return nil, fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", gasFeeCap, gasTipCap)
	}

	//// Estimate GasPrice
	//price, err := w.client.SuggestGasPrice(w.ctx)
	//if err != nil {
	//	return nil, fmt.Errorf("SuggestGasPrice error: %w, account address: %v", err, account)
	//}

	// Estimate GasLimit
	msg := ethereum.CallMsg{
		From: account,
		To:   &contractAddress,
		// GasPrice:  price,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      calldata,
		Value:     value,
	}

	gas, err := w.EstimateGasAPI(msg)
	if err != nil {
		return nil, wrapError(err)
	}

	// extend gas limit 20%
	gasLimit := decimal.NewFromUint64(gas).Mul(decimal.NewFromFloat(1.2))

	nextNonce := decimal.Zero
	_ = nextNonce

	for {
		nonce, err := w.client.PendingNonceAt(w.ctx, account)
		if err != nil {
			return nil, wrapError(err)
		}

		nextNonce = decimal.NewFromUint64(nonce)

		latestNonce, err := w.walletState.LatestNonce(nil, w.ChainId(), account)
		if err != nil {
			return nil, err
		}

		if nextNonce.Cmp(latestNonce) >= 0 {
			break
		}

		time.Sleep(200 * time.Millisecond)
	}

	//latestNonce := w.nonce.Load()
	//if latestNonce >= nextNonce {
	//	nextNonce = latestNonce + 1
	//	w.nonce.Store(nextNonce)
	//}

	chainId := w.ChainID()

	baseTx := &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nextNonce.BigInt().Uint64(),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit.BigInt().Uint64(),
		To:        &contractAddress,
		Value:     value,
		Data:      calldata,
	}

	tx := types.NewTx(baseTx)
	err = w.walletState.CreateTx(
		nil,
		ty,
		seqID,
		account,
		tx,
	)

	return tx, err
}

func (w *WalletClient) ExtendTxGas(tx *types.Transaction) *types.Transaction {
	oldTx := tx

	gasLimit := oldTx.Gas()
	if gasLimit < defaultGasLimit {
		gasLimit = defaultGasLimit
	}

	// extend 20%
	gasTipCap := decimal.NewFromFloat(1.2).Mul(decimal.NewFromBigInt(oldTx.GasTipCap(), 0)).BigInt()
	baseTx := &types.DynamicFeeTx{
		To:        oldTx.To(),
		Nonce:     oldTx.Nonce(),
		GasTipCap: gasTipCap,
		GasFeeCap: oldTx.GasFeeCap(),
		Gas:       gasLimit * 2,
		Value:     oldTx.Value(),
		Data:      oldTx.Data(),
	}
	unSignedTx := types.NewTx(baseTx)

	return unSignedTx
}

func (w *WalletClient) SignOperationNewTx(unSignedTx *types.Transaction) ([]byte, error) {
	signer := types.LatestSignerForChainID(unSignedTx.ChainId())
	txHash := signer.Hash(unSignedTx).Bytes()
	sig, err := crypto.Sign(signer.Hash(unSignedTx).Bytes(), w.submitterPrivateKey)
	log.Debugf("SignOperationNewTx:sig %x, tx hash: %x,address: %v，chain id: %v", sig, txHash, crypto.PubkeyToAddress(w.submitterPrivateKey.PublicKey), unSignedTx.ChainId().Uint64())

	return sig, err
}

func (w *WalletClient) sendTransaction(ctx *TxContext) error {
	tx, err := ctx.UnSignTx().WithSignature(w.Signer(), ctx.sig)
	if err != nil {
		return err
	}

	err = w.client.SendTransaction(w.ctx, tx)
	if err != nil {
		err = errors.Join(ErrSendTransaction, wrapError(err))
		_, is := lo.Find(failErrorList, func(item error) bool { return errors.Is(err, item) })

		if is {
			dbErr := w.walletState.UpdateFailTx(tx.Hash(), err)
			err = errors.Join(err, dbErr)
		}

		return err
	}

	return w.walletState.UpdatePendingTx(tx.Hash())
}

// SpeedSendTx
// 1. resend order tx
// 2. speed resend.
func (w *WalletClient) SpeedSendTx(ctx *TxContext) error {
	oldOrderTx := ctx.tx
	unSignedTx := w.ExtendTxGas(oldOrderTx)

	err := w.walletState.CreateTx(
		nil,
		ctx.ty,
		ctx.SeqID(),
		w.From(ctx.tx),
		unSignedTx,
	)
	if err != nil {
		return err
	}

	ctx.tx = unSignedTx

	w.pendingTx.Delete(ctx.DigestHash())
	w.pendingTx.Store(unSignedTx.Hash(), ctx)

	err = w.signTx(ctx) // todo
	if err != nil {
		return wrapError(err)
	}

	w.pendingTx.Store(unSignedTx.Hash(), ctx)

	return w.SendSignedTx(ctx)
}

func (w *WalletClient) SendUnSignTx(ctx *TxContext) error {
	oldTx := ctx.UnSignTx()
	tx := oldTx
	sender := w.From(tx)

	var err error

	ctx.tx, err = w.BuildUnSignTx(sender, *tx.To(), tx.Value(), tx.Data(), ctx.ty, ctx.SeqID())
	if err != nil {
		return err
	}

	w.pendingTx.Delete(ctx.DigestHash())
	w.pendingTx.Store(ctx.DigestHash(), ctx)

	err = w.signTx(ctx)
	if err != nil {
		return err
	}

	return w.SendSignedTx(ctx)
}

func (w *WalletClient) RawTxBytes(tx *types.Transaction) []byte {
	signer := w.Signer()
	return signer.Hash(tx).Bytes()
}

func (w *WalletClient) Signer() types.Signer {
	chainId := w.ChainID()

	signer := types.LatestSignerForChainID(chainId)

	return signer
}

func (w *WalletClient) From(tx *types.Transaction) common.Address {
	signer := w.Signer()
	address, err := signer.Sender(tx)
	utils.Assert(err)

	return address
}

func (w *WalletClient) TransactionWithSignature(tx *types.Transaction, signature []byte) (*types.Transaction, error) {
	signer := w.Signer()

	return tx.WithSignature(signer, signature)
}

func (w *WalletClient) WaitTxSuccess(txHash common.Hash) (*types.Receipt, error) {
	begin := time.Now()
	defer func() {
		log.Infof("WaitTxSuccess, duration_ms: %v", time.Since(begin).Milliseconds())
	}()

	count := 60
	for count > 0 {
		r, err := w.client.TransactionReceipt(w.ctx, txHash)
		// TransactionReceipt can is nil
		if err != nil {
			if errors.Is(err, ethereum.NotFound) {
				time.Sleep(time.Second)

				count--
			} else {
				return nil, errors.Join(fmt.Errorf("call TransactionReceipt method error, txHash: %s", txHash), err, ErrWallet)
			}
		} else {
			dbErr := w.walletState.UpdateBookedTx(txHash)
			return r, dbErr
		}
	}

	return nil, ErrTxFoundTimeOut
}

func (w *WalletClient) IsOnChain(txHash common.Hash) bool {
	receipt, err := w.client.TransactionReceipt(w.ctx, txHash)
	if receipt != nil && err == nil {
		return true
	}

	return false
}

func (w *WalletClient) IsPending(txHash common.Hash) bool {
	_, ok := w.pendingTx.Load(txHash)
	if ok {
		return ok
	}

	_, isPending, _ := w.client.TransactionByHash(w.ctx, txHash)

	return isPending
}

// IsOnline = IsPending + IsOnChain.
func (w *WalletClient) IsOnline(txHash common.Hash) bool {
	tx, isPending, _ := w.client.TransactionByHash(w.ctx, txHash)
	return isPending || tx != nil
}

func (w *WalletClient) IsCanProcess(txHash common.Hash) bool {
	_, isPending := w.pendingTx.Load(txHash)
	if !isPending {
		tx, _, _ := w.client.TransactionByHash(w.ctx, txHash)
		return tx == nil
	}

	return false
}

func wrapError(err error) error {
	for _, wrap := range wrapErrorList {
		if utils.ContainErr(err, wrap) {
			return errors.Join(err, wrap, ErrWallet)
		}
	}

	return err
}

func (w *WalletClient) tickerRetrySendTx() {
	ticker := time.NewTicker(30 * time.Second)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-w.ctx.Done():
				log.Info("evm wallet tickerRetrySendTx done")
				return
			case <-ticker.C:
				w.tickerRetryUpdateTx()
			}
		}
	}()
}

func (w *WalletClient) tickerRetryUpdateTx() {
	txs, _ := w.walletState.PendingBlockchainTransactions(nil, w.ChainId())
	group := sync.WaitGroup{}

	for _, tx := range txs {
		group.Add(1)

		go func() {
			defer group.Done()

			if w.IsCanProcess(tx.TxHash) {
				ctx := w.NewTxContext(tx.Type, tx.SeqID, tx.Tx())
				w.pendingTx.Store(ctx.TxHash(), ctx)
				defer w.pendingTx.Delete(ctx.TxHash())

				err := w.SendUnSignTx(ctx)
				if err != nil {
					log.Infof("evm wallet SendUnSignTx: tx hash:%v, err:%v", ctx.TxHash(), err)
					return
				}

				_, err = w.WaitTxSuccess(ctx.TxHash())
				if err != nil {
					log.Infof("evm wallet WaitTxSuccess: err:%v", err)
				}
			}
		}()
	}

	group.Wait()
}
