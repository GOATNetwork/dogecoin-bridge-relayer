package btc

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

func SpentP2wsh(netwk *chaincfg.Params, tssGroupKey *btcec.PrivateKey, evmAddresses [][]byte, prevTxIds []string,
	prevTxouts []int, amounts []int64, fee int64,
) (string, error) {
	if len(prevTxIds) != len(prevTxouts) || len(prevTxouts) != len(amounts) || len(evmAddresses) != len(prevTxIds) {
		return "", fmt.Errorf("mismatched input lengths")
	}

	newtx := wire.NewMsgTx(2)
	totalInputAmount := int64(0)

	for _, amount := range amounts {
		totalInputAmount += amount
	}

	tssGroupAddress, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(tssGroupKey.PubKey().SerializeCompressed()), netwk)
	if err != nil {
		return "", fmt.Errorf("failed to build tssGroupAddress: %w", err)
	}

	outputAddress, err := txscript.PayToAddrScript(tssGroupAddress)
	if err != nil {
		return "", fmt.Errorf("failed to build outputAddress: %w", err)
	}

	txout := wire.NewTxOut(totalInputAmount-fee, outputAddress)
	newtx.AddTxOut(txout)

	for i, prevTxout := range prevTxouts {
		prevTxid, err := chainhash.NewHashFromStr(prevTxIds[i])
		if err != nil {
			return "", fmt.Errorf("failed to build prevTxid: %w", err)
		}

		txin := wire.NewTxIn(wire.NewOutPoint(prevTxid, uint32(prevTxout)), nil, nil)
		newtx.AddTxIn(txin)
	}

	for i, evmAddress := range evmAddresses {
		subScript, err := txscript.NewScriptBuilder().
			AddData(evmAddress).
			AddOp(txscript.OP_DROP).
			AddData(tssGroupKey.PubKey().SerializeCompressed()).
			AddOp(txscript.OP_CHECKSIG).Script()
		if err != nil {
			return "", fmt.Errorf("failed to build subScript: %w", err)
		}

		witnessProg := sha256.Sum256(subScript)

		prevPkScript, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(witnessProg[:]).Script()
		if err != nil {
			return "", fmt.Errorf("failed to build prevPkScript: %w", err)
		}

		sigHashes := txscript.NewTxSigHashes(newtx,
			txscript.NewCannedPrevOutputFetcher(prevPkScript, amounts[i]))

		witSig, err := txscript.RawTxInWitnessSignature(
			newtx, sigHashes, i,
			amounts[i], subScript,
			txscript.SigHashAll, tssGroupKey,
		)
		if err != nil {
			return "", fmt.Errorf("failed to build witSig: %w", err)
		}

		newtx.TxIn[i].Witness = wire.TxWitness{witSig, subScript}
	}

	txbuf := bytes.NewBuffer(nil)
	if err := newtx.Serialize(txbuf); err != nil {
		return "", fmt.Errorf("failed to build txbuf: %w", err)
	}

	return hex.EncodeToString(txbuf.Bytes()), nil
}

func (w *WalletClient) ProcessConsolidation(task *db.Consolidation) {
	hotAddress := address.HotAddressOfBtc(w.tss.ECPoint(w.ChainType()))
	log.Infof("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, task.FromAddress, task.Amount)

	txClient := NewTxClient(w.ctx, time.Minute*5, &chaincfg.MainNetParams, w.tss, w.chainInfo)
	txClient.AddTask(task)

	err := txClient.BuildTx(task.FromAddress, hotAddress, task.Amount.BigInt().Int64())
	if err != nil {
		log.Errorf("Build btc Tx err: %v", err)
		return
	}

	err = txClient.Sign()
	if err != nil {
		log.Errorf("Sign btc Tx err: %v", err)
		return
	}

	w.txContext.Store(txClient.TxHash(), txClient)
}
