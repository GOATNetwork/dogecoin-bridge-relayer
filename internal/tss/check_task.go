package tss

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/shopspring/decimal"
)

type TxStatusResponse struct {
	Status string `json:"status"`
}

func (m *Scheduler) checkTask(task pool.Task[uint64]) (bool, int, error) {
	switch taskData := task.(type) {
	case *db.DepositTask:
		hashCheckStatus, err := checkTxStatus(taskData.TxHash)
		if err != nil {
			return false, -1, err
		}

		if hashCheckStatus != "success" {
			return true, db.TaskErrorCodeCheckTxFailed, err
		}

		inscriptionMintb, err := m.stateDB.GetInscriptionMintb(taskData.TxHash)
		if err != nil || inscriptionMintb == nil {
			return true, db.TaskErrorCodeCheckInscriptionFailed, err
		}

		if inscriptionMintb.Amount != taskData.Amount {
			return true, db.TaskErrorCodeCheckAmountFailed, err
		}

		asset, err := m.stateDB.GetAsset(inscriptionMintb.Ticker)
		if err != nil || asset == nil {
			return true, db.TaskErrorCodeCheckAssetFailed, err
		}

		if !asset.DepositEnabled {
			return true, db.TaskErrorCodeDepositAssetNotEnabled, err
		}

		if taskData.Amount.Cmp(decimal.NewFromUint64(asset.MinDepositAmount)) == -1 {
			return true, db.TaskErrorCodeDepositAmountTooLow, err
		}

		tokenInfo, err := m.stateDB.GetTokenInfo(inscriptionMintb.Ticker)
		if err != nil || tokenInfo == nil {
			return true, db.TaskErrorCodeDepositTokenNotSupported, err
		}

		if !tokenInfo.IsActive {
			return true, db.TaskErrorCodeDepositTokenNotActive, err
		}

		return true, db.TaskErrorCodeSuccess, nil
	case *db.WithdrawalTask:
		hashCheckStatus, err := checkTxStatus(taskData.TxHash)
		if err != nil {
			return false, -1, err
		}

		if hashCheckStatus != "success" {
			return true, db.TaskErrorCodeCheckTxFailed, err
		}

		inscriptionBurnb, err := m.stateDB.GetInscriptionBurnb(taskData.TxHash)
		if err != nil || inscriptionBurnb == nil {
			return true, db.TaskErrorCodeCheckInscriptionFailed, err
		}

		if inscriptionBurnb.Amount != taskData.Amount {
			return true, db.TaskErrorCodeCheckAmountFailed, err
		}

		asset, err := m.stateDB.GetAsset(inscriptionBurnb.Ticker)
		if err != nil || asset == nil {
			return true, db.TaskErrorCodeCheckAssetFailed, err
		}

		if !asset.WithdrawalEnabled {
			return true, db.TaskErrorCodeWithdrawalAssetNotEnabled, err
		}

		if taskData.Amount.Cmp(decimal.NewFromUint64(asset.MinWithdrawAmount)) == -1 {
			return true, db.TaskErrorCodeWithdrawalAmountTooLow, err
		}

		tokenInfo, err := m.stateDB.GetTokenInfo(inscriptionBurnb.Ticker)
		if err != nil || tokenInfo == nil {
			return true, db.TaskErrorCodeWithdrawalTokenNotSupported, err
		}

		if !tokenInfo.IsActive {
			return true, db.TaskErrorCodeWithdrawalTokenNotActive, err
		}

		assetType := tokenInfo.AssetType

		switch taskData.Chain {
		case types.CoinTypeBTC:
			switch assetType {
			case types.AssetTypeMain:

			case types.AssetTypeErc20:
			default:
				return true, db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", assetType)
			}
		case types.ChainEthereum:
			switch assetType {
			case types.AssetTypeMain:
				// @todo check main asset
			case types.AssetTypeErc20:
				// @todo check erc20 asset
			default:
				return true, db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", assetType)
			}
		case types.ChainSolana:
			switch assetType {
			case types.AssetTypeMain:
			case types.AssetTypeErc20:
			default:
				return true, db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", assetType)
			}
		case types.ChainSui:
			switch assetType {
			case types.AssetTypeMain:

			case types.AssetTypeErc20:
			default:
				return true, db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", assetType)
			}
		default:
			return true, db.TaskErrorCodeChainNotSupported, fmt.Errorf("unknown Chain type: %v", taskData.Chain)
		}
	default:
		panic(fmt.Errorf("error pending task id: %v", task.TaskID()))
	}

	return false, db.TaskErrorCodePending, nil
}

func checkTxStatus(txHash string) (string, error) {
	url := fmt.Sprintf("https://nip-api.testnet.nudex.io/%s/status", txHash)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var txStatus TxStatusResponse
	if err := json.Unmarshal(body, &txStatus); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	return txStatus.Status, nil
}
