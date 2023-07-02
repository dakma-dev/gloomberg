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

// CreatorCoreERC721MetaData contains all meta data concerning the CreatorCoreERC721 contract.
var CreatorCoreERC721MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"name\":\"ApproveTransferUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"DefaultRoyaltiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"ExtensionApproveTransferUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ExtensionBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ExtensionRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"ExtensionRoyaltiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ExtensionUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"permissions\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MintPermissionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"RoyaltiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"approveAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"name\":\"blacklistExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApproveTransfer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExtensions\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"extensions\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFeeBps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFeeRecipients\",\"outputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFees\",\"outputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyalties\",\"outputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"mintBaseBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"mintBaseBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintExtension\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"data\",\"type\":\"uint80\"}],\"name\":\"mintExtension\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintExtension\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"mintExtensionBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint80[]\",\"name\":\"data\",\"type\":\"uint80[]\"}],\"name\":\"mintExtensionBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"mintExtensionBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"registerExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"baseURIIdentical\",\"type\":\"bool\"}],\"name\":\"registerExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"revokeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"name\":\"setApproveTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setApproveTransferExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setBaseTokenURIExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"identical\",\"type\":\"bool\"}],\"name\":\"setBaseTokenURIExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permissions\",\"type\":\"address\"}],\"name\":\"setMintPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"setRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"setRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"basisPoints\",\"type\":\"uint256[]\"}],\"name\":\"setRoyaltiesExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"setTokenURIExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setTokenURIExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"}],\"name\":\"setTokenURIPrefix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"}],\"name\":\"setTokenURIPrefixExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenExtension\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"extension\",\"type\":\"address\"}],\"name\":\"unregisterExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CreatorCoreERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use CreatorCoreERC721MetaData.ABI instead.
var CreatorCoreERC721ABI = CreatorCoreERC721MetaData.ABI

// CreatorCoreERC721 is an auto generated Go binding around an Ethereum contract.
type CreatorCoreERC721 struct {
	CreatorCoreERC721Caller     // Read-only binding to the contract
	CreatorCoreERC721Transactor // Write-only binding to the contract
	CreatorCoreERC721Filterer   // Log filterer for contract events
}

// CreatorCoreERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type CreatorCoreERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorCoreERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CreatorCoreERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorCoreERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreatorCoreERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorCoreERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreatorCoreERC721Session struct {
	Contract     *CreatorCoreERC721 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CreatorCoreERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreatorCoreERC721CallerSession struct {
	Contract *CreatorCoreERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CreatorCoreERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreatorCoreERC721TransactorSession struct {
	Contract     *CreatorCoreERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CreatorCoreERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type CreatorCoreERC721Raw struct {
	Contract *CreatorCoreERC721 // Generic contract binding to access the raw methods on
}

// CreatorCoreERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreatorCoreERC721CallerRaw struct {
	Contract *CreatorCoreERC721Caller // Generic read-only contract binding to access the raw methods on
}

// CreatorCoreERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreatorCoreERC721TransactorRaw struct {
	Contract *CreatorCoreERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreatorCoreERC721 creates a new instance of CreatorCoreERC721, bound to a specific deployed contract.
func NewCreatorCoreERC721(address common.Address, backend bind.ContractBackend) (*CreatorCoreERC721, error) {
	contract, err := bindCreatorCoreERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721{CreatorCoreERC721Caller: CreatorCoreERC721Caller{contract: contract}, CreatorCoreERC721Transactor: CreatorCoreERC721Transactor{contract: contract}, CreatorCoreERC721Filterer: CreatorCoreERC721Filterer{contract: contract}}, nil
}

// NewCreatorCoreERC721Caller creates a new read-only instance of CreatorCoreERC721, bound to a specific deployed contract.
func NewCreatorCoreERC721Caller(address common.Address, caller bind.ContractCaller) (*CreatorCoreERC721Caller, error) {
	contract, err := bindCreatorCoreERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721Caller{contract: contract}, nil
}

// NewCreatorCoreERC721Transactor creates a new write-only instance of CreatorCoreERC721, bound to a specific deployed contract.
func NewCreatorCoreERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*CreatorCoreERC721Transactor, error) {
	contract, err := bindCreatorCoreERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721Transactor{contract: contract}, nil
}

// NewCreatorCoreERC721Filterer creates a new log filterer instance of CreatorCoreERC721, bound to a specific deployed contract.
func NewCreatorCoreERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*CreatorCoreERC721Filterer, error) {
	contract, err := bindCreatorCoreERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721Filterer{contract: contract}, nil
}

// bindCreatorCoreERC721 binds a generic wrapper to an already deployed contract.
func bindCreatorCoreERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreatorCoreERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreatorCoreERC721 *CreatorCoreERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreatorCoreERC721.Contract.CreatorCoreERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreatorCoreERC721 *CreatorCoreERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.CreatorCoreERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreatorCoreERC721 *CreatorCoreERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.CreatorCoreERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreatorCoreERC721 *CreatorCoreERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreatorCoreERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) VERSION() (*big.Int, error) {
	return _CreatorCoreERC721.Contract.VERSION(&_CreatorCoreERC721.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) VERSION() (*big.Int, error) {
	return _CreatorCoreERC721.Contract.VERSION(&_CreatorCoreERC721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CreatorCoreERC721.Contract.BalanceOf(&_CreatorCoreERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CreatorCoreERC721.Contract.BalanceOf(&_CreatorCoreERC721.CallOpts, owner)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) GetAdmins() ([]common.Address, error) {
	return _CreatorCoreERC721.Contract.GetAdmins(&_CreatorCoreERC721.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) GetAdmins() ([]common.Address, error) {
	return _CreatorCoreERC721.Contract.GetAdmins(&_CreatorCoreERC721.CallOpts)
}

// GetApproveTransfer is a free data retrieval call binding the contract method 0x22f374d0.
//
// Solidity: function getApproveTransfer() view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) GetApproveTransfer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "getApproveTransfer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproveTransfer is a free data retrieval call binding the contract method 0x22f374d0.
//
// Solidity: function getApproveTransfer() view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) GetApproveTransfer() (common.Address, error) {
	return _CreatorCoreERC721.Contract.GetApproveTransfer(&_CreatorCoreERC721.CallOpts)
}

// GetApproveTransfer is a free data retrieval call binding the contract method 0x22f374d0.
//
// Solidity: function getApproveTransfer() view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) GetApproveTransfer() (common.Address, error) {
	return _CreatorCoreERC721.Contract.GetApproveTransfer(&_CreatorCoreERC721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CreatorCoreERC721.Contract.GetApproved(&_CreatorCoreERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CreatorCoreERC721.Contract.GetApproved(&_CreatorCoreERC721.CallOpts, tokenId)
}

// GetExtensions is a free data retrieval call binding the contract method 0x83b7db63.
//
// Solidity: function getExtensions() view returns(address[] extensions)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) GetExtensions(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "getExtensions")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetExtensions is a free data retrieval call binding the contract method 0x83b7db63.
//
// Solidity: function getExtensions() view returns(address[] extensions)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) GetExtensions() ([]common.Address, error) {
	return _CreatorCoreERC721.Contract.GetExtensions(&_CreatorCoreERC721.CallOpts)
}

// GetExtensions is a free data retrieval call binding the contract method 0x83b7db63.
//
// Solidity: function getExtensions() view returns(address[] extensions)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) GetExtensions() ([]common.Address, error) {
	return _CreatorCoreERC721.Contract.GetExtensions(&_CreatorCoreERC721.CallOpts)
}

// GetFeeBps is a free data retrieval call binding the contract method 0x0ebd4c7f.
//
// Solidity: function getFeeBps(uint256 tokenId) view returns(uint256[])
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) GetFeeBps(opts *bind.CallOpts, tokenId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "getFeeBps", tokenId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFeeBps is a free data retrieval call binding the contract method 0x0ebd4c7f.
//
// Solidity: function getFeeBps(uint256 tokenId) view returns(uint256[])
func (_CreatorCoreERC721 *CreatorCoreERC721Session) GetFeeBps(tokenId *big.Int) ([]*big.Int, error) {
	return _CreatorCoreERC721.Contract.GetFeeBps(&_CreatorCoreERC721.CallOpts, tokenId)
}

// GetFeeBps is a free data retrieval call binding the contract method 0x0ebd4c7f.
//
// Solidity: function getFeeBps(uint256 tokenId) view returns(uint256[])
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) GetFeeBps(tokenId *big.Int) ([]*big.Int, error) {
	return _CreatorCoreERC721.Contract.GetFeeBps(&_CreatorCoreERC721.CallOpts, tokenId)
}

// GetFeeRecipients is a free data retrieval call binding the contract method 0xb9c4d9fb.
//
// Solidity: function getFeeRecipients(uint256 tokenId) view returns(address[])
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) GetFeeRecipients(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "getFeeRecipients", tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeeRecipients is a free data retrieval call binding the contract method 0xb9c4d9fb.
//
// Solidity: function getFeeRecipients(uint256 tokenId) view returns(address[])
func (_CreatorCoreERC721 *CreatorCoreERC721Session) GetFeeRecipients(tokenId *big.Int) ([]common.Address, error) {
	return _CreatorCoreERC721.Contract.GetFeeRecipients(&_CreatorCoreERC721.CallOpts, tokenId)
}

// GetFeeRecipients is a free data retrieval call binding the contract method 0xb9c4d9fb.
//
// Solidity: function getFeeRecipients(uint256 tokenId) view returns(address[])
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) GetFeeRecipients(tokenId *big.Int) ([]common.Address, error) {
	return _CreatorCoreERC721.Contract.GetFeeRecipients(&_CreatorCoreERC721.CallOpts, tokenId)
}

