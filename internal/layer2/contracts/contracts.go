package contracts

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts/codec"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

func Pack(meta *bind.MetaData, method string, params ...interface{}) []byte {
	a, err := meta.GetAbi()
	if err != nil {
		panic(err)
	}

	// Otherwise pack up the parameters and invoke the contract
	input, err := a.Pack(method, params...)
	if err != nil {
		panic(err)
	}

	return input
}

func PackEvent(meta *bind.MetaData, name string, args ...interface{}) []byte {
	a, err := meta.GetAbi()
	if err != nil {
		panic(err)
	}
	// Otherwise pack up the parameters and invoke the contract
	event, exist := a.Events[name]
	if !exist {
		panic(fmt.Errorf("event '%s' not found", name))
	}

	arguments, err := event.Inputs.Pack(args...)
	if err != nil {
		panic(err)
	}
	// Pack up the event ID too if not a constructor and return
	return append(event.ID.Bytes(), arguments...)
}

func UnpackEvent(meta *bind.MetaData, name string, event any, data []byte) {
	parsedABI, err := ParseABI(meta.ABI)
	utils.Assert(err)
	utils.Assert(parsedABI.UnpackIntoInterface(event, name, data))
}

func EncodeFun(abiStr string, method string, params ...interface{}) []byte {
	return Pack(
		&bind.MetaData{
			ABI: abiStr,
		},
		method,
		params...,
	)
}

func ParseABI(abiJSON string) (abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("unable to parse abi json: %w", err)
	}

	return parsed, nil
}

func MustParseABI(abiJSON string) abi.ABI {
	parsed, err := ParseABI(abiJSON)
	if err != nil {
		panic(err)
	}

	return parsed
}

type Contract struct {
	ABI abi.ABI
}

func NewContract(abiMeta *bind.MetaData) *Contract {
	abi, err := ParseABI(abiMeta.ABI)
	if err != nil {
		panic(err)
	}

	return &Contract{
		ABI: abi,
	}
}

func (c *Contract) Encode(method string, args ...any) ([]byte, error) {
	// Otherwise pack up the parameters and invoke the contract
	return c.ABI.Pack(method, args...)
}

func (c *Contract) Decode(method string, data []byte) ([]any, error) {
	return c.ABI.Unpack(method, data)
}

func (c *Contract) MethodID(method string) string {
	m := c.ABI.Methods[method]
	return hex.EncodeToString(m.ID)
}

func (c *Contract) UnPackInput(method string, input any, data []byte) error {
	m := c.ABI.Methods[method]

	unpacked, err := m.Inputs.Unpack(data)
	if err != nil {
		return err
	}

	return m.Inputs.Copy(input, unpacked)
}

func (c *Contract) EventTopicHash(eventName string) (common.Hash, error) {
	ev, ok := c.ABI.Events[eventName]
	if !ok {
		return common.Hash{}, fmt.Errorf("ethcontract: event '%s' not found in contract abi", eventName)
	}

	h := crypto.Keccak256Hash([]byte(ev.Sig))

	return h, nil
}

func EncodeTransferOfERC20(from, to common.Address, amount *big.Int) []byte {
	return EncodeFun(ERC20ABI, "transferFrom", from, to, amount)
}

func EncodeVerifyAndCall(_target common.Address, _data []byte, _signature []byte) []byte {
	return EncodeFun(VotingManagerContractMetaData.ABI, "verifyAndCall", _target, _data, _signature)
}

func EncodeSubmitTaskReceipt(taskId *big.Int, result []byte, signature []byte) []byte {
	return EncodeFun(VotingManagerContractMetaData.ABI, "submitTaskReceipt", taskId, result, signature)
}

var (
	errNoEventSignature       = errors.New("no event signature")
	errEventSignatureMismatch = errors.New("event signature mismatch")
)

// UnpackEventLog unpacks a retrieved log into the provided output structure.
func UnpackEventLog(meta *bind.MetaData, out interface{}, event string, log types.Log) {
	err := unpackEventLog(meta, out, event, log)
	if err != nil {
		panic(err)
	}
}

func unpackEventLog(meta *bind.MetaData, out interface{}, event string, log types.Log) error {
	a, err := meta.GetAbi()
	if err != nil {
		return err
	}

	// Anonymous events are not supported.
	if len(log.Topics) == 0 {
		return errNoEventSignature
	}

	if log.Topics[0] != a.Events[event].ID {
		return errEventSignatureMismatch
	}

	if len(log.Data) > 0 {
		if err := a.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}

	var indexed abi.Arguments

	for _, arg := range a.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	return abi.ParseTopics(out, indexed, log.Topics[1:])
}

func EncodeOperation(operations []codec.Operation, tssNonce *big.Int, chainId *big.Int) []byte {
	return EncodeFun(codec.VoterCodecMetaData.ABI, "", operations, tssNonce, chainId)
}

func EventTopic(meta *bind.MetaData, eventName string) common.Hash {
	abi, err := meta.GetAbi()
	if err != nil {
		panic(err)
	}

	ev, ok := abi.Events[eventName]
	if !ok {
		panic(fmt.Errorf("ethcontract: event '%s' not found in contract abi", eventName))
	}

	return crypto.Keccak256Hash([]byte(ev.Sig))
}

func MethodID(meta *bind.MetaData, method string) string {
	abi, err := meta.GetAbi()
	if err != nil {
		panic(err)
	}

	ev, ok := abi.Methods[method]
	if !ok {
		panic(fmt.Errorf("ethcontract: event '%s' not found in contract abi", method))
	}

	return hex.EncodeToString(ev.ID)
}
