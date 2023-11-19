// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenArt721Core

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

// ArtblocksMetaData contains all meta data concerning the Artblocks contract.
var ArtblocksMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_randomizerContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addMintWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_projectName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_artistAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerTokenInWei\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_dynamic\",\"type\":\"bool\"}],\"name\":\"addProject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_script\",\"type\":\"string\"}],\"name\":\"addProjectScript\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"artblocksAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"artblocksPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"clearTokenIpfsImageUri\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyaltyData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"artistAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"additionalPayee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalPayeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"royaltyFeeByID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMintWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_by\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextProjectId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ipfsHash\",\"type\":\"string\"}],\"name\":\"overrideTokenDynamicImageWithIpfsLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"artist\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"license\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"dynamic\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectIdToAdditionalPayee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectIdToAdditionalPayeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectIdToArtistAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectIdToCurrencyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectIdToCurrencySymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectIdToPricePerTokenInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectIdToSecondaryMarketRoyaltyPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"projectScriptByIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectScriptInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"scriptJSON\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"scriptCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useHashString\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectShowAllTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectTokenInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"artistAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerTokenInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invocations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxInvocations\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"additionalPayee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalPayeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"currencyAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectURIInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"projectBaseURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"projectBaseIpfsURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useIpfs\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"randomizerContract\",\"outputs\":[{\"internalType\":\"contractRandomizer\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeMintWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"removeProjectLastScript\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"staticIpfsImageLink\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsDynamic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectUseHashString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectUseIpfsForStatic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToProjectId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_artblocksAddress\",\"type\":\"address\"}],\"name\":\"updateArtblocksAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_artblocksPercentage\",\"type\":\"uint256\"}],\"name\":\"updateArtblocksPercentage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_additionalPayee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_additionalPayeePercentage\",\"type\":\"uint256\"}],\"name\":\"updateProjectAdditionalPayeeInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_artistAddress\",\"type\":\"address\"}],\"name\":\"updateProjectArtistAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectArtistName\",\"type\":\"string\"}],\"name\":\"updateProjectArtistName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectBaseIpfsURI\",\"type\":\"string\"}],\"name\":\"updateProjectBaseIpfsURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newBaseURI\",\"type\":\"string\"}],\"name\":\"updateProjectBaseURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_currencySymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"}],\"name\":\"updateProjectCurrencyInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectDescription\",\"type\":\"string\"}],\"name\":\"updateProjectDescription\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ipfsHash\",\"type\":\"string\"}],\"name\":\"updateProjectIpfsHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectLicense\",\"type\":\"string\"}],\"name\":\"updateProjectLicense\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxInvocations\",\"type\":\"uint256\"}],\"name\":\"updateProjectMaxInvocations\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectName\",\"type\":\"string\"}],\"name\":\"updateProjectName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerTokenInWei\",\"type\":\"uint256\"}],\"name\":\"updateProjectPricePerTokenInWei\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_scriptId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_script\",\"type\":\"string\"}],\"name\":\"updateProjectScript\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectScriptJSON\",\"type\":\"string\"}],\"name\":\"updateProjectScriptJSON\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondMarketRoyalty\",\"type\":\"uint256\"}],\"name\":\"updateProjectSecondaryMarketRoyaltyPercentage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectWebsite\",\"type\":\"string\"}],\"name\":\"updateProjectWebsite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_randomizerAddress\",\"type\":\"address\"}],\"name\":\"updateRandomizerAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ArtblocksABI is the input ABI used to generate the binding from.
// Deprecated: Use ArtblocksMetaData.ABI instead.
var ArtblocksABI = ArtblocksMetaData.ABI

// Artblocks is an auto generated Go binding around an Ethereum contract.
type Artblocks struct {
	ArtblocksCaller     // Read-only binding to the contract
	ArtblocksTransactor // Write-only binding to the contract
	ArtblocksFilterer   // Log filterer for contract events
}

// ArtblocksCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArtblocksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtblocksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArtblocksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtblocksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArtblocksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtblocksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArtblocksSession struct {
	Contract     *Artblocks        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtblocksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArtblocksCallerSession struct {
	Contract *ArtblocksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArtblocksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArtblocksTransactorSession struct {
	Contract     *ArtblocksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArtblocksRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArtblocksRaw struct {
	Contract *Artblocks // Generic contract binding to access the raw methods on
}

// ArtblocksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArtblocksCallerRaw struct {
	Contract *ArtblocksCaller // Generic read-only contract binding to access the raw methods on
}

// ArtblocksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArtblocksTransactorRaw struct {
	Contract *ArtblocksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArtblocks creates a new instance of Artblocks, bound to a specific deployed contract.
func NewArtblocks(address common.Address, backend bind.ContractBackend) (*Artblocks, error) {
	contract, err := bindArtblocks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Artblocks{ArtblocksCaller: ArtblocksCaller{contract: contract}, ArtblocksTransactor: ArtblocksTransactor{contract: contract}, ArtblocksFilterer: ArtblocksFilterer{contract: contract}}, nil
}

// NewArtblocksCaller creates a new read-only instance of Artblocks, bound to a specific deployed contract.
func NewArtblocksCaller(address common.Address, caller bind.ContractCaller) (*ArtblocksCaller, error) {
	contract, err := bindArtblocks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArtblocksCaller{contract: contract}, nil
}

// NewArtblocksTransactor creates a new write-only instance of Artblocks, bound to a specific deployed contract.
func NewArtblocksTransactor(address common.Address, transactor bind.ContractTransactor) (*ArtblocksTransactor, error) {
	contract, err := bindArtblocks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArtblocksTransactor{contract: contract}, nil
}

// NewArtblocksFilterer creates a new log filterer instance of Artblocks, bound to a specific deployed contract.
func NewArtblocksFilterer(address common.Address, filterer bind.ContractFilterer) (*ArtblocksFilterer, error) {
	contract, err := bindArtblocks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArtblocksFilterer{contract: contract}, nil
}

// bindArtblocks binds a generic wrapper to an already deployed contract.
func bindArtblocks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArtblocksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Artblocks *ArtblocksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Artblocks.Contract.ArtblocksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Artblocks *ArtblocksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artblocks.Contract.ArtblocksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Artblocks *ArtblocksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Artblocks.Contract.ArtblocksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Artblocks *ArtblocksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Artblocks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Artblocks *ArtblocksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artblocks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Artblocks *ArtblocksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Artblocks.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Artblocks *ArtblocksCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Artblocks *ArtblocksSession) Admin() (common.Address, error) {
	return _Artblocks.Contract.Admin(&_Artblocks.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Artblocks *ArtblocksCallerSession) Admin() (common.Address, error) {
	return _Artblocks.Contract.Admin(&_Artblocks.CallOpts)
}

// ArtblocksAddress is a free data retrieval call binding the contract method 0x3949f906.
//
// Solidity: function artblocksAddress() view returns(address)
func (_Artblocks *ArtblocksCaller) ArtblocksAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "artblocksAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArtblocksAddress is a free data retrieval call binding the contract method 0x3949f906.
//
// Solidity: function artblocksAddress() view returns(address)
func (_Artblocks *ArtblocksSession) ArtblocksAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksAddress(&_Artblocks.CallOpts)
}

// ArtblocksAddress is a free data retrieval call binding the contract method 0x3949f906.
//
// Solidity: function artblocksAddress() view returns(address)
func (_Artblocks *ArtblocksCallerSession) ArtblocksAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksAddress(&_Artblocks.CallOpts)
}

// ArtblocksPercentage is a free data retrieval call binding the contract method 0x4f029c39.
//
// Solidity: function artblocksPercentage() view returns(uint256)
func (_Artblocks *ArtblocksCaller) ArtblocksPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "artblocksPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArtblocksPercentage is a free data retrieval call binding the contract method 0x4f029c39.
//
// Solidity: function artblocksPercentage() view returns(uint256)
func (_Artblocks *ArtblocksSession) ArtblocksPercentage() (*big.Int, error) {
	return _Artblocks.Contract.ArtblocksPercentage(&_Artblocks.CallOpts)
}

