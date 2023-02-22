// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

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
)

// TransitStructsCallbytesDescription is an auto generated low-level Go binding around an user-defined struct.
type TransitStructsCallbytesDescription struct {
	Flag      uint8
	SrcToken  common.Address
	Calldatas []byte
}

// TransitStructsTransitSwapDescription is an auto generated low-level Go binding around an user-defined struct.
type TransitStructsTransitSwapDescription struct {
	SwapType        uint8
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceiver     common.Address
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Channel         string
	ToChainID       *big.Int
	WrappedNative   common.Address
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transitSwap_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"transitCross_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"transitFees_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8[]\",\"name\":\"types\",\"type\":\"uint8[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"newModes\",\"type\":\"bool[]\"}],\"name\":\"ChangeSwapTypeMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousTransit\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTransit\",\"type\":\"address\"}],\"name\":\"ChangeTransitCross\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousTransitFees\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTransitFees\",\"type\":\"address\"}],\"name\":\"ChangeTransitFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousTransit\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTransit\",\"type\":\"address\"}],\"name\":\"ChangeTransitSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"wrappeds\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"newAllowed\",\"type\":\"bool[]\"}],\"name\":\"ChangeWrappedAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousExecutor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newExecutor\",\"type\":\"address\"}],\"name\":\"ExecutorshipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousExecutor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newExecutor\",\"type\":\"address\"}],\"name\":\"ExecutorshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Receipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"feeMode\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TransitSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptExecutorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"changePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"swapTypes\",\"type\":\"uint8[]\"}],\"name\":\"changeSwapTypeMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTransit\",\"type\":\"address\"}],\"name\":\"changeTransitCross\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTransitFees\",\"type\":\"address\"}],\"name\":\"changeTransitFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTransit\",\"type\":\"address\"}],\"name\":\"changeTransitSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"wrappeds\",\"type\":\"address[]\"}],\"name\":\"changeWrappedAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"internalType\":\"structTransitStructs.TransitSwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldatas\",\"type\":\"bytes\"}],\"internalType\":\"structTransitStructs.CallbytesDescription\",\"name\":\"callbytesDesc\",\"type\":\"tuple\"}],\"name\":\"cross\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"internalType\":\"structTransitStructs.TransitSwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldatas\",\"type\":\"bytes\"}],\"internalType\":\"structTransitStructs.CallbytesDescription\",\"name\":\"callbytesDesc\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"swapType\",\"type\":\"uint8\"}],\"name\":\"swapTypeMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newExecutor\",\"type\":\"address\"}],\"name\":\"transferExecutorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transitCross\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transitFees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transitSwap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrapped\",\"type\":\"address\"}],\"name\":\"wrappedAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Router *RouterCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Router *RouterSession) Executor() (common.Address, error) {
	return _Router.Contract.Executor(&_Router.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Router *RouterCallerSession) Executor() (common.Address, error) {
	return _Router.Contract.Executor(&_Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCallerSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Router *RouterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Router *RouterSession) Paused() (bool, error) {
	return _Router.Contract.Paused(&_Router.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Router *RouterCallerSession) Paused() (bool, error) {
	return _Router.Contract.Paused(&_Router.CallOpts)
}

// PendingExecutor is a free data retrieval call binding the contract method 0xd63234e0.
//
// Solidity: function pendingExecutor() view returns(address)
func (_Router *RouterCaller) PendingExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "pendingExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingExecutor is a free data retrieval call binding the contract method 0xd63234e0.
//
// Solidity: function pendingExecutor() view returns(address)
func (_Router *RouterSession) PendingExecutor() (common.Address, error) {
	return _Router.Contract.PendingExecutor(&_Router.CallOpts)
}

// PendingExecutor is a free data retrieval call binding the contract method 0xd63234e0.
//
// Solidity: function pendingExecutor() view returns(address)
func (_Router *RouterCallerSession) PendingExecutor() (common.Address, error) {
	return _Router.Contract.PendingExecutor(&_Router.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Router *RouterCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Router *RouterSession) PendingOwner() (common.Address, error) {
	return _Router.Contract.PendingOwner(&_Router.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Router *RouterCallerSession) PendingOwner() (common.Address, error) {
	return _Router.Contract.PendingOwner(&_Router.CallOpts)
}

// SwapTypeMode is a free data retrieval call binding the contract method 0xc58dff5c.
//
// Solidity: function swapTypeMode(uint8 swapType) view returns(bool)
func (_Router *RouterCaller) SwapTypeMode(opts *bind.CallOpts, swapType uint8) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "swapTypeMode", swapType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapTypeMode is a free data retrieval call binding the contract method 0xc58dff5c.
//
// Solidity: function swapTypeMode(uint8 swapType) view returns(bool)
func (_Router *RouterSession) SwapTypeMode(swapType uint8) (bool, error) {
	return _Router.Contract.SwapTypeMode(&_Router.CallOpts, swapType)
}

// SwapTypeMode is a free data retrieval call binding the contract method 0xc58dff5c.
//
// Solidity: function swapTypeMode(uint8 swapType) view returns(bool)
func (_Router *RouterCallerSession) SwapTypeMode(swapType uint8) (bool, error) {
	return _Router.Contract.SwapTypeMode(&_Router.CallOpts, swapType)
}

// TransitCross is a free data retrieval call binding the contract method 0x94752e82.
//
// Solidity: function transitCross() view returns(address)
func (_Router *RouterCaller) TransitCross(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "transitCross")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransitCross is a free data retrieval call binding the contract method 0x94752e82.
//
// Solidity: function transitCross() view returns(address)
func (_Router *RouterSession) TransitCross() (common.Address, error) {
	return _Router.Contract.TransitCross(&_Router.CallOpts)
}

// TransitCross is a free data retrieval call binding the contract method 0x94752e82.
//
// Solidity: function transitCross() view returns(address)
func (_Router *RouterCallerSession) TransitCross() (common.Address, error) {
	return _Router.Contract.TransitCross(&_Router.CallOpts)
}

// TransitFees is a free data retrieval call binding the contract method 0x51205349.
//
// Solidity: function transitFees() view returns(address)
func (_Router *RouterCaller) TransitFees(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "transitFees")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransitFees is a free data retrieval call binding the contract method 0x51205349.
//
// Solidity: function transitFees() view returns(address)
func (_Router *RouterSession) TransitFees() (common.Address, error) {
	return _Router.Contract.TransitFees(&_Router.CallOpts)
}

// TransitFees is a free data retrieval call binding the contract method 0x51205349.
//
// Solidity: function transitFees() view returns(address)
func (_Router *RouterCallerSession) TransitFees() (common.Address, error) {
	return _Router.Contract.TransitFees(&_Router.CallOpts)
}

// TransitSwap is a free data retrieval call binding the contract method 0x845b1f77.
//
// Solidity: function transitSwap() view returns(address)
func (_Router *RouterCaller) TransitSwap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "transitSwap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransitSwap is a free data retrieval call binding the contract method 0x845b1f77.
//
// Solidity: function transitSwap() view returns(address)
func (_Router *RouterSession) TransitSwap() (common.Address, error) {
	return _Router.Contract.TransitSwap(&_Router.CallOpts)
}

// TransitSwap is a free data retrieval call binding the contract method 0x845b1f77.
//
// Solidity: function transitSwap() view returns(address)
func (_Router *RouterCallerSession) TransitSwap() (common.Address, error) {
	return _Router.Contract.TransitSwap(&_Router.CallOpts)
}

// WrappedAllowed is a free data retrieval call binding the contract method 0x3932a93c.
//
// Solidity: function wrappedAllowed(address wrapped) view returns(bool)
func (_Router *RouterCaller) WrappedAllowed(opts *bind.CallOpts, wrapped common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "wrappedAllowed", wrapped)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WrappedAllowed is a free data retrieval call binding the contract method 0x3932a93c.
//
// Solidity: function wrappedAllowed(address wrapped) view returns(bool)
func (_Router *RouterSession) WrappedAllowed(wrapped common.Address) (bool, error) {
	return _Router.Contract.WrappedAllowed(&_Router.CallOpts, wrapped)
}

// WrappedAllowed is a free data retrieval call binding the contract method 0x3932a93c.
//
// Solidity: function wrappedAllowed(address wrapped) view returns(bool)
func (_Router *RouterCallerSession) WrappedAllowed(wrapped common.Address) (bool, error) {
	return _Router.Contract.WrappedAllowed(&_Router.CallOpts, wrapped)
}

// AcceptExecutorship is a paid mutator transaction binding the contract method 0x94d3d793.
//
// Solidity: function acceptExecutorship() returns()
func (_Router *RouterTransactor) AcceptExecutorship(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "acceptExecutorship")
}

// AcceptExecutorship is a paid mutator transaction binding the contract method 0x94d3d793.
//
// Solidity: function acceptExecutorship() returns()
func (_Router *RouterSession) AcceptExecutorship() (*types.Transaction, error) {
	return _Router.Contract.AcceptExecutorship(&_Router.TransactOpts)
}

// AcceptExecutorship is a paid mutator transaction binding the contract method 0x94d3d793.
//
// Solidity: function acceptExecutorship() returns()
func (_Router *RouterTransactorSession) AcceptExecutorship() (*types.Transaction, error) {
	return _Router.Contract.AcceptExecutorship(&_Router.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Router *RouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Router *RouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _Router.Contract.AcceptOwnership(&_Router.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Router *RouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Router.Contract.AcceptOwnership(&_Router.TransactOpts)
}

// ChangePause is a paid mutator transaction binding the contract method 0x5d4fead3.
//
// Solidity: function changePause(bool paused) returns()
func (_Router *RouterTransactor) ChangePause(opts *bind.TransactOpts, paused bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "changePause", paused)
}

// ChangePause is a paid mutator transaction binding the contract method 0x5d4fead3.
//
// Solidity: function changePause(bool paused) returns()
func (_Router *RouterSession) ChangePause(paused bool) (*types.Transaction, error) {
	return _Router.Contract.ChangePause(&_Router.TransactOpts, paused)
}

// ChangePause is a paid mutator transaction binding the contract method 0x5d4fead3.
//
// Solidity: function changePause(bool paused) returns()
func (_Router *RouterTransactorSession) ChangePause(paused bool) (*types.Transaction, error) {
	return _Router.Contract.ChangePause(&_Router.TransactOpts, paused)
}

// ChangeSwapTypeMode is a paid mutator transaction binding the contract method 0xc3dc51fe.
//
// Solidity: function changeSwapTypeMode(uint8[] swapTypes) returns()
func (_Router *RouterTransactor) ChangeSwapTypeMode(opts *bind.TransactOpts, swapTypes []uint8) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "changeSwapTypeMode", swapTypes)
}

// ChangeSwapTypeMode is a paid mutator transaction binding the contract method 0xc3dc51fe.
//
// Solidity: function changeSwapTypeMode(uint8[] swapTypes) returns()
func (_Router *RouterSession) ChangeSwapTypeMode(swapTypes []uint8) (*types.Transaction, error) {
	return _Router.Contract.ChangeSwapTypeMode(&_Router.TransactOpts, swapTypes)
}

// ChangeSwapTypeMode is a paid mutator transaction binding the contract method 0xc3dc51fe.
//
// Solidity: function changeSwapTypeMode(uint8[] swapTypes) returns()
func (_Router *RouterTransactorSession) ChangeSwapTypeMode(swapTypes []uint8) (*types.Transaction, error) {
	return _Router.Contract.ChangeSwapTypeMode(&_Router.TransactOpts, swapTypes)
}

// ChangeTransitCross is a paid mutator transaction binding the contract method 0xe8e1bc5d.
//
// Solidity: function changeTransitCross(address newTransit) returns()
func (_Router *RouterTransactor) ChangeTransitCross(opts *bind.TransactOpts, newTransit common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "changeTransitCross", newTransit)
}

// ChangeTransitCross is a paid mutator transaction binding the contract method 0xe8e1bc5d.
//
// Solidity: function changeTransitCross(address newTransit) returns()
func (_Router *RouterSession) ChangeTransitCross(newTransit common.Address) (*types.Transaction, error) {
	return _Router.Contract.ChangeTransitCross(&_Router.TransactOpts, newTransit)
}

// ChangeTransitCross is a paid mutator transaction binding the contract method 0xe8e1bc5d.
//
// Solidity: function changeTransitCross(address newTransit) returns()
func (_Router *RouterTransactorSession) ChangeTransitCross(newTransit common.Address) (*types.Transaction, error) {
	return _Router.Contract.ChangeTransitCross(&_Router.TransactOpts, newTransit)
}

// ChangeTransitFees is a paid mutator transaction binding the contract method 0xc6c46552.
//
// Solidity: function changeTransitFees(address newTransitFees) returns()
func (_Router *RouterTransactor) ChangeTransitFees(opts *bind.TransactOpts, newTransitFees common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "changeTransitFees", newTransitFees)
}

// ChangeTransitFees is a paid mutator transaction binding the contract method 0xc6c46552.
//
// Solidity: function changeTransitFees(address newTransitFees) returns()
func (_Router *RouterSession) ChangeTransitFees(newTransitFees common.Address) (*types.Transaction, error) {
	return _Router.Contract.ChangeTransitFees(&_Router.TransactOpts, newTransitFees)
}

// ChangeTransitFees is a paid mutator transaction binding the contract method 0xc6c46552.
//
// Solidity: function changeTransitFees(address newTransitFees) returns()
func (_Router *RouterTransactorSession) ChangeTransitFees(newTransitFees common.Address) (*types.Transaction, error) {
	return _Router.Contract.ChangeTransitFees(&_Router.TransactOpts, newTransitFees)
}

// ChangeTransitSwap is a paid mutator transaction binding the contract method 0xdc69bbad.
//
// Solidity: function changeTransitSwap(address newTransit) returns()
func (_Router *RouterTransactor) ChangeTransitSwap(opts *bind.TransactOpts, newTransit common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "changeTransitSwap", newTransit)
}

