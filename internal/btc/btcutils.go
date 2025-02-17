package btc

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/nuvosphere/nudex-voter/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

func GenerateSPVProof(txHash string, txHashes []string) ([]byte, []byte, uint32, error) {
	// Find the transaction's position in the block
	var txIndex int

	for i, hash := range txHashes {
		if hash == txHash {
			txIndex = i
			break
		}
	}

	// Generate merkle root and proof
	txHashesPtrs := make([]*chainhash.Hash, len(txHashes))

	for i, hashStr := range txHashes {
		hash, err := chainhash.NewHashFromStr(hashStr)
		if err != nil {
			return nil, nil, 0, fmt.Errorf("failed to parse transaction hash: %w", err)
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

	return merkleRoot.CloneBytes(), buf.Bytes(), uint32(txIndex), nil
}

func GenerateSPVProofByTx(msgTx *wire.MsgTx) (string, error) {
	// Open or create the local storage
	dbPath := filepath.Join(config.AppConfig.DB.DbRootDir, "btc_cache.db")

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatalf("Failed to open local storage: %v", err)
	}

	defer db.Close()

	txHash := msgTx.TxHash()

	// Get block hash
	var blockHashBytes []byte

	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Next() {
		key := iter.Key()
		if bytes.HasPrefix(key, []byte("utxo:")) {
			value := iter.Value()
			if bytes.Equal(value, txHash[:]) {
				blockHashBytes = key[len("utxo:") : len("utxo:")+64]
				break
			}
		}
	}

	if blockHashBytes == nil {
		return "", fmt.Errorf("failed to find block hash for tx: %v", txHash)
	}

	blockHash, err := chainhash.NewHash(blockHashBytes)
	if err != nil {
		return "", fmt.Errorf("invalid block hash: %w", err)
	}

	// Get block header
	headerBytes, err := db.Get([]byte("header:"+blockHash.String()), nil)
	if err != nil {
		return "", fmt.Errorf("failed to get block header from db: %w", err)
	}

	var header wire.BlockHeader

	err = header.Deserialize(bytes.NewReader(headerBytes))
	if err != nil {
		return "", fmt.Errorf("failed to deserialize block header: %w", err)
	}

	// Get transaction hash list
	txHashesBytes, err := db.Get([]byte("txhashes:"+blockHash.String()), nil)
	if err != nil {
		return "", fmt.Errorf("failed to get tx hashes from db: %w", err)
	}

	var txHashes []chainhash.Hash

	err = json.Unmarshal(txHashesBytes, &txHashes)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal tx hashes: %w", err)
	}

	// Find the transaction's position in the block
	var txIndex int

	for i, hash := range txHashes {
		if hash == txHash {
			txIndex = i
			break
		}
	}

	// Generate Merkle proof
	txHashesPtrs := make([]*chainhash.Hash, len(txHashes))
	for i := range txHashes {
		txHashesPtrs[i] = &txHashes[i]
	}

	var proof []*chainhash.Hash
	merkleRoot := ComputeMerkleRootAndProof(txHashesPtrs, txIndex, &proof)

	// Serialize Merkle proof
	var buf bytes.Buffer

	buf.Write(txHash[:])

	for _, p := range proof {
		buf.Write(p[:])
	}

	buf.Write(merkleRoot[:])

	return hex.EncodeToString(buf.Bytes()), nil
}
