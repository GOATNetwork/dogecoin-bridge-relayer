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

// ParticipantManagerContractMetaData contains all meta data concerning the ParticipantManagerContract contract.
var ParticipantManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_nuvoLock\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ENTRYPOINT_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SUBMITTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addParticipant\",\"inputs\":[{\"name\":\"_newParticipant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getParticipants\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRandomParticipant\",\"inputs\":[{\"name\":\"_salt\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"randParticipant\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialParticipants\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isParticipant\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nuvoLock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINuvoLock\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"participants\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeParticipant\",\"inputs\":[{\"name\":\"_participant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetParticipants\",\"inputs\":[{\"name\":\"_newParticipants\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitAddParticipantTask\",\"inputs\":[{\"name\":\"_newParticipant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRemoveParticipantTask\",\"inputs\":[{\"name\":\"_participant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitResetParticipantsTask\",\"inputs\":[{\"name\":\"_newParticipants\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParticipantAdded\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParticipantRemoved\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParticipantsReset\",\"inputs\":[{\"name\":\"participants\",\"type\":\"address[]\",\"indexed\":true,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestAddParticipant\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"participants\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestRemoveParticipant\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"participants\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestResetParticipants\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"participants\",\"type\":\"address[]\",\"indexed\":true,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyParticipant\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotEnoughParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// ParticipantManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ParticipantManagerContractMetaData.ABI instead.
var ParticipantManagerContractABI = ParticipantManagerContractMetaData.ABI

// ParticipantManagerContract is an auto generated Go binding around an Ethereum contract.
type ParticipantManagerContract struct {
	ParticipantManagerContractCaller     // Read-only binding to the contract
	ParticipantManagerContractTransactor // Write-only binding to the contract
	ParticipantManagerContractFilterer   // Log filterer for contract events
}

// ParticipantManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParticipantManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParticipantManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParticipantManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParticipantManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParticipantManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParticipantManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParticipantManagerContractSession struct {
	Contract     *ParticipantManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ParticipantManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParticipantManagerContractCallerSession struct {
	Contract *ParticipantManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ParticipantManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParticipantManagerContractTransactorSession struct {
	Contract     *ParticipantManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ParticipantManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParticipantManagerContractRaw struct {
	Contract *ParticipantManagerContract // Generic contract binding to access the raw methods on
}

// ParticipantManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParticipantManagerContractCallerRaw struct {
	Contract *ParticipantManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// ParticipantManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParticipantManagerContractTransactorRaw struct {
	Contract *ParticipantManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParticipantManagerContract creates a new instance of ParticipantManagerContract, bound to a specific deployed contract.
func NewParticipantManagerContract(address common.Address, backend bind.ContractBackend) (*ParticipantManagerContract, error) {
	contract, err := bindParticipantManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContract{ParticipantManagerContractCaller: ParticipantManagerContractCaller{contract: contract}, ParticipantManagerContractTransactor: ParticipantManagerContractTransactor{contract: contract}, ParticipantManagerContractFilterer: ParticipantManagerContractFilterer{contract: contract}}, nil
}

// NewParticipantManagerContractCaller creates a new read-only instance of ParticipantManagerContract, bound to a specific deployed contract.
func NewParticipantManagerContractCaller(address common.Address, caller bind.ContractCaller) (*ParticipantManagerContractCaller, error) {
	contract, err := bindParticipantManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractCaller{contract: contract}, nil
}

// NewParticipantManagerContractTransactor creates a new write-only instance of ParticipantManagerContract, bound to a specific deployed contract.
func NewParticipantManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ParticipantManagerContractTransactor, error) {
	contract, err := bindParticipantManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractTransactor{contract: contract}, nil
}

// NewParticipantManagerContractFilterer creates a new log filterer instance of ParticipantManagerContract, bound to a specific deployed contract.
func NewParticipantManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ParticipantManagerContractFilterer, error) {
	contract, err := bindParticipantManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractFilterer{contract: contract}, nil
}

