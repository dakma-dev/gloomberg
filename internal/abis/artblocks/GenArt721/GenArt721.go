// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenArt721

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
	_ = abi.ConvertType
)

// GenArt721MetaData contains all meta data concerning the GenArt721 contract.
var GenArt721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_invocations\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pricePerTokenInWei\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_dynamic\",\"type\":\"bool\"}],\"name\":\"addProject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_script\",\"type\":\"string\"}],\"name\":\"addProjectScript\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"artblocksAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"artblocksPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"clearTokenIpfsImageUri\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyaltyData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"artistAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"additionalPayee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalPayeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"royaltyFeeByID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextProjectId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ipfsHash\",\"type\":\"string\"}],\"name\":\"overrideTokenDynamicImageWithIpfsLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"artist\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"license\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"dynamic\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"projectScriptByIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectScriptInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"scriptJSON\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"scriptCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hashes\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectShowAllTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectTokenInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"artistAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerTokenInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invocations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxInvocations\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"additionalPayee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalPayeePercentage\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectURIInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"projectBaseURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"projectBaseIpfsURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useIpfs\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"purchase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"purchaseTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"removeProjectLastScript\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"showTokenHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"staticIpfsImageLink\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsDynamic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectUseIpfsForStatic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToProjectId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_artblocksAddress\",\"type\":\"address\"}],\"name\":\"updateArtblocksAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_artblocksPercentage\",\"type\":\"uint256\"}],\"name\":\"updateArtblocksPercentage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_additionalPayee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_additionalPayeePercentage\",\"type\":\"uint256\"}],\"name\":\"updateProjectAdditionalPayeeInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_artistAddress\",\"type\":\"address\"}],\"name\":\"updateProjectArtistAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectArtistName\",\"type\":\"string\"}],\"name\":\"updateProjectArtistName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectBaseIpfsURI\",\"type\":\"string\"}],\"name\":\"updateProjectBaseIpfsURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newBaseURI\",\"type\":\"string\"}],\"name\":\"updateProjectBaseURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectDescription\",\"type\":\"string\"}],\"name\":\"updateProjectDescription\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hashes\",\"type\":\"uint256\"}],\"name\":\"updateProjectHashesGenerated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ipfsHash\",\"type\":\"string\"}],\"name\":\"updateProjectIpfsHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectLicense\",\"type\":\"string\"}],\"name\":\"updateProjectLicense\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxInvocations\",\"type\":\"uint256\"}],\"name\":\"updateProjectMaxInvocations\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectName\",\"type\":\"string\"}],\"name\":\"updateProjectName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerTokenInWei\",\"type\":\"uint256\"}],\"name\":\"updateProjectPricePerTokenInWei\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_scriptId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_script\",\"type\":\"string\"}],\"name\":\"updateProjectScript\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectScriptJSON\",\"type\":\"string\"}],\"name\":\"updateProjectScriptJSON\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondMarketRoyalty\",\"type\":\"uint256\"}],\"name\":\"updateProjectSecondaryMarketRoyaltyPercentage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectWebsite\",\"type\":\"string\"}],\"name\":\"updateProjectWebsite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GenArt721ABI is the input ABI used to generate the binding from.
// Deprecated: Use GenArt721MetaData.ABI instead.
var GenArt721ABI = GenArt721MetaData.ABI

// GenArt721 is an auto generated Go binding around an Ethereum contract.
type GenArt721 struct {
	GenArt721Caller     // Read-only binding to the contract
	GenArt721Transactor // Write-only binding to the contract
	GenArt721Filterer   // Log filterer for contract events
}

