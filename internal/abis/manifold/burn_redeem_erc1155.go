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

// IBurnRedeemCoreBurnGroup is an auto generated low-level Go binding around an user-defined struct.
type IBurnRedeemCoreBurnGroup struct {
	RequiredCount *big.Int
	Items         []IBurnRedeemCoreBurnItem
}

// IBurnRedeemCoreBurnItem is an auto generated low-level Go binding around an user-defined struct.
type IBurnRedeemCoreBurnItem struct {
	ValidationType  uint8
	ContractAddress common.Address
	TokenSpec       uint8
	BurnSpec        uint8
	Amount          *big.Int
	MinTokenId      *big.Int
	MaxTokenId      *big.Int
	MerkleRoot      [32]byte
}

// IBurnRedeemCoreBurnRedeem is an auto generated low-level Go binding around an user-defined struct.
type IBurnRedeemCoreBurnRedeem struct {
	PaymentReceiver common.Address
	StorageProtocol uint8
	RedeemedCount   uint32
	RedeemAmount    uint16
	TotalSupply     uint32
	ContractVersion uint8
	StartDate       *big.Int
	EndDate         *big.Int
	Cost            *big.Int
	Location        string
	BurnSet         []IBurnRedeemCoreBurnGroup
}

// IBurnRedeemCoreBurnRedeemParameters is an auto generated low-level Go binding around an user-defined struct.
type IBurnRedeemCoreBurnRedeemParameters struct {
	PaymentReceiver common.Address
	StorageProtocol uint8
	RedeemAmount    uint16
	TotalSupply     uint32
	StartDate       *big.Int
	EndDate         *big.Int
	Cost            *big.Int
	Location        string
	BurnSet         []IBurnRedeemCoreBurnGroup
}

// IBurnRedeemCoreBurnToken is an auto generated low-level Go binding around an user-defined struct.
type IBurnRedeemCoreBurnToken struct {
	GroupIndex      *big.Int
	ItemIndex       *big.Int
	ContractAddress common.Address
	Id              *big.Int
	MerkleProof     [][32]byte
}

// BurnRedeemERC1155MetaData contains all meta data concerning the BurnRedeemERC1155 contract.
var BurnRedeemERC1155MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AdminRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_BURN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"amounts\",\"type\":\"uint32[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"approveAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"creatorContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"instanceIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32[]\",\"name\":\"burnRedeemCounts\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint48\",\"name\":\"groupIndex\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"itemIndex\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnToken[][]\",\"name\":\"burnTokens\",\"type\":\"tuple[][]\"}],\"name\":\"burnRedeem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"burnRedeemCount\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint48\",\"name\":\"groupIndex\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"itemIndex\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnToken[]\",\"name\":\"burnTokens\",\"type\":\"tuple[]\"}],\"name\":\"burnRedeem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"}],\"name\":\"getBurnRedeem\",\"outputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"enumIBurnRedeemCore.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"redeemedCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"redeemAmount\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"totalSupply\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"contractVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"uint160\",\"name\":\"cost\",\"type\":\"uint160\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumIBurnRedeemCore.ValidationType\",\"name\":\"validationType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"enumIBurnRedeemCore.TokenSpec\",\"name\":\"tokenSpec\",\"type\":\"uint8\"},{\"internalType\":\"enumIBurnRedeemCore.BurnSpec\",\"name\":\"burnSpec\",\"type\":\"uint8\"},{\"internalType\":\"uint72\",\"name\":\"amount\",\"type\":\"uint72\"},{\"internalType\":\"uint256\",\"name\":\"minTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIBurnRedeemCore.BurnItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnGroup[]\",\"name\":\"burnSet\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnRedeem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBurnRedeemForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"enumIBurnRedeemCore.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"redeemedCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"redeemAmount\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"totalSupply\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"contractVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"uint160\",\"name\":\"cost\",\"type\":\"uint160\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumIBurnRedeemCore.ValidationType\",\"name\":\"validationType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"enumIBurnRedeemCore.TokenSpec\",\"name\":\"tokenSpec\",\"type\":\"uint8\"},{\"internalType\":\"enumIBurnRedeemCore.BurnSpec\",\"name\":\"burnSpec\",\"type\":\"uint8\"},{\"internalType\":\"uint72\",\"name\":\"amount\",\"type\":\"uint72\"},{\"internalType\":\"uint256\",\"name\":\"minTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIBurnRedeemCore.BurnItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnGroup[]\",\"name\":\"burnSet\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnRedeem\",\"name\":\"burnRedeem\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"enumIBurnRedeemCore.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"redeemAmount\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"totalSupply\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"uint160\",\"name\":\"cost\",\"type\":\"uint160\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumIBurnRedeemCore.ValidationType\",\"name\":\"validationType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"enumIBurnRedeemCore.TokenSpec\",\"name\":\"tokenSpec\",\"type\":\"uint8\"},{\"internalType\":\"enumIBurnRedeemCore.BurnSpec\",\"name\":\"burnSpec\",\"type\":\"uint8\"},{\"internalType\":\"uint72\",\"name\":\"amount\",\"type\":\"uint72\"},{\"internalType\":\"uint256\",\"name\":\"minTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIBurnRedeemCore.BurnItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnGroup[]\",\"name\":\"burnSet\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnRedeemParameters\",\"name\":\"burnRedeemParameters\",\"type\":\"tuple\"}],\"name\":\"initializeBurnRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manifoldMembershipContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"revokeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setMembershipAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"enumIBurnRedeemCore.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"redeemAmount\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"totalSupply\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"startDate\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endDate\",\"type\":\"uint48\"},{\"internalType\":\"uint160\",\"name\":\"cost\",\"type\":\"uint160\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumIBurnRedeemCore.ValidationType\",\"name\":\"validationType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"enumIBurnRedeemCore.TokenSpec\",\"name\":\"tokenSpec\",\"type\":\"uint8\"},{\"internalType\":\"enumIBurnRedeemCore.BurnSpec\",\"name\":\"burnSpec\",\"type\":\"uint8\"},{\"internalType\":\"uint72\",\"name\":\"amount\",\"type\":\"uint72\"},{\"internalType\":\"uint256\",\"name\":\"minTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIBurnRedeemCore.BurnItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnGroup[]\",\"name\":\"burnSet\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBurnRedeemCore.BurnRedeemParameters\",\"name\":\"burnRedeemParameters\",\"type\":\"tuple\"}],\"name\":\"updateBurnRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instanceId\",\"type\":\"uint256\"},{\"internalType\":\"enumIBurnRedeemCore.StorageProtocol\",\"name\":\"storageProtocol\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"updateURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BurnRedeemERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use BurnRedeemERC1155MetaData.ABI instead.
var BurnRedeemERC1155ABI = BurnRedeemERC1155MetaData.ABI

