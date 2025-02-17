package tss

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper/testutil"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	_ "github.com/nuvosphere/nudex-voter/internal/wallet/btc"
	_ "github.com/nuvosphere/nudex-voter/internal/wallet/dog"
	_ "github.com/nuvosphere/nudex-voter/internal/wallet/evm"
	_ "github.com/nuvosphere/nudex-voter/internal/wallet/solana"
	_ "github.com/nuvosphere/nudex-voter/internal/wallet/sui"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func createSubmitter(t *testing.T) common.Address {
	pk, err := crypto.GenerateKey()
	assert.Nil(t, err)

	return crypto.PubkeyToAddress(pk.PublicKey)
}

func createDB(t *testing.T, index int) *gorm.DB {
	basePath := filepath.Join("./", strconv.Itoa(index))
	err := os.MkdirAll(basePath, os.ModePerm)
	assert.Nil(t, err)

	path := filepath.Join(basePath, "relayer_data.db")

	relayerDb, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger:         gormlogger.Default.LogMode(gormlogger.Warn),
		TranslateError: true, // https://gorm.golang.ac.cn/docs/error_handling.html
	})
	assert.Nil(t, err)
	db.SetConnParam(relayerDb)

	return relayerDb
}

type Account struct {
	PK      string
	PubKey  string
	Address common.Address
}

var accounts = []Account{
	{
		PK:      "76cbb08e5321cec5f584b2b40b4666d9bbbee59eb3022e80d804e8310b17a105",
		PubKey:  "020b537f46c6da81f84824ce1409bab1f9825fb58b57dcafbf4f4b074e90a0c040",
		Address: common.HexToAddress("0x3a818294ca1F3C27d7588b123Ec43F2546fa07f4"),
	},
	{
		PK:      "ffab86884b5f4696c503e8d0cef97f818d122f44017528c24ce3ac580f12b876",
		PubKey:  "02a8fd23c439e9226f422e94911f06788e0019aa1f8efd4f498f75e4f1d5ef7c0a",
		Address: common.HexToAddress("0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"),
	},
	{
		PK:      "5d0ca3f7b4e63f3308a73537001065ee1d6ff3e217115444b148018a1bcbfaf7",
		PubKey:  "02f82403b0337c908478d381f88582e1051c2a9da22a34cd0a1a5b1d10a85b6256",
		Address: common.HexToAddress("0xf6D37CE75dB465DcDb4c7097bEB9c1D46b171037"),
	},
	{
		PK:      "dd4ae923532c8b47440db5497bf0591769969d3da3ed6ac1d7c2a037033404e9",
		PubKey:  "0349b0799d14fcfd9e0726e037523302515e1082b0b1f23d2d876647aa31ef107d",
		Address: common.HexToAddress("0x1D2cd50A3cF3c55a7982AD54F9f364C1e953Bc57"),
	},
	{
		PK:      "ccb83c6d8cf4d1400ca1d90df2f9c9fafe4b1947ba51c13617603af3bef18590",
		PubKey:  "038801d4a8877f5285a9b3048e6b2e36dbe1b5e00ce4ff98c9f723763d15883c0b",
		Address: common.HexToAddress("0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2"),
	},
}

func TestCreateAddress1(t *testing.T) {
	for _, account := range accounts {
		data, _ := hex.DecodeString(account.PK)
		pk := crypto.ToECDSAUnsafe(data)
		address := crypto.PubkeyToAddress(pk.PublicKey)
		t.Log("pk", hex.EncodeToString(crypto.FromECDSA(pk)), "address:", address, "publicKey:", hex.EncodeToString(crypto.CompressPubkey(&pk.PublicKey)))
	}
}

func TestCreateAddress(t *testing.T) {
	for i := 0; i < testutil.TestPartyCount; i++ {
		pk, err := crypto.GenerateKey()
		assert.Nil(t, err)

		address := crypto.PubkeyToAddress(pk.PublicKey)
		t.Log("pk", hex.EncodeToString(crypto.FromECDSA(pk)), "address:", address)
	}
}

