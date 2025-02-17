package btc

import (
	"encoding/json"
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func (w *WalletClient) ChainType() uint8 {
	return types.ChainBitcoin
}

// Verify ExtraData：wire.OutPoint.
func (w *WalletClient) Verify(reqId uint64, ty string, signDigest string, extraData []byte) error {
	bound := &TxBound{}

	err := json.Unmarshal(extraData, bound)
	if err != nil {
		return fmt.Errorf("unmarshal err: %w", err)
	}

	//ctx, ok := w.txContext.Load(reqId)
	//if !ok {
	//	return fmt.Errorf("tx id %d is not found", reqId)
	//}
	//// txCtx, is := ctx.(*TxContext)
	//_, is := ctx.(*TxClient)
	//if !is {
	//	return fmt.Errorf("tx id %d is not TxContext", reqId)
	//}

	// todo
	return nil
}

func (w *WalletClient) ReceiveSignature(res *suite.SignRes) {
	if res.Type == types.SignTxSessionType {
		w.processTxSignResult(res)
	}
}
