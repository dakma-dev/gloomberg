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

// CreatorCoreERC1155MetaData contains all meta data concerning the CreatorCoreERC1155 contract.
var CreatorCoreERC1155MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"name\":\"ApproveTransferUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"DefaultRoyaltiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"ExtensionApproveTransferUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ExtensionBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ExtensionRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"ExtensionRoyaltiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ExtensionUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"permissions\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MintPermissionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"RoyaltiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"approveAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"name\":\"blacklistExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApproveTransfer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExtensions\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"extensions\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFeeBps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFeeRecipients\",\"outputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFees\",\"outputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyalties\",\"outputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mintBaseExisting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"mintBaseNew\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mintExtensionExisting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"mintExtensionNew\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"registerExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"baseURIIdentical\",\"type\":\"bool\"}],\"name\":\"registerExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"revokeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"name\":\"setApproveTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setApproveTransferExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"setBaseTokenURIExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"identical\",\"type\":\"bool\"}],\"name\":\"setBaseTokenURIExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permissions\",\"type\":\"address\"}],\"name\":\"setMintPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"setRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"setRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"setRoyaltiesExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"setTokenURIExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"setTokenURIExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"}],\"name\":\"setTokenURIPrefix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"}],\"name\":\"setTokenURIPrefixExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenExtension\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"name\":\"unregisterExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreatorCoreERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use CreatorCoreERC1155MetaData.ABI instead.
var CreatorCoreERC1155ABI = CreatorCoreERC1155MetaData.ABI

// CreatorCoreERC1155 is an auto generated Go binding around an Ethereum contract.
type CreatorCoreERC1155 struct {
	CreatorCoreERC1155Caller     // Read-only binding to the contract
	CreatorCoreERC1155Transactor // Write-only binding to the contract
	CreatorCoreERC1155Filterer   // Log filterer for contract events
}

// CreatorCoreERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type CreatorCoreERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorCoreERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CreatorCoreERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorCoreERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreatorCoreERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorCoreERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreatorCoreERC1155Session struct {
	Contract     *CreatorCoreERC1155 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CreatorCoreERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreatorCoreERC1155CallerSession struct {
	Contract *CreatorCoreERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CreatorCoreERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreatorCoreERC1155TransactorSession struct {
	Contract     *CreatorCoreERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CreatorCoreERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type CreatorCoreERC1155Raw struct {
	Contract *CreatorCoreERC1155 // Generic contract binding to access the raw methods on
}

// CreatorCoreERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreatorCoreERC1155CallerRaw struct {
	Contract *CreatorCoreERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// CreatorCoreERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreatorCoreERC1155TransactorRaw struct {
	Contract *CreatorCoreERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreatorCoreERC1155 creates a new instance of CreatorCoreERC1155, bound to a specific deployed contract.
func NewCreatorCoreERC1155(address common.Address, backend bind.ContractBackend) (*CreatorCoreERC1155, error) {
	contract, err := bindCreatorCoreERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155{CreatorCoreERC1155Caller: CreatorCoreERC1155Caller{contract: contract}, CreatorCoreERC1155Transactor: CreatorCoreERC1155Transactor{contract: contract}, CreatorCoreERC1155Filterer: CreatorCoreERC1155Filterer{contract: contract}}, nil
}

// NewCreatorCoreERC1155Caller creates a new read-only instance of CreatorCoreERC1155, bound to a specific deployed contract.
func NewCreatorCoreERC1155Caller(address common.Address, caller bind.ContractCaller) (*CreatorCoreERC1155Caller, error) {
	contract, err := bindCreatorCoreERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155Caller{contract: contract}, nil
}

// NewCreatorCoreERC1155Transactor creates a new write-only instance of CreatorCoreERC1155, bound to a specific deployed contract.
func NewCreatorCoreERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*CreatorCoreERC1155Transactor, error) {
	contract, err := bindCreatorCoreERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155Transactor{contract: contract}, nil
}

// NewCreatorCoreERC1155Filterer creates a new log filterer instance of CreatorCoreERC1155, bound to a specific deployed contract.
func NewCreatorCoreERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*CreatorCoreERC1155Filterer, error) {
	contract, err := bindCreatorCoreERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155Filterer{contract: contract}, nil
}

