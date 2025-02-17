package evm

import (
	"time"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) processTask() {
	if !w.tss.IsCanProposal() {
		return
	}

	log.Info("Start process task")

	tasks := w.GetTopN(2 * TopN)
	for _, task := range tasks {
		var t any = task
		switch v := t.(type) {
		case db.DetailTask:
			if v.ChainId() == w.ChainId() {
				if w.IsDiscussingTask(v.TaskID()) {
					return
				}

				go func(v db.DetailTask) {
					log.Debugf("processTask: received chain id:%v task from layer2 log scan: %v", w.ChainId(), v)
					w.AddDiscussingTask(v.TaskID())

					switch v.Status() {
					case db.Created:
						w.ProcessCreatedTask(v)
					case db.Pending:
						w.ProcessPendingTask(v)
					default:
						log.Errorf("taskID: %d, invalid task status: %v", v.TaskID(), v.Status())
					}
				}(v)
			}
		default:
			log.Warnf("unexpected type received in receiveL2TaskLoop: %T", v)
		}
	}
}

// receiveL2TaskLoop processes tasks received from layer2 log scan and handles them based on their status.
func (w *WalletClient) receiveL2TaskLoop() {
	taskEvent := w.Bus().Subscribe(eventbus.EventTask{})
	now := time.Now()
	isProcessing := false

	go func() {
		for {
			select {
			case <-w.ctx.Done():
				log.Info("evm wallet receive L2 task loop done")
				return
			case data := <-taskEvent:
				switch v := data.(type) {
				case db.DetailTask:
					if v.ChainId() == w.ChainId() {
						log.Debugf("received chain id:%v task from layer2 log scan: %v", w.ChainId(), v)

						switch v.Status() {
						case db.Created, db.Pending:
							w.AddTask(v)
						case db.Completed, db.Failed:
							w.RemoveTask(v.TaskID())
						default:
							log.Errorf("taskID: %d, invalid task status: %v", v.TaskID(), v.Status())
						}
					}
				case bool:
					if w.tss.IsProposer() {
						if !isProcessing || time.Since(now) > 30*time.Second {
							isProcessing = true
							now = time.Now()

							w.processTask()
						}
					} else {
						isProcessing = false
					}
				default:
					log.Warnf("unexpected type received in receiveL2TaskLoop: %T", v)
				}
			}
		}
	}()
}

// ProcessCreatedAddressTask processes a task to create a wallet address and submits the result.
func (w *WalletClient) ProcessCreatedAddressTask(task *db.CreateWallet) {
	log.Debugf("process created address task taskID: %v", task.TaskID())

	result := &db.TaskResult{
		TaskId: task.TaskID(),
	}
	coinType := types.GetCoinTypeByChain(task.Chain)
	userAddress := w.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)
	result.ExtraData = w.EncodeString(userAddress)
	result.State = db.Completed
	w.SubmitTask(result)
}

// ProcessCreatedTask processes tasks based on their type and delegates to specific handlers.
func (w *WalletClient) ProcessCreatedTask(detailTask pool.Task[uint64]) {
	log.Debugf("ProcessCreatedTask taskID: %v", detailTask.TaskID())

	switch task := detailTask.(type) {
	case *db.CreateWallet:
		w.ProcessCreatedAddressTask(task)
	case *db.Deposit:
		log.Warn("ProcessCreatedTask for Deposit is not implemented yet")
	case *db.Withdrawal:
		w.ProcessWithdraw(task)
	case *db.Consolidation:
		w.ProcessConsolidation(task)
	case *db.Transfer:
		w.ProcessTransfer(task)
	default:
		log.Errorf("unhandled task type in ProcessCreatedTask: %T", task)
	}
}

// ProcessPendingTask confirms the transaction hash for pending tasks.
func (w *WalletClient) ProcessPendingTask(detailTask db.DetailTask) {
	if detailTask.TransactionHash() == "" {
		log.Errorf("missing transaction hash for taskID: %d", detailTask.TaskID())
		return
	}

	w.ConfirmTaskTxHash(detailTask.TaskID(), detailTask.TransactionHash())
}