// ArtblocksPercentage is a free data retrieval call binding the contract method 0x4f029c39.
//
// Solidity: function artblocksPercentage() view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) ArtblocksPercentage() (*big.Int, error) {
	return _Artblocks.Contract.ArtblocksPercentage(&_Artblocks.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Artblocks *ArtblocksCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Artblocks *ArtblocksSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Artblocks.Contract.BalanceOf(&_Artblocks.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Artblocks.Contract.BalanceOf(&_Artblocks.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Artblocks *ArtblocksCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Artblocks *ArtblocksSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.GetApproved(&_Artblocks.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Artblocks *ArtblocksCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.GetApproved(&_Artblocks.CallOpts, tokenId)
}

// GetRoyaltyData is a free data retrieval call binding the contract method 0xa65ff74c.
//
// Solidity: function getRoyaltyData(uint256 _tokenId) view returns(address artistAddress, address additionalPayee, uint256 additionalPayeePercentage, uint256 royaltyFeeByID)
func (_Artblocks *ArtblocksCaller) GetRoyaltyData(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	ArtistAddress             common.Address
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
	RoyaltyFeeByID            *big.Int
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "getRoyaltyData", _tokenId)

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
func (_Artblocks *ArtblocksSession) GetRoyaltyData(_tokenId *big.Int) (struct {
	ArtistAddress             common.Address
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
	RoyaltyFeeByID            *big.Int
}, error) {
	return _Artblocks.Contract.GetRoyaltyData(&_Artblocks.CallOpts, _tokenId)
}

// GetRoyaltyData is a free data retrieval call binding the contract method 0xa65ff74c.
//
// Solidity: function getRoyaltyData(uint256 _tokenId) view returns(address artistAddress, address additionalPayee, uint256 additionalPayeePercentage, uint256 royaltyFeeByID)
func (_Artblocks *ArtblocksCallerSession) GetRoyaltyData(_tokenId *big.Int) (struct {
	ArtistAddress             common.Address
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
	RoyaltyFeeByID            *big.Int
}, error) {
	return _Artblocks.Contract.GetRoyaltyData(&_Artblocks.CallOpts, _tokenId)
}

// HashToTokenId is a free data retrieval call binding the contract method 0xf51f74a9.
//
// Solidity: function hashToTokenId(bytes32 ) view returns(uint256)
func (_Artblocks *ArtblocksCaller) HashToTokenId(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "hashToTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashToTokenId is a free data retrieval call binding the contract method 0xf51f74a9.
//
// Solidity: function hashToTokenId(bytes32 ) view returns(uint256)
func (_Artblocks *ArtblocksSession) HashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _Artblocks.Contract.HashToTokenId(&_Artblocks.CallOpts, arg0)
}

// HashToTokenId is a free data retrieval call binding the contract method 0xf51f74a9.
//
// Solidity: function hashToTokenId(bytes32 ) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) HashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _Artblocks.Contract.HashToTokenId(&_Artblocks.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Artblocks *ArtblocksCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Artblocks *ArtblocksSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Artblocks.Contract.IsApprovedForAll(&_Artblocks.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Artblocks *ArtblocksCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Artblocks.Contract.IsApprovedForAll(&_Artblocks.CallOpts, owner, operator)
}

// IsMintWhitelisted is a free data retrieval call binding the contract method 0xad0305ce.
//
// Solidity: function isMintWhitelisted(address ) view returns(bool)
func (_Artblocks *ArtblocksCaller) IsMintWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "isMintWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMintWhitelisted is a free data retrieval call binding the contract method 0xad0305ce.
//
// Solidity: function isMintWhitelisted(address ) view returns(bool)
func (_Artblocks *ArtblocksSession) IsMintWhitelisted(arg0 common.Address) (bool, error) {
	return _Artblocks.Contract.IsMintWhitelisted(&_Artblocks.CallOpts, arg0)
}

// IsMintWhitelisted is a free data retrieval call binding the contract method 0xad0305ce.
//
// Solidity: function isMintWhitelisted(address ) view returns(bool)
func (_Artblocks *ArtblocksCallerSession) IsMintWhitelisted(arg0 common.Address) (bool, error) {
	return _Artblocks.Contract.IsMintWhitelisted(&_Artblocks.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_Artblocks *ArtblocksCaller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_Artblocks *ArtblocksSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Artblocks.Contract.IsWhitelisted(&_Artblocks.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_Artblocks *ArtblocksCallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Artblocks.Contract.IsWhitelisted(&_Artblocks.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Artblocks *ArtblocksCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Artblocks *ArtblocksSession) Name() (string, error) {
	return _Artblocks.Contract.Name(&_Artblocks.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Artblocks *ArtblocksCallerSession) Name() (string, error) {
	return _Artblocks.Contract.Name(&_Artblocks.CallOpts)
}

// NextProjectId is a free data retrieval call binding the contract method 0xe935b7b1.
//
// Solidity: function nextProjectId() view returns(uint256)
func (_Artblocks *ArtblocksCaller) NextProjectId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "nextProjectId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextProjectId is a free data retrieval call binding the contract method 0xe935b7b1.
//
// Solidity: function nextProjectId() view returns(uint256)
func (_Artblocks *ArtblocksSession) NextProjectId() (*big.Int, error) {
	return _Artblocks.Contract.NextProjectId(&_Artblocks.CallOpts)
}

// NextProjectId is a free data retrieval call binding the contract method 0xe935b7b1.
//
// Solidity: function nextProjectId() view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) NextProjectId() (*big.Int, error) {
	return _Artblocks.Contract.NextProjectId(&_Artblocks.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Artblocks *ArtblocksCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Artblocks *ArtblocksSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.OwnerOf(&_Artblocks.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Artblocks *ArtblocksCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.OwnerOf(&_Artblocks.CallOpts, tokenId)
}

// ProjectDetails is a free data retrieval call binding the contract method 0x8dd91a56.
//
// Solidity: function projectDetails(uint256 _projectId) view returns(string projectName, string artist, string description, string website, string license, bool dynamic)
func (_Artblocks *ArtblocksCaller) ProjectDetails(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ProjectName string
	Artist      string
	Description string
	Website     string
	License     string
	Dynamic     bool
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectDetails", _projectId)

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
func (_Artblocks *ArtblocksSession) ProjectDetails(_projectId *big.Int) (struct {
	ProjectName string
	Artist      string
	Description string
	Website     string
	License     string
	Dynamic     bool
}, error) {
	return _Artblocks.Contract.ProjectDetails(&_Artblocks.CallOpts, _projectId)
}

// ProjectDetails is a free data retrieval call binding the contract method 0x8dd91a56.
//
// Solidity: function projectDetails(uint256 _projectId) view returns(string projectName, string artist, string description, string website, string license, bool dynamic)
func (_Artblocks *ArtblocksCallerSession) ProjectDetails(_projectId *big.Int) (struct {
	ProjectName string
	Artist      string
	Description string
	Website     string
	License     string
	Dynamic     bool
}, error) {
	return _Artblocks.Contract.ProjectDetails(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToAdditionalPayee is a free data retrieval call binding the contract method 0xd7b044b6.
//
// Solidity: function projectIdToAdditionalPayee(uint256 ) view returns(address)
func (_Artblocks *ArtblocksCaller) ProjectIdToAdditionalPayee(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToAdditionalPayee", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectIdToAdditionalPayee is a free data retrieval call binding the contract method 0xd7b044b6.
//
// Solidity: function projectIdToAdditionalPayee(uint256 ) view returns(address)
func (_Artblocks *ArtblocksSession) ProjectIdToAdditionalPayee(arg0 *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayee(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToAdditionalPayee is a free data retrieval call binding the contract method 0xd7b044b6.
//
// Solidity: function projectIdToAdditionalPayee(uint256 ) view returns(address)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToAdditionalPayee(arg0 *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayee(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToAdditionalPayeePercentage is a free data retrieval call binding the contract method 0xcc74234b.
//
// Solidity: function projectIdToAdditionalPayeePercentage(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksCaller) ProjectIdToAdditionalPayeePercentage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToAdditionalPayeePercentage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectIdToAdditionalPayeePercentage is a free data retrieval call binding the contract method 0xcc74234b.
//
// Solidity: function projectIdToAdditionalPayeePercentage(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksSession) ProjectIdToAdditionalPayeePercentage(arg0 *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeePercentage(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToAdditionalPayeePercentage is a free data retrieval call binding the contract method 0xcc74234b.
//
// Solidity: function projectIdToAdditionalPayeePercentage(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToAdditionalPayeePercentage(arg0 *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeePercentage(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToArtistAddress is a free data retrieval call binding the contract method 0xa47d29cb.
//
// Solidity: function projectIdToArtistAddress(uint256 ) view returns(address)
func (_Artblocks *ArtblocksCaller) ProjectIdToArtistAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToArtistAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectIdToArtistAddress is a free data retrieval call binding the contract method 0xa47d29cb.
//
// Solidity: function projectIdToArtistAddress(uint256 ) view returns(address)
func (_Artblocks *ArtblocksSession) ProjectIdToArtistAddress(arg0 *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToArtistAddress(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToArtistAddress is a free data retrieval call binding the contract method 0xa47d29cb.
//
// Solidity: function projectIdToArtistAddress(uint256 ) view returns(address)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToArtistAddress(arg0 *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToArtistAddress(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToCurrencyAddress is a free data retrieval call binding the contract method 0x498dd0c1.
//
// Solidity: function projectIdToCurrencyAddress(uint256 ) view returns(address)
func (_Artblocks *ArtblocksCaller) ProjectIdToCurrencyAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToCurrencyAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectIdToCurrencyAddress is a free data retrieval call binding the contract method 0x498dd0c1.
//
// Solidity: function projectIdToCurrencyAddress(uint256 ) view returns(address)
func (_Artblocks *ArtblocksSession) ProjectIdToCurrencyAddress(arg0 *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToCurrencyAddress(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToCurrencyAddress is a free data retrieval call binding the contract method 0x498dd0c1.
//
// Solidity: function projectIdToCurrencyAddress(uint256 ) view returns(address)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToCurrencyAddress(arg0 *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToCurrencyAddress(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToCurrencySymbol is a free data retrieval call binding the contract method 0x20927ec9.
//
// Solidity: function projectIdToCurrencySymbol(uint256 ) view returns(string)
func (_Artblocks *ArtblocksCaller) ProjectIdToCurrencySymbol(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToCurrencySymbol", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProjectIdToCurrencySymbol is a free data retrieval call binding the contract method 0x20927ec9.
//
// Solidity: function projectIdToCurrencySymbol(uint256 ) view returns(string)
func (_Artblocks *ArtblocksSession) ProjectIdToCurrencySymbol(arg0 *big.Int) (string, error) {
	return _Artblocks.Contract.ProjectIdToCurrencySymbol(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToCurrencySymbol is a free data retrieval call binding the contract method 0x20927ec9.
//
// Solidity: function projectIdToCurrencySymbol(uint256 ) view returns(string)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToCurrencySymbol(arg0 *big.Int) (string, error) {
	return _Artblocks.Contract.ProjectIdToCurrencySymbol(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToPricePerTokenInWei is a free data retrieval call binding the contract method 0xf70c0f04.
//
// Solidity: function projectIdToPricePerTokenInWei(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksCaller) ProjectIdToPricePerTokenInWei(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToPricePerTokenInWei", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectIdToPricePerTokenInWei is a free data retrieval call binding the contract method 0xf70c0f04.
//
// Solidity: function projectIdToPricePerTokenInWei(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksSession) ProjectIdToPricePerTokenInWei(arg0 *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToPricePerTokenInWei(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToPricePerTokenInWei is a free data retrieval call binding the contract method 0xf70c0f04.
//
// Solidity: function projectIdToPricePerTokenInWei(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToPricePerTokenInWei(arg0 *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToPricePerTokenInWei(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToSecondaryMarketRoyaltyPercentage is a free data retrieval call binding the contract method 0xed8abfda.
//
// Solidity: function projectIdToSecondaryMarketRoyaltyPercentage(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksCaller) ProjectIdToSecondaryMarketRoyaltyPercentage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToSecondaryMarketRoyaltyPercentage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectIdToSecondaryMarketRoyaltyPercentage is a free data retrieval call binding the contract method 0xed8abfda.
//
// Solidity: function projectIdToSecondaryMarketRoyaltyPercentage(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksSession) ProjectIdToSecondaryMarketRoyaltyPercentage(arg0 *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToSecondaryMarketRoyaltyPercentage(&_Artblocks.CallOpts, arg0)
}

// ProjectIdToSecondaryMarketRoyaltyPercentage is a free data retrieval call binding the contract method 0xed8abfda.
//
// Solidity: function projectIdToSecondaryMarketRoyaltyPercentage(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToSecondaryMarketRoyaltyPercentage(arg0 *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToSecondaryMarketRoyaltyPercentage(&_Artblocks.CallOpts, arg0)
}

// ProjectScriptByIndex is a free data retrieval call binding the contract method 0x8c3c9cdd.
//
// Solidity: function projectScriptByIndex(uint256 _projectId, uint256 _index) view returns(string)
func (_Artblocks *ArtblocksCaller) ProjectScriptByIndex(opts *bind.CallOpts, _projectId *big.Int, _index *big.Int) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectScriptByIndex", _projectId, _index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProjectScriptByIndex is a free data retrieval call binding the contract method 0x8c3c9cdd.
//
// Solidity: function projectScriptByIndex(uint256 _projectId, uint256 _index) view returns(string)
func (_Artblocks *ArtblocksSession) ProjectScriptByIndex(_projectId *big.Int, _index *big.Int) (string, error) {
	return _Artblocks.Contract.ProjectScriptByIndex(&_Artblocks.CallOpts, _projectId, _index)
}

// ProjectScriptByIndex is a free data retrieval call binding the contract method 0x8c3c9cdd.
//
// Solidity: function projectScriptByIndex(uint256 _projectId, uint256 _index) view returns(string)
func (_Artblocks *ArtblocksCallerSession) ProjectScriptByIndex(_projectId *big.Int, _index *big.Int) (string, error) {
	return _Artblocks.Contract.ProjectScriptByIndex(&_Artblocks.CallOpts, _projectId, _index)
}

// ProjectScriptInfo is a free data retrieval call binding the contract method 0x4aa6d417.
//
// Solidity: function projectScriptInfo(uint256 _projectId) view returns(string scriptJSON, uint256 scriptCount, bool useHashString, string ipfsHash, bool locked, bool paused)
func (_Artblocks *ArtblocksCaller) ProjectScriptInfo(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ScriptJSON    string
	ScriptCount   *big.Int
	UseHashString bool
	IpfsHash      string
	Locked        bool
	Paused        bool
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectScriptInfo", _projectId)

	outstruct := new(struct {
		ScriptJSON    string
		ScriptCount   *big.Int
		UseHashString bool
		IpfsHash      string
		Locked        bool
		Paused        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScriptJSON = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ScriptCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UseHashString = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.IpfsHash = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Locked = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Paused = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// ProjectScriptInfo is a free data retrieval call binding the contract method 0x4aa6d417.
//
// Solidity: function projectScriptInfo(uint256 _projectId) view returns(string scriptJSON, uint256 scriptCount, bool useHashString, string ipfsHash, bool locked, bool paused)
func (_Artblocks *ArtblocksSession) ProjectScriptInfo(_projectId *big.Int) (struct {
	ScriptJSON    string
	ScriptCount   *big.Int
	UseHashString bool
	IpfsHash      string
	Locked        bool
	Paused        bool
}, error) {
	return _Artblocks.Contract.ProjectScriptInfo(&_Artblocks.CallOpts, _projectId)
}

// ProjectScriptInfo is a free data retrieval call binding the contract method 0x4aa6d417.
//
// Solidity: function projectScriptInfo(uint256 _projectId) view returns(string scriptJSON, uint256 scriptCount, bool useHashString, string ipfsHash, bool locked, bool paused)
func (_Artblocks *ArtblocksCallerSession) ProjectScriptInfo(_projectId *big.Int) (struct {
	ScriptJSON    string
	ScriptCount   *big.Int
	UseHashString bool
	IpfsHash      string
	Locked        bool
	Paused        bool
}, error) {
	return _Artblocks.Contract.ProjectScriptInfo(&_Artblocks.CallOpts, _projectId)
}

// ProjectShowAllTokens is a free data retrieval call binding the contract method 0xbee04f9c.
//
// Solidity: function projectShowAllTokens(uint256 _projectId) view returns(uint256[])
func (_Artblocks *ArtblocksCaller) ProjectShowAllTokens(opts *bind.CallOpts, _projectId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectShowAllTokens", _projectId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ProjectShowAllTokens is a free data retrieval call binding the contract method 0xbee04f9c.
//
// Solidity: function projectShowAllTokens(uint256 _projectId) view returns(uint256[])
func (_Artblocks *ArtblocksSession) ProjectShowAllTokens(_projectId *big.Int) ([]*big.Int, error) {
	return _Artblocks.Contract.ProjectShowAllTokens(&_Artblocks.CallOpts, _projectId)
}

// ProjectShowAllTokens is a free data retrieval call binding the contract method 0xbee04f9c.
//
// Solidity: function projectShowAllTokens(uint256 _projectId) view returns(uint256[])
func (_Artblocks *ArtblocksCallerSession) ProjectShowAllTokens(_projectId *big.Int) ([]*big.Int, error) {
	return _Artblocks.Contract.ProjectShowAllTokens(&_Artblocks.CallOpts, _projectId)
}

// ProjectTokenInfo is a free data retrieval call binding the contract method 0x8c2c3622.
//
// Solidity: function projectTokenInfo(uint256 _projectId) view returns(address artistAddress, uint256 pricePerTokenInWei, uint256 invocations, uint256 maxInvocations, bool active, address additionalPayee, uint256 additionalPayeePercentage, string currency, address currencyAddress)
func (_Artblocks *ArtblocksCaller) ProjectTokenInfo(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ArtistAddress             common.Address
	PricePerTokenInWei        *big.Int
	Invocations               *big.Int
	MaxInvocations            *big.Int
	Active                    bool
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
	Currency                  string
	CurrencyAddress           common.Address
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectTokenInfo", _projectId)

	outstruct := new(struct {
		ArtistAddress             common.Address
		PricePerTokenInWei        *big.Int
		Invocations               *big.Int
		MaxInvocations            *big.Int
		Active                    bool
		AdditionalPayee           common.Address
		AdditionalPayeePercentage *big.Int
		Currency                  string
		CurrencyAddress           common.Address
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
	outstruct.Currency = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.CurrencyAddress = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ProjectTokenInfo is a free data retrieval call binding the contract method 0x8c2c3622.
//
// Solidity: function projectTokenInfo(uint256 _projectId) view returns(address artistAddress, uint256 pricePerTokenInWei, uint256 invocations, uint256 maxInvocations, bool active, address additionalPayee, uint256 additionalPayeePercentage, string currency, address currencyAddress)
func (_Artblocks *ArtblocksSession) ProjectTokenInfo(_projectId *big.Int) (struct {
	ArtistAddress             common.Address
	PricePerTokenInWei        *big.Int
	Invocations               *big.Int
	MaxInvocations            *big.Int
	Active                    bool
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
	Currency                  string
	CurrencyAddress           common.Address
}, error) {
	return _Artblocks.Contract.ProjectTokenInfo(&_Artblocks.CallOpts, _projectId)
}

// ProjectTokenInfo is a free data retrieval call binding the contract method 0x8c2c3622.
//
// Solidity: function projectTokenInfo(uint256 _projectId) view returns(address artistAddress, uint256 pricePerTokenInWei, uint256 invocations, uint256 maxInvocations, bool active, address additionalPayee, uint256 additionalPayeePercentage, string currency, address currencyAddress)
func (_Artblocks *ArtblocksCallerSession) ProjectTokenInfo(_projectId *big.Int) (struct {
	ArtistAddress             common.Address
	PricePerTokenInWei        *big.Int
	Invocations               *big.Int
	MaxInvocations            *big.Int
	Active                    bool
	AdditionalPayee           common.Address
	AdditionalPayeePercentage *big.Int
	Currency                  string
	CurrencyAddress           common.Address
}, error) {
	return _Artblocks.Contract.ProjectTokenInfo(&_Artblocks.CallOpts, _projectId)
}

// ProjectURIInfo is a free data retrieval call binding the contract method 0x2d9c0205.
//
// Solidity: function projectURIInfo(uint256 _projectId) view returns(string projectBaseURI, string projectBaseIpfsURI, bool useIpfs)
func (_Artblocks *ArtblocksCaller) ProjectURIInfo(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ProjectBaseURI     string
	ProjectBaseIpfsURI string
	UseIpfs            bool
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectURIInfo", _projectId)

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
func (_Artblocks *ArtblocksSession) ProjectURIInfo(_projectId *big.Int) (struct {
	ProjectBaseURI     string
	ProjectBaseIpfsURI string
	UseIpfs            bool
}, error) {
	return _Artblocks.Contract.ProjectURIInfo(&_Artblocks.CallOpts, _projectId)
}

// ProjectURIInfo is a free data retrieval call binding the contract method 0x2d9c0205.
//
// Solidity: function projectURIInfo(uint256 _projectId) view returns(string projectBaseURI, string projectBaseIpfsURI, bool useIpfs)
func (_Artblocks *ArtblocksCallerSession) ProjectURIInfo(_projectId *big.Int) (struct {
	ProjectBaseURI     string
	ProjectBaseIpfsURI string
	UseIpfs            bool
}, error) {
	return _Artblocks.Contract.ProjectURIInfo(&_Artblocks.CallOpts, _projectId)
}

// RandomizerContract is a free data retrieval call binding the contract method 0x36c7c12c.
//
// Solidity: function randomizerContract() view returns(address)
func (_Artblocks *ArtblocksCaller) RandomizerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "randomizerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomizerContract is a free data retrieval call binding the contract method 0x36c7c12c.
//
// Solidity: function randomizerContract() view returns(address)
func (_Artblocks *ArtblocksSession) RandomizerContract() (common.Address, error) {
	return _Artblocks.Contract.RandomizerContract(&_Artblocks.CallOpts)
}

// RandomizerContract is a free data retrieval call binding the contract method 0x36c7c12c.
//
// Solidity: function randomizerContract() view returns(address)
func (_Artblocks *ArtblocksCallerSession) RandomizerContract() (common.Address, error) {
	return _Artblocks.Contract.RandomizerContract(&_Artblocks.CallOpts)
}

// StaticIpfsImageLink is a free data retrieval call binding the contract method 0x261eb4e5.
//
// Solidity: function staticIpfsImageLink(uint256 ) view returns(string)
func (_Artblocks *ArtblocksCaller) StaticIpfsImageLink(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "staticIpfsImageLink", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StaticIpfsImageLink is a free data retrieval call binding the contract method 0x261eb4e5.
//
// Solidity: function staticIpfsImageLink(uint256 ) view returns(string)
func (_Artblocks *ArtblocksSession) StaticIpfsImageLink(arg0 *big.Int) (string, error) {
	return _Artblocks.Contract.StaticIpfsImageLink(&_Artblocks.CallOpts, arg0)
}

// StaticIpfsImageLink is a free data retrieval call binding the contract method 0x261eb4e5.
//
// Solidity: function staticIpfsImageLink(uint256 ) view returns(string)
func (_Artblocks *ArtblocksCallerSession) StaticIpfsImageLink(arg0 *big.Int) (string, error) {
	return _Artblocks.Contract.StaticIpfsImageLink(&_Artblocks.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Artblocks *ArtblocksCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Artblocks *ArtblocksSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Artblocks.Contract.SupportsInterface(&_Artblocks.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Artblocks *ArtblocksCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Artblocks.Contract.SupportsInterface(&_Artblocks.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Artblocks *ArtblocksCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Artblocks *ArtblocksSession) Symbol() (string, error) {
	return _Artblocks.Contract.Symbol(&_Artblocks.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Artblocks *ArtblocksCallerSession) Symbol() (string, error) {
	return _Artblocks.Contract.Symbol(&_Artblocks.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Artblocks *ArtblocksCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Artblocks *ArtblocksSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.TokenByIndex(&_Artblocks.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.TokenByIndex(&_Artblocks.CallOpts, index)
}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 ) view returns(bytes32)
func (_Artblocks *ArtblocksCaller) TokenIdToHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "tokenIdToHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 ) view returns(bytes32)
func (_Artblocks *ArtblocksSession) TokenIdToHash(arg0 *big.Int) ([32]byte, error) {
	return _Artblocks.Contract.TokenIdToHash(&_Artblocks.CallOpts, arg0)
}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 ) view returns(bytes32)
func (_Artblocks *ArtblocksCallerSession) TokenIdToHash(arg0 *big.Int) ([32]byte, error) {
	return _Artblocks.Contract.TokenIdToHash(&_Artblocks.CallOpts, arg0)
}

// TokenIdToProjectId is a free data retrieval call binding the contract method 0x1b689c0b.
//
// Solidity: function tokenIdToProjectId(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksCaller) TokenIdToProjectId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "tokenIdToProjectId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToProjectId is a free data retrieval call binding the contract method 0x1b689c0b.
//
// Solidity: function tokenIdToProjectId(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksSession) TokenIdToProjectId(arg0 *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.TokenIdToProjectId(&_Artblocks.CallOpts, arg0)
}

// TokenIdToProjectId is a free data retrieval call binding the contract method 0x1b689c0b.
//
// Solidity: function tokenIdToProjectId(uint256 ) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) TokenIdToProjectId(arg0 *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.TokenIdToProjectId(&_Artblocks.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Artblocks *ArtblocksCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Artblocks *ArtblocksSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.TokenOfOwnerByIndex(&_Artblocks.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.TokenOfOwnerByIndex(&_Artblocks.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Artblocks *ArtblocksCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Artblocks *ArtblocksSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Artblocks.Contract.TokenURI(&_Artblocks.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Artblocks *ArtblocksCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Artblocks.Contract.TokenURI(&_Artblocks.CallOpts, _tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Artblocks *ArtblocksCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Artblocks *ArtblocksSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Artblocks.Contract.TokensOfOwner(&_Artblocks.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Artblocks *ArtblocksCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Artblocks.Contract.TokensOfOwner(&_Artblocks.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Artblocks *ArtblocksCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Artblocks *ArtblocksSession) TotalSupply() (*big.Int, error) {
	return _Artblocks.Contract.TotalSupply(&_Artblocks.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) TotalSupply() (*big.Int, error) {
	return _Artblocks.Contract.TotalSupply(&_Artblocks.CallOpts)
}

// AddMintWhitelisted is a paid mutator transaction binding the contract method 0x8bddb0a6.
//
// Solidity: function addMintWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksTransactor) AddMintWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "addMintWhitelisted", _address)
}

// AddMintWhitelisted is a paid mutator transaction binding the contract method 0x8bddb0a6.
//
// Solidity: function addMintWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksSession) AddMintWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.AddMintWhitelisted(&_Artblocks.TransactOpts, _address)
}

// AddMintWhitelisted is a paid mutator transaction binding the contract method 0x8bddb0a6.
//
// Solidity: function addMintWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksTransactorSession) AddMintWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.AddMintWhitelisted(&_Artblocks.TransactOpts, _address)
}

// AddProject is a paid mutator transaction binding the contract method 0xe3f59c44.
//
// Solidity: function addProject(string _projectName, address _artistAddress, uint256 _pricePerTokenInWei, bool _dynamic) returns()
func (_Artblocks *ArtblocksTransactor) AddProject(opts *bind.TransactOpts, _projectName string, _artistAddress common.Address, _pricePerTokenInWei *big.Int, _dynamic bool) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "addProject", _projectName, _artistAddress, _pricePerTokenInWei, _dynamic)
}

// AddProject is a paid mutator transaction binding the contract method 0xe3f59c44.
//
// Solidity: function addProject(string _projectName, address _artistAddress, uint256 _pricePerTokenInWei, bool _dynamic) returns()
func (_Artblocks *ArtblocksSession) AddProject(_projectName string, _artistAddress common.Address, _pricePerTokenInWei *big.Int, _dynamic bool) (*types.Transaction, error) {
	return _Artblocks.Contract.AddProject(&_Artblocks.TransactOpts, _projectName, _artistAddress, _pricePerTokenInWei, _dynamic)
}

// AddProject is a paid mutator transaction binding the contract method 0xe3f59c44.
//
// Solidity: function addProject(string _projectName, address _artistAddress, uint256 _pricePerTokenInWei, bool _dynamic) returns()
func (_Artblocks *ArtblocksTransactorSession) AddProject(_projectName string, _artistAddress common.Address, _pricePerTokenInWei *big.Int, _dynamic bool) (*types.Transaction, error) {
	return _Artblocks.Contract.AddProject(&_Artblocks.TransactOpts, _projectName, _artistAddress, _pricePerTokenInWei, _dynamic)
}

// AddProjectScript is a paid mutator transaction binding the contract method 0xacad0124.
//
// Solidity: function addProjectScript(uint256 _projectId, string _script) returns()
func (_Artblocks *ArtblocksTransactor) AddProjectScript(opts *bind.TransactOpts, _projectId *big.Int, _script string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "addProjectScript", _projectId, _script)
}

// AddProjectScript is a paid mutator transaction binding the contract method 0xacad0124.
//
// Solidity: function addProjectScript(uint256 _projectId, string _script) returns()
func (_Artblocks *ArtblocksSession) AddProjectScript(_projectId *big.Int, _script string) (*types.Transaction, error) {
	return _Artblocks.Contract.AddProjectScript(&_Artblocks.TransactOpts, _projectId, _script)
}

// AddProjectScript is a paid mutator transaction binding the contract method 0xacad0124.
//
// Solidity: function addProjectScript(uint256 _projectId, string _script) returns()
func (_Artblocks *ArtblocksTransactorSession) AddProjectScript(_projectId *big.Int, _script string) (*types.Transaction, error) {
	return _Artblocks.Contract.AddProjectScript(&_Artblocks.TransactOpts, _projectId, _script)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksTransactor) AddWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "addWhitelisted", _address)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksSession) AddWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.AddWhitelisted(&_Artblocks.TransactOpts, _address)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksTransactorSession) AddWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.AddWhitelisted(&_Artblocks.TransactOpts, _address)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Artblocks *ArtblocksTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Artblocks *ArtblocksSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.Approve(&_Artblocks.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Artblocks *ArtblocksTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.Approve(&_Artblocks.TransactOpts, to, tokenId)
}

// ClearTokenIpfsImageUri is a paid mutator transaction binding the contract method 0x27901822.
//
// Solidity: function clearTokenIpfsImageUri(uint256 _tokenId) returns()
func (_Artblocks *ArtblocksTransactor) ClearTokenIpfsImageUri(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "clearTokenIpfsImageUri", _tokenId)
}

// ClearTokenIpfsImageUri is a paid mutator transaction binding the contract method 0x27901822.
//
// Solidity: function clearTokenIpfsImageUri(uint256 _tokenId) returns()
func (_Artblocks *ArtblocksSession) ClearTokenIpfsImageUri(_tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ClearTokenIpfsImageUri(&_Artblocks.TransactOpts, _tokenId)
}

// ClearTokenIpfsImageUri is a paid mutator transaction binding the contract method 0x27901822.
//
// Solidity: function clearTokenIpfsImageUri(uint256 _tokenId) returns()
func (_Artblocks *ArtblocksTransactorSession) ClearTokenIpfsImageUri(_tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ClearTokenIpfsImageUri(&_Artblocks.TransactOpts, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x0d4d1513.
//
// Solidity: function mint(address _to, uint256 _projectId, address _by) returns(uint256 _tokenId)
func (_Artblocks *ArtblocksTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _projectId *big.Int, _by common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "mint", _to, _projectId, _by)
}

// Mint is a paid mutator transaction binding the contract method 0x0d4d1513.
//
// Solidity: function mint(address _to, uint256 _projectId, address _by) returns(uint256 _tokenId)
func (_Artblocks *ArtblocksSession) Mint(_to common.Address, _projectId *big.Int, _by common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.Mint(&_Artblocks.TransactOpts, _to, _projectId, _by)
}

// Mint is a paid mutator transaction binding the contract method 0x0d4d1513.
//
// Solidity: function mint(address _to, uint256 _projectId, address _by) returns(uint256 _tokenId)
func (_Artblocks *ArtblocksTransactorSession) Mint(_to common.Address, _projectId *big.Int, _by common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.Mint(&_Artblocks.TransactOpts, _to, _projectId, _by)
}

// OverrideTokenDynamicImageWithIpfsLink is a paid mutator transaction binding the contract method 0x93961c66.
//
// Solidity: function overrideTokenDynamicImageWithIpfsLink(uint256 _tokenId, string _ipfsHash) returns()
func (_Artblocks *ArtblocksTransactor) OverrideTokenDynamicImageWithIpfsLink(opts *bind.TransactOpts, _tokenId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "overrideTokenDynamicImageWithIpfsLink", _tokenId, _ipfsHash)
}

// OverrideTokenDynamicImageWithIpfsLink is a paid mutator transaction binding the contract method 0x93961c66.
//
// Solidity: function overrideTokenDynamicImageWithIpfsLink(uint256 _tokenId, string _ipfsHash) returns()
func (_Artblocks *ArtblocksSession) OverrideTokenDynamicImageWithIpfsLink(_tokenId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _Artblocks.Contract.OverrideTokenDynamicImageWithIpfsLink(&_Artblocks.TransactOpts, _tokenId, _ipfsHash)
}

// OverrideTokenDynamicImageWithIpfsLink is a paid mutator transaction binding the contract method 0x93961c66.
//
// Solidity: function overrideTokenDynamicImageWithIpfsLink(uint256 _tokenId, string _ipfsHash) returns()
func (_Artblocks *ArtblocksTransactorSession) OverrideTokenDynamicImageWithIpfsLink(_tokenId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _Artblocks.Contract.OverrideTokenDynamicImageWithIpfsLink(&_Artblocks.TransactOpts, _tokenId, _ipfsHash)
}

// RemoveMintWhitelisted is a paid mutator transaction binding the contract method 0x867f1a3b.
//
// Solidity: function removeMintWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksTransactor) RemoveMintWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "removeMintWhitelisted", _address)
}

// RemoveMintWhitelisted is a paid mutator transaction binding the contract method 0x867f1a3b.
//
// Solidity: function removeMintWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksSession) RemoveMintWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.RemoveMintWhitelisted(&_Artblocks.TransactOpts, _address)
}

// RemoveMintWhitelisted is a paid mutator transaction binding the contract method 0x867f1a3b.
//
// Solidity: function removeMintWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksTransactorSession) RemoveMintWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.RemoveMintWhitelisted(&_Artblocks.TransactOpts, _address)
}

// RemoveProjectLastScript is a paid mutator transaction binding the contract method 0xdb2ff861.
//
// Solidity: function removeProjectLastScript(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactor) RemoveProjectLastScript(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "removeProjectLastScript", _projectId)
}

// RemoveProjectLastScript is a paid mutator transaction binding the contract method 0xdb2ff861.
//
// Solidity: function removeProjectLastScript(uint256 _projectId) returns()
func (_Artblocks *ArtblocksSession) RemoveProjectLastScript(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.RemoveProjectLastScript(&_Artblocks.TransactOpts, _projectId)
}

// RemoveProjectLastScript is a paid mutator transaction binding the contract method 0xdb2ff861.
//
// Solidity: function removeProjectLastScript(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactorSession) RemoveProjectLastScript(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.RemoveProjectLastScript(&_Artblocks.TransactOpts, _projectId)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksTransactor) RemoveWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "removeWhitelisted", _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksSession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.RemoveWhitelisted(&_Artblocks.TransactOpts, _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_Artblocks *ArtblocksTransactorSession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.RemoveWhitelisted(&_Artblocks.TransactOpts, _address)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Artblocks *ArtblocksTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Artblocks *ArtblocksSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.SafeTransferFrom(&_Artblocks.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Artblocks *ArtblocksTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.SafeTransferFrom(&_Artblocks.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Artblocks *ArtblocksTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Artblocks *ArtblocksSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Artblocks.Contract.SafeTransferFrom0(&_Artblocks.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Artblocks *ArtblocksTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Artblocks.Contract.SafeTransferFrom0(&_Artblocks.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Artblocks *ArtblocksTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Artblocks *ArtblocksSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _Artblocks.Contract.SetApprovalForAll(&_Artblocks.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Artblocks *ArtblocksTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _Artblocks.Contract.SetApprovalForAll(&_Artblocks.TransactOpts, to, approved)
}

// ToggleProjectIsActive is a paid mutator transaction binding the contract method 0xd03c390c.
//
// Solidity: function toggleProjectIsActive(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactor) ToggleProjectIsActive(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "toggleProjectIsActive", _projectId)
}

// ToggleProjectIsActive is a paid mutator transaction binding the contract method 0xd03c390c.
//
// Solidity: function toggleProjectIsActive(uint256 _projectId) returns()
func (_Artblocks *ArtblocksSession) ToggleProjectIsActive(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectIsActive(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectIsActive is a paid mutator transaction binding the contract method 0xd03c390c.
//
// Solidity: function toggleProjectIsActive(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactorSession) ToggleProjectIsActive(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectIsActive(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectIsDynamic is a paid mutator transaction binding the contract method 0x3bdbd5c4.
//
// Solidity: function toggleProjectIsDynamic(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactor) ToggleProjectIsDynamic(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "toggleProjectIsDynamic", _projectId)
}

// ToggleProjectIsDynamic is a paid mutator transaction binding the contract method 0x3bdbd5c4.
//
// Solidity: function toggleProjectIsDynamic(uint256 _projectId) returns()
func (_Artblocks *ArtblocksSession) ToggleProjectIsDynamic(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectIsDynamic(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectIsDynamic is a paid mutator transaction binding the contract method 0x3bdbd5c4.
//
// Solidity: function toggleProjectIsDynamic(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactorSession) ToggleProjectIsDynamic(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectIsDynamic(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectIsLocked is a paid mutator transaction binding the contract method 0x8ba8f14d.
//
// Solidity: function toggleProjectIsLocked(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactor) ToggleProjectIsLocked(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "toggleProjectIsLocked", _projectId)
}

// ToggleProjectIsLocked is a paid mutator transaction binding the contract method 0x8ba8f14d.
//
// Solidity: function toggleProjectIsLocked(uint256 _projectId) returns()
func (_Artblocks *ArtblocksSession) ToggleProjectIsLocked(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectIsLocked(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectIsLocked is a paid mutator transaction binding the contract method 0x8ba8f14d.
//
// Solidity: function toggleProjectIsLocked(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactorSession) ToggleProjectIsLocked(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectIsLocked(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectIsPaused is a paid mutator transaction binding the contract method 0xa11ec70a.
//
// Solidity: function toggleProjectIsPaused(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactor) ToggleProjectIsPaused(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "toggleProjectIsPaused", _projectId)
}

// ToggleProjectIsPaused is a paid mutator transaction binding the contract method 0xa11ec70a.
//
// Solidity: function toggleProjectIsPaused(uint256 _projectId) returns()
func (_Artblocks *ArtblocksSession) ToggleProjectIsPaused(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectIsPaused(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectIsPaused is a paid mutator transaction binding the contract method 0xa11ec70a.
//
// Solidity: function toggleProjectIsPaused(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactorSession) ToggleProjectIsPaused(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectIsPaused(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectUseHashString is a paid mutator transaction binding the contract method 0xdce5d858.
//
// Solidity: function toggleProjectUseHashString(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactor) ToggleProjectUseHashString(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "toggleProjectUseHashString", _projectId)
}

// ToggleProjectUseHashString is a paid mutator transaction binding the contract method 0xdce5d858.
//
// Solidity: function toggleProjectUseHashString(uint256 _projectId) returns()
func (_Artblocks *ArtblocksSession) ToggleProjectUseHashString(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectUseHashString(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectUseHashString is a paid mutator transaction binding the contract method 0xdce5d858.
//
// Solidity: function toggleProjectUseHashString(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactorSession) ToggleProjectUseHashString(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectUseHashString(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectUseIpfsForStatic is a paid mutator transaction binding the contract method 0x5c088dcc.
//
// Solidity: function toggleProjectUseIpfsForStatic(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactor) ToggleProjectUseIpfsForStatic(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "toggleProjectUseIpfsForStatic", _projectId)
}

// ToggleProjectUseIpfsForStatic is a paid mutator transaction binding the contract method 0x5c088dcc.
//
// Solidity: function toggleProjectUseIpfsForStatic(uint256 _projectId) returns()
func (_Artblocks *ArtblocksSession) ToggleProjectUseIpfsForStatic(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectUseIpfsForStatic(&_Artblocks.TransactOpts, _projectId)
}

// ToggleProjectUseIpfsForStatic is a paid mutator transaction binding the contract method 0x5c088dcc.
//
// Solidity: function toggleProjectUseIpfsForStatic(uint256 _projectId) returns()
func (_Artblocks *ArtblocksTransactorSession) ToggleProjectUseIpfsForStatic(_projectId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ToggleProjectUseIpfsForStatic(&_Artblocks.TransactOpts, _projectId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Artblocks *ArtblocksTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Artblocks *ArtblocksSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.TransferFrom(&_Artblocks.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Artblocks *ArtblocksTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.TransferFrom(&_Artblocks.TransactOpts, from, to, tokenId)
}

// UpdateArtblocksAddress is a paid mutator transaction binding the contract method 0x06e1db17.
//
// Solidity: function updateArtblocksAddress(address _artblocksAddress) returns()
func (_Artblocks *ArtblocksTransactor) UpdateArtblocksAddress(opts *bind.TransactOpts, _artblocksAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateArtblocksAddress", _artblocksAddress)
}

// UpdateArtblocksAddress is a paid mutator transaction binding the contract method 0x06e1db17.
//
// Solidity: function updateArtblocksAddress(address _artblocksAddress) returns()
func (_Artblocks *ArtblocksSession) UpdateArtblocksAddress(_artblocksAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksAddress(&_Artblocks.TransactOpts, _artblocksAddress)
}

// UpdateArtblocksAddress is a paid mutator transaction binding the contract method 0x06e1db17.
//
// Solidity: function updateArtblocksAddress(address _artblocksAddress) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateArtblocksAddress(_artblocksAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksAddress(&_Artblocks.TransactOpts, _artblocksAddress)
}

// UpdateArtblocksPercentage is a paid mutator transaction binding the contract method 0xed6df982.
//
// Solidity: function updateArtblocksPercentage(uint256 _artblocksPercentage) returns()
func (_Artblocks *ArtblocksTransactor) UpdateArtblocksPercentage(opts *bind.TransactOpts, _artblocksPercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateArtblocksPercentage", _artblocksPercentage)
}

// UpdateArtblocksPercentage is a paid mutator transaction binding the contract method 0xed6df982.
//
// Solidity: function updateArtblocksPercentage(uint256 _artblocksPercentage) returns()
func (_Artblocks *ArtblocksSession) UpdateArtblocksPercentage(_artblocksPercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksPercentage(&_Artblocks.TransactOpts, _artblocksPercentage)
}

// UpdateArtblocksPercentage is a paid mutator transaction binding the contract method 0xed6df982.
//
// Solidity: function updateArtblocksPercentage(uint256 _artblocksPercentage) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateArtblocksPercentage(_artblocksPercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksPercentage(&_Artblocks.TransactOpts, _artblocksPercentage)
}

// UpdateProjectAdditionalPayeeInfo is a paid mutator transaction binding the contract method 0xe13208b4.
//
// Solidity: function updateProjectAdditionalPayeeInfo(uint256 _projectId, address _additionalPayee, uint256 _additionalPayeePercentage) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectAdditionalPayeeInfo(opts *bind.TransactOpts, _projectId *big.Int, _additionalPayee common.Address, _additionalPayeePercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectAdditionalPayeeInfo", _projectId, _additionalPayee, _additionalPayeePercentage)
}

// UpdateProjectAdditionalPayeeInfo is a paid mutator transaction binding the contract method 0xe13208b4.
//
// Solidity: function updateProjectAdditionalPayeeInfo(uint256 _projectId, address _additionalPayee, uint256 _additionalPayeePercentage) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectAdditionalPayeeInfo(_projectId *big.Int, _additionalPayee common.Address, _additionalPayeePercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectAdditionalPayeeInfo(&_Artblocks.TransactOpts, _projectId, _additionalPayee, _additionalPayeePercentage)
}

// UpdateProjectAdditionalPayeeInfo is a paid mutator transaction binding the contract method 0xe13208b4.
//
// Solidity: function updateProjectAdditionalPayeeInfo(uint256 _projectId, address _additionalPayee, uint256 _additionalPayeePercentage) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectAdditionalPayeeInfo(_projectId *big.Int, _additionalPayee common.Address, _additionalPayeePercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectAdditionalPayeeInfo(&_Artblocks.TransactOpts, _projectId, _additionalPayee, _additionalPayeePercentage)
}

// UpdateProjectArtistAddress is a paid mutator transaction binding the contract method 0x69d14faf.
//
// Solidity: function updateProjectArtistAddress(uint256 _projectId, address _artistAddress) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectArtistAddress(opts *bind.TransactOpts, _projectId *big.Int, _artistAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectArtistAddress", _projectId, _artistAddress)
}

// UpdateProjectArtistAddress is a paid mutator transaction binding the contract method 0x69d14faf.
//
// Solidity: function updateProjectArtistAddress(uint256 _projectId, address _artistAddress) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectArtistAddress(_projectId *big.Int, _artistAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectArtistAddress(&_Artblocks.TransactOpts, _projectId, _artistAddress)
}

// UpdateProjectArtistAddress is a paid mutator transaction binding the contract method 0x69d14faf.
//
// Solidity: function updateProjectArtistAddress(uint256 _projectId, address _artistAddress) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectArtistAddress(_projectId *big.Int, _artistAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectArtistAddress(&_Artblocks.TransactOpts, _projectId, _artistAddress)
}

// UpdateProjectArtistName is a paid mutator transaction binding the contract method 0xb7b04fae.
//
// Solidity: function updateProjectArtistName(uint256 _projectId, string _projectArtistName) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectArtistName(opts *bind.TransactOpts, _projectId *big.Int, _projectArtistName string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectArtistName", _projectId, _projectArtistName)
}

// UpdateProjectArtistName is a paid mutator transaction binding the contract method 0xb7b04fae.
//
// Solidity: function updateProjectArtistName(uint256 _projectId, string _projectArtistName) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectArtistName(_projectId *big.Int, _projectArtistName string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectArtistName(&_Artblocks.TransactOpts, _projectId, _projectArtistName)
}

// UpdateProjectArtistName is a paid mutator transaction binding the contract method 0xb7b04fae.
//
// Solidity: function updateProjectArtistName(uint256 _projectId, string _projectArtistName) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectArtistName(_projectId *big.Int, _projectArtistName string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectArtistName(&_Artblocks.TransactOpts, _projectId, _projectArtistName)
}

// UpdateProjectBaseIpfsURI is a paid mutator transaction binding the contract method 0x6bd5d591.
//
// Solidity: function updateProjectBaseIpfsURI(uint256 _projectId, string _projectBaseIpfsURI) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectBaseIpfsURI(opts *bind.TransactOpts, _projectId *big.Int, _projectBaseIpfsURI string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectBaseIpfsURI", _projectId, _projectBaseIpfsURI)
}

// UpdateProjectBaseIpfsURI is a paid mutator transaction binding the contract method 0x6bd5d591.
//
// Solidity: function updateProjectBaseIpfsURI(uint256 _projectId, string _projectBaseIpfsURI) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectBaseIpfsURI(_projectId *big.Int, _projectBaseIpfsURI string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectBaseIpfsURI(&_Artblocks.TransactOpts, _projectId, _projectBaseIpfsURI)
}

// UpdateProjectBaseIpfsURI is a paid mutator transaction binding the contract method 0x6bd5d591.
//
// Solidity: function updateProjectBaseIpfsURI(uint256 _projectId, string _projectBaseIpfsURI) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectBaseIpfsURI(_projectId *big.Int, _projectBaseIpfsURI string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectBaseIpfsURI(&_Artblocks.TransactOpts, _projectId, _projectBaseIpfsURI)
}

// UpdateProjectBaseURI is a paid mutator transaction binding the contract method 0x3e48e848.
//
// Solidity: function updateProjectBaseURI(uint256 _projectId, string _newBaseURI) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectBaseURI(opts *bind.TransactOpts, _projectId *big.Int, _newBaseURI string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectBaseURI", _projectId, _newBaseURI)
}

// UpdateProjectBaseURI is a paid mutator transaction binding the contract method 0x3e48e848.
//
// Solidity: function updateProjectBaseURI(uint256 _projectId, string _newBaseURI) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectBaseURI(_projectId *big.Int, _newBaseURI string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectBaseURI(&_Artblocks.TransactOpts, _projectId, _newBaseURI)
}

// UpdateProjectBaseURI is a paid mutator transaction binding the contract method 0x3e48e848.
//
// Solidity: function updateProjectBaseURI(uint256 _projectId, string _newBaseURI) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectBaseURI(_projectId *big.Int, _newBaseURI string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectBaseURI(&_Artblocks.TransactOpts, _projectId, _newBaseURI)
}

// UpdateProjectCurrencyInfo is a paid mutator transaction binding the contract method 0xd195b365.
//
// Solidity: function updateProjectCurrencyInfo(uint256 _projectId, string _currencySymbol, address _currencyAddress) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectCurrencyInfo(opts *bind.TransactOpts, _projectId *big.Int, _currencySymbol string, _currencyAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectCurrencyInfo", _projectId, _currencySymbol, _currencyAddress)
}

// UpdateProjectCurrencyInfo is a paid mutator transaction binding the contract method 0xd195b365.
//
// Solidity: function updateProjectCurrencyInfo(uint256 _projectId, string _currencySymbol, address _currencyAddress) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectCurrencyInfo(_projectId *big.Int, _currencySymbol string, _currencyAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectCurrencyInfo(&_Artblocks.TransactOpts, _projectId, _currencySymbol, _currencyAddress)
}

// UpdateProjectCurrencyInfo is a paid mutator transaction binding the contract method 0xd195b365.
//
// Solidity: function updateProjectCurrencyInfo(uint256 _projectId, string _currencySymbol, address _currencyAddress) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectCurrencyInfo(_projectId *big.Int, _currencySymbol string, _currencyAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectCurrencyInfo(&_Artblocks.TransactOpts, _projectId, _currencySymbol, _currencyAddress)
}

// UpdateProjectDescription is a paid mutator transaction binding the contract method 0xa3b2cca6.
//
// Solidity: function updateProjectDescription(uint256 _projectId, string _projectDescription) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectDescription(opts *bind.TransactOpts, _projectId *big.Int, _projectDescription string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectDescription", _projectId, _projectDescription)
}

// UpdateProjectDescription is a paid mutator transaction binding the contract method 0xa3b2cca6.
//
// Solidity: function updateProjectDescription(uint256 _projectId, string _projectDescription) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectDescription(_projectId *big.Int, _projectDescription string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectDescription(&_Artblocks.TransactOpts, _projectId, _projectDescription)
}