// bindCreatorCoreERC1155 binds a generic wrapper to an already deployed contract.
func bindCreatorCoreERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreatorCoreERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreatorCoreERC1155 *CreatorCoreERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreatorCoreERC1155.Contract.CreatorCoreERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreatorCoreERC1155 *CreatorCoreERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.CreatorCoreERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreatorCoreERC1155 *CreatorCoreERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.CreatorCoreERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreatorCoreERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) VERSION() (*big.Int, error) {
	return _CreatorCoreERC1155.Contract.VERSION(&_CreatorCoreERC1155.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) VERSION() (*big.Int, error) {
	return _CreatorCoreERC1155.Contract.VERSION(&_CreatorCoreERC1155.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _CreatorCoreERC1155.Contract.BalanceOf(&_CreatorCoreERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _CreatorCoreERC1155.Contract.BalanceOf(&_CreatorCoreERC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _CreatorCoreERC1155.Contract.BalanceOfBatch(&_CreatorCoreERC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _CreatorCoreERC1155.Contract.BalanceOfBatch(&_CreatorCoreERC1155.CallOpts, accounts, ids)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) GetAdmins() ([]common.Address, error) {
	return _CreatorCoreERC1155.Contract.GetAdmins(&_CreatorCoreERC1155.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) GetAdmins() ([]common.Address, error) {
	return _CreatorCoreERC1155.Contract.GetAdmins(&_CreatorCoreERC1155.CallOpts)
}

// GetApproveTransfer is a free data retrieval call binding the contract method 0x22f374d0.
//
// Solidity: function getApproveTransfer() view returns(address)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) GetApproveTransfer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "getApproveTransfer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproveTransfer is a free data retrieval call binding the contract method 0x22f374d0.
//
// Solidity: function getApproveTransfer() view returns(address)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) GetApproveTransfer() (common.Address, error) {
	return _CreatorCoreERC1155.Contract.GetApproveTransfer(&_CreatorCoreERC1155.CallOpts)
}

// GetApproveTransfer is a free data retrieval call binding the contract method 0x22f374d0.
//
// Solidity: function getApproveTransfer() view returns(address)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) GetApproveTransfer() (common.Address, error) {
	return _CreatorCoreERC1155.Contract.GetApproveTransfer(&_CreatorCoreERC1155.CallOpts)
}

// GetExtensions is a free data retrieval call binding the contract method 0x83b7db63.
//
// Solidity: function getExtensions() view returns(address[] extensions)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) GetExtensions(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "getExtensions")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetExtensions is a free data retrieval call binding the contract method 0x83b7db63.
//
// Solidity: function getExtensions() view returns(address[] extensions)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) GetExtensions() ([]common.Address, error) {
	return _CreatorCoreERC1155.Contract.GetExtensions(&_CreatorCoreERC1155.CallOpts)
}

// GetExtensions is a free data retrieval call binding the contract method 0x83b7db63.
//
// Solidity: function getExtensions() view returns(address[] extensions)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) GetExtensions() ([]common.Address, error) {
	return _CreatorCoreERC1155.Contract.GetExtensions(&_CreatorCoreERC1155.CallOpts)
}

// GetFeeBps is a free data retrieval call binding the contract method 0x0ebd4c7f.
//
// Solidity: function getFeeBps(uint256 tokenId) view returns(uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) GetFeeBps(opts *bind.CallOpts, tokenId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "getFeeBps", tokenId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFeeBps is a free data retrieval call binding the contract method 0x0ebd4c7f.
//
// Solidity: function getFeeBps(uint256 tokenId) view returns(uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) GetFeeBps(tokenId *big.Int) ([]*big.Int, error) {
	return _CreatorCoreERC1155.Contract.GetFeeBps(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// GetFeeBps is a free data retrieval call binding the contract method 0x0ebd4c7f.
//
// Solidity: function getFeeBps(uint256 tokenId) view returns(uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) GetFeeBps(tokenId *big.Int) ([]*big.Int, error) {
	return _CreatorCoreERC1155.Contract.GetFeeBps(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// GetFeeRecipients is a free data retrieval call binding the contract method 0xb9c4d9fb.
//
// Solidity: function getFeeRecipients(uint256 tokenId) view returns(address[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) GetFeeRecipients(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "getFeeRecipients", tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeeRecipients is a free data retrieval call binding the contract method 0xb9c4d9fb.
//
// Solidity: function getFeeRecipients(uint256 tokenId) view returns(address[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) GetFeeRecipients(tokenId *big.Int) ([]common.Address, error) {
	return _CreatorCoreERC1155.Contract.GetFeeRecipients(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// GetFeeRecipients is a free data retrieval call binding the contract method 0xb9c4d9fb.
//
// Solidity: function getFeeRecipients(uint256 tokenId) view returns(address[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) GetFeeRecipients(tokenId *big.Int) ([]common.Address, error) {
	return _CreatorCoreERC1155.Contract.GetFeeRecipients(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// GetFees is a free data retrieval call binding the contract method 0xd5a06d4c.
//
// Solidity: function getFees(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) GetFees(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "getFees", tokenId)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetFees is a free data retrieval call binding the contract method 0xd5a06d4c.
//
// Solidity: function getFees(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) GetFees(tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	return _CreatorCoreERC1155.Contract.GetFees(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// GetFees is a free data retrieval call binding the contract method 0xd5a06d4c.
//
// Solidity: function getFees(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) GetFees(tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	return _CreatorCoreERC1155.Contract.GetFees(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) GetRoyalties(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "getRoyalties", tokenId)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) GetRoyalties(tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	return _CreatorCoreERC1155.Contract.GetRoyalties(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) GetRoyalties(tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	return _CreatorCoreERC1155.Contract.GetRoyalties(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) IsAdmin(opts *bind.CallOpts, admin common.Address) (bool, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "isAdmin", admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) IsAdmin(admin common.Address) (bool, error) {
	return _CreatorCoreERC1155.Contract.IsAdmin(&_CreatorCoreERC1155.CallOpts, admin)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) IsAdmin(admin common.Address) (bool, error) {
	return _CreatorCoreERC1155.Contract.IsAdmin(&_CreatorCoreERC1155.CallOpts, admin)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _CreatorCoreERC1155.Contract.IsApprovedForAll(&_CreatorCoreERC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _CreatorCoreERC1155.Contract.IsApprovedForAll(&_CreatorCoreERC1155.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) Name() (string, error) {
	return _CreatorCoreERC1155.Contract.Name(&_CreatorCoreERC1155.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) Name() (string, error) {
	return _CreatorCoreERC1155.Contract.Name(&_CreatorCoreERC1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) Owner() (common.Address, error) {
	return _CreatorCoreERC1155.Contract.Owner(&_CreatorCoreERC1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) Owner() (common.Address, error) {
	return _CreatorCoreERC1155.Contract.Owner(&_CreatorCoreERC1155.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address, uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, value *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "royaltyInfo", tokenId, value)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address, uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) RoyaltyInfo(tokenId *big.Int, value *big.Int) (common.Address, *big.Int, error) {
	return _CreatorCoreERC1155.Contract.RoyaltyInfo(&_CreatorCoreERC1155.CallOpts, tokenId, value)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address, uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) RoyaltyInfo(tokenId *big.Int, value *big.Int) (common.Address, *big.Int, error) {
	return _CreatorCoreERC1155.Contract.RoyaltyInfo(&_CreatorCoreERC1155.CallOpts, tokenId, value)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CreatorCoreERC1155.Contract.SupportsInterface(&_CreatorCoreERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CreatorCoreERC1155.Contract.SupportsInterface(&_CreatorCoreERC1155.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) Symbol() (string, error) {
	return _CreatorCoreERC1155.Contract.Symbol(&_CreatorCoreERC1155.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) Symbol() (string, error) {
	return _CreatorCoreERC1155.Contract.Symbol(&_CreatorCoreERC1155.CallOpts)
}

// TokenExtension is a free data retrieval call binding the contract method 0x239be317.
//
// Solidity: function tokenExtension(uint256 tokenId) view returns(address extension)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) TokenExtension(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "tokenExtension", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenExtension is a free data retrieval call binding the contract method 0x239be317.
//
// Solidity: function tokenExtension(uint256 tokenId) view returns(address extension)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) TokenExtension(tokenId *big.Int) (common.Address, error) {
	return _CreatorCoreERC1155.Contract.TokenExtension(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// TokenExtension is a free data retrieval call binding the contract method 0x239be317.
//
// Solidity: function tokenExtension(uint256 tokenId) view returns(address extension)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) TokenExtension(tokenId *big.Int) (common.Address, error) {
	return _CreatorCoreERC1155.Contract.TokenExtension(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) TotalSupply(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "totalSupply", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) TotalSupply(tokenId *big.Int) (*big.Int, error) {
	return _CreatorCoreERC1155.Contract.TotalSupply(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) TotalSupply(tokenId *big.Int) (*big.Int, error) {
	return _CreatorCoreERC1155.Contract.TotalSupply(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Caller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CreatorCoreERC1155.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) Uri(tokenId *big.Int) (string, error) {
	return _CreatorCoreERC1155.Contract.Uri(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_CreatorCoreERC1155 *CreatorCoreERC1155CallerSession) Uri(tokenId *big.Int) (string, error) {
	return _CreatorCoreERC1155.Contract.Uri(&_CreatorCoreERC1155.CallOpts, tokenId)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) ApproveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "approveAdmin", admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.ApproveAdmin(&_CreatorCoreERC1155.TransactOpts, admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.ApproveAdmin(&_CreatorCoreERC1155.TransactOpts, admin)
}

// BlacklistExtension is a paid mutator transaction binding the contract method 0x02e7afb7.
//
// Solidity: function blacklistExtension(address extension) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) BlacklistExtension(opts *bind.TransactOpts, extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "blacklistExtension", extension)
}

// BlacklistExtension is a paid mutator transaction binding the contract method 0x02e7afb7.
//
// Solidity: function blacklistExtension(address extension) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) BlacklistExtension(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.BlacklistExtension(&_CreatorCoreERC1155.TransactOpts, extension)
}

// BlacklistExtension is a paid mutator transaction binding the contract method 0x02e7afb7.
//
// Solidity: function blacklistExtension(address extension) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) BlacklistExtension(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.BlacklistExtension(&_CreatorCoreERC1155.TransactOpts, extension)
}

// Burn is a paid mutator transaction binding the contract method 0x3db0f8ab.
//
// Solidity: function burn(address account, uint256[] tokenIds, uint256[] amounts) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) Burn(opts *bind.TransactOpts, account common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "burn", account, tokenIds, amounts)
}

// Burn is a paid mutator transaction binding the contract method 0x3db0f8ab.
//
// Solidity: function burn(address account, uint256[] tokenIds, uint256[] amounts) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) Burn(account common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.Burn(&_CreatorCoreERC1155.TransactOpts, account, tokenIds, amounts)
}

// Burn is a paid mutator transaction binding the contract method 0x3db0f8ab.
//
// Solidity: function burn(address account, uint256[] tokenIds, uint256[] amounts) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) Burn(account common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.Burn(&_CreatorCoreERC1155.TransactOpts, account, tokenIds, amounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "initialize", _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.Initialize(&_CreatorCoreERC1155.TransactOpts, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.Initialize(&_CreatorCoreERC1155.TransactOpts, _name, _symbol)
}

// MintBaseExisting is a paid mutator transaction binding the contract method 0x695c96e6.
//
// Solidity: function mintBaseExisting(address[] to, uint256[] tokenIds, uint256[] amounts) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) MintBaseExisting(opts *bind.TransactOpts, to []common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "mintBaseExisting", to, tokenIds, amounts)
}

// MintBaseExisting is a paid mutator transaction binding the contract method 0x695c96e6.
//
// Solidity: function mintBaseExisting(address[] to, uint256[] tokenIds, uint256[] amounts) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) MintBaseExisting(to []common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.MintBaseExisting(&_CreatorCoreERC1155.TransactOpts, to, tokenIds, amounts)
}

// MintBaseExisting is a paid mutator transaction binding the contract method 0x695c96e6.
//
// Solidity: function mintBaseExisting(address[] to, uint256[] tokenIds, uint256[] amounts) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) MintBaseExisting(to []common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.MintBaseExisting(&_CreatorCoreERC1155.TransactOpts, to, tokenIds, amounts)
}

// MintBaseNew is a paid mutator transaction binding the contract method 0xfeeb5a9a.
//
// Solidity: function mintBaseNew(address[] to, uint256[] amounts, string[] uris) returns(uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) MintBaseNew(opts *bind.TransactOpts, to []common.Address, amounts []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "mintBaseNew", to, amounts, uris)
}

// MintBaseNew is a paid mutator transaction binding the contract method 0xfeeb5a9a.
//
// Solidity: function mintBaseNew(address[] to, uint256[] amounts, string[] uris) returns(uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) MintBaseNew(to []common.Address, amounts []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.MintBaseNew(&_CreatorCoreERC1155.TransactOpts, to, amounts, uris)
}

// MintBaseNew is a paid mutator transaction binding the contract method 0xfeeb5a9a.
//
// Solidity: function mintBaseNew(address[] to, uint256[] amounts, string[] uris) returns(uint256[])
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) MintBaseNew(to []common.Address, amounts []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.MintBaseNew(&_CreatorCoreERC1155.TransactOpts, to, amounts, uris)
}

// MintExtensionExisting is a paid mutator transaction binding the contract method 0xe6c884dc.
//
// Solidity: function mintExtensionExisting(address[] to, uint256[] tokenIds, uint256[] amounts) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) MintExtensionExisting(opts *bind.TransactOpts, to []common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "mintExtensionExisting", to, tokenIds, amounts)
}

// MintExtensionExisting is a paid mutator transaction binding the contract method 0xe6c884dc.
//
// Solidity: function mintExtensionExisting(address[] to, uint256[] tokenIds, uint256[] amounts) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) MintExtensionExisting(to []common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.MintExtensionExisting(&_CreatorCoreERC1155.TransactOpts, to, tokenIds, amounts)
}

// MintExtensionExisting is a paid mutator transaction binding the contract method 0xe6c884dc.
//
// Solidity: function mintExtensionExisting(address[] to, uint256[] tokenIds, uint256[] amounts) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) MintExtensionExisting(to []common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.MintExtensionExisting(&_CreatorCoreERC1155.TransactOpts, to, tokenIds, amounts)
}

// MintExtensionNew is a paid mutator transaction binding the contract method 0x8c6e8472.
//
// Solidity: function mintExtensionNew(address[] to, uint256[] amounts, string[] uris) returns(uint256[] tokenIds)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) MintExtensionNew(opts *bind.TransactOpts, to []common.Address, amounts []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "mintExtensionNew", to, amounts, uris)
}

// MintExtensionNew is a paid mutator transaction binding the contract method 0x8c6e8472.
//
// Solidity: function mintExtensionNew(address[] to, uint256[] amounts, string[] uris) returns(uint256[] tokenIds)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) MintExtensionNew(to []common.Address, amounts []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.MintExtensionNew(&_CreatorCoreERC1155.TransactOpts, to, amounts, uris)
}

