package evm

import (
	tssCrypto "github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/ethereum/go-ethereum/crypto"
	ty "github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
)

func init() {
	address.RegisterAddress(ty.CoinTypeEVM, EthAddress)
}

func EthAddress(p *tssCrypto.ECPoint) string {
	return crypto.PubkeyToAddress(*p.ToECDSAPubKey()).String()
}