// BurnRedeemERC1155 is an auto generated Go binding around an Ethereum contract.
type BurnRedeemERC1155 struct {
	BurnRedeemERC1155Caller     // Read-only binding to the contract
	BurnRedeemERC1155Transactor // Write-only binding to the contract
	BurnRedeemERC1155Filterer   // Log filterer for contract events
}

// BurnRedeemERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type BurnRedeemERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnRedeemERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnRedeemERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnRedeemERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnRedeemERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnRedeemERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnRedeemERC1155Session struct {
	Contract     *BurnRedeemERC1155 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BurnRedeemERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnRedeemERC1155CallerSession struct {
	Contract *BurnRedeemERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BurnRedeemERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnRedeemERC1155TransactorSession struct {
	Contract     *BurnRedeemERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BurnRedeemERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type BurnRedeemERC1155Raw struct {
	Contract *BurnRedeemERC1155 // Generic contract binding to access the raw methods on
}

// BurnRedeemERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnRedeemERC1155CallerRaw struct {
	Contract *BurnRedeemERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// BurnRedeemERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnRedeemERC1155TransactorRaw struct {
	Contract *BurnRedeemERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnRedeemERC1155 creates a new instance of BurnRedeemERC1155, bound to a specific deployed contract.
func NewBurnRedeemERC1155(address common.Address, backend bind.ContractBackend) (*BurnRedeemERC1155, error) {
	contract, err := bindBurnRedeemERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnRedeemERC1155{BurnRedeemERC1155Caller: BurnRedeemERC1155Caller{contract: contract}, BurnRedeemERC1155Transactor: BurnRedeemERC1155Transactor{contract: contract}, BurnRedeemERC1155Filterer: BurnRedeemERC1155Filterer{contract: contract}}, nil
}

// NewBurnRedeemERC1155Caller creates a new read-only instance of BurnRedeemERC1155, bound to a specific deployed contract.
func NewBurnRedeemERC1155Caller(address common.Address, caller bind.ContractCaller) (*BurnRedeemERC1155Caller, error) {
	contract, err := bindBurnRedeemERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnRedeemERC1155Caller{contract: contract}, nil
}

// NewBurnRedeemERC1155Transactor creates a new write-only instance of BurnRedeemERC1155, bound to a specific deployed contract.
func NewBurnRedeemERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*BurnRedeemERC1155Transactor, error) {
	contract, err := bindBurnRedeemERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnRedeemERC1155Transactor{contract: contract}, nil
}

// NewBurnRedeemERC1155Filterer creates a new log filterer instance of BurnRedeemERC1155, bound to a specific deployed contract.
func NewBurnRedeemERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*BurnRedeemERC1155Filterer, error) {
	contract, err := bindBurnRedeemERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnRedeemERC1155Filterer{contract: contract}, nil
}

