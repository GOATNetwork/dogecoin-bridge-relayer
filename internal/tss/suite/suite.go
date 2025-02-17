package suite

import (
	tssCrypto "github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
)

type SignReq struct {
	SeqId      uint64
	Type       string
	ChainId    uint64
	Signer     string
	DataDigest string
	SignData   []byte
	ExtraData  []byte
}

type SignRes struct {
	SeqId      uint64
	Type       string
	DataDigest string
	Signature  []byte
	ExtraData  []byte
	Err        error
}

type TssService interface {
	ECPoint(chainType uint8) *tssCrypto.ECPoint
	GetUserAddress(coinType, account, index uint32) string
	GetPublicKey(address string) crypto.PublicKey
	TssSigner() common.Address
	IsMeeting(signDigest string) bool
	Sign(req *SignReq)
	IsProposer() bool
	IsCanProposal() bool
	Proposer() common.Address
	LocalSubmitter() common.Address
	RegisterTssClient(client TssClient)
}

type TssClient interface {
	Verify(reqId uint64, ty string, signDigest string, extraData []byte) error
	ReceiveSignature(res *SignRes)
	ChainId() uint64
}
