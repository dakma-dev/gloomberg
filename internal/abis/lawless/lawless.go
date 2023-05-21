// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lawless

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

// LawlessData is an auto generated low-level Go binding around an user-defined struct.
type LawlessData struct {
	Id        *big.Int
	ModelId   *big.Int
	PaletteId *big.Int
	Owner     common.Address
	Version   uint16
	Details   *big.Int
}

// Model is an auto generated low-level Go binding around an user-defined struct.
type Model struct {
	Width         uint8
	Height        uint8
	AniX          uint8
	AniY          uint8
	AniWidth      uint8
	AniHeight     uint8
	AniDelay1     uint8
	AniDelay2     uint8
	AniDelay3     uint8
	StaticWidth   uint8
	StaticHeight  uint8
	StaticOffsetX uint8
	StaticOffsetY uint8
	MaxScale      uint8
	F1            []byte
	F2            []byte
	F3            []byte
	F4            []byte
}

// LawlessMetaData contains all meta data concerning the Lawless contract.
var LawlessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"RoleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CodexAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addMetadataMod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"addRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"enumRole[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"name\":\"addRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"width\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"height\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniX\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniY\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniWidth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniHeight\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniDelay1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniDelay2\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniDelay3\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticWidth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticHeight\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticOffsetX\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticOffsetY\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxScale\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"f1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"f2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"f3\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"f4\",\"type\":\"bytes\"}],\"internalType\":\"structModel\",\"name\":\"model\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"palette\",\"type\":\"bytes\"}],\"name\":\"animatedGIF\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"authorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint48[]\",\"name\":\"details\",\"type\":\"uint48[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCustomAttributes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paletteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"details\",\"type\":\"uint48\"}],\"internalType\":\"structLawlessData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLawlessId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getModel\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"width\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"height\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniX\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniY\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniWidth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniHeight\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniDelay1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniDelay2\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniDelay3\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticWidth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticHeight\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticOffsetX\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticOffsetY\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxScale\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"f1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"f2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"f3\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"f4\",\"type\":\"bytes\"}],\"internalType\":\"structModel\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getPalette\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"incrementVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"base\",\"type\":\"bool\"}],\"name\":\"lawlessGIF\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"lockRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumRole[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"name\":\"lockRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"details\",\"type\":\"uint48\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"removeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"enumRole[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"name\":\"removeRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"resolverClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"roleLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setB64EncodeURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smashFlask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"width\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"height\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniX\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniY\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniWidth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniHeight\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniDelay1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniDelay2\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"aniDelay3\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticWidth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticHeight\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticOffsetX\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"staticOffsetY\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxScale\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"f1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"f2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"f3\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"f4\",\"type\":\"bytes\"}],\"internalType\":\"structModel\",\"name\":\"model\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"palette\",\"type\":\"bytes\"}],\"name\":\"staticGIF\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"toPctString1000x\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pct\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"base\",\"type\":\"bool\"}],\"name\":\"tokenGIF\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"Ex\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"details\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"incVersion\",\"type\":\"bool\"}],\"name\":\"updateDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"count\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uploadModels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"count\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uploadPalettes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawForeignERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawForeignERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LawlessABI is the input ABI used to generate the binding from.
// Deprecated: Use LawlessMetaData.ABI instead.
var LawlessABI = LawlessMetaData.ABI

// Lawless is an auto generated Go binding around an Ethereum contract.
type Lawless struct {
	LawlessCaller     // Read-only binding to the contract
	LawlessTransactor // Write-only binding to the contract
	LawlessFilterer   // Log filterer for contract events
}

// LawlessCaller is an auto generated read-only Go binding around an Ethereum contract.
type LawlessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LawlessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LawlessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LawlessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LawlessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LawlessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LawlessSession struct {
	Contract     *Lawless          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LawlessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LawlessCallerSession struct {
	Contract *LawlessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LawlessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LawlessTransactorSession struct {
	Contract     *LawlessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LawlessRaw is an auto generated low-level Go binding around an Ethereum contract.
type LawlessRaw struct {
	Contract *Lawless // Generic contract binding to access the raw methods on
}

// LawlessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LawlessCallerRaw struct {
	Contract *LawlessCaller // Generic read-only contract binding to access the raw methods on
}

// LawlessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LawlessTransactorRaw struct {
	Contract *LawlessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLawless creates a new instance of Lawless, bound to a specific deployed contract.
func NewLawless(address common.Address, backend bind.ContractBackend) (*Lawless, error) {
	contract, err := bindLawless(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lawless{LawlessCaller: LawlessCaller{contract: contract}, LawlessTransactor: LawlessTransactor{contract: contract}, LawlessFilterer: LawlessFilterer{contract: contract}}, nil
}

// NewLawlessCaller creates a new read-only instance of Lawless, bound to a specific deployed contract.
func NewLawlessCaller(address common.Address, caller bind.ContractCaller) (*LawlessCaller, error) {
	contract, err := bindLawless(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LawlessCaller{contract: contract}, nil
}

// NewLawlessTransactor creates a new write-only instance of Lawless, bound to a specific deployed contract.
func NewLawlessTransactor(address common.Address, transactor bind.ContractTransactor) (*LawlessTransactor, error) {
	contract, err := bindLawless(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LawlessTransactor{contract: contract}, nil
}

// NewLawlessFilterer creates a new log filterer instance of Lawless, bound to a specific deployed contract.
func NewLawlessFilterer(address common.Address, filterer bind.ContractFilterer) (*LawlessFilterer, error) {
	contract, err := bindLawless(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LawlessFilterer{contract: contract}, nil
}

// bindLawless binds a generic wrapper to an already deployed contract.
func bindLawless(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LawlessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lawless *LawlessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lawless.Contract.LawlessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lawless *LawlessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lawless.Contract.LawlessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lawless *LawlessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lawless.Contract.LawlessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lawless *LawlessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lawless.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lawless *LawlessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lawless.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lawless *LawlessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lawless.Contract.contract.Transact(opts, method, params...)
}

// CodexAddress is a free data retrieval call binding the contract method 0xf9d4b071.
//
// Solidity: function CodexAddress() view returns(address)
func (_Lawless *LawlessCaller) CodexAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "CodexAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CodexAddress is a free data retrieval call binding the contract method 0xf9d4b071.
//
// Solidity: function CodexAddress() view returns(address)
func (_Lawless *LawlessSession) CodexAddress() (common.Address, error) {
	return _Lawless.Contract.CodexAddress(&_Lawless.CallOpts)
}

// CodexAddress is a free data retrieval call binding the contract method 0xf9d4b071.
//
// Solidity: function CodexAddress() view returns(address)
func (_Lawless *LawlessCallerSession) CodexAddress() (common.Address, error) {
	return _Lawless.Contract.CodexAddress(&_Lawless.CallOpts)
}

// AnimatedGIF is a free data retrieval call binding the contract method 0xc1a9b313.
//
// Solidity: function animatedGIF((uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,bytes,bytes,bytes,bytes) model, bytes palette) view returns(string)
func (_Lawless *LawlessCaller) AnimatedGIF(opts *bind.CallOpts, model Model, palette []byte) (string, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "animatedGIF", model, palette)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AnimatedGIF is a free data retrieval call binding the contract method 0xc1a9b313.
//
// Solidity: function animatedGIF((uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,bytes,bytes,bytes,bytes) model, bytes palette) view returns(string)
func (_Lawless *LawlessSession) AnimatedGIF(model Model, palette []byte) (string, error) {
	return _Lawless.Contract.AnimatedGIF(&_Lawless.CallOpts, model, palette)
}

// AnimatedGIF is a free data retrieval call binding the contract method 0xc1a9b313.
//
// Solidity: function animatedGIF((uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,bytes,bytes,bytes,bytes) model, bytes palette) view returns(string)
func (_Lawless *LawlessCallerSession) AnimatedGIF(model Model, palette []byte) (string, error) {
	return _Lawless.Contract.AnimatedGIF(&_Lawless.CallOpts, model, palette)
}

// Authorized is a free data retrieval call binding the contract method 0x5ed7660e.
//
// Solidity: function authorized(address operator, uint256 tokenId) view returns(bool)
func (_Lawless *LawlessCaller) Authorized(opts *bind.CallOpts, operator common.Address, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "authorized", operator, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Authorized is a free data retrieval call binding the contract method 0x5ed7660e.
//
// Solidity: function authorized(address operator, uint256 tokenId) view returns(bool)
func (_Lawless *LawlessSession) Authorized(operator common.Address, tokenId *big.Int) (bool, error) {
	return _Lawless.Contract.Authorized(&_Lawless.CallOpts, operator, tokenId)
}

// Authorized is a free data retrieval call binding the contract method 0x5ed7660e.
//
// Solidity: function authorized(address operator, uint256 tokenId) view returns(bool)
func (_Lawless *LawlessCallerSession) Authorized(operator common.Address, tokenId *big.Int) (bool, error) {
	return _Lawless.Contract.Authorized(&_Lawless.CallOpts, operator, tokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Lawless *LawlessCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Lawless *LawlessSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Lawless.Contract.BalanceOf(&_Lawless.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Lawless *LawlessCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Lawless.Contract.BalanceOf(&_Lawless.CallOpts, owner)
}

// DelegationEnabled is a free data retrieval call binding the contract method 0x54b8c601.
//
// Solidity: function delegationEnabled() view returns(bool)
func (_Lawless *LawlessCaller) DelegationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "delegationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationEnabled is a free data retrieval call binding the contract method 0x54b8c601.
//
// Solidity: function delegationEnabled() view returns(bool)
func (_Lawless *LawlessSession) DelegationEnabled() (bool, error) {
	return _Lawless.Contract.DelegationEnabled(&_Lawless.CallOpts)
}

// DelegationEnabled is a free data retrieval call binding the contract method 0x54b8c601.
//
// Solidity: function delegationEnabled() view returns(bool)
func (_Lawless *LawlessCallerSession) DelegationEnabled() (bool, error) {
	return _Lawless.Contract.DelegationEnabled(&_Lawless.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Lawless *LawlessCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Lawless *LawlessSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Lawless.Contract.GetApproved(&_Lawless.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Lawless *LawlessCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Lawless.Contract.GetApproved(&_Lawless.CallOpts, tokenId)
}

// GetCustomAttributes is a free data retrieval call binding the contract method 0x65990b9a.
//
// Solidity: function getCustomAttributes() view returns(bytes)
func (_Lawless *LawlessCaller) GetCustomAttributes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "getCustomAttributes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCustomAttributes is a free data retrieval call binding the contract method 0x65990b9a.
//
// Solidity: function getCustomAttributes() view returns(bytes)
func (_Lawless *LawlessSession) GetCustomAttributes() ([]byte, error) {
	return _Lawless.Contract.GetCustomAttributes(&_Lawless.CallOpts)
}

// GetCustomAttributes is a free data retrieval call binding the contract method 0x65990b9a.
//
// Solidity: function getCustomAttributes() view returns(bytes)
func (_Lawless *LawlessCallerSession) GetCustomAttributes() ([]byte, error) {
	return _Lawless.Contract.GetCustomAttributes(&_Lawless.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 id) view returns((uint256,uint256,uint256,address,uint16,uint48))
func (_Lawless *LawlessCaller) GetData(opts *bind.CallOpts, id *big.Int) (LawlessData, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "getData", id)

	if err != nil {
		return *new(LawlessData), err
	}

	out0 := *abi.ConvertType(out[0], new(LawlessData)).(*LawlessData)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 id) view returns((uint256,uint256,uint256,address,uint16,uint48))
func (_Lawless *LawlessSession) GetData(id *big.Int) (LawlessData, error) {
	return _Lawless.Contract.GetData(&_Lawless.CallOpts, id)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 id) view returns((uint256,uint256,uint256,address,uint16,uint48))
func (_Lawless *LawlessCallerSession) GetData(id *big.Int) (LawlessData, error) {
	return _Lawless.Contract.GetData(&_Lawless.CallOpts, id)
}

// GetDetails is a free data retrieval call binding the contract method 0xb93a89f7.
//
// Solidity: function getDetails(uint256 tokenId) view returns(uint256, address, uint48)
func (_Lawless *LawlessCaller) GetDetails(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, common.Address, *big.Int, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "getDetails", tokenId)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetDetails is a free data retrieval call binding the contract method 0xb93a89f7.
//
// Solidity: function getDetails(uint256 tokenId) view returns(uint256, address, uint48)
func (_Lawless *LawlessSession) GetDetails(tokenId *big.Int) (*big.Int, common.Address, *big.Int, error) {
	return _Lawless.Contract.GetDetails(&_Lawless.CallOpts, tokenId)
}

// GetDetails is a free data retrieval call binding the contract method 0xb93a89f7.
//
// Solidity: function getDetails(uint256 tokenId) view returns(uint256, address, uint48)
func (_Lawless *LawlessCallerSession) GetDetails(tokenId *big.Int) (*big.Int, common.Address, *big.Int, error) {
	return _Lawless.Contract.GetDetails(&_Lawless.CallOpts, tokenId)
}

// GetLawlessId is a free data retrieval call binding the contract method 0xf9753fa7.
//
// Solidity: function getLawlessId(uint256 tokenId) view returns(uint256 id)
func (_Lawless *LawlessCaller) GetLawlessId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "getLawlessId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLawlessId is a free data retrieval call binding the contract method 0xf9753fa7.
//
// Solidity: function getLawlessId(uint256 tokenId) view returns(uint256 id)
func (_Lawless *LawlessSession) GetLawlessId(tokenId *big.Int) (*big.Int, error) {
	return _Lawless.Contract.GetLawlessId(&_Lawless.CallOpts, tokenId)
}

// GetLawlessId is a free data retrieval call binding the contract method 0xf9753fa7.
//
// Solidity: function getLawlessId(uint256 tokenId) view returns(uint256 id)
func (_Lawless *LawlessCallerSession) GetLawlessId(tokenId *big.Int) (*big.Int, error) {
	return _Lawless.Contract.GetLawlessId(&_Lawless.CallOpts, tokenId)
}

// GetModel is a free data retrieval call binding the contract method 0x6d361694.
//
// Solidity: function getModel(uint256 id) view returns((uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,bytes,bytes,bytes,bytes))
func (_Lawless *LawlessCaller) GetModel(opts *bind.CallOpts, id *big.Int) (Model, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "getModel", id)

	if err != nil {
		return *new(Model), err
	}

	out0 := *abi.ConvertType(out[0], new(Model)).(*Model)

	return out0, err

}

// GetModel is a free data retrieval call binding the contract method 0x6d361694.
//
// Solidity: function getModel(uint256 id) view returns((uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,bytes,bytes,bytes,bytes))
func (_Lawless *LawlessSession) GetModel(id *big.Int) (Model, error) {
	return _Lawless.Contract.GetModel(&_Lawless.CallOpts, id)
}

// GetModel is a free data retrieval call binding the contract method 0x6d361694.
//
// Solidity: function getModel(uint256 id) view returns((uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,bytes,bytes,bytes,bytes))
func (_Lawless *LawlessCallerSession) GetModel(id *big.Int) (Model, error) {
	return _Lawless.Contract.GetModel(&_Lawless.CallOpts, id)
}

// GetPalette is a free data retrieval call binding the contract method 0x505e570a.
//
// Solidity: function getPalette(uint256 id) view returns(bytes)
func (_Lawless *LawlessCaller) GetPalette(opts *bind.CallOpts, id *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "getPalette", id)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPalette is a free data retrieval call binding the contract method 0x505e570a.
//
// Solidity: function getPalette(uint256 id) view returns(bytes)
func (_Lawless *LawlessSession) GetPalette(id *big.Int) ([]byte, error) {
	return _Lawless.Contract.GetPalette(&_Lawless.CallOpts, id)
}

// GetPalette is a free data retrieval call binding the contract method 0x505e570a.
//
// Solidity: function getPalette(uint256 id) view returns(bytes)
func (_Lawless *LawlessCallerSession) GetPalette(id *big.Int) ([]byte, error) {
	return _Lawless.Contract.GetPalette(&_Lawless.CallOpts, id)
}

// GetTokenId is a free data retrieval call binding the contract method 0x14ff5ea3.
//
// Solidity: function getTokenId(uint256 id) view returns(uint256)
func (_Lawless *LawlessCaller) GetTokenId(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "getTokenId", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenId is a free data retrieval call binding the contract method 0x14ff5ea3.
//
// Solidity: function getTokenId(uint256 id) view returns(uint256)
func (_Lawless *LawlessSession) GetTokenId(id *big.Int) (*big.Int, error) {
	return _Lawless.Contract.GetTokenId(&_Lawless.CallOpts, id)
}

// GetTokenId is a free data retrieval call binding the contract method 0x14ff5ea3.
//
// Solidity: function getTokenId(uint256 id) view returns(uint256)
func (_Lawless *LawlessCallerSession) GetTokenId(id *big.Int) (*big.Int, error) {
	return _Lawless.Contract.GetTokenId(&_Lawless.CallOpts, id)
}

// HasRole is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address user, uint8 role) view returns(bool)
func (_Lawless *LawlessCaller) HasRole(opts *bind.CallOpts, user common.Address, role uint8) (bool, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "hasRole", user, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address user, uint8 role) view returns(bool)
func (_Lawless *LawlessSession) HasRole(user common.Address, role uint8) (bool, error) {
	return _Lawless.Contract.HasRole(&_Lawless.CallOpts, user, role)
}

// HasRole is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address user, uint8 role) view returns(bool)
func (_Lawless *LawlessCallerSession) HasRole(user common.Address, role uint8) (bool, error) {
	return _Lawless.Contract.HasRole(&_Lawless.CallOpts, user, role)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Lawless *LawlessCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Lawless *LawlessSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Lawless.Contract.IsApprovedForAll(&_Lawless.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Lawless *LawlessCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Lawless.Contract.IsApprovedForAll(&_Lawless.CallOpts, owner, operator)
}

// LawlessGIF is a free data retrieval call binding the contract method 0x25f399cb.
//
// Solidity: function lawlessGIF(uint256 id, bool base) view returns(string)
func (_Lawless *LawlessCaller) LawlessGIF(opts *bind.CallOpts, id *big.Int, base bool) (string, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "lawlessGIF", id, base)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LawlessGIF is a free data retrieval call binding the contract method 0x25f399cb.
//
// Solidity: function lawlessGIF(uint256 id, bool base) view returns(string)
func (_Lawless *LawlessSession) LawlessGIF(id *big.Int, base bool) (string, error) {
	return _Lawless.Contract.LawlessGIF(&_Lawless.CallOpts, id, base)
}

// LawlessGIF is a free data retrieval call binding the contract method 0x25f399cb.
//
// Solidity: function lawlessGIF(uint256 id, bool base) view returns(string)
func (_Lawless *LawlessCallerSession) LawlessGIF(id *big.Int, base bool) (string, error) {
	return _Lawless.Contract.LawlessGIF(&_Lawless.CallOpts, id, base)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lawless *LawlessCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lawless *LawlessSession) Name() (string, error) {
	return _Lawless.Contract.Name(&_Lawless.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lawless *LawlessCallerSession) Name() (string, error) {
	return _Lawless.Contract.Name(&_Lawless.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lawless *LawlessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lawless *LawlessSession) Owner() (common.Address, error) {
	return _Lawless.Contract.Owner(&_Lawless.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lawless *LawlessCallerSession) Owner() (common.Address, error) {
	return _Lawless.Contract.Owner(&_Lawless.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Lawless *LawlessCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Lawless *LawlessSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Lawless.Contract.OwnerOf(&_Lawless.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Lawless *LawlessCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Lawless.Contract.OwnerOf(&_Lawless.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lawless *LawlessCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lawless *LawlessSession) Paused() (bool, error) {
	return _Lawless.Contract.Paused(&_Lawless.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lawless *LawlessCallerSession) Paused() (bool, error) {
	return _Lawless.Contract.Paused(&_Lawless.CallOpts)
}

// RoleLocked is a free data retrieval call binding the contract method 0x3c4bec2a.
//
// Solidity: function roleLocked(uint8 role) view returns(bool)
func (_Lawless *LawlessCaller) RoleLocked(opts *bind.CallOpts, role uint8) (bool, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "roleLocked", role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RoleLocked is a free data retrieval call binding the contract method 0x3c4bec2a.
//
// Solidity: function roleLocked(uint8 role) view returns(bool)
func (_Lawless *LawlessSession) RoleLocked(role uint8) (bool, error) {
	return _Lawless.Contract.RoleLocked(&_Lawless.CallOpts, role)
}

// RoleLocked is a free data retrieval call binding the contract method 0x3c4bec2a.
//
// Solidity: function roleLocked(uint8 role) view returns(bool)
func (_Lawless *LawlessCallerSession) RoleLocked(role uint8) (bool, error) {
	return _Lawless.Contract.RoleLocked(&_Lawless.CallOpts, role)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 ) view returns(address, uint256)
func (_Lawless *LawlessCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "royaltyInfo", arg0, arg1)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 ) view returns(address, uint256)
func (_Lawless *LawlessSession) RoyaltyInfo(arg0 *big.Int, arg1 *big.Int) (common.Address, *big.Int, error) {
	return _Lawless.Contract.RoyaltyInfo(&_Lawless.CallOpts, arg0, arg1)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 ) view returns(address, uint256)
func (_Lawless *LawlessCallerSession) RoyaltyInfo(arg0 *big.Int, arg1 *big.Int) (common.Address, *big.Int, error) {
	return _Lawless.Contract.RoyaltyInfo(&_Lawless.CallOpts, arg0, arg1)
}

// StaticGIF is a free data retrieval call binding the contract method 0x5b90e10c.
//
// Solidity: function staticGIF((uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,bytes,bytes,bytes,bytes) model, bytes palette) view returns(string)
func (_Lawless *LawlessCaller) StaticGIF(opts *bind.CallOpts, model Model, palette []byte) (string, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "staticGIF", model, palette)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StaticGIF is a free data retrieval call binding the contract method 0x5b90e10c.
//
// Solidity: function staticGIF((uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,bytes,bytes,bytes,bytes) model, bytes palette) view returns(string)
func (_Lawless *LawlessSession) StaticGIF(model Model, palette []byte) (string, error) {
	return _Lawless.Contract.StaticGIF(&_Lawless.CallOpts, model, palette)
}

// StaticGIF is a free data retrieval call binding the contract method 0x5b90e10c.
//
// Solidity: function staticGIF((uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,bytes,bytes,bytes,bytes) model, bytes palette) view returns(string)
func (_Lawless *LawlessCallerSession) StaticGIF(model Model, palette []byte) (string, error) {
	return _Lawless.Contract.StaticGIF(&_Lawless.CallOpts, model, palette)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lawless *LawlessCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lawless *LawlessSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lawless.Contract.SupportsInterface(&_Lawless.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lawless *LawlessCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lawless.Contract.SupportsInterface(&_Lawless.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lawless *LawlessCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lawless *LawlessSession) Symbol() (string, error) {
	return _Lawless.Contract.Symbol(&_Lawless.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lawless *LawlessCallerSession) Symbol() (string, error) {
	return _Lawless.Contract.Symbol(&_Lawless.CallOpts)
}

// ToPctString1000x is a free data retrieval call binding the contract method 0xeb5a2680.
//
// Solidity: function toPctString1000x(uint256 value) pure returns(string pct)
func (_Lawless *LawlessCaller) ToPctString1000x(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "toPctString1000x", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToPctString1000x is a free data retrieval call binding the contract method 0xeb5a2680.
//
// Solidity: function toPctString1000x(uint256 value) pure returns(string pct)
func (_Lawless *LawlessSession) ToPctString1000x(value *big.Int) (string, error) {
	return _Lawless.Contract.ToPctString1000x(&_Lawless.CallOpts, value)
}

// ToPctString1000x is a free data retrieval call binding the contract method 0xeb5a2680.
//
// Solidity: function toPctString1000x(uint256 value) pure returns(string pct)
func (_Lawless *LawlessCallerSession) ToPctString1000x(value *big.Int) (string, error) {
	return _Lawless.Contract.ToPctString1000x(&_Lawless.CallOpts, value)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Lawless *LawlessCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Lawless *LawlessSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Lawless.Contract.TokenByIndex(&_Lawless.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Lawless *LawlessCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Lawless.Contract.TokenByIndex(&_Lawless.CallOpts, index)
}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool exists)
func (_Lawless *LawlessCaller) TokenExists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "tokenExists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool exists)
func (_Lawless *LawlessSession) TokenExists(tokenId *big.Int) (bool, error) {
	return _Lawless.Contract.TokenExists(&_Lawless.CallOpts, tokenId)
}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool exists)
func (_Lawless *LawlessCallerSession) TokenExists(tokenId *big.Int) (bool, error) {
	return _Lawless.Contract.TokenExists(&_Lawless.CallOpts, tokenId)
}

// TokenGIF is a free data retrieval call binding the contract method 0x7383a066.
//
// Solidity: function tokenGIF(uint256 tokenId, bool base) view returns(string)
func (_Lawless *LawlessCaller) TokenGIF(opts *bind.CallOpts, tokenId *big.Int, base bool) (string, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "tokenGIF", tokenId, base)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenGIF is a free data retrieval call binding the contract method 0x7383a066.
//
// Solidity: function tokenGIF(uint256 tokenId, bool base) view returns(string)
func (_Lawless *LawlessSession) TokenGIF(tokenId *big.Int, base bool) (string, error) {
	return _Lawless.Contract.TokenGIF(&_Lawless.CallOpts, tokenId, base)
}

// TokenGIF is a free data retrieval call binding the contract method 0x7383a066.
//
// Solidity: function tokenGIF(uint256 tokenId, bool base) view returns(string)
func (_Lawless *LawlessCallerSession) TokenGIF(tokenId *big.Int, base bool) (string, error) {
	return _Lawless.Contract.TokenGIF(&_Lawless.CallOpts, tokenId, base)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_Lawless *LawlessCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_Lawless *LawlessSession) TokenId() (*big.Int, error) {
	return _Lawless.Contract.TokenId(&_Lawless.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_Lawless *LawlessCallerSession) TokenId() (*big.Int, error) {
	return _Lawless.Contract.TokenId(&_Lawless.CallOpts)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Lawless *LawlessCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Lawless *LawlessSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Lawless.Contract.TokenOfOwnerByIndex(&_Lawless.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Lawless *LawlessCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Lawless.Contract.TokenOfOwnerByIndex(&_Lawless.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Lawless *LawlessCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Lawless *LawlessSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Lawless.Contract.TokenURI(&_Lawless.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Lawless *LawlessCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Lawless.Contract.TokenURI(&_Lawless.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lawless *LawlessCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lawless.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lawless *LawlessSession) TotalSupply() (*big.Int, error) {
	return _Lawless.Contract.TotalSupply(&_Lawless.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lawless *LawlessCallerSession) TotalSupply() (*big.Int, error) {
	return _Lawless.Contract.TotalSupply(&_Lawless.CallOpts)
}

// AddMetadataMod is a paid mutator transaction binding the contract method 0x90fab70b.
//
// Solidity: function addMetadataMod(address addr) returns()
func (_Lawless *LawlessTransactor) AddMetadataMod(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "addMetadataMod", addr)
}

// AddMetadataMod is a paid mutator transaction binding the contract method 0x90fab70b.
//
// Solidity: function addMetadataMod(address addr) returns()
func (_Lawless *LawlessSession) AddMetadataMod(addr common.Address) (*types.Transaction, error) {
	return _Lawless.Contract.AddMetadataMod(&_Lawless.TransactOpts, addr)
}

// AddMetadataMod is a paid mutator transaction binding the contract method 0x90fab70b.
//
// Solidity: function addMetadataMod(address addr) returns()
func (_Lawless *LawlessTransactorSession) AddMetadataMod(addr common.Address) (*types.Transaction, error) {
	return _Lawless.Contract.AddMetadataMod(&_Lawless.TransactOpts, addr)
}

// AddRole is a paid mutator transaction binding the contract method 0x44deb6f3.
//
// Solidity: function addRole(address user, uint8 role) returns()
func (_Lawless *LawlessTransactor) AddRole(opts *bind.TransactOpts, user common.Address, role uint8) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "addRole", user, role)
}

// AddRole is a paid mutator transaction binding the contract method 0x44deb6f3.
//
// Solidity: function addRole(address user, uint8 role) returns()
func (_Lawless *LawlessSession) AddRole(user common.Address, role uint8) (*types.Transaction, error) {
	return _Lawless.Contract.AddRole(&_Lawless.TransactOpts, user, role)
}

// AddRole is a paid mutator transaction binding the contract method 0x44deb6f3.
//
// Solidity: function addRole(address user, uint8 role) returns()
func (_Lawless *LawlessTransactorSession) AddRole(user common.Address, role uint8) (*types.Transaction, error) {
	return _Lawless.Contract.AddRole(&_Lawless.TransactOpts, user, role)
}

// AddRoles is a paid mutator transaction binding the contract method 0xe04388c4.
//
// Solidity: function addRoles(address[] users, uint8[] roles) returns()
func (_Lawless *LawlessTransactor) AddRoles(opts *bind.TransactOpts, users []common.Address, roles []uint8) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "addRoles", users, roles)
}

// AddRoles is a paid mutator transaction binding the contract method 0xe04388c4.
//
// Solidity: function addRoles(address[] users, uint8[] roles) returns()
func (_Lawless *LawlessSession) AddRoles(users []common.Address, roles []uint8) (*types.Transaction, error) {
	return _Lawless.Contract.AddRoles(&_Lawless.TransactOpts, users, roles)
}

// AddRoles is a paid mutator transaction binding the contract method 0xe04388c4.
//
// Solidity: function addRoles(address[] users, uint8[] roles) returns()
func (_Lawless *LawlessTransactorSession) AddRoles(users []common.Address, roles []uint8) (*types.Transaction, error) {
	return _Lawless.Contract.AddRoles(&_Lawless.TransactOpts, users, roles)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Lawless *LawlessTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Lawless *LawlessSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.Approve(&_Lawless.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Lawless *LawlessTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.Approve(&_Lawless.TransactOpts, to, tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb892e7ea.
//
// Solidity: function batchMint(uint256 seed, address[] to, uint48[] details) returns()
func (_Lawless *LawlessTransactor) BatchMint(opts *bind.TransactOpts, seed *big.Int, to []common.Address, details []*big.Int) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "batchMint", seed, to, details)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb892e7ea.
//
// Solidity: function batchMint(uint256 seed, address[] to, uint48[] details) returns()
func (_Lawless *LawlessSession) BatchMint(seed *big.Int, to []common.Address, details []*big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.BatchMint(&_Lawless.TransactOpts, seed, to, details)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb892e7ea.
//
// Solidity: function batchMint(uint256 seed, address[] to, uint48[] details) returns()
func (_Lawless *LawlessTransactorSession) BatchMint(seed *big.Int, to []common.Address, details []*big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.BatchMint(&_Lawless.TransactOpts, seed, to, details)
}

// IncrementVersion is a paid mutator transaction binding the contract method 0xb5e1ec54.
//
// Solidity: function incrementVersion(address operator, uint256 id) returns()
func (_Lawless *LawlessTransactor) IncrementVersion(opts *bind.TransactOpts, operator common.Address, id *big.Int) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "incrementVersion", operator, id)
}

// IncrementVersion is a paid mutator transaction binding the contract method 0xb5e1ec54.
//
// Solidity: function incrementVersion(address operator, uint256 id) returns()
func (_Lawless *LawlessSession) IncrementVersion(operator common.Address, id *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.IncrementVersion(&_Lawless.TransactOpts, operator, id)
}

// IncrementVersion is a paid mutator transaction binding the contract method 0xb5e1ec54.
//
// Solidity: function incrementVersion(address operator, uint256 id) returns()
func (_Lawless *LawlessTransactorSession) IncrementVersion(operator common.Address, id *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.IncrementVersion(&_Lawless.TransactOpts, operator, id)
}

// LockRole is a paid mutator transaction binding the contract method 0xcc162102.
//
// Solidity: function lockRole(uint8 role) returns()
func (_Lawless *LawlessTransactor) LockRole(opts *bind.TransactOpts, role uint8) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "lockRole", role)
}

// LockRole is a paid mutator transaction binding the contract method 0xcc162102.
//
// Solidity: function lockRole(uint8 role) returns()
func (_Lawless *LawlessSession) LockRole(role uint8) (*types.Transaction, error) {
	return _Lawless.Contract.LockRole(&_Lawless.TransactOpts, role)
}

// LockRole is a paid mutator transaction binding the contract method 0xcc162102.
//
// Solidity: function lockRole(uint8 role) returns()
func (_Lawless *LawlessTransactorSession) LockRole(role uint8) (*types.Transaction, error) {
	return _Lawless.Contract.LockRole(&_Lawless.TransactOpts, role)
}

// LockRoles is a paid mutator transaction binding the contract method 0x550bb2f1.
//
// Solidity: function lockRoles(uint8[] roles) returns()
func (_Lawless *LawlessTransactor) LockRoles(opts *bind.TransactOpts, roles []uint8) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "lockRoles", roles)
}

// LockRoles is a paid mutator transaction binding the contract method 0x550bb2f1.
//
// Solidity: function lockRoles(uint8[] roles) returns()
func (_Lawless *LawlessSession) LockRoles(roles []uint8) (*types.Transaction, error) {
	return _Lawless.Contract.LockRoles(&_Lawless.TransactOpts, roles)
}

// LockRoles is a paid mutator transaction binding the contract method 0x550bb2f1.
//
// Solidity: function lockRoles(uint8[] roles) returns()
func (_Lawless *LawlessTransactorSession) LockRoles(roles []uint8) (*types.Transaction, error) {
	return _Lawless.Contract.LockRoles(&_Lawless.TransactOpts, roles)
}

// Mint is a paid mutator transaction binding the contract method 0x1f2c02f7.
//
// Solidity: function mint(uint256 seed, address to, uint48 details) returns()
func (_Lawless *LawlessTransactor) Mint(opts *bind.TransactOpts, seed *big.Int, to common.Address, details *big.Int) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "mint", seed, to, details)
}

// Mint is a paid mutator transaction binding the contract method 0x1f2c02f7.
//
// Solidity: function mint(uint256 seed, address to, uint48 details) returns()
func (_Lawless *LawlessSession) Mint(seed *big.Int, to common.Address, details *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.Mint(&_Lawless.TransactOpts, seed, to, details)
}

// Mint is a paid mutator transaction binding the contract method 0x1f2c02f7.
//
// Solidity: function mint(uint256 seed, address to, uint48 details) returns()
func (_Lawless *LawlessTransactorSession) Mint(seed *big.Int, to common.Address, details *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.Mint(&_Lawless.TransactOpts, seed, to, details)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lawless *LawlessTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lawless *LawlessSession) Pause() (*types.Transaction, error) {
	return _Lawless.Contract.Pause(&_Lawless.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lawless *LawlessTransactorSession) Pause() (*types.Transaction, error) {
	return _Lawless.Contract.Pause(&_Lawless.TransactOpts)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xb74866fb.
//
// Solidity: function removeRole(address user, uint8 role) returns()
func (_Lawless *LawlessTransactor) RemoveRole(opts *bind.TransactOpts, user common.Address, role uint8) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "removeRole", user, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xb74866fb.
//
// Solidity: function removeRole(address user, uint8 role) returns()
func (_Lawless *LawlessSession) RemoveRole(user common.Address, role uint8) (*types.Transaction, error) {
	return _Lawless.Contract.RemoveRole(&_Lawless.TransactOpts, user, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xb74866fb.
//
// Solidity: function removeRole(address user, uint8 role) returns()
func (_Lawless *LawlessTransactorSession) RemoveRole(user common.Address, role uint8) (*types.Transaction, error) {
	return _Lawless.Contract.RemoveRole(&_Lawless.TransactOpts, user, role)
}

// RemoveRoles is a paid mutator transaction binding the contract method 0xd84a5223.
//
// Solidity: function removeRoles(address[] users, uint8[] roles) returns()
func (_Lawless *LawlessTransactor) RemoveRoles(opts *bind.TransactOpts, users []common.Address, roles []uint8) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "removeRoles", users, roles)
}

// RemoveRoles is a paid mutator transaction binding the contract method 0xd84a5223.
//
// Solidity: function removeRoles(address[] users, uint8[] roles) returns()
func (_Lawless *LawlessSession) RemoveRoles(users []common.Address, roles []uint8) (*types.Transaction, error) {
	return _Lawless.Contract.RemoveRoles(&_Lawless.TransactOpts, users, roles)
}

// RemoveRoles is a paid mutator transaction binding the contract method 0xd84a5223.
//
// Solidity: function removeRoles(address[] users, uint8[] roles) returns()
func (_Lawless *LawlessTransactorSession) RemoveRoles(users []common.Address, roles []uint8) (*types.Transaction, error) {
	return _Lawless.Contract.RemoveRoles(&_Lawless.TransactOpts, users, roles)
}

// ResolverClaim is a paid mutator transaction binding the contract method 0x34fae557.
//
// Solidity: function resolverClaim(address newOwner) returns()
func (_Lawless *LawlessTransactor) ResolverClaim(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "resolverClaim", newOwner)
}

// ResolverClaim is a paid mutator transaction binding the contract method 0x34fae557.
//
// Solidity: function resolverClaim(address newOwner) returns()
func (_Lawless *LawlessSession) ResolverClaim(newOwner common.Address) (*types.Transaction, error) {
	return _Lawless.Contract.ResolverClaim(&_Lawless.TransactOpts, newOwner)
}

// ResolverClaim is a paid mutator transaction binding the contract method 0x34fae557.
//
// Solidity: function resolverClaim(address newOwner) returns()
func (_Lawless *LawlessTransactorSession) ResolverClaim(newOwner common.Address) (*types.Transaction, error) {
	return _Lawless.Contract.ResolverClaim(&_Lawless.TransactOpts, newOwner)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Lawless *LawlessTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Lawless *LawlessSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.SafeTransferFrom(&_Lawless.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Lawless *LawlessTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.SafeTransferFrom(&_Lawless.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Lawless *LawlessTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Lawless *LawlessSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Lawless.Contract.SafeTransferFrom0(&_Lawless.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Lawless *LawlessTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Lawless.Contract.SafeTransferFrom0(&_Lawless.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Lawless *LawlessTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Lawless *LawlessSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Lawless.Contract.SetApprovalForAll(&_Lawless.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Lawless *LawlessTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Lawless.Contract.SetApprovalForAll(&_Lawless.TransactOpts, operator, approved)
}

// SetB64EncodeURI is a paid mutator transaction binding the contract method 0x94e83b9b.
//
// Solidity: function setB64EncodeURI(bool value) returns()
func (_Lawless *LawlessTransactor) SetB64EncodeURI(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "setB64EncodeURI", value)
}

// SetB64EncodeURI is a paid mutator transaction binding the contract method 0x94e83b9b.
//
// Solidity: function setB64EncodeURI(bool value) returns()
func (_Lawless *LawlessSession) SetB64EncodeURI(value bool) (*types.Transaction, error) {
	return _Lawless.Contract.SetB64EncodeURI(&_Lawless.TransactOpts, value)
}

// SetB64EncodeURI is a paid mutator transaction binding the contract method 0x94e83b9b.
//
// Solidity: function setB64EncodeURI(bool value) returns()
func (_Lawless *LawlessTransactorSession) SetB64EncodeURI(value bool) (*types.Transaction, error) {
	return _Lawless.Contract.SetB64EncodeURI(&_Lawless.TransactOpts, value)
}

// SetMetadata is a paid mutator transaction binding the contract method 0xee57e36f.
//
// Solidity: function setMetadata(bytes metadata) returns()
func (_Lawless *LawlessTransactor) SetMetadata(opts *bind.TransactOpts, metadata []byte) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "setMetadata", metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0xee57e36f.
//
// Solidity: function setMetadata(bytes metadata) returns()
func (_Lawless *LawlessSession) SetMetadata(metadata []byte) (*types.Transaction, error) {
	return _Lawless.Contract.SetMetadata(&_Lawless.TransactOpts, metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0xee57e36f.
//
// Solidity: function setMetadata(bytes metadata) returns()
func (_Lawless *LawlessTransactorSession) SetMetadata(metadata []byte) (*types.Transaction, error) {
	return _Lawless.Contract.SetMetadata(&_Lawless.TransactOpts, metadata)
}

// SmashFlask is a paid mutator transaction binding the contract method 0xaf7460f7.
//
// Solidity: function smashFlask() returns()
func (_Lawless *LawlessTransactor) SmashFlask(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "smashFlask")
}

// SmashFlask is a paid mutator transaction binding the contract method 0xaf7460f7.
//
// Solidity: function smashFlask() returns()
func (_Lawless *LawlessSession) SmashFlask() (*types.Transaction, error) {
	return _Lawless.Contract.SmashFlask(&_Lawless.TransactOpts)
}

// SmashFlask is a paid mutator transaction binding the contract method 0xaf7460f7.
//
// Solidity: function smashFlask() returns()
func (_Lawless *LawlessTransactorSession) SmashFlask() (*types.Transaction, error) {
	return _Lawless.Contract.SmashFlask(&_Lawless.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Lawless *LawlessTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Lawless *LawlessSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.TransferFrom(&_Lawless.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Lawless *LawlessTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.TransferFrom(&_Lawless.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lawless *LawlessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lawless *LawlessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lawless.Contract.TransferOwnership(&_Lawless.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lawless *LawlessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lawless.Contract.TransferOwnership(&_Lawless.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lawless *LawlessTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lawless *LawlessSession) Unpause() (*types.Transaction, error) {
	return _Lawless.Contract.Unpause(&_Lawless.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lawless *LawlessTransactorSession) Unpause() (*types.Transaction, error) {
	return _Lawless.Contract.Unpause(&_Lawless.TransactOpts)
}

// UpdateDetails is a paid mutator transaction binding the contract method 0x104907e8.
//
// Solidity: function updateDetails(address operator, uint256 id, uint48 details, bool incVersion) returns()
func (_Lawless *LawlessTransactor) UpdateDetails(opts *bind.TransactOpts, operator common.Address, id *big.Int, details *big.Int, incVersion bool) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "updateDetails", operator, id, details, incVersion)
}

// UpdateDetails is a paid mutator transaction binding the contract method 0x104907e8.
//
// Solidity: function updateDetails(address operator, uint256 id, uint48 details, bool incVersion) returns()
func (_Lawless *LawlessSession) UpdateDetails(operator common.Address, id *big.Int, details *big.Int, incVersion bool) (*types.Transaction, error) {
	return _Lawless.Contract.UpdateDetails(&_Lawless.TransactOpts, operator, id, details, incVersion)
}

// UpdateDetails is a paid mutator transaction binding the contract method 0x104907e8.
//
// Solidity: function updateDetails(address operator, uint256 id, uint48 details, bool incVersion) returns()
func (_Lawless *LawlessTransactorSession) UpdateDetails(operator common.Address, id *big.Int, details *big.Int, incVersion bool) (*types.Transaction, error) {
	return _Lawless.Contract.UpdateDetails(&_Lawless.TransactOpts, operator, id, details, incVersion)
}

// UploadModels is a paid mutator transaction binding the contract method 0xb12c0299.
//
// Solidity: function uploadModels(uint48 count, bytes data) returns()
func (_Lawless *LawlessTransactor) UploadModels(opts *bind.TransactOpts, count *big.Int, data []byte) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "uploadModels", count, data)
}

// UploadModels is a paid mutator transaction binding the contract method 0xb12c0299.
//
// Solidity: function uploadModels(uint48 count, bytes data) returns()
func (_Lawless *LawlessSession) UploadModels(count *big.Int, data []byte) (*types.Transaction, error) {
	return _Lawless.Contract.UploadModels(&_Lawless.TransactOpts, count, data)
}

// UploadModels is a paid mutator transaction binding the contract method 0xb12c0299.
//
// Solidity: function uploadModels(uint48 count, bytes data) returns()
func (_Lawless *LawlessTransactorSession) UploadModels(count *big.Int, data []byte) (*types.Transaction, error) {
	return _Lawless.Contract.UploadModels(&_Lawless.TransactOpts, count, data)
}

// UploadPalettes is a paid mutator transaction binding the contract method 0x1ef9744c.
//
// Solidity: function uploadPalettes(uint48 count, bytes data) returns()
func (_Lawless *LawlessTransactor) UploadPalettes(opts *bind.TransactOpts, count *big.Int, data []byte) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "uploadPalettes", count, data)
}

// UploadPalettes is a paid mutator transaction binding the contract method 0x1ef9744c.
//
// Solidity: function uploadPalettes(uint48 count, bytes data) returns()
func (_Lawless *LawlessSession) UploadPalettes(count *big.Int, data []byte) (*types.Transaction, error) {
	return _Lawless.Contract.UploadPalettes(&_Lawless.TransactOpts, count, data)
}

// UploadPalettes is a paid mutator transaction binding the contract method 0x1ef9744c.
//
// Solidity: function uploadPalettes(uint48 count, bytes data) returns()
func (_Lawless *LawlessTransactorSession) UploadPalettes(count *big.Int, data []byte) (*types.Transaction, error) {
	return _Lawless.Contract.UploadPalettes(&_Lawless.TransactOpts, count, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Lawless *LawlessTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Lawless *LawlessSession) Withdraw() (*types.Transaction, error) {
	return _Lawless.Contract.Withdraw(&_Lawless.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Lawless *LawlessTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Lawless.Contract.Withdraw(&_Lawless.TransactOpts)
}

// WithdrawForeignERC20 is a paid mutator transaction binding the contract method 0x5c471995.
//
// Solidity: function withdrawForeignERC20(address tokenContract) returns()
func (_Lawless *LawlessTransactor) WithdrawForeignERC20(opts *bind.TransactOpts, tokenContract common.Address) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "withdrawForeignERC20", tokenContract)
}

// WithdrawForeignERC20 is a paid mutator transaction binding the contract method 0x5c471995.
//
// Solidity: function withdrawForeignERC20(address tokenContract) returns()
func (_Lawless *LawlessSession) WithdrawForeignERC20(tokenContract common.Address) (*types.Transaction, error) {
	return _Lawless.Contract.WithdrawForeignERC20(&_Lawless.TransactOpts, tokenContract)
}

// WithdrawForeignERC20 is a paid mutator transaction binding the contract method 0x5c471995.
//
// Solidity: function withdrawForeignERC20(address tokenContract) returns()
func (_Lawless *LawlessTransactorSession) WithdrawForeignERC20(tokenContract common.Address) (*types.Transaction, error) {
	return _Lawless.Contract.WithdrawForeignERC20(&_Lawless.TransactOpts, tokenContract)
}

// WithdrawForeignERC721 is a paid mutator transaction binding the contract method 0x0ce06b68.
//
// Solidity: function withdrawForeignERC721(address tokenContract, uint256 _tokenId) returns()
func (_Lawless *LawlessTransactor) WithdrawForeignERC721(opts *bind.TransactOpts, tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.contract.Transact(opts, "withdrawForeignERC721", tokenContract, _tokenId)
}

// WithdrawForeignERC721 is a paid mutator transaction binding the contract method 0x0ce06b68.
//
// Solidity: function withdrawForeignERC721(address tokenContract, uint256 _tokenId) returns()
func (_Lawless *LawlessSession) WithdrawForeignERC721(tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.WithdrawForeignERC721(&_Lawless.TransactOpts, tokenContract, _tokenId)
}

// WithdrawForeignERC721 is a paid mutator transaction binding the contract method 0x0ce06b68.
//
// Solidity: function withdrawForeignERC721(address tokenContract, uint256 _tokenId) returns()
func (_Lawless *LawlessTransactorSession) WithdrawForeignERC721(tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Lawless.Contract.WithdrawForeignERC721(&_Lawless.TransactOpts, tokenContract, _tokenId)
}

// LawlessApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Lawless contract.
type LawlessApprovalIterator struct {
	Event *LawlessApproval // Event containing the contract specifics and raw log

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
func (it *LawlessApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LawlessApproval)
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
		it.Event = new(LawlessApproval)
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
func (it *LawlessApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LawlessApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LawlessApproval represents a Approval event raised by the Lawless contract.
type LawlessApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Lawless *LawlessFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LawlessApprovalIterator, error) {

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

	logs, sub, err := _Lawless.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LawlessApprovalIterator{contract: _Lawless.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Lawless *LawlessFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LawlessApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Lawless.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LawlessApproval)
				if err := _Lawless.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Lawless *LawlessFilterer) ParseApproval(log types.Log) (*LawlessApproval, error) {
	event := new(LawlessApproval)
	if err := _Lawless.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LawlessApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Lawless contract.
type LawlessApprovalForAllIterator struct {
	Event *LawlessApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LawlessApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LawlessApprovalForAll)
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
		it.Event = new(LawlessApprovalForAll)
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
func (it *LawlessApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LawlessApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LawlessApprovalForAll represents a ApprovalForAll event raised by the Lawless contract.
type LawlessApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Lawless *LawlessFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LawlessApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Lawless.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LawlessApprovalForAllIterator{contract: _Lawless.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Lawless *LawlessFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LawlessApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Lawless.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LawlessApprovalForAll)
				if err := _Lawless.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Lawless *LawlessFilterer) ParseApprovalForAll(log types.Log) (*LawlessApprovalForAll, error) {
	event := new(LawlessApprovalForAll)
	if err := _Lawless.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LawlessPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Lawless contract.
type LawlessPausedIterator struct {
	Event *LawlessPaused // Event containing the contract specifics and raw log

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
func (it *LawlessPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LawlessPaused)
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
		it.Event = new(LawlessPaused)
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
func (it *LawlessPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LawlessPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LawlessPaused represents a Paused event raised by the Lawless contract.
type LawlessPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lawless *LawlessFilterer) FilterPaused(opts *bind.FilterOpts) (*LawlessPausedIterator, error) {

	logs, sub, err := _Lawless.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LawlessPausedIterator{contract: _Lawless.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lawless *LawlessFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LawlessPaused) (event.Subscription, error) {

	logs, sub, err := _Lawless.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LawlessPaused)
				if err := _Lawless.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Lawless *LawlessFilterer) ParsePaused(log types.Log) (*LawlessPaused, error) {
	event := new(LawlessPaused)
	if err := _Lawless.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LawlessRoleLockedIterator is returned from FilterRoleLocked and is used to iterate over the raw logs and unpacked data for RoleLocked events raised by the Lawless contract.
type LawlessRoleLockedIterator struct {
	Event *LawlessRoleLocked // Event containing the contract specifics and raw log

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
func (it *LawlessRoleLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LawlessRoleLocked)
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
		it.Event = new(LawlessRoleLocked)
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
func (it *LawlessRoleLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LawlessRoleLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LawlessRoleLocked represents a RoleLocked event raised by the Lawless contract.
type LawlessRoleLocked struct {
	Role uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRoleLocked is a free log retrieval operation binding the contract event 0x2a4c011a3eb263cc32455711057131757df48518d49547f77ab832d8681a53ff.
//
// Solidity: event RoleLocked(uint8 indexed role)
func (_Lawless *LawlessFilterer) FilterRoleLocked(opts *bind.FilterOpts, role []uint8) (*LawlessRoleLockedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Lawless.contract.FilterLogs(opts, "RoleLocked", roleRule)
	if err != nil {
		return nil, err
	}
	return &LawlessRoleLockedIterator{contract: _Lawless.contract, event: "RoleLocked", logs: logs, sub: sub}, nil
}

// WatchRoleLocked is a free log subscription operation binding the contract event 0x2a4c011a3eb263cc32455711057131757df48518d49547f77ab832d8681a53ff.
//
// Solidity: event RoleLocked(uint8 indexed role)
func (_Lawless *LawlessFilterer) WatchRoleLocked(opts *bind.WatchOpts, sink chan<- *LawlessRoleLocked, role []uint8) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Lawless.contract.WatchLogs(opts, "RoleLocked", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LawlessRoleLocked)
				if err := _Lawless.contract.UnpackLog(event, "RoleLocked", log); err != nil {
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

// ParseRoleLocked is a log parse operation binding the contract event 0x2a4c011a3eb263cc32455711057131757df48518d49547f77ab832d8681a53ff.
//
// Solidity: event RoleLocked(uint8 indexed role)
func (_Lawless *LawlessFilterer) ParseRoleLocked(log types.Log) (*LawlessRoleLocked, error) {
	event := new(LawlessRoleLocked)
	if err := _Lawless.contract.UnpackLog(event, "RoleLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LawlessRoleUpdatedIterator is returned from FilterRoleUpdated and is used to iterate over the raw logs and unpacked data for RoleUpdated events raised by the Lawless contract.
type LawlessRoleUpdatedIterator struct {
	Event *LawlessRoleUpdated // Event containing the contract specifics and raw log

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
func (it *LawlessRoleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LawlessRoleUpdated)
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
		it.Event = new(LawlessRoleUpdated)
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
func (it *LawlessRoleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LawlessRoleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LawlessRoleUpdated represents a RoleUpdated event raised by the Lawless contract.
type LawlessRoleUpdated struct {
	User    common.Address
	Role    uint8
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleUpdated is a free log retrieval operation binding the contract event 0x25cef78a4b665fcf0923e3f0e15c3449d405e0fa46d991b78246e8aaf19571ab.
//
// Solidity: event RoleUpdated(address indexed user, uint8 indexed role, bool enabled)
func (_Lawless *LawlessFilterer) FilterRoleUpdated(opts *bind.FilterOpts, user []common.Address, role []uint8) (*LawlessRoleUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Lawless.contract.FilterLogs(opts, "RoleUpdated", userRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &LawlessRoleUpdatedIterator{contract: _Lawless.contract, event: "RoleUpdated", logs: logs, sub: sub}, nil
}

// WatchRoleUpdated is a free log subscription operation binding the contract event 0x25cef78a4b665fcf0923e3f0e15c3449d405e0fa46d991b78246e8aaf19571ab.
//
// Solidity: event RoleUpdated(address indexed user, uint8 indexed role, bool enabled)
func (_Lawless *LawlessFilterer) WatchRoleUpdated(opts *bind.WatchOpts, sink chan<- *LawlessRoleUpdated, user []common.Address, role []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Lawless.contract.WatchLogs(opts, "RoleUpdated", userRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LawlessRoleUpdated)
				if err := _Lawless.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
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

// ParseRoleUpdated is a log parse operation binding the contract event 0x25cef78a4b665fcf0923e3f0e15c3449d405e0fa46d991b78246e8aaf19571ab.
//
// Solidity: event RoleUpdated(address indexed user, uint8 indexed role, bool enabled)
func (_Lawless *LawlessFilterer) ParseRoleUpdated(log types.Log) (*LawlessRoleUpdated, error) {
	event := new(LawlessRoleUpdated)
	if err := _Lawless.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LawlessTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Lawless contract.
type LawlessTransferIterator struct {
	Event *LawlessTransfer // Event containing the contract specifics and raw log

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
func (it *LawlessTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LawlessTransfer)
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
		it.Event = new(LawlessTransfer)
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
func (it *LawlessTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LawlessTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LawlessTransfer represents a Transfer event raised by the Lawless contract.
type LawlessTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Lawless *LawlessFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LawlessTransferIterator, error) {

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

	logs, sub, err := _Lawless.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LawlessTransferIterator{contract: _Lawless.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Lawless *LawlessFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LawlessTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Lawless.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LawlessTransfer)
				if err := _Lawless.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Lawless *LawlessFilterer) ParseTransfer(log types.Log) (*LawlessTransfer, error) {
	event := new(LawlessTransfer)
	if err := _Lawless.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LawlessUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Lawless contract.
type LawlessUnpausedIterator struct {
	Event *LawlessUnpaused // Event containing the contract specifics and raw log

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
func (it *LawlessUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LawlessUnpaused)
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
		it.Event = new(LawlessUnpaused)
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
func (it *LawlessUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LawlessUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LawlessUnpaused represents a Unpaused event raised by the Lawless contract.
type LawlessUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lawless *LawlessFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LawlessUnpausedIterator, error) {

	logs, sub, err := _Lawless.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LawlessUnpausedIterator{contract: _Lawless.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lawless *LawlessFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LawlessUnpaused) (event.Subscription, error) {

	logs, sub, err := _Lawless.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LawlessUnpaused)
				if err := _Lawless.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Lawless *LawlessFilterer) ParseUnpaused(log types.Log) (*LawlessUnpaused, error) {
	event := new(LawlessUnpaused)
	if err := _Lawless.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
