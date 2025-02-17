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

// AccountRegistrationTaskParam is an auto generated low-level Go binding around an user-defined struct.
type AccountRegistrationTaskParam struct {
	UserAddr common.Address
	Account  uint32
	Chain    uint8
	Index    uint32
}

// AccountManagerContractMetaData contains all meta data concerning the AccountManagerContract contract.
var AccountManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ENTRYPOINT_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SUBMITTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressRecord\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressRecord\",\"inputs\":[{\"name\":\"_account\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_chain\",\"type\":\"uint8\",\"internalType\":\"enumAddressCategory\"},{\"name\":\"_index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNewAddress\",\"inputs\":[{\"name\":\"_userAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_chain\",\"type\":\"uint8\",\"internalType\":\"enumAddressCategory\"},{\"name\":\"_index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_address\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRegisterTask\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple[]\",\"internalType\":\"structAccountRegistrationTaskParam[]\",\"components\":[{\"name\":\"userAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chain\",\"type\":\"uint8\",\"internalType\":\"enumAddressCategory\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"taskIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userAccounts\",\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"userAccount\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userMapping\",\"inputs\":[{\"name\":\"depositAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumAddressCategory\"}],\"outputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddressRegistered\",\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"chain\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumAddressCategory\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestRegisterAddress\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"taskParam\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAccountRegistrationTaskParam\",\"components\":[{\"name\":\"userAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chain\",\"type\":\"uint8\",\"internalType\":\"enumAddressCategory\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUserAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedAccount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RegisteredAccount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// AccountManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountManagerContractMetaData.ABI instead.
var AccountManagerContractABI = AccountManagerContractMetaData.ABI

// AccountManagerContract is an auto generated Go binding around an Ethereum contract.
type AccountManagerContract struct {
	AccountManagerContractCaller     // Read-only binding to the contract
	AccountManagerContractTransactor // Write-only binding to the contract
	AccountManagerContractFilterer   // Log filterer for contract events
}

// AccountManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountManagerContractSession struct {
	Contract     *AccountManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AccountManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountManagerContractCallerSession struct {
	Contract *AccountManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AccountManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountManagerContractTransactorSession struct {
	Contract     *AccountManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AccountManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountManagerContractRaw struct {
	Contract *AccountManagerContract // Generic contract binding to access the raw methods on
}

// AccountManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountManagerContractCallerRaw struct {
	Contract *AccountManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountManagerContractTransactorRaw struct {
	Contract *AccountManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManagerContract creates a new instance of AccountManagerContract, bound to a specific deployed contract.
func NewAccountManagerContract(address common.Address, backend bind.ContractBackend) (*AccountManagerContract, error) {
	contract, err := bindAccountManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContract{AccountManagerContractCaller: AccountManagerContractCaller{contract: contract}, AccountManagerContractTransactor: AccountManagerContractTransactor{contract: contract}, AccountManagerContractFilterer: AccountManagerContractFilterer{contract: contract}}, nil
}

// NewAccountManagerContractCaller creates a new read-only instance of AccountManagerContract, bound to a specific deployed contract.
func NewAccountManagerContractCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerContractCaller, error) {
	contract, err := bindAccountManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractCaller{contract: contract}, nil
}

// NewAccountManagerContractTransactor creates a new write-only instance of AccountManagerContract, bound to a specific deployed contract.
func NewAccountManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerContractTransactor, error) {
	contract, err := bindAccountManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractTransactor{contract: contract}, nil
}

// NewAccountManagerContractFilterer creates a new log filterer instance of AccountManagerContract, bound to a specific deployed contract.
func NewAccountManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerContractFilterer, error) {
	contract, err := bindAccountManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractFilterer{contract: contract}, nil
}

