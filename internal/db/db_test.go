package db

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

func TestTask(t *testing.T) {
	utils.SkipCI(t)

	config.AppConfig.DB.DbRootDir = "./"
	dbm := NewDatabaseManager()
	dbm.initDB()

	var taskId uint64 = 13
	task := CreateWalletTask{
		TaskId:  taskId,
		Account: 0,
		Chain:   0,
		Index:   0,
	}
	db := dbm.GetContractDB().Debug()
	db.DryRun = true
	// err := db.Create(&task).Error
	// err := db.Save(&task).Error
	// assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))
	err := db.Model(&CreateWalletTask{}).Where("task_id", taskId).Last(&task).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))

	task = CreateWalletTask{}
	// db.DryRun = tr
	err = db.Model(&task).Preload(clause.Associations).Where("id", taskId).Last(&task).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))

	baseTask := Task{}
	err = db.Model(&Task{}).Preload(clause.Associations).Where("id", taskId).Last(&baseTask).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(baseTask))
	// err = db.Preload(clause.Associations).Where("task_id", taskId).Last(&baseTask).Error
	// assert.Nil(t, err)
	// t.Log(utils.FormatJSON(baseTask))
	t.Log("end")
}

func TestUniqueTask(t *testing.T) {
	utils.SkipCI(t)

	config.AppConfig.DB.DbRootDir = "./"
	dbm := NewDatabaseManager()
	dbm.initDB()

	var taskId uint64 = 26
	task := CreateWalletTask{
		TaskId:      taskId,
		UserAddress: common.Address{},
		Account:     0,
		Chain:       0,
		Index:       0,
	}
	db := dbm.GetContractDB().Debug()
	db.DryRun = true
	err := db.Save(&task).Error
	// err := db.Save(&task).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))
	task = CreateWalletTask{}
	err = db.Model(&CreateWalletTask{}).Where("task_id", 25).Last(&task).Error
	assert.Nil(t, err)
	t.Log(utils.FormatJSON(task))
	t.Log("end")
}

func TestSubmitterChosenUniqueTask(t *testing.T) {
	utils.SkipCI(t)

	config.AppConfig.DB.DbRootDir = "./"
	dbm := NewDatabaseManager()
	dbm.initDB()

	s := SubmitterChosen{
		BlockNumber: 12,
		Submitter:   "122",
	}
	db := dbm.GetContractDB().Debug()
	db.DryRun = true
	err := db.Clauses(clause.OnConflict{UpdateAll: true}).Save(&s).Error
	assert.Nil(t, err)
	t.Log("end")
}
