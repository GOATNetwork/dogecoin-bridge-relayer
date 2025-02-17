package solana

import (
	"math/big"

	soltypes "github.com/blocto/solana-go-sdk/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

// TxContext encapsulates an unsigned transaction and a task for processing transactional operations.
type TxContext struct {
	tx   *UnSignTx
	task pool.Task[uint64]
}

func (w *WalletClient) ProcessWithdraw(task *db.Withdrawal) {
	hotAddress := address.HotAddressOfSolana(w.tss.ECPoint(w.ChainType()))
	log.Infof("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, task.ToAddress, task.Amount)

	var (
		tx  *UnSignTx
		err error
	)

	if task.IsNative() {
		tx, err = w.client.BuildSolTransferWithAddress(hotAddress, task.ToAddress, task.Amount.BigInt().Uint64())
	} else {
		tx, err = w.client.BuildTokenTransfer(task.ContractAddress, hotAddress, task.ToAddress, task.Amount.BigInt().Uint64(), task.Decimals)
	}

	if err != nil {
		log.Errorf("failed to build unsign tx: %v", err)
		return
	}

	raw, err := tx.RawData()
	if err != nil {
		log.Errorf("failed to build unsign tx: %v", err)
		return
	}

	log.Infof("raw: %x", raw)
	proposal := new(big.Int).SetBytes(raw)

	w.txContext.Store(task.TaskID(), &TxContext{
		tx:   tx,
		task: task,
	})

	req := &suite.SignReq{
		SeqId:      task.TaskID(),
		Type:       types.SignTxSessionType,
		ChainId:    w.ChainId(),
		Signer:     hotAddress,
		DataDigest: proposal.String(), // todo
		SignData:   proposal.Bytes(),
	}
	w.tss.Sign(req)
}

func (w *WalletClient) processTxSignResult(res *suite.SignRes) {
	taskID := res.SeqId
	txCtx, ok := w.txContext.Load(taskID)

	if ok {
		defer w.txContext.Delete(taskID)

		switch ctx := txCtx.(type) {
		case *TxContext:
			unSignRawData, _ := ctx.tx.RawData()
			log.Debugf("SolTxContext: unSignRawData: %x, signature: %x", unSignRawData, res.Signature)

			sig, err := w.client.SendTransaction(w.ctx, (*soltypes.Transaction)(ctx.tx.BuildTxWithSignature(res.Signature)))
			if err != nil {
				log.Errorf("send transaction err: %v", err)
				return
			}

			w.SubmitPendingTask(taskID, sig)

			log.Infof("successfully submitted transaction for taskId: %d,txHash: %v", ctx.task.TaskID(), sig)
		}
	}
}
