package state

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EvmWalletState struct {
	pool *gorm.DB
}

func NewEvmWalletState(pool *gorm.DB) *EvmWalletState {
	return &EvmWalletState{pool: pool}
}

func (d *EvmWalletState) tx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		tx = d.pool
	}

	return tx
}

func (d *EvmWalletState) CreateTx(tx *gorm.DB,
	ty int,
	seqID uint64,
	from common.Address,
	unSign *types.Transaction,
) error {
	tx = d.tx(tx)
	signer := types.LatestSignerForChainID(unSign.ChainId())
	signatureHash := signer.Hash(unSign)

	txJsonData, err := json.Marshal(unSign)
	if err != nil {
		return err
	}

	dbTx := &db.EvmTransaction{
		ChainId:    unSign.ChainId().Uint64(),
		SignHash:   signatureHash,
		TxJsonData: txJsonData,
		TxNonce:    decimal.NewFromUint64(unSign.Nonce()),
		Sender:     from,
		Status:     db.Created,
		Error:      "",
		Type:       ty,
		SeqID:      seqID,
	}

	return tx.Create(dbTx).Error
}

func (d *EvmWalletState) UpdateTransaction(tx *gorm.DB, signTx *types.Transaction) error {
	tx = d.tx(tx)
	signer := types.LatestSignerForChainID(signTx.ChainId())
	signatureHash := signer.Hash(signTx)

	txJsonData, err := json.Marshal(signTx)
	if err != nil {
		return err
	}

	//tx.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "chain_id"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"block_number"}),
	//}).Create(&db.BlockScanned{
	//	ChainId:     chainId,
	//	BlockNumber: blockNumber,
	//}).Error

	return tx.Model(&db.EvmTransaction{}).
		Where("sign_hash = ?", signatureHash).
		Updates(&map[string]interface{}{
			"tx_hash":      signTx.Hash(),
			"tx_json_data": txJsonData,
			"status":       db.Pending,
			"error":        "",
		}).Error
}

func (d *EvmWalletState) PendingBlockchainTransaction(tx *gorm.DB, txHash common.Hash) (*db.EvmTransaction, error) {
	tx = d.tx(tx)
	bt := &db.EvmTransaction{}

	err := tx.
		Model(bt).
		Where("tx_hash = ? AND status = ?", txHash, db.Pending).
		Last(&bt).
		Error

	return bt, err
}

func (d *EvmWalletState) PendingBlockchainTransactions(tx *gorm.DB, chainId uint64) (txs []db.EvmTransaction, err error) {
	tx = d.tx(tx)

	err = tx.
		Model(&db.EvmTransaction{}).
		Where("chain_id = ?  and status = ?", chainId, db.Pending).
		Find(&txs).
		Error

	return txs, err
}

func (d *EvmWalletState) LatestNonce(tx *gorm.DB, chainId uint64, account common.Address) (decimal.Decimal, error) {
	tx = d.tx(tx)
	bt := &db.EvmTransaction{}
	result := tx.
		Model(bt).
		Where("chain_id = ? AND sender = ? AND status IN ?", chainId, account, []int{db.Created, db.Pending, db.Completed}).
		Last(&bt)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return decimal.NewFromInt(-1), nil
	}

	return bt.TxNonce, result.Error
}

func (d *EvmWalletState) UpdateTx(tx *gorm.DB, txHash common.Hash, status int, err error) error {
	db := d.tx(tx).
		Model(&db.EvmTransaction{}).
		Where("tx_hash = ? AND status != ?", txHash, db.Completed)
	if err != nil {
		db = db.Updates(&map[string]interface{}{
			"status": status,
			"error":  err.Error(),
		})
	} else {
		db = db.Updates(&map[string]interface{}{
			"status": status,
		})
	}

	return db.Error
}

func (d *EvmWalletState) UpdateFailTx(txHash common.Hash, err error) error {
	return d.UpdateTx(nil, txHash, db.Failed, err)
}

func (d *EvmWalletState) UpdatePendingTx(txHash common.Hash) error {
	return d.UpdateTx(nil, txHash, db.Pending, nil)
}

func (d *EvmWalletState) UpdateBookedTx(txHash common.Hash) error {
	return d.UpdateTx(
		nil,
		txHash,
		db.Completed,
		nil,
	)
}

func (d *EvmWalletState) SaveOperations(tx *gorm.DB, op *db.Operations) error {
	tx = d.tx(tx)

	err := tx.Save(op).Error // todo
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return fmt.Errorf("operation nonce:%d already exists: %w", op.TssNonce, err)
	}

	return err
}

func (d *EvmWalletState) FindOperationByNonce(tx *gorm.DB, tssNonce uint64) (*db.Operations, error) {
	tx = d.tx(tx)

	var operation db.Operations

	err := tx.Model(&db.Operations{}).Where("tss_nonce = ?", tssNonce).First(&operation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &operation, err
}
