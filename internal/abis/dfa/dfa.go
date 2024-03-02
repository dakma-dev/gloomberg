// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dfa

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

// IBaseArgs is an auto generated low-level Go binding around an user-defined struct.
type IBaseArgs struct {
	Payees             []common.Address
	Shares             []*big.Int
	RoyaltiesRecipient common.Address
	RoyaltyValue       *big.Int
	Stages             []IBaseStageConfig
	TokenConfig        IBaseTokenConfig
}

// IBaseStageConfig is an auto generated low-level Go binding around an user-defined struct.
type IBaseStageConfig struct {
	Limit          *big.Int
	Price          *big.Int
	MerkleTreeRoot [32]byte
}

// IBaseTokenConfig is an auto generated low-level Go binding around an user-defined struct.
type IBaseTokenConfig struct {
	Name   string
	Symbol string
	Supply *big.Int
	Prefix string
	Suffix string
}

// IERC721AUpgradeableTokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type IERC721AUpgradeableTokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
	Burned         bool
	ExtraData      *big.Int
}

// DfaMetaData contains all meta data concerning the Dfa contract.
var DfaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQueryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20PaymentReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"PayeeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RoyaltiesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stageId_\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"accounts_\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"adminMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stageId_\",\"type\":\"uint8\"}],\"name\":\"balanceLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stageId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"currentBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"explicitOwnershipOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721AUpgradeable.TokenOwnership\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"explicitOwnershipsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721AUpgradeable.TokenOwnership[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"royaltiesRecipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyValue\",\"type\":\"uint96\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleTreeRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIBase.StageConfig[]\",\"name\":\"stages\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"suffix\",\"type\":\"string\"}],\"internalType\":\"structIBase.TokenConfig\",\"name\":\"tokenConfig\",\"type\":\"tuple\"}],\"internalType\":\"structIBase.Args\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isSoulBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stageId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof_\",\"type\":\"bytes32[]\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"manuallyApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stageId_\",\"type\":\"uint8\"}],\"name\":\"merkleTreeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof_\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"stage_\",\"type\":\"uint8\"}],\"name\":\"multiStageMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"payee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payeesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prefix\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stage_\",\"type\":\"uint8\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"publicMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"releasable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"releasable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"releaseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stageId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"remainingBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stageId_\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"accounts_\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable_\",\"type\":\"bool\"}],\"name\":\"setAdminPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply_\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix_\",\"type\":\"string\"}],\"name\":\"setPrefix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stage_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stage_\",\"type\":\"uint8\"}],\"name\":\"setStage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"suffix_\",\"type\":\"string\"}],\"name\":\"setSuffix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suffix\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"toggleSoulBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stop\",\"type\":\"uint256\"}],\"name\":\"tokensOfOwnerIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"totalReleased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReleased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stageId_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"updateBalanceLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stageId_\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"merkleTreeRoot_\",\"type\":\"bytes32\"}],\"name\":\"updateMerkleTreeRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DfaABI is the input ABI used to generate the binding from.
// Deprecated: Use DfaMetaData.ABI instead.
var DfaABI = DfaMetaData.ABI

// Dfa is an auto generated Go binding around an Ethereum contract.
type Dfa struct {
	DfaCaller     // Read-only binding to the contract
	DfaTransactor // Write-only binding to the contract
	DfaFilterer   // Log filterer for contract events
}

// DfaCaller is an auto generated read-only Go binding around an Ethereum contract.
type DfaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DfaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DfaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DfaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DfaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DfaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DfaSession struct {
	Contract     *Dfa              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DfaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DfaCallerSession struct {
	Contract *DfaCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DfaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DfaTransactorSession struct {
	Contract     *DfaTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DfaRaw is an auto generated low-level Go binding around an Ethereum contract.
type DfaRaw struct {
	Contract *Dfa // Generic contract binding to access the raw methods on
}

// DfaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DfaCallerRaw struct {
	Contract *DfaCaller // Generic read-only contract binding to access the raw methods on
}

// DfaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DfaTransactorRaw struct {
	Contract *DfaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDfa creates a new instance of Dfa, bound to a specific deployed contract.
func NewDfa(address common.Address, backend bind.ContractBackend) (*Dfa, error) {
	contract, err := bindDfa(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dfa{DfaCaller: DfaCaller{contract: contract}, DfaTransactor: DfaTransactor{contract: contract}, DfaFilterer: DfaFilterer{contract: contract}}, nil
}

// NewDfaCaller creates a new read-only instance of Dfa, bound to a specific deployed contract.
func NewDfaCaller(address common.Address, caller bind.ContractCaller) (*DfaCaller, error) {
	contract, err := bindDfa(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DfaCaller{contract: contract}, nil
}

// NewDfaTransactor creates a new write-only instance of Dfa, bound to a specific deployed contract.
func NewDfaTransactor(address common.Address, transactor bind.ContractTransactor) (*DfaTransactor, error) {
	contract, err := bindDfa(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DfaTransactor{contract: contract}, nil
}

// NewDfaFilterer creates a new log filterer instance of Dfa, bound to a specific deployed contract.
func NewDfaFilterer(address common.Address, filterer bind.ContractFilterer) (*DfaFilterer, error) {
	contract, err := bindDfa(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DfaFilterer{contract: contract}, nil
}

// bindDfa binds a generic wrapper to an already deployed contract.
func bindDfa(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DfaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dfa *DfaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dfa.Contract.DfaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dfa *DfaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfa.Contract.DfaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dfa *DfaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dfa.Contract.DfaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dfa *DfaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dfa.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dfa *DfaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfa.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dfa *DfaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dfa.Contract.contract.Transact(opts, method, params...)
}

// BalanceLimit is a free data retrieval call binding the contract method 0xb1ba72d6.
//
// Solidity: function balanceLimit(uint8 stageId_) view returns(uint256)
func (_Dfa *DfaCaller) BalanceLimit(opts *bind.CallOpts, stageId_ uint8) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "balanceLimit", stageId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceLimit is a free data retrieval call binding the contract method 0xb1ba72d6.
//
// Solidity: function balanceLimit(uint8 stageId_) view returns(uint256)
func (_Dfa *DfaSession) BalanceLimit(stageId_ uint8) (*big.Int, error) {
	return _Dfa.Contract.BalanceLimit(&_Dfa.CallOpts, stageId_)
}

// BalanceLimit is a free data retrieval call binding the contract method 0xb1ba72d6.
//
// Solidity: function balanceLimit(uint8 stageId_) view returns(uint256)
func (_Dfa *DfaCallerSession) BalanceLimit(stageId_ uint8) (*big.Int, error) {
	return _Dfa.Contract.BalanceLimit(&_Dfa.CallOpts, stageId_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dfa *DfaCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dfa *DfaSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Dfa.Contract.BalanceOf(&_Dfa.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dfa *DfaCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Dfa.Contract.BalanceOf(&_Dfa.CallOpts, owner)
}

// CurrentBalance is a free data retrieval call binding the contract method 0xe4ab4bb9.
//
// Solidity: function currentBalance(uint8 stageId_, address account_) view returns(uint256)
func (_Dfa *DfaCaller) CurrentBalance(opts *bind.CallOpts, stageId_ uint8, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "currentBalance", stageId_, account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBalance is a free data retrieval call binding the contract method 0xe4ab4bb9.
//
// Solidity: function currentBalance(uint8 stageId_, address account_) view returns(uint256)
func (_Dfa *DfaSession) CurrentBalance(stageId_ uint8, account_ common.Address) (*big.Int, error) {
	return _Dfa.Contract.CurrentBalance(&_Dfa.CallOpts, stageId_, account_)
}

// CurrentBalance is a free data retrieval call binding the contract method 0xe4ab4bb9.
//
// Solidity: function currentBalance(uint8 stageId_, address account_) view returns(uint256)
func (_Dfa *DfaCallerSession) CurrentBalance(stageId_ uint8, account_ common.Address) (*big.Int, error) {
	return _Dfa.Contract.CurrentBalance(&_Dfa.CallOpts, stageId_, account_)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Dfa *DfaCaller) ExplicitOwnershipOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "explicitOwnershipOf", tokenId)

	if err != nil {
		return *new(IERC721AUpgradeableTokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721AUpgradeableTokenOwnership)).(*IERC721AUpgradeableTokenOwnership)

	return out0, err

}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Dfa *DfaSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	return _Dfa.Contract.ExplicitOwnershipOf(&_Dfa.CallOpts, tokenId)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Dfa *DfaCallerSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	return _Dfa.Contract.ExplicitOwnershipOf(&_Dfa.CallOpts, tokenId)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Dfa *DfaCaller) ExplicitOwnershipsOf(opts *bind.CallOpts, tokenIds []*big.Int) ([]IERC721AUpgradeableTokenOwnership, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "explicitOwnershipsOf", tokenIds)

	if err != nil {
		return *new([]IERC721AUpgradeableTokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC721AUpgradeableTokenOwnership)).(*[]IERC721AUpgradeableTokenOwnership)

	return out0, err

}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Dfa *DfaSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721AUpgradeableTokenOwnership, error) {
	return _Dfa.Contract.ExplicitOwnershipsOf(&_Dfa.CallOpts, tokenIds)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Dfa *DfaCallerSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721AUpgradeableTokenOwnership, error) {
	return _Dfa.Contract.ExplicitOwnershipsOf(&_Dfa.CallOpts, tokenIds)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dfa *DfaCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dfa *DfaSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Dfa.Contract.GetApproved(&_Dfa.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dfa *DfaCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Dfa.Contract.GetApproved(&_Dfa.CallOpts, tokenId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account_) view returns(bool)
func (_Dfa *DfaCaller) IsAdmin(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "isAdmin", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account_) view returns(bool)
func (_Dfa *DfaSession) IsAdmin(account_ common.Address) (bool, error) {
	return _Dfa.Contract.IsAdmin(&_Dfa.CallOpts, account_)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account_) view returns(bool)
func (_Dfa *DfaCallerSession) IsAdmin(account_ common.Address) (bool, error) {
	return _Dfa.Contract.IsAdmin(&_Dfa.CallOpts, account_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dfa *DfaCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dfa *DfaSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Dfa.Contract.IsApprovedForAll(&_Dfa.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dfa *DfaCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Dfa.Contract.IsApprovedForAll(&_Dfa.CallOpts, owner, operator)
}

// IsSoulBound is a free data retrieval call binding the contract method 0xd508a212.
//
// Solidity: function isSoulBound(uint256 ) view returns(bool)
func (_Dfa *DfaCaller) IsSoulBound(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "isSoulBound", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSoulBound is a free data retrieval call binding the contract method 0xd508a212.
//
// Solidity: function isSoulBound(uint256 ) view returns(bool)
func (_Dfa *DfaSession) IsSoulBound(arg0 *big.Int) (bool, error) {
	return _Dfa.Contract.IsSoulBound(&_Dfa.CallOpts, arg0)
}

// IsSoulBound is a free data retrieval call binding the contract method 0xd508a212.
//
// Solidity: function isSoulBound(uint256 ) view returns(bool)
func (_Dfa *DfaCallerSession) IsSoulBound(arg0 *big.Int) (bool, error) {
	return _Dfa.Contract.IsSoulBound(&_Dfa.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0xcb3afdb6.
//
// Solidity: function isWhitelisted(uint8 stageId_, address account_, bytes32[] proof_) view returns(bool)
func (_Dfa *DfaCaller) IsWhitelisted(opts *bind.CallOpts, stageId_ uint8, account_ common.Address, proof_ [][32]byte) (bool, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "isWhitelisted", stageId_, account_, proof_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0xcb3afdb6.
//
// Solidity: function isWhitelisted(uint8 stageId_, address account_, bytes32[] proof_) view returns(bool)
func (_Dfa *DfaSession) IsWhitelisted(stageId_ uint8, account_ common.Address, proof_ [][32]byte) (bool, error) {
	return _Dfa.Contract.IsWhitelisted(&_Dfa.CallOpts, stageId_, account_, proof_)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0xcb3afdb6.
//
// Solidity: function isWhitelisted(uint8 stageId_, address account_, bytes32[] proof_) view returns(bool)
func (_Dfa *DfaCallerSession) IsWhitelisted(stageId_ uint8, account_ common.Address, proof_ [][32]byte) (bool, error) {
	return _Dfa.Contract.IsWhitelisted(&_Dfa.CallOpts, stageId_, account_, proof_)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Dfa *DfaCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Dfa *DfaSession) MaxSupply() (*big.Int, error) {
	return _Dfa.Contract.MaxSupply(&_Dfa.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Dfa *DfaCallerSession) MaxSupply() (*big.Int, error) {
	return _Dfa.Contract.MaxSupply(&_Dfa.CallOpts)
}

// MerkleTreeRoot is a free data retrieval call binding the contract method 0xad0127f0.
//
// Solidity: function merkleTreeRoot(uint8 stageId_) view returns(bytes32)
func (_Dfa *DfaCaller) MerkleTreeRoot(opts *bind.CallOpts, stageId_ uint8) ([32]byte, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "merkleTreeRoot", stageId_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleTreeRoot is a free data retrieval call binding the contract method 0xad0127f0.
//
// Solidity: function merkleTreeRoot(uint8 stageId_) view returns(bytes32)
func (_Dfa *DfaSession) MerkleTreeRoot(stageId_ uint8) ([32]byte, error) {
	return _Dfa.Contract.MerkleTreeRoot(&_Dfa.CallOpts, stageId_)
}

// MerkleTreeRoot is a free data retrieval call binding the contract method 0xad0127f0.
//
// Solidity: function merkleTreeRoot(uint8 stageId_) view returns(bytes32)
func (_Dfa *DfaCallerSession) MerkleTreeRoot(stageId_ uint8) ([32]byte, error) {
	return _Dfa.Contract.MerkleTreeRoot(&_Dfa.CallOpts, stageId_)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dfa *DfaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dfa *DfaSession) Name() (string, error) {
	return _Dfa.Contract.Name(&_Dfa.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dfa *DfaCallerSession) Name() (string, error) {
	return _Dfa.Contract.Name(&_Dfa.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dfa *DfaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dfa *DfaSession) Owner() (common.Address, error) {
	return _Dfa.Contract.Owner(&_Dfa.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dfa *DfaCallerSession) Owner() (common.Address, error) {
	return _Dfa.Contract.Owner(&_Dfa.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dfa *DfaCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dfa *DfaSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Dfa.Contract.OwnerOf(&_Dfa.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dfa *DfaCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Dfa.Contract.OwnerOf(&_Dfa.CallOpts, tokenId)
}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_Dfa *DfaCaller) Payee(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "payee", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_Dfa *DfaSession) Payee(index *big.Int) (common.Address, error) {
	return _Dfa.Contract.Payee(&_Dfa.CallOpts, index)
}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_Dfa *DfaCallerSession) Payee(index *big.Int) (common.Address, error) {
	return _Dfa.Contract.Payee(&_Dfa.CallOpts, index)
}

// PayeesLength is a free data retrieval call binding the contract method 0xe919ecad.
//
// Solidity: function payeesLength() view returns(uint256)
func (_Dfa *DfaCaller) PayeesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "payeesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayeesLength is a free data retrieval call binding the contract method 0xe919ecad.
//
// Solidity: function payeesLength() view returns(uint256)
func (_Dfa *DfaSession) PayeesLength() (*big.Int, error) {
	return _Dfa.Contract.PayeesLength(&_Dfa.CallOpts)
}

// PayeesLength is a free data retrieval call binding the contract method 0xe919ecad.
//
// Solidity: function payeesLength() view returns(uint256)
func (_Dfa *DfaCallerSession) PayeesLength() (*big.Int, error) {
	return _Dfa.Contract.PayeesLength(&_Dfa.CallOpts)
}

// Prefix is a free data retrieval call binding the contract method 0x75dadb32.
//
// Solidity: function prefix() view returns(string)
func (_Dfa *DfaCaller) Prefix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "prefix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Prefix is a free data retrieval call binding the contract method 0x75dadb32.
//
// Solidity: function prefix() view returns(string)
func (_Dfa *DfaSession) Prefix() (string, error) {
	return _Dfa.Contract.Prefix(&_Dfa.CallOpts)
}

// Prefix is a free data retrieval call binding the contract method 0x75dadb32.
//
// Solidity: function prefix() view returns(string)
func (_Dfa *DfaCallerSession) Prefix() (string, error) {
	return _Dfa.Contract.Prefix(&_Dfa.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xb7fafcd7.
//
// Solidity: function price(uint8 stage_) view returns(uint256)
func (_Dfa *DfaCaller) Price(opts *bind.CallOpts, stage_ uint8) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "price", stage_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xb7fafcd7.
//
// Solidity: function price(uint8 stage_) view returns(uint256)
func (_Dfa *DfaSession) Price(stage_ uint8) (*big.Int, error) {
	return _Dfa.Contract.Price(&_Dfa.CallOpts, stage_)
}

// Price is a free data retrieval call binding the contract method 0xb7fafcd7.
//
// Solidity: function price(uint8 stage_) view returns(uint256)
func (_Dfa *DfaCallerSession) Price(stage_ uint8) (*big.Int, error) {
	return _Dfa.Contract.Price(&_Dfa.CallOpts, stage_)
}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address account) view returns(uint256)
func (_Dfa *DfaCaller) Releasable(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "releasable", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address account) view returns(uint256)
func (_Dfa *DfaSession) Releasable(account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Releasable(&_Dfa.CallOpts, account)
}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address account) view returns(uint256)
func (_Dfa *DfaCallerSession) Releasable(account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Releasable(&_Dfa.CallOpts, account)
}

// Releasable0 is a free data retrieval call binding the contract method 0xc45ac050.
//
// Solidity: function releasable(address token, address account) view returns(uint256)
func (_Dfa *DfaCaller) Releasable0(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "releasable0", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Releasable0 is a free data retrieval call binding the contract method 0xc45ac050.
//
// Solidity: function releasable(address token, address account) view returns(uint256)
func (_Dfa *DfaSession) Releasable0(token common.Address, account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Releasable0(&_Dfa.CallOpts, token, account)
}

// Releasable0 is a free data retrieval call binding the contract method 0xc45ac050.
//
// Solidity: function releasable(address token, address account) view returns(uint256)
func (_Dfa *DfaCallerSession) Releasable0(token common.Address, account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Releasable0(&_Dfa.CallOpts, token, account)
}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_Dfa *DfaCaller) Released(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "released", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_Dfa *DfaSession) Released(token common.Address, account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Released(&_Dfa.CallOpts, token, account)
}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_Dfa *DfaCallerSession) Released(token common.Address, account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Released(&_Dfa.CallOpts, token, account)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_Dfa *DfaCaller) Released0(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "released0", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_Dfa *DfaSession) Released0(account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Released0(&_Dfa.CallOpts, account)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_Dfa *DfaCallerSession) Released0(account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Released0(&_Dfa.CallOpts, account)
}

// RemainingBalance is a free data retrieval call binding the contract method 0x6dba1163.
//
// Solidity: function remainingBalance(uint8 stageId_, address account_) view returns(uint256)
func (_Dfa *DfaCaller) RemainingBalance(opts *bind.CallOpts, stageId_ uint8, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "remainingBalance", stageId_, account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingBalance is a free data retrieval call binding the contract method 0x6dba1163.
//
// Solidity: function remainingBalance(uint8 stageId_, address account_) view returns(uint256)
func (_Dfa *DfaSession) RemainingBalance(stageId_ uint8, account_ common.Address) (*big.Int, error) {
	return _Dfa.Contract.RemainingBalance(&_Dfa.CallOpts, stageId_, account_)
}

// RemainingBalance is a free data retrieval call binding the contract method 0x6dba1163.
//
// Solidity: function remainingBalance(uint8 stageId_, address account_) view returns(uint256)
func (_Dfa *DfaCallerSession) RemainingBalance(stageId_ uint8, account_ common.Address) (*big.Int, error) {
	return _Dfa.Contract.RemainingBalance(&_Dfa.CallOpts, stageId_, account_)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 value_) view returns(address receiver, uint256 royaltyAmount)
func (_Dfa *DfaCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, value_ *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "royaltyInfo", arg0, value_)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 value_) view returns(address receiver, uint256 royaltyAmount)
func (_Dfa *DfaSession) RoyaltyInfo(arg0 *big.Int, value_ *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Dfa.Contract.RoyaltyInfo(&_Dfa.CallOpts, arg0, value_)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 value_) view returns(address receiver, uint256 royaltyAmount)
func (_Dfa *DfaCallerSession) RoyaltyInfo(arg0 *big.Int, value_ *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Dfa.Contract.RoyaltyInfo(&_Dfa.CallOpts, arg0, value_)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_Dfa *DfaCaller) Shares(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "shares", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_Dfa *DfaSession) Shares(account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Shares(&_Dfa.CallOpts, account)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_Dfa *DfaCallerSession) Shares(account common.Address) (*big.Int, error) {
	return _Dfa.Contract.Shares(&_Dfa.CallOpts, account)
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() view returns(uint8)
func (_Dfa *DfaCaller) Stage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "stage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() view returns(uint8)
func (_Dfa *DfaSession) Stage() (uint8, error) {
	return _Dfa.Contract.Stage(&_Dfa.CallOpts)
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() view returns(uint8)
func (_Dfa *DfaCallerSession) Stage() (uint8, error) {
	return _Dfa.Contract.Stage(&_Dfa.CallOpts)
}

// Suffix is a free data retrieval call binding the contract method 0xf7073c3a.
//
// Solidity: function suffix() view returns(string)
func (_Dfa *DfaCaller) Suffix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "suffix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Suffix is a free data retrieval call binding the contract method 0xf7073c3a.
//
// Solidity: function suffix() view returns(string)
func (_Dfa *DfaSession) Suffix() (string, error) {
	return _Dfa.Contract.Suffix(&_Dfa.CallOpts)
}

// Suffix is a free data retrieval call binding the contract method 0xf7073c3a.
//
// Solidity: function suffix() view returns(string)
func (_Dfa *DfaCallerSession) Suffix() (string, error) {
	return _Dfa.Contract.Suffix(&_Dfa.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dfa *DfaCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dfa *DfaSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dfa.Contract.SupportsInterface(&_Dfa.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dfa *DfaCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dfa.Contract.SupportsInterface(&_Dfa.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dfa *DfaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dfa *DfaSession) Symbol() (string, error) {
	return _Dfa.Contract.Symbol(&_Dfa.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dfa *DfaCallerSession) Symbol() (string, error) {
	return _Dfa.Contract.Symbol(&_Dfa.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Dfa *DfaCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Dfa *DfaSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Dfa.Contract.TokenURI(&_Dfa.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Dfa *DfaCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Dfa.Contract.TokenURI(&_Dfa.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Dfa *DfaCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Dfa *DfaSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Dfa.Contract.TokensOfOwner(&_Dfa.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Dfa *DfaCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Dfa.Contract.TokensOfOwner(&_Dfa.CallOpts, owner)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Dfa *DfaCaller) TokensOfOwnerIn(opts *bind.CallOpts, owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "tokensOfOwnerIn", owner, start, stop)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Dfa *DfaSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Dfa.Contract.TokensOfOwnerIn(&_Dfa.CallOpts, owner, start, stop)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Dfa *DfaCallerSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Dfa.Contract.TokensOfOwnerIn(&_Dfa.CallOpts, owner, start, stop)
}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_Dfa *DfaCaller) TotalReleased(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "totalReleased", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_Dfa *DfaSession) TotalReleased(token common.Address) (*big.Int, error) {
	return _Dfa.Contract.TotalReleased(&_Dfa.CallOpts, token)
}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_Dfa *DfaCallerSession) TotalReleased(token common.Address) (*big.Int, error) {
	return _Dfa.Contract.TotalReleased(&_Dfa.CallOpts, token)
}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Dfa *DfaCaller) TotalReleased0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "totalReleased0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Dfa *DfaSession) TotalReleased0() (*big.Int, error) {
	return _Dfa.Contract.TotalReleased0(&_Dfa.CallOpts)
}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Dfa *DfaCallerSession) TotalReleased0() (*big.Int, error) {
	return _Dfa.Contract.TotalReleased0(&_Dfa.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Dfa *DfaCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Dfa *DfaSession) TotalShares() (*big.Int, error) {
	return _Dfa.Contract.TotalShares(&_Dfa.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Dfa *DfaCallerSession) TotalShares() (*big.Int, error) {
	return _Dfa.Contract.TotalShares(&_Dfa.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dfa *DfaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfa.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dfa *DfaSession) TotalSupply() (*big.Int, error) {
	return _Dfa.Contract.TotalSupply(&_Dfa.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dfa *DfaCallerSession) TotalSupply() (*big.Int, error) {
	return _Dfa.Contract.TotalSupply(&_Dfa.CallOpts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x21a588de.
//
// Solidity: function addToWhitelist(uint8 stageId_, address[] accounts_) returns()
func (_Dfa *DfaTransactor) AddToWhitelist(opts *bind.TransactOpts, stageId_ uint8, accounts_ []common.Address) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "addToWhitelist", stageId_, accounts_)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x21a588de.
//
// Solidity: function addToWhitelist(uint8 stageId_, address[] accounts_) returns()
func (_Dfa *DfaSession) AddToWhitelist(stageId_ uint8, accounts_ []common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.AddToWhitelist(&_Dfa.TransactOpts, stageId_, accounts_)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x21a588de.
//
// Solidity: function addToWhitelist(uint8 stageId_, address[] accounts_) returns()
func (_Dfa *DfaTransactorSession) AddToWhitelist(stageId_ uint8, accounts_ []common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.AddToWhitelist(&_Dfa.TransactOpts, stageId_, accounts_)
}

// AdminMint is a paid mutator transaction binding the contract method 0xa49340cc.
//
// Solidity: function adminMint(address[] accounts_, uint256[] amounts_) returns()
func (_Dfa *DfaTransactor) AdminMint(opts *bind.TransactOpts, accounts_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "adminMint", accounts_, amounts_)
}

// AdminMint is a paid mutator transaction binding the contract method 0xa49340cc.
//
// Solidity: function adminMint(address[] accounts_, uint256[] amounts_) returns()
func (_Dfa *DfaSession) AdminMint(accounts_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.AdminMint(&_Dfa.TransactOpts, accounts_, amounts_)
}

// AdminMint is a paid mutator transaction binding the contract method 0xa49340cc.
//
// Solidity: function adminMint(address[] accounts_, uint256[] amounts_) returns()
func (_Dfa *DfaTransactorSession) AdminMint(accounts_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.AdminMint(&_Dfa.TransactOpts, accounts_, amounts_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Dfa *DfaTransactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Dfa *DfaSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.Approve(&_Dfa.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Dfa *DfaTransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.Approve(&_Dfa.TransactOpts, operator, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x9635b16a.
//
// Solidity: function initialize((address[],uint256[],address,uint96,(uint256,uint256,bytes32)[],(string,string,uint256,string,string)) args) returns()
func (_Dfa *DfaTransactor) Initialize(opts *bind.TransactOpts, args IBaseArgs) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "initialize", args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9635b16a.
//
// Solidity: function initialize((address[],uint256[],address,uint96,(uint256,uint256,bytes32)[],(string,string,uint256,string,string)) args) returns()
func (_Dfa *DfaSession) Initialize(args IBaseArgs) (*types.Transaction, error) {
	return _Dfa.Contract.Initialize(&_Dfa.TransactOpts, args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9635b16a.
//
// Solidity: function initialize((address[],uint256[],address,uint96,(uint256,uint256,bytes32)[],(string,string,uint256,string,string)) args) returns()
func (_Dfa *DfaTransactorSession) Initialize(args IBaseArgs) (*types.Transaction, error) {
	return _Dfa.Contract.Initialize(&_Dfa.TransactOpts, args)
}

// ManuallyApprove is a paid mutator transaction binding the contract method 0xd3570687.
//
// Solidity: function manuallyApprove(address admin_) returns()
func (_Dfa *DfaTransactor) ManuallyApprove(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "manuallyApprove", admin_)
}

// ManuallyApprove is a paid mutator transaction binding the contract method 0xd3570687.
//
// Solidity: function manuallyApprove(address admin_) returns()
func (_Dfa *DfaSession) ManuallyApprove(admin_ common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.ManuallyApprove(&_Dfa.TransactOpts, admin_)
}

// ManuallyApprove is a paid mutator transaction binding the contract method 0xd3570687.
//
// Solidity: function manuallyApprove(address admin_) returns()
func (_Dfa *DfaTransactorSession) ManuallyApprove(admin_ common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.ManuallyApprove(&_Dfa.TransactOpts, admin_)
}

// MultiStageMint is a paid mutator transaction binding the contract method 0xb03bc27c.
//
// Solidity: function multiStageMint(uint256 amount_, bytes32[] proof_, uint8 stage_) payable returns()
func (_Dfa *DfaTransactor) MultiStageMint(opts *bind.TransactOpts, amount_ *big.Int, proof_ [][32]byte, stage_ uint8) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "multiStageMint", amount_, proof_, stage_)
}

// MultiStageMint is a paid mutator transaction binding the contract method 0xb03bc27c.
//
// Solidity: function multiStageMint(uint256 amount_, bytes32[] proof_, uint8 stage_) payable returns()
func (_Dfa *DfaSession) MultiStageMint(amount_ *big.Int, proof_ [][32]byte, stage_ uint8) (*types.Transaction, error) {
	return _Dfa.Contract.MultiStageMint(&_Dfa.TransactOpts, amount_, proof_, stage_)
}

// MultiStageMint is a paid mutator transaction binding the contract method 0xb03bc27c.
//
// Solidity: function multiStageMint(uint256 amount_, bytes32[] proof_, uint8 stage_) payable returns()
func (_Dfa *DfaTransactorSession) MultiStageMint(amount_ *big.Int, proof_ [][32]byte, stage_ uint8) (*types.Transaction, error) {
	return _Dfa.Contract.MultiStageMint(&_Dfa.TransactOpts, amount_, proof_, stage_)
}

// PublicMint is a paid mutator transaction binding the contract method 0x2db11544.
//
// Solidity: function publicMint(uint256 amount_) payable returns()
func (_Dfa *DfaTransactor) PublicMint(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "publicMint", amount_)
}

// PublicMint is a paid mutator transaction binding the contract method 0x2db11544.
//
// Solidity: function publicMint(uint256 amount_) payable returns()
func (_Dfa *DfaSession) PublicMint(amount_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.PublicMint(&_Dfa.TransactOpts, amount_)
}

// PublicMint is a paid mutator transaction binding the contract method 0x2db11544.
//
// Solidity: function publicMint(uint256 amount_) payable returns()
func (_Dfa *DfaTransactorSession) PublicMint(amount_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.PublicMint(&_Dfa.TransactOpts, amount_)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_Dfa *DfaTransactor) Release(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "release", account)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_Dfa *DfaSession) Release(account common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.Release(&_Dfa.TransactOpts, account)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_Dfa *DfaTransactorSession) Release(account common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.Release(&_Dfa.TransactOpts, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_Dfa *DfaTransactor) Release0(opts *bind.TransactOpts, token common.Address, account common.Address) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "release0", token, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_Dfa *DfaSession) Release0(token common.Address, account common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.Release0(&_Dfa.TransactOpts, token, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_Dfa *DfaTransactorSession) Release0(token common.Address, account common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.Release0(&_Dfa.TransactOpts, token, account)
}

// ReleaseAll is a paid mutator transaction binding the contract method 0x580fc80a.
//
// Solidity: function releaseAll(address token) returns()
func (_Dfa *DfaTransactor) ReleaseAll(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "releaseAll", token)
}

// ReleaseAll is a paid mutator transaction binding the contract method 0x580fc80a.
//
// Solidity: function releaseAll(address token) returns()
func (_Dfa *DfaSession) ReleaseAll(token common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.ReleaseAll(&_Dfa.TransactOpts, token)
}

// ReleaseAll is a paid mutator transaction binding the contract method 0x580fc80a.
//
// Solidity: function releaseAll(address token) returns()
func (_Dfa *DfaTransactorSession) ReleaseAll(token common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.ReleaseAll(&_Dfa.TransactOpts, token)
}

// ReleaseAll0 is a paid mutator transaction binding the contract method 0x5be7fde8.
//
// Solidity: function releaseAll() returns()
func (_Dfa *DfaTransactor) ReleaseAll0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "releaseAll0")
}

// ReleaseAll0 is a paid mutator transaction binding the contract method 0x5be7fde8.
//
// Solidity: function releaseAll() returns()
func (_Dfa *DfaSession) ReleaseAll0() (*types.Transaction, error) {
	return _Dfa.Contract.ReleaseAll0(&_Dfa.TransactOpts)
}

// ReleaseAll0 is a paid mutator transaction binding the contract method 0x5be7fde8.
//
// Solidity: function releaseAll() returns()
func (_Dfa *DfaTransactorSession) ReleaseAll0() (*types.Transaction, error) {
	return _Dfa.Contract.ReleaseAll0(&_Dfa.TransactOpts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x2f59f741.
//
// Solidity: function removeFromWhitelist(uint8 stageId_, address[] accounts_) returns()
func (_Dfa *DfaTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, stageId_ uint8, accounts_ []common.Address) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "removeFromWhitelist", stageId_, accounts_)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x2f59f741.
//
// Solidity: function removeFromWhitelist(uint8 stageId_, address[] accounts_) returns()
func (_Dfa *DfaSession) RemoveFromWhitelist(stageId_ uint8, accounts_ []common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.RemoveFromWhitelist(&_Dfa.TransactOpts, stageId_, accounts_)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x2f59f741.
//
// Solidity: function removeFromWhitelist(uint8 stageId_, address[] accounts_) returns()
func (_Dfa *DfaTransactorSession) RemoveFromWhitelist(stageId_ uint8, accounts_ []common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.RemoveFromWhitelist(&_Dfa.TransactOpts, stageId_, accounts_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dfa *DfaTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dfa *DfaSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dfa.Contract.RenounceOwnership(&_Dfa.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dfa *DfaTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dfa.Contract.RenounceOwnership(&_Dfa.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Dfa *DfaTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Dfa *DfaSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.SafeTransferFrom(&_Dfa.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Dfa *DfaTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.SafeTransferFrom(&_Dfa.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Dfa *DfaTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Dfa *DfaSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Dfa.Contract.SafeTransferFrom0(&_Dfa.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Dfa *DfaTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Dfa.Contract.SafeTransferFrom0(&_Dfa.TransactOpts, from, to, tokenId, data)
}

// SetAdminPermissions is a paid mutator transaction binding the contract method 0x240ff27f.
//
// Solidity: function setAdminPermissions(address account_, bool enable_) returns()
func (_Dfa *DfaTransactor) SetAdminPermissions(opts *bind.TransactOpts, account_ common.Address, enable_ bool) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "setAdminPermissions", account_, enable_)
}

// SetAdminPermissions is a paid mutator transaction binding the contract method 0x240ff27f.
//
// Solidity: function setAdminPermissions(address account_, bool enable_) returns()
func (_Dfa *DfaSession) SetAdminPermissions(account_ common.Address, enable_ bool) (*types.Transaction, error) {
	return _Dfa.Contract.SetAdminPermissions(&_Dfa.TransactOpts, account_, enable_)
}

// SetAdminPermissions is a paid mutator transaction binding the contract method 0x240ff27f.
//
// Solidity: function setAdminPermissions(address account_, bool enable_) returns()
func (_Dfa *DfaTransactorSession) SetAdminPermissions(account_ common.Address, enable_ bool) (*types.Transaction, error) {
	return _Dfa.Contract.SetAdminPermissions(&_Dfa.TransactOpts, account_, enable_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dfa *DfaTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dfa *DfaSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dfa.Contract.SetApprovalForAll(&_Dfa.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dfa *DfaTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dfa.Contract.SetApprovalForAll(&_Dfa.TransactOpts, operator, approved)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_Dfa *DfaTransactor) SetMaxSupply(opts *bind.TransactOpts, maxSupply_ *big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "setMaxSupply", maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_Dfa *DfaSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.SetMaxSupply(&_Dfa.TransactOpts, maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_Dfa *DfaTransactorSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.SetMaxSupply(&_Dfa.TransactOpts, maxSupply_)
}

// SetPrefix is a paid mutator transaction binding the contract method 0x85cb593b.
//
// Solidity: function setPrefix(string prefix_) returns()
func (_Dfa *DfaTransactor) SetPrefix(opts *bind.TransactOpts, prefix_ string) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "setPrefix", prefix_)
}

// SetPrefix is a paid mutator transaction binding the contract method 0x85cb593b.
//
// Solidity: function setPrefix(string prefix_) returns()
func (_Dfa *DfaSession) SetPrefix(prefix_ string) (*types.Transaction, error) {
	return _Dfa.Contract.SetPrefix(&_Dfa.TransactOpts, prefix_)
}

// SetPrefix is a paid mutator transaction binding the contract method 0x85cb593b.
//
// Solidity: function setPrefix(string prefix_) returns()
func (_Dfa *DfaTransactorSession) SetPrefix(prefix_ string) (*types.Transaction, error) {
	return _Dfa.Contract.SetPrefix(&_Dfa.TransactOpts, prefix_)
}

// SetPrice is a paid mutator transaction binding the contract method 0x6a00670b.
//
// Solidity: function setPrice(uint8 stage_, uint256 value_) returns()
func (_Dfa *DfaTransactor) SetPrice(opts *bind.TransactOpts, stage_ uint8, value_ *big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "setPrice", stage_, value_)
}

// SetPrice is a paid mutator transaction binding the contract method 0x6a00670b.
//
// Solidity: function setPrice(uint8 stage_, uint256 value_) returns()
func (_Dfa *DfaSession) SetPrice(stage_ uint8, value_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.SetPrice(&_Dfa.TransactOpts, stage_, value_)
}

// SetPrice is a paid mutator transaction binding the contract method 0x6a00670b.
//
// Solidity: function setPrice(uint8 stage_, uint256 value_) returns()
func (_Dfa *DfaTransactorSession) SetPrice(stage_ uint8, value_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.SetPrice(&_Dfa.TransactOpts, stage_, value_)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x8c7ea24b.
//
// Solidity: function setRoyalties(address recipient_, uint256 amount_) returns()
func (_Dfa *DfaTransactor) SetRoyalties(opts *bind.TransactOpts, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "setRoyalties", recipient_, amount_)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x8c7ea24b.
//
// Solidity: function setRoyalties(address recipient_, uint256 amount_) returns()
func (_Dfa *DfaSession) SetRoyalties(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.SetRoyalties(&_Dfa.TransactOpts, recipient_, amount_)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x8c7ea24b.
//
// Solidity: function setRoyalties(address recipient_, uint256 amount_) returns()
func (_Dfa *DfaTransactorSession) SetRoyalties(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.SetRoyalties(&_Dfa.TransactOpts, recipient_, amount_)
}

// SetStage is a paid mutator transaction binding the contract method 0xce3cd997.
//
// Solidity: function setStage(uint8 stage_) returns()
func (_Dfa *DfaTransactor) SetStage(opts *bind.TransactOpts, stage_ uint8) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "setStage", stage_)
}

// SetStage is a paid mutator transaction binding the contract method 0xce3cd997.
//
// Solidity: function setStage(uint8 stage_) returns()
func (_Dfa *DfaSession) SetStage(stage_ uint8) (*types.Transaction, error) {
	return _Dfa.Contract.SetStage(&_Dfa.TransactOpts, stage_)
}

// SetStage is a paid mutator transaction binding the contract method 0xce3cd997.
//
// Solidity: function setStage(uint8 stage_) returns()
func (_Dfa *DfaTransactorSession) SetStage(stage_ uint8) (*types.Transaction, error) {
	return _Dfa.Contract.SetStage(&_Dfa.TransactOpts, stage_)
}

// SetSuffix is a paid mutator transaction binding the contract method 0x75d5ae9f.
//
// Solidity: function setSuffix(string suffix_) returns()
func (_Dfa *DfaTransactor) SetSuffix(opts *bind.TransactOpts, suffix_ string) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "setSuffix", suffix_)
}

// SetSuffix is a paid mutator transaction binding the contract method 0x75d5ae9f.
//
// Solidity: function setSuffix(string suffix_) returns()
func (_Dfa *DfaSession) SetSuffix(suffix_ string) (*types.Transaction, error) {
	return _Dfa.Contract.SetSuffix(&_Dfa.TransactOpts, suffix_)
}

// SetSuffix is a paid mutator transaction binding the contract method 0x75d5ae9f.
//
// Solidity: function setSuffix(string suffix_) returns()
func (_Dfa *DfaTransactorSession) SetSuffix(suffix_ string) (*types.Transaction, error) {
	return _Dfa.Contract.SetSuffix(&_Dfa.TransactOpts, suffix_)
}

// ToggleSoulBound is a paid mutator transaction binding the contract method 0x555684c7.
//
// Solidity: function toggleSoulBound(uint256 tokenId) returns()
func (_Dfa *DfaTransactor) ToggleSoulBound(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "toggleSoulBound", tokenId)
}

// ToggleSoulBound is a paid mutator transaction binding the contract method 0x555684c7.
//
// Solidity: function toggleSoulBound(uint256 tokenId) returns()
func (_Dfa *DfaSession) ToggleSoulBound(tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.ToggleSoulBound(&_Dfa.TransactOpts, tokenId)
}

// ToggleSoulBound is a paid mutator transaction binding the contract method 0x555684c7.
//
// Solidity: function toggleSoulBound(uint256 tokenId) returns()
func (_Dfa *DfaTransactorSession) ToggleSoulBound(tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.ToggleSoulBound(&_Dfa.TransactOpts, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Dfa *DfaTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Dfa *DfaSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.TransferFrom(&_Dfa.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Dfa *DfaTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.TransferFrom(&_Dfa.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dfa *DfaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dfa *DfaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.TransferOwnership(&_Dfa.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dfa *DfaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dfa.Contract.TransferOwnership(&_Dfa.TransactOpts, newOwner)
}

// UpdateBalanceLimit is a paid mutator transaction binding the contract method 0xc519cd1c.
//
// Solidity: function updateBalanceLimit(uint8 stageId_, uint256 limit_) returns()
func (_Dfa *DfaTransactor) UpdateBalanceLimit(opts *bind.TransactOpts, stageId_ uint8, limit_ *big.Int) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "updateBalanceLimit", stageId_, limit_)
}

// UpdateBalanceLimit is a paid mutator transaction binding the contract method 0xc519cd1c.
//
// Solidity: function updateBalanceLimit(uint8 stageId_, uint256 limit_) returns()
func (_Dfa *DfaSession) UpdateBalanceLimit(stageId_ uint8, limit_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.UpdateBalanceLimit(&_Dfa.TransactOpts, stageId_, limit_)
}

// UpdateBalanceLimit is a paid mutator transaction binding the contract method 0xc519cd1c.
//
// Solidity: function updateBalanceLimit(uint8 stageId_, uint256 limit_) returns()
func (_Dfa *DfaTransactorSession) UpdateBalanceLimit(stageId_ uint8, limit_ *big.Int) (*types.Transaction, error) {
	return _Dfa.Contract.UpdateBalanceLimit(&_Dfa.TransactOpts, stageId_, limit_)
}

// UpdateMerkleTreeRoot is a paid mutator transaction binding the contract method 0x5ee54e23.
//
// Solidity: function updateMerkleTreeRoot(uint8 stageId_, bytes32 merkleTreeRoot_) returns()
func (_Dfa *DfaTransactor) UpdateMerkleTreeRoot(opts *bind.TransactOpts, stageId_ uint8, merkleTreeRoot_ [32]byte) (*types.Transaction, error) {
	return _Dfa.contract.Transact(opts, "updateMerkleTreeRoot", stageId_, merkleTreeRoot_)
}

// UpdateMerkleTreeRoot is a paid mutator transaction binding the contract method 0x5ee54e23.
//
// Solidity: function updateMerkleTreeRoot(uint8 stageId_, bytes32 merkleTreeRoot_) returns()
func (_Dfa *DfaSession) UpdateMerkleTreeRoot(stageId_ uint8, merkleTreeRoot_ [32]byte) (*types.Transaction, error) {
	return _Dfa.Contract.UpdateMerkleTreeRoot(&_Dfa.TransactOpts, stageId_, merkleTreeRoot_)
}

// UpdateMerkleTreeRoot is a paid mutator transaction binding the contract method 0x5ee54e23.
//
// Solidity: function updateMerkleTreeRoot(uint8 stageId_, bytes32 merkleTreeRoot_) returns()
func (_Dfa *DfaTransactorSession) UpdateMerkleTreeRoot(stageId_ uint8, merkleTreeRoot_ [32]byte) (*types.Transaction, error) {
	return _Dfa.Contract.UpdateMerkleTreeRoot(&_Dfa.TransactOpts, stageId_, merkleTreeRoot_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dfa *DfaTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfa.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dfa *DfaSession) Receive() (*types.Transaction, error) {
	return _Dfa.Contract.Receive(&_Dfa.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dfa *DfaTransactorSession) Receive() (*types.Transaction, error) {
	return _Dfa.Contract.Receive(&_Dfa.TransactOpts)
}

// DfaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Dfa contract.
type DfaApprovalIterator struct {
	Event *DfaApproval // Event containing the contract specifics and raw log

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
func (it *DfaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaApproval)
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
		it.Event = new(DfaApproval)
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
func (it *DfaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaApproval represents a Approval event raised by the Dfa contract.
type DfaApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dfa *DfaFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DfaApprovalIterator, error) {

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

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DfaApprovalIterator{contract: _Dfa.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dfa *DfaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DfaApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaApproval)
				if err := _Dfa.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Dfa *DfaFilterer) ParseApproval(log types.Log) (*DfaApproval, error) {
	event := new(DfaApproval)
	if err := _Dfa.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Dfa contract.
type DfaApprovalForAllIterator struct {
	Event *DfaApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DfaApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaApprovalForAll)
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
		it.Event = new(DfaApprovalForAll)
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
func (it *DfaApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaApprovalForAll represents a ApprovalForAll event raised by the Dfa contract.
type DfaApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dfa *DfaFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DfaApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DfaApprovalForAllIterator{contract: _Dfa.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dfa *DfaFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DfaApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaApprovalForAll)
				if err := _Dfa.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Dfa *DfaFilterer) ParseApprovalForAll(log types.Log) (*DfaApprovalForAll, error) {
	event := new(DfaApprovalForAll)
	if err := _Dfa.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Dfa contract.
type DfaConsecutiveTransferIterator struct {
	Event *DfaConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *DfaConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaConsecutiveTransfer)
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
		it.Event = new(DfaConsecutiveTransfer)
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
func (it *DfaConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Dfa contract.
type DfaConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Dfa *DfaFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*DfaConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DfaConsecutiveTransferIterator{contract: _Dfa.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Dfa *DfaFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *DfaConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaConsecutiveTransfer)
				if err := _Dfa.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Dfa *DfaFilterer) ParseConsecutiveTransfer(log types.Log) (*DfaConsecutiveTransfer, error) {
	event := new(DfaConsecutiveTransfer)
	if err := _Dfa.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaERC20PaymentReleasedIterator is returned from FilterERC20PaymentReleased and is used to iterate over the raw logs and unpacked data for ERC20PaymentReleased events raised by the Dfa contract.
type DfaERC20PaymentReleasedIterator struct {
	Event *DfaERC20PaymentReleased // Event containing the contract specifics and raw log

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
func (it *DfaERC20PaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaERC20PaymentReleased)
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
		it.Event = new(DfaERC20PaymentReleased)
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
func (it *DfaERC20PaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaERC20PaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaERC20PaymentReleased represents a ERC20PaymentReleased event raised by the Dfa contract.
type DfaERC20PaymentReleased struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20PaymentReleased is a free log retrieval operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_Dfa *DfaFilterer) FilterERC20PaymentReleased(opts *bind.FilterOpts, token []common.Address) (*DfaERC20PaymentReleasedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "ERC20PaymentReleased", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DfaERC20PaymentReleasedIterator{contract: _Dfa.contract, event: "ERC20PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchERC20PaymentReleased is a free log subscription operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_Dfa *DfaFilterer) WatchERC20PaymentReleased(opts *bind.WatchOpts, sink chan<- *DfaERC20PaymentReleased, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "ERC20PaymentReleased", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaERC20PaymentReleased)
				if err := _Dfa.contract.UnpackLog(event, "ERC20PaymentReleased", log); err != nil {
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

// ParseERC20PaymentReleased is a log parse operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_Dfa *DfaFilterer) ParseERC20PaymentReleased(log types.Log) (*DfaERC20PaymentReleased, error) {
	event := new(DfaERC20PaymentReleased)
	if err := _Dfa.contract.UnpackLog(event, "ERC20PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dfa contract.
type DfaInitializedIterator struct {
	Event *DfaInitialized // Event containing the contract specifics and raw log

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
func (it *DfaInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaInitialized)
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
		it.Event = new(DfaInitialized)
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
func (it *DfaInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaInitialized represents a Initialized event raised by the Dfa contract.
type DfaInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dfa *DfaFilterer) FilterInitialized(opts *bind.FilterOpts) (*DfaInitializedIterator, error) {

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DfaInitializedIterator{contract: _Dfa.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dfa *DfaFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DfaInitialized) (event.Subscription, error) {

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaInitialized)
				if err := _Dfa.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dfa *DfaFilterer) ParseInitialized(log types.Log) (*DfaInitialized, error) {
	event := new(DfaInitialized)
	if err := _Dfa.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dfa contract.
type DfaOwnershipTransferredIterator struct {
	Event *DfaOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DfaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaOwnershipTransferred)
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
		it.Event = new(DfaOwnershipTransferred)
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
func (it *DfaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaOwnershipTransferred represents a OwnershipTransferred event raised by the Dfa contract.
type DfaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dfa *DfaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DfaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DfaOwnershipTransferredIterator{contract: _Dfa.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dfa *DfaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DfaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaOwnershipTransferred)
				if err := _Dfa.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dfa *DfaFilterer) ParseOwnershipTransferred(log types.Log) (*DfaOwnershipTransferred, error) {
	event := new(DfaOwnershipTransferred)
	if err := _Dfa.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaPayeeAddedIterator is returned from FilterPayeeAdded and is used to iterate over the raw logs and unpacked data for PayeeAdded events raised by the Dfa contract.
type DfaPayeeAddedIterator struct {
	Event *DfaPayeeAdded // Event containing the contract specifics and raw log

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
func (it *DfaPayeeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaPayeeAdded)
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
		it.Event = new(DfaPayeeAdded)
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
func (it *DfaPayeeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaPayeeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaPayeeAdded represents a PayeeAdded event raised by the Dfa contract.
type DfaPayeeAdded struct {
	Account common.Address
	Shares  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayeeAdded is a free log retrieval operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_Dfa *DfaFilterer) FilterPayeeAdded(opts *bind.FilterOpts) (*DfaPayeeAddedIterator, error) {

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "PayeeAdded")
	if err != nil {
		return nil, err
	}
	return &DfaPayeeAddedIterator{contract: _Dfa.contract, event: "PayeeAdded", logs: logs, sub: sub}, nil
}

// WatchPayeeAdded is a free log subscription operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_Dfa *DfaFilterer) WatchPayeeAdded(opts *bind.WatchOpts, sink chan<- *DfaPayeeAdded) (event.Subscription, error) {

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "PayeeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaPayeeAdded)
				if err := _Dfa.contract.UnpackLog(event, "PayeeAdded", log); err != nil {
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

// ParsePayeeAdded is a log parse operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_Dfa *DfaFilterer) ParsePayeeAdded(log types.Log) (*DfaPayeeAdded, error) {
	event := new(DfaPayeeAdded)
	if err := _Dfa.contract.UnpackLog(event, "PayeeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the Dfa contract.
type DfaPaymentReceivedIterator struct {
	Event *DfaPaymentReceived // Event containing the contract specifics and raw log

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
func (it *DfaPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaPaymentReceived)
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
		it.Event = new(DfaPaymentReceived)
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
func (it *DfaPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaPaymentReceived represents a PaymentReceived event raised by the Dfa contract.
type DfaPaymentReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Dfa *DfaFilterer) FilterPaymentReceived(opts *bind.FilterOpts) (*DfaPaymentReceivedIterator, error) {

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return &DfaPaymentReceivedIterator{contract: _Dfa.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Dfa *DfaFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *DfaPaymentReceived) (event.Subscription, error) {

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaPaymentReceived)
				if err := _Dfa.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
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

// ParsePaymentReceived is a log parse operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Dfa *DfaFilterer) ParsePaymentReceived(log types.Log) (*DfaPaymentReceived, error) {
	event := new(DfaPaymentReceived)
	if err := _Dfa.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaPaymentReleasedIterator is returned from FilterPaymentReleased and is used to iterate over the raw logs and unpacked data for PaymentReleased events raised by the Dfa contract.
type DfaPaymentReleasedIterator struct {
	Event *DfaPaymentReleased // Event containing the contract specifics and raw log

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
func (it *DfaPaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaPaymentReleased)
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
		it.Event = new(DfaPaymentReleased)
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
func (it *DfaPaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaPaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaPaymentReleased represents a PaymentReleased event raised by the Dfa contract.
type DfaPaymentReleased struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReleased is a free log retrieval operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Dfa *DfaFilterer) FilterPaymentReleased(opts *bind.FilterOpts) (*DfaPaymentReleasedIterator, error) {

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return &DfaPaymentReleasedIterator{contract: _Dfa.contract, event: "PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchPaymentReleased is a free log subscription operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Dfa *DfaFilterer) WatchPaymentReleased(opts *bind.WatchOpts, sink chan<- *DfaPaymentReleased) (event.Subscription, error) {

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaPaymentReleased)
				if err := _Dfa.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
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

// ParsePaymentReleased is a log parse operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Dfa *DfaFilterer) ParsePaymentReleased(log types.Log) (*DfaPaymentReleased, error) {
	event := new(DfaPaymentReleased)
	if err := _Dfa.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaRoyaltiesSetIterator is returned from FilterRoyaltiesSet and is used to iterate over the raw logs and unpacked data for RoyaltiesSet events raised by the Dfa contract.
type DfaRoyaltiesSetIterator struct {
	Event *DfaRoyaltiesSet // Event containing the contract specifics and raw log

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
func (it *DfaRoyaltiesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaRoyaltiesSet)
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
		it.Event = new(DfaRoyaltiesSet)
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
func (it *DfaRoyaltiesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaRoyaltiesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaRoyaltiesSet represents a RoyaltiesSet event raised by the Dfa contract.
type DfaRoyaltiesSet struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRoyaltiesSet is a free log retrieval operation binding the contract event 0x908669f35f6fb3977a956ba70597841fe541d1e8491ca3c025161e258d3bfdb6.
//
// Solidity: event RoyaltiesSet(address recipient, uint256 amount)
func (_Dfa *DfaFilterer) FilterRoyaltiesSet(opts *bind.FilterOpts) (*DfaRoyaltiesSetIterator, error) {

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "RoyaltiesSet")
	if err != nil {
		return nil, err
	}
	return &DfaRoyaltiesSetIterator{contract: _Dfa.contract, event: "RoyaltiesSet", logs: logs, sub: sub}, nil
}

// WatchRoyaltiesSet is a free log subscription operation binding the contract event 0x908669f35f6fb3977a956ba70597841fe541d1e8491ca3c025161e258d3bfdb6.
//
// Solidity: event RoyaltiesSet(address recipient, uint256 amount)
func (_Dfa *DfaFilterer) WatchRoyaltiesSet(opts *bind.WatchOpts, sink chan<- *DfaRoyaltiesSet) (event.Subscription, error) {

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "RoyaltiesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaRoyaltiesSet)
				if err := _Dfa.contract.UnpackLog(event, "RoyaltiesSet", log); err != nil {
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

// ParseRoyaltiesSet is a log parse operation binding the contract event 0x908669f35f6fb3977a956ba70597841fe541d1e8491ca3c025161e258d3bfdb6.
//
// Solidity: event RoyaltiesSet(address recipient, uint256 amount)
func (_Dfa *DfaFilterer) ParseRoyaltiesSet(log types.Log) (*DfaRoyaltiesSet, error) {
	event := new(DfaRoyaltiesSet)
	if err := _Dfa.contract.UnpackLog(event, "RoyaltiesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dfa contract.
type DfaTransferIterator struct {
	Event *DfaTransfer // Event containing the contract specifics and raw log

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
func (it *DfaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfaTransfer)
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
		it.Event = new(DfaTransfer)
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
func (it *DfaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfaTransfer represents a Transfer event raised by the Dfa contract.
type DfaTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dfa *DfaFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DfaTransferIterator, error) {

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

	logs, sub, err := _Dfa.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DfaTransferIterator{contract: _Dfa.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dfa *DfaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DfaTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Dfa.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfaTransfer)
				if err := _Dfa.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Dfa *DfaFilterer) ParseTransfer(log types.Log) (*DfaTransfer, error) {
	event := new(DfaTransfer)
	if err := _Dfa.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
