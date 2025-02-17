package wallet

//func (w *BaseWallet) ConfirmTaskTxHash(taskID uint64, txHash string) {
//	panic("unimplemented")
//}

//func (w *BaseWallet) ProcessPendingTask(detailTask db.DetailTask) {
//	w.ConfirmTaskTxHash(detailTask.TaskID(), detailTask.TransactionHash())
//}

//func (w *BaseWallet) receiveL2TaskLoop() {
//	taskEvent := w.Bus().Subscribe(eventbus.EventTask{})
//
//	go func() {
//		for {
//			select {
//			case <-w.ctx.Done():
//				log.Info("evm wallet receive task event done")
//			case data := <-taskEvent: // from layer2 log scan
//				log.Info("received task from layer2 log scan: ", data)
//
//				switch v := data.(type) {
//				case db.DetailTask:
//					if v.ChainId() == w.ChainId() {
//						switch v.Status() {
//						case db.Created: // todo
//							w.AddTask(v)
//
//							if !w.IsSyncing() {
//								go w.subWallet.ProcessCreatedTask(v)
//							}
//						case db.Pending:
//							w.AddTask(v)
//
//							if !w.IsSyncing() {
//								go w.subWallet.ProcessPendingTask(v)
//							}
//						case db.Completed, db.Failed:
//							w.RemoveTask(v.TaskID())
//						default:
//							log.Errorf("taskID: %d, invalid task walletState : %v", v.TaskID(), v.Status())
//						}
//					}
//				}
//			}
//		}
//	}()
//}
//
//func (w *BaseWallet) ProcessCreatedTask(detailTask pool.Task[uint64]) {
//	switch task := detailTask.(type) {
//	case *db.Deposit:
//		// todo
//		// w.submitTask()
//
//	case *db.Withdrawal:
//		w.subWallet.ProcessWithdraw(task)
//	case *db.Consolidation:
//		w.subWallet.ProcessConsolidation(task)
//	case *db.Transfer:
//		w.subWallet.ProcessTransfer(task)
//	default:
//		log.Errorf("unhandled default case")
//	}
//}