// GenArt721Caller is an auto generated read-only Go binding around an Ethereum contract.
type GenArt721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenArt721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type GenArt721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenArt721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenArt721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenArt721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenArt721Session struct {
	Contract     *GenArt721        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenArt721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenArt721CallerSession struct {
	Contract *GenArt721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GenArt721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenArt721TransactorSession struct {
	Contract     *GenArt721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GenArt721Raw is an auto generated low-level Go binding around an Ethereum contract.
type GenArt721Raw struct {
	Contract *GenArt721 // Generic contract binding to access the raw methods on
}

// GenArt721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenArt721CallerRaw struct {
	Contract *GenArt721Caller // Generic read-only contract binding to access the raw methods on
}

// GenArt721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenArt721TransactorRaw struct {
	Contract *GenArt721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGenArt721 creates a new instance of GenArt721, bound to a specific deployed contract.
func NewGenArt721(address common.Address, backend bind.ContractBackend) (*GenArt721, error) {
	contract, err := bindGenArt721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenArt721{GenArt721Caller: GenArt721Caller{contract: contract}, GenArt721Transactor: GenArt721Transactor{contract: contract}, GenArt721Filterer: GenArt721Filterer{contract: contract}}, nil
}

// NewGenArt721Caller creates a new read-only instance of GenArt721, bound to a specific deployed contract.
func NewGenArt721Caller(address common.Address, caller bind.ContractCaller) (*GenArt721Caller, error) {
	contract, err := bindGenArt721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenArt721Caller{contract: contract}, nil
}

// NewGenArt721Transactor creates a new write-only instance of GenArt721, bound to a specific deployed contract.
func NewGenArt721Transactor(address common.Address, transactor bind.ContractTransactor) (*GenArt721Transactor, error) {
	contract, err := bindGenArt721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenArt721Transactor{contract: contract}, nil
}

// NewGenArt721Filterer creates a new log filterer instance of GenArt721, bound to a specific deployed contract.
func NewGenArt721Filterer(address common.Address, filterer bind.ContractFilterer) (*GenArt721Filterer, error) {
	contract, err := bindGenArt721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenArt721Filterer{contract: contract}, nil
}

// bindGenArt721 binds a generic wrapper to an already deployed contract.
func bindGenArt721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GenArt721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenArt721 *GenArt721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenArt721.Contract.GenArt721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenArt721 *GenArt721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenArt721.Contract.GenArt721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenArt721 *GenArt721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenArt721.Contract.GenArt721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenArt721 *GenArt721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenArt721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenArt721 *GenArt721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenArt721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenArt721 *GenArt721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenArt721.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GenArt721 *GenArt721Caller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GenArt721 *GenArt721Session) Admin() (common.Address, error) {
	return _GenArt721.Contract.Admin(&_GenArt721.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GenArt721 *GenArt721CallerSession) Admin() (common.Address, error) {
	return _GenArt721.Contract.Admin(&_GenArt721.CallOpts)
}

// ArtblocksAddress is a free data retrieval call binding the contract method 0x3949f906.
//
// Solidity: function artblocksAddress() view returns(address)
func (_GenArt721 *GenArt721Caller) ArtblocksAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "artblocksAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArtblocksAddress is a free data retrieval call binding the contract method 0x3949f906.
//
// Solidity: function artblocksAddress() view returns(address)
func (_GenArt721 *GenArt721Session) ArtblocksAddress() (common.Address, error) {
	return _GenArt721.Contract.ArtblocksAddress(&_GenArt721.CallOpts)
}

// ArtblocksAddress is a free data retrieval call binding the contract method 0x3949f906.
//
// Solidity: function artblocksAddress() view returns(address)
func (_GenArt721 *GenArt721CallerSession) ArtblocksAddress() (common.Address, error) {
	return _GenArt721.Contract.ArtblocksAddress(&_GenArt721.CallOpts)
}

// ArtblocksPercentage is a free data retrieval call binding the contract method 0x4f029c39.
//
// Solidity: function artblocksPercentage() view returns(uint256)
func (_GenArt721 *GenArt721Caller) ArtblocksPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "artblocksPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArtblocksPercentage is a free data retrieval call binding the contract method 0x4f029c39.
//
// Solidity: function artblocksPercentage() view returns(uint256)
func (_GenArt721 *GenArt721Session) ArtblocksPercentage() (*big.Int, error) {
	return _GenArt721.Contract.ArtblocksPercentage(&_GenArt721.CallOpts)
}

// ArtblocksPercentage is a free data retrieval call binding the contract method 0x4f029c39.
//
// Solidity: function artblocksPercentage() view returns(uint256)
func (_GenArt721 *GenArt721CallerSession) ArtblocksPercentage() (*big.Int, error) {
	return _GenArt721.Contract.ArtblocksPercentage(&_GenArt721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GenArt721 *GenArt721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GenArt721 *GenArt721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GenArt721.Contract.BalanceOf(&_GenArt721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GenArt721 *GenArt721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GenArt721.Contract.BalanceOf(&_GenArt721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GenArt721 *GenArt721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GenArt721 *GenArt721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GenArt721.Contract.GetApproved(&_GenArt721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GenArt721 *GenArt721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GenArt721.Contract.GetApproved(&_GenArt721.CallOpts, tokenId)
}

// GetRoyaltyData is a free data retrieval call binding the contract method 0xa65ff74c.
//
// Solidity: function getRoyaltyData(uint256 _tokenId) view returns(address artistAddress, address additionalPayee, uint256 additionalPayeePercentage, uint256 royaltyFeeByID)
func (_GenArt721 *GenArt721Caller) GetRoyaltyData(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	ArtistAddress             common.Address
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
	RoyaltyFeeByID            *big.Int
}, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "getRoyaltyData", _tokenId)

	outstruct := new(struct {
		ArtistAddress             common.Address
		AdditionalPayee           common.Address
		AdditionalPayeePercentage *big.Int
		RoyaltyFeeByID            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ArtistAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AdditionalPayee = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AdditionalPayeePercentage = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RoyaltyFeeByID = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoyaltyData is a free data retrieval call binding the contract method 0xa65ff74c.
//
// Solidity: function getRoyaltyData(uint256 _tokenId) view returns(address artistAddress, address additionalPayee, uint256 additionalPayeePercentage, uint256 royaltyFeeByID)
func (_GenArt721 *GenArt721Session) GetRoyaltyData(_tokenId *big.Int) (struct {
	ArtistAddress             common.Address
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
	RoyaltyFeeByID            *big.Int
}, error) {
	return _GenArt721.Contract.GetRoyaltyData(&_GenArt721.CallOpts, _tokenId)
}

// GetRoyaltyData is a free data retrieval call binding the contract method 0xa65ff74c.
//
// Solidity: function getRoyaltyData(uint256 _tokenId) view returns(address artistAddress, address additionalPayee, uint256 additionalPayeePercentage, uint256 royaltyFeeByID)
func (_GenArt721 *GenArt721CallerSession) GetRoyaltyData(_tokenId *big.Int) (struct {
	ArtistAddress             common.Address
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
	RoyaltyFeeByID            *big.Int
}, error) {
	return _GenArt721.Contract.GetRoyaltyData(&_GenArt721.CallOpts, _tokenId)
}

// HashToTokenId is a free data retrieval call binding the contract method 0xf51f74a9.
//
// Solidity: function hashToTokenId(bytes32 ) view returns(uint256)
func (_GenArt721 *GenArt721Caller) HashToTokenId(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "hashToTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashToTokenId is a free data retrieval call binding the contract method 0xf51f74a9.
//
// Solidity: function hashToTokenId(bytes32 ) view returns(uint256)
func (_GenArt721 *GenArt721Session) HashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _GenArt721.Contract.HashToTokenId(&_GenArt721.CallOpts, arg0)
}

// HashToTokenId is a free data retrieval call binding the contract method 0xf51f74a9.
//
// Solidity: function hashToTokenId(bytes32 ) view returns(uint256)
func (_GenArt721 *GenArt721CallerSession) HashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _GenArt721.Contract.HashToTokenId(&_GenArt721.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GenArt721 *GenArt721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GenArt721 *GenArt721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GenArt721.Contract.IsApprovedForAll(&_GenArt721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GenArt721 *GenArt721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GenArt721.Contract.IsApprovedForAll(&_GenArt721.CallOpts, owner, operator)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_GenArt721 *GenArt721Caller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_GenArt721 *GenArt721Session) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _GenArt721.Contract.IsWhitelisted(&_GenArt721.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_GenArt721 *GenArt721CallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _GenArt721.Contract.IsWhitelisted(&_GenArt721.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GenArt721 *GenArt721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GenArt721 *GenArt721Session) Name() (string, error) {
	return _GenArt721.Contract.Name(&_GenArt721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GenArt721 *GenArt721CallerSession) Name() (string, error) {
	return _GenArt721.Contract.Name(&_GenArt721.CallOpts)
}

// NextProjectId is a free data retrieval call binding the contract method 0xe935b7b1.
//
// Solidity: function nextProjectId() view returns(uint256)
func (_GenArt721 *GenArt721Caller) NextProjectId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "nextProjectId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextProjectId is a free data retrieval call binding the contract method 0xe935b7b1.
//
// Solidity: function nextProjectId() view returns(uint256)
func (_GenArt721 *GenArt721Session) NextProjectId() (*big.Int, error) {
	return _GenArt721.Contract.NextProjectId(&_GenArt721.CallOpts)
}

// NextProjectId is a free data retrieval call binding the contract method 0xe935b7b1.
//
// Solidity: function nextProjectId() view returns(uint256)
func (_GenArt721 *GenArt721CallerSession) NextProjectId() (*big.Int, error) {
	return _GenArt721.Contract.NextProjectId(&_GenArt721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GenArt721 *GenArt721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GenArt721 *GenArt721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GenArt721.Contract.OwnerOf(&_GenArt721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GenArt721 *GenArt721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GenArt721.Contract.OwnerOf(&_GenArt721.CallOpts, tokenId)
}

// ProjectDetails is a free data retrieval call binding the contract method 0x8dd91a56.
//
// Solidity: function projectDetails(uint256 _projectId) view returns(string projectName, string artist, string description, string website, string license, bool dynamic)
func (_GenArt721 *GenArt721Caller) ProjectDetails(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ProjectName string
	Artist      string
	Description string
	Website     string
	License     string
	Dynamic     bool
}, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "projectDetails", _projectId)

	outstruct := new(struct {
		ProjectName string
		Artist      string
		Description string
		Website     string
		License     string
		Dynamic     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProjectName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Artist = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Website = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.License = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Dynamic = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// ProjectDetails is a free data retrieval call binding the contract method 0x8dd91a56.
//
// Solidity: function projectDetails(uint256 _projectId) view returns(string projectName, string artist, string description, string website, string license, bool dynamic)
func (_GenArt721 *GenArt721Session) ProjectDetails(_projectId *big.Int) (struct {
	ProjectName string
	Artist      string
	Description string
	Website     string
	License     string
	Dynamic     bool
}, error) {
	return _GenArt721.Contract.ProjectDetails(&_GenArt721.CallOpts, _projectId)
}

// ProjectDetails is a free data retrieval call binding the contract method 0x8dd91a56.
//
// Solidity: function projectDetails(uint256 _projectId) view returns(string projectName, string artist, string description, string website, string license, bool dynamic)
func (_GenArt721 *GenArt721CallerSession) ProjectDetails(_projectId *big.Int) (struct {
	ProjectName string
	Artist      string
	Description string
	Website     string
	License     string
	Dynamic     bool
}, error) {
	return _GenArt721.Contract.ProjectDetails(&_GenArt721.CallOpts, _projectId)
}

// ProjectScriptByIndex is a free data retrieval call binding the contract method 0x8c3c9cdd.
//
// Solidity: function projectScriptByIndex(uint256 _projectId, uint256 _index) view returns(string)
func (_GenArt721 *GenArt721Caller) ProjectScriptByIndex(opts *bind.CallOpts, _projectId *big.Int, _index *big.Int) (string, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "projectScriptByIndex", _projectId, _index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProjectScriptByIndex is a free data retrieval call binding the contract method 0x8c3c9cdd.
//
// Solidity: function projectScriptByIndex(uint256 _projectId, uint256 _index) view returns(string)
func (_GenArt721 *GenArt721Session) ProjectScriptByIndex(_projectId *big.Int, _index *big.Int) (string, error) {
	return _GenArt721.Contract.ProjectScriptByIndex(&_GenArt721.CallOpts, _projectId, _index)
}

// ProjectScriptByIndex is a free data retrieval call binding the contract method 0x8c3c9cdd.
//
// Solidity: function projectScriptByIndex(uint256 _projectId, uint256 _index) view returns(string)
func (_GenArt721 *GenArt721CallerSession) ProjectScriptByIndex(_projectId *big.Int, _index *big.Int) (string, error) {
	return _GenArt721.Contract.ProjectScriptByIndex(&_GenArt721.CallOpts, _projectId, _index)
}

// ProjectScriptInfo is a free data retrieval call binding the contract method 0x4aa6d417.
//
// Solidity: function projectScriptInfo(uint256 _projectId) view returns(string scriptJSON, uint256 scriptCount, uint256 hashes, string ipfsHash, bool locked, bool paused)
func (_GenArt721 *GenArt721Caller) ProjectScriptInfo(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ScriptJSON  string
	ScriptCount *big.Int
	Hashes      *big.Int
	IpfsHash    string
	Locked      bool
	Paused      bool
}, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "projectScriptInfo", _projectId)

	outstruct := new(struct {
		ScriptJSON  string
		ScriptCount *big.Int
		Hashes      *big.Int
		IpfsHash    string
		Locked      bool
		Paused      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScriptJSON = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ScriptCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Hashes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IpfsHash = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Locked = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Paused = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// ProjectScriptInfo is a free data retrieval call binding the contract method 0x4aa6d417.
//
// Solidity: function projectScriptInfo(uint256 _projectId) view returns(string scriptJSON, uint256 scriptCount, uint256 hashes, string ipfsHash, bool locked, bool paused)
func (_GenArt721 *GenArt721Session) ProjectScriptInfo(_projectId *big.Int) (struct {
	ScriptJSON  string
	ScriptCount *big.Int
	Hashes      *big.Int
	IpfsHash    string
	Locked      bool
	Paused      bool
}, error) {
	return _GenArt721.Contract.ProjectScriptInfo(&_GenArt721.CallOpts, _projectId)
}

// ProjectScriptInfo is a free data retrieval call binding the contract method 0x4aa6d417.
//
// Solidity: function projectScriptInfo(uint256 _projectId) view returns(string scriptJSON, uint256 scriptCount, uint256 hashes, string ipfsHash, bool locked, bool paused)
func (_GenArt721 *GenArt721CallerSession) ProjectScriptInfo(_projectId *big.Int) (struct {
	ScriptJSON  string
	ScriptCount *big.Int
	Hashes      *big.Int
	IpfsHash    string
	Locked      bool
	Paused      bool
}, error) {
	return _GenArt721.Contract.ProjectScriptInfo(&_GenArt721.CallOpts, _projectId)
}

// ProjectShowAllTokens is a free data retrieval call binding the contract method 0xbee04f9c.
//
// Solidity: function projectShowAllTokens(uint256 _projectId) view returns(uint256[])
func (_GenArt721 *GenArt721Caller) ProjectShowAllTokens(opts *bind.CallOpts, _projectId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "projectShowAllTokens", _projectId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ProjectShowAllTokens is a free data retrieval call binding the contract method 0xbee04f9c.
//
// Solidity: function projectShowAllTokens(uint256 _projectId) view returns(uint256[])
func (_GenArt721 *GenArt721Session) ProjectShowAllTokens(_projectId *big.Int) ([]*big.Int, error) {
	return _GenArt721.Contract.ProjectShowAllTokens(&_GenArt721.CallOpts, _projectId)
}

// ProjectShowAllTokens is a free data retrieval call binding the contract method 0xbee04f9c.
//
// Solidity: function projectShowAllTokens(uint256 _projectId) view returns(uint256[])
func (_GenArt721 *GenArt721CallerSession) ProjectShowAllTokens(_projectId *big.Int) ([]*big.Int, error) {
	return _GenArt721.Contract.ProjectShowAllTokens(&_GenArt721.CallOpts, _projectId)
}

// ProjectTokenInfo is a free data retrieval call binding the contract method 0x8c2c3622.
//
// Solidity: function projectTokenInfo(uint256 _projectId) view returns(address artistAddress, uint256 pricePerTokenInWei, uint256 invocations, uint256 maxInvocations, bool active, address additionalPayee, uint256 additionalPayeePercentage)
func (_GenArt721 *GenArt721Caller) ProjectTokenInfo(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ArtistAddress             common.Address
	PricePerTokenInWei        *big.Int
	Invocations               *big.Int
	MaxInvocations            *big.Int
	Active                    bool
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
}, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "projectTokenInfo", _projectId)

	outstruct := new(struct {
		ArtistAddress             common.Address
		PricePerTokenInWei        *big.Int
		Invocations               *big.Int
		MaxInvocations            *big.Int
		Active                    bool
		AdditionalPayee           common.Address
		AdditionalPayeePercentage *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ArtistAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PricePerTokenInWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Invocations = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxInvocations = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.AdditionalPayee = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.AdditionalPayeePercentage = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProjectTokenInfo is a free data retrieval call binding the contract method 0x8c2c3622.
//
// Solidity: function projectTokenInfo(uint256 _projectId) view returns(address artistAddress, uint256 pricePerTokenInWei, uint256 invocations, uint256 maxInvocations, bool active, address additionalPayee, uint256 additionalPayeePercentage)
func (_GenArt721 *GenArt721Session) ProjectTokenInfo(_projectId *big.Int) (struct {
	ArtistAddress             common.Address
	PricePerTokenInWei        *big.Int
	Invocations               *big.Int
	MaxInvocations            *big.Int
	Active                    bool
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
}, error) {
	return _GenArt721.Contract.ProjectTokenInfo(&_GenArt721.CallOpts, _projectId)
}

// ProjectTokenInfo is a free data retrieval call binding the contract method 0x8c2c3622.
//
// Solidity: function projectTokenInfo(uint256 _projectId) view returns(address artistAddress, uint256 pricePerTokenInWei, uint256 invocations, uint256 maxInvocations, bool active, address additionalPayee, uint256 additionalPayeePercentage)
func (_GenArt721 *GenArt721CallerSession) ProjectTokenInfo(_projectId *big.Int) (struct {
	ArtistAddress             common.Address
	PricePerTokenInWei        *big.Int
	Invocations               *big.Int
	MaxInvocations            *big.Int
	Active                    bool
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
}, error) {
	return _GenArt721.Contract.ProjectTokenInfo(&_GenArt721.CallOpts, _projectId)
}

// ProjectURIInfo is a free data retrieval call binding the contract method 0x2d9c0205.
//
// Solidity: function projectURIInfo(uint256 _projectId) view returns(string projectBaseURI, string projectBaseIpfsURI, bool useIpfs)
func (_GenArt721 *GenArt721Caller) ProjectURIInfo(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ProjectBaseURI     string
	ProjectBaseIpfsURI string
	UseIpfs            bool
}, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "projectURIInfo", _projectId)

	outstruct := new(struct {
		ProjectBaseURI     string
		ProjectBaseIpfsURI string
		UseIpfs            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProjectBaseURI = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ProjectBaseIpfsURI = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.UseIpfs = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// ProjectURIInfo is a free data retrieval call binding the contract method 0x2d9c0205.
//
// Solidity: function projectURIInfo(uint256 _projectId) view returns(string projectBaseURI, string projectBaseIpfsURI, bool useIpfs)
func (_GenArt721 *GenArt721Session) ProjectURIInfo(_projectId *big.Int) (struct {
	ProjectBaseURI     string
	ProjectBaseIpfsURI string
	UseIpfs            bool
}, error) {
	return _GenArt721.Contract.ProjectURIInfo(&_GenArt721.CallOpts, _projectId)
}

// ProjectURIInfo is a free data retrieval call binding the contract method 0x2d9c0205.
//
// Solidity: function projectURIInfo(uint256 _projectId) view returns(string projectBaseURI, string projectBaseIpfsURI, bool useIpfs)
func (_GenArt721 *GenArt721CallerSession) ProjectURIInfo(_projectId *big.Int) (struct {
	ProjectBaseURI     string
	ProjectBaseIpfsURI string
	UseIpfs            bool
}, error) {
	return _GenArt721.Contract.ProjectURIInfo(&_GenArt721.CallOpts, _projectId)
}

// ShowTokenHashes is a free data retrieval call binding the contract method 0x271aaab4.
//
// Solidity: function showTokenHashes(uint256 _tokenId) view returns(bytes32[])
func (_GenArt721 *GenArt721Caller) ShowTokenHashes(opts *bind.CallOpts, _tokenId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "showTokenHashes", _tokenId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ShowTokenHashes is a free data retrieval call binding the contract method 0x271aaab4.
//
// Solidity: function showTokenHashes(uint256 _tokenId) view returns(bytes32[])
func (_GenArt721 *GenArt721Session) ShowTokenHashes(_tokenId *big.Int) ([][32]byte, error) {
	return _GenArt721.Contract.ShowTokenHashes(&_GenArt721.CallOpts, _tokenId)
}

// ShowTokenHashes is a free data retrieval call binding the contract method 0x271aaab4.
//
// Solidity: function showTokenHashes(uint256 _tokenId) view returns(bytes32[])
func (_GenArt721 *GenArt721CallerSession) ShowTokenHashes(_tokenId *big.Int) ([][32]byte, error) {
	return _GenArt721.Contract.ShowTokenHashes(&_GenArt721.CallOpts, _tokenId)
}

// StaticIpfsImageLink is a free data retrieval call binding the contract method 0x261eb4e5.
//
// Solidity: function staticIpfsImageLink(uint256 ) view returns(string)
func (_GenArt721 *GenArt721Caller) StaticIpfsImageLink(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "staticIpfsImageLink", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StaticIpfsImageLink is a free data retrieval call binding the contract method 0x261eb4e5.
//
// Solidity: function staticIpfsImageLink(uint256 ) view returns(string)
func (_GenArt721 *GenArt721Session) StaticIpfsImageLink(arg0 *big.Int) (string, error) {
	return _GenArt721.Contract.StaticIpfsImageLink(&_GenArt721.CallOpts, arg0)
}

// StaticIpfsImageLink is a free data retrieval call binding the contract method 0x261eb4e5.
//
// Solidity: function staticIpfsImageLink(uint256 ) view returns(string)
func (_GenArt721 *GenArt721CallerSession) StaticIpfsImageLink(arg0 *big.Int) (string, error) {
	return _GenArt721.Contract.StaticIpfsImageLink(&_GenArt721.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GenArt721 *GenArt721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GenArt721 *GenArt721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GenArt721.Contract.SupportsInterface(&_GenArt721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GenArt721 *GenArt721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GenArt721.Contract.SupportsInterface(&_GenArt721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GenArt721 *GenArt721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GenArt721 *GenArt721Session) Symbol() (string, error) {
	return _GenArt721.Contract.Symbol(&_GenArt721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GenArt721 *GenArt721CallerSession) Symbol() (string, error) {
	return _GenArt721.Contract.Symbol(&_GenArt721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_GenArt721 *GenArt721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_GenArt721 *GenArt721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _GenArt721.Contract.TokenByIndex(&_GenArt721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_GenArt721 *GenArt721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _GenArt721.Contract.TokenByIndex(&_GenArt721.CallOpts, index)
}

// TokenIdToProjectId is a free data retrieval call binding the contract method 0x1b689c0b.
//
// Solidity: function tokenIdToProjectId(uint256 ) view returns(uint256)
func (_GenArt721 *GenArt721Caller) TokenIdToProjectId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "tokenIdToProjectId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToProjectId is a free data retrieval call binding the contract method 0x1b689c0b.
//
// Solidity: function tokenIdToProjectId(uint256 ) view returns(uint256)
func (_GenArt721 *GenArt721Session) TokenIdToProjectId(arg0 *big.Int) (*big.Int, error) {
	return _GenArt721.Contract.TokenIdToProjectId(&_GenArt721.CallOpts, arg0)
}

// TokenIdToProjectId is a free data retrieval call binding the contract method 0x1b689c0b.
//
// Solidity: function tokenIdToProjectId(uint256 ) view returns(uint256)
func (_GenArt721 *GenArt721CallerSession) TokenIdToProjectId(arg0 *big.Int) (*big.Int, error) {
	return _GenArt721.Contract.TokenIdToProjectId(&_GenArt721.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_GenArt721 *GenArt721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_GenArt721 *GenArt721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _GenArt721.Contract.TokenOfOwnerByIndex(&_GenArt721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_GenArt721 *GenArt721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _GenArt721.Contract.TokenOfOwnerByIndex(&_GenArt721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_GenArt721 *GenArt721Caller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_GenArt721 *GenArt721Session) TokenURI(_tokenId *big.Int) (string, error) {
	return _GenArt721.Contract.TokenURI(&_GenArt721.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_GenArt721 *GenArt721CallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _GenArt721.Contract.TokenURI(&_GenArt721.CallOpts, _tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_GenArt721 *GenArt721Caller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_GenArt721 *GenArt721Session) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _GenArt721.Contract.TokensOfOwner(&_GenArt721.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_GenArt721 *GenArt721CallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _GenArt721.Contract.TokensOfOwner(&_GenArt721.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GenArt721 *GenArt721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GenArt721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GenArt721 *GenArt721Session) TotalSupply() (*big.Int, error) {
	return _GenArt721.Contract.TotalSupply(&_GenArt721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GenArt721 *GenArt721CallerSession) TotalSupply() (*big.Int, error) {
	return _GenArt721.Contract.TotalSupply(&_GenArt721.CallOpts)
}

// AddProject is a paid mutator transaction binding the contract method 0x4ac03ec2.
//
// Solidity: function addProject(uint256 _pricePerTokenInWei, bool _dynamic) returns()
func (_GenArt721 *GenArt721Transactor) AddProject(opts *bind.TransactOpts, _pricePerTokenInWei *big.Int, _dynamic bool) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "addProject", _pricePerTokenInWei, _dynamic)
}

// AddProject is a paid mutator transaction binding the contract method 0x4ac03ec2.
//
// Solidity: function addProject(uint256 _pricePerTokenInWei, bool _dynamic) returns()
func (_GenArt721 *GenArt721Session) AddProject(_pricePerTokenInWei *big.Int, _dynamic bool) (*types.Transaction, error) {
	return _GenArt721.Contract.AddProject(&_GenArt721.TransactOpts, _pricePerTokenInWei, _dynamic)
}

// AddProject is a paid mutator transaction binding the contract method 0x4ac03ec2.
//
// Solidity: function addProject(uint256 _pricePerTokenInWei, bool _dynamic) returns()
func (_GenArt721 *GenArt721TransactorSession) AddProject(_pricePerTokenInWei *big.Int, _dynamic bool) (*types.Transaction, error) {
	return _GenArt721.Contract.AddProject(&_GenArt721.TransactOpts, _pricePerTokenInWei, _dynamic)
}

// AddProjectScript is a paid mutator transaction binding the contract method 0xacad0124.
//
// Solidity: function addProjectScript(uint256 _projectId, string _script) returns()
func (_GenArt721 *GenArt721Transactor) AddProjectScript(opts *bind.TransactOpts, _projectId *big.Int, _script string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "addProjectScript", _projectId, _script)
}

// AddProjectScript is a paid mutator transaction binding the contract method 0xacad0124.
//
// Solidity: function addProjectScript(uint256 _projectId, string _script) returns()
func (_GenArt721 *GenArt721Session) AddProjectScript(_projectId *big.Int, _script string) (*types.Transaction, error) {
	return _GenArt721.Contract.AddProjectScript(&_GenArt721.TransactOpts, _projectId, _script)
}

// AddProjectScript is a paid mutator transaction binding the contract method 0xacad0124.
//
// Solidity: function addProjectScript(uint256 _projectId, string _script) returns()
func (_GenArt721 *GenArt721TransactorSession) AddProjectScript(_projectId *big.Int, _script string) (*types.Transaction, error) {
	return _GenArt721.Contract.AddProjectScript(&_GenArt721.TransactOpts, _projectId, _script)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_GenArt721 *GenArt721Transactor) AddWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "addWhitelisted", _address)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_GenArt721 *GenArt721Session) AddWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _GenArt721.Contract.AddWhitelisted(&_GenArt721.TransactOpts, _address)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_GenArt721 *GenArt721TransactorSession) AddWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _GenArt721.Contract.AddWhitelisted(&_GenArt721.TransactOpts, _address)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GenArt721 *GenArt721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GenArt721 *GenArt721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.Approve(&_GenArt721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GenArt721 *GenArt721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.Approve(&_GenArt721.TransactOpts, to, tokenId)
}

// ClearTokenIpfsImageUri is a paid mutator transaction binding the contract method 0x27901822.
//
// Solidity: function clearTokenIpfsImageUri(uint256 _tokenId) returns()
func (_GenArt721 *GenArt721Transactor) ClearTokenIpfsImageUri(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "clearTokenIpfsImageUri", _tokenId)
}

// ClearTokenIpfsImageUri is a paid mutator transaction binding the contract method 0x27901822.
//
// Solidity: function clearTokenIpfsImageUri(uint256 _tokenId) returns()
func (_GenArt721 *GenArt721Session) ClearTokenIpfsImageUri(_tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ClearTokenIpfsImageUri(&_GenArt721.TransactOpts, _tokenId)
}

// ClearTokenIpfsImageUri is a paid mutator transaction binding the contract method 0x27901822.
//
// Solidity: function clearTokenIpfsImageUri(uint256 _tokenId) returns()
func (_GenArt721 *GenArt721TransactorSession) ClearTokenIpfsImageUri(_tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ClearTokenIpfsImageUri(&_GenArt721.TransactOpts, _tokenId)
}

// OverrideTokenDynamicImageWithIpfsLink is a paid mutator transaction binding the contract method 0x93961c66.
//
// Solidity: function overrideTokenDynamicImageWithIpfsLink(uint256 _tokenId, string _ipfsHash) returns()
func (_GenArt721 *GenArt721Transactor) OverrideTokenDynamicImageWithIpfsLink(opts *bind.TransactOpts, _tokenId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "overrideTokenDynamicImageWithIpfsLink", _tokenId, _ipfsHash)
}

// OverrideTokenDynamicImageWithIpfsLink is a paid mutator transaction binding the contract method 0x93961c66.
//
// Solidity: function overrideTokenDynamicImageWithIpfsLink(uint256 _tokenId, string _ipfsHash) returns()
func (_GenArt721 *GenArt721Session) OverrideTokenDynamicImageWithIpfsLink(_tokenId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _GenArt721.Contract.OverrideTokenDynamicImageWithIpfsLink(&_GenArt721.TransactOpts, _tokenId, _ipfsHash)
}

// OverrideTokenDynamicImageWithIpfsLink is a paid mutator transaction binding the contract method 0x93961c66.
//
// Solidity: function overrideTokenDynamicImageWithIpfsLink(uint256 _tokenId, string _ipfsHash) returns()
func (_GenArt721 *GenArt721TransactorSession) OverrideTokenDynamicImageWithIpfsLink(_tokenId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _GenArt721.Contract.OverrideTokenDynamicImageWithIpfsLink(&_GenArt721.TransactOpts, _tokenId, _ipfsHash)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 _projectId) payable returns(uint256 _tokenId)
func (_GenArt721 *GenArt721Transactor) Purchase(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "purchase", _projectId)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 _projectId) payable returns(uint256 _tokenId)
func (_GenArt721 *GenArt721Session) Purchase(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.Purchase(&_GenArt721.TransactOpts, _projectId)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 _projectId) payable returns(uint256 _tokenId)
func (_GenArt721 *GenArt721TransactorSession) Purchase(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.Purchase(&_GenArt721.TransactOpts, _projectId)
}

// PurchaseTo is a paid mutator transaction binding the contract method 0x891407c0.
//
// Solidity: function purchaseTo(address _to, uint256 _projectId) payable returns(uint256 _tokenId)
func (_GenArt721 *GenArt721Transactor) PurchaseTo(opts *bind.TransactOpts, _to common.Address, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "purchaseTo", _to, _projectId)
}

// PurchaseTo is a paid mutator transaction binding the contract method 0x891407c0.
//
// Solidity: function purchaseTo(address _to, uint256 _projectId) payable returns(uint256 _tokenId)
func (_GenArt721 *GenArt721Session) PurchaseTo(_to common.Address, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.PurchaseTo(&_GenArt721.TransactOpts, _to, _projectId)
}

// PurchaseTo is a paid mutator transaction binding the contract method 0x891407c0.
//
// Solidity: function purchaseTo(address _to, uint256 _projectId) payable returns(uint256 _tokenId)
func (_GenArt721 *GenArt721TransactorSession) PurchaseTo(_to common.Address, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.PurchaseTo(&_GenArt721.TransactOpts, _to, _projectId)
}

// RemoveProjectLastScript is a paid mutator transaction binding the contract method 0xdb2ff861.
//
// Solidity: function removeProjectLastScript(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Transactor) RemoveProjectLastScript(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "removeProjectLastScript", _projectId)
}

// RemoveProjectLastScript is a paid mutator transaction binding the contract method 0xdb2ff861.
//
// Solidity: function removeProjectLastScript(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Session) RemoveProjectLastScript(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.RemoveProjectLastScript(&_GenArt721.TransactOpts, _projectId)
}

// RemoveProjectLastScript is a paid mutator transaction binding the contract method 0xdb2ff861.
//
// Solidity: function removeProjectLastScript(uint256 _projectId) returns()
func (_GenArt721 *GenArt721TransactorSession) RemoveProjectLastScript(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.RemoveProjectLastScript(&_GenArt721.TransactOpts, _projectId)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_GenArt721 *GenArt721Transactor) RemoveWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "removeWhitelisted", _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_GenArt721 *GenArt721Session) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _GenArt721.Contract.RemoveWhitelisted(&_GenArt721.TransactOpts, _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_GenArt721 *GenArt721TransactorSession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _GenArt721.Contract.RemoveWhitelisted(&_GenArt721.TransactOpts, _address)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GenArt721 *GenArt721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GenArt721 *GenArt721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.SafeTransferFrom(&_GenArt721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GenArt721 *GenArt721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.SafeTransferFrom(&_GenArt721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_GenArt721 *GenArt721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_GenArt721 *GenArt721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _GenArt721.Contract.SafeTransferFrom0(&_GenArt721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_GenArt721 *GenArt721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _GenArt721.Contract.SafeTransferFrom0(&_GenArt721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_GenArt721 *GenArt721Transactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_GenArt721 *GenArt721Session) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _GenArt721.Contract.SetApprovalForAll(&_GenArt721.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_GenArt721 *GenArt721TransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _GenArt721.Contract.SetApprovalForAll(&_GenArt721.TransactOpts, to, approved)
}

// ToggleProjectIsActive is a paid mutator transaction binding the contract method 0xd03c390c.
//
// Solidity: function toggleProjectIsActive(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Transactor) ToggleProjectIsActive(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "toggleProjectIsActive", _projectId)
}

// ToggleProjectIsActive is a paid mutator transaction binding the contract method 0xd03c390c.
//
// Solidity: function toggleProjectIsActive(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Session) ToggleProjectIsActive(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectIsActive(&_GenArt721.TransactOpts, _projectId)
}

// ToggleProjectIsActive is a paid mutator transaction binding the contract method 0xd03c390c.
//
// Solidity: function toggleProjectIsActive(uint256 _projectId) returns()
func (_GenArt721 *GenArt721TransactorSession) ToggleProjectIsActive(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectIsActive(&_GenArt721.TransactOpts, _projectId)
}

// ToggleProjectIsDynamic is a paid mutator transaction binding the contract method 0x3bdbd5c4.
//
// Solidity: function toggleProjectIsDynamic(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Transactor) ToggleProjectIsDynamic(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "toggleProjectIsDynamic", _projectId)
}

// ToggleProjectIsDynamic is a paid mutator transaction binding the contract method 0x3bdbd5c4.
//
// Solidity: function toggleProjectIsDynamic(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Session) ToggleProjectIsDynamic(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectIsDynamic(&_GenArt721.TransactOpts, _projectId)
}

// ToggleProjectIsDynamic is a paid mutator transaction binding the contract method 0x3bdbd5c4.
//
// Solidity: function toggleProjectIsDynamic(uint256 _projectId) returns()
func (_GenArt721 *GenArt721TransactorSession) ToggleProjectIsDynamic(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectIsDynamic(&_GenArt721.TransactOpts, _projectId)
}

// ToggleProjectIsLocked is a paid mutator transaction binding the contract method 0x8ba8f14d.
//
// Solidity: function toggleProjectIsLocked(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Transactor) ToggleProjectIsLocked(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "toggleProjectIsLocked", _projectId)
}

// ToggleProjectIsLocked is a paid mutator transaction binding the contract method 0x8ba8f14d.
//
// Solidity: function toggleProjectIsLocked(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Session) ToggleProjectIsLocked(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectIsLocked(&_GenArt721.TransactOpts, _projectId)
}

// ToggleProjectIsLocked is a paid mutator transaction binding the contract method 0x8ba8f14d.
//
// Solidity: function toggleProjectIsLocked(uint256 _projectId) returns()
func (_GenArt721 *GenArt721TransactorSession) ToggleProjectIsLocked(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectIsLocked(&_GenArt721.TransactOpts, _projectId)
}

// ToggleProjectIsPaused is a paid mutator transaction binding the contract method 0xa11ec70a.
//
// Solidity: function toggleProjectIsPaused(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Transactor) ToggleProjectIsPaused(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "toggleProjectIsPaused", _projectId)
}

// ToggleProjectIsPaused is a paid mutator transaction binding the contract method 0xa11ec70a.
//
// Solidity: function toggleProjectIsPaused(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Session) ToggleProjectIsPaused(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectIsPaused(&_GenArt721.TransactOpts, _projectId)
}

// ToggleProjectIsPaused is a paid mutator transaction binding the contract method 0xa11ec70a.
//
// Solidity: function toggleProjectIsPaused(uint256 _projectId) returns()
func (_GenArt721 *GenArt721TransactorSession) ToggleProjectIsPaused(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectIsPaused(&_GenArt721.TransactOpts, _projectId)
}

// ToggleProjectUseIpfsForStatic is a paid mutator transaction binding the contract method 0x5c088dcc.
//
// Solidity: function toggleProjectUseIpfsForStatic(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Transactor) ToggleProjectUseIpfsForStatic(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "toggleProjectUseIpfsForStatic", _projectId)
}

// ToggleProjectUseIpfsForStatic is a paid mutator transaction binding the contract method 0x5c088dcc.
//
// Solidity: function toggleProjectUseIpfsForStatic(uint256 _projectId) returns()
func (_GenArt721 *GenArt721Session) ToggleProjectUseIpfsForStatic(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectUseIpfsForStatic(&_GenArt721.TransactOpts, _projectId)
}

// ToggleProjectUseIpfsForStatic is a paid mutator transaction binding the contract method 0x5c088dcc.
//
// Solidity: function toggleProjectUseIpfsForStatic(uint256 _projectId) returns()
func (_GenArt721 *GenArt721TransactorSession) ToggleProjectUseIpfsForStatic(_projectId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.ToggleProjectUseIpfsForStatic(&_GenArt721.TransactOpts, _projectId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GenArt721 *GenArt721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GenArt721 *GenArt721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.TransferFrom(&_GenArt721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GenArt721 *GenArt721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.TransferFrom(&_GenArt721.TransactOpts, from, to, tokenId)
}

// UpdateArtblocksAddress is a paid mutator transaction binding the contract method 0x06e1db17.
//
// Solidity: function updateArtblocksAddress(address _artblocksAddress) returns()
func (_GenArt721 *GenArt721Transactor) UpdateArtblocksAddress(opts *bind.TransactOpts, _artblocksAddress common.Address) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateArtblocksAddress", _artblocksAddress)
}

// UpdateArtblocksAddress is a paid mutator transaction binding the contract method 0x06e1db17.
//
// Solidity: function updateArtblocksAddress(address _artblocksAddress) returns()
func (_GenArt721 *GenArt721Session) UpdateArtblocksAddress(_artblocksAddress common.Address) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateArtblocksAddress(&_GenArt721.TransactOpts, _artblocksAddress)
}

// UpdateArtblocksAddress is a paid mutator transaction binding the contract method 0x06e1db17.
//
// Solidity: function updateArtblocksAddress(address _artblocksAddress) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateArtblocksAddress(_artblocksAddress common.Address) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateArtblocksAddress(&_GenArt721.TransactOpts, _artblocksAddress)
}

// UpdateArtblocksPercentage is a paid mutator transaction binding the contract method 0xed6df982.
//
// Solidity: function updateArtblocksPercentage(uint256 _artblocksPercentage) returns()
func (_GenArt721 *GenArt721Transactor) UpdateArtblocksPercentage(opts *bind.TransactOpts, _artblocksPercentage *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateArtblocksPercentage", _artblocksPercentage)
}

// UpdateArtblocksPercentage is a paid mutator transaction binding the contract method 0xed6df982.
//
// Solidity: function updateArtblocksPercentage(uint256 _artblocksPercentage) returns()
func (_GenArt721 *GenArt721Session) UpdateArtblocksPercentage(_artblocksPercentage *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateArtblocksPercentage(&_GenArt721.TransactOpts, _artblocksPercentage)
}

// UpdateArtblocksPercentage is a paid mutator transaction binding the contract method 0xed6df982.
//
// Solidity: function updateArtblocksPercentage(uint256 _artblocksPercentage) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateArtblocksPercentage(_artblocksPercentage *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateArtblocksPercentage(&_GenArt721.TransactOpts, _artblocksPercentage)
}

// UpdateProjectAdditionalPayeeInfo is a paid mutator transaction binding the contract method 0xe13208b4.
//
// Solidity: function updateProjectAdditionalPayeeInfo(uint256 _projectId, address _additionalPayee, uint256 _additionalPayeePercentage) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectAdditionalPayeeInfo(opts *bind.TransactOpts, _projectId *big.Int, _additionalPayee common.Address, _additionalPayeePercentage *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectAdditionalPayeeInfo", _projectId, _additionalPayee, _additionalPayeePercentage)
}

// UpdateProjectAdditionalPayeeInfo is a paid mutator transaction binding the contract method 0xe13208b4.
//
// Solidity: function updateProjectAdditionalPayeeInfo(uint256 _projectId, address _additionalPayee, uint256 _additionalPayeePercentage) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectAdditionalPayeeInfo(_projectId *big.Int, _additionalPayee common.Address, _additionalPayeePercentage *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectAdditionalPayeeInfo(&_GenArt721.TransactOpts, _projectId, _additionalPayee, _additionalPayeePercentage)
}

// UpdateProjectAdditionalPayeeInfo is a paid mutator transaction binding the contract method 0xe13208b4.
//
// Solidity: function updateProjectAdditionalPayeeInfo(uint256 _projectId, address _additionalPayee, uint256 _additionalPayeePercentage) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectAdditionalPayeeInfo(_projectId *big.Int, _additionalPayee common.Address, _additionalPayeePercentage *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectAdditionalPayeeInfo(&_GenArt721.TransactOpts, _projectId, _additionalPayee, _additionalPayeePercentage)
}

// UpdateProjectArtistAddress is a paid mutator transaction binding the contract method 0x69d14faf.
//
// Solidity: function updateProjectArtistAddress(uint256 _projectId, address _artistAddress) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectArtistAddress(opts *bind.TransactOpts, _projectId *big.Int, _artistAddress common.Address) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectArtistAddress", _projectId, _artistAddress)
}

// UpdateProjectArtistAddress is a paid mutator transaction binding the contract method 0x69d14faf.
//
// Solidity: function updateProjectArtistAddress(uint256 _projectId, address _artistAddress) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectArtistAddress(_projectId *big.Int, _artistAddress common.Address) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectArtistAddress(&_GenArt721.TransactOpts, _projectId, _artistAddress)
}

// UpdateProjectArtistAddress is a paid mutator transaction binding the contract method 0x69d14faf.
//
// Solidity: function updateProjectArtistAddress(uint256 _projectId, address _artistAddress) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectArtistAddress(_projectId *big.Int, _artistAddress common.Address) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectArtistAddress(&_GenArt721.TransactOpts, _projectId, _artistAddress)
}

// UpdateProjectArtistName is a paid mutator transaction binding the contract method 0xb7b04fae.
//
// Solidity: function updateProjectArtistName(uint256 _projectId, string _projectArtistName) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectArtistName(opts *bind.TransactOpts, _projectId *big.Int, _projectArtistName string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectArtistName", _projectId, _projectArtistName)
}

// UpdateProjectArtistName is a paid mutator transaction binding the contract method 0xb7b04fae.
//
// Solidity: function updateProjectArtistName(uint256 _projectId, string _projectArtistName) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectArtistName(_projectId *big.Int, _projectArtistName string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectArtistName(&_GenArt721.TransactOpts, _projectId, _projectArtistName)
}

// UpdateProjectArtistName is a paid mutator transaction binding the contract method 0xb7b04fae.
//
// Solidity: function updateProjectArtistName(uint256 _projectId, string _projectArtistName) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectArtistName(_projectId *big.Int, _projectArtistName string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectArtistName(&_GenArt721.TransactOpts, _projectId, _projectArtistName)
}

// UpdateProjectBaseIpfsURI is a paid mutator transaction binding the contract method 0x6bd5d591.
//
// Solidity: function updateProjectBaseIpfsURI(uint256 _projectId, string _projectBaseIpfsURI) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectBaseIpfsURI(opts *bind.TransactOpts, _projectId *big.Int, _projectBaseIpfsURI string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectBaseIpfsURI", _projectId, _projectBaseIpfsURI)
}

// UpdateProjectBaseIpfsURI is a paid mutator transaction binding the contract method 0x6bd5d591.
//
// Solidity: function updateProjectBaseIpfsURI(uint256 _projectId, string _projectBaseIpfsURI) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectBaseIpfsURI(_projectId *big.Int, _projectBaseIpfsURI string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectBaseIpfsURI(&_GenArt721.TransactOpts, _projectId, _projectBaseIpfsURI)
}

// UpdateProjectBaseIpfsURI is a paid mutator transaction binding the contract method 0x6bd5d591.
//
// Solidity: function updateProjectBaseIpfsURI(uint256 _projectId, string _projectBaseIpfsURI) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectBaseIpfsURI(_projectId *big.Int, _projectBaseIpfsURI string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectBaseIpfsURI(&_GenArt721.TransactOpts, _projectId, _projectBaseIpfsURI)
}

// UpdateProjectBaseURI is a paid mutator transaction binding the contract method 0x3e48e848.
//
// Solidity: function updateProjectBaseURI(uint256 _projectId, string _newBaseURI) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectBaseURI(opts *bind.TransactOpts, _projectId *big.Int, _newBaseURI string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectBaseURI", _projectId, _newBaseURI)
}

