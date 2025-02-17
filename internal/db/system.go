package db

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	vtypes "github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type EVMSyncStatus struct {
	gorm.Model
	LastSyncBlock uint64 `gorm:"not null" json:"last_sync_block"`
}

func (EVMSyncStatus) TableName() string {
	return "evm_sync_status"
}

// SubmitterChosen contains block number and current submitter.
type SubmitterChosen struct {
	gorm.Model
	Submitter   string   `gorm:"index:submitter_block_number_unique,unique" json:"submitter"`
	BlockNumber uint64   `gorm:"index:submitter_block_number_unique,unique" json:"block_number"`
	LogIndex    LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (SubmitterChosen) TableName() string {
	return "submitter"
}

// Participant save all participants.
type Participant struct {
	gorm.Model
	Address string `gorm:"uniqueIndex;not null" json:"address"`
}

func (Participant) TableName() string {
	return "participant"
}

// ParticipantEvent save all participants.
type ParticipantEvent struct {
	gorm.Model
	EventName   string   `json:"eventName"` // event name
	Address     string   `gorm:"index;not null"       json:"address"`
	BlockNumber uint64   `gorm:"index;not null"       json:"block_number"`
	LogIndex    LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (ParticipantEvent) TableName() string {
	return "participant_event"
}

func (e ParticipantEvent) ParticipantEvent() vtypes.ParticipantEvent {
	return vtypes.ParticipantEvent{
		EventName: e.EventName,
		Address: lo.Map(strings.Split(e.Address, ","), func(address string, index int) common.Address {
			return common.HexToAddress(address)
		}),
	}
}
