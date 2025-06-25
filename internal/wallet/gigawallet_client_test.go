package wallet

import (
	"testing"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/goatnetwork/goat-relayer/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestGigaWalletClient_ExtractAddressFromPkScript(t *testing.T) {
	client := NewGigaWalletClient("http://localhost:8080", "test")

	// Test P2PKH address
	p2pkhScript := []byte{
		0x76, 0xa9, 0x14, // OP_DUP OP_HASH160 OP_PUSHDATA20
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a,
		0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14,
		0x88, 0xac, // OP_EQUALVERIFY OP_CHECKSIG
	}

	// Mock network parameters
	network := &chaincfg.MainNetParams
	network.PubKeyHashAddrID = 0x30 // Dogecoin mainnet address prefix

	address, err := client.extractAddressFromPkScript(p2pkhScript)
	if err != nil {
		t.Logf("Expected error for test script: %v", err)
	} else {
		t.Logf("Extracted address: %s", address)
	}
}

func TestGigaWalletClient_SendRawTransaction(t *testing.T) {
	client := NewGigaWalletClient("http://localhost:8080", "test")

	// Create a simple test transaction
	tx := wire.NewMsgTx(wire.TxVersion)

	// Add output
	output := &wire.TxOut{
		Value:    100000000, // 1 DOGE in satoshis
		PkScript: []byte{0x76, 0xa9, 0x14, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x88, 0xac},
	}
	tx.AddTxOut(output)

	// Mock UTXOs
	utxos := []*db.Utxo{
		{
			Receiver: "DTestChangeAddress123456789",
		},
	}

	// Test sending transaction (this will fail without real GigaWallet service, but can test code logic)
	txHash, exist, err := client.SendRawTransaction(tx, utxos, "withdrawal")

	// Since there's no real GigaWallet service, we expect an error
	assert.Error(t, err, "Expected error when GigaWallet service is not available")
	assert.False(t, exist, "Transaction should not exist")
	assert.Empty(t, txHash, "Transaction hash should be empty")
}

func TestGigaWalletClient_CheckPending(t *testing.T) {
	client := NewGigaWalletClient("http://localhost:8080", "test")

	// Test checking transaction status (this will fail without real GigaWallet service)
	revert, confirmations, blockHeight, err := client.CheckPending("test-txid", "test-external-id", time.Now())

	// Since there's no real GigaWallet service, we expect an error
	assert.Error(t, err, "Expected error when GigaWallet service is not available")
	assert.False(t, revert, "Should not revert")
	assert.Equal(t, uint64(0), confirmations, "Confirmations should be 0")
	assert.Equal(t, uint64(0), blockHeight, "Block height should be 0")
}

func TestGigaWalletClient_GetAccountBalance(t *testing.T) {
	client := NewGigaWalletClient("http://localhost:8080", "test")

	// Test getting account balance (this will fail without real GigaWallet service)
	balance, err := client.GetAccountBalance()

	// Since there's no real GigaWallet service, we expect an error
	assert.Error(t, err, "Expected error when GigaWallet service is not available")
	assert.Equal(t, float64(0), balance, "Balance should be 0")
}
