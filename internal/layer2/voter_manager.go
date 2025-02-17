package layer2

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts/codec"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	log "github.com/sirupsen/logrus"
)

type ContractVotingManager interface {
	TssNonce() (*big.Int, error)
	Proposer() (common.Address, error)
	NextSubmitter() (common.Address, error)
	TssSigner() (common.Address, error)
	LastSubmissionTime() (*big.Int, error)
	EncodeVerifyAndCall(operations []codec.Operation, signature []byte) []byte
	GenerateVerifyTaskUnSignMsg(operations []codec.Operation, nonce *big.Int) (common.Hash, common.Hash)
}

func (l *Layer2Listener) TssSigner() (common.Address, error) {
	return l.contractVotingManager.TssSigner(nil)
}

func (l *Layer2Listener) LastSubmissionTime() (*big.Int, error) {
	return l.contractVotingManager.LastSubmissionTime(nil)
}

func (l *Layer2Listener) TssNonce() (*big.Int, error) {
	return l.contractVotingManager.TssNonce(nil)
}

func (l *Layer2Listener) ContractVotingManager() *contracts.VotingManagerContract {
	return l.contractVotingManager
}

func (l *Layer2Listener) Proposer() (common.Address, error) {
	return l.NextSubmitter()
}

func (l *Layer2Listener) GenerateVerifyTaskUnSignMsg(operations []codec.Operation, nonce *big.Int) (common.Hash, common.Hash) {
	chainId := l.ChainID(context.Background())

	encodeData := contracts.EncodeOperation(operations, nonce, chainId)

	dataHash := crypto.Keccak256Hash(encodeData)
	hash := utils.PersonalMsgHash(dataHash)

	return dataHash, hash
}

func (l *Layer2Listener) NextSubmitter() (common.Address, error) {
	return l.contractVotingManager.NextSubmitter(nil)
}

func (l *Layer2Listener) EncodeVerifyAndCall(operations []codec.Operation, signature []byte) []byte {
	log.Debugf("EncodeVerifyAndCall: %v", utils.FormatJSON(operations))
	return contracts.EncodeFun(contracts.VotingManagerContractMetaData.ABI, "verifyAndCall", operations, signature)
}
