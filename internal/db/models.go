package db

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	vtypes "github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type LogIndex struct {
	gorm.Model
	Address     common.Address `gorm:"index;size:160"                json:"address"`
	EventName   string         `json:"eventName"`                                         // event name
	Log         *types.Log     `gorm:"serializer:json"               json:"log"`          // event content
	TxHash      common.Hash    `gorm:"index;size:256"                json:"tx_hash"`      // tx hash
	ChainId     uint64         `gorm:"index:log_index_unique,unique" json:"chain_id"`     // chainId
	BlockNumber uint64         `gorm:"index:log_index_unique,unique" json:"block_number"` // block number of the tx
	LogIndex    uint64         `gorm:"index:log_index_unique,unique" json:"log_index"`    // block log index
	ForeignID   uint           `gorm:"index"                         json:"foreign_id"`   // task table ID;submitter table ID;participant_event table ID;...
}

func (LogIndex) TableName() string {
	return "log_index"
}

// Account save all accounts.
type Account struct {
	gorm.Model
	UserAddress common.Address `gorm:"index;not null"        json:"user_address"`
	Account     uint64         `gorm:"index;not null"        json:"account"`
	Chain       uint8          `gorm:"not null"              json:"chain"`
	Index       uint32         `gorm:"not null"              json:"index"`
	Address     string         `gorm:"uniqueIndex; not null" json:"address"`
	LogIndex    LogIndex       `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (Account) TableName() string {
	return "account"
}

type DepositRecord struct {
	gorm.Model
	UserAddress    common.Address  `gorm:"index;not null" json:"user_address"`
	DepositAddress string          `gorm:"not null"       json:"deposit_address"`
	Amount         decimal.Decimal `gorm:"not null"       json:"amount"`
	ChainId        uint64          `gorm:"not null"       json:"chain_id"`
	TxHash         string
	BlockHeight    uint64
	LogTxIndex     uint64
	LogIndex       LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (DepositRecord) TableName() string {
	return "deposit_record"
}

type WithdrawalRecord struct {
	gorm.Model
	ChainId        uint64          `gorm:"not null"             json:"chain_id"`
	UserAddress    common.Address  `gorm:"index;not null"       json:"user_address"`
	DepositAddress string          `gorm:"index;not null"       json:"deposit_address"`
	ToAddress      string          `gorm:"index;not null"       json:"to_address"`
	Amount         decimal.Decimal `gorm:"not null"             json:"amount"`
	TxHash         string          `json:"tx_hash"`
	LogIndex       LogIndex        `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (WithdrawalRecord) TableName() string {
	return "withdrawal_record"
}

type AddressBalance struct {
	gorm.Model
	ChainId uint64          `gorm:"not null"                            json:"chain_id"`
	Address string          `gorm:"uniqueIndex:address_token; not null" json:"address"`
	Token   string          `gorm:"uniqueIndex:address_token; not null" json:"token"`
	Amount  decimal.Decimal `gorm:"not null"                            json:"amount"`
}

func (AddressBalance) TableName() string {
	return "address_balance"
}

type InscriptionMintb struct {
	gorm.Model
	Recipient string          `gorm:"not null"             json:"recipient"`
	Ticker    vtypes.Byte32   `gorm:"not null"             json:"ticker"`
	Amount    decimal.Decimal `gorm:"not null"             json:"amount"`
	LogIndex  LogIndex        `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (InscriptionMintb) TableName() string {
	return "inscription_mintb"
}

type InscriptionBurnb struct {
	gorm.Model
	From     string          `gorm:"not null"             json:"from"`
	Ticker   vtypes.Byte32   `gorm:"not null"             json:"ticker"`
	Amount   decimal.Decimal `gorm:"not null"             json:"amount"`
	LogIndex LogIndex        `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (InscriptionBurnb) TableName() string {
	return "inscription_burnb"
}

type Asset struct {
	gorm.Model
	Ticker            vtypes.Byte32 `gorm:"uniqueIndex;not null" json:"ticker"`
	AssetType         uint8         `gorm:"not null"             json:"asset_type"`
	Decimals          uint8         `gorm:"not null"             json:"decimals"`
	DepositEnabled    bool          `gorm:"not null"             json:"deposit_enabled"`
	WithdrawalEnabled bool          `gorm:"not null"             json:"withdrawal_enabled"`
	MinDepositAmount  uint64        `gorm:"not null"             json:"min_deposit_amount"`
	MinWithdrawAmount uint64        `gorm:"not null"             json:"min_withdraw_amount"`
	AssetAlias        string        `gorm:"not null"             json:"asset_alias"`
	AssetLogo         string        `gorm:"not null"             json:"asset_logo"`
}

func (Asset) TableName() string {
	return "asset"
}

type TokenInfo struct {
	gorm.Model
	ChainId         uint64          `gorm:"index:chain_id_ticker_unique,unique" json:"chain_id"`
	Ticker          vtypes.Byte32   `gorm:"index:chain_id_ticker_unique,unique" json:"ticker"`
	IsActive        bool            `gorm:"not null"                            json:"is_active"`
	AssetType       uint8           `gorm:"not null"                            json:"asset_type"`
	Decimals        uint8           `gorm:"not null"                            json:"decimals"`
	ContractAddress string          `gorm:"index;not null"                      json:"contract_address"`
	Symbol          string          `gorm:"not null"                            json:"symbol"`
	WithdrawFee     decimal.Decimal `gorm:"not null"                            json:"withdraw_fee"`
}

func (TokenInfo) TableName() string {
	return "token_info"
}
