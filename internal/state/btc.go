package state

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"gorm.io/gorm"
)

type BtcWalletState struct {
	pool    *gorm.DB
	chainId uint64
}

func NewBtcWalletState(pool *gorm.DB, chainId uint64) *BtcWalletState {
	return &BtcWalletState{pool: pool, chainId: chainId}
}

func (d *BtcWalletState) TX(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		tx = d.pool
	}

	return tx
}

func (d *BtcWalletState) AddNewUTXO(txId string, index uint32, amount int64, from, to string, receiveHeight uint64) error {
	return d.pool.Create(&db.UTXO{
		Txid:         txId,
		OutIndex:     int(index),
		PkScript:     nil,
		SubScript:    nil,
		Amount:       amount,
		Receiver:     to,
		Sender:       from,
		Source:       "",
		ReceiverType: "P2WPKH",
		ReceiveBlock: receiveHeight,
		SpentBlock:   0,
		Status:       0,
	}).Error
}

func (d *BtcWalletState) SpentUTXO(txId string, index uint32, spentHeight uint64) error {
	return d.pool.
		Model(&db.UTXO{}).
		Where("txid = ? AND out_index = ? AND status = 1", txId, index).
		Updates(&db.UTXO{
			SpentBlock: spentHeight,
			Status:     3,
		}).
		Error
}

func (d *BtcWalletState) GetUTXO(txId string, index uint32) (*db.UTXO, error) {
	utxo := &db.UTXO{}

	return utxo, d.pool.
		Model(&db.UTXO{}).
		Where("txid = ? AND out_index = ?", txId, index).
		Last(utxo).
		Error
}

func (d *BtcWalletState) IsExistUTXO(txId string, index uint32) bool {
	_, err := d.GetUTXO(txId, index)
	return err == nil
}

func (d *BtcWalletState) GetUTXOByOwner(owner string) ([]db.UTXO, error) {
	var utxos []db.UTXO
	return utxos, d.pool.Model(&db.UTXO{}).Where("owner = ?", owner).Find(&utxos).Error
}