// bindBurnRedeemERC1155 binds a generic wrapper to an already deployed contract.
func bindBurnRedeemERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnRedeemERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnRedeemERC1155 *BurnRedeemERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnRedeemERC1155.Contract.BurnRedeemERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnRedeemERC1155 *BurnRedeemERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.BurnRedeemERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnRedeemERC1155 *BurnRedeemERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.BurnRedeemERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnRedeemERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.contract.Transact(opts, method, params...)
}

// BURNFEE is a free data retrieval call binding the contract method 0x480df058.
//
// Solidity: function BURN_FEE() view returns(uint256)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) BURNFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "BURN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BURNFEE is a free data retrieval call binding the contract method 0x480df058.
//
// Solidity: function BURN_FEE() view returns(uint256)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) BURNFEE() (*big.Int, error) {
	return _BurnRedeemERC1155.Contract.BURNFEE(&_BurnRedeemERC1155.CallOpts)
}

// BURNFEE is a free data retrieval call binding the contract method 0x480df058.
//
// Solidity: function BURN_FEE() view returns(uint256)
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) BURNFEE() (*big.Int, error) {
	return _BurnRedeemERC1155.Contract.BURNFEE(&_BurnRedeemERC1155.CallOpts)
}

// MULTIBURNFEE is a free data retrieval call binding the contract method 0x95bc312b.
//
// Solidity: function MULTI_BURN_FEE() view returns(uint256)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) MULTIBURNFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "MULTI_BURN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIBURNFEE is a free data retrieval call binding the contract method 0x95bc312b.
//
// Solidity: function MULTI_BURN_FEE() view returns(uint256)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) MULTIBURNFEE() (*big.Int, error) {
	return _BurnRedeemERC1155.Contract.MULTIBURNFEE(&_BurnRedeemERC1155.CallOpts)
}

// MULTIBURNFEE is a free data retrieval call binding the contract method 0x95bc312b.
//
// Solidity: function MULTI_BURN_FEE() view returns(uint256)
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) MULTIBURNFEE() (*big.Int, error) {
	return _BurnRedeemERC1155.Contract.MULTIBURNFEE(&_BurnRedeemERC1155.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) GetAdmins() ([]common.Address, error) {
	return _BurnRedeemERC1155.Contract.GetAdmins(&_BurnRedeemERC1155.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[] admins)
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) GetAdmins() ([]common.Address, error) {
	return _BurnRedeemERC1155.Contract.GetAdmins(&_BurnRedeemERC1155.CallOpts)
}

// GetBurnRedeem is a free data retrieval call binding the contract method 0xb058d2f8.
//
// Solidity: function getBurnRedeem(address creatorContractAddress, uint256 instanceId) view returns((address,uint8,uint32,uint16,uint32,uint8,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]))
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) GetBurnRedeem(opts *bind.CallOpts, creatorContractAddress common.Address, instanceId *big.Int) (IBurnRedeemCoreBurnRedeem, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "getBurnRedeem", creatorContractAddress, instanceId)

	if err != nil {
		return *new(IBurnRedeemCoreBurnRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new(IBurnRedeemCoreBurnRedeem)).(*IBurnRedeemCoreBurnRedeem)

	return out0, err

}

// GetBurnRedeem is a free data retrieval call binding the contract method 0xb058d2f8.
//
// Solidity: function getBurnRedeem(address creatorContractAddress, uint256 instanceId) view returns((address,uint8,uint32,uint16,uint32,uint8,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]))
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) GetBurnRedeem(creatorContractAddress common.Address, instanceId *big.Int) (IBurnRedeemCoreBurnRedeem, error) {
	return _BurnRedeemERC1155.Contract.GetBurnRedeem(&_BurnRedeemERC1155.CallOpts, creatorContractAddress, instanceId)
}

// GetBurnRedeem is a free data retrieval call binding the contract method 0xb058d2f8.
//
// Solidity: function getBurnRedeem(address creatorContractAddress, uint256 instanceId) view returns((address,uint8,uint32,uint16,uint32,uint8,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]))
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) GetBurnRedeem(creatorContractAddress common.Address, instanceId *big.Int) (IBurnRedeemCoreBurnRedeem, error) {
	return _BurnRedeemERC1155.Contract.GetBurnRedeem(&_BurnRedeemERC1155.CallOpts, creatorContractAddress, instanceId)
}