// bindParticipantManagerContract binds a generic wrapper to an already deployed contract.
func bindParticipantManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParticipantManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParticipantManagerContract *ParticipantManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParticipantManagerContract.Contract.ParticipantManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParticipantManagerContract *ParticipantManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.ParticipantManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParticipantManagerContract *ParticipantManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.ParticipantManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParticipantManagerContract *ParticipantManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParticipantManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParticipantManagerContract *ParticipantManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParticipantManagerContract *ParticipantManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ParticipantManagerContract.Contract.DEFAULTADMINROLE(&_ParticipantManagerContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ParticipantManagerContract.Contract.DEFAULTADMINROLE(&_ParticipantManagerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) ENTRYPOINTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "ENTRYPOINT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _ParticipantManagerContract.Contract.ENTRYPOINTROLE(&_ParticipantManagerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _ParticipantManagerContract.Contract.ENTRYPOINTROLE(&_ParticipantManagerContract.CallOpts)
}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) SUBMITTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "SUBMITTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractSession) SUBMITTERROLE() ([32]byte, error) {
	return _ParticipantManagerContract.Contract.SUBMITTERROLE(&_ParticipantManagerContract.CallOpts)
}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) SUBMITTERROLE() ([32]byte, error) {
	return _ParticipantManagerContract.Contract.SUBMITTERROLE(&_ParticipantManagerContract.CallOpts)
}

// GetParticipants is a free data retrieval call binding the contract method 0x5aa68ac0.
//
// Solidity: function getParticipants() view returns(address[])
func (_ParticipantManagerContract *ParticipantManagerContractCaller) GetParticipants(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "getParticipants")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetParticipants is a free data retrieval call binding the contract method 0x5aa68ac0.
//
// Solidity: function getParticipants() view returns(address[])
func (_ParticipantManagerContract *ParticipantManagerContractSession) GetParticipants() ([]common.Address, error) {
	return _ParticipantManagerContract.Contract.GetParticipants(&_ParticipantManagerContract.CallOpts)
}

// GetParticipants is a free data retrieval call binding the contract method 0x5aa68ac0.
//
// Solidity: function getParticipants() view returns(address[])
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) GetParticipants() ([]common.Address, error) {
	return _ParticipantManagerContract.Contract.GetParticipants(&_ParticipantManagerContract.CallOpts)
}

// GetRandomParticipant is a free data retrieval call binding the contract method 0x6235b415.
//
// Solidity: function getRandomParticipant(address _salt) view returns(address randParticipant)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) GetRandomParticipant(opts *bind.CallOpts, _salt common.Address) (common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "getRandomParticipant", _salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRandomParticipant is a free data retrieval call binding the contract method 0x6235b415.
//
// Solidity: function getRandomParticipant(address _salt) view returns(address randParticipant)
func (_ParticipantManagerContract *ParticipantManagerContractSession) GetRandomParticipant(_salt common.Address) (common.Address, error) {
	return _ParticipantManagerContract.Contract.GetRandomParticipant(&_ParticipantManagerContract.CallOpts, _salt)
}

