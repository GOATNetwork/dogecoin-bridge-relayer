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

// DepositInfo is an auto generated low-level Go binding around an user-defined struct.
type DepositInfo struct {
	UserAddress    common.Address
	ChainId        uint64
	Ticker         [32]byte
	DepositAddress string
	Amount         *big.Int
	TxHash         string
	BlockHeight    *big.Int
	LogIndex       *big.Int
}

// WithdrawalInfo is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalInfo struct {
	UserAddress common.Address
	ChainId     uint64
	Ticker      [32]byte
	ToAddress   string
	Amount      *big.Int
}

// DepositManagerContractMetaData contains all meta data concerning the DepositManagerContract contract.
var DepositManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_assetHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ENTRYPOINT_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SUBMITTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assetHandler\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAssetHandler\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposit\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDepositInfo\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structDepositInfo[]\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawal\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structWithdrawalInfo\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"toAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawals\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structWithdrawalInfo[]\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"toAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseState\",\"inputs\":[{\"name\":\"pauseType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordDeposit\",\"inputs\":[{\"name\":\"_param\",\"type\":\"tuple\",\"internalType\":\"structDepositInfo\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordWithdrawal\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_toAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_txHash\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauseState\",\"inputs\":[{\"name\":\"_condition\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_newState\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitDepositTask\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple[]\",\"internalType\":\"structDepositInfo[]\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"taskIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitWithdrawTask\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple[]\",\"internalType\":\"structWithdrawalInfo[]\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"toAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"taskIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"toAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositRecorded\",\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"depositAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NIP20TokenEvent_burnb\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NIP20TokenEvent_mintb\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewPauseState\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newState\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestDeposit\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"depositInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structDepositInfo\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestWithdrawal\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"withdrawalInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structWithdrawalInfo\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"toAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRecorded\",\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"toAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Paused\",\"inputs\":[]}]",
}

// DepositManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositManagerContractMetaData.ABI instead.
var DepositManagerContractABI = DepositManagerContractMetaData.ABI

// DepositManagerContract is an auto generated Go binding around an Ethereum contract.
type DepositManagerContract struct {
	DepositManagerContractCaller     // Read-only binding to the contract
	DepositManagerContractTransactor // Write-only binding to the contract
	DepositManagerContractFilterer   // Log filterer for contract events
}

// DepositManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositManagerContractSession struct {
	Contract     *DepositManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DepositManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositManagerContractCallerSession struct {
	Contract *DepositManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DepositManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositManagerContractTransactorSession struct {
	Contract     *DepositManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DepositManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositManagerContractRaw struct {
	Contract *DepositManagerContract // Generic contract binding to access the raw methods on
}

// DepositManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositManagerContractCallerRaw struct {
	Contract *DepositManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// DepositManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositManagerContractTransactorRaw struct {
	Contract *DepositManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositManagerContract creates a new instance of DepositManagerContract, bound to a specific deployed contract.
func NewDepositManagerContract(address common.Address, backend bind.ContractBackend) (*DepositManagerContract, error) {
	contract, err := bindDepositManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContract{DepositManagerContractCaller: DepositManagerContractCaller{contract: contract}, DepositManagerContractTransactor: DepositManagerContractTransactor{contract: contract}, DepositManagerContractFilterer: DepositManagerContractFilterer{contract: contract}}, nil
}

// NewDepositManagerContractCaller creates a new read-only instance of DepositManagerContract, bound to a specific deployed contract.
func NewDepositManagerContractCaller(address common.Address, caller bind.ContractCaller) (*DepositManagerContractCaller, error) {
	contract, err := bindDepositManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractCaller{contract: contract}, nil
}

// NewDepositManagerContractTransactor creates a new write-only instance of DepositManagerContract, bound to a specific deployed contract.
func NewDepositManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositManagerContractTransactor, error) {
	contract, err := bindDepositManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractTransactor{contract: contract}, nil
}

// NewDepositManagerContractFilterer creates a new log filterer instance of DepositManagerContract, bound to a specific deployed contract.
func NewDepositManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositManagerContractFilterer, error) {
	contract, err := bindDepositManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractFilterer{contract: contract}, nil
}

// bindDepositManagerContract binds a generic wrapper to an already deployed contract.
func bindDepositManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManagerContract *DepositManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositManagerContract.Contract.DepositManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManagerContract *DepositManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.DepositManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManagerContract *DepositManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.DepositManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManagerContract *DepositManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManagerContract *DepositManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManagerContract *DepositManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DepositManagerContract.Contract.DEFAULTADMINROLE(&_DepositManagerContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DepositManagerContract.Contract.DEFAULTADMINROLE(&_DepositManagerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractCaller) ENTRYPOINTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "ENTRYPOINT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _DepositManagerContract.Contract.ENTRYPOINTROLE(&_DepositManagerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractCallerSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _DepositManagerContract.Contract.ENTRYPOINTROLE(&_DepositManagerContract.CallOpts)
}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractCaller) SUBMITTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "SUBMITTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractSession) SUBMITTERROLE() ([32]byte, error) {
	return _DepositManagerContract.Contract.SUBMITTERROLE(&_DepositManagerContract.CallOpts)
}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractCallerSession) SUBMITTERROLE() ([32]byte, error) {
	return _DepositManagerContract.Contract.SUBMITTERROLE(&_DepositManagerContract.CallOpts)
}

// AssetHandler is a free data retrieval call binding the contract method 0xb18a6b93.
//
// Solidity: function assetHandler() view returns(address)
func (_DepositManagerContract *DepositManagerContractCaller) AssetHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "assetHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetHandler is a free data retrieval call binding the contract method 0xb18a6b93.
//
// Solidity: function assetHandler() view returns(address)
func (_DepositManagerContract *DepositManagerContractSession) AssetHandler() (common.Address, error) {
	return _DepositManagerContract.Contract.AssetHandler(&_DepositManagerContract.CallOpts)
}

