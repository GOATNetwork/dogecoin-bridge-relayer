package evm

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds  = errors.New("insufficient funds for")              // https://blog.csdn.net/wo541075754/article/details/79537043
	ErrNonceTooLow        = errors.New("nonce too low")                       // -32000
	ErrReplacement        = errors.New("replacement transaction underpriced") // -32000
	ErrIntrinsicGasTooLow = errors.New("intrinsic gas too low")               // -32000
	ErrAlreadyKnown       = errors.New("already known")                       // -32000
	ErrGasLimit           = errors.New("exceeds block gas limit")             // -32000
	ErrExecutionReverted  = errors.New("execution reverted")
	ErrWallet             = errors.New("wallet error")
	ErrSendTransaction    = errors.Join(fmt.Errorf("send transaction"), ErrWallet)
	ErrEstimateGas        = errors.Join(fmt.Errorf("estimate gas"), ErrWallet)
	ErrTxFoundTimeOut     = errors.Join(fmt.Errorf("tx found time out"), ErrWallet)
	ErrTxPending          = fmt.Errorf("tx pending: %w", ErrWallet)
	ErrTxCompleted        = fmt.Errorf("tx completed: %w", ErrWallet)
)

var wrapErrorList = append(failErrorList, ErrAlreadyKnown)

var failErrorList = []error{
	ErrNonceTooLow,
	ErrReplacement,
	ErrIntrinsicGasTooLow,
	ErrInsufficientFunds,
	ErrExecutionReverted,
	ErrGasLimit,
}

func (w *WalletClient) HandlerError() {
}