func TestSchedulerOfNewGroup(t *testing.T) {
	utils.SkipCI(t)
	log.SetLevel(log.DebugLevel)

	schedulerList := make([]*Scheduler, 0, len(accounts))

	bus := eventbus.NewBus()
	p2pMocker := NewP2PMocker(bus)
	p2pMocker.SetOnlinePeerCount(testutil.TestPartyCount)

	submitters := types.Participants{}

	var proposer common.Address

	for i := 0; i < testutil.TestPartyCount; i++ {
		submitter := accounts[i].Address
		if i == 2 {
			proposer = submitter
		}

		submitters = append(submitters, submitter)
	}

	t.Log("submitters", submitters)

	createNode := func(index int, submitter common.Address) *Scheduler {
		stateDB := createDB(t, index)
		t.Logf("index: %d, submitter:%v", index, submitter)

		voterContractMocker := NewVoterContractMocker()
		copyParts := types.Participants{}
		copyParts = append(copyParts, submitters...)
		t.Logf("copyParts: %v", copyParts)
		voterContractMocker.SetParticipants(copyParts)
		voterContractMocker.SetProposer(proposer)

		s := NewScheduler(false, p2pMocker, bus, state.NewContractState(stateDB), voterContractMocker, submitter)
		basePath := filepath.Join("./", strconv.Itoa(index))
		err := os.MkdirAll(basePath, os.ModePerm)
		assert.NoError(t, err)

		s.partyData.basePath = basePath
		schedulerList = append(schedulerList, s)

		return s
	}

	// create node
	for i, submitter := range submitters {
		createNode(i, submitter) // specifies the run mode for the first time
	}

	// run node
	for _, s := range schedulerList {
		go s.Start()
	}

	time.Sleep(10 * time.Minute)
	lo.ForEach(schedulerList, func(item *Scheduler, index int) { item.Stop() })
}

var addAccount = Account{
	PK:      "a451cf94141706b5f426dab712cf99753f7f3101abb3125ad6541cd661f35230",
	Address: common.HexToAddress("0xF5D09AE932101D53DDe91659686285F316e4C613"),
}

// add one submitter.
func TestSchedulerOfReGroupForAddAccount(t *testing.T) {
	utils.SkipCI(t)
	log.SetLevel(log.DebugLevel)

	schedulerList := make([]*Scheduler, 0, len(accounts))

	bus := eventbus.NewBus()
	p2pMocker := NewP2PMocker(bus)
	p2pMocker.SetOnlinePeerCount(testutil.TestPartyCount)

	submitters := types.Participants{}

	var proposer common.Address

	for i := 0; i < testutil.TestPartyCount; i++ {
		submitter := accounts[i].Address
		if i == 2 {
			proposer = submitter
		}

		submitters = append(submitters, submitter)
	}

	t.Log("submitters", submitters)

	createNode := func(index int, submitter common.Address) *Scheduler {
		stateDB := createDB(t, index)
		t.Logf("index: %d, submitter:%v", index, submitter)

		voterContractMocker := NewVoterContractMocker()
		copyParts := types.Participants{}
		copyParts = append(copyParts, submitters...)
		t.Logf("copyParts: %v", copyParts)
		voterContractMocker.SetParticipants(copyParts)
		voterContractMocker.SetProposer(proposer)

		s := NewScheduler(false, p2pMocker, bus, state.NewContractState(stateDB), voterContractMocker, submitter)
		basePath := filepath.Join("./", strconv.Itoa(index))
		err := os.MkdirAll(basePath, os.ModePerm)
		assert.NoError(t, err)

		s.partyData.basePath = basePath
		schedulerList = append(schedulerList, s)

		return s
	}

	// 1.create old node
	for i, submitter := range submitters {
		createNode(i, submitter)
	}

	// 2.run old node
	for _, s := range schedulerList {
		go s.Start()
	}

	time.Sleep(5 * time.Second)

	t.Log("new node join")
	// 3.create new node
	s := createNode(5, addAccount.Address) // specifies the run mode for the join node
	t.Logf("new node: %v", s.Participants())
	time.Sleep(1 * time.Second)

	// 4.run new node
	go s.Start()

	// 5.send new participant tx to contact online by owner
	// generate `ParticipantEvent`
	event := &db.ParticipantEvent{
		EventName:   layer2.ParticipantAdded,
		Address:     addAccount.Address.String(),
		BlockNumber: 10,
	}

	time.Sleep(5 * time.Second)

	t.Log("send new join node event")
	// 6.leader(proposer) listen contact event(ParticipantEvent)
	// start regroup
	bus.Publish(eventbus.EventTestTask{}, event)

	// 7.wait end
	time.Sleep(10 * time.Minute)
	lo.ForEach(schedulerList, func(item *Scheduler, index int) { item.Stop() })
}

