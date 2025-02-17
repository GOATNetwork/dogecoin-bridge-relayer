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

// Task is an auto generated low-level Go binding around an user-defined struct.
type Task struct {
	Id        uint64
	State     uint8
	Submitter common.Address
	Handler   common.Address
	CreatedAt uint32
	UpdatedAt uint32
	Result    []byte
}

// TaskManagerContractMetaData contains all meta data concerning the TaskManagerContract contract.
var TaskManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ENTRYPOINT_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestTask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTask\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTask\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTask\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskState\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUncompletedTasks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structTask[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taskHandlers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextCreatedTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitTask\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateTask\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"_result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskSubmitted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"handler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskUpdated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumState\"},{\"name\":\"updateTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EmptyTask\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
}

// TaskManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskManagerContractMetaData.ABI instead.
var TaskManagerContractABI = TaskManagerContractMetaData.ABI

// TaskManagerContract is an auto generated Go binding around an Ethereum contract.
type TaskManagerContract struct {
	TaskManagerContractCaller     // Read-only binding to the contract
	TaskManagerContractTransactor // Write-only binding to the contract
	TaskManagerContractFilterer   // Log filterer for contract events
}

// TaskManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskManagerContractSession struct {
	Contract     *TaskManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TaskManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskManagerContractCallerSession struct {
	Contract *TaskManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TaskManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskManagerContractTransactorSession struct {
	Contract     *TaskManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TaskManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskManagerContractRaw struct {
	Contract *TaskManagerContract // Generic contract binding to access the raw methods on
}

// TaskManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskManagerContractCallerRaw struct {
	Contract *TaskManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// TaskManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskManagerContractTransactorRaw struct {
	Contract *TaskManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskManagerContract creates a new instance of TaskManagerContract, bound to a specific deployed contract.
func NewTaskManagerContract(address common.Address, backend bind.ContractBackend) (*TaskManagerContract, error) {
	contract, err := bindTaskManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContract{TaskManagerContractCaller: TaskManagerContractCaller{contract: contract}, TaskManagerContractTransactor: TaskManagerContractTransactor{contract: contract}, TaskManagerContractFilterer: TaskManagerContractFilterer{contract: contract}}, nil
}

// NewTaskManagerContractCaller creates a new read-only instance of TaskManagerContract, bound to a specific deployed contract.
func NewTaskManagerContractCaller(address common.Address, caller bind.ContractCaller) (*TaskManagerContractCaller, error) {
	contract, err := bindTaskManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractCaller{contract: contract}, nil
}

// NewTaskManagerContractTransactor creates a new write-only instance of TaskManagerContract, bound to a specific deployed contract.
func NewTaskManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskManagerContractTransactor, error) {
	contract, err := bindTaskManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractTransactor{contract: contract}, nil
}

// NewTaskManagerContractFilterer creates a new log filterer instance of TaskManagerContract, bound to a specific deployed contract.
func NewTaskManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskManagerContractFilterer, error) {
	contract, err := bindTaskManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractFilterer{contract: contract}, nil
}

// bindTaskManagerContract binds a generic wrapper to an already deployed contract.
func bindTaskManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskManagerContract *TaskManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskManagerContract.Contract.TaskManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskManagerContract *TaskManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.TaskManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskManagerContract *TaskManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.TaskManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskManagerContract *TaskManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskManagerContract *TaskManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskManagerContract *TaskManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TaskManagerContract.Contract.DEFAULTADMINROLE(&_TaskManagerContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TaskManagerContract.Contract.DEFAULTADMINROLE(&_TaskManagerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractCaller) ENTRYPOINTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "ENTRYPOINT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _TaskManagerContract.Contract.ENTRYPOINTROLE(&_TaskManagerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractCallerSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _TaskManagerContract.Contract.ENTRYPOINTROLE(&_TaskManagerContract.CallOpts)
}

