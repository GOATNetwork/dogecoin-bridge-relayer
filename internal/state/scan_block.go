package state

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetScannedNumber(d *gorm.DB, chainId, startBlock uint64) (uint64, error) {
	bs := db.BlockScanned{}

	err := d.
		Where("chain_id = ?", chainId).
		Attrs(db.BlockScanned{BlockNumber: startBlock}).
		FirstOrCreate(&bs).
		Error

	return bs.BlockNumber, err
}

func UpdateLatestScannedHeight(d *gorm.DB, chainId, blockNumber uint64) error {
	// Update columns to new value on `chain_id` conflict
	return d.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chain_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"block_number"}),
	}).Create(&db.BlockScanned{
		ChainId:     chainId,
		BlockNumber: blockNumber,
	}).Error
}
