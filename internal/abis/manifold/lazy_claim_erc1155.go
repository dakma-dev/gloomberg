// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package manifold

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

// IERC1155LazyPayableClaimClaim is an auto generated low-level Go binding around an user-defined struct.
type IERC1155LazyPayableClaimClaim struct {
	Total           uint32
	TotalMax        uint32
	WalletMax       uint32
	StartDate       *big.Int
	EndDate         *big.Int
	StorageProtocol uint8
	MerkleRoot      [32]byte
	Location        string
	TokenId         *big.Int
	Cost            *big.Int
	PaymentReceiver common.Address
	Erc20           common.Address
}

// IERC1155LazyPayableClaimClaimParameters is an auto generated low-level Go binding around an user-defined struct.
type IERC1155LazyPayableClaimClaimParameters struct {
	TotalMax        uint32
	WalletMax       uint32
	StartDate       *big.Int
	EndDate         *big.Int
	StorageProtocol uint8
	MerkleRoot      [32]byte
	Location        string
	Cost            *big.Int
	PaymentReceiver common.Address
	Erc20           common.Address
}

// LazyClaimERC1155MetaData contains all meta data concerning the LazyClaimERC1155 contract.
var LazyClaimERC1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegationRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"}],\"name\":\"ClaimInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"}],\"name\":\"ClaimMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"mintCount\",\"type\":\"uint16\"}],\"name\":\"ClaimMintBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"mintCount\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintFor\",\"type\":\"address\"}],\"name\":\"ClaimMintProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"}],\"name\":\"ClaimUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_REGISTRY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MEMBERSHIP_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_FEE_MERKLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"approveAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"mintIndex\",\"type\":\"uint32\"}],\"name\":\"checkMintIndex\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint32[]\",\"name\":\"mintIndices\",\"type\":\"uint32[]\"}],\"name\":\"checkMintIndices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"minted\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"locationChunk\",\"type\":\"string\"}],\"name\":\"extendTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"}],\"name\":\"getClaim\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"total\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalMax\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"walletMax\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"internalType\":\"structIERC1155LazyPayableClaim.Claim\",\"name\":\"claim\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getClaimForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"total\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalMax\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"walletMax\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"internalType\":\"structIERC1155LazyPayableClaim.Claim\",\"name\":\"claim\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"}],\"name\":\"getTotalMints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"totalMax\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"walletMax\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"internalType\":\"structIERC1155LazyPayableClaim.ClaimParameters\",\"name\":\"claimParameters\",\"type\":\"tuple\"}],\"name\":\"initializeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"mintIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"mintFor\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"mintCount\",\"type\":\"uint16\"},{\"internalType\":\"uint32[]\",\"name\":\"mintIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"merkleProofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"address\",\"name\":\"mintFor\",\"type\":\"address\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"mintCount\",\"type\":\"uint16\"},{\"internalType\":\"uint32[]\",\"name\":\"mintIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"merkleProofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"address\",\"name\":\"mintFor\",\"type\":\"address\"}],\"name\":\"mintProxy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"revokeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"membershipAddress\",\"type\":\"address\"}],\"name\":\"setMembershipAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"totalMax\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"walletMax\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"internalType\":\"structIERC1155LazyPayableClaim.ClaimParameters\",\"name\":\"claimParameters\",\"type\":\"tuple\"}],\"name\":\"updateClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"updateTokenURIParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LazyClaimERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use LazyClaimERC1155MetaData.ABI instead.
var LazyClaimERC1155ABI = LazyClaimERC1155MetaData.ABI

// LazyClaimERC1155 is an auto generated Go binding around an Ethereum contract.
type LazyClaimERC1155 struct {
	LazyClaimERC1155Caller     // Read-only binding to the contract
	LazyClaimERC1155Transactor // Write-only binding to the contract
	LazyClaimERC1155Filterer   // Log filterer for contract events
}

// LazyClaimERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type LazyClaimERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LazyClaimERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type LazyClaimERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LazyClaimERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LazyClaimERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LazyClaimERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LazyClaimERC1155Session struct {
	Contract     *LazyClaimERC1155 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LazyClaimERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LazyClaimERC1155CallerSession struct {
	Contract *LazyClaimERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LazyClaimERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LazyClaimERC1155TransactorSession struct {
	Contract     *LazyClaimERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LazyClaimERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type LazyClaimERC1155Raw struct {
	Contract *LazyClaimERC1155 // Generic contract binding to access the raw methods on
}

// LazyClaimERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LazyClaimERC1155CallerRaw struct {
	Contract *LazyClaimERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// LazyClaimERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LazyClaimERC1155TransactorRaw struct {
	Contract *LazyClaimERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewLazyClaimERC1155 creates a new instance of LazyClaimERC1155, bound to a specific deployed contract.
func NewLazyClaimERC1155(address common.Address, backend bind.ContractBackend) (*LazyClaimERC1155, error) {
	contract, err := bindLazyClaimERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155{LazyClaimERC1155Caller: LazyClaimERC1155Caller{contract: contract}, LazyClaimERC1155Transactor: LazyClaimERC1155Transactor{contract: contract}, LazyClaimERC1155Filterer: LazyClaimERC1155Filterer{contract: contract}}, nil
}

// NewLazyClaimERC1155Caller creates a new read-only instance of LazyClaimERC1155, bound to a specific deployed contract.
func NewLazyClaimERC1155Caller(address common.Address, caller bind.ContractCaller) (*LazyClaimERC1155Caller, error) {
	contract, err := bindLazyClaimERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155Caller{contract: contract}, nil
}

// NewLazyClaimERC1155Transactor creates a new write-only instance of LazyClaimERC1155, bound to a specific deployed contract.
func NewLazyClaimERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*LazyClaimERC1155Transactor, error) {
	contract, err := bindLazyClaimERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155Transactor{contract: contract}, nil
}

// NewLazyClaimERC1155Filterer creates a new log filterer instance of LazyClaimERC1155, bound to a specific deployed contract.
func NewLazyClaimERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*LazyClaimERC1155Filterer, error) {
	contract, err := bindLazyClaimERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155Filterer{contract: contract}, nil
}

