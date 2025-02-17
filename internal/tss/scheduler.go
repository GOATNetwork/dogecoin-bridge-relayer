package tss

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	. "github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/crypto/ckd"
	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	ecdsaSigning "github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	eddsaSigning "github.com/bnb-chain/tss-lib/v2/eddsa/signing"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet/bip44"
	"github.com/patrickmn/go-cache"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

type Scheduler struct {
	isProd          bool
	signTimeout     time.Duration
	p2p             p2p.P2PService
	bus             eventbus.Bus
	ctx             context.Context
	cancel          context.CancelFunc
	grw             sync.RWMutex
	groups          map[party.GroupID]*Group
	srw             sync.RWMutex
	sessions        map[party.SessionID]Session[ProposalID]
	proposalSession map[ProposalID]Session[ProposalID]
	crw             sync.RWMutex
	tssClients      map[uint64]suite.TssClient
	signerRw        sync.RWMutex
	signerCtx       map[string]*SignerContext // address:SigContext
	sigInToOut      chan *SessionResult[ProposalID, *tsscommon.SignatureData]
	senateInToOut   chan *SessionResult[ProposalID, *LocalPartySaveData]
	partyData       *PartyData
	localSubmitter  common.Address
	submitterChosen *db.SubmitterChosen
	proposer        *atomic.Value // current submitter
	partners        *atomic.Value // types.Participants
	ecCount         *atomic.Int64
	newGroup        *atomic.Value // *NewGroup
	voter           layer2.Voter
	// only used test
	taskQueue          *pool.Pool[uint64] // created state task
	stateDB            *state.ContractState
	discussedTaskCache *cache.Cache
	testExit           chan error
}

func (m *Scheduler) TssSigner() common.Address {
	return common.HexToAddress(m.tssSigner().Address())
}

func (m *Scheduler) tssSigner() *SignerContext {
	return &SignerContext{
		chainType:          types.ChainEthereum,
		localData:          *m.partyData.ECDSALocalData(),
		keyDerivationDelta: nil,
	}
}

func (m *Scheduler) IsMeeting(signDigest string) bool {
	m.srw.RLock()
	defer m.srw.RUnlock()
	_, ok := m.proposalSession[signDigest]

	return ok
}

func (m *Scheduler) GetPublicKey(address string) crypto.PublicKey {
	signer := m.GetSigner(address)
	if signer == nil {
		return nil
	}

	localData := signer.LocalData()

	return localData.PublicKey()
}

func (m *Scheduler) RegisterTssClient(client suite.TssClient) {
	defer m.crw.Unlock()
	m.crw.Lock()
	m.tssClients[client.ChainId()] = client
}