// MintExtensionNew is a paid mutator transaction binding the contract method 0x8c6e8472.
//
// Solidity: function mintExtensionNew(address[] to, uint256[] amounts, string[] uris) returns(uint256[] tokenIds)
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) MintExtensionNew(to []common.Address, amounts []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.MintExtensionNew(&_CreatorCoreERC1155.TransactOpts, to, amounts, uris)
}

// RegisterExtension is a paid mutator transaction binding the contract method 0x3071a0f9.
//
// Solidity: function registerExtension(address extension, string baseURI) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) RegisterExtension(opts *bind.TransactOpts, extension common.Address, baseURI string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "registerExtension", extension, baseURI)
}

// RegisterExtension is a paid mutator transaction binding the contract method 0x3071a0f9.
//
// Solidity: function registerExtension(address extension, string baseURI) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) RegisterExtension(extension common.Address, baseURI string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.RegisterExtension(&_CreatorCoreERC1155.TransactOpts, extension, baseURI)
}

// RegisterExtension is a paid mutator transaction binding the contract method 0x3071a0f9.
//
// Solidity: function registerExtension(address extension, string baseURI) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) RegisterExtension(extension common.Address, baseURI string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.RegisterExtension(&_CreatorCoreERC1155.TransactOpts, extension, baseURI)
}

// RegisterExtension0 is a paid mutator transaction binding the contract method 0x3f0f37f6.
//
// Solidity: function registerExtension(address extension, string baseURI, bool baseURIIdentical) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) RegisterExtension0(opts *bind.TransactOpts, extension common.Address, baseURI string, baseURIIdentical bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "registerExtension0", extension, baseURI, baseURIIdentical)
}

// RegisterExtension0 is a paid mutator transaction binding the contract method 0x3f0f37f6.
//
// Solidity: function registerExtension(address extension, string baseURI, bool baseURIIdentical) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) RegisterExtension0(extension common.Address, baseURI string, baseURIIdentical bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.RegisterExtension0(&_CreatorCoreERC1155.TransactOpts, extension, baseURI, baseURIIdentical)
}

// RegisterExtension0 is a paid mutator transaction binding the contract method 0x3f0f37f6.
//
// Solidity: function registerExtension(address extension, string baseURI, bool baseURIIdentical) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) RegisterExtension0(extension common.Address, baseURI string, baseURIIdentical bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.RegisterExtension0(&_CreatorCoreERC1155.TransactOpts, extension, baseURI, baseURIIdentical)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) RenounceOwnership() (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.RenounceOwnership(&_CreatorCoreERC1155.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.RenounceOwnership(&_CreatorCoreERC1155.TransactOpts)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) RevokeAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "revokeAdmin", admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.RevokeAdmin(&_CreatorCoreERC1155.TransactOpts, admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.RevokeAdmin(&_CreatorCoreERC1155.TransactOpts, admin)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SafeBatchTransferFrom(&_CreatorCoreERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SafeBatchTransferFrom(&_CreatorCoreERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SafeTransferFrom(&_CreatorCoreERC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SafeTransferFrom(&_CreatorCoreERC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetApprovalForAll(&_CreatorCoreERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetApprovalForAll(&_CreatorCoreERC1155.TransactOpts, operator, approved)
}

// SetApproveTransfer is a paid mutator transaction binding the contract method 0x596798ad.
//
// Solidity: function setApproveTransfer(address extension) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetApproveTransfer(opts *bind.TransactOpts, extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setApproveTransfer", extension)
}