// bindLazyClaimERC1155 binds a generic wrapper to an already deployed contract.
func bindLazyClaimERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LazyClaimERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LazyClaimERC1155 *LazyClaimERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LazyClaimERC1155.Contract.LazyClaimERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LazyClaimERC1155 *LazyClaimERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.LazyClaimERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LazyClaimERC1155 *LazyClaimERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.LazyClaimERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LazyClaimERC1155 *LazyClaimERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LazyClaimERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONREGISTRY is a free data retrieval call binding the contract method 0x4daadff7.
//
// Solidity: function DELEGATION_REGISTRY() view returns(address)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) DELEGATIONREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "DELEGATION_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DELEGATIONREGISTRY is a free data retrieval call binding the contract method 0x4daadff7.
//
// Solidity: function DELEGATION_REGISTRY() view returns(address)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) DELEGATIONREGISTRY() (common.Address, error) {
	return _LazyClaimERC1155.Contract.DELEGATIONREGISTRY(&_LazyClaimERC1155.CallOpts)
}

// DELEGATIONREGISTRY is a free data retrieval call binding the contract method 0x4daadff7.
//
// Solidity: function DELEGATION_REGISTRY() view returns(address)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) DELEGATIONREGISTRY() (common.Address, error) {
	return _LazyClaimERC1155.Contract.DELEGATIONREGISTRY(&_LazyClaimERC1155.CallOpts)
}

// MEMBERSHIPADDRESS is a free data retrieval call binding the contract method 0x4baa62bf.
//
// Solidity: function MEMBERSHIP_ADDRESS() view returns(address)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) MEMBERSHIPADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "MEMBERSHIP_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MEMBERSHIPADDRESS is a free data retrieval call binding the contract method 0x4baa62bf.
//
// Solidity: function MEMBERSHIP_ADDRESS() view returns(address)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) MEMBERSHIPADDRESS() (common.Address, error) {
	return _LazyClaimERC1155.Contract.MEMBERSHIPADDRESS(&_LazyClaimERC1155.CallOpts)
}

// MEMBERSHIPADDRESS is a free data retrieval call binding the contract method 0x4baa62bf.
//
// Solidity: function MEMBERSHIP_ADDRESS() view returns(address)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) MEMBERSHIPADDRESS() (common.Address, error) {
	return _LazyClaimERC1155.Contract.MEMBERSHIPADDRESS(&_LazyClaimERC1155.CallOpts)
}

// MINTFEE is a free data retrieval call binding the contract method 0xd7bf81a3.
//
// Solidity: function MINT_FEE() view returns(uint256)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) MINTFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "MINT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTFEE is a free data retrieval call binding the contract method 0xd7bf81a3.
//
// Solidity: function MINT_FEE() view returns(uint256)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) MINTFEE() (*big.Int, error) {
	return _LazyClaimERC1155.Contract.MINTFEE(&_LazyClaimERC1155.CallOpts)
}

// MINTFEE is a free data retrieval call binding the contract method 0xd7bf81a3.
//
// Solidity: function MINT_FEE() view returns(uint256)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) MINTFEE() (*big.Int, error) {
	return _LazyClaimERC1155.Contract.MINTFEE(&_LazyClaimERC1155.CallOpts)
}

// MINTFEEMERKLE is a free data retrieval call binding the contract method 0xcb799716.
//
// Solidity: function MINT_FEE_MERKLE() view returns(uint256)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) MINTFEEMERKLE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "MINT_FEE_MERKLE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTFEEMERKLE is a free data retrieval call binding the contract method 0xcb799716.
//
// Solidity: function MINT_FEE_MERKLE() view returns(uint256)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) MINTFEEMERKLE() (*big.Int, error) {
	return _LazyClaimERC1155.Contract.MINTFEEMERKLE(&_LazyClaimERC1155.CallOpts)
}

// MINTFEEMERKLE is a free data retrieval call binding the contract method 0xcb799716.
//
// Solidity: function MINT_FEE_MERKLE() view returns(uint256)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) MINTFEEMERKLE() (*big.Int, error) {
	return _LazyClaimERC1155.Contract.MINTFEEMERKLE(&_LazyClaimERC1155.CallOpts)
}

// CheckMintIndex is a free data retrieval call binding the contract method 0xcda08536.
//
// Solidity: function checkMintIndex(address creatorContractAddress, uint256 instanceId, uint32 mintIndex) view returns(bool)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) CheckMintIndex(opts *bind.CallOpts, creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32) (bool, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "checkMintIndex", creatorContractAddress, instanceId, mintIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMintIndex is a free data retrieval call binding the contract method 0xcda08536.
//
// Solidity: function checkMintIndex(address creatorContractAddress, uint256 instanceId, uint32 mintIndex) view returns(bool)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) CheckMintIndex(creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32) (bool, error) {
	return _LazyClaimERC1155.Contract.CheckMintIndex(&_LazyClaimERC1155.CallOpts, creatorContractAddress, instanceId, mintIndex)
}

// CheckMintIndex is a free data retrieval call binding the contract method 0xcda08536.
//
// Solidity: function checkMintIndex(address creatorContractAddress, uint256 instanceId, uint32 mintIndex) view returns(bool)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) CheckMintIndex(creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32) (bool, error) {
	return _LazyClaimERC1155.Contract.CheckMintIndex(&_LazyClaimERC1155.CallOpts, creatorContractAddress, instanceId, mintIndex)
}

