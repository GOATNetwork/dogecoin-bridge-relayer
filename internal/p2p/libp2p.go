package p2p

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/multiformats/go-multiaddr"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

const (
	handshakeProtocol  = "/nudex/voter/handshake/1.0.0"
	messageProtocol    = "/nudex/voter/message/1.0.0"
	expectedHandshake  = "nudexvoternudexbest"
	messageTopicName   = "gossip-topic"
	heartbeatTopicName = "heartbeat-topic"
	privKeyFile        = "node_private_key.pem"
)

type P2PService interface {
	Bind(msgType MessageType, event eventbus.Event)
	PublishMessage(ctx context.Context, msg any) error
	OnlinePeerCount() int
	IsOnline(submitter string) bool
	UpdateParticipants(partners types.Participants)
}

type Service struct {
	messageTopic  *pubsub.Topic
	state         *state.State
	typeBindEvent sync.Map // MessageType:eventbus.Event

	submitterBindPeerID map[string]peer.ID // submitter:peer.ID
	peerIDBindSubmitter map[peer.ID]string // peer.ID:submitter
	onlineList          map[peer.ID]bool   // peer.ID:bool
	rw                  sync.RWMutex
	localSubmitter      common.Address
	selfPeerID          peer.ID
	partners            *atomic.Value // types.Participants
}

func NewLibP2PService(state *state.State, localSubmitterPrivateKey *ecdsa.PrivateKey) *Service {
	localSubmitter := ethCrypto.PubkeyToAddress(config.SubmitterPrivateKey.PublicKey)

	return &Service{
		state:               state,
		typeBindEvent:       sync.Map{},
		submitterBindPeerID: make(map[string]peer.ID),
		peerIDBindSubmitter: make(map[peer.ID]string),
		onlineList:          make(map[peer.ID]bool),
		rw:                  sync.RWMutex{},
		localSubmitter:      localSubmitter,
		partners:            &atomic.Value{},
	}
}

func (lp *Service) addPeerInfo(peerID peer.ID, submitter string) {
	defer lp.rw.Unlock()
	lp.rw.Lock()

	log.Debugf("addPeerInfo:peerID = %v, submitter = %v", peerID, submitter)
	submitter = strings.ToLower(submitter)
	lp.submitterBindPeerID[submitter] = peerID
	lp.peerIDBindSubmitter[peerID] = submitter
	lp.onlineList[peerID] = true
}

func (lp *Service) OnlinePeerCount() int {
	defer lp.rw.RUnlock()
	lp.rw.RLock()
	return len(lp.onlineList)
}

func (lp *Service) IsOnline(submitter string) bool {
	defer lp.rw.RUnlock()
	lp.rw.RLock()

	peerID, ok := lp.submitterBindPeerID[strings.ToLower(submitter)]
	if ok {
		_, ok = lp.onlineList[peerID]
		return ok
	}

	return false
}

func (lp *Service) UpdateParticipants(partners types.Participants) {
	lp.partners.Store(partners)
}

func (lp *Service) Participants() types.Participants {
	return lp.partners.Load().(types.Participants)
}

func (lp *Service) GroupID() common.Hash {
	return lp.Participants().GroupID()
}

func (lp *Service) IsPartner(address common.Address) bool {
	return lp.Participants().Contains(address)
}

func (lp *Service) online(remotePeerID peer.ID) {
	defer lp.rw.Unlock()
	lp.rw.Lock()
	lp.onlineList[remotePeerID] = true
}

func (lp *Service) offline(remotePeerID peer.ID) {
	defer lp.rw.Unlock()
	lp.rw.Lock()
	delete(lp.onlineList, remotePeerID)
}

func (lp *Service) removePeer(remotePeerID peer.ID) {
	defer lp.rw.Unlock()
	lp.rw.Lock()

	submitter, ok := lp.peerIDBindSubmitter[remotePeerID]
	if ok {
		delete(lp.onlineList, remotePeerID)
		delete(lp.peerIDBindSubmitter, remotePeerID)
		delete(lp.submitterBindPeerID, strings.ToLower(submitter))
	}
}

