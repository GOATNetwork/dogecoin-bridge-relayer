package evm

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts/codec"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

type Operations struct {
	Nonce     *big.Int
	Operation []codec.Operation
	Hash      common.Hash
	DataHash  common.Hash
	Signature []byte
}

func (o *Operations) TaskID() uint64 {
	return o.Nonce.Uint64()
}

func (o *Operations) Type() int {
	return db.TaskTypeOperations
}

func (m *WalletClient) GetDiscussedOperation(id uint64) *Operations {
	ops := m.operationsQueue.Get(id)
	if ops == nil {
		return nil
	}

	return ops.(*Operations)
}

func (w *WalletClient) Operation(task pool.Task[uint64]) *codec.Operation {
	operation := &codec.Operation{
		TaskId: task.TaskID(),
	}
	switch value := task.(type) {
	case *db.TaskResult:
		operation.State = value.State
		operation.ExtraData = value.ExtraData
	default:
		log.Errorf("Unhandled task type: %T", value)
		return nil
	}

	log.Debugf("Operation: %v", utils.FormatJSON(operation))

	return operation
}

func (w *WalletClient) saveOperations(nonce *big.Int, ops []codec.Operation, dataHash, hash common.Hash) {
	w.currentVoterNonce.Store(nonce.Uint64())
	operations := &Operations{
		Nonce:     nonce,
		Operation: ops,
		Hash:      hash,
		DataHash:  dataHash,
	}

	lo.ForEach(ops, func(item codec.Operation, _ int) { w.AddDiscussingTask(item.TaskId) })
	w.operationsQueue.Add(operations)
}

func (w *WalletClient) loopProcessOperation() {
	taskEvent := w.Bus().Subscribe(eventbus.EventTask{})

	go func() {
		now := time.Now()
		isProcessing := false

		for {
			select {
			case <-w.ctx.Done():
				log.Info("approve proposal done")
				return // Add return to exit the goroutine when context is done

			case synced := <-taskEvent:
				log.Debugf("received synced completed event: %v", synced)

				if w.tss.IsProposer() {
					if !isProcessing || time.Since(now) > 30*time.Second {
						isProcessing = true
						now = time.Now()

						w.processOperation()
					}
				} else {
					isProcessing = false
				}
			}
		}
	}()
}

const TopN = 20 // TopN defines the maximum number of tasks to process in a batch

func (w *WalletClient) processOperation() {
	if !w.tss.IsCanProposal() {
		return
	}

	nonce, err := w.Voter.TssNonce()
	if err != nil {
		log.Errorf("Error getting tss nonce: %v", err)
		return
	}

	if nonce.Cmp(big.NewInt(int64(w.currentVoterNonce.Load()))) < 0 {
		log.Errorf("tss nonce: %v, currentVoterNonce: %v", nonce, w.currentVoterNonce.Load())
		return
	}

	if w.operationsQueue.IsExist(nonce.Uint64()) {
		log.Warnf("operations is exist: %d", nonce.Uint64())

		return
	}

	log.Info("Starting batch proposal processing")

	tasks := w.submitTaskQueue.GetTopN(TopN)

	log.Debugf("Retrieved tasks: %v", tasks)

	operations := make([]codec.Operation, 0)

	var dataHash, msg common.Hash

	if w.operationsQueue.IsExist(nonce.Uint64()) {
		log.Warnf("operations is exist: %d", nonce.Uint64())
		ops := w.operationsQueue.Get(nonce.Uint64()).(*Operations)
		dataHash = ops.DataHash
		msg = ops.Hash
	} else {
		for _, task := range tasks {
			operations = append(operations, *w.Operation(task))
		}

		if len(operations) == 0 {
			log.Warn("No operations to process, operationsQueue is empty")
			return
		}

		dataHash, msg = w.GenerateVerifyTaskUnSignMsg(operations, nonce)
		err = w.walletState.SaveOperations(nil, &db.Operations{
			TssNonce: nonce.Uint64(),
			DataHash: dataHash,
			Data:     contracts.PackEvent(codec.VoterCodecMetaData, "operations", operations, nonce, big.NewInt(int64(w.ChainId()))),
		})
		if err != nil {
			log.Errorf("Error saving operations: %v", err)
			return
		}
	}

	log.Infof("Generated nonce: %v, dataHash: %v, msg: %v", nonce, dataHash, msg)

	data := lo.Map(operations, func(item codec.Operation, index int) uint64 { return item.TaskId })
	batchData := types.BatchData{Ids: data}

	if !w.tss.IsMeeting(msg.String()) {
		signReq := &suite.SignReq{
			SeqId:      nonce.Uint64(),
			Type:       types.SignOperationSessionType,
			ChainId:    w.ChainId(),
			Signer:     w.tss.TssSigner().String(),
			DataDigest: msg.String(),
			SignData:   msg.Bytes(),
			ExtraData:  batchData.Bytes(),
		}
		w.tss.Sign(signReq)
		w.saveOperations(nonce, operations, dataHash, msg)
	} else {
		log.Warn("Batch task signTx already processed")
	}
}