// CheckMintIndices is a free data retrieval call binding the contract method 0xf8a6137b.
//
// Solidity: function checkMintIndices(address creatorContractAddress, uint256 instanceId, uint32[] mintIndices) view returns(bool[] minted)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) CheckMintIndices(opts *bind.CallOpts, creatorContractAddress common.Address, instanceId *big.Int, mintIndices []uint32) ([]bool, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "checkMintIndices", creatorContractAddress, instanceId, mintIndices)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckMintIndices is a free data retrieval call binding the contract method 0xf8a6137b.
//
// Solidity: function checkMintIndices(address creatorContractAddress, uint256 instanceId, uint32[] mintIndices) view returns(bool[] minted)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) CheckMintIndices(creatorContractAddress common.Address, instanceId *big.Int, mintIndices []uint32) ([]bool, error) {
	return _LazyClaimERC1155.Contract.CheckMintIndices(&_LazyClaimERC1155.CallOpts, creatorContractAddress, instanceId, mintIndices)
}

// CheckMintIndices is a free data retrieval call binding the contract method 0xf8a6137b.
//
// Solidity: function checkMintIndices(address creatorContractAddress, uint256 instanceId, uint32[] mintIndices) view returns(bool[] minted)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) CheckMintIndices(creatorContractAddress common.Address, instanceId *big.Int, mintIndices []uint32) ([]bool, error) {
	return _LazyClaimERC1155.Contract.CheckMintIndices(&_LazyClaimERC1155.CallOpts, creatorContractAddress, instanceId, mintIndices)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) GetAdmins() ([]common.Address, error) {
	return _LazyClaimERC1155.Contract.GetAdmins(&_LazyClaimERC1155.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) GetAdmins() ([]common.Address, error) {
	return _LazyClaimERC1155.Contract.GetAdmins(&_LazyClaimERC1155.CallOpts)
}

// GetClaim is a free data retrieval call binding the contract method 0x0f79ab39.
//
// Solidity: function getClaim(address creatorContractAddress, uint256 instanceId) view returns((uint32,uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,uint256,address,address) claim)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) GetClaim(opts *bind.CallOpts, creatorContractAddress common.Address, instanceId *big.Int) (IERC1155LazyPayableClaimClaim, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "getClaim", creatorContractAddress, instanceId)

	if err != nil {
		return *new(IERC1155LazyPayableClaimClaim), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC1155LazyPayableClaimClaim)).(*IERC1155LazyPayableClaimClaim)

	return out0, err

}

// GetClaim is a free data retrieval call binding the contract method 0x0f79ab39.
//
// Solidity: function getClaim(address creatorContractAddress, uint256 instanceId) view returns((uint32,uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,uint256,address,address) claim)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) GetClaim(creatorContractAddress common.Address, instanceId *big.Int) (IERC1155LazyPayableClaimClaim, error) {
	return _LazyClaimERC1155.Contract.GetClaim(&_LazyClaimERC1155.CallOpts, creatorContractAddress, instanceId)
}

// GetClaim is a free data retrieval call binding the contract method 0x0f79ab39.
//
// Solidity: function getClaim(address creatorContractAddress, uint256 instanceId) view returns((uint32,uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,uint256,address,address) claim)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) GetClaim(creatorContractAddress common.Address, instanceId *big.Int) (IERC1155LazyPayableClaimClaim, error) {
	return _LazyClaimERC1155.Contract.GetClaim(&_LazyClaimERC1155.CallOpts, creatorContractAddress, instanceId)
}

// GetClaimForToken is a free data retrieval call binding the contract method 0x895696f2.
//
// Solidity: function getClaimForToken(address creatorContractAddress, uint256 tokenId) view returns(uint256 instanceId, (uint32,uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,uint256,address,address) claim)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) GetClaimForToken(opts *bind.CallOpts, creatorContractAddress common.Address, tokenId *big.Int) (struct {
	InstanceId *big.Int
	Claim      IERC1155LazyPayableClaimClaim
}, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "getClaimForToken", creatorContractAddress, tokenId)

	outstruct := new(struct {
		InstanceId *big.Int
		Claim      IERC1155LazyPayableClaimClaim
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InstanceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Claim = *abi.ConvertType(out[1], new(IERC1155LazyPayableClaimClaim)).(*IERC1155LazyPayableClaimClaim)

	return *outstruct, err

}

// GetClaimForToken is a free data retrieval call binding the contract method 0x895696f2.
//
// Solidity: function getClaimForToken(address creatorContractAddress, uint256 tokenId) view returns(uint256 instanceId, (uint32,uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,uint256,address,address) claim)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) GetClaimForToken(creatorContractAddress common.Address, tokenId *big.Int) (struct {
	InstanceId *big.Int
	Claim      IERC1155LazyPayableClaimClaim
}, error) {
	return _LazyClaimERC1155.Contract.GetClaimForToken(&_LazyClaimERC1155.CallOpts, creatorContractAddress, tokenId)
}

// GetClaimForToken is a free data retrieval call binding the contract method 0x895696f2.
//
// Solidity: function getClaimForToken(address creatorContractAddress, uint256 tokenId) view returns(uint256 instanceId, (uint32,uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,uint256,address,address) claim)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) GetClaimForToken(creatorContractAddress common.Address, tokenId *big.Int) (struct {
	InstanceId *big.Int
	Claim      IERC1155LazyPayableClaimClaim
}, error) {
	return _LazyClaimERC1155.Contract.GetClaimForToken(&_LazyClaimERC1155.CallOpts, creatorContractAddress, tokenId)
}

