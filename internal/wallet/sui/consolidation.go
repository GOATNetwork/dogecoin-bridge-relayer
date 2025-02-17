package sui

import (
	"fmt"
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) ProcessConsolidation(task *db.Consolidation) {
	balance, err := w.ContractState().GetAddressBalance(task.FromAddress, task.ContractAddress)
	if err != nil {
		log.Errorf("Failed to get address balance for %v: %v", task.FromAddress, err)
		return
	}

	if balance.Cmp(task.Amount) == 1 {
		err := w.signConsolidationTask(task)
		if err != nil {
			log.Errorf("Failed to sign task for %v: %v", task.Decimals, err)
		}
	}
}

func (w *WalletClient) signConsolidationTask(task *db.Consolidation) error {
	hotAddress := address.HotAddressOfSui(w.tss.ECPoint(w.ChainType()))
	log.Debugf("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, task.FromAddress, task.Amount)

	unSignTx, err := w.client.BuildCollectFoundTx(CoinType(task.ContractAddress, task.Symbol), task.FromAddress, hotAddress)
	if err != nil {
		return fmt.Errorf("failed to build unSignTx: %w", err)
	}

	proposal := new(big.Int).SetBytes(unSignTx.Blake2bHash())
	w.txContext.Store(task.TaskID(), &TxContext{
		signerPubKey: w.tss.GetPublicKey(task.FromAddress).SerializeCompressed(),
		tx:           unSignTx,
		task:         task,
	})

	req := &suite.SignReq{
		SeqId:      task.TaskID(),
		Type:       types.SignTxSessionType,
		ChainId:    w.ChainId(),
		Signer:     task.FromAddress,
		DataDigest: proposal.String(),
		SignData:   proposal.Bytes(),
		ExtraData:  nil,
	}

	w.tss.Sign(req)

	return nil
}
