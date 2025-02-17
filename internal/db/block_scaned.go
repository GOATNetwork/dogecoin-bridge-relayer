package db

import "gorm.io/gorm"

// BlockScanned the latest block scanned.
type BlockScanned struct {
	gorm.Model
	ChainId     uint64 `gorm:"index:unique"`
	BlockNumber uint64 // block number of latest scanned
}