// bindAccountManagerContract binds a generic wrapper to an already deployed contract.
func bindAccountManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManagerContract *AccountManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManagerContract.Contract.AccountManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManagerContract *AccountManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.AccountManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManagerContract *AccountManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.AccountManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManagerContract *AccountManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManagerContract *AccountManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManagerContract *AccountManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccountManagerContract.Contract.DEFAULTADMINROLE(&_AccountManagerContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccountManagerContract.Contract.DEFAULTADMINROLE(&_AccountManagerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractCaller) ENTRYPOINTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "ENTRYPOINT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _AccountManagerContract.Contract.ENTRYPOINTROLE(&_AccountManagerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractCallerSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _AccountManagerContract.Contract.ENTRYPOINTROLE(&_AccountManagerContract.CallOpts)
}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractCaller) SUBMITTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "SUBMITTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractSession) SUBMITTERROLE() ([32]byte, error) {
	return _AccountManagerContract.Contract.SUBMITTERROLE(&_AccountManagerContract.CallOpts)
}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractCallerSession) SUBMITTERROLE() ([32]byte, error) {
	return _AccountManagerContract.Contract.SUBMITTERROLE(&_AccountManagerContract.CallOpts)
}

// AddressRecord is a free data retrieval call binding the contract method 0x2d0adbca.
//
// Solidity: function addressRecord(bytes32 ) view returns(string)
func (_AccountManagerContract *AccountManagerContractCaller) AddressRecord(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "addressRecord", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddressRecord is a free data retrieval call binding the contract method 0x2d0adbca.
//
// Solidity: function addressRecord(bytes32 ) view returns(string)
func (_AccountManagerContract *AccountManagerContractSession) AddressRecord(arg0 [32]byte) (string, error) {
	return _AccountManagerContract.Contract.AddressRecord(&_AccountManagerContract.CallOpts, arg0)
}

// AddressRecord is a free data retrieval call binding the contract method 0x2d0adbca.
//
// Solidity: function addressRecord(bytes32 ) view returns(string)
func (_AccountManagerContract *AccountManagerContractCallerSession) AddressRecord(arg0 [32]byte) (string, error) {
	return _AccountManagerContract.Contract.AddressRecord(&_AccountManagerContract.CallOpts, arg0)
}

// GetAddressRecord is a free data retrieval call binding the contract method 0xe1797f20.
//
// Solidity: function getAddressRecord(uint32 _account, uint8 _chain, uint32 _index) view returns(string)
func (_AccountManagerContract *AccountManagerContractCaller) GetAddressRecord(opts *bind.CallOpts, _account uint32, _chain uint8, _index uint32) (string, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "getAddressRecord", _account, _chain, _index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAddressRecord is a free data retrieval call binding the contract method 0xe1797f20.
//
// Solidity: function getAddressRecord(uint32 _account, uint8 _chain, uint32 _index) view returns(string)
func (_AccountManagerContract *AccountManagerContractSession) GetAddressRecord(_account uint32, _chain uint8, _index uint32) (string, error) {
	return _AccountManagerContract.Contract.GetAddressRecord(&_AccountManagerContract.CallOpts, _account, _chain, _index)
}

// GetAddressRecord is a free data retrieval call binding the contract method 0xe1797f20.
//
// Solidity: function getAddressRecord(uint32 _account, uint8 _chain, uint32 _index) view returns(string)
func (_AccountManagerContract *AccountManagerContractCallerSession) GetAddressRecord(_account uint32, _chain uint8, _index uint32) (string, error) {
	return _AccountManagerContract.Contract.GetAddressRecord(&_AccountManagerContract.CallOpts, _account, _chain, _index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccountManagerContract.Contract.GetRoleAdmin(&_AccountManagerContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccountManagerContract *AccountManagerContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccountManagerContract.Contract.GetRoleAdmin(&_AccountManagerContract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccountManagerContract *AccountManagerContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccountManagerContract *AccountManagerContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccountManagerContract.Contract.HasRole(&_AccountManagerContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccountManagerContract *AccountManagerContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccountManagerContract.Contract.HasRole(&_AccountManagerContract.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccountManagerContract *AccountManagerContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccountManagerContract *AccountManagerContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccountManagerContract.Contract.SupportsInterface(&_AccountManagerContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccountManagerContract *AccountManagerContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccountManagerContract.Contract.SupportsInterface(&_AccountManagerContract.CallOpts, interfaceId)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_AccountManagerContract *AccountManagerContractCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_AccountManagerContract *AccountManagerContractSession) TaskManager() (common.Address, error) {
	return _AccountManagerContract.Contract.TaskManager(&_AccountManagerContract.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_AccountManagerContract *AccountManagerContractCallerSession) TaskManager() (common.Address, error) {
	return _AccountManagerContract.Contract.TaskManager(&_AccountManagerContract.CallOpts)
}

// UserAccounts is a free data retrieval call binding the contract method 0x5251d91c.
//
// Solidity: function userAccounts(address userAddress) view returns(uint32 userAccount)
func (_AccountManagerContract *AccountManagerContractCaller) UserAccounts(opts *bind.CallOpts, userAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "userAccounts", userAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// UserAccounts is a free data retrieval call binding the contract method 0x5251d91c.
//
// Solidity: function userAccounts(address userAddress) view returns(uint32 userAccount)
func (_AccountManagerContract *AccountManagerContractSession) UserAccounts(userAddress common.Address) (uint32, error) {
	return _AccountManagerContract.Contract.UserAccounts(&_AccountManagerContract.CallOpts, userAddress)
}

// UserAccounts is a free data retrieval call binding the contract method 0x5251d91c.
//
// Solidity: function userAccounts(address userAddress) view returns(uint32 userAccount)
func (_AccountManagerContract *AccountManagerContractCallerSession) UserAccounts(userAddress common.Address) (uint32, error) {
	return _AccountManagerContract.Contract.UserAccounts(&_AccountManagerContract.CallOpts, userAddress)
}

// UserMapping is a free data retrieval call binding the contract method 0xa0873ccb.
//
// Solidity: function userMapping(string depositAddress, uint8 ) view returns(address account)
func (_AccountManagerContract *AccountManagerContractCaller) UserMapping(opts *bind.CallOpts, depositAddress string, arg1 uint8) (common.Address, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "userMapping", depositAddress, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserMapping is a free data retrieval call binding the contract method 0xa0873ccb.
//
// Solidity: function userMapping(string depositAddress, uint8 ) view returns(address account)
func (_AccountManagerContract *AccountManagerContractSession) UserMapping(depositAddress string, arg1 uint8) (common.Address, error) {
	return _AccountManagerContract.Contract.UserMapping(&_AccountManagerContract.CallOpts, depositAddress, arg1)
}

// UserMapping is a free data retrieval call binding the contract method 0xa0873ccb.
//
// Solidity: function userMapping(string depositAddress, uint8 ) view returns(address account)
func (_AccountManagerContract *AccountManagerContractCallerSession) UserMapping(depositAddress string, arg1 uint8) (common.Address, error) {
	return _AccountManagerContract.Contract.UserMapping(&_AccountManagerContract.CallOpts, depositAddress, arg1)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccountManagerContract *AccountManagerContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccountManagerContract *AccountManagerContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.GrantRole(&_AccountManagerContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccountManagerContract *AccountManagerContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.GrantRole(&_AccountManagerContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter) returns()
func (_AccountManagerContract *AccountManagerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _entryPoint common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "initialize", _owner, _entryPoint, _submitter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter) returns()
func (_AccountManagerContract *AccountManagerContractSession) Initialize(_owner common.Address, _entryPoint common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.Initialize(&_AccountManagerContract.TransactOpts, _owner, _entryPoint, _submitter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter) returns()
func (_AccountManagerContract *AccountManagerContractTransactorSession) Initialize(_owner common.Address, _entryPoint common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.Initialize(&_AccountManagerContract.TransactOpts, _owner, _entryPoint, _submitter)
}

// RegisterNewAddress is a paid mutator transaction binding the contract method 0x4c1d1565.
//
// Solidity: function registerNewAddress(address _userAddr, uint32 _account, uint8 _chain, uint32 _index, string _address) returns(bytes)
func (_AccountManagerContract *AccountManagerContractTransactor) RegisterNewAddress(opts *bind.TransactOpts, _userAddr common.Address, _account uint32, _chain uint8, _index uint32, _address string) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "registerNewAddress", _userAddr, _account, _chain, _index, _address)
}

// RegisterNewAddress is a paid mutator transaction binding the contract method 0x4c1d1565.
//
// Solidity: function registerNewAddress(address _userAddr, uint32 _account, uint8 _chain, uint32 _index, string _address) returns(bytes)
func (_AccountManagerContract *AccountManagerContractSession) RegisterNewAddress(_userAddr common.Address, _account uint32, _chain uint8, _index uint32, _address string) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RegisterNewAddress(&_AccountManagerContract.TransactOpts, _userAddr, _account, _chain, _index, _address)
}

// RegisterNewAddress is a paid mutator transaction binding the contract method 0x4c1d1565.
//
// Solidity: function registerNewAddress(address _userAddr, uint32 _account, uint8 _chain, uint32 _index, string _address) returns(bytes)
func (_AccountManagerContract *AccountManagerContractTransactorSession) RegisterNewAddress(_userAddr common.Address, _account uint32, _chain uint8, _index uint32, _address string) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RegisterNewAddress(&_AccountManagerContract.TransactOpts, _userAddr, _account, _chain, _index, _address)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccountManagerContract *AccountManagerContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccountManagerContract *AccountManagerContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RenounceRole(&_AccountManagerContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccountManagerContract *AccountManagerContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RenounceRole(&_AccountManagerContract.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccountManagerContract *AccountManagerContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccountManagerContract *AccountManagerContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RevokeRole(&_AccountManagerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccountManagerContract *AccountManagerContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RevokeRole(&_AccountManagerContract.TransactOpts, role, account)
}

// SubmitRegisterTask is a paid mutator transaction binding the contract method 0xb5d1bb3b.
//
// Solidity: function submitRegisterTask((address,uint32,uint8,uint32)[] _params) returns(uint64[] taskIds)
func (_AccountManagerContract *AccountManagerContractTransactor) SubmitRegisterTask(opts *bind.TransactOpts, _params []AccountRegistrationTaskParam) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "submitRegisterTask", _params)
}

// SubmitRegisterTask is a paid mutator transaction binding the contract method 0xb5d1bb3b.
//
// Solidity: function submitRegisterTask((address,uint32,uint8,uint32)[] _params) returns(uint64[] taskIds)
func (_AccountManagerContract *AccountManagerContractSession) SubmitRegisterTask(_params []AccountRegistrationTaskParam) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.SubmitRegisterTask(&_AccountManagerContract.TransactOpts, _params)
}

// SubmitRegisterTask is a paid mutator transaction binding the contract method 0xb5d1bb3b.
//
// Solidity: function submitRegisterTask((address,uint32,uint8,uint32)[] _params) returns(uint64[] taskIds)
func (_AccountManagerContract *AccountManagerContractTransactorSession) SubmitRegisterTask(_params []AccountRegistrationTaskParam) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.SubmitRegisterTask(&_AccountManagerContract.TransactOpts, _params)
}

// AccountManagerContractAddressRegisteredIterator is returned from FilterAddressRegistered and is used to iterate over the raw logs and unpacked data for AddressRegistered events raised by the AccountManagerContract contract.
type AccountManagerContractAddressRegisteredIterator struct {
	Event *AccountManagerContractAddressRegistered // Event containing the contract specifics and raw log

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
func (it *AccountManagerContractAddressRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerContractAddressRegistered)
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
		it.Event = new(AccountManagerContractAddressRegistered)
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
func (it *AccountManagerContractAddressRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerContractAddressRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerContractAddressRegistered represents a AddressRegistered event raised by the AccountManagerContract contract.
type AccountManagerContractAddressRegistered struct {
	UserAddr   common.Address
	Account    *big.Int
	Chain      uint8
	Index      *big.Int
	NewAddress string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressRegistered is a free log retrieval operation binding the contract event 0x0ab661710c67363885e0e51920050375aff9dcd587adf3e2e468e060ee8f0e1e.
//
// Solidity: event AddressRegistered(address userAddr, uint256 indexed account, uint8 indexed chain, uint256 indexed index, string newAddress)
func (_AccountManagerContract *AccountManagerContractFilterer) FilterAddressRegistered(opts *bind.FilterOpts, account []*big.Int, chain []uint8, index []*big.Int) (*AccountManagerContractAddressRegisteredIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AccountManagerContract.contract.FilterLogs(opts, "AddressRegistered", accountRule, chainRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractAddressRegisteredIterator{contract: _AccountManagerContract.contract, event: "AddressRegistered", logs: logs, sub: sub}, nil
}

// WatchAddressRegistered is a free log subscription operation binding the contract event 0x0ab661710c67363885e0e51920050375aff9dcd587adf3e2e468e060ee8f0e1e.
//
// Solidity: event AddressRegistered(address userAddr, uint256 indexed account, uint8 indexed chain, uint256 indexed index, string newAddress)
func (_AccountManagerContract *AccountManagerContractFilterer) WatchAddressRegistered(opts *bind.WatchOpts, sink chan<- *AccountManagerContractAddressRegistered, account []*big.Int, chain []uint8, index []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AccountManagerContract.contract.WatchLogs(opts, "AddressRegistered", accountRule, chainRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerContractAddressRegistered)
				if err := _AccountManagerContract.contract.UnpackLog(event, "AddressRegistered", log); err != nil {
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

// ParseAddressRegistered is a log parse operation binding the contract event 0x0ab661710c67363885e0e51920050375aff9dcd587adf3e2e468e060ee8f0e1e.
//
// Solidity: event AddressRegistered(address userAddr, uint256 indexed account, uint8 indexed chain, uint256 indexed index, string newAddress)
func (_AccountManagerContract *AccountManagerContractFilterer) ParseAddressRegistered(log types.Log) (*AccountManagerContractAddressRegistered, error) {
	event := new(AccountManagerContractAddressRegistered)
	if err := _AccountManagerContract.contract.UnpackLog(event, "AddressRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AccountManagerContract contract.
type AccountManagerContractInitializedIterator struct {
	Event *AccountManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *AccountManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerContractInitialized)
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
		it.Event = new(AccountManagerContractInitialized)
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
func (it *AccountManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerContractInitialized represents a Initialized event raised by the AccountManagerContract contract.
type AccountManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccountManagerContract *AccountManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountManagerContractInitializedIterator, error) {

	logs, sub, err := _AccountManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractInitializedIterator{contract: _AccountManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccountManagerContract *AccountManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _AccountManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerContractInitialized)
				if err := _AccountManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AccountManagerContract *AccountManagerContractFilterer) ParseInitialized(log types.Log) (*AccountManagerContractInitialized, error) {
	event := new(AccountManagerContractInitialized)
	if err := _AccountManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerContractRequestRegisterAddressIterator is returned from FilterRequestRegisterAddress and is used to iterate over the raw logs and unpacked data for RequestRegisterAddress events raised by the AccountManagerContract contract.
type AccountManagerContractRequestRegisterAddressIterator struct {
	Event *AccountManagerContractRequestRegisterAddress // Event containing the contract specifics and raw log

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
func (it *AccountManagerContractRequestRegisterAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerContractRequestRegisterAddress)
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
		it.Event = new(AccountManagerContractRequestRegisterAddress)
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
func (it *AccountManagerContractRequestRegisterAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerContractRequestRegisterAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerContractRequestRegisterAddress represents a RequestRegisterAddress event raised by the AccountManagerContract contract.
type AccountManagerContractRequestRegisterAddress struct {
	TaskId    uint64
	TaskParam AccountRegistrationTaskParam
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestRegisterAddress is a free log retrieval operation binding the contract event 0xc995ff8d6b3ea3422817922078e4f90f4ea4a5d279964f8581d2d6f6e9da6f5e.
//
// Solidity: event RequestRegisterAddress(uint64 taskId, (address,uint32,uint8,uint32) taskParam)
func (_AccountManagerContract *AccountManagerContractFilterer) FilterRequestRegisterAddress(opts *bind.FilterOpts) (*AccountManagerContractRequestRegisterAddressIterator, error) {

	logs, sub, err := _AccountManagerContract.contract.FilterLogs(opts, "RequestRegisterAddress")
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractRequestRegisterAddressIterator{contract: _AccountManagerContract.contract, event: "RequestRegisterAddress", logs: logs, sub: sub}, nil
}

// WatchRequestRegisterAddress is a free log subscription operation binding the contract event 0xc995ff8d6b3ea3422817922078e4f90f4ea4a5d279964f8581d2d6f6e9da6f5e.
//
// Solidity: event RequestRegisterAddress(uint64 taskId, (address,uint32,uint8,uint32) taskParam)
func (_AccountManagerContract *AccountManagerContractFilterer) WatchRequestRegisterAddress(opts *bind.WatchOpts, sink chan<- *AccountManagerContractRequestRegisterAddress) (event.Subscription, error) {

	logs, sub, err := _AccountManagerContract.contract.WatchLogs(opts, "RequestRegisterAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerContractRequestRegisterAddress)
				if err := _AccountManagerContract.contract.UnpackLog(event, "RequestRegisterAddress", log); err != nil {
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

// ParseRequestRegisterAddress is a log parse operation binding the contract event 0xc995ff8d6b3ea3422817922078e4f90f4ea4a5d279964f8581d2d6f6e9da6f5e.
//
// Solidity: event RequestRegisterAddress(uint64 taskId, (address,uint32,uint8,uint32) taskParam)
func (_AccountManagerContract *AccountManagerContractFilterer) ParseRequestRegisterAddress(log types.Log) (*AccountManagerContractRequestRegisterAddress, error) {
	event := new(AccountManagerContractRequestRegisterAddress)
	if err := _AccountManagerContract.contract.UnpackLog(event, "RequestRegisterAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccountManagerContract contract.
type AccountManagerContractRoleAdminChangedIterator struct {
	Event *AccountManagerContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccountManagerContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerContractRoleAdminChanged)
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
		it.Event = new(AccountManagerContractRoleAdminChanged)
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
func (it *AccountManagerContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerContractRoleAdminChanged represents a RoleAdminChanged event raised by the AccountManagerContract contract.
type AccountManagerContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccountManagerContract *AccountManagerContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccountManagerContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AccountManagerContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractRoleAdminChangedIterator{contract: _AccountManagerContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccountManagerContract *AccountManagerContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccountManagerContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AccountManagerContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerContractRoleAdminChanged)
				if err := _AccountManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AccountManagerContract *AccountManagerContractFilterer) ParseRoleAdminChanged(log types.Log) (*AccountManagerContractRoleAdminChanged, error) {
	event := new(AccountManagerContractRoleAdminChanged)
	if err := _AccountManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccountManagerContract contract.
type AccountManagerContractRoleGrantedIterator struct {
	Event *AccountManagerContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccountManagerContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerContractRoleGranted)
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
		it.Event = new(AccountManagerContractRoleGranted)
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
func (it *AccountManagerContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerContractRoleGranted represents a RoleGranted event raised by the AccountManagerContract contract.
type AccountManagerContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccountManagerContract *AccountManagerContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccountManagerContractRoleGrantedIterator, error) {

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

	logs, sub, err := _AccountManagerContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractRoleGrantedIterator{contract: _AccountManagerContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccountManagerContract *AccountManagerContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccountManagerContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccountManagerContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerContractRoleGranted)
				if err := _AccountManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AccountManagerContract *AccountManagerContractFilterer) ParseRoleGranted(log types.Log) (*AccountManagerContractRoleGranted, error) {
	event := new(AccountManagerContractRoleGranted)
	if err := _AccountManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccountManagerContract contract.
type AccountManagerContractRoleRevokedIterator struct {
	Event *AccountManagerContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccountManagerContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerContractRoleRevoked)
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
		it.Event = new(AccountManagerContractRoleRevoked)
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
func (it *AccountManagerContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerContractRoleRevoked represents a RoleRevoked event raised by the AccountManagerContract contract.
type AccountManagerContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccountManagerContract *AccountManagerContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccountManagerContractRoleRevokedIterator, error) {

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

	logs, sub, err := _AccountManagerContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractRoleRevokedIterator{contract: _AccountManagerContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccountManagerContract *AccountManagerContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccountManagerContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccountManagerContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerContractRoleRevoked)
				if err := _AccountManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AccountManagerContract *AccountManagerContractFilterer) ParseRoleRevoked(log types.Log) (*AccountManagerContractRoleRevoked, error) {
	event := new(AccountManagerContractRoleRevoked)
	if err := _AccountManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
