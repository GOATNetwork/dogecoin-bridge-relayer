package tss

import (
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/wallet/solana"
	"github.com/nuvosphere/nudex-voter/internal/wallet/sui"
)

type SolTxContext struct {
	c    *solana.SolClient
	tx   *solana.UnSignTx
	task pool.Task[uint64]
}

type SuiTxContext struct {
	signerPubKey []byte
	c            *sui.SuiClient
	tx           *sui.UnSignTx
	task         pool.Task[uint64]
}
