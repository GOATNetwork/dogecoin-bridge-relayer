package btc

import (
	"encoding/json"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	log "github.com/sirupsen/logrus"
)

type TxContext struct {
	c TxClient
}

func (w *WalletClient) ProcessWithdraw(task *db.Withdrawal) {
	hotAddress := address.HotAddressOfBtc(w.tss.ECPoint(w.ChainType()))
	log.Infof("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, task.ToAddress, task.Amount)

	txClient := NewTxClient(w.ctx, time.Minute*5, &chaincfg.MainNetParams, w.tss, w.chainInfo)
	txClient.AddTask(task)

	err := txClient.BuildTx(hotAddress, task.ToAddress, task.Amount.BigInt().Int64())
	if err != nil {
		log.Errorf("Build btc Tx err: %v", err)
		return
	}

	err = txClient.Sign()
	if err != nil {
		log.Errorf("Sign btc Tx err: %v", err)
		return
	}

	w.txContext.Store(txClient.TxHash(), txClient)
}

func (w *WalletClient) processTxSignResult(res *suite.SignRes) {
	bound := &TxBound{}

	err := json.Unmarshal(res.ExtraData, bound)
	utils.Assert(err)

	txCtx, ok := w.txContext.Load(bound.TxHash)
	if ok {
		defer w.txContext.Delete(bound.TxHash)

		if txClient, ok := txCtx.(*TxClient); ok {
			hash, err := chainhash.NewHashFromStr(res.DataDigest)
			if err != nil {
				log.Errorf("NewHashFromStr err: %v", err)
				return
			}

			is := txClient.AddWitnessSignature(hash[:], res.Signature)
			if is {
				err = txClient.SendTx()
				if err != nil {
					log.Errorf("Send btc Tx err: %v", err)
					return
				}
			}

			for _, id := range bound.TaskIds {
				w.SubmitPendingTask(id, hash.String())
			}
		}
	}
}