// UpdateProjectDescription is a paid mutator transaction binding the contract method 0xa3b2cca6.
//
// Solidity: function updateProjectDescription(uint256 _projectId, string _projectDescription) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectDescription(_projectId *big.Int, _projectDescription string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectDescription(&_Artblocks.TransactOpts, _projectId, _projectDescription)
}

// UpdateProjectIpfsHash is a paid mutator transaction binding the contract method 0x3fef6c2a.
//
// Solidity: function updateProjectIpfsHash(uint256 _projectId, string _ipfsHash) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectIpfsHash(opts *bind.TransactOpts, _projectId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectIpfsHash", _projectId, _ipfsHash)
}

// UpdateProjectIpfsHash is a paid mutator transaction binding the contract method 0x3fef6c2a.
//
// Solidity: function updateProjectIpfsHash(uint256 _projectId, string _ipfsHash) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectIpfsHash(_projectId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectIpfsHash(&_Artblocks.TransactOpts, _projectId, _ipfsHash)
}

// UpdateProjectIpfsHash is a paid mutator transaction binding the contract method 0x3fef6c2a.
//
// Solidity: function updateProjectIpfsHash(uint256 _projectId, string _ipfsHash) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectIpfsHash(_projectId *big.Int, _ipfsHash string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectIpfsHash(&_Artblocks.TransactOpts, _projectId, _ipfsHash)
}

