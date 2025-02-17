package types

import (
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
)

// CalculateThreshold (threshold + 1) / length >= 2/3.
func CalculateThreshold(length int) int {
	threshold := decimal.NewFromUint64(uint64(length)).
		Mul(decimal.NewFromInt(2)).
		Div(decimal.NewFromInt(3)).
		Ceil().
		IntPart()

	return int(threshold - 1)
}

type Participants []common.Address

func (p Participants) Threshold() int {
	if p.Len() <= 1 {
		return 1
	}

	return CalculateThreshold(len(p))
}

func (p Participants) Len() int {
	return len(p)
}

func (p Participants) GroupID() common.Hash {
	slices.SortStableFunc(p, func(a, b common.Address) int { return a.Big().Cmp(b.Big()) })

	var data []byte
	for _, a := range p {
		data = append(data, a.Bytes()...)
	}

	return crypto.Keccak256Hash(data)
}

func (p Participants) Contains(address common.Address) bool {
	return slices.Contains(p, address)
}

type ParticipantEvent struct {
	EventName string
	Address   []common.Address
}