var removeAccount = Account{
	PK:      "a451cf94141706b5f426dab712cf99753f7f3101abb3125ad6541cd661f35230",
	Address: common.HexToAddress("0xF5D09AE932101D53DDe91659686285F316e4C613"),
}

// remove one submitter.
func TestSchedulerOfReGroupForRemoveAccount(t *testing.T) {
	utils.SkipCI(t)
	log.SetLevel(log.DebugLevel)

	schedulerList := make([]*Scheduler, 0, len(accounts))

	bus := eventbus.NewBus()
	p2pMocker := NewP2PMocker(bus)
	p2pMocker.SetOnlinePeerCount(testutil.TestPartyCount)

	submitters := types.Participants{}

	var proposer common.Address

	for i := 0; i < testutil.TestPartyCount; i++ {
		submitter := accounts[i].Address
		if i == 2 {
			proposer = submitter
		}

		submitters = append(submitters, submitter)
	}

	submitters = append(submitters, removeAccount.Address)

	t.Log("submitters", submitters)

	createNode := func(index int, submitter common.Address) *Scheduler {
		stateDB := createDB(t, index)
		voterContractMocker := NewVoterContractMocker()
		copyParts := types.Participants{}
		copyParts = append(copyParts, submitters...)
		voterContractMocker.SetParticipants(copyParts)
		voterContractMocker.SetProposer(proposer)

		s := NewScheduler(false, p2pMocker, bus, state.NewContractState(stateDB), voterContractMocker, submitter)
		basePath := filepath.Join("./", strconv.Itoa(index))
		err := os.MkdirAll(basePath, os.ModePerm)
		assert.NoError(t, err)

		s.partyData.basePath = basePath
		schedulerList = append(schedulerList, s)

		return s
	}

	// 1.create old node
	for i, submitter := range submitters {
		createNode(i, submitter)
	}

	// 2.run old node
	for _, s := range schedulerList {
		go s.Start()
	}

	time.Sleep(5 * time.Second)

	// 5.send new participant tx to contact online by owner
	// generate `ParticipantEvent`
	event := &db.ParticipantEvent{
		EventName:   layer2.ParticipantRemoved,
		Address:     removeAccount.Address.String(),
		BlockNumber: 10,
	}

	time.Sleep(5 * time.Second)

	t.Log("send new join node event")
	// 6.leader(proposer) listen contact event(ParticipantEvent)
	// start regroup
	bus.Publish(eventbus.EventTestTask{}, event)

	// 7.wait end
	time.Sleep(10 * time.Minute)
	lo.ForEach(schedulerList, func(item *Scheduler, index int) { item.Stop() })
}

// generate signature.
func TestSchedulerSignature(t *testing.T) {
	utils.SkipCI(t)
	log.SetLevel(log.DebugLevel)

	schedulerList := make([]*Scheduler, 0, len(accounts))

	bus := eventbus.NewBus()
	p2pMocker := NewP2PMocker(bus)
	p2pMocker.SetOnlinePeerCount(testutil.TestPartyCount)

	submitters := types.Participants{}

	var proposer common.Address

	for i := 0; i < testutil.TestPartyCount; i++ {
		submitter := accounts[i].Address
		if i == 2 {
			proposer = submitter
		}

		submitters = append(submitters, submitter)
	}

	t.Log("submitters", submitters)

	createNode := func(index int, submitter common.Address) *Scheduler {
		stateDB := createDB(t, index)
		voterContractMocker := NewVoterContractMocker()
		copyParts := types.Participants{}
		copyParts = append(copyParts, submitters...)
		voterContractMocker.SetParticipants(copyParts)
		voterContractMocker.SetProposer(proposer)

		basePath := filepath.Join("./", strconv.Itoa(index))
		err := os.MkdirAll(basePath, os.ModePerm)
		assert.NoError(t, err)

		config.AppConfig.DB.DbRootDir = basePath
		s := NewScheduler(false, p2pMocker, bus, state.NewContractState(stateDB), voterContractMocker, submitter)

		schedulerList = append(schedulerList, s)

		return s
	}

	// 1.create node
	for i, submitter := range submitters {
		createNode(i, submitter)
	}

	// 2.run node
	for _, s := range schedulerList {
		s.Start()
	}

	// time.Sleep(10 * time.Second)

	// 3.send tx to contact online by owner

	// generate `CreateWalletTask`
	task := &db.CreateWalletTask{
		TaskId:  99999,
		Account: 1,
		Chain:   1, // eth
		Index:   0,
	}

	t.Log("send create wallet task")
	// 6.leader(proposer) listen contact task(CreateWalletTask)
	bus.Publish(eventbus.EventTestTask{}, task)

	// 7.wait end
	for _, scheduler := range schedulerList {
		err := <-scheduler.testExit
		assert.NoError(t, err)
	}

	lo.ForEach(schedulerList, func(item *Scheduler, index int) { item.Stop() })
}