// UpdateProjectBaseURI is a paid mutator transaction binding the contract method 0x3e48e848.
//
// Solidity: function updateProjectBaseURI(uint256 _projectId, string _newBaseURI) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectBaseURI(_projectId *big.Int, _newBaseURI string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectBaseURI(&_GenArt721.TransactOpts, _projectId, _newBaseURI)
}

// UpdateProjectBaseURI is a paid mutator transaction binding the contract method 0x3e48e848.
//
// Solidity: function updateProjectBaseURI(uint256 _projectId, string _newBaseURI) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectBaseURI(_projectId *big.Int, _newBaseURI string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectBaseURI(&_GenArt721.TransactOpts, _projectId, _newBaseURI)
}

// UpdateProjectDescription is a paid mutator transaction binding the contract method 0xa3b2cca6.
//
// Solidity: function updateProjectDescription(uint256 _projectId, string _projectDescription) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectDescription(opts *bind.TransactOpts, _projectId *big.Int, _projectDescription string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectDescription", _projectId, _projectDescription)
}

// UpdateProjectDescription is a paid mutator transaction binding the contract method 0xa3b2cca6.
//
// Solidity: function updateProjectDescription(uint256 _projectId, string _projectDescription) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectDescription(_projectId *big.Int, _projectDescription string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectDescription(&_GenArt721.TransactOpts, _projectId, _projectDescription)
}

