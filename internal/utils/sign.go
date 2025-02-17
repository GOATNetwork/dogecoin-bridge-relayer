package utils

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// ETHMessagePadding x19 to avoid collision with rlp encode. x01 version byte defined in EIP-191.
var ETHMessagePadding = []byte{0x19, 0x01}

func VerifySig(hash common.Hash, signature []byte, owner common.Address) error {
	// Get the bytes of the signed message
	// decodedMessage := hexutil.MustDecode(signature)
	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}

	// Recover a public key from the signed message
	pubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if pubKey == nil {
		err = errors.New("could not get a public get from the message signature")
	}

	if err != nil {
		return err
	}

	owner2 := crypto.PubkeyToAddress(*pubKey)
	if owner != owner2 {
		return errors.New("signature is invalid")
	}

	return nil
}

func Verify(hash common.Hash, sig []byte, sender common.Address) bool {
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false
	}

	return crypto.PubkeyToAddress(*pubKey) == sender
}

// PersonalMsgHash Returns a hash.
func PersonalMsgHash(messageHash common.Hash) common.Hash {
	// fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n32%s", hex.EncodeToString(messageHash[:]))
	// return crypto.Keccak256Hash([]byte(fullMessage))
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n32%s", messageHash[:])
	return crypto.Keccak256Hash([]byte(msg))
}

// PersonalSign Returns a signature string.
func PersonalSign(messageHash common.Hash, privateKey *ecdsa.PrivateKey) (string, error) {
	hash := PersonalMsgHash(messageHash)

	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}

	signatureBytes[64] += 27

	return hexutil.Encode(signatureBytes), nil
}