// solana withdraw signature.
func TestSchedulerSolanaWithdrawSignature(t *testing.T) {
	utils.SkipCI(t)
	log.SetLevel(log.DebugLevel)

	schedulerList := make([]*Scheduler, 0, len(accounts))

	bus := eventbus.NewBus()
	p2pMocker := NewP2PMocker(bus)
	p2pMocker.SetOnlinePeerCount(testutil.TestPartyCount)

	submitters := types.Participants{}

	var proposer common.Address

	for i := 0; i < testutil.TestPartyCount; i++ {
		submitter := accounts[i].Address
		if i == 2 {
			proposer = submitter
		}

		submitters = append(submitters, submitter)
	}

	t.Log("submitters", submitters)

	createNode := func(index int, submitter common.Address) *Scheduler {
		stateDB := createDB(t, index)
		voterContractMocker := NewVoterContractMocker()
		copyParts := types.Participants{}
		copyParts = append(copyParts, submitters...)
		voterContractMocker.SetParticipants(copyParts)
		voterContractMocker.SetProposer(proposer)

		s := NewScheduler(false, p2pMocker, bus, state.NewContractState(stateDB), voterContractMocker, submitter)
		basePath := filepath.Join("./", strconv.Itoa(index))
		err := os.MkdirAll(basePath, os.ModePerm)
		assert.NoError(t, err)

		s.partyData.basePath = basePath

		schedulerList = append(schedulerList, s)

		return s
	}

	// 1.create node
	for i, submitter := range submitters {
		createNode(i, submitter)
	}

	// 2.run node
	for _, s := range schedulerList {
		go s.Start()
	}

	time.Sleep(5 * time.Second)

	// 3.send tx to contact online by owner
	// generate pending `WithdrawalTask`
	task := &db.WithdrawalTask{
		TaskId: 1,
		// TargetAddress: "2cz1TgTjQSdmGSjUiL9Z1QupEAUD3S46AX4KB4Uefr59",
		Amount: decimal.NewFromInt(3000),
		Chain:  types.ChainSolana,
		// ChainID:         0,
		// Ticker:          "SOL",
	}

	t.Log("send WithdrawalTask")
	// 6.leader(proposer) listen contact task(WithdrawalTask)
	bus.Publish(eventbus.EventTestTask{}, task)

	// 7.wait end
	time.Sleep(10 * time.Minute)
	lo.ForEach(schedulerList, func(item *Scheduler, index int) { item.Stop() })
}

