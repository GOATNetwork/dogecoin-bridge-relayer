package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

const (
	WALLET_TYPE_P2WPKH = "P2WPKH"
	WALLET_TYPE_P2PKH  = "P2PKH"
	WALLET_TYPE_P2SH   = "P2SH"
	WALLET_TYPE_P2WSH  = "P2WSH"
	WALLET_TYPE_P2TR   = "P2TR"

	SMALL_UTXO_DEFINE = 50000000 // 0.5 BTC
)

// MsgUtxoDeposit defines deposit UTXO broadcast to p2p which received in relayer rpc.
type MsgUtxoDeposit struct {
	RawTx       string `json:"raw_tx"`
	TxId        string `json:"tx_id"`
	OutputIndex int    `json:"output_index"`
	SignVersion uint32 `json:"sign_version"`
	EvmAddr     string `json:"evm_addr"`
	Amount      int64  `json:"amount"`
	Timestamp   int64  `json:"timestamp"`
}

type MsgSendOrderBroadcasted struct {
	TxId         string `json:"tx_id"`
	ExternalTxId string `json:"external_tx_id"`
}

// MsgUtxoWithdraw defines withdraw UTXO broadcast to p2p which received in relayer rpc.
type MsgUtxoWithdraw struct {
	TxId      string `json:"tx_id"`
	EvmAddr   string `json:"evm_addr"`
	Timestamp int64  `json:"timestamp"`
}

type BtcBlockExt struct {
	wire.MsgBlock

	BlockNumber uint64
}

func isOpReturn(txOut *wire.TxOut) bool {
	// Ensure the PkScript is not empty and starts with OP_RETURN
	return len(txOut.PkScript) > 0 && txOut.PkScript[0] == txscript.OP_RETURN
}

func convertVoutToTxOut(vout btcjson.Vout) (*wire.TxOut, error) {
	// Decode the ScriptPubKey hex string into bytes
	scriptPubKeyBytes, err := hex.DecodeString(vout.ScriptPubKey.Hex)
	if err != nil {
		return nil, err
	}

	// Create wire.TxOut
	txOut := &wire.TxOut{
		Value:    int64(vout.Value * 1e8), // Convert BTC to satoshis
		PkScript: scriptPubKeyBytes,
	}

	return txOut, nil
}

func parseOpReturnGoatMagic(data []byte, magicBytes []byte) (common.Address, error) {
	// Ensure the data is long enough to contain the magic bytes
	if len(data) < len(magicBytes)+1 {
		return common.Address{}, fmt.Errorf("data is too short, expected at least %d bytes, got %d", len(magicBytes), len(data))
	}

	dataLen := uint32(data[0])
	if dataLen != 24 {
		return common.Address{}, fmt.Errorf("data length is not expected 24, got %d", dataLen)
	}

	data = data[1:]
	// Check if the data starts with GOAT_MAGIC_BYTES
	if !bytes.HasPrefix(data, magicBytes) {
		return common.Address{}, errors.New("data does not start with magic bytes")
	}

	log.Debugf("Parsed OP_RETURN as GTT0: %v", data)
	remainingBytes := data[len(magicBytes):]
	// Check if the remaining bytes match the expected EVM address length (20 bytes)
	if len(remainingBytes) != 20 {
		return common.Address{}, fmt.Errorf("invalid data length for EVM address, expected 20 bytes, got %d", len(remainingBytes))
	}

	evmAddr := common.BytesToAddress(remainingBytes)
	log.Debugf("Parsed OP_RETURN EVM address: %s", evmAddr.Hex())

	return evmAddr, nil
}