// GetRandomParticipant is a free data retrieval call binding the contract method 0x6235b415.
//
// Solidity: function getRandomParticipant(address _salt) view returns(address randParticipant)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) GetRandomParticipant(_salt common.Address) (common.Address, error) {
	return _ParticipantManagerContract.Contract.GetRandomParticipant(&_ParticipantManagerContract.CallOpts, _salt)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ParticipantManagerContract.Contract.GetRoleAdmin(&_ParticipantManagerContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ParticipantManagerContract.Contract.GetRoleAdmin(&_ParticipantManagerContract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ParticipantManagerContract.Contract.HasRole(&_ParticipantManagerContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ParticipantManagerContract.Contract.HasRole(&_ParticipantManagerContract.CallOpts, role, account)
}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address ) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) IsParticipant(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "isParticipant", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address ) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractSession) IsParticipant(arg0 common.Address) (bool, error) {
	return _ParticipantManagerContract.Contract.IsParticipant(&_ParticipantManagerContract.CallOpts, arg0)
}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address ) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) IsParticipant(arg0 common.Address) (bool, error) {
	return _ParticipantManagerContract.Contract.IsParticipant(&_ParticipantManagerContract.CallOpts, arg0)
}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) NuvoLock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "nuvoLock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractSession) NuvoLock() (common.Address, error) {
	return _ParticipantManagerContract.Contract.NuvoLock(&_ParticipantManagerContract.CallOpts)
}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) NuvoLock() (common.Address, error) {
	return _ParticipantManagerContract.Contract.NuvoLock(&_ParticipantManagerContract.CallOpts)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "participants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _ParticipantManagerContract.Contract.Participants(&_ParticipantManagerContract.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _ParticipantManagerContract.Contract.Participants(&_ParticipantManagerContract.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ParticipantManagerContract.Contract.SupportsInterface(&_ParticipantManagerContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ParticipantManagerContract.Contract.SupportsInterface(&_ParticipantManagerContract.CallOpts, interfaceId)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractSession) TaskManager() (common.Address, error) {
	return _ParticipantManagerContract.Contract.TaskManager(&_ParticipantManagerContract.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) TaskManager() (common.Address, error) {
	return _ParticipantManagerContract.Contract.TaskManager(&_ParticipantManagerContract.CallOpts)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address _newParticipant) returns(bytes)
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) AddParticipant(opts *bind.TransactOpts, _newParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "addParticipant", _newParticipant)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address _newParticipant) returns(bytes)
func (_ParticipantManagerContract *ParticipantManagerContractSession) AddParticipant(_newParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.AddParticipant(&_ParticipantManagerContract.TransactOpts, _newParticipant)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address _newParticipant) returns(bytes)
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) AddParticipant(_newParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.AddParticipant(&_ParticipantManagerContract.TransactOpts, _newParticipant)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ParticipantManagerContract *ParticipantManagerContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.GrantRole(&_ParticipantManagerContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.GrantRole(&_ParticipantManagerContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xe6bfbfd8.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter, address[] _initialParticipants) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _entryPoint common.Address, _submitter common.Address, _initialParticipants []common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "initialize", _owner, _entryPoint, _submitter, _initialParticipants)
}

// Initialize is a paid mutator transaction binding the contract method 0xe6bfbfd8.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter, address[] _initialParticipants) returns()
func (_ParticipantManagerContract *ParticipantManagerContractSession) Initialize(_owner common.Address, _entryPoint common.Address, _submitter common.Address, _initialParticipants []common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.Initialize(&_ParticipantManagerContract.TransactOpts, _owner, _entryPoint, _submitter, _initialParticipants)
}

// Initialize is a paid mutator transaction binding the contract method 0xe6bfbfd8.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter, address[] _initialParticipants) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) Initialize(_owner common.Address, _entryPoint common.Address, _submitter common.Address, _initialParticipants []common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.Initialize(&_ParticipantManagerContract.TransactOpts, _owner, _entryPoint, _submitter, _initialParticipants)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0x668a2001.
//
// Solidity: function removeParticipant(address _participant) returns(bytes)
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) RemoveParticipant(opts *bind.TransactOpts, _participant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "removeParticipant", _participant)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0x668a2001.
//
// Solidity: function removeParticipant(address _participant) returns(bytes)
func (_ParticipantManagerContract *ParticipantManagerContractSession) RemoveParticipant(_participant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RemoveParticipant(&_ParticipantManagerContract.TransactOpts, _participant)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0x668a2001.
//
// Solidity: function removeParticipant(address _participant) returns(bytes)
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) RemoveParticipant(_participant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RemoveParticipant(&_ParticipantManagerContract.TransactOpts, _participant)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ParticipantManagerContract *ParticipantManagerContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RenounceRole(&_ParticipantManagerContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RenounceRole(&_ParticipantManagerContract.TransactOpts, role, callerConfirmation)
}

// ResetParticipants is a paid mutator transaction binding the contract method 0x15c47d12.
//
// Solidity: function resetParticipants(address[] _newParticipants) returns(bytes)
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) ResetParticipants(opts *bind.TransactOpts, _newParticipants []common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "resetParticipants", _newParticipants)
}

