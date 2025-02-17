package types

import (
	"bytes"
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Lengths of hashes and addresses in bytes.
const (
	// Byte32Length is the expected length of the hash.
	Byte32Length = 32
)

var Byte32T = reflect.TypeOf(Byte32{})

// Byte32 represents the 32 byte Keccak256 hash of arbitrary data.
type Byte32 [32]byte

// BytesToByte32 sets b to hash.
// If b is larger than len(h), b will be cropped from the left.
func BytesToByte32(b []byte) Byte32 {
	var h Byte32

	h.SetBytes(b)

	return h
}

// BigToByte32 sets byte representation of b to hash.
// If b is larger than len(h), b will be cropped from the left.
func BigToByte32(b *big.Int) Byte32 { return BytesToByte32(b.Bytes()) }

// HexToByte32 sets byte representation of s to hash.
// If b is larger than len(h), b will be cropped from the left.
func HexToByte32(s string) Byte32 { return BytesToByte32(common.FromHex(s)) }

// Cmp compares two hashes.
func (h Byte32) Cmp(other Byte32) int {
	return bytes.Compare(h[:], other[:])
}

// Bytes gets the byte representation of the underlying hash.
func (h Byte32) Bytes() []byte { return h[:] }

// Big converts a hash to a big integer.
func (h Byte32) Big() *big.Int { return new(big.Int).SetBytes(h[:]) }

// Hex converts a hash to a hex string.
func (h Byte32) Hex() string { return hexutil.Encode(h[:]) }

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (h Byte32) TerminalString() string {
	return fmt.Sprintf("%x..%x", h[:3], h[29:])
}

// String implements the stringer interface and is used also by the logger when
// doing full logging into a file.
func (h Byte32) String() string {
	return h.Hex()
}

// Format implements fmt.Formatter.
// Byte32 supports the %v, %s, %q, %x, %X and %d format verbs.
func (h Byte32) Format(s fmt.State, c rune) {
	hexb := make([]byte, 2+len(h)*2)
	copy(hexb, "0x")
	hex.Encode(hexb[2:], h[:])

	switch c {
	case 'x', 'X':
		if !s.Flag('#') {
			hexb = hexb[2:]
		}

		if c == 'X' {
			hexb = bytes.ToUpper(hexb)
		}

		fallthrough
	case 'v', 's':
		_, _ = s.Write(hexb)
	case 'q':
		q := []byte{'"'}
		_, _ = s.Write(q)
		_, _ = s.Write(hexb)
		_, _ = s.Write(q)
	case 'd':
		_, _ = fmt.Fprint(s, ([len(h)]byte)(h))
	default:
		_, _ = fmt.Fprintf(s, "%%!%c(hash=%x)", c, h)
	}
}

// UnmarshalText parses a hash in hex syntax.
func (h *Byte32) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("Byte32", input, h[:])
}

// UnmarshalJSON parses a hash in hex syntax.
func (h *Byte32) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(Byte32T, input, h[:])
}

// MarshalText returns the hex representation of h.
func (h Byte32) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

// SetBytes sets the hash to the value of b.
// If b is larger than len(h), b will be cropped from the left.
func (h *Byte32) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-Byte32Length:]
	}

	copy(h[Byte32Length-len(b):], b)
}

// Generate implements testing/quick.Generator.
func (h Byte32) Generate(rand *rand.Rand, size int) reflect.Value {
	m := rand.Intn(len(h))
	for i := len(h) - 1; i > m; i-- {
		h[i] = byte(rand.Uint32())
	}

	return reflect.ValueOf(h)
}

// Scan implements Scanner for database/sql.
func (h *Byte32) Scan(src interface{}) error {
	srcB, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can't scan %T into Byte32", src)
	}

	if len(srcB) != Byte32Length {
		return fmt.Errorf("can't scan []byte of len %d into Byte32, want %d", len(srcB), Byte32Length)
	}

	copy(h[:], srcB)

	return nil
}

// Value implements valuer for database/sql.
func (h Byte32) Value() (driver.Value, error) {
	return h[:], nil
}

// ImplementsGraphQLType returns true if Byte32 implements the specified GraphQL type.
func (Byte32) ImplementsGraphQLType(name string) bool { return name == "Bytes32" }

// UnmarshalGraphQL unmarshals the provided GraphQL query data.
func (h *Byte32) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		err = h.UnmarshalText([]byte(input))
	default:
		err = fmt.Errorf("unexpected type %T for Byte32", input)
	}

	return err
}

func EthHashToByte32(hash common.Hash) Byte32 {
	return BytesToByte32(hash.Bytes())
}

func Byte32ToEthHash(hash common.Hash) common.Hash {
	return common.BytesToHash(hash.Bytes())
}

func BtcHashToByte32(hash chainhash.Hash) Byte32 {
	return Byte32(hash)
}

func Byte32ToBtcHash(hash Byte32) chainhash.Hash {
	return chainhash.Hash(hash)
}

//func SolHashToByte32(hash sola.Hash) Byte32 {
//	return BytesToByte32(hash.Bytes())
//}
