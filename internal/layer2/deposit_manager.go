package layer2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
)

type DepositManager interface {
	EncodeRecordDeposit(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) []byte
	EncodeRecordWithdrawal(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) []byte
}

func (l *Layer2Listener) EncodeRecordDeposit(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) []byte {
	return contracts.EncodeFun(contracts.DepositManagerContractMetaData.ABI, "recordDeposit", _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}

func (l *Layer2Listener) EncodeRecordWithdrawal(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) []byte {
	return contracts.EncodeFun(contracts.DepositManagerContractMetaData.ABI, "recordWithdrawal", _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}
