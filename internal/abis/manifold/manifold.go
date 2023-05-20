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

// IERC721LazyPayableClaimClaim is an auto generated low-level Go binding around an user-defined struct.
type IERC721LazyPayableClaimClaim struct {
	Total           uint32
	TotalMax        uint32
	WalletMax       uint32
	StartDate       *big.Int
	EndDate         *big.Int
	StorageProtocol uint8
	ContractVersion uint8
	Identical       bool
	MerkleRoot      [32]byte
	Location        string
	Cost            *big.Int
	PaymentReceiver common.Address
	Erc20           common.Address
}

// IERC721LazyPayableClaimClaimParameters is an auto generated low-level Go binding around an user-defined struct.
type IERC721LazyPayableClaimClaimParameters struct {
	TotalMax        uint32
	WalletMax       uint32
	StartDate       *big.Int
	EndDate         *big.Int
	StorageProtocol uint8
	Identical       bool
	MerkleRoot      [32]byte
	Location        string
	Cost            *big.Int
	PaymentReceiver common.Address
	Erc20           common.Address
}

// ManifoldMetaData contains all meta data concerning the Manifold contract.
var ManifoldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegationRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"}],\"name\":\"ClaimInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"}],\"name\":\"ClaimMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"mintCount\",\"type\":\"uint16\"}],\"name\":\"ClaimMintBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"mintCount\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintFor\",\"type\":\"address\"}],\"name\":\"ClaimMintProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"}],\"name\":\"ClaimUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_REGISTRY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MEMBERSHIP_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_FEE_MERKLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint16[]\",\"name\":\"amounts\",\"type\":\"uint16[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"approveAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"mintIndex\",\"type\":\"uint32\"}],\"name\":\"checkMintIndex\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint32[]\",\"name\":\"mintIndices\",\"type\":\"uint32[]\"}],\"name\":\"checkMintIndices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"minted\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"locationChunk\",\"type\":\"string\"}],\"name\":\"extendTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"}],\"name\":\"getClaim\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"total\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalMax\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"walletMax\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"contractVersion\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"identical\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"internalType\":\"structIERC721LazyPayableClaim.Claim\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getClaimForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"total\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalMax\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"walletMax\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"contractVersion\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"identical\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"internalType\":\"structIERC721LazyPayableClaim.Claim\",\"name\":\"claim\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"}],\"name\":\"getTotalMints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"totalMax\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"walletMax\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"identical\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"internalType\":\"structIERC721LazyPayableClaim.ClaimParameters\",\"name\":\"claimParameters\",\"type\":\"tuple\"}],\"name\":\"initializeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"mintIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"mintFor\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"mintCount\",\"type\":\"uint16\"},{\"internalType\":\"uint32[]\",\"name\":\"mintIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"merkleProofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"address\",\"name\":\"mintFor\",\"type\":\"address\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"mintCount\",\"type\":\"uint16\"},{\"internalType\":\"uint32[]\",\"name\":\"mintIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"merkleProofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"address\",\"name\":\"mintFor\",\"type\":\"address\"}],\"name\":\"mintProxy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"revokeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"membershipAddress\",\"type\":\"address\"}],\"name\":\"setMembershipAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"totalMax\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"walletMax\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"identical\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"internalType\":\"structIERC721LazyPayableClaim.ClaimParameters\",\"name\":\"claimParameters\",\"type\":\"tuple\"}],\"name\":\"updateClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"enumILazyPayableClaim.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"identical\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"updateTokenURIParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ManifoldABI is the input ABI used to generate the binding from.
// Deprecated: Use ManifoldMetaData.ABI instead.
var ManifoldABI = ManifoldMetaData.ABI

// Manifold is an auto generated Go binding around an Ethereum contract.
type Manifold struct {
	ManifoldCaller     // Read-only binding to the contract
	ManifoldTransactor // Write-only binding to the contract
	ManifoldFilterer   // Log filterer for contract events
}

// ManifoldCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManifoldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManifoldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManifoldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManifoldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManifoldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManifoldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManifoldSession struct {
	Contract     *Manifold         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManifoldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManifoldCallerSession struct {
	Contract *ManifoldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ManifoldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManifoldTransactorSession struct {
	Contract     *ManifoldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ManifoldRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManifoldRaw struct {
	Contract *Manifold // Generic contract binding to access the raw methods on
}

// ManifoldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManifoldCallerRaw struct {
	Contract *ManifoldCaller // Generic read-only contract binding to access the raw methods on
}

// ManifoldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManifoldTransactorRaw struct {
	Contract *ManifoldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManifold creates a new instance of Manifold, bound to a specific deployed contract.
func NewManifold(address common.Address, backend bind.ContractBackend) (*Manifold, error) {
	contract, err := bindManifold(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manifold{ManifoldCaller: ManifoldCaller{contract: contract}, ManifoldTransactor: ManifoldTransactor{contract: contract}, ManifoldFilterer: ManifoldFilterer{contract: contract}}, nil
}

// NewManifoldCaller creates a new read-only instance of Manifold, bound to a specific deployed contract.
func NewManifoldCaller(address common.Address, caller bind.ContractCaller) (*ManifoldCaller, error) {
	contract, err := bindManifold(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManifoldCaller{contract: contract}, nil
}

// NewManifoldTransactor creates a new write-only instance of Manifold, bound to a specific deployed contract.
func NewManifoldTransactor(address common.Address, transactor bind.ContractTransactor) (*ManifoldTransactor, error) {
	contract, err := bindManifold(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManifoldTransactor{contract: contract}, nil
}

// NewManifoldFilterer creates a new log filterer instance of Manifold, bound to a specific deployed contract.
func NewManifoldFilterer(address common.Address, filterer bind.ContractFilterer) (*ManifoldFilterer, error) {
	contract, err := bindManifold(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManifoldFilterer{contract: contract}, nil
}

// bindManifold binds a generic wrapper to an already deployed contract.
func bindManifold(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManifoldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manifold *ManifoldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manifold.Contract.ManifoldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manifold *ManifoldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manifold.Contract.ManifoldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manifold *ManifoldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manifold.Contract.ManifoldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manifold *ManifoldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manifold.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manifold *ManifoldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manifold.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manifold *ManifoldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manifold.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONREGISTRY is a free data retrieval call binding the contract method 0x4daadff7.
//
// Solidity: function DELEGATION_REGISTRY() view returns(address)
func (_Manifold *ManifoldCaller) DELEGATIONREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "DELEGATION_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DELEGATIONREGISTRY is a free data retrieval call binding the contract method 0x4daadff7.
//
// Solidity: function DELEGATION_REGISTRY() view returns(address)
func (_Manifold *ManifoldSession) DELEGATIONREGISTRY() (common.Address, error) {
	return _Manifold.Contract.DELEGATIONREGISTRY(&_Manifold.CallOpts)
}

// DELEGATIONREGISTRY is a free data retrieval call binding the contract method 0x4daadff7.
//
// Solidity: function DELEGATION_REGISTRY() view returns(address)
func (_Manifold *ManifoldCallerSession) DELEGATIONREGISTRY() (common.Address, error) {
	return _Manifold.Contract.DELEGATIONREGISTRY(&_Manifold.CallOpts)
}

// MEMBERSHIPADDRESS is a free data retrieval call binding the contract method 0x4baa62bf.
//
// Solidity: function MEMBERSHIP_ADDRESS() view returns(address)
func (_Manifold *ManifoldCaller) MEMBERSHIPADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "MEMBERSHIP_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MEMBERSHIPADDRESS is a free data retrieval call binding the contract method 0x4baa62bf.
//
// Solidity: function MEMBERSHIP_ADDRESS() view returns(address)
func (_Manifold *ManifoldSession) MEMBERSHIPADDRESS() (common.Address, error) {
	return _Manifold.Contract.MEMBERSHIPADDRESS(&_Manifold.CallOpts)
}

// MEMBERSHIPADDRESS is a free data retrieval call binding the contract method 0x4baa62bf.
//
// Solidity: function MEMBERSHIP_ADDRESS() view returns(address)
func (_Manifold *ManifoldCallerSession) MEMBERSHIPADDRESS() (common.Address, error) {
	return _Manifold.Contract.MEMBERSHIPADDRESS(&_Manifold.CallOpts)
}

// MINTFEE is a free data retrieval call binding the contract method 0xd7bf81a3.
//
// Solidity: function MINT_FEE() view returns(uint256)
func (_Manifold *ManifoldCaller) MINTFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "MINT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTFEE is a free data retrieval call binding the contract method 0xd7bf81a3.
//
// Solidity: function MINT_FEE() view returns(uint256)
func (_Manifold *ManifoldSession) MINTFEE() (*big.Int, error) {
	return _Manifold.Contract.MINTFEE(&_Manifold.CallOpts)
}

// MINTFEE is a free data retrieval call binding the contract method 0xd7bf81a3.
//
// Solidity: function MINT_FEE() view returns(uint256)
func (_Manifold *ManifoldCallerSession) MINTFEE() (*big.Int, error) {
	return _Manifold.Contract.MINTFEE(&_Manifold.CallOpts)
}

// MINTFEEMERKLE is a free data retrieval call binding the contract method 0xcb799716.
//
// Solidity: function MINT_FEE_MERKLE() view returns(uint256)
func (_Manifold *ManifoldCaller) MINTFEEMERKLE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "MINT_FEE_MERKLE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTFEEMERKLE is a free data retrieval call binding the contract method 0xcb799716.
//
// Solidity: function MINT_FEE_MERKLE() view returns(uint256)
func (_Manifold *ManifoldSession) MINTFEEMERKLE() (*big.Int, error) {
	return _Manifold.Contract.MINTFEEMERKLE(&_Manifold.CallOpts)
}

// MINTFEEMERKLE is a free data retrieval call binding the contract method 0xcb799716.
//
// Solidity: function MINT_FEE_MERKLE() view returns(uint256)
func (_Manifold *ManifoldCallerSession) MINTFEEMERKLE() (*big.Int, error) {
	return _Manifold.Contract.MINTFEEMERKLE(&_Manifold.CallOpts)
}

// CheckMintIndex is a free data retrieval call binding the contract method 0xcda08536.
//
// Solidity: function checkMintIndex(address creatorContractAddress, uint256 instanceId, uint32 mintIndex) view returns(bool)
func (_Manifold *ManifoldCaller) CheckMintIndex(opts *bind.CallOpts, creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32) (bool, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "checkMintIndex", creatorContractAddress, instanceId, mintIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMintIndex is a free data retrieval call binding the contract method 0xcda08536.
//
// Solidity: function checkMintIndex(address creatorContractAddress, uint256 instanceId, uint32 mintIndex) view returns(bool)
func (_Manifold *ManifoldSession) CheckMintIndex(creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32) (bool, error) {
	return _Manifold.Contract.CheckMintIndex(&_Manifold.CallOpts, creatorContractAddress, instanceId, mintIndex)
}

// CheckMintIndex is a free data retrieval call binding the contract method 0xcda08536.
//
// Solidity: function checkMintIndex(address creatorContractAddress, uint256 instanceId, uint32 mintIndex) view returns(bool)
func (_Manifold *ManifoldCallerSession) CheckMintIndex(creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32) (bool, error) {
	return _Manifold.Contract.CheckMintIndex(&_Manifold.CallOpts, creatorContractAddress, instanceId, mintIndex)
}

// CheckMintIndices is a free data retrieval call binding the contract method 0xf8a6137b.
//
// Solidity: function checkMintIndices(address creatorContractAddress, uint256 instanceId, uint32[] mintIndices) view returns(bool[] minted)
func (_Manifold *ManifoldCaller) CheckMintIndices(opts *bind.CallOpts, creatorContractAddress common.Address, instanceId *big.Int, mintIndices []uint32) ([]bool, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "checkMintIndices", creatorContractAddress, instanceId, mintIndices)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckMintIndices is a free data retrieval call binding the contract method 0xf8a6137b.
//
// Solidity: function checkMintIndices(address creatorContractAddress, uint256 instanceId, uint32[] mintIndices) view returns(bool[] minted)
func (_Manifold *ManifoldSession) CheckMintIndices(creatorContractAddress common.Address, instanceId *big.Int, mintIndices []uint32) ([]bool, error) {
	return _Manifold.Contract.CheckMintIndices(&_Manifold.CallOpts, creatorContractAddress, instanceId, mintIndices)
}

// CheckMintIndices is a free data retrieval call binding the contract method 0xf8a6137b.
//
// Solidity: function checkMintIndices(address creatorContractAddress, uint256 instanceId, uint32[] mintIndices) view returns(bool[] minted)
func (_Manifold *ManifoldCallerSession) CheckMintIndices(creatorContractAddress common.Address, instanceId *big.Int, mintIndices []uint32) ([]bool, error) {
	return _Manifold.Contract.CheckMintIndices(&_Manifold.CallOpts, creatorContractAddress, instanceId, mintIndices)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_Manifold *ManifoldCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_Manifold *ManifoldSession) GetAdmins() ([]common.Address, error) {
	return _Manifold.Contract.GetAdmins(&_Manifold.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_Manifold *ManifoldCallerSession) GetAdmins() ([]common.Address, error) {
	return _Manifold.Contract.GetAdmins(&_Manifold.CallOpts)
}

// GetClaim is a free data retrieval call binding the contract method 0x0f79ab39.
//
// Solidity: function getClaim(address creatorContractAddress, uint256 instanceId) view returns((uint32,uint32,uint32,uint48,uint48,uint8,uint8,bool,bytes32,string,uint256,address,address))
func (_Manifold *ManifoldCaller) GetClaim(opts *bind.CallOpts, creatorContractAddress common.Address, instanceId *big.Int) (IERC721LazyPayableClaimClaim, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "getClaim", creatorContractAddress, instanceId)

	if err != nil {
		return *new(IERC721LazyPayableClaimClaim), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721LazyPayableClaimClaim)).(*IERC721LazyPayableClaimClaim)

	return out0, err

}

// GetClaim is a free data retrieval call binding the contract method 0x0f79ab39.
//
// Solidity: function getClaim(address creatorContractAddress, uint256 instanceId) view returns((uint32,uint32,uint32,uint48,uint48,uint8,uint8,bool,bytes32,string,uint256,address,address))
func (_Manifold *ManifoldSession) GetClaim(creatorContractAddress common.Address, instanceId *big.Int) (IERC721LazyPayableClaimClaim, error) {
	return _Manifold.Contract.GetClaim(&_Manifold.CallOpts, creatorContractAddress, instanceId)
}

// GetClaim is a free data retrieval call binding the contract method 0x0f79ab39.
//
// Solidity: function getClaim(address creatorContractAddress, uint256 instanceId) view returns((uint32,uint32,uint32,uint48,uint48,uint8,uint8,bool,bytes32,string,uint256,address,address))
func (_Manifold *ManifoldCallerSession) GetClaim(creatorContractAddress common.Address, instanceId *big.Int) (IERC721LazyPayableClaimClaim, error) {
	return _Manifold.Contract.GetClaim(&_Manifold.CallOpts, creatorContractAddress, instanceId)
}

// GetClaimForToken is a free data retrieval call binding the contract method 0x895696f2.
//
// Solidity: function getClaimForToken(address creatorContractAddress, uint256 tokenId) view returns(uint256 instanceId, (uint32,uint32,uint32,uint48,uint48,uint8,uint8,bool,bytes32,string,uint256,address,address) claim)
func (_Manifold *ManifoldCaller) GetClaimForToken(opts *bind.CallOpts, creatorContractAddress common.Address, tokenId *big.Int) (struct {
	InstanceId *big.Int
	Claim      IERC721LazyPayableClaimClaim
}, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "getClaimForToken", creatorContractAddress, tokenId)

	outstruct := new(struct {
		InstanceId *big.Int
		Claim      IERC721LazyPayableClaimClaim
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InstanceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Claim = *abi.ConvertType(out[1], new(IERC721LazyPayableClaimClaim)).(*IERC721LazyPayableClaimClaim)

	return *outstruct, err

}

// GetClaimForToken is a free data retrieval call binding the contract method 0x895696f2.
//
// Solidity: function getClaimForToken(address creatorContractAddress, uint256 tokenId) view returns(uint256 instanceId, (uint32,uint32,uint32,uint48,uint48,uint8,uint8,bool,bytes32,string,uint256,address,address) claim)
func (_Manifold *ManifoldSession) GetClaimForToken(creatorContractAddress common.Address, tokenId *big.Int) (struct {
	InstanceId *big.Int
	Claim      IERC721LazyPayableClaimClaim
}, error) {
	return _Manifold.Contract.GetClaimForToken(&_Manifold.CallOpts, creatorContractAddress, tokenId)
}

// GetClaimForToken is a free data retrieval call binding the contract method 0x895696f2.
//
// Solidity: function getClaimForToken(address creatorContractAddress, uint256 tokenId) view returns(uint256 instanceId, (uint32,uint32,uint32,uint48,uint48,uint8,uint8,bool,bytes32,string,uint256,address,address) claim)
func (_Manifold *ManifoldCallerSession) GetClaimForToken(creatorContractAddress common.Address, tokenId *big.Int) (struct {
	InstanceId *big.Int
	Claim      IERC721LazyPayableClaimClaim
}, error) {
	return _Manifold.Contract.GetClaimForToken(&_Manifold.CallOpts, creatorContractAddress, tokenId)
}

// GetTotalMints is a free data retrieval call binding the contract method 0x42f3bef4.
//
// Solidity: function getTotalMints(address minter, address creatorContractAddress, uint256 instanceId) view returns(uint32)
func (_Manifold *ManifoldCaller) GetTotalMints(opts *bind.CallOpts, minter common.Address, creatorContractAddress common.Address, instanceId *big.Int) (uint32, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "getTotalMints", minter, creatorContractAddress, instanceId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTotalMints is a free data retrieval call binding the contract method 0x42f3bef4.
//
// Solidity: function getTotalMints(address minter, address creatorContractAddress, uint256 instanceId) view returns(uint32)
func (_Manifold *ManifoldSession) GetTotalMints(minter common.Address, creatorContractAddress common.Address, instanceId *big.Int) (uint32, error) {
	return _Manifold.Contract.GetTotalMints(&_Manifold.CallOpts, minter, creatorContractAddress, instanceId)
}

// GetTotalMints is a free data retrieval call binding the contract method 0x42f3bef4.
//
// Solidity: function getTotalMints(address minter, address creatorContractAddress, uint256 instanceId) view returns(uint32)
func (_Manifold *ManifoldCallerSession) GetTotalMints(minter common.Address, creatorContractAddress common.Address, instanceId *big.Int) (uint32, error) {
	return _Manifold.Contract.GetTotalMints(&_Manifold.CallOpts, minter, creatorContractAddress, instanceId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_Manifold *ManifoldCaller) IsAdmin(opts *bind.CallOpts, admin common.Address) (bool, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "isAdmin", admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_Manifold *ManifoldSession) IsAdmin(admin common.Address) (bool, error) {
	return _Manifold.Contract.IsAdmin(&_Manifold.CallOpts, admin)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_Manifold *ManifoldCallerSession) IsAdmin(admin common.Address) (bool, error) {
	return _Manifold.Contract.IsAdmin(&_Manifold.CallOpts, admin)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Manifold *ManifoldCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Manifold *ManifoldSession) Owner() (common.Address, error) {
	return _Manifold.Contract.Owner(&_Manifold.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Manifold *ManifoldCallerSession) Owner() (common.Address, error) {
	return _Manifold.Contract.Owner(&_Manifold.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Manifold *ManifoldCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Manifold *ManifoldSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Manifold.Contract.SupportsInterface(&_Manifold.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Manifold *ManifoldCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Manifold.Contract.SupportsInterface(&_Manifold.CallOpts, interfaceId)
}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address creatorContractAddress, uint256 tokenId) view returns(string uri)
func (_Manifold *ManifoldCaller) TokenURI(opts *bind.CallOpts, creatorContractAddress common.Address, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Manifold.contract.Call(opts, &out, "tokenURI", creatorContractAddress, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address creatorContractAddress, uint256 tokenId) view returns(string uri)
func (_Manifold *ManifoldSession) TokenURI(creatorContractAddress common.Address, tokenId *big.Int) (string, error) {
	return _Manifold.Contract.TokenURI(&_Manifold.CallOpts, creatorContractAddress, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address creatorContractAddress, uint256 tokenId) view returns(string uri)
func (_Manifold *ManifoldCallerSession) TokenURI(creatorContractAddress common.Address, tokenId *big.Int) (string, error) {
	return _Manifold.Contract.TokenURI(&_Manifold.CallOpts, creatorContractAddress, tokenId)
}

// Airdrop is a paid mutator transaction binding the contract method 0x5f2f5129.
//
// Solidity: function airdrop(address creatorContractAddress, uint256 instanceId, address[] recipients, uint16[] amounts) returns()
func (_Manifold *ManifoldTransactor) Airdrop(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, recipients []common.Address, amounts []uint16) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "airdrop", creatorContractAddress, instanceId, recipients, amounts)
}

// Airdrop is a paid mutator transaction binding the contract method 0x5f2f5129.
//
// Solidity: function airdrop(address creatorContractAddress, uint256 instanceId, address[] recipients, uint16[] amounts) returns()
func (_Manifold *ManifoldSession) Airdrop(creatorContractAddress common.Address, instanceId *big.Int, recipients []common.Address, amounts []uint16) (*types.Transaction, error) {
	return _Manifold.Contract.Airdrop(&_Manifold.TransactOpts, creatorContractAddress, instanceId, recipients, amounts)
}

// Airdrop is a paid mutator transaction binding the contract method 0x5f2f5129.
//
// Solidity: function airdrop(address creatorContractAddress, uint256 instanceId, address[] recipients, uint16[] amounts) returns()
func (_Manifold *ManifoldTransactorSession) Airdrop(creatorContractAddress common.Address, instanceId *big.Int, recipients []common.Address, amounts []uint16) (*types.Transaction, error) {
	return _Manifold.Contract.Airdrop(&_Manifold.TransactOpts, creatorContractAddress, instanceId, recipients, amounts)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_Manifold *ManifoldTransactor) ApproveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "approveAdmin", admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_Manifold *ManifoldSession) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.ApproveAdmin(&_Manifold.TransactOpts, admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_Manifold *ManifoldTransactorSession) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.ApproveAdmin(&_Manifold.TransactOpts, admin)
}

// ExtendTokenURI is a paid mutator transaction binding the contract method 0xb93aa86c.
//
// Solidity: function extendTokenURI(address creatorContractAddress, uint256 instanceId, string locationChunk) returns()
func (_Manifold *ManifoldTransactor) ExtendTokenURI(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, locationChunk string) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "extendTokenURI", creatorContractAddress, instanceId, locationChunk)
}

// ExtendTokenURI is a paid mutator transaction binding the contract method 0xb93aa86c.
//
// Solidity: function extendTokenURI(address creatorContractAddress, uint256 instanceId, string locationChunk) returns()
func (_Manifold *ManifoldSession) ExtendTokenURI(creatorContractAddress common.Address, instanceId *big.Int, locationChunk string) (*types.Transaction, error) {
	return _Manifold.Contract.ExtendTokenURI(&_Manifold.TransactOpts, creatorContractAddress, instanceId, locationChunk)
}

// ExtendTokenURI is a paid mutator transaction binding the contract method 0xb93aa86c.
//
// Solidity: function extendTokenURI(address creatorContractAddress, uint256 instanceId, string locationChunk) returns()
func (_Manifold *ManifoldTransactorSession) ExtendTokenURI(creatorContractAddress common.Address, instanceId *big.Int, locationChunk string) (*types.Transaction, error) {
	return _Manifold.Contract.ExtendTokenURI(&_Manifold.TransactOpts, creatorContractAddress, instanceId, locationChunk)
}

// InitializeClaim is a paid mutator transaction binding the contract method 0x975b4d95.
//
// Solidity: function initializeClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bool,bytes32,string,uint256,address,address) claimParameters) returns()
func (_Manifold *ManifoldTransactor) InitializeClaim(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC721LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "initializeClaim", creatorContractAddress, instanceId, claimParameters)
}

// InitializeClaim is a paid mutator transaction binding the contract method 0x975b4d95.
//
// Solidity: function initializeClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bool,bytes32,string,uint256,address,address) claimParameters) returns()
func (_Manifold *ManifoldSession) InitializeClaim(creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC721LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _Manifold.Contract.InitializeClaim(&_Manifold.TransactOpts, creatorContractAddress, instanceId, claimParameters)
}

// InitializeClaim is a paid mutator transaction binding the contract method 0x975b4d95.
//
// Solidity: function initializeClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bool,bytes32,string,uint256,address,address) claimParameters) returns()
func (_Manifold *ManifoldTransactorSession) InitializeClaim(creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC721LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _Manifold.Contract.InitializeClaim(&_Manifold.TransactOpts, creatorContractAddress, instanceId, claimParameters)
}

// Mint is a paid mutator transaction binding the contract method 0xfa2b068f.
//
// Solidity: function mint(address creatorContractAddress, uint256 instanceId, uint32 mintIndex, bytes32[] merkleProof, address mintFor) payable returns()
func (_Manifold *ManifoldTransactor) Mint(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32, merkleProof [][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "mint", creatorContractAddress, instanceId, mintIndex, merkleProof, mintFor)
}

// Mint is a paid mutator transaction binding the contract method 0xfa2b068f.
//
// Solidity: function mint(address creatorContractAddress, uint256 instanceId, uint32 mintIndex, bytes32[] merkleProof, address mintFor) payable returns()
func (_Manifold *ManifoldSession) Mint(creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32, merkleProof [][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.Mint(&_Manifold.TransactOpts, creatorContractAddress, instanceId, mintIndex, merkleProof, mintFor)
}

// Mint is a paid mutator transaction binding the contract method 0xfa2b068f.
//
// Solidity: function mint(address creatorContractAddress, uint256 instanceId, uint32 mintIndex, bytes32[] merkleProof, address mintFor) payable returns()
func (_Manifold *ManifoldTransactorSession) Mint(creatorContractAddress common.Address, instanceId *big.Int, mintIndex uint32, merkleProof [][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.Mint(&_Manifold.TransactOpts, creatorContractAddress, instanceId, mintIndex, merkleProof, mintFor)
}

// MintBatch is a paid mutator transaction binding the contract method 0x26c858a4.
//
// Solidity: function mintBatch(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_Manifold *ManifoldTransactor) MintBatch(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "mintBatch", creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintBatch is a paid mutator transaction binding the contract method 0x26c858a4.
//
// Solidity: function mintBatch(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_Manifold *ManifoldSession) MintBatch(creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.MintBatch(&_Manifold.TransactOpts, creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintBatch is a paid mutator transaction binding the contract method 0x26c858a4.
//
// Solidity: function mintBatch(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_Manifold *ManifoldTransactorSession) MintBatch(creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.MintBatch(&_Manifold.TransactOpts, creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintProxy is a paid mutator transaction binding the contract method 0x07591acc.
//
// Solidity: function mintProxy(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_Manifold *ManifoldTransactor) MintProxy(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "mintProxy", creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintProxy is a paid mutator transaction binding the contract method 0x07591acc.
//
// Solidity: function mintProxy(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_Manifold *ManifoldSession) MintProxy(creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.MintProxy(&_Manifold.TransactOpts, creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// MintProxy is a paid mutator transaction binding the contract method 0x07591acc.
//
// Solidity: function mintProxy(address creatorContractAddress, uint256 instanceId, uint16 mintCount, uint32[] mintIndices, bytes32[][] merkleProofs, address mintFor) payable returns()
func (_Manifold *ManifoldTransactorSession) MintProxy(creatorContractAddress common.Address, instanceId *big.Int, mintCount uint16, mintIndices []uint32, merkleProofs [][][32]byte, mintFor common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.MintProxy(&_Manifold.TransactOpts, creatorContractAddress, instanceId, mintCount, mintIndices, merkleProofs, mintFor)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manifold *ManifoldTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manifold *ManifoldSession) RenounceOwnership() (*types.Transaction, error) {
	return _Manifold.Contract.RenounceOwnership(&_Manifold.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manifold *ManifoldTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Manifold.Contract.RenounceOwnership(&_Manifold.TransactOpts)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_Manifold *ManifoldTransactor) RevokeAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "revokeAdmin", admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_Manifold *ManifoldSession) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.RevokeAdmin(&_Manifold.TransactOpts, admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_Manifold *ManifoldTransactorSession) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.RevokeAdmin(&_Manifold.TransactOpts, admin)
}

// SetMembershipAddress is a paid mutator transaction binding the contract method 0x7ab39392.
//
// Solidity: function setMembershipAddress(address membershipAddress) returns()
func (_Manifold *ManifoldTransactor) SetMembershipAddress(opts *bind.TransactOpts, membershipAddress common.Address) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "setMembershipAddress", membershipAddress)
}

// SetMembershipAddress is a paid mutator transaction binding the contract method 0x7ab39392.
//
// Solidity: function setMembershipAddress(address membershipAddress) returns()
func (_Manifold *ManifoldSession) SetMembershipAddress(membershipAddress common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.SetMembershipAddress(&_Manifold.TransactOpts, membershipAddress)
}

// SetMembershipAddress is a paid mutator transaction binding the contract method 0x7ab39392.
//
// Solidity: function setMembershipAddress(address membershipAddress) returns()
func (_Manifold *ManifoldTransactorSession) SetMembershipAddress(membershipAddress common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.SetMembershipAddress(&_Manifold.TransactOpts, membershipAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manifold *ManifoldTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manifold *ManifoldSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.TransferOwnership(&_Manifold.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manifold *ManifoldTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manifold.Contract.TransferOwnership(&_Manifold.TransactOpts, newOwner)
}

// UpdateClaim is a paid mutator transaction binding the contract method 0x9c8eb489.
//
// Solidity: function updateClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bool,bytes32,string,uint256,address,address) claimParameters) returns()
func (_Manifold *ManifoldTransactor) UpdateClaim(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC721LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "updateClaim", creatorContractAddress, instanceId, claimParameters)
}

// UpdateClaim is a paid mutator transaction binding the contract method 0x9c8eb489.
//
// Solidity: function updateClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bool,bytes32,string,uint256,address,address) claimParameters) returns()
func (_Manifold *ManifoldSession) UpdateClaim(creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC721LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _Manifold.Contract.UpdateClaim(&_Manifold.TransactOpts, creatorContractAddress, instanceId, claimParameters)
}

// UpdateClaim is a paid mutator transaction binding the contract method 0x9c8eb489.
//
// Solidity: function updateClaim(address creatorContractAddress, uint256 instanceId, (uint32,uint32,uint48,uint48,uint8,bool,bytes32,string,uint256,address,address) claimParameters) returns()
func (_Manifold *ManifoldTransactorSession) UpdateClaim(creatorContractAddress common.Address, instanceId *big.Int, claimParameters IERC721LazyPayableClaimClaimParameters) (*types.Transaction, error) {
	return _Manifold.Contract.UpdateClaim(&_Manifold.TransactOpts, creatorContractAddress, instanceId, claimParameters)
}

// UpdateTokenURIParams is a paid mutator transaction binding the contract method 0x99c14347.
//
// Solidity: function updateTokenURIParams(address creatorContractAddress, uint256 instanceId, uint8 storageProtocol, bool identical, string location) returns()
func (_Manifold *ManifoldTransactor) UpdateTokenURIParams(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, storageProtocol uint8, identical bool, location string) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "updateTokenURIParams", creatorContractAddress, instanceId, storageProtocol, identical, location)
}

// UpdateTokenURIParams is a paid mutator transaction binding the contract method 0x99c14347.
//
// Solidity: function updateTokenURIParams(address creatorContractAddress, uint256 instanceId, uint8 storageProtocol, bool identical, string location) returns()
func (_Manifold *ManifoldSession) UpdateTokenURIParams(creatorContractAddress common.Address, instanceId *big.Int, storageProtocol uint8, identical bool, location string) (*types.Transaction, error) {
	return _Manifold.Contract.UpdateTokenURIParams(&_Manifold.TransactOpts, creatorContractAddress, instanceId, storageProtocol, identical, location)
}

// UpdateTokenURIParams is a paid mutator transaction binding the contract method 0x99c14347.
//
// Solidity: function updateTokenURIParams(address creatorContractAddress, uint256 instanceId, uint8 storageProtocol, bool identical, string location) returns()
func (_Manifold *ManifoldTransactorSession) UpdateTokenURIParams(creatorContractAddress common.Address, instanceId *big.Int, storageProtocol uint8, identical bool, location string) (*types.Transaction, error) {
	return _Manifold.Contract.UpdateTokenURIParams(&_Manifold.TransactOpts, creatorContractAddress, instanceId, storageProtocol, identical, location)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_Manifold *ManifoldTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Manifold.contract.Transact(opts, "withdraw", receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_Manifold *ManifoldSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Manifold.Contract.Withdraw(&_Manifold.TransactOpts, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_Manifold *ManifoldTransactorSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Manifold.Contract.Withdraw(&_Manifold.TransactOpts, receiver, amount)
}

// ManifoldAdminApprovedIterator is returned from FilterAdminApproved and is used to iterate over the raw logs and unpacked data for AdminApproved events raised by the Manifold contract.
type ManifoldAdminApprovedIterator struct {
	Event *ManifoldAdminApproved // Event containing the contract specifics and raw log

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
func (it *ManifoldAdminApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManifoldAdminApproved)
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
		it.Event = new(ManifoldAdminApproved)
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
func (it *ManifoldAdminApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManifoldAdminApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManifoldAdminApproved represents a AdminApproved event raised by the Manifold contract.
type ManifoldAdminApproved struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminApproved is a free log retrieval operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_Manifold *ManifoldFilterer) FilterAdminApproved(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*ManifoldAdminApprovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Manifold.contract.FilterLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ManifoldAdminApprovedIterator{contract: _Manifold.contract, event: "AdminApproved", logs: logs, sub: sub}, nil
}

// WatchAdminApproved is a free log subscription operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_Manifold *ManifoldFilterer) WatchAdminApproved(opts *bind.WatchOpts, sink chan<- *ManifoldAdminApproved, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Manifold.contract.WatchLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManifoldAdminApproved)
				if err := _Manifold.contract.UnpackLog(event, "AdminApproved", log); err != nil {
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
func (_Manifold *ManifoldFilterer) ParseAdminApproved(log types.Log) (*ManifoldAdminApproved, error) {
	event := new(ManifoldAdminApproved)
	if err := _Manifold.contract.UnpackLog(event, "AdminApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManifoldAdminRevokedIterator is returned from FilterAdminRevoked and is used to iterate over the raw logs and unpacked data for AdminRevoked events raised by the Manifold contract.
type ManifoldAdminRevokedIterator struct {
	Event *ManifoldAdminRevoked // Event containing the contract specifics and raw log

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
func (it *ManifoldAdminRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManifoldAdminRevoked)
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
		it.Event = new(ManifoldAdminRevoked)
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
func (it *ManifoldAdminRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManifoldAdminRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManifoldAdminRevoked represents a AdminRevoked event raised by the Manifold contract.
type ManifoldAdminRevoked struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRevoked is a free log retrieval operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_Manifold *ManifoldFilterer) FilterAdminRevoked(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*ManifoldAdminRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Manifold.contract.FilterLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ManifoldAdminRevokedIterator{contract: _Manifold.contract, event: "AdminRevoked", logs: logs, sub: sub}, nil
}

// WatchAdminRevoked is a free log subscription operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_Manifold *ManifoldFilterer) WatchAdminRevoked(opts *bind.WatchOpts, sink chan<- *ManifoldAdminRevoked, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Manifold.contract.WatchLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManifoldAdminRevoked)
				if err := _Manifold.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
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
func (_Manifold *ManifoldFilterer) ParseAdminRevoked(log types.Log) (*ManifoldAdminRevoked, error) {
	event := new(ManifoldAdminRevoked)
	if err := _Manifold.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManifoldClaimInitializedIterator is returned from FilterClaimInitialized and is used to iterate over the raw logs and unpacked data for ClaimInitialized events raised by the Manifold contract.
type ManifoldClaimInitializedIterator struct {
	Event *ManifoldClaimInitialized // Event containing the contract specifics and raw log

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
func (it *ManifoldClaimInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManifoldClaimInitialized)
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
		it.Event = new(ManifoldClaimInitialized)
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
func (it *ManifoldClaimInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManifoldClaimInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManifoldClaimInitialized represents a ClaimInitialized event raised by the Manifold contract.
type ManifoldClaimInitialized struct {
	CreatorContract common.Address
	InstanceId      *big.Int
	Initializer     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimInitialized is a free log retrieval operation binding the contract event 0xd02727da4c6c6c111e00310108209a4de39f6817414df43ca1a10730d47c6a34.
//
// Solidity: event ClaimInitialized(address indexed creatorContract, uint256 indexed instanceId, address initializer)
func (_Manifold *ManifoldFilterer) FilterClaimInitialized(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*ManifoldClaimInitializedIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.FilterLogs(opts, "ClaimInitialized", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &ManifoldClaimInitializedIterator{contract: _Manifold.contract, event: "ClaimInitialized", logs: logs, sub: sub}, nil
}

// WatchClaimInitialized is a free log subscription operation binding the contract event 0xd02727da4c6c6c111e00310108209a4de39f6817414df43ca1a10730d47c6a34.
//
// Solidity: event ClaimInitialized(address indexed creatorContract, uint256 indexed instanceId, address initializer)
func (_Manifold *ManifoldFilterer) WatchClaimInitialized(opts *bind.WatchOpts, sink chan<- *ManifoldClaimInitialized, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.WatchLogs(opts, "ClaimInitialized", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManifoldClaimInitialized)
				if err := _Manifold.contract.UnpackLog(event, "ClaimInitialized", log); err != nil {
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
func (_Manifold *ManifoldFilterer) ParseClaimInitialized(log types.Log) (*ManifoldClaimInitialized, error) {
	event := new(ManifoldClaimInitialized)
	if err := _Manifold.contract.UnpackLog(event, "ClaimInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManifoldClaimMintIterator is returned from FilterClaimMint and is used to iterate over the raw logs and unpacked data for ClaimMint events raised by the Manifold contract.
type ManifoldClaimMintIterator struct {
	Event *ManifoldClaimMint // Event containing the contract specifics and raw log

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
func (it *ManifoldClaimMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManifoldClaimMint)
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
		it.Event = new(ManifoldClaimMint)
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
func (it *ManifoldClaimMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManifoldClaimMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManifoldClaimMint represents a ClaimMint event raised by the Manifold contract.
type ManifoldClaimMint struct {
	CreatorContract common.Address
	InstanceId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimMint is a free log retrieval operation binding the contract event 0x5d404f369772cfab2b65717fca9bc2077efeab89a0dbec036bf0c13783154eb1.
//
// Solidity: event ClaimMint(address indexed creatorContract, uint256 indexed instanceId)
func (_Manifold *ManifoldFilterer) FilterClaimMint(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*ManifoldClaimMintIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.FilterLogs(opts, "ClaimMint", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &ManifoldClaimMintIterator{contract: _Manifold.contract, event: "ClaimMint", logs: logs, sub: sub}, nil
}

// WatchClaimMint is a free log subscription operation binding the contract event 0x5d404f369772cfab2b65717fca9bc2077efeab89a0dbec036bf0c13783154eb1.
//
// Solidity: event ClaimMint(address indexed creatorContract, uint256 indexed instanceId)
func (_Manifold *ManifoldFilterer) WatchClaimMint(opts *bind.WatchOpts, sink chan<- *ManifoldClaimMint, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.WatchLogs(opts, "ClaimMint", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManifoldClaimMint)
				if err := _Manifold.contract.UnpackLog(event, "ClaimMint", log); err != nil {
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
func (_Manifold *ManifoldFilterer) ParseClaimMint(log types.Log) (*ManifoldClaimMint, error) {
	event := new(ManifoldClaimMint)
	if err := _Manifold.contract.UnpackLog(event, "ClaimMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManifoldClaimMintBatchIterator is returned from FilterClaimMintBatch and is used to iterate over the raw logs and unpacked data for ClaimMintBatch events raised by the Manifold contract.
type ManifoldClaimMintBatchIterator struct {
	Event *ManifoldClaimMintBatch // Event containing the contract specifics and raw log

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
func (it *ManifoldClaimMintBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManifoldClaimMintBatch)
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
		it.Event = new(ManifoldClaimMintBatch)
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
func (it *ManifoldClaimMintBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManifoldClaimMintBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManifoldClaimMintBatch represents a ClaimMintBatch event raised by the Manifold contract.
type ManifoldClaimMintBatch struct {
	CreatorContract common.Address
	InstanceId      *big.Int
	MintCount       uint16
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimMintBatch is a free log retrieval operation binding the contract event 0x74f5d3254dfa39a7b1217a27d5d9b3e061eafe11720eca1cf499da2dc1eb1259.
//
// Solidity: event ClaimMintBatch(address indexed creatorContract, uint256 indexed instanceId, uint16 mintCount)
func (_Manifold *ManifoldFilterer) FilterClaimMintBatch(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*ManifoldClaimMintBatchIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.FilterLogs(opts, "ClaimMintBatch", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &ManifoldClaimMintBatchIterator{contract: _Manifold.contract, event: "ClaimMintBatch", logs: logs, sub: sub}, nil
}

// WatchClaimMintBatch is a free log subscription operation binding the contract event 0x74f5d3254dfa39a7b1217a27d5d9b3e061eafe11720eca1cf499da2dc1eb1259.
//
// Solidity: event ClaimMintBatch(address indexed creatorContract, uint256 indexed instanceId, uint16 mintCount)
func (_Manifold *ManifoldFilterer) WatchClaimMintBatch(opts *bind.WatchOpts, sink chan<- *ManifoldClaimMintBatch, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.WatchLogs(opts, "ClaimMintBatch", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManifoldClaimMintBatch)
				if err := _Manifold.contract.UnpackLog(event, "ClaimMintBatch", log); err != nil {
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
func (_Manifold *ManifoldFilterer) ParseClaimMintBatch(log types.Log) (*ManifoldClaimMintBatch, error) {
	event := new(ManifoldClaimMintBatch)
	if err := _Manifold.contract.UnpackLog(event, "ClaimMintBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManifoldClaimMintProxyIterator is returned from FilterClaimMintProxy and is used to iterate over the raw logs and unpacked data for ClaimMintProxy events raised by the Manifold contract.
type ManifoldClaimMintProxyIterator struct {
	Event *ManifoldClaimMintProxy // Event containing the contract specifics and raw log

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
func (it *ManifoldClaimMintProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManifoldClaimMintProxy)
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
		it.Event = new(ManifoldClaimMintProxy)
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
func (it *ManifoldClaimMintProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManifoldClaimMintProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManifoldClaimMintProxy represents a ClaimMintProxy event raised by the Manifold contract.
type ManifoldClaimMintProxy struct {
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
func (_Manifold *ManifoldFilterer) FilterClaimMintProxy(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*ManifoldClaimMintProxyIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.FilterLogs(opts, "ClaimMintProxy", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &ManifoldClaimMintProxyIterator{contract: _Manifold.contract, event: "ClaimMintProxy", logs: logs, sub: sub}, nil
}

// WatchClaimMintProxy is a free log subscription operation binding the contract event 0x61039ad47d0b05ec206a4450fd164cc2055af66ac594c12b8dd747e8803a90de.
//
// Solidity: event ClaimMintProxy(address indexed creatorContract, uint256 indexed instanceId, uint16 mintCount, address proxy, address mintFor)
func (_Manifold *ManifoldFilterer) WatchClaimMintProxy(opts *bind.WatchOpts, sink chan<- *ManifoldClaimMintProxy, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.WatchLogs(opts, "ClaimMintProxy", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManifoldClaimMintProxy)
				if err := _Manifold.contract.UnpackLog(event, "ClaimMintProxy", log); err != nil {
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
func (_Manifold *ManifoldFilterer) ParseClaimMintProxy(log types.Log) (*ManifoldClaimMintProxy, error) {
	event := new(ManifoldClaimMintProxy)
	if err := _Manifold.contract.UnpackLog(event, "ClaimMintProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManifoldClaimUpdatedIterator is returned from FilterClaimUpdated and is used to iterate over the raw logs and unpacked data for ClaimUpdated events raised by the Manifold contract.
type ManifoldClaimUpdatedIterator struct {
	Event *ManifoldClaimUpdated // Event containing the contract specifics and raw log

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
func (it *ManifoldClaimUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManifoldClaimUpdated)
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
		it.Event = new(ManifoldClaimUpdated)
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
func (it *ManifoldClaimUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManifoldClaimUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManifoldClaimUpdated represents a ClaimUpdated event raised by the Manifold contract.
type ManifoldClaimUpdated struct {
	CreatorContract common.Address
	InstanceId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimUpdated is a free log retrieval operation binding the contract event 0x657336af9bb6c51d60c05491508d7d3026a24ee549d7a0af42e44c75bfaec47c.
//
// Solidity: event ClaimUpdated(address indexed creatorContract, uint256 indexed instanceId)
func (_Manifold *ManifoldFilterer) FilterClaimUpdated(opts *bind.FilterOpts, creatorContract []common.Address, instanceId []*big.Int) (*ManifoldClaimUpdatedIterator, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.FilterLogs(opts, "ClaimUpdated", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return &ManifoldClaimUpdatedIterator{contract: _Manifold.contract, event: "ClaimUpdated", logs: logs, sub: sub}, nil
}

// WatchClaimUpdated is a free log subscription operation binding the contract event 0x657336af9bb6c51d60c05491508d7d3026a24ee549d7a0af42e44c75bfaec47c.
//
// Solidity: event ClaimUpdated(address indexed creatorContract, uint256 indexed instanceId)
func (_Manifold *ManifoldFilterer) WatchClaimUpdated(opts *bind.WatchOpts, sink chan<- *ManifoldClaimUpdated, creatorContract []common.Address, instanceId []*big.Int) (event.Subscription, error) {

	var creatorContractRule []interface{}
	for _, creatorContractItem := range creatorContract {
		creatorContractRule = append(creatorContractRule, creatorContractItem)
	}
	var instanceIdRule []interface{}
	for _, instanceIdItem := range instanceId {
		instanceIdRule = append(instanceIdRule, instanceIdItem)
	}

	logs, sub, err := _Manifold.contract.WatchLogs(opts, "ClaimUpdated", creatorContractRule, instanceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManifoldClaimUpdated)
				if err := _Manifold.contract.UnpackLog(event, "ClaimUpdated", log); err != nil {
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
func (_Manifold *ManifoldFilterer) ParseClaimUpdated(log types.Log) (*ManifoldClaimUpdated, error) {
	event := new(ManifoldClaimUpdated)
	if err := _Manifold.contract.UnpackLog(event, "ClaimUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManifoldOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Manifold contract.
type ManifoldOwnershipTransferredIterator struct {
	Event *ManifoldOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ManifoldOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManifoldOwnershipTransferred)
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
		it.Event = new(ManifoldOwnershipTransferred)
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
func (it *ManifoldOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManifoldOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManifoldOwnershipTransferred represents a OwnershipTransferred event raised by the Manifold contract.
type ManifoldOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Manifold *ManifoldFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ManifoldOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manifold.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ManifoldOwnershipTransferredIterator{contract: _Manifold.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Manifold *ManifoldFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ManifoldOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manifold.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManifoldOwnershipTransferred)
				if err := _Manifold.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Manifold *ManifoldFilterer) ParseOwnershipTransferred(log types.Log) (*ManifoldOwnershipTransferred, error) {
	event := new(ManifoldOwnershipTransferred)
	if err := _Manifold.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