// SetApproveTransfer is a paid mutator transaction binding the contract method 0x596798ad.
//
// Solidity: function setApproveTransfer(address extension) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetApproveTransfer(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetApproveTransfer(&_CreatorCoreERC1155.TransactOpts, extension)
}

// SetApproveTransfer is a paid mutator transaction binding the contract method 0x596798ad.
//
// Solidity: function setApproveTransfer(address extension) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetApproveTransfer(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetApproveTransfer(&_CreatorCoreERC1155.TransactOpts, extension)
}

// SetApproveTransferExtension is a paid mutator transaction binding the contract method 0xac0c8cfa.
//
// Solidity: function setApproveTransferExtension(bool enabled) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetApproveTransferExtension(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setApproveTransferExtension", enabled)
}

// SetApproveTransferExtension is a paid mutator transaction binding the contract method 0xac0c8cfa.
//
// Solidity: function setApproveTransferExtension(bool enabled) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetApproveTransferExtension(enabled bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetApproveTransferExtension(&_CreatorCoreERC1155.TransactOpts, enabled)
}

// SetApproveTransferExtension is a paid mutator transaction binding the contract method 0xac0c8cfa.
//
// Solidity: function setApproveTransferExtension(bool enabled) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetApproveTransferExtension(enabled bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetApproveTransferExtension(&_CreatorCoreERC1155.TransactOpts, enabled)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetBaseTokenURI(opts *bind.TransactOpts, uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setBaseTokenURI", uri_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetBaseTokenURI(uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetBaseTokenURI(&_CreatorCoreERC1155.TransactOpts, uri_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetBaseTokenURI(uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetBaseTokenURI(&_CreatorCoreERC1155.TransactOpts, uri_)
}

// SetBaseTokenURIExtension is a paid mutator transaction binding the contract method 0x3e6134b8.
//
// Solidity: function setBaseTokenURIExtension(string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetBaseTokenURIExtension(opts *bind.TransactOpts, uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setBaseTokenURIExtension", uri_)
}

// SetBaseTokenURIExtension is a paid mutator transaction binding the contract method 0x3e6134b8.
//
// Solidity: function setBaseTokenURIExtension(string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetBaseTokenURIExtension(uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetBaseTokenURIExtension(&_CreatorCoreERC1155.TransactOpts, uri_)
}

// SetBaseTokenURIExtension is a paid mutator transaction binding the contract method 0x3e6134b8.
//
// Solidity: function setBaseTokenURIExtension(string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetBaseTokenURIExtension(uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetBaseTokenURIExtension(&_CreatorCoreERC1155.TransactOpts, uri_)
}

// SetBaseTokenURIExtension0 is a paid mutator transaction binding the contract method 0x82dcc0c8.
//
// Solidity: function setBaseTokenURIExtension(string uri_, bool identical) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetBaseTokenURIExtension0(opts *bind.TransactOpts, uri_ string, identical bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setBaseTokenURIExtension0", uri_, identical)
}

// SetBaseTokenURIExtension0 is a paid mutator transaction binding the contract method 0x82dcc0c8.
//
// Solidity: function setBaseTokenURIExtension(string uri_, bool identical) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetBaseTokenURIExtension0(uri_ string, identical bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetBaseTokenURIExtension0(&_CreatorCoreERC1155.TransactOpts, uri_, identical)
}

// SetBaseTokenURIExtension0 is a paid mutator transaction binding the contract method 0x82dcc0c8.
//
// Solidity: function setBaseTokenURIExtension(string uri_, bool identical) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetBaseTokenURIExtension0(uri_ string, identical bool) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetBaseTokenURIExtension0(&_CreatorCoreERC1155.TransactOpts, uri_, identical)
}

// SetMintPermissions is a paid mutator transaction binding the contract method 0xf0cdc499.
//
// Solidity: function setMintPermissions(address extension, address permissions) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetMintPermissions(opts *bind.TransactOpts, extension common.Address, permissions common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setMintPermissions", extension, permissions)
}

// SetMintPermissions is a paid mutator transaction binding the contract method 0xf0cdc499.
//
// Solidity: function setMintPermissions(address extension, address permissions) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetMintPermissions(extension common.Address, permissions common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetMintPermissions(&_CreatorCoreERC1155.TransactOpts, extension, permissions)
}

// SetMintPermissions is a paid mutator transaction binding the contract method 0xf0cdc499.
//
// Solidity: function setMintPermissions(address extension, address permissions) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetMintPermissions(extension common.Address, permissions common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetMintPermissions(&_CreatorCoreERC1155.TransactOpts, extension, permissions)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x20e4afe2.
//
// Solidity: function setRoyalties(uint256 tokenId, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetRoyalties(opts *bind.TransactOpts, tokenId *big.Int, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setRoyalties", tokenId, receivers, basisPoints)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x20e4afe2.
//
// Solidity: function setRoyalties(uint256 tokenId, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetRoyalties(tokenId *big.Int, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetRoyalties(&_CreatorCoreERC1155.TransactOpts, tokenId, receivers, basisPoints)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x20e4afe2.
//
// Solidity: function setRoyalties(uint256 tokenId, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetRoyalties(tokenId *big.Int, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetRoyalties(&_CreatorCoreERC1155.TransactOpts, tokenId, receivers, basisPoints)
}

// SetRoyalties0 is a paid mutator transaction binding the contract method 0x332dd1ae.
//
// Solidity: function setRoyalties(address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetRoyalties0(opts *bind.TransactOpts, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setRoyalties0", receivers, basisPoints)
}

// SetRoyalties0 is a paid mutator transaction binding the contract method 0x332dd1ae.
//
// Solidity: function setRoyalties(address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetRoyalties0(receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetRoyalties0(&_CreatorCoreERC1155.TransactOpts, receivers, basisPoints)
}

// SetRoyalties0 is a paid mutator transaction binding the contract method 0x332dd1ae.
//
// Solidity: function setRoyalties(address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetRoyalties0(receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetRoyalties0(&_CreatorCoreERC1155.TransactOpts, receivers, basisPoints)
}

// SetRoyaltiesExtension is a paid mutator transaction binding the contract method 0xb0fe87c9.
//
// Solidity: function setRoyaltiesExtension(address extension, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetRoyaltiesExtension(opts *bind.TransactOpts, extension common.Address, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setRoyaltiesExtension", extension, receivers, basisPoints)
}

// SetRoyaltiesExtension is a paid mutator transaction binding the contract method 0xb0fe87c9.
//
// Solidity: function setRoyaltiesExtension(address extension, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetRoyaltiesExtension(extension common.Address, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetRoyaltiesExtension(&_CreatorCoreERC1155.TransactOpts, extension, receivers, basisPoints)
}

// SetRoyaltiesExtension is a paid mutator transaction binding the contract method 0xb0fe87c9.
//
// Solidity: function setRoyaltiesExtension(address extension, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetRoyaltiesExtension(extension common.Address, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetRoyaltiesExtension(&_CreatorCoreERC1155.TransactOpts, extension, receivers, basisPoints)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetTokenURI(opts *bind.TransactOpts, tokenId *big.Int, uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setTokenURI", tokenId, uri_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetTokenURI(tokenId *big.Int, uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURI(&_CreatorCoreERC1155.TransactOpts, tokenId, uri_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetTokenURI(tokenId *big.Int, uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURI(&_CreatorCoreERC1155.TransactOpts, tokenId, uri_)
}

// SetTokenURI0 is a paid mutator transaction binding the contract method 0xaafb2d44.
//
// Solidity: function setTokenURI(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetTokenURI0(opts *bind.TransactOpts, tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setTokenURI0", tokenIds, uris)
}