// GetTotalMints is a free data retrieval call binding the contract method 0x42f3bef4.
//
// Solidity: function getTotalMints(address minter, address creatorContractAddress, uint256 instanceId) view returns(uint32)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) GetTotalMints(opts *bind.CallOpts, minter common.Address, creatorContractAddress common.Address, instanceId *big.Int) (uint32, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "getTotalMints", minter, creatorContractAddress, instanceId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTotalMints is a free data retrieval call binding the contract method 0x42f3bef4.
//
// Solidity: function getTotalMints(address minter, address creatorContractAddress, uint256 instanceId) view returns(uint32)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) GetTotalMints(minter common.Address, creatorContractAddress common.Address, instanceId *big.Int) (uint32, error) {
	return _LazyClaimERC1155.Contract.GetTotalMints(&_LazyClaimERC1155.CallOpts, minter, creatorContractAddress, instanceId)
}

// GetTotalMints is a free data retrieval call binding the contract method 0x42f3bef4.
//
// Solidity: function getTotalMints(address minter, address creatorContractAddress, uint256 instanceId) view returns(uint32)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) GetTotalMints(minter common.Address, creatorContractAddress common.Address, instanceId *big.Int) (uint32, error) {
	return _LazyClaimERC1155.Contract.GetTotalMints(&_LazyClaimERC1155.CallOpts, minter, creatorContractAddress, instanceId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) IsAdmin(opts *bind.CallOpts, admin common.Address) (bool, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "isAdmin", admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) IsAdmin(admin common.Address) (bool, error) {
	return _LazyClaimERC1155.Contract.IsAdmin(&_LazyClaimERC1155.CallOpts, admin)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) IsAdmin(admin common.Address) (bool, error) {
	return _LazyClaimERC1155.Contract.IsAdmin(&_LazyClaimERC1155.CallOpts, admin)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) Owner() (common.Address, error) {
	return _LazyClaimERC1155.Contract.Owner(&_LazyClaimERC1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) Owner() (common.Address, error) {
	return _LazyClaimERC1155.Contract.Owner(&_LazyClaimERC1155.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LazyClaimERC1155.Contract.SupportsInterface(&_LazyClaimERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LazyClaimERC1155.Contract.SupportsInterface(&_LazyClaimERC1155.CallOpts, interfaceId)
}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address creatorContractAddress, uint256 tokenId) view returns(string uri)
func (_LazyClaimERC1155 *LazyClaimERC1155Caller) TokenURI(opts *bind.CallOpts, creatorContractAddress common.Address, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LazyClaimERC1155.contract.Call(opts, &out, "tokenURI", creatorContractAddress, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address creatorContractAddress, uint256 tokenId) view returns(string uri)
func (_LazyClaimERC1155 *LazyClaimERC1155Session) TokenURI(creatorContractAddress common.Address, tokenId *big.Int) (string, error) {
	return _LazyClaimERC1155.Contract.TokenURI(&_LazyClaimERC1155.CallOpts, creatorContractAddress, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address creatorContractAddress, uint256 tokenId) view returns(string uri)
func (_LazyClaimERC1155 *LazyClaimERC1155CallerSession) TokenURI(creatorContractAddress common.Address, tokenId *big.Int) (string, error) {
	return _LazyClaimERC1155.Contract.TokenURI(&_LazyClaimERC1155.CallOpts, creatorContractAddress, tokenId)
}

// Airdrop is a paid mutator transaction binding the contract method 0xbd04e411.
//
// Solidity: function airdrop(address creatorContractAddress, uint256 instanceId, address[] recipients, uint256[] amounts) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) Airdrop(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "airdrop", creatorContractAddress, instanceId, recipients, amounts)
}

// Airdrop is a paid mutator transaction binding the contract method 0xbd04e411.
//
// Solidity: function airdrop(address creatorContractAddress, uint256 instanceId, address[] recipients, uint256[] amounts) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) Airdrop(creatorContractAddress common.Address, instanceId *big.Int, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.Airdrop(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, recipients, amounts)
}

// Airdrop is a paid mutator transaction binding the contract method 0xbd04e411.
//
// Solidity: function airdrop(address creatorContractAddress, uint256 instanceId, address[] recipients, uint256[] amounts) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) Airdrop(creatorContractAddress common.Address, instanceId *big.Int, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.Airdrop(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, recipients, amounts)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) ApproveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "approveAdmin", admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.ApproveAdmin(&_LazyClaimERC1155.TransactOpts, admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.ApproveAdmin(&_LazyClaimERC1155.TransactOpts, admin)
}

// ExtendTokenURI is a paid mutator transaction binding the contract method 0xb93aa86c.
//
// Solidity: function extendTokenURI(address creatorContractAddress, uint256 instanceId, string locationChunk) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) ExtendTokenURI(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, locationChunk string) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "extendTokenURI", creatorContractAddress, instanceId, locationChunk)
}

// ExtendTokenURI is a paid mutator transaction binding the contract method 0xb93aa86c.
//
// Solidity: function extendTokenURI(address creatorContractAddress, uint256 instanceId, string locationChunk) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) ExtendTokenURI(creatorContractAddress common.Address, instanceId *big.Int, locationChunk string) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.ExtendTokenURI(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, locationChunk)
}

// ExtendTokenURI is a paid mutator transaction binding the contract method 0xb93aa86c.
//
// Solidity: function extendTokenURI(address creatorContractAddress, uint256 instanceId, string locationChunk) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) ExtendTokenURI(creatorContractAddress common.Address, instanceId *big.Int, locationChunk string) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.ExtendTokenURI(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, locationChunk)
}