func NewScheduler(isProd bool, p p2p.P2PService, bus eventbus.Bus, stateDB *state.ContractState, voter layer2.Voter, localSubmitter common.Address) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	pp := atomic.Value{}

	proposer, err := voter.Proposer()
	if err != nil {
		log.Warnf("get proposer error, %s", err.Error())
		log.Infof("TssPublicKeys: %v", len(config.TssPublicKeys))
		proposer = ethcrypto.PubkeyToAddress(*config.TssPublicKeys[0]) // genesis
		pp.Store(proposer)
	} else {
		pp.Store(proposer)
	}

	ps := atomic.Value{}

	partners, err := voter.Participants()
	if err != nil {
		log.Warnf("get partners error, %s", err.Error())

		partners = lo.Map(config.TssPublicKeys, func(item *ecdsa.PublicKey, _ int) common.Address { return ethcrypto.PubkeyToAddress(*item) })
		ps.Store(partners)
	} else {
		ps.Store(partners)
	}

	log.Infof("partners: %v", partners)
	p.UpdateParticipants(partners)

	currentNonce := &atomic.Uint64{}
	nonce, _ := voter.TssNonce()

	if nonce != nil {
		currentNonce.Store(nonce.Uint64())
	}

	newGroup := &atomic.Value{}
	newGroup.Store(nullNewGroup)

	signTimeout := config.AppConfig.Tss.SignTimeout
	if signTimeout == 0 {
		signTimeout = 60
	}

	return &Scheduler{
		ctx:                ctx,
		cancel:             cancel,
		isProd:             isProd,
		signTimeout:        signTimeout * time.Second,
		p2p:                p,
		bus:                bus,
		srw:                sync.RWMutex{},
		grw:                sync.RWMutex{},
		groups:             make(map[party.GroupID]*Group),
		sessions:           make(map[party.SessionID]Session[ProposalID]),
		proposalSession:    make(map[ProposalID]Session[ProposalID]),
		crw:                sync.RWMutex{},
		tssClients:         make(map[uint64]suite.TssClient),
		signerRw:           sync.RWMutex{},
		signerCtx:          make(map[string]*SignerContext),
		sigInToOut:         make(chan *SessionResult[ProposalID, *tsscommon.SignatureData], 1024),
		senateInToOut:      make(chan *SessionResult[ProposalID, *LocalPartySaveData], 1024),
		localSubmitter:     localSubmitter,
		proposer:           &pp,
		partners:           &ps,
		newGroup:           newGroup,
		discussedTaskCache: cache.New(time.Minute*10, time.Minute),
		stateDB:            stateDB,
		voter:              voter,
		partyData:          NewPartyData(config.AppConfig.DB.DbRootDir),
		taskQueue:          pool.NewTaskPool[uint64](),
		testExit:           make(chan error),
	}
}

func (m *Scheduler) Start() {
	m.p2pLoop()
	m.systemProposalLoop()
	m.BlockDetectionThreshold()

	if m.IsGenesis() {
		if m.IsCanProposal() {
			log.Info("TSS keygen process started ", "leader:", m.LocalSubmitter(), " proposer: ", m.Proposer())
			// leader
			m.Genesis() // build senate session
		} else {
			log.Info("TSS keygen process started ", "Candidate:", m.LocalSubmitter(), " proposer: ", m.Proposer())
		}

		m.saveSenateData()
		log.Info("TSS keygen success!", "localSubmitter:", m.LocalSubmitter(), " proposer: ", m.Proposer(), " ECDSA PublicKey: ", m.partyData.ECDSALocalData().PublicKeyBase58(), " EDDSA PublicKey: ", m.partyData.EDDSALocalData().PublicKeyBase58())
	} else {
		log.Info("local data already exists: scheduler begin running")
		log.Info("ECDSA PublicKey: ", m.partyData.ECDSALocalData().PublicKeyBase58(), " EDDSA PublicKey: ", m.partyData.EDDSALocalData().PublicKeyBase58())
	}

	m.initKnownSigner()
	log.Infof("********Scheduler master tss ecdsa address********: %v", m.partyData.GetData(crypto.ECDSA).TssSigner())
	log.Infof("localSubmitter: %v, proposer: %v", m.LocalSubmitter(), m.Proposer())
	m.reGroupResultLoop()
	m.loopSigInToOut()
	m.loopDetectionCondition()
	m.proposalLoopForTest()
	log.Info("Scheduler stared success!")
}

func (m *Scheduler) BlockDetectionLatestState() {
	for m.voter.IsSyncing() {
		time.Sleep(1 * time.Second)
		log.Warn("l2 info is syncing")
	}
}

func (m *Scheduler) SaveSenateSessionResult(sessionResult *SessionResult[ProposalID, *LocalPartySaveData]) {
	if sessionResult.Err != nil {
		panic(sessionResult.Err)
	}

	err := m.partyData.SaveLocalData(sessionResult.SignData)
	utils.Assert(err)
	log.Info("TSS keygen success! SaveSenateSessionResult: ", "localSubmitter:", m.LocalSubmitter())
}

