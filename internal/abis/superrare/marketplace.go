// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package superrare

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

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"addresspayable[]\",\"name\":\"_splitAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8[]\",\"name\":\"_splitRatios\",\"type\":\"uint8[]\"}],\"name\":\"AcceptOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_startedAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newAuctionLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_previousBidder\",\"type\":\"address\"}],\"name\":\"AuctionBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"AuctionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_auctionCreator\",\"type\":\"address\"}],\"name\":\"CancelAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"CancelOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_auctionCreator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minimumBid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lengthOfAuction\",\"type\":\"uint256\"}],\"name\":\"NewAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_convertible\",\"type\":\"bool\"}],\"name\":\"OfferPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"addresspayable[]\",\"name\":\"_splitRecipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8[]\",\"name\":\"_splitRatios\",\"type\":\"uint8[]\"}],\"name\":\"SetSalePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Sold\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLDIE_AUCTION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NO_AUCTION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SCHEDULED_AUCTION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable[]\",\"name\":\"_splitAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_splitRatios\",\"type\":\"uint8[]\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvedTokenRegistry\",\"outputs\":[{\"internalType\":\"contractIApprovedTokenRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionBids\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"marketplaceFee\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionLengthExtension\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_auctionType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lengthOfAuction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable[]\",\"name\":\"_splitAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_splitRatios\",\"type\":\"uint8[]\"}],\"name\":\"configureAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lengthOfAuction\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable[]\",\"name\":\"_splitAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_splitRatios\",\"type\":\"uint8[]\"}],\"name\":\"convertOfferToAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuctionDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"getSalePrice\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyEngine\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_superRareMarketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_superRareAuctionHouse\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spaceOperatorRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_approvedTokenRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_payments\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkBeneficiary\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplaceSettings\",\"outputs\":[{\"internalType\":\"contractIMarketplaceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAuctionLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumBidIncreasePercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkBeneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_convertible\",\"type\":\"bool\"}],\"name\":\"offer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offerCancelationDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"contractIPayments\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"removeSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyEngine\",\"outputs\":[{\"internalType\":\"contractIRoyaltyEngineV1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyRegistry\",\"outputs\":[{\"internalType\":\"contractIERC721CreatorRoyalty\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_approvedTokenRegistry\",\"type\":\"address\"}],\"name\":\"setApprovedTokenRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionLengthExtension\",\"type\":\"uint256\"}],\"name\":\"setAuctionLengthExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceSettings\",\"type\":\"address\"}],\"name\":\"setMarketplaceSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_maxAuctionLength\",\"type\":\"uint8\"}],\"name\":\"setMaxAuctionLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_minimumBidIncreasePercentage\",\"type\":\"uint8\"}],\"name\":\"setMinimumBidIncreasePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_networkBeneficiary\",\"type\":\"address\"}],\"name\":\"setNetworkBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offerCancelationDelay\",\"type\":\"uint256\"}],\"name\":\"setOfferCancelationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payments\",\"type\":\"address\"}],\"name\":\"setPayments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyEngine\",\"type\":\"address\"}],\"name\":\"setRoyaltyEngine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyRegistry\",\"type\":\"address\"}],\"name\":\"setRoyaltyRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_listPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"_splitAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_splitRatios\",\"type\":\"uint8[]\"}],\"name\":\"setSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spaceOperatorRegistry\",\"type\":\"address\"}],\"name\":\"setSpaceOperatorRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingRegistry\",\"type\":\"address\"}],\"name\":\"setStakingRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_superRareAuctionHouse\",\"type\":\"address\"}],\"name\":\"setSuperRareAuctionHouse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_superRareMarketplace\",\"type\":\"address\"}],\"name\":\"setSuperRareMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"settleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spaceOperatorRegistry\",\"outputs\":[{\"internalType\":\"contractISpaceOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superRareAuctionHouse\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superRareMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenAuctions\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"auctionCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lengthOfAuction\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionType\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenCurrentOffers\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"marketplaceFee\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"convertible\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenSalePrices\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// COLDIEAUCTION is a free data retrieval call binding the contract method 0xc90b8714.
//
// Solidity: function COLDIE_AUCTION() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) COLDIEAUCTION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "COLDIE_AUCTION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// COLDIEAUCTION is a free data retrieval call binding the contract method 0xc90b8714.
//
// Solidity: function COLDIE_AUCTION() view returns(bytes32)
func (_Marketplace *MarketplaceSession) COLDIEAUCTION() ([32]byte, error) {
	return _Marketplace.Contract.COLDIEAUCTION(&_Marketplace.CallOpts)
}

// COLDIEAUCTION is a free data retrieval call binding the contract method 0xc90b8714.
//
// Solidity: function COLDIE_AUCTION() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) COLDIEAUCTION() ([32]byte, error) {
	return _Marketplace.Contract.COLDIEAUCTION(&_Marketplace.CallOpts)
}

// NOAUCTION is a free data retrieval call binding the contract method 0x155a56b1.
//
// Solidity: function NO_AUCTION() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) NOAUCTION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "NO_AUCTION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NOAUCTION is a free data retrieval call binding the contract method 0x155a56b1.
//
// Solidity: function NO_AUCTION() view returns(bytes32)
func (_Marketplace *MarketplaceSession) NOAUCTION() ([32]byte, error) {
	return _Marketplace.Contract.NOAUCTION(&_Marketplace.CallOpts)
}

// NOAUCTION is a free data retrieval call binding the contract method 0x155a56b1.
//
// Solidity: function NO_AUCTION() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) NOAUCTION() ([32]byte, error) {
	return _Marketplace.Contract.NOAUCTION(&_Marketplace.CallOpts)
}

// SCHEDULEDAUCTION is a free data retrieval call binding the contract method 0xb23afc26.
//
// Solidity: function SCHEDULED_AUCTION() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) SCHEDULEDAUCTION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "SCHEDULED_AUCTION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SCHEDULEDAUCTION is a free data retrieval call binding the contract method 0xb23afc26.
//
// Solidity: function SCHEDULED_AUCTION() view returns(bytes32)
func (_Marketplace *MarketplaceSession) SCHEDULEDAUCTION() ([32]byte, error) {
	return _Marketplace.Contract.SCHEDULEDAUCTION(&_Marketplace.CallOpts)
}

// SCHEDULEDAUCTION is a free data retrieval call binding the contract method 0xb23afc26.
//
// Solidity: function SCHEDULED_AUCTION() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) SCHEDULEDAUCTION() ([32]byte, error) {
	return _Marketplace.Contract.SCHEDULEDAUCTION(&_Marketplace.CallOpts)
}

// ApprovedTokenRegistry is a free data retrieval call binding the contract method 0x6240cd1c.
//
// Solidity: function approvedTokenRegistry() view returns(address)
func (_Marketplace *MarketplaceCaller) ApprovedTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "approvedTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApprovedTokenRegistry is a free data retrieval call binding the contract method 0x6240cd1c.
//
// Solidity: function approvedTokenRegistry() view returns(address)
func (_Marketplace *MarketplaceSession) ApprovedTokenRegistry() (common.Address, error) {
	return _Marketplace.Contract.ApprovedTokenRegistry(&_Marketplace.CallOpts)
}

// ApprovedTokenRegistry is a free data retrieval call binding the contract method 0x6240cd1c.
//
// Solidity: function approvedTokenRegistry() view returns(address)
func (_Marketplace *MarketplaceCallerSession) ApprovedTokenRegistry() (common.Address, error) {
	return _Marketplace.Contract.ApprovedTokenRegistry(&_Marketplace.CallOpts)
}

// AuctionBids is a free data retrieval call binding the contract method 0x299a0e1e.
//
// Solidity: function auctionBids(address , uint256 ) view returns(address bidder, address currencyAddress, uint256 amount, uint8 marketplaceFee)
func (_Marketplace *MarketplaceCaller) AuctionBids(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bidder          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
	MarketplaceFee  uint8
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "auctionBids", arg0, arg1)

	outstruct := new(struct {
		Bidder          common.Address
		CurrencyAddress common.Address
		Amount          *big.Int
		MarketplaceFee  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CurrencyAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MarketplaceFee = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// AuctionBids is a free data retrieval call binding the contract method 0x299a0e1e.
//
// Solidity: function auctionBids(address , uint256 ) view returns(address bidder, address currencyAddress, uint256 amount, uint8 marketplaceFee)
func (_Marketplace *MarketplaceSession) AuctionBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
	MarketplaceFee  uint8
}, error) {
	return _Marketplace.Contract.AuctionBids(&_Marketplace.CallOpts, arg0, arg1)
}

// AuctionBids is a free data retrieval call binding the contract method 0x299a0e1e.
//
// Solidity: function auctionBids(address , uint256 ) view returns(address bidder, address currencyAddress, uint256 amount, uint8 marketplaceFee)
func (_Marketplace *MarketplaceCallerSession) AuctionBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
	MarketplaceFee  uint8
}, error) {
	return _Marketplace.Contract.AuctionBids(&_Marketplace.CallOpts, arg0, arg1)
}

// AuctionLengthExtension is a free data retrieval call binding the contract method 0xdaa26499.
//
// Solidity: function auctionLengthExtension() view returns(uint256)
func (_Marketplace *MarketplaceCaller) AuctionLengthExtension(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "auctionLengthExtension")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionLengthExtension is a free data retrieval call binding the contract method 0xdaa26499.
//
// Solidity: function auctionLengthExtension() view returns(uint256)
func (_Marketplace *MarketplaceSession) AuctionLengthExtension() (*big.Int, error) {
	return _Marketplace.Contract.AuctionLengthExtension(&_Marketplace.CallOpts)
}

// AuctionLengthExtension is a free data retrieval call binding the contract method 0xdaa26499.
//
// Solidity: function auctionLengthExtension() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) AuctionLengthExtension() (*big.Int, error) {
	return _Marketplace.Contract.AuctionLengthExtension(&_Marketplace.CallOpts)
}