// SetTokenURI0 is a paid mutator transaction binding the contract method 0xaafb2d44.
//
// Solidity: function setTokenURI(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetTokenURI0(tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURI0(&_CreatorCoreERC1155.TransactOpts, tokenIds, uris)
}

// SetTokenURI0 is a paid mutator transaction binding the contract method 0xaafb2d44.
//
// Solidity: function setTokenURI(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetTokenURI0(tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURI0(&_CreatorCoreERC1155.TransactOpts, tokenIds, uris)
}

// SetTokenURIExtension is a paid mutator transaction binding the contract method 0x61e5bc6b.
//
// Solidity: function setTokenURIExtension(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetTokenURIExtension(opts *bind.TransactOpts, tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setTokenURIExtension", tokenIds, uris)
}

// SetTokenURIExtension is a paid mutator transaction binding the contract method 0x61e5bc6b.
//
// Solidity: function setTokenURIExtension(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetTokenURIExtension(tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURIExtension(&_CreatorCoreERC1155.TransactOpts, tokenIds, uris)
}

// SetTokenURIExtension is a paid mutator transaction binding the contract method 0x61e5bc6b.
//
// Solidity: function setTokenURIExtension(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetTokenURIExtension(tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURIExtension(&_CreatorCoreERC1155.TransactOpts, tokenIds, uris)
}

// SetTokenURIExtension0 is a paid mutator transaction binding the contract method 0xe92a89f6.
//
// Solidity: function setTokenURIExtension(uint256 tokenId, string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetTokenURIExtension0(opts *bind.TransactOpts, tokenId *big.Int, uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setTokenURIExtension0", tokenId, uri_)
}

// SetTokenURIExtension0 is a paid mutator transaction binding the contract method 0xe92a89f6.
//
// Solidity: function setTokenURIExtension(uint256 tokenId, string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetTokenURIExtension0(tokenId *big.Int, uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURIExtension0(&_CreatorCoreERC1155.TransactOpts, tokenId, uri_)
}

// SetTokenURIExtension0 is a paid mutator transaction binding the contract method 0xe92a89f6.
//
// Solidity: function setTokenURIExtension(uint256 tokenId, string uri_) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetTokenURIExtension0(tokenId *big.Int, uri_ string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURIExtension0(&_CreatorCoreERC1155.TransactOpts, tokenId, uri_)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetTokenURIPrefix(opts *bind.TransactOpts, prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setTokenURIPrefix", prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURIPrefix(&_CreatorCoreERC1155.TransactOpts, prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURIPrefix(&_CreatorCoreERC1155.TransactOpts, prefix)
}

// SetTokenURIPrefixExtension is a paid mutator transaction binding the contract method 0x66d1e9d0.
//
// Solidity: function setTokenURIPrefixExtension(string prefix) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) SetTokenURIPrefixExtension(opts *bind.TransactOpts, prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "setTokenURIPrefixExtension", prefix)
}

// SetTokenURIPrefixExtension is a paid mutator transaction binding the contract method 0x66d1e9d0.
//
// Solidity: function setTokenURIPrefixExtension(string prefix) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) SetTokenURIPrefixExtension(prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURIPrefixExtension(&_CreatorCoreERC1155.TransactOpts, prefix)
}

// SetTokenURIPrefixExtension is a paid mutator transaction binding the contract method 0x66d1e9d0.
//
// Solidity: function setTokenURIPrefixExtension(string prefix) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) SetTokenURIPrefixExtension(prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.SetTokenURIPrefixExtension(&_CreatorCoreERC1155.TransactOpts, prefix)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.TransferOwnership(&_CreatorCoreERC1155.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.TransferOwnership(&_CreatorCoreERC1155.TransactOpts, newOwner)
}

// UnregisterExtension is a paid mutator transaction binding the contract method 0xce8aee9d.
//
// Solidity: function unregisterExtension(address extension) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Transactor) UnregisterExtension(opts *bind.TransactOpts, extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.contract.Transact(opts, "unregisterExtension", extension)
}

// UnregisterExtension is a paid mutator transaction binding the contract method 0xce8aee9d.
//
// Solidity: function unregisterExtension(address extension) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155Session) UnregisterExtension(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.UnregisterExtension(&_CreatorCoreERC1155.TransactOpts, extension)
}

// UnregisterExtension is a paid mutator transaction binding the contract method 0xce8aee9d.
//
// Solidity: function unregisterExtension(address extension) returns()
func (_CreatorCoreERC1155 *CreatorCoreERC1155TransactorSession) UnregisterExtension(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC1155.Contract.UnregisterExtension(&_CreatorCoreERC1155.TransactOpts, extension)
}