func (m *Scheduler) Stop() {
	m.cancel()
}

func (m *Scheduler) Genesis() {
	_ = m.NewGenerateKeySession(
		crypto.ECDSA,
		SenateProposalIDOfECDSA,
		SenateSessionIDOfECDSA,
		SenateProposal,
	)
	_ = m.NewGenerateKeySession(
		crypto.EDDSA,
		SenateProposalIDOfEDDSA,
		SenateSessionIDOfEDDSA,
		SenateProposal,
	)
}

func (m *Scheduler) saveSenateData() {
	sessionResult := <-m.senateInToOut
	m.SaveSenateSessionResult(sessionResult)
	sessionResult = <-m.senateInToOut
	m.SaveSenateSessionResult(sessionResult)
}

func (m *Scheduler) IsGenesis() bool {
	return !m.partyData.LoadData()
}

func (m *Scheduler) GetUserAddress(coinType, account, index uint32) string {
	return address.GenerateAddressByPath(m.partyData.GetData(types.GetCurveTypeByCoinType(int(coinType))).ECPoint(), coinType, account, index)
}

func (m *Scheduler) initKnownSigner() {
	// tss signer
	m.AddSigner(m.tssSigner())

	// add hot address
	coins := []int{types.CoinTypeBTC, types.CoinTypeEVM, types.CoinTypeSOL, types.CoinTypeSUI}
	accounts := []uint32{0, 1}

	for _, coin := range coins {
		for _, account := range accounts {
			localPartySaveData, key := m.GenerateDerivationWalletProposal(uint32(coin), account, 0)
			m.AddSigner(&SignerContext{
				chainType:          types.GetChainByCoinType(coin),
				localData:          localPartySaveData,
				keyDerivationDelta: key,
			})
		}
	}
}

func (m *Scheduler) BlockDetectionThreshold() {
L:
	for {
		select {
		case <-m.ctx.Done():
			log.Info("DetectionThreshold context done")
		default:
			count := m.p2p.OnlinePeerCount()
			threshold := m.Threshold()
			if count > 0 && threshold > 0 && count > threshold {
				if m.IsGenesis() {
					if count >= m.Participants().Len() {
						break L
					}
				} else {
					break L
				}
			}
			log.Infof("detection online peer count:%d, threshold:%d", count, threshold)
			time.Sleep(time.Second)
		}
	}
}

func (m *Scheduler) Threshold() int {
	return m.Participants().Threshold()
}

func (m *Scheduler) ECPoint(chainType uint8) *ECPoint {
	return m.partyData.GetDataByChain(chainType).ECPoint()
}

func (m *Scheduler) AddSigner(signer *SignerContext) {
	defer m.signerRw.Unlock()
	m.signerRw.Lock()
	m.signerCtx[signer.Address()] = signer
}

func (m *Scheduler) GetSigner(address string) *SignerContext {
	defer m.signerRw.RUnlock()
	m.signerRw.RLock()

	signer := m.signerCtx[strings.ToLower(address)]
	if signer != nil {
		return signer
	}

	account, err := m.stateDB.Account(address)
	if err != nil {
		return nil
	}

	local, keyDerivationDelta := m.GenerateDerivationWalletProposal(uint32(types.GetCoinTypeByChain(account.Chain)), uint32(account.Account), account.Index)

	signer = &SignerContext{
		chainType:          account.Chain,
		localData:          local,
		keyDerivationDelta: keyDerivationDelta,
	}

	m.signerCtx[signer.Address()] = signer

	return signer
}

func (m *Scheduler) AddGroup(group *Group) {
	m.grw.Lock()
	defer m.grw.Unlock()
	m.groups[group.GroupID()] = group
}