// GetAuctionDetails is a free data retrieval call binding the contract method 0x0cd87c68.
//
// Solidity: function getAuctionDetails(address _originContract, uint256 _tokenId) view returns(address, uint256, uint256, uint256, address, uint256, bytes32, address[], uint8[])
func (_Marketplace *MarketplaceCaller) GetAuctionDetails(opts *bind.CallOpts, _originContract common.Address, _tokenId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, common.Address, *big.Int, [32]byte, []common.Address, []uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getAuctionDetails", _originContract, _tokenId)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(common.Address), *new(*big.Int), *new([32]byte), *new([]common.Address), *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	out7 := *abi.ConvertType(out[7], new([]common.Address)).(*[]common.Address)
	out8 := *abi.ConvertType(out[8], new([]uint8)).(*[]uint8)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAuctionDetails is a free data retrieval call binding the contract method 0x0cd87c68.
//
// Solidity: function getAuctionDetails(address _originContract, uint256 _tokenId) view returns(address, uint256, uint256, uint256, address, uint256, bytes32, address[], uint8[])
func (_Marketplace *MarketplaceSession) GetAuctionDetails(_originContract common.Address, _tokenId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, common.Address, *big.Int, [32]byte, []common.Address, []uint8, error) {
	return _Marketplace.Contract.GetAuctionDetails(&_Marketplace.CallOpts, _originContract, _tokenId)
}

// GetAuctionDetails is a free data retrieval call binding the contract method 0x0cd87c68.
//
// Solidity: function getAuctionDetails(address _originContract, uint256 _tokenId) view returns(address, uint256, uint256, uint256, address, uint256, bytes32, address[], uint8[])
func (_Marketplace *MarketplaceCallerSession) GetAuctionDetails(_originContract common.Address, _tokenId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, common.Address, *big.Int, [32]byte, []common.Address, []uint8, error) {
	return _Marketplace.Contract.GetAuctionDetails(&_Marketplace.CallOpts, _originContract, _tokenId)
}

// GetSalePrice is a free data retrieval call binding the contract method 0x369679a4.
//
// Solidity: function getSalePrice(address _originContract, uint256 _tokenId, address _target) view returns(address, address, uint256, address[], uint8[])
func (_Marketplace *MarketplaceCaller) GetSalePrice(opts *bind.CallOpts, _originContract common.Address, _tokenId *big.Int, _target common.Address) (common.Address, common.Address, *big.Int, []common.Address, []uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getSalePrice", _originContract, _tokenId, _target)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(*big.Int), *new([]common.Address), *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)
	out4 := *abi.ConvertType(out[4], new([]uint8)).(*[]uint8)

	return out0, out1, out2, out3, out4, err

}

// GetSalePrice is a free data retrieval call binding the contract method 0x369679a4.
//
// Solidity: function getSalePrice(address _originContract, uint256 _tokenId, address _target) view returns(address, address, uint256, address[], uint8[])
func (_Marketplace *MarketplaceSession) GetSalePrice(_originContract common.Address, _tokenId *big.Int, _target common.Address) (common.Address, common.Address, *big.Int, []common.Address, []uint8, error) {
	return _Marketplace.Contract.GetSalePrice(&_Marketplace.CallOpts, _originContract, _tokenId, _target)
}

// GetSalePrice is a free data retrieval call binding the contract method 0x369679a4.
//
// Solidity: function getSalePrice(address _originContract, uint256 _tokenId, address _target) view returns(address, address, uint256, address[], uint8[])
func (_Marketplace *MarketplaceCallerSession) GetSalePrice(_originContract common.Address, _tokenId *big.Int, _target common.Address) (common.Address, common.Address, *big.Int, []common.Address, []uint8, error) {
	return _Marketplace.Contract.GetSalePrice(&_Marketplace.CallOpts, _originContract, _tokenId, _target)
}

// MarketplaceSettings is a free data retrieval call binding the contract method 0xba50b632.
//
// Solidity: function marketplaceSettings() view returns(address)
func (_Marketplace *MarketplaceCaller) MarketplaceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "marketplaceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketplaceSettings is a free data retrieval call binding the contract method 0xba50b632.
//
// Solidity: function marketplaceSettings() view returns(address)
func (_Marketplace *MarketplaceSession) MarketplaceSettings() (common.Address, error) {
	return _Marketplace.Contract.MarketplaceSettings(&_Marketplace.CallOpts)
}

// MarketplaceSettings is a free data retrieval call binding the contract method 0xba50b632.
//
// Solidity: function marketplaceSettings() view returns(address)
func (_Marketplace *MarketplaceCallerSession) MarketplaceSettings() (common.Address, error) {
	return _Marketplace.Contract.MarketplaceSettings(&_Marketplace.CallOpts)
}

// MaxAuctionLength is a free data retrieval call binding the contract method 0x0e519ef9.
//
// Solidity: function maxAuctionLength() view returns(uint256)
func (_Marketplace *MarketplaceCaller) MaxAuctionLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "maxAuctionLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAuctionLength is a free data retrieval call binding the contract method 0x0e519ef9.
//
// Solidity: function maxAuctionLength() view returns(uint256)
func (_Marketplace *MarketplaceSession) MaxAuctionLength() (*big.Int, error) {
	return _Marketplace.Contract.MaxAuctionLength(&_Marketplace.CallOpts)
}

// MaxAuctionLength is a free data retrieval call binding the contract method 0x0e519ef9.
//
// Solidity: function maxAuctionLength() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) MaxAuctionLength() (*big.Int, error) {
	return _Marketplace.Contract.MaxAuctionLength(&_Marketplace.CallOpts)
}

// MinimumBidIncreasePercentage is a free data retrieval call binding the contract method 0x6fe9f44c.
//
// Solidity: function minimumBidIncreasePercentage() view returns(uint8)
func (_Marketplace *MarketplaceCaller) MinimumBidIncreasePercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "minimumBidIncreasePercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinimumBidIncreasePercentage is a free data retrieval call binding the contract method 0x6fe9f44c.
//
// Solidity: function minimumBidIncreasePercentage() view returns(uint8)
func (_Marketplace *MarketplaceSession) MinimumBidIncreasePercentage() (uint8, error) {
	return _Marketplace.Contract.MinimumBidIncreasePercentage(&_Marketplace.CallOpts)
}

// MinimumBidIncreasePercentage is a free data retrieval call binding the contract method 0x6fe9f44c.
//
// Solidity: function minimumBidIncreasePercentage() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) MinimumBidIncreasePercentage() (uint8, error) {
	return _Marketplace.Contract.MinimumBidIncreasePercentage(&_Marketplace.CallOpts)
}

// NetworkBeneficiary is a free data retrieval call binding the contract method 0x6b534ed0.
//
// Solidity: function networkBeneficiary() view returns(address)
func (_Marketplace *MarketplaceCaller) NetworkBeneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "networkBeneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkBeneficiary is a free data retrieval call binding the contract method 0x6b534ed0.
//
// Solidity: function networkBeneficiary() view returns(address)
func (_Marketplace *MarketplaceSession) NetworkBeneficiary() (common.Address, error) {
	return _Marketplace.Contract.NetworkBeneficiary(&_Marketplace.CallOpts)
}

// NetworkBeneficiary is a free data retrieval call binding the contract method 0x6b534ed0.
//
// Solidity: function networkBeneficiary() view returns(address)
func (_Marketplace *MarketplaceCallerSession) NetworkBeneficiary() (common.Address, error) {
	return _Marketplace.Contract.NetworkBeneficiary(&_Marketplace.CallOpts)
}

// OfferCancelationDelay is a free data retrieval call binding the contract method 0x0a5c4ed5.
//
// Solidity: function offerCancelationDelay() view returns(uint256)
func (_Marketplace *MarketplaceCaller) OfferCancelationDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "offerCancelationDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfferCancelationDelay is a free data retrieval call binding the contract method 0x0a5c4ed5.
//
// Solidity: function offerCancelationDelay() view returns(uint256)
func (_Marketplace *MarketplaceSession) OfferCancelationDelay() (*big.Int, error) {
	return _Marketplace.Contract.OfferCancelationDelay(&_Marketplace.CallOpts)
}

