package tss

import (
	"math/big"
	"strings"

	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

type SignerContext struct {
	chainType          uint8
	localData          LocalPartySaveData
	keyDerivationDelta *big.Int
}

func (c *SignerContext) Address() string {
	return strings.ToLower(c.localData.Address(c.ChainType()))
}

func (c *SignerContext) ChainType() uint8 {
	return c.chainType
}

func (c *SignerContext) CurveType() crypto.CurveType {
	return types.GetCurveTypeByChain(c.ChainType())
}

func (c *SignerContext) KeyDerivationDelta() *big.Int {
	return c.keyDerivationDelta
}

func (c *SignerContext) LocalData() LocalPartySaveData {
	return c.localData
}

func (c *SignerContext) IsTssSinger() bool {
	return c.KeyDerivationDelta() == nil && c.ChainType() == uint8(types.ChainEthereum)
}
