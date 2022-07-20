// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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

// ProxyMetaData contains all meta data concerning the Proxy contract.
var ProxyMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"}],\"name\":\"proxy\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"}],\"name\":\"proxyAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyMetaData.ABI instead.
var ProxyABI = ProxyMetaData.ABI

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// Proxy is a paid mutator transaction binding the contract method 0x1b0f7ba9.
//
// Solidity: function proxy(address dest, uint8 howToCall, bytes calldata) returns(bool success)
func (_Proxy *ProxyTransactor) Proxy(opts *bind.TransactOpts, dest common.Address, howToCall uint8, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "proxy", dest, howToCall, calldata)
}

// Proxy is a paid mutator transaction binding the contract method 0x1b0f7ba9.
//
// Solidity: function proxy(address dest, uint8 howToCall, bytes calldata) returns(bool success)
func (_Proxy *ProxySession) Proxy(dest common.Address, howToCall uint8, calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Proxy(&_Proxy.TransactOpts, dest, howToCall, calldata)
}

// Proxy is a paid mutator transaction binding the contract method 0x1b0f7ba9.
//
// Solidity: function proxy(address dest, uint8 howToCall, bytes calldata) returns(bool success)
func (_Proxy *ProxyTransactorSession) Proxy(dest common.Address, howToCall uint8, calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Proxy(&_Proxy.TransactOpts, dest, howToCall, calldata)
}

// ProxyAssert is a paid mutator transaction binding the contract method 0x3f801f91.
//
// Solidity: function proxyAssert(address dest, uint8 howToCall, bytes calldata) returns()
func (_Proxy *ProxyTransactor) ProxyAssert(opts *bind.TransactOpts, dest common.Address, howToCall uint8, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "proxyAssert", dest, howToCall, calldata)
}

// ProxyAssert is a paid mutator transaction binding the contract method 0x3f801f91.
//
// Solidity: function proxyAssert(address dest, uint8 howToCall, bytes calldata) returns()
func (_Proxy *ProxySession) ProxyAssert(dest common.Address, howToCall uint8, calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyAssert(&_Proxy.TransactOpts, dest, howToCall, calldata)
}

// ProxyAssert is a paid mutator transaction binding the contract method 0x3f801f91.
//
// Solidity: function proxyAssert(address dest, uint8 howToCall, bytes calldata) returns()
func (_Proxy *ProxyTransactorSession) ProxyAssert(dest common.Address, howToCall uint8, calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyAssert(&_Proxy.TransactOpts, dest, howToCall, calldata)
}