// ChangeTransitSwap is a paid mutator transaction binding the contract method 0xdc69bbad.
//
// Solidity: function changeTransitSwap(address newTransit) returns()
func (_Router *RouterSession) ChangeTransitSwap(newTransit common.Address) (*types.Transaction, error) {
	return _Router.Contract.ChangeTransitSwap(&_Router.TransactOpts, newTransit)
}

// ChangeTransitSwap is a paid mutator transaction binding the contract method 0xdc69bbad.
//
// Solidity: function changeTransitSwap(address newTransit) returns()
func (_Router *RouterTransactorSession) ChangeTransitSwap(newTransit common.Address) (*types.Transaction, error) {
	return _Router.Contract.ChangeTransitSwap(&_Router.TransactOpts, newTransit)
}

// ChangeWrappedAllowed is a paid mutator transaction binding the contract method 0xf8c79f73.
//
// Solidity: function changeWrappedAllowed(address[] wrappeds) returns()
func (_Router *RouterTransactor) ChangeWrappedAllowed(opts *bind.TransactOpts, wrappeds []common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "changeWrappedAllowed", wrappeds)
}

// ChangeWrappedAllowed is a paid mutator transaction binding the contract method 0xf8c79f73.
//
// Solidity: function changeWrappedAllowed(address[] wrappeds) returns()
func (_Router *RouterSession) ChangeWrappedAllowed(wrappeds []common.Address) (*types.Transaction, error) {
	return _Router.Contract.ChangeWrappedAllowed(&_Router.TransactOpts, wrappeds)
}

