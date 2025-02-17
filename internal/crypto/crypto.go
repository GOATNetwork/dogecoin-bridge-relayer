package crypto

import (
	"crypto/elliptic"

	"github.com/bnb-chain/tss-lib/v2/tss"
)

type PublicKey interface {
	SerializeCompressed() []byte
	SerializeUncompressed() []byte
}

//type ExtendPublicKey interface {
//	PublicKey
//	Address() string
//	GetType() int
//}

type PrivateKey interface {
	Serialize() []byte
}

type CurveType int

func (e *CurveType) EC() elliptic.Curve {
	switch *e {
	case EDDSA:
		return tss.Edwards()
	case ECDSA:
		return tss.S256()
	default:
		panic("implement me")
	}
}

func (e *CurveType) CurveName() string {
	switch *e {
	case EDDSA:
		return "ed25519"
	case ECDSA:
		return "secp256k1"
	default:
		panic("implement me")
	}
}

func (e *CurveType) String() string {
	return e.CurveName()
}

const (
	ECDSA CurveType = iota
	EDDSA
)
