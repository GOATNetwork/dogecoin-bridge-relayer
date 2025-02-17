package solana

import (
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) ProcessTransfer(task *db.Transfer) {
	var (
		tx  *UnSignTx
		err error
	)

	if task.IsNative() {
		tx, err = w.client.BuildSolTransferWithAddress(task.FromAddress, task.ToAddress, task.Amount.BigInt().Uint64())
	} else {
		tx, err = w.client.BuildTokenTransfer(task.ContractAddress, task.FromAddress, task.ToAddress, task.Amount.BigInt().Uint64(), task.Decimals)
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
		Signer:     task.FromAddress,
		DataDigest: proposal.String(), // todo
		SignData:   proposal.Bytes(),
	}
	w.tss.Sign(req)
}