// UpdateProjectDescription is a paid mutator transaction binding the contract method 0xa3b2cca6.
//
// Solidity: function updateProjectDescription(uint256 _projectId, string _projectDescription) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectDescription(_projectId *big.Int, _projectDescription string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectDescription(&_GenArt721.TransactOpts, _projectId, _projectDescription)
}

// UpdateProjectHashesGenerated is a paid mutator transaction binding the contract method 0x77444f62.
//
// Solidity: function updateProjectHashesGenerated(uint256 _projectId, uint256 _hashes) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectHashesGenerated(opts *bind.TransactOpts, _projectId *big.Int, _hashes *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectHashesGenerated", _projectId, _hashes)
}

// UpdateProjectHashesGenerated is a paid mutator transaction binding the contract method 0x77444f62.
//
// Solidity: function updateProjectHashesGenerated(uint256 _projectId, uint256 _hashes) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectHashesGenerated(_projectId *big.Int, _hashes *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectHashesGenerated(&_GenArt721.TransactOpts, _projectId, _hashes)
}

// UpdateProjectHashesGenerated is a paid mutator transaction binding the contract method 0x77444f62.
//
// Solidity: function updateProjectHashesGenerated(uint256 _projectId, uint256 _hashes) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectHashesGenerated(_projectId *big.Int, _hashes *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectHashesGenerated(&_GenArt721.TransactOpts, _projectId, _hashes)
}