// HANDLERROLE is a free data retrieval call binding the contract method 0x7afa9e9c.
//
// Solidity: function HANDLER_ROLE() view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractCaller) HANDLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "HANDLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HANDLERROLE is a free data retrieval call binding the contract method 0x7afa9e9c.
//
// Solidity: function HANDLER_ROLE() view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractSession) HANDLERROLE() ([32]byte, error) {
	return _TaskManagerContract.Contract.HANDLERROLE(&_TaskManagerContract.CallOpts)
}

// HANDLERROLE is a free data retrieval call binding the contract method 0x7afa9e9c.
//
// Solidity: function HANDLER_ROLE() view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractCallerSession) HANDLERROLE() ([32]byte, error) {
	return _TaskManagerContract.Contract.HANDLERROLE(&_TaskManagerContract.CallOpts)
}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint64,uint8,address,address,uint32,uint32,bytes))
func (_TaskManagerContract *TaskManagerContractCaller) GetLatestTask(opts *bind.CallOpts) (Task, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "getLatestTask")

	if err != nil {
		return *new(Task), err
	}

	out0 := *abi.ConvertType(out[0], new(Task)).(*Task)

	return out0, err

}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint64,uint8,address,address,uint32,uint32,bytes))
func (_TaskManagerContract *TaskManagerContractSession) GetLatestTask() (Task, error) {
	return _TaskManagerContract.Contract.GetLatestTask(&_TaskManagerContract.CallOpts)
}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint64,uint8,address,address,uint32,uint32,bytes))
func (_TaskManagerContract *TaskManagerContractCallerSession) GetLatestTask() (Task, error) {
	return _TaskManagerContract.Contract.GetLatestTask(&_TaskManagerContract.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TaskManagerContract.Contract.GetRoleAdmin(&_TaskManagerContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TaskManagerContract *TaskManagerContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TaskManagerContract.Contract.GetRoleAdmin(&_TaskManagerContract.CallOpts, role)
}

// GetTask is a free data retrieval call binding the contract method 0xae7b4a03.
//
// Solidity: function getTask(uint64 _taskId) view returns((uint64,uint8,address,address,uint32,uint32,bytes))
func (_TaskManagerContract *TaskManagerContractCaller) GetTask(opts *bind.CallOpts, _taskId uint64) (Task, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "getTask", _taskId)

	if err != nil {
		return *new(Task), err
	}

	out0 := *abi.ConvertType(out[0], new(Task)).(*Task)

	return out0, err

}

// GetTask is a free data retrieval call binding the contract method 0xae7b4a03.
//
// Solidity: function getTask(uint64 _taskId) view returns((uint64,uint8,address,address,uint32,uint32,bytes))
func (_TaskManagerContract *TaskManagerContractSession) GetTask(_taskId uint64) (Task, error) {
	return _TaskManagerContract.Contract.GetTask(&_TaskManagerContract.CallOpts, _taskId)
}

// GetTask is a free data retrieval call binding the contract method 0xae7b4a03.
//
// Solidity: function getTask(uint64 _taskId) view returns((uint64,uint8,address,address,uint32,uint32,bytes))
func (_TaskManagerContract *TaskManagerContractCallerSession) GetTask(_taskId uint64) (Task, error) {
	return _TaskManagerContract.Contract.GetTask(&_TaskManagerContract.CallOpts, _taskId)
}

// GetTaskState is a free data retrieval call binding the contract method 0xffcfc50a.
//
// Solidity: function getTaskState(uint64 _taskId) view returns(uint8)
func (_TaskManagerContract *TaskManagerContractCaller) GetTaskState(opts *bind.CallOpts, _taskId uint64) (uint8, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "getTaskState", _taskId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTaskState is a free data retrieval call binding the contract method 0xffcfc50a.
//
// Solidity: function getTaskState(uint64 _taskId) view returns(uint8)
func (_TaskManagerContract *TaskManagerContractSession) GetTaskState(_taskId uint64) (uint8, error) {
	return _TaskManagerContract.Contract.GetTaskState(&_TaskManagerContract.CallOpts, _taskId)
}

// GetTaskState is a free data retrieval call binding the contract method 0xffcfc50a.
//
// Solidity: function getTaskState(uint64 _taskId) view returns(uint8)
func (_TaskManagerContract *TaskManagerContractCallerSession) GetTaskState(_taskId uint64) (uint8, error) {
	return _TaskManagerContract.Contract.GetTaskState(&_TaskManagerContract.CallOpts, _taskId)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint64,uint8,address,address,uint32,uint32,bytes)[])
func (_TaskManagerContract *TaskManagerContractCaller) GetUncompletedTasks(opts *bind.CallOpts) ([]Task, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "getUncompletedTasks")

	if err != nil {
		return *new([]Task), err
	}

	out0 := *abi.ConvertType(out[0], new([]Task)).(*[]Task)

	return out0, err

}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint64,uint8,address,address,uint32,uint32,bytes)[])
func (_TaskManagerContract *TaskManagerContractSession) GetUncompletedTasks() ([]Task, error) {
	return _TaskManagerContract.Contract.GetUncompletedTasks(&_TaskManagerContract.CallOpts)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint64,uint8,address,address,uint32,uint32,bytes)[])
