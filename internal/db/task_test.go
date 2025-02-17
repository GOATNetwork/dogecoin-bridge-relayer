package db

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
	assert.True(t, utils.IsZero(0))

	values := []any{
		WithdrawalTask{},
		DepositTask{},
		Deposit{},
		DepositRecord{},
		Withdrawal{},
		WithdrawalTask{},
		WithdrawalRequest{},
		common.Hash{},
		common.Address{},
	}

	for _, value := range values {
		assert.True(t, utils.IsZero(value))
	}

	value := DepositRequest{
		UserAddress:    common.Address{},
		ChainId:        1,
		Ticker:         types.Byte32{},
		DepositAddress: "",
		Amount:         nil,
		TxHash:         "",
		BlockHeight:    nil,
		TxIndex:        nil,
	}

	assert.True(t, !utils.IsZero(value))
}
