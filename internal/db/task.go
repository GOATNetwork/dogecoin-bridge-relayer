package db

import (
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const (
	TaskTypeUnknown = iota
	TaskTypeCreateWallet
	TaskTypeDeposit
	TaskTypeWithdrawal
	TaskTypeConsolidation
	TaskTypeTransfer
	TaskTypeOperations
)

const (
	Created = iota
	Pending
	Completed
	Failed
)

type WalletCreationRequest struct {
	UserAddress common.Address
	Account     uint32
	AddressType uint8 // chain type
	Index       uint32
}

type DepositRequest struct {
	UserAddress    common.Address
	ChainId        uint64
	Ticker         types.Byte32
	DepositAddress string
	Amount         *big.Int // uint256
	TxHash         string
	BlockHeight    *big.Int
	TxIndex        *big.Int
}

type WithdrawalRequest struct {
	UserAddress common.Address
	ChainId     uint64
	Ticker      types.Byte32
	ToAddress   string
	Amount      *big.Int // uint256
	TxHash      string   // todo
}

type ConsolidateRequest struct {
	FromAddress string
	Ticker      types.Byte32
	ChainId     uint64
	Amount      *big.Int // uint256
	TxHash      string   // todo
}

type TransferRequest struct {
	ChainId     uint64
	Ticker      types.Byte32
	FromAddress string
	ToAddress   string
	Amount      *big.Int // uint256
	TxHash      string   // todo
}

type DetailTask interface {
	types.ChainType
	types.ChainID
	pool.Task[uint64]
	Status() int
	SetStatus(int)
	IsWalletTask() bool
	TransactionHash() string
}

type Task struct {
	ID        uint64 `gorm:"primarykey;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Ty        int            `gorm:"omitempty"`
	ChainID   uint64         `gorm:"omitempty"`
	Context   []byte         `gorm:"omitempty"            json:"context"`
	Submitter string         `gorm:"omitempty"            json:"submitter"`
	State     int            `gorm:"omitempty;default:0"  json:"status"` // 0:Created; 1:pending; 2:Completed; 3:Failed
	Result    []byte         `gorm:"omitempty"`
	LogIndex  LogIndex       `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (t *Task) SetStatus(status int) {
	t.State = status
}

func (t *Task) ChainId() uint64 {
	return t.ChainID
}

func (*Task) TableName() string {
	return "task"
}

func (t *Task) Type() int {
	return t.Ty
}

func (t *Task) TaskID() uint64 {
	return t.ID
}

func (t *Task) DetailTask() DetailTask {
	detailTask := DecodeTaskOfEvent(t.TaskID(), t.Context, t.TransactionHash())
	detailTask.SetStatus(t.Status())

	return detailTask
}

func (t *Task) TransactionHash() string {
	if t.Status() == Pending {
		return string(t.Result)
	}

	return ""
}

func (t *Task) IsWalletTask() bool {
	return t.Ty == TaskTypeCreateWallet
}

func (t *Task) ChainType() uint8 {
	return t.DetailTask().ChainType()
}

func (t *Task) Status() int {
	return t.State
}

type CreateWalletTask struct {
	TaskId      uint64
	State       int
	UserAddress common.Address
	Account     uint32
	Chain       uint8
	Index       uint32
}

func (t *CreateWalletTask) SetStatus(status int) {
	t.State = status
}

func (t *CreateWalletTask) ChainType() uint8 {
	return t.Chain
}

func (t *CreateWalletTask) ChainId() uint64 {
	return config.AppConfig.MasterChainId
}

func (t *CreateWalletTask) Status() int {
	return t.State
}

func (t *CreateWalletTask) IsWalletTask() bool {
	return true
}

func (t *CreateWalletTask) TransactionHash() string {
	// TODO implement me
	panic("implement me")
}

func (t *CreateWalletTask) Type() int {
	return TaskTypeCreateWallet
}

func (t *CreateWalletTask) TaskID() uint64 {
	return t.TaskId
}

type DepositTask struct {
	TaskId         uint64
	State          int
	UserAddress    common.Address
	DepositAddress string
	Amount         decimal.Decimal
	Chain          uint8
	ChainID        uint64
	Ticker         types.Byte32
	TxHash         string
	TxIndex        uint64
	BlockHeight    uint64
}

func (t *DepositTask) SetStatus(status int) {
	t.State = status
}

func (t *DepositTask) ChainType() uint8 {
	return t.Chain
}

func (t *DepositTask) ChainId() uint64 {
	return t.ChainID
}

func (t *DepositTask) Status() int {
	return t.State
}

func (t *DepositTask) IsWalletTask() bool {
	return false
}

func (t *DepositTask) TransactionHash() string {
	return t.TxHash
}

func (t *DepositTask) Type() int {
	return TaskTypeDeposit
}

func (t *DepositTask) TaskID() uint64 {
	return t.TaskId
}

type WithdrawalTask struct {
	TaskId      uint64
	State       int
	UserAddress common.Address
	ToAddress   string
	Amount      decimal.Decimal
	Chain       uint8
	ChainID     uint64
	Ticker      types.Byte32
	TxHash      string
}

func (t *WithdrawalTask) SetStatus(status int) {
	t.State = status
}

func (t *WithdrawalTask) ChainType() uint8 {
	return t.Chain
}

func (t *WithdrawalTask) ChainId() uint64 {
	return t.ChainID
}

func (t *WithdrawalTask) Status() int {
	return t.State
}

func (t *WithdrawalTask) IsWalletTask() bool {
	return false
}

func (t *WithdrawalTask) TransactionHash() string {
	return t.TxHash
}

func (t *WithdrawalTask) Type() int {
	return TaskTypeWithdrawal
}

func (t *WithdrawalTask) TaskID() uint64 {
	return t.TaskId
}

type ConsolidationTask struct {
	TaskId      uint64
	State       int
	FromAddress string
	Amount      decimal.Decimal
	Chain       uint8
	ChainID     uint64
	Ticker      types.Byte32
	TxHash      string
}

func (t *ConsolidationTask) SetStatus(status int) {
	t.State = status
}

func (t *ConsolidationTask) ChainType() uint8 {
	return t.Chain
}

func (t *ConsolidationTask) ChainId() uint64 {
	return t.ChainID
}

func (t *ConsolidationTask) Status() int {
	return t.State
}

func (t *ConsolidationTask) IsWalletTask() bool {
	return false
}

func (t *ConsolidationTask) TransactionHash() string {
	return t.TxHash
}

func (t *ConsolidationTask) Type() int {
	return TaskTypeConsolidation
}

func (t *ConsolidationTask) TaskID() uint64 {
	return t.TaskId
}

type TransferTask struct {
	TaskId      uint64
	State       int
	FromAddress string
	ToAddress   string
	Amount      decimal.Decimal
	Chain       uint8 // todo
	ChainID     uint64
	Ticker      types.Byte32
	TxHash      string
}

func (t *TransferTask) SetStatus(status int) {
	t.State = status
}

func (t *TransferTask) ChainType() uint8 {
	return t.Chain
}

func (t *TransferTask) ChainId() uint64 {
	return t.ChainID
}

func (t *TransferTask) Status() int {
	return t.State
}

func (t *TransferTask) IsWalletTask() bool {
	return false
}

func (t *TransferTask) TransactionHash() string {
	return t.TxHash
}

func (t *TransferTask) Type() int {
	return TaskTypeTransfer
}

func (t *TransferTask) TaskID() uint64 {
	return t.TaskId
}

const (
	TaskVersionInitial = iota
	TaskVersionV1
)

const (
	TaskErrorCodeSuccess = iota
	TaskErrorCodePending
	TaskErrorCodeChainNotSupported
	TaskErrorCodeAssetNotSupported
	TaskErrorCodeCheckTxFailed
	TaskErrorCodeCheckInscriptionFailed
	TaskErrorCodeCheckAmountFailed
	TaskErrorCodeCheckAssetFailed
	TaskErrorCodeDepositAssetNotEnabled
	TaskErrorCodeDepositAmountTooLow
	TaskErrorCodeDepositTokenNotSupported
	TaskErrorCodeDepositTokenNotActive
	TaskErrorCodeWithdrawalAssetNotEnabled
	TaskErrorCodeWithdrawalAmountTooLow
	TaskErrorCodeWithdrawalTokenNotSupported
	TaskErrorCodeWithdrawalTokenNotActive
	TaskErrorCodeCheckWithdrawalBalanceFailed
)

type TaskUpdatedEvent struct {
	ID         uint64 `gorm:"primarykey;autoIncrement:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Submitter  string         `gorm:"not null"`
	UpdateTime int64          `gorm:"not null"`
	State      uint8          `gorm:"state"`
	Result     []byte
	LogIndex   LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

type TaskResult struct {
	TaskId    uint64
	State     uint8
	ExtraData []byte
}

func (t *TaskResult) Type() int {
	// TODO implement me
	panic("implement me")
}

func (t *TaskResult) TaskID() uint64 {
	return t.TaskId
}

type CreateWallet struct {
	TaskId      uint64
	UserAddress common.Address
	State       int
	Account     uint32
	Chain       uint8 // evm_tss btc solana sui
	Index       uint32
}

func (t *CreateWallet) SetStatus(status int) {
	t.State = status
}

func (t *CreateWallet) ChainId() uint64 {
	return config.AppConfig.MasterChainId // todo
}

func (t *CreateWallet) IsWalletTask() bool {
	return true
}

func (t *CreateWallet) ChainType() uint8 {
	return types.ChainEthereum
}

func (t *CreateWallet) Type() int {
	return TaskTypeCreateWallet
}

func (t *CreateWallet) TaskID() uint64 {
	return t.TaskId
}

func (t *CreateWallet) TransactionHash() string {
	// TODO implement me
	panic("implement me")
}

func (t *CreateWallet) Status() int {
	return Created
}

type Common struct {
	TaskId          uint64
	State           int
	Chain           uint8
	ChainID         uint64
	ContractAddress string
	Amount          decimal.Decimal
	Decimals        uint8
	Symbol          string
	TxHash          string
}

func (c *Common) SetStatus(status int) {
	c.State = status
}

func (c *Common) TransactionHash() string {
	return c.TxHash
}

func (c *Common) IsNative() bool {
	return strings.TrimSpace(c.ContractAddress) == "" // todo
}

func (c *Common) TaskID() uint64 {
	return c.TaskId
}

func (c *Common) IsWalletTask() bool {
	return false
}

func (c *Common) ChainType() uint8 {
	return c.Chain
}

func (c *Common) ChainId() uint64 {
	return c.ChainID
}

func (c *Common) Status() int {
	return c.State
}

type Deposit struct {
	UserAddress common.Address
	Common
	ToAddress   string
	BlockHeight uint64
	TxIndex     uint64
}

func (t *Deposit) Type() int {
	return TaskTypeDeposit
}

type Withdrawal struct {
	Common
	WithdrawalFee decimal.Decimal
	FromAddress   string
	ToAddress     string
}

func (t *Withdrawal) Type() int {
	return TaskTypeWithdrawal
}

type Consolidation struct {
	FromAddress string
	Common
}

func (t *Consolidation) Type() int {
	return TaskTypeConsolidation
}

type Transfer struct {
	Common
	FromAddress string
	ToAddress   string
}

func (t *Transfer) Type() int {
	return TaskTypeTransfer
}
