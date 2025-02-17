package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) ProcessTransfer(task *db.Transfer) {
	log.Debugf("processTransferTxSign taskId: %v", task.TaskID())

	from := common.HexToAddress(task.FromAddress)
	to := common.HexToAddress(task.ToAddress)

	txHash, err := w.signTask(
		task.IsNative(),
		from,
		to,
		common.HexToAddress(task.ContractAddress),
		task.Amount.BigInt(),
		task.TaskID(),
		task.Type(),
	)
	if err != nil {
		log.Errorf("Error signing task %d: %v", task.TaskID(), err)
		return
	}

	w.SubmitPendingTask(task.TaskID(), txHash.String())
	log.Debugf("Successfully processed transfer for taskId: %v with txHash: %v", task.TaskID(), txHash.String())
}