func (m *Scheduler) AddSession(session Session[ProposalID]) bool {
	//m.grw.Lock()
	//_, ok := m.groups[session.GroupID()] // todo
	//m.grw.Unlock()
	//
	//if ok {
	//	m.srw.Lock()
	//	m.sessions[session.SessionID()] = session
	//	m.proposalSession[session.ProposalID()] = session
	//	m.srw.Unlock()
	//}
	m.srw.Lock()
	defer m.srw.Unlock()

	if _, ok := m.proposalSession[session.ProposalID()]; ok {
		return false
	}

	m.sessions[session.SessionID()] = session
	m.proposalSession[session.ProposalID()] = session

	return true
}

func (m *Scheduler) GetGroup(groupID party.GroupID) *Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return m.groups[groupID]
}

func (m *Scheduler) GetSession(sessionID party.SessionID) Session[ProposalID] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return m.sessions[sessionID]
}

func (m *Scheduler) GetGroups() []*Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return lo.MapToSlice(m.groups, func(_ party.GroupID, group *Group) *Group { return group })
}

func (m *Scheduler) GetSessions() []Session[ProposalID] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return lo.MapToSlice(m.sessions, func(_ party.SessionID, session Session[ProposalID]) Session[ProposalID] { return session })
}

func (m *Scheduler) ReleaseGroup(groupID party.GroupID) {
	m.grw.Lock()
	defer m.grw.Unlock()
	delete(m.groups, groupID)
}

func (m *Scheduler) SessionRelease(sessionID party.SessionID) {
	m.srw.Lock()
	defer m.srw.Unlock()

	s, ok := m.sessions[sessionID]
	if ok {
		delete(m.sessions, sessionID)
		delete(m.proposalSession, s.ProposalID())
		s.Release()
	}
}

func (m *Scheduler) Release() {
	m.grw.Lock()
	m.groups = make(map[party.GroupID]*Group)
	m.grw.Unlock()
	m.srw.Lock()
	for _, s := range m.sessions {
		s.Release()
	}

	m.sessions = make(map[party.SessionID]Session[ProposalID])
	m.proposalSession = make(map[ProposalID]Session[ProposalID])
	m.srw.Unlock()
	close(m.sigInToOut)
}

func (m *Scheduler) IsDiscussed(taskID uint64) bool {
	_, ok := m.discussedTaskCache.Get(fmt.Sprintf("%d", taskID))
	return ok
}

func (m *Scheduler) AddDiscussedTask(taskID uint64) {
	m.discussedTaskCache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (m *Scheduler) LocalSubmitter() common.Address {
	return m.localSubmitter
}

func (m *Scheduler) Proposer() common.Address {
	p := m.proposer.Load()
	if p != nil {
		return p.(common.Address)
	}

	proposer, err := m.voter.Proposer()
	if !utils.IsZero(proposer) && err == nil {
		m.proposer.Store(proposer)
	}

	return proposer
}

func (m *Scheduler) IsProposer() bool {
	return m.Proposer() == m.LocalSubmitter()
}

func (m *Scheduler) p2pLoop() {
	m.p2p.Bind(p2p.MessageTypeTssMsg, eventbus.EventTssMsg{})
	tssMsgCh := m.bus.Subscribe(eventbus.EventTssMsg{})

	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("Signer stopping...")
				return
			case event := <-tssMsgCh: // from p2p network
				log.Debugf("Received m msg event")

				e := event.(p2p.Message[json.RawMessage])
				proposal := ConvertP2PMsgData(e).(SessionMessage[ProposalID, Proposal])

				err := m.processReceivedProposal(proposal)
				if err != nil {
					log.Warnf("handle session msg error, %v", err)
				}
			}
		}
	}()
	log.Info("p2p loop started")
}