// UpdateProjectLicense is a paid mutator transaction binding the contract method 0x25b75d68.
//
// Solidity: function updateProjectLicense(uint256 _projectId, string _projectLicense) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectLicense(opts *bind.TransactOpts, _projectId *big.Int, _projectLicense string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectLicense", _projectId, _projectLicense)
}

// UpdateProjectLicense is a paid mutator transaction binding the contract method 0x25b75d68.
//
// Solidity: function updateProjectLicense(uint256 _projectId, string _projectLicense) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectLicense(_projectId *big.Int, _projectLicense string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectLicense(&_Artblocks.TransactOpts, _projectId, _projectLicense)
}

// UpdateProjectLicense is a paid mutator transaction binding the contract method 0x25b75d68.
//
// Solidity: function updateProjectLicense(uint256 _projectId, string _projectLicense) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectLicense(_projectId *big.Int, _projectLicense string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectLicense(&_Artblocks.TransactOpts, _projectId, _projectLicense)
}

// UpdateProjectMaxInvocations is a paid mutator transaction binding the contract method 0x826fc391.
//
// Solidity: function updateProjectMaxInvocations(uint256 _projectId, uint256 _maxInvocations) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectMaxInvocations(opts *bind.TransactOpts, _projectId *big.Int, _maxInvocations *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectMaxInvocations", _projectId, _maxInvocations)
}

