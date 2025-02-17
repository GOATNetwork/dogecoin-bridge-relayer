package types

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
)

func DefaultPartners() Participants {
	// todo online contact get address list
	return lo.Map(
		config.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return ethCrypto.PubkeyToAddress(*pubKey) },
	)
}

func PartyKey(ec crypto.CurveType, participants Participants, address common.Address) *big.Int {
	key := new(big.Int).Add(address.Big(), big.NewInt(int64(ec)))
	return key.Add(key, participants.GroupID().Big())
}

func PartyID(ec crypto.CurveType, participants Participants, address common.Address) string {
	return PartyKey(ec, participants, address).Text(16)
}

type TxClient interface {
	Post(hash, signature []byte)
	ChainType() int
	BuildTx(to string, amount int64) error
	SendTx() error
	WaitTxSuccess() error
	TxHash() []byte
	NextSignTask() []byte
}

func UnmarshalJson[T any](data json.RawMessage) T {
	var obj T

	err := json.Unmarshal(data, &obj)
	if err != nil || data == nil {
		panic(fmt.Errorf("unmarshal data:%v, error: %w", data, err))
	}

	return obj
}

type BatchData struct {
	Ids []uint64 `json:"ids"`
}

func (b *BatchData) Bytes() []byte {
	data, err := json.Marshal(b)
	utils.Assert(err)

	return data
}

func (b *BatchData) FromBytes(data []byte) {
	*b = UnmarshalJson[BatchData](data)
}
