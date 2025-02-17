// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// InscriptionContractMetaData contains all meta data concerning the InscriptionContract contract.
var InscriptionContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"NIP20TokenEvent_burnb\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NIP20TokenEvent_mintb\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// InscriptionContractABI is the input ABI used to generate the binding from.
// Deprecated: Use InscriptionContractMetaData.ABI instead.
var InscriptionContractABI = InscriptionContractMetaData.ABI

// InscriptionContract is an auto generated Go binding around an Ethereum contract.
type InscriptionContract struct {
	InscriptionContractCaller     // Read-only binding to the contract
	InscriptionContractTransactor // Write-only binding to the contract
	InscriptionContractFilterer   // Log filterer for contract events
}

// InscriptionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type InscriptionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InscriptionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InscriptionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InscriptionContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InscriptionContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InscriptionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InscriptionContractSession struct {
	Contract     *InscriptionContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InscriptionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InscriptionContractCallerSession struct {
	Contract *InscriptionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// InscriptionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InscriptionContractTransactorSession struct {
	Contract     *InscriptionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// InscriptionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type InscriptionContractRaw struct {
	Contract *InscriptionContract // Generic contract binding to access the raw methods on
}

// InscriptionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InscriptionContractCallerRaw struct {
	Contract *InscriptionContractCaller // Generic read-only contract binding to access the raw methods on
}

// InscriptionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InscriptionContractTransactorRaw struct {
	Contract *InscriptionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInscriptionContract creates a new instance of InscriptionContract, bound to a specific deployed contract.
func NewInscriptionContract(address common.Address, backend bind.ContractBackend) (*InscriptionContract, error) {
	contract, err := bindInscriptionContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InscriptionContract{InscriptionContractCaller: InscriptionContractCaller{contract: contract}, InscriptionContractTransactor: InscriptionContractTransactor{contract: contract}, InscriptionContractFilterer: InscriptionContractFilterer{contract: contract}}, nil
}

// NewInscriptionContractCaller creates a new read-only instance of InscriptionContract, bound to a specific deployed contract.
func NewInscriptionContractCaller(address common.Address, caller bind.ContractCaller) (*InscriptionContractCaller, error) {
	contract, err := bindInscriptionContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InscriptionContractCaller{contract: contract}, nil
}

// NewInscriptionContractTransactor creates a new write-only instance of InscriptionContract, bound to a specific deployed contract.
func NewInscriptionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*InscriptionContractTransactor, error) {
	contract, err := bindInscriptionContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InscriptionContractTransactor{contract: contract}, nil
}

// NewInscriptionContractFilterer creates a new log filterer instance of InscriptionContract, bound to a specific deployed contract.
func NewInscriptionContractFilterer(address common.Address, filterer bind.ContractFilterer) (*InscriptionContractFilterer, error) {
	contract, err := bindInscriptionContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InscriptionContractFilterer{contract: contract}, nil
}

// bindInscriptionContract binds a generic wrapper to an already deployed contract.
func bindInscriptionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InscriptionContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InscriptionContract *InscriptionContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InscriptionContract.Contract.InscriptionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InscriptionContract *InscriptionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InscriptionContract.Contract.InscriptionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InscriptionContract *InscriptionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InscriptionContract.Contract.InscriptionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InscriptionContract *InscriptionContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InscriptionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InscriptionContract *InscriptionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InscriptionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InscriptionContract *InscriptionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InscriptionContract.Contract.contract.Transact(opts, method, params...)
}

