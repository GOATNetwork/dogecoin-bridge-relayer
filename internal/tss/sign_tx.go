package tss

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (m *Scheduler) loopSigInToOut() {
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("tss signature read result loop stopped")
			case result := <-m.sigInToOut:
				log.Infof("finish consensus, sessionID:%s", result.SessionID)
				info := fmt.Sprintf("tss signature sessionID=%v, groupID=%v, ProposalID=%v", result.SessionID, result.GroupID, result.ProposalID)

				switch result.Type {
				case types.SignTxSessionType, types.SignOperationSessionType:
					log.Debugf("result.SignData.Signature: len: %d, result.SignData.Signature: %x", len(result.SignData.Signature), result.SignData.Signature)
					log.Debugf("result.SignData.SignatureRecovery: len: %d, result.SignData.SignatureRecovery: %x", len(result.SignData.SignatureRecovery), result.SignData.SignatureRecovery)

					signature := result.SignData.Signature
					if config.AppConfig.IsEVM(result.ChainId) {
						signature = secp256k1Signature(result.SignData)
					}

					if result.Err != nil {
						// todo
						log.Errorf("result error:%v, info: %v", result.Err, info)
					}

					m.PostClient(
						result.ChainId,
						result.SeqId,
						result.Type,
						result.ProposalID,
						signature,
						result.ExtraData,
						result.Err,
					)
				case types.SignTaskSessionType:
					m.testExit <- result.Err

				default:
					log.Infof("tss signature result: %v", result)
				}
			}
		}
	}()
}

func (m *Scheduler) PostClient(chainId, SeqId uint64, ty, signDigest string, signature, data []byte, err error) {
	defer m.crw.RUnlock()
	m.crw.RLock()

	c, ok := m.tssClients[chainId]
	if ok {
		c.ReceiveSignature(&suite.SignRes{
			SeqId:      SeqId,
			Type:       ty,
			DataDigest: signDigest,
			Signature:  signature,
			ExtraData:  data,
			Err:        err,
		})
	}
}

func (m *Scheduler) verifySign(msg *SessionMessage[ProposalID, Proposal]) error {
	defer m.crw.RUnlock()
	m.crw.RLock()

	c, ok := m.tssClients[msg.ChainId]
	if ok {
		return c.Verify(msg.SeqId, msg.Type, msg.ProposalID, msg.Data)
	}

	return errors.New("tss client not found")
}

func (m *Scheduler) Sign(req *suite.SignReq) {
	signerCtx := m.GetSigner(req.Signer)

	if m.IsCanProposal() {
		log.Warn("sign proposal")
		m.NewSignSessionWitKey(
			req.ChainId,
			ZeroSessionID,
			req.SeqId,
			req.DataDigest,
			new(big.Int).SetBytes(req.SignData),
			req.Type,
			req.ExtraData,
			signerCtx,
		)
	} else {
		log.Warn("not sign proposal")
	}
}
