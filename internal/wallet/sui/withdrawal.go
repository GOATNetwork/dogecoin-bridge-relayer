package sui

import (
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

type TxContext struct {
	signerPubKey []byte
	tx           *UnSignTx
	task         pool.Task[uint64]
}

func (w *WalletClient) ProcessWithdraw(task *db.Withdrawal) {
	hotAddress := address.HotAddressOfSui(w.tss.ECPoint(w.ChainType()))
	log.Infof("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, task.ToAddress, task.Amount)

	unSignTx, err := w.client.BuildPaySuiTx(CoinType(task.ContractAddress, task.Symbol), hotAddress, []Recipient{
		{
			Recipient: task.ToAddress,
			Amount:    task.Amount.String(),
		},
	})
	if err != nil {
		log.Errorf("failed to build unsign tx: %v", err)
		return
	}

	proposal := new(big.Int).SetBytes(unSignTx.Blake2bHash())
	w.txContext.Store(task.TaskID(), &TxContext{
		signerPubKey: w.tss.GetPublicKey(hotAddress).SerializeCompressed(),
		tx:           unSignTx,
		task:         task,
	})

	req := &suite.SignReq{
		SeqId:      task.TaskID(),
		Type:       types.SignTxSessionType,
		ChainId:    w.ChainId(),
		Signer:     hotAddress,
		DataDigest: proposal.String(),
		SignData:   proposal.Bytes(),
		ExtraData:  nil,
	}

	w.tss.Sign(req)
}

func (w *WalletClient) ProcessTxSignResult(res *suite.SignRes) {
	taskID := res.SeqId
	txCtx, ok := w.txContext.Load(taskID)

	if ok {
		defer w.txContext.Delete(taskID)

		switch ctx := txCtx.(type) {
		case *TxContext:
			signTx := ctx.tx.SerializedSigWith(res.Signature, ctx.signerPubKey)
			log.Debugf("SuiTxContext: signTx: %v, signature: %x", signTx, res.Signature)

			digest, err := w.client.SendTx((*SignedTx)(signTx))
			if err != nil {
				log.Errorf("send transaction err: %v", err)
				return
			}

			w.SubmitPendingTask(taskID, digest)

			//err = w.client.WaitTxSuccess(digest)
			//if err != nil {
			//	log.Errorf("check transaction success err: %v", err)
			//	return
			//}
			log.Infof("successfully submitted transaction for taskId: %d,txHash: %v", ctx.task.TaskID(), digest)
		}
	}
}
