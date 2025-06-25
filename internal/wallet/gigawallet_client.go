package wallet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/goatnetwork/goat-relayer/internal/config"
	"github.com/goatnetwork/goat-relayer/internal/db"
	"github.com/goatnetwork/goat-relayer/internal/types"
	log "github.com/sirupsen/logrus"
)

// GigaWalletClient GigaWallet client
type GigaWalletClient struct {
	baseURL    string
	accountID  string
	httpClient *http.Client
}

// PayoutRequest GigaWallet payout request structure
type PayoutRequest struct {
	Address string  `json:"address"` // Target Dogecoin address
	Amount  float64 `json:"amount"`  // Amount in DOGE (floating point)
	Label   string  `json:"label,omitempty"`
}

// PayoutResponse GigaWallet payout response structure
type PayoutResponse struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	Transaction string `json:"transaction,omitempty"`
	Error       string `json:"error,omitempty"`
}

// TransactionStatus GigaWallet transaction status response
type TransactionStatus struct {
	ID            string  `json:"id"`
	Status        string  `json:"status"`
	TransactionID string  `json:"transaction_id,omitempty"`
	Confirmations int     `json:"confirmations,omitempty"`
	BlockHeight   int64   `json:"block_height,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	Fee           float64 `json:"fee,omitempty"`
}

// NewGigaWalletClient Create new GigaWallet client
func NewGigaWalletClient(baseURL, accountID string) *GigaWalletClient {
	return &GigaWalletClient{
		baseURL:   baseURL,
		accountID: accountID,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SendRawTransaction Send transaction through GigaWallet
// This method converts BTC transactions to GigaWallet payout requests
func (c *GigaWalletClient) SendRawTransaction(tx *wire.MsgTx, utxos []*db.Utxo, orderType string) (txHash string, exist bool, err error) {
	// Extract target address and amount from transaction outputs
	// Find non-change address outputs
	targetAddress := ""
	targetAmount := float64(0)

	for _, txOut := range tx.TxOut {
		// Skip OP_RETURN outputs
		if len(txOut.PkScript) > 0 && txOut.PkScript[0] == txscript.OP_RETURN {
			continue
		}

		// Try to extract address
		address, err := c.extractAddressFromPkScript(txOut.PkScript)
		if err != nil {
			log.Errorf("Failed to extract address from PkScript: %v", err)
			continue
		}

		// Check if it's a change address (usually change address is the sender's address)
		isChangeAddress := false
		for _, utxo := range utxos {
			if address == utxo.Receiver {
				isChangeAddress = true
				break
			}
		}

		// If it's not a change address, consider it as target address
		if !isChangeAddress {
			targetAddress = address
			targetAmount = float64(txOut.Value) / 100000000.0 // Convert to DOGE (assuming 8 decimal places)
			break
		}
	}

	if targetAddress == "" {
		return "", false, fmt.Errorf("no valid target address found in transaction")
	}

	// Construct payout request
	payout := PayoutRequest{
		Address: targetAddress,
		Amount:  targetAmount,
		Label:   fmt.Sprintf("%s-%s", orderType, tx.TxHash().String()),
	}

	// Send request to GigaWallet
	txHash, err = c.sendPayoutRequest(payout)
	if err != nil {
		return "", false, err
	}

	return txHash, false, nil
}

// CheckPending Check transaction status
func (c *GigaWalletClient) CheckPending(txid string, externalTxId string, updatedAt time.Time) (revert bool, confirmations uint64, blockHeight uint64, err error) {
	// Use externalTxId (GigaWallet transaction ID) to query status
	status, err := c.getTransactionStatus(externalTxId)
	if err != nil {
		return false, 0, 0, err
	}

	// Check transaction status
	switch status.Status {
	case "completed":
		return false, uint64(status.Confirmations), uint64(status.BlockHeight), nil
	case "failed", "cancelled", "rejected":
		return true, 0, 0, nil
	case "pending":
		return false, 0, 0, nil
	default:
		return false, 0, 0, fmt.Errorf("unknown transaction status: %s", status.Status)
	}
}

// sendPayoutRequest Send payout request to GigaWallet
func (c *GigaWalletClient) sendPayoutRequest(payout PayoutRequest) (string, error) {
	payload, err := json.Marshal(payout)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payout request: %v", err)
	}

	url := fmt.Sprintf("%s/account/%s/payout", c.baseURL, c.accountID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GigaWallet API error: %s, status: %d", string(body), resp.StatusCode)
	}

	var payoutResp PayoutResponse
	if err := json.Unmarshal(body, &payoutResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if payoutResp.Error != "" {
		return "", fmt.Errorf("GigaWallet payout error: %s", payoutResp.Error)
	}

	return payoutResp.ID, nil
}

// getTransactionStatus Get transaction status
func (c *GigaWalletClient) getTransactionStatus(txID string) (*TransactionStatus, error) {
	url := fmt.Sprintf("%s/account/%s/transaction/%s", c.baseURL, c.accountID, txID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GigaWallet API error: %s, status: %d", string(body), resp.StatusCode)
	}

	var status TransactionStatus
	if err := json.Unmarshal(body, &status); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &status, nil
}

// extractAddressFromPkScript Extract address from PkScript
func (c *GigaWalletClient) extractAddressFromPkScript(pkScript []byte) (string, error) {
	// Use btcd's address parsing functionality
	network := types.GetBTCNetwork(config.AppConfig.BTCNetworkType)

	// Try to parse address
	_, addresses, _, err := txscript.ExtractPkScriptAddrs(pkScript, network)
	if err != nil {
		return "", fmt.Errorf("failed to extract address from PkScript: %v", err)
	}

	if len(addresses) == 0 {
		return "", fmt.Errorf("no addresses found in PkScript")
	}

	// Return the first address
	return addresses[0].EncodeAddress(), nil
}

// GetAccountBalance Get account balance
func (c *GigaWalletClient) GetAccountBalance() (float64, error) {
	url := fmt.Sprintf("%s/account/%s/balance", c.baseURL, c.accountID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("GigaWallet API error: %s, status: %d", string(body), resp.StatusCode)
	}

	var balance struct {
		Balance float64 `json:"balance"`
	}
	if err := json.Unmarshal(body, &balance); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return balance.Balance, nil
}