// OfferCancelationDelay is a free data retrieval call binding the contract method 0x0a5c4ed5.
//
// Solidity: function offerCancelationDelay() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) OfferCancelationDelay() (*big.Int, error) {
	return _Marketplace.Contract.OfferCancelationDelay(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0xa6d23e10.
//
// Solidity: function payments() view returns(address)
func (_Marketplace *MarketplaceCaller) Payments(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "payments")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Payments is a free data retrieval call binding the contract method 0xa6d23e10.
//
// Solidity: function payments() view returns(address)
func (_Marketplace *MarketplaceSession) Payments() (common.Address, error) {
	return _Marketplace.Contract.Payments(&_Marketplace.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0xa6d23e10.
//
// Solidity: function payments() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Payments() (common.Address, error) {
	return _Marketplace.Contract.Payments(&_Marketplace.CallOpts)
}

// RoyaltyEngine is a free data retrieval call binding the contract method 0x4c94c90c.
//
// Solidity: function royaltyEngine() view returns(address)
func (_Marketplace *MarketplaceCaller) RoyaltyEngine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "royaltyEngine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyEngine is a free data retrieval call binding the contract method 0x4c94c90c.
//
// Solidity: function royaltyEngine() view returns(address)
func (_Marketplace *MarketplaceSession) RoyaltyEngine() (common.Address, error) {
	return _Marketplace.Contract.RoyaltyEngine(&_Marketplace.CallOpts)
}

// RoyaltyEngine is a free data retrieval call binding the contract method 0x4c94c90c.
//
// Solidity: function royaltyEngine() view returns(address)
func (_Marketplace *MarketplaceCallerSession) RoyaltyEngine() (common.Address, error) {
	return _Marketplace.Contract.RoyaltyEngine(&_Marketplace.CallOpts)
}

// RoyaltyRegistry is a free data retrieval call binding the contract method 0xa11b0712.
//
// Solidity: function royaltyRegistry() view returns(address)
func (_Marketplace *MarketplaceCaller) RoyaltyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "royaltyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyRegistry is a free data retrieval call binding the contract method 0xa11b0712.
//
// Solidity: function royaltyRegistry() view returns(address)
func (_Marketplace *MarketplaceSession) RoyaltyRegistry() (common.Address, error) {
	return _Marketplace.Contract.RoyaltyRegistry(&_Marketplace.CallOpts)
}

// RoyaltyRegistry is a free data retrieval call binding the contract method 0xa11b0712.
//
// Solidity: function royaltyRegistry() view returns(address)
func (_Marketplace *MarketplaceCallerSession) RoyaltyRegistry() (common.Address, error) {
	return _Marketplace.Contract.RoyaltyRegistry(&_Marketplace.CallOpts)
}

// SpaceOperatorRegistry is a free data retrieval call binding the contract method 0x0bcba09d.
//
// Solidity: function spaceOperatorRegistry() view returns(address)
func (_Marketplace *MarketplaceCaller) SpaceOperatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "spaceOperatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpaceOperatorRegistry is a free data retrieval call binding the contract method 0x0bcba09d.
//
// Solidity: function spaceOperatorRegistry() view returns(address)
func (_Marketplace *MarketplaceSession) SpaceOperatorRegistry() (common.Address, error) {
	return _Marketplace.Contract.SpaceOperatorRegistry(&_Marketplace.CallOpts)
}

// SpaceOperatorRegistry is a free data retrieval call binding the contract method 0x0bcba09d.
//
// Solidity: function spaceOperatorRegistry() view returns(address)
func (_Marketplace *MarketplaceCallerSession) SpaceOperatorRegistry() (common.Address, error) {
	return _Marketplace.Contract.SpaceOperatorRegistry(&_Marketplace.CallOpts)
}

// StakingRegistry is a free data retrieval call binding the contract method 0x009d9aa9.
//
// Solidity: function stakingRegistry() view returns(address)
func (_Marketplace *MarketplaceCaller) StakingRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "stakingRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingRegistry is a free data retrieval call binding the contract method 0x009d9aa9.
//
// Solidity: function stakingRegistry() view returns(address)
func (_Marketplace *MarketplaceSession) StakingRegistry() (common.Address, error) {
	return _Marketplace.Contract.StakingRegistry(&_Marketplace.CallOpts)
}

// StakingRegistry is a free data retrieval call binding the contract method 0x009d9aa9.
//
// Solidity: function stakingRegistry() view returns(address)
func (_Marketplace *MarketplaceCallerSession) StakingRegistry() (common.Address, error) {
	return _Marketplace.Contract.StakingRegistry(&_Marketplace.CallOpts)
}

// SuperRareAuctionHouse is a free data retrieval call binding the contract method 0x3bc3d9be.
//
// Solidity: function superRareAuctionHouse() view returns(address)
func (_Marketplace *MarketplaceCaller) SuperRareAuctionHouse(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "superRareAuctionHouse")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperRareAuctionHouse is a free data retrieval call binding the contract method 0x3bc3d9be.
//
// Solidity: function superRareAuctionHouse() view returns(address)
func (_Marketplace *MarketplaceSession) SuperRareAuctionHouse() (common.Address, error) {
	return _Marketplace.Contract.SuperRareAuctionHouse(&_Marketplace.CallOpts)
}

// SuperRareAuctionHouse is a free data retrieval call binding the contract method 0x3bc3d9be.
//
// Solidity: function superRareAuctionHouse() view returns(address)
func (_Marketplace *MarketplaceCallerSession) SuperRareAuctionHouse() (common.Address, error) {
	return _Marketplace.Contract.SuperRareAuctionHouse(&_Marketplace.CallOpts)
}

// SuperRareMarketplace is a free data retrieval call binding the contract method 0x0141c590.
//
// Solidity: function superRareMarketplace() view returns(address)
func (_Marketplace *MarketplaceCaller) SuperRareMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "superRareMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperRareMarketplace is a free data retrieval call binding the contract method 0x0141c590.
//
// Solidity: function superRareMarketplace() view returns(address)
func (_Marketplace *MarketplaceSession) SuperRareMarketplace() (common.Address, error) {
	return _Marketplace.Contract.SuperRareMarketplace(&_Marketplace.CallOpts)
}

// SuperRareMarketplace is a free data retrieval call binding the contract method 0x0141c590.
//
// Solidity: function superRareMarketplace() view returns(address)
func (_Marketplace *MarketplaceCallerSession) SuperRareMarketplace() (common.Address, error) {
	return _Marketplace.Contract.SuperRareMarketplace(&_Marketplace.CallOpts)
}

// TokenAuctions is a free data retrieval call binding the contract method 0xc47c35c1.
//
// Solidity: function tokenAuctions(address , uint256 ) view returns(address auctionCreator, uint256 creationBlock, uint256 startingTime, uint256 lengthOfAuction, address currencyAddress, uint256 minimumBid, bytes32 auctionType)
func (_Marketplace *MarketplaceCaller) TokenAuctions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	AuctionCreator  common.Address
	CreationBlock   *big.Int
	StartingTime    *big.Int
	LengthOfAuction *big.Int
	CurrencyAddress common.Address
	MinimumBid      *big.Int
	AuctionType     [32]byte
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "tokenAuctions", arg0, arg1)

	outstruct := new(struct {
		AuctionCreator  common.Address
		CreationBlock   *big.Int
		StartingTime    *big.Int
		LengthOfAuction *big.Int
		CurrencyAddress common.Address
		MinimumBid      *big.Int
		AuctionType     [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AuctionCreator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CreationBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartingTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LengthOfAuction = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CurrencyAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.MinimumBid = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.AuctionType = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// TokenAuctions is a free data retrieval call binding the contract method 0xc47c35c1.
//
// Solidity: function tokenAuctions(address , uint256 ) view returns(address auctionCreator, uint256 creationBlock, uint256 startingTime, uint256 lengthOfAuction, address currencyAddress, uint256 minimumBid, bytes32 auctionType)
func (_Marketplace *MarketplaceSession) TokenAuctions(arg0 common.Address, arg1 *big.Int) (struct {
	AuctionCreator  common.Address
	CreationBlock   *big.Int
	StartingTime    *big.Int
	LengthOfAuction *big.Int
	CurrencyAddress common.Address
	MinimumBid      *big.Int
	AuctionType     [32]byte
}, error) {
	return _Marketplace.Contract.TokenAuctions(&_Marketplace.CallOpts, arg0, arg1)
}

// TokenAuctions is a free data retrieval call binding the contract method 0xc47c35c1.
//
// Solidity: function tokenAuctions(address , uint256 ) view returns(address auctionCreator, uint256 creationBlock, uint256 startingTime, uint256 lengthOfAuction, address currencyAddress, uint256 minimumBid, bytes32 auctionType)
func (_Marketplace *MarketplaceCallerSession) TokenAuctions(arg0 common.Address, arg1 *big.Int) (struct {
	AuctionCreator  common.Address
	CreationBlock   *big.Int
	StartingTime    *big.Int
	LengthOfAuction *big.Int
	CurrencyAddress common.Address
	MinimumBid      *big.Int
	AuctionType     [32]byte
}, error) {
	return _Marketplace.Contract.TokenAuctions(&_Marketplace.CallOpts, arg0, arg1)
}

// TokenCurrentOffers is a free data retrieval call binding the contract method 0x2c419053.
//
// Solidity: function tokenCurrentOffers(address , uint256 , address ) view returns(address buyer, uint256 amount, uint256 timestamp, uint8 marketplaceFee, bool convertible)
func (_Marketplace *MarketplaceCaller) TokenCurrentOffers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Buyer          common.Address
	Amount         *big.Int
	Timestamp      *big.Int
	MarketplaceFee uint8
	Convertible    bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "tokenCurrentOffers", arg0, arg1, arg2)

	outstruct := new(struct {
		Buyer          common.Address
		Amount         *big.Int
		Timestamp      *big.Int
		MarketplaceFee uint8
		Convertible    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MarketplaceFee = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Convertible = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// TokenCurrentOffers is a free data retrieval call binding the contract method 0x2c419053.
//
// Solidity: function tokenCurrentOffers(address , uint256 , address ) view returns(address buyer, uint256 amount, uint256 timestamp, uint8 marketplaceFee, bool convertible)
func (_Marketplace *MarketplaceSession) TokenCurrentOffers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Buyer          common.Address
	Amount         *big.Int
	Timestamp      *big.Int
	MarketplaceFee uint8
	Convertible    bool
}, error) {
	return _Marketplace.Contract.TokenCurrentOffers(&_Marketplace.CallOpts, arg0, arg1, arg2)
}

// TokenCurrentOffers is a free data retrieval call binding the contract method 0x2c419053.
//
// Solidity: function tokenCurrentOffers(address , uint256 , address ) view returns(address buyer, uint256 amount, uint256 timestamp, uint8 marketplaceFee, bool convertible)
func (_Marketplace *MarketplaceCallerSession) TokenCurrentOffers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Buyer          common.Address
	Amount         *big.Int
	Timestamp      *big.Int
	MarketplaceFee uint8
	Convertible    bool
}, error) {
	return _Marketplace.Contract.TokenCurrentOffers(&_Marketplace.CallOpts, arg0, arg1, arg2)
}

// TokenSalePrices is a free data retrieval call binding the contract method 0x1a2ac30f.
//
// Solidity: function tokenSalePrices(address , uint256 , address ) view returns(address seller, address currencyAddress, uint256 amount)
func (_Marketplace *MarketplaceCaller) TokenSalePrices(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Seller          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "tokenSalePrices", arg0, arg1, arg2)

	outstruct := new(struct {
		Seller          common.Address
		CurrencyAddress common.Address
		Amount          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CurrencyAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenSalePrices is a free data retrieval call binding the contract method 0x1a2ac30f.
//
// Solidity: function tokenSalePrices(address , uint256 , address ) view returns(address seller, address currencyAddress, uint256 amount)
func (_Marketplace *MarketplaceSession) TokenSalePrices(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Seller          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
}, error) {
	return _Marketplace.Contract.TokenSalePrices(&_Marketplace.CallOpts, arg0, arg1, arg2)
}

// TokenSalePrices is a free data retrieval call binding the contract method 0x1a2ac30f.
//
// Solidity: function tokenSalePrices(address , uint256 , address ) view returns(address seller, address currencyAddress, uint256 amount)
func (_Marketplace *MarketplaceCallerSession) TokenSalePrices(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Seller          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
}, error) {
	return _Marketplace.Contract.TokenSalePrices(&_Marketplace.CallOpts, arg0, arg1, arg2)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x534665e9.
//
// Solidity: function acceptOffer(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceTransactor) AcceptOffer(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "acceptOffer", _originContract, _tokenId, _currencyAddress, _amount, _splitAddresses, _splitRatios)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x534665e9.
//
// Solidity: function acceptOffer(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceSession) AcceptOffer(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.AcceptOffer(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount, _splitAddresses, _splitRatios)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x534665e9.
//
// Solidity: function acceptOffer(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceTransactorSession) AcceptOffer(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.AcceptOffer(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount, _splitAddresses, _splitRatios)
}

// Bid is a paid mutator transaction binding the contract method 0xb5678588.
//
// Solidity: function bid(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount) payable returns()
func (_Marketplace *MarketplaceTransactor) Bid(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "bid", _originContract, _tokenId, _currencyAddress, _amount)
}

// Bid is a paid mutator transaction binding the contract method 0xb5678588.
//
// Solidity: function bid(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount) payable returns()
func (_Marketplace *MarketplaceSession) Bid(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Bid(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount)
}