// InitializeClaim is a paid mutator transaction binding the contract method 0xd670c080.
//
// Solidity: function initializeClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,address,address) claimParameters) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) InitializeClaim(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC1155LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "initializeClaim", creatorContractAddress, instanceId, claimParameters)
}

// InitializeClaim is a paid mutator transaction binding the contract method 0xd670c080.
//
// Solidity: function initializeClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,address,address) claimParameters) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) InitializeClaim(creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC1155LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.InitializeClaim(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, claimParameters)
}

// InitializeClaim is a paid mutator transaction binding the contract method 0xd670c080.
//
// Solidity: function initializeClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,address,address) claimParameters) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) InitializeClaim(creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC1155LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.InitializeClaim(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, claimParameters)
}

// Mint is a paid mutator transaction binding the contract method 0xfa2b068f.
//
// Solidity: function mint(address creatorContractAddress, uint256 instanceId, uint32 mintIndex, bytes32[] merkleProof, address mintFor) payable returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) Mint(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32, merkleProof [][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "mint", creatorContractAddress, instanceId, mintIndex, merkleProof, mintFor)
}

// Mint is a paid mutator transaction binding the contract method 0xfa2b068f.
//
// Solidity: function mint(address creatorContractAddress, uint256 instanceId, uint32 mintIndex, bytes32[] merkleProof, address mintFor) payable returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) Mint(creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32, merkleProof [][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.Mint(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, mintIndex, merkleProof, mintFor)
}

// Mint is a paid mutator transaction binding the contract method 0xfa2b068f.
//
// Solidity: function mint(address creatorContractAddress, uint256 instanceId, uint32 mintIndex, bytes32[] merkleProof, address mintFor) payable returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) Mint(creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32, merkleProof [][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.Mint(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, mintIndex, merkleProof, mintFor)
}

// MintBatch is a paid mutator transaction binding the contract method 0x26c858a4.
//
// Solidity: function mintBatch(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) MintBatch(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "mintBatch", creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintBatch is a paid mutator transaction binding the contract method 0x26c858a4.
//
// Solidity: function mintBatch(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) MintBatch(creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.MintBatch(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintBatch is a paid mutator transaction binding the contract method 0x26c858a4.
//
// Solidity: function mintBatch(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) MintBatch(creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.MintBatch(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintProxy is a paid mutator transaction binding the contract method 0x07591acc.
//
// Solidity: function mintProxy(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) MintProxy(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "mintProxy", creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintProxy is a paid mutator transaction binding the contract method 0x07591acc.
//
// Solidity: function mintProxy(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) MintProxy(creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.MintProxy(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintProxy is a paid mutator transaction binding the contract method 0x07591acc.
//
// Solidity: function mintProxy(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) MintProxy(creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.MintProxy(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) RenounceOwnership() (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.RenounceOwnership(&_LazyClaimERC1155.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.RenounceOwnership(&_LazyClaimERC1155.TransactOpts)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) RevokeAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "revokeAdmin", admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.RevokeAdmin(&_LazyClaimERC1155.TransactOpts, admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.RevokeAdmin(&_LazyClaimERC1155.TransactOpts, admin)
}

// SetMembershipAddress is a paid mutator transaction binding the contract method 0x7ab39392.
//
// Solidity: function setMembershipAddress(address membershipAddress) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) SetMembershipAddress(opts *bind.TransactOpts, membershipAddress common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "setMembershipAddress", membershipAddress)
}

// SetMembershipAddress is a paid mutator transaction binding the contract method 0x7ab39392.
//
// Solidity: function setMembershipAddress(address membershipAddress) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) SetMembershipAddress(membershipAddress common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.SetMembershipAddress(&_LazyClaimERC1155.TransactOpts, membershipAddress)
}

// SetMembershipAddress is a paid mutator transaction binding the contract method 0x7ab39392.
//
// Solidity: function setMembershipAddress(address membershipAddress) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) SetMembershipAddress(membershipAddress common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.SetMembershipAddress(&_LazyClaimERC1155.TransactOpts, membershipAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.TransferOwnership(&_LazyClaimERC1155.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.TransferOwnership(&_LazyClaimERC1155.TransactOpts, newOwner)
}

// UpdateClaim is a paid mutator transaction binding the contract method 0x0a6330b8.
//
// Solidity: function updateClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,address,address) claimParameters) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) UpdateClaim(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC1155LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "updateClaim", creatorContractAddress, instanceId, claimParameters)
}

// UpdateClaim is a paid mutator transaction binding the contract method 0x0a6330b8.
//
// Solidity: function updateClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,address,address) claimParameters) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) UpdateClaim(creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC1155LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.UpdateClaim(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, claimParameters)
}

// UpdateClaim is a paid mutator transaction binding the contract method 0x0a6330b8.
//
// Solidity: function updateClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bytes32,string,uint256,address,address) claimParameters) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) UpdateClaim(creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC1155LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.UpdateClaim(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, claimParameters)
}

// UpdateTokenURIParams is a paid mutator transaction binding the contract method 0x6e12056a.
//
// Solidity: function updateTokenURIParams(address creatorContractAddress, uint256 instanceId, uint8 storageProtocol, string location) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) UpdateTokenURIParams(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, storageProtocol uint8, location string) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "updateTokenURIParams", creatorContractAddress, instanceId, storageProtocol, location)
}

// UpdateTokenURIParams is a paid mutator transaction binding the contract method 0x6e12056a.
//
// Solidity: function updateTokenURIParams(address creatorContractAddress, uint256 instanceId, uint8 storageProtocol, string location) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) UpdateTokenURIParams(creatorContractAddress common.Address, instanceId *big.Int, storageProtocol uint8, location string) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.UpdateTokenURIParams(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, storageProtocol, location)
}