// ResetParticipants is a paid mutator transaction binding the contract method 0x15c47d12.
//
// Solidity: function resetParticipants(address[] _newParticipants) returns(bytes)
func (_ParticipantManagerContract *ParticipantManagerContractSession) ResetParticipants(_newParticipants []common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.ResetParticipants(&_ParticipantManagerContract.TransactOpts, _newParticipants)
}

// ResetParticipants is a paid mutator transaction binding the contract method 0x15c47d12.
//
// Solidity: function resetParticipants(address[] _newParticipants) returns(bytes)
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) ResetParticipants(_newParticipants []common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.ResetParticipants(&_ParticipantManagerContract.TransactOpts, _newParticipants)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ParticipantManagerContract *ParticipantManagerContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RevokeRole(&_ParticipantManagerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RevokeRole(&_ParticipantManagerContract.TransactOpts, role, account)
}

// SubmitAddParticipantTask is a paid mutator transaction binding the contract method 0xf8ecc362.
//
// Solidity: function submitAddParticipantTask(address _newParticipant) returns(uint64 taskId)
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) SubmitAddParticipantTask(opts *bind.TransactOpts, _newParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "submitAddParticipantTask", _newParticipant)
}

// SubmitAddParticipantTask is a paid mutator transaction binding the contract method 0xf8ecc362.
//
// Solidity: function submitAddParticipantTask(address _newParticipant) returns(uint64 taskId)
func (_ParticipantManagerContract *ParticipantManagerContractSession) SubmitAddParticipantTask(_newParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.SubmitAddParticipantTask(&_ParticipantManagerContract.TransactOpts, _newParticipant)
}

// SubmitAddParticipantTask is a paid mutator transaction binding the contract method 0xf8ecc362.
//
// Solidity: function submitAddParticipantTask(address _newParticipant) returns(uint64 taskId)
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) SubmitAddParticipantTask(_newParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.SubmitAddParticipantTask(&_ParticipantManagerContract.TransactOpts, _newParticipant)
}

// SubmitRemoveParticipantTask is a paid mutator transaction binding the contract method 0xb82727f3.
//
// Solidity: function submitRemoveParticipantTask(address _participant) returns(uint64 taskId)
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) SubmitRemoveParticipantTask(opts *bind.TransactOpts, _participant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "submitRemoveParticipantTask", _participant)
}

// SubmitRemoveParticipantTask is a paid mutator transaction binding the contract method 0xb82727f3.
//
// Solidity: function submitRemoveParticipantTask(address _participant) returns(uint64 taskId)
func (_ParticipantManagerContract *ParticipantManagerContractSession) SubmitRemoveParticipantTask(_participant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.SubmitRemoveParticipantTask(&_ParticipantManagerContract.TransactOpts, _participant)
}

// SubmitRemoveParticipantTask is a paid mutator transaction binding the contract method 0xb82727f3.
//
// Solidity: function submitRemoveParticipantTask(address _participant) returns(uint64 taskId)
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) SubmitRemoveParticipantTask(_participant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.SubmitRemoveParticipantTask(&_ParticipantManagerContract.TransactOpts, _participant)
}

// SubmitResetParticipantsTask is a paid mutator transaction binding the contract method 0x2b40e050.
//
// Solidity: function submitResetParticipantsTask(address[] _newParticipants) returns(uint64 taskId)
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) SubmitResetParticipantsTask(opts *bind.TransactOpts, _newParticipants []common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "submitResetParticipantsTask", _newParticipants)
}