func (w *WalletClient) processOperationSignResult(operations *Operations) {
	// 1. save db
	// 2. update status
	if w.tss.IsProposer() {
		log.Info("proposer submit signature")

		calldata := w.EncodeVerifyAndCall(operations.Operation, operations.Signature)
		log.Infof("calldata: %x, signature: %x, nonce: %v, DataHash: %v, hash: %v", calldata, operations.Signature, operations.Nonce, operations.DataHash, operations.Hash)

		tx, err := w.BuildUnSignTx(
			w.tss.LocalSubmitter(),
			common.HexToAddress(config.AppConfig.Contract.Voter),
			big.NewInt(0),
			calldata,
			operations.Type(),
			operations.TaskID(),
		)
		if err != nil {
			log.Errorf("failed to build unsigned transaction: %v", err)
			return
		}

		ctx := w.NewTxContext(operations.Type(), operations.TaskID(), tx)
		w.pendingTx.Store(ctx.DigestHash(), ctx)

		defer w.pendingTx.Delete(ctx.DigestHash())

		err = w.signTx(ctx)
		if err != nil {
			log.Errorf("failed to sign transaction: %v", err)
			return
		}

		err = w.SendSignedTx(ctx)
		if err != nil {
			log.Errorf("failed to send transaction: %v", err)
			return
		}
		// updated status to pending
		receipt, err := w.WaitTxSuccess(ctx.TxHash())
		if err != nil {
			log.Errorf("failed to wait for transaction success: %v", err)
			return
		}

		if receipt.Status == 0 {
			// updated status to fail
			log.Errorf("failed to submit transaction for taskId: %d, txHash: %s", operations.TaskID(), ctx.TxHash().String())
		} else {
			// updated status to completed
			log.Infof("successfully submitted transaction for taskId: %d, txHash: %s", operations.TaskID(), ctx.TxHash().String())
		}
	}
}

func (w *WalletClient) receiveSubmitTaskLoop() {
	taskEvent := w.Bus().Subscribe(eventbus.EventSubmitTask{})

	go func() {
		for {
			select {
			case <-w.ctx.Done():
				log.Info("EVM wallet receive task event loop exiting")
				return

			case result := <-taskEvent:
				log.Debugf("Received task event: %v", result)

				t, ok := result.(*db.TaskResult)
				if !ok {
					log.Warn("Received task event with unexpected type")
					continue
				}

				w.submitTaskQueue.Add(t)
			}
		}
	}()
}

// removeSubmitTaskTaskLoop listens for task events and removes completed or failed tasks from the submitTaskQueue.
func (w *WalletClient) removeSubmitTaskTaskLoop() {
	taskEvent := w.Bus().Subscribe(eventbus.EventTask{})

	go func() {
		for {
			select {
			case <-w.ctx.Done():
				log.Info("evm wallet remove submit task loop done")
				return
			case data := <-taskEvent:
				switch v := data.(type) {
				case db.DetailTask:
					if v.Status() == db.Completed || v.Status() == db.Failed {
						w.submitTaskQueue.Remove(v.TaskID())
					}
				default:
					log.Warnf("unexpected type received in removeSubmitTaskTaskLoop: %T", v)
				}
			}
		}
	}()
}
