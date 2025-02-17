package db

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EvmTransaction struct {
	gorm.Model
	ChainId    uint64          `gorm:"index"                json:"chain_id"`
	TxHash     common.Hash     `gorm:"uniqueIndex;size:256" json:"tx_hash"`   // tx hash
	SignHash   common.Hash     `gorm:"uniqueIndex;size:256" json:"sign_hash"` // sign hash
	TxJsonData []byte          `json:"tx"`                                    // blockchain origin tx of json format
	TxNonce    decimal.Decimal `gorm:"index:sender_nonce"   json:"tx_nonce"`  // tx nonce
	Sender     common.Address  `json:"sender"`
	Status     int             `json:"status"` // 0: new，1:booked
	Error      string          `json:"error"`
	Type       int             // operation、withdraw、consolidation、transfer
	SeqID      uint64
}

func (*EvmTransaction) TableName() string {
	return "evm_transactions"
}

func (e *EvmTransaction) Tx() *types.Transaction {
	tx := new(types.Transaction)
	err := json.Unmarshal(e.TxJsonData, tx)
	utils.Assert(err)

	return tx
}