// from layer2 log event.
func (m *Scheduler) systemProposalLoop() {
	go func() {
		pendingTask := m.bus.Subscribe(eventbus.EventTask{})

		for {
			select {
			case <-m.ctx.Done():
				log.Info("proposal loop stopping...")
				return
			case data := <-pendingTask: // from layer2 log sca
				switch v := data.(type) {
				case *types.ParticipantEvent: // regroup
					log.Info("received ParticipantEvent task from layer2 log scan: ", v)
					m.processReGroupProposal(v)

				case *db.SubmitterChosen: // charge proposer
					log.Info("received SubmitterChosen task from layer2 log scan: ", v)
					m.submitterChosen = v
					m.proposer.Store(common.HexToAddress(v.Submitter))
				}
			}
		}
	}()

	log.Info("proposal loop started")
}

const TopN = 20

func (m *Scheduler) loopDetectionCondition() {
	ticker := time.NewTicker(20 * time.Second)

	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("detection condition loop stopped")
			case <-ticker.C:
				latestProposer, err := m.voter.Proposer()
				if err != nil {
					log.Errorf("voter.Proposer err: %v", err)
				} else {
					proposer := m.Proposer()
					if proposer != latestProposer {
						m.proposer.Store(latestProposer)
					}
				}
			}
		}
	}()
}

func (m *Scheduler) IsCanProposal() bool {
	proposer, err := m.voter.Proposer()
	if err != nil || utils.IsZero(proposer) {
		proposer = m.Proposer()
	}

	log.Debugf("proposer is: %v; local submitter: %v", proposer, m.LocalSubmitter())

	is := m.LocalSubmitter() == proposer && m.isJoined()
	if is {
		m.BlockDetectionThreshold()
		m.BlockDetectionLatestState()

		return true
	}

	return false
}

func (m *Scheduler) isJoined() bool {
	return m.Participants().Contains(m.LocalSubmitter())
}

func (m *Scheduler) IsNewJoined() bool {
	return m.newGroup.Load().(*NewGroup).IsNewJoined(m.LocalSubmitter())
}

func (m *Scheduler) Participants() types.Participants {
	if val := m.partners.Load(); val != nil {
		return val.(types.Participants)
	}

	panic("participants is empty")
}

func (m *Scheduler) GenerateDerivationWalletProposal(coinType, account, index uint32) (LocalPartySaveData, *big.Int) {
	// coinType := types.GetCoinTypeByChain(coinType)
	path := bip44.Bip44DerivationPath(coinType, account, index)
	param, err := path.ToParams()
	utils.Assert(err)

	ec := types.GetCurveTypeByCoinType(int(coinType))
	localPartySaveData := m.partyData.GetData(ec)
	l := *localPartySaveData

	chainCode := big.NewInt(int64(coinType)).Bytes() // todo
	keyDerivationDelta, extendedChildPk, err := ckd.DerivingPubkeyFromPath(l.ECPoint(), chainCode, param.Indexes(), ec.EC())
	utils.Assert(err)

	switch ec {
	case crypto.ECDSA:
		data := []ecdsaKeygen.LocalPartySaveData{*l.ECDSAData()}
		err = ecdsaSigning.UpdatePublicKeyAndAdjustBigXj(
			keyDerivationDelta,
			data,
			extendedChildPk.PublicKey,
			ec.EC(),
		)
		utils.Assert(err)
		l.SetData(&data[0])

		return l, keyDerivationDelta
	case crypto.EDDSA:
		data := []eddsaKeygen.LocalPartySaveData{*l.EDDSAData()}
		err = eddsaSigning.UpdatePublicKeyAndAdjustBigXj(
			keyDerivationDelta,
			data,
			extendedChildPk.PublicKey,
			ec.EC(),
		)
		utils.Assert(err)
		l.SetData(&data[0])

		return l, keyDerivationDelta

	default:
		panic(fmt.Errorf("unknown EC type: %v", ec))
	}
}

type NewGroup struct {
	Event    *types.ParticipantEvent
	NewParts types.Participants
	OldParts types.Participants
}

func (g *NewGroup) IsNewJoined(address common.Address) bool {
	return g.NewParts.Contains(address)
}

var nullNewGroup *NewGroup
