package state

import (
	"testing"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestContractState(t *testing.T) {
	utils.SkipCI(t)

	// config.AppConfig.DB.DbRootDir = "../test/node1/data/db"
	config.AppConfig.DB.DbRootDir = "/Users/suyanlong/github/metis/nudex-voter/test/node1/data/db"
	dbm := db.NewDatabaseManager()

	c := dbm.GetContractDB()

	state := NewContractState(c)
	//
	tasks, err := state.GetUnCompletedTasks()
	assert.NoError(t, err)
	t.Log(utils.FormatJSON(tasks))
	// task, err := state.GetUnCompletedTask(0)
	// assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))
	// var createTask []db.Task
	// err = c.Model(&db.Task{}).Preload(clause.Associations).Find(&createTask).Error
	// assert.NoError(t, err)
	// t.Log(utils.FormatJSON(createTask))
}