func IsUtxoGoatDepositV1(tx *wire.MsgTx, tssAddress []btcutil.Address, net *chaincfg.Params, minDepositAmount int64, magicBytes []byte) (isTrue bool, evmAddr string, outIdxToAmount map[int]int64) {
	// Ensure there are at least 2 outputs, one of them is OP_RETURN
	outIdxToAmount = make(map[int]int64)

	if len(tx.TxOut) < 2 {
		return false, "", outIdxToAmount
	}
	// Check if tx.TxOut[1] is OP_RETURN and tx.TxOut[0] is not OP_RETURN
	if !isOpReturn(tx.TxOut[1]) || isOpReturn(tx.TxOut[0]) {
		return false, "", outIdxToAmount
	}
	// Extract addresses from tx.TxOut[0]
	_, addresses, requireSigs, err := txscript.ExtractPkScriptAddrs(tx.TxOut[0].PkScript, net)
	if err != nil || addresses == nil || requireSigs > 1 {
		log.Debugf("Cannot extract PkScript addresses from TxOut[0]: %v", err)
		return false, "", outIdxToAmount
	}
	// Check if any of the addresses match tssAddress
	for _, address := range tssAddress {
		if address.EncodeAddress() == addresses[0].EncodeAddress() && tx.TxOut[0].Value >= minDepositAmount {
			// check if tx.TxOut[1] OP_RETURN rule: https://www.goat.network/docs/deposit/v1
			// Process OP_RETURN to extract EVM address
			data := tx.TxOut[1].PkScript[1:] // Assuming OP_RETURN opcode is at index 0

			evmAddr, err := parseOpReturnGoatMagic(data, magicBytes)
			if err != nil {
				log.Debugf("Cannot parse OP_RETURN in TxOut[1]: %v", err)
				return false, "", outIdxToAmount
			}

			outIdxToAmount[0] = tx.TxOut[0].Value

			return true, evmAddr.Hex(), outIdxToAmount
		}
	}

	return false, "", outIdxToAmount
}

func IsUtxoGoatDepositV0(tx *wire.MsgTx, tssAddress []btcutil.Address, net *chaincfg.Params, minDepositAmount int64) (isTrue bool, outIdxToAmount map[int]int64) {
	// Ensure there are at least 1 output
	outIdxToAmount = make(map[int]int64)

	if len(tx.TxOut) < 1 {
		return false, outIdxToAmount
	}

	// Extract addresses from tx.TxOut[0]
	for idx, txOut := range tx.TxOut {
		if isOpReturn(txOut) {
			continue
		}

		_, addresses, requireSigs, err := txscript.ExtractPkScriptAddrs(txOut.PkScript, net)
		if err != nil || addresses == nil || requireSigs > 1 {
			log.Debugf("Cannot extract PkScript addresses from TxOut[0]: %v", err)
			continue
		}

		for _, addr := range tssAddress {
			if addr.EncodeAddress() == addresses[0].EncodeAddress() && txOut.Value >= minDepositAmount {
				outIdxToAmount[idx] = txOut.Value
			}
		}
	}

	if len(outIdxToAmount) == 0 {
		return false, outIdxToAmount
	}

	return true, outIdxToAmount
}

func IsUtxoGoatDepositV0Json(tx *btcjson.TxRawResult, tssAddress []btcutil.Address, net *chaincfg.Params) (isV0 bool, outputIndex int, amount int64, pkScript []byte) {
	// Ensure there are at least 1 output
	if len(tx.Vout) < 1 {
		return false, -1, 0, nil
	}

	// Extract addresses from tx.TxOut[0]
	for idx, vout := range tx.Vout {
		txOut, err := convertVoutToTxOut(vout)
		if err != nil {
			log.Debugf("Cannot convert Vout to TxOut: %v", err)
			continue
		}

		if isOpReturn(txOut) {
			continue
		}

		_, addresses, requireSigs, err := txscript.ExtractPkScriptAddrs(txOut.PkScript, net)
		if err != nil || addresses == nil || requireSigs > 1 {
			log.Debugf("Cannot extract PkScript addresses from TxOut[0]: %v", err)
			continue
		}

		for _, address := range tssAddress {
			if address.EncodeAddress() == addresses[0].EncodeAddress() {
				return true, idx, txOut.Value, txOut.PkScript
			}
		}
	}

	return false, -1, 0, nil
}