// ChangeWrappedAllowed is a paid mutator transaction binding the contract method 0xf8c79f73.
//
// Solidity: function changeWrappedAllowed(address[] wrappeds) returns()
func (_Router *RouterTransactorSession) ChangeWrappedAllowed(wrappeds []common.Address) (*types.Transaction, error) {
	return _Router.Contract.ChangeWrappedAllowed(&_Router.TransactOpts, wrappeds)
}

// Cross is a paid mutator transaction binding the contract method 0x017efb28.
//
// Solidity: function cross((uint8,address,address,address,address,uint256,uint256,string,uint256,address) desc, (uint8,address,bytes) callbytesDesc) payable returns()
func (_Router *RouterTransactor) Cross(opts *bind.TransactOpts, desc TransitStructsTransitSwapDescription, callbytesDesc TransitStructsCallbytesDescription) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "cross", desc, callbytesDesc)
}

// Cross is a paid mutator transaction binding the contract method 0x017efb28.
//
// Solidity: function cross((uint8,address,address,address,address,uint256,uint256,string,uint256,address) desc, (uint8,address,bytes) callbytesDesc) payable returns()
func (_Router *RouterSession) Cross(desc TransitStructsTransitSwapDescription, callbytesDesc TransitStructsCallbytesDescription) (*types.Transaction, error) {
	return _Router.Contract.Cross(&_Router.TransactOpts, desc, callbytesDesc)
}

// Cross is a paid mutator transaction binding the contract method 0x017efb28.
//
// Solidity: function cross((uint8,address,address,address,address,uint256,uint256,string,uint256,address) desc, (uint8,address,bytes) callbytesDesc) payable returns()
func (_Router *RouterTransactorSession) Cross(desc TransitStructsTransitSwapDescription, callbytesDesc TransitStructsCallbytesDescription) (*types.Transaction, error) {
	return _Router.Contract.Cross(&_Router.TransactOpts, desc, callbytesDesc)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Router.Contract.RenounceOwnership(&_Router.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Router.Contract.RenounceOwnership(&_Router.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xc10bea5c.
//
// Solidity: function swap((uint8,address,address,address,address,uint256,uint256,string,uint256,address) desc, (uint8,address,bytes) callbytesDesc) payable returns()
func (_Router *RouterTransactor) Swap(opts *bind.TransactOpts, desc TransitStructsTransitSwapDescription, callbytesDesc TransitStructsCallbytesDescription) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swap", desc, callbytesDesc)
}

// Swap is a paid mutator transaction binding the contract method 0xc10bea5c.
//
// Solidity: function swap((uint8,address,address,address,address,uint256,uint256,string,uint256,address) desc, (uint8,address,bytes) callbytesDesc) payable returns()
func (_Router *RouterSession) Swap(desc TransitStructsTransitSwapDescription, callbytesDesc TransitStructsCallbytesDescription) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, desc, callbytesDesc)
}

// Swap is a paid mutator transaction binding the contract method 0xc10bea5c.
//
// Solidity: function swap((uint8,address,address,address,address,uint256,uint256,string,uint256,address) desc, (uint8,address,bytes) callbytesDesc) payable returns()
func (_Router *RouterTransactorSession) Swap(desc TransitStructsTransitSwapDescription, callbytesDesc TransitStructsCallbytesDescription) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, desc, callbytesDesc)
}

// TransferExecutorship is a paid mutator transaction binding the contract method 0xafed2d0e.
//
// Solidity: function transferExecutorship(address newExecutor) returns()
func (_Router *RouterTransactor) TransferExecutorship(opts *bind.TransactOpts, newExecutor common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "transferExecutorship", newExecutor)
}

// TransferExecutorship is a paid mutator transaction binding the contract method 0xafed2d0e.
//
// Solidity: function transferExecutorship(address newExecutor) returns()
func (_Router *RouterSession) TransferExecutorship(newExecutor common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferExecutorship(&_Router.TransactOpts, newExecutor)
}

