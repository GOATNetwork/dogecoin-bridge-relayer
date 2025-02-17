package db

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type DatabaseManager struct {
	systemDB   *gorm.DB
	contractDB *gorm.DB
	taskDB     *gorm.DB
	tokenDB    *gorm.DB
	walletDB   map[uint64]*gorm.DB
}

func NewDatabaseManager() *DatabaseManager {
	dm := &DatabaseManager{
		walletDB: make(map[uint64]*gorm.DB),
	}
	dm.initDB()

	return dm
}

func SetConnParam(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

// todo split db by chain id.
func (dm *DatabaseManager) initDB() {
	dbDir := config.AppConfig.DB.DbRootDir
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create database directory: %v", err)
	}

	databaseConfigs := []struct {
		dbPath string
		dbRef  **gorm.DB
		dbName string
	}{
		{filepath.Join(dbDir, "system.db"), &dm.systemDB, "Database system"},
		//{filepath.Join(dbDir, "l2_sync.db"), &dm.l2SyncDB, "Database l2_sync"},
		{filepath.Join(dbDir, "contract.db"), &dm.contractDB, "Database contract"},
		{filepath.Join(dbDir, "task.db"), &dm.taskDB, "Database task"},
		{filepath.Join(dbDir, "token.db"), &dm.tokenDB, "Database token"},
		//{filepath.Join(dbDir, "btc_light.db"), &dm.btcLightDb, "Database btc_light"},
		//{filepath.Join(dbDir, "wallet.db"), &dm.walletDb, "Database wallet"},
		//{filepath.Join(dbDir, "btc_cache.db"), &dm.btcCacheDb, "Database btc_cache"},
	}

	for _, dbConfig := range databaseConfigs {
		if err := dm.connectDatabase(dbConfig.dbPath, dbConfig.dbRef, dbConfig.dbName); err != nil {
			log.Fatalf("Failed to connect to %s: %v", dbConfig.dbName, err)
		}
	}

	dm.autoMigrate()

	for _, chain := range config.AppConfig.Chains {
		var (
			walletDB *gorm.DB
			err      error
		)

		subDir := fmt.Sprintf("chain_id_%d", chain.ChainId)
		dir := filepath.Join(dbDir, subDir)

		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create database directory: %v", err)
		}

		err = dm.connectDatabase(filepath.Join(dir, "wallet.db"), &walletDB, fmt.Sprintf("%s db", subDir))
		if err != nil {
			log.Fatalf("Failed to connect to %d db error: %v", chain.ChainId, err)
		}

		// AutoMigrate
		switch chain.ChainType {
		case types.ChainBitcoin:
			err = walletDB.AutoMigrate(&UTXO{}, &Vin{}, &Vout{}, &BTCTransaction{})
		case types.ChainEthereum:
			err = walletDB.AutoMigrate(&LogIndex{}, &EvmTransaction{}, &Operations{})
		case types.ChainSolana:
			// todo
		case types.ChainSui:
			// todo
		default:
			panic(fmt.Sprintf("Unknown chain type: %d", chain.ChainType))
		}

		if err != nil {
			log.Fatalf("Failed to migrate database %d error: %v", chain.ChainId, err)
		}

		dm.walletDB[chain.ChainId] = walletDB
	}

	log.Debugf("Database init completed successfully")
}

func (dm *DatabaseManager) GetSystemDB() *gorm.DB {
	return dm.systemDB
}

func (dm *DatabaseManager) GetContractDB() *gorm.DB {
	return dm.contractDB
}

func (dm *DatabaseManager) GetTaskDB() *gorm.DB {
	return dm.taskDB
}

func (dm *DatabaseManager) GetTokenDB() *gorm.DB {
	return dm.tokenDB
}

func (dm *DatabaseManager) GetBtcLightDB() *gorm.DB {
	return nil // todo
}

func (dm *DatabaseManager) GetBtcCacheDB() *gorm.DB {
	return nil // todo
}

func (dm *DatabaseManager) GetWalletDB(chainId uint64) *gorm.DB {
	return dm.walletDB[chainId]
}

func (dm *DatabaseManager) connectDatabase(dbPath string, dbRef **gorm.DB, dbName string) error {
	// open database and set WAL mode
	db, err := gorm.Open(sqlite.Open(dbPath+"?_journal_mode=WAL"), &gorm.Config{
		Logger:         gormlogger.Default.LogMode(gormlogger.Info),
		TranslateError: true, // https://gorm.golang.ac.cn/docs/error_handling.html
	})
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", dbName, err)
	}

	SetConnParam(db)
	*dbRef = db

	log.Debugf("%s connected successfully in WAL mode, path: %s", dbName, dbPath)

	return nil
}

func (dm *DatabaseManager) autoMigrate() {
	if err := dm.systemDB.AutoMigrate(
		&LogIndex{},
		&SubmitterChosen{},
		&Participant{},
		&ParticipantEvent{},
	); err != nil {
		log.Fatalf("Failed to migrate system database: %v", err)
	}

	if err := dm.contractDB.AutoMigrate(
		&LogIndex{},
		&EVMSyncStatus{},
		&Account{},
		&DepositRecord{},
		&WithdrawalRecord{},
		&Task{},
		&CreateWalletTask{},
		&DepositTask{},
		&WithdrawalTask{},
		&TransferTask{},
		&ConsolidationTask{},
		&TaskUpdatedEvent{},
		&AddressBalance{},
		&InscriptionMintb{},
		&InscriptionBurnb{},
		&Asset{},
		&TokenInfo{},
	); err != nil {
		log.Fatalf("Failed to migrate contract database: %v", err)
	}

	if err := dm.taskDB.AutoMigrate(
		&LogIndex{},
		&Task{},
		&CreateWalletTask{},
		&DepositTask{},
		&WithdrawalTask{},
		&TransferTask{},
		&TaskUpdatedEvent{},
	); err != nil {
		log.Fatalf("Failed to migrate task database: %v", err)
	}

	if err := dm.tokenDB.AutoMigrate(
		&LogIndex{},
		&Asset{},
		&TokenInfo{},
	); err != nil {
		log.Fatalf("Failed to migrate contract database: %v", err)
	}
}
