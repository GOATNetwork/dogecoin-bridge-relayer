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

// TaskPayloadContractMetaData contains all meta data concerning the TaskPayloadContract contract.
var TaskPayloadContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ticker\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"ConsolidationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ticker\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"depositAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"}],\"name\":\"DepositRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"errorCode\",\"type\":\"uint8\"}],\"name\":\"DepositResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"errorCode\",\"type\":\"uint8\"}],\"name\":\"TaskResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ticker\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"TransferRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_account\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"enumTaskPayload.AddressCategory\",\"name\":\"_chain\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_index\",\"type\":\"uint32\"}],\"name\":\"WalletCreationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"errorCode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"walletAddress\",\"type\":\"string\"}],\"name\":\"WalletCreationResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ticker\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"errorCode\",\"type\":\"uint8\"}],\"name\":\"WithdrawalResult\",\"type\":\"event\"}]",
}

// TaskPayloadContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskPayloadContractMetaData.ABI instead.
var TaskPayloadContractABI = TaskPayloadContractMetaData.ABI

// TaskPayloadContract is an auto generated Go binding around an Ethereum contract.
type TaskPayloadContract struct {
	TaskPayloadContractCaller     // Read-only binding to the contract
	TaskPayloadContractTransactor // Write-only binding to the contract
	TaskPayloadContractFilterer   // Log filterer for contract events
}

// TaskPayloadContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskPayloadContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskPayloadContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskPayloadContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskPayloadContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskPayloadContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskPayloadContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskPayloadContractSession struct {
	Contract     *TaskPayloadContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TaskPayloadContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskPayloadContractCallerSession struct {
	Contract *TaskPayloadContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TaskPayloadContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskPayloadContractTransactorSession struct {
	Contract     *TaskPayloadContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TaskPayloadContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskPayloadContractRaw struct {
	Contract *TaskPayloadContract // Generic contract binding to access the raw methods on
}

// TaskPayloadContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskPayloadContractCallerRaw struct {
	Contract *TaskPayloadContractCaller // Generic read-only contract binding to access the raw methods on
}

// TaskPayloadContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskPayloadContractTransactorRaw struct {
	Contract *TaskPayloadContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskPayloadContract creates a new instance of TaskPayloadContract, bound to a specific deployed contract.
func NewTaskPayloadContract(address common.Address, backend bind.ContractBackend) (*TaskPayloadContract, error) {
	contract, err := bindTaskPayloadContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContract{TaskPayloadContractCaller: TaskPayloadContractCaller{contract: contract}, TaskPayloadContractTransactor: TaskPayloadContractTransactor{contract: contract}, TaskPayloadContractFilterer: TaskPayloadContractFilterer{contract: contract}}, nil
}

// NewTaskPayloadContractCaller creates a new read-only instance of TaskPayloadContract, bound to a specific deployed contract.
func NewTaskPayloadContractCaller(address common.Address, caller bind.ContractCaller) (*TaskPayloadContractCaller, error) {
	contract, err := bindTaskPayloadContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractCaller{contract: contract}, nil
}

// NewTaskPayloadContractTransactor creates a new write-only instance of TaskPayloadContract, bound to a specific deployed contract.
func NewTaskPayloadContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskPayloadContractTransactor, error) {
	contract, err := bindTaskPayloadContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractTransactor{contract: contract}, nil
}

// NewTaskPayloadContractFilterer creates a new log filterer instance of TaskPayloadContract, bound to a specific deployed contract.
func NewTaskPayloadContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskPayloadContractFilterer, error) {
	contract, err := bindTaskPayloadContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractFilterer{contract: contract}, nil
}

// bindTaskPayloadContract binds a generic wrapper to an already deployed contract.
func bindTaskPayloadContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskPayloadContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskPayloadContract *TaskPayloadContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskPayloadContract.Contract.TaskPayloadContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskPayloadContract *TaskPayloadContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskPayloadContract.Contract.TaskPayloadContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskPayloadContract *TaskPayloadContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskPayloadContract.Contract.TaskPayloadContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskPayloadContract *TaskPayloadContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskPayloadContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskPayloadContract *TaskPayloadContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskPayloadContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskPayloadContract *TaskPayloadContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskPayloadContract.Contract.contract.Transact(opts, method, params...)
}