// CreatorCoreERC1155AdminApprovedIterator is returned from FilterAdminApproved and is used to iterate over the raw logs and unpacked data for AdminApproved events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155AdminApprovedIterator struct {
	Event *CreatorCoreERC1155AdminApproved // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155AdminApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155AdminApproved)
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
		it.Event = new(CreatorCoreERC1155AdminApproved)
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
func (it *CreatorCoreERC1155AdminApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155AdminApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155AdminApproved represents a AdminApproved event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155AdminApproved struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminApproved is a free log retrieval operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterAdminApproved(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*CreatorCoreERC1155AdminApprovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155AdminApprovedIterator{contract: _CreatorCoreERC1155.contract, event: "AdminApproved", logs: logs, sub: sub}, nil
}

// WatchAdminApproved is a free log subscription operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchAdminApproved(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155AdminApproved, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155AdminApproved)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "AdminApproved", log); err != nil {
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
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseAdminApproved(log types.Log) (*CreatorCoreERC1155AdminApproved, error) {
	event := new(CreatorCoreERC1155AdminApproved)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "AdminApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155AdminRevokedIterator is returned from FilterAdminRevoked and is used to iterate over the raw logs and unpacked data for AdminRevoked events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155AdminRevokedIterator struct {
	Event *CreatorCoreERC1155AdminRevoked // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155AdminRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155AdminRevoked)
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
		it.Event = new(CreatorCoreERC1155AdminRevoked)
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
func (it *CreatorCoreERC1155AdminRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155AdminRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155AdminRevoked represents a AdminRevoked event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155AdminRevoked struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRevoked is a free log retrieval operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterAdminRevoked(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*CreatorCoreERC1155AdminRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155AdminRevokedIterator{contract: _CreatorCoreERC1155.contract, event: "AdminRevoked", logs: logs, sub: sub}, nil
}

// WatchAdminRevoked is a free log subscription operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchAdminRevoked(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155AdminRevoked, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155AdminRevoked)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
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
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseAdminRevoked(log types.Log) (*CreatorCoreERC1155AdminRevoked, error) {
	event := new(CreatorCoreERC1155AdminRevoked)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ApprovalForAllIterator struct {
	Event *CreatorCoreERC1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155ApprovalForAll)
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
		it.Event = new(CreatorCoreERC1155ApprovalForAll)
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
func (it *CreatorCoreERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155ApprovalForAll represents a ApprovalForAll event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*CreatorCoreERC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155ApprovalForAllIterator{contract: _CreatorCoreERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155ApprovalForAll)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseApprovalForAll(log types.Log) (*CreatorCoreERC1155ApprovalForAll, error) {
	event := new(CreatorCoreERC1155ApprovalForAll)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155ApproveTransferUpdatedIterator is returned from FilterApproveTransferUpdated and is used to iterate over the raw logs and unpacked data for ApproveTransferUpdated events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ApproveTransferUpdatedIterator struct {
	Event *CreatorCoreERC1155ApproveTransferUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155ApproveTransferUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155ApproveTransferUpdated)
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
		it.Event = new(CreatorCoreERC1155ApproveTransferUpdated)
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
func (it *CreatorCoreERC1155ApproveTransferUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155ApproveTransferUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155ApproveTransferUpdated represents a ApproveTransferUpdated event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ApproveTransferUpdated struct {
	Extension common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApproveTransferUpdated is a free log retrieval operation binding the contract event 0x959c0e47a2fe3cf01e237ba4892e2cc3194d77cbfb33e434e40873225d6b595f.
//
// Solidity: event ApproveTransferUpdated(address extension)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterApproveTransferUpdated(opts *bind.FilterOpts) (*CreatorCoreERC1155ApproveTransferUpdatedIterator, error) {

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "ApproveTransferUpdated")
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155ApproveTransferUpdatedIterator{contract: _CreatorCoreERC1155.contract, event: "ApproveTransferUpdated", logs: logs, sub: sub}, nil
}

// WatchApproveTransferUpdated is a free log subscription operation binding the contract event 0x959c0e47a2fe3cf01e237ba4892e2cc3194d77cbfb33e434e40873225d6b595f.
//
// Solidity: event ApproveTransferUpdated(address extension)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchApproveTransferUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155ApproveTransferUpdated) (event.Subscription, error) {

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "ApproveTransferUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155ApproveTransferUpdated)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ApproveTransferUpdated", log); err != nil {
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

// ParseApproveTransferUpdated is a log parse operation binding the contract event 0x959c0e47a2fe3cf01e237ba4892e2cc3194d77cbfb33e434e40873225d6b595f.
//
// Solidity: event ApproveTransferUpdated(address extension)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseApproveTransferUpdated(log types.Log) (*CreatorCoreERC1155ApproveTransferUpdated, error) {
	event := new(CreatorCoreERC1155ApproveTransferUpdated)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ApproveTransferUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155DefaultRoyaltiesUpdatedIterator is returned from FilterDefaultRoyaltiesUpdated and is used to iterate over the raw logs and unpacked data for DefaultRoyaltiesUpdated events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155DefaultRoyaltiesUpdatedIterator struct {
	Event *CreatorCoreERC1155DefaultRoyaltiesUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155DefaultRoyaltiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155DefaultRoyaltiesUpdated)
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
		it.Event = new(CreatorCoreERC1155DefaultRoyaltiesUpdated)
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
func (it *CreatorCoreERC1155DefaultRoyaltiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155DefaultRoyaltiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155DefaultRoyaltiesUpdated represents a DefaultRoyaltiesUpdated event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155DefaultRoyaltiesUpdated struct {
	Receivers   []common.Address
	BasisPoints []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyaltiesUpdated is a free log retrieval operation binding the contract event 0x2b6849d5976d799a5b0ca4dfd6b40a3d7afe9ea72c091fa01a958594f9a2659b.
//
// Solidity: event DefaultRoyaltiesUpdated(address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterDefaultRoyaltiesUpdated(opts *bind.FilterOpts) (*CreatorCoreERC1155DefaultRoyaltiesUpdatedIterator, error) {

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "DefaultRoyaltiesUpdated")
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155DefaultRoyaltiesUpdatedIterator{contract: _CreatorCoreERC1155.contract, event: "DefaultRoyaltiesUpdated", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyaltiesUpdated is a free log subscription operation binding the contract event 0x2b6849d5976d799a5b0ca4dfd6b40a3d7afe9ea72c091fa01a958594f9a2659b.
//
// Solidity: event DefaultRoyaltiesUpdated(address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchDefaultRoyaltiesUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155DefaultRoyaltiesUpdated) (event.Subscription, error) {

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "DefaultRoyaltiesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155DefaultRoyaltiesUpdated)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "DefaultRoyaltiesUpdated", log); err != nil {
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

// ParseDefaultRoyaltiesUpdated is a log parse operation binding the contract event 0x2b6849d5976d799a5b0ca4dfd6b40a3d7afe9ea72c091fa01a958594f9a2659b.
//
// Solidity: event DefaultRoyaltiesUpdated(address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseDefaultRoyaltiesUpdated(log types.Log) (*CreatorCoreERC1155DefaultRoyaltiesUpdated, error) {
	event := new(CreatorCoreERC1155DefaultRoyaltiesUpdated)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "DefaultRoyaltiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155ExtensionApproveTransferUpdatedIterator is returned from FilterExtensionApproveTransferUpdated and is used to iterate over the raw logs and unpacked data for ExtensionApproveTransferUpdated events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionApproveTransferUpdatedIterator struct {
	Event *CreatorCoreERC1155ExtensionApproveTransferUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155ExtensionApproveTransferUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155ExtensionApproveTransferUpdated)
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
		it.Event = new(CreatorCoreERC1155ExtensionApproveTransferUpdated)
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
func (it *CreatorCoreERC1155ExtensionApproveTransferUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155ExtensionApproveTransferUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155ExtensionApproveTransferUpdated represents a ExtensionApproveTransferUpdated event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionApproveTransferUpdated struct {
	Extension common.Address
	Enabled   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExtensionApproveTransferUpdated is a free log retrieval operation binding the contract event 0x072a7592283e2c2d1d56d21517ff6013325e0f55483f4828373ff4d98b0a1a36.
//
// Solidity: event ExtensionApproveTransferUpdated(address indexed extension, bool enabled)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterExtensionApproveTransferUpdated(opts *bind.FilterOpts, extension []common.Address) (*CreatorCoreERC1155ExtensionApproveTransferUpdatedIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "ExtensionApproveTransferUpdated", extensionRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155ExtensionApproveTransferUpdatedIterator{contract: _CreatorCoreERC1155.contract, event: "ExtensionApproveTransferUpdated", logs: logs, sub: sub}, nil
}

// WatchExtensionApproveTransferUpdated is a free log subscription operation binding the contract event 0x072a7592283e2c2d1d56d21517ff6013325e0f55483f4828373ff4d98b0a1a36.
//
// Solidity: event ExtensionApproveTransferUpdated(address indexed extension, bool enabled)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchExtensionApproveTransferUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155ExtensionApproveTransferUpdated, extension []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "ExtensionApproveTransferUpdated", extensionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155ExtensionApproveTransferUpdated)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionApproveTransferUpdated", log); err != nil {
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

// ParseExtensionApproveTransferUpdated is a log parse operation binding the contract event 0x072a7592283e2c2d1d56d21517ff6013325e0f55483f4828373ff4d98b0a1a36.
//
// Solidity: event ExtensionApproveTransferUpdated(address indexed extension, bool enabled)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseExtensionApproveTransferUpdated(log types.Log) (*CreatorCoreERC1155ExtensionApproveTransferUpdated, error) {
	event := new(CreatorCoreERC1155ExtensionApproveTransferUpdated)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionApproveTransferUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155ExtensionBlacklistedIterator is returned from FilterExtensionBlacklisted and is used to iterate over the raw logs and unpacked data for ExtensionBlacklisted events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionBlacklistedIterator struct {
	Event *CreatorCoreERC1155ExtensionBlacklisted // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155ExtensionBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155ExtensionBlacklisted)
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
		it.Event = new(CreatorCoreERC1155ExtensionBlacklisted)
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
func (it *CreatorCoreERC1155ExtensionBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155ExtensionBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155ExtensionBlacklisted represents a ExtensionBlacklisted event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionBlacklisted struct {
	Extension common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExtensionBlacklisted is a free log retrieval operation binding the contract event 0x05ac7bc5a606cd92a63365f9fda244499b9add0526b22d99937b6bd88181059c.
//
// Solidity: event ExtensionBlacklisted(address indexed extension, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterExtensionBlacklisted(opts *bind.FilterOpts, extension []common.Address, sender []common.Address) (*CreatorCoreERC1155ExtensionBlacklistedIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "ExtensionBlacklisted", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155ExtensionBlacklistedIterator{contract: _CreatorCoreERC1155.contract, event: "ExtensionBlacklisted", logs: logs, sub: sub}, nil
}

// WatchExtensionBlacklisted is a free log subscription operation binding the contract event 0x05ac7bc5a606cd92a63365f9fda244499b9add0526b22d99937b6bd88181059c.
//
// Solidity: event ExtensionBlacklisted(address indexed extension, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchExtensionBlacklisted(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155ExtensionBlacklisted, extension []common.Address, sender []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "ExtensionBlacklisted", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155ExtensionBlacklisted)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionBlacklisted", log); err != nil {
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

// ParseExtensionBlacklisted is a log parse operation binding the contract event 0x05ac7bc5a606cd92a63365f9fda244499b9add0526b22d99937b6bd88181059c.
//
// Solidity: event ExtensionBlacklisted(address indexed extension, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseExtensionBlacklisted(log types.Log) (*CreatorCoreERC1155ExtensionBlacklisted, error) {
	event := new(CreatorCoreERC1155ExtensionBlacklisted)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155ExtensionRegisteredIterator is returned from FilterExtensionRegistered and is used to iterate over the raw logs and unpacked data for ExtensionRegistered events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionRegisteredIterator struct {
	Event *CreatorCoreERC1155ExtensionRegistered // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155ExtensionRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155ExtensionRegistered)
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
		it.Event = new(CreatorCoreERC1155ExtensionRegistered)
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
func (it *CreatorCoreERC1155ExtensionRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155ExtensionRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155ExtensionRegistered represents a ExtensionRegistered event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionRegistered struct {
	Extension common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExtensionRegistered is a free log retrieval operation binding the contract event 0xd8cb8ba4086944eabf43c5535b7712015e4d4c714b24bf812c040ea5b7a3e42a.
//
// Solidity: event ExtensionRegistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterExtensionRegistered(opts *bind.FilterOpts, extension []common.Address, sender []common.Address) (*CreatorCoreERC1155ExtensionRegisteredIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "ExtensionRegistered", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155ExtensionRegisteredIterator{contract: _CreatorCoreERC1155.contract, event: "ExtensionRegistered", logs: logs, sub: sub}, nil
}

// WatchExtensionRegistered is a free log subscription operation binding the contract event 0xd8cb8ba4086944eabf43c5535b7712015e4d4c714b24bf812c040ea5b7a3e42a.
//
// Solidity: event ExtensionRegistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchExtensionRegistered(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155ExtensionRegistered, extension []common.Address, sender []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "ExtensionRegistered", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155ExtensionRegistered)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionRegistered", log); err != nil {
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

// ParseExtensionRegistered is a log parse operation binding the contract event 0xd8cb8ba4086944eabf43c5535b7712015e4d4c714b24bf812c040ea5b7a3e42a.
//
// Solidity: event ExtensionRegistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseExtensionRegistered(log types.Log) (*CreatorCoreERC1155ExtensionRegistered, error) {
	event := new(CreatorCoreERC1155ExtensionRegistered)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155ExtensionRoyaltiesUpdatedIterator is returned from FilterExtensionRoyaltiesUpdated and is used to iterate over the raw logs and unpacked data for ExtensionRoyaltiesUpdated events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionRoyaltiesUpdatedIterator struct {
	Event *CreatorCoreERC1155ExtensionRoyaltiesUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155ExtensionRoyaltiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155ExtensionRoyaltiesUpdated)
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
		it.Event = new(CreatorCoreERC1155ExtensionRoyaltiesUpdated)
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
func (it *CreatorCoreERC1155ExtensionRoyaltiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155ExtensionRoyaltiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155ExtensionRoyaltiesUpdated represents a ExtensionRoyaltiesUpdated event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionRoyaltiesUpdated struct {
	Extension   common.Address
	Receivers   []common.Address
	BasisPoints []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtensionRoyaltiesUpdated is a free log retrieval operation binding the contract event 0x535a93d2cb000582c0ebeaa9be4890ec6a287f98eb2df00c54c300612fd78d8f.
//
// Solidity: event ExtensionRoyaltiesUpdated(address indexed extension, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterExtensionRoyaltiesUpdated(opts *bind.FilterOpts, extension []common.Address) (*CreatorCoreERC1155ExtensionRoyaltiesUpdatedIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "ExtensionRoyaltiesUpdated", extensionRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155ExtensionRoyaltiesUpdatedIterator{contract: _CreatorCoreERC1155.contract, event: "ExtensionRoyaltiesUpdated", logs: logs, sub: sub}, nil
}

// WatchExtensionRoyaltiesUpdated is a free log subscription operation binding the contract event 0x535a93d2cb000582c0ebeaa9be4890ec6a287f98eb2df00c54c300612fd78d8f.
//
// Solidity: event ExtensionRoyaltiesUpdated(address indexed extension, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchExtensionRoyaltiesUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155ExtensionRoyaltiesUpdated, extension []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "ExtensionRoyaltiesUpdated", extensionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155ExtensionRoyaltiesUpdated)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionRoyaltiesUpdated", log); err != nil {
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

// ParseExtensionRoyaltiesUpdated is a log parse operation binding the contract event 0x535a93d2cb000582c0ebeaa9be4890ec6a287f98eb2df00c54c300612fd78d8f.
//
// Solidity: event ExtensionRoyaltiesUpdated(address indexed extension, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseExtensionRoyaltiesUpdated(log types.Log) (*CreatorCoreERC1155ExtensionRoyaltiesUpdated, error) {
	event := new(CreatorCoreERC1155ExtensionRoyaltiesUpdated)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionRoyaltiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155ExtensionUnregisteredIterator is returned from FilterExtensionUnregistered and is used to iterate over the raw logs and unpacked data for ExtensionUnregistered events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionUnregisteredIterator struct {
	Event *CreatorCoreERC1155ExtensionUnregistered // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155ExtensionUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155ExtensionUnregistered)
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
		it.Event = new(CreatorCoreERC1155ExtensionUnregistered)
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
func (it *CreatorCoreERC1155ExtensionUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155ExtensionUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155ExtensionUnregistered represents a ExtensionUnregistered event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155ExtensionUnregistered struct {
	Extension common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExtensionUnregistered is a free log retrieval operation binding the contract event 0xd19cf84cf0fec6bec9ddfa29c63adf83a55707c712f32c8285d6180a78901479.
//
// Solidity: event ExtensionUnregistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterExtensionUnregistered(opts *bind.FilterOpts, extension []common.Address, sender []common.Address) (*CreatorCoreERC1155ExtensionUnregisteredIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "ExtensionUnregistered", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155ExtensionUnregisteredIterator{contract: _CreatorCoreERC1155.contract, event: "ExtensionUnregistered", logs: logs, sub: sub}, nil
}

// WatchExtensionUnregistered is a free log subscription operation binding the contract event 0xd19cf84cf0fec6bec9ddfa29c63adf83a55707c712f32c8285d6180a78901479.
//
// Solidity: event ExtensionUnregistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchExtensionUnregistered(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155ExtensionUnregistered, extension []common.Address, sender []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "ExtensionUnregistered", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155ExtensionUnregistered)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionUnregistered", log); err != nil {
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

// ParseExtensionUnregistered is a log parse operation binding the contract event 0xd19cf84cf0fec6bec9ddfa29c63adf83a55707c712f32c8285d6180a78901479.
//
// Solidity: event ExtensionUnregistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseExtensionUnregistered(log types.Log) (*CreatorCoreERC1155ExtensionUnregistered, error) {
	event := new(CreatorCoreERC1155ExtensionUnregistered)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "ExtensionUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155InitializedIterator struct {
	Event *CreatorCoreERC1155Initialized // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155Initialized)
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
		it.Event = new(CreatorCoreERC1155Initialized)
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
func (it *CreatorCoreERC1155InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155Initialized represents a Initialized event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterInitialized(opts *bind.FilterOpts) (*CreatorCoreERC1155InitializedIterator, error) {

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155InitializedIterator{contract: _CreatorCoreERC1155.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155Initialized) (event.Subscription, error) {

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155Initialized)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseInitialized(log types.Log) (*CreatorCoreERC1155Initialized, error) {
	event := new(CreatorCoreERC1155Initialized)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155MintPermissionsUpdatedIterator is returned from FilterMintPermissionsUpdated and is used to iterate over the raw logs and unpacked data for MintPermissionsUpdated events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155MintPermissionsUpdatedIterator struct {
	Event *CreatorCoreERC1155MintPermissionsUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155MintPermissionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155MintPermissionsUpdated)
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
		it.Event = new(CreatorCoreERC1155MintPermissionsUpdated)
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
func (it *CreatorCoreERC1155MintPermissionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155MintPermissionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155MintPermissionsUpdated represents a MintPermissionsUpdated event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155MintPermissionsUpdated struct {
	Extension   common.Address
	Permissions common.Address
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMintPermissionsUpdated is a free log retrieval operation binding the contract event 0x6a835c4fcf7e0d398db3762332fdaa1471814ad39f1e2d6d0b3fdabf8efee3e0.
//
// Solidity: event MintPermissionsUpdated(address indexed extension, address indexed permissions, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterMintPermissionsUpdated(opts *bind.FilterOpts, extension []common.Address, permissions []common.Address, sender []common.Address) (*CreatorCoreERC1155MintPermissionsUpdatedIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var permissionsRule []interface{}
	for _, permissionsItem := range permissions {
		permissionsRule = append(permissionsRule, permissionsItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "MintPermissionsUpdated", extensionRule, permissionsRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155MintPermissionsUpdatedIterator{contract: _CreatorCoreERC1155.contract, event: "MintPermissionsUpdated", logs: logs, sub: sub}, nil
}

// WatchMintPermissionsUpdated is a free log subscription operation binding the contract event 0x6a835c4fcf7e0d398db3762332fdaa1471814ad39f1e2d6d0b3fdabf8efee3e0.
//
// Solidity: event MintPermissionsUpdated(address indexed extension, address indexed permissions, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchMintPermissionsUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155MintPermissionsUpdated, extension []common.Address, permissions []common.Address, sender []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var permissionsRule []interface{}
	for _, permissionsItem := range permissions {
		permissionsRule = append(permissionsRule, permissionsItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "MintPermissionsUpdated", extensionRule, permissionsRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155MintPermissionsUpdated)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "MintPermissionsUpdated", log); err != nil {
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

// ParseMintPermissionsUpdated is a log parse operation binding the contract event 0x6a835c4fcf7e0d398db3762332fdaa1471814ad39f1e2d6d0b3fdabf8efee3e0.
//
// Solidity: event MintPermissionsUpdated(address indexed extension, address indexed permissions, address indexed sender)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseMintPermissionsUpdated(log types.Log) (*CreatorCoreERC1155MintPermissionsUpdated, error) {
	event := new(CreatorCoreERC1155MintPermissionsUpdated)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "MintPermissionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155OwnershipTransferredIterator struct {
	Event *CreatorCoreERC1155OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155OwnershipTransferred)
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
		it.Event = new(CreatorCoreERC1155OwnershipTransferred)
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
func (it *CreatorCoreERC1155OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155OwnershipTransferred represents a OwnershipTransferred event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CreatorCoreERC1155OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155OwnershipTransferredIterator{contract: _CreatorCoreERC1155.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155OwnershipTransferred)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseOwnershipTransferred(log types.Log) (*CreatorCoreERC1155OwnershipTransferred, error) {
	event := new(CreatorCoreERC1155OwnershipTransferred)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155RoyaltiesUpdatedIterator is returned from FilterRoyaltiesUpdated and is used to iterate over the raw logs and unpacked data for RoyaltiesUpdated events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155RoyaltiesUpdatedIterator struct {
	Event *CreatorCoreERC1155RoyaltiesUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155RoyaltiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155RoyaltiesUpdated)
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
		it.Event = new(CreatorCoreERC1155RoyaltiesUpdated)
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
func (it *CreatorCoreERC1155RoyaltiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155RoyaltiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155RoyaltiesUpdated represents a RoyaltiesUpdated event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155RoyaltiesUpdated struct {
	TokenId     *big.Int
	Receivers   []common.Address
	BasisPoints []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRoyaltiesUpdated is a free log retrieval operation binding the contract event 0xabb46fe0761d77584bde75697647804ffd8113abd4d8d06bc664150395eccdee.
//
// Solidity: event RoyaltiesUpdated(uint256 indexed tokenId, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterRoyaltiesUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*CreatorCoreERC1155RoyaltiesUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "RoyaltiesUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155RoyaltiesUpdatedIterator{contract: _CreatorCoreERC1155.contract, event: "RoyaltiesUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltiesUpdated is a free log subscription operation binding the contract event 0xabb46fe0761d77584bde75697647804ffd8113abd4d8d06bc664150395eccdee.
//
// Solidity: event RoyaltiesUpdated(uint256 indexed tokenId, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchRoyaltiesUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155RoyaltiesUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "RoyaltiesUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155RoyaltiesUpdated)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "RoyaltiesUpdated", log); err != nil {
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

// ParseRoyaltiesUpdated is a log parse operation binding the contract event 0xabb46fe0761d77584bde75697647804ffd8113abd4d8d06bc664150395eccdee.
//
// Solidity: event RoyaltiesUpdated(uint256 indexed tokenId, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseRoyaltiesUpdated(log types.Log) (*CreatorCoreERC1155RoyaltiesUpdated, error) {
	event := new(CreatorCoreERC1155RoyaltiesUpdated)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "RoyaltiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155TransferBatchIterator struct {
	Event *CreatorCoreERC1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155TransferBatch)
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
		it.Event = new(CreatorCoreERC1155TransferBatch)
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
func (it *CreatorCoreERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155TransferBatch represents a TransferBatch event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*CreatorCoreERC1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155TransferBatchIterator{contract: _CreatorCoreERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155TransferBatch)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseTransferBatch(log types.Log) (*CreatorCoreERC1155TransferBatch, error) {
	event := new(CreatorCoreERC1155TransferBatch)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155TransferSingleIterator struct {
	Event *CreatorCoreERC1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155TransferSingle)
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
		it.Event = new(CreatorCoreERC1155TransferSingle)
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
func (it *CreatorCoreERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155TransferSingle represents a TransferSingle event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*CreatorCoreERC1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155TransferSingleIterator{contract: _CreatorCoreERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155TransferSingle)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseTransferSingle(log types.Log) (*CreatorCoreERC1155TransferSingle, error) {
	event := new(CreatorCoreERC1155TransferSingle)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155URIIterator struct {
	Event *CreatorCoreERC1155URI // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC1155URI)
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
		it.Event = new(CreatorCoreERC1155URI)
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
func (it *CreatorCoreERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC1155URI represents a URI event raised by the CreatorCoreERC1155 contract.
type CreatorCoreERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*CreatorCoreERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC1155URIIterator{contract: _CreatorCoreERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CreatorCoreERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC1155URI)
				if err := _CreatorCoreERC1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_CreatorCoreERC1155 *CreatorCoreERC1155Filterer) ParseURI(log types.Log) (*CreatorCoreERC1155URI, error) {
	event := new(CreatorCoreERC1155URI)
	if err := _CreatorCoreERC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