// Bid is a paid mutator transaction binding the contract method 0xb5678588.
//
// Solidity: function bid(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount) payable returns()
func (_Marketplace *MarketplaceTransactorSession) Bid(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Bid(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0xb3ffb760.
//
// Solidity: function buy(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount) payable returns()
func (_Marketplace *MarketplaceTransactor) Buy(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buy", _originContract, _tokenId, _currencyAddress, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0xb3ffb760.
//
// Solidity: function buy(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount) payable returns()
func (_Marketplace *MarketplaceSession) Buy(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0xb3ffb760.
//
// Solidity: function buy(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount) payable returns()
func (_Marketplace *MarketplaceTransactorSession) Buy(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _originContract, uint256 _tokenId) returns()
func (_Marketplace *MarketplaceTransactor) CancelAuction(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelAuction", _originContract, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _originContract, uint256 _tokenId) returns()
func (_Marketplace *MarketplaceSession) CancelAuction(_originContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelAuction(&_Marketplace.TransactOpts, _originContract, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _originContract, uint256 _tokenId) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelAuction(_originContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelAuction(&_Marketplace.TransactOpts, _originContract, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xe92f94d1.
//
// Solidity: function cancelOffer(address _originContract, uint256 _tokenId, address _currencyAddress) returns()
func (_Marketplace *MarketplaceTransactor) CancelOffer(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelOffer", _originContract, _tokenId, _currencyAddress)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xe92f94d1.
//
// Solidity: function cancelOffer(address _originContract, uint256 _tokenId, address _currencyAddress) returns()
func (_Marketplace *MarketplaceSession) CancelOffer(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelOffer(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xe92f94d1.
//
// Solidity: function cancelOffer(address _originContract, uint256 _tokenId, address _currencyAddress) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelOffer(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelOffer(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress)
}

// ConfigureAuction is a paid mutator transaction binding the contract method 0x9041a0ec.
//
// Solidity: function configureAuction(bytes32 _auctionType, address _originContract, uint256 _tokenId, uint256 _startingAmount, address _currencyAddress, uint256 _lengthOfAuction, uint256 _startTime, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceTransactor) ConfigureAuction(opts *bind.TransactOpts, _auctionType [32]byte, _originContract common.Address, _tokenId *big.Int, _startingAmount *big.Int, _currencyAddress common.Address, _lengthOfAuction *big.Int, _startTime *big.Int, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "configureAuction", _auctionType, _originContract, _tokenId, _startingAmount, _currencyAddress, _lengthOfAuction, _startTime, _splitAddresses, _splitRatios)
}

// ConfigureAuction is a paid mutator transaction binding the contract method 0x9041a0ec.
//
// Solidity: function configureAuction(bytes32 _auctionType, address _originContract, uint256 _tokenId, uint256 _startingAmount, address _currencyAddress, uint256 _lengthOfAuction, uint256 _startTime, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceSession) ConfigureAuction(_auctionType [32]byte, _originContract common.Address, _tokenId *big.Int, _startingAmount *big.Int, _currencyAddress common.Address, _lengthOfAuction *big.Int, _startTime *big.Int, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.ConfigureAuction(&_Marketplace.TransactOpts, _auctionType, _originContract, _tokenId, _startingAmount, _currencyAddress, _lengthOfAuction, _startTime, _splitAddresses, _splitRatios)
}

// ConfigureAuction is a paid mutator transaction binding the contract method 0x9041a0ec.
//
// Solidity: function configureAuction(bytes32 _auctionType, address _originContract, uint256 _tokenId, uint256 _startingAmount, address _currencyAddress, uint256 _lengthOfAuction, uint256 _startTime, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceTransactorSession) ConfigureAuction(_auctionType [32]byte, _originContract common.Address, _tokenId *big.Int, _startingAmount *big.Int, _currencyAddress common.Address, _lengthOfAuction *big.Int, _startTime *big.Int, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.ConfigureAuction(&_Marketplace.TransactOpts, _auctionType, _originContract, _tokenId, _startingAmount, _currencyAddress, _lengthOfAuction, _startTime, _splitAddresses, _splitRatios)
}

// ConvertOfferToAuction is a paid mutator transaction binding the contract method 0x060d9eeb.
//
// Solidity: function convertOfferToAuction(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount, uint256 _lengthOfAuction, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceTransactor) ConvertOfferToAuction(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int, _lengthOfAuction *big.Int, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "convertOfferToAuction", _originContract, _tokenId, _currencyAddress, _amount, _lengthOfAuction, _splitAddresses, _splitRatios)
}

// ConvertOfferToAuction is a paid mutator transaction binding the contract method 0x060d9eeb.
//
// Solidity: function convertOfferToAuction(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount, uint256 _lengthOfAuction, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceSession) ConvertOfferToAuction(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int, _lengthOfAuction *big.Int, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.ConvertOfferToAuction(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount, _lengthOfAuction, _splitAddresses, _splitRatios)
}

// ConvertOfferToAuction is a paid mutator transaction binding the contract method 0x060d9eeb.
//
// Solidity: function convertOfferToAuction(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount, uint256 _lengthOfAuction, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceTransactorSession) ConvertOfferToAuction(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int, _lengthOfAuction *big.Int, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.ConvertOfferToAuction(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount, _lengthOfAuction, _splitAddresses, _splitRatios)
}

// Initialize is a paid mutator transaction binding the contract method 0xc306b378.
//
// Solidity: function initialize(address _marketplaceSettings, address _royaltyRegistry, address _royaltyEngine, address _superRareMarketplace, address _superRareAuctionHouse, address _spaceOperatorRegistry, address _approvedTokenRegistry, address _payments, address _stakingRegistry, address _networkBeneficiary) returns()
func (_Marketplace *MarketplaceTransactor) Initialize(opts *bind.TransactOpts, _marketplaceSettings common.Address, _royaltyRegistry common.Address, _royaltyEngine common.Address, _superRareMarketplace common.Address, _superRareAuctionHouse common.Address, _spaceOperatorRegistry common.Address, _approvedTokenRegistry common.Address, _payments common.Address, _stakingRegistry common.Address, _networkBeneficiary common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "initialize", _marketplaceSettings, _royaltyRegistry, _royaltyEngine, _superRareMarketplace, _superRareAuctionHouse, _spaceOperatorRegistry, _approvedTokenRegistry, _payments, _stakingRegistry, _networkBeneficiary)
}

// Initialize is a paid mutator transaction binding the contract method 0xc306b378.
//
// Solidity: function initialize(address _marketplaceSettings, address _royaltyRegistry, address _royaltyEngine, address _superRareMarketplace, address _superRareAuctionHouse, address _spaceOperatorRegistry, address _approvedTokenRegistry, address _payments, address _stakingRegistry, address _networkBeneficiary) returns()
func (_Marketplace *MarketplaceSession) Initialize(_marketplaceSettings common.Address, _royaltyRegistry common.Address, _royaltyEngine common.Address, _superRareMarketplace common.Address, _superRareAuctionHouse common.Address, _spaceOperatorRegistry common.Address, _approvedTokenRegistry common.Address, _payments common.Address, _stakingRegistry common.Address, _networkBeneficiary common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _marketplaceSettings, _royaltyRegistry, _royaltyEngine, _superRareMarketplace, _superRareAuctionHouse, _spaceOperatorRegistry, _approvedTokenRegistry, _payments, _stakingRegistry, _networkBeneficiary)
}

// Initialize is a paid mutator transaction binding the contract method 0xc306b378.
//
// Solidity: function initialize(address _marketplaceSettings, address _royaltyRegistry, address _royaltyEngine, address _superRareMarketplace, address _superRareAuctionHouse, address _spaceOperatorRegistry, address _approvedTokenRegistry, address _payments, address _stakingRegistry, address _networkBeneficiary) returns()
func (_Marketplace *MarketplaceTransactorSession) Initialize(_marketplaceSettings common.Address, _royaltyRegistry common.Address, _royaltyEngine common.Address, _superRareMarketplace common.Address, _superRareAuctionHouse common.Address, _spaceOperatorRegistry common.Address, _approvedTokenRegistry common.Address, _payments common.Address, _stakingRegistry common.Address, _networkBeneficiary common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _marketplaceSettings, _royaltyRegistry, _royaltyEngine, _superRareMarketplace, _superRareAuctionHouse, _spaceOperatorRegistry, _approvedTokenRegistry, _payments, _stakingRegistry, _networkBeneficiary)
}

// Offer is a paid mutator transaction binding the contract method 0x0f2b2532.
//
// Solidity: function offer(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount, bool _convertible) payable returns()
func (_Marketplace *MarketplaceTransactor) Offer(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int, _convertible bool) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "offer", _originContract, _tokenId, _currencyAddress, _amount, _convertible)
}

// Offer is a paid mutator transaction binding the contract method 0x0f2b2532.
//
// Solidity: function offer(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount, bool _convertible) payable returns()
func (_Marketplace *MarketplaceSession) Offer(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int, _convertible bool) (*types.Transaction, error) {
	return _Marketplace.Contract.Offer(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount, _convertible)
}

// Offer is a paid mutator transaction binding the contract method 0x0f2b2532.
//
// Solidity: function offer(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _amount, bool _convertible) payable returns()
func (_Marketplace *MarketplaceTransactorSession) Offer(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _amount *big.Int, _convertible bool) (*types.Transaction, error) {
	return _Marketplace.Contract.Offer(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _amount, _convertible)
}

// RemoveSalePrice is a paid mutator transaction binding the contract method 0xf7cfaad0.
//
// Solidity: function removeSalePrice(address _originContract, uint256 _tokenId, address _target) returns()
func (_Marketplace *MarketplaceTransactor) RemoveSalePrice(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int, _target common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "removeSalePrice", _originContract, _tokenId, _target)
}

// RemoveSalePrice is a paid mutator transaction binding the contract method 0xf7cfaad0.
//
// Solidity: function removeSalePrice(address _originContract, uint256 _tokenId, address _target) returns()
func (_Marketplace *MarketplaceSession) RemoveSalePrice(_originContract common.Address, _tokenId *big.Int, _target common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveSalePrice(&_Marketplace.TransactOpts, _originContract, _tokenId, _target)
}

