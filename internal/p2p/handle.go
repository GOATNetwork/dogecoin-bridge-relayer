package p2p

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	log "github.com/sirupsen/logrus"
)

var ErrHandshake = errors.New("hand shake error")

func (lp *Service) handleReadHandshake(s network.Stream, self host.Host) (*HandshakeMessage, error) {
	buf := make([]byte, 1024)

	n, err := s.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("error reading handshake message: %w", err)
	}

	handshakeMsg := buf[:n]
	log.Infof("Received handshake message: %s", string(handshakeMsg))

	handShake := HandshakeMessage{}
	err = json.Unmarshal(handshakeMsg, &handShake)
	utils.Assert(err)

	remotePeerID := s.Conn().RemotePeer()
	if handShake.Handshake != expectedHandshake || remotePeerID.String() != handShake.PeerID {
		log.Warn("Invalid handshake message received, closing connection")

		_ = s.Reset()
		// disconnect peer
		// s.Conn().Close()
		err = self.Network().ClosePeer(remotePeerID)

		return nil, errors.Join(err, ErrHandshake)
	}

	return &handShake, nil
}

func (lp *Service) handleWriteHandshake(s network.Stream, self host.Host) error {
	handshakeMsg := lp.handshakeMessage()

	_, err := s.Write(handshakeMsg)
	if err != nil {
		return fmt.Errorf("error writing handshake message: %w", err)
	}

	return nil
}

// handleHandshake: echo.
func (lp *Service) handleHandshake(s network.Stream, self host.Host) error {
	remotePeerPublicKey, err := s.Conn().RemotePeer().ExtractPublicKey()
	if err != nil {
		return fmt.Errorf("error extracting public key: %w", err)
	}

	if remotePeerPublicKey.Type() != pb.KeyType_Secp256k1 {
		return fmt.Errorf("%w: %s", ErrHandshake, remotePeerPublicKey.Type())
	}

	remotePeerStdPublicKey, err := crypto.PubKeyToStdKey(remotePeerPublicKey)
	if err != nil {
		return fmt.Errorf("error extracting std public key: %w", err)
	}

	secp256k1PublicKey := remotePeerStdPublicKey.(*crypto.Secp256k1PublicKey)

	res, err := secp256k1PublicKey.Raw()
	if err != nil {
		return fmt.Errorf("error extracting secp256k1PublicKey: %w", err)
	}

	publicKey, err := ethCrypto.DecompressPubkey(res)
	if err != nil {
		return fmt.Errorf("error decompressing secp256k1PublicKey: %w", err)
	}

	remoteSubmitter := ethCrypto.PubkeyToAddress(*publicKey)

	if !lp.IsPartner(remoteSubmitter) {
		// todo
		// return fmt.Errorf("%w: remoteSubmitter: %v", ErrHandshake, remoteSubmitter)
		log.Errorf("%v: remoteSubmitter: %v", ErrHandshake, remoteSubmitter)
	}

	handShake, err := lp.handleReadHandshake(s, self)
	if err != nil {
		return err
	}

	err = lp.handleWriteHandshake(s, self)
	if err != nil {
		return err
	}

	id, err := peer.Decode(handShake.PeerID)
	utils.Assert(err)
	lp.addPeerInfo(id, handShake.Submitter)
	log.Info("handleHandshake successful")

	return nil
}

func (lp *Service) Bind(msgType MessageType, event eventbus.Event) {
	lp.typeBindEvent.Store(msgType, event)
}

func (lp *Service) PublishMessage(ctx context.Context, msg any) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return errors.New("failed to marshal message")
	}

	if lp.messageTopic == nil {
		startTime := time.Now()
		if time.Since(startTime) >= 10*time.Second {
			return errors.New("message topic is nil, cannot publish message")
		}

		if lp.messageTopic == nil {
			time.Sleep(1 * time.Second)
		}
	}

	return lp.messageTopic.Publish(ctx, msgBytes)
}

func (lp *Service) handlePubSubMessages(ctx context.Context, sub *pubsub.Subscription) {
	for {
		select {
		case <-ctx.Done():
			log.Info("Context cancelled, exiting handlePubSubMessages")
			return
		default:
			msg, err := sub.Next(ctx)
			if err != nil {
				log.Errorf("Error reading message from pubsub: %v", err)
				continue
			}

			if msg.ReceivedFrom == lp.selfPeerID {
				log.Debug("Received message from self, ignore")
				continue
			}

			var receivedMsg Message[json.RawMessage]
			if err := json.Unmarshal(msg.Data, &receivedMsg); err != nil {
				log.Errorf("Error unmarshaling pubsub message: %v", err)
				continue
			}

			dataStr := fmt.Sprintf("%v", receivedMsg.Data)
			if len(dataStr) > 200 {
				dataStr = dataStr[:200] + "..."
			}

			log.Debugf("Received message via pubsub: ID=%d, RequestId=%s, Data=%v", receivedMsg.MessageType, receivedMsg.RequestId, dataStr)

			event, ok := lp.typeBindEvent.Load(receivedMsg.MessageType)
			if ok {
				lp.state.Bus().Publish(event, receivedMsg)
			} else {
				log.Warnf("Unknown message type: %d", receivedMsg.MessageType)
			}
		}
	}
}

func (lp *Service) handleHeartbeatMessages(ctx context.Context, sub *pubsub.Subscription) {
	for {
		select {
		case <-ctx.Done():
			log.Info("Context cancelled, exiting handleHeartbeatMessages")
			return
		default:
			msg, err := sub.Next(ctx)
			if err != nil {
				log.Errorf("Error reading heartbeat message from pubsub: %v", err)
				continue
			}

			if msg.ReceivedFrom == lp.selfPeerID {
				log.Debug("Received heartbeat from self, ignore")
				continue
			}

			var hbMsg HeartbeatMessage
			if err := json.Unmarshal(msg.Data, &hbMsg); err != nil {
				log.Errorf("Error unmarshaling heartbeat message: %v", err)
				continue
			}

			id, err := peer.Decode(hbMsg.PeerID)
			if err != nil {
				log.Errorf("Error decoding peer ID from heartbeat message: %v", err)
				continue
			}

			lp.addPeerInfo(id, hbMsg.Message)
			log.Infof("Received heartbeat from %d-%s: %s", hbMsg.Timestamp, hbMsg.PeerID, hbMsg.Message)
		}
	}
}

func (lp *Service) startHeartbeat(ctx context.Context, topic *pubsub.Topic) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			hbMsg := HeartbeatMessage{
				PeerID:    lp.selfPeerID.String(),
				Message:   lp.localSubmitter.Hex(),
				Timestamp: time.Now().Unix(),
			}

			msgBytes, err := json.Marshal(hbMsg)
			if err != nil {
				log.Errorf("Failed to marshal heartbeat message: %v", err)
				continue
			}

			if err := topic.Publish(ctx, msgBytes); err != nil {
				log.Errorf("Failed to publish heartbeat message: %v", err)
			} else {
				log.Infof("Heartbeat message sent by %s", hbMsg.PeerID)
			}

		case <-ctx.Done():
			return
		}
	}
}
