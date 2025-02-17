package evm

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) tickerProcess(f func()) {
	if w.tss.IsProposer() {
		f()

		go func() {
			ticker := time.NewTicker(20 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-w.ctx.Done():
					log.Info("ticker process task done")
					return // Add return to exit the goroutine when context is done
				case <-ticker.C:
					if w.tss.IsProposer() {
						f()
					} else {
						return
					}
				}
			}
		}()
	}
}

func (w *WalletClient) LoadUnCompletedTasks() {
	for {
		if w.IsSyncing() {
			time.Sleep(time.Second * 2)
		} else {
			tasks, _ := w.ContractState().GetUnCompletedTasks()
			for _, task := range tasks {
				log.Debugf("tickerLoopUnCompletedTasks: task: %v", task)
				w.PostTask(task.DetailTask())
			}

			return
		}
	}
}
