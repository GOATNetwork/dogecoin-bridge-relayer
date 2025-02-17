package btc

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

const precision = 1e8

type RawSignTx struct{}

type UnSignTx struct{}

type SignedTx struct{}

type SignatureCtx struct {
	Hash, Signature []byte
}

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
	client *rpcclient.Client
	params *chaincfg.Params
}

func NewClient(
	ctx context.Context,
	timeout time.Duration,
	params *chaincfg.Params,
	chainInfo *config.Chain,
) *Client {
	connConfig := &rpcclient.ConnConfig{
		Host:         chainInfo.Rpc.Url,
		User:         chainInfo.Rpc.User,
		Pass:         chainInfo.Rpc.Password,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connConfig, nil)
	utils.Assert(err)

	ctx, cancel := context.WithTimeout(ctx, timeout)

	return &Client{
		ctx:    ctx,
		cancel: cancel,
		client: client,
		params: params,
	}
}

func (c *Client) WaitTxSuccess(txHash chainhash.Hash) error {
	begin := time.Now()

	defer func() {
		log.Infof("waitTxSuccess, duration_ms: %v", time.Since(begin).Milliseconds())
	}()

	count := 60
	for count > 0 {
		res, err := c.client.GetRawTransactionVerbose(&txHash)
		if err != nil {
			return fmt.Errorf("get raw transaction: %w", err)
		}

		if res == nil {
			count--

			time.Sleep(time.Second)
		} else {
			return nil
		}
	}

	return fmt.Errorf("get raw transaction fail")
}

func (c *Client) Close() error {
	c.cancel()
	return nil
}

func (c *Client) GetUTXOs(from btcutil.Address) ([]btcjson.ListUnspentResult, error) {
	// todo
	return c.client.ListUnspentMinMaxAddresses(0, 50, []btcutil.Address{from})
}

func (c *Client) SendTx(tx *wire.MsgTx) error {
	_, err := c.sendTx(tx, false)
	return err
}

type TxBound struct {
	OutPoint wire.OutPoint `json:"out_put_tx_hash,omitempty"`
	TxHash   string        `json:"tx_hash,omitempty"`
	TaskIds  []uint64      `json:"task_ids,omitempty"`
}

type SignHashCtx struct {
	signHash chainhash.Hash
	OutPoint wire.OutPoint
	signer   string
}

type TxClient struct {
	*Client
	tss             suite.TssService
	tasks           *pool.Pool[uint64]
	tx              *wire.MsgTx
	signHashCounter atomic.Int64
	signHashes      []*SignHashCtx
}

func NewTxClient(
	ctx context.Context,
	timeout time.Duration,
	params *chaincfg.Params,
	tssService suite.TssService,
	chainInfo *config.Chain,
) *TxClient {
	client := NewClient(ctx, timeout, params, chainInfo)

	return &TxClient{
		Client:          client,
		tx:              wire.NewMsgTx(wire.TxVersion),
		tss:             tssService,
		signHashCounter: atomic.Int64{},
		signHashes:      []*SignHashCtx{},
	}
}

// buildTxOut TxOut https://www.mengbin.top/2024-07-24-btcd_raw_tx/
func (c *TxClient) buildTxOut(addr string, amount int64) error {
	destinationAddress, err := btcutil.DecodeAddress(addr, c.params)
	if err != nil {
		return err
	}

	pkScript, err := txscript.PayToAddrScript(destinationAddress)
	if err != nil {
		return err
	}

	// satoshis
	c.tx.AddTxOut(wire.NewTxOut(amount, pkScript))

	return nil
}