func (_TaskManagerContract *TaskManagerContractCallerSession) GetUncompletedTasks() ([]Task, error) {
	return _TaskManagerContract.Contract.GetUncompletedTasks(&_TaskManagerContract.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TaskManagerContract *TaskManagerContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TaskManagerContract *TaskManagerContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TaskManagerContract.Contract.HasRole(&_TaskManagerContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TaskManagerContract *TaskManagerContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TaskManagerContract.Contract.HasRole(&_TaskManagerContract.CallOpts, role, account)
}

// NextCreatedTaskId is a free data retrieval call binding the contract method 0x4a882c4c.
//
// Solidity: function nextCreatedTaskId() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCaller) NextCreatedTaskId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "nextCreatedTaskId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextCreatedTaskId is a free data retrieval call binding the contract method 0x4a882c4c.
//
// Solidity: function nextCreatedTaskId() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractSession) NextCreatedTaskId() (uint64, error) {
	return _TaskManagerContract.Contract.NextCreatedTaskId(&_TaskManagerContract.CallOpts)
}

// NextCreatedTaskId is a free data retrieval call binding the contract method 0x4a882c4c.
//
// Solidity: function nextCreatedTaskId() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCallerSession) NextCreatedTaskId() (uint64, error) {
	return _TaskManagerContract.Contract.NextCreatedTaskId(&_TaskManagerContract.CallOpts)
}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCaller) NextTaskId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "nextTaskId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractSession) NextTaskId() (uint64, error) {
	return _TaskManagerContract.Contract.NextTaskId(&_TaskManagerContract.CallOpts)
}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCallerSession) NextTaskId() (uint64, error) {
	return _TaskManagerContract.Contract.NextTaskId(&_TaskManagerContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TaskManagerContract *TaskManagerContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TaskManagerContract *TaskManagerContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TaskManagerContract.Contract.SupportsInterface(&_TaskManagerContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TaskManagerContract *TaskManagerContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TaskManagerContract.Contract.SupportsInterface(&_TaskManagerContract.CallOpts, interfaceId)
}

// Tasks is a free data retrieval call binding the contract method 0x6b895ab7.
//
// Solidity: function tasks(uint64 ) view returns(uint64 id, uint8 state, address submitter, address handler, uint32 createdAt, uint32 updatedAt, bytes result)
func (_TaskManagerContract *TaskManagerContractCaller) Tasks(opts *bind.CallOpts, arg0 uint64) (struct {
	Id        uint64
	State     uint8
	Submitter common.Address
	Handler   common.Address
	CreatedAt uint32
	UpdatedAt uint32
	Result    []byte
}, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Id        uint64
		State     uint8
		Submitter common.Address
		Handler   common.Address
		CreatedAt uint32
		UpdatedAt uint32
		Result    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.State = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Submitter = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Handler = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.UpdatedAt = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.Result = *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x6b895ab7.
//
// Solidity: function tasks(uint64 ) view returns(uint64 id, uint8 state, address submitter, address handler, uint32 createdAt, uint32 updatedAt, bytes result)
func (_TaskManagerContract *TaskManagerContractSession) Tasks(arg0 uint64) (struct {
	Id        uint64
	State     uint8
	Submitter common.Address
	Handler   common.Address
	CreatedAt uint32
	UpdatedAt uint32
	Result    []byte
}, error) {
	return _TaskManagerContract.Contract.Tasks(&_TaskManagerContract.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x6b895ab7.
//
// Solidity: function tasks(uint64 ) view returns(uint64 id, uint8 state, address submitter, address handler, uint32 createdAt, uint32 updatedAt, bytes result)
func (_TaskManagerContract *TaskManagerContractCallerSession) Tasks(arg0 uint64) (struct {
	Id        uint64
	State     uint8
	Submitter common.Address
	Handler   common.Address
	CreatedAt uint32
	UpdatedAt uint32
	Result    []byte
}, error) {
	return _TaskManagerContract.Contract.Tasks(&_TaskManagerContract.CallOpts, arg0)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TaskManagerContract *TaskManagerContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.GrantRole(&_TaskManagerContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.GrantRole(&_TaskManagerContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x77a24f36.
//
// Solidity: function initialize(address _owner, address _entryPoint, address[] _taskHandlers) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _entryPoint common.Address, _taskHandlers []common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "initialize", _owner, _entryPoint, _taskHandlers)
}

// Initialize is a paid mutator transaction binding the contract method 0x77a24f36.
//
// Solidity: function initialize(address _owner, address _entryPoint, address[] _taskHandlers) returns()
func (_TaskManagerContract *TaskManagerContractSession) Initialize(_owner common.Address, _entryPoint common.Address, _taskHandlers []common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.Initialize(&_TaskManagerContract.TransactOpts, _owner, _entryPoint, _taskHandlers)
}

// Initialize is a paid mutator transaction binding the contract method 0x77a24f36.
//
// Solidity: function initialize(address _owner, address _entryPoint, address[] _taskHandlers) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) Initialize(_owner common.Address, _entryPoint common.Address, _taskHandlers []common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.Initialize(&_TaskManagerContract.TransactOpts, _owner, _entryPoint, _taskHandlers)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TaskManagerContract *TaskManagerContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.RenounceRole(&_TaskManagerContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.RenounceRole(&_TaskManagerContract.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TaskManagerContract *TaskManagerContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.RevokeRole(&_TaskManagerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.RevokeRole(&_TaskManagerContract.TransactOpts, role, account)
}

// SubmitTask is a paid mutator transaction binding the contract method 0xe2a944e6.
//
// Solidity: function submitTask(address _submitter, bytes _data) returns(uint64)
func (_TaskManagerContract *TaskManagerContractTransactor) SubmitTask(opts *bind.TransactOpts, _submitter common.Address, _data []byte) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "submitTask", _submitter, _data)
}

// SubmitTask is a paid mutator transaction binding the contract method 0xe2a944e6.
//
// Solidity: function submitTask(address _submitter, bytes _data) returns(uint64)
func (_TaskManagerContract *TaskManagerContractSession) SubmitTask(_submitter common.Address, _data []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.SubmitTask(&_TaskManagerContract.TransactOpts, _submitter, _data)
}

// SubmitTask is a paid mutator transaction binding the contract method 0xe2a944e6.
//
// Solidity: function submitTask(address _submitter, bytes _data) returns(uint64)
func (_TaskManagerContract *TaskManagerContractTransactorSession) SubmitTask(_submitter common.Address, _data []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.SubmitTask(&_TaskManagerContract.TransactOpts, _submitter, _data)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x875d4e7d.
//
// Solidity: function updateTask(uint64 _taskId, uint8 _state, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) UpdateTask(opts *bind.TransactOpts, _taskId uint64, _state uint8, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "updateTask", _taskId, _state, _result)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x875d4e7d.
//
// Solidity: function updateTask(uint64 _taskId, uint8 _state, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractSession) UpdateTask(_taskId uint64, _state uint8, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.UpdateTask(&_TaskManagerContract.TransactOpts, _taskId, _state, _result)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x875d4e7d.
//
// Solidity: function updateTask(uint64 _taskId, uint8 _state, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) UpdateTask(_taskId uint64, _state uint8, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.UpdateTask(&_TaskManagerContract.TransactOpts, _taskId, _state, _result)
}

// TaskManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaskManagerContract contract.
type TaskManagerContractInitializedIterator struct {
	Event *TaskManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractInitialized)
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
		it.Event = new(TaskManagerContractInitialized)
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
func (it *TaskManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractInitialized represents a Initialized event raised by the TaskManagerContract contract.
type TaskManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaskManagerContractInitializedIterator, error) {

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractInitializedIterator{contract: _TaskManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaskManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractInitialized)
				if err := _TaskManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaskManagerContract *TaskManagerContractFilterer) ParseInitialized(log types.Log) (*TaskManagerContractInitialized, error) {
	event := new(TaskManagerContractInitialized)
	if err := _TaskManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskManagerContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TaskManagerContract contract.
type TaskManagerContractRoleAdminChangedIterator struct {
	Event *TaskManagerContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractRoleAdminChanged)
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
		it.Event = new(TaskManagerContractRoleAdminChanged)
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
func (it *TaskManagerContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractRoleAdminChanged represents a RoleAdminChanged event raised by the TaskManagerContract contract.
type TaskManagerContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TaskManagerContractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractRoleAdminChangedIterator{contract: _TaskManagerContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TaskManagerContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractRoleAdminChanged)
				if err := _TaskManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TaskManagerContract *TaskManagerContractFilterer) ParseRoleAdminChanged(log types.Log) (*TaskManagerContractRoleAdminChanged, error) {
	event := new(TaskManagerContractRoleAdminChanged)
	if err := _TaskManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskManagerContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TaskManagerContract contract.
type TaskManagerContractRoleGrantedIterator struct {
	Event *TaskManagerContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractRoleGranted)
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
		it.Event = new(TaskManagerContractRoleGranted)
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
func (it *TaskManagerContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractRoleGranted represents a RoleGranted event raised by the TaskManagerContract contract.
type TaskManagerContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TaskManagerContractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractRoleGrantedIterator{contract: _TaskManagerContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TaskManagerContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractRoleGranted)
				if err := _TaskManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TaskManagerContract *TaskManagerContractFilterer) ParseRoleGranted(log types.Log) (*TaskManagerContractRoleGranted, error) {
	event := new(TaskManagerContractRoleGranted)
	if err := _TaskManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskManagerContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TaskManagerContract contract.
type TaskManagerContractRoleRevokedIterator struct {
	Event *TaskManagerContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractRoleRevoked)
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
		it.Event = new(TaskManagerContractRoleRevoked)
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
func (it *TaskManagerContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractRoleRevoked represents a RoleRevoked event raised by the TaskManagerContract contract.
type TaskManagerContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TaskManagerContractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractRoleRevokedIterator{contract: _TaskManagerContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TaskManagerContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractRoleRevoked)
				if err := _TaskManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TaskManagerContract *TaskManagerContractFilterer) ParseRoleRevoked(log types.Log) (*TaskManagerContractRoleRevoked, error) {
	event := new(TaskManagerContractRoleRevoked)
	if err := _TaskManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskManagerContractTaskSubmittedIterator is returned from FilterTaskSubmitted and is used to iterate over the raw logs and unpacked data for TaskSubmitted events raised by the TaskManagerContract contract.
type TaskManagerContractTaskSubmittedIterator struct {
	Event *TaskManagerContractTaskSubmitted // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractTaskSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractTaskSubmitted)
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
		it.Event = new(TaskManagerContractTaskSubmitted)
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
func (it *TaskManagerContractTaskSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractTaskSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractTaskSubmitted represents a TaskSubmitted event raised by the TaskManagerContract contract.
type TaskManagerContractTaskSubmitted struct {
	TaskId    uint64
	Submitter common.Address
	Handler   common.Address
	CallData  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0xc43825e1299c5fbb71e3a2b57d3a8d8178eaddc032e876d32dc37514160bd376.
//
// Solidity: event TaskSubmitted(uint64 indexed taskId, address indexed submitter, address indexed handler, bytes callData)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterTaskSubmitted(opts *bind.FilterOpts, taskId []uint64, submitter []common.Address, handler []common.Address) (*TaskManagerContractTaskSubmittedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "TaskSubmitted", taskIdRule, submitterRule, handlerRule)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractTaskSubmittedIterator{contract: _TaskManagerContract.contract, event: "TaskSubmitted", logs: logs, sub: sub}, nil
}

// WatchTaskSubmitted is a free log subscription operation binding the contract event 0xc43825e1299c5fbb71e3a2b57d3a8d8178eaddc032e876d32dc37514160bd376.
//
// Solidity: event TaskSubmitted(uint64 indexed taskId, address indexed submitter, address indexed handler, bytes callData)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *TaskManagerContractTaskSubmitted, taskId []uint64, submitter []common.Address, handler []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "TaskSubmitted", taskIdRule, submitterRule, handlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractTaskSubmitted)
				if err := _TaskManagerContract.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
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

// ParseTaskSubmitted is a log parse operation binding the contract event 0xc43825e1299c5fbb71e3a2b57d3a8d8178eaddc032e876d32dc37514160bd376.
//
// Solidity: event TaskSubmitted(uint64 indexed taskId, address indexed submitter, address indexed handler, bytes callData)
func (_TaskManagerContract *TaskManagerContractFilterer) ParseTaskSubmitted(log types.Log) (*TaskManagerContractTaskSubmitted, error) {
	event := new(TaskManagerContractTaskSubmitted)
	if err := _TaskManagerContract.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskManagerContractTaskUpdatedIterator is returned from FilterTaskUpdated and is used to iterate over the raw logs and unpacked data for TaskUpdated events raised by the TaskManagerContract contract.
type TaskManagerContractTaskUpdatedIterator struct {
	Event *TaskManagerContractTaskUpdated // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractTaskUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractTaskUpdated)
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
		it.Event = new(TaskManagerContractTaskUpdated)
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
func (it *TaskManagerContractTaskUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractTaskUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractTaskUpdated represents a TaskUpdated event raised by the TaskManagerContract contract.
type TaskManagerContractTaskUpdated struct {
	TaskId     uint64
	Submitter  common.Address
	State      uint8
	UpdateTime *big.Int
	Result     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskUpdated is a free log retrieval operation binding the contract event 0x3357e4d22dded01b8f5952481c84ccbcd9818d31adbb1bb1e8df6557f58e1ab7.
//
// Solidity: event TaskUpdated(uint64 indexed taskId, address indexed submitter, uint8 indexed state, uint256 updateTime, bytes result)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterTaskUpdated(opts *bind.FilterOpts, taskId []uint64, submitter []common.Address, state []uint8) (*TaskManagerContractTaskUpdatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "TaskUpdated", taskIdRule, submitterRule, stateRule)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractTaskUpdatedIterator{contract: _TaskManagerContract.contract, event: "TaskUpdated", logs: logs, sub: sub}, nil
}

// WatchTaskUpdated is a free log subscription operation binding the contract event 0x3357e4d22dded01b8f5952481c84ccbcd9818d31adbb1bb1e8df6557f58e1ab7.
//
// Solidity: event TaskUpdated(uint64 indexed taskId, address indexed submitter, uint8 indexed state, uint256 updateTime, bytes result)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchTaskUpdated(opts *bind.WatchOpts, sink chan<- *TaskManagerContractTaskUpdated, taskId []uint64, submitter []common.Address, state []uint8) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "TaskUpdated", taskIdRule, submitterRule, stateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractTaskUpdated)
				if err := _TaskManagerContract.contract.UnpackLog(event, "TaskUpdated", log); err != nil {
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

// ParseTaskUpdated is a log parse operation binding the contract event 0x3357e4d22dded01b8f5952481c84ccbcd9818d31adbb1bb1e8df6557f58e1ab7.
//
// Solidity: event TaskUpdated(uint64 indexed taskId, address indexed submitter, uint8 indexed state, uint256 updateTime, bytes result)
func (_TaskManagerContract *TaskManagerContractFilterer) ParseTaskUpdated(log types.Log) (*TaskManagerContractTaskUpdated, error) {
	event := new(TaskManagerContractTaskUpdated)
	if err := _TaskManagerContract.contract.UnpackLog(event, "TaskUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