// RemoveSalePrice is a paid mutator transaction binding the contract method 0xf7cfaad0.
//
// Solidity: function removeSalePrice(address _originContract, uint256 _tokenId, address _target) returns()
func (_Marketplace *MarketplaceTransactorSession) RemoveSalePrice(_originContract common.Address, _tokenId *big.Int, _target common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveSalePrice(&_Marketplace.TransactOpts, _originContract, _tokenId, _target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// SetApprovedTokenRegistry is a paid mutator transaction binding the contract method 0xe4e87e3b.
//
// Solidity: function setApprovedTokenRegistry(address _approvedTokenRegistry) returns()
func (_Marketplace *MarketplaceTransactor) SetApprovedTokenRegistry(opts *bind.TransactOpts, _approvedTokenRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setApprovedTokenRegistry", _approvedTokenRegistry)
}

// SetApprovedTokenRegistry is a paid mutator transaction binding the contract method 0xe4e87e3b.
//
// Solidity: function setApprovedTokenRegistry(address _approvedTokenRegistry) returns()
func (_Marketplace *MarketplaceSession) SetApprovedTokenRegistry(_approvedTokenRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetApprovedTokenRegistry(&_Marketplace.TransactOpts, _approvedTokenRegistry)
}

// SetApprovedTokenRegistry is a paid mutator transaction binding the contract method 0xe4e87e3b.
//
// Solidity: function setApprovedTokenRegistry(address _approvedTokenRegistry) returns()
func (_Marketplace *MarketplaceTransactorSession) SetApprovedTokenRegistry(_approvedTokenRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetApprovedTokenRegistry(&_Marketplace.TransactOpts, _approvedTokenRegistry)
}

// SetAuctionLengthExtension is a paid mutator transaction binding the contract method 0x10f79789.
//
// Solidity: function setAuctionLengthExtension(uint256 _auctionLengthExtension) returns()
func (_Marketplace *MarketplaceTransactor) SetAuctionLengthExtension(opts *bind.TransactOpts, _auctionLengthExtension *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setAuctionLengthExtension", _auctionLengthExtension)
}

// SetAuctionLengthExtension is a paid mutator transaction binding the contract method 0x10f79789.
//
// Solidity: function setAuctionLengthExtension(uint256 _auctionLengthExtension) returns()
func (_Marketplace *MarketplaceSession) SetAuctionLengthExtension(_auctionLengthExtension *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetAuctionLengthExtension(&_Marketplace.TransactOpts, _auctionLengthExtension)
}

// SetAuctionLengthExtension is a paid mutator transaction binding the contract method 0x10f79789.
//
// Solidity: function setAuctionLengthExtension(uint256 _auctionLengthExtension) returns()
func (_Marketplace *MarketplaceTransactorSession) SetAuctionLengthExtension(_auctionLengthExtension *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetAuctionLengthExtension(&_Marketplace.TransactOpts, _auctionLengthExtension)
}

// SetMarketplaceSettings is a paid mutator transaction binding the contract method 0x176ab440.
//
// Solidity: function setMarketplaceSettings(address _marketplaceSettings) returns()
func (_Marketplace *MarketplaceTransactor) SetMarketplaceSettings(opts *bind.TransactOpts, _marketplaceSettings common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setMarketplaceSettings", _marketplaceSettings)
}

// SetMarketplaceSettings is a paid mutator transaction binding the contract method 0x176ab440.
//
// Solidity: function setMarketplaceSettings(address _marketplaceSettings) returns()
func (_Marketplace *MarketplaceSession) SetMarketplaceSettings(_marketplaceSettings common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMarketplaceSettings(&_Marketplace.TransactOpts, _marketplaceSettings)
}

// SetMarketplaceSettings is a paid mutator transaction binding the contract method 0x176ab440.
//
// Solidity: function setMarketplaceSettings(address _marketplaceSettings) returns()
func (_Marketplace *MarketplaceTransactorSession) SetMarketplaceSettings(_marketplaceSettings common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMarketplaceSettings(&_Marketplace.TransactOpts, _marketplaceSettings)
}

// SetMaxAuctionLength is a paid mutator transaction binding the contract method 0xdce96bf5.
//
// Solidity: function setMaxAuctionLength(uint8 _maxAuctionLength) returns()
func (_Marketplace *MarketplaceTransactor) SetMaxAuctionLength(opts *bind.TransactOpts, _maxAuctionLength uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setMaxAuctionLength", _maxAuctionLength)
}

// SetMaxAuctionLength is a paid mutator transaction binding the contract method 0xdce96bf5.
//
// Solidity: function setMaxAuctionLength(uint8 _maxAuctionLength) returns()
func (_Marketplace *MarketplaceSession) SetMaxAuctionLength(_maxAuctionLength uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMaxAuctionLength(&_Marketplace.TransactOpts, _maxAuctionLength)
}

// SetMaxAuctionLength is a paid mutator transaction binding the contract method 0xdce96bf5.
//
// Solidity: function setMaxAuctionLength(uint8 _maxAuctionLength) returns()
func (_Marketplace *MarketplaceTransactorSession) SetMaxAuctionLength(_maxAuctionLength uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMaxAuctionLength(&_Marketplace.TransactOpts, _maxAuctionLength)
}

// SetMinimumBidIncreasePercentage is a paid mutator transaction binding the contract method 0x48626b90.
//
// Solidity: function setMinimumBidIncreasePercentage(uint8 _minimumBidIncreasePercentage) returns()
func (_Marketplace *MarketplaceTransactor) SetMinimumBidIncreasePercentage(opts *bind.TransactOpts, _minimumBidIncreasePercentage uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setMinimumBidIncreasePercentage", _minimumBidIncreasePercentage)
}

// SetMinimumBidIncreasePercentage is a paid mutator transaction binding the contract method 0x48626b90.
//
// Solidity: function setMinimumBidIncreasePercentage(uint8 _minimumBidIncreasePercentage) returns()
func (_Marketplace *MarketplaceSession) SetMinimumBidIncreasePercentage(_minimumBidIncreasePercentage uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMinimumBidIncreasePercentage(&_Marketplace.TransactOpts, _minimumBidIncreasePercentage)
}

// SetMinimumBidIncreasePercentage is a paid mutator transaction binding the contract method 0x48626b90.
//
// Solidity: function setMinimumBidIncreasePercentage(uint8 _minimumBidIncreasePercentage) returns()
func (_Marketplace *MarketplaceTransactorSession) SetMinimumBidIncreasePercentage(_minimumBidIncreasePercentage uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMinimumBidIncreasePercentage(&_Marketplace.TransactOpts, _minimumBidIncreasePercentage)
}

// SetNetworkBeneficiary is a paid mutator transaction binding the contract method 0x3492e5a8.
//
// Solidity: function setNetworkBeneficiary(address _networkBeneficiary) returns()
func (_Marketplace *MarketplaceTransactor) SetNetworkBeneficiary(opts *bind.TransactOpts, _networkBeneficiary common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setNetworkBeneficiary", _networkBeneficiary)
}

// SetNetworkBeneficiary is a paid mutator transaction binding the contract method 0x3492e5a8.
//
// Solidity: function setNetworkBeneficiary(address _networkBeneficiary) returns()
func (_Marketplace *MarketplaceSession) SetNetworkBeneficiary(_networkBeneficiary common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetNetworkBeneficiary(&_Marketplace.TransactOpts, _networkBeneficiary)
}

// SetNetworkBeneficiary is a paid mutator transaction binding the contract method 0x3492e5a8.
//
// Solidity: function setNetworkBeneficiary(address _networkBeneficiary) returns()
func (_Marketplace *MarketplaceTransactorSession) SetNetworkBeneficiary(_networkBeneficiary common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetNetworkBeneficiary(&_Marketplace.TransactOpts, _networkBeneficiary)
}

// SetOfferCancelationDelay is a paid mutator transaction binding the contract method 0x7a544792.
//
// Solidity: function setOfferCancelationDelay(uint256 _offerCancelationDelay) returns()
func (_Marketplace *MarketplaceTransactor) SetOfferCancelationDelay(opts *bind.TransactOpts, _offerCancelationDelay *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setOfferCancelationDelay", _offerCancelationDelay)
}

// SetOfferCancelationDelay is a paid mutator transaction binding the contract method 0x7a544792.
//
// Solidity: function setOfferCancelationDelay(uint256 _offerCancelationDelay) returns()
func (_Marketplace *MarketplaceSession) SetOfferCancelationDelay(_offerCancelationDelay *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetOfferCancelationDelay(&_Marketplace.TransactOpts, _offerCancelationDelay)
}

// SetOfferCancelationDelay is a paid mutator transaction binding the contract method 0x7a544792.
//
// Solidity: function setOfferCancelationDelay(uint256 _offerCancelationDelay) returns()
func (_Marketplace *MarketplaceTransactorSession) SetOfferCancelationDelay(_offerCancelationDelay *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetOfferCancelationDelay(&_Marketplace.TransactOpts, _offerCancelationDelay)
}

// SetPayments is a paid mutator transaction binding the contract method 0xaf231a58.
//
// Solidity: function setPayments(address _payments) returns()
func (_Marketplace *MarketplaceTransactor) SetPayments(opts *bind.TransactOpts, _payments common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setPayments", _payments)
}

// SetPayments is a paid mutator transaction binding the contract method 0xaf231a58.
//
// Solidity: function setPayments(address _payments) returns()
func (_Marketplace *MarketplaceSession) SetPayments(_payments common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPayments(&_Marketplace.TransactOpts, _payments)
}

// SetPayments is a paid mutator transaction binding the contract method 0xaf231a58.
//
// Solidity: function setPayments(address _payments) returns()
func (_Marketplace *MarketplaceTransactorSession) SetPayments(_payments common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPayments(&_Marketplace.TransactOpts, _payments)
}

// SetRoyaltyEngine is a paid mutator transaction binding the contract method 0x21ede032.
//
// Solidity: function setRoyaltyEngine(address _royaltyEngine) returns()
func (_Marketplace *MarketplaceTransactor) SetRoyaltyEngine(opts *bind.TransactOpts, _royaltyEngine common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setRoyaltyEngine", _royaltyEngine)
}

// SetRoyaltyEngine is a paid mutator transaction binding the contract method 0x21ede032.
//
// Solidity: function setRoyaltyEngine(address _royaltyEngine) returns()
func (_Marketplace *MarketplaceSession) SetRoyaltyEngine(_royaltyEngine common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRoyaltyEngine(&_Marketplace.TransactOpts, _royaltyEngine)
}

// SetRoyaltyEngine is a paid mutator transaction binding the contract method 0x21ede032.
//
// Solidity: function setRoyaltyEngine(address _royaltyEngine) returns()
func (_Marketplace *MarketplaceTransactorSession) SetRoyaltyEngine(_royaltyEngine common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRoyaltyEngine(&_Marketplace.TransactOpts, _royaltyEngine)
}

// SetRoyaltyRegistry is a paid mutator transaction binding the contract method 0x84a608e2.
//
// Solidity: function setRoyaltyRegistry(address _royaltyRegistry) returns()
func (_Marketplace *MarketplaceTransactor) SetRoyaltyRegistry(opts *bind.TransactOpts, _royaltyRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setRoyaltyRegistry", _royaltyRegistry)
}

// SetRoyaltyRegistry is a paid mutator transaction binding the contract method 0x84a608e2.
//
// Solidity: function setRoyaltyRegistry(address _royaltyRegistry) returns()
func (_Marketplace *MarketplaceSession) SetRoyaltyRegistry(_royaltyRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRoyaltyRegistry(&_Marketplace.TransactOpts, _royaltyRegistry)
}

// SetRoyaltyRegistry is a paid mutator transaction binding the contract method 0x84a608e2.
//
// Solidity: function setRoyaltyRegistry(address _royaltyRegistry) returns()
func (_Marketplace *MarketplaceTransactorSession) SetRoyaltyRegistry(_royaltyRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRoyaltyRegistry(&_Marketplace.TransactOpts, _royaltyRegistry)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0xc8f94f4e.
//
// Solidity: function setSalePrice(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _listPrice, address _target, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceTransactor) SetSalePrice(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _listPrice *big.Int, _target common.Address, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setSalePrice", _originContract, _tokenId, _currencyAddress, _listPrice, _target, _splitAddresses, _splitRatios)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0xc8f94f4e.
//
// Solidity: function setSalePrice(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _listPrice, address _target, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceSession) SetSalePrice(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _listPrice *big.Int, _target common.Address, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.SetSalePrice(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _listPrice, _target, _splitAddresses, _splitRatios)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0xc8f94f4e.
//
// Solidity: function setSalePrice(address _originContract, uint256 _tokenId, address _currencyAddress, uint256 _listPrice, address _target, address[] _splitAddresses, uint8[] _splitRatios) returns()
func (_Marketplace *MarketplaceTransactorSession) SetSalePrice(_originContract common.Address, _tokenId *big.Int, _currencyAddress common.Address, _listPrice *big.Int, _target common.Address, _splitAddresses []common.Address, _splitRatios []uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.SetSalePrice(&_Marketplace.TransactOpts, _originContract, _tokenId, _currencyAddress, _listPrice, _target, _splitAddresses, _splitRatios)
}

// SetSpaceOperatorRegistry is a paid mutator transaction binding the contract method 0x7f358230.
//
// Solidity: function setSpaceOperatorRegistry(address _spaceOperatorRegistry) returns()
func (_Marketplace *MarketplaceTransactor) SetSpaceOperatorRegistry(opts *bind.TransactOpts, _spaceOperatorRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setSpaceOperatorRegistry", _spaceOperatorRegistry)
}

// SetSpaceOperatorRegistry is a paid mutator transaction binding the contract method 0x7f358230.
//
// Solidity: function setSpaceOperatorRegistry(address _spaceOperatorRegistry) returns()
func (_Marketplace *MarketplaceSession) SetSpaceOperatorRegistry(_spaceOperatorRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetSpaceOperatorRegistry(&_Marketplace.TransactOpts, _spaceOperatorRegistry)
}

// SetSpaceOperatorRegistry is a paid mutator transaction binding the contract method 0x7f358230.
//
// Solidity: function setSpaceOperatorRegistry(address _spaceOperatorRegistry) returns()
func (_Marketplace *MarketplaceTransactorSession) SetSpaceOperatorRegistry(_spaceOperatorRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetSpaceOperatorRegistry(&_Marketplace.TransactOpts, _spaceOperatorRegistry)
}

// SetStakingRegistry is a paid mutator transaction binding the contract method 0x2c740844.
//
// Solidity: function setStakingRegistry(address _stakingRegistry) returns()
func (_Marketplace *MarketplaceTransactor) SetStakingRegistry(opts *bind.TransactOpts, _stakingRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setStakingRegistry", _stakingRegistry)
}

// SetStakingRegistry is a paid mutator transaction binding the contract method 0x2c740844.
//
// Solidity: function setStakingRegistry(address _stakingRegistry) returns()
func (_Marketplace *MarketplaceSession) SetStakingRegistry(_stakingRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetStakingRegistry(&_Marketplace.TransactOpts, _stakingRegistry)
}

// SetStakingRegistry is a paid mutator transaction binding the contract method 0x2c740844.
//
// Solidity: function setStakingRegistry(address _stakingRegistry) returns()
func (_Marketplace *MarketplaceTransactorSession) SetStakingRegistry(_stakingRegistry common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetStakingRegistry(&_Marketplace.TransactOpts, _stakingRegistry)
}

// SetSuperRareAuctionHouse is a paid mutator transaction binding the contract method 0x2a2a326c.
//
// Solidity: function setSuperRareAuctionHouse(address _superRareAuctionHouse) returns()
func (_Marketplace *MarketplaceTransactor) SetSuperRareAuctionHouse(opts *bind.TransactOpts, _superRareAuctionHouse common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setSuperRareAuctionHouse", _superRareAuctionHouse)
}

// SetSuperRareAuctionHouse is a paid mutator transaction binding the contract method 0x2a2a326c.
//
// Solidity: function setSuperRareAuctionHouse(address _superRareAuctionHouse) returns()
func (_Marketplace *MarketplaceSession) SetSuperRareAuctionHouse(_superRareAuctionHouse common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetSuperRareAuctionHouse(&_Marketplace.TransactOpts, _superRareAuctionHouse)
}

// SetSuperRareAuctionHouse is a paid mutator transaction binding the contract method 0x2a2a326c.
//
// Solidity: function setSuperRareAuctionHouse(address _superRareAuctionHouse) returns()
func (_Marketplace *MarketplaceTransactorSession) SetSuperRareAuctionHouse(_superRareAuctionHouse common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetSuperRareAuctionHouse(&_Marketplace.TransactOpts, _superRareAuctionHouse)
}

// SetSuperRareMarketplace is a paid mutator transaction binding the contract method 0x9c883af2.
//
// Solidity: function setSuperRareMarketplace(address _superRareMarketplace) returns()
func (_Marketplace *MarketplaceTransactor) SetSuperRareMarketplace(opts *bind.TransactOpts, _superRareMarketplace common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setSuperRareMarketplace", _superRareMarketplace)
}

// SetSuperRareMarketplace is a paid mutator transaction binding the contract method 0x9c883af2.
//
// Solidity: function setSuperRareMarketplace(address _superRareMarketplace) returns()
func (_Marketplace *MarketplaceSession) SetSuperRareMarketplace(_superRareMarketplace common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetSuperRareMarketplace(&_Marketplace.TransactOpts, _superRareMarketplace)
}

// SetSuperRareMarketplace is a paid mutator transaction binding the contract method 0x9c883af2.
//
// Solidity: function setSuperRareMarketplace(address _superRareMarketplace) returns()
func (_Marketplace *MarketplaceTransactorSession) SetSuperRareMarketplace(_superRareMarketplace common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetSuperRareMarketplace(&_Marketplace.TransactOpts, _superRareMarketplace)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x5138b08c.
//
// Solidity: function settleAuction(address _originContract, uint256 _tokenId) returns()
func (_Marketplace *MarketplaceTransactor) SettleAuction(opts *bind.TransactOpts, _originContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "settleAuction", _originContract, _tokenId)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x5138b08c.
//
// Solidity: function settleAuction(address _originContract, uint256 _tokenId) returns()
func (_Marketplace *MarketplaceSession) SettleAuction(_originContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SettleAuction(&_Marketplace.TransactOpts, _originContract, _tokenId)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x5138b08c.
//
// Solidity: function settleAuction(address _originContract, uint256 _tokenId) returns()
func (_Marketplace *MarketplaceTransactorSession) SettleAuction(_originContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SettleAuction(&_Marketplace.TransactOpts, _originContract, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// MarketplaceAcceptOfferIterator is returned from FilterAcceptOffer and is used to iterate over the raw logs and unpacked data for AcceptOffer events raised by the Marketplace contract.
type MarketplaceAcceptOfferIterator struct {
	Event *MarketplaceAcceptOffer // Event containing the contract specifics and raw log

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
func (it *MarketplaceAcceptOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceAcceptOffer)
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
		it.Event = new(MarketplaceAcceptOffer)
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
func (it *MarketplaceAcceptOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceAcceptOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceAcceptOffer represents a AcceptOffer event raised by the Marketplace contract.
type MarketplaceAcceptOffer struct {
	OriginContract  common.Address
	Bidder          common.Address
	Seller          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
	TokenId         *big.Int
	SplitAddresses  []common.Address
	SplitRatios     []uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAcceptOffer is a free log retrieval operation binding the contract event 0x97c3d2068ce177bc33d84acecc45eededcf298c4a9d4340ae03d4afbb3993f7b.
//
// Solidity: event AcceptOffer(address indexed _originContract, address indexed _bidder, address indexed _seller, address _currencyAddress, uint256 _amount, uint256 _tokenId, address[] _splitAddresses, uint8[] _splitRatios)
func (_Marketplace *MarketplaceFilterer) FilterAcceptOffer(opts *bind.FilterOpts, _originContract []common.Address, _bidder []common.Address, _seller []common.Address) (*MarketplaceAcceptOfferIterator, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "AcceptOffer", _originContractRule, _bidderRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceAcceptOfferIterator{contract: _Marketplace.contract, event: "AcceptOffer", logs: logs, sub: sub}, nil
}

// WatchAcceptOffer is a free log subscription operation binding the contract event 0x97c3d2068ce177bc33d84acecc45eededcf298c4a9d4340ae03d4afbb3993f7b.
//
// Solidity: event AcceptOffer(address indexed _originContract, address indexed _bidder, address indexed _seller, address _currencyAddress, uint256 _amount, uint256 _tokenId, address[] _splitAddresses, uint8[] _splitRatios)
func (_Marketplace *MarketplaceFilterer) WatchAcceptOffer(opts *bind.WatchOpts, sink chan<- *MarketplaceAcceptOffer, _originContract []common.Address, _bidder []common.Address, _seller []common.Address) (event.Subscription, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "AcceptOffer", _originContractRule, _bidderRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceAcceptOffer)
				if err := _Marketplace.contract.UnpackLog(event, "AcceptOffer", log); err != nil {
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

// ParseAcceptOffer is a log parse operation binding the contract event 0x97c3d2068ce177bc33d84acecc45eededcf298c4a9d4340ae03d4afbb3993f7b.
//
// Solidity: event AcceptOffer(address indexed _originContract, address indexed _bidder, address indexed _seller, address _currencyAddress, uint256 _amount, uint256 _tokenId, address[] _splitAddresses, uint8[] _splitRatios)
func (_Marketplace *MarketplaceFilterer) ParseAcceptOffer(log types.Log) (*MarketplaceAcceptOffer, error) {
	event := new(MarketplaceAcceptOffer)
	if err := _Marketplace.contract.UnpackLog(event, "AcceptOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceAuctionBidIterator is returned from FilterAuctionBid and is used to iterate over the raw logs and unpacked data for AuctionBid events raised by the Marketplace contract.
type MarketplaceAuctionBidIterator struct {
	Event *MarketplaceAuctionBid // Event containing the contract specifics and raw log

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
func (it *MarketplaceAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceAuctionBid)
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
		it.Event = new(MarketplaceAuctionBid)
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
func (it *MarketplaceAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceAuctionBid represents a AuctionBid event raised by the Marketplace contract.
type MarketplaceAuctionBid struct {
	ContractAddress  common.Address
	Bidder           common.Address
	TokenId          *big.Int
	CurrencyAddress  common.Address
	Amount           *big.Int
	StartedAuction   bool
	NewAuctionLength *big.Int
	PreviousBidder   common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAuctionBid is a free log retrieval operation binding the contract event 0x189a468e632a450afd491f84aa9ae94addc6813c7e309a7b1b1223f88802a77d.
//
// Solidity: event AuctionBid(address indexed _contractAddress, address indexed _bidder, uint256 indexed _tokenId, address _currencyAddress, uint256 _amount, bool _startedAuction, uint256 _newAuctionLength, address _previousBidder)
func (_Marketplace *MarketplaceFilterer) FilterAuctionBid(opts *bind.FilterOpts, _contractAddress []common.Address, _bidder []common.Address, _tokenId []*big.Int) (*MarketplaceAuctionBidIterator, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "AuctionBid", _contractAddressRule, _bidderRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceAuctionBidIterator{contract: _Marketplace.contract, event: "AuctionBid", logs: logs, sub: sub}, nil
}

// WatchAuctionBid is a free log subscription operation binding the contract event 0x189a468e632a450afd491f84aa9ae94addc6813c7e309a7b1b1223f88802a77d.
//
// Solidity: event AuctionBid(address indexed _contractAddress, address indexed _bidder, uint256 indexed _tokenId, address _currencyAddress, uint256 _amount, bool _startedAuction, uint256 _newAuctionLength, address _previousBidder)
func (_Marketplace *MarketplaceFilterer) WatchAuctionBid(opts *bind.WatchOpts, sink chan<- *MarketplaceAuctionBid, _contractAddress []common.Address, _bidder []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "AuctionBid", _contractAddressRule, _bidderRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceAuctionBid)
				if err := _Marketplace.contract.UnpackLog(event, "AuctionBid", log); err != nil {
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

// ParseAuctionBid is a log parse operation binding the contract event 0x189a468e632a450afd491f84aa9ae94addc6813c7e309a7b1b1223f88802a77d.
//
// Solidity: event AuctionBid(address indexed _contractAddress, address indexed _bidder, uint256 indexed _tokenId, address _currencyAddress, uint256 _amount, bool _startedAuction, uint256 _newAuctionLength, address _previousBidder)
func (_Marketplace *MarketplaceFilterer) ParseAuctionBid(log types.Log) (*MarketplaceAuctionBid, error) {
	event := new(MarketplaceAuctionBid)
	if err := _Marketplace.contract.UnpackLog(event, "AuctionBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceAuctionSettledIterator is returned from FilterAuctionSettled and is used to iterate over the raw logs and unpacked data for AuctionSettled events raised by the Marketplace contract.
type MarketplaceAuctionSettledIterator struct {
	Event *MarketplaceAuctionSettled // Event containing the contract specifics and raw log

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
func (it *MarketplaceAuctionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceAuctionSettled)
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
		it.Event = new(MarketplaceAuctionSettled)
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
func (it *MarketplaceAuctionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceAuctionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceAuctionSettled represents a AuctionSettled event raised by the Marketplace contract.
type MarketplaceAuctionSettled struct {
	ContractAddress common.Address
	Bidder          common.Address
	Seller          common.Address
	TokenId         *big.Int
	CurrencyAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettled is a free log retrieval operation binding the contract event 0xef4e2262a841641690bb931801dc0d1923e6b417cd217f91f8049d8aa9f5f086.
//
// Solidity: event AuctionSettled(address indexed _contractAddress, address indexed _bidder, address _seller, uint256 indexed _tokenId, address _currencyAddress, uint256 _amount)
func (_Marketplace *MarketplaceFilterer) FilterAuctionSettled(opts *bind.FilterOpts, _contractAddress []common.Address, _bidder []common.Address, _tokenId []*big.Int) (*MarketplaceAuctionSettledIterator, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "AuctionSettled", _contractAddressRule, _bidderRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceAuctionSettledIterator{contract: _Marketplace.contract, event: "AuctionSettled", logs: logs, sub: sub}, nil
}

// WatchAuctionSettled is a free log subscription operation binding the contract event 0xef4e2262a841641690bb931801dc0d1923e6b417cd217f91f8049d8aa9f5f086.
//
// Solidity: event AuctionSettled(address indexed _contractAddress, address indexed _bidder, address _seller, uint256 indexed _tokenId, address _currencyAddress, uint256 _amount)
func (_Marketplace *MarketplaceFilterer) WatchAuctionSettled(opts *bind.WatchOpts, sink chan<- *MarketplaceAuctionSettled, _contractAddress []common.Address, _bidder []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "AuctionSettled", _contractAddressRule, _bidderRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceAuctionSettled)
				if err := _Marketplace.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
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

// ParseAuctionSettled is a log parse operation binding the contract event 0xef4e2262a841641690bb931801dc0d1923e6b417cd217f91f8049d8aa9f5f086.
//
// Solidity: event AuctionSettled(address indexed _contractAddress, address indexed _bidder, address _seller, uint256 indexed _tokenId, address _currencyAddress, uint256 _amount)
func (_Marketplace *MarketplaceFilterer) ParseAuctionSettled(log types.Log) (*MarketplaceAuctionSettled, error) {
	event := new(MarketplaceAuctionSettled)
	if err := _Marketplace.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCancelAuctionIterator is returned from FilterCancelAuction and is used to iterate over the raw logs and unpacked data for CancelAuction events raised by the Marketplace contract.
type MarketplaceCancelAuctionIterator struct {
	Event *MarketplaceCancelAuction // Event containing the contract specifics and raw log

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
func (it *MarketplaceCancelAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCancelAuction)
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
		it.Event = new(MarketplaceCancelAuction)
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
func (it *MarketplaceCancelAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCancelAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCancelAuction represents a CancelAuction event raised by the Marketplace contract.
type MarketplaceCancelAuction struct {
	ContractAddress common.Address
	TokenId         *big.Int
	AuctionCreator  common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelAuction is a free log retrieval operation binding the contract event 0x26d4510b556e779d6507640413e013206e44c8f5d018c7c74ed8926f3f024a9c.
//
// Solidity: event CancelAuction(address indexed _contractAddress, uint256 indexed _tokenId, address indexed _auctionCreator)
func (_Marketplace *MarketplaceFilterer) FilterCancelAuction(opts *bind.FilterOpts, _contractAddress []common.Address, _tokenId []*big.Int, _auctionCreator []common.Address) (*MarketplaceCancelAuctionIterator, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _auctionCreatorRule []interface{}
	for _, _auctionCreatorItem := range _auctionCreator {
		_auctionCreatorRule = append(_auctionCreatorRule, _auctionCreatorItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CancelAuction", _contractAddressRule, _tokenIdRule, _auctionCreatorRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCancelAuctionIterator{contract: _Marketplace.contract, event: "CancelAuction", logs: logs, sub: sub}, nil
}

// WatchCancelAuction is a free log subscription operation binding the contract event 0x26d4510b556e779d6507640413e013206e44c8f5d018c7c74ed8926f3f024a9c.
//
// Solidity: event CancelAuction(address indexed _contractAddress, uint256 indexed _tokenId, address indexed _auctionCreator)
func (_Marketplace *MarketplaceFilterer) WatchCancelAuction(opts *bind.WatchOpts, sink chan<- *MarketplaceCancelAuction, _contractAddress []common.Address, _tokenId []*big.Int, _auctionCreator []common.Address) (event.Subscription, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _auctionCreatorRule []interface{}
	for _, _auctionCreatorItem := range _auctionCreator {
		_auctionCreatorRule = append(_auctionCreatorRule, _auctionCreatorItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CancelAuction", _contractAddressRule, _tokenIdRule, _auctionCreatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCancelAuction)
				if err := _Marketplace.contract.UnpackLog(event, "CancelAuction", log); err != nil {
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

// ParseCancelAuction is a log parse operation binding the contract event 0x26d4510b556e779d6507640413e013206e44c8f5d018c7c74ed8926f3f024a9c.
//
// Solidity: event CancelAuction(address indexed _contractAddress, uint256 indexed _tokenId, address indexed _auctionCreator)
func (_Marketplace *MarketplaceFilterer) ParseCancelAuction(log types.Log) (*MarketplaceCancelAuction, error) {
	event := new(MarketplaceCancelAuction)
	if err := _Marketplace.contract.UnpackLog(event, "CancelAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCancelOfferIterator is returned from FilterCancelOffer and is used to iterate over the raw logs and unpacked data for CancelOffer events raised by the Marketplace contract.
type MarketplaceCancelOfferIterator struct {
	Event *MarketplaceCancelOffer // Event containing the contract specifics and raw log

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
func (it *MarketplaceCancelOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCancelOffer)
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
		it.Event = new(MarketplaceCancelOffer)
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
func (it *MarketplaceCancelOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCancelOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCancelOffer represents a CancelOffer event raised by the Marketplace contract.
type MarketplaceCancelOffer struct {
	OriginContract  common.Address
	Bidder          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
	TokenId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelOffer is a free log retrieval operation binding the contract event 0xb9a071fe7d38dc86fbc448d440311b6bd67e5e09de8b1b62c72f5fe344100453.
//
// Solidity: event CancelOffer(address indexed _originContract, address indexed _bidder, address indexed _currencyAddress, uint256 _amount, uint256 _tokenId)
func (_Marketplace *MarketplaceFilterer) FilterCancelOffer(opts *bind.FilterOpts, _originContract []common.Address, _bidder []common.Address, _currencyAddress []common.Address) (*MarketplaceCancelOfferIterator, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}
	var _currencyAddressRule []interface{}
	for _, _currencyAddressItem := range _currencyAddress {
		_currencyAddressRule = append(_currencyAddressRule, _currencyAddressItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CancelOffer", _originContractRule, _bidderRule, _currencyAddressRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCancelOfferIterator{contract: _Marketplace.contract, event: "CancelOffer", logs: logs, sub: sub}, nil
}

// WatchCancelOffer is a free log subscription operation binding the contract event 0xb9a071fe7d38dc86fbc448d440311b6bd67e5e09de8b1b62c72f5fe344100453.
//
// Solidity: event CancelOffer(address indexed _originContract, address indexed _bidder, address indexed _currencyAddress, uint256 _amount, uint256 _tokenId)
func (_Marketplace *MarketplaceFilterer) WatchCancelOffer(opts *bind.WatchOpts, sink chan<- *MarketplaceCancelOffer, _originContract []common.Address, _bidder []common.Address, _currencyAddress []common.Address) (event.Subscription, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}
	var _currencyAddressRule []interface{}
	for _, _currencyAddressItem := range _currencyAddress {
		_currencyAddressRule = append(_currencyAddressRule, _currencyAddressItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CancelOffer", _originContractRule, _bidderRule, _currencyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCancelOffer)
				if err := _Marketplace.contract.UnpackLog(event, "CancelOffer", log); err != nil {
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

// ParseCancelOffer is a log parse operation binding the contract event 0xb9a071fe7d38dc86fbc448d440311b6bd67e5e09de8b1b62c72f5fe344100453.
//
// Solidity: event CancelOffer(address indexed _originContract, address indexed _bidder, address indexed _currencyAddress, uint256 _amount, uint256 _tokenId)
func (_Marketplace *MarketplaceFilterer) ParseCancelOffer(log types.Log) (*MarketplaceCancelOffer, error) {
	event := new(MarketplaceCancelOffer)
	if err := _Marketplace.contract.UnpackLog(event, "CancelOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewAuctionIterator is returned from FilterNewAuction and is used to iterate over the raw logs and unpacked data for NewAuction events raised by the Marketplace contract.
type MarketplaceNewAuctionIterator struct {
	Event *MarketplaceNewAuction // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewAuction)
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
		it.Event = new(MarketplaceNewAuction)
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
func (it *MarketplaceNewAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewAuction represents a NewAuction event raised by the Marketplace contract.
type MarketplaceNewAuction struct {
	ContractAddress common.Address
	TokenId         *big.Int
	AuctionCreator  common.Address
	CurrencyAddress common.Address
	StartingTime    *big.Int
	MinimumBid      *big.Int
	LengthOfAuction *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAuction is a free log retrieval operation binding the contract event 0xdb89081d9a5399380ffbcc376a961ed023027fca462e08b3d23146e4c6ac62f2.
//
// Solidity: event NewAuction(address indexed _contractAddress, uint256 indexed _tokenId, address indexed _auctionCreator, address _currencyAddress, uint256 _startingTime, uint256 _minimumBid, uint256 _lengthOfAuction)
func (_Marketplace *MarketplaceFilterer) FilterNewAuction(opts *bind.FilterOpts, _contractAddress []common.Address, _tokenId []*big.Int, _auctionCreator []common.Address) (*MarketplaceNewAuctionIterator, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _auctionCreatorRule []interface{}
	for _, _auctionCreatorItem := range _auctionCreator {
		_auctionCreatorRule = append(_auctionCreatorRule, _auctionCreatorItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewAuction", _contractAddressRule, _tokenIdRule, _auctionCreatorRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewAuctionIterator{contract: _Marketplace.contract, event: "NewAuction", logs: logs, sub: sub}, nil
}

// WatchNewAuction is a free log subscription operation binding the contract event 0xdb89081d9a5399380ffbcc376a961ed023027fca462e08b3d23146e4c6ac62f2.
//
// Solidity: event NewAuction(address indexed _contractAddress, uint256 indexed _tokenId, address indexed _auctionCreator, address _currencyAddress, uint256 _startingTime, uint256 _minimumBid, uint256 _lengthOfAuction)
func (_Marketplace *MarketplaceFilterer) WatchNewAuction(opts *bind.WatchOpts, sink chan<- *MarketplaceNewAuction, _contractAddress []common.Address, _tokenId []*big.Int, _auctionCreator []common.Address) (event.Subscription, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _auctionCreatorRule []interface{}
	for _, _auctionCreatorItem := range _auctionCreator {
		_auctionCreatorRule = append(_auctionCreatorRule, _auctionCreatorItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewAuction", _contractAddressRule, _tokenIdRule, _auctionCreatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewAuction)
				if err := _Marketplace.contract.UnpackLog(event, "NewAuction", log); err != nil {
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

// ParseNewAuction is a log parse operation binding the contract event 0xdb89081d9a5399380ffbcc376a961ed023027fca462e08b3d23146e4c6ac62f2.
//
// Solidity: event NewAuction(address indexed _contractAddress, uint256 indexed _tokenId, address indexed _auctionCreator, address _currencyAddress, uint256 _startingTime, uint256 _minimumBid, uint256 _lengthOfAuction)
func (_Marketplace *MarketplaceFilterer) ParseNewAuction(log types.Log) (*MarketplaceNewAuction, error) {
	event := new(MarketplaceNewAuction)
	if err := _Marketplace.contract.UnpackLog(event, "NewAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceOfferPlacedIterator is returned from FilterOfferPlaced and is used to iterate over the raw logs and unpacked data for OfferPlaced events raised by the Marketplace contract.
type MarketplaceOfferPlacedIterator struct {
	Event *MarketplaceOfferPlaced // Event containing the contract specifics and raw log

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
func (it *MarketplaceOfferPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOfferPlaced)
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
		it.Event = new(MarketplaceOfferPlaced)
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
func (it *MarketplaceOfferPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOfferPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOfferPlaced represents a OfferPlaced event raised by the Marketplace contract.
type MarketplaceOfferPlaced struct {
	OriginContract  common.Address
	Bidder          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
	TokenId         *big.Int
	Convertible     bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOfferPlaced is a free log retrieval operation binding the contract event 0xacbe96367816f728b84abf006ca777225b704072592132845e9a3dbd7b023691.
//
// Solidity: event OfferPlaced(address indexed _originContract, address indexed _bidder, address indexed _currencyAddress, uint256 _amount, uint256 _tokenId, bool _convertible)
func (_Marketplace *MarketplaceFilterer) FilterOfferPlaced(opts *bind.FilterOpts, _originContract []common.Address, _bidder []common.Address, _currencyAddress []common.Address) (*MarketplaceOfferPlacedIterator, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}
	var _currencyAddressRule []interface{}
	for _, _currencyAddressItem := range _currencyAddress {
		_currencyAddressRule = append(_currencyAddressRule, _currencyAddressItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OfferPlaced", _originContractRule, _bidderRule, _currencyAddressRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceOfferPlacedIterator{contract: _Marketplace.contract, event: "OfferPlaced", logs: logs, sub: sub}, nil
}

// WatchOfferPlaced is a free log subscription operation binding the contract event 0xacbe96367816f728b84abf006ca777225b704072592132845e9a3dbd7b023691.
//
// Solidity: event OfferPlaced(address indexed _originContract, address indexed _bidder, address indexed _currencyAddress, uint256 _amount, uint256 _tokenId, bool _convertible)
func (_Marketplace *MarketplaceFilterer) WatchOfferPlaced(opts *bind.WatchOpts, sink chan<- *MarketplaceOfferPlaced, _originContract []common.Address, _bidder []common.Address, _currencyAddress []common.Address) (event.Subscription, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}
	var _currencyAddressRule []interface{}
	for _, _currencyAddressItem := range _currencyAddress {
		_currencyAddressRule = append(_currencyAddressRule, _currencyAddressItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OfferPlaced", _originContractRule, _bidderRule, _currencyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOfferPlaced)
				if err := _Marketplace.contract.UnpackLog(event, "OfferPlaced", log); err != nil {
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

// ParseOfferPlaced is a log parse operation binding the contract event 0xacbe96367816f728b84abf006ca777225b704072592132845e9a3dbd7b023691.
//
// Solidity: event OfferPlaced(address indexed _originContract, address indexed _bidder, address indexed _currencyAddress, uint256 _amount, uint256 _tokenId, bool _convertible)
func (_Marketplace *MarketplaceFilterer) ParseOfferPlaced(log types.Log) (*MarketplaceOfferPlaced, error) {
	event := new(MarketplaceOfferPlaced)
	if err := _Marketplace.contract.UnpackLog(event, "OfferPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Marketplace contract.
type MarketplaceOwnershipTransferredIterator struct {
	Event *MarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOwnershipTransferred)
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
		it.Event = new(MarketplaceOwnershipTransferred)
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
func (it *MarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the Marketplace contract.
type MarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceOwnershipTransferredIterator{contract: _Marketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOwnershipTransferred)
				if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*MarketplaceOwnershipTransferred, error) {
	event := new(MarketplaceOwnershipTransferred)
	if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceSetSalePriceIterator is returned from FilterSetSalePrice and is used to iterate over the raw logs and unpacked data for SetSalePrice events raised by the Marketplace contract.
type MarketplaceSetSalePriceIterator struct {
	Event *MarketplaceSetSalePrice // Event containing the contract specifics and raw log

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
func (it *MarketplaceSetSalePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceSetSalePrice)
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
		it.Event = new(MarketplaceSetSalePrice)
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
func (it *MarketplaceSetSalePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceSetSalePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceSetSalePrice represents a SetSalePrice event raised by the Marketplace contract.
type MarketplaceSetSalePrice struct {
	OriginContract  common.Address
	CurrencyAddress common.Address
	Target          common.Address
	Amount          *big.Int
	TokenId         *big.Int
	SplitRecipients []common.Address
	SplitRatios     []uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetSalePrice is a free log retrieval operation binding the contract event 0xb6039ff1edf80efca6bc48b89f5415ba07fecb2d321058dae9ce6369b2ff964b.
//
// Solidity: event SetSalePrice(address indexed _originContract, address indexed _currencyAddress, address _target, uint256 _amount, uint256 _tokenId, address[] _splitRecipients, uint8[] _splitRatios)
func (_Marketplace *MarketplaceFilterer) FilterSetSalePrice(opts *bind.FilterOpts, _originContract []common.Address, _currencyAddress []common.Address) (*MarketplaceSetSalePriceIterator, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _currencyAddressRule []interface{}
	for _, _currencyAddressItem := range _currencyAddress {
		_currencyAddressRule = append(_currencyAddressRule, _currencyAddressItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "SetSalePrice", _originContractRule, _currencyAddressRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceSetSalePriceIterator{contract: _Marketplace.contract, event: "SetSalePrice", logs: logs, sub: sub}, nil
}

// WatchSetSalePrice is a free log subscription operation binding the contract event 0xb6039ff1edf80efca6bc48b89f5415ba07fecb2d321058dae9ce6369b2ff964b.
//
// Solidity: event SetSalePrice(address indexed _originContract, address indexed _currencyAddress, address _target, uint256 _amount, uint256 _tokenId, address[] _splitRecipients, uint8[] _splitRatios)
func (_Marketplace *MarketplaceFilterer) WatchSetSalePrice(opts *bind.WatchOpts, sink chan<- *MarketplaceSetSalePrice, _originContract []common.Address, _currencyAddress []common.Address) (event.Subscription, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _currencyAddressRule []interface{}
	for _, _currencyAddressItem := range _currencyAddress {
		_currencyAddressRule = append(_currencyAddressRule, _currencyAddressItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "SetSalePrice", _originContractRule, _currencyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceSetSalePrice)
				if err := _Marketplace.contract.UnpackLog(event, "SetSalePrice", log); err != nil {
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

// ParseSetSalePrice is a log parse operation binding the contract event 0xb6039ff1edf80efca6bc48b89f5415ba07fecb2d321058dae9ce6369b2ff964b.
//
// Solidity: event SetSalePrice(address indexed _originContract, address indexed _currencyAddress, address _target, uint256 _amount, uint256 _tokenId, address[] _splitRecipients, uint8[] _splitRatios)
func (_Marketplace *MarketplaceFilterer) ParseSetSalePrice(log types.Log) (*MarketplaceSetSalePrice, error) {
	event := new(MarketplaceSetSalePrice)
	if err := _Marketplace.contract.UnpackLog(event, "SetSalePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceSoldIterator is returned from FilterSold and is used to iterate over the raw logs and unpacked data for Sold events raised by the Marketplace contract.
type MarketplaceSoldIterator struct {
	Event *MarketplaceSold // Event containing the contract specifics and raw log

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
func (it *MarketplaceSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceSold)
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
		it.Event = new(MarketplaceSold)
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
func (it *MarketplaceSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceSold represents a Sold event raised by the Marketplace contract.
type MarketplaceSold struct {
	OriginContract  common.Address
	Buyer           common.Address
	Seller          common.Address
	CurrencyAddress common.Address
	Amount          *big.Int
	TokenId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSold is a free log retrieval operation binding the contract event 0x6f9e7bc841408072f4a49e469f90e1a634b85251803662bc8e5c220b28782472.
//
// Solidity: event Sold(address indexed _originContract, address indexed _buyer, address indexed _seller, address _currencyAddress, uint256 _amount, uint256 _tokenId)
func (_Marketplace *MarketplaceFilterer) FilterSold(opts *bind.FilterOpts, _originContract []common.Address, _buyer []common.Address, _seller []common.Address) (*MarketplaceSoldIterator, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Sold", _originContractRule, _buyerRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceSoldIterator{contract: _Marketplace.contract, event: "Sold", logs: logs, sub: sub}, nil
}

// WatchSold is a free log subscription operation binding the contract event 0x6f9e7bc841408072f4a49e469f90e1a634b85251803662bc8e5c220b28782472.
//
// Solidity: event Sold(address indexed _originContract, address indexed _buyer, address indexed _seller, address _currencyAddress, uint256 _amount, uint256 _tokenId)
func (_Marketplace *MarketplaceFilterer) WatchSold(opts *bind.WatchOpts, sink chan<- *MarketplaceSold, _originContract []common.Address, _buyer []common.Address, _seller []common.Address) (event.Subscription, error) {

	var _originContractRule []interface{}
	for _, _originContractItem := range _originContract {
		_originContractRule = append(_originContractRule, _originContractItem)
	}
	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Sold", _originContractRule, _buyerRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceSold)
				if err := _Marketplace.contract.UnpackLog(event, "Sold", log); err != nil {
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

// ParseSold is a log parse operation binding the contract event 0x6f9e7bc841408072f4a49e469f90e1a634b85251803662bc8e5c220b28782472.
//
// Solidity: event Sold(address indexed _originContract, address indexed _buyer, address indexed _seller, address _currencyAddress, uint256 _amount, uint256 _tokenId)
func (_Marketplace *MarketplaceFilterer) ParseSold(log types.Log) (*MarketplaceSold, error) {
	event := new(MarketplaceSold)
	if err := _Marketplace.contract.UnpackLog(event, "Sold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