// UpdateProjectMaxInvocations is a paid mutator transaction binding the contract method 0x826fc391.
//
// Solidity: function updateProjectMaxInvocations(uint256 _projectId, uint256 _maxInvocations) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectMaxInvocations(_projectId *big.Int, _maxInvocations *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectMaxInvocations(&_Artblocks.TransactOpts, _projectId, _maxInvocations)
}

// UpdateProjectMaxInvocations is a paid mutator transaction binding the contract method 0x826fc391.
//
// Solidity: function updateProjectMaxInvocations(uint256 _projectId, uint256 _maxInvocations) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectMaxInvocations(_projectId *big.Int, _maxInvocations *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectMaxInvocations(&_Artblocks.TransactOpts, _projectId, _maxInvocations)
}

// UpdateProjectName is a paid mutator transaction binding the contract method 0x0d170673.
//
// Solidity: function updateProjectName(uint256 _projectId, string _projectName) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectName(opts *bind.TransactOpts, _projectId *big.Int, _projectName string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectName", _projectId, _projectName)
}

// UpdateProjectName is a paid mutator transaction binding the contract method 0x0d170673.
//
// Solidity: function updateProjectName(uint256 _projectId, string _projectName) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectName(_projectId *big.Int, _projectName string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectName(&_Artblocks.TransactOpts, _projectId, _projectName)
}