// TransferExecutorship is a paid mutator transaction binding the contract method 0xafed2d0e.
//
// Solidity: function transferExecutorship(address newExecutor) returns()
func (_Router *RouterTransactorSession) TransferExecutorship(newExecutor common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferExecutorship(&_Router.TransactOpts, newExecutor)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x0e8cc705.
//
// Solidity: function withdrawTokens(address[] tokens, address recipient) returns()
func (_Router *RouterTransactor) WithdrawTokens(opts *bind.TransactOpts, tokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "withdrawTokens", tokens, recipient)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x0e8cc705.
//
// Solidity: function withdrawTokens(address[] tokens, address recipient) returns()
func (_Router *RouterSession) WithdrawTokens(tokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Router.Contract.WithdrawTokens(&_Router.TransactOpts, tokens, recipient)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x0e8cc705.
//
// Solidity: function withdrawTokens(address[] tokens, address recipient) returns()
func (_Router *RouterTransactorSession) WithdrawTokens(tokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Router.Contract.WithdrawTokens(&_Router.TransactOpts, tokens, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// RouterChangeSwapTypeModeIterator is returned from FilterChangeSwapTypeMode and is used to iterate over the raw logs and unpacked data for ChangeSwapTypeMode events raised by the Router contract.
type RouterChangeSwapTypeModeIterator struct {
	Event *RouterChangeSwapTypeMode // Event containing the contract specifics and raw log

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
func (it *RouterChangeSwapTypeModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterChangeSwapTypeMode)
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
		it.Event = new(RouterChangeSwapTypeMode)
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
func (it *RouterChangeSwapTypeModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterChangeSwapTypeModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterChangeSwapTypeMode represents a ChangeSwapTypeMode event raised by the Router contract.
type RouterChangeSwapTypeMode struct {
	Types    []uint8
	NewModes []bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeSwapTypeMode is a free log retrieval operation binding the contract event 0x5e69b8fb6404283a560ee2205dc0f225019578c56e3a6c52137e0aa995235b66.
//
// Solidity: event ChangeSwapTypeMode(uint8[] types, bool[] newModes)
func (_Router *RouterFilterer) FilterChangeSwapTypeMode(opts *bind.FilterOpts) (*RouterChangeSwapTypeModeIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "ChangeSwapTypeMode")
	if err != nil {
		return nil, err
	}
	return &RouterChangeSwapTypeModeIterator{contract: _Router.contract, event: "ChangeSwapTypeMode", logs: logs, sub: sub}, nil
}

// WatchChangeSwapTypeMode is a free log subscription operation binding the contract event 0x5e69b8fb6404283a560ee2205dc0f225019578c56e3a6c52137e0aa995235b66.
//
// Solidity: event ChangeSwapTypeMode(uint8[] types, bool[] newModes)
func (_Router *RouterFilterer) WatchChangeSwapTypeMode(opts *bind.WatchOpts, sink chan<- *RouterChangeSwapTypeMode) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "ChangeSwapTypeMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterChangeSwapTypeMode)
				if err := _Router.contract.UnpackLog(event, "ChangeSwapTypeMode", log); err != nil {
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

// ParseChangeSwapTypeMode is a log parse operation binding the contract event 0x5e69b8fb6404283a560ee2205dc0f225019578c56e3a6c52137e0aa995235b66.
//
// Solidity: event ChangeSwapTypeMode(uint8[] types, bool[] newModes)
func (_Router *RouterFilterer) ParseChangeSwapTypeMode(log types.Log) (*RouterChangeSwapTypeMode, error) {
	event := new(RouterChangeSwapTypeMode)
	if err := _Router.contract.UnpackLog(event, "ChangeSwapTypeMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterChangeTransitCrossIterator is returned from FilterChangeTransitCross and is used to iterate over the raw logs and unpacked data for ChangeTransitCross events raised by the Router contract.
type RouterChangeTransitCrossIterator struct {
	Event *RouterChangeTransitCross // Event containing the contract specifics and raw log

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
func (it *RouterChangeTransitCrossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterChangeTransitCross)
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
		it.Event = new(RouterChangeTransitCross)
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
func (it *RouterChangeTransitCrossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterChangeTransitCrossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterChangeTransitCross represents a ChangeTransitCross event raised by the Router contract.
type RouterChangeTransitCross struct {
	PreviousTransit common.Address
	NewTransit      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChangeTransitCross is a free log retrieval operation binding the contract event 0xf2ef974bc2aedf0c322b1a97e8627e85114a789cadf8d172b5efb5d6c1d94f8d.
//
// Solidity: event ChangeTransitCross(address indexed previousTransit, address indexed newTransit)
func (_Router *RouterFilterer) FilterChangeTransitCross(opts *bind.FilterOpts, previousTransit []common.Address, newTransit []common.Address) (*RouterChangeTransitCrossIterator, error) {

	var previousTransitRule []interface{}
	for _, previousTransitItem := range previousTransit {
		previousTransitRule = append(previousTransitRule, previousTransitItem)
	}
	var newTransitRule []interface{}
	for _, newTransitItem := range newTransit {
		newTransitRule = append(newTransitRule, newTransitItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "ChangeTransitCross", previousTransitRule, newTransitRule)
	if err != nil {
		return nil, err
	}
	return &RouterChangeTransitCrossIterator{contract: _Router.contract, event: "ChangeTransitCross", logs: logs, sub: sub}, nil
}

// WatchChangeTransitCross is a free log subscription operation binding the contract event 0xf2ef974bc2aedf0c322b1a97e8627e85114a789cadf8d172b5efb5d6c1d94f8d.
//
// Solidity: event ChangeTransitCross(address indexed previousTransit, address indexed newTransit)
func (_Router *RouterFilterer) WatchChangeTransitCross(opts *bind.WatchOpts, sink chan<- *RouterChangeTransitCross, previousTransit []common.Address, newTransit []common.Address) (event.Subscription, error) {

	var previousTransitRule []interface{}
	for _, previousTransitItem := range previousTransit {
		previousTransitRule = append(previousTransitRule, previousTransitItem)
	}
	var newTransitRule []interface{}
	for _, newTransitItem := range newTransit {
		newTransitRule = append(newTransitRule, newTransitItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "ChangeTransitCross", previousTransitRule, newTransitRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterChangeTransitCross)
				if err := _Router.contract.UnpackLog(event, "ChangeTransitCross", log); err != nil {
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

// ParseChangeTransitCross is a log parse operation binding the contract event 0xf2ef974bc2aedf0c322b1a97e8627e85114a789cadf8d172b5efb5d6c1d94f8d.
//
// Solidity: event ChangeTransitCross(address indexed previousTransit, address indexed newTransit)
func (_Router *RouterFilterer) ParseChangeTransitCross(log types.Log) (*RouterChangeTransitCross, error) {
	event := new(RouterChangeTransitCross)
	if err := _Router.contract.UnpackLog(event, "ChangeTransitCross", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterChangeTransitFeesIterator is returned from FilterChangeTransitFees and is used to iterate over the raw logs and unpacked data for ChangeTransitFees events raised by the Router contract.
type RouterChangeTransitFeesIterator struct {
	Event *RouterChangeTransitFees // Event containing the contract specifics and raw log

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
func (it *RouterChangeTransitFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterChangeTransitFees)
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
		it.Event = new(RouterChangeTransitFees)
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
func (it *RouterChangeTransitFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterChangeTransitFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterChangeTransitFees represents a ChangeTransitFees event raised by the Router contract.
type RouterChangeTransitFees struct {
	PreviousTransitFees common.Address
	NewTransitFees      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterChangeTransitFees is a free log retrieval operation binding the contract event 0xaf855a7e4074e8930caf657b3505942d4ba0e69b6b6e63553c653d7b1ffcc061.
//
// Solidity: event ChangeTransitFees(address indexed previousTransitFees, address indexed newTransitFees)
func (_Router *RouterFilterer) FilterChangeTransitFees(opts *bind.FilterOpts, previousTransitFees []common.Address, newTransitFees []common.Address) (*RouterChangeTransitFeesIterator, error) {

	var previousTransitFeesRule []interface{}
	for _, previousTransitFeesItem := range previousTransitFees {
		previousTransitFeesRule = append(previousTransitFeesRule, previousTransitFeesItem)
	}
	var newTransitFeesRule []interface{}
	for _, newTransitFeesItem := range newTransitFees {
		newTransitFeesRule = append(newTransitFeesRule, newTransitFeesItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "ChangeTransitFees", previousTransitFeesRule, newTransitFeesRule)
	if err != nil {
		return nil, err
	}
	return &RouterChangeTransitFeesIterator{contract: _Router.contract, event: "ChangeTransitFees", logs: logs, sub: sub}, nil
}

// WatchChangeTransitFees is a free log subscription operation binding the contract event 0xaf855a7e4074e8930caf657b3505942d4ba0e69b6b6e63553c653d7b1ffcc061.
//
// Solidity: event ChangeTransitFees(address indexed previousTransitFees, address indexed newTransitFees)
func (_Router *RouterFilterer) WatchChangeTransitFees(opts *bind.WatchOpts, sink chan<- *RouterChangeTransitFees, previousTransitFees []common.Address, newTransitFees []common.Address) (event.Subscription, error) {

	var previousTransitFeesRule []interface{}
	for _, previousTransitFeesItem := range previousTransitFees {
		previousTransitFeesRule = append(previousTransitFeesRule, previousTransitFeesItem)
	}
	var newTransitFeesRule []interface{}
	for _, newTransitFeesItem := range newTransitFees {
		newTransitFeesRule = append(newTransitFeesRule, newTransitFeesItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "ChangeTransitFees", previousTransitFeesRule, newTransitFeesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterChangeTransitFees)
				if err := _Router.contract.UnpackLog(event, "ChangeTransitFees", log); err != nil {
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

// ParseChangeTransitFees is a log parse operation binding the contract event 0xaf855a7e4074e8930caf657b3505942d4ba0e69b6b6e63553c653d7b1ffcc061.
//
// Solidity: event ChangeTransitFees(address indexed previousTransitFees, address indexed newTransitFees)
func (_Router *RouterFilterer) ParseChangeTransitFees(log types.Log) (*RouterChangeTransitFees, error) {
	event := new(RouterChangeTransitFees)
	if err := _Router.contract.UnpackLog(event, "ChangeTransitFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterChangeTransitSwapIterator is returned from FilterChangeTransitSwap and is used to iterate over the raw logs and unpacked data for ChangeTransitSwap events raised by the Router contract.
type RouterChangeTransitSwapIterator struct {
	Event *RouterChangeTransitSwap // Event containing the contract specifics and raw log

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
func (it *RouterChangeTransitSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterChangeTransitSwap)
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
		it.Event = new(RouterChangeTransitSwap)
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
func (it *RouterChangeTransitSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterChangeTransitSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterChangeTransitSwap represents a ChangeTransitSwap event raised by the Router contract.
type RouterChangeTransitSwap struct {
	PreviousTransit common.Address
	NewTransit      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChangeTransitSwap is a free log retrieval operation binding the contract event 0xbc071769e275808816d09db519c1888c355f7557cb63095ce21e26a97a5f0e9c.
//
// Solidity: event ChangeTransitSwap(address indexed previousTransit, address indexed newTransit)
func (_Router *RouterFilterer) FilterChangeTransitSwap(opts *bind.FilterOpts, previousTransit []common.Address, newTransit []common.Address) (*RouterChangeTransitSwapIterator, error) {

	var previousTransitRule []interface{}
	for _, previousTransitItem := range previousTransit {
		previousTransitRule = append(previousTransitRule, previousTransitItem)
	}
	var newTransitRule []interface{}
	for _, newTransitItem := range newTransit {
		newTransitRule = append(newTransitRule, newTransitItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "ChangeTransitSwap", previousTransitRule, newTransitRule)
	if err != nil {
		return nil, err
	}
	return &RouterChangeTransitSwapIterator{contract: _Router.contract, event: "ChangeTransitSwap", logs: logs, sub: sub}, nil
}

// WatchChangeTransitSwap is a free log subscription operation binding the contract event 0xbc071769e275808816d09db519c1888c355f7557cb63095ce21e26a97a5f0e9c.
//
// Solidity: event ChangeTransitSwap(address indexed previousTransit, address indexed newTransit)
func (_Router *RouterFilterer) WatchChangeTransitSwap(opts *bind.WatchOpts, sink chan<- *RouterChangeTransitSwap, previousTransit []common.Address, newTransit []common.Address) (event.Subscription, error) {

	var previousTransitRule []interface{}
	for _, previousTransitItem := range previousTransit {
		previousTransitRule = append(previousTransitRule, previousTransitItem)
	}
	var newTransitRule []interface{}
	for _, newTransitItem := range newTransit {
		newTransitRule = append(newTransitRule, newTransitItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "ChangeTransitSwap", previousTransitRule, newTransitRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterChangeTransitSwap)
				if err := _Router.contract.UnpackLog(event, "ChangeTransitSwap", log); err != nil {
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

// ParseChangeTransitSwap is a log parse operation binding the contract event 0xbc071769e275808816d09db519c1888c355f7557cb63095ce21e26a97a5f0e9c.
//
// Solidity: event ChangeTransitSwap(address indexed previousTransit, address indexed newTransit)
func (_Router *RouterFilterer) ParseChangeTransitSwap(log types.Log) (*RouterChangeTransitSwap, error) {
	event := new(RouterChangeTransitSwap)
	if err := _Router.contract.UnpackLog(event, "ChangeTransitSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterChangeWrappedAllowedIterator is returned from FilterChangeWrappedAllowed and is used to iterate over the raw logs and unpacked data for ChangeWrappedAllowed events raised by the Router contract.
type RouterChangeWrappedAllowedIterator struct {
	Event *RouterChangeWrappedAllowed // Event containing the contract specifics and raw log

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
func (it *RouterChangeWrappedAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterChangeWrappedAllowed)
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
		it.Event = new(RouterChangeWrappedAllowed)
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
func (it *RouterChangeWrappedAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterChangeWrappedAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterChangeWrappedAllowed represents a ChangeWrappedAllowed event raised by the Router contract.
type RouterChangeWrappedAllowed struct {
	Wrappeds   []common.Address
	NewAllowed []bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChangeWrappedAllowed is a free log retrieval operation binding the contract event 0x4a28b173d9bc739be3886d172e07fef80392184787fc6b92406ce0f0c05b7e63.
//
// Solidity: event ChangeWrappedAllowed(address[] wrappeds, bool[] newAllowed)
func (_Router *RouterFilterer) FilterChangeWrappedAllowed(opts *bind.FilterOpts) (*RouterChangeWrappedAllowedIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "ChangeWrappedAllowed")
	if err != nil {
		return nil, err
	}
	return &RouterChangeWrappedAllowedIterator{contract: _Router.contract, event: "ChangeWrappedAllowed", logs: logs, sub: sub}, nil
}

// WatchChangeWrappedAllowed is a free log subscription operation binding the contract event 0x4a28b173d9bc739be3886d172e07fef80392184787fc6b92406ce0f0c05b7e63.
//
// Solidity: event ChangeWrappedAllowed(address[] wrappeds, bool[] newAllowed)
func (_Router *RouterFilterer) WatchChangeWrappedAllowed(opts *bind.WatchOpts, sink chan<- *RouterChangeWrappedAllowed) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "ChangeWrappedAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterChangeWrappedAllowed)
				if err := _Router.contract.UnpackLog(event, "ChangeWrappedAllowed", log); err != nil {
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

// ParseChangeWrappedAllowed is a log parse operation binding the contract event 0x4a28b173d9bc739be3886d172e07fef80392184787fc6b92406ce0f0c05b7e63.
//
// Solidity: event ChangeWrappedAllowed(address[] wrappeds, bool[] newAllowed)
func (_Router *RouterFilterer) ParseChangeWrappedAllowed(log types.Log) (*RouterChangeWrappedAllowed, error) {
	event := new(RouterChangeWrappedAllowed)
	if err := _Router.contract.UnpackLog(event, "ChangeWrappedAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterExecutorshipTransferStartedIterator is returned from FilterExecutorshipTransferStarted and is used to iterate over the raw logs and unpacked data for ExecutorshipTransferStarted events raised by the Router contract.
type RouterExecutorshipTransferStartedIterator struct {
	Event *RouterExecutorshipTransferStarted // Event containing the contract specifics and raw log

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
func (it *RouterExecutorshipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterExecutorshipTransferStarted)
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
		it.Event = new(RouterExecutorshipTransferStarted)
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
func (it *RouterExecutorshipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterExecutorshipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterExecutorshipTransferStarted represents a ExecutorshipTransferStarted event raised by the Router contract.
type RouterExecutorshipTransferStarted struct {
	PreviousExecutor common.Address
	NewExecutor      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutorshipTransferStarted is a free log retrieval operation binding the contract event 0xdd01547fc40682edc3cd8d164d53f5a1ae6b46138a83f045658ed760823ddba8.
//
// Solidity: event ExecutorshipTransferStarted(address indexed previousExecutor, address indexed newExecutor)
func (_Router *RouterFilterer) FilterExecutorshipTransferStarted(opts *bind.FilterOpts, previousExecutor []common.Address, newExecutor []common.Address) (*RouterExecutorshipTransferStartedIterator, error) {

	var previousExecutorRule []interface{}
	for _, previousExecutorItem := range previousExecutor {
		previousExecutorRule = append(previousExecutorRule, previousExecutorItem)
	}
	var newExecutorRule []interface{}
	for _, newExecutorItem := range newExecutor {
		newExecutorRule = append(newExecutorRule, newExecutorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "ExecutorshipTransferStarted", previousExecutorRule, newExecutorRule)
	if err != nil {
		return nil, err
	}
	return &RouterExecutorshipTransferStartedIterator{contract: _Router.contract, event: "ExecutorshipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchExecutorshipTransferStarted is a free log subscription operation binding the contract event 0xdd01547fc40682edc3cd8d164d53f5a1ae6b46138a83f045658ed760823ddba8.
//
// Solidity: event ExecutorshipTransferStarted(address indexed previousExecutor, address indexed newExecutor)
func (_Router *RouterFilterer) WatchExecutorshipTransferStarted(opts *bind.WatchOpts, sink chan<- *RouterExecutorshipTransferStarted, previousExecutor []common.Address, newExecutor []common.Address) (event.Subscription, error) {

	var previousExecutorRule []interface{}
	for _, previousExecutorItem := range previousExecutor {
		previousExecutorRule = append(previousExecutorRule, previousExecutorItem)
	}
	var newExecutorRule []interface{}
	for _, newExecutorItem := range newExecutor {
		newExecutorRule = append(newExecutorRule, newExecutorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "ExecutorshipTransferStarted", previousExecutorRule, newExecutorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterExecutorshipTransferStarted)
				if err := _Router.contract.UnpackLog(event, "ExecutorshipTransferStarted", log); err != nil {
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

// ParseExecutorshipTransferStarted is a log parse operation binding the contract event 0xdd01547fc40682edc3cd8d164d53f5a1ae6b46138a83f045658ed760823ddba8.
//
// Solidity: event ExecutorshipTransferStarted(address indexed previousExecutor, address indexed newExecutor)
func (_Router *RouterFilterer) ParseExecutorshipTransferStarted(log types.Log) (*RouterExecutorshipTransferStarted, error) {
	event := new(RouterExecutorshipTransferStarted)
	if err := _Router.contract.UnpackLog(event, "ExecutorshipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterExecutorshipTransferredIterator is returned from FilterExecutorshipTransferred and is used to iterate over the raw logs and unpacked data for ExecutorshipTransferred events raised by the Router contract.
type RouterExecutorshipTransferredIterator struct {
	Event *RouterExecutorshipTransferred // Event containing the contract specifics and raw log

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
func (it *RouterExecutorshipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterExecutorshipTransferred)
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
		it.Event = new(RouterExecutorshipTransferred)
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
func (it *RouterExecutorshipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterExecutorshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterExecutorshipTransferred represents a ExecutorshipTransferred event raised by the Router contract.
type RouterExecutorshipTransferred struct {
	PreviousExecutor common.Address
	NewExecutor      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutorshipTransferred is a free log retrieval operation binding the contract event 0x88436636ea40d5bb1bcc55ff9cd54788af71da886f4147a87f199adcca733d4d.
//
// Solidity: event ExecutorshipTransferred(address indexed previousExecutor, address indexed newExecutor)
func (_Router *RouterFilterer) FilterExecutorshipTransferred(opts *bind.FilterOpts, previousExecutor []common.Address, newExecutor []common.Address) (*RouterExecutorshipTransferredIterator, error) {

	var previousExecutorRule []interface{}
	for _, previousExecutorItem := range previousExecutor {
		previousExecutorRule = append(previousExecutorRule, previousExecutorItem)
	}
	var newExecutorRule []interface{}
	for _, newExecutorItem := range newExecutor {
		newExecutorRule = append(newExecutorRule, newExecutorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "ExecutorshipTransferred", previousExecutorRule, newExecutorRule)
	if err != nil {
		return nil, err
	}
	return &RouterExecutorshipTransferredIterator{contract: _Router.contract, event: "ExecutorshipTransferred", logs: logs, sub: sub}, nil
}

// WatchExecutorshipTransferred is a free log subscription operation binding the contract event 0x88436636ea40d5bb1bcc55ff9cd54788af71da886f4147a87f199adcca733d4d.
//
// Solidity: event ExecutorshipTransferred(address indexed previousExecutor, address indexed newExecutor)
func (_Router *RouterFilterer) WatchExecutorshipTransferred(opts *bind.WatchOpts, sink chan<- *RouterExecutorshipTransferred, previousExecutor []common.Address, newExecutor []common.Address) (event.Subscription, error) {

	var previousExecutorRule []interface{}
	for _, previousExecutorItem := range previousExecutor {
		previousExecutorRule = append(previousExecutorRule, previousExecutorItem)
	}
	var newExecutorRule []interface{}
	for _, newExecutorItem := range newExecutor {
		newExecutorRule = append(newExecutorRule, newExecutorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "ExecutorshipTransferred", previousExecutorRule, newExecutorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterExecutorshipTransferred)
				if err := _Router.contract.UnpackLog(event, "ExecutorshipTransferred", log); err != nil {
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

// ParseExecutorshipTransferred is a log parse operation binding the contract event 0x88436636ea40d5bb1bcc55ff9cd54788af71da886f4147a87f199adcca733d4d.
//
// Solidity: event ExecutorshipTransferred(address indexed previousExecutor, address indexed newExecutor)
func (_Router *RouterFilterer) ParseExecutorshipTransferred(log types.Log) (*RouterExecutorshipTransferred, error) {
	event := new(RouterExecutorshipTransferred)
	if err := _Router.contract.UnpackLog(event, "ExecutorshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Router contract.
type RouterOwnershipTransferStartedIterator struct {
	Event *RouterOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *RouterOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnershipTransferStarted)
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
		it.Event = new(RouterOwnershipTransferStarted)
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
func (it *RouterOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Router contract.
type RouterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RouterOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnershipTransferStartedIterator{contract: _Router.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterOwnershipTransferStarted)
				if err := _Router.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) ParseOwnershipTransferStarted(log types.Log) (*RouterOwnershipTransferStarted, error) {
	event := new(RouterOwnershipTransferStarted)
	if err := _Router.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Router contract.
type RouterOwnershipTransferredIterator struct {
	Event *RouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnershipTransferred)
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
		it.Event = new(RouterOwnershipTransferred)
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
func (it *RouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterOwnershipTransferred represents a OwnershipTransferred event raised by the Router contract.
type RouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnershipTransferredIterator{contract: _Router.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterOwnershipTransferred)
				if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) ParseOwnershipTransferred(log types.Log) (*RouterOwnershipTransferred, error) {
	event := new(RouterOwnershipTransferred)
	if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Router contract.
type RouterPausedIterator struct {
	Event *RouterPaused // Event containing the contract specifics and raw log

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
func (it *RouterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterPaused)
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
		it.Event = new(RouterPaused)
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
func (it *RouterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterPaused represents a Paused event raised by the Router contract.
type RouterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Router *RouterFilterer) FilterPaused(opts *bind.FilterOpts) (*RouterPausedIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RouterPausedIterator{contract: _Router.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Router *RouterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RouterPaused) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterPaused)
				if err := _Router.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Router *RouterFilterer) ParsePaused(log types.Log) (*RouterPaused, error) {
	event := new(RouterPaused)
	if err := _Router.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterReceiptIterator is returned from FilterReceipt and is used to iterate over the raw logs and unpacked data for Receipt events raised by the Router contract.
type RouterReceiptIterator struct {
	Event *RouterReceipt // Event containing the contract specifics and raw log

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
func (it *RouterReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterReceipt)
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
		it.Event = new(RouterReceipt)
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
func (it *RouterReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterReceipt represents a Receipt event raised by the Router contract.
type RouterReceipt struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceipt is a free log retrieval operation binding the contract event 0x7784f8d436dc514f0690e472c7e2d7f660a73e504c69b2350f6be5a5f02432ef.
//
// Solidity: event Receipt(address from, uint256 amount)
func (_Router *RouterFilterer) FilterReceipt(opts *bind.FilterOpts) (*RouterReceiptIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "Receipt")
	if err != nil {
		return nil, err
	}
	return &RouterReceiptIterator{contract: _Router.contract, event: "Receipt", logs: logs, sub: sub}, nil
}

// WatchReceipt is a free log subscription operation binding the contract event 0x7784f8d436dc514f0690e472c7e2d7f660a73e504c69b2350f6be5a5f02432ef.
//
// Solidity: event Receipt(address from, uint256 amount)
func (_Router *RouterFilterer) WatchReceipt(opts *bind.WatchOpts, sink chan<- *RouterReceipt) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "Receipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterReceipt)
				if err := _Router.contract.UnpackLog(event, "Receipt", log); err != nil {
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

// ParseReceipt is a log parse operation binding the contract event 0x7784f8d436dc514f0690e472c7e2d7f660a73e504c69b2350f6be5a5f02432ef.
//
// Solidity: event Receipt(address from, uint256 amount)
func (_Router *RouterFilterer) ParseReceipt(log types.Log) (*RouterReceipt, error) {
	event := new(RouterReceipt)
	if err := _Router.contract.UnpackLog(event, "Receipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterTransitSwappedIterator is returned from FilterTransitSwapped and is used to iterate over the raw logs and unpacked data for TransitSwapped events raised by the Router contract.
type RouterTransitSwappedIterator struct {
	Event *RouterTransitSwapped // Event containing the contract specifics and raw log

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
func (it *RouterTransitSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterTransitSwapped)
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
		it.Event = new(RouterTransitSwapped)
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
func (it *RouterTransitSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterTransitSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterTransitSwapped represents a TransitSwapped event raised by the Router contract.
type RouterTransitSwapped struct {
	SrcToken        common.Address
	DstToken        common.Address
	DstReceiver     common.Address
	Trader          common.Address
	FeeMode         bool
	Amount          *big.Int
	ReturnAmount    *big.Int
	MinReturnAmount *big.Int
	Fee             *big.Int
	ToChainID       *big.Int
	Channel         string
	Time            *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransitSwapped is a free log retrieval operation binding the contract event 0x7055e3d08e2c20429c6b162f3e3bee3f426d59896e66084c3580dc353e54129d.
//
// Solidity: event TransitSwapped(address indexed srcToken, address indexed dstToken, address indexed dstReceiver, address trader, bool feeMode, uint256 amount, uint256 returnAmount, uint256 minReturnAmount, uint256 fee, uint256 toChainID, string channel, uint256 time)
func (_Router *RouterFilterer) FilterTransitSwapped(opts *bind.FilterOpts, srcToken []common.Address, dstToken []common.Address, dstReceiver []common.Address) (*RouterTransitSwappedIterator, error) {

	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var dstTokenRule []interface{}
	for _, dstTokenItem := range dstToken {
		dstTokenRule = append(dstTokenRule, dstTokenItem)
	}
	var dstReceiverRule []interface{}
	for _, dstReceiverItem := range dstReceiver {
		dstReceiverRule = append(dstReceiverRule, dstReceiverItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "TransitSwapped", srcTokenRule, dstTokenRule, dstReceiverRule)
	if err != nil {
		return nil, err
	}
	return &RouterTransitSwappedIterator{contract: _Router.contract, event: "TransitSwapped", logs: logs, sub: sub}, nil
}

// WatchTransitSwapped is a free log subscription operation binding the contract event 0x7055e3d08e2c20429c6b162f3e3bee3f426d59896e66084c3580dc353e54129d.
//
// Solidity: event TransitSwapped(address indexed srcToken, address indexed dstToken, address indexed dstReceiver, address trader, bool feeMode, uint256 amount, uint256 returnAmount, uint256 minReturnAmount, uint256 fee, uint256 toChainID, string channel, uint256 time)
func (_Router *RouterFilterer) WatchTransitSwapped(opts *bind.WatchOpts, sink chan<- *RouterTransitSwapped, srcToken []common.Address, dstToken []common.Address, dstReceiver []common.Address) (event.Subscription, error) {

	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var dstTokenRule []interface{}
	for _, dstTokenItem := range dstToken {
		dstTokenRule = append(dstTokenRule, dstTokenItem)
	}
	var dstReceiverRule []interface{}
	for _, dstReceiverItem := range dstReceiver {
		dstReceiverRule = append(dstReceiverRule, dstReceiverItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "TransitSwapped", srcTokenRule, dstTokenRule, dstReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterTransitSwapped)
				if err := _Router.contract.UnpackLog(event, "TransitSwapped", log); err != nil {
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

// ParseTransitSwapped is a log parse operation binding the contract event 0x7055e3d08e2c20429c6b162f3e3bee3f426d59896e66084c3580dc353e54129d.
//
// Solidity: event TransitSwapped(address indexed srcToken, address indexed dstToken, address indexed dstReceiver, address trader, bool feeMode, uint256 amount, uint256 returnAmount, uint256 minReturnAmount, uint256 fee, uint256 toChainID, string channel, uint256 time)
func (_Router *RouterFilterer) ParseTransitSwapped(log types.Log) (*RouterTransitSwapped, error) {
	event := new(RouterTransitSwapped)
	if err := _Router.contract.UnpackLog(event, "TransitSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Router contract.
type RouterUnpausedIterator struct {
	Event *RouterUnpaused // Event containing the contract specifics and raw log

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
func (it *RouterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterUnpaused)
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
		it.Event = new(RouterUnpaused)
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
func (it *RouterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterUnpaused represents a Unpaused event raised by the Router contract.
type RouterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Router *RouterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RouterUnpausedIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RouterUnpausedIterator{contract: _Router.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Router *RouterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RouterUnpaused) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterUnpaused)
				if err := _Router.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Router *RouterFilterer) ParseUnpaused(log types.Log) (*RouterUnpaused, error) {
	event := new(RouterUnpaused)
	if err := _Router.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Router contract.
type RouterWithdrawIterator struct {
	Event *RouterWithdraw // Event containing the contract specifics and raw log

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
func (it *RouterWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterWithdraw)
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
		it.Event = new(RouterWithdraw)
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
func (it *RouterWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterWithdraw represents a Withdraw event raised by the Router contract.
type RouterWithdraw struct {
	Token     common.Address
	Executor  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed token, address indexed executor, address indexed recipient, uint256 amount)
func (_Router *RouterFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, executor []common.Address, recipient []common.Address) (*RouterWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "Withdraw", tokenRule, executorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RouterWithdrawIterator{contract: _Router.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed token, address indexed executor, address indexed recipient, uint256 amount)
func (_Router *RouterFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *RouterWithdraw, token []common.Address, executor []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "Withdraw", tokenRule, executorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterWithdraw)
				if err := _Router.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed token, address indexed executor, address indexed recipient, uint256 amount)
func (_Router *RouterFilterer) ParseWithdraw(log types.Log) (*RouterWithdraw, error) {
	event := new(RouterWithdraw)
	if err := _Router.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