// GetFees is a free data retrieval call binding the contract method 0xd5a06d4c.
//
// Solidity: function getFees(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) GetFees(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "getFees", tokenId)

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
func (_CreatorCoreERC721 *CreatorCoreERC721Session) GetFees(tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	return _CreatorCoreERC721.Contract.GetFees(&_CreatorCoreERC721.CallOpts, tokenId)
}

// GetFees is a free data retrieval call binding the contract method 0xd5a06d4c.
//
// Solidity: function getFees(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) GetFees(tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	return _CreatorCoreERC721.Contract.GetFees(&_CreatorCoreERC721.CallOpts, tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) GetRoyalties(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "getRoyalties", tokenId)

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
func (_CreatorCoreERC721 *CreatorCoreERC721Session) GetRoyalties(tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	return _CreatorCoreERC721.Contract.GetRoyalties(&_CreatorCoreERC721.CallOpts, tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns(address[], uint256[])
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) GetRoyalties(tokenId *big.Int) ([]common.Address, []*big.Int, error) {
	return _CreatorCoreERC721.Contract.GetRoyalties(&_CreatorCoreERC721.CallOpts, tokenId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) IsAdmin(opts *bind.CallOpts, admin common.Address) (bool, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "isAdmin", admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) IsAdmin(admin common.Address) (bool, error) {
	return _CreatorCoreERC721.Contract.IsAdmin(&_CreatorCoreERC721.CallOpts, admin)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) IsAdmin(admin common.Address) (bool, error) {
	return _CreatorCoreERC721.Contract.IsAdmin(&_CreatorCoreERC721.CallOpts, admin)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CreatorCoreERC721.Contract.IsApprovedForAll(&_CreatorCoreERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CreatorCoreERC721.Contract.IsApprovedForAll(&_CreatorCoreERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) Name() (string, error) {
	return _CreatorCoreERC721.Contract.Name(&_CreatorCoreERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) Name() (string, error) {
	return _CreatorCoreERC721.Contract.Name(&_CreatorCoreERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) Owner() (common.Address, error) {
	return _CreatorCoreERC721.Contract.Owner(&_CreatorCoreERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) Owner() (common.Address, error) {
	return _CreatorCoreERC721.Contract.Owner(&_CreatorCoreERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CreatorCoreERC721.Contract.OwnerOf(&_CreatorCoreERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CreatorCoreERC721.Contract.OwnerOf(&_CreatorCoreERC721.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address, uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, value *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "royaltyInfo", tokenId, value)

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
func (_CreatorCoreERC721 *CreatorCoreERC721Session) RoyaltyInfo(tokenId *big.Int, value *big.Int) (common.Address, *big.Int, error) {
	return _CreatorCoreERC721.Contract.RoyaltyInfo(&_CreatorCoreERC721.CallOpts, tokenId, value)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address, uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) RoyaltyInfo(tokenId *big.Int, value *big.Int) (common.Address, *big.Int, error) {
	return _CreatorCoreERC721.Contract.RoyaltyInfo(&_CreatorCoreERC721.CallOpts, tokenId, value)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CreatorCoreERC721.Contract.SupportsInterface(&_CreatorCoreERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CreatorCoreERC721.Contract.SupportsInterface(&_CreatorCoreERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) Symbol() (string, error) {
	return _CreatorCoreERC721.Contract.Symbol(&_CreatorCoreERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) Symbol() (string, error) {
	return _CreatorCoreERC721.Contract.Symbol(&_CreatorCoreERC721.CallOpts)
}

// TokenData is a free data retrieval call binding the contract method 0xb4b5b48f.
//
// Solidity: function tokenData(uint256 tokenId) view returns(uint80)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) TokenData(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "tokenData", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenData is a free data retrieval call binding the contract method 0xb4b5b48f.
//
// Solidity: function tokenData(uint256 tokenId) view returns(uint80)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) TokenData(tokenId *big.Int) (*big.Int, error) {
	return _CreatorCoreERC721.Contract.TokenData(&_CreatorCoreERC721.CallOpts, tokenId)
}

// TokenData is a free data retrieval call binding the contract method 0xb4b5b48f.
//
// Solidity: function tokenData(uint256 tokenId) view returns(uint80)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) TokenData(tokenId *big.Int) (*big.Int, error) {
	return _CreatorCoreERC721.Contract.TokenData(&_CreatorCoreERC721.CallOpts, tokenId)
}

// TokenExtension is a free data retrieval call binding the contract method 0x239be317.
//
// Solidity: function tokenExtension(uint256 tokenId) view returns(address extension)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) TokenExtension(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "tokenExtension", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenExtension is a free data retrieval call binding the contract method 0x239be317.
//
// Solidity: function tokenExtension(uint256 tokenId) view returns(address extension)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) TokenExtension(tokenId *big.Int) (common.Address, error) {
	return _CreatorCoreERC721.Contract.TokenExtension(&_CreatorCoreERC721.CallOpts, tokenId)
}

// TokenExtension is a free data retrieval call binding the contract method 0x239be317.
//
// Solidity: function tokenExtension(uint256 tokenId) view returns(address extension)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) TokenExtension(tokenId *big.Int) (common.Address, error) {
	return _CreatorCoreERC721.Contract.TokenExtension(&_CreatorCoreERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CreatorCoreERC721 *CreatorCoreERC721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CreatorCoreERC721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _CreatorCoreERC721.Contract.TokenURI(&_CreatorCoreERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CreatorCoreERC721 *CreatorCoreERC721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CreatorCoreERC721.Contract.TokenURI(&_CreatorCoreERC721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.Approve(&_CreatorCoreERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.Approve(&_CreatorCoreERC721.TransactOpts, to, tokenId)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) ApproveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "approveAdmin", admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.ApproveAdmin(&_CreatorCoreERC721.TransactOpts, admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.ApproveAdmin(&_CreatorCoreERC721.TransactOpts, admin)
}

// BlacklistExtension is a paid mutator transaction binding the contract method 0x02e7afb7.
//
// Solidity: function blacklistExtension(address extension) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) BlacklistExtension(opts *bind.TransactOpts, extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "blacklistExtension", extension)
}

// BlacklistExtension is a paid mutator transaction binding the contract method 0x02e7afb7.
//
// Solidity: function blacklistExtension(address extension) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) BlacklistExtension(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.BlacklistExtension(&_CreatorCoreERC721.TransactOpts, extension)
}

// BlacklistExtension is a paid mutator transaction binding the contract method 0x02e7afb7.
//
// Solidity: function blacklistExtension(address extension) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) BlacklistExtension(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.BlacklistExtension(&_CreatorCoreERC721.TransactOpts, extension)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.Burn(&_CreatorCoreERC721.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.Burn(&_CreatorCoreERC721.TransactOpts, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "initialize", _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.Initialize(&_CreatorCoreERC721.TransactOpts, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.Initialize(&_CreatorCoreERC721.TransactOpts, _name, _symbol)
}

// MintBase is a paid mutator transaction binding the contract method 0x72ff03d3.
//
// Solidity: function mintBase(address to) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintBase(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintBase", to)
}

// MintBase is a paid mutator transaction binding the contract method 0x72ff03d3.
//
// Solidity: function mintBase(address to) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintBase(to common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintBase(&_CreatorCoreERC721.TransactOpts, to)
}

// MintBase is a paid mutator transaction binding the contract method 0x72ff03d3.
//
// Solidity: function mintBase(address to) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintBase(to common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintBase(&_CreatorCoreERC721.TransactOpts, to)
}

// MintBase0 is a paid mutator transaction binding the contract method 0x7884af44.
//
// Solidity: function mintBase(address to, string uri) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintBase0(opts *bind.TransactOpts, to common.Address, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintBase0", to, uri)
}

// MintBase0 is a paid mutator transaction binding the contract method 0x7884af44.
//
// Solidity: function mintBase(address to, string uri) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintBase0(to common.Address, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintBase0(&_CreatorCoreERC721.TransactOpts, to, uri)
}

// MintBase0 is a paid mutator transaction binding the contract method 0x7884af44.
//
// Solidity: function mintBase(address to, string uri) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintBase0(to common.Address, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintBase0(&_CreatorCoreERC721.TransactOpts, to, uri)
}

// MintBaseBatch is a paid mutator transaction binding the contract method 0x7aa15f16.
//
// Solidity: function mintBaseBatch(address to, string[] uris) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintBaseBatch(opts *bind.TransactOpts, to common.Address, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintBaseBatch", to, uris)
}

// MintBaseBatch is a paid mutator transaction binding the contract method 0x7aa15f16.
//
// Solidity: function mintBaseBatch(address to, string[] uris) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintBaseBatch(to common.Address, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintBaseBatch(&_CreatorCoreERC721.TransactOpts, to, uris)
}

// MintBaseBatch is a paid mutator transaction binding the contract method 0x7aa15f16.
//
// Solidity: function mintBaseBatch(address to, string[] uris) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintBaseBatch(to common.Address, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintBaseBatch(&_CreatorCoreERC721.TransactOpts, to, uris)
}

// MintBaseBatch0 is a paid mutator transaction binding the contract method 0xad2d0ddd.
//
// Solidity: function mintBaseBatch(address to, uint16 count) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintBaseBatch0(opts *bind.TransactOpts, to common.Address, count uint16) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintBaseBatch0", to, count)
}