// UpdateProjectName is a paid mutator transaction binding the contract method 0x0d170673.
//
// Solidity: function updateProjectName(uint256 _projectId, string _projectName) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectName(_projectId *big.Int, _projectName string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectName(&_Artblocks.TransactOpts, _projectId, _projectName)
}

// UpdateProjectPricePerTokenInWei is a paid mutator transaction binding the contract method 0x97dc86cf.
//
// Solidity: function updateProjectPricePerTokenInWei(uint256 _projectId, uint256 _pricePerTokenInWei) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectPricePerTokenInWei(opts *bind.TransactOpts, _projectId *big.Int, _pricePerTokenInWei *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectPricePerTokenInWei", _projectId, _pricePerTokenInWei)
}

// UpdateProjectPricePerTokenInWei is a paid mutator transaction binding the contract method 0x97dc86cf.
//
// Solidity: function updateProjectPricePerTokenInWei(uint256 _projectId, uint256 _pricePerTokenInWei) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectPricePerTokenInWei(_projectId *big.Int, _pricePerTokenInWei *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectPricePerTokenInWei(&_Artblocks.TransactOpts, _projectId, _pricePerTokenInWei)
}

// UpdateProjectPricePerTokenInWei is a paid mutator transaction binding the contract method 0x97dc86cf.
//
// Solidity: function updateProjectPricePerTokenInWei(uint256 _projectId, uint256 _pricePerTokenInWei) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectPricePerTokenInWei(_projectId *big.Int, _pricePerTokenInWei *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectPricePerTokenInWei(&_Artblocks.TransactOpts, _projectId, _pricePerTokenInWei)
}