// UpdateProjectIpfsHash is a paid mutator transaction binding the contract method 0x3fef6c2a.
//
// Solidity: function updateProjectIpfsHash(uint256 _projectId, string _ipfsHash) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectIpfsHash(opts *bind.TransactOpts, _projectId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectIpfsHash", _projectId, _ipfsHash)
}

// UpdateProjectIpfsHash is a paid mutator transaction binding the contract method 0x3fef6c2a.
//
// Solidity: function updateProjectIpfsHash(uint256 _projectId, string _ipfsHash) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectIpfsHash(_projectId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectIpfsHash(&_GenArt721.TransactOpts, _projectId, _ipfsHash)
}

// UpdateProjectIpfsHash is a paid mutator transaction binding the contract method 0x3fef6c2a.
//
// Solidity: function updateProjectIpfsHash(uint256 _projectId, string _ipfsHash) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectIpfsHash(_projectId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectIpfsHash(&_GenArt721.TransactOpts, _projectId, _ipfsHash)
}

// UpdateProjectLicense is a paid mutator transaction binding the contract method 0x25b75d68.
//
// Solidity: function updateProjectLicense(uint256 _projectId, string _projectLicense) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectLicense(opts *bind.TransactOpts, _projectId *big.Int, _projectLicense string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectLicense", _projectId, _projectLicense)
}