// MintBaseBatch0 is a paid mutator transaction binding the contract method 0xad2d0ddd.
//
// Solidity: function mintBaseBatch(address to, uint16 count) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintBaseBatch0(to common.Address, count uint16) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintBaseBatch0(&_CreatorCoreERC721.TransactOpts, to, count)
}

// MintBaseBatch0 is a paid mutator transaction binding the contract method 0xad2d0ddd.
//
// Solidity: function mintBaseBatch(address to, uint16 count) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintBaseBatch0(to common.Address, count uint16) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintBaseBatch0(&_CreatorCoreERC721.TransactOpts, to, count)
}

// MintExtension is a paid mutator transaction binding the contract method 0x2928ca58.
//
// Solidity: function mintExtension(address to) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintExtension(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintExtension", to)
}

// MintExtension is a paid mutator transaction binding the contract method 0x2928ca58.
//
// Solidity: function mintExtension(address to) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintExtension(to common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtension(&_CreatorCoreERC721.TransactOpts, to)
}

// MintExtension is a paid mutator transaction binding the contract method 0x2928ca58.
//
// Solidity: function mintExtension(address to) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintExtension(to common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtension(&_CreatorCoreERC721.TransactOpts, to)
}

// MintExtension0 is a paid mutator transaction binding the contract method 0xd3973719.
//
// Solidity: function mintExtension(address to, uint80 data) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintExtension0(opts *bind.TransactOpts, to common.Address, data *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintExtension0", to, data)
}

// MintExtension0 is a paid mutator transaction binding the contract method 0xd3973719.
//
// Solidity: function mintExtension(address to, uint80 data) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintExtension0(to common.Address, data *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtension0(&_CreatorCoreERC721.TransactOpts, to, data)
}

// MintExtension0 is a paid mutator transaction binding the contract method 0xd3973719.
//
// Solidity: function mintExtension(address to, uint80 data) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintExtension0(to common.Address, data *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtension0(&_CreatorCoreERC721.TransactOpts, to, data)
}

// MintExtension1 is a paid mutator transaction binding the contract method 0xfe2e1f58.
//
// Solidity: function mintExtension(address to, string uri) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintExtension1(opts *bind.TransactOpts, to common.Address, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintExtension1", to, uri)
}

// MintExtension1 is a paid mutator transaction binding the contract method 0xfe2e1f58.
//
// Solidity: function mintExtension(address to, string uri) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintExtension1(to common.Address, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtension1(&_CreatorCoreERC721.TransactOpts, to, uri)
}

// MintExtension1 is a paid mutator transaction binding the contract method 0xfe2e1f58.
//
// Solidity: function mintExtension(address to, string uri) returns(uint256)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintExtension1(to common.Address, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtension1(&_CreatorCoreERC721.TransactOpts, to, uri)
}

// MintExtensionBatch is a paid mutator transaction binding the contract method 0x38e52e78.
//
// Solidity: function mintExtensionBatch(address to, string[] uris) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintExtensionBatch(opts *bind.TransactOpts, to common.Address, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintExtensionBatch", to, uris)
}

// MintExtensionBatch is a paid mutator transaction binding the contract method 0x38e52e78.
//
// Solidity: function mintExtensionBatch(address to, string[] uris) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintExtensionBatch(to common.Address, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtensionBatch(&_CreatorCoreERC721.TransactOpts, to, uris)
}

// MintExtensionBatch is a paid mutator transaction binding the contract method 0x38e52e78.
//
// Solidity: function mintExtensionBatch(address to, string[] uris) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintExtensionBatch(to common.Address, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtensionBatch(&_CreatorCoreERC721.TransactOpts, to, uris)
}

// MintExtensionBatch0 is a paid mutator transaction binding the contract method 0x4278330e.
//
// Solidity: function mintExtensionBatch(address to, uint80[] data) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintExtensionBatch0(opts *bind.TransactOpts, to common.Address, data []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintExtensionBatch0", to, data)
}

// MintExtensionBatch0 is a paid mutator transaction binding the contract method 0x4278330e.
//
// Solidity: function mintExtensionBatch(address to, uint80[] data) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintExtensionBatch0(to common.Address, data []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtensionBatch0(&_CreatorCoreERC721.TransactOpts, to, data)
}

// MintExtensionBatch0 is a paid mutator transaction binding the contract method 0x4278330e.
//
// Solidity: function mintExtensionBatch(address to, uint80[] data) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintExtensionBatch0(to common.Address, data []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtensionBatch0(&_CreatorCoreERC721.TransactOpts, to, data)
}