// GetBurnRedeemForToken is a free data retrieval call binding the contract method 0x41d29f07.
//
// Solidity: function getBurnRedeemForToken(address creatorContractAddress, uint256 tokenId) view returns(uint256 instanceId, (address,uint8,uint32,uint16,uint32,uint8,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]) burnRedeem)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) GetBurnRedeemForToken(opts *bind.CallOpts, creatorContractAddress common.Address, tokenId *big.Int) (struct {
	InstanceId *big.Int
	BurnRedeem IBurnRedeemCoreBurnRedeem
}, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "getBurnRedeemForToken", creatorContractAddress, tokenId)

	outstruct := new(struct {
		InstanceId *big.Int
		BurnRedeem IBurnRedeemCoreBurnRedeem
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InstanceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BurnRedeem = *abi.ConvertType(out[1], new(IBurnRedeemCoreBurnRedeem)).(*IBurnRedeemCoreBurnRedeem)

	return *outstruct, err

}

// GetBurnRedeemForToken is a free data retrieval call binding the contract method 0x41d29f07.
//
// Solidity: function getBurnRedeemForToken(address creatorContractAddress, uint256 tokenId) view returns(uint256 instanceId, (address,uint8,uint32,uint16,uint32,uint8,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]) burnRedeem)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) GetBurnRedeemForToken(creatorContractAddress common.Address, tokenId *big.Int) (struct {
	InstanceId *big.Int
	BurnRedeem IBurnRedeemCoreBurnRedeem
}, error) {
	return _BurnRedeemERC1155.Contract.GetBurnRedeemForToken(&_BurnRedeemERC1155.CallOpts, creatorContractAddress, tokenId)
}

// GetBurnRedeemForToken is a free data retrieval call binding the contract method 0x41d29f07.
//
// Solidity: function getBurnRedeemForToken(address creatorContractAddress, uint256 tokenId) view returns(uint256 instanceId, (address,uint8,uint32,uint16,uint32,uint8,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]) burnRedeem)
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) GetBurnRedeemForToken(creatorContractAddress common.Address, tokenId *big.Int) (struct {
	InstanceId *big.Int
	BurnRedeem IBurnRedeemCoreBurnRedeem
}, error) {
	return _BurnRedeemERC1155.Contract.GetBurnRedeemForToken(&_BurnRedeemERC1155.CallOpts, creatorContractAddress, tokenId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) IsAdmin(opts *bind.CallOpts, admin common.Address) (bool, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "isAdmin", admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) IsAdmin(admin common.Address) (bool, error) {
	return _BurnRedeemERC1155.Contract.IsAdmin(&_BurnRedeemERC1155.CallOpts, admin)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin) view returns(bool)
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) IsAdmin(admin common.Address) (bool, error) {
	return _BurnRedeemERC1155.Contract.IsAdmin(&_BurnRedeemERC1155.CallOpts, admin)
}

// ManifoldMembershipContract is a free data retrieval call binding the contract method 0x50d10839.
//
// Solidity: function manifoldMembershipContract() view returns(address)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) ManifoldMembershipContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "manifoldMembershipContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManifoldMembershipContract is a free data retrieval call binding the contract method 0x50d10839.
//
// Solidity: function manifoldMembershipContract() view returns(address)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) ManifoldMembershipContract() (common.Address, error) {
	return _BurnRedeemERC1155.Contract.ManifoldMembershipContract(&_BurnRedeemERC1155.CallOpts)
}

