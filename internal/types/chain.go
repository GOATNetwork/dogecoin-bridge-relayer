package types

type ChainType interface {
	ChainType() uint8
}

type ChainID interface {
	ChainId() uint64
}