func GetDustAmount(txPrice int64) int64 {
	return txPrice * 31 * 3
}

func GetAddressType(addressStr string, net *chaincfg.Params) (string, error) {
	address, err := btcutil.DecodeAddress(addressStr, net)
	if err != nil {
		return "", fmt.Errorf("invalid Bitcoin address: %w", err)
	}

	switch address.(type) {
	case *btcutil.AddressPubKeyHash:
		return WALLET_TYPE_P2PKH, nil
	case *btcutil.AddressScriptHash:
		return WALLET_TYPE_P2SH, nil
	case *btcutil.AddressWitnessPubKeyHash:
		return WALLET_TYPE_P2WPKH, nil
	case *btcutil.AddressWitnessScriptHash:
		return WALLET_TYPE_P2WSH, nil
	case *btcutil.AddressTaproot:
		return WALLET_TYPE_P2TR, nil
	default:
		return "", nil
	}
}

// TransactionSizeEstimate estimates the size of a transaction in bytes.
func TransactionSizeEstimate(numInputs int, receiverTypes []string, numOutputs int, utxoTypes []string) int64 {
	var totalSize int64 = 10 // Base transaction size (version, locktime, etc.)

	// Add inputs size
	for _, utxoType := range utxoTypes {
		switch utxoType {
		case WALLET_TYPE_P2WPKH:
			totalSize += 41 // P2WPKH input size without witness (32 bytes txid + 4 bytes vout + 1 byte script length + 4 bytes sequence)
		case WALLET_TYPE_P2PKH:
			totalSize += 148 // P2PKH input size
		case WALLET_TYPE_P2WSH:
			totalSize += 41 // P2WSH input size without witness (32 bytes txid + 4 bytes vout + 1 byte script length + 4 bytes sequence)
		case WALLET_TYPE_P2SH:
			totalSize += 296 // P2SH input size
		case WALLET_TYPE_P2TR:
			totalSize += 41 // P2TR input size without witness (32 bytes txid + 4 bytes vout + 1 byte script length + 4 bytes sequence)
		}
	}

	// Each output (P2PKH: 34 bytes, P2WPKH: 31 bytes, P2SH: 32 bytes, P2WSH: 43 bytes, P2TR: 42 bytes)
	for _, receiverType := range receiverTypes {
		switch receiverType {
		case WALLET_TYPE_P2PKH:
			totalSize += 34
		case WALLET_TYPE_P2WPKH:
			totalSize += 31
		case WALLET_TYPE_P2SH:
			totalSize += 32
		case WALLET_TYPE_P2WSH:
			totalSize += 43
		case WALLET_TYPE_P2TR:
			totalSize += 42
		}
	}

	if len(receiverTypes) < numOutputs {
		// change output as P2WPKH
		totalSize += int64(31 * (numOutputs - len(receiverTypes)))
	}

	return totalSize
}