// InscriptionContractNIP20TokenEventBurnbIterator is returned from FilterNIP20TokenEventBurnb and is used to iterate over the raw logs and unpacked data for NIP20TokenEventBurnb events raised by the InscriptionContract contract.
type InscriptionContractNIP20TokenEventBurnbIterator struct {
	Event *InscriptionContractNIP20TokenEventBurnb // Event containing the contract specifics and raw log

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
func (it *InscriptionContractNIP20TokenEventBurnbIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InscriptionContractNIP20TokenEventBurnb)
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
		it.Event = new(InscriptionContractNIP20TokenEventBurnb)
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
func (it *InscriptionContractNIP20TokenEventBurnbIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InscriptionContractNIP20TokenEventBurnbIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InscriptionContractNIP20TokenEventBurnb represents a NIP20TokenEventBurnb event raised by the InscriptionContract contract.
type InscriptionContractNIP20TokenEventBurnb struct {
	From   common.Address
	Ticker [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNIP20TokenEventBurnb is a free log retrieval operation binding the contract event 0xebe23dd93b970477278ceb9abd3082df92d977d6131fb0ef75f18c3d353b565a.
//
// Solidity: event NIP20TokenEvent_burnb(address from, bytes32 ticker, uint256 amount)
func (_InscriptionContract *InscriptionContractFilterer) FilterNIP20TokenEventBurnb(opts *bind.FilterOpts) (*InscriptionContractNIP20TokenEventBurnbIterator, error) {

	logs, sub, err := _InscriptionContract.contract.FilterLogs(opts, "NIP20TokenEvent_burnb")
	if err != nil {
		return nil, err
	}
	return &InscriptionContractNIP20TokenEventBurnbIterator{contract: _InscriptionContract.contract, event: "NIP20TokenEvent_burnb", logs: logs, sub: sub}, nil
}

// WatchNIP20TokenEventBurnb is a free log subscription operation binding the contract event 0xebe23dd93b970477278ceb9abd3082df92d977d6131fb0ef75f18c3d353b565a.
//
// Solidity: event NIP20TokenEvent_burnb(address from, bytes32 ticker, uint256 amount)
func (_InscriptionContract *InscriptionContractFilterer) WatchNIP20TokenEventBurnb(opts *bind.WatchOpts, sink chan<- *InscriptionContractNIP20TokenEventBurnb) (event.Subscription, error) {

	logs, sub, err := _InscriptionContract.contract.WatchLogs(opts, "NIP20TokenEvent_burnb")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InscriptionContractNIP20TokenEventBurnb)
				if err := _InscriptionContract.contract.UnpackLog(event, "NIP20TokenEvent_burnb", log); err != nil {
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

// ParseNIP20TokenEventBurnb is a log parse operation binding the contract event 0xebe23dd93b970477278ceb9abd3082df92d977d6131fb0ef75f18c3d353b565a.
//
// Solidity: event NIP20TokenEvent_burnb(address from, bytes32 ticker, uint256 amount)
func (_InscriptionContract *InscriptionContractFilterer) ParseNIP20TokenEventBurnb(log types.Log) (*InscriptionContractNIP20TokenEventBurnb, error) {
	event := new(InscriptionContractNIP20TokenEventBurnb)
	if err := _InscriptionContract.contract.UnpackLog(event, "NIP20TokenEvent_burnb", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InscriptionContractNIP20TokenEventMintbIterator is returned from FilterNIP20TokenEventMintb and is used to iterate over the raw logs and unpacked data for NIP20TokenEventMintb events raised by the InscriptionContract contract.
type InscriptionContractNIP20TokenEventMintbIterator struct {
	Event *InscriptionContractNIP20TokenEventMintb // Event containing the contract specifics and raw log

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
func (it *InscriptionContractNIP20TokenEventMintbIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InscriptionContractNIP20TokenEventMintb)
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
		it.Event = new(InscriptionContractNIP20TokenEventMintb)
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
func (it *InscriptionContractNIP20TokenEventMintbIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InscriptionContractNIP20TokenEventMintbIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InscriptionContractNIP20TokenEventMintb represents a NIP20TokenEventMintb event raised by the InscriptionContract contract.
type InscriptionContractNIP20TokenEventMintb struct {
	Recipient common.Address
	Ticker    [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNIP20TokenEventMintb is a free log retrieval operation binding the contract event 0x685c530c280ee1f7a4e96d082303ee9ebf21cec512259c6a943eda3854e05102.
//
// Solidity: event NIP20TokenEvent_mintb(address recipient, bytes32 ticker, uint256 amount)
func (_InscriptionContract *InscriptionContractFilterer) FilterNIP20TokenEventMintb(opts *bind.FilterOpts) (*InscriptionContractNIP20TokenEventMintbIterator, error) {

	logs, sub, err := _InscriptionContract.contract.FilterLogs(opts, "NIP20TokenEvent_mintb")
	if err != nil {
		return nil, err
	}
	return &InscriptionContractNIP20TokenEventMintbIterator{contract: _InscriptionContract.contract, event: "NIP20TokenEvent_mintb", logs: logs, sub: sub}, nil
}

// WatchNIP20TokenEventMintb is a free log subscription operation binding the contract event 0x685c530c280ee1f7a4e96d082303ee9ebf21cec512259c6a943eda3854e05102.
//
// Solidity: event NIP20TokenEvent_mintb(address recipient, bytes32 ticker, uint256 amount)
func (_InscriptionContract *InscriptionContractFilterer) WatchNIP20TokenEventMintb(opts *bind.WatchOpts, sink chan<- *InscriptionContractNIP20TokenEventMintb) (event.Subscription, error) {

	logs, sub, err := _InscriptionContract.contract.WatchLogs(opts, "NIP20TokenEvent_mintb")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InscriptionContractNIP20TokenEventMintb)
				if err := _InscriptionContract.contract.UnpackLog(event, "NIP20TokenEvent_mintb", log); err != nil {
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

// ParseNIP20TokenEventMintb is a log parse operation binding the contract event 0x685c530c280ee1f7a4e96d082303ee9ebf21cec512259c6a943eda3854e05102.
//
// Solidity: event NIP20TokenEvent_mintb(address recipient, bytes32 ticker, uint256 amount)
func (_InscriptionContract *InscriptionContractFilterer) ParseNIP20TokenEventMintb(log types.Log) (*InscriptionContractNIP20TokenEventMintb, error) {
	event := new(InscriptionContractNIP20TokenEventMintb)
	if err := _InscriptionContract.contract.UnpackLog(event, "NIP20TokenEvent_mintb", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
