package tss

import (
	"context"
	"encoding/json"
	"math/big"
	"slices"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts/codec"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	log "github.com/sirupsen/logrus"
)

type P2PMocker struct {
	typeBindEvent *sync.Map // MessageType:eventbus.Event
	bus           eventbus.Bus
	nodeCount     int
	partners      types.Participants
}

func NewP2PMocker(bus eventbus.Bus) *P2PMocker {
	return &P2PMocker{
		bus:           bus,
		typeBindEvent: &sync.Map{},
	}
}

func (p *P2PMocker) Bind(msgType p2p.MessageType, event eventbus.Event) {
	p.typeBindEvent.Store(msgType, event)
}

func (p *P2PMocker) UpdateParticipants(partners types.Participants) {
	p.partners = partners
}

func (p *P2PMocker) PublishMessage(ctx context.Context, msg any) error {
	data, err := json.Marshal(msg)
	utils.Assert(err)

	var receivedMsg p2p.Message[json.RawMessage]
	err = json.Unmarshal(data, &receivedMsg)
	utils.Assert(err)

	event, ok := p.typeBindEvent.Load(receivedMsg.MessageType)
	if ok {
		p.bus.Publish(event, receivedMsg)
		log.Infof("receivedMsg.MessageType: %v", receivedMsg.MessageType)
	} else {
		log.Errorf("receivedMsg.MessageType error: %v", receivedMsg.MessageType)
	}

	return nil
}

func (p *P2PMocker) SetOnlinePeerCount(nodeCount int) {
	p.nodeCount = nodeCount
}

func (p *P2PMocker) OnlinePeerCount() int {
	return p.nodeCount
}

func (p *P2PMocker) IsOnline(submitter string) bool {
	return true
}

type VoterContractMocker struct {
	nonce        big.Int
	proposer     common.Address
	participants atomic.Value // types.Participants
}

func (v *VoterContractMocker) SubmitTask(taskResult *db.TaskResult) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) PostTask(task any) {
	// TODO implement me
	panic("implement me")
}

func NewVoterContractMocker() *VoterContractMocker {
	return &VoterContractMocker{
		nonce:        *big.NewInt(1),
		participants: atomic.Value{},
	}
}

func (v *VoterContractMocker) IsSyncing() bool {
	return false
}

func (v *VoterContractMocker) TssNonce() (*big.Int, error) {
	return &v.nonce, nil
}

func (v *VoterContractMocker) SetProposer(p common.Address) {
	v.proposer = p
}

func (v *VoterContractMocker) Proposer() (common.Address, error) {
	return v.proposer, nil
}

func (v *VoterContractMocker) NextSubmitter() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) TssSigner() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) LastSubmissionTime() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) ForcedRotationWindow() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) TaskCompletionThreshold() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) EncodeVerifyAndCall(operations []codec.Operation, signature []byte) []byte {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) GenerateVerifyTaskUnSignMsg(operations []codec.Operation, nonce *big.Int) (common.Hash, common.Hash) {
	encodeData := contracts.EncodeOperation(operations, nonce, big.NewInt(31337))

	dataHash := crypto.Keccak256Hash(encodeData)
	hash := utils.PersonalMsgHash(dataHash)

	return dataHash, hash
}

func (v *VoterContractMocker) SetParticipants(pp types.Participants) {
	v.participants.Store(pp)
}

func (v *VoterContractMocker) Participants() (types.Participants, error) {
	return v.participants.Load().(types.Participants), nil
}

func (v *VoterContractMocker) IsParticipant(participant common.Address) (bool, error) {
	participants := v.participants.Load().(types.Participants)
	return slices.Contains(participants, participant), nil
}

func (v *VoterContractMocker) GetRandomParticipant(_salt common.Address) (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) GetLatestTask() (contracts.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) IsTaskCompleted(taskId uint64) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) GetTaskState(taskId uint64) (uint8, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) GetUncompletedTasks() ([]contracts.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) EncodeSubmitTask(submitter common.Address, context []byte) []byte {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) NextTaskId() (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) Tasks(taskId uint64) (contracts.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) TaskSubmitter() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) EncodeMarkTaskCompleted(taskId *big.Int, result []byte) []byte {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) EncodeRegisterNewAddress(_account uint32, _chain uint8, _index uint32, _address string) []byte {
	return contracts.EncodeFun(contracts.AccountManagerContractABI, "registerNewAddress", _account, _chain, _index, _address)
}

func (v *VoterContractMocker) GetAddressRecord(_account uint32, _chain uint8, _index uint32) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) EncodeRecordDeposit(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) []byte {
	return contracts.EncodeFun(contracts.DepositManagerContractMetaData.ABI, "recordWithdrawal", _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}

func (v *VoterContractMocker) EncodeRecordWithdrawal(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) []byte {
	return contracts.EncodeFun(contracts.DepositManagerContractMetaData.ABI, "recordWithdrawal", _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}
