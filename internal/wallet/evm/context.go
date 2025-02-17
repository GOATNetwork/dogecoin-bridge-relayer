package evm

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	vtypes "github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

type TxContext struct {
	err    error
	ty     int // operation、withdraw、consolidation、transfer
	seqID  uint64
	tx     *types.Transaction
	sig    []byte
	notify chan error
	ctx    context.Context
	cancel context.CancelFunc
}

func (t *TxContext) TxHash() common.Hash {
	return t.tx.Hash()
}

func (t *TxContext) SeqID() uint64 {
	return t.seqID
}

func (t *TxContext) IsSig() bool {
	return t.sig != nil
}

func (t *TxContext) UnSignTx() *types.Transaction {
	return t.tx
}

func (t *TxContext) DigestHash() common.Hash {
	tx := t.tx
	signer := types.LatestSignerForChainID(tx.ChainId())
	msgHash := signer.Hash(tx)

	return msgHash
}

func (t *TxContext) From() common.Address {
	tx := t.tx
	signer := types.LatestSignerForChainID(tx.ChainId())
	from, err := signer.Sender(tx)
	utils.Assert(err)

	return from
}

func (t *TxContext) SignType() string {
	switch t.ty {
	case db.TaskTypeConsolidation, db.TaskTypeWithdrawal, db.TaskTypeTransfer:
		return vtypes.SignTxSessionType
	case db.TaskTypeOperations:
		return vtypes.SignOperationSessionType
	default:
		panic("unknown task type")
	}
}