// UpdateTokenURIParams is a paid mutator transaction binding the contract method 0x6e12056a.
//
// Solidity: function updateTokenURIParams(address creatorContractAddress, uint256 instanceId, uint8 storageProtocol, string location) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) UpdateTokenURIParams(creatorContractAddress common.Address, instanceId *big.Int, storageProtocol uint8, location string) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.UpdateTokenURIParams(&_LazyClaimERC1155.TransactOpts, creatorContractAddress, instanceId, storageProtocol, location)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Transactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LazyClaimERC1155.contract.Transact(opts, "withdraw", receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155Session) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.Withdraw(&_LazyClaimERC1155.TransactOpts, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_LazyClaimERC1155 *LazyClaimERC1155TransactorSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LazyClaimERC1155.Contract.Withdraw(&_LazyClaimERC1155.TransactOpts, receiver, amount)
}

// LazyClaimERC1155AdminApprovedIterator is returned from FilterAdminApproved and is used to iterate over the raw logs and unpacked data for AdminApproved events raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155AdminApprovedIterator struct {
	Event *LazyClaimERC1155AdminApproved // Event containing the contract specifics and raw log

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
func (it *LazyClaimERC1155AdminApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LazyClaimERC1155AdminApproved)
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
		it.Event = new(LazyClaimERC1155AdminApproved)
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
func (it *LazyClaimERC1155AdminApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LazyClaimERC1155AdminApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LazyClaimERC1155AdminApproved represents a AdminApproved event raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155AdminApproved struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminApproved is a free log retrieval operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) FilterAdminApproved(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*LazyClaimERC1155AdminApprovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.FilterLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155AdminApprovedIterator{contract: _LazyClaimERC1155.contract, event: "AdminApproved", logs: logs, sub: sub}, nil
}

