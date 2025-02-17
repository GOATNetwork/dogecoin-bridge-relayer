package btc

import (
	"errors"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/nuvosphere/nudex-voter/internal/state"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetBTCNetwork(networkType string) *chaincfg.Params {
	switch networkType {
	case "devnet":
		return &chaincfg.RegressionNetParams
	case "testnet":
		return &chaincfg.TestNet3Params
	default:
		return &chaincfg.MainNetParams
	}
}

func (m *WalletClient) ScanBlocks() {
	m.loadState()
	log.Infof("Resuming BTC scan from block %d...", m.lastScannedBlock)

	started := false

	for {
		select {
		case <-m.ctx.Done():
			log.Infof("Stopping BTC block scanning.")
			return
		default:
			if started {
				time.Sleep(m.chainInfo.ScanInterval * time.Second)
			} else {
				started = true
			}
			// Fetch the latest block height
			latestBlock, err := m.client.client.GetBlockCount()
			if err != nil {
				log.Errorf("Failed to get latest block count: %v", err)
				time.Sleep(10 * time.Second)

				continue
			}

			latestBlock = latestBlock - int64(m.chainInfo.Confirmations)
			log.Infof("BTC latest confirmed block: %d", latestBlock)

			fromBlock := m.lastScannedBlock + 1
			if fromBlock > uint64(latestBlock) {
				log.Debugf("BTC no new blocks to scan. Latest block: %d", latestBlock)
				continue
			}

			rangeSize := m.chainInfo.MaxBlockRange
			if rangeSize == 0 {
				rangeSize = 500
			}

			toBlock := fromBlock + uint64(rangeSize) - 1
			if toBlock > uint64(latestBlock) {
				toBlock = uint64(latestBlock)
			}

			log.Infof("Scanning BTC blocks from %d to %d", fromBlock, toBlock)

			for fromBlock <= toBlock {
				blockHash, err := m.client.client.GetBlockHash(int64(fromBlock))
				if err != nil {
					log.Errorf("Failed to get block hash %d: %v", fromBlock, err)
					break
				}

				block, err := m.client.client.GetBlock(blockHash)
				if err != nil {
					log.Errorf("Failed to get block %d: %v", fromBlock, err)
					break
				}

				log.Infof("Scanning BTC block %d...", fromBlock)

				err = m.scanBlock(fromBlock, block)
				if err != nil {
					log.Errorf("Failed to scan BTC block %d: %v", fromBlock, err)
					break
				}

				// Save the state after scanning each block
				err = m.saveState(fromBlock)
				if err != nil {
					log.Errorf("Failed to save state: %v", err)
					break
				}

				log.Infof("BTC scanned up to block %d", fromBlock)

				m.lastScannedBlock = fromBlock
				fromBlock++

				time.Sleep(100 * time.Millisecond)
			}

			// Wait before checking for new blocks
			time.Sleep(2 * time.Second)
		}
	}
}

func (m *WalletClient) loadState() {
	latestSyncHeight, err := state.GetScannedNumber(m.state.TX(nil), m.chainInfo.ChainId, uint64(m.chainInfo.StartHeight))
	if err != nil {
		log.Fatalf("Failed to fetch scan info: %v", err)
	}

	lastScannedBlock := latestSyncHeight
	if m.chainInfo.StartHeight > 0 && lastScannedBlock < uint64(m.chainInfo.StartHeight)-1 {
		lastScannedBlock = uint64(m.chainInfo.StartHeight) - 1
	}

	m.lastScannedBlock = lastScannedBlock
}

func (m *WalletClient) saveState(blockNumber uint64) error {
	err := state.UpdateLatestScannedHeight(m.state.TX(nil), m.chainInfo.ChainId, blockNumber)
	if err != nil {
		return fmt.Errorf("failed to save btc scan info: %w", err)
	}

	m.lastScannedBlock = blockNumber

	return nil
}

func (w *WalletClient) IsExist(address string) bool {
	_, err := w.ContractState().Account(address) // todo
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (w *WalletClient) scanBlock(blockNumber uint64, block *wire.MsgBlock) error {
	log.Infof("Scanning BTC block %d, %d transactions...", blockNumber, len(block.Transactions))

	network := GetBTCNetwork(w.chainInfo.Network)

	for txIndex, tx := range block.Transactions {
		log.Debugf("BTC tx: %s", tx.TxHash().String())
		txId := tx.TxHash().String()
		sender, receiver := "", ""
		_ = txId

		for _, vin := range tx.TxIn {
			// check utxo spent
			scriptType, addresses, requireSigs, err := txscript.ExtractPkScriptAddrs(vin.SignatureScript, network)
			if err != nil {
				log.Warnf("Error extracting input address, %v", err)
				continue
			}

			if len(addresses) == 0 {
				// ignore coinbase or other tx without address
				continue
			}

			if requireSigs > 1 {
				// ignore multi sigs
				continue
			}

			if scriptType != txscript.WitnessV0PubKeyHashTy {
				// only accept p2wphk
				continue
			}

			sender = addresses[0].EncodeAddress()

			is := w.IsExist(sender)
			if is {
				// todo
				// wallet address spent utxo
				// update utxo spent
				log.Debug("spend tx hash", txId, "tx index:", txIndex)

				err := w.state.SpentUTXO(vin.PreviousOutPoint.Hash.String(), vin.PreviousOutPoint.Index, blockNumber)
				if err != nil {
					return err
				}
			}
		}

		for txIndex, vout := range tx.TxOut {
			// check new deposit
			scriptType, addresses, requireSigs, err := txscript.ExtractPkScriptAddrs(vout.PkScript, network)
			if err != nil {
				log.Warnf("Error extracting output address, %v", err)
				continue
			}

			if len(addresses) == 0 {
				// ignore coinbase or other tx without address
				continue
			}

			if requireSigs > 1 {
				// ignore multi sigs
				continue
			}

			if scriptType != txscript.WitnessV0PubKeyHashTy {
				// only accept p2wphk
				continue
			}

			receiver = addresses[0].EncodeAddress()

			is := w.IsExist(receiver)
			if is {
				// todo
				// save tx Hash and vout txIndex
				// add utxo
				log.Debug("receiver utxo tx hash", txId, "tx index:", txIndex)

				err := w.state.AddNewUTXO(txId, uint32(txIndex), vout.Value, sender, receiver, blockNumber)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