// UpdateProjectScript is a paid mutator transaction binding the contract method 0xb1656ba3.
//
// Solidity: function updateProjectScript(uint256 _projectId, uint256 _scriptId, string _script) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectScript(opts *bind.TransactOpts, _projectId *big.Int, _scriptId *big.Int, _script string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectScript", _projectId, _scriptId, _script)
}

// UpdateProjectScript is a paid mutator transaction binding the contract method 0xb1656ba3.
//
// Solidity: function updateProjectScript(uint256 _projectId, uint256 _scriptId, string _script) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectScript(_projectId *big.Int, _scriptId *big.Int, _script string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectScript(&_Artblocks.TransactOpts, _projectId, _scriptId, _script)
}

// UpdateProjectScript is a paid mutator transaction binding the contract method 0xb1656ba3.
//
// Solidity: function updateProjectScript(uint256 _projectId, uint256 _scriptId, string _script) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectScript(_projectId *big.Int, _scriptId *big.Int, _script string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectScript(&_Artblocks.TransactOpts, _projectId, _scriptId, _script)
}

// UpdateProjectScriptJSON is a paid mutator transaction binding the contract method 0xc6d73231.
//
// Solidity: function updateProjectScriptJSON(uint256 _projectId, string _projectScriptJSON) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectScriptJSON(opts *bind.TransactOpts, _projectId *big.Int, _projectScriptJSON string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectScriptJSON", _projectId, _projectScriptJSON)
}