// TaskPayloadContractConsolidationRequestIterator is returned from FilterConsolidationRequest and is used to iterate over the raw logs and unpacked data for ConsolidationRequest events raised by the TaskPayloadContract contract.
type TaskPayloadContractConsolidationRequestIterator struct {
	Event *TaskPayloadContractConsolidationRequest // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractConsolidationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractConsolidationRequest)
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
		it.Event = new(TaskPayloadContractConsolidationRequest)
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
func (it *TaskPayloadContractConsolidationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractConsolidationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractConsolidationRequest represents a ConsolidationRequest event raised by the TaskPayloadContract contract.
type TaskPayloadContractConsolidationRequest struct {
	FromAddress string
	Ticker      [32]byte
	ChainId     uint64
	Amount      *big.Int
	Offset      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsolidationRequest is a free log retrieval operation binding the contract event 0x34e87801f39043cc3dae17b66d4877b18013efe557bb7e69245a32b8eea415e2.
//
// Solidity: event ConsolidationRequest(string fromAddress, bytes32 ticker, uint64 chainId, uint256 amount, uint256 offset)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterConsolidationRequest(opts *bind.FilterOpts) (*TaskPayloadContractConsolidationRequestIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "ConsolidationRequest")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractConsolidationRequestIterator{contract: _TaskPayloadContract.contract, event: "ConsolidationRequest", logs: logs, sub: sub}, nil
}

// WatchConsolidationRequest is a free log subscription operation binding the contract event 0x34e87801f39043cc3dae17b66d4877b18013efe557bb7e69245a32b8eea415e2.
//
// Solidity: event ConsolidationRequest(string fromAddress, bytes32 ticker, uint64 chainId, uint256 amount, uint256 offset)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchConsolidationRequest(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractConsolidationRequest) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "ConsolidationRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractConsolidationRequest)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "ConsolidationRequest", log); err != nil {
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

// ParseConsolidationRequest is a log parse operation binding the contract event 0x34e87801f39043cc3dae17b66d4877b18013efe557bb7e69245a32b8eea415e2.
//
// Solidity: event ConsolidationRequest(string fromAddress, bytes32 ticker, uint64 chainId, uint256 amount, uint256 offset)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseConsolidationRequest(log types.Log) (*TaskPayloadContractConsolidationRequest, error) {
	event := new(TaskPayloadContractConsolidationRequest)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "ConsolidationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractDepositRequestIterator is returned from FilterDepositRequest and is used to iterate over the raw logs and unpacked data for DepositRequest events raised by the TaskPayloadContract contract.
type TaskPayloadContractDepositRequestIterator struct {
	Event *TaskPayloadContractDepositRequest // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractDepositRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractDepositRequest)
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
		it.Event = new(TaskPayloadContractDepositRequest)
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
func (it *TaskPayloadContractDepositRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractDepositRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractDepositRequest represents a DepositRequest event raised by the TaskPayloadContract contract.
type TaskPayloadContractDepositRequest struct {
	UserAddress    common.Address
	ChainId        uint64
	Ticker         [32]byte
	DepositAddress string
	Amount         *big.Int
	TxHash         string
	BlockHeight    *big.Int
	LogIndex       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDepositRequest is a free log retrieval operation binding the contract event 0x46eb35c874886d4b8befed8b686ac3486e433d9dc0ba1f1b329034afa00012e0.
//
// Solidity: event DepositRequest(address userAddress, uint64 chainId, bytes32 ticker, string depositAddress, uint256 amount, string txHash, uint256 blockHeight, uint256 logIndex)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterDepositRequest(opts *bind.FilterOpts) (*TaskPayloadContractDepositRequestIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "DepositRequest")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractDepositRequestIterator{contract: _TaskPayloadContract.contract, event: "DepositRequest", logs: logs, sub: sub}, nil
}

// WatchDepositRequest is a free log subscription operation binding the contract event 0x46eb35c874886d4b8befed8b686ac3486e433d9dc0ba1f1b329034afa00012e0.
//
// Solidity: event DepositRequest(address userAddress, uint64 chainId, bytes32 ticker, string depositAddress, uint256 amount, string txHash, uint256 blockHeight, uint256 logIndex)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchDepositRequest(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractDepositRequest) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "DepositRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractDepositRequest)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "DepositRequest", log); err != nil {
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

// ParseDepositRequest is a log parse operation binding the contract event 0x46eb35c874886d4b8befed8b686ac3486e433d9dc0ba1f1b329034afa00012e0.
//
// Solidity: event DepositRequest(address userAddress, uint64 chainId, bytes32 ticker, string depositAddress, uint256 amount, string txHash, uint256 blockHeight, uint256 logIndex)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseDepositRequest(log types.Log) (*TaskPayloadContractDepositRequest, error) {
	event := new(TaskPayloadContractDepositRequest)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "DepositRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractDepositResultIterator is returned from FilterDepositResult and is used to iterate over the raw logs and unpacked data for DepositResult events raised by the TaskPayloadContract contract.
type TaskPayloadContractDepositResultIterator struct {
	Event *TaskPayloadContractDepositResult // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractDepositResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractDepositResult)
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
		it.Event = new(TaskPayloadContractDepositResult)
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
func (it *TaskPayloadContractDepositResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractDepositResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractDepositResult represents a DepositResult event raised by the TaskPayloadContract contract.
type TaskPayloadContractDepositResult struct {
	Version   uint8
	Success   bool
	ErrorCode uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositResult is a free log retrieval operation binding the contract event 0xae9e6016838d9912f513c2adb0656673485ababaddbf853b28d121bf2ce24b9e.
//
// Solidity: event DepositResult(uint8 version, bool success, uint8 errorCode)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterDepositResult(opts *bind.FilterOpts) (*TaskPayloadContractDepositResultIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "DepositResult")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractDepositResultIterator{contract: _TaskPayloadContract.contract, event: "DepositResult", logs: logs, sub: sub}, nil
}

// WatchDepositResult is a free log subscription operation binding the contract event 0xae9e6016838d9912f513c2adb0656673485ababaddbf853b28d121bf2ce24b9e.
//
// Solidity: event DepositResult(uint8 version, bool success, uint8 errorCode)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchDepositResult(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractDepositResult) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "DepositResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractDepositResult)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "DepositResult", log); err != nil {
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

// ParseDepositResult is a log parse operation binding the contract event 0xae9e6016838d9912f513c2adb0656673485ababaddbf853b28d121bf2ce24b9e.
//
// Solidity: event DepositResult(uint8 version, bool success, uint8 errorCode)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseDepositResult(log types.Log) (*TaskPayloadContractDepositResult, error) {
	event := new(TaskPayloadContractDepositResult)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "DepositResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractTaskResultIterator is returned from FilterTaskResult and is used to iterate over the raw logs and unpacked data for TaskResult events raised by the TaskPayloadContract contract.
type TaskPayloadContractTaskResultIterator struct {
	Event *TaskPayloadContractTaskResult // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractTaskResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractTaskResult)
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
		it.Event = new(TaskPayloadContractTaskResult)
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
func (it *TaskPayloadContractTaskResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractTaskResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractTaskResult represents a TaskResult event raised by the TaskPayloadContract contract.
type TaskPayloadContractTaskResult struct {
	Version   uint8
	Success   bool
	ErrorCode uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskResult is a free log retrieval operation binding the contract event 0x168d1dedeecc1984c4fb8e094dadd280c50fcd66b8f07b73148cd62f65e847c6.
//
// Solidity: event TaskResult(uint8 version, bool success, uint8 errorCode)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterTaskResult(opts *bind.FilterOpts) (*TaskPayloadContractTaskResultIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "TaskResult")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractTaskResultIterator{contract: _TaskPayloadContract.contract, event: "TaskResult", logs: logs, sub: sub}, nil
}

// WatchTaskResult is a free log subscription operation binding the contract event 0x168d1dedeecc1984c4fb8e094dadd280c50fcd66b8f07b73148cd62f65e847c6.
//
// Solidity: event TaskResult(uint8 version, bool success, uint8 errorCode)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchTaskResult(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractTaskResult) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "TaskResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractTaskResult)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "TaskResult", log); err != nil {
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

// ParseTaskResult is a log parse operation binding the contract event 0x168d1dedeecc1984c4fb8e094dadd280c50fcd66b8f07b73148cd62f65e847c6.
//
// Solidity: event TaskResult(uint8 version, bool success, uint8 errorCode)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseTaskResult(log types.Log) (*TaskPayloadContractTaskResult, error) {
	event := new(TaskPayloadContractTaskResult)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "TaskResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractTransferRequestIterator is returned from FilterTransferRequest and is used to iterate over the raw logs and unpacked data for TransferRequest events raised by the TaskPayloadContract contract.
type TaskPayloadContractTransferRequestIterator struct {
	Event *TaskPayloadContractTransferRequest // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractTransferRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractTransferRequest)
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
		it.Event = new(TaskPayloadContractTransferRequest)
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
func (it *TaskPayloadContractTransferRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractTransferRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractTransferRequest represents a TransferRequest event raised by the TaskPayloadContract contract.
type TaskPayloadContractTransferRequest struct {
	FromAddress string
	ToAddress   string
	Ticker      [32]byte
	ChainId     uint64
	Amount      *big.Int
	Offset      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferRequest is a free log retrieval operation binding the contract event 0x3d6d85aed1e63ef460b1191c1793b551609790450966f5a12ee4bc84b23eb53c.
//
// Solidity: event TransferRequest(string fromAddress, string toAddress, bytes32 ticker, uint64 chainId, uint256 amount, uint256 offset)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterTransferRequest(opts *bind.FilterOpts) (*TaskPayloadContractTransferRequestIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "TransferRequest")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractTransferRequestIterator{contract: _TaskPayloadContract.contract, event: "TransferRequest", logs: logs, sub: sub}, nil
}

// WatchTransferRequest is a free log subscription operation binding the contract event 0x3d6d85aed1e63ef460b1191c1793b551609790450966f5a12ee4bc84b23eb53c.
//
// Solidity: event TransferRequest(string fromAddress, string toAddress, bytes32 ticker, uint64 chainId, uint256 amount, uint256 offset)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchTransferRequest(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractTransferRequest) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "TransferRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractTransferRequest)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "TransferRequest", log); err != nil {
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

// ParseTransferRequest is a log parse operation binding the contract event 0x3d6d85aed1e63ef460b1191c1793b551609790450966f5a12ee4bc84b23eb53c.
//
// Solidity: event TransferRequest(string fromAddress, string toAddress, bytes32 ticker, uint64 chainId, uint256 amount, uint256 offset)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseTransferRequest(log types.Log) (*TaskPayloadContractTransferRequest, error) {
	event := new(TaskPayloadContractTransferRequest)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "TransferRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractWalletCreationRequestIterator is returned from FilterWalletCreationRequest and is used to iterate over the raw logs and unpacked data for WalletCreationRequest events raised by the TaskPayloadContract contract.
type TaskPayloadContractWalletCreationRequestIterator struct {
	Event *TaskPayloadContractWalletCreationRequest // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractWalletCreationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractWalletCreationRequest)
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
		it.Event = new(TaskPayloadContractWalletCreationRequest)
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
func (it *TaskPayloadContractWalletCreationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractWalletCreationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractWalletCreationRequest represents a WalletCreationRequest event raised by the TaskPayloadContract contract.
type TaskPayloadContractWalletCreationRequest struct {
	UserAddr common.Address
	Account  uint32
	Chain    uint8
	Index    uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletCreationRequest is a free log retrieval operation binding the contract event 0x7da35a471eda0737635babc5325016e3fb598528bc3905f4ed60674d17abffd6.
//
// Solidity: event WalletCreationRequest(address _userAddr, uint32 _account, uint8 _chain, uint32 _index)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterWalletCreationRequest(opts *bind.FilterOpts) (*TaskPayloadContractWalletCreationRequestIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "WalletCreationRequest")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractWalletCreationRequestIterator{contract: _TaskPayloadContract.contract, event: "WalletCreationRequest", logs: logs, sub: sub}, nil
}

// WatchWalletCreationRequest is a free log subscription operation binding the contract event 0x7da35a471eda0737635babc5325016e3fb598528bc3905f4ed60674d17abffd6.
//
// Solidity: event WalletCreationRequest(address _userAddr, uint32 _account, uint8 _chain, uint32 _index)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchWalletCreationRequest(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractWalletCreationRequest) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "WalletCreationRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractWalletCreationRequest)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "WalletCreationRequest", log); err != nil {
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

// ParseWalletCreationRequest is a log parse operation binding the contract event 0x7da35a471eda0737635babc5325016e3fb598528bc3905f4ed60674d17abffd6.
//
// Solidity: event WalletCreationRequest(address _userAddr, uint32 _account, uint8 _chain, uint32 _index)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseWalletCreationRequest(log types.Log) (*TaskPayloadContractWalletCreationRequest, error) {
	event := new(TaskPayloadContractWalletCreationRequest)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "WalletCreationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractWalletCreationResultIterator is returned from FilterWalletCreationResult and is used to iterate over the raw logs and unpacked data for WalletCreationResult events raised by the TaskPayloadContract contract.
type TaskPayloadContractWalletCreationResultIterator struct {
	Event *TaskPayloadContractWalletCreationResult // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractWalletCreationResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractWalletCreationResult)
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
		it.Event = new(TaskPayloadContractWalletCreationResult)
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
func (it *TaskPayloadContractWalletCreationResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractWalletCreationResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractWalletCreationResult represents a WalletCreationResult event raised by the TaskPayloadContract contract.
type TaskPayloadContractWalletCreationResult struct {
	Version       uint8
	Success       bool
	ErrorCode     uint8
	WalletAddress string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWalletCreationResult is a free log retrieval operation binding the contract event 0x440691550bb1f6d18c60b1a17fff36325a996ba3ab5917f3003445984c5302cf.
//
// Solidity: event WalletCreationResult(uint8 version, bool success, uint8 errorCode, string walletAddress)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterWalletCreationResult(opts *bind.FilterOpts) (*TaskPayloadContractWalletCreationResultIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "WalletCreationResult")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractWalletCreationResultIterator{contract: _TaskPayloadContract.contract, event: "WalletCreationResult", logs: logs, sub: sub}, nil
}

// WatchWalletCreationResult is a free log subscription operation binding the contract event 0x440691550bb1f6d18c60b1a17fff36325a996ba3ab5917f3003445984c5302cf.
//
// Solidity: event WalletCreationResult(uint8 version, bool success, uint8 errorCode, string walletAddress)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchWalletCreationResult(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractWalletCreationResult) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "WalletCreationResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractWalletCreationResult)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "WalletCreationResult", log); err != nil {
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

// ParseWalletCreationResult is a log parse operation binding the contract event 0x440691550bb1f6d18c60b1a17fff36325a996ba3ab5917f3003445984c5302cf.
//
// Solidity: event WalletCreationResult(uint8 version, bool success, uint8 errorCode, string walletAddress)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseWalletCreationResult(log types.Log) (*TaskPayloadContractWalletCreationResult, error) {
	event := new(TaskPayloadContractWalletCreationResult)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "WalletCreationResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractWithdrawalRequestIterator is returned from FilterWithdrawalRequest and is used to iterate over the raw logs and unpacked data for WithdrawalRequest events raised by the TaskPayloadContract contract.
type TaskPayloadContractWithdrawalRequestIterator struct {
	Event *TaskPayloadContractWithdrawalRequest // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractWithdrawalRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractWithdrawalRequest)
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
		it.Event = new(TaskPayloadContractWithdrawalRequest)
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
func (it *TaskPayloadContractWithdrawalRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractWithdrawalRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractWithdrawalRequest represents a WithdrawalRequest event raised by the TaskPayloadContract contract.
type TaskPayloadContractWithdrawalRequest struct {
	UserAddress common.Address
	ChainId     uint64
	Ticker      [32]byte
	ToAddress   string
	Amount      *big.Int
	Offset      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequest is a free log retrieval operation binding the contract event 0xb2a3265f378dfb4ac5a52b79e3e7cb3b836b5aef9974f009b7f00d9ce2c67190.
//
// Solidity: event WithdrawalRequest(address userAddress, uint64 chainId, bytes32 ticker, string toAddress, uint256 amount, uint256 offset)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterWithdrawalRequest(opts *bind.FilterOpts) (*TaskPayloadContractWithdrawalRequestIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "WithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractWithdrawalRequestIterator{contract: _TaskPayloadContract.contract, event: "WithdrawalRequest", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequest is a free log subscription operation binding the contract event 0xb2a3265f378dfb4ac5a52b79e3e7cb3b836b5aef9974f009b7f00d9ce2c67190.
//
// Solidity: event WithdrawalRequest(address userAddress, uint64 chainId, bytes32 ticker, string toAddress, uint256 amount, uint256 offset)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchWithdrawalRequest(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractWithdrawalRequest) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "WithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractWithdrawalRequest)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "WithdrawalRequest", log); err != nil {
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

// ParseWithdrawalRequest is a log parse operation binding the contract event 0xb2a3265f378dfb4ac5a52b79e3e7cb3b836b5aef9974f009b7f00d9ce2c67190.
//
// Solidity: event WithdrawalRequest(address userAddress, uint64 chainId, bytes32 ticker, string toAddress, uint256 amount, uint256 offset)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseWithdrawalRequest(log types.Log) (*TaskPayloadContractWithdrawalRequest, error) {
	event := new(TaskPayloadContractWithdrawalRequest)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "WithdrawalRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractWithdrawalResultIterator is returned from FilterWithdrawalResult and is used to iterate over the raw logs and unpacked data for WithdrawalResult events raised by the TaskPayloadContract contract.
type TaskPayloadContractWithdrawalResultIterator struct {
	Event *TaskPayloadContractWithdrawalResult // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractWithdrawalResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractWithdrawalResult)
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
		it.Event = new(TaskPayloadContractWithdrawalResult)
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
func (it *TaskPayloadContractWithdrawalResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractWithdrawalResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractWithdrawalResult represents a WithdrawalResult event raised by the TaskPayloadContract contract.
type TaskPayloadContractWithdrawalResult struct {
	Version   uint8
	Success   bool
	ErrorCode uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalResult is a free log retrieval operation binding the contract event 0x9a474499969867585df13ccda2ed8f3f9ad89cd1704e038cb941e1fbdc1c08fe.
//
// Solidity: event WithdrawalResult(uint8 version, bool success, uint8 errorCode)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterWithdrawalResult(opts *bind.FilterOpts) (*TaskPayloadContractWithdrawalResultIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "WithdrawalResult")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractWithdrawalResultIterator{contract: _TaskPayloadContract.contract, event: "WithdrawalResult", logs: logs, sub: sub}, nil
}

// WatchWithdrawalResult is a free log subscription operation binding the contract event 0x9a474499969867585df13ccda2ed8f3f9ad89cd1704e038cb941e1fbdc1c08fe.
//
// Solidity: event WithdrawalResult(uint8 version, bool success, uint8 errorCode)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchWithdrawalResult(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractWithdrawalResult) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "WithdrawalResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractWithdrawalResult)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "WithdrawalResult", log); err != nil {
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

// ParseWithdrawalResult is a log parse operation binding the contract event 0x9a474499969867585df13ccda2ed8f3f9ad89cd1704e038cb941e1fbdc1c08fe.
//
// Solidity: event WithdrawalResult(uint8 version, bool success, uint8 errorCode)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseWithdrawalResult(log types.Log) (*TaskPayloadContractWithdrawalResult, error) {
	event := new(TaskPayloadContractWithdrawalResult)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "WithdrawalResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
