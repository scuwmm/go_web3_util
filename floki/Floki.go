// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package floki

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

// FlokiMetaData contains all meta data concerning the Floki contract.
var FlokiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"taxHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasuryHandlerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentDelegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint224\",\"name\":\"oldVotes\",\"type\":\"uint224\"},{\"indexed\":false,\"internalType\":\"uint224\",\"name\":\"newVotes\",\"type\":\"uint224\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"TaxHandlerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"TreasuryHandlerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint224\",\"name\":\"votes\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getVotesAtBlock\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taxHandlerAddress\",\"type\":\"address\"}],\"name\":\"setTaxHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasuryHandlerAddress\",\"type\":\"address\"}],\"name\":\"setTreasuryHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxHandler\",\"outputs\":[{\"internalType\":\"contractITaxHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryHandler\",\"outputs\":[{\"internalType\":\"contractITreasuryHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FlokiABI is the input ABI used to generate the binding from.
// Deprecated: Use FlokiMetaData.ABI instead.
var FlokiABI = FlokiMetaData.ABI

// Floki is an auto generated Go binding around an Ethereum contract.
type Floki struct {
	FlokiCaller     // Read-only binding to the contract
	FlokiTransactor // Write-only binding to the contract
	FlokiFilterer   // Log filterer for contract events
}

// FlokiCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlokiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlokiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlokiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlokiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlokiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlokiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlokiSession struct {
	Contract     *Floki            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlokiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlokiCallerSession struct {
	Contract *FlokiCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FlokiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlokiTransactorSession struct {
	Contract     *FlokiTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlokiRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlokiRaw struct {
	Contract *Floki // Generic contract binding to access the raw methods on
}

// FlokiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlokiCallerRaw struct {
	Contract *FlokiCaller // Generic read-only contract binding to access the raw methods on
}

// FlokiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlokiTransactorRaw struct {
	Contract *FlokiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFloki creates a new instance of Floki, bound to a specific deployed contract.
func NewFloki(address common.Address, backend bind.ContractBackend) (*Floki, error) {
	contract, err := bindFloki(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Floki{FlokiCaller: FlokiCaller{contract: contract}, FlokiTransactor: FlokiTransactor{contract: contract}, FlokiFilterer: FlokiFilterer{contract: contract}}, nil
}

// NewFlokiCaller creates a new read-only instance of Floki, bound to a specific deployed contract.
func NewFlokiCaller(address common.Address, caller bind.ContractCaller) (*FlokiCaller, error) {
	contract, err := bindFloki(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlokiCaller{contract: contract}, nil
}

// NewFlokiTransactor creates a new write-only instance of Floki, bound to a specific deployed contract.
func NewFlokiTransactor(address common.Address, transactor bind.ContractTransactor) (*FlokiTransactor, error) {
	contract, err := bindFloki(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlokiTransactor{contract: contract}, nil
}

// NewFlokiFilterer creates a new log filterer instance of Floki, bound to a specific deployed contract.
func NewFlokiFilterer(address common.Address, filterer bind.ContractFilterer) (*FlokiFilterer, error) {
	contract, err := bindFloki(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlokiFilterer{contract: contract}, nil
}

// bindFloki binds a generic wrapper to an already deployed contract.
func bindFloki(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlokiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Floki *FlokiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Floki.Contract.FlokiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Floki *FlokiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Floki.Contract.FlokiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Floki *FlokiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Floki.Contract.FlokiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Floki *FlokiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Floki.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Floki *FlokiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Floki.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Floki *FlokiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Floki.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Floki *FlokiCaller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Floki *FlokiSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Floki.Contract.DELEGATIONTYPEHASH(&_Floki.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Floki *FlokiCallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Floki.Contract.DELEGATIONTYPEHASH(&_Floki.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Floki *FlokiCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Floki *FlokiSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Floki.Contract.DOMAINTYPEHASH(&_Floki.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Floki *FlokiCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Floki.Contract.DOMAINTYPEHASH(&_Floki.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Floki *FlokiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Floki *FlokiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Floki.Contract.Allowance(&_Floki.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Floki *FlokiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Floki.Contract.Allowance(&_Floki.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Floki *FlokiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Floki *FlokiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Floki.Contract.BalanceOf(&_Floki.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Floki *FlokiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Floki.Contract.BalanceOf(&_Floki.CallOpts, account)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 blockNumber, uint224 votes)
func (_Floki *FlokiCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	BlockNumber uint32
	Votes       *big.Int
}, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		BlockNumber uint32
		Votes       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 blockNumber, uint224 votes)
func (_Floki *FlokiSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	BlockNumber uint32
	Votes       *big.Int
}, error) {
	return _Floki.Contract.Checkpoints(&_Floki.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 blockNumber, uint224 votes)
func (_Floki *FlokiCallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	BlockNumber uint32
	Votes       *big.Int
}, error) {
	return _Floki.Contract.Checkpoints(&_Floki.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Floki *FlokiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Floki *FlokiSession) Decimals() (uint8, error) {
	return _Floki.Contract.Decimals(&_Floki.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Floki *FlokiCallerSession) Decimals() (uint8, error) {
	return _Floki.Contract.Decimals(&_Floki.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_Floki *FlokiCaller) Delegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "delegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_Floki *FlokiSession) Delegates(arg0 common.Address) (common.Address, error) {
	return _Floki.Contract.Delegates(&_Floki.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_Floki *FlokiCallerSession) Delegates(arg0 common.Address) (common.Address, error) {
	return _Floki.Contract.Delegates(&_Floki.CallOpts, arg0)
}

// GetVotesAtBlock is a free data retrieval call binding the contract method 0x271a4529.
//
// Solidity: function getVotesAtBlock(address account, uint32 blockNumber) view returns(uint224)
func (_Floki *FlokiCaller) GetVotesAtBlock(opts *bind.CallOpts, account common.Address, blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "getVotesAtBlock", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesAtBlock is a free data retrieval call binding the contract method 0x271a4529.
//
// Solidity: function getVotesAtBlock(address account, uint32 blockNumber) view returns(uint224)
func (_Floki *FlokiSession) GetVotesAtBlock(account common.Address, blockNumber uint32) (*big.Int, error) {
	return _Floki.Contract.GetVotesAtBlock(&_Floki.CallOpts, account, blockNumber)
}

// GetVotesAtBlock is a free data retrieval call binding the contract method 0x271a4529.
//
// Solidity: function getVotesAtBlock(address account, uint32 blockNumber) view returns(uint224)
func (_Floki *FlokiCallerSession) GetVotesAtBlock(account common.Address, blockNumber uint32) (*big.Int, error) {
	return _Floki.Contract.GetVotesAtBlock(&_Floki.CallOpts, account, blockNumber)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Floki *FlokiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Floki *FlokiSession) Name() (string, error) {
	return _Floki.Contract.Name(&_Floki.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Floki *FlokiCallerSession) Name() (string, error) {
	return _Floki.Contract.Name(&_Floki.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Floki *FlokiCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Floki *FlokiSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Floki.Contract.Nonces(&_Floki.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Floki *FlokiCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Floki.Contract.Nonces(&_Floki.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Floki *FlokiCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Floki *FlokiSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Floki.Contract.NumCheckpoints(&_Floki.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Floki *FlokiCallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Floki.Contract.NumCheckpoints(&_Floki.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Floki *FlokiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Floki *FlokiSession) Owner() (common.Address, error) {
	return _Floki.Contract.Owner(&_Floki.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Floki *FlokiCallerSession) Owner() (common.Address, error) {
	return _Floki.Contract.Owner(&_Floki.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Floki *FlokiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Floki *FlokiSession) Symbol() (string, error) {
	return _Floki.Contract.Symbol(&_Floki.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Floki *FlokiCallerSession) Symbol() (string, error) {
	return _Floki.Contract.Symbol(&_Floki.CallOpts)
}

// TaxHandler is a free data retrieval call binding the contract method 0x12280ba8.
//
// Solidity: function taxHandler() view returns(address)
func (_Floki *FlokiCaller) TaxHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "taxHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaxHandler is a free data retrieval call binding the contract method 0x12280ba8.
//
// Solidity: function taxHandler() view returns(address)
func (_Floki *FlokiSession) TaxHandler() (common.Address, error) {
	return _Floki.Contract.TaxHandler(&_Floki.CallOpts)
}

// TaxHandler is a free data retrieval call binding the contract method 0x12280ba8.
//
// Solidity: function taxHandler() view returns(address)
func (_Floki *FlokiCallerSession) TaxHandler() (common.Address, error) {
	return _Floki.Contract.TaxHandler(&_Floki.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_Floki *FlokiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_Floki *FlokiSession) TotalSupply() (*big.Int, error) {
	return _Floki.Contract.TotalSupply(&_Floki.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_Floki *FlokiCallerSession) TotalSupply() (*big.Int, error) {
	return _Floki.Contract.TotalSupply(&_Floki.CallOpts)
}

// TreasuryHandler is a free data retrieval call binding the contract method 0x17889633.
//
// Solidity: function treasuryHandler() view returns(address)
func (_Floki *FlokiCaller) TreasuryHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Floki.contract.Call(opts, &out, "treasuryHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryHandler is a free data retrieval call binding the contract method 0x17889633.
//
// Solidity: function treasuryHandler() view returns(address)
func (_Floki *FlokiSession) TreasuryHandler() (common.Address, error) {
	return _Floki.Contract.TreasuryHandler(&_Floki.CallOpts)
}

// TreasuryHandler is a free data retrieval call binding the contract method 0x17889633.
//
// Solidity: function treasuryHandler() view returns(address)
func (_Floki *FlokiCallerSession) TreasuryHandler() (common.Address, error) {
	return _Floki.Contract.TreasuryHandler(&_Floki.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Floki *FlokiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Floki *FlokiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.Approve(&_Floki.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Floki *FlokiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.Approve(&_Floki.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Floki *FlokiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Floki *FlokiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.DecreaseAllowance(&_Floki.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Floki *FlokiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.DecreaseAllowance(&_Floki.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Floki *FlokiTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Floki *FlokiSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Floki.Contract.Delegate(&_Floki.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Floki *FlokiTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Floki.Contract.Delegate(&_Floki.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Floki *FlokiTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Floki *FlokiSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Floki.Contract.DelegateBySig(&_Floki.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Floki *FlokiTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Floki.Contract.DelegateBySig(&_Floki.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Floki *FlokiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Floki *FlokiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.IncreaseAllowance(&_Floki.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Floki *FlokiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.IncreaseAllowance(&_Floki.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Floki *FlokiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Floki *FlokiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Floki.Contract.RenounceOwnership(&_Floki.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Floki *FlokiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Floki.Contract.RenounceOwnership(&_Floki.TransactOpts)
}

// SetTaxHandler is a paid mutator transaction binding the contract method 0x488d4a51.
//
// Solidity: function setTaxHandler(address taxHandlerAddress) returns()
func (_Floki *FlokiTransactor) SetTaxHandler(opts *bind.TransactOpts, taxHandlerAddress common.Address) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "setTaxHandler", taxHandlerAddress)
}

// SetTaxHandler is a paid mutator transaction binding the contract method 0x488d4a51.
//
// Solidity: function setTaxHandler(address taxHandlerAddress) returns()
func (_Floki *FlokiSession) SetTaxHandler(taxHandlerAddress common.Address) (*types.Transaction, error) {
	return _Floki.Contract.SetTaxHandler(&_Floki.TransactOpts, taxHandlerAddress)
}

// SetTaxHandler is a paid mutator transaction binding the contract method 0x488d4a51.
//
// Solidity: function setTaxHandler(address taxHandlerAddress) returns()
func (_Floki *FlokiTransactorSession) SetTaxHandler(taxHandlerAddress common.Address) (*types.Transaction, error) {
	return _Floki.Contract.SetTaxHandler(&_Floki.TransactOpts, taxHandlerAddress)
}

// SetTreasuryHandler is a paid mutator transaction binding the contract method 0xa9373b7b.
//
// Solidity: function setTreasuryHandler(address treasuryHandlerAddress) returns()
func (_Floki *FlokiTransactor) SetTreasuryHandler(opts *bind.TransactOpts, treasuryHandlerAddress common.Address) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "setTreasuryHandler", treasuryHandlerAddress)
}

// SetTreasuryHandler is a paid mutator transaction binding the contract method 0xa9373b7b.
//
// Solidity: function setTreasuryHandler(address treasuryHandlerAddress) returns()
func (_Floki *FlokiSession) SetTreasuryHandler(treasuryHandlerAddress common.Address) (*types.Transaction, error) {
	return _Floki.Contract.SetTreasuryHandler(&_Floki.TransactOpts, treasuryHandlerAddress)
}

// SetTreasuryHandler is a paid mutator transaction binding the contract method 0xa9373b7b.
//
// Solidity: function setTreasuryHandler(address treasuryHandlerAddress) returns()
func (_Floki *FlokiTransactorSession) SetTreasuryHandler(treasuryHandlerAddress common.Address) (*types.Transaction, error) {
	return _Floki.Contract.SetTreasuryHandler(&_Floki.TransactOpts, treasuryHandlerAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Floki *FlokiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Floki *FlokiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.Transfer(&_Floki.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Floki *FlokiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.Transfer(&_Floki.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Floki *FlokiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Floki *FlokiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.TransferFrom(&_Floki.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Floki *FlokiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Floki.Contract.TransferFrom(&_Floki.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Floki *FlokiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Floki.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Floki *FlokiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Floki.Contract.TransferOwnership(&_Floki.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Floki *FlokiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Floki.Contract.TransferOwnership(&_Floki.TransactOpts, newOwner)
}

// FlokiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Floki contract.
type FlokiApprovalIterator struct {
	Event *FlokiApproval // Event containing the contract specifics and raw log

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
func (it *FlokiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlokiApproval)
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
		it.Event = new(FlokiApproval)
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
func (it *FlokiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlokiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlokiApproval represents a Approval event raised by the Floki contract.
type FlokiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Floki *FlokiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FlokiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Floki.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FlokiApprovalIterator{contract: _Floki.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Floki *FlokiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FlokiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Floki.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlokiApproval)
				if err := _Floki.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Floki *FlokiFilterer) ParseApproval(log types.Log) (*FlokiApproval, error) {
	event := new(FlokiApproval)
	if err := _Floki.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlokiDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Floki contract.
type FlokiDelegateChangedIterator struct {
	Event *FlokiDelegateChanged // Event containing the contract specifics and raw log

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
func (it *FlokiDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlokiDelegateChanged)
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
		it.Event = new(FlokiDelegateChanged)
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
func (it *FlokiDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlokiDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlokiDelegateChanged represents a DelegateChanged event raised by the Floki contract.
type FlokiDelegateChanged struct {
	Delegator       common.Address
	CurrentDelegate common.Address
	NewDelegate     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address delegator, address currentDelegate, address newDelegate)
func (_Floki *FlokiFilterer) FilterDelegateChanged(opts *bind.FilterOpts) (*FlokiDelegateChangedIterator, error) {

	logs, sub, err := _Floki.contract.FilterLogs(opts, "DelegateChanged")
	if err != nil {
		return nil, err
	}
	return &FlokiDelegateChangedIterator{contract: _Floki.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address delegator, address currentDelegate, address newDelegate)
func (_Floki *FlokiFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *FlokiDelegateChanged) (event.Subscription, error) {

	logs, sub, err := _Floki.contract.WatchLogs(opts, "DelegateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlokiDelegateChanged)
				if err := _Floki.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address delegator, address currentDelegate, address newDelegate)
func (_Floki *FlokiFilterer) ParseDelegateChanged(log types.Log) (*FlokiDelegateChanged, error) {
	event := new(FlokiDelegateChanged)
	if err := _Floki.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlokiDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the Floki contract.
type FlokiDelegateVotesChangedIterator struct {
	Event *FlokiDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *FlokiDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlokiDelegateVotesChanged)
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
		it.Event = new(FlokiDelegateVotesChanged)
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
func (it *FlokiDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlokiDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlokiDelegateVotesChanged represents a DelegateVotesChanged event raised by the Floki contract.
type FlokiDelegateVotesChanged struct {
	Delegatee common.Address
	OldVotes  *big.Int
	NewVotes  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xda5a64c2947c0b7bf4d6e7bf736c6f84d9d1c5f991770f88bbeb3fe19c85a134.
//
// Solidity: event DelegateVotesChanged(address delegatee, uint224 oldVotes, uint224 newVotes)
func (_Floki *FlokiFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts) (*FlokiDelegateVotesChangedIterator, error) {

	logs, sub, err := _Floki.contract.FilterLogs(opts, "DelegateVotesChanged")
	if err != nil {
		return nil, err
	}
	return &FlokiDelegateVotesChangedIterator{contract: _Floki.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xda5a64c2947c0b7bf4d6e7bf736c6f84d9d1c5f991770f88bbeb3fe19c85a134.
//
// Solidity: event DelegateVotesChanged(address delegatee, uint224 oldVotes, uint224 newVotes)
func (_Floki *FlokiFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *FlokiDelegateVotesChanged) (event.Subscription, error) {

	logs, sub, err := _Floki.contract.WatchLogs(opts, "DelegateVotesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlokiDelegateVotesChanged)
				if err := _Floki.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xda5a64c2947c0b7bf4d6e7bf736c6f84d9d1c5f991770f88bbeb3fe19c85a134.
//
// Solidity: event DelegateVotesChanged(address delegatee, uint224 oldVotes, uint224 newVotes)
func (_Floki *FlokiFilterer) ParseDelegateVotesChanged(log types.Log) (*FlokiDelegateVotesChanged, error) {
	event := new(FlokiDelegateVotesChanged)
	if err := _Floki.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlokiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Floki contract.
type FlokiOwnershipTransferredIterator struct {
	Event *FlokiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FlokiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlokiOwnershipTransferred)
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
		it.Event = new(FlokiOwnershipTransferred)
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
func (it *FlokiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlokiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlokiOwnershipTransferred represents a OwnershipTransferred event raised by the Floki contract.
type FlokiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Floki *FlokiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FlokiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Floki.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FlokiOwnershipTransferredIterator{contract: _Floki.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Floki *FlokiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FlokiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Floki.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlokiOwnershipTransferred)
				if err := _Floki.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Floki *FlokiFilterer) ParseOwnershipTransferred(log types.Log) (*FlokiOwnershipTransferred, error) {
	event := new(FlokiOwnershipTransferred)
	if err := _Floki.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlokiTaxHandlerChangedIterator is returned from FilterTaxHandlerChanged and is used to iterate over the raw logs and unpacked data for TaxHandlerChanged events raised by the Floki contract.
type FlokiTaxHandlerChangedIterator struct {
	Event *FlokiTaxHandlerChanged // Event containing the contract specifics and raw log

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
func (it *FlokiTaxHandlerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlokiTaxHandlerChanged)
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
		it.Event = new(FlokiTaxHandlerChanged)
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
func (it *FlokiTaxHandlerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlokiTaxHandlerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlokiTaxHandlerChanged represents a TaxHandlerChanged event raised by the Floki contract.
type FlokiTaxHandlerChanged struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaxHandlerChanged is a free log retrieval operation binding the contract event 0x00d910c9481701ba32afe0c247572aaece27072f230c8ec769bf245fc0b38de6.
//
// Solidity: event TaxHandlerChanged(address oldAddress, address newAddress)
func (_Floki *FlokiFilterer) FilterTaxHandlerChanged(opts *bind.FilterOpts) (*FlokiTaxHandlerChangedIterator, error) {

	logs, sub, err := _Floki.contract.FilterLogs(opts, "TaxHandlerChanged")
	if err != nil {
		return nil, err
	}
	return &FlokiTaxHandlerChangedIterator{contract: _Floki.contract, event: "TaxHandlerChanged", logs: logs, sub: sub}, nil
}

// WatchTaxHandlerChanged is a free log subscription operation binding the contract event 0x00d910c9481701ba32afe0c247572aaece27072f230c8ec769bf245fc0b38de6.
//
// Solidity: event TaxHandlerChanged(address oldAddress, address newAddress)
func (_Floki *FlokiFilterer) WatchTaxHandlerChanged(opts *bind.WatchOpts, sink chan<- *FlokiTaxHandlerChanged) (event.Subscription, error) {

	logs, sub, err := _Floki.contract.WatchLogs(opts, "TaxHandlerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlokiTaxHandlerChanged)
				if err := _Floki.contract.UnpackLog(event, "TaxHandlerChanged", log); err != nil {
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

// ParseTaxHandlerChanged is a log parse operation binding the contract event 0x00d910c9481701ba32afe0c247572aaece27072f230c8ec769bf245fc0b38de6.
//
// Solidity: event TaxHandlerChanged(address oldAddress, address newAddress)
func (_Floki *FlokiFilterer) ParseTaxHandlerChanged(log types.Log) (*FlokiTaxHandlerChanged, error) {
	event := new(FlokiTaxHandlerChanged)
	if err := _Floki.contract.UnpackLog(event, "TaxHandlerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlokiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Floki contract.
type FlokiTransferIterator struct {
	Event *FlokiTransfer // Event containing the contract specifics and raw log

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
func (it *FlokiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlokiTransfer)
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
		it.Event = new(FlokiTransfer)
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
func (it *FlokiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlokiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlokiTransfer represents a Transfer event raised by the Floki contract.
type FlokiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Floki *FlokiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FlokiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Floki.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FlokiTransferIterator{contract: _Floki.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Floki *FlokiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FlokiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Floki.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlokiTransfer)
				if err := _Floki.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Floki *FlokiFilterer) ParseTransfer(log types.Log) (*FlokiTransfer, error) {
	event := new(FlokiTransfer)
	if err := _Floki.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlokiTreasuryHandlerChangedIterator is returned from FilterTreasuryHandlerChanged and is used to iterate over the raw logs and unpacked data for TreasuryHandlerChanged events raised by the Floki contract.
type FlokiTreasuryHandlerChangedIterator struct {
	Event *FlokiTreasuryHandlerChanged // Event containing the contract specifics and raw log

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
func (it *FlokiTreasuryHandlerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlokiTreasuryHandlerChanged)
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
		it.Event = new(FlokiTreasuryHandlerChanged)
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
func (it *FlokiTreasuryHandlerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlokiTreasuryHandlerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlokiTreasuryHandlerChanged represents a TreasuryHandlerChanged event raised by the Floki contract.
type FlokiTreasuryHandlerChanged struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTreasuryHandlerChanged is a free log retrieval operation binding the contract event 0x1bf87992a35ee29395ab494f9adb9a500a7fa60c3082cba0ef02701bb35900d9.
//
// Solidity: event TreasuryHandlerChanged(address oldAddress, address newAddress)
func (_Floki *FlokiFilterer) FilterTreasuryHandlerChanged(opts *bind.FilterOpts) (*FlokiTreasuryHandlerChangedIterator, error) {

	logs, sub, err := _Floki.contract.FilterLogs(opts, "TreasuryHandlerChanged")
	if err != nil {
		return nil, err
	}
	return &FlokiTreasuryHandlerChangedIterator{contract: _Floki.contract, event: "TreasuryHandlerChanged", logs: logs, sub: sub}, nil
}

// WatchTreasuryHandlerChanged is a free log subscription operation binding the contract event 0x1bf87992a35ee29395ab494f9adb9a500a7fa60c3082cba0ef02701bb35900d9.
//
// Solidity: event TreasuryHandlerChanged(address oldAddress, address newAddress)
func (_Floki *FlokiFilterer) WatchTreasuryHandlerChanged(opts *bind.WatchOpts, sink chan<- *FlokiTreasuryHandlerChanged) (event.Subscription, error) {

	logs, sub, err := _Floki.contract.WatchLogs(opts, "TreasuryHandlerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlokiTreasuryHandlerChanged)
				if err := _Floki.contract.UnpackLog(event, "TreasuryHandlerChanged", log); err != nil {
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

// ParseTreasuryHandlerChanged is a log parse operation binding the contract event 0x1bf87992a35ee29395ab494f9adb9a500a7fa60c3082cba0ef02701bb35900d9.
//
// Solidity: event TreasuryHandlerChanged(address oldAddress, address newAddress)
func (_Floki *FlokiFilterer) ParseTreasuryHandlerChanged(log types.Log) (*FlokiTreasuryHandlerChanged, error) {
	event := new(FlokiTreasuryHandlerChanged)
	if err := _Floki.contract.UnpackLog(event, "TreasuryHandlerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
