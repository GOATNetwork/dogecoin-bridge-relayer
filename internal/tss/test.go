package tss

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

// only test.
func (m *Scheduler) proposalLoopForTest() {
	testPendingTask := m.bus.Subscribe(eventbus.EventTestTask{})

	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("proposal loop stopping...")
				return
			case data := <-testPendingTask: // from test task
				log.Info("received task from layer2 log scan: ", data)

				switch v := data.(type) {
				case *types.ParticipantEvent: // regroup
					m.processReGroupProposal(v)

				case *db.SubmitterChosen: // charge proposer
					m.submitterChosen = v
					m.proposer.Store(common.HexToAddress(v.Submitter))

				case db.DetailTask:
					if v.Status() == db.Created {
						m.taskQueue.Add(v)

						if m.IsCanProposal() {
							log.Info("proposal task", v)
							m.processTaskProposal(v)
						}
					} else {
						log.Infof("taskID: %d completed on blockchain", v.TaskID())
						m.taskQueue.Remove(v.TaskID())
						m.AddDiscussedTask(v.TaskID())
					}
				}
			}
		}
	}()
}

func (m *Scheduler) GetTask(taskID uint64) (pool.Task[uint64], error) {
	t := m.taskQueue.Get(taskID)
	if t != nil {
		return t, nil
	}

	panic("todo")
}

// only used test.
func (m *Scheduler) joinSignTaskSession(msg SessionMessage[ProposalID, Proposal], task pool.Task[uint64]) error {
	log.Debugf("JoinSignTaskSession: session id: %v, task id:%v, task type: %v", msg.SessionID, task.TaskID(), task.Type())

	switch v := task.(type) {
	case *db.CreateWalletTask:
		log.Debugf("joinSignTaskSession: tss signer: %v", msg.Signer)

		_, unSignMsg := m.createUserAddressProposal(v)
		if unSignMsg.String() != msg.Proposal.String() {
			return fmt.Errorf("SignTaskSessionType: %w", ErrTaskSignatureMsgWrong)
		}

		m.NewSignSession(
			msg.SessionID,
			msg.SeqId,
			unSignMsg.String(),
			unSignMsg,
			m.tssSigner(),
		)
	case *db.DepositTask:
	case *db.WithdrawalTask:

	default:
		return fmt.Errorf("taskID %d: %w: %v", task.TaskID(), ErrTaskIdWrong, task.Type())
	}

	return nil
}

// only used test.
func (m *Scheduler) createUserAddressProposal(task *db.CreateWalletTask) (LocalPartySaveData, *big.Int) {
	coinType := types.GetCoinTypeByChain(task.Chain)

	ec := types.GetCurveTypeByCoinType(coinType)

	switch ec {
	case crypto.ECDSA:
		localPartySaveData := m.partyData.GetData(ec)
		userAddress := address.GenerateAddressByPath(localPartySaveData.ECPoint(), uint32(coinType), task.Account, task.Index)
		// msg := m.voter.EncodeRegisterNewAddress(task.Account, task.Chain, task.Index, strings.ToLower(userAddress))
		hash := ethCrypto.Keccak256Hash([]byte(userAddress)) // todo

		return *localPartySaveData, hash.Big()
	default:
		panic(fmt.Errorf("unknown EC type: %v", ec))
	}
}

// only used test.
func (m *Scheduler) processTaskProposal(task pool.Task[uint64]) {
	switch taskData := task.(type) {
	case *db.CreateWalletTask:
		_, unSignMsg := m.createUserAddressProposal(taskData)
		log.Debugf("processTaskProposal: tss signer: %v", m.tssSigner().Address())
		tssSigner := m.GetSigner(m.tssSigner().Address())
		log.Debugf("processTaskProposal: other tssSigner: %v", tssSigner.Address())
		m.NewSignSession(
			ZeroSessionID,
			taskData.TaskId,
			unSignMsg.String(),
			unSignMsg,
			m.tssSigner(),
		)
	case *db.DepositTask:

	case *db.WithdrawalTask:
	}
}
