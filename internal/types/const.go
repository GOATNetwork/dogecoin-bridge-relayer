package types

import (
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/crypto"
)

const (
	ChainBitcoin = iota
	ChainEthereum
	ChainSolana
	ChainSui
	ChainDOG
)

func IsEvm(chainId uint64) bool {
	return chainId == ChainIdEthereum
}

// https://docs.expand.network/ids/chain-ids
const (
	ChainIdBitcoin  = 1800
	ChainIdEthereum = 1
	ChainIdSolana   = 900
	ChainIdSui      = 101
	ChainIdDog      = 1900
)

const (
	CoinTypeBTC = 0
	CoinTypeEVM = 60
	CoinTypeSOL = 501
	CoinTypeSUI = 784
	CoinTypeDOG = 1 // todo
)

const (
	AssetTypeMain = iota
	AssetTypeErc20
)

var ErrCoinType = fmt.Errorf("error coin type")

func GetCoinTypeByChain(chain uint8) int {
	switch chain {
	case ChainBitcoin:
		return CoinTypeBTC
	case ChainEthereum:
		return CoinTypeEVM
	case ChainSolana:
		return CoinTypeSOL
	case ChainSui:
		return CoinTypeSUI
	case ChainDOG:
		return CoinTypeDOG
	default:
		panic(ErrCoinType)
	}
}

func GetChainByCoinType(coinType int) uint8 {
	switch coinType {
	case CoinTypeBTC:
		return ChainBitcoin
	case CoinTypeEVM:
		return ChainEthereum
	case CoinTypeSOL:
		return ChainSolana
	case CoinTypeSUI:
		return ChainSui
	case CoinTypeDOG:
		return ChainDOG
	default:
		panic(ErrCoinType)
	}
}

func GetCurveTypeByChain(chain uint8) crypto.CurveType {
	switch chain {
	case ChainBitcoin, ChainEthereum, ChainDOG:
		return crypto.ECDSA
	case ChainSolana, ChainSui:
		return crypto.EDDSA
	default:
		panic("implement me")
	}
}

func GetCurveTypeByCoinType(coinType int) crypto.CurveType {
	return GetCurveTypeByChain(GetChainByCoinType(coinType))
}

const (
	GenKeySessionType        = "GenerateKeySession"
	ReShareGroupSessionType  = "ReShareGroupSession"
	SignOperationSessionType = "SignOperationSession"
	SignTxSessionType        = "SignTxSession"
	SignTaskSessionType      = "SignTaskSession" // only used test
)