func (c *TxClient) buildTxIn(from string, amount int64) error {
	// SerializeCompressed p2wpkh address
	fromAddr, err := btcutil.NewAddressWitnessPubKeyHash(
		btcutil.Hash160(c.tss.GetPublicKey(from).SerializeCompressed()),
		c.params,
	)
	if err != nil {
		return err
	}

	utxos, err := c.GetUTXOs(fromAddr)
	if err != nil {
		return err
	}

	// satoshis
	totalInput := int64(0)
	for _, utxo := range utxos {
		if totalInput > amount {
			break
		}

		txHash, err := chainhash.NewHashFromStr(utxo.TxID)
		if err != nil {
			return err
		}

		txIn := wire.NewTxIn(&wire.OutPoint{Hash: *txHash, Index: utxo.Vout}, nil, nil)
		c.tx.AddTxIn(txIn)

		totalInput += int64(utxo.Amount * precision)
	}

	// tx fee
	fee := int64(c.tx.SerializeSize())
	change := totalInput - amount

	if change > fee {
		changePkScript, err := txscript.PayToAddrScript(fromAddr)
		if err != nil {
			return err
		}

		txOut := wire.NewTxOut(change-fee, changePkScript)
		c.tx.AddTxOut(txOut)
	}

	for i := range c.tx.TxIn {
		prevOutputScript, err := hex.DecodeString(utxos[i].ScriptPubKey)
		if err != nil {
			return err
		}

		txHash, err := chainhash.NewHashFromStr(utxos[i].TxID)
		if err != nil {
			return err
		}

		outPoint := wire.OutPoint{Hash: *txHash, Index: utxos[i].Vout}
		prevOutputFetcher := txscript.NewMultiPrevOutFetcher(
			map[wire.OutPoint]*wire.TxOut{outPoint: {Value: int64(utxos[i].Amount * precision), PkScript: prevOutputScript}},
		)
		sigHashes := txscript.NewTxSigHashes(c.tx, prevOutputFetcher)

		hash, err := txscript.CalcWitnessSigHash(prevOutputScript, sigHashes, txscript.SigHashAll, c.tx, int(utxos[i].Vout), int64(utxos[i].Amount*precision))
		if err != nil {
			return err
		}

		c.signHashes = append(c.signHashes, &SignHashCtx{
			signHash: chainhash.Hash(hash),
			OutPoint: outPoint,
			signer:   utxos[i].Address,
		})
		c.signHashCounter.Add(1)
		// signature, err := c.sign(hash) //todo
		//
		//	if err != nil {
		//		return err
		//	}
		//
		// txscript.SigHashAll, // https://www.btcstudy.org/2021/11/09/bitcoin-signature-types-sighash/
		// signature = append(signature, byte(txscript.SigHashAll)) //todo
		//
		// txIn.Witness = wire.TxWitness{signature, c.publicKey.SerializeCompressed()}
	}

	return nil
}

func (c *TxClient) AddWitnessSignature(signHash, signature []byte) bool {
	for i, signCtx := range c.signHashes {
		if bytes.Equal(signCtx.signHash[:], signHash) {
			// txscript.SigHashAll, // https://www.btcstudy.org/2021/11/09/bitcoin-signature-types-sighash/
			signature = append(signature, byte(txscript.SigHashAll)) // todo
			c.tx.TxIn[i].Witness = wire.TxWitness{signature, c.tss.GetPublicKey(signCtx.signer).SerializeCompressed()}
			c.signHashCounter.Add(-1)
		}
	}

	return c.signHashCounter.Load() == 0
}

func (c *TxClient) IsHaveWitnessSignature(hash []byte) bool {
	for _, inHash := range c.signHashes {
		if bytes.Equal(inHash.signHash[:], hash) {
			return true
		}
	}

	return false
}

func (c *TxClient) Sign() error {
	txHash := c.TxHash()
	TaskIds := lo.Map(c.tasks.GetTopN(100), func(task pool.Task[uint64], index int) uint64 { return task.TaskID() }) // todo

	errs := make([]error, 0)

	for _, signCtx := range c.signHashes {
		bound := &TxBound{
			OutPoint: signCtx.OutPoint,
			TxHash:   txHash.String(),
			TaskIds:  TaskIds,
		}
		data, err := json.Marshal(bound)
		errs = append(errs, err)

		c.tss.Sign(&suite.SignReq{
			SeqId:      0, // todo
			Type:       types.SignTxSessionType,
			ChainId:    types.ChainIdBitcoin,
			Signer:     signCtx.signer,
			DataDigest: signCtx.signHash.String(),
			SignData:   signCtx.signHash[:],
			ExtraData:  data,
		})
	}

	return errors.Join(errs...)
}

func (c *Client) sendTx(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error) {
	hash, err := c.client.SendRawTransaction(tx, allowHighFees)
	if err != nil {
		return nil, fmt.Errorf("send raw transaction: %w", err)
	}

	return hash, nil
}

func (c *TxClient) BuildTx(from, to string, amount int64) error {
	return errors.Join(c.buildTxOut(to, amount), c.buildTxIn(from, amount))
}

type Receipt struct {
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}

func (c *TxClient) BuildOneToManyTx(from string, receipt []Receipt) error {
	errs := make([]error, 0)

	var amount uint64

	for _, r := range receipt {
		errs = append(errs, c.buildTxOut(r.To, int64(r.Amount)))
		amount += r.Amount
	}

	if amount > 0 {
		errs = append(errs, c.buildTxIn(from, int64(amount)))
	}

	return errors.Join(errs...)
}

func (c *TxClient) BuildCollectionTx(from []string, to string) error {
	errs := make([]error, 0)
	//for _, r := range from {
	//	errs = append(errs, c.buildTxIn(r, 0))
	//}
	return errors.Join(errs...)
}

func (c *TxClient) SendTx() error {
	return c.Client.SendTx(c.tx)
}

func (c *TxClient) WaitTxSuccess() error {
	return c.Client.WaitTxSuccess(c.tx.TxHash())
}

func (c *TxClient) TxHash() chainhash.Hash {
	return c.tx.TxHash()
}

func (c *TxClient) IsExist(id uint64) bool {
	return c.tasks.IsExist(id)
}

func (c *TxClient) AddTask(task pool.Task[uint64]) {
	c.tasks.Add(task)
}