func (lp *Service) Start(ctx context.Context) {
	self, ps, err := lp.createNodeWithPubSub(ctx)
	if err != nil {
		log.Fatalf("Failed to create libp2p self: %v", err)
	}

	lp.selfPeerID = self.ID()

	PrintNodeAddrInfo(self)

	self.SetStreamHandler(handshakeProtocol, func(s network.Stream) {
		log.Println("New handshake stream")

		err := lp.handleHandshake(s, self)
		if err != nil {
			log.Errorf("Failed to handle handshake message: %v", err)
		}

		s.Close()
	})

	go lp.connectToBootNodes(ctx, self)

	lp.messageTopic, err = ps.Join(messageTopicName)
	if err != nil {
		log.Fatalf("Failed to join message topic: %v", err)
	}

	msgSub, err := lp.messageTopic.Subscribe()
	if err != nil {
		log.Fatalf("Failed to subscribe to message topic: %v", err)
	}

	heartbeatTopic, err := ps.Join(heartbeatTopicName)
	if err != nil {
		log.Fatalf("Failed to join heartbeat topic: %v", err)
	}

	heartbeatSub, err := heartbeatTopic.Subscribe()
	if err != nil {
		log.Fatalf("Failed to subscribe to heartbeat topic: %v", err)
	}

	go lp.handlePubSubMessages(ctx, msgSub)
	go lp.handleHeartbeatMessages(ctx, heartbeatSub)
	go lp.startHeartbeat(ctx, heartbeatTopic)

	go func() {
		msg := Message[string]{
			RequestId:   "1",
			MessageType: MessageTypeUnknown,
			Data:        "Hello, nudex voter libp2p PubSub network with handshake!",
		}
		err = lp.PublishMessage(ctx, msg)
		utils.Assert(err)
	}()

	lp.addPeerInfo(self.ID(), lp.localSubmitter.Hex())

	<-ctx.Done()

	log.Info("LibP2PService is stopping...")

	if err := self.Close(); err != nil {
		log.Errorf("Error closing libp2p self: %v", err)
	}

	log.Info("LibP2PService has stopped.")
}

func (lp *Service) Stop(ctx context.Context) {
	_ = lp.messageTopic.Close()
}

func (lp *Service) connectToBootNodes(ctx context.Context, self host.Host) {
	bootNodeAddList := lo.FilterMap(
		config.AppConfig.P2P.BootNodes,
		func(addr string, index int) (*peer.AddrInfo, bool) {
			peerInfo, err := parseAddr(addr)
			if err != nil {
				return nil, false
			}

			return peerInfo, true
		},
	)

	var wg sync.WaitGroup

	for _, addr := range bootNodeAddList {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					log.Infof("Context cancelled, stopping connection attempts to %s", addr)
					return
				default:
					err := lp.connectToBootNode(ctx, self, addr)
					if err != nil {
						lp.offline(addr.ID)
						log.Errorf("Failed to connect to bootnode %s: %v", addr, err)
						time.Sleep(10 * time.Second)
					} else {
						log.Infof("Successfully connected to bootnode %v", addr)
						lp.online(addr.ID)
						lp.monitorConnection(ctx, self, addr)

						return
					}
				}
			}
		}()
	}

	wg.Wait()
}

func (lp *Service) connectToBootNode(ctx context.Context, self host.Host, peerInfo *peer.AddrInfo) error {
	if err := self.Connect(ctx, *peerInfo); err != nil {
		return fmt.Errorf("failed to connect to bootnode %s: %w", peerInfo, err)
	}

	s, err := self.NewStream(ctx, peerInfo.ID, handshakeProtocol)
	if err != nil {
		return fmt.Errorf("failed to create handshake stream with %s: %w", peerInfo, err)
	}
	defer s.Close()

	err = lp.sendHandshake(s, self)
	if err != nil {
		return fmt.Errorf("failed to send handshake to %s: %w", peerInfo, err)
	}

	return nil
}