// MintExtensionBatch1 is a paid mutator transaction binding the contract method 0xe00aab4b.
//
// Solidity: function mintExtensionBatch(address to, uint16 count) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) MintExtensionBatch1(opts *bind.TransactOpts, to common.Address, count uint16) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "mintExtensionBatch1", to, count)
}

// MintExtensionBatch1 is a paid mutator transaction binding the contract method 0xe00aab4b.
//
// Solidity: function mintExtensionBatch(address to, uint16 count) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721Session) MintExtensionBatch1(to common.Address, count uint16) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtensionBatch1(&_CreatorCoreERC721.TransactOpts, to, count)
}

// MintExtensionBatch1 is a paid mutator transaction binding the contract method 0xe00aab4b.
//
// Solidity: function mintExtensionBatch(address to, uint16 count) returns(uint256[] tokenIds)
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) MintExtensionBatch1(to common.Address, count uint16) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.MintExtensionBatch1(&_CreatorCoreERC721.TransactOpts, to, count)
}

// RegisterExtension is a paid mutator transaction binding the contract method 0x3071a0f9.
//
// Solidity: function registerExtension(address extension, string baseURI) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) RegisterExtension(opts *bind.TransactOpts, extension common.Address, baseURI string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "registerExtension", extension, baseURI)
}

// RegisterExtension is a paid mutator transaction binding the contract method 0x3071a0f9.
//
// Solidity: function registerExtension(address extension, string baseURI) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) RegisterExtension(extension common.Address, baseURI string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.RegisterExtension(&_CreatorCoreERC721.TransactOpts, extension, baseURI)
}

// RegisterExtension is a paid mutator transaction binding the contract method 0x3071a0f9.
//
// Solidity: function registerExtension(address extension, string baseURI) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) RegisterExtension(extension common.Address, baseURI string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.RegisterExtension(&_CreatorCoreERC721.TransactOpts, extension, baseURI)
}

// RegisterExtension0 is a paid mutator transaction binding the contract method 0x3f0f37f6.
//
// Solidity: function registerExtension(address extension, string baseURI, bool baseURIIdentical) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) RegisterExtension0(opts *bind.TransactOpts, extension common.Address, baseURI string, baseURIIdentical bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "registerExtension0", extension, baseURI, baseURIIdentical)
}

// RegisterExtension0 is a paid mutator transaction binding the contract method 0x3f0f37f6.
//
// Solidity: function registerExtension(address extension, string baseURI, bool baseURIIdentical) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) RegisterExtension0(extension common.Address, baseURI string, baseURIIdentical bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.RegisterExtension0(&_CreatorCoreERC721.TransactOpts, extension, baseURI, baseURIIdentical)
}

// RegisterExtension0 is a paid mutator transaction binding the contract method 0x3f0f37f6.
//
// Solidity: function registerExtension(address extension, string baseURI, bool baseURIIdentical) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) RegisterExtension0(extension common.Address, baseURI string, baseURIIdentical bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.RegisterExtension0(&_CreatorCoreERC721.TransactOpts, extension, baseURI, baseURIIdentical)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) RenounceOwnership() (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.RenounceOwnership(&_CreatorCoreERC721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.RenounceOwnership(&_CreatorCoreERC721.TransactOpts)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) RevokeAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "revokeAdmin", admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.RevokeAdmin(&_CreatorCoreERC721.TransactOpts, admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.RevokeAdmin(&_CreatorCoreERC721.TransactOpts, admin)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SafeTransferFrom(&_CreatorCoreERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SafeTransferFrom(&_CreatorCoreERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SafeTransferFrom0(&_CreatorCoreERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SafeTransferFrom0(&_CreatorCoreERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetApprovalForAll(&_CreatorCoreERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetApprovalForAll(&_CreatorCoreERC721.TransactOpts, operator, approved)
}

// SetApproveTransfer is a paid mutator transaction binding the contract method 0x596798ad.
//
// Solidity: function setApproveTransfer(address extension) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetApproveTransfer(opts *bind.TransactOpts, extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setApproveTransfer", extension)
}

// SetApproveTransfer is a paid mutator transaction binding the contract method 0x596798ad.
//
// Solidity: function setApproveTransfer(address extension) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetApproveTransfer(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetApproveTransfer(&_CreatorCoreERC721.TransactOpts, extension)
}

// SetApproveTransfer is a paid mutator transaction binding the contract method 0x596798ad.
//
// Solidity: function setApproveTransfer(address extension) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetApproveTransfer(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetApproveTransfer(&_CreatorCoreERC721.TransactOpts, extension)
}

// SetApproveTransferExtension is a paid mutator transaction binding the contract method 0xac0c8cfa.
//
// Solidity: function setApproveTransferExtension(bool enabled) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetApproveTransferExtension(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setApproveTransferExtension", enabled)
}

// SetApproveTransferExtension is a paid mutator transaction binding the contract method 0xac0c8cfa.
//
// Solidity: function setApproveTransferExtension(bool enabled) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetApproveTransferExtension(enabled bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetApproveTransferExtension(&_CreatorCoreERC721.TransactOpts, enabled)
}

// SetApproveTransferExtension is a paid mutator transaction binding the contract method 0xac0c8cfa.
//
// Solidity: function setApproveTransferExtension(bool enabled) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetApproveTransferExtension(enabled bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetApproveTransferExtension(&_CreatorCoreERC721.TransactOpts, enabled)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetBaseTokenURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setBaseTokenURI", uri)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetBaseTokenURI(uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetBaseTokenURI(&_CreatorCoreERC721.TransactOpts, uri)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetBaseTokenURI(uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetBaseTokenURI(&_CreatorCoreERC721.TransactOpts, uri)
}

// SetBaseTokenURIExtension is a paid mutator transaction binding the contract method 0x3e6134b8.
//
// Solidity: function setBaseTokenURIExtension(string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetBaseTokenURIExtension(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setBaseTokenURIExtension", uri)
}

// SetBaseTokenURIExtension is a paid mutator transaction binding the contract method 0x3e6134b8.
//
// Solidity: function setBaseTokenURIExtension(string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetBaseTokenURIExtension(uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetBaseTokenURIExtension(&_CreatorCoreERC721.TransactOpts, uri)
}

// SetBaseTokenURIExtension is a paid mutator transaction binding the contract method 0x3e6134b8.
//
// Solidity: function setBaseTokenURIExtension(string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetBaseTokenURIExtension(uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetBaseTokenURIExtension(&_CreatorCoreERC721.TransactOpts, uri)
}

// SetBaseTokenURIExtension0 is a paid mutator transaction binding the contract method 0x82dcc0c8.
//
// Solidity: function setBaseTokenURIExtension(string uri, bool identical) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetBaseTokenURIExtension0(opts *bind.TransactOpts, uri string, identical bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setBaseTokenURIExtension0", uri, identical)
}

// SetBaseTokenURIExtension0 is a paid mutator transaction binding the contract method 0x82dcc0c8.
//
// Solidity: function setBaseTokenURIExtension(string uri, bool identical) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetBaseTokenURIExtension0(uri string, identical bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetBaseTokenURIExtension0(&_CreatorCoreERC721.TransactOpts, uri, identical)
}

// SetBaseTokenURIExtension0 is a paid mutator transaction binding the contract method 0x82dcc0c8.
//
// Solidity: function setBaseTokenURIExtension(string uri, bool identical) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetBaseTokenURIExtension0(uri string, identical bool) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetBaseTokenURIExtension0(&_CreatorCoreERC721.TransactOpts, uri, identical)
}

// SetMintPermissions is a paid mutator transaction binding the contract method 0xf0cdc499.
//
// Solidity: function setMintPermissions(address extension, address permissions) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetMintPermissions(opts *bind.TransactOpts, extension common.Address, permissions common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setMintPermissions", extension, permissions)
}

