package db

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"gorm.io/gorm"
)

type Operations struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	TssNonce  uint64         `gorm:"primarykey"`
	DataHash  common.Hash
	Data      []byte
}

func (*Operations) TableName() string {
	return "operations"
}

func (o *Operations) TaskID() uint64 {
	return o.TssNonce
}

func (o *Operations) Type() int {
	return TaskTypeOperations
}

type TaskOperations struct {
	Operation []contracts.TaskOperation
}