// UpdateProjectLicense is a paid mutator transaction binding the contract method 0x25b75d68.
//
// Solidity: function updateProjectLicense(uint256 _projectId, string _projectLicense) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectLicense(_projectId *big.Int, _projectLicense string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectLicense(&_GenArt721.TransactOpts, _projectId, _projectLicense)
}

// UpdateProjectLicense is a paid mutator transaction binding the contract method 0x25b75d68.
//
// Solidity: function updateProjectLicense(uint256 _projectId, string _projectLicense) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectLicense(_projectId *big.Int, _projectLicense string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectLicense(&_GenArt721.TransactOpts, _projectId, _projectLicense)
}

// UpdateProjectMaxInvocations is a paid mutator transaction binding the contract method 0x826fc391.
//
// Solidity: function updateProjectMaxInvocations(uint256 _projectId, uint256 _maxInvocations) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectMaxInvocations(opts *bind.TransactOpts, _projectId *big.Int, _maxInvocations *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectMaxInvocations", _projectId, _maxInvocations)
}

// UpdateProjectMaxInvocations is a paid mutator transaction binding the contract method 0x826fc391.
//
// Solidity: function updateProjectMaxInvocations(uint256 _projectId, uint256 _maxInvocations) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectMaxInvocations(_projectId *big.Int, _maxInvocations *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectMaxInvocations(&_GenArt721.TransactOpts, _projectId, _maxInvocations)
}