func (lp *Service) monitorConnection(ctx context.Context, self host.Host, addr *peer.AddrInfo) {
	for {
		select {
		case <-ctx.Done():
			log.Infof("Context cancelled, stopping monitoring of %s", addr)
			return
		default:
			if self.Network().Connectedness(addr.ID) != network.Connected {
				log.Warnf("Disconnected from %s, attempting to reconnect", addr)

				err := lp.connectToBootNode(ctx, self, addr)
				if err != nil {
					log.Errorf("Failed to reconnect to %s: %v", addr, err)
					lp.offline(addr.ID)
					time.Sleep(5 * time.Second)

					continue
				}

				lp.online(addr.ID)
				log.Infof("Successfully reconnected to %s", addr)
			}

			time.Sleep(20 * time.Second)
		}
	}
}

func parseAddr(addrStr string) (*peer.AddrInfo, error) {
	addr, err := multiaddr.NewMultiaddr(addrStr)
	if err != nil {
		return nil, err
	}

	return peer.AddrInfoFromP2pAddr(addr)
}

func (lp *Service) handshakeMessage() []byte {
	msg := HandshakeMessage{
		PeerID:    lp.selfPeerID.String(),
		Submitter: lp.localSubmitter.Hex(),
		Handshake: expectedHandshake,
		Timestamp: time.Now().Unix(),
	}
	handshakeMsg, err := json.Marshal(&msg)
	utils.Assert(err)

	return handshakeMsg
}

func (lp *Service) sendHandshake(s network.Stream, self host.Host) error {
	err := lp.handleWriteHandshake(s, self)
	if err != nil {
		return err
	}

	handShake, err := lp.handleReadHandshake(s, self)
	if err != nil {
		return err
	}

	id, err := peer.Decode(handShake.PeerID)
	utils.Assert(err)
	lp.addPeerInfo(id, handShake.Submitter)
	log.Info("sendHandshake successful")

	return nil
}

func ListenAddr() string {
	return fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", config.AppConfig.P2P.Port)
}

func (lp *Service) createNodeWithPubSub(ctx context.Context) (host.Host, *pubsub.PubSub, error) {
	secp256k1PrivateKey, err := crypto.UnmarshalSecp256k1PrivateKey(ethCrypto.FromECDSA(config.SubmitterPrivateKey))
	if err != nil {
		return nil, nil, err
	}

	listenAddr := ListenAddr()

	node, err := libp2p.New(
		libp2p.Identity(secp256k1PrivateKey),
		libp2p.Transport(tcp.NewTCPTransport), // TCP only
		libp2p.ListenAddrStrings(listenAddr),  // ipv4 only
	)
	if err != nil {
		return nil, nil, err
	}

	ps, err := pubsub.NewGossipSub(ctx, node)
	if err != nil {
		return nil, nil, err
	}

	return node, ps, nil
}

func loadOrCreatePrivateKey(fileName string) (crypto.PrivKey, error) {
	dbDir := config.AppConfig.DB.DbRootDir
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create database directory: %v", err)
	}

	pemPath := filepath.Join(dbDir, fileName)
	if _, err := os.Stat(pemPath); err == nil {
		privKeyBytes, err := os.ReadFile(pemPath)
		if err != nil {
			return nil, err
		}

		privKey, err := crypto.UnmarshalPrivateKey(privKeyBytes)
		if err != nil {
			return nil, err
		}

		return privKey, nil
	}

	privKey, _, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, 2048, rand.Reader)
	if err != nil {
		return nil, err
	}

	privKeyBytes, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(pemPath, privKeyBytes, 0o600); err != nil {
		return nil, err
	}

	return privKey, nil
}

func FullAddr(addr ma.Multiaddr, peerID peer.ID) string {
	return fmt.Sprintf("%s/p2p/%s", addr.String(), peerID.String())
}

func PrintNodeAddrInfo(node host.Host) {
	addrs := node.Addrs()
	peerID := node.ID()

	for _, addr := range addrs {
		fullAddr := FullAddr(addr, peerID)
		log.Infof("Bootnode address: %s", fullAddr)
	}
}
