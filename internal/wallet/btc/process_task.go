package btc

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) receiveL2TaskLoop() {
	taskEvent := w.Bus().Subscribe(eventbus.EventTask{})

	go func() {
		for {
			select {
			case <-w.ctx.Done():
				log.Info("evm wallet receive task event done")
			case data := <-taskEvent: // from layer2 log scan
				log.Info("received task from layer2 log scan: ", data)

				switch v := data.(type) {
				case db.DetailTask:
					if v.ChainId() == w.ChainId() {
						switch v.Status() {
						case db.Created: // todo
							w.AddTask(v)

							if w.tss.IsCanProposal() {
								go w.ProcessCreatedTask(v)
							}
						case db.Pending:
							w.AddTask(v)

							if w.tss.IsCanProposal() {
								go w.ProcessPendingTask(v)
							}
						case db.Completed, db.Failed:
							w.RemoveTask(v.TaskID())
						default:
							log.Errorf("taskID: %d, invalid task walletState : %v", v.TaskID(), v.Status())
						}
					}
				}
			}
		}
	}()
}

func (w *WalletClient) ProcessCreatedTask(detailTask pool.Task[uint64]) {
	switch task := detailTask.(type) {
	case *db.Deposit:
		// todo
		// w.submitTask()

	case *db.Withdrawal:
		w.ProcessWithdraw(task)
	case *db.Consolidation:
		w.ProcessConsolidation(task)
	}
}

func (w *WalletClient) ProcessPendingTask(detailTask db.DetailTask) {
	w.ConfirmTaskTxHash(detailTask.TaskID(), detailTask.TransactionHash())
}

func (w *WalletClient) ConfirmTaskTxHash(taskID uint64, txHash string) {
	hash, err := chainhash.NewHashFromStr(txHash)
	if err != nil {
		log.Errorf("confirmTaskTxHash: invalid tx hash: %v", txHash)
		return
	}

	err = w.client.WaitTxSuccess(*hash)
	if err != nil {
		log.Errorf("Failed to confirm transaction for %v: %v", taskID, err)
		return
	}

	// updated status to completed
	w.SubmitCompletedTask(taskID)
}
