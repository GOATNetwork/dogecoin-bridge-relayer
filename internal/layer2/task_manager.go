package layer2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
)

type TaskManager interface {
	GetLatestTask() (contracts.Task, error)
	IsTaskCompleted(taskId uint64) (bool, error)
	GetTaskState(taskId uint64) (uint8, error)
	GetUncompletedTasks() ([]contracts.Task, error)
	EncodeSubmitTask(submitter common.Address, context []byte) []byte
	NextTaskId() (uint64, error)
	Tasks(taskId uint64) (contracts.Task, error)
	EncodeMarkTaskCompleted(taskId *big.Int, result []byte) []byte
}

func (l *Layer2Listener) GetLatestTask() (contracts.Task, error) {
	return l.taskManager.GetLatestTask(nil)
}

func (l *Layer2Listener) IsTaskCompleted(taskId uint64) (bool, error) {
	state, err := l.GetTaskState(taskId)
	return state == db.Completed, err
}

func (l *Layer2Listener) GetTaskState(taskId uint64) (uint8, error) {
	return l.taskManager.GetTaskState(nil, taskId)
}

func (l *Layer2Listener) GetUncompletedTasks() ([]contracts.Task, error) {
	return l.taskManager.GetUncompletedTasks(nil)
}

func (l *Layer2Listener) NextTaskId() (uint64, error) {
	return l.taskManager.NextTaskId(nil)
}

func (l *Layer2Listener) Tasks(taskId uint64) (contracts.Task, error) {
	return l.taskManager.Tasks(nil, taskId)
}

func (l *Layer2Listener) EncodeSubmitTask(submitter common.Address, context []byte) []byte {
	return contracts.EncodeFun(contracts.TaskManagerContractMetaData.ABI, "submitTask", submitter, context)
}

func (l *Layer2Listener) EncodeMarkTaskCompleted(taskId *big.Int, result []byte) []byte {
	return contracts.EncodeFun(contracts.TaskManagerContractMetaData.ABI, "markTaskCompleted", taskId, result)
}

func (l *Layer2Listener) EncodeMarkTaskCompletedBatch(taskIds []*big.Int, results [][]byte) []byte {
	return contracts.EncodeFun(contracts.TaskManagerContractMetaData.ABI, "markTaskCompleted_Batch", taskIds, results)
}
