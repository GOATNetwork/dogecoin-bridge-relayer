package address_test

import (
	"testing"

	_ "github.com/nuvosphere/nudex-voter/internal/wallet/evm"
)

func TestBip44GenerateAddress(t *testing.T) {
	// localData := testutil.ReadTestKey(1)
	// t.Log("master address: ", ethcrypto.PubkeyToAddress(*localData.ECDSAPub.ToECDSAPubKey()))
	// address := address.GenerateAddressByPath(localData.ECDSAPub, types.CoinTypeEVM, 0, 0)
	// t.Log("address: ", address)
	// assert.Equal(t, strings.ToLower("0xf1cbea0b78f0083530056b88c4cea93e5ff3b5a7"), strings.ToLower(address))
}
