// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package codec

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Operation is an auto generated low-level Go binding around an user-defined struct.
type Operation struct {
	TaskId    uint64
	State     uint8
	ExtraData []byte
}

// VoterCodecMetaData contains all meta data concerning the VoterCodec contract.
var VoterCodecMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"taskId\",\"type\":\"uint64\"},{\"internalType\":\"enumState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOperation[]\",\"name\":\"opts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"taskId\",\"type\":\"uint64\"},{\"internalType\":\"enumState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOperation[]\",\"name\":\"opts\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"operations\",\"type\":\"event\"}]",
}

// VoterCodecABI is the input ABI used to generate the binding from.
// Deprecated: Use VoterCodecMetaData.ABI instead.
var VoterCodecABI = VoterCodecMetaData.ABI

// VoterCodec is an auto generated Go binding around an Ethereum contract.
type VoterCodec struct {
	VoterCodecCaller     // Read-only binding to the contract
	VoterCodecTransactor // Write-only binding to the contract
	VoterCodecFilterer   // Log filterer for contract events
}

// VoterCodecCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterCodecCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterCodecTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterCodecTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterCodecFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoterCodecFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterCodecSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterCodecSession struct {
	Contract     *VoterCodec       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterCodecCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterCodecCallerSession struct {
	Contract *VoterCodecCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VoterCodecTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterCodecTransactorSession struct {
	Contract     *VoterCodecTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VoterCodecRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterCodecRaw struct {
	Contract *VoterCodec // Generic contract binding to access the raw methods on
}

// VoterCodecCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterCodecCallerRaw struct {
	Contract *VoterCodecCaller // Generic read-only contract binding to access the raw methods on
}

// VoterCodecTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterCodecTransactorRaw struct {
	Contract *VoterCodecTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoterCodec creates a new instance of VoterCodec, bound to a specific deployed contract.
func NewVoterCodec(address common.Address, backend bind.ContractBackend) (*VoterCodec, error) {
	contract, err := bindVoterCodec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoterCodec{VoterCodecCaller: VoterCodecCaller{contract: contract}, VoterCodecTransactor: VoterCodecTransactor{contract: contract}, VoterCodecFilterer: VoterCodecFilterer{contract: contract}}, nil
}

// NewVoterCodecCaller creates a new read-only instance of VoterCodec, bound to a specific deployed contract.
func NewVoterCodecCaller(address common.Address, caller bind.ContractCaller) (*VoterCodecCaller, error) {
	contract, err := bindVoterCodec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoterCodecCaller{contract: contract}, nil
}

// NewVoterCodecTransactor creates a new write-only instance of VoterCodec, bound to a specific deployed contract.
func NewVoterCodecTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterCodecTransactor, error) {
	contract, err := bindVoterCodec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoterCodecTransactor{contract: contract}, nil
}

// NewVoterCodecFilterer creates a new log filterer instance of VoterCodec, bound to a specific deployed contract.
func NewVoterCodecFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterCodecFilterer, error) {
	contract, err := bindVoterCodec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoterCodecFilterer{contract: contract}, nil
}

// bindVoterCodec binds a generic wrapper to an already deployed contract.
func bindVoterCodec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoterCodecMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterCodec *VoterCodecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoterCodec.Contract.VoterCodecCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterCodec *VoterCodecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterCodec.Contract.VoterCodecTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterCodec *VoterCodecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterCodec.Contract.VoterCodecTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterCodec *VoterCodecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoterCodec.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterCodec *VoterCodecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterCodec.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterCodec *VoterCodecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterCodec.Contract.contract.Transact(opts, method, params...)
}

// VoterCodecOperationsIterator is returned from FilterOperations and is used to iterate over the raw logs and unpacked data for Operations events raised by the VoterCodec contract.
type VoterCodecOperationsIterator struct {
	Event *VoterCodecOperations // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoterCodecOperationsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterCodecOperations)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoterCodecOperations)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoterCodecOperationsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterCodecOperationsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterCodecOperations represents a Operations event raised by the VoterCodec contract.
type VoterCodecOperations struct {
	Opts    []Operation
	Nonce   *big.Int
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOperations is a free log retrieval operation binding the contract event 0xd13b7446735b33f1e2b0bb372153f282e84181aa7e64ebaf82818283798dfaa6.
//
// Solidity: event operations((uint64,uint8,bytes)[] opts, uint256 nonce, uint256 chainId)
func (_VoterCodec *VoterCodecFilterer) FilterOperations(opts *bind.FilterOpts) (*VoterCodecOperationsIterator, error) {

	logs, sub, err := _VoterCodec.contract.FilterLogs(opts, "operations")
	if err != nil {
		return nil, err
	}
	return &VoterCodecOperationsIterator{contract: _VoterCodec.contract, event: "operations", logs: logs, sub: sub}, nil
}

// WatchOperations is a free log subscription operation binding the contract event 0xd13b7446735b33f1e2b0bb372153f282e84181aa7e64ebaf82818283798dfaa6.
//
// Solidity: event operations((uint64,uint8,bytes)[] opts, uint256 nonce, uint256 chainId)
func (_VoterCodec *VoterCodecFilterer) WatchOperations(opts *bind.WatchOpts, sink chan<- *VoterCodecOperations) (event.Subscription, error) {

	logs, sub, err := _VoterCodec.contract.WatchLogs(opts, "operations")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterCodecOperations)
				if err := _VoterCodec.contract.UnpackLog(event, "operations", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperations is a log parse operation binding the contract event 0xd13b7446735b33f1e2b0bb372153f282e84181aa7e64ebaf82818283798dfaa6.
//
// Solidity: event operations((uint64,uint8,bytes)[] opts, uint256 nonce, uint256 chainId)
func (_VoterCodec *VoterCodecFilterer) ParseOperations(log types.Log) (*VoterCodecOperations, error) {
	event := new(VoterCodecOperations)
	if err := _VoterCodec.contract.UnpackLog(event, "operations", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