// sui withdraw signature.
func TestSchedulerSuiWithdrawSignature(t *testing.T) {
	utils.SkipCI(t)
	log.SetLevel(log.DebugLevel)

	schedulerList := make([]*Scheduler, 0, len(accounts))

	bus := eventbus.NewBus()
	p2pMocker := NewP2PMocker(bus)
	p2pMocker.SetOnlinePeerCount(testutil.TestPartyCount)

	submitters := types.Participants{}

	var proposer common.Address

	for i := 0; i < testutil.TestPartyCount; i++ {
		submitter := accounts[i].Address
		if i == 2 {
			proposer = submitter
		}

		submitters = append(submitters, submitter)
	}

	t.Log("submitters", submitters)

	createNode := func(index int, submitter common.Address) *Scheduler {
		stateDB := createDB(t, index)
		voterContractMocker := NewVoterContractMocker()
		copyParts := types.Participants{}
		copyParts = append(copyParts, submitters...)
		voterContractMocker.SetParticipants(copyParts)
		voterContractMocker.SetProposer(proposer)

		s := NewScheduler(false, p2pMocker, bus, state.NewContractState(stateDB), voterContractMocker, submitter)
		basePath := filepath.Join("./", strconv.Itoa(index))
		err := os.MkdirAll(basePath, os.ModePerm)
		assert.NoError(t, err)

		s.partyData.basePath = basePath

		schedulerList = append(schedulerList, s)

		return s
	}

	// 1.create node
	for i, submitter := range submitters {
		createNode(i, submitter)
	}

	// 2.run node
	for _, s := range schedulerList {
		go s.Start()
	}

	time.Sleep(5 * time.Second)

	// 3.send tx to contact online by owner
	// generate pending `WithdrawalTask`
	task := &db.WithdrawalTask{
		TaskId: 1,
		// from 0x6e3b7823972bffb6107abc008ff257c89a268fb0e3dfb27d503741fe9a38749f
		// TargetAddress: "0x9099b85cce1e63a584f981390bf3457611df3f7778c0d77de3f16cb57951bcf9",
		Amount:  decimal.NewFromInt(3000),
		Chain:   types.ChainSui,
		ChainID: 0,
		// Ticker:          "sui",
	}

	t.Log("send WithdrawalTask")
	// 6.leader(proposer) listen contact task(WithdrawalTask)
	bus.Publish(eventbus.EventTestTask{}, task)

	// 7.wait end
	time.Sleep(10 * time.Minute)
	lo.ForEach(schedulerList, func(item *Scheduler, index int) { item.Stop() })
}

func TestValue(t *testing.T) {
	newGroup := NewGroup{
		Event:    nil,
		NewParts: []common.Address{addAccount.Address},
		OldParts: nil,
	}

	value := atomic.Value{}
	ll := value.Load()
	assert.Nil(t, ll)
	t.Log(ll)
	value.Store(&newGroup)
	t.Log(value.Load())
	t.Logf("&loadValue = %p, &newGroup = %p", value.Load(), &newGroup)

	var null *NewGroup

	t.Logf("null %v", null)

	if null == nil {
		t.Log("null == nil")
	}

	var obj any = null

	t.Logf("obj %v", obj)
	// assert.Equal(t, nil, obj) !!!
	assert.NotEqual(t, obj, nil)
	assert.Equal(t, obj, null)
	value.Store(null)
	assert.Nil(t, value.Load())

	if value.Load() == null {
		t.Log("value.Load() == null, var null *NewGroup")
	}

	if value.Load() != nil {
		t.Log("value.Load() != nil")
	}

	t.Log(value.Load())
	t.Logf("&loadValue = %p, &newGroup = %p", value.Load(), &newGroup)

	otherValue := atomic.Value{}
	otherValue.Store(newGroup)
	loadValue := otherValue.Load().(NewGroup)
	assert.Equal(t, loadValue, newGroup)
	assert.Equal(t, otherValue.Load(), newGroup)

	if &loadValue == &newGroup {
		t.Log("&loadValue == &newGroup")
	}

	otherGroup := NewGroup{
		Event:    nil,
		NewParts: nil,
		OldParts: nil,
	}

	if &loadValue != &otherGroup {
		t.Log("&loadValue != &otherGroup")
		t.Logf("&loadValue = %p, &otherGroup = %p", &loadValue, &otherGroup)
	}

	assert.Equal(t, &loadValue, &newGroup)
	t.Logf("&loadValue = %p, &newGroup = %p", &loadValue, &newGroup)
}

func TestVerifySig(t *testing.T) {
	signature, err := hexutil.Decode("0x7b5271a558d9319c395ac0c0403baa3a1cff47c2790870fb327fe89b8d801ea7531563c7763bc364dce8df59cf22d59f59f0272a506f01715985ecd3562f597801")
	assert.NoError(t, err)
	err = utils.VerifySig(common.HexToHash("0x0ff92700e1f5c45afab5763ddda39c503dd2fba606aef046278882b13d14ee50"), signature, common.HexToAddress("0xB43EB0e9Ec8040737FFcc144073C72Cf68bC4bab"))
	assert.NoError(t, err)
}