// Deserialize transaction.
func DeserializeTransaction(data []byte) (*wire.MsgTx, error) {
	var tx wire.MsgTx

	buf := bytes.NewReader(data)

	err := tx.Deserialize(buf)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

// Serialize transaction to bytes (with witness data).
func SerializeTransaction(tx *wire.MsgTx) ([]byte, error) {
	var buf bytes.Buffer

	err := tx.Serialize(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Serialize transaction to bytes (without witness data).
func SerializeTransactionNoWitness(tx *wire.MsgTx) ([]byte, error) {
	var buf bytes.Buffer

	err := tx.SerializeNoWitness(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func ConvertTxRawResultToMsgTx(txResult *btcjson.TxRawResult) (*wire.MsgTx, error) {
	// Decode the hex-encoded transaction
	txBytes, err := hex.DecodeString(txResult.Hex)
	if err != nil {
		return nil, err
	}

	// Deserialize the transaction
	msgTx := wire.NewMsgTx(wire.TxVersion)

	err = msgTx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		return nil, err
	}

	return msgTx, nil
}

func IsTargetP2PKHAddress(script []byte, targetAddress btcutil.Address, net *chaincfg.Params) bool {
	addressHash, err := btcutil.NewAddressPubKeyHash(script[3:23], net)
	if err != nil {
		return false
	}

	return addressHash.EncodeAddress() == targetAddress.EncodeAddress()
}

func IsTargetP2WPKHAddress(script []byte, targetAddress btcutil.Address, net *chaincfg.Params) bool {
	// P2WPKH is 22 bytes (0x00 + 0x14 + 20 hash)
	if len(script) != 22 || script[0] != 0x00 || script[1] != 0x14 {
		return false
	}

	pubKeyHash := script[2:22]

	address, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, net)
	if err != nil {
		return false
	}

	return address.EncodeAddress() == targetAddress.EncodeAddress()
}

func IsP2WSHAddress(script []byte, net *chaincfg.Params) (bool, string) {
	// P2WSH is 34 bytes (0x00 + 0x20 + 32 hash)
	if len(script) != 34 || script[0] != 0x00 || script[1] != 0x20 {
		return false, ""
	}

	witnessHash := script[2:34]

	address, err := btcutil.NewAddressWitnessScriptHash(witnessHash, net)
	if err != nil {
		return false, ""
	}

	return true, address.EncodeAddress()
}

func GenerateP2PKHAddress(pubKey []byte, net *chaincfg.Params) (*btcutil.AddressPubKeyHash, error) {
	pubKeyHash := btcutil.Hash160(pubKey)

	address, err := btcutil.NewAddressPubKeyHash(pubKeyHash, net)
	if err != nil {
		log.Errorf("Error generating P2PKH address: %v", err)
		return nil, err
	}

	return address, nil
}

func GenerateP2WPKHAddress(pubKey []byte, net *chaincfg.Params) (*btcutil.AddressWitnessPubKeyHash, error) {
	pubKeyHash := btcutil.Hash160(pubKey)

	address, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, net)
	if err != nil {
		log.Errorf("Error generating P2WPKH address: %v", err)
		return nil, err
	}

	return address, nil
}

func GenerateV0P2WSHAddress(pubKey []byte, evmAddress string, net *chaincfg.Params) (*btcutil.AddressWitnessScriptHash, error) {
	subScript, err := BuildSubScriptForP2WSH(evmAddress, pubKey)
	if err != nil {
		return nil, err
	}

	witnessProg := sha256.Sum256(subScript)

	p2wsh, err := btcutil.NewAddressWitnessScriptHash(witnessProg[:], net)
	if err != nil {
		return nil, fmt.Errorf("failed to create v0 p2wsh address: %w", err)
	}

	return p2wsh, nil
}

func GenerateSPVProof(txHash string, txHashes []string) ([]byte, []byte, int, error) {
	// Find the transaction's position in the block
	txIndex := -1

	for i, hash := range txHashes {
		if hash == txHash {
			txIndex = i
			break
		}
	}

	if txIndex == -1 {
		return nil, nil, -1, fmt.Errorf("transaction hash not found in block, expected txid: %s, found txhashes: %s", txHash, txHashes)
	}

	// Generate merkle root and proof
	txHashesPtrs := make([]*chainhash.Hash, len(txHashes))

	for i, hashStr := range txHashes {
		hash, err := chainhash.NewHashFromStr(hashStr)
		if err != nil {
			return nil, nil, -1, fmt.Errorf("failed to parse transaction hash: %w", err)
		}

		txHashesPtrs[i] = hash
	}

	var proof []*chainhash.Hash
	merkleRoot := ComputeMerkleRootAndProof(txHashesPtrs, txIndex, &proof)

	// Serialize immediate proof
	var buf bytes.Buffer
	for _, p := range proof {
		buf.Write(p[:])
	}

	return merkleRoot.CloneBytes(), buf.Bytes(), txIndex, nil
}

func VerifyBlockSPV(btcBlock BtcBlockExt) error {
	// get merkle root from header
	expectedMerkleRoot := btcBlock.Header.MerkleRoot

	// generate actual merkle root from transactions
	txHashes := make([]*chainhash.Hash, 0, len(btcBlock.Transactions))

	for _, tx := range btcBlock.Transactions {
		txHash := tx.TxHash()
		txHashes = append(txHashes, &txHash)
	}

	actualMerkleRoot := buildMerkleRoot(txHashes)

	// check merkle root is match
	if !actualMerkleRoot.IsEqual(&expectedMerkleRoot) {
		return fmt.Errorf("merkle root mismatch: expected %s, got %s",
			expectedMerkleRoot, actualMerkleRoot)
	}

	// check header hash is match
	headerHash := btcBlock.Header.BlockHash()
	blockHash := btcBlock.BlockHash()

	if !headerHash.IsEqual(&blockHash) {
		return fmt.Errorf("block hash mismatch: expected %s, got %s", blockHash, headerHash)
	}

	log.Infof("Block %d SPV verification successful: Merkle root and block hash match", btcBlock.BlockNumber)

	return nil
}

// buildMerkleRoot builds the Merkle tree and returns the root hash.
func buildMerkleRoot(txHashes []*chainhash.Hash) *chainhash.Hash {
	if len(txHashes) == 0 {
		return nil
	}

	// Merkle Root calculation loop
	for len(txHashes) > 1 {
		var newLevel []*chainhash.Hash

		// combine hashes two by two
		for i := 0; i < len(txHashes); i += 2 {
			var combined []byte
			if i+1 < len(txHashes) {
				// normal case: combine two by two
				combined = append(txHashes[i][:], txHashes[i+1][:]...)
			} else {
				// odd case: duplicate the last transaction hash
				combined = append(txHashes[i][:], txHashes[i][:]...)
			}

			newHash := chainhash.DoubleHashH(combined)
			newLevel = append(newLevel, &newHash)
		}

		// prepare for the next level
		txHashes = newLevel
	}

	// return the final root hash
	return txHashes[0]
}

func SerializeNoWitnessTx(rawTransaction []byte) ([]byte, error) {
	// Parse the raw transaction
	rawTx := wire.NewMsgTx(wire.TxVersion)

	err := rawTx.Deserialize(bytes.NewReader(rawTransaction))
	if err != nil {
		return nil, fmt.Errorf("failed to parse raw transaction: %w", err)
	}

	// Create a new transaction without witness data
	noWitnessTx := wire.NewMsgTx(rawTx.Version)

	// Copy transaction inputs, excluding witness data
	for _, txIn := range rawTx.TxIn {
		newTxIn := wire.NewTxIn(&txIn.PreviousOutPoint, nil, nil)
		newTxIn.Sequence = txIn.Sequence
		noWitnessTx.AddTxIn(newTxIn)
	}

	// Copy transaction outputs
	for _, txOut := range rawTx.TxOut {
		noWitnessTx.AddTxOut(txOut)
	}

	// Set lock time
	noWitnessTx.LockTime = rawTx.LockTime

	// Serialize the transaction without witness data
	var buf bytes.Buffer

	err = noWitnessTx.Serialize(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize transaction without witness data: %w", err)
	}

	return buf.Bytes(), nil
}

func BuildSubScriptForP2WSH(evmAddress string, pubKey []byte) ([]byte, error) {
	posPubkey, err := btcec.ParsePubKey(pubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	evmAddress = strings.TrimPrefix(evmAddress, "0x")

	addr, err := hex.DecodeString(evmAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to decode evmAddress: %w", err)
	}

	subScript, err := txscript.NewScriptBuilder().
		AddData(addr).
		AddOp(txscript.OP_DROP).
		AddData(posPubkey.SerializeCompressed()).
		AddOp(txscript.OP_CHECKSIG).Script()
	if err != nil {
		return nil, fmt.Errorf("failed to build subscript: %w", err)
	}

	return subScript, nil
}
