// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenArt721CoreV3

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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_randomizerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_adminACLContract\",\"type\":\"address\"},{\"internalType\":\"uint248\",\"name\":\"_startingProjectId\",\"type\":\"uint248\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"AcceptedArtistAddressesAndSplits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_currentMinter\",\"type\":\"address\"}],\"name\":\"MinterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_field\",\"type\":\"bytes32\"}],\"name\":\"PlatformUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_update\",\"type\":\"bytes32\"}],\"name\":\"ProjectUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_artistAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_additionalPayeePrimarySales\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_additionalPayeePrimarySalesPercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_additionalPayeeSecondarySales\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_additionalPayeeSecondarySalesPercentage\",\"type\":\"uint256\"}],\"name\":\"ProposedArtistAddressesAndSplits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ART_BLOCKS_ERC721TOKEN_ADDRESS_V0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ART_BLOCKS_ERC721TOKEN_ADDRESS_V1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_projectName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_artistAddress\",\"type\":\"address\"}],\"name\":\"addProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_script\",\"type\":\"string\"}],\"name\":\"addProjectScript\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"adminACLAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminACLContract\",\"outputs\":[{\"internalType\":\"contractIAdminACLV0\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_artistAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_additionalPayeePrimarySales\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_additionalPayeePrimarySalesPercentage\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_additionalPayeeSecondarySales\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_additionalPayeeSecondarySalesPercentage\",\"type\":\"uint256\"}],\"name\":\"adminAcceptArtistAddressesAndSplits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artblocksAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artblocksCurationRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artblocksDependencyRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artblocksPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artblocksPrimarySalesAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artblocksPrimarySalesPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artblocksSecondarySalesAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artblocksSecondarySalesBPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coreType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coreVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forbidNewProjects\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getHistoricalRandomizerAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"getPrimaryRevenueSplits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"artblocksRevenue_\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"artblocksAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"artistRevenue_\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"artistAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalPayeePrimaryRevenue_\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"additionalPayeePrimaryAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyalties\",\"outputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"bps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyaltyData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"artistAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"additionalPayee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalPayeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"royaltyFeeByID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"isMintWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_by\",\"type\":\"address\"}],\"name\":\"mint_Ecf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minterContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newProjectsForbidden\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextProjectId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numHistoricalRandomizers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectArtistPaymentInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"artistAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"additionalPayeePrimarySales\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalPayeePrimarySalesPercentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"additionalPayeeSecondarySales\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalPayeeSecondarySalesPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryMarketRoyaltyPercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"artist\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"license\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectIdToAdditionalPayeePrimarySales\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectIdToAdditionalPayeePrimarySalesPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectIdToAdditionalPayeeSecondarySales\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectIdToAdditionalPayeeSecondarySalesPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectIdToArtistAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectIdToSecondaryMarketRoyaltyPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"projectScriptByIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"projectScriptBytecodeAddressByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectScriptDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"scriptTypeAndVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"aspectRatio\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"scriptCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectStateData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"invocations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxInvocations\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"completedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"projectURIInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"projectBaseURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_artistAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_additionalPayeePrimarySales\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_additionalPayeePrimarySalesPercentage\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_additionalPayeeSecondarySales\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_additionalPayeeSecondarySalesPercentage\",\"type\":\"uint256\"}],\"name\":\"proposeArtistPaymentAddressesAndSplits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposedArtistAddressesAndSplitsHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomizerContract\",\"outputs\":[{\"internalType\":\"contractIRandomizerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"removeProjectLastScript\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hashSeed\",\"type\":\"bytes32\"}],\"name\":\"setTokenHash_8PT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingProjectId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"toggleProjectIsPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToProjectId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_artblocksCurationRegistryAddress\",\"type\":\"address\"}],\"name\":\"updateArtblocksCurationRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_artblocksDependencyRegistryAddress\",\"type\":\"address\"}],\"name\":\"updateArtblocksDependencyRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_artblocksPrimarySalesAddress\",\"type\":\"address\"}],\"name\":\"updateArtblocksPrimarySalesAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"artblocksPrimarySalesPercentage_\",\"type\":\"uint256\"}],\"name\":\"updateArtblocksPrimarySalesPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_artblocksSecondarySalesAddress\",\"type\":\"address\"}],\"name\":\"updateArtblocksSecondarySalesAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_artblocksSecondarySalesBPS\",\"type\":\"uint256\"}],\"name\":\"updateArtblocksSecondarySalesBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_defaultBaseURI\",\"type\":\"string\"}],\"name\":\"updateDefaultBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateMinterContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_artistAddress\",\"type\":\"address\"}],\"name\":\"updateProjectArtistAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectArtistName\",\"type\":\"string\"}],\"name\":\"updateProjectArtistName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_aspectRatio\",\"type\":\"string\"}],\"name\":\"updateProjectAspectRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newBaseURI\",\"type\":\"string\"}],\"name\":\"updateProjectBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectDescription\",\"type\":\"string\"}],\"name\":\"updateProjectDescription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectLicense\",\"type\":\"string\"}],\"name\":\"updateProjectLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"_maxInvocations\",\"type\":\"uint24\"}],\"name\":\"updateProjectMaxInvocations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectName\",\"type\":\"string\"}],\"name\":\"updateProjectName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_scriptId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_script\",\"type\":\"string\"}],\"name\":\"updateProjectScript\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_scriptTypeAndVersion\",\"type\":\"bytes32\"}],\"name\":\"updateProjectScriptType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondMarketRoyalty\",\"type\":\"uint256\"}],\"name\":\"updateProjectSecondaryMarketRoyaltyPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectWebsite\",\"type\":\"string\"}],\"name\":\"updateProjectWebsite\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_randomizerAddress\",\"type\":\"address\"}],\"name\":\"updateRandomizerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// ARTBLOCKSERC721TOKENADDRESSV0 is a free data retrieval call binding the contract method 0x14fc8f2d.
//
// Solidity: function ART_BLOCKS_ERC721TOKEN_ADDRESS_V0() view returns(address)
func (_Artblocks *ArtblocksCaller) ARTBLOCKSERC721TOKENADDRESSV0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "ART_BLOCKS_ERC721TOKEN_ADDRESS_V0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ARTBLOCKSERC721TOKENADDRESSV0 is a free data retrieval call binding the contract method 0x14fc8f2d.
//
// Solidity: function ART_BLOCKS_ERC721TOKEN_ADDRESS_V0() view returns(address)
func (_Artblocks *ArtblocksSession) ARTBLOCKSERC721TOKENADDRESSV0() (common.Address, error) {
	return _Artblocks.Contract.ARTBLOCKSERC721TOKENADDRESSV0(&_Artblocks.CallOpts)
}

// ARTBLOCKSERC721TOKENADDRESSV0 is a free data retrieval call binding the contract method 0x14fc8f2d.
//
// Solidity: function ART_BLOCKS_ERC721TOKEN_ADDRESS_V0() view returns(address)
func (_Artblocks *ArtblocksCallerSession) ARTBLOCKSERC721TOKENADDRESSV0() (common.Address, error) {
	return _Artblocks.Contract.ARTBLOCKSERC721TOKENADDRESSV0(&_Artblocks.CallOpts)
}

// ARTBLOCKSERC721TOKENADDRESSV1 is a free data retrieval call binding the contract method 0x9afc2be5.
//
// Solidity: function ART_BLOCKS_ERC721TOKEN_ADDRESS_V1() view returns(address)
func (_Artblocks *ArtblocksCaller) ARTBLOCKSERC721TOKENADDRESSV1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "ART_BLOCKS_ERC721TOKEN_ADDRESS_V1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ARTBLOCKSERC721TOKENADDRESSV1 is a free data retrieval call binding the contract method 0x9afc2be5.
//
// Solidity: function ART_BLOCKS_ERC721TOKEN_ADDRESS_V1() view returns(address)
func (_Artblocks *ArtblocksSession) ARTBLOCKSERC721TOKENADDRESSV1() (common.Address, error) {
	return _Artblocks.Contract.ARTBLOCKSERC721TOKENADDRESSV1(&_Artblocks.CallOpts)
}

// ARTBLOCKSERC721TOKENADDRESSV1 is a free data retrieval call binding the contract method 0x9afc2be5.
//
// Solidity: function ART_BLOCKS_ERC721TOKEN_ADDRESS_V1() view returns(address)
func (_Artblocks *ArtblocksCallerSession) ARTBLOCKSERC721TOKENADDRESSV1() (common.Address, error) {
	return _Artblocks.Contract.ARTBLOCKSERC721TOKENADDRESSV1(&_Artblocks.CallOpts)
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

// AdminACLContract is a free data retrieval call binding the contract method 0x1e9bef46.
//
// Solidity: function adminACLContract() view returns(address)
func (_Artblocks *ArtblocksCaller) AdminACLContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "adminACLContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminACLContract is a free data retrieval call binding the contract method 0x1e9bef46.
//
// Solidity: function adminACLContract() view returns(address)
func (_Artblocks *ArtblocksSession) AdminACLContract() (common.Address, error) {
	return _Artblocks.Contract.AdminACLContract(&_Artblocks.CallOpts)
}

// AdminACLContract is a free data retrieval call binding the contract method 0x1e9bef46.
//
// Solidity: function adminACLContract() view returns(address)
func (_Artblocks *ArtblocksCallerSession) AdminACLContract() (common.Address, error) {
	return _Artblocks.Contract.AdminACLContract(&_Artblocks.CallOpts)
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

// ArtblocksCurationRegistryAddress is a free data retrieval call binding the contract method 0xbba4448a.
//
// Solidity: function artblocksCurationRegistryAddress() view returns(address)
func (_Artblocks *ArtblocksCaller) ArtblocksCurationRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "artblocksCurationRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArtblocksCurationRegistryAddress is a free data retrieval call binding the contract method 0xbba4448a.
//
// Solidity: function artblocksCurationRegistryAddress() view returns(address)
func (_Artblocks *ArtblocksSession) ArtblocksCurationRegistryAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksCurationRegistryAddress(&_Artblocks.CallOpts)
}

// ArtblocksCurationRegistryAddress is a free data retrieval call binding the contract method 0xbba4448a.
//
// Solidity: function artblocksCurationRegistryAddress() view returns(address)
func (_Artblocks *ArtblocksCallerSession) ArtblocksCurationRegistryAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksCurationRegistryAddress(&_Artblocks.CallOpts)
}

// ArtblocksDependencyRegistryAddress is a free data retrieval call binding the contract method 0x17df5366.
//
// Solidity: function artblocksDependencyRegistryAddress() view returns(address)
func (_Artblocks *ArtblocksCaller) ArtblocksDependencyRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "artblocksDependencyRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArtblocksDependencyRegistryAddress is a free data retrieval call binding the contract method 0x17df5366.
//
// Solidity: function artblocksDependencyRegistryAddress() view returns(address)
func (_Artblocks *ArtblocksSession) ArtblocksDependencyRegistryAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksDependencyRegistryAddress(&_Artblocks.CallOpts)
}

// ArtblocksDependencyRegistryAddress is a free data retrieval call binding the contract method 0x17df5366.
//
// Solidity: function artblocksDependencyRegistryAddress() view returns(address)
func (_Artblocks *ArtblocksCallerSession) ArtblocksDependencyRegistryAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksDependencyRegistryAddress(&_Artblocks.CallOpts)
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

// ArtblocksPrimarySalesAddress is a free data retrieval call binding the contract method 0xddd0ee0f.
//
// Solidity: function artblocksPrimarySalesAddress() view returns(address)
func (_Artblocks *ArtblocksCaller) ArtblocksPrimarySalesAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "artblocksPrimarySalesAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArtblocksPrimarySalesAddress is a free data retrieval call binding the contract method 0xddd0ee0f.
//
// Solidity: function artblocksPrimarySalesAddress() view returns(address)
func (_Artblocks *ArtblocksSession) ArtblocksPrimarySalesAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksPrimarySalesAddress(&_Artblocks.CallOpts)
}

