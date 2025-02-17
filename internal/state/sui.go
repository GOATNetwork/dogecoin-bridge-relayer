package state

import "gorm.io/gorm"

type SuiWalletState struct {
	pool *gorm.DB
}

func NewSuiWalletState(pool *gorm.DB) *SuiWalletState {
	return &SuiWalletState{pool: pool}
}

func (d *SuiWalletState) tx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		tx = d.pool
	}

	return tx
}