// ManifoldMembershipContract is a free data retrieval call binding the contract method 0x50d10839.
//
// Solidity: function manifoldMembershipContract() view returns(address)
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) ManifoldMembershipContract() (common.Address, error) {
	return _BurnRedeemERC1155.Contract.ManifoldMembershipContract(&_BurnRedeemERC1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) Owner() (common.Address, error) {
	return _BurnRedeemERC1155.Contract.Owner(&_BurnRedeemERC1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) Owner() (common.Address, error) {
	return _BurnRedeemERC1155.Contract.Owner(&_BurnRedeemERC1155.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnRedeemERC1155.Contract.SupportsInterface(&_BurnRedeemERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnRedeemERC1155.Contract.SupportsInterface(&_BurnRedeemERC1155.CallOpts, interfaceId)
}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address creatorContractAddress, uint256 tokenId) view returns(string uri)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Caller) TokenURI(opts *bind.CallOpts, creatorContractAddress common.Address, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _BurnRedeemERC1155.contract.Call(opts, &out, "tokenURI", creatorContractAddress, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address creatorContractAddress, uint256 tokenId) view returns(string uri)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) TokenURI(creatorContractAddress common.Address, tokenId *big.Int) (string, error) {
	return _BurnRedeemERC1155.Contract.TokenURI(&_BurnRedeemERC1155.CallOpts, creatorContractAddress, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address creatorContractAddress, uint256 tokenId) view returns(string uri)
func (_BurnRedeemERC1155 *BurnRedeemERC1155CallerSession) TokenURI(creatorContractAddress common.Address, tokenId *big.Int) (string, error) {
	return _BurnRedeemERC1155.Contract.TokenURI(&_BurnRedeemERC1155.CallOpts, creatorContractAddress, tokenId)
}

// Airdrop is a paid mutator transaction binding the contract method 0xc2065c10.
//
// Solidity: function airdrop(address creatorContractAddress, uint256 instanceId, address[] recipients, uint32[] amounts) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) Airdrop(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, recipients []common.Address, amounts []uint32) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "airdrop", creatorContractAddress, instanceId, recipients, amounts)
}

// Airdrop is a paid mutator transaction binding the contract method 0xc2065c10.
//
// Solidity: function airdrop(address creatorContractAddress, uint256 instanceId, address[] recipients, uint32[] amounts) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) Airdrop(creatorContractAddress common.Address, instanceId *big.Int, recipients []common.Address, amounts []uint32) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.Airdrop(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, recipients, amounts)
}

// Airdrop is a paid mutator transaction binding the contract method 0xc2065c10.
//
// Solidity: function airdrop(address creatorContractAddress, uint256 instanceId, address[] recipients, uint32[] amounts) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) Airdrop(creatorContractAddress common.Address, instanceId *big.Int, recipients []common.Address, amounts []uint32) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.Airdrop(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, recipients, amounts)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) ApproveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "approveAdmin", admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.ApproveAdmin(&_BurnRedeemERC1155.TransactOpts, admin)
}

// ApproveAdmin is a paid mutator transaction binding the contract method 0x6d73e669.
//
// Solidity: function approveAdmin(address admin) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) ApproveAdmin(admin common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.ApproveAdmin(&_BurnRedeemERC1155.TransactOpts, admin)
}

// BurnRedeem is a paid mutator transaction binding the contract method 0x42b54fee.
//
// Solidity: function burnRedeem(address[] creatorContractAddresses, uint256[] instanceIds, uint32[] burnRedeemCounts, (uint48,uint48,address,uint256,bytes32[])[][] burnTokens) payable returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) BurnRedeem(opts *bind.TransactOpts, creatorContractAddresses []common.Address, instanceIds []*big.Int, burnRedeemCounts []uint32, burnTokens [][]IBurnRedeemCoreBurnToken) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "burnRedeem", creatorContractAddresses, instanceIds, burnRedeemCounts, burnTokens)
}

// BurnRedeem is a paid mutator transaction binding the contract method 0x42b54fee.
//
// Solidity: function burnRedeem(address[] creatorContractAddresses, uint256[] instanceIds, uint32[] burnRedeemCounts, (uint48,uint48,address,uint256,bytes32[])[][] burnTokens) payable returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) BurnRedeem(creatorContractAddresses []common.Address, instanceIds []*big.Int, burnRedeemCounts []uint32, burnTokens [][]IBurnRedeemCoreBurnToken) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.BurnRedeem(&_BurnRedeemERC1155.TransactOpts, creatorContractAddresses, instanceIds, burnRedeemCounts, burnTokens)
}

// BurnRedeem is a paid mutator transaction binding the contract method 0x42b54fee.
//
// Solidity: function burnRedeem(address[] creatorContractAddresses, uint256[] instanceIds, uint32[] burnRedeemCounts, (uint48,uint48,address,uint256,bytes32[])[][] burnTokens) payable returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) BurnRedeem(creatorContractAddresses []common.Address, instanceIds []*big.Int, burnRedeemCounts []uint32, burnTokens [][]IBurnRedeemCoreBurnToken) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.BurnRedeem(&_BurnRedeemERC1155.TransactOpts, creatorContractAddresses, instanceIds, burnRedeemCounts, burnTokens)
}

// BurnRedeem0 is a paid mutator transaction binding the contract method 0xc9dad696.
//
// Solidity: function burnRedeem(address creatorContractAddress, uint256 instanceId, uint32 burnRedeemCount, (uint48,uint48,address,uint256,bytes32[])[] burnTokens) payable returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) BurnRedeem0(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, burnRedeemCount uint32, burnTokens []IBurnRedeemCoreBurnToken) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "burnRedeem0", creatorContractAddress, instanceId, burnRedeemCount, burnTokens)
}

// BurnRedeem0 is a paid mutator transaction binding the contract method 0xc9dad696.
//
// Solidity: function burnRedeem(address creatorContractAddress, uint256 instanceId, uint32 burnRedeemCount, (uint48,uint48,address,uint256,bytes32[])[] burnTokens) payable returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) BurnRedeem0(creatorContractAddress common.Address, instanceId *big.Int, burnRedeemCount uint32, burnTokens []IBurnRedeemCoreBurnToken) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.BurnRedeem0(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, burnRedeemCount, burnTokens)
}

