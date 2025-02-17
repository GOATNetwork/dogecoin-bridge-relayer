package evm

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts/codec"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) ChainType() uint8 {
	return types.ChainEthereum
}

func (w *WalletClient) Verify(reqId uint64, ty string, signDigest string, extraData []byte) error {
	switch ty {
	case types.SignTxSessionType:
		if err := w.verifyTxSession(reqId, signDigest); err != nil {
			return err
		}
	case types.SignOperationSessionType:
		// todo
	}

	return nil
}

func (w *WalletClient) verifyTxSession(reqId uint64, signDigest string) error {
	ctx, ok := w.pendingTx.Load(signDigest)
	if !ok {
		return fmt.Errorf("tx id %d is not found", reqId)
	}

	txCtx, is := ctx.(*TxContext)
	if !is {
		return fmt.Errorf("tx id %d is not TxContext", reqId)
	}

	if txCtx.TxHash() != common.HexToHash(signDigest) {
		return fmt.Errorf("tx id %d hash does not match", reqId)
	}

	return nil
}

func (w *WalletClient) ReceiveSignature(res *suite.SignRes) {
	log.Debugf("ReceiveSignature: type: %v, req id: %d", res.Type, res.SeqId)

	switch res.Type {
	case types.SignOperationSessionType:
		w.handleSignOperationSession(res)
	case types.SignTxSessionType:
		w.processTxSignResult(res)
	}
}

func (w *WalletClient) handleSignOperationSession(res *suite.SignRes) {
	op := w.operationsQueue.Get(res.SeqId)
	if op != nil {
		w.operationsQueue.Remove(res.SeqId)
		w.operationsQueue.RemoveTopN(res.SeqId)

		operations := op.(*Operations)
		operations.Signature = res.Signature
		w.processOperationSignResult(operations)
		lo.ForEach(operations.Operation, func(item codec.Operation, _ int) {
			w.RemoveDiscussingTask(item.TaskId)
		})
	}
}

func (w *WalletClient) signTx(ctx *TxContext) error {
	var err error

	switch ctx.ty {
	case db.TaskTypeWithdrawal, db.TaskTypeConsolidation, db.TaskTypeTransfer:
		err = w.signStandardTx(ctx)
	case db.TaskTypeOperations:
		err = w.signOperationsTx(ctx)
	default:
		err = fmt.Errorf("unknown task:%d, type %d", ctx.SeqID(), ctx.ty)
	}

	if err != nil {
		return err
	}

	signTx, err := w.TransactionWithSignature(ctx.UnSignTx(), ctx.sig)
	if err != nil {
		return err
	}

	ctx.tx = signTx

	return w.walletState.UpdateTransaction(nil, signTx)
}

func (w *WalletClient) signStandardTx(ctx *TxContext) error {
	signer, err := w.Signer().Sender(ctx.tx)
	if err != nil {
		return err
	}

	digest := ctx.DigestHash()
	req := &suite.SignReq{
		SeqId:      ctx.SeqID(),
		Type:       types.SignTxSessionType,
		ChainId:    w.ChainId(),
		Signer:     strings.ToLower(signer.String()),
		DataDigest: digest.String(),
		SignData:   digest.Bytes(),
		ExtraData:  nil,
	}

	w.tss.Sign(req)
	select {
	case <-w.ctx.Done():
		return w.ctx.Err()
	case err := <-ctx.notify:
		return err
	}
}

func (w *WalletClient) signOperationsTx(ctx *TxContext) error {
	sig, err := w.SignOperationNewTx(ctx.UnSignTx())
	if err != nil {
		return err
	}

	ctx.sig = sig

	return nil
}

func (w *WalletClient) NewTxContext(ty int, reqId uint64, tx *etypes.Transaction) *TxContext {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	return &TxContext{
		err:    nil,
		ty:     ty,
		seqID:  reqId,
		tx:     tx,
		sig:    nil,
		notify: make(chan error, 1),
		ctx:    ctx,
		cancel: cancel,
	}
}
