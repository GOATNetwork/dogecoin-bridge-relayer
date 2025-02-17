package layer2

import (
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
)

type AccountManager interface {
	EncodeRegisterNewAddress(_account uint32, _chain uint8, _index uint32, _address string) []byte
	GetAddressRecord(_account uint32, _chain uint8, _index uint32) (string, error) // address
}

func (l *Layer2Listener) EncodeRegisterNewAddress(_account uint32, _chain uint8, _index uint32, _address string) []byte {
	return contracts.EncodeFun(contracts.AccountManagerContractMetaData.ABI, "registerNewAddress", _account, _chain, _index, _address)
}

func (l *Layer2Listener) GetAddressRecord(_account uint32, _chain uint8, _index uint32) (string, error) {
	return l.accountManager.GetAddressRecord(nil, _account, _chain, _index)
}