// BurnRedeem0 is a paid mutator transaction binding the contract method 0xc9dad696.
//
// Solidity: function burnRedeem(address creatorContractAddress, uint256 instanceId, uint32 burnRedeemCount, (uint48,uint48,address,uint256,bytes32[])[] burnTokens) payable returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) BurnRedeem0(creatorContractAddress common.Address, instanceId *big.Int, burnRedeemCount uint32, burnTokens []IBurnRedeemCoreBurnToken) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.BurnRedeem0(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, burnRedeemCount, burnTokens)
}

// InitializeBurnRedeem is a paid mutator transaction binding the contract method 0x38ec8995.
//
// Solidity: function initializeBurnRedeem(address creatorContractAddress, uint256 instanceId, (address,uint8,uint16,uint32,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]) burnRedeemParameters) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) InitializeBurnRedeem(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, burnRedeemParameters IBurnRedeemCoreBurnRedeemParameters) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "initializeBurnRedeem", creatorContractAddress, instanceId, burnRedeemParameters)
}

// InitializeBurnRedeem is a paid mutator transaction binding the contract method 0x38ec8995.
//
// Solidity: function initializeBurnRedeem(address creatorContractAddress, uint256 instanceId, (address,uint8,uint16,uint32,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]) burnRedeemParameters) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) InitializeBurnRedeem(creatorContractAddress common.Address, instanceId *big.Int, burnRedeemParameters IBurnRedeemCoreBurnRedeemParameters) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.InitializeBurnRedeem(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, burnRedeemParameters)
}

// InitializeBurnRedeem is a paid mutator transaction binding the contract method 0x38ec8995.
//
// Solidity: function initializeBurnRedeem(address creatorContractAddress, uint256 instanceId, (address,uint8,uint16,uint32,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]) burnRedeemParameters) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) InitializeBurnRedeem(creatorContractAddress common.Address, instanceId *big.Int, burnRedeemParameters IBurnRedeemCoreBurnRedeemParameters) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.InitializeBurnRedeem(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, burnRedeemParameters)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "onERC1155BatchReceived", arg0, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) OnERC1155BatchReceived(arg0 common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.OnERC1155BatchReceived(&_BurnRedeemERC1155.TransactOpts, arg0, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) OnERC1155BatchReceived(arg0 common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.OnERC1155BatchReceived(&_BurnRedeemERC1155.TransactOpts, arg0, from, ids, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "onERC1155Received", arg0, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) OnERC1155Received(arg0 common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.OnERC1155Received(&_BurnRedeemERC1155.TransactOpts, arg0, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) OnERC1155Received(arg0 common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.OnERC1155Received(&_BurnRedeemERC1155.TransactOpts, arg0, from, id, value, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 id, bytes data) returns(bytes4)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, from common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "onERC721Received", arg0, from, id, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 id, bytes data) returns(bytes4)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) OnERC721Received(arg0 common.Address, from common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.OnERC721Received(&_BurnRedeemERC1155.TransactOpts, arg0, from, id, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 id, bytes data) returns(bytes4)
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) OnERC721Received(arg0 common.Address, from common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.OnERC721Received(&_BurnRedeemERC1155.TransactOpts, arg0, from, id, data)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0xf0e9fcd1.
//
// Solidity: function recoverERC721(address tokenAddress, uint256 tokenId, address destination) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) RecoverERC721(opts *bind.TransactOpts, tokenAddress common.Address, tokenId *big.Int, destination common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "recoverERC721", tokenAddress, tokenId, destination)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0xf0e9fcd1.
//
// Solidity: function recoverERC721(address tokenAddress, uint256 tokenId, address destination) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) RecoverERC721(tokenAddress common.Address, tokenId *big.Int, destination common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.RecoverERC721(&_BurnRedeemERC1155.TransactOpts, tokenAddress, tokenId, destination)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0xf0e9fcd1.
//
// Solidity: function recoverERC721(address tokenAddress, uint256 tokenId, address destination) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) RecoverERC721(tokenAddress common.Address, tokenId *big.Int, destination common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.RecoverERC721(&_BurnRedeemERC1155.TransactOpts, tokenAddress, tokenId, destination)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) RenounceOwnership() (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.RenounceOwnership(&_BurnRedeemERC1155.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.RenounceOwnership(&_BurnRedeemERC1155.TransactOpts)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) RevokeAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "revokeAdmin", admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.RevokeAdmin(&_BurnRedeemERC1155.TransactOpts, admin)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address admin) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) RevokeAdmin(admin common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.RevokeAdmin(&_BurnRedeemERC1155.TransactOpts, admin)
}