// SetMintPermissions is a paid mutator transaction binding the contract method 0xf0cdc499.
//
// Solidity: function setMintPermissions(address extension, address permissions) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetMintPermissions(extension common.Address, permissions common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetMintPermissions(&_CreatorCoreERC721.TransactOpts, extension, permissions)
}

// SetMintPermissions is a paid mutator transaction binding the contract method 0xf0cdc499.
//
// Solidity: function setMintPermissions(address extension, address permissions) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetMintPermissions(extension common.Address, permissions common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetMintPermissions(&_CreatorCoreERC721.TransactOpts, extension, permissions)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x20e4afe2.
//
// Solidity: function setRoyalties(uint256 tokenId, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetRoyalties(opts *bind.TransactOpts, tokenId *big.Int, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setRoyalties", tokenId, receivers, basisPoints)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x20e4afe2.
//
// Solidity: function setRoyalties(uint256 tokenId, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetRoyalties(tokenId *big.Int, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetRoyalties(&_CreatorCoreERC721.TransactOpts, tokenId, receivers, basisPoints)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x20e4afe2.
//
// Solidity: function setRoyalties(uint256 tokenId, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetRoyalties(tokenId *big.Int, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetRoyalties(&_CreatorCoreERC721.TransactOpts, tokenId, receivers, basisPoints)
}

// SetRoyalties0 is a paid mutator transaction binding the contract method 0x332dd1ae.
//
// Solidity: function setRoyalties(address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetRoyalties0(opts *bind.TransactOpts, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setRoyalties0", receivers, basisPoints)
}

// SetRoyalties0 is a paid mutator transaction binding the contract method 0x332dd1ae.
//
// Solidity: function setRoyalties(address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetRoyalties0(receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetRoyalties0(&_CreatorCoreERC721.TransactOpts, receivers, basisPoints)
}

// SetRoyalties0 is a paid mutator transaction binding the contract method 0x332dd1ae.
//
// Solidity: function setRoyalties(address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetRoyalties0(receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetRoyalties0(&_CreatorCoreERC721.TransactOpts, receivers, basisPoints)
}

// SetRoyaltiesExtension is a paid mutator transaction binding the contract method 0xb0fe87c9.
//
// Solidity: function setRoyaltiesExtension(address extension, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetRoyaltiesExtension(opts *bind.TransactOpts, extension common.Address, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setRoyaltiesExtension", extension, receivers, basisPoints)
}

// SetRoyaltiesExtension is a paid mutator transaction binding the contract method 0xb0fe87c9.
//
// Solidity: function setRoyaltiesExtension(address extension, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetRoyaltiesExtension(extension common.Address, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetRoyaltiesExtension(&_CreatorCoreERC721.TransactOpts, extension, receivers, basisPoints)
}

// SetRoyaltiesExtension is a paid mutator transaction binding the contract method 0xb0fe87c9.
//
// Solidity: function setRoyaltiesExtension(address extension, address[] receivers, uint256[] basisPoints) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetRoyaltiesExtension(extension common.Address, receivers []common.Address, basisPoints []*big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetRoyaltiesExtension(&_CreatorCoreERC721.TransactOpts, extension, receivers, basisPoints)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetTokenURI(opts *bind.TransactOpts, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setTokenURI", tokenId, uri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetTokenURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURI(&_CreatorCoreERC721.TransactOpts, tokenId, uri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetTokenURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURI(&_CreatorCoreERC721.TransactOpts, tokenId, uri)
}

// SetTokenURI0 is a paid mutator transaction binding the contract method 0xaafb2d44.
//
// Solidity: function setTokenURI(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetTokenURI0(opts *bind.TransactOpts, tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setTokenURI0", tokenIds, uris)
}

// SetTokenURI0 is a paid mutator transaction binding the contract method 0xaafb2d44.
//
// Solidity: function setTokenURI(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetTokenURI0(tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURI0(&_CreatorCoreERC721.TransactOpts, tokenIds, uris)
}

// SetTokenURI0 is a paid mutator transaction binding the contract method 0xaafb2d44.
//
// Solidity: function setTokenURI(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetTokenURI0(tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURI0(&_CreatorCoreERC721.TransactOpts, tokenIds, uris)
}

// SetTokenURIExtension is a paid mutator transaction binding the contract method 0x61e5bc6b.
//
// Solidity: function setTokenURIExtension(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetTokenURIExtension(opts *bind.TransactOpts, tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setTokenURIExtension", tokenIds, uris)
}

// SetTokenURIExtension is a paid mutator transaction binding the contract method 0x61e5bc6b.
//
// Solidity: function setTokenURIExtension(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetTokenURIExtension(tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURIExtension(&_CreatorCoreERC721.TransactOpts, tokenIds, uris)
}

// SetTokenURIExtension is a paid mutator transaction binding the contract method 0x61e5bc6b.
//
// Solidity: function setTokenURIExtension(uint256[] tokenIds, string[] uris) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetTokenURIExtension(tokenIds []*big.Int, uris []string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURIExtension(&_CreatorCoreERC721.TransactOpts, tokenIds, uris)
}

// SetTokenURIExtension0 is a paid mutator transaction binding the contract method 0xe92a89f6.
//
// Solidity: function setTokenURIExtension(uint256 tokenId, string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetTokenURIExtension0(opts *bind.TransactOpts, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setTokenURIExtension0", tokenId, uri)
}

// SetTokenURIExtension0 is a paid mutator transaction binding the contract method 0xe92a89f6.
//
// Solidity: function setTokenURIExtension(uint256 tokenId, string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetTokenURIExtension0(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURIExtension0(&_CreatorCoreERC721.TransactOpts, tokenId, uri)
}

// SetTokenURIExtension0 is a paid mutator transaction binding the contract method 0xe92a89f6.
//
// Solidity: function setTokenURIExtension(uint256 tokenId, string uri) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetTokenURIExtension0(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURIExtension0(&_CreatorCoreERC721.TransactOpts, tokenId, uri)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetTokenURIPrefix(opts *bind.TransactOpts, prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setTokenURIPrefix", prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURIPrefix(&_CreatorCoreERC721.TransactOpts, prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURIPrefix(&_CreatorCoreERC721.TransactOpts, prefix)
}

// SetTokenURIPrefixExtension is a paid mutator transaction binding the contract method 0x66d1e9d0.
//
// Solidity: function setTokenURIPrefixExtension(string prefix) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) SetTokenURIPrefixExtension(opts *bind.TransactOpts, prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "setTokenURIPrefixExtension", prefix)
}

// SetTokenURIPrefixExtension is a paid mutator transaction binding the contract method 0x66d1e9d0.
//
// Solidity: function setTokenURIPrefixExtension(string prefix) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) SetTokenURIPrefixExtension(prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURIPrefixExtension(&_CreatorCoreERC721.TransactOpts, prefix)
}

// SetTokenURIPrefixExtension is a paid mutator transaction binding the contract method 0x66d1e9d0.
//
// Solidity: function setTokenURIPrefixExtension(string prefix) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) SetTokenURIPrefixExtension(prefix string) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.SetTokenURIPrefixExtension(&_CreatorCoreERC721.TransactOpts, prefix)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.TransferFrom(&_CreatorCoreERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.TransferFrom(&_CreatorCoreERC721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.TransferOwnership(&_CreatorCoreERC721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.TransferOwnership(&_CreatorCoreERC721.TransactOpts, newOwner)
}

// UnregisterExtension is a paid mutator transaction binding the contract method 0xce8aee9d.
//
// Solidity: function unregisterExtension(address extension) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Transactor) UnregisterExtension(opts *bind.TransactOpts, extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.contract.Transact(opts, "unregisterExtension", extension)
}

// UnregisterExtension is a paid mutator transaction binding the contract method 0xce8aee9d.
//
// Solidity: function unregisterExtension(address extension) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721Session) UnregisterExtension(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.UnregisterExtension(&_CreatorCoreERC721.TransactOpts, extension)
}