// ArtblocksPrimarySalesAddress is a free data retrieval call binding the contract method 0xddd0ee0f.
//
// Solidity: function artblocksPrimarySalesAddress() view returns(address)
func (_Artblocks *ArtblocksCallerSession) ArtblocksPrimarySalesAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksPrimarySalesAddress(&_Artblocks.CallOpts)
}

// ArtblocksPrimarySalesPercentage is a free data retrieval call binding the contract method 0xa87ac619.
//
// Solidity: function artblocksPrimarySalesPercentage() view returns(uint256)
func (_Artblocks *ArtblocksCaller) ArtblocksPrimarySalesPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "artblocksPrimarySalesPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArtblocksPrimarySalesPercentage is a free data retrieval call binding the contract method 0xa87ac619.
//
// Solidity: function artblocksPrimarySalesPercentage() view returns(uint256)
func (_Artblocks *ArtblocksSession) ArtblocksPrimarySalesPercentage() (*big.Int, error) {
	return _Artblocks.Contract.ArtblocksPrimarySalesPercentage(&_Artblocks.CallOpts)
}

// ArtblocksPrimarySalesPercentage is a free data retrieval call binding the contract method 0xa87ac619.
//
// Solidity: function artblocksPrimarySalesPercentage() view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) ArtblocksPrimarySalesPercentage() (*big.Int, error) {
	return _Artblocks.Contract.ArtblocksPrimarySalesPercentage(&_Artblocks.CallOpts)
}

// ArtblocksSecondarySalesAddress is a free data retrieval call binding the contract method 0x94535b99.
//
// Solidity: function artblocksSecondarySalesAddress() view returns(address)
func (_Artblocks *ArtblocksCaller) ArtblocksSecondarySalesAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "artblocksSecondarySalesAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArtblocksSecondarySalesAddress is a free data retrieval call binding the contract method 0x94535b99.
//
// Solidity: function artblocksSecondarySalesAddress() view returns(address)
func (_Artblocks *ArtblocksSession) ArtblocksSecondarySalesAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksSecondarySalesAddress(&_Artblocks.CallOpts)
}

// ArtblocksSecondarySalesAddress is a free data retrieval call binding the contract method 0x94535b99.
//
// Solidity: function artblocksSecondarySalesAddress() view returns(address)
func (_Artblocks *ArtblocksCallerSession) ArtblocksSecondarySalesAddress() (common.Address, error) {
	return _Artblocks.Contract.ArtblocksSecondarySalesAddress(&_Artblocks.CallOpts)
}

// ArtblocksSecondarySalesBPS is a free data retrieval call binding the contract method 0xaccd17f0.
//
// Solidity: function artblocksSecondarySalesBPS() view returns(uint256)
func (_Artblocks *ArtblocksCaller) ArtblocksSecondarySalesBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "artblocksSecondarySalesBPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArtblocksSecondarySalesBPS is a free data retrieval call binding the contract method 0xaccd17f0.
//
// Solidity: function artblocksSecondarySalesBPS() view returns(uint256)
func (_Artblocks *ArtblocksSession) ArtblocksSecondarySalesBPS() (*big.Int, error) {
	return _Artblocks.Contract.ArtblocksSecondarySalesBPS(&_Artblocks.CallOpts)
}

// ArtblocksSecondarySalesBPS is a free data retrieval call binding the contract method 0xaccd17f0.
//
// Solidity: function artblocksSecondarySalesBPS() view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) ArtblocksSecondarySalesBPS() (*big.Int, error) {
	return _Artblocks.Contract.ArtblocksSecondarySalesBPS(&_Artblocks.CallOpts)
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

// CoreType is a free data retrieval call binding the contract method 0xae45ad98.
//
// Solidity: function coreType() view returns(string)
func (_Artblocks *ArtblocksCaller) CoreType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "coreType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CoreType is a free data retrieval call binding the contract method 0xae45ad98.
//
// Solidity: function coreType() view returns(string)
func (_Artblocks *ArtblocksSession) CoreType() (string, error) {
	return _Artblocks.Contract.CoreType(&_Artblocks.CallOpts)
}

// CoreType is a free data retrieval call binding the contract method 0xae45ad98.
//
// Solidity: function coreType() view returns(string)
func (_Artblocks *ArtblocksCallerSession) CoreType() (string, error) {
	return _Artblocks.Contract.CoreType(&_Artblocks.CallOpts)
}

// CoreVersion is a free data retrieval call binding the contract method 0x4e1d64af.
//
// Solidity: function coreVersion() view returns(string)
func (_Artblocks *ArtblocksCaller) CoreVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "coreVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CoreVersion is a free data retrieval call binding the contract method 0x4e1d64af.
//
// Solidity: function coreVersion() view returns(string)
func (_Artblocks *ArtblocksSession) CoreVersion() (string, error) {
	return _Artblocks.Contract.CoreVersion(&_Artblocks.CallOpts)
}

// CoreVersion is a free data retrieval call binding the contract method 0x4e1d64af.
//
// Solidity: function coreVersion() view returns(string)
func (_Artblocks *ArtblocksCallerSession) CoreVersion() (string, error) {
	return _Artblocks.Contract.CoreVersion(&_Artblocks.CallOpts)
}

// DefaultBaseURI is a free data retrieval call binding the contract method 0xabcbb7b4.
//
// Solidity: function defaultBaseURI() view returns(string)
func (_Artblocks *ArtblocksCaller) DefaultBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "defaultBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DefaultBaseURI is a free data retrieval call binding the contract method 0xabcbb7b4.
//
// Solidity: function defaultBaseURI() view returns(string)
func (_Artblocks *ArtblocksSession) DefaultBaseURI() (string, error) {
	return _Artblocks.Contract.DefaultBaseURI(&_Artblocks.CallOpts)
}

