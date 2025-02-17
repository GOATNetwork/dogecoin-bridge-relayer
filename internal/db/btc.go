package db

import (
	"time"

	"gorm.io/gorm"
)

type BTCTransaction struct {
	gorm.Model
	TxID        string    `gorm:"uniqueIndex;not null" json:"tx_id"`
	RawTxData   string    `gorm:"type:text;not null"   json:"raw_tx_data"`
	ReceivedAt  time.Time `gorm:"not null"             json:"received_at"`
	Processed   bool      `gorm:"default:false"        json:"processed"`
	ProcessedAt time.Time `json:"processed_at"`
}

type BtcSyncStatus struct {
	gorm.Model
	UnconfirmHeight int64 `gorm:"not null" json:"unconfirm_height"`
	ConfirmedHeight int64 `gorm:"not null" json:"confirmed_height"`
}

type BtcBlockData struct {
	gorm.Model
	BlockHeight  uint64 `gorm:"unique;not null" json:"block_height"`
	BlockHash    string `gorm:"unique;not null" json:"block_hash"`
	Header       []byte `json:"header"`
	Difficulty   uint32 `json:"difficulty"`
	RandomNumber uint32 `json:"random_number"`
	MerkleRoot   string `json:"merkle_root"`
	BlockTime    int64  `json:"block_time"`
	TxHashes     string `json:"tx_hashes"`
}

type BtcTXOutput struct {
	gorm.Model
	BlockID  uint   `json:"block_data_id"`
	TxHash   string `json:"tx_hash"`
	Value    uint64 `json:"value"`
	PkScript []byte `json:"pk_script"`
}

// BtcBlock model.
type BtcBlock struct {
	ID        uint      `gorm:"primaryKey"                            json:"id"`
	Height    uint64    `gorm:"not null;uniqueIndex"                  json:"height"`
	Hash      string    `gorm:"not null"                              json:"hash"`
	Status    string    `gorm:"not null;index:btc_block_status_index" json:"status"` // "unconfirm", "confirmed", "signing", "pending", "processed"
	UpdatedAt time.Time `gorm:"not null"                              json:"updated_at"`
}

// UTXO model (wallet UTXO).
type UTXO struct {
	gorm.Model
	Txid         string `gorm:"not null;index:unique_txid_out_index,unique" json:"txid"`
	OutIndex     int    `gorm:"not null;index:unique_txid_out_index,unique" json:"out_index"`
	PkScript     []byte `json:"pk_script"`
	SubScript    []byte `json:"sub_script"`                                                       // P2WSH Type
	Amount       int64  `gorm:"not null;index:utxo_amount_index"            json:"amount"`        // BTC precision up to 8 decimal places
	Receiver     string `gorm:"not null;index:utxo_receiver_index"          json:"receiver"`      // utxo owner(to): it is MPC p2wpkh address here, or p2wsh (need collect)
	Sender       string `gorm:"not null"                                    json:"sender"`        // from
	Source       string `gorm:"not null"                                    json:"source"`        // "deposit", "unknown"
	ReceiverType string `gorm:"not null"                                    json:"receiver_type"` // P2PKH P2SH P2WSH P2WPKH P2TR
	ReceiveBlock uint64 `gorm:"not null"                                    json:"receive_block"` // receive at BTC block height
	SpentBlock   uint64 `gorm:"not null"                                    json:"spent_block"`   // spent at BTC block height
	Status       int    `gorm:"not null"                                    json:"status"`        // 0:"unconfirm", 1:"confirmed", 2:"pending (spend out)", 3:"spent"
}

// BtcDepositResult model, it save deposit data from layer2 events.
type BtcDepositResult struct {
	ID                 uint   `gorm:"primaryKey"                                             json:"id"`
	Txid               string `gorm:"uniqueIndex:idx_txid_tx_out"                            json:"txid"`
	TxOut              uint64 `gorm:"uniqueIndex:idx_txid_tx_out"                            json:"tx_out"`
	Address            string `gorm:"not null"                                               json:"address"`
	Amount             uint64 `gorm:"not null"                                               json:"amount"`
	BlockHash          string `gorm:"not null"                                               json:"block_hash"`
	NeedFetchSubScript bool   `gorm:"not null;default:false;index:idx_need_fetch_sub_script" json:"need_fetch_sub_script"` // if true, need fetch sub script from BTC client, or fetch not exist utxo then save
}

// BtcWithdraw model (for managing withdrawals).
type BtcWithdraw struct {
	ID        uint      `gorm:"primaryKey"                            json:"id"`
	RequestId uint64    `gorm:"not null;uniqueIndex"                  json:"request_id"`
	GoatBlock uint64    `gorm:"not null"                              json:"goat_block"` // Goat block height
	Amount    uint64    `gorm:"not null;index:withdraw_amount_index"  json:"amount"`     // withdraw BTC satoshis, build tx out should minus tx fee
	TxPrice   uint64    `gorm:"not null;index:withdraw_txprice_index" json:"tx_price"`   // Unit is satoshis
	TxFee     uint64    `gorm:"not null"                              json:"tx_fee"`     // will update when aggregating build
	From      string    `gorm:"not null"                              json:"from"`
	To        string    `gorm:"not null"                              json:"to"`       // BTC address, support all 4 types
	Status    string    `gorm:"not null;index:withdraw_status_index"  json:"status"`   // "create", "aggregating", "init", "signing", "pending", "unconfirm", "confirmed", "processed", "closed" - means user cancel
	OrderId   string    `gorm:"not null;index:withdraw_orderid_index" json:"order_id"` // update when signing, it always can be query from SendOrder by BTC txid
	Txid      string    `gorm:"not null;index:withdraw_txid_index"    json:"txid"`     // update when signing
	Reason    string    `gorm:"not null"                              json:"reason"`   // reason for closed
	CreatedAt time.Time `gorm:"not null"                              json:"created_at"`
	UpdatedAt time.Time `gorm:"not null"                              json:"updated_at"`
}

