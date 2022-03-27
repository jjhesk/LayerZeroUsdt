// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package layer_zero_dollar

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

// LayerZeroDollarMetaData contains all meta data concerning the LayerZeroDollar contract.
var LayerZeroDollarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_layerZeroEndpoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evil\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evil\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rawAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"remotes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"noevil\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_dstOmniChainTokenAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_qty\",\"type\":\"uint256\"}],\"name\":\"sendTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"setRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LayerZeroDollarABI is the input ABI used to generate the binding from.
// Deprecated: Use LayerZeroDollarMetaData.ABI instead.
var LayerZeroDollarABI = LayerZeroDollarMetaData.ABI

// LayerZeroDollar is an auto generated Go binding around an Ethereum contract.
type LayerZeroDollar struct {
	LayerZeroDollarCaller     // Read-only binding to the contract
	LayerZeroDollarTransactor // Write-only binding to the contract
	LayerZeroDollarFilterer   // Log filterer for contract events
}

// LayerZeroDollarCaller is an auto generated read-only Go binding around an Ethereum contract.
type LayerZeroDollarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LayerZeroDollarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LayerZeroDollarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LayerZeroDollarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LayerZeroDollarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LayerZeroDollarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LayerZeroDollarSession struct {
	Contract     *LayerZeroDollar  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LayerZeroDollarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LayerZeroDollarCallerSession struct {
	Contract *LayerZeroDollarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LayerZeroDollarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LayerZeroDollarTransactorSession struct {
	Contract     *LayerZeroDollarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LayerZeroDollarRaw is an auto generated low-level Go binding around an Ethereum contract.
type LayerZeroDollarRaw struct {
	Contract *LayerZeroDollar // Generic contract binding to access the raw methods on
}

// LayerZeroDollarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LayerZeroDollarCallerRaw struct {
	Contract *LayerZeroDollarCaller // Generic read-only contract binding to access the raw methods on
}

// LayerZeroDollarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LayerZeroDollarTransactorRaw struct {
	Contract *LayerZeroDollarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLayerZeroDollar creates a new instance of LayerZeroDollar, bound to a specific deployed contract.
func NewLayerZeroDollar(address common.Address, backend bind.ContractBackend) (*LayerZeroDollar, error) {
	contract, err := bindLayerZeroDollar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LayerZeroDollar{LayerZeroDollarCaller: LayerZeroDollarCaller{contract: contract}, LayerZeroDollarTransactor: LayerZeroDollarTransactor{contract: contract}, LayerZeroDollarFilterer: LayerZeroDollarFilterer{contract: contract}}, nil
}

// NewLayerZeroDollarCaller creates a new read-only instance of LayerZeroDollar, bound to a specific deployed contract.
func NewLayerZeroDollarCaller(address common.Address, caller bind.ContractCaller) (*LayerZeroDollarCaller, error) {
	contract, err := bindLayerZeroDollar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LayerZeroDollarCaller{contract: contract}, nil
}

// NewLayerZeroDollarTransactor creates a new write-only instance of LayerZeroDollar, bound to a specific deployed contract.
func NewLayerZeroDollarTransactor(address common.Address, transactor bind.ContractTransactor) (*LayerZeroDollarTransactor, error) {
	contract, err := bindLayerZeroDollar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LayerZeroDollarTransactor{contract: contract}, nil
}

// NewLayerZeroDollarFilterer creates a new log filterer instance of LayerZeroDollar, bound to a specific deployed contract.
func NewLayerZeroDollarFilterer(address common.Address, filterer bind.ContractFilterer) (*LayerZeroDollarFilterer, error) {
	contract, err := bindLayerZeroDollar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LayerZeroDollarFilterer{contract: contract}, nil
}

// bindLayerZeroDollar binds a generic wrapper to an already deployed contract.
func bindLayerZeroDollar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LayerZeroDollarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LayerZeroDollar *LayerZeroDollarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LayerZeroDollar.Contract.LayerZeroDollarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LayerZeroDollar *LayerZeroDollarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.LayerZeroDollarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LayerZeroDollar *LayerZeroDollarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.LayerZeroDollarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LayerZeroDollar *LayerZeroDollarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LayerZeroDollar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LayerZeroDollar *LayerZeroDollarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LayerZeroDollar *LayerZeroDollarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.contract.Transact(opts, method, params...)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_LayerZeroDollar *LayerZeroDollarCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_LayerZeroDollar *LayerZeroDollarSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _LayerZeroDollar.Contract.DOMAINTYPEHASH(&_LayerZeroDollar.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _LayerZeroDollar.Contract.DOMAINTYPEHASH(&_LayerZeroDollar.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_LayerZeroDollar *LayerZeroDollarCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_LayerZeroDollar *LayerZeroDollarSession) PERMITTYPEHASH() ([32]byte, error) {
	return _LayerZeroDollar.Contract.PERMITTYPEHASH(&_LayerZeroDollar.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _LayerZeroDollar.Contract.PERMITTYPEHASH(&_LayerZeroDollar.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LayerZeroDollar.Contract.Allowance(&_LayerZeroDollar.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LayerZeroDollar.Contract.Allowance(&_LayerZeroDollar.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LayerZeroDollar.Contract.BalanceOf(&_LayerZeroDollar.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LayerZeroDollar.Contract.BalanceOf(&_LayerZeroDollar.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LayerZeroDollar *LayerZeroDollarCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LayerZeroDollar *LayerZeroDollarSession) Decimals() (uint8, error) {
	return _LayerZeroDollar.Contract.Decimals(&_LayerZeroDollar.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) Decimals() (uint8, error) {
	return _LayerZeroDollar.Contract.Decimals(&_LayerZeroDollar.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_LayerZeroDollar *LayerZeroDollarCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_LayerZeroDollar *LayerZeroDollarSession) Endpoint() (common.Address, error) {
	return _LayerZeroDollar.Contract.Endpoint(&_LayerZeroDollar.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) Endpoint() (common.Address, error) {
	return _LayerZeroDollar.Contract.Endpoint(&_LayerZeroDollar.CallOpts)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_LayerZeroDollar *LayerZeroDollarCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_LayerZeroDollar *LayerZeroDollarSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _LayerZeroDollar.Contract.GetBlackListStatus(&_LayerZeroDollar.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _LayerZeroDollar.Contract.GetBlackListStatus(&_LayerZeroDollar.CallOpts, _maker)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_LayerZeroDollar *LayerZeroDollarCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_LayerZeroDollar *LayerZeroDollarSession) Governance() (common.Address, error) {
	return _LayerZeroDollar.Contract.Governance(&_LayerZeroDollar.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) Governance() (common.Address, error) {
	return _LayerZeroDollar.Contract.Governance(&_LayerZeroDollar.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_LayerZeroDollar *LayerZeroDollarCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_LayerZeroDollar *LayerZeroDollarSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _LayerZeroDollar.Contract.IsBlackListed(&_LayerZeroDollar.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _LayerZeroDollar.Contract.IsBlackListed(&_LayerZeroDollar.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_LayerZeroDollar *LayerZeroDollarCaller) Minters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_LayerZeroDollar *LayerZeroDollarSession) Minters(arg0 common.Address) (bool, error) {
	return _LayerZeroDollar.Contract.Minters(&_LayerZeroDollar.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) Minters(arg0 common.Address) (bool, error) {
	return _LayerZeroDollar.Contract.Minters(&_LayerZeroDollar.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LayerZeroDollar *LayerZeroDollarCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LayerZeroDollar *LayerZeroDollarSession) Name() (string, error) {
	return _LayerZeroDollar.Contract.Name(&_LayerZeroDollar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) Name() (string, error) {
	return _LayerZeroDollar.Contract.Name(&_LayerZeroDollar.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _LayerZeroDollar.Contract.Nonces(&_LayerZeroDollar.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _LayerZeroDollar.Contract.Nonces(&_LayerZeroDollar.CallOpts, arg0)
}

// Remotes is a free data retrieval call binding the contract method 0x9d1f6113.
//
// Solidity: function remotes(uint16 ) view returns(bytes)
func (_LayerZeroDollar *LayerZeroDollarCaller) Remotes(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "remotes", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Remotes is a free data retrieval call binding the contract method 0x9d1f6113.
//
// Solidity: function remotes(uint16 ) view returns(bytes)
func (_LayerZeroDollar *LayerZeroDollarSession) Remotes(arg0 uint16) ([]byte, error) {
	return _LayerZeroDollar.Contract.Remotes(&_LayerZeroDollar.CallOpts, arg0)
}

// Remotes is a free data retrieval call binding the contract method 0x9d1f6113.
//
// Solidity: function remotes(uint16 ) view returns(bytes)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) Remotes(arg0 uint16) ([]byte, error) {
	return _LayerZeroDollar.Contract.Remotes(&_LayerZeroDollar.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LayerZeroDollar *LayerZeroDollarCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LayerZeroDollar *LayerZeroDollarSession) Symbol() (string, error) {
	return _LayerZeroDollar.Contract.Symbol(&_LayerZeroDollar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) Symbol() (string, error) {
	return _LayerZeroDollar.Contract.Symbol(&_LayerZeroDollar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LayerZeroDollar.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarSession) TotalSupply() (*big.Int, error) {
	return _LayerZeroDollar.Contract.TotalSupply(&_LayerZeroDollar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LayerZeroDollar *LayerZeroDollarCallerSession) TotalSupply() (*big.Int, error) {
	return _LayerZeroDollar.Contract.TotalSupply(&_LayerZeroDollar.CallOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address evil) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactor) AddBlackList(opts *bind.TransactOpts, evil common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "addBlackList", evil)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address evil) returns()
func (_LayerZeroDollar *LayerZeroDollarSession) AddBlackList(evil common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.AddBlackList(&_LayerZeroDollar.TransactOpts, evil)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address evil) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) AddBlackList(evil common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.AddBlackList(&_LayerZeroDollar.TransactOpts, evil)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.Approve(&_LayerZeroDollar.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.Approve(&_LayerZeroDollar.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.DecreaseAllowance(&_LayerZeroDollar.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.DecreaseAllowance(&_LayerZeroDollar.TransactOpts, spender, subtractedValue)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address evil) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactor) DestroyBlackFunds(opts *bind.TransactOpts, evil common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "destroyBlackFunds", evil)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address evil) returns()
func (_LayerZeroDollar *LayerZeroDollarSession) DestroyBlackFunds(evil common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.DestroyBlackFunds(&_LayerZeroDollar.TransactOpts, evil)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address evil) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) DestroyBlackFunds(evil common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.DestroyBlackFunds(&_LayerZeroDollar.TransactOpts, evil)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.IncreaseAllowance(&_LayerZeroDollar.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.IncreaseAllowance(&_LayerZeroDollar.TransactOpts, spender, addedValue)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 , bytes _payload) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, arg2 uint64, _payload []byte) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, arg2, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 , bytes _payload) returns()
func (_LayerZeroDollar *LayerZeroDollarSession) LzReceive(_srcChainId uint16, _srcAddress []byte, arg2 uint64, _payload []byte) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.LzReceive(&_LayerZeroDollar.TransactOpts, _srcChainId, _srcAddress, arg2, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 , bytes _payload) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, arg2 uint64, _payload []byte) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.LzReceive(&_LayerZeroDollar.TransactOpts, _srcChainId, _srcAddress, arg2, _payload)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 rawAmount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, rawAmount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "permit", owner, spender, rawAmount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 rawAmount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LayerZeroDollar *LayerZeroDollarSession) Permit(owner common.Address, spender common.Address, rawAmount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.Permit(&_LayerZeroDollar.TransactOpts, owner, spender, rawAmount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 rawAmount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) Permit(owner common.Address, spender common.Address, rawAmount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.Permit(&_LayerZeroDollar.TransactOpts, owner, spender, rawAmount, deadline, v, r, s)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address noevil) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactor) RemoveBlackList(opts *bind.TransactOpts, noevil common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "removeBlackList", noevil)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address noevil) returns()
func (_LayerZeroDollar *LayerZeroDollarSession) RemoveBlackList(noevil common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.RemoveBlackList(&_LayerZeroDollar.TransactOpts, noevil)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address noevil) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) RemoveBlackList(noevil common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.RemoveBlackList(&_LayerZeroDollar.TransactOpts, noevil)
}

// SendTokens is a paid mutator transaction binding the contract method 0x36da46c9.
//
// Solidity: function sendTokens(uint16 _chainId, bytes _dstOmniChainTokenAddr, uint256 _qty) payable returns()
func (_LayerZeroDollar *LayerZeroDollarTransactor) SendTokens(opts *bind.TransactOpts, _chainId uint16, _dstOmniChainTokenAddr []byte, _qty *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "sendTokens", _chainId, _dstOmniChainTokenAddr, _qty)
}

// SendTokens is a paid mutator transaction binding the contract method 0x36da46c9.
//
// Solidity: function sendTokens(uint16 _chainId, bytes _dstOmniChainTokenAddr, uint256 _qty) payable returns()
func (_LayerZeroDollar *LayerZeroDollarSession) SendTokens(_chainId uint16, _dstOmniChainTokenAddr []byte, _qty *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.SendTokens(&_LayerZeroDollar.TransactOpts, _chainId, _dstOmniChainTokenAddr, _qty)
}

// SendTokens is a paid mutator transaction binding the contract method 0x36da46c9.
//
// Solidity: function sendTokens(uint16 _chainId, bytes _dstOmniChainTokenAddr, uint256 _qty) payable returns()
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) SendTokens(_chainId uint16, _dstOmniChainTokenAddr []byte, _qty *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.SendTokens(&_LayerZeroDollar.TransactOpts, _chainId, _dstOmniChainTokenAddr, _qty)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_LayerZeroDollar *LayerZeroDollarSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.SetGovernance(&_LayerZeroDollar.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.SetGovernance(&_LayerZeroDollar.TransactOpts, _governance)
}

// SetRemote is a paid mutator transaction binding the contract method 0x20cdd0a1.
//
// Solidity: function setRemote(uint16 _chainId, bytes _remoteAddress) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactor) SetRemote(opts *bind.TransactOpts, _chainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "setRemote", _chainId, _remoteAddress)
}

// SetRemote is a paid mutator transaction binding the contract method 0x20cdd0a1.
//
// Solidity: function setRemote(uint16 _chainId, bytes _remoteAddress) returns()
func (_LayerZeroDollar *LayerZeroDollarSession) SetRemote(_chainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.SetRemote(&_LayerZeroDollar.TransactOpts, _chainId, _remoteAddress)
}

// SetRemote is a paid mutator transaction binding the contract method 0x20cdd0a1.
//
// Solidity: function setRemote(uint16 _chainId, bytes _remoteAddress) returns()
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) SetRemote(_chainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.SetRemote(&_LayerZeroDollar.TransactOpts, _chainId, _remoteAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.Transfer(&_LayerZeroDollar.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.Transfer(&_LayerZeroDollar.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.TransferFrom(&_LayerZeroDollar.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LayerZeroDollar *LayerZeroDollarTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LayerZeroDollar.Contract.TransferFrom(&_LayerZeroDollar.TransactOpts, sender, recipient, amount)
}

// LayerZeroDollarAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the LayerZeroDollar contract.
type LayerZeroDollarAddedBlackListIterator struct {
	Event *LayerZeroDollarAddedBlackList // Event containing the contract specifics and raw log

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
func (it *LayerZeroDollarAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerZeroDollarAddedBlackList)
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
		it.Event = new(LayerZeroDollarAddedBlackList)
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
func (it *LayerZeroDollarAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerZeroDollarAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerZeroDollarAddedBlackList represents a AddedBlackList event raised by the LayerZeroDollar contract.
type LayerZeroDollarAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_LayerZeroDollar *LayerZeroDollarFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*LayerZeroDollarAddedBlackListIterator, error) {

	logs, sub, err := _LayerZeroDollar.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &LayerZeroDollarAddedBlackListIterator{contract: _LayerZeroDollar.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_LayerZeroDollar *LayerZeroDollarFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *LayerZeroDollarAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _LayerZeroDollar.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerZeroDollarAddedBlackList)
				if err := _LayerZeroDollar.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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

// ParseAddedBlackList is a log parse operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_LayerZeroDollar *LayerZeroDollarFilterer) ParseAddedBlackList(log types.Log) (*LayerZeroDollarAddedBlackList, error) {
	event := new(LayerZeroDollarAddedBlackList)
	if err := _LayerZeroDollar.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerZeroDollarApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LayerZeroDollar contract.
type LayerZeroDollarApprovalIterator struct {
	Event *LayerZeroDollarApproval // Event containing the contract specifics and raw log

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
func (it *LayerZeroDollarApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerZeroDollarApproval)
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
		it.Event = new(LayerZeroDollarApproval)
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
func (it *LayerZeroDollarApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerZeroDollarApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerZeroDollarApproval represents a Approval event raised by the LayerZeroDollar contract.
type LayerZeroDollarApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LayerZeroDollar *LayerZeroDollarFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LayerZeroDollarApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LayerZeroDollar.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LayerZeroDollarApprovalIterator{contract: _LayerZeroDollar.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LayerZeroDollar *LayerZeroDollarFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LayerZeroDollarApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LayerZeroDollar.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerZeroDollarApproval)
				if err := _LayerZeroDollar.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LayerZeroDollar *LayerZeroDollarFilterer) ParseApproval(log types.Log) (*LayerZeroDollarApproval, error) {
	event := new(LayerZeroDollarApproval)
	if err := _LayerZeroDollar.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerZeroDollarDestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the LayerZeroDollar contract.
type LayerZeroDollarDestroyedBlackFundsIterator struct {
	Event *LayerZeroDollarDestroyedBlackFunds // Event containing the contract specifics and raw log

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
func (it *LayerZeroDollarDestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerZeroDollarDestroyedBlackFunds)
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
		it.Event = new(LayerZeroDollarDestroyedBlackFunds)
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
func (it *LayerZeroDollarDestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerZeroDollarDestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerZeroDollarDestroyedBlackFunds represents a DestroyedBlackFunds event raised by the LayerZeroDollar contract.
type LayerZeroDollarDestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_LayerZeroDollar *LayerZeroDollarFilterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*LayerZeroDollarDestroyedBlackFundsIterator, error) {

	logs, sub, err := _LayerZeroDollar.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &LayerZeroDollarDestroyedBlackFundsIterator{contract: _LayerZeroDollar.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_LayerZeroDollar *LayerZeroDollarFilterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *LayerZeroDollarDestroyedBlackFunds) (event.Subscription, error) {

	logs, sub, err := _LayerZeroDollar.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerZeroDollarDestroyedBlackFunds)
				if err := _LayerZeroDollar.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
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

// ParseDestroyedBlackFunds is a log parse operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_LayerZeroDollar *LayerZeroDollarFilterer) ParseDestroyedBlackFunds(log types.Log) (*LayerZeroDollarDestroyedBlackFunds, error) {
	event := new(LayerZeroDollarDestroyedBlackFunds)
	if err := _LayerZeroDollar.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerZeroDollarRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the LayerZeroDollar contract.
type LayerZeroDollarRemovedBlackListIterator struct {
	Event *LayerZeroDollarRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *LayerZeroDollarRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerZeroDollarRemovedBlackList)
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
		it.Event = new(LayerZeroDollarRemovedBlackList)
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
func (it *LayerZeroDollarRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerZeroDollarRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerZeroDollarRemovedBlackList represents a RemovedBlackList event raised by the LayerZeroDollar contract.
type LayerZeroDollarRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_LayerZeroDollar *LayerZeroDollarFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*LayerZeroDollarRemovedBlackListIterator, error) {

	logs, sub, err := _LayerZeroDollar.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &LayerZeroDollarRemovedBlackListIterator{contract: _LayerZeroDollar.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_LayerZeroDollar *LayerZeroDollarFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *LayerZeroDollarRemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _LayerZeroDollar.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerZeroDollarRemovedBlackList)
				if err := _LayerZeroDollar.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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

// ParseRemovedBlackList is a log parse operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_LayerZeroDollar *LayerZeroDollarFilterer) ParseRemovedBlackList(log types.Log) (*LayerZeroDollarRemovedBlackList, error) {
	event := new(LayerZeroDollarRemovedBlackList)
	if err := _LayerZeroDollar.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerZeroDollarTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LayerZeroDollar contract.
type LayerZeroDollarTransferIterator struct {
	Event *LayerZeroDollarTransfer // Event containing the contract specifics and raw log

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
func (it *LayerZeroDollarTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerZeroDollarTransfer)
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
		it.Event = new(LayerZeroDollarTransfer)
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
func (it *LayerZeroDollarTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerZeroDollarTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerZeroDollarTransfer represents a Transfer event raised by the LayerZeroDollar contract.
type LayerZeroDollarTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LayerZeroDollar *LayerZeroDollarFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LayerZeroDollarTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LayerZeroDollar.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LayerZeroDollarTransferIterator{contract: _LayerZeroDollar.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LayerZeroDollar *LayerZeroDollarFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LayerZeroDollarTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LayerZeroDollar.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerZeroDollarTransfer)
				if err := _LayerZeroDollar.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LayerZeroDollar *LayerZeroDollarFilterer) ParseTransfer(log types.Log) (*LayerZeroDollarTransfer, error) {
	event := new(LayerZeroDollarTransfer)
	if err := _LayerZeroDollar.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