// AssetHandler is a free data retrieval call binding the contract method 0xb18a6b93.
//
// Solidity: function assetHandler() view returns(address)
func (_DepositManagerContract *DepositManagerContractCallerSession) AssetHandler() (common.Address, error) {
	return _DepositManagerContract.Contract.AssetHandler(&_DepositManagerContract.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0xd6d68177.
//
// Solidity: function deposits(address userAddr, uint256 ) view returns(address userAddress, uint64 chainId, bytes32 ticker, string depositAddress, uint256 amount, string txHash, uint256 blockHeight, uint256 logIndex)
func (_DepositManagerContract *DepositManagerContractCaller) Deposits(opts *bind.CallOpts, userAddr common.Address, arg1 *big.Int) (struct {
	UserAddress    common.Address
	ChainId        uint64
	Ticker         [32]byte
	DepositAddress string
	Amount         *big.Int
	TxHash         string
	BlockHeight    *big.Int
	LogIndex       *big.Int
}, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "deposits", userAddr, arg1)

	outstruct := new(struct {
		UserAddress    common.Address
		ChainId        uint64
		Ticker         [32]byte
		DepositAddress string
		Amount         *big.Int
		TxHash         string
		BlockHeight    *big.Int
		LogIndex       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Ticker = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.DepositAddress = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TxHash = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.BlockHeight = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LogIndex = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xd6d68177.
//
// Solidity: function deposits(address userAddr, uint256 ) view returns(address userAddress, uint64 chainId, bytes32 ticker, string depositAddress, uint256 amount, string txHash, uint256 blockHeight, uint256 logIndex)
func (_DepositManagerContract *DepositManagerContractSession) Deposits(userAddr common.Address, arg1 *big.Int) (struct {
	UserAddress    common.Address
	ChainId        uint64
	Ticker         [32]byte
	DepositAddress string
	Amount         *big.Int
	TxHash         string
	BlockHeight    *big.Int
	LogIndex       *big.Int
}, error) {
	return _DepositManagerContract.Contract.Deposits(&_DepositManagerContract.CallOpts, userAddr, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0xd6d68177.
//
// Solidity: function deposits(address userAddr, uint256 ) view returns(address userAddress, uint64 chainId, bytes32 ticker, string depositAddress, uint256 amount, string txHash, uint256 blockHeight, uint256 logIndex)
func (_DepositManagerContract *DepositManagerContractCallerSession) Deposits(userAddr common.Address, arg1 *big.Int) (struct {
	UserAddress    common.Address
	ChainId        uint64
	Ticker         [32]byte
	DepositAddress string
	Amount         *big.Int
	TxHash         string
	BlockHeight    *big.Int
	LogIndex       *big.Int
}, error) {
	return _DepositManagerContract.Contract.Deposits(&_DepositManagerContract.CallOpts, userAddr, arg1)
}

// GetDeposit is a free data retrieval call binding the contract method 0x2726b506.
//
// Solidity: function getDeposit(address _userAddress, uint256 _index) view returns((address,uint64,bytes32,string,uint256,string,uint256,uint256))
func (_DepositManagerContract *DepositManagerContractCaller) GetDeposit(opts *bind.CallOpts, _userAddress common.Address, _index *big.Int) (DepositInfo, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "getDeposit", _userAddress, _index)

	if err != nil {
		return *new(DepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DepositInfo)).(*DepositInfo)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0x2726b506.
//
// Solidity: function getDeposit(address _userAddress, uint256 _index) view returns((address,uint64,bytes32,string,uint256,string,uint256,uint256))
func (_DepositManagerContract *DepositManagerContractSession) GetDeposit(_userAddress common.Address, _index *big.Int) (DepositInfo, error) {
	return _DepositManagerContract.Contract.GetDeposit(&_DepositManagerContract.CallOpts, _userAddress, _index)
}

// GetDeposit is a free data retrieval call binding the contract method 0x2726b506.
//
// Solidity: function getDeposit(address _userAddress, uint256 _index) view returns((address,uint64,bytes32,string,uint256,string,uint256,uint256))
func (_DepositManagerContract *DepositManagerContractCallerSession) GetDeposit(_userAddress common.Address, _index *big.Int) (DepositInfo, error) {
	return _DepositManagerContract.Contract.GetDeposit(&_DepositManagerContract.CallOpts, _userAddress, _index)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address _userAddress) view returns((address,uint64,bytes32,string,uint256,string,uint256,uint256)[])
func (_DepositManagerContract *DepositManagerContractCaller) GetDeposits(opts *bind.CallOpts, _userAddress common.Address) ([]DepositInfo, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "getDeposits", _userAddress)

	if err != nil {
		return *new([]DepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DepositInfo)).(*[]DepositInfo)

	return out0, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address _userAddress) view returns((address,uint64,bytes32,string,uint256,string,uint256,uint256)[])
func (_DepositManagerContract *DepositManagerContractSession) GetDeposits(_userAddress common.Address) ([]DepositInfo, error) {
	return _DepositManagerContract.Contract.GetDeposits(&_DepositManagerContract.CallOpts, _userAddress)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address _userAddress) view returns((address,uint64,bytes32,string,uint256,string,uint256,uint256)[])
func (_DepositManagerContract *DepositManagerContractCallerSession) GetDeposits(_userAddress common.Address) ([]DepositInfo, error) {
	return _DepositManagerContract.Contract.GetDeposits(&_DepositManagerContract.CallOpts, _userAddress)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DepositManagerContract.Contract.GetRoleAdmin(&_DepositManagerContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DepositManagerContract *DepositManagerContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DepositManagerContract.Contract.GetRoleAdmin(&_DepositManagerContract.CallOpts, role)
}

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(address _userAddress, uint256 _index) view returns((address,uint64,bytes32,string,uint256))
func (_DepositManagerContract *DepositManagerContractCaller) GetWithdrawal(opts *bind.CallOpts, _userAddress common.Address, _index *big.Int) (WithdrawalInfo, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "getWithdrawal", _userAddress, _index)

	if err != nil {
		return *new(WithdrawalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(WithdrawalInfo)).(*WithdrawalInfo)

	return out0, err

}

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(address _userAddress, uint256 _index) view returns((address,uint64,bytes32,string,uint256))
func (_DepositManagerContract *DepositManagerContractSession) GetWithdrawal(_userAddress common.Address, _index *big.Int) (WithdrawalInfo, error) {
	return _DepositManagerContract.Contract.GetWithdrawal(&_DepositManagerContract.CallOpts, _userAddress, _index)
}

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(address _userAddress, uint256 _index) view returns((address,uint64,bytes32,string,uint256))
func (_DepositManagerContract *DepositManagerContractCallerSession) GetWithdrawal(_userAddress common.Address, _index *big.Int) (WithdrawalInfo, error) {
	return _DepositManagerContract.Contract.GetWithdrawal(&_DepositManagerContract.CallOpts, _userAddress, _index)
}

// GetWithdrawals is a free data retrieval call binding the contract method 0x3a2b643a.
//
// Solidity: function getWithdrawals(address _userAddress) view returns((address,uint64,bytes32,string,uint256)[])
func (_DepositManagerContract *DepositManagerContractCaller) GetWithdrawals(opts *bind.CallOpts, _userAddress common.Address) ([]WithdrawalInfo, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "getWithdrawals", _userAddress)

	if err != nil {
		return *new([]WithdrawalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]WithdrawalInfo)).(*[]WithdrawalInfo)

	return out0, err

}

// GetWithdrawals is a free data retrieval call binding the contract method 0x3a2b643a.
//
// Solidity: function getWithdrawals(address _userAddress) view returns((address,uint64,bytes32,string,uint256)[])
func (_DepositManagerContract *DepositManagerContractSession) GetWithdrawals(_userAddress common.Address) ([]WithdrawalInfo, error) {
	return _DepositManagerContract.Contract.GetWithdrawals(&_DepositManagerContract.CallOpts, _userAddress)
}

// GetWithdrawals is a free data retrieval call binding the contract method 0x3a2b643a.
//
// Solidity: function getWithdrawals(address _userAddress) view returns((address,uint64,bytes32,string,uint256)[])
func (_DepositManagerContract *DepositManagerContractCallerSession) GetWithdrawals(_userAddress common.Address) ([]WithdrawalInfo, error) {
	return _DepositManagerContract.Contract.GetWithdrawals(&_DepositManagerContract.CallOpts, _userAddress)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DepositManagerContract *DepositManagerContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DepositManagerContract *DepositManagerContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DepositManagerContract.Contract.HasRole(&_DepositManagerContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DepositManagerContract *DepositManagerContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DepositManagerContract.Contract.HasRole(&_DepositManagerContract.CallOpts, role, account)
}

// PauseState is a free data retrieval call binding the contract method 0xc94db23c.
//
// Solidity: function pauseState(bytes32 pauseType) view returns(bool isPaused)
func (_DepositManagerContract *DepositManagerContractCaller) PauseState(opts *bind.CallOpts, pauseType [32]byte) (bool, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "pauseState", pauseType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PauseState is a free data retrieval call binding the contract method 0xc94db23c.
//
// Solidity: function pauseState(bytes32 pauseType) view returns(bool isPaused)
func (_DepositManagerContract *DepositManagerContractSession) PauseState(pauseType [32]byte) (bool, error) {
	return _DepositManagerContract.Contract.PauseState(&_DepositManagerContract.CallOpts, pauseType)
}

// PauseState is a free data retrieval call binding the contract method 0xc94db23c.
//
// Solidity: function pauseState(bytes32 pauseType) view returns(bool isPaused)
func (_DepositManagerContract *DepositManagerContractCallerSession) PauseState(pauseType [32]byte) (bool, error) {
	return _DepositManagerContract.Contract.PauseState(&_DepositManagerContract.CallOpts, pauseType)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DepositManagerContract *DepositManagerContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DepositManagerContract *DepositManagerContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DepositManagerContract.Contract.SupportsInterface(&_DepositManagerContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DepositManagerContract *DepositManagerContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DepositManagerContract.Contract.SupportsInterface(&_DepositManagerContract.CallOpts, interfaceId)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_DepositManagerContract *DepositManagerContractCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_DepositManagerContract *DepositManagerContractSession) TaskManager() (common.Address, error) {
	return _DepositManagerContract.Contract.TaskManager(&_DepositManagerContract.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_DepositManagerContract *DepositManagerContractCallerSession) TaskManager() (common.Address, error) {
	return _DepositManagerContract.Contract.TaskManager(&_DepositManagerContract.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0x422b1077.
//
// Solidity: function withdrawals(address userAddr, uint256 ) view returns(address userAddress, uint64 chainId, bytes32 ticker, string toAddress, uint256 amount)
func (_DepositManagerContract *DepositManagerContractCaller) Withdrawals(opts *bind.CallOpts, userAddr common.Address, arg1 *big.Int) (struct {
	UserAddress common.Address
	ChainId     uint64
	Ticker      [32]byte
	ToAddress   string
	Amount      *big.Int
}, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "withdrawals", userAddr, arg1)

	outstruct := new(struct {
		UserAddress common.Address
		ChainId     uint64
		Ticker      [32]byte
		ToAddress   string
		Amount      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Ticker = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.ToAddress = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Withdrawals is a free data retrieval call binding the contract method 0x422b1077.
//
// Solidity: function withdrawals(address userAddr, uint256 ) view returns(address userAddress, uint64 chainId, bytes32 ticker, string toAddress, uint256 amount)
func (_DepositManagerContract *DepositManagerContractSession) Withdrawals(userAddr common.Address, arg1 *big.Int) (struct {
	UserAddress common.Address
	ChainId     uint64
	Ticker      [32]byte
	ToAddress   string
	Amount      *big.Int
}, error) {
	return _DepositManagerContract.Contract.Withdrawals(&_DepositManagerContract.CallOpts, userAddr, arg1)
}

// Withdrawals is a free data retrieval call binding the contract method 0x422b1077.
//
// Solidity: function withdrawals(address userAddr, uint256 ) view returns(address userAddress, uint64 chainId, bytes32 ticker, string toAddress, uint256 amount)
func (_DepositManagerContract *DepositManagerContractCallerSession) Withdrawals(userAddr common.Address, arg1 *big.Int) (struct {
	UserAddress common.Address
	ChainId     uint64
	Ticker      [32]byte
	ToAddress   string
	Amount      *big.Int
}, error) {
	return _DepositManagerContract.Contract.Withdrawals(&_DepositManagerContract.CallOpts, userAddr, arg1)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DepositManagerContract *DepositManagerContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DepositManagerContract *DepositManagerContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.GrantRole(&_DepositManagerContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DepositManagerContract *DepositManagerContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.GrantRole(&_DepositManagerContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter) returns()
func (_DepositManagerContract *DepositManagerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _entryPoint common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "initialize", _owner, _entryPoint, _submitter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter) returns()
func (_DepositManagerContract *DepositManagerContractSession) Initialize(_owner common.Address, _entryPoint common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.Initialize(&_DepositManagerContract.TransactOpts, _owner, _entryPoint, _submitter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter) returns()
func (_DepositManagerContract *DepositManagerContractTransactorSession) Initialize(_owner common.Address, _entryPoint common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.Initialize(&_DepositManagerContract.TransactOpts, _owner, _entryPoint, _submitter)
}

// RecordDeposit is a paid mutator transaction binding the contract method 0xdabfb4e1.
//
// Solidity: function recordDeposit((address,uint64,bytes32,string,uint256,string,uint256,uint256) _param) returns(bytes)
func (_DepositManagerContract *DepositManagerContractTransactor) RecordDeposit(opts *bind.TransactOpts, _param DepositInfo) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "recordDeposit", _param)
}

// RecordDeposit is a paid mutator transaction binding the contract method 0xdabfb4e1.
//
// Solidity: function recordDeposit((address,uint64,bytes32,string,uint256,string,uint256,uint256) _param) returns(bytes)
func (_DepositManagerContract *DepositManagerContractSession) RecordDeposit(_param DepositInfo) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RecordDeposit(&_DepositManagerContract.TransactOpts, _param)
}

// RecordDeposit is a paid mutator transaction binding the contract method 0xdabfb4e1.
//
// Solidity: function recordDeposit((address,uint64,bytes32,string,uint256,string,uint256,uint256) _param) returns(bytes)
func (_DepositManagerContract *DepositManagerContractTransactorSession) RecordDeposit(_param DepositInfo) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RecordDeposit(&_DepositManagerContract.TransactOpts, _param)
}

// RecordWithdrawal is a paid mutator transaction binding the contract method 0x39441427.
//
// Solidity: function recordWithdrawal(address _userAddress, uint64 _chainId, bytes32 _ticker, string _toAddress, uint256 _amount, string _txHash) returns(bytes)
func (_DepositManagerContract *DepositManagerContractTransactor) RecordWithdrawal(opts *bind.TransactOpts, _userAddress common.Address, _chainId uint64, _ticker [32]byte, _toAddress string, _amount *big.Int, _txHash string) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "recordWithdrawal", _userAddress, _chainId, _ticker, _toAddress, _amount, _txHash)
}

// RecordWithdrawal is a paid mutator transaction binding the contract method 0x39441427.
//
// Solidity: function recordWithdrawal(address _userAddress, uint64 _chainId, bytes32 _ticker, string _toAddress, uint256 _amount, string _txHash) returns(bytes)
func (_DepositManagerContract *DepositManagerContractSession) RecordWithdrawal(_userAddress common.Address, _chainId uint64, _ticker [32]byte, _toAddress string, _amount *big.Int, _txHash string) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RecordWithdrawal(&_DepositManagerContract.TransactOpts, _userAddress, _chainId, _ticker, _toAddress, _amount, _txHash)
}

// RecordWithdrawal is a paid mutator transaction binding the contract method 0x39441427.
//
// Solidity: function recordWithdrawal(address _userAddress, uint64 _chainId, bytes32 _ticker, string _toAddress, uint256 _amount, string _txHash) returns(bytes)
func (_DepositManagerContract *DepositManagerContractTransactorSession) RecordWithdrawal(_userAddress common.Address, _chainId uint64, _ticker [32]byte, _toAddress string, _amount *big.Int, _txHash string) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RecordWithdrawal(&_DepositManagerContract.TransactOpts, _userAddress, _chainId, _ticker, _toAddress, _amount, _txHash)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DepositManagerContract *DepositManagerContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DepositManagerContract *DepositManagerContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RenounceRole(&_DepositManagerContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DepositManagerContract *DepositManagerContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RenounceRole(&_DepositManagerContract.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DepositManagerContract *DepositManagerContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DepositManagerContract *DepositManagerContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RevokeRole(&_DepositManagerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DepositManagerContract *DepositManagerContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RevokeRole(&_DepositManagerContract.TransactOpts, role, account)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xa9b92ccf.
//
// Solidity: function setPauseState(bytes32 _condition, bool _newState) returns()
func (_DepositManagerContract *DepositManagerContractTransactor) SetPauseState(opts *bind.TransactOpts, _condition [32]byte, _newState bool) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "setPauseState", _condition, _newState)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xa9b92ccf.
//
// Solidity: function setPauseState(bytes32 _condition, bool _newState) returns()
func (_DepositManagerContract *DepositManagerContractSession) SetPauseState(_condition [32]byte, _newState bool) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.SetPauseState(&_DepositManagerContract.TransactOpts, _condition, _newState)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xa9b92ccf.
//
// Solidity: function setPauseState(bytes32 _condition, bool _newState) returns()
func (_DepositManagerContract *DepositManagerContractTransactorSession) SetPauseState(_condition [32]byte, _newState bool) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.SetPauseState(&_DepositManagerContract.TransactOpts, _condition, _newState)
}

// SubmitDepositTask is a paid mutator transaction binding the contract method 0x760cd65c.
//
// Solidity: function submitDepositTask((address,uint64,bytes32,string,uint256,string,uint256,uint256)[] _params) returns(uint64[] taskIds)
func (_DepositManagerContract *DepositManagerContractTransactor) SubmitDepositTask(opts *bind.TransactOpts, _params []DepositInfo) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "submitDepositTask", _params)
}

// SubmitDepositTask is a paid mutator transaction binding the contract method 0x760cd65c.
//
// Solidity: function submitDepositTask((address,uint64,bytes32,string,uint256,string,uint256,uint256)[] _params) returns(uint64[] taskIds)
func (_DepositManagerContract *DepositManagerContractSession) SubmitDepositTask(_params []DepositInfo) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.SubmitDepositTask(&_DepositManagerContract.TransactOpts, _params)
}

// SubmitDepositTask is a paid mutator transaction binding the contract method 0x760cd65c.
//
// Solidity: function submitDepositTask((address,uint64,bytes32,string,uint256,string,uint256,uint256)[] _params) returns(uint64[] taskIds)
func (_DepositManagerContract *DepositManagerContractTransactorSession) SubmitDepositTask(_params []DepositInfo) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.SubmitDepositTask(&_DepositManagerContract.TransactOpts, _params)
}

// SubmitWithdrawTask is a paid mutator transaction binding the contract method 0xbc6e90cd.
//
// Solidity: function submitWithdrawTask((address,uint64,bytes32,string,uint256)[] _params) returns(uint64[] taskIds)
func (_DepositManagerContract *DepositManagerContractTransactor) SubmitWithdrawTask(opts *bind.TransactOpts, _params []WithdrawalInfo) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "submitWithdrawTask", _params)
}

// SubmitWithdrawTask is a paid mutator transaction binding the contract method 0xbc6e90cd.
//
// Solidity: function submitWithdrawTask((address,uint64,bytes32,string,uint256)[] _params) returns(uint64[] taskIds)
func (_DepositManagerContract *DepositManagerContractSession) SubmitWithdrawTask(_params []WithdrawalInfo) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.SubmitWithdrawTask(&_DepositManagerContract.TransactOpts, _params)
}

// SubmitWithdrawTask is a paid mutator transaction binding the contract method 0xbc6e90cd.
//
// Solidity: function submitWithdrawTask((address,uint64,bytes32,string,uint256)[] _params) returns(uint64[] taskIds)
func (_DepositManagerContract *DepositManagerContractTransactorSession) SubmitWithdrawTask(_params []WithdrawalInfo) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.SubmitWithdrawTask(&_DepositManagerContract.TransactOpts, _params)
}

// DepositManagerContractDepositRecordedIterator is returned from FilterDepositRecorded and is used to iterate over the raw logs and unpacked data for DepositRecorded events raised by the DepositManagerContract contract.
type DepositManagerContractDepositRecordedIterator struct {
	Event *DepositManagerContractDepositRecorded // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractDepositRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractDepositRecorded)
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
		it.Event = new(DepositManagerContractDepositRecorded)
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
func (it *DepositManagerContractDepositRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractDepositRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractDepositRecorded represents a DepositRecorded event raised by the DepositManagerContract contract.
type DepositManagerContractDepositRecorded struct {
	UserAddress    common.Address
	Ticker         [32]byte
	ChainId        uint64
	DepositAddress string
	Amount         *big.Int
	TxHash         string
	BlockHeight    *big.Int
	LogIndex       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDepositRecorded is a free log retrieval operation binding the contract event 0xb707b36bff7df581533b6b1939ae22c92c6fb817647b5a7bc7b7a42914dd7927.
//
// Solidity: event DepositRecorded(address indexed userAddress, bytes32 indexed ticker, uint64 indexed chainId, string depositAddress, uint256 amount, string txHash, uint256 blockHeight, uint256 logIndex)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterDepositRecorded(opts *bind.FilterOpts, userAddress []common.Address, ticker [][32]byte, chainId []uint64) (*DepositManagerContractDepositRecordedIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "DepositRecorded", userAddressRule, tickerRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractDepositRecordedIterator{contract: _DepositManagerContract.contract, event: "DepositRecorded", logs: logs, sub: sub}, nil
}

// WatchDepositRecorded is a free log subscription operation binding the contract event 0xb707b36bff7df581533b6b1939ae22c92c6fb817647b5a7bc7b7a42914dd7927.
//
// Solidity: event DepositRecorded(address indexed userAddress, bytes32 indexed ticker, uint64 indexed chainId, string depositAddress, uint256 amount, string txHash, uint256 blockHeight, uint256 logIndex)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchDepositRecorded(opts *bind.WatchOpts, sink chan<- *DepositManagerContractDepositRecorded, userAddress []common.Address, ticker [][32]byte, chainId []uint64) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "DepositRecorded", userAddressRule, tickerRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractDepositRecorded)
				if err := _DepositManagerContract.contract.UnpackLog(event, "DepositRecorded", log); err != nil {
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

// ParseDepositRecorded is a log parse operation binding the contract event 0xb707b36bff7df581533b6b1939ae22c92c6fb817647b5a7bc7b7a42914dd7927.
//
// Solidity: event DepositRecorded(address indexed userAddress, bytes32 indexed ticker, uint64 indexed chainId, string depositAddress, uint256 amount, string txHash, uint256 blockHeight, uint256 logIndex)
func (_DepositManagerContract *DepositManagerContractFilterer) ParseDepositRecorded(log types.Log) (*DepositManagerContractDepositRecorded, error) {
	event := new(DepositManagerContractDepositRecorded)
	if err := _DepositManagerContract.contract.UnpackLog(event, "DepositRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DepositManagerContract contract.
type DepositManagerContractInitializedIterator struct {
	Event *DepositManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractInitialized)
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
		it.Event = new(DepositManagerContractInitialized)
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
func (it *DepositManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractInitialized represents a Initialized event raised by the DepositManagerContract contract.
type DepositManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*DepositManagerContractInitializedIterator, error) {

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractInitializedIterator{contract: _DepositManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DepositManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractInitialized)
				if err := _DepositManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DepositManagerContract *DepositManagerContractFilterer) ParseInitialized(log types.Log) (*DepositManagerContractInitialized, error) {
	event := new(DepositManagerContractInitialized)
	if err := _DepositManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractNIP20TokenEventBurnbIterator is returned from FilterNIP20TokenEventBurnb and is used to iterate over the raw logs and unpacked data for NIP20TokenEventBurnb events raised by the DepositManagerContract contract.
type DepositManagerContractNIP20TokenEventBurnbIterator struct {
	Event *DepositManagerContractNIP20TokenEventBurnb // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractNIP20TokenEventBurnbIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractNIP20TokenEventBurnb)
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
		it.Event = new(DepositManagerContractNIP20TokenEventBurnb)
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
func (it *DepositManagerContractNIP20TokenEventBurnbIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractNIP20TokenEventBurnbIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractNIP20TokenEventBurnb represents a NIP20TokenEventBurnb event raised by the DepositManagerContract contract.
type DepositManagerContractNIP20TokenEventBurnb struct {
	From   common.Address
	Ticker [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNIP20TokenEventBurnb is a free log retrieval operation binding the contract event 0xebe23dd93b970477278ceb9abd3082df92d977d6131fb0ef75f18c3d353b565a.
//
// Solidity: event NIP20TokenEvent_burnb(address indexed from, bytes32 indexed ticker, uint256 amount)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterNIP20TokenEventBurnb(opts *bind.FilterOpts, from []common.Address, ticker [][32]byte) (*DepositManagerContractNIP20TokenEventBurnbIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "NIP20TokenEvent_burnb", fromRule, tickerRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractNIP20TokenEventBurnbIterator{contract: _DepositManagerContract.contract, event: "NIP20TokenEvent_burnb", logs: logs, sub: sub}, nil
}

// WatchNIP20TokenEventBurnb is a free log subscription operation binding the contract event 0xebe23dd93b970477278ceb9abd3082df92d977d6131fb0ef75f18c3d353b565a.
//
// Solidity: event NIP20TokenEvent_burnb(address indexed from, bytes32 indexed ticker, uint256 amount)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchNIP20TokenEventBurnb(opts *bind.WatchOpts, sink chan<- *DepositManagerContractNIP20TokenEventBurnb, from []common.Address, ticker [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "NIP20TokenEvent_burnb", fromRule, tickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractNIP20TokenEventBurnb)
				if err := _DepositManagerContract.contract.UnpackLog(event, "NIP20TokenEvent_burnb", log); err != nil {
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
// Solidity: event NIP20TokenEvent_burnb(address indexed from, bytes32 indexed ticker, uint256 amount)
func (_DepositManagerContract *DepositManagerContractFilterer) ParseNIP20TokenEventBurnb(log types.Log) (*DepositManagerContractNIP20TokenEventBurnb, error) {
	event := new(DepositManagerContractNIP20TokenEventBurnb)
	if err := _DepositManagerContract.contract.UnpackLog(event, "NIP20TokenEvent_burnb", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractNIP20TokenEventMintbIterator is returned from FilterNIP20TokenEventMintb and is used to iterate over the raw logs and unpacked data for NIP20TokenEventMintb events raised by the DepositManagerContract contract.
type DepositManagerContractNIP20TokenEventMintbIterator struct {
	Event *DepositManagerContractNIP20TokenEventMintb // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractNIP20TokenEventMintbIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractNIP20TokenEventMintb)
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
		it.Event = new(DepositManagerContractNIP20TokenEventMintb)
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
func (it *DepositManagerContractNIP20TokenEventMintbIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractNIP20TokenEventMintbIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractNIP20TokenEventMintb represents a NIP20TokenEventMintb event raised by the DepositManagerContract contract.
type DepositManagerContractNIP20TokenEventMintb struct {
	Recipient common.Address
	Ticker    [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNIP20TokenEventMintb is a free log retrieval operation binding the contract event 0x685c530c280ee1f7a4e96d082303ee9ebf21cec512259c6a943eda3854e05102.
//
// Solidity: event NIP20TokenEvent_mintb(address indexed recipient, bytes32 indexed ticker, uint256 amount)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterNIP20TokenEventMintb(opts *bind.FilterOpts, recipient []common.Address, ticker [][32]byte) (*DepositManagerContractNIP20TokenEventMintbIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "NIP20TokenEvent_mintb", recipientRule, tickerRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractNIP20TokenEventMintbIterator{contract: _DepositManagerContract.contract, event: "NIP20TokenEvent_mintb", logs: logs, sub: sub}, nil
}

// WatchNIP20TokenEventMintb is a free log subscription operation binding the contract event 0x685c530c280ee1f7a4e96d082303ee9ebf21cec512259c6a943eda3854e05102.
//
// Solidity: event NIP20TokenEvent_mintb(address indexed recipient, bytes32 indexed ticker, uint256 amount)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchNIP20TokenEventMintb(opts *bind.WatchOpts, sink chan<- *DepositManagerContractNIP20TokenEventMintb, recipient []common.Address, ticker [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "NIP20TokenEvent_mintb", recipientRule, tickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractNIP20TokenEventMintb)
				if err := _DepositManagerContract.contract.UnpackLog(event, "NIP20TokenEvent_mintb", log); err != nil {
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
// Solidity: event NIP20TokenEvent_mintb(address indexed recipient, bytes32 indexed ticker, uint256 amount)
func (_DepositManagerContract *DepositManagerContractFilterer) ParseNIP20TokenEventMintb(log types.Log) (*DepositManagerContractNIP20TokenEventMintb, error) {
	event := new(DepositManagerContractNIP20TokenEventMintb)
	if err := _DepositManagerContract.contract.UnpackLog(event, "NIP20TokenEvent_mintb", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractNewPauseStateIterator is returned from FilterNewPauseState and is used to iterate over the raw logs and unpacked data for NewPauseState events raised by the DepositManagerContract contract.
type DepositManagerContractNewPauseStateIterator struct {
	Event *DepositManagerContractNewPauseState // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractNewPauseStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractNewPauseState)
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
		it.Event = new(DepositManagerContractNewPauseState)
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
func (it *DepositManagerContractNewPauseStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractNewPauseStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractNewPauseState represents a NewPauseState event raised by the DepositManagerContract contract.
type DepositManagerContractNewPauseState struct {
	Condition [32]byte
	NewState  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewPauseState is a free log retrieval operation binding the contract event 0xfa6e53b4fcee476e09253f70545d292a28f0bd8254ea05c37283b9b8ab040402.
//
// Solidity: event NewPauseState(bytes32 indexed condition, bool indexed newState)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterNewPauseState(opts *bind.FilterOpts, condition [][32]byte, newState []bool) (*DepositManagerContractNewPauseStateIterator, error) {

	var conditionRule []interface{}
	for _, conditionItem := range condition {
		conditionRule = append(conditionRule, conditionItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "NewPauseState", conditionRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractNewPauseStateIterator{contract: _DepositManagerContract.contract, event: "NewPauseState", logs: logs, sub: sub}, nil
}

// WatchNewPauseState is a free log subscription operation binding the contract event 0xfa6e53b4fcee476e09253f70545d292a28f0bd8254ea05c37283b9b8ab040402.
//
// Solidity: event NewPauseState(bytes32 indexed condition, bool indexed newState)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchNewPauseState(opts *bind.WatchOpts, sink chan<- *DepositManagerContractNewPauseState, condition [][32]byte, newState []bool) (event.Subscription, error) {

	var conditionRule []interface{}
	for _, conditionItem := range condition {
		conditionRule = append(conditionRule, conditionItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "NewPauseState", conditionRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractNewPauseState)
				if err := _DepositManagerContract.contract.UnpackLog(event, "NewPauseState", log); err != nil {
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

// ParseNewPauseState is a log parse operation binding the contract event 0xfa6e53b4fcee476e09253f70545d292a28f0bd8254ea05c37283b9b8ab040402.
//
// Solidity: event NewPauseState(bytes32 indexed condition, bool indexed newState)
func (_DepositManagerContract *DepositManagerContractFilterer) ParseNewPauseState(log types.Log) (*DepositManagerContractNewPauseState, error) {
	event := new(DepositManagerContractNewPauseState)
	if err := _DepositManagerContract.contract.UnpackLog(event, "NewPauseState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractRequestDepositIterator is returned from FilterRequestDeposit and is used to iterate over the raw logs and unpacked data for RequestDeposit events raised by the DepositManagerContract contract.
type DepositManagerContractRequestDepositIterator struct {
	Event *DepositManagerContractRequestDeposit // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractRequestDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractRequestDeposit)
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
		it.Event = new(DepositManagerContractRequestDeposit)
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
func (it *DepositManagerContractRequestDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractRequestDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractRequestDeposit represents a RequestDeposit event raised by the DepositManagerContract contract.
type DepositManagerContractRequestDeposit struct {
	TaskId      uint64
	DepositInfo DepositInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestDeposit is a free log retrieval operation binding the contract event 0xe84fe8b7649d8ddef364d5de9c5006de2baa5aed97143b907af093cf81a0dca4.
//
// Solidity: event RequestDeposit(uint64 taskId, (address,uint64,bytes32,string,uint256,string,uint256,uint256) depositInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterRequestDeposit(opts *bind.FilterOpts) (*DepositManagerContractRequestDepositIterator, error) {

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "RequestDeposit")
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractRequestDepositIterator{contract: _DepositManagerContract.contract, event: "RequestDeposit", logs: logs, sub: sub}, nil
}

// WatchRequestDeposit is a free log subscription operation binding the contract event 0xe84fe8b7649d8ddef364d5de9c5006de2baa5aed97143b907af093cf81a0dca4.
//
// Solidity: event RequestDeposit(uint64 taskId, (address,uint64,bytes32,string,uint256,string,uint256,uint256) depositInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchRequestDeposit(opts *bind.WatchOpts, sink chan<- *DepositManagerContractRequestDeposit) (event.Subscription, error) {

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "RequestDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractRequestDeposit)
				if err := _DepositManagerContract.contract.UnpackLog(event, "RequestDeposit", log); err != nil {
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

// ParseRequestDeposit is a log parse operation binding the contract event 0xe84fe8b7649d8ddef364d5de9c5006de2baa5aed97143b907af093cf81a0dca4.
//
// Solidity: event RequestDeposit(uint64 taskId, (address,uint64,bytes32,string,uint256,string,uint256,uint256) depositInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) ParseRequestDeposit(log types.Log) (*DepositManagerContractRequestDeposit, error) {
	event := new(DepositManagerContractRequestDeposit)
	if err := _DepositManagerContract.contract.UnpackLog(event, "RequestDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractRequestWithdrawalIterator is returned from FilterRequestWithdrawal and is used to iterate over the raw logs and unpacked data for RequestWithdrawal events raised by the DepositManagerContract contract.
type DepositManagerContractRequestWithdrawalIterator struct {
	Event *DepositManagerContractRequestWithdrawal // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractRequestWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractRequestWithdrawal)
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
		it.Event = new(DepositManagerContractRequestWithdrawal)
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
func (it *DepositManagerContractRequestWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractRequestWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractRequestWithdrawal represents a RequestWithdrawal event raised by the DepositManagerContract contract.
type DepositManagerContractRequestWithdrawal struct {
	TaskId         uint64
	WithdrawalInfo WithdrawalInfo
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdrawal is a free log retrieval operation binding the contract event 0x73995490aece71003a054222db40613bf2ebe8293cd52844d155a29eb47930f9.
//
// Solidity: event RequestWithdrawal(uint64 taskId, (address,uint64,bytes32,string,uint256) withdrawalInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterRequestWithdrawal(opts *bind.FilterOpts) (*DepositManagerContractRequestWithdrawalIterator, error) {

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "RequestWithdrawal")
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractRequestWithdrawalIterator{contract: _DepositManagerContract.contract, event: "RequestWithdrawal", logs: logs, sub: sub}, nil
}

// WatchRequestWithdrawal is a free log subscription operation binding the contract event 0x73995490aece71003a054222db40613bf2ebe8293cd52844d155a29eb47930f9.
//
// Solidity: event RequestWithdrawal(uint64 taskId, (address,uint64,bytes32,string,uint256) withdrawalInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchRequestWithdrawal(opts *bind.WatchOpts, sink chan<- *DepositManagerContractRequestWithdrawal) (event.Subscription, error) {

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "RequestWithdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractRequestWithdrawal)
				if err := _DepositManagerContract.contract.UnpackLog(event, "RequestWithdrawal", log); err != nil {
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

// ParseRequestWithdrawal is a log parse operation binding the contract event 0x73995490aece71003a054222db40613bf2ebe8293cd52844d155a29eb47930f9.
//
// Solidity: event RequestWithdrawal(uint64 taskId, (address,uint64,bytes32,string,uint256) withdrawalInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) ParseRequestWithdrawal(log types.Log) (*DepositManagerContractRequestWithdrawal, error) {
	event := new(DepositManagerContractRequestWithdrawal)
	if err := _DepositManagerContract.contract.UnpackLog(event, "RequestWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DepositManagerContract contract.
type DepositManagerContractRoleAdminChangedIterator struct {
	Event *DepositManagerContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractRoleAdminChanged)
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
		it.Event = new(DepositManagerContractRoleAdminChanged)
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
func (it *DepositManagerContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractRoleAdminChanged represents a RoleAdminChanged event raised by the DepositManagerContract contract.
type DepositManagerContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DepositManagerContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractRoleAdminChangedIterator{contract: _DepositManagerContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DepositManagerContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractRoleAdminChanged)
				if err := _DepositManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_DepositManagerContract *DepositManagerContractFilterer) ParseRoleAdminChanged(log types.Log) (*DepositManagerContractRoleAdminChanged, error) {
	event := new(DepositManagerContractRoleAdminChanged)
	if err := _DepositManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DepositManagerContract contract.
type DepositManagerContractRoleGrantedIterator struct {
	Event *DepositManagerContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractRoleGranted)
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
		it.Event = new(DepositManagerContractRoleGranted)
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
func (it *DepositManagerContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractRoleGranted represents a RoleGranted event raised by the DepositManagerContract contract.
type DepositManagerContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DepositManagerContractRoleGrantedIterator, error) {

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

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractRoleGrantedIterator{contract: _DepositManagerContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DepositManagerContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractRoleGranted)
				if err := _DepositManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_DepositManagerContract *DepositManagerContractFilterer) ParseRoleGranted(log types.Log) (*DepositManagerContractRoleGranted, error) {
	event := new(DepositManagerContractRoleGranted)
	if err := _DepositManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DepositManagerContract contract.
type DepositManagerContractRoleRevokedIterator struct {
	Event *DepositManagerContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractRoleRevoked)
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
		it.Event = new(DepositManagerContractRoleRevoked)
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
func (it *DepositManagerContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractRoleRevoked represents a RoleRevoked event raised by the DepositManagerContract contract.
type DepositManagerContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DepositManagerContractRoleRevokedIterator, error) {

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

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractRoleRevokedIterator{contract: _DepositManagerContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DepositManagerContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractRoleRevoked)
				if err := _DepositManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_DepositManagerContract *DepositManagerContractFilterer) ParseRoleRevoked(log types.Log) (*DepositManagerContractRoleRevoked, error) {
	event := new(DepositManagerContractRoleRevoked)
	if err := _DepositManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractWithdrawalRecordedIterator is returned from FilterWithdrawalRecorded and is used to iterate over the raw logs and unpacked data for WithdrawalRecorded events raised by the DepositManagerContract contract.
type DepositManagerContractWithdrawalRecordedIterator struct {
	Event *DepositManagerContractWithdrawalRecorded // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractWithdrawalRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractWithdrawalRecorded)
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
		it.Event = new(DepositManagerContractWithdrawalRecorded)
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
func (it *DepositManagerContractWithdrawalRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractWithdrawalRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractWithdrawalRecorded represents a WithdrawalRecorded event raised by the DepositManagerContract contract.
type DepositManagerContractWithdrawalRecorded struct {
	UserAddress common.Address
	Ticker      [32]byte
	ChainId     uint64
	ToAddress   string
	Amount      *big.Int
	TxHash      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRecorded is a free log retrieval operation binding the contract event 0x59a01e2a8864dd3fae2368280a0cb84e97d4960b6c44fdde8fda335c715ce9e4.
//
// Solidity: event WithdrawalRecorded(address indexed userAddress, bytes32 indexed ticker, uint64 indexed chainId, string toAddress, uint256 amount, string txHash)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterWithdrawalRecorded(opts *bind.FilterOpts, userAddress []common.Address, ticker [][32]byte, chainId []uint64) (*DepositManagerContractWithdrawalRecordedIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "WithdrawalRecorded", userAddressRule, tickerRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractWithdrawalRecordedIterator{contract: _DepositManagerContract.contract, event: "WithdrawalRecorded", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRecorded is a free log subscription operation binding the contract event 0x59a01e2a8864dd3fae2368280a0cb84e97d4960b6c44fdde8fda335c715ce9e4.
//
// Solidity: event WithdrawalRecorded(address indexed userAddress, bytes32 indexed ticker, uint64 indexed chainId, string toAddress, uint256 amount, string txHash)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchWithdrawalRecorded(opts *bind.WatchOpts, sink chan<- *DepositManagerContractWithdrawalRecorded, userAddress []common.Address, ticker [][32]byte, chainId []uint64) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "WithdrawalRecorded", userAddressRule, tickerRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractWithdrawalRecorded)
				if err := _DepositManagerContract.contract.UnpackLog(event, "WithdrawalRecorded", log); err != nil {
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

// ParseWithdrawalRecorded is a log parse operation binding the contract event 0x59a01e2a8864dd3fae2368280a0cb84e97d4960b6c44fdde8fda335c715ce9e4.
//
// Solidity: event WithdrawalRecorded(address indexed userAddress, bytes32 indexed ticker, uint64 indexed chainId, string toAddress, uint256 amount, string txHash)
func (_DepositManagerContract *DepositManagerContractFilterer) ParseWithdrawalRecorded(log types.Log) (*DepositManagerContractWithdrawalRecorded, error) {
	event := new(DepositManagerContractWithdrawalRecorded)
	if err := _DepositManagerContract.contract.UnpackLog(event, "WithdrawalRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