// UpdateProjectMaxInvocations is a paid mutator transaction binding the contract method 0x826fc391.
//
// Solidity: function updateProjectMaxInvocations(uint256 _projectId, uint256 _maxInvocations) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectMaxInvocations(_projectId *big.Int, _maxInvocations *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectMaxInvocations(&_GenArt721.TransactOpts, _projectId, _maxInvocations)
}

// UpdateProjectName is a paid mutator transaction binding the contract method 0x0d170673.
//
// Solidity: function updateProjectName(uint256 _projectId, string _projectName) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectName(opts *bind.TransactOpts, _projectId *big.Int, _projectName string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectName", _projectId, _projectName)
}

// UpdateProjectName is a paid mutator transaction binding the contract method 0x0d170673.
//
// Solidity: function updateProjectName(uint256 _projectId, string _projectName) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectName(_projectId *big.Int, _projectName string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectName(&_GenArt721.TransactOpts, _projectId, _projectName)
}

// UpdateProjectName is a paid mutator transaction binding the contract method 0x0d170673.
//
// Solidity: function updateProjectName(uint256 _projectId, string _projectName) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectName(_projectId *big.Int, _projectName string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectName(&_GenArt721.TransactOpts, _projectId, _projectName)
}

// UpdateProjectPricePerTokenInWei is a paid mutator transaction binding the contract method 0x97dc86cf.
//
// Solidity: function updateProjectPricePerTokenInWei(uint256 _projectId, uint256 _pricePerTokenInWei) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectPricePerTokenInWei(opts *bind.TransactOpts, _projectId *big.Int, _pricePerTokenInWei *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectPricePerTokenInWei", _projectId, _pricePerTokenInWei)
}

// UpdateProjectPricePerTokenInWei is a paid mutator transaction binding the contract method 0x97dc86cf.
//
// Solidity: function updateProjectPricePerTokenInWei(uint256 _projectId, uint256 _pricePerTokenInWei) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectPricePerTokenInWei(_projectId *big.Int, _pricePerTokenInWei *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectPricePerTokenInWei(&_GenArt721.TransactOpts, _projectId, _pricePerTokenInWei)
}

// UpdateProjectPricePerTokenInWei is a paid mutator transaction binding the contract method 0x97dc86cf.
//
// Solidity: function updateProjectPricePerTokenInWei(uint256 _projectId, uint256 _pricePerTokenInWei) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectPricePerTokenInWei(_projectId *big.Int, _pricePerTokenInWei *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectPricePerTokenInWei(&_GenArt721.TransactOpts, _projectId, _pricePerTokenInWei)
}

// UpdateProjectScript is a paid mutator transaction binding the contract method 0xb1656ba3.
//
// Solidity: function updateProjectScript(uint256 _projectId, uint256 _scriptId, string _script) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectScript(opts *bind.TransactOpts, _projectId *big.Int, _scriptId *big.Int, _script string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectScript", _projectId, _scriptId, _script)
}

// UpdateProjectScript is a paid mutator transaction binding the contract method 0xb1656ba3.
//
// Solidity: function updateProjectScript(uint256 _projectId, uint256 _scriptId, string _script) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectScript(_projectId *big.Int, _scriptId *big.Int, _script string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectScript(&_GenArt721.TransactOpts, _projectId, _scriptId, _script)
}

// UpdateProjectScript is a paid mutator transaction binding the contract method 0xb1656ba3.
//
// Solidity: function updateProjectScript(uint256 _projectId, uint256 _scriptId, string _script) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectScript(_projectId *big.Int, _scriptId *big.Int, _script string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectScript(&_GenArt721.TransactOpts, _projectId, _scriptId, _script)
}

// UpdateProjectScriptJSON is a paid mutator transaction binding the contract method 0xc6d73231.
//
// Solidity: function updateProjectScriptJSON(uint256 _projectId, string _projectScriptJSON) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectScriptJSON(opts *bind.TransactOpts, _projectId *big.Int, _projectScriptJSON string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectScriptJSON", _projectId, _projectScriptJSON)
}

