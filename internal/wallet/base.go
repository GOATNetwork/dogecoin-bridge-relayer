package wallet

import (
	"context"
	"fmt"
	"time"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/patrickmn/go-cache"
)

type Wallet interface {
	ProcessCreatedTask(detailTask pool.Task[uint64])
	ProcessPendingTask(detailTask pool.Task[uint64])
	ConfirmTaskTxHash(taskID uint64, txHash string)
	ProcessDeposit(task *db.Deposit)
	ProcessWithdraw(task *db.Withdrawal)
	ProcessConsolidation(task *db.Consolidation)
	ProcessTransfer(task *db.Transfer)
}

type BaseWallet struct {
	layer2.Voter
	subWallet           Wallet
	ctx                 context.Context
	cancel              context.CancelFunc
	bus                 eventbus.Bus
	stateDB             *state.ContractState
	discussedTaskCache  *cache.Cache
	discussingTaskCache *cache.Cache
	taskQueue           *pool.Pool[uint64] // receive l2 task
	chainId             uint64
}

func NewBaseWallet(bus eventbus.Bus, stateDB *state.ContractState, voter layer2.Voter, chainId uint64) *BaseWallet {
	return &BaseWallet{
		bus:                 bus,
		stateDB:             stateDB,
		Voter:               voter,
		discussedTaskCache:  cache.New(5*time.Minute, 10*time.Minute),
		discussingTaskCache: cache.New(5*time.Minute, 10*time.Minute),
		taskQueue:           pool.NewTaskPool[uint64](),
		chainId:             chainId,
	}
}

func (w *BaseWallet) Ctx() context.Context {
	return w.ctx
}

func (w *BaseWallet) IsProd() bool {
	return config.AppConfig.IsProd()
}

func (w *BaseWallet) SetSubWallet(wallet Wallet) {
	w.subWallet = wallet
}

func (w *BaseWallet) AddTask(task pool.Task[uint64]) {
	w.taskQueue.Add(task)
}

func (w *BaseWallet) RemoveTask(taskId uint64) {
	w.taskQueue.Remove(taskId)
}

func (w *BaseWallet) GetTask(taskID uint64) (pool.Task[uint64], error) {
	t := w.taskQueue.Get(taskID)
	if t != nil {
		return t, nil
	}

	task, err := w.stateDB.GetUnCompletedTask(taskID)
	//todo
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	return w.GetOnlineTask(taskID)
	//}
	if err != nil {
		return nil, fmt.Errorf("taskID:%v, %w", taskID, err)
	}

	return task, err
}

//func (w *BaseWallet) GetOnlineTask(taskId uint64) (pool.Task[uint64], error) {
//	t, err := w.Voter.Tasks(taskId)
//	if err != nil {
//		return nil, err
//	}
//
//	detailTask := layer2.DecodeTaskOfEvent(t.Id, t.Result)
//
//	return detailTask, nil
//}

func (w *BaseWallet) IsDiscussed(taskID uint64) bool {
	_, ok := w.discussedTaskCache.Get(fmt.Sprintf("%d", taskID))
	return ok
}

func (w *BaseWallet) ContractState() *state.ContractState {
	return w.stateDB
}

func (w *BaseWallet) AddDiscussedTask(taskID uint64) {
	w.discussedTaskCache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (w *BaseWallet) AddDiscussingTask(taskID uint64) {
	w.discussingTaskCache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (w *BaseWallet) IsDiscussingTask(taskID uint64) bool {
	_, is := w.discussingTaskCache.Get(fmt.Sprintf("%d", taskID))
	return is
}

func (w *BaseWallet) RemoveDiscussingTask(taskID uint64) {
	w.discussingTaskCache.Delete(fmt.Sprintf("%d", taskID))
}

func (w *BaseWallet) GetAddressBalance(chainID uint64, minAmount uint64) []db.AddressBalance {
	address, _ := w.stateDB.GetAddressBalanceByCondition(chainID, minAmount)
	return address
}

func (w *BaseWallet) GetToken(ticker types.Byte32) (*db.TokenInfo, error) {
	return w.stateDB.GetTokenInfo(ticker)
}

func (w *BaseWallet) SubmitPendingTask(taskID uint64, txHash string) {
	w.SubmitTask(&db.TaskResult{
		TaskId:    taskID,
		State:     db.Pending,
		ExtraData: w.EncodeString(txHash),
	})
}

func (w *BaseWallet) EncodeString(txHash string) []byte {
	return utils.EncodeString(txHash)
}

func (w *BaseWallet) SubmitCompletedTask(taskID uint64) {
	w.SubmitTask(&db.TaskResult{
		TaskId: taskID,
		State:  db.Completed,
	})
}

func (w *BaseWallet) SubmitFailTask(taskID uint64) {
	w.SubmitTask(&db.TaskResult{
		TaskId: taskID,
		State:  db.Failed,
	})
}

func (w *BaseWallet) ChainId() uint64 {
	return w.chainId
}

func (w *BaseWallet) IsCanProcess() bool {
	return !w.Voter.IsSyncing()
}

func (w *BaseWallet) Bus() eventbus.Bus {
	return w.bus
}

func (w *BaseWallet) GetTopN(n int64) []pool.Task[uint64] {
	return w.taskQueue.GetTopN(n)
}