// SetMembershipAddress is a paid mutator transaction binding the contract method 0x7ab39392.
//
// Solidity: function setMembershipAddress(address addr) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) SetMembershipAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "setMembershipAddress", addr)
}

// SetMembershipAddress is a paid mutator transaction binding the contract method 0x7ab39392.
//
// Solidity: function setMembershipAddress(address addr) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) SetMembershipAddress(addr common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.SetMembershipAddress(&_BurnRedeemERC1155.TransactOpts, addr)
}

// SetMembershipAddress is a paid mutator transaction binding the contract method 0x7ab39392.
//
// Solidity: function setMembershipAddress(address addr) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) SetMembershipAddress(addr common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.SetMembershipAddress(&_BurnRedeemERC1155.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.TransferOwnership(&_BurnRedeemERC1155.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.TransferOwnership(&_BurnRedeemERC1155.TransactOpts, newOwner)
}

// UpdateBurnRedeem is a paid mutator transaction binding the contract method 0x625888da.
//
// Solidity: function updateBurnRedeem(address creatorContractAddress, uint256 instanceId, (address,uint8,uint16,uint32,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]) burnRedeemParameters) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) UpdateBurnRedeem(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, burnRedeemParameters IBurnRedeemCoreBurnRedeemParameters) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "updateBurnRedeem", creatorContractAddress, instanceId, burnRedeemParameters)
}

// UpdateBurnRedeem is a paid mutator transaction binding the contract method 0x625888da.
//
// Solidity: function updateBurnRedeem(address creatorContractAddress, uint256 instanceId, (address,uint8,uint16,uint32,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]) burnRedeemParameters) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) UpdateBurnRedeem(creatorContractAddress common.Address, instanceId *big.Int, burnRedeemParameters IBurnRedeemCoreBurnRedeemParameters) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.UpdateBurnRedeem(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, burnRedeemParameters)
}

// UpdateBurnRedeem is a paid mutator transaction binding the contract method 0x625888da.
//
// Solidity: function updateBurnRedeem(address creatorContractAddress, uint256 instanceId, (address,uint8,uint16,uint32,uint48,uint48,uint160,string,(uint256,(uint8,address,uint8,uint8,uint72,uint256,uint256,bytes32)[])[]) burnRedeemParameters) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) UpdateBurnRedeem(creatorContractAddress common.Address, instanceId *big.Int, burnRedeemParameters IBurnRedeemCoreBurnRedeemParameters) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.UpdateBurnRedeem(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, burnRedeemParameters)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x04ea1e90.
//
// Solidity: function updateURI(address creatorContractAddress, uint256 instanceId, uint8 storageProtocol, string location) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) UpdateURI(opts *bind.TransactOpts, creatorContractAddress common.Address, instanceId *big.Int, storageProtocol uint8, location string) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "updateURI", creatorContractAddress, instanceId, storageProtocol, location)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x04ea1e90.
//
// Solidity: function updateURI(address creatorContractAddress, uint256 instanceId, uint8 storageProtocol, string location) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) UpdateURI(creatorContractAddress common.Address, instanceId *big.Int, storageProtocol uint8, location string) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.UpdateURI(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, storageProtocol, location)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x04ea1e90.
//
// Solidity: function updateURI(address creatorContractAddress, uint256 instanceId, uint8 storageProtocol, string location) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) UpdateURI(creatorContractAddress common.Address, instanceId *big.Int, storageProtocol uint8, location string) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.UpdateURI(&_BurnRedeemERC1155.TransactOpts, creatorContractAddress, instanceId, storageProtocol, location)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Transactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnRedeemERC1155.contract.Transact(opts, "withdraw", recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155Session) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.Withdraw(&_BurnRedeemERC1155.TransactOpts, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_BurnRedeemERC1155 *BurnRedeemERC1155TransactorSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnRedeemERC1155.Contract.Withdraw(&_BurnRedeemERC1155.TransactOpts, recipient, amount)
}