// SubmitResetParticipantsTask is a paid mutator transaction binding the contract method 0x2b40e050.
//
// Solidity: function submitResetParticipantsTask(address[] _newParticipants) returns(uint64 taskId)
func (_ParticipantManagerContract *ParticipantManagerContractSession) SubmitResetParticipantsTask(_newParticipants []common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.SubmitResetParticipantsTask(&_ParticipantManagerContract.TransactOpts, _newParticipants)
}

// SubmitResetParticipantsTask is a paid mutator transaction binding the contract method 0x2b40e050.
//
// Solidity: function submitResetParticipantsTask(address[] _newParticipants) returns(uint64 taskId)
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) SubmitResetParticipantsTask(_newParticipants []common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.SubmitResetParticipantsTask(&_ParticipantManagerContract.TransactOpts, _newParticipants)
}

// ParticipantManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractInitializedIterator struct {
	Event *ParticipantManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractInitialized)
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
		it.Event = new(ParticipantManagerContractInitialized)
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
func (it *ParticipantManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractInitialized represents a Initialized event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ParticipantManagerContractInitializedIterator, error) {

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractInitializedIterator{contract: _ParticipantManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractInitialized)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseInitialized(log types.Log) (*ParticipantManagerContractInitialized, error) {
	event := new(ParticipantManagerContractInitialized)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractParticipantAddedIterator is returned from FilterParticipantAdded and is used to iterate over the raw logs and unpacked data for ParticipantAdded events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantAddedIterator struct {
	Event *ParticipantManagerContractParticipantAdded // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractParticipantAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractParticipantAdded)
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
		it.Event = new(ParticipantManagerContractParticipantAdded)
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
func (it *ParticipantManagerContractParticipantAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractParticipantAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractParticipantAdded represents a ParticipantAdded event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantAdded struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantAdded is a free log retrieval operation binding the contract event 0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b.
//
// Solidity: event ParticipantAdded(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterParticipantAdded(opts *bind.FilterOpts, participant []common.Address) (*ParticipantManagerContractParticipantAddedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "ParticipantAdded", participantRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractParticipantAddedIterator{contract: _ParticipantManagerContract.contract, event: "ParticipantAdded", logs: logs, sub: sub}, nil
}

// WatchParticipantAdded is a free log subscription operation binding the contract event 0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b.
//
// Solidity: event ParticipantAdded(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchParticipantAdded(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractParticipantAdded, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "ParticipantAdded", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractParticipantAdded)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantAdded", log); err != nil {
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

// ParseParticipantAdded is a log parse operation binding the contract event 0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b.
//
// Solidity: event ParticipantAdded(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseParticipantAdded(log types.Log) (*ParticipantManagerContractParticipantAdded, error) {
	event := new(ParticipantManagerContractParticipantAdded)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractParticipantRemovedIterator is returned from FilterParticipantRemoved and is used to iterate over the raw logs and unpacked data for ParticipantRemoved events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantRemovedIterator struct {
	Event *ParticipantManagerContractParticipantRemoved // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractParticipantRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractParticipantRemoved)
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
		it.Event = new(ParticipantManagerContractParticipantRemoved)
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
func (it *ParticipantManagerContractParticipantRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractParticipantRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractParticipantRemoved represents a ParticipantRemoved event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantRemoved struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantRemoved is a free log retrieval operation binding the contract event 0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc.
//
// Solidity: event ParticipantRemoved(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterParticipantRemoved(opts *bind.FilterOpts, participant []common.Address) (*ParticipantManagerContractParticipantRemovedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "ParticipantRemoved", participantRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractParticipantRemovedIterator{contract: _ParticipantManagerContract.contract, event: "ParticipantRemoved", logs: logs, sub: sub}, nil
}

// WatchParticipantRemoved is a free log subscription operation binding the contract event 0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc.
//
// Solidity: event ParticipantRemoved(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchParticipantRemoved(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractParticipantRemoved, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "ParticipantRemoved", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractParticipantRemoved)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantRemoved", log); err != nil {
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

// ParseParticipantRemoved is a log parse operation binding the contract event 0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc.
//
// Solidity: event ParticipantRemoved(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseParticipantRemoved(log types.Log) (*ParticipantManagerContractParticipantRemoved, error) {
	event := new(ParticipantManagerContractParticipantRemoved)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractParticipantsResetIterator is returned from FilterParticipantsReset and is used to iterate over the raw logs and unpacked data for ParticipantsReset events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantsResetIterator struct {
	Event *ParticipantManagerContractParticipantsReset // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractParticipantsResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractParticipantsReset)
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
		it.Event = new(ParticipantManagerContractParticipantsReset)
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
func (it *ParticipantManagerContractParticipantsResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractParticipantsResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractParticipantsReset represents a ParticipantsReset event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantsReset struct {
	Participants []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterParticipantsReset is a free log retrieval operation binding the contract event 0x32e9d8d19fb1e71c8dc610e5f45fd7f1e2f81babf8ea90e267475a708e09c35e.
//
// Solidity: event ParticipantsReset(address[] indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterParticipantsReset(opts *bind.FilterOpts, participants [][]common.Address) (*ParticipantManagerContractParticipantsResetIterator, error) {

	var participantsRule []interface{}
	for _, participantsItem := range participants {
		participantsRule = append(participantsRule, participantsItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "ParticipantsReset", participantsRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractParticipantsResetIterator{contract: _ParticipantManagerContract.contract, event: "ParticipantsReset", logs: logs, sub: sub}, nil
}

// WatchParticipantsReset is a free log subscription operation binding the contract event 0x32e9d8d19fb1e71c8dc610e5f45fd7f1e2f81babf8ea90e267475a708e09c35e.
//
// Solidity: event ParticipantsReset(address[] indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchParticipantsReset(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractParticipantsReset, participants [][]common.Address) (event.Subscription, error) {

	var participantsRule []interface{}
	for _, participantsItem := range participants {
		participantsRule = append(participantsRule, participantsItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "ParticipantsReset", participantsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractParticipantsReset)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantsReset", log); err != nil {
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

// ParseParticipantsReset is a log parse operation binding the contract event 0x32e9d8d19fb1e71c8dc610e5f45fd7f1e2f81babf8ea90e267475a708e09c35e.
//
// Solidity: event ParticipantsReset(address[] indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseParticipantsReset(log types.Log) (*ParticipantManagerContractParticipantsReset, error) {
	event := new(ParticipantManagerContractParticipantsReset)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantsReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractRequestAddParticipantIterator is returned from FilterRequestAddParticipant and is used to iterate over the raw logs and unpacked data for RequestAddParticipant events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRequestAddParticipantIterator struct {
	Event *ParticipantManagerContractRequestAddParticipant // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractRequestAddParticipantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractRequestAddParticipant)
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
		it.Event = new(ParticipantManagerContractRequestAddParticipant)
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
func (it *ParticipantManagerContractRequestAddParticipantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractRequestAddParticipantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractRequestAddParticipant represents a RequestAddParticipant event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRequestAddParticipant struct {
	TaskId       uint64
	Participants common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestAddParticipant is a free log retrieval operation binding the contract event 0xa38a906b00a07b0bae891d2bd0e45e461c09d7bb442f90ea1b2c83fbdf69aa8a.
//
// Solidity: event RequestAddParticipant(uint64 taskId, address indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterRequestAddParticipant(opts *bind.FilterOpts, participants []common.Address) (*ParticipantManagerContractRequestAddParticipantIterator, error) {

	var participantsRule []interface{}
	for _, participantsItem := range participants {
		participantsRule = append(participantsRule, participantsItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "RequestAddParticipant", participantsRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractRequestAddParticipantIterator{contract: _ParticipantManagerContract.contract, event: "RequestAddParticipant", logs: logs, sub: sub}, nil
}

// WatchRequestAddParticipant is a free log subscription operation binding the contract event 0xa38a906b00a07b0bae891d2bd0e45e461c09d7bb442f90ea1b2c83fbdf69aa8a.
//
// Solidity: event RequestAddParticipant(uint64 taskId, address indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchRequestAddParticipant(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractRequestAddParticipant, participants []common.Address) (event.Subscription, error) {

	var participantsRule []interface{}
	for _, participantsItem := range participants {
		participantsRule = append(participantsRule, participantsItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "RequestAddParticipant", participantsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractRequestAddParticipant)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "RequestAddParticipant", log); err != nil {
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

// ParseRequestAddParticipant is a log parse operation binding the contract event 0xa38a906b00a07b0bae891d2bd0e45e461c09d7bb442f90ea1b2c83fbdf69aa8a.
//
// Solidity: event RequestAddParticipant(uint64 taskId, address indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseRequestAddParticipant(log types.Log) (*ParticipantManagerContractRequestAddParticipant, error) {
	event := new(ParticipantManagerContractRequestAddParticipant)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "RequestAddParticipant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractRequestRemoveParticipantIterator is returned from FilterRequestRemoveParticipant and is used to iterate over the raw logs and unpacked data for RequestRemoveParticipant events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRequestRemoveParticipantIterator struct {
	Event *ParticipantManagerContractRequestRemoveParticipant // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractRequestRemoveParticipantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractRequestRemoveParticipant)
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
		it.Event = new(ParticipantManagerContractRequestRemoveParticipant)
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
func (it *ParticipantManagerContractRequestRemoveParticipantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractRequestRemoveParticipantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractRequestRemoveParticipant represents a RequestRemoveParticipant event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRequestRemoveParticipant struct {
	TaskId       uint64
	Participants common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestRemoveParticipant is a free log retrieval operation binding the contract event 0xa702c6c08114153641251bbaa959f68ed4350162994c98b82f1dde49923826ed.
//
// Solidity: event RequestRemoveParticipant(uint64 taskId, address indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterRequestRemoveParticipant(opts *bind.FilterOpts, participants []common.Address) (*ParticipantManagerContractRequestRemoveParticipantIterator, error) {

	var participantsRule []interface{}
	for _, participantsItem := range participants {
		participantsRule = append(participantsRule, participantsItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "RequestRemoveParticipant", participantsRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractRequestRemoveParticipantIterator{contract: _ParticipantManagerContract.contract, event: "RequestRemoveParticipant", logs: logs, sub: sub}, nil
}

// WatchRequestRemoveParticipant is a free log subscription operation binding the contract event 0xa702c6c08114153641251bbaa959f68ed4350162994c98b82f1dde49923826ed.
//
// Solidity: event RequestRemoveParticipant(uint64 taskId, address indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchRequestRemoveParticipant(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractRequestRemoveParticipant, participants []common.Address) (event.Subscription, error) {

	var participantsRule []interface{}
	for _, participantsItem := range participants {
		participantsRule = append(participantsRule, participantsItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "RequestRemoveParticipant", participantsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractRequestRemoveParticipant)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "RequestRemoveParticipant", log); err != nil {
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

// ParseRequestRemoveParticipant is a log parse operation binding the contract event 0xa702c6c08114153641251bbaa959f68ed4350162994c98b82f1dde49923826ed.
//
// Solidity: event RequestRemoveParticipant(uint64 taskId, address indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseRequestRemoveParticipant(log types.Log) (*ParticipantManagerContractRequestRemoveParticipant, error) {
	event := new(ParticipantManagerContractRequestRemoveParticipant)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "RequestRemoveParticipant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractRequestResetParticipantsIterator is returned from FilterRequestResetParticipants and is used to iterate over the raw logs and unpacked data for RequestResetParticipants events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRequestResetParticipantsIterator struct {
	Event *ParticipantManagerContractRequestResetParticipants // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractRequestResetParticipantsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractRequestResetParticipants)
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
		it.Event = new(ParticipantManagerContractRequestResetParticipants)
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
func (it *ParticipantManagerContractRequestResetParticipantsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractRequestResetParticipantsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractRequestResetParticipants represents a RequestResetParticipants event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRequestResetParticipants struct {
	TaskId       uint64
	Participants []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestResetParticipants is a free log retrieval operation binding the contract event 0xcbbbc90ca424df467059e1af72e17dc379369c37abde7095e4ca4bf77af5d7f3.
//
// Solidity: event RequestResetParticipants(uint64 taskId, address[] indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterRequestResetParticipants(opts *bind.FilterOpts, participants [][]common.Address) (*ParticipantManagerContractRequestResetParticipantsIterator, error) {

	var participantsRule []interface{}
	for _, participantsItem := range participants {
		participantsRule = append(participantsRule, participantsItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "RequestResetParticipants", participantsRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractRequestResetParticipantsIterator{contract: _ParticipantManagerContract.contract, event: "RequestResetParticipants", logs: logs, sub: sub}, nil
}

// WatchRequestResetParticipants is a free log subscription operation binding the contract event 0xcbbbc90ca424df467059e1af72e17dc379369c37abde7095e4ca4bf77af5d7f3.
//
// Solidity: event RequestResetParticipants(uint64 taskId, address[] indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchRequestResetParticipants(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractRequestResetParticipants, participants [][]common.Address) (event.Subscription, error) {

	var participantsRule []interface{}
	for _, participantsItem := range participants {
		participantsRule = append(participantsRule, participantsItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "RequestResetParticipants", participantsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractRequestResetParticipants)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "RequestResetParticipants", log); err != nil {
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

// ParseRequestResetParticipants is a log parse operation binding the contract event 0xcbbbc90ca424df467059e1af72e17dc379369c37abde7095e4ca4bf77af5d7f3.
//
// Solidity: event RequestResetParticipants(uint64 taskId, address[] indexed participants)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseRequestResetParticipants(log types.Log) (*ParticipantManagerContractRequestResetParticipants, error) {
	event := new(ParticipantManagerContractRequestResetParticipants)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "RequestResetParticipants", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRoleAdminChangedIterator struct {
	Event *ParticipantManagerContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractRoleAdminChanged)
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
		it.Event = new(ParticipantManagerContractRoleAdminChanged)
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
func (it *ParticipantManagerContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractRoleAdminChanged represents a RoleAdminChanged event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ParticipantManagerContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractRoleAdminChangedIterator{contract: _ParticipantManagerContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractRoleAdminChanged)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseRoleAdminChanged(log types.Log) (*ParticipantManagerContractRoleAdminChanged, error) {
	event := new(ParticipantManagerContractRoleAdminChanged)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRoleGrantedIterator struct {
	Event *ParticipantManagerContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractRoleGranted)
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
		it.Event = new(ParticipantManagerContractRoleGranted)
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
func (it *ParticipantManagerContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractRoleGranted represents a RoleGranted event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ParticipantManagerContractRoleGrantedIterator, error) {

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

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractRoleGrantedIterator{contract: _ParticipantManagerContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractRoleGranted)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseRoleGranted(log types.Log) (*ParticipantManagerContractRoleGranted, error) {
	event := new(ParticipantManagerContractRoleGranted)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRoleRevokedIterator struct {
	Event *ParticipantManagerContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ParticipantManagerContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractRoleRevoked)
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
		it.Event = new(ParticipantManagerContractRoleRevoked)
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
func (it *ParticipantManagerContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractRoleRevoked represents a RoleRevoked event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ParticipantManagerContractRoleRevokedIterator, error) {

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

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractRoleRevokedIterator{contract: _ParticipantManagerContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractRoleRevoked)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseRoleRevoked(log types.Log) (*ParticipantManagerContractRoleRevoked, error) {
	event := new(ParticipantManagerContractRoleRevoked)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