// UpdateProjectScriptJSON is a paid mutator transaction binding the contract method 0xc6d73231.
//
// Solidity: function updateProjectScriptJSON(uint256 _projectId, string _projectScriptJSON) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectScriptJSON(_projectId *big.Int, _projectScriptJSON string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectScriptJSON(&_GenArt721.TransactOpts, _projectId, _projectScriptJSON)
}

// UpdateProjectScriptJSON is a paid mutator transaction binding the contract method 0xc6d73231.
//
// Solidity: function updateProjectScriptJSON(uint256 _projectId, string _projectScriptJSON) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectScriptJSON(_projectId *big.Int, _projectScriptJSON string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectScriptJSON(&_GenArt721.TransactOpts, _projectId, _projectScriptJSON)
}

// UpdateProjectSecondaryMarketRoyaltyPercentage is a paid mutator transaction binding the contract method 0xc34a03b5.
//
// Solidity: function updateProjectSecondaryMarketRoyaltyPercentage(uint256 _projectId, uint256 _secondMarketRoyalty) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectSecondaryMarketRoyaltyPercentage(opts *bind.TransactOpts, _projectId *big.Int, _secondMarketRoyalty *big.Int) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectSecondaryMarketRoyaltyPercentage", _projectId, _secondMarketRoyalty)
}

// UpdateProjectSecondaryMarketRoyaltyPercentage is a paid mutator transaction binding the contract method 0xc34a03b5.
//
// Solidity: function updateProjectSecondaryMarketRoyaltyPercentage(uint256 _projectId, uint256 _secondMarketRoyalty) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectSecondaryMarketRoyaltyPercentage(_projectId *big.Int, _secondMarketRoyalty *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectSecondaryMarketRoyaltyPercentage(&_GenArt721.TransactOpts, _projectId, _secondMarketRoyalty)
}

// UpdateProjectSecondaryMarketRoyaltyPercentage is a paid mutator transaction binding the contract method 0xc34a03b5.
//
// Solidity: function updateProjectSecondaryMarketRoyaltyPercentage(uint256 _projectId, uint256 _secondMarketRoyalty) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectSecondaryMarketRoyaltyPercentage(_projectId *big.Int, _secondMarketRoyalty *big.Int) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectSecondaryMarketRoyaltyPercentage(&_GenArt721.TransactOpts, _projectId, _secondMarketRoyalty)
}

// UpdateProjectWebsite is a paid mutator transaction binding the contract method 0x37859963.
//
// Solidity: function updateProjectWebsite(uint256 _projectId, string _projectWebsite) returns()
func (_GenArt721 *GenArt721Transactor) UpdateProjectWebsite(opts *bind.TransactOpts, _projectId *big.Int, _projectWebsite string) (*types.Transaction, error) {
	return _GenArt721.contract.Transact(opts, "updateProjectWebsite", _projectId, _projectWebsite)
}

// UpdateProjectWebsite is a paid mutator transaction binding the contract method 0x37859963.
//
// Solidity: function updateProjectWebsite(uint256 _projectId, string _projectWebsite) returns()
func (_GenArt721 *GenArt721Session) UpdateProjectWebsite(_projectId *big.Int, _projectWebsite string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectWebsite(&_GenArt721.TransactOpts, _projectId, _projectWebsite)
}

// UpdateProjectWebsite is a paid mutator transaction binding the contract method 0x37859963.
//
// Solidity: function updateProjectWebsite(uint256 _projectId, string _projectWebsite) returns()
func (_GenArt721 *GenArt721TransactorSession) UpdateProjectWebsite(_projectId *big.Int, _projectWebsite string) (*types.Transaction, error) {
	return _GenArt721.Contract.UpdateProjectWebsite(&_GenArt721.TransactOpts, _projectId, _projectWebsite)
}

// GenArt721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GenArt721 contract.
type GenArt721ApprovalIterator struct {
	Event *GenArt721Approval // Event containing the contract specifics and raw log

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
func (it *GenArt721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenArt721Approval)
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
		it.Event = new(GenArt721Approval)
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
func (it *GenArt721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenArt721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenArt721Approval represents a Approval event raised by the GenArt721 contract.
type GenArt721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GenArt721 *GenArt721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*GenArt721ApprovalIterator, error) {

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

	logs, sub, err := _GenArt721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GenArt721ApprovalIterator{contract: _GenArt721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GenArt721 *GenArt721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GenArt721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _GenArt721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenArt721Approval)
				if err := _GenArt721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_GenArt721 *GenArt721Filterer) ParseApproval(log types.Log) (*GenArt721Approval, error) {
	event := new(GenArt721Approval)
	if err := _GenArt721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenArt721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the GenArt721 contract.
type GenArt721ApprovalForAllIterator struct {
	Event *GenArt721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *GenArt721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenArt721ApprovalForAll)
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
		it.Event = new(GenArt721ApprovalForAll)
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
func (it *GenArt721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenArt721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenArt721ApprovalForAll represents a ApprovalForAll event raised by the GenArt721 contract.
type GenArt721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GenArt721 *GenArt721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*GenArt721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GenArt721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GenArt721ApprovalForAllIterator{contract: _GenArt721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GenArt721 *GenArt721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GenArt721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GenArt721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenArt721ApprovalForAll)
				if err := _GenArt721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_GenArt721 *GenArt721Filterer) ParseApprovalForAll(log types.Log) (*GenArt721ApprovalForAll, error) {
	event := new(GenArt721ApprovalForAll)
	if err := _GenArt721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenArt721MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the GenArt721 contract.
type GenArt721MintIterator struct {
	Event *GenArt721Mint // Event containing the contract specifics and raw log

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
func (it *GenArt721MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenArt721Mint)
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
		it.Event = new(GenArt721Mint)
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
func (it *GenArt721MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenArt721MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenArt721Mint represents a Mint event raised by the GenArt721 contract.
type GenArt721Mint struct {
	To          common.Address
	TokenId     *big.Int
	ProjectId   *big.Int
	Invocations *big.Int
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x94c792774c59479f7bd68442f3af3691c02123a5aabee8b6f9116d8af8aa6669.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId, uint256 indexed _projectId, uint256 _invocations, uint256 _value)
func (_GenArt721 *GenArt721Filterer) FilterMint(opts *bind.FilterOpts, _to []common.Address, _tokenId []*big.Int, _projectId []*big.Int) (*GenArt721MintIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _projectIdRule []interface{}
	for _, _projectIdItem := range _projectId {
		_projectIdRule = append(_projectIdRule, _projectIdItem)
	}

	logs, sub, err := _GenArt721.contract.FilterLogs(opts, "Mint", _toRule, _tokenIdRule, _projectIdRule)
	if err != nil {
		return nil, err
	}
	return &GenArt721MintIterator{contract: _GenArt721.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x94c792774c59479f7bd68442f3af3691c02123a5aabee8b6f9116d8af8aa6669.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId, uint256 indexed _projectId, uint256 _invocations, uint256 _value)
func (_GenArt721 *GenArt721Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *GenArt721Mint, _to []common.Address, _tokenId []*big.Int, _projectId []*big.Int) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _projectIdRule []interface{}
	for _, _projectIdItem := range _projectId {
		_projectIdRule = append(_projectIdRule, _projectIdItem)
	}

	logs, sub, err := _GenArt721.contract.WatchLogs(opts, "Mint", _toRule, _tokenIdRule, _projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenArt721Mint)
				if err := _GenArt721.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x94c792774c59479f7bd68442f3af3691c02123a5aabee8b6f9116d8af8aa6669.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId, uint256 indexed _projectId, uint256 _invocations, uint256 _value)
func (_GenArt721 *GenArt721Filterer) ParseMint(log types.Log) (*GenArt721Mint, error) {
	event := new(GenArt721Mint)
	if err := _GenArt721.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenArt721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GenArt721 contract.
type GenArt721TransferIterator struct {
	Event *GenArt721Transfer // Event containing the contract specifics and raw log

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
func (it *GenArt721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenArt721Transfer)
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
		it.Event = new(GenArt721Transfer)
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
func (it *GenArt721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenArt721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenArt721Transfer represents a Transfer event raised by the GenArt721 contract.
type GenArt721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GenArt721 *GenArt721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*GenArt721TransferIterator, error) {

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

	logs, sub, err := _GenArt721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GenArt721TransferIterator{contract: _GenArt721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GenArt721 *GenArt721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GenArt721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _GenArt721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenArt721Transfer)
				if err := _GenArt721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_GenArt721 *GenArt721Filterer) ParseTransfer(log types.Log) (*GenArt721Transfer, error) {
	event := new(GenArt721Transfer)
	if err := _GenArt721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