// BurnRedeemERC1155AdminApprovedIterator is returned from FilterAdminApproved and is used to iterate over the raw logs and unpacked data for AdminApproved events raised by the BurnRedeemERC1155 contract.
type BurnRedeemERC1155AdminApprovedIterator struct {
	Event *BurnRedeemERC1155AdminApproved // Event containing the contract specifics and raw log

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
func (it *BurnRedeemERC1155AdminApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnRedeemERC1155AdminApproved)
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
		it.Event = new(BurnRedeemERC1155AdminApproved)
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
func (it *BurnRedeemERC1155AdminApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnRedeemERC1155AdminApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnRedeemERC1155AdminApproved represents a AdminApproved event raised by the BurnRedeemERC1155 contract.
type BurnRedeemERC1155AdminApproved struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminApproved is a free log retrieval operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Filterer) FilterAdminApproved(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*BurnRedeemERC1155AdminApprovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnRedeemERC1155.contract.FilterLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnRedeemERC1155AdminApprovedIterator{contract: _BurnRedeemERC1155.contract, event: "AdminApproved", logs: logs, sub: sub}, nil
}

// WatchAdminApproved is a free log subscription operation binding the contract event 0x7e1a1a08d52e4ba0e21554733d66165fd5151f99460116223d9e3a608eec5cb1.
//
// Solidity: event AdminApproved(address indexed account, address indexed sender)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Filterer) WatchAdminApproved(opts *bind.WatchOpts, sink chan<- *BurnRedeemERC1155AdminApproved, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnRedeemERC1155.contract.WatchLogs(opts, "AdminApproved", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnRedeemERC1155AdminApproved)
				if err := _BurnRedeemERC1155.contract.UnpackLog(event, "AdminApproved", log); err != nil {
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
func (_BurnRedeemERC1155 *BurnRedeemERC1155Filterer) ParseAdminApproved(log types.Log) (*BurnRedeemERC1155AdminApproved, error) {
	event := new(BurnRedeemERC1155AdminApproved)
	if err := _BurnRedeemERC1155.contract.UnpackLog(event, "AdminApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnRedeemERC1155AdminRevokedIterator is returned from FilterAdminRevoked and is used to iterate over the raw logs and unpacked data for AdminRevoked events raised by the BurnRedeemERC1155 contract.
type BurnRedeemERC1155AdminRevokedIterator struct {
	Event *BurnRedeemERC1155AdminRevoked // Event containing the contract specifics and raw log

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
func (it *BurnRedeemERC1155AdminRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnRedeemERC1155AdminRevoked)
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
		it.Event = new(BurnRedeemERC1155AdminRevoked)
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
func (it *BurnRedeemERC1155AdminRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnRedeemERC1155AdminRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnRedeemERC1155AdminRevoked represents a AdminRevoked event raised by the BurnRedeemERC1155 contract.
type BurnRedeemERC1155AdminRevoked struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRevoked is a free log retrieval operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Filterer) FilterAdminRevoked(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*BurnRedeemERC1155AdminRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnRedeemERC1155.contract.FilterLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnRedeemERC1155AdminRevokedIterator{contract: _BurnRedeemERC1155.contract, event: "AdminRevoked", logs: logs, sub: sub}, nil
}

// WatchAdminRevoked is a free log subscription operation binding the contract event 0x7c0c3c84c67c85fcac635147348bfe374c24a1a93d0366d1cfe9d8853cbf89d5.
//
// Solidity: event AdminRevoked(address indexed account, address indexed sender)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Filterer) WatchAdminRevoked(opts *bind.WatchOpts, sink chan<- *BurnRedeemERC1155AdminRevoked, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnRedeemERC1155.contract.WatchLogs(opts, "AdminRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnRedeemERC1155AdminRevoked)
				if err := _BurnRedeemERC1155.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
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
func (_BurnRedeemERC1155 *BurnRedeemERC1155Filterer) ParseAdminRevoked(log types.Log) (*BurnRedeemERC1155AdminRevoked, error) {
	event := new(BurnRedeemERC1155AdminRevoked)
	if err := _BurnRedeemERC1155.contract.UnpackLog(event, "AdminRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnRedeemERC1155OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BurnRedeemERC1155 contract.
type BurnRedeemERC1155OwnershipTransferredIterator struct {
	Event *BurnRedeemERC1155OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BurnRedeemERC1155OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnRedeemERC1155OwnershipTransferred)
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
		it.Event = new(BurnRedeemERC1155OwnershipTransferred)
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
func (it *BurnRedeemERC1155OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnRedeemERC1155OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnRedeemERC1155OwnershipTransferred represents a OwnershipTransferred event raised by the BurnRedeemERC1155 contract.
type BurnRedeemERC1155OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BurnRedeemERC1155OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BurnRedeemERC1155.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnRedeemERC1155OwnershipTransferredIterator{contract: _BurnRedeemERC1155.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BurnRedeemERC1155 *BurnRedeemERC1155Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnRedeemERC1155OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BurnRedeemERC1155.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnRedeemERC1155OwnershipTransferred)
				if err := _BurnRedeemERC1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BurnRedeemERC1155 *BurnRedeemERC1155Filterer) ParseOwnershipTransferred(log types.Log) (*BurnRedeemERC1155OwnershipTransferred, error) {
	event := new(BurnRedeemERC1155OwnershipTransferred)
	if err := _BurnRedeemERC1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