// SendOrder model (should send withdraw, vin, vout via off-chain consensus).
type SendOrder struct {
	ID           uint      `gorm:"primaryKey"                               json:"id"`
	OrderId      string    `gorm:"not null;uniqueIndex"                     json:"order_id"`
	Proposer     string    `gorm:"not null"                                 json:"proposer"`
	Pid          uint64    `gorm:"not null"                                 json:"pid"`
	Amount       uint64    `gorm:"not null"                                 json:"amount"` // BTC precision up to 8 decimal places
	TxPrice      uint64    `gorm:"not null;index:sendorder_txprice_index"   json:"tx_price"`
	Status       string    `gorm:"not null;index:sendorder_status_index"    json:"status"`     // "aggregating", "init", "signing", "pending", "rbf-request", "unconfirm", "confirmed", "processed", "closed" - means not in use, should rollback withdraw, vin, vout
	OrderType    string    `gorm:"not null;index:sendorder_ordertype_index" json:"order_type"` // "withdrawal", "consolidation"
	BtcBlock     uint64    `gorm:"not null"                                 json:"btc_block"`  // BTC block height
	Txid         string    `gorm:"not null;index:sendorder_txid_index"      json:"txid"`       // txid will update after signing status
	NoWitnessTx  []byte    `json:"no_witness_tx"`                                              // no witness tx after tx build
	TxFee        uint64    `gorm:"not null"                                 json:"tx_fee"`     // the real tx fee will update after tx built
	ExternalTxId string    `json:"external_tx_id"`                                             // fireblocks will return its special transaction id
	UpdatedAt    time.Time `gorm:"not null"                                 json:"updated_at"`
}

// Vin model (sent transaction input).
type Vin struct {
	ID           uint      `gorm:"primaryKey"                       json:"id"`
	OrderId      string    `gorm:"not null;index:vin_orderid_index" json:"order_id"`
	BtcHeight    uint64    `gorm:"not null"                         json:"btc_height"`
	Txid         string    `gorm:"not null;index:vin_txid_index"    json:"txid"`
	OutIndex     int       `gorm:"not null;index:vin_out_index"     json:"out_index"`
	SigScript    []byte    `json:"sig_script"`
	SubScript    []byte    `json:"sub_script"` // P2WSH Type
	Sender       string    `json:"sender"`
	ReceiverType string    `gorm:"not null"                         json:"receiver_type"` // P2PKH P2SH P2WSH P2WPKH P2TR
	Source       string    `gorm:"not null"                         json:"source"`        // "withdraw", "unknown"
	Status       string    `gorm:"not null;index:vin_status_index"  json:"status"`        // "aggregating", "init", "signing", "pending", "unconfirm", "confirmed", "processed", "closed"
	UpdatedAt    time.Time `gorm:"not null"                         json:"updated_at"`
}

// Vout model (sent transaction output).
type Vout struct {
	ID         uint      `gorm:"primaryKey"                         json:"id"`
	OrderId    string    `gorm:"not null;index:vout_orderid_index"  json:"order_id"`
	BtcHeight  uint64    `gorm:"not null"                           json:"btc_height"`
	Txid       string    `gorm:"not null;index:vout_txid_out_index" json:"txid"`
	OutIndex   int       `gorm:"not null;index:vout_txid_out_index" json:"out_index"`
	WithdrawId string    `json:"withdraw_id"`                                        // EvmTxId
	Amount     int64     `gorm:"not null"                           json:"amount"`   // BTC precision up to 8 decimal places
	Receiver   string    `gorm:"not null"                           json:"receiver"` // withdraw To
	PkScript   []byte    `json:"pk_script"`
	Sender     string    `json:"sender"`                                           // MPC address
	Source     string    `gorm:"not null"                           json:"source"` // "withdraw", "unknown"
	Status     string    `gorm:"not null;index:vout_status_index"   json:"status"` // "aggregating", "init", "signing", "pending", "unconfirm", "confirmed", "processed", "closed"
	UpdatedAt  time.Time `gorm:"not null"                           json:"updated_at"`
}

// BtcDeposit model (for managing deposits).
type BtcDeposit struct {
	ID          uint      `gorm:"primaryKey"                                 json:"id"`
	TxHash      string    `gorm:"not null;index:deposit_txhash_output_index" json:"tx_hash"`
	Amount      int64     `gorm:"not null;default:0"                         json:"amount"`
	RawTx       string    `gorm:"not null"                                   json:"raw_tx"`
	EvmAddr     string    `gorm:"not null"                                   json:"evm_addr"`
	BlockHash   string    `gorm:"not null;index:deposit_blockhash_index"     json:"block_hash"`
	BlockHeight uint64    `gorm:"not null;index:deposit_blockhash_height"    json:"block_height"`
	TxIndex     int       `gorm:"not null;index:deposit_txindex_index"       json:"tx_index"`
	OutputIndex int       `gorm:"not null;index:deposit_txhash_output_index" json:"output_index"`
	MerkleRoot  []byte    `json:"merkle_root"`
	Proof       []byte    `json:"proof"`
	SignVersion uint32    `gorm:"not null"                                   json:"sign_version"`
	Status      string    `gorm:"not null;index:deposit_status_index"        json:"status"` // "unconfirm", "confirmed", "signing", "pending", "processed"
	CreatedAt   time.Time `gorm:"index:deposit_created_index"                json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null"                                   json:"updated_at"`
}