// WatchAdminApproved is a free log subscription operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) WatchAdminApproved(opts *bind.WatchOpts, sink chan<- *LazyClaimERC1155AdminApproved, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.WatchLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LazyClaimERC1155AdminApproved)
				if err := _LazyClaimERC1155.contract.UnpackLog(event, "AdminApproved", log); err != nil {
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

// ParseAdminApproved is a log parse operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) ParseAdminApproved(log types.Log) (*LazyClaimERC1155AdminApproved, error) {
	event := new(LazyClaimERC1155AdminApproved)
	if err := _LazyClaimERC1155.contract.UnpackLog(event, "AdminApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LazyClaimERC1155AdminRevokedIterator is returned from FilterAdminRevoked and is used to iterate over the raw logs and unpacked data for AdminRevoked events raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155AdminRevokedIterator struct {
	Event *LazyClaimERC1155AdminRevoked // Event containing the contract specifics and raw log

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
func (it *LazyClaimERC1155AdminRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LazyClaimERC1155AdminRevoked)
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
		it.Event = new(LazyClaimERC1155AdminRevoked)
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
func (it *LazyClaimERC1155AdminRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LazyClaimERC1155AdminRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LazyClaimERC1155AdminRevoked represents a AdminRevoked event raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155AdminRevoked struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRevoked is a free log retrieval operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) FilterAdminRevoked(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*LazyClaimERC1155AdminRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.FilterLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155AdminRevokedIterator{contract: _LazyClaimERC1155.contract, event: "AdminRevoked", logs: logs, sub: sub}, nil
}

// WatchAdminRevoked is a free log subscription operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) WatchAdminRevoked(opts *bind.WatchOpts, sink chan<- *LazyClaimERC1155AdminRevoked, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.WatchLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LazyClaimERC1155AdminRevoked)
				if err := _LazyClaimERC1155.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
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

// ParseAdminRevoked is a log parse operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) ParseAdminRevoked(log types.Log) (*LazyClaimERC1155AdminRevoked, error) {
	event := new(LazyClaimERC1155AdminRevoked)
	if err := _LazyClaimERC1155.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LazyClaimERC1155ClaimInitializedIterator is returned from FilterClaimInitialized and is used to iterate over the raw logs and unpacked data for ClaimInitialized events raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimInitializedIterator struct {
	Event *LazyClaimERC1155ClaimInitialized // Event containing the contract specifics and raw log

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
func (it *LazyClaimERC1155ClaimInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LazyClaimERC1155ClaimInitialized)
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
		it.Event = new(LazyClaimERC1155ClaimInitialized)
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
func (it *LazyClaimERC1155ClaimInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LazyClaimERC1155ClaimInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LazyClaimERC1155ClaimInitialized represents a ClaimInitialized event raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimInitialized struct {
	CreatorContract common.Address
	InstanceId      *big.Int
	Initializer     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimInitialized is a free log retrieval operation binding the contract event 0xd02727da4c6c6c111e00310108209a4de39f6817414df43ca1a10730d47c6a34.
//
// Solidity: event ClaimInitialized(address indexed creatorContract, uint256 indexed instanceId, address initializer)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) FilterClaimInitialized(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*LazyClaimERC1155ClaimInitializedIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.FilterLogs(opts, "ClaimInitialized", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155ClaimInitializedIterator{contract: _LazyClaimERC1155.contract, event: "ClaimInitialized", logs: logs, sub: sub}, nil
}

// WatchClaimInitialized is a free log subscription operation binding the contract event 0xd02727da4c6c6c111e00310108209a4de39f6817414df43ca1a10730d47c6a34.
//
// Solidity: event ClaimInitialized(address indexed creatorContract, uint256 indexed instanceId, address initializer)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) WatchClaimInitialized(opts *bind.WatchOpts, sink chan<- *LazyClaimERC1155ClaimInitialized, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.WatchLogs(opts, "ClaimInitialized", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LazyClaimERC1155ClaimInitialized)
				if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimInitialized", log); err != nil {
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

// ParseClaimInitialized is a log parse operation binding the contract event 0xd02727da4c6c6c111e00310108209a4de39f6817414df43ca1a10730d47c6a34.
//
// Solidity: event ClaimInitialized(address indexed creatorContract, uint256 indexed instanceId, address initializer)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) ParseClaimInitialized(log types.Log) (*LazyClaimERC1155ClaimInitialized, error) {
	event := new(LazyClaimERC1155ClaimInitialized)
	if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LazyClaimERC1155ClaimMintIterator is returned from FilterClaimMint and is used to iterate over the raw logs and unpacked data for ClaimMint events raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimMintIterator struct {
	Event *LazyClaimERC1155ClaimMint // Event containing the contract specifics and raw log

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
func (it *LazyClaimERC1155ClaimMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LazyClaimERC1155ClaimMint)
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
		it.Event = new(LazyClaimERC1155ClaimMint)
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
func (it *LazyClaimERC1155ClaimMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LazyClaimERC1155ClaimMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LazyClaimERC1155ClaimMint represents a ClaimMint event raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimMint struct {
	CreatorContract common.Address
	InstanceId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimMint is a free log retrieval operation binding the contract event 0x5d404f369772cfab2b65717fca9bc2077efeab89a0dbec036bf0c13783154eb1.
//
// Solidity: event ClaimMint(address indexed creatorContract, uint256 indexed instanceId)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) FilterClaimMint(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*LazyClaimERC1155ClaimMintIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.FilterLogs(opts, "ClaimMint", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155ClaimMintIterator{contract: _LazyClaimERC1155.contract, event: "ClaimMint", logs: logs, sub: sub}, nil
}

// WatchClaimMint is a free log subscription operation binding the contract event 0x5d404f369772cfab2b65717fca9bc2077efeab89a0dbec036bf0c13783154eb1.
//
// Solidity: event ClaimMint(address indexed creatorContract, uint256 indexed instanceId)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) WatchClaimMint(opts *bind.WatchOpts, sink chan<- *LazyClaimERC1155ClaimMint, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.WatchLogs(opts, "ClaimMint", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LazyClaimERC1155ClaimMint)
				if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimMint", log); err != nil {
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

// ParseClaimMint is a log parse operation binding the contract event 0x5d404f369772cfab2b65717fca9bc2077efeab89a0dbec036bf0c13783154eb1.
//
// Solidity: event ClaimMint(address indexed creatorContract, uint256 indexed instanceId)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) ParseClaimMint(log types.Log) (*LazyClaimERC1155ClaimMint, error) {
	event := new(LazyClaimERC1155ClaimMint)
	if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LazyClaimERC1155ClaimMintBatchIterator is returned from FilterClaimMintBatch and is used to iterate over the raw logs and unpacked data for ClaimMintBatch events raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimMintBatchIterator struct {
	Event *LazyClaimERC1155ClaimMintBatch // Event containing the contract specifics and raw log

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
func (it *LazyClaimERC1155ClaimMintBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LazyClaimERC1155ClaimMintBatch)
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
		it.Event = new(LazyClaimERC1155ClaimMintBatch)
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
func (it *LazyClaimERC1155ClaimMintBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LazyClaimERC1155ClaimMintBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LazyClaimERC1155ClaimMintBatch represents a ClaimMintBatch event raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimMintBatch struct {
	CreatorContract common.Address
	InstanceId      *big.Int
	MintCount       uint16
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimMintBatch is a free log retrieval operation binding the contract event 0x74f5d3254dfa39a7b1217a27d5d9b3e061eafe11720eca1cf499da2dc1eb1259.
//
// Solidity: event ClaimMintBatch(address indexed creatorContract, uint256 indexed instanceId, uint16 mintCount)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) FilterClaimMintBatch(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*LazyClaimERC1155ClaimMintBatchIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.FilterLogs(opts, "ClaimMintBatch", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155ClaimMintBatchIterator{contract: _LazyClaimERC1155.contract, event: "ClaimMintBatch", logs: logs, sub: sub}, nil
}

// WatchClaimMintBatch is a free log subscription operation binding the contract event 0x74f5d3254dfa39a7b1217a27d5d9b3e061eafe11720eca1cf499da2dc1eb1259.
//
// Solidity: event ClaimMintBatch(address indexed creatorContract, uint256 indexed instanceId, uint16 mintCount)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) WatchClaimMintBatch(opts *bind.WatchOpts, sink chan<- *LazyClaimERC1155ClaimMintBatch, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.WatchLogs(opts, "ClaimMintBatch", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LazyClaimERC1155ClaimMintBatch)
				if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimMintBatch", log); err != nil {
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

// ParseClaimMintBatch is a log parse operation binding the contract event 0x74f5d3254dfa39a7b1217a27d5d9b3e061eafe11720eca1cf499da2dc1eb1259.
//
// Solidity: event ClaimMintBatch(address indexed creatorContract, uint256 indexed instanceId, uint16 mintCount)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) ParseClaimMintBatch(log types.Log) (*LazyClaimERC1155ClaimMintBatch, error) {
	event := new(LazyClaimERC1155ClaimMintBatch)
	if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimMintBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LazyClaimERC1155ClaimMintProxyIterator is returned from FilterClaimMintProxy and is used to iterate over the raw logs and unpacked data for ClaimMintProxy events raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimMintProxyIterator struct {
	Event *LazyClaimERC1155ClaimMintProxy // Event containing the contract specifics and raw log

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
func (it *LazyClaimERC1155ClaimMintProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LazyClaimERC1155ClaimMintProxy)
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
		it.Event = new(LazyClaimERC1155ClaimMintProxy)
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
func (it *LazyClaimERC1155ClaimMintProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LazyClaimERC1155ClaimMintProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LazyClaimERC1155ClaimMintProxy represents a ClaimMintProxy event raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimMintProxy struct {
	CreatorContract common.Address
	InstanceId      *big.Int
	MintCount       uint16
	Proxy           common.Address
	MintFor         common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimMintProxy is a free log retrieval operation binding the contract event 0x61039ad47d0b05ec206a4450fd164cc2055af66ac594c12b8dd747e8803a90de.
//
// Solidity: event ClaimMintProxy(address indexed creatorContract, uint256 indexed instanceId, uint16 mintCount, address proxy, address mintFor)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) FilterClaimMintProxy(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*LazyClaimERC1155ClaimMintProxyIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.FilterLogs(opts, "ClaimMintProxy", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155ClaimMintProxyIterator{contract: _LazyClaimERC1155.contract, event: "ClaimMintProxy", logs: logs, sub: sub}, nil
}

// WatchClaimMintProxy is a free log subscription operation binding the contract event 0x61039ad47d0b05ec206a4450fd164cc2055af66ac594c12b8dd747e8803a90de.
//
// Solidity: event ClaimMintProxy(address indexed creatorContract, uint256 indexed instanceId, uint16 mintCount, address proxy, address mintFor)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) WatchClaimMintProxy(opts *bind.WatchOpts, sink chan<- *LazyClaimERC1155ClaimMintProxy, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.WatchLogs(opts, "ClaimMintProxy", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LazyClaimERC1155ClaimMintProxy)
				if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimMintProxy", log); err != nil {
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

// ParseClaimMintProxy is a log parse operation binding the contract event 0x61039ad47d0b05ec206a4450fd164cc2055af66ac594c12b8dd747e8803a90de.
//
// Solidity: event ClaimMintProxy(address indexed creatorContract, uint256 indexed instanceId, uint16 mintCount, address proxy, address mintFor)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) ParseClaimMintProxy(log types.Log) (*LazyClaimERC1155ClaimMintProxy, error) {
	event := new(LazyClaimERC1155ClaimMintProxy)
	if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimMintProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LazyClaimERC1155ClaimUpdatedIterator is returned from FilterClaimUpdated and is used to iterate over the raw logs and unpacked data for ClaimUpdated events raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimUpdatedIterator struct {
	Event *LazyClaimERC1155ClaimUpdated // Event containing the contract specifics and raw log

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
func (it *LazyClaimERC1155ClaimUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LazyClaimERC1155ClaimUpdated)
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
		it.Event = new(LazyClaimERC1155ClaimUpdated)
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
func (it *LazyClaimERC1155ClaimUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LazyClaimERC1155ClaimUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LazyClaimERC1155ClaimUpdated represents a ClaimUpdated event raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155ClaimUpdated struct {
	CreatorContract common.Address
	InstanceId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimUpdated is a free log retrieval operation binding the contract event 0x657336af9bb6c51d60c05491508d7d3026a24ee549d7a0af42e44c75bfaec47c.
//
// Solidity: event ClaimUpdated(address indexed creatorContract, uint256 indexed instanceId)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) FilterClaimUpdated(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*LazyClaimERC1155ClaimUpdatedIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.FilterLogs(opts, "ClaimUpdated", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155ClaimUpdatedIterator{contract: _LazyClaimERC1155.contract, event: "ClaimUpdated", logs: logs, sub: sub}, nil
}

// WatchClaimUpdated is a free log subscription operation binding the contract event 0x657336af9bb6c51d60c05491508d7d3026a24ee549d7a0af42e44c75bfaec47c.
//
// Solidity: event ClaimUpdated(address indexed creatorContract, uint256 indexed instanceId)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) WatchClaimUpdated(opts *bind.WatchOpts, sink chan<- *LazyClaimERC1155ClaimUpdated, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.WatchLogs(opts, "ClaimUpdated", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LazyClaimERC1155ClaimUpdated)
				if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimUpdated", log); err != nil {
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

// ParseClaimUpdated is a log parse operation binding the contract event 0x657336af9bb6c51d60c05491508d7d3026a24ee549d7a0af42e44c75bfaec47c.
//
// Solidity: event ClaimUpdated(address indexed creatorContract, uint256 indexed instanceId)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) ParseClaimUpdated(log types.Log) (*LazyClaimERC1155ClaimUpdated, error) {
	event := new(LazyClaimERC1155ClaimUpdated)
	if err := _LazyClaimERC1155.contract.UnpackLog(event, "ClaimUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LazyClaimERC1155OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155OwnershipTransferredIterator struct {
	Event *LazyClaimERC1155OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LazyClaimERC1155OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LazyClaimERC1155OwnershipTransferred)
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
		it.Event = new(LazyClaimERC1155OwnershipTransferred)
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
func (it *LazyClaimERC1155OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LazyClaimERC1155OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LazyClaimERC1155OwnershipTransferred represents a OwnershipTransferred event raised by the LazyClaimERC1155 contract.
type LazyClaimERC1155OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LazyClaimERC1155OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LazyClaimERC1155OwnershipTransferredIterator{contract: _LazyClaimERC1155.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LazyClaimERC1155OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LazyClaimERC1155.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LazyClaimERC1155OwnershipTransferred)
				if err := _LazyClaimERC1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LazyClaimERC1155 *LazyClaimERC1155Filterer) ParseOwnershipTransferred(log types.Log) (*LazyClaimERC1155OwnershipTransferred, error) {
	event := new(LazyClaimERC1155OwnershipTransferred)
	if err := _LazyClaimERC1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
