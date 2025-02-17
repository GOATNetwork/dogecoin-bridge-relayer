package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) ProcessConsolidation(task *db.Consolidation) {
	balance, err := w.ContractState().GetAddressBalance(task.FromAddress, task.ContractAddress)
	if err != nil {
		log.Errorf("Failed to get address balance for %v: %v", task.FromAddress, err)
		return
	}

	if balance.Cmp(task.Amount) >= 0 {
		to := common.HexToAddress(address.HotAddressOfEth(w.tss.ECPoint(w.ChainType())))

		txHash, err := w.signTask(task.IsNative(), common.HexToAddress(task.FromAddress), to, common.HexToAddress(task.ContractAddress), task.Amount.BigInt(), task.TaskID(), task.Type())
		if err != nil {
			log.Errorf("Failed to sign task for %v: %v", task.TaskID(), err)
			return
		}

		w.SubmitPendingTask(task.TaskID(), txHash.String())
	}
}

func (w *WalletClient) ConfirmTaskTxHash(taskID uint64, txHash string) {
	receipt, err := w.WaitTxSuccess(common.HexToHash(txHash))
	if err != nil {
		log.Errorf("Failed to confirm transaction for %v: %v", taskID, err)
		return
	}

	state := db.Completed

	if receipt.Status == 0 {
		// updated status to fail
		log.Errorf("failed to submit transaction for taskId: %d, txHash: %s", taskID, txHash)

		// todo rollback task
		state = db.Failed
	}
	// updated status to completed
	w.SubmitTask(&db.TaskResult{
		TaskId: taskID,
		State:  uint8(state),
	})
}
