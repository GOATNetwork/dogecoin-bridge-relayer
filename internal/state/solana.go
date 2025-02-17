package state

import "gorm.io/gorm"

type SolWalletState struct {
	pool *gorm.DB
}

func NewSolWalletState(pool *gorm.DB) *SolWalletState {
	return &SolWalletState{pool: pool}
}

func (d *SolWalletState) tx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		tx = d.pool
	}

	return tx
}