// UnregisterExtension is a paid mutator transaction binding the contract method 0xce8aee9d.
//
// Solidity: function unregisterExtension(address extension) returns()
func (_CreatorCoreERC721 *CreatorCoreERC721TransactorSession) UnregisterExtension(extension common.Address) (*types.Transaction, error) {
	return _CreatorCoreERC721.Contract.UnregisterExtension(&_CreatorCoreERC721.TransactOpts, extension)
}

// CreatorCoreERC721AdminApprovedIterator is returned from FilterAdminApproved and is used to iterate over the raw logs and unpacked data for AdminApproved events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721AdminApprovedIterator struct {
	Event *CreatorCoreERC721AdminApproved // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721AdminApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721AdminApproved)
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
		it.Event = new(CreatorCoreERC721AdminApproved)
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
func (it *CreatorCoreERC721AdminApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721AdminApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721AdminApproved represents a AdminApproved event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721AdminApproved struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminApproved is a free log retrieval operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterAdminApproved(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*CreatorCoreERC721AdminApprovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721AdminApprovedIterator{contract: _CreatorCoreERC721.contract, event: "AdminApproved", logs: logs, sub: sub}, nil
}

// WatchAdminApproved is a free log subscription operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchAdminApproved(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721AdminApproved, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721AdminApproved)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "AdminApproved", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseAdminApproved(log types.Log) (*CreatorCoreERC721AdminApproved, error) {
	event := new(CreatorCoreERC721AdminApproved)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "AdminApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721AdminRevokedIterator is returned from FilterAdminRevoked and is used to iterate over the raw logs and unpacked data for AdminRevoked events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721AdminRevokedIterator struct {
	Event *CreatorCoreERC721AdminRevoked // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721AdminRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721AdminRevoked)
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
		it.Event = new(CreatorCoreERC721AdminRevoked)
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
func (it *CreatorCoreERC721AdminRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721AdminRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721AdminRevoked represents a AdminRevoked event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721AdminRevoked struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRevoked is a free log retrieval operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterAdminRevoked(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*CreatorCoreERC721AdminRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721AdminRevokedIterator{contract: _CreatorCoreERC721.contract, event: "AdminRevoked", logs: logs, sub: sub}, nil
}

// WatchAdminRevoked is a free log subscription operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchAdminRevoked(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721AdminRevoked, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721AdminRevoked)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseAdminRevoked(log types.Log) (*CreatorCoreERC721AdminRevoked, error) {
	event := new(CreatorCoreERC721AdminRevoked)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ApprovalIterator struct {
	Event *CreatorCoreERC721Approval // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721Approval)
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
		it.Event = new(CreatorCoreERC721Approval)
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
func (it *CreatorCoreERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721Approval represents a Approval event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CreatorCoreERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721ApprovalIterator{contract: _CreatorCoreERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721Approval)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseApproval(log types.Log) (*CreatorCoreERC721Approval, error) {
	event := new(CreatorCoreERC721Approval)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ApprovalForAllIterator struct {
	Event *CreatorCoreERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721ApprovalForAll)
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
		it.Event = new(CreatorCoreERC721ApprovalForAll)
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
func (it *CreatorCoreERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721ApprovalForAll represents a ApprovalForAll event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CreatorCoreERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721ApprovalForAllIterator{contract: _CreatorCoreERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721ApprovalForAll)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseApprovalForAll(log types.Log) (*CreatorCoreERC721ApprovalForAll, error) {
	event := new(CreatorCoreERC721ApprovalForAll)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721ApproveTransferUpdatedIterator is returned from FilterApproveTransferUpdated and is used to iterate over the raw logs and unpacked data for ApproveTransferUpdated events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ApproveTransferUpdatedIterator struct {
	Event *CreatorCoreERC721ApproveTransferUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721ApproveTransferUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721ApproveTransferUpdated)
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
		it.Event = new(CreatorCoreERC721ApproveTransferUpdated)
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
func (it *CreatorCoreERC721ApproveTransferUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721ApproveTransferUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721ApproveTransferUpdated represents a ApproveTransferUpdated event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ApproveTransferUpdated struct {
	Extension common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApproveTransferUpdated is a free log retrieval operation binding the contract event 0x959c0e47a2fe3cf01e237ba4892e2cc3194d77cbfb33e434e40873225d6b595f.
//
// Solidity: event ApproveTransferUpdated(address extension)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterApproveTransferUpdated(opts *bind.FilterOpts) (*CreatorCoreERC721ApproveTransferUpdatedIterator, error) {

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "ApproveTransferUpdated")
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721ApproveTransferUpdatedIterator{contract: _CreatorCoreERC721.contract, event: "ApproveTransferUpdated", logs: logs, sub: sub}, nil
}

// WatchApproveTransferUpdated is a free log subscription operation binding the contract event 0x959c0e47a2fe3cf01e237ba4892e2cc3194d77cbfb33e434e40873225d6b595f.
//
// Solidity: event ApproveTransferUpdated(address extension)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchApproveTransferUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721ApproveTransferUpdated) (event.Subscription, error) {

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "ApproveTransferUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721ApproveTransferUpdated)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "ApproveTransferUpdated", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseApproveTransferUpdated(log types.Log) (*CreatorCoreERC721ApproveTransferUpdated, error) {
	event := new(CreatorCoreERC721ApproveTransferUpdated)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "ApproveTransferUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721DefaultRoyaltiesUpdatedIterator is returned from FilterDefaultRoyaltiesUpdated and is used to iterate over the raw logs and unpacked data for DefaultRoyaltiesUpdated events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721DefaultRoyaltiesUpdatedIterator struct {
	Event *CreatorCoreERC721DefaultRoyaltiesUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721DefaultRoyaltiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721DefaultRoyaltiesUpdated)
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
		it.Event = new(CreatorCoreERC721DefaultRoyaltiesUpdated)
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
func (it *CreatorCoreERC721DefaultRoyaltiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721DefaultRoyaltiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721DefaultRoyaltiesUpdated represents a DefaultRoyaltiesUpdated event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721DefaultRoyaltiesUpdated struct {
	Receivers   []common.Address
	BasisPoints []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyaltiesUpdated is a free log retrieval operation binding the contract event 0x2b6849d5976d799a5b0ca4dfd6b40a3d7afe9ea72c091fa01a958594f9a2659b.
//
// Solidity: event DefaultRoyaltiesUpdated(address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterDefaultRoyaltiesUpdated(opts *bind.FilterOpts) (*CreatorCoreERC721DefaultRoyaltiesUpdatedIterator, error) {

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "DefaultRoyaltiesUpdated")
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721DefaultRoyaltiesUpdatedIterator{contract: _CreatorCoreERC721.contract, event: "DefaultRoyaltiesUpdated", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyaltiesUpdated is a free log subscription operation binding the contract event 0x2b6849d5976d799a5b0ca4dfd6b40a3d7afe9ea72c091fa01a958594f9a2659b.
//
// Solidity: event DefaultRoyaltiesUpdated(address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchDefaultRoyaltiesUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721DefaultRoyaltiesUpdated) (event.Subscription, error) {

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "DefaultRoyaltiesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721DefaultRoyaltiesUpdated)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "DefaultRoyaltiesUpdated", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseDefaultRoyaltiesUpdated(log types.Log) (*CreatorCoreERC721DefaultRoyaltiesUpdated, error) {
	event := new(CreatorCoreERC721DefaultRoyaltiesUpdated)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "DefaultRoyaltiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721ExtensionApproveTransferUpdatedIterator is returned from FilterExtensionApproveTransferUpdated and is used to iterate over the raw logs and unpacked data for ExtensionApproveTransferUpdated events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionApproveTransferUpdatedIterator struct {
	Event *CreatorCoreERC721ExtensionApproveTransferUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721ExtensionApproveTransferUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721ExtensionApproveTransferUpdated)
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
		it.Event = new(CreatorCoreERC721ExtensionApproveTransferUpdated)
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
func (it *CreatorCoreERC721ExtensionApproveTransferUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721ExtensionApproveTransferUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721ExtensionApproveTransferUpdated represents a ExtensionApproveTransferUpdated event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionApproveTransferUpdated struct {
	Extension common.Address
	Enabled   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExtensionApproveTransferUpdated is a free log retrieval operation binding the contract event 0x072a7592283e2c2d1d56d21517ff6013325e0f55483f4828373ff4d98b0a1a36.
//
// Solidity: event ExtensionApproveTransferUpdated(address indexed extension, bool enabled)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterExtensionApproveTransferUpdated(opts *bind.FilterOpts, extension []common.Address) (*CreatorCoreERC721ExtensionApproveTransferUpdatedIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "ExtensionApproveTransferUpdated", extensionRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721ExtensionApproveTransferUpdatedIterator{contract: _CreatorCoreERC721.contract, event: "ExtensionApproveTransferUpdated", logs: logs, sub: sub}, nil
}

// WatchExtensionApproveTransferUpdated is a free log subscription operation binding the contract event 0x072a7592283e2c2d1d56d21517ff6013325e0f55483f4828373ff4d98b0a1a36.
//
// Solidity: event ExtensionApproveTransferUpdated(address indexed extension, bool enabled)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchExtensionApproveTransferUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721ExtensionApproveTransferUpdated, extension []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "ExtensionApproveTransferUpdated", extensionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721ExtensionApproveTransferUpdated)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionApproveTransferUpdated", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseExtensionApproveTransferUpdated(log types.Log) (*CreatorCoreERC721ExtensionApproveTransferUpdated, error) {
	event := new(CreatorCoreERC721ExtensionApproveTransferUpdated)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionApproveTransferUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721ExtensionBlacklistedIterator is returned from FilterExtensionBlacklisted and is used to iterate over the raw logs and unpacked data for ExtensionBlacklisted events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionBlacklistedIterator struct {
	Event *CreatorCoreERC721ExtensionBlacklisted // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721ExtensionBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721ExtensionBlacklisted)
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
		it.Event = new(CreatorCoreERC721ExtensionBlacklisted)
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
func (it *CreatorCoreERC721ExtensionBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721ExtensionBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721ExtensionBlacklisted represents a ExtensionBlacklisted event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionBlacklisted struct {
	Extension common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExtensionBlacklisted is a free log retrieval operation binding the contract event 0x05ac7bc5a606cd92a63365f9fda244499b9add0526b22d99937b6bd88181059c.
//
// Solidity: event ExtensionBlacklisted(address indexed extension, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterExtensionBlacklisted(opts *bind.FilterOpts, extension []common.Address, sender []common.Address) (*CreatorCoreERC721ExtensionBlacklistedIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "ExtensionBlacklisted", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721ExtensionBlacklistedIterator{contract: _CreatorCoreERC721.contract, event: "ExtensionBlacklisted", logs: logs, sub: sub}, nil
}

// WatchExtensionBlacklisted is a free log subscription operation binding the contract event 0x05ac7bc5a606cd92a63365f9fda244499b9add0526b22d99937b6bd88181059c.
//
// Solidity: event ExtensionBlacklisted(address indexed extension, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchExtensionBlacklisted(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721ExtensionBlacklisted, extension []common.Address, sender []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "ExtensionBlacklisted", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721ExtensionBlacklisted)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionBlacklisted", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseExtensionBlacklisted(log types.Log) (*CreatorCoreERC721ExtensionBlacklisted, error) {
	event := new(CreatorCoreERC721ExtensionBlacklisted)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721ExtensionRegisteredIterator is returned from FilterExtensionRegistered and is used to iterate over the raw logs and unpacked data for ExtensionRegistered events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionRegisteredIterator struct {
	Event *CreatorCoreERC721ExtensionRegistered // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721ExtensionRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721ExtensionRegistered)
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
		it.Event = new(CreatorCoreERC721ExtensionRegistered)
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
func (it *CreatorCoreERC721ExtensionRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721ExtensionRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721ExtensionRegistered represents a ExtensionRegistered event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionRegistered struct {
	Extension common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExtensionRegistered is a free log retrieval operation binding the contract event 0xd8cb8ba4086944eabf43c5535b7712015e4d4c714b24bf812c040ea5b7a3e42a.
//
// Solidity: event ExtensionRegistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterExtensionRegistered(opts *bind.FilterOpts, extension []common.Address, sender []common.Address) (*CreatorCoreERC721ExtensionRegisteredIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "ExtensionRegistered", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721ExtensionRegisteredIterator{contract: _CreatorCoreERC721.contract, event: "ExtensionRegistered", logs: logs, sub: sub}, nil
}

// WatchExtensionRegistered is a free log subscription operation binding the contract event 0xd8cb8ba4086944eabf43c5535b7712015e4d4c714b24bf812c040ea5b7a3e42a.
//
// Solidity: event ExtensionRegistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchExtensionRegistered(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721ExtensionRegistered, extension []common.Address, sender []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "ExtensionRegistered", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721ExtensionRegistered)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionRegistered", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseExtensionRegistered(log types.Log) (*CreatorCoreERC721ExtensionRegistered, error) {
	event := new(CreatorCoreERC721ExtensionRegistered)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721ExtensionRoyaltiesUpdatedIterator is returned from FilterExtensionRoyaltiesUpdated and is used to iterate over the raw logs and unpacked data for ExtensionRoyaltiesUpdated events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionRoyaltiesUpdatedIterator struct {
	Event *CreatorCoreERC721ExtensionRoyaltiesUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721ExtensionRoyaltiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721ExtensionRoyaltiesUpdated)
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
		it.Event = new(CreatorCoreERC721ExtensionRoyaltiesUpdated)
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
func (it *CreatorCoreERC721ExtensionRoyaltiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721ExtensionRoyaltiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721ExtensionRoyaltiesUpdated represents a ExtensionRoyaltiesUpdated event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionRoyaltiesUpdated struct {
	Extension   common.Address
	Receivers   []common.Address
	BasisPoints []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtensionRoyaltiesUpdated is a free log retrieval operation binding the contract event 0x535a93d2cb000582c0ebeaa9be4890ec6a287f98eb2df00c54c300612fd78d8f.
//
// Solidity: event ExtensionRoyaltiesUpdated(address indexed extension, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterExtensionRoyaltiesUpdated(opts *bind.FilterOpts, extension []common.Address) (*CreatorCoreERC721ExtensionRoyaltiesUpdatedIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "ExtensionRoyaltiesUpdated", extensionRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721ExtensionRoyaltiesUpdatedIterator{contract: _CreatorCoreERC721.contract, event: "ExtensionRoyaltiesUpdated", logs: logs, sub: sub}, nil
}

// WatchExtensionRoyaltiesUpdated is a free log subscription operation binding the contract event 0x535a93d2cb000582c0ebeaa9be4890ec6a287f98eb2df00c54c300612fd78d8f.
//
// Solidity: event ExtensionRoyaltiesUpdated(address indexed extension, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchExtensionRoyaltiesUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721ExtensionRoyaltiesUpdated, extension []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "ExtensionRoyaltiesUpdated", extensionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721ExtensionRoyaltiesUpdated)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionRoyaltiesUpdated", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseExtensionRoyaltiesUpdated(log types.Log) (*CreatorCoreERC721ExtensionRoyaltiesUpdated, error) {
	event := new(CreatorCoreERC721ExtensionRoyaltiesUpdated)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionRoyaltiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721ExtensionUnregisteredIterator is returned from FilterExtensionUnregistered and is used to iterate over the raw logs and unpacked data for ExtensionUnregistered events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionUnregisteredIterator struct {
	Event *CreatorCoreERC721ExtensionUnregistered // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721ExtensionUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721ExtensionUnregistered)
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
		it.Event = new(CreatorCoreERC721ExtensionUnregistered)
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
func (it *CreatorCoreERC721ExtensionUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721ExtensionUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721ExtensionUnregistered represents a ExtensionUnregistered event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721ExtensionUnregistered struct {
	Extension common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExtensionUnregistered is a free log retrieval operation binding the contract event 0xd19cf84cf0fec6bec9ddfa29c63adf83a55707c712f32c8285d6180a78901479.
//
// Solidity: event ExtensionUnregistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterExtensionUnregistered(opts *bind.FilterOpts, extension []common.Address, sender []common.Address) (*CreatorCoreERC721ExtensionUnregisteredIterator, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "ExtensionUnregistered", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721ExtensionUnregisteredIterator{contract: _CreatorCoreERC721.contract, event: "ExtensionUnregistered", logs: logs, sub: sub}, nil
}

// WatchExtensionUnregistered is a free log subscription operation binding the contract event 0xd19cf84cf0fec6bec9ddfa29c63adf83a55707c712f32c8285d6180a78901479.
//
// Solidity: event ExtensionUnregistered(address indexed extension, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchExtensionUnregistered(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721ExtensionUnregistered, extension []common.Address, sender []common.Address) (event.Subscription, error) {

	var extensionRule []interface{}
	for _, extensionItem := range extension {
		extensionRule = append(extensionRule, extensionItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "ExtensionUnregistered", extensionRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721ExtensionUnregistered)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionUnregistered", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseExtensionUnregistered(log types.Log) (*CreatorCoreERC721ExtensionUnregistered, error) {
	event := new(CreatorCoreERC721ExtensionUnregistered)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "ExtensionUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721InitializedIterator struct {
	Event *CreatorCoreERC721Initialized // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721Initialized)
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
		it.Event = new(CreatorCoreERC721Initialized)
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
func (it *CreatorCoreERC721InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721Initialized represents a Initialized event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterInitialized(opts *bind.FilterOpts) (*CreatorCoreERC721InitializedIterator, error) {

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721InitializedIterator{contract: _CreatorCoreERC721.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721Initialized) (event.Subscription, error) {

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721Initialized)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseInitialized(log types.Log) (*CreatorCoreERC721Initialized, error) {
	event := new(CreatorCoreERC721Initialized)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721MintPermissionsUpdatedIterator is returned from FilterMintPermissionsUpdated and is used to iterate over the raw logs and unpacked data for MintPermissionsUpdated events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721MintPermissionsUpdatedIterator struct {
	Event *CreatorCoreERC721MintPermissionsUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721MintPermissionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721MintPermissionsUpdated)
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
		it.Event = new(CreatorCoreERC721MintPermissionsUpdated)
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
func (it *CreatorCoreERC721MintPermissionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721MintPermissionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721MintPermissionsUpdated represents a MintPermissionsUpdated event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721MintPermissionsUpdated struct {
	Extension   common.Address
	Permissions common.Address
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMintPermissionsUpdated is a free log retrieval operation binding the contract event 0x6a835c4fcf7e0d398db3762332fdaa1471814ad39f1e2d6d0b3fdabf8efee3e0.
//
// Solidity: event MintPermissionsUpdated(address indexed extension, address indexed permissions, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterMintPermissionsUpdated(opts *bind.FilterOpts, extension []common.Address, permissions []common.Address, sender []common.Address) (*CreatorCoreERC721MintPermissionsUpdatedIterator, error) {

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

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "MintPermissionsUpdated", extensionRule, permissionsRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721MintPermissionsUpdatedIterator{contract: _CreatorCoreERC721.contract, event: "MintPermissionsUpdated", logs: logs, sub: sub}, nil
}

// WatchMintPermissionsUpdated is a free log subscription operation binding the contract event 0x6a835c4fcf7e0d398db3762332fdaa1471814ad39f1e2d6d0b3fdabf8efee3e0.
//
// Solidity: event MintPermissionsUpdated(address indexed extension, address indexed permissions, address indexed sender)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchMintPermissionsUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721MintPermissionsUpdated, extension []common.Address, permissions []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "MintPermissionsUpdated", extensionRule, permissionsRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721MintPermissionsUpdated)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "MintPermissionsUpdated", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseMintPermissionsUpdated(log types.Log) (*CreatorCoreERC721MintPermissionsUpdated, error) {
	event := new(CreatorCoreERC721MintPermissionsUpdated)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "MintPermissionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721OwnershipTransferredIterator struct {
	Event *CreatorCoreERC721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721OwnershipTransferred)
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
		it.Event = new(CreatorCoreERC721OwnershipTransferred)
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
func (it *CreatorCoreERC721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721OwnershipTransferred represents a OwnershipTransferred event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CreatorCoreERC721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721OwnershipTransferredIterator{contract: _CreatorCoreERC721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721OwnershipTransferred)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseOwnershipTransferred(log types.Log) (*CreatorCoreERC721OwnershipTransferred, error) {
	event := new(CreatorCoreERC721OwnershipTransferred)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721RoyaltiesUpdatedIterator is returned from FilterRoyaltiesUpdated and is used to iterate over the raw logs and unpacked data for RoyaltiesUpdated events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721RoyaltiesUpdatedIterator struct {
	Event *CreatorCoreERC721RoyaltiesUpdated // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721RoyaltiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721RoyaltiesUpdated)
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
		it.Event = new(CreatorCoreERC721RoyaltiesUpdated)
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
func (it *CreatorCoreERC721RoyaltiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721RoyaltiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721RoyaltiesUpdated represents a RoyaltiesUpdated event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721RoyaltiesUpdated struct {
	TokenId     *big.Int
	Receivers   []common.Address
	BasisPoints []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRoyaltiesUpdated is a free log retrieval operation binding the contract event 0xabb46fe0761d77584bde75697647804ffd8113abd4d8d06bc664150395eccdee.
//
// Solidity: event RoyaltiesUpdated(uint256 indexed tokenId, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterRoyaltiesUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*CreatorCoreERC721RoyaltiesUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "RoyaltiesUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721RoyaltiesUpdatedIterator{contract: _CreatorCoreERC721.contract, event: "RoyaltiesUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltiesUpdated is a free log subscription operation binding the contract event 0xabb46fe0761d77584bde75697647804ffd8113abd4d8d06bc664150395eccdee.
//
// Solidity: event RoyaltiesUpdated(uint256 indexed tokenId, address[] receivers, uint256[] basisPoints)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchRoyaltiesUpdated(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721RoyaltiesUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "RoyaltiesUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721RoyaltiesUpdated)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "RoyaltiesUpdated", log); err != nil {
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
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseRoyaltiesUpdated(log types.Log) (*CreatorCoreERC721RoyaltiesUpdated, error) {
	event := new(CreatorCoreERC721RoyaltiesUpdated)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "RoyaltiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorCoreERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721TransferIterator struct {
	Event *CreatorCoreERC721Transfer // Event containing the contract specifics and raw log

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
func (it *CreatorCoreERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorCoreERC721Transfer)
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
		it.Event = new(CreatorCoreERC721Transfer)
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
func (it *CreatorCoreERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorCoreERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorCoreERC721Transfer represents a Transfer event raised by the CreatorCoreERC721 contract.
type CreatorCoreERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CreatorCoreERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CreatorCoreERC721TransferIterator{contract: _CreatorCoreERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CreatorCoreERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CreatorCoreERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorCoreERC721Transfer)
				if err := _CreatorCoreERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CreatorCoreERC721 *CreatorCoreERC721Filterer) ParseTransfer(log types.Log) (*CreatorCoreERC721Transfer, error) {
	event := new(CreatorCoreERC721Transfer)
	if err := _CreatorCoreERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
