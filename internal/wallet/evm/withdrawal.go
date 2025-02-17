package evm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) signTask(isNative bool, from, to, contractAddress common.Address, amount *big.Int, taskId uint64, ty int) (common.Hash, error) {
	var (
		tx  *types.Transaction
		err error
	)

	if isNative {
		tx, err = w.BuildUnSignTx(
			from,
			to,
			amount,
			nil,
			ty,
			taskId,
		)
	} else {
		tx, err = w.BuildUnSignTx(
			from,
			contractAddress,
			nil,
			contracts.EncodeTransferOfERC20(from, to, amount),
			ty,
			taskId,
		)
	}

	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to build unsigned tx: %w", err)
	}

	ctx := w.NewTxContext(ty, taskId, tx)
	w.pendingTx.Store(ctx.TxHash(), ctx)
	defer w.pendingTx.Delete(ctx.TxHash())

	err = w.signTx(ctx)
	if err != nil {
		return ctx.TxHash(), fmt.Errorf("failed to sign tx: %w", err)
	}

	return ctx.TxHash(), w.SendSignedTx(ctx)
}

func (w *WalletClient) ProcessWithdraw(task *db.Withdrawal) {
	log.Debugf("processWithdrawTxSign taskId: %v", task.TaskID())

	hotAddress := common.HexToAddress(address.HotAddressOfEth(w.tss.ECPoint(w.ChainType())))
	to := common.HexToAddress(task.ToAddress)

	txHash, err := w.signTask(task.IsNative(), hotAddress, to, common.HexToAddress(task.ContractAddress), task.Amount.BigInt(), task.TaskID(), task.Type())
	if err != nil {
		log.Errorf("failed to sign task %d: %v", task.TaskID(), err)
		return
	}

	w.SubmitPendingTask(task.TaskID(), txHash.String())
}

func (w *WalletClient) processTxSignResult(res *suite.SignRes) {
	txCtx, ok := w.pendingTx.Load(res.DataDigest)
	if ok {
		switch ctx := txCtx.(type) {
		case *TxContext:
			ctx.sig = res.Signature
			ctx.notify <- res.Err
		}
	}
}