// DefaultBaseURI is a free data retrieval call binding the contract method 0xabcbb7b4.
//
// Solidity: function defaultBaseURI() view returns(string)
func (_Artblocks *ArtblocksCallerSession) DefaultBaseURI() (string, error) {
	return _Artblocks.Contract.DefaultBaseURI(&_Artblocks.CallOpts)
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

// GetHistoricalRandomizerAt is a free data retrieval call binding the contract method 0xb1687622.
//
// Solidity: function getHistoricalRandomizerAt(uint256 _index) view returns(address)
func (_Artblocks *ArtblocksCaller) GetHistoricalRandomizerAt(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "getHistoricalRandomizerAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHistoricalRandomizerAt is a free data retrieval call binding the contract method 0xb1687622.
//
// Solidity: function getHistoricalRandomizerAt(uint256 _index) view returns(address)
func (_Artblocks *ArtblocksSession) GetHistoricalRandomizerAt(_index *big.Int) (common.Address, error) {
	return _Artblocks.Contract.GetHistoricalRandomizerAt(&_Artblocks.CallOpts, _index)
}

// GetHistoricalRandomizerAt is a free data retrieval call binding the contract method 0xb1687622.
//
// Solidity: function getHistoricalRandomizerAt(uint256 _index) view returns(address)
func (_Artblocks *ArtblocksCallerSession) GetHistoricalRandomizerAt(_index *big.Int) (common.Address, error) {
	return _Artblocks.Contract.GetHistoricalRandomizerAt(&_Artblocks.CallOpts, _index)
}

// GetPrimaryRevenueSplits is a free data retrieval call binding the contract method 0x8639415b.
//
// Solidity: function getPrimaryRevenueSplits(uint256 _projectId, uint256 _price) view returns(uint256 artblocksRevenue_, address artblocksAddress_, uint256 artistRevenue_, address artistAddress_, uint256 additionalPayeePrimaryRevenue_, address additionalPayeePrimaryAddress_)
func (_Artblocks *ArtblocksCaller) GetPrimaryRevenueSplits(opts *bind.CallOpts, _projectId *big.Int, _price *big.Int) (struct {
	ArtblocksRevenue              *big.Int
	ArtblocksAddress              common.Address
	ArtistRevenue                 *big.Int
	ArtistAddress                 common.Address
	AdditionalPayeePrimaryRevenue *big.Int
	AdditionalPayeePrimaryAddress common.Address
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "getPrimaryRevenueSplits", _projectId, _price)

	outstruct := new(struct {
		ArtblocksRevenue              *big.Int
		ArtblocksAddress              common.Address
		ArtistRevenue                 *big.Int
		ArtistAddress                 common.Address
		AdditionalPayeePrimaryRevenue *big.Int
		AdditionalPayeePrimaryAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ArtblocksRevenue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ArtblocksAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ArtistRevenue = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ArtistAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.AdditionalPayeePrimaryRevenue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AdditionalPayeePrimaryAddress = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetPrimaryRevenueSplits is a free data retrieval call binding the contract method 0x8639415b.
//
// Solidity: function getPrimaryRevenueSplits(uint256 _projectId, uint256 _price) view returns(uint256 artblocksRevenue_, address artblocksAddress_, uint256 artistRevenue_, address artistAddress_, uint256 additionalPayeePrimaryRevenue_, address additionalPayeePrimaryAddress_)
func (_Artblocks *ArtblocksSession) GetPrimaryRevenueSplits(_projectId *big.Int, _price *big.Int) (struct {
	ArtblocksRevenue              *big.Int
	ArtblocksAddress              common.Address
	ArtistRevenue                 *big.Int
	ArtistAddress                 common.Address
	AdditionalPayeePrimaryRevenue *big.Int
	AdditionalPayeePrimaryAddress common.Address
}, error) {
	return _Artblocks.Contract.GetPrimaryRevenueSplits(&_Artblocks.CallOpts, _projectId, _price)
}

// GetPrimaryRevenueSplits is a free data retrieval call binding the contract method 0x8639415b.
//
// Solidity: function getPrimaryRevenueSplits(uint256 _projectId, uint256 _price) view returns(uint256 artblocksRevenue_, address artblocksAddress_, uint256 artistRevenue_, address artistAddress_, uint256 additionalPayeePrimaryRevenue_, address additionalPayeePrimaryAddress_)
func (_Artblocks *ArtblocksCallerSession) GetPrimaryRevenueSplits(_projectId *big.Int, _price *big.Int) (struct {
	ArtblocksRevenue              *big.Int
	ArtblocksAddress              common.Address
	ArtistRevenue                 *big.Int
	ArtistAddress                 common.Address
	AdditionalPayeePrimaryRevenue *big.Int
	AdditionalPayeePrimaryAddress common.Address
}, error) {
	return _Artblocks.Contract.GetPrimaryRevenueSplits(&_Artblocks.CallOpts, _projectId, _price)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 _tokenId) view returns(address[] recipients, uint256[] bps)
func (_Artblocks *ArtblocksCaller) GetRoyalties(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Recipients []common.Address
	Bps        []*big.Int
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "getRoyalties", _tokenId)

	outstruct := new(struct {
		Recipients []common.Address
		Bps        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipients = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Bps = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 _tokenId) view returns(address[] recipients, uint256[] bps)
func (_Artblocks *ArtblocksSession) GetRoyalties(_tokenId *big.Int) (struct {
	Recipients []common.Address
	Bps        []*big.Int
}, error) {
	return _Artblocks.Contract.GetRoyalties(&_Artblocks.CallOpts, _tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 _tokenId) view returns(address[] recipients, uint256[] bps)
func (_Artblocks *ArtblocksCallerSession) GetRoyalties(_tokenId *big.Int) (struct {
	Recipients []common.Address
	Bps        []*big.Int
}, error) {
	return _Artblocks.Contract.GetRoyalties(&_Artblocks.CallOpts, _tokenId)
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
// Solidity: function isMintWhitelisted(address _minter) view returns(bool)
func (_Artblocks *ArtblocksCaller) IsMintWhitelisted(opts *bind.CallOpts, _minter common.Address) (bool, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "isMintWhitelisted", _minter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMintWhitelisted is a free data retrieval call binding the contract method 0xad0305ce.
//
// Solidity: function isMintWhitelisted(address _minter) view returns(bool)
func (_Artblocks *ArtblocksSession) IsMintWhitelisted(_minter common.Address) (bool, error) {
	return _Artblocks.Contract.IsMintWhitelisted(&_Artblocks.CallOpts, _minter)
}

// IsMintWhitelisted is a free data retrieval call binding the contract method 0xad0305ce.
//
// Solidity: function isMintWhitelisted(address _minter) view returns(bool)
func (_Artblocks *ArtblocksCallerSession) IsMintWhitelisted(_minter common.Address) (bool, error) {
	return _Artblocks.Contract.IsMintWhitelisted(&_Artblocks.CallOpts, _minter)
}

// MinterContract is a free data retrieval call binding the contract method 0x92f00233.
//
// Solidity: function minterContract() view returns(address)
func (_Artblocks *ArtblocksCaller) MinterContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "minterContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MinterContract is a free data retrieval call binding the contract method 0x92f00233.
//
// Solidity: function minterContract() view returns(address)
func (_Artblocks *ArtblocksSession) MinterContract() (common.Address, error) {
	return _Artblocks.Contract.MinterContract(&_Artblocks.CallOpts)
}

// MinterContract is a free data retrieval call binding the contract method 0x92f00233.
//
// Solidity: function minterContract() view returns(address)
func (_Artblocks *ArtblocksCallerSession) MinterContract() (common.Address, error) {
	return _Artblocks.Contract.MinterContract(&_Artblocks.CallOpts)
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

// NewProjectsForbidden is a free data retrieval call binding the contract method 0x5508fd52.
//
// Solidity: function newProjectsForbidden() view returns(bool)
func (_Artblocks *ArtblocksCaller) NewProjectsForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "newProjectsForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NewProjectsForbidden is a free data retrieval call binding the contract method 0x5508fd52.
//
// Solidity: function newProjectsForbidden() view returns(bool)
func (_Artblocks *ArtblocksSession) NewProjectsForbidden() (bool, error) {
	return _Artblocks.Contract.NewProjectsForbidden(&_Artblocks.CallOpts)
}

// NewProjectsForbidden is a free data retrieval call binding the contract method 0x5508fd52.
//
// Solidity: function newProjectsForbidden() view returns(bool)
func (_Artblocks *ArtblocksCallerSession) NewProjectsForbidden() (bool, error) {
	return _Artblocks.Contract.NewProjectsForbidden(&_Artblocks.CallOpts)
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

// NumHistoricalRandomizers is a free data retrieval call binding the contract method 0xb9711368.
//
// Solidity: function numHistoricalRandomizers() view returns(uint256)
func (_Artblocks *ArtblocksCaller) NumHistoricalRandomizers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "numHistoricalRandomizers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumHistoricalRandomizers is a free data retrieval call binding the contract method 0xb9711368.
//
// Solidity: function numHistoricalRandomizers() view returns(uint256)
func (_Artblocks *ArtblocksSession) NumHistoricalRandomizers() (*big.Int, error) {
	return _Artblocks.Contract.NumHistoricalRandomizers(&_Artblocks.CallOpts)
}

// NumHistoricalRandomizers is a free data retrieval call binding the contract method 0xb9711368.
//
// Solidity: function numHistoricalRandomizers() view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) NumHistoricalRandomizers() (*big.Int, error) {
	return _Artblocks.Contract.NumHistoricalRandomizers(&_Artblocks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Artblocks *ArtblocksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Artblocks *ArtblocksSession) Owner() (common.Address, error) {
	return _Artblocks.Contract.Owner(&_Artblocks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Artblocks *ArtblocksCallerSession) Owner() (common.Address, error) {
	return _Artblocks.Contract.Owner(&_Artblocks.CallOpts)
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

// ProjectArtistPaymentInfo is a free data retrieval call binding the contract method 0xf313d733.
//
// Solidity: function projectArtistPaymentInfo(uint256 _projectId) view returns(address artistAddress, address additionalPayeePrimarySales, uint256 additionalPayeePrimarySalesPercentage, address additionalPayeeSecondarySales, uint256 additionalPayeeSecondarySalesPercentage, uint256 secondaryMarketRoyaltyPercentage)
func (_Artblocks *ArtblocksCaller) ProjectArtistPaymentInfo(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ArtistAddress                           common.Address
	AdditionalPayeePrimarySales             common.Address
	AdditionalPayeePrimarySalesPercentage   *big.Int
	AdditionalPayeeSecondarySales           common.Address
	AdditionalPayeeSecondarySalesPercentage *big.Int
	SecondaryMarketRoyaltyPercentage        *big.Int
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectArtistPaymentInfo", _projectId)

	outstruct := new(struct {
		ArtistAddress                           common.Address
		AdditionalPayeePrimarySales             common.Address
		AdditionalPayeePrimarySalesPercentage   *big.Int
		AdditionalPayeeSecondarySales           common.Address
		AdditionalPayeeSecondarySalesPercentage *big.Int
		SecondaryMarketRoyaltyPercentage        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ArtistAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AdditionalPayeePrimarySales = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AdditionalPayeePrimarySalesPercentage = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AdditionalPayeeSecondarySales = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.AdditionalPayeeSecondarySalesPercentage = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.SecondaryMarketRoyaltyPercentage = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProjectArtistPaymentInfo is a free data retrieval call binding the contract method 0xf313d733.
//
// Solidity: function projectArtistPaymentInfo(uint256 _projectId) view returns(address artistAddress, address additionalPayeePrimarySales, uint256 additionalPayeePrimarySalesPercentage, address additionalPayeeSecondarySales, uint256 additionalPayeeSecondarySalesPercentage, uint256 secondaryMarketRoyaltyPercentage)
func (_Artblocks *ArtblocksSession) ProjectArtistPaymentInfo(_projectId *big.Int) (struct {
	ArtistAddress                           common.Address
	AdditionalPayeePrimarySales             common.Address
	AdditionalPayeePrimarySalesPercentage   *big.Int
	AdditionalPayeeSecondarySales           common.Address
	AdditionalPayeeSecondarySalesPercentage *big.Int
	SecondaryMarketRoyaltyPercentage        *big.Int
}, error) {
	return _Artblocks.Contract.ProjectArtistPaymentInfo(&_Artblocks.CallOpts, _projectId)
}

// ProjectArtistPaymentInfo is a free data retrieval call binding the contract method 0xf313d733.
//
// Solidity: function projectArtistPaymentInfo(uint256 _projectId) view returns(address artistAddress, address additionalPayeePrimarySales, uint256 additionalPayeePrimarySalesPercentage, address additionalPayeeSecondarySales, uint256 additionalPayeeSecondarySalesPercentage, uint256 secondaryMarketRoyaltyPercentage)
func (_Artblocks *ArtblocksCallerSession) ProjectArtistPaymentInfo(_projectId *big.Int) (struct {
	ArtistAddress                           common.Address
	AdditionalPayeePrimarySales             common.Address
	AdditionalPayeePrimarySalesPercentage   *big.Int
	AdditionalPayeeSecondarySales           common.Address
	AdditionalPayeeSecondarySalesPercentage *big.Int
	SecondaryMarketRoyaltyPercentage        *big.Int
}, error) {
	return _Artblocks.Contract.ProjectArtistPaymentInfo(&_Artblocks.CallOpts, _projectId)
}

// ProjectDetails is a free data retrieval call binding the contract method 0x8dd91a56.
//
// Solidity: function projectDetails(uint256 _projectId) view returns(string projectName, string artist, string description, string website, string license)
func (_Artblocks *ArtblocksCaller) ProjectDetails(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ProjectName string
	Artist      string
	Description string
	Website     string
	License     string
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectDetails", _projectId)

	outstruct := new(struct {
		ProjectName string
		Artist      string
		Description string
		Website     string
		License     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProjectName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Artist = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Website = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.License = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// ProjectDetails is a free data retrieval call binding the contract method 0x8dd91a56.
//
// Solidity: function projectDetails(uint256 _projectId) view returns(string projectName, string artist, string description, string website, string license)
func (_Artblocks *ArtblocksSession) ProjectDetails(_projectId *big.Int) (struct {
	ProjectName string
	Artist      string
	Description string
	Website     string
	License     string
}, error) {
	return _Artblocks.Contract.ProjectDetails(&_Artblocks.CallOpts, _projectId)
}

// ProjectDetails is a free data retrieval call binding the contract method 0x8dd91a56.
//
// Solidity: function projectDetails(uint256 _projectId) view returns(string projectName, string artist, string description, string website, string license)
func (_Artblocks *ArtblocksCallerSession) ProjectDetails(_projectId *big.Int) (struct {
	ProjectName string
	Artist      string
	Description string
	Website     string
	License     string
}, error) {
	return _Artblocks.Contract.ProjectDetails(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToAdditionalPayeePrimarySales is a free data retrieval call binding the contract method 0x0ebeb0ee.
//
// Solidity: function projectIdToAdditionalPayeePrimarySales(uint256 _projectId) view returns(address)
func (_Artblocks *ArtblocksCaller) ProjectIdToAdditionalPayeePrimarySales(opts *bind.CallOpts, _projectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToAdditionalPayeePrimarySales", _projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectIdToAdditionalPayeePrimarySales is a free data retrieval call binding the contract method 0x0ebeb0ee.
//
// Solidity: function projectIdToAdditionalPayeePrimarySales(uint256 _projectId) view returns(address)
func (_Artblocks *ArtblocksSession) ProjectIdToAdditionalPayeePrimarySales(_projectId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeePrimarySales(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToAdditionalPayeePrimarySales is a free data retrieval call binding the contract method 0x0ebeb0ee.
//
// Solidity: function projectIdToAdditionalPayeePrimarySales(uint256 _projectId) view returns(address)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToAdditionalPayeePrimarySales(_projectId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeePrimarySales(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToAdditionalPayeePrimarySalesPercentage is a free data retrieval call binding the contract method 0xad2cdfc4.
//
// Solidity: function projectIdToAdditionalPayeePrimarySalesPercentage(uint256 _projectId) view returns(uint256)
func (_Artblocks *ArtblocksCaller) ProjectIdToAdditionalPayeePrimarySalesPercentage(opts *bind.CallOpts, _projectId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToAdditionalPayeePrimarySalesPercentage", _projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectIdToAdditionalPayeePrimarySalesPercentage is a free data retrieval call binding the contract method 0xad2cdfc4.
//
// Solidity: function projectIdToAdditionalPayeePrimarySalesPercentage(uint256 _projectId) view returns(uint256)
func (_Artblocks *ArtblocksSession) ProjectIdToAdditionalPayeePrimarySalesPercentage(_projectId *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeePrimarySalesPercentage(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToAdditionalPayeePrimarySalesPercentage is a free data retrieval call binding the contract method 0xad2cdfc4.
//
// Solidity: function projectIdToAdditionalPayeePrimarySalesPercentage(uint256 _projectId) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToAdditionalPayeePrimarySalesPercentage(_projectId *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeePrimarySalesPercentage(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToAdditionalPayeeSecondarySales is a free data retrieval call binding the contract method 0x38c8e468.
//
// Solidity: function projectIdToAdditionalPayeeSecondarySales(uint256 _projectId) view returns(address)
func (_Artblocks *ArtblocksCaller) ProjectIdToAdditionalPayeeSecondarySales(opts *bind.CallOpts, _projectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToAdditionalPayeeSecondarySales", _projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectIdToAdditionalPayeeSecondarySales is a free data retrieval call binding the contract method 0x38c8e468.
//
// Solidity: function projectIdToAdditionalPayeeSecondarySales(uint256 _projectId) view returns(address)
func (_Artblocks *ArtblocksSession) ProjectIdToAdditionalPayeeSecondarySales(_projectId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeeSecondarySales(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToAdditionalPayeeSecondarySales is a free data retrieval call binding the contract method 0x38c8e468.
//
// Solidity: function projectIdToAdditionalPayeeSecondarySales(uint256 _projectId) view returns(address)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToAdditionalPayeeSecondarySales(_projectId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeeSecondarySales(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToAdditionalPayeeSecondarySalesPercentage is a free data retrieval call binding the contract method 0x28ec8f50.
//
// Solidity: function projectIdToAdditionalPayeeSecondarySalesPercentage(uint256 _projectId) view returns(uint256)
func (_Artblocks *ArtblocksCaller) ProjectIdToAdditionalPayeeSecondarySalesPercentage(opts *bind.CallOpts, _projectId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToAdditionalPayeeSecondarySalesPercentage", _projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectIdToAdditionalPayeeSecondarySalesPercentage is a free data retrieval call binding the contract method 0x28ec8f50.
//
// Solidity: function projectIdToAdditionalPayeeSecondarySalesPercentage(uint256 _projectId) view returns(uint256)
func (_Artblocks *ArtblocksSession) ProjectIdToAdditionalPayeeSecondarySalesPercentage(_projectId *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeeSecondarySalesPercentage(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToAdditionalPayeeSecondarySalesPercentage is a free data retrieval call binding the contract method 0x28ec8f50.
//
// Solidity: function projectIdToAdditionalPayeeSecondarySalesPercentage(uint256 _projectId) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToAdditionalPayeeSecondarySalesPercentage(_projectId *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToAdditionalPayeeSecondarySalesPercentage(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToArtistAddress is a free data retrieval call binding the contract method 0xa47d29cb.
//
// Solidity: function projectIdToArtistAddress(uint256 _projectId) view returns(address)
func (_Artblocks *ArtblocksCaller) ProjectIdToArtistAddress(opts *bind.CallOpts, _projectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToArtistAddress", _projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectIdToArtistAddress is a free data retrieval call binding the contract method 0xa47d29cb.
//
// Solidity: function projectIdToArtistAddress(uint256 _projectId) view returns(address)
func (_Artblocks *ArtblocksSession) ProjectIdToArtistAddress(_projectId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToArtistAddress(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToArtistAddress is a free data retrieval call binding the contract method 0xa47d29cb.
//
// Solidity: function projectIdToArtistAddress(uint256 _projectId) view returns(address)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToArtistAddress(_projectId *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectIdToArtistAddress(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToSecondaryMarketRoyaltyPercentage is a free data retrieval call binding the contract method 0xed8abfda.
//
// Solidity: function projectIdToSecondaryMarketRoyaltyPercentage(uint256 _projectId) view returns(uint256)
func (_Artblocks *ArtblocksCaller) ProjectIdToSecondaryMarketRoyaltyPercentage(opts *bind.CallOpts, _projectId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectIdToSecondaryMarketRoyaltyPercentage", _projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectIdToSecondaryMarketRoyaltyPercentage is a free data retrieval call binding the contract method 0xed8abfda.
//
// Solidity: function projectIdToSecondaryMarketRoyaltyPercentage(uint256 _projectId) view returns(uint256)
func (_Artblocks *ArtblocksSession) ProjectIdToSecondaryMarketRoyaltyPercentage(_projectId *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToSecondaryMarketRoyaltyPercentage(&_Artblocks.CallOpts, _projectId)
}

// ProjectIdToSecondaryMarketRoyaltyPercentage is a free data retrieval call binding the contract method 0xed8abfda.
//
// Solidity: function projectIdToSecondaryMarketRoyaltyPercentage(uint256 _projectId) view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) ProjectIdToSecondaryMarketRoyaltyPercentage(_projectId *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.ProjectIdToSecondaryMarketRoyaltyPercentage(&_Artblocks.CallOpts, _projectId)
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

// ProjectScriptBytecodeAddressByIndex is a free data retrieval call binding the contract method 0x58b9a5a9.
//
// Solidity: function projectScriptBytecodeAddressByIndex(uint256 _projectId, uint256 _index) view returns(address)
func (_Artblocks *ArtblocksCaller) ProjectScriptBytecodeAddressByIndex(opts *bind.CallOpts, _projectId *big.Int, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectScriptBytecodeAddressByIndex", _projectId, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectScriptBytecodeAddressByIndex is a free data retrieval call binding the contract method 0x58b9a5a9.
//
// Solidity: function projectScriptBytecodeAddressByIndex(uint256 _projectId, uint256 _index) view returns(address)
func (_Artblocks *ArtblocksSession) ProjectScriptBytecodeAddressByIndex(_projectId *big.Int, _index *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectScriptBytecodeAddressByIndex(&_Artblocks.CallOpts, _projectId, _index)
}

// ProjectScriptBytecodeAddressByIndex is a free data retrieval call binding the contract method 0x58b9a5a9.
//
// Solidity: function projectScriptBytecodeAddressByIndex(uint256 _projectId, uint256 _index) view returns(address)
func (_Artblocks *ArtblocksCallerSession) ProjectScriptBytecodeAddressByIndex(_projectId *big.Int, _index *big.Int) (common.Address, error) {
	return _Artblocks.Contract.ProjectScriptBytecodeAddressByIndex(&_Artblocks.CallOpts, _projectId, _index)
}

// ProjectScriptDetails is a free data retrieval call binding the contract method 0xeb9cd5d4.
//
// Solidity: function projectScriptDetails(uint256 _projectId) view returns(string scriptTypeAndVersion, string aspectRatio, uint256 scriptCount)
func (_Artblocks *ArtblocksCaller) ProjectScriptDetails(opts *bind.CallOpts, _projectId *big.Int) (struct {
	ScriptTypeAndVersion string
	AspectRatio          string
	ScriptCount          *big.Int
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectScriptDetails", _projectId)

	outstruct := new(struct {
		ScriptTypeAndVersion string
		AspectRatio          string
		ScriptCount          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScriptTypeAndVersion = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AspectRatio = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ScriptCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProjectScriptDetails is a free data retrieval call binding the contract method 0xeb9cd5d4.
//
// Solidity: function projectScriptDetails(uint256 _projectId) view returns(string scriptTypeAndVersion, string aspectRatio, uint256 scriptCount)
func (_Artblocks *ArtblocksSession) ProjectScriptDetails(_projectId *big.Int) (struct {
	ScriptTypeAndVersion string
	AspectRatio          string
	ScriptCount          *big.Int
}, error) {
	return _Artblocks.Contract.ProjectScriptDetails(&_Artblocks.CallOpts, _projectId)
}

// ProjectScriptDetails is a free data retrieval call binding the contract method 0xeb9cd5d4.
//
// Solidity: function projectScriptDetails(uint256 _projectId) view returns(string scriptTypeAndVersion, string aspectRatio, uint256 scriptCount)
func (_Artblocks *ArtblocksCallerSession) ProjectScriptDetails(_projectId *big.Int) (struct {
	ScriptTypeAndVersion string
	AspectRatio          string
	ScriptCount          *big.Int
}, error) {
	return _Artblocks.Contract.ProjectScriptDetails(&_Artblocks.CallOpts, _projectId)
}

// ProjectStateData is a free data retrieval call binding the contract method 0x0ea5613f.
//
// Solidity: function projectStateData(uint256 _projectId) view returns(uint256 invocations, uint256 maxInvocations, bool active, bool paused, uint256 completedTimestamp, bool locked)
func (_Artblocks *ArtblocksCaller) ProjectStateData(opts *bind.CallOpts, _projectId *big.Int) (struct {
	Invocations        *big.Int
	MaxInvocations     *big.Int
	Active             bool
	Paused             bool
	CompletedTimestamp *big.Int
	Locked             bool
}, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectStateData", _projectId)

	outstruct := new(struct {
		Invocations        *big.Int
		MaxInvocations     *big.Int
		Active             bool
		Paused             bool
		CompletedTimestamp *big.Int
		Locked             bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Invocations = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxInvocations = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Paused = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.CompletedTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Locked = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// ProjectStateData is a free data retrieval call binding the contract method 0x0ea5613f.
//
// Solidity: function projectStateData(uint256 _projectId) view returns(uint256 invocations, uint256 maxInvocations, bool active, bool paused, uint256 completedTimestamp, bool locked)
func (_Artblocks *ArtblocksSession) ProjectStateData(_projectId *big.Int) (struct {
	Invocations        *big.Int
	MaxInvocations     *big.Int
	Active             bool
	Paused             bool
	CompletedTimestamp *big.Int
	Locked             bool
}, error) {
	return _Artblocks.Contract.ProjectStateData(&_Artblocks.CallOpts, _projectId)
}

// ProjectStateData is a free data retrieval call binding the contract method 0x0ea5613f.
//
// Solidity: function projectStateData(uint256 _projectId) view returns(uint256 invocations, uint256 maxInvocations, bool active, bool paused, uint256 completedTimestamp, bool locked)
func (_Artblocks *ArtblocksCallerSession) ProjectStateData(_projectId *big.Int) (struct {
	Invocations        *big.Int
	MaxInvocations     *big.Int
	Active             bool
	Paused             bool
	CompletedTimestamp *big.Int
	Locked             bool
}, error) {
	return _Artblocks.Contract.ProjectStateData(&_Artblocks.CallOpts, _projectId)
}

// ProjectURIInfo is a free data retrieval call binding the contract method 0x2d9c0205.
//
// Solidity: function projectURIInfo(uint256 _projectId) view returns(string projectBaseURI)
func (_Artblocks *ArtblocksCaller) ProjectURIInfo(opts *bind.CallOpts, _projectId *big.Int) (string, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "projectURIInfo", _projectId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProjectURIInfo is a free data retrieval call binding the contract method 0x2d9c0205.
//
// Solidity: function projectURIInfo(uint256 _projectId) view returns(string projectBaseURI)
func (_Artblocks *ArtblocksSession) ProjectURIInfo(_projectId *big.Int) (string, error) {
	return _Artblocks.Contract.ProjectURIInfo(&_Artblocks.CallOpts, _projectId)
}

// ProjectURIInfo is a free data retrieval call binding the contract method 0x2d9c0205.
//
// Solidity: function projectURIInfo(uint256 _projectId) view returns(string projectBaseURI)
func (_Artblocks *ArtblocksCallerSession) ProjectURIInfo(_projectId *big.Int) (string, error) {
	return _Artblocks.Contract.ProjectURIInfo(&_Artblocks.CallOpts, _projectId)
}

// ProposedArtistAddressesAndSplitsHash is a free data retrieval call binding the contract method 0xac11fa1c.
//
// Solidity: function proposedArtistAddressesAndSplitsHash(uint256 ) view returns(bytes32)
func (_Artblocks *ArtblocksCaller) ProposedArtistAddressesAndSplitsHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "proposedArtistAddressesAndSplitsHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProposedArtistAddressesAndSplitsHash is a free data retrieval call binding the contract method 0xac11fa1c.
//
// Solidity: function proposedArtistAddressesAndSplitsHash(uint256 ) view returns(bytes32)
func (_Artblocks *ArtblocksSession) ProposedArtistAddressesAndSplitsHash(arg0 *big.Int) ([32]byte, error) {
	return _Artblocks.Contract.ProposedArtistAddressesAndSplitsHash(&_Artblocks.CallOpts, arg0)
}

// ProposedArtistAddressesAndSplitsHash is a free data retrieval call binding the contract method 0xac11fa1c.
//
// Solidity: function proposedArtistAddressesAndSplitsHash(uint256 ) view returns(bytes32)
func (_Artblocks *ArtblocksCallerSession) ProposedArtistAddressesAndSplitsHash(arg0 *big.Int) ([32]byte, error) {
	return _Artblocks.Contract.ProposedArtistAddressesAndSplitsHash(&_Artblocks.CallOpts, arg0)
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

// StartingProjectId is a free data retrieval call binding the contract method 0xf893c07b.
//
// Solidity: function startingProjectId() view returns(uint256)
func (_Artblocks *ArtblocksCaller) StartingProjectId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "startingProjectId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingProjectId is a free data retrieval call binding the contract method 0xf893c07b.
//
// Solidity: function startingProjectId() view returns(uint256)
func (_Artblocks *ArtblocksSession) StartingProjectId() (*big.Int, error) {
	return _Artblocks.Contract.StartingProjectId(&_Artblocks.CallOpts)
}

// StartingProjectId is a free data retrieval call binding the contract method 0xf893c07b.
//
// Solidity: function startingProjectId() view returns(uint256)
func (_Artblocks *ArtblocksCallerSession) StartingProjectId() (*big.Int, error) {
	return _Artblocks.Contract.StartingProjectId(&_Artblocks.CallOpts)
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

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 _tokenId) view returns(bytes32)
func (_Artblocks *ArtblocksCaller) TokenIdToHash(opts *bind.CallOpts, _tokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "tokenIdToHash", _tokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 _tokenId) view returns(bytes32)
func (_Artblocks *ArtblocksSession) TokenIdToHash(_tokenId *big.Int) ([32]byte, error) {
	return _Artblocks.Contract.TokenIdToHash(&_Artblocks.CallOpts, _tokenId)
}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 _tokenId) view returns(bytes32)
func (_Artblocks *ArtblocksCallerSession) TokenIdToHash(_tokenId *big.Int) ([32]byte, error) {
	return _Artblocks.Contract.TokenIdToHash(&_Artblocks.CallOpts, _tokenId)
}

// TokenIdToProjectId is a free data retrieval call binding the contract method 0x1b689c0b.
//
// Solidity: function tokenIdToProjectId(uint256 _tokenId) pure returns(uint256 _projectId)
func (_Artblocks *ArtblocksCaller) TokenIdToProjectId(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artblocks.contract.Call(opts, &out, "tokenIdToProjectId", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToProjectId is a free data retrieval call binding the contract method 0x1b689c0b.
//
// Solidity: function tokenIdToProjectId(uint256 _tokenId) pure returns(uint256 _projectId)
func (_Artblocks *ArtblocksSession) TokenIdToProjectId(_tokenId *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.TokenIdToProjectId(&_Artblocks.CallOpts, _tokenId)
}

// TokenIdToProjectId is a free data retrieval call binding the contract method 0x1b689c0b.
//
// Solidity: function tokenIdToProjectId(uint256 _tokenId) pure returns(uint256 _projectId)
func (_Artblocks *ArtblocksCallerSession) TokenIdToProjectId(_tokenId *big.Int) (*big.Int, error) {
	return _Artblocks.Contract.TokenIdToProjectId(&_Artblocks.CallOpts, _tokenId)
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

// AddProject is a paid mutator transaction binding the contract method 0xcc90e725.
//
// Solidity: function addProject(string _projectName, address _artistAddress) returns()
func (_Artblocks *ArtblocksTransactor) AddProject(opts *bind.TransactOpts, _projectName string, _artistAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "addProject", _projectName, _artistAddress)
}

// AddProject is a paid mutator transaction binding the contract method 0xcc90e725.
//
// Solidity: function addProject(string _projectName, address _artistAddress) returns()
func (_Artblocks *ArtblocksSession) AddProject(_projectName string, _artistAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.AddProject(&_Artblocks.TransactOpts, _projectName, _artistAddress)
}

// AddProject is a paid mutator transaction binding the contract method 0xcc90e725.
//
// Solidity: function addProject(string _projectName, address _artistAddress) returns()
func (_Artblocks *ArtblocksTransactorSession) AddProject(_projectName string, _artistAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.AddProject(&_Artblocks.TransactOpts, _projectName, _artistAddress)
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

// AdminACLAllowed is a paid mutator transaction binding the contract method 0x230448b1.
//
// Solidity: function adminACLAllowed(address _sender, address _contract, bytes4 _selector) returns(bool)
func (_Artblocks *ArtblocksTransactor) AdminACLAllowed(opts *bind.TransactOpts, _sender common.Address, _contract common.Address, _selector [4]byte) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "adminACLAllowed", _sender, _contract, _selector)
}

// AdminACLAllowed is a paid mutator transaction binding the contract method 0x230448b1.
//
// Solidity: function adminACLAllowed(address _sender, address _contract, bytes4 _selector) returns(bool)
func (_Artblocks *ArtblocksSession) AdminACLAllowed(_sender common.Address, _contract common.Address, _selector [4]byte) (*types.Transaction, error) {
	return _Artblocks.Contract.AdminACLAllowed(&_Artblocks.TransactOpts, _sender, _contract, _selector)
}

// AdminACLAllowed is a paid mutator transaction binding the contract method 0x230448b1.
//
// Solidity: function adminACLAllowed(address _sender, address _contract, bytes4 _selector) returns(bool)
func (_Artblocks *ArtblocksTransactorSession) AdminACLAllowed(_sender common.Address, _contract common.Address, _selector [4]byte) (*types.Transaction, error) {
	return _Artblocks.Contract.AdminACLAllowed(&_Artblocks.TransactOpts, _sender, _contract, _selector)
}

// AdminAcceptArtistAddressesAndSplits is a paid mutator transaction binding the contract method 0x76ee6fab.
//
// Solidity: function adminAcceptArtistAddressesAndSplits(uint256 _projectId, address _artistAddress, address _additionalPayeePrimarySales, uint256 _additionalPayeePrimarySalesPercentage, address _additionalPayeeSecondarySales, uint256 _additionalPayeeSecondarySalesPercentage) returns()
func (_Artblocks *ArtblocksTransactor) AdminAcceptArtistAddressesAndSplits(opts *bind.TransactOpts, _projectId *big.Int, _artistAddress common.Address, _additionalPayeePrimarySales common.Address, _additionalPayeePrimarySalesPercentage *big.Int, _additionalPayeeSecondarySales common.Address, _additionalPayeeSecondarySalesPercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "adminAcceptArtistAddressesAndSplits", _projectId, _artistAddress, _additionalPayeePrimarySales, _additionalPayeePrimarySalesPercentage, _additionalPayeeSecondarySales, _additionalPayeeSecondarySalesPercentage)
}

// AdminAcceptArtistAddressesAndSplits is a paid mutator transaction binding the contract method 0x76ee6fab.
//
// Solidity: function adminAcceptArtistAddressesAndSplits(uint256 _projectId, address _artistAddress, address _additionalPayeePrimarySales, uint256 _additionalPayeePrimarySalesPercentage, address _additionalPayeeSecondarySales, uint256 _additionalPayeeSecondarySalesPercentage) returns()
func (_Artblocks *ArtblocksSession) AdminAcceptArtistAddressesAndSplits(_projectId *big.Int, _artistAddress common.Address, _additionalPayeePrimarySales common.Address, _additionalPayeePrimarySalesPercentage *big.Int, _additionalPayeeSecondarySales common.Address, _additionalPayeeSecondarySalesPercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.AdminAcceptArtistAddressesAndSplits(&_Artblocks.TransactOpts, _projectId, _artistAddress, _additionalPayeePrimarySales, _additionalPayeePrimarySalesPercentage, _additionalPayeeSecondarySales, _additionalPayeeSecondarySalesPercentage)
}

// AdminAcceptArtistAddressesAndSplits is a paid mutator transaction binding the contract method 0x76ee6fab.
//
// Solidity: function adminAcceptArtistAddressesAndSplits(uint256 _projectId, address _artistAddress, address _additionalPayeePrimarySales, uint256 _additionalPayeePrimarySalesPercentage, address _additionalPayeeSecondarySales, uint256 _additionalPayeeSecondarySalesPercentage) returns()
func (_Artblocks *ArtblocksTransactorSession) AdminAcceptArtistAddressesAndSplits(_projectId *big.Int, _artistAddress common.Address, _additionalPayeePrimarySales common.Address, _additionalPayeePrimarySalesPercentage *big.Int, _additionalPayeeSecondarySales common.Address, _additionalPayeeSecondarySalesPercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.AdminAcceptArtistAddressesAndSplits(&_Artblocks.TransactOpts, _projectId, _artistAddress, _additionalPayeePrimarySales, _additionalPayeePrimarySalesPercentage, _additionalPayeeSecondarySales, _additionalPayeeSecondarySalesPercentage)
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

// ForbidNewProjects is a paid mutator transaction binding the contract method 0x04143a5c.
//
// Solidity: function forbidNewProjects() returns()
func (_Artblocks *ArtblocksTransactor) ForbidNewProjects(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "forbidNewProjects")
}

// ForbidNewProjects is a paid mutator transaction binding the contract method 0x04143a5c.
//
// Solidity: function forbidNewProjects() returns()
func (_Artblocks *ArtblocksSession) ForbidNewProjects() (*types.Transaction, error) {
	return _Artblocks.Contract.ForbidNewProjects(&_Artblocks.TransactOpts)
}

// ForbidNewProjects is a paid mutator transaction binding the contract method 0x04143a5c.
//
// Solidity: function forbidNewProjects() returns()
func (_Artblocks *ArtblocksTransactorSession) ForbidNewProjects() (*types.Transaction, error) {
	return _Artblocks.Contract.ForbidNewProjects(&_Artblocks.TransactOpts)
}

// MintEcf is a paid mutator transaction binding the contract method 0x00005de5.
//
// Solidity: function mint_Ecf(address _to, uint256 _projectId, address _by) returns(uint256 _tokenId)
func (_Artblocks *ArtblocksTransactor) MintEcf(opts *bind.TransactOpts, _to common.Address, _projectId *big.Int, _by common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "mint_Ecf", _to, _projectId, _by)
}

// MintEcf is a paid mutator transaction binding the contract method 0x00005de5.
//
// Solidity: function mint_Ecf(address _to, uint256 _projectId, address _by) returns(uint256 _tokenId)
func (_Artblocks *ArtblocksSession) MintEcf(_to common.Address, _projectId *big.Int, _by common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.MintEcf(&_Artblocks.TransactOpts, _to, _projectId, _by)
}

// MintEcf is a paid mutator transaction binding the contract method 0x00005de5.
//
// Solidity: function mint_Ecf(address _to, uint256 _projectId, address _by) returns(uint256 _tokenId)
func (_Artblocks *ArtblocksTransactorSession) MintEcf(_to common.Address, _projectId *big.Int, _by common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.MintEcf(&_Artblocks.TransactOpts, _to, _projectId, _by)
}

// ProposeArtistPaymentAddressesAndSplits is a paid mutator transaction binding the contract method 0x2b65e67d.
//
// Solidity: function proposeArtistPaymentAddressesAndSplits(uint256 _projectId, address _artistAddress, address _additionalPayeePrimarySales, uint256 _additionalPayeePrimarySalesPercentage, address _additionalPayeeSecondarySales, uint256 _additionalPayeeSecondarySalesPercentage) returns()
func (_Artblocks *ArtblocksTransactor) ProposeArtistPaymentAddressesAndSplits(opts *bind.TransactOpts, _projectId *big.Int, _artistAddress common.Address, _additionalPayeePrimarySales common.Address, _additionalPayeePrimarySalesPercentage *big.Int, _additionalPayeeSecondarySales common.Address, _additionalPayeeSecondarySalesPercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "proposeArtistPaymentAddressesAndSplits", _projectId, _artistAddress, _additionalPayeePrimarySales, _additionalPayeePrimarySalesPercentage, _additionalPayeeSecondarySales, _additionalPayeeSecondarySalesPercentage)
}

// ProposeArtistPaymentAddressesAndSplits is a paid mutator transaction binding the contract method 0x2b65e67d.
//
// Solidity: function proposeArtistPaymentAddressesAndSplits(uint256 _projectId, address _artistAddress, address _additionalPayeePrimarySales, uint256 _additionalPayeePrimarySalesPercentage, address _additionalPayeeSecondarySales, uint256 _additionalPayeeSecondarySalesPercentage) returns()
func (_Artblocks *ArtblocksSession) ProposeArtistPaymentAddressesAndSplits(_projectId *big.Int, _artistAddress common.Address, _additionalPayeePrimarySales common.Address, _additionalPayeePrimarySalesPercentage *big.Int, _additionalPayeeSecondarySales common.Address, _additionalPayeeSecondarySalesPercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ProposeArtistPaymentAddressesAndSplits(&_Artblocks.TransactOpts, _projectId, _artistAddress, _additionalPayeePrimarySales, _additionalPayeePrimarySalesPercentage, _additionalPayeeSecondarySales, _additionalPayeeSecondarySalesPercentage)
}

// ProposeArtistPaymentAddressesAndSplits is a paid mutator transaction binding the contract method 0x2b65e67d.
//
// Solidity: function proposeArtistPaymentAddressesAndSplits(uint256 _projectId, address _artistAddress, address _additionalPayeePrimarySales, uint256 _additionalPayeePrimarySalesPercentage, address _additionalPayeeSecondarySales, uint256 _additionalPayeeSecondarySalesPercentage) returns()
func (_Artblocks *ArtblocksTransactorSession) ProposeArtistPaymentAddressesAndSplits(_projectId *big.Int, _artistAddress common.Address, _additionalPayeePrimarySales common.Address, _additionalPayeePrimarySalesPercentage *big.Int, _additionalPayeeSecondarySales common.Address, _additionalPayeeSecondarySalesPercentage *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.ProposeArtistPaymentAddressesAndSplits(&_Artblocks.TransactOpts, _projectId, _artistAddress, _additionalPayeePrimarySales, _additionalPayeePrimarySalesPercentage, _additionalPayeeSecondarySales, _additionalPayeeSecondarySalesPercentage)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Artblocks *ArtblocksTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Artblocks *ArtblocksSession) RenounceOwnership() (*types.Transaction, error) {
	return _Artblocks.Contract.RenounceOwnership(&_Artblocks.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Artblocks *ArtblocksTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Artblocks.Contract.RenounceOwnership(&_Artblocks.TransactOpts)
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
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Artblocks *ArtblocksTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Artblocks *ArtblocksSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Artblocks.Contract.SafeTransferFrom0(&_Artblocks.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Artblocks *ArtblocksTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Artblocks.Contract.SafeTransferFrom0(&_Artblocks.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Artblocks *ArtblocksTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Artblocks *ArtblocksSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Artblocks.Contract.SetApprovalForAll(&_Artblocks.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Artblocks *ArtblocksTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Artblocks.Contract.SetApprovalForAll(&_Artblocks.TransactOpts, operator, approved)
}

// SetTokenHash8PT is a paid mutator transaction binding the contract method 0x00001e3c.
//
// Solidity: function setTokenHash_8PT(uint256 _tokenId, bytes32 _hashSeed) returns()
func (_Artblocks *ArtblocksTransactor) SetTokenHash8PT(opts *bind.TransactOpts, _tokenId *big.Int, _hashSeed [32]byte) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "setTokenHash_8PT", _tokenId, _hashSeed)
}

// SetTokenHash8PT is a paid mutator transaction binding the contract method 0x00001e3c.
//
// Solidity: function setTokenHash_8PT(uint256 _tokenId, bytes32 _hashSeed) returns()
func (_Artblocks *ArtblocksSession) SetTokenHash8PT(_tokenId *big.Int, _hashSeed [32]byte) (*types.Transaction, error) {
	return _Artblocks.Contract.SetTokenHash8PT(&_Artblocks.TransactOpts, _tokenId, _hashSeed)
}

// SetTokenHash8PT is a paid mutator transaction binding the contract method 0x00001e3c.
//
// Solidity: function setTokenHash_8PT(uint256 _tokenId, bytes32 _hashSeed) returns()
func (_Artblocks *ArtblocksTransactorSession) SetTokenHash8PT(_tokenId *big.Int, _hashSeed [32]byte) (*types.Transaction, error) {
	return _Artblocks.Contract.SetTokenHash8PT(&_Artblocks.TransactOpts, _tokenId, _hashSeed)
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

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Artblocks *ArtblocksTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Artblocks *ArtblocksSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.TransferOwnership(&_Artblocks.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Artblocks *ArtblocksTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.TransferOwnership(&_Artblocks.TransactOpts, newOwner)
}

// UpdateArtblocksCurationRegistryAddress is a paid mutator transaction binding the contract method 0xa0bee564.
//
// Solidity: function updateArtblocksCurationRegistryAddress(address _artblocksCurationRegistryAddress) returns()
func (_Artblocks *ArtblocksTransactor) UpdateArtblocksCurationRegistryAddress(opts *bind.TransactOpts, _artblocksCurationRegistryAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateArtblocksCurationRegistryAddress", _artblocksCurationRegistryAddress)
}

// UpdateArtblocksCurationRegistryAddress is a paid mutator transaction binding the contract method 0xa0bee564.
//
// Solidity: function updateArtblocksCurationRegistryAddress(address _artblocksCurationRegistryAddress) returns()
func (_Artblocks *ArtblocksSession) UpdateArtblocksCurationRegistryAddress(_artblocksCurationRegistryAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksCurationRegistryAddress(&_Artblocks.TransactOpts, _artblocksCurationRegistryAddress)
}

// UpdateArtblocksCurationRegistryAddress is a paid mutator transaction binding the contract method 0xa0bee564.
//
// Solidity: function updateArtblocksCurationRegistryAddress(address _artblocksCurationRegistryAddress) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateArtblocksCurationRegistryAddress(_artblocksCurationRegistryAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksCurationRegistryAddress(&_Artblocks.TransactOpts, _artblocksCurationRegistryAddress)
}

// UpdateArtblocksDependencyRegistryAddress is a paid mutator transaction binding the contract method 0x2b274166.
//
// Solidity: function updateArtblocksDependencyRegistryAddress(address _artblocksDependencyRegistryAddress) returns()
func (_Artblocks *ArtblocksTransactor) UpdateArtblocksDependencyRegistryAddress(opts *bind.TransactOpts, _artblocksDependencyRegistryAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateArtblocksDependencyRegistryAddress", _artblocksDependencyRegistryAddress)
}

// UpdateArtblocksDependencyRegistryAddress is a paid mutator transaction binding the contract method 0x2b274166.
//
// Solidity: function updateArtblocksDependencyRegistryAddress(address _artblocksDependencyRegistryAddress) returns()
func (_Artblocks *ArtblocksSession) UpdateArtblocksDependencyRegistryAddress(_artblocksDependencyRegistryAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksDependencyRegistryAddress(&_Artblocks.TransactOpts, _artblocksDependencyRegistryAddress)
}

// UpdateArtblocksDependencyRegistryAddress is a paid mutator transaction binding the contract method 0x2b274166.
//
// Solidity: function updateArtblocksDependencyRegistryAddress(address _artblocksDependencyRegistryAddress) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateArtblocksDependencyRegistryAddress(_artblocksDependencyRegistryAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksDependencyRegistryAddress(&_Artblocks.TransactOpts, _artblocksDependencyRegistryAddress)
}

// UpdateArtblocksPrimarySalesAddress is a paid mutator transaction binding the contract method 0x0c5b1ad4.
//
// Solidity: function updateArtblocksPrimarySalesAddress(address _artblocksPrimarySalesAddress) returns()
func (_Artblocks *ArtblocksTransactor) UpdateArtblocksPrimarySalesAddress(opts *bind.TransactOpts, _artblocksPrimarySalesAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateArtblocksPrimarySalesAddress", _artblocksPrimarySalesAddress)
}

// UpdateArtblocksPrimarySalesAddress is a paid mutator transaction binding the contract method 0x0c5b1ad4.
//
// Solidity: function updateArtblocksPrimarySalesAddress(address _artblocksPrimarySalesAddress) returns()
func (_Artblocks *ArtblocksSession) UpdateArtblocksPrimarySalesAddress(_artblocksPrimarySalesAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksPrimarySalesAddress(&_Artblocks.TransactOpts, _artblocksPrimarySalesAddress)
}

// UpdateArtblocksPrimarySalesAddress is a paid mutator transaction binding the contract method 0x0c5b1ad4.
//
// Solidity: function updateArtblocksPrimarySalesAddress(address _artblocksPrimarySalesAddress) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateArtblocksPrimarySalesAddress(_artblocksPrimarySalesAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksPrimarySalesAddress(&_Artblocks.TransactOpts, _artblocksPrimarySalesAddress)
}

// UpdateArtblocksPrimarySalesPercentage is a paid mutator transaction binding the contract method 0x4bbc4ff0.
//
// Solidity: function updateArtblocksPrimarySalesPercentage(uint256 artblocksPrimarySalesPercentage_) returns()
func (_Artblocks *ArtblocksTransactor) UpdateArtblocksPrimarySalesPercentage(opts *bind.TransactOpts, artblocksPrimarySalesPercentage_ *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateArtblocksPrimarySalesPercentage", artblocksPrimarySalesPercentage_)
}

// UpdateArtblocksPrimarySalesPercentage is a paid mutator transaction binding the contract method 0x4bbc4ff0.
//
// Solidity: function updateArtblocksPrimarySalesPercentage(uint256 artblocksPrimarySalesPercentage_) returns()
func (_Artblocks *ArtblocksSession) UpdateArtblocksPrimarySalesPercentage(artblocksPrimarySalesPercentage_ *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksPrimarySalesPercentage(&_Artblocks.TransactOpts, artblocksPrimarySalesPercentage_)
}

// UpdateArtblocksPrimarySalesPercentage is a paid mutator transaction binding the contract method 0x4bbc4ff0.
//
// Solidity: function updateArtblocksPrimarySalesPercentage(uint256 artblocksPrimarySalesPercentage_) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateArtblocksPrimarySalesPercentage(artblocksPrimarySalesPercentage_ *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksPrimarySalesPercentage(&_Artblocks.TransactOpts, artblocksPrimarySalesPercentage_)
}

// UpdateArtblocksSecondarySalesAddress is a paid mutator transaction binding the contract method 0x2b6cfc8d.
//
// Solidity: function updateArtblocksSecondarySalesAddress(address _artblocksSecondarySalesAddress) returns()
func (_Artblocks *ArtblocksTransactor) UpdateArtblocksSecondarySalesAddress(opts *bind.TransactOpts, _artblocksSecondarySalesAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateArtblocksSecondarySalesAddress", _artblocksSecondarySalesAddress)
}

// UpdateArtblocksSecondarySalesAddress is a paid mutator transaction binding the contract method 0x2b6cfc8d.
//
// Solidity: function updateArtblocksSecondarySalesAddress(address _artblocksSecondarySalesAddress) returns()
func (_Artblocks *ArtblocksSession) UpdateArtblocksSecondarySalesAddress(_artblocksSecondarySalesAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksSecondarySalesAddress(&_Artblocks.TransactOpts, _artblocksSecondarySalesAddress)
}

// UpdateArtblocksSecondarySalesAddress is a paid mutator transaction binding the contract method 0x2b6cfc8d.
//
// Solidity: function updateArtblocksSecondarySalesAddress(address _artblocksSecondarySalesAddress) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateArtblocksSecondarySalesAddress(_artblocksSecondarySalesAddress common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksSecondarySalesAddress(&_Artblocks.TransactOpts, _artblocksSecondarySalesAddress)
}

// UpdateArtblocksSecondarySalesBPS is a paid mutator transaction binding the contract method 0x9ab31a2d.
//
// Solidity: function updateArtblocksSecondarySalesBPS(uint256 _artblocksSecondarySalesBPS) returns()
func (_Artblocks *ArtblocksTransactor) UpdateArtblocksSecondarySalesBPS(opts *bind.TransactOpts, _artblocksSecondarySalesBPS *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateArtblocksSecondarySalesBPS", _artblocksSecondarySalesBPS)
}

// UpdateArtblocksSecondarySalesBPS is a paid mutator transaction binding the contract method 0x9ab31a2d.
//
// Solidity: function updateArtblocksSecondarySalesBPS(uint256 _artblocksSecondarySalesBPS) returns()
func (_Artblocks *ArtblocksSession) UpdateArtblocksSecondarySalesBPS(_artblocksSecondarySalesBPS *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksSecondarySalesBPS(&_Artblocks.TransactOpts, _artblocksSecondarySalesBPS)
}

// UpdateArtblocksSecondarySalesBPS is a paid mutator transaction binding the contract method 0x9ab31a2d.
//
// Solidity: function updateArtblocksSecondarySalesBPS(uint256 _artblocksSecondarySalesBPS) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateArtblocksSecondarySalesBPS(_artblocksSecondarySalesBPS *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateArtblocksSecondarySalesBPS(&_Artblocks.TransactOpts, _artblocksSecondarySalesBPS)
}

// UpdateDefaultBaseURI is a paid mutator transaction binding the contract method 0x2302cbda.
//
// Solidity: function updateDefaultBaseURI(string _defaultBaseURI) returns()
func (_Artblocks *ArtblocksTransactor) UpdateDefaultBaseURI(opts *bind.TransactOpts, _defaultBaseURI string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateDefaultBaseURI", _defaultBaseURI)
}

// UpdateDefaultBaseURI is a paid mutator transaction binding the contract method 0x2302cbda.
//
// Solidity: function updateDefaultBaseURI(string _defaultBaseURI) returns()
func (_Artblocks *ArtblocksSession) UpdateDefaultBaseURI(_defaultBaseURI string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateDefaultBaseURI(&_Artblocks.TransactOpts, _defaultBaseURI)
}

// UpdateDefaultBaseURI is a paid mutator transaction binding the contract method 0x2302cbda.
//
// Solidity: function updateDefaultBaseURI(string _defaultBaseURI) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateDefaultBaseURI(_defaultBaseURI string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateDefaultBaseURI(&_Artblocks.TransactOpts, _defaultBaseURI)
}

// UpdateMinterContract is a paid mutator transaction binding the contract method 0x48337282.
//
// Solidity: function updateMinterContract(address _address) returns()
func (_Artblocks *ArtblocksTransactor) UpdateMinterContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateMinterContract", _address)
}

// UpdateMinterContract is a paid mutator transaction binding the contract method 0x48337282.
//
// Solidity: function updateMinterContract(address _address) returns()
func (_Artblocks *ArtblocksSession) UpdateMinterContract(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateMinterContract(&_Artblocks.TransactOpts, _address)
}

// UpdateMinterContract is a paid mutator transaction binding the contract method 0x48337282.
//
// Solidity: function updateMinterContract(address _address) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateMinterContract(_address common.Address) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateMinterContract(&_Artblocks.TransactOpts, _address)
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

// UpdateProjectAspectRatio is a paid mutator transaction binding the contract method 0x0e79c928.
//
// Solidity: function updateProjectAspectRatio(uint256 _projectId, string _aspectRatio) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectAspectRatio(opts *bind.TransactOpts, _projectId *big.Int, _aspectRatio string) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectAspectRatio", _projectId, _aspectRatio)
}

// UpdateProjectAspectRatio is a paid mutator transaction binding the contract method 0x0e79c928.
//
// Solidity: function updateProjectAspectRatio(uint256 _projectId, string _aspectRatio) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectAspectRatio(_projectId *big.Int, _aspectRatio string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectAspectRatio(&_Artblocks.TransactOpts, _projectId, _aspectRatio)
}

// UpdateProjectAspectRatio is a paid mutator transaction binding the contract method 0x0e79c928.
//
// Solidity: function updateProjectAspectRatio(uint256 _projectId, string _aspectRatio) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectAspectRatio(_projectId *big.Int, _aspectRatio string) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectAspectRatio(&_Artblocks.TransactOpts, _projectId, _aspectRatio)
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

// UpdateProjectMaxInvocations is a paid mutator transaction binding the contract method 0x0132c697.
//
// Solidity: function updateProjectMaxInvocations(uint256 _projectId, uint24 _maxInvocations) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectMaxInvocations(opts *bind.TransactOpts, _projectId *big.Int, _maxInvocations *big.Int) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectMaxInvocations", _projectId, _maxInvocations)
}

// UpdateProjectMaxInvocations is a paid mutator transaction binding the contract method 0x0132c697.
//
// Solidity: function updateProjectMaxInvocations(uint256 _projectId, uint24 _maxInvocations) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectMaxInvocations(_projectId *big.Int, _maxInvocations *big.Int) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectMaxInvocations(&_Artblocks.TransactOpts, _projectId, _maxInvocations)
}

// UpdateProjectMaxInvocations is a paid mutator transaction binding the contract method 0x0132c697.
//
// Solidity: function updateProjectMaxInvocations(uint256 _projectId, uint24 _maxInvocations) returns()
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

// UpdateProjectScriptType is a paid mutator transaction binding the contract method 0x01856fd4.
//
// Solidity: function updateProjectScriptType(uint256 _projectId, bytes32 _scriptTypeAndVersion) returns()
func (_Artblocks *ArtblocksTransactor) UpdateProjectScriptType(opts *bind.TransactOpts, _projectId *big.Int, _scriptTypeAndVersion [32]byte) (*types.Transaction, error) {
	return _Artblocks.contract.Transact(opts, "updateProjectScriptType", _projectId, _scriptTypeAndVersion)
}

// UpdateProjectScriptType is a paid mutator transaction binding the contract method 0x01856fd4.
//
// Solidity: function updateProjectScriptType(uint256 _projectId, bytes32 _scriptTypeAndVersion) returns()
func (_Artblocks *ArtblocksSession) UpdateProjectScriptType(_projectId *big.Int, _scriptTypeAndVersion [32]byte) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectScriptType(&_Artblocks.TransactOpts, _projectId, _scriptTypeAndVersion)
}

// UpdateProjectScriptType is a paid mutator transaction binding the contract method 0x01856fd4.
//
// Solidity: function updateProjectScriptType(uint256 _projectId, bytes32 _scriptTypeAndVersion) returns()
func (_Artblocks *ArtblocksTransactorSession) UpdateProjectScriptType(_projectId *big.Int, _scriptTypeAndVersion [32]byte) (*types.Transaction, error) {
	return _Artblocks.Contract.UpdateProjectScriptType(&_Artblocks.TransactOpts, _projectId, _scriptTypeAndVersion)
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

// ArtblocksAcceptedArtistAddressesAndSplitsIterator is returned from FilterAcceptedArtistAddressesAndSplits and is used to iterate over the raw logs and unpacked data for AcceptedArtistAddressesAndSplits events raised by the Artblocks contract.
type ArtblocksAcceptedArtistAddressesAndSplitsIterator struct {
	Event *ArtblocksAcceptedArtistAddressesAndSplits // Event containing the contract specifics and raw log

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
func (it *ArtblocksAcceptedArtistAddressesAndSplitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksAcceptedArtistAddressesAndSplits)
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
		it.Event = new(ArtblocksAcceptedArtistAddressesAndSplits)
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
func (it *ArtblocksAcceptedArtistAddressesAndSplitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksAcceptedArtistAddressesAndSplitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksAcceptedArtistAddressesAndSplits represents a AcceptedArtistAddressesAndSplits event raised by the Artblocks contract.
type ArtblocksAcceptedArtistAddressesAndSplits struct {
	ProjectId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAcceptedArtistAddressesAndSplits is a free log retrieval operation binding the contract event 0xc582d05e1da854143bd3271ef4529d79cf5a69fc6057ae320f357acfd291b738.
//
// Solidity: event AcceptedArtistAddressesAndSplits(uint256 indexed _projectId)
func (_Artblocks *ArtblocksFilterer) FilterAcceptedArtistAddressesAndSplits(opts *bind.FilterOpts, _projectId []*big.Int) (*ArtblocksAcceptedArtistAddressesAndSplitsIterator, error) {

	var _projectIdRule []interface{}
	for _, _projectIdItem := range _projectId {
		_projectIdRule = append(_projectIdRule, _projectIdItem)
	}

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "AcceptedArtistAddressesAndSplits", _projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksAcceptedArtistAddressesAndSplitsIterator{contract: _Artblocks.contract, event: "AcceptedArtistAddressesAndSplits", logs: logs, sub: sub}, nil
}

// WatchAcceptedArtistAddressesAndSplits is a free log subscription operation binding the contract event 0xc582d05e1da854143bd3271ef4529d79cf5a69fc6057ae320f357acfd291b738.
//
// Solidity: event AcceptedArtistAddressesAndSplits(uint256 indexed _projectId)
func (_Artblocks *ArtblocksFilterer) WatchAcceptedArtistAddressesAndSplits(opts *bind.WatchOpts, sink chan<- *ArtblocksAcceptedArtistAddressesAndSplits, _projectId []*big.Int) (event.Subscription, error) {

	var _projectIdRule []interface{}
	for _, _projectIdItem := range _projectId {
		_projectIdRule = append(_projectIdRule, _projectIdItem)
	}

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "AcceptedArtistAddressesAndSplits", _projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksAcceptedArtistAddressesAndSplits)
				if err := _Artblocks.contract.UnpackLog(event, "AcceptedArtistAddressesAndSplits", log); err != nil {
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

// ParseAcceptedArtistAddressesAndSplits is a log parse operation binding the contract event 0xc582d05e1da854143bd3271ef4529d79cf5a69fc6057ae320f357acfd291b738.
//
// Solidity: event AcceptedArtistAddressesAndSplits(uint256 indexed _projectId)
func (_Artblocks *ArtblocksFilterer) ParseAcceptedArtistAddressesAndSplits(log types.Log) (*ArtblocksAcceptedArtistAddressesAndSplits, error) {
	event := new(ArtblocksAcceptedArtistAddressesAndSplits)
	if err := _Artblocks.contract.UnpackLog(event, "AcceptedArtistAddressesAndSplits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_Artblocks *ArtblocksFilterer) FilterMint(opts *bind.FilterOpts, _to []common.Address, _tokenId []*big.Int) (*ArtblocksMintIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "Mint", _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksMintIterator{contract: _Artblocks.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_Artblocks *ArtblocksFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ArtblocksMint, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "Mint", _toRule, _tokenIdRule)
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_Artblocks *ArtblocksFilterer) ParseMint(log types.Log) (*ArtblocksMint, error) {
	event := new(ArtblocksMint)
	if err := _Artblocks.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtblocksMinterUpdatedIterator is returned from FilterMinterUpdated and is used to iterate over the raw logs and unpacked data for MinterUpdated events raised by the Artblocks contract.
type ArtblocksMinterUpdatedIterator struct {
	Event *ArtblocksMinterUpdated // Event containing the contract specifics and raw log

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
func (it *ArtblocksMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksMinterUpdated)
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
		it.Event = new(ArtblocksMinterUpdated)
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
func (it *ArtblocksMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksMinterUpdated represents a MinterUpdated event raised by the Artblocks contract.
type ArtblocksMinterUpdated struct {
	CurrentMinter common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMinterUpdated is a free log retrieval operation binding the contract event 0xad0f299ec81a386c98df0ac27dae11dd020ed1b56963c53a7292e7a3a314539a.
//
// Solidity: event MinterUpdated(address indexed _currentMinter)
func (_Artblocks *ArtblocksFilterer) FilterMinterUpdated(opts *bind.FilterOpts, _currentMinter []common.Address) (*ArtblocksMinterUpdatedIterator, error) {

	var _currentMinterRule []interface{}
	for _, _currentMinterItem := range _currentMinter {
		_currentMinterRule = append(_currentMinterRule, _currentMinterItem)
	}

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "MinterUpdated", _currentMinterRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksMinterUpdatedIterator{contract: _Artblocks.contract, event: "MinterUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterUpdated is a free log subscription operation binding the contract event 0xad0f299ec81a386c98df0ac27dae11dd020ed1b56963c53a7292e7a3a314539a.
//
// Solidity: event MinterUpdated(address indexed _currentMinter)
func (_Artblocks *ArtblocksFilterer) WatchMinterUpdated(opts *bind.WatchOpts, sink chan<- *ArtblocksMinterUpdated, _currentMinter []common.Address) (event.Subscription, error) {

	var _currentMinterRule []interface{}
	for _, _currentMinterItem := range _currentMinter {
		_currentMinterRule = append(_currentMinterRule, _currentMinterItem)
	}

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "MinterUpdated", _currentMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksMinterUpdated)
				if err := _Artblocks.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
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

// ParseMinterUpdated is a log parse operation binding the contract event 0xad0f299ec81a386c98df0ac27dae11dd020ed1b56963c53a7292e7a3a314539a.
//
// Solidity: event MinterUpdated(address indexed _currentMinter)
func (_Artblocks *ArtblocksFilterer) ParseMinterUpdated(log types.Log) (*ArtblocksMinterUpdated, error) {
	event := new(ArtblocksMinterUpdated)
	if err := _Artblocks.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtblocksOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Artblocks contract.
type ArtblocksOwnershipTransferredIterator struct {
	Event *ArtblocksOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArtblocksOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksOwnershipTransferred)
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
		it.Event = new(ArtblocksOwnershipTransferred)
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
func (it *ArtblocksOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksOwnershipTransferred represents a OwnershipTransferred event raised by the Artblocks contract.
type ArtblocksOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Artblocks *ArtblocksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArtblocksOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksOwnershipTransferredIterator{contract: _Artblocks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Artblocks *ArtblocksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArtblocksOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksOwnershipTransferred)
				if err := _Artblocks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Artblocks *ArtblocksFilterer) ParseOwnershipTransferred(log types.Log) (*ArtblocksOwnershipTransferred, error) {
	event := new(ArtblocksOwnershipTransferred)
	if err := _Artblocks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtblocksPlatformUpdatedIterator is returned from FilterPlatformUpdated and is used to iterate over the raw logs and unpacked data for PlatformUpdated events raised by the Artblocks contract.
type ArtblocksPlatformUpdatedIterator struct {
	Event *ArtblocksPlatformUpdated // Event containing the contract specifics and raw log

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
func (it *ArtblocksPlatformUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksPlatformUpdated)
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
		it.Event = new(ArtblocksPlatformUpdated)
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
func (it *ArtblocksPlatformUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksPlatformUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksPlatformUpdated represents a PlatformUpdated event raised by the Artblocks contract.
type ArtblocksPlatformUpdated struct {
	Field [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPlatformUpdated is a free log retrieval operation binding the contract event 0x8b810f233ce7ee6e962ab4d98bf0277751de1f5589de3dcc812ac2047994d009.
//
// Solidity: event PlatformUpdated(bytes32 indexed _field)
func (_Artblocks *ArtblocksFilterer) FilterPlatformUpdated(opts *bind.FilterOpts, _field [][32]byte) (*ArtblocksPlatformUpdatedIterator, error) {

	var _fieldRule []interface{}
	for _, _fieldItem := range _field {
		_fieldRule = append(_fieldRule, _fieldItem)
	}

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "PlatformUpdated", _fieldRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksPlatformUpdatedIterator{contract: _Artblocks.contract, event: "PlatformUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformUpdated is a free log subscription operation binding the contract event 0x8b810f233ce7ee6e962ab4d98bf0277751de1f5589de3dcc812ac2047994d009.
//
// Solidity: event PlatformUpdated(bytes32 indexed _field)
func (_Artblocks *ArtblocksFilterer) WatchPlatformUpdated(opts *bind.WatchOpts, sink chan<- *ArtblocksPlatformUpdated, _field [][32]byte) (event.Subscription, error) {

	var _fieldRule []interface{}
	for _, _fieldItem := range _field {
		_fieldRule = append(_fieldRule, _fieldItem)
	}

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "PlatformUpdated", _fieldRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksPlatformUpdated)
				if err := _Artblocks.contract.UnpackLog(event, "PlatformUpdated", log); err != nil {
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

// ParsePlatformUpdated is a log parse operation binding the contract event 0x8b810f233ce7ee6e962ab4d98bf0277751de1f5589de3dcc812ac2047994d009.
//
// Solidity: event PlatformUpdated(bytes32 indexed _field)
func (_Artblocks *ArtblocksFilterer) ParsePlatformUpdated(log types.Log) (*ArtblocksPlatformUpdated, error) {
	event := new(ArtblocksPlatformUpdated)
	if err := _Artblocks.contract.UnpackLog(event, "PlatformUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtblocksProjectUpdatedIterator is returned from FilterProjectUpdated and is used to iterate over the raw logs and unpacked data for ProjectUpdated events raised by the Artblocks contract.
type ArtblocksProjectUpdatedIterator struct {
	Event *ArtblocksProjectUpdated // Event containing the contract specifics and raw log

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
func (it *ArtblocksProjectUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksProjectUpdated)
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
		it.Event = new(ArtblocksProjectUpdated)
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
func (it *ArtblocksProjectUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksProjectUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksProjectUpdated represents a ProjectUpdated event raised by the Artblocks contract.
type ArtblocksProjectUpdated struct {
	ProjectId *big.Int
	Update    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectUpdated is a free log retrieval operation binding the contract event 0xb96a30340e86d03ce4be42f94ac02d7b27b4a4cdae942beb69026718dfe66afc.
//
// Solidity: event ProjectUpdated(uint256 indexed _projectId, bytes32 indexed _update)
func (_Artblocks *ArtblocksFilterer) FilterProjectUpdated(opts *bind.FilterOpts, _projectId []*big.Int, _update [][32]byte) (*ArtblocksProjectUpdatedIterator, error) {

	var _projectIdRule []interface{}
	for _, _projectIdItem := range _projectId {
		_projectIdRule = append(_projectIdRule, _projectIdItem)
	}
	var _updateRule []interface{}
	for _, _updateItem := range _update {
		_updateRule = append(_updateRule, _updateItem)
	}

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "ProjectUpdated", _projectIdRule, _updateRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksProjectUpdatedIterator{contract: _Artblocks.contract, event: "ProjectUpdated", logs: logs, sub: sub}, nil
}

// WatchProjectUpdated is a free log subscription operation binding the contract event 0xb96a30340e86d03ce4be42f94ac02d7b27b4a4cdae942beb69026718dfe66afc.
//
// Solidity: event ProjectUpdated(uint256 indexed _projectId, bytes32 indexed _update)
func (_Artblocks *ArtblocksFilterer) WatchProjectUpdated(opts *bind.WatchOpts, sink chan<- *ArtblocksProjectUpdated, _projectId []*big.Int, _update [][32]byte) (event.Subscription, error) {

	var _projectIdRule []interface{}
	for _, _projectIdItem := range _projectId {
		_projectIdRule = append(_projectIdRule, _projectIdItem)
	}
	var _updateRule []interface{}
	for _, _updateItem := range _update {
		_updateRule = append(_updateRule, _updateItem)
	}

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "ProjectUpdated", _projectIdRule, _updateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksProjectUpdated)
				if err := _Artblocks.contract.UnpackLog(event, "ProjectUpdated", log); err != nil {
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

// ParseProjectUpdated is a log parse operation binding the contract event 0xb96a30340e86d03ce4be42f94ac02d7b27b4a4cdae942beb69026718dfe66afc.
//
// Solidity: event ProjectUpdated(uint256 indexed _projectId, bytes32 indexed _update)
func (_Artblocks *ArtblocksFilterer) ParseProjectUpdated(log types.Log) (*ArtblocksProjectUpdated, error) {
	event := new(ArtblocksProjectUpdated)
	if err := _Artblocks.contract.UnpackLog(event, "ProjectUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtblocksProposedArtistAddressesAndSplitsIterator is returned from FilterProposedArtistAddressesAndSplits and is used to iterate over the raw logs and unpacked data for ProposedArtistAddressesAndSplits events raised by the Artblocks contract.
type ArtblocksProposedArtistAddressesAndSplitsIterator struct {
	Event *ArtblocksProposedArtistAddressesAndSplits // Event containing the contract specifics and raw log

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
func (it *ArtblocksProposedArtistAddressesAndSplitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtblocksProposedArtistAddressesAndSplits)
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
		it.Event = new(ArtblocksProposedArtistAddressesAndSplits)
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
func (it *ArtblocksProposedArtistAddressesAndSplitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtblocksProposedArtistAddressesAndSplitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtblocksProposedArtistAddressesAndSplits represents a ProposedArtistAddressesAndSplits event raised by the Artblocks contract.
type ArtblocksProposedArtistAddressesAndSplits struct {
	ProjectId                               *big.Int
	ArtistAddress                           common.Address
	AdditionalPayeePrimarySales             common.Address
	AdditionalPayeePrimarySalesPercentage   *big.Int
	AdditionalPayeeSecondarySales           common.Address
	AdditionalPayeeSecondarySalesPercentage *big.Int
	Raw                                     types.Log // Blockchain specific contextual infos
}

// FilterProposedArtistAddressesAndSplits is a free log retrieval operation binding the contract event 0x6ff7d102bb3657a26dcbbcd299d821a066718a7cf76ae7cd98279f18b74da8ac.
//
// Solidity: event ProposedArtistAddressesAndSplits(uint256 indexed _projectId, address _artistAddress, address _additionalPayeePrimarySales, uint256 _additionalPayeePrimarySalesPercentage, address _additionalPayeeSecondarySales, uint256 _additionalPayeeSecondarySalesPercentage)
func (_Artblocks *ArtblocksFilterer) FilterProposedArtistAddressesAndSplits(opts *bind.FilterOpts, _projectId []*big.Int) (*ArtblocksProposedArtistAddressesAndSplitsIterator, error) {

	var _projectIdRule []interface{}
	for _, _projectIdItem := range _projectId {
		_projectIdRule = append(_projectIdRule, _projectIdItem)
	}

	logs, sub, err := _Artblocks.contract.FilterLogs(opts, "ProposedArtistAddressesAndSplits", _projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtblocksProposedArtistAddressesAndSplitsIterator{contract: _Artblocks.contract, event: "ProposedArtistAddressesAndSplits", logs: logs, sub: sub}, nil
}

// WatchProposedArtistAddressesAndSplits is a free log subscription operation binding the contract event 0x6ff7d102bb3657a26dcbbcd299d821a066718a7cf76ae7cd98279f18b74da8ac.
//
// Solidity: event ProposedArtistAddressesAndSplits(uint256 indexed _projectId, address _artistAddress, address _additionalPayeePrimarySales, uint256 _additionalPayeePrimarySalesPercentage, address _additionalPayeeSecondarySales, uint256 _additionalPayeeSecondarySalesPercentage)
func (_Artblocks *ArtblocksFilterer) WatchProposedArtistAddressesAndSplits(opts *bind.WatchOpts, sink chan<- *ArtblocksProposedArtistAddressesAndSplits, _projectId []*big.Int) (event.Subscription, error) {

	var _projectIdRule []interface{}
	for _, _projectIdItem := range _projectId {
		_projectIdRule = append(_projectIdRule, _projectIdItem)
	}

	logs, sub, err := _Artblocks.contract.WatchLogs(opts, "ProposedArtistAddressesAndSplits", _projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtblocksProposedArtistAddressesAndSplits)
				if err := _Artblocks.contract.UnpackLog(event, "ProposedArtistAddressesAndSplits", log); err != nil {
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

// ParseProposedArtistAddressesAndSplits is a log parse operation binding the contract event 0x6ff7d102bb3657a26dcbbcd299d821a066718a7cf76ae7cd98279f18b74da8ac.
//
// Solidity: event ProposedArtistAddressesAndSplits(uint256 indexed _projectId, address _artistAddress, address _additionalPayeePrimarySales, uint256 _additionalPayeePrimarySalesPercentage, address _additionalPayeeSecondarySales, uint256 _additionalPayeeSecondarySalesPercentage)
func (_Artblocks *ArtblocksFilterer) ParseProposedArtistAddressesAndSplits(log types.Log) (*ArtblocksProposedArtistAddressesAndSplits, error) {
	event := new(ArtblocksProposedArtistAddressesAndSplits)
	if err := _Artblocks.contract.UnpackLog(event, "ProposedArtistAddressesAndSplits", log); err != nil {
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