// UpdateProjectScriptJSON is a paid mutator transaction binding the contract method 0xc6d73231.
//
// Solidity: function updateProjectScriptJSON(uint256 _projectId, string _projectScriptJSON) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectScriptJSON(_projectId *big.Int, _projectScriptJSON string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectScriptJSON(&_Artblocks.TransactOpts, _projectId, _projectScriptJSON)
}

// UpdateProjectScriptJSON is a paid mutator transaction binding the contract method 0xc6d73231.
//
// Solidity: function updateProjectScriptJSON(uint256 _projectId, string _projectScriptJSON) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectScriptJSON(_projectId *big.Int, _projectScriptJSON string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectScriptJSON(&_Artblocks.TransactOpts, _projectId, _projectScriptJSON)
}

// UpdateProjectSecondaryMarketRoyaltyPercentage is a paid mutator transaction binding the contract method 0xc34a03b5.
//
// Solidity: function updateProjectSecondaryMarketRoyaltyPercentage(uint256 _projectId, uint256 _secondMarketRoyalty) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectSecondaryMarketRoyaltyPercentage(opts *bind.TransactOpts, _projectId *big.Int, _secondMarketRoyalty *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectSecondaryMarketRoyaltyPercentage", _projectId, _secondMarketRoyalty)
}

// UpdateProjectSecondaryMarketRoyaltyPercentage is a paid mutator transaction binding the contract method 0xc34a03b5.
//
// Solidity: function updateProjectSecondaryMarketRoyaltyPercentage(uint256 _projectId, uint256 _secondMarketRoyalty) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectSecondaryMarketRoyaltyPercentage(_projectId *big.Int, _secondMarketRoyalty *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectSecondaryMarketRoyaltyPercentage(&_Artblocks.TransactOpts, _projectId, _secondMarketRoyalty)
}

// UpdateProjectSecondaryMarketRoyaltyPercentage is a paid mutator transaction binding the contract method 0xc34a03b5.
//
// Solidity: function updateProjectSecondaryMarketRoyaltyPercentage(uint256 _projectId, uint256 _secondMarketRoyalty) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectSecondaryMarketRoyaltyPercentage(_projectId *big.Int, _secondMarketRoyalty *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectSecondaryMarketRoyaltyPercentage(&_Artblocks.TransactOpts, _projectId, _secondMarketRoyalty)
}

// UpdateProjectWebsite is a paid mutator transaction binding the contract method 0x37859963.
//
// Solidity: function updateProjectWebsite(uint256 _projectId, string _projectWebsite) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectWebsite(opts *bind.TransactOpts, _projectId *big.Int, _projectWebsite string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectWebsite", _projectId, _projectWebsite)
}

// UpdateProjectWebsite is a paid mutator transaction binding the contract method 0x37859963.
//
// Solidity: function updateProjectWebsite(uint256 _projectId, string _projectWebsite) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectWebsite(_projectId *big.Int, _projectWebsite string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectWebsite(&_Artblocks.TransactOpts, _projectId, _projectWebsite)
}

// UpdateProjectWebsite is a paid mutator transaction binding the contract method 0x37859963.
//
// Solidity: function updateProjectWebsite(uint256 _projectId, string _projectWebsite) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectWebsite(_projectId *big.Int, _projectWebsite string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectWebsite(&_Artblocks.TransactOpts, _projectId, _projectWebsite)
}

// UpdateRandomizerAddress is a paid mutator transaction binding the contract method 0x6c907b7f.
//
// Solidity: function updateRandomizerAddress(address _randomizerAddress) returns()
func (_Artblocks *ArtblocksTransactor) UpdateRandomizerAddress(opts *bind.TransactOpts, _randomizerAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateRandomizerAddress", _randomizerAddress)
}

// UpdateRandomizerAddress is a paid mutator transaction binding the contract method 0x6c907b7f.
//
// Solidity: function updateRandomizerAddress(address _randomizerAddress) returns()
func (_Artblocks *ArtblocksSession) UpdateRandomizerAddress(_randomizerAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateRandomizerAddress(&_Artblocks.TransactOpts, _randomizerAddress)
}

// UpdateRandomizerAddress is a paid mutator transaction binding the contract method 0x6c907b7f.
//
// Solidity: function updateRandomizerAddress(address _randomizerAddress) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateRandomizerAddress(_randomizerAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateRandomizerAddress(&_Artblocks.TransactOpts, _randomizerAddress)
}

// ArtblocksApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Artblocks contract.
type ArtblocksApprovalIterator struct {
	Event *ArtblocksApproval // Event containing the contract specifics and raw log

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
func (it *ArtblocksApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksApproval)
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
		it.Event = new(ArtblocksApproval)
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
func (it *ArtblocksApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksApproval represents a Approval event raised by the Artblocks contract.
type ArtblocksApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Artblocks *ArtblocksFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ArtblocksApprovalIterator, error) {

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

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksApprovalIterator{contract: _Artblocks.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Artblocks *ArtblocksFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ArtblocksApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksApproval)
				if err := _Artblocks.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Artblocks *ArtblocksFilterer) ParseApproval(log types.Log) (*ArtblocksApproval, error) {
	event := new(ArtblocksApproval)
	if err := _Artblocks.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtblocksApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Artblocks contract.
type ArtblocksApprovalForAllIterator struct {
	Event *ArtblocksApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ArtblocksApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksApprovalForAll)
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
		it.Event = new(ArtblocksApprovalForAll)
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
func (it *ArtblocksApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksApprovalForAll represents a ApprovalForAll event raised by the Artblocks contract.
type ArtblocksApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Artblocks *ArtblocksFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ArtblocksApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksApprovalForAllIterator{contract: _Artblocks.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Artblocks *ArtblocksFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ArtblocksApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksApprovalForAll)
				if err := _Artblocks.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Artblocks *ArtblocksFilterer) ParseApprovalForAll(log types.Log) (*ArtblocksApprovalForAll, error) {
	event := new(ArtblocksApprovalForAll)
	if err := _Artblocks.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtblocksMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Artblocks contract.
type ArtblocksMintIterator struct {
	Event *ArtblocksMint // Event containing the contract specifics and raw log

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
func (it *ArtblocksMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksMint)
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
		it.Event = new(ArtblocksMint)
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
func (it *ArtblocksMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksMint represents a Mint event raised by the Artblocks contract.
type ArtblocksMint struct {
	To        common.Address
	TokenId   *big.Int
	ProjectId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId, uint256 indexed _projectId)
func (_Artblocks *ArtblocksFilterer) FilterMint(opts *bind.FilterOpts, _to []common.Address, _tokenId []*big.Int, _projectId []*big.Int) (*ArtblocksMintIterator, error) {

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

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "Mint", _toRule, _tokenIdRule, _projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksMintIterator{contract: _Artblocks.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId, uint256 indexed _projectId)
func (_Artblocks *ArtblocksFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ArtblocksMint, _to []common.Address, _tokenId []*big.Int, _projectId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "Mint", _toRule, _tokenIdRule, _projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksMint)
				if err := _Artblocks.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId, uint256 indexed _projectId)
func (_Artblocks *ArtblocksFilterer) ParseMint(log types.Log) (*ArtblocksMint, error) {
	event := new(ArtblocksMint)
	if err := _Artblocks.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtblocksTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Artblocks contract.
type ArtblocksTransferIterator struct {
	Event *ArtblocksTransfer // Event containing the contract specifics and raw log

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
func (it *ArtblocksTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksTransfer)
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
		it.Event = new(ArtblocksTransfer)
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
func (it *ArtblocksTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksTransfer represents a Transfer event raised by the Artblocks contract.
type ArtblocksTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Artblocks *ArtblocksFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ArtblocksTransferIterator, error) {

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

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksTransferIterator{contract: _Artblocks.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Artblocks *ArtblocksFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ArtblocksTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksTransfer)
				if err := _Artblocks.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Artblocks *ArtblocksFilterer) ParseTransfer(log types.Log) (*ArtblocksTransfer, error) {
	event := new(ArtblocksTransfer)
	if err := _Artblocks.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
