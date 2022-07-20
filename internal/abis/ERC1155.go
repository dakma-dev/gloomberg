// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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

// ERC1155MetaData contains all meta data concerning the ERC1155 contract.
var ERC1155MetaData = &bind.MetaData{
	ABI: "[{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_block\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_storage\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_oldContract\",\"type\":\"address\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_block\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_nextContract\",\"type\":\"address\"}],\"name\":\"Retire\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UpdateDecimals\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UpdateName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UpdateSymbol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"SetURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"Assign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"AcceptAssignment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_isNonFungible\",\"type\":\"bool\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxMeltFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UpdateMeltFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"OperatorApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_feeId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_feeValue\",\"type\":\"uint256\"}],\"name\":\"TransferFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxTransferFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UpdateTransferable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UpdateTransferFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_whitelisted\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_on\",\"type\":\"bool\"}],\"name\":\"Whitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Melt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"DeployERCAdapter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_tradeId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_firstParty\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_secondParty\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_escrowedEnjFirstParty\",\"type\":\"uint256\"}],\"name\":\"CreateTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_tradeId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_firstParty\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_secondParty\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_receivedEnjFirstParty\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_changeEnjFirstParty\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_receivedEnjSecondParty\",\"type\":\"uint256\"}],\"name\":\"CompleteTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_tradeId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_firstParty\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_receivedEnjFirstParty\",\"type\":\"uint256\"}],\"name\":\"CancelTrade\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"name\":\"_initialReserve\",\"type\":\"uint256\"},{\"name\":\"_supplyModel\",\"type\":\"address\"},{\"name\":\"_meltValue\",\"type\":\"uint256\"},{\"name\":\"_meltFeeRatio\",\"type\":\"uint16\"},{\"name\":\"_transferable\",\"type\":\"uint8\"},{\"name\":\"_transferFeeSettings\",\"type\":\"uint256[3]\"},{\"name\":\"_nonFungible\",\"type\":\"bool\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_initialReserve\",\"type\":\"uint256\"}],\"name\":\"minMeltValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"mintFungibles\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address[]\"}],\"name\":\"mintNonFungibles\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address[]\"},{\"name\":\"_data\",\"type\":\"uint128[]\"}],\"name\":\"mintNonFungiblesWithData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"reserve\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint128\"}],\"name\":\"releaseReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"updateName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"assign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"acceptAssignment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_whitelisted\",\"type\":\"address\"},{\"name\":\"_on\",\"type\":\"bool\"}],\"name\":\"setWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_transferable\",\"type\":\"uint8\"}],\"name\":\"setTransferable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint16\"}],\"name\":\"setMeltFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint16\"}],\"name\":\"decreaseMaxMeltFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setTransferFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"decreaseMaxTransferFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"deployERC20Adapter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"deployERC721Adapter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"addLog\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"typeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"typeByIndex\",\"outputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleTypeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"nonFungibleTypeByIndex\",\"outputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fungibleTypeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"fungibleTypeByIndex\",\"outputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"typeData\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_creator\",\"type\":\"address\"},{\"name\":\"_meltValue\",\"type\":\"uint256\"},{\"name\":\"_meltFeeRatio\",\"type\":\"uint16\"},{\"name\":\"_meltFeeMaxRatio\",\"type\":\"uint16\"},{\"name\":\"_supplyModel\",\"type\":\"address\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"name\":\"_circulatingSupply\",\"type\":\"uint256\"},{\"name\":\"_reserve\",\"type\":\"uint256\"},{\"name\":\"_transferable\",\"type\":\"uint8\"},{\"name\":\"_transferFeeData\",\"type\":\"uint256[4]\"},{\"name\":\"_nonFungible\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"transferSettings\",\"outputs\":[{\"name\":\"_transferable\",\"type\":\"uint8\"},{\"name\":\"_transferFeeType\",\"type\":\"uint8\"},{\"name\":\"_transferFeeCurrency\",\"type\":\"uint256\"},{\"name\":\"_transferFeeValue\",\"type\":\"uint256\"},{\"name\":\"_transferFeeMaxValue\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"isCreatorOf\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_whitelisted\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"mintableSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"circulatingSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getERC20Adapter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getERC721Adapter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_msgSender\",\"type\":\"address\"}],\"name\":\"transferAdapter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_msgSender\",\"type\":\"address\"}],\"name\":\"transferFromAdapter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address[]\"},{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"multicastTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address[]\"},{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeMulticastTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address[]\"},{\"name\":\"_to\",\"type\":\"address[]\"},{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"multicastTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_currentValue\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_currentValue\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_msgSender\",\"type\":\"address\"}],\"name\":\"approveAdapter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_currentValues\",\"type\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"batchApprove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_approved\",\"type\":\"bool\"},{\"name\":\"_msgSender\",\"type\":\"address\"}],\"name\":\"setApprovalAdapter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferFees\",\"outputs\":[{\"name\":\"_transferValue\",\"type\":\"uint256\"},{\"name\":\"_minTransferValue\",\"type\":\"uint256\"},{\"name\":\"_transferFeeCurrency\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_maxFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_askingIds\",\"type\":\"uint256[]\"},{\"name\":\"_askingValues\",\"type\":\"uint128[]\"},{\"name\":\"_offeringIds\",\"type\":\"uint256[]\"},{\"name\":\"_offeringValues\",\"type\":\"uint128[]\"},{\"name\":\"_secondParty\",\"type\":\"address\"}],\"name\":\"createTrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"tradeCompletable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"completeTrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"cancelTrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"melt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"nonFungibleCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"nonFungibleByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"nonFungibleOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isNonFungible\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_erc20ContractAddress\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"releaseERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"releaseETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_erc721ContractAddress\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"uint256\"}],\"name\":\"releaseERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_erc1155ContractAddress\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"releaseERC1155\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storage\",\"type\":\"address\"},{\"name\":\"_oldContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nextContract\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155MetaData.ABI instead.
var ERC1155ABI = ERC1155MetaData.ABI

// ERC1155 is an auto generated Go binding around an Ethereum contract.
type ERC1155 struct {
	ERC1155Caller     // Read-only binding to the contract
	ERC1155Transactor // Write-only binding to the contract
	ERC1155Filterer   // Log filterer for contract events
}

// ERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155Session struct {
	Contract     *ERC1155          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155CallerSession struct {
	Contract *ERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155TransactorSession struct {
	Contract     *ERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155Raw struct {
	Contract *ERC1155 // Generic contract binding to access the raw methods on
}

// ERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155CallerRaw struct {
	Contract *ERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155TransactorRaw struct {
	Contract *ERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155 creates a new instance of ERC1155, bound to a specific deployed contract.
func NewERC1155(address common.Address, backend bind.ContractBackend) (*ERC1155, error) {
	contract, err := bindERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155{ERC1155Caller: ERC1155Caller{contract: contract}, ERC1155Transactor: ERC1155Transactor{contract: contract}, ERC1155Filterer: ERC1155Filterer{contract: contract}}, nil
}

// NewERC1155Caller creates a new read-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Caller(address common.Address, caller bind.ContractCaller) (*ERC1155Caller, error) {
	contract, err := bindERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Caller{contract: contract}, nil
}

// NewERC1155Transactor creates a new write-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155Transactor, error) {
	contract, err := bindERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Transactor{contract: contract}, nil
}

// NewERC1155Filterer creates a new log filterer instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155Filterer, error) {
	contract, err := bindERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155Filterer{contract: contract}, nil
}

// bindERC1155 binds a generic wrapper to an already deployed contract.
func bindERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.ERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x0d550b75.
//
// Solidity: function allowance(uint256 _id, address _owner, address _spender) view returns(uint256)
func (_ERC1155 *ERC1155Caller) Allowance(opts *bind.CallOpts, _id *big.Int, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "allowance", _id, _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x0d550b75.
//
// Solidity: function allowance(uint256 _id, address _owner, address _spender) view returns(uint256)
func (_ERC1155 *ERC1155Session) Allowance(_id *big.Int, _owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC1155.Contract.Allowance(&_ERC1155.CallOpts, _id, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0x0d550b75.
//
// Solidity: function allowance(uint256 _id, address _owner, address _spender) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) Allowance(_id *big.Int, _owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC1155.Contract.Allowance(&_ERC1155.CallOpts, _id, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x3656eec2.
//
// Solidity: function balanceOf(uint256 _id, address _owner) view returns(uint256)
func (_ERC1155 *ERC1155Caller) BalanceOf(opts *bind.CallOpts, _id *big.Int, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOf", _id, _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x3656eec2.
//
// Solidity: function balanceOf(uint256 _id, address _owner) view returns(uint256)
func (_ERC1155 *ERC1155Session) BalanceOf(_id *big.Int, _owner common.Address) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, _id, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x3656eec2.
//
// Solidity: function balanceOf(uint256 _id, address _owner) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) BalanceOf(_id *big.Int, _owner common.Address) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, _id, _owner)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x92ff6aea.
//
// Solidity: function circulatingSupply(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) CirculatingSupply(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "circulatingSupply", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CirculatingSupply is a free data retrieval call binding the contract method 0x92ff6aea.
//
// Solidity: function circulatingSupply(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Session) CirculatingSupply(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.CirculatingSupply(&_ERC1155.CallOpts, _id)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x92ff6aea.
//
// Solidity: function circulatingSupply(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) CirculatingSupply(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.CirculatingSupply(&_ERC1155.CallOpts, _id)
}

// Decimals is a free data retrieval call binding the contract method 0x3f47e662.
//
// Solidity: function decimals(uint256 _id) view returns(uint8)
func (_ERC1155 *ERC1155Caller) Decimals(opts *bind.CallOpts, _id *big.Int) (uint8, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "decimals", _id)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x3f47e662.
//
// Solidity: function decimals(uint256 _id) view returns(uint8)
func (_ERC1155 *ERC1155Session) Decimals(_id *big.Int) (uint8, error) {
	return _ERC1155.Contract.Decimals(&_ERC1155.CallOpts, _id)
}

// Decimals is a free data retrieval call binding the contract method 0x3f47e662.
//
// Solidity: function decimals(uint256 _id) view returns(uint8)
func (_ERC1155 *ERC1155CallerSession) Decimals(_id *big.Int) (uint8, error) {
	return _ERC1155.Contract.Decimals(&_ERC1155.CallOpts, _id)
}

// FungibleTypeByIndex is a free data retrieval call binding the contract method 0x230df5d7.
//
// Solidity: function fungibleTypeByIndex(uint256 _index) view returns(uint256 _id)
func (_ERC1155 *ERC1155Caller) FungibleTypeByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "fungibleTypeByIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FungibleTypeByIndex is a free data retrieval call binding the contract method 0x230df5d7.
//
// Solidity: function fungibleTypeByIndex(uint256 _index) view returns(uint256 _id)
func (_ERC1155 *ERC1155Session) FungibleTypeByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.FungibleTypeByIndex(&_ERC1155.CallOpts, _index)
}

// FungibleTypeByIndex is a free data retrieval call binding the contract method 0x230df5d7.
//
// Solidity: function fungibleTypeByIndex(uint256 _index) view returns(uint256 _id)
func (_ERC1155 *ERC1155CallerSession) FungibleTypeByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.FungibleTypeByIndex(&_ERC1155.CallOpts, _index)
}

// FungibleTypeCount is a free data retrieval call binding the contract method 0x2827616f.
//
// Solidity: function fungibleTypeCount() view returns(uint256)
func (_ERC1155 *ERC1155Caller) FungibleTypeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "fungibleTypeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FungibleTypeCount is a free data retrieval call binding the contract method 0x2827616f.
//
// Solidity: function fungibleTypeCount() view returns(uint256)
func (_ERC1155 *ERC1155Session) FungibleTypeCount() (*big.Int, error) {
	return _ERC1155.Contract.FungibleTypeCount(&_ERC1155.CallOpts)
}

// FungibleTypeCount is a free data retrieval call binding the contract method 0x2827616f.
//
// Solidity: function fungibleTypeCount() view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) FungibleTypeCount() (*big.Int, error) {
	return _ERC1155.Contract.FungibleTypeCount(&_ERC1155.CallOpts)
}

// GetERC20Adapter is a free data retrieval call binding the contract method 0x82f1dc3a.
//
// Solidity: function getERC20Adapter(uint256 _id) view returns(address)
func (_ERC1155 *ERC1155Caller) GetERC20Adapter(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "getERC20Adapter", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetERC20Adapter is a free data retrieval call binding the contract method 0x82f1dc3a.
//
// Solidity: function getERC20Adapter(uint256 _id) view returns(address)
func (_ERC1155 *ERC1155Session) GetERC20Adapter(_id *big.Int) (common.Address, error) {
	return _ERC1155.Contract.GetERC20Adapter(&_ERC1155.CallOpts, _id)
}

// GetERC20Adapter is a free data retrieval call binding the contract method 0x82f1dc3a.
//
// Solidity: function getERC20Adapter(uint256 _id) view returns(address)
func (_ERC1155 *ERC1155CallerSession) GetERC20Adapter(_id *big.Int) (common.Address, error) {
	return _ERC1155.Contract.GetERC20Adapter(&_ERC1155.CallOpts, _id)
}

// GetERC721Adapter is a free data retrieval call binding the contract method 0xade46d12.
//
// Solidity: function getERC721Adapter(uint256 _id) view returns(address)
func (_ERC1155 *ERC1155Caller) GetERC721Adapter(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "getERC721Adapter", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetERC721Adapter is a free data retrieval call binding the contract method 0xade46d12.
//
// Solidity: function getERC721Adapter(uint256 _id) view returns(address)
func (_ERC1155 *ERC1155Session) GetERC721Adapter(_id *big.Int) (common.Address, error) {
	return _ERC1155.Contract.GetERC721Adapter(&_ERC1155.CallOpts, _id)
}

// GetERC721Adapter is a free data retrieval call binding the contract method 0xade46d12.
//
// Solidity: function getERC721Adapter(uint256 _id) view returns(address)
func (_ERC1155 *ERC1155CallerSession) GetERC721Adapter(_id *big.Int) (common.Address, error) {
	return _ERC1155.Contract.GetERC721Adapter(&_ERC1155.CallOpts, _id)
}

// IsApproved is a free data retrieval call binding the contract method 0xe5af48d8.
//
// Solidity: function isApproved(address _owner, address _operator, uint256 _id) view returns(bool)
func (_ERC1155 *ERC1155Caller) IsApproved(opts *bind.CallOpts, _owner common.Address, _operator common.Address, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isApproved", _owner, _operator, _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0xe5af48d8.
//
// Solidity: function isApproved(address _owner, address _operator, uint256 _id) view returns(bool)
func (_ERC1155 *ERC1155Session) IsApproved(_owner common.Address, _operator common.Address, _id *big.Int) (bool, error) {
	return _ERC1155.Contract.IsApproved(&_ERC1155.CallOpts, _owner, _operator, _id)
}

// IsApproved is a free data retrieval call binding the contract method 0xe5af48d8.
//
// Solidity: function isApproved(address _owner, address _operator, uint256 _id) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsApproved(_owner common.Address, _operator common.Address, _id *big.Int) (bool, error) {
	return _ERC1155.Contract.IsApproved(&_ERC1155.CallOpts, _owner, _operator, _id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_ERC1155 *ERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_ERC1155 *ERC1155Session) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, _owner, _operator)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address _addr) view returns(bool)
func (_ERC1155 *ERC1155Caller) IsContract(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isContract", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address _addr) view returns(bool)
func (_ERC1155 *ERC1155Session) IsContract(_addr common.Address) (bool, error) {
	return _ERC1155.Contract.IsContract(&_ERC1155.CallOpts, _addr)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address _addr) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsContract(_addr common.Address) (bool, error) {
	return _ERC1155.Contract.IsContract(&_ERC1155.CallOpts, _addr)
}

// IsCreatorOf is a free data retrieval call binding the contract method 0xea6df23f.
//
// Solidity: function isCreatorOf(uint256 _id, address _creator) view returns(bool)
func (_ERC1155 *ERC1155Caller) IsCreatorOf(opts *bind.CallOpts, _id *big.Int, _creator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isCreatorOf", _id, _creator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCreatorOf is a free data retrieval call binding the contract method 0xea6df23f.
//
// Solidity: function isCreatorOf(uint256 _id, address _creator) view returns(bool)
func (_ERC1155 *ERC1155Session) IsCreatorOf(_id *big.Int, _creator common.Address) (bool, error) {
	return _ERC1155.Contract.IsCreatorOf(&_ERC1155.CallOpts, _id, _creator)
}

// IsCreatorOf is a free data retrieval call binding the contract method 0xea6df23f.
//
// Solidity: function isCreatorOf(uint256 _id, address _creator) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsCreatorOf(_id *big.Int, _creator common.Address) (bool, error) {
	return _ERC1155.Contract.IsCreatorOf(&_ERC1155.CallOpts, _id, _creator)
}

// IsNonFungible is a free data retrieval call binding the contract method 0xe44591f0.
//
// Solidity: function isNonFungible(uint256 _id) pure returns(bool)
func (_ERC1155 *ERC1155Caller) IsNonFungible(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isNonFungible", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonFungible is a free data retrieval call binding the contract method 0xe44591f0.
//
// Solidity: function isNonFungible(uint256 _id) pure returns(bool)
func (_ERC1155 *ERC1155Session) IsNonFungible(_id *big.Int) (bool, error) {
	return _ERC1155.Contract.IsNonFungible(&_ERC1155.CallOpts, _id)
}

// IsNonFungible is a free data retrieval call binding the contract method 0xe44591f0.
//
// Solidity: function isNonFungible(uint256 _id) pure returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsNonFungible(_id *big.Int) (bool, error) {
	return _ERC1155.Contract.IsNonFungible(&_ERC1155.CallOpts, _id)
}

// MinMeltValue is a free data retrieval call binding the contract method 0x11f7e404.
//
// Solidity: function minMeltValue(uint256 _initialReserve) view returns(uint256)
func (_ERC1155 *ERC1155Caller) MinMeltValue(opts *bind.CallOpts, _initialReserve *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "minMeltValue", _initialReserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinMeltValue is a free data retrieval call binding the contract method 0x11f7e404.
//
// Solidity: function minMeltValue(uint256 _initialReserve) view returns(uint256)
func (_ERC1155 *ERC1155Session) MinMeltValue(_initialReserve *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.MinMeltValue(&_ERC1155.CallOpts, _initialReserve)
}

// MinMeltValue is a free data retrieval call binding the contract method 0x11f7e404.
//
// Solidity: function minMeltValue(uint256 _initialReserve) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) MinMeltValue(_initialReserve *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.MinMeltValue(&_ERC1155.CallOpts, _initialReserve)
}

// MintableSupply is a free data retrieval call binding the contract method 0x443bf984.
//
// Solidity: function mintableSupply(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) MintableSupply(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "mintableSupply", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintableSupply is a free data retrieval call binding the contract method 0x443bf984.
//
// Solidity: function mintableSupply(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Session) MintableSupply(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.MintableSupply(&_ERC1155.CallOpts, _id)
}

// MintableSupply is a free data retrieval call binding the contract method 0x443bf984.
//
// Solidity: function mintableSupply(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) MintableSupply(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.MintableSupply(&_ERC1155.CallOpts, _id)
}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 _id) view returns(string)
func (_ERC1155 *ERC1155Caller) Name(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "name", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 _id) view returns(string)
func (_ERC1155 *ERC1155Session) Name(_id *big.Int) (string, error) {
	return _ERC1155.Contract.Name(&_ERC1155.CallOpts, _id)
}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 _id) view returns(string)
func (_ERC1155 *ERC1155CallerSession) Name(_id *big.Int) (string, error) {
	return _ERC1155.Contract.Name(&_ERC1155.CallOpts, _id)
}

// NonFungibleByIndex is a free data retrieval call binding the contract method 0xd56e9c5a.
//
// Solidity: function nonFungibleByIndex(uint256 _id, uint256 _index) view returns(uint256)
func (_ERC1155 *ERC1155Caller) NonFungibleByIndex(opts *bind.CallOpts, _id *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "nonFungibleByIndex", _id, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonFungibleByIndex is a free data retrieval call binding the contract method 0xd56e9c5a.
//
// Solidity: function nonFungibleByIndex(uint256 _id, uint256 _index) view returns(uint256)
func (_ERC1155 *ERC1155Session) NonFungibleByIndex(_id *big.Int, _index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleByIndex(&_ERC1155.CallOpts, _id, _index)
}

// NonFungibleByIndex is a free data retrieval call binding the contract method 0xd56e9c5a.
//
// Solidity: function nonFungibleByIndex(uint256 _id, uint256 _index) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) NonFungibleByIndex(_id *big.Int, _index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleByIndex(&_ERC1155.CallOpts, _id, _index)
}

// NonFungibleCount is a free data retrieval call binding the contract method 0x11f108f6.
//
// Solidity: function nonFungibleCount(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) NonFungibleCount(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "nonFungibleCount", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonFungibleCount is a free data retrieval call binding the contract method 0x11f108f6.
//
// Solidity: function nonFungibleCount(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Session) NonFungibleCount(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleCount(&_ERC1155.CallOpts, _id)
}

// NonFungibleCount is a free data retrieval call binding the contract method 0x11f108f6.
//
// Solidity: function nonFungibleCount(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) NonFungibleCount(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleCount(&_ERC1155.CallOpts, _id)
}

// NonFungibleOfOwnerByIndex is a free data retrieval call binding the contract method 0x7520d98c.
//
// Solidity: function nonFungibleOfOwnerByIndex(uint256 _id, address _owner, uint256 _index) view returns(uint256)
func (_ERC1155 *ERC1155Caller) NonFungibleOfOwnerByIndex(opts *bind.CallOpts, _id *big.Int, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "nonFungibleOfOwnerByIndex", _id, _owner, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonFungibleOfOwnerByIndex is a free data retrieval call binding the contract method 0x7520d98c.
//
// Solidity: function nonFungibleOfOwnerByIndex(uint256 _id, address _owner, uint256 _index) view returns(uint256)
func (_ERC1155 *ERC1155Session) NonFungibleOfOwnerByIndex(_id *big.Int, _owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleOfOwnerByIndex(&_ERC1155.CallOpts, _id, _owner, _index)
}

// NonFungibleOfOwnerByIndex is a free data retrieval call binding the contract method 0x7520d98c.
//
// Solidity: function nonFungibleOfOwnerByIndex(uint256 _id, address _owner, uint256 _index) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) NonFungibleOfOwnerByIndex(_id *big.Int, _owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleOfOwnerByIndex(&_ERC1155.CallOpts, _id, _owner, _index)
}

// NonFungibleTypeByIndex is a free data retrieval call binding the contract method 0x2da3c83b.
//
// Solidity: function nonFungibleTypeByIndex(uint256 _index) view returns(uint256 _id)
func (_ERC1155 *ERC1155Caller) NonFungibleTypeByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "nonFungibleTypeByIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonFungibleTypeByIndex is a free data retrieval call binding the contract method 0x2da3c83b.
//
// Solidity: function nonFungibleTypeByIndex(uint256 _index) view returns(uint256 _id)
func (_ERC1155 *ERC1155Session) NonFungibleTypeByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleTypeByIndex(&_ERC1155.CallOpts, _index)
}

// NonFungibleTypeByIndex is a free data retrieval call binding the contract method 0x2da3c83b.
//
// Solidity: function nonFungibleTypeByIndex(uint256 _index) view returns(uint256 _id)
func (_ERC1155 *ERC1155CallerSession) NonFungibleTypeByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleTypeByIndex(&_ERC1155.CallOpts, _index)
}

// NonFungibleTypeCount is a free data retrieval call binding the contract method 0xe94361d7.
//
// Solidity: function nonFungibleTypeCount() view returns(uint256)
func (_ERC1155 *ERC1155Caller) NonFungibleTypeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "nonFungibleTypeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonFungibleTypeCount is a free data retrieval call binding the contract method 0xe94361d7.
//
// Solidity: function nonFungibleTypeCount() view returns(uint256)
func (_ERC1155 *ERC1155Session) NonFungibleTypeCount() (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleTypeCount(&_ERC1155.CallOpts)
}

// NonFungibleTypeCount is a free data retrieval call binding the contract method 0xe94361d7.
//
// Solidity: function nonFungibleTypeCount() view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) NonFungibleTypeCount() (*big.Int, error) {
	return _ERC1155.Contract.NonFungibleTypeCount(&_ERC1155.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _id) view returns(address)
func (_ERC1155 *ERC1155Caller) OwnerOf(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "ownerOf", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _id) view returns(address)
func (_ERC1155 *ERC1155Session) OwnerOf(_id *big.Int) (common.Address, error) {
	return _ERC1155.Contract.OwnerOf(&_ERC1155.CallOpts, _id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _id) view returns(address)
func (_ERC1155 *ERC1155CallerSession) OwnerOf(_id *big.Int) (common.Address, error) {
	return _ERC1155.Contract.OwnerOf(&_ERC1155.CallOpts, _id)
}

// Reserve is a free data retrieval call binding the contract method 0x819b25ba.
//
// Solidity: function reserve(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) Reserve(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "reserve", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve is a free data retrieval call binding the contract method 0x819b25ba.
//
// Solidity: function reserve(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Session) Reserve(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.Reserve(&_ERC1155.CallOpts, _id)
}

// Reserve is a free data retrieval call binding the contract method 0x819b25ba.
//
// Solidity: function reserve(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) Reserve(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.Reserve(&_ERC1155.CallOpts, _id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_ERC1155 *ERC1155Caller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_ERC1155 *ERC1155Session) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_ERC1155 *ERC1155CallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x4e41a1fb.
//
// Solidity: function symbol(uint256 _id) view returns(string)
func (_ERC1155 *ERC1155Caller) Symbol(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "symbol", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x4e41a1fb.
//
// Solidity: function symbol(uint256 _id) view returns(string)
func (_ERC1155 *ERC1155Session) Symbol(_id *big.Int) (string, error) {
	return _ERC1155.Contract.Symbol(&_ERC1155.CallOpts, _id)
}

// Symbol is a free data retrieval call binding the contract method 0x4e41a1fb.
//
// Solidity: function symbol(uint256 _id) view returns(string)
func (_ERC1155 *ERC1155CallerSession) Symbol(_id *big.Int) (string, error) {
	return _ERC1155.Contract.Symbol(&_ERC1155.CallOpts, _id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) TotalSupply(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "totalSupply", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Session) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.TotalSupply(&_ERC1155.CallOpts, _id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.TotalSupply(&_ERC1155.CallOpts, _id)
}

// TradeCompletable is a free data retrieval call binding the contract method 0xb8f7e585.
//
// Solidity: function tradeCompletable(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) TradeCompletable(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "tradeCompletable", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradeCompletable is a free data retrieval call binding the contract method 0xb8f7e585.
//
// Solidity: function tradeCompletable(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155Session) TradeCompletable(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.TradeCompletable(&_ERC1155.CallOpts, _id)
}

// TradeCompletable is a free data retrieval call binding the contract method 0xb8f7e585.
//
// Solidity: function tradeCompletable(uint256 _id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) TradeCompletable(_id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.TradeCompletable(&_ERC1155.CallOpts, _id)
}

// TransferFees is a free data retrieval call binding the contract method 0x12ab550f.
//
// Solidity: function transferFees(uint256 _id, uint256 _value, address _from, address _to) view returns(uint256 _transferValue, uint256 _minTransferValue, uint256 _transferFeeCurrency, uint256 _fee, uint256 _maxFee)
func (_ERC1155 *ERC1155Caller) TransferFees(opts *bind.CallOpts, _id *big.Int, _value *big.Int, _from common.Address, _to common.Address) (struct {
	TransferValue       *big.Int
	MinTransferValue    *big.Int
	TransferFeeCurrency *big.Int
	Fee                 *big.Int
	MaxFee              *big.Int
}, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "transferFees", _id, _value, _from, _to)

	outstruct := new(struct {
		TransferValue       *big.Int
		MinTransferValue    *big.Int
		TransferFeeCurrency *big.Int
		Fee                 *big.Int
		MaxFee              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TransferValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MinTransferValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TransferFeeCurrency = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TransferFees is a free data retrieval call binding the contract method 0x12ab550f.
//
// Solidity: function transferFees(uint256 _id, uint256 _value, address _from, address _to) view returns(uint256 _transferValue, uint256 _minTransferValue, uint256 _transferFeeCurrency, uint256 _fee, uint256 _maxFee)
func (_ERC1155 *ERC1155Session) TransferFees(_id *big.Int, _value *big.Int, _from common.Address, _to common.Address) (struct {
	TransferValue       *big.Int
	MinTransferValue    *big.Int
	TransferFeeCurrency *big.Int
	Fee                 *big.Int
	MaxFee              *big.Int
}, error) {
	return _ERC1155.Contract.TransferFees(&_ERC1155.CallOpts, _id, _value, _from, _to)
}

// TransferFees is a free data retrieval call binding the contract method 0x12ab550f.
//
// Solidity: function transferFees(uint256 _id, uint256 _value, address _from, address _to) view returns(uint256 _transferValue, uint256 _minTransferValue, uint256 _transferFeeCurrency, uint256 _fee, uint256 _maxFee)
func (_ERC1155 *ERC1155CallerSession) TransferFees(_id *big.Int, _value *big.Int, _from common.Address, _to common.Address) (struct {
	TransferValue       *big.Int
	MinTransferValue    *big.Int
	TransferFeeCurrency *big.Int
	Fee                 *big.Int
	MaxFee              *big.Int
}, error) {
	return _ERC1155.Contract.TransferFees(&_ERC1155.CallOpts, _id, _value, _from, _to)
}

// TransferSettings is a free data retrieval call binding the contract method 0x6683de90.
//
// Solidity: function transferSettings(uint256 _id) view returns(uint8 _transferable, uint8 _transferFeeType, uint256 _transferFeeCurrency, uint256 _transferFeeValue, uint256 _transferFeeMaxValue)
func (_ERC1155 *ERC1155Caller) TransferSettings(opts *bind.CallOpts, _id *big.Int) (struct {
	Transferable        uint8
	TransferFeeType     uint8
	TransferFeeCurrency *big.Int
	TransferFeeValue    *big.Int
	TransferFeeMaxValue *big.Int
}, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "transferSettings", _id)

	outstruct := new(struct {
		Transferable        uint8
		TransferFeeType     uint8
		TransferFeeCurrency *big.Int
		TransferFeeValue    *big.Int
		TransferFeeMaxValue *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Transferable = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.TransferFeeType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.TransferFeeCurrency = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TransferFeeValue = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TransferFeeMaxValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TransferSettings is a free data retrieval call binding the contract method 0x6683de90.
//
// Solidity: function transferSettings(uint256 _id) view returns(uint8 _transferable, uint8 _transferFeeType, uint256 _transferFeeCurrency, uint256 _transferFeeValue, uint256 _transferFeeMaxValue)
func (_ERC1155 *ERC1155Session) TransferSettings(_id *big.Int) (struct {
	Transferable        uint8
	TransferFeeType     uint8
	TransferFeeCurrency *big.Int
	TransferFeeValue    *big.Int
	TransferFeeMaxValue *big.Int
}, error) {
	return _ERC1155.Contract.TransferSettings(&_ERC1155.CallOpts, _id)
}

// TransferSettings is a free data retrieval call binding the contract method 0x6683de90.
//
// Solidity: function transferSettings(uint256 _id) view returns(uint8 _transferable, uint8 _transferFeeType, uint256 _transferFeeCurrency, uint256 _transferFeeValue, uint256 _transferFeeMaxValue)
func (_ERC1155 *ERC1155CallerSession) TransferSettings(_id *big.Int) (struct {
	Transferable        uint8
	TransferFeeType     uint8
	TransferFeeCurrency *big.Int
	TransferFeeValue    *big.Int
	TransferFeeMaxValue *big.Int
}, error) {
	return _ERC1155.Contract.TransferSettings(&_ERC1155.CallOpts, _id)
}

// TypeByIndex is a free data retrieval call binding the contract method 0x5b7852c9.
//
// Solidity: function typeByIndex(uint256 _index) view returns(uint256 _id)
func (_ERC1155 *ERC1155Caller) TypeByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "typeByIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TypeByIndex is a free data retrieval call binding the contract method 0x5b7852c9.
//
// Solidity: function typeByIndex(uint256 _index) view returns(uint256 _id)
func (_ERC1155 *ERC1155Session) TypeByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.TypeByIndex(&_ERC1155.CallOpts, _index)
}

// TypeByIndex is a free data retrieval call binding the contract method 0x5b7852c9.
//
// Solidity: function typeByIndex(uint256 _index) view returns(uint256 _id)
func (_ERC1155 *ERC1155CallerSession) TypeByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.TypeByIndex(&_ERC1155.CallOpts, _index)
}

// TypeCount is a free data retrieval call binding the contract method 0x602d7fe6.
//
// Solidity: function typeCount() view returns(uint256)
func (_ERC1155 *ERC1155Caller) TypeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "typeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TypeCount is a free data retrieval call binding the contract method 0x602d7fe6.
//
// Solidity: function typeCount() view returns(uint256)
func (_ERC1155 *ERC1155Session) TypeCount() (*big.Int, error) {
	return _ERC1155.Contract.TypeCount(&_ERC1155.CallOpts)
}

// TypeCount is a free data retrieval call binding the contract method 0x602d7fe6.
//
// Solidity: function typeCount() view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) TypeCount() (*big.Int, error) {
	return _ERC1155.Contract.TypeCount(&_ERC1155.CallOpts)
}

// TypeData is a free data retrieval call binding the contract method 0x4341963e.
//
// Solidity: function typeData(uint256 _id) view returns(string _name, address _creator, uint256 _meltValue, uint16 _meltFeeRatio, uint16 _meltFeeMaxRatio, address _supplyModel, uint256 _totalSupply, uint256 _circulatingSupply, uint256 _reserve, uint8 _transferable, uint256[4] _transferFeeData, bool _nonFungible)
func (_ERC1155 *ERC1155Caller) TypeData(opts *bind.CallOpts, _id *big.Int) (struct {
	Name              string
	Creator           common.Address
	MeltValue         *big.Int
	MeltFeeRatio      uint16
	MeltFeeMaxRatio   uint16
	SupplyModel       common.Address
	TotalSupply       *big.Int
	CirculatingSupply *big.Int
	Reserve           *big.Int
	Transferable      uint8
	TransferFeeData   [4]*big.Int
	NonFungible       bool
}, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "typeData", _id)

	outstruct := new(struct {
		Name              string
		Creator           common.Address
		MeltValue         *big.Int
		MeltFeeRatio      uint16
		MeltFeeMaxRatio   uint16
		SupplyModel       common.Address
		TotalSupply       *big.Int
		CirculatingSupply *big.Int
		Reserve           *big.Int
		Transferable      uint8
		TransferFeeData   [4]*big.Int
		NonFungible       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.MeltValue = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MeltFeeRatio = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.MeltFeeMaxRatio = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.SupplyModel = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.TotalSupply = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CirculatingSupply = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Reserve = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Transferable = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.TransferFeeData = *abi.ConvertType(out[10], new([4]*big.Int)).(*[4]*big.Int)
	outstruct.NonFungible = *abi.ConvertType(out[11], new(bool)).(*bool)

	return *outstruct, err

}

// TypeData is a free data retrieval call binding the contract method 0x4341963e.
//
// Solidity: function typeData(uint256 _id) view returns(string _name, address _creator, uint256 _meltValue, uint16 _meltFeeRatio, uint16 _meltFeeMaxRatio, address _supplyModel, uint256 _totalSupply, uint256 _circulatingSupply, uint256 _reserve, uint8 _transferable, uint256[4] _transferFeeData, bool _nonFungible)
func (_ERC1155 *ERC1155Session) TypeData(_id *big.Int) (struct {
	Name              string
	Creator           common.Address
	MeltValue         *big.Int
	MeltFeeRatio      uint16
	MeltFeeMaxRatio   uint16
	SupplyModel       common.Address
	TotalSupply       *big.Int
	CirculatingSupply *big.Int
	Reserve           *big.Int
	Transferable      uint8
	TransferFeeData   [4]*big.Int
	NonFungible       bool
}, error) {
	return _ERC1155.Contract.TypeData(&_ERC1155.CallOpts, _id)
}

// TypeData is a free data retrieval call binding the contract method 0x4341963e.
//
// Solidity: function typeData(uint256 _id) view returns(string _name, address _creator, uint256 _meltValue, uint16 _meltFeeRatio, uint16 _meltFeeMaxRatio, address _supplyModel, uint256 _totalSupply, uint256 _circulatingSupply, uint256 _reserve, uint8 _transferable, uint256[4] _transferFeeData, bool _nonFungible)
func (_ERC1155 *ERC1155CallerSession) TypeData(_id *big.Int) (struct {
	Name              string
	Creator           common.Address
	MeltValue         *big.Int
	MeltFeeRatio      uint16
	MeltFeeMaxRatio   uint16
	SupplyModel       common.Address
	TotalSupply       *big.Int
	CirculatingSupply *big.Int
	Reserve           *big.Int
	Transferable      uint8
	TransferFeeData   [4]*big.Int
	NonFungible       bool
}, error) {
	return _ERC1155.Contract.TypeData(&_ERC1155.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_ERC1155 *ERC1155Caller) Uri(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "uri", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_ERC1155 *ERC1155Session) Uri(_id *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_ERC1155 *ERC1155CallerSession) Uri(_id *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, _id)
}

// Whitelisted is a free data retrieval call binding the contract method 0xb2aac5fb.
//
// Solidity: function whitelisted(uint256 _id, address _account, address _whitelisted) view returns(bool)
func (_ERC1155 *ERC1155Caller) Whitelisted(opts *bind.CallOpts, _id *big.Int, _account common.Address, _whitelisted common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "whitelisted", _id, _account, _whitelisted)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xb2aac5fb.
//
// Solidity: function whitelisted(uint256 _id, address _account, address _whitelisted) view returns(bool)
func (_ERC1155 *ERC1155Session) Whitelisted(_id *big.Int, _account common.Address, _whitelisted common.Address) (bool, error) {
	return _ERC1155.Contract.Whitelisted(&_ERC1155.CallOpts, _id, _account, _whitelisted)
}

// Whitelisted is a free data retrieval call binding the contract method 0xb2aac5fb.
//
// Solidity: function whitelisted(uint256 _id, address _account, address _whitelisted) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) Whitelisted(_id *big.Int, _account common.Address, _whitelisted common.Address) (bool, error) {
	return _ERC1155.Contract.Whitelisted(&_ERC1155.CallOpts, _id, _account, _whitelisted)
}

// AcceptAssignment is a paid mutator transaction binding the contract method 0x6907be85.
//
// Solidity: function acceptAssignment(uint256 _id) returns()
func (_ERC1155 *ERC1155Transactor) AcceptAssignment(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "acceptAssignment", _id)
}

// AcceptAssignment is a paid mutator transaction binding the contract method 0x6907be85.
//
// Solidity: function acceptAssignment(uint256 _id) returns()
func (_ERC1155 *ERC1155Session) AcceptAssignment(_id *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.AcceptAssignment(&_ERC1155.TransactOpts, _id)
}

// AcceptAssignment is a paid mutator transaction binding the contract method 0x6907be85.
//
// Solidity: function acceptAssignment(uint256 _id) returns()
func (_ERC1155 *ERC1155TransactorSession) AcceptAssignment(_id *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.AcceptAssignment(&_ERC1155.TransactOpts, _id)
}

// AddLog is a paid mutator transaction binding the contract method 0x36d10002.
//
// Solidity: function addLog(uint256 _id, string _data) returns()
func (_ERC1155 *ERC1155Transactor) AddLog(opts *bind.TransactOpts, _id *big.Int, _data string) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "addLog", _id, _data)
}

// AddLog is a paid mutator transaction binding the contract method 0x36d10002.
//
// Solidity: function addLog(uint256 _id, string _data) returns()
func (_ERC1155 *ERC1155Session) AddLog(_id *big.Int, _data string) (*types.Transaction, error) {
	return _ERC1155.Contract.AddLog(&_ERC1155.TransactOpts, _id, _data)
}

// AddLog is a paid mutator transaction binding the contract method 0x36d10002.
//
// Solidity: function addLog(uint256 _id, string _data) returns()
func (_ERC1155 *ERC1155TransactorSession) AddLog(_id *big.Int, _data string) (*types.Transaction, error) {
	return _ERC1155.Contract.AddLog(&_ERC1155.TransactOpts, _id, _data)
}

// Approve is a paid mutator transaction binding the contract method 0x4f4df442.
//
// Solidity: function approve(address _spender, uint256 _id, uint256 _currentValue, uint256 _value) returns()
func (_ERC1155 *ERC1155Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _id *big.Int, _currentValue *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "approve", _spender, _id, _currentValue, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x4f4df442.
//
// Solidity: function approve(address _spender, uint256 _id, uint256 _currentValue, uint256 _value) returns()
func (_ERC1155 *ERC1155Session) Approve(_spender common.Address, _id *big.Int, _currentValue *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Approve(&_ERC1155.TransactOpts, _spender, _id, _currentValue, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x4f4df442.
//
// Solidity: function approve(address _spender, uint256 _id, uint256 _currentValue, uint256 _value) returns()
func (_ERC1155 *ERC1155TransactorSession) Approve(_spender common.Address, _id *big.Int, _currentValue *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Approve(&_ERC1155.TransactOpts, _spender, _id, _currentValue, _value)
}

// ApproveAdapter is a paid mutator transaction binding the contract method 0x438df628.
//
// Solidity: function approveAdapter(address _spender, uint256 _id, uint256 _currentValue, uint256 _value, address _msgSender) returns()
func (_ERC1155 *ERC1155Transactor) ApproveAdapter(opts *bind.TransactOpts, _spender common.Address, _id *big.Int, _currentValue *big.Int, _value *big.Int, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "approveAdapter", _spender, _id, _currentValue, _value, _msgSender)
}

// ApproveAdapter is a paid mutator transaction binding the contract method 0x438df628.
//
// Solidity: function approveAdapter(address _spender, uint256 _id, uint256 _currentValue, uint256 _value, address _msgSender) returns()
func (_ERC1155 *ERC1155Session) ApproveAdapter(_spender common.Address, _id *big.Int, _currentValue *big.Int, _value *big.Int, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.ApproveAdapter(&_ERC1155.TransactOpts, _spender, _id, _currentValue, _value, _msgSender)
}

// ApproveAdapter is a paid mutator transaction binding the contract method 0x438df628.
//
// Solidity: function approveAdapter(address _spender, uint256 _id, uint256 _currentValue, uint256 _value, address _msgSender) returns()
func (_ERC1155 *ERC1155TransactorSession) ApproveAdapter(_spender common.Address, _id *big.Int, _currentValue *big.Int, _value *big.Int, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.ApproveAdapter(&_ERC1155.TransactOpts, _spender, _id, _currentValue, _value, _msgSender)
}

// Assign is a paid mutator transaction binding the contract method 0xe07d3b5a.
//
// Solidity: function assign(uint256 _id, address _creator) returns()
func (_ERC1155 *ERC1155Transactor) Assign(opts *bind.TransactOpts, _id *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "assign", _id, _creator)
}

// Assign is a paid mutator transaction binding the contract method 0xe07d3b5a.
//
// Solidity: function assign(uint256 _id, address _creator) returns()
func (_ERC1155 *ERC1155Session) Assign(_id *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.Assign(&_ERC1155.TransactOpts, _id, _creator)
}

// Assign is a paid mutator transaction binding the contract method 0xe07d3b5a.
//
// Solidity: function assign(uint256 _id, address _creator) returns()
func (_ERC1155 *ERC1155TransactorSession) Assign(_id *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.Assign(&_ERC1155.TransactOpts, _id, _creator)
}

// BatchApprove is a paid mutator transaction binding the contract method 0xd1bab4cc.
//
// Solidity: function batchApprove(address _spender, uint256[] _ids, uint256[] _currentValues, uint256[] _values) returns()
func (_ERC1155 *ERC1155Transactor) BatchApprove(opts *bind.TransactOpts, _spender common.Address, _ids []*big.Int, _currentValues []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "batchApprove", _spender, _ids, _currentValues, _values)
}

// BatchApprove is a paid mutator transaction binding the contract method 0xd1bab4cc.
//
// Solidity: function batchApprove(address _spender, uint256[] _ids, uint256[] _currentValues, uint256[] _values) returns()
func (_ERC1155 *ERC1155Session) BatchApprove(_spender common.Address, _ids []*big.Int, _currentValues []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.BatchApprove(&_ERC1155.TransactOpts, _spender, _ids, _currentValues, _values)
}

// BatchApprove is a paid mutator transaction binding the contract method 0xd1bab4cc.
//
// Solidity: function batchApprove(address _spender, uint256[] _ids, uint256[] _currentValues, uint256[] _values) returns()
func (_ERC1155 *ERC1155TransactorSession) BatchApprove(_spender common.Address, _ids []*big.Int, _currentValues []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.BatchApprove(&_ERC1155.TransactOpts, _spender, _ids, _currentValues, _values)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0xe39c6d80.
//
// Solidity: function batchTransfer(address _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Transactor) BatchTransfer(opts *bind.TransactOpts, _to common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "batchTransfer", _to, _ids, _values)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0xe39c6d80.
//
// Solidity: function batchTransfer(address _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Session) BatchTransfer(_to common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.BatchTransfer(&_ERC1155.TransactOpts, _to, _ids, _values)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0xe39c6d80.
//
// Solidity: function batchTransfer(address _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155TransactorSession) BatchTransfer(_to common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.BatchTransfer(&_ERC1155.TransactOpts, _to, _ids, _values)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x17fad7fc.
//
// Solidity: function batchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Transactor) BatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "batchTransferFrom", _from, _to, _ids, _values)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x17fad7fc.
//
// Solidity: function batchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Session) BatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.BatchTransferFrom(&_ERC1155.TransactOpts, _from, _to, _ids, _values)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x17fad7fc.
//
// Solidity: function batchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155TransactorSession) BatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.BatchTransferFrom(&_ERC1155.TransactOpts, _from, _to, _ids, _values)
}

// CancelTrade is a paid mutator transaction binding the contract method 0x09ec6cc7.
//
// Solidity: function cancelTrade(uint256 _id) returns()
func (_ERC1155 *ERC1155Transactor) CancelTrade(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "cancelTrade", _id)
}

// CancelTrade is a paid mutator transaction binding the contract method 0x09ec6cc7.
//
// Solidity: function cancelTrade(uint256 _id) returns()
func (_ERC1155 *ERC1155Session) CancelTrade(_id *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.CancelTrade(&_ERC1155.TransactOpts, _id)
}

// CancelTrade is a paid mutator transaction binding the contract method 0x09ec6cc7.
//
// Solidity: function cancelTrade(uint256 _id) returns()
func (_ERC1155 *ERC1155TransactorSession) CancelTrade(_id *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.CancelTrade(&_ERC1155.TransactOpts, _id)
}

// CompleteTrade is a paid mutator transaction binding the contract method 0xad221551.
//
// Solidity: function completeTrade(uint256 _id) returns()
func (_ERC1155 *ERC1155Transactor) CompleteTrade(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "completeTrade", _id)
}

// CompleteTrade is a paid mutator transaction binding the contract method 0xad221551.
//
// Solidity: function completeTrade(uint256 _id) returns()
func (_ERC1155 *ERC1155Session) CompleteTrade(_id *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.CompleteTrade(&_ERC1155.TransactOpts, _id)
}

// CompleteTrade is a paid mutator transaction binding the contract method 0xad221551.
//
// Solidity: function completeTrade(uint256 _id) returns()
func (_ERC1155 *ERC1155TransactorSession) CompleteTrade(_id *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.CompleteTrade(&_ERC1155.TransactOpts, _id)
}

// Create is a paid mutator transaction binding the contract method 0xcd23dde0.
//
// Solidity: function create(string _name, uint256 _totalSupply, uint256 _initialReserve, address _supplyModel, uint256 _meltValue, uint16 _meltFeeRatio, uint8 _transferable, uint256[3] _transferFeeSettings, bool _nonFungible) returns()
func (_ERC1155 *ERC1155Transactor) Create(opts *bind.TransactOpts, _name string, _totalSupply *big.Int, _initialReserve *big.Int, _supplyModel common.Address, _meltValue *big.Int, _meltFeeRatio uint16, _transferable uint8, _transferFeeSettings [3]*big.Int, _nonFungible bool) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "create", _name, _totalSupply, _initialReserve, _supplyModel, _meltValue, _meltFeeRatio, _transferable, _transferFeeSettings, _nonFungible)
}

// Create is a paid mutator transaction binding the contract method 0xcd23dde0.
//
// Solidity: function create(string _name, uint256 _totalSupply, uint256 _initialReserve, address _supplyModel, uint256 _meltValue, uint16 _meltFeeRatio, uint8 _transferable, uint256[3] _transferFeeSettings, bool _nonFungible) returns()
func (_ERC1155 *ERC1155Session) Create(_name string, _totalSupply *big.Int, _initialReserve *big.Int, _supplyModel common.Address, _meltValue *big.Int, _meltFeeRatio uint16, _transferable uint8, _transferFeeSettings [3]*big.Int, _nonFungible bool) (*types.Transaction, error) {
	return _ERC1155.Contract.Create(&_ERC1155.TransactOpts, _name, _totalSupply, _initialReserve, _supplyModel, _meltValue, _meltFeeRatio, _transferable, _transferFeeSettings, _nonFungible)
}

// Create is a paid mutator transaction binding the contract method 0xcd23dde0.
//
// Solidity: function create(string _name, uint256 _totalSupply, uint256 _initialReserve, address _supplyModel, uint256 _meltValue, uint16 _meltFeeRatio, uint8 _transferable, uint256[3] _transferFeeSettings, bool _nonFungible) returns()
func (_ERC1155 *ERC1155TransactorSession) Create(_name string, _totalSupply *big.Int, _initialReserve *big.Int, _supplyModel common.Address, _meltValue *big.Int, _meltFeeRatio uint16, _transferable uint8, _transferFeeSettings [3]*big.Int, _nonFungible bool) (*types.Transaction, error) {
	return _ERC1155.Contract.Create(&_ERC1155.TransactOpts, _name, _totalSupply, _initialReserve, _supplyModel, _meltValue, _meltFeeRatio, _transferable, _transferFeeSettings, _nonFungible)
}

// CreateTrade is a paid mutator transaction binding the contract method 0xf637ee09.
//
// Solidity: function createTrade(uint256[] _askingIds, uint128[] _askingValues, uint256[] _offeringIds, uint128[] _offeringValues, address _secondParty) returns()
func (_ERC1155 *ERC1155Transactor) CreateTrade(opts *bind.TransactOpts, _askingIds []*big.Int, _askingValues []*big.Int, _offeringIds []*big.Int, _offeringValues []*big.Int, _secondParty common.Address) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "createTrade", _askingIds, _askingValues, _offeringIds, _offeringValues, _secondParty)
}

// CreateTrade is a paid mutator transaction binding the contract method 0xf637ee09.
//
// Solidity: function createTrade(uint256[] _askingIds, uint128[] _askingValues, uint256[] _offeringIds, uint128[] _offeringValues, address _secondParty) returns()
func (_ERC1155 *ERC1155Session) CreateTrade(_askingIds []*big.Int, _askingValues []*big.Int, _offeringIds []*big.Int, _offeringValues []*big.Int, _secondParty common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.CreateTrade(&_ERC1155.TransactOpts, _askingIds, _askingValues, _offeringIds, _offeringValues, _secondParty)
}

// CreateTrade is a paid mutator transaction binding the contract method 0xf637ee09.
//
// Solidity: function createTrade(uint256[] _askingIds, uint128[] _askingValues, uint256[] _offeringIds, uint128[] _offeringValues, address _secondParty) returns()
func (_ERC1155 *ERC1155TransactorSession) CreateTrade(_askingIds []*big.Int, _askingValues []*big.Int, _offeringIds []*big.Int, _offeringValues []*big.Int, _secondParty common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.CreateTrade(&_ERC1155.TransactOpts, _askingIds, _askingValues, _offeringIds, _offeringValues, _secondParty)
}

// DecreaseMaxMeltFee is a paid mutator transaction binding the contract method 0x8fa1cc21.
//
// Solidity: function decreaseMaxMeltFee(uint256 _id, uint16 _fee) returns()
func (_ERC1155 *ERC1155Transactor) DecreaseMaxMeltFee(opts *bind.TransactOpts, _id *big.Int, _fee uint16) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "decreaseMaxMeltFee", _id, _fee)
}

// DecreaseMaxMeltFee is a paid mutator transaction binding the contract method 0x8fa1cc21.
//
// Solidity: function decreaseMaxMeltFee(uint256 _id, uint16 _fee) returns()
func (_ERC1155 *ERC1155Session) DecreaseMaxMeltFee(_id *big.Int, _fee uint16) (*types.Transaction, error) {
	return _ERC1155.Contract.DecreaseMaxMeltFee(&_ERC1155.TransactOpts, _id, _fee)
}

// DecreaseMaxMeltFee is a paid mutator transaction binding the contract method 0x8fa1cc21.
//
// Solidity: function decreaseMaxMeltFee(uint256 _id, uint16 _fee) returns()
func (_ERC1155 *ERC1155TransactorSession) DecreaseMaxMeltFee(_id *big.Int, _fee uint16) (*types.Transaction, error) {
	return _ERC1155.Contract.DecreaseMaxMeltFee(&_ERC1155.TransactOpts, _id, _fee)
}

// DecreaseMaxTransferFee is a paid mutator transaction binding the contract method 0xc5817e1e.
//
// Solidity: function decreaseMaxTransferFee(uint256 _id, uint256 _fee) returns()
func (_ERC1155 *ERC1155Transactor) DecreaseMaxTransferFee(opts *bind.TransactOpts, _id *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "decreaseMaxTransferFee", _id, _fee)
}

// DecreaseMaxTransferFee is a paid mutator transaction binding the contract method 0xc5817e1e.
//
// Solidity: function decreaseMaxTransferFee(uint256 _id, uint256 _fee) returns()
func (_ERC1155 *ERC1155Session) DecreaseMaxTransferFee(_id *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.DecreaseMaxTransferFee(&_ERC1155.TransactOpts, _id, _fee)
}

// DecreaseMaxTransferFee is a paid mutator transaction binding the contract method 0xc5817e1e.
//
// Solidity: function decreaseMaxTransferFee(uint256 _id, uint256 _fee) returns()
func (_ERC1155 *ERC1155TransactorSession) DecreaseMaxTransferFee(_id *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.DecreaseMaxTransferFee(&_ERC1155.TransactOpts, _id, _fee)
}

// DeployERC20Adapter is a paid mutator transaction binding the contract method 0xaaee660e.
//
// Solidity: function deployERC20Adapter(uint256 _id, uint8 _decimals, string _symbol) returns(address)
func (_ERC1155 *ERC1155Transactor) DeployERC20Adapter(opts *bind.TransactOpts, _id *big.Int, _decimals uint8, _symbol string) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "deployERC20Adapter", _id, _decimals, _symbol)
}

// DeployERC20Adapter is a paid mutator transaction binding the contract method 0xaaee660e.
//
// Solidity: function deployERC20Adapter(uint256 _id, uint8 _decimals, string _symbol) returns(address)
func (_ERC1155 *ERC1155Session) DeployERC20Adapter(_id *big.Int, _decimals uint8, _symbol string) (*types.Transaction, error) {
	return _ERC1155.Contract.DeployERC20Adapter(&_ERC1155.TransactOpts, _id, _decimals, _symbol)
}

// DeployERC20Adapter is a paid mutator transaction binding the contract method 0xaaee660e.
//
// Solidity: function deployERC20Adapter(uint256 _id, uint8 _decimals, string _symbol) returns(address)
func (_ERC1155 *ERC1155TransactorSession) DeployERC20Adapter(_id *big.Int, _decimals uint8, _symbol string) (*types.Transaction, error) {
	return _ERC1155.Contract.DeployERC20Adapter(&_ERC1155.TransactOpts, _id, _decimals, _symbol)
}

// DeployERC721Adapter is a paid mutator transaction binding the contract method 0x6ab98c15.
//
// Solidity: function deployERC721Adapter(uint256 _id, string _symbol) returns(address)
func (_ERC1155 *ERC1155Transactor) DeployERC721Adapter(opts *bind.TransactOpts, _id *big.Int, _symbol string) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "deployERC721Adapter", _id, _symbol)
}

// DeployERC721Adapter is a paid mutator transaction binding the contract method 0x6ab98c15.
//
// Solidity: function deployERC721Adapter(uint256 _id, string _symbol) returns(address)
func (_ERC1155 *ERC1155Session) DeployERC721Adapter(_id *big.Int, _symbol string) (*types.Transaction, error) {
	return _ERC1155.Contract.DeployERC721Adapter(&_ERC1155.TransactOpts, _id, _symbol)
}

// DeployERC721Adapter is a paid mutator transaction binding the contract method 0x6ab98c15.
//
// Solidity: function deployERC721Adapter(uint256 _id, string _symbol) returns(address)
func (_ERC1155 *ERC1155TransactorSession) DeployERC721Adapter(_id *big.Int, _symbol string) (*types.Transaction, error) {
	return _ERC1155.Contract.DeployERC721Adapter(&_ERC1155.TransactOpts, _id, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _storage, address _oldContract) returns()
func (_ERC1155 *ERC1155Transactor) Initialize(opts *bind.TransactOpts, _storage common.Address, _oldContract common.Address) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "initialize", _storage, _oldContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _storage, address _oldContract) returns()
func (_ERC1155 *ERC1155Session) Initialize(_storage common.Address, _oldContract common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.Initialize(&_ERC1155.TransactOpts, _storage, _oldContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _storage, address _oldContract) returns()
func (_ERC1155 *ERC1155TransactorSession) Initialize(_storage common.Address, _oldContract common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.Initialize(&_ERC1155.TransactOpts, _storage, _oldContract)
}

// Melt is a paid mutator transaction binding the contract method 0xf6089e12.
//
// Solidity: function melt(uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Transactor) Melt(opts *bind.TransactOpts, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "melt", _ids, _values)
}

// Melt is a paid mutator transaction binding the contract method 0xf6089e12.
//
// Solidity: function melt(uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Session) Melt(_ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Melt(&_ERC1155.TransactOpts, _ids, _values)
}

// Melt is a paid mutator transaction binding the contract method 0xf6089e12.
//
// Solidity: function melt(uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155TransactorSession) Melt(_ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Melt(&_ERC1155.TransactOpts, _ids, _values)
}

// MintFungibles is a paid mutator transaction binding the contract method 0x3d7d20a4.
//
// Solidity: function mintFungibles(uint256 _id, address[] _to, uint256[] _values) returns()
func (_ERC1155 *ERC1155Transactor) MintFungibles(opts *bind.TransactOpts, _id *big.Int, _to []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "mintFungibles", _id, _to, _values)
}

// MintFungibles is a paid mutator transaction binding the contract method 0x3d7d20a4.
//
// Solidity: function mintFungibles(uint256 _id, address[] _to, uint256[] _values) returns()
func (_ERC1155 *ERC1155Session) MintFungibles(_id *big.Int, _to []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.MintFungibles(&_ERC1155.TransactOpts, _id, _to, _values)
}

// MintFungibles is a paid mutator transaction binding the contract method 0x3d7d20a4.
//
// Solidity: function mintFungibles(uint256 _id, address[] _to, uint256[] _values) returns()
func (_ERC1155 *ERC1155TransactorSession) MintFungibles(_id *big.Int, _to []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.MintFungibles(&_ERC1155.TransactOpts, _id, _to, _values)
}

// MintNonFungibles is a paid mutator transaction binding the contract method 0x00549c2b.
//
// Solidity: function mintNonFungibles(uint256 _id, address[] _to) returns()
func (_ERC1155 *ERC1155Transactor) MintNonFungibles(opts *bind.TransactOpts, _id *big.Int, _to []common.Address) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "mintNonFungibles", _id, _to)
}

// MintNonFungibles is a paid mutator transaction binding the contract method 0x00549c2b.
//
// Solidity: function mintNonFungibles(uint256 _id, address[] _to) returns()
func (_ERC1155 *ERC1155Session) MintNonFungibles(_id *big.Int, _to []common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.MintNonFungibles(&_ERC1155.TransactOpts, _id, _to)
}

// MintNonFungibles is a paid mutator transaction binding the contract method 0x00549c2b.
//
// Solidity: function mintNonFungibles(uint256 _id, address[] _to) returns()
func (_ERC1155 *ERC1155TransactorSession) MintNonFungibles(_id *big.Int, _to []common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.MintNonFungibles(&_ERC1155.TransactOpts, _id, _to)
}

// MintNonFungiblesWithData is a paid mutator transaction binding the contract method 0x176ea852.
//
// Solidity: function mintNonFungiblesWithData(uint256 _id, address[] _to, uint128[] _data) returns()
func (_ERC1155 *ERC1155Transactor) MintNonFungiblesWithData(opts *bind.TransactOpts, _id *big.Int, _to []common.Address, _data []*big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "mintNonFungiblesWithData", _id, _to, _data)
}

// MintNonFungiblesWithData is a paid mutator transaction binding the contract method 0x176ea852.
//
// Solidity: function mintNonFungiblesWithData(uint256 _id, address[] _to, uint128[] _data) returns()
func (_ERC1155 *ERC1155Session) MintNonFungiblesWithData(_id *big.Int, _to []common.Address, _data []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.MintNonFungiblesWithData(&_ERC1155.TransactOpts, _id, _to, _data)
}

// MintNonFungiblesWithData is a paid mutator transaction binding the contract method 0x176ea852.
//
// Solidity: function mintNonFungiblesWithData(uint256 _id, address[] _to, uint128[] _data) returns()
func (_ERC1155 *ERC1155TransactorSession) MintNonFungiblesWithData(_id *big.Int, _to []common.Address, _data []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.MintNonFungiblesWithData(&_ERC1155.TransactOpts, _id, _to, _data)
}

// MulticastTransfer is a paid mutator transaction binding the contract method 0x67fd9da3.
//
// Solidity: function multicastTransfer(address[] _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Transactor) MulticastTransfer(opts *bind.TransactOpts, _to []common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "multicastTransfer", _to, _ids, _values)
}

// MulticastTransfer is a paid mutator transaction binding the contract method 0x67fd9da3.
//
// Solidity: function multicastTransfer(address[] _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Session) MulticastTransfer(_to []common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.MulticastTransfer(&_ERC1155.TransactOpts, _to, _ids, _values)
}

// MulticastTransfer is a paid mutator transaction binding the contract method 0x67fd9da3.
//
// Solidity: function multicastTransfer(address[] _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155TransactorSession) MulticastTransfer(_to []common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.MulticastTransfer(&_ERC1155.TransactOpts, _to, _ids, _values)
}

// MulticastTransferFrom is a paid mutator transaction binding the contract method 0x8f76ccf7.
//
// Solidity: function multicastTransferFrom(address[] _from, address[] _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Transactor) MulticastTransferFrom(opts *bind.TransactOpts, _from []common.Address, _to []common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "multicastTransferFrom", _from, _to, _ids, _values)
}

// MulticastTransferFrom is a paid mutator transaction binding the contract method 0x8f76ccf7.
//
// Solidity: function multicastTransferFrom(address[] _from, address[] _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155Session) MulticastTransferFrom(_from []common.Address, _to []common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.MulticastTransferFrom(&_ERC1155.TransactOpts, _from, _to, _ids, _values)
}

// MulticastTransferFrom is a paid mutator transaction binding the contract method 0x8f76ccf7.
//
// Solidity: function multicastTransferFrom(address[] _from, address[] _to, uint256[] _ids, uint256[] _values) returns()
func (_ERC1155 *ERC1155TransactorSession) MulticastTransferFrom(_from []common.Address, _to []common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.MulticastTransferFrom(&_ERC1155.TransactOpts, _from, _to, _ids, _values)
}

// ReleaseERC1155 is a paid mutator transaction binding the contract method 0xa9d27375.
//
// Solidity: function releaseERC1155(address _erc1155ContractAddress, address _to, uint256 _id, uint256 _value) returns()
func (_ERC1155 *ERC1155Transactor) ReleaseERC1155(opts *bind.TransactOpts, _erc1155ContractAddress common.Address, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "releaseERC1155", _erc1155ContractAddress, _to, _id, _value)
}

// ReleaseERC1155 is a paid mutator transaction binding the contract method 0xa9d27375.
//
// Solidity: function releaseERC1155(address _erc1155ContractAddress, address _to, uint256 _id, uint256 _value) returns()
func (_ERC1155 *ERC1155Session) ReleaseERC1155(_erc1155ContractAddress common.Address, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseERC1155(&_ERC1155.TransactOpts, _erc1155ContractAddress, _to, _id, _value)
}

// ReleaseERC1155 is a paid mutator transaction binding the contract method 0xa9d27375.
//
// Solidity: function releaseERC1155(address _erc1155ContractAddress, address _to, uint256 _id, uint256 _value) returns()
func (_ERC1155 *ERC1155TransactorSession) ReleaseERC1155(_erc1155ContractAddress common.Address, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseERC1155(&_ERC1155.TransactOpts, _erc1155ContractAddress, _to, _id, _value)
}

// ReleaseERC20 is a paid mutator transaction binding the contract method 0x6cd533d8.
//
// Solidity: function releaseERC20(address _erc20ContractAddress, address _to, uint256 _value) returns()
func (_ERC1155 *ERC1155Transactor) ReleaseERC20(opts *bind.TransactOpts, _erc20ContractAddress common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "releaseERC20", _erc20ContractAddress, _to, _value)
}

// ReleaseERC20 is a paid mutator transaction binding the contract method 0x6cd533d8.
//
// Solidity: function releaseERC20(address _erc20ContractAddress, address _to, uint256 _value) returns()
func (_ERC1155 *ERC1155Session) ReleaseERC20(_erc20ContractAddress common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseERC20(&_ERC1155.TransactOpts, _erc20ContractAddress, _to, _value)
}

// ReleaseERC20 is a paid mutator transaction binding the contract method 0x6cd533d8.
//
// Solidity: function releaseERC20(address _erc20ContractAddress, address _to, uint256 _value) returns()
func (_ERC1155 *ERC1155TransactorSession) ReleaseERC20(_erc20ContractAddress common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseERC20(&_ERC1155.TransactOpts, _erc20ContractAddress, _to, _value)
}

// ReleaseERC721 is a paid mutator transaction binding the contract method 0xd8265846.
//
// Solidity: function releaseERC721(address _erc721ContractAddress, address _to, uint256 _token) returns()
func (_ERC1155 *ERC1155Transactor) ReleaseERC721(opts *bind.TransactOpts, _erc721ContractAddress common.Address, _to common.Address, _token *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "releaseERC721", _erc721ContractAddress, _to, _token)
}

// ReleaseERC721 is a paid mutator transaction binding the contract method 0xd8265846.
//
// Solidity: function releaseERC721(address _erc721ContractAddress, address _to, uint256 _token) returns()
func (_ERC1155 *ERC1155Session) ReleaseERC721(_erc721ContractAddress common.Address, _to common.Address, _token *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseERC721(&_ERC1155.TransactOpts, _erc721ContractAddress, _to, _token)
}

// ReleaseERC721 is a paid mutator transaction binding the contract method 0xd8265846.
//
// Solidity: function releaseERC721(address _erc721ContractAddress, address _to, uint256 _token) returns()
func (_ERC1155 *ERC1155TransactorSession) ReleaseERC721(_erc721ContractAddress common.Address, _to common.Address, _token *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseERC721(&_ERC1155.TransactOpts, _erc721ContractAddress, _to, _token)
}

// ReleaseETH is a paid mutator transaction binding the contract method 0xe8e71f0c.
//
// Solidity: function releaseETH(address _to, uint256 _value) returns()
func (_ERC1155 *ERC1155Transactor) ReleaseETH(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "releaseETH", _to, _value)
}

// ReleaseETH is a paid mutator transaction binding the contract method 0xe8e71f0c.
//
// Solidity: function releaseETH(address _to, uint256 _value) returns()
func (_ERC1155 *ERC1155Session) ReleaseETH(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseETH(&_ERC1155.TransactOpts, _to, _value)
}

// ReleaseETH is a paid mutator transaction binding the contract method 0xe8e71f0c.
//
// Solidity: function releaseETH(address _to, uint256 _value) returns()
func (_ERC1155 *ERC1155TransactorSession) ReleaseETH(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseETH(&_ERC1155.TransactOpts, _to, _value)
}

// ReleaseReserve is a paid mutator transaction binding the contract method 0xc2eed0de.
//
// Solidity: function releaseReserve(uint256 _id, uint128 _value) returns()
func (_ERC1155 *ERC1155Transactor) ReleaseReserve(opts *bind.TransactOpts, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "releaseReserve", _id, _value)
}

// ReleaseReserve is a paid mutator transaction binding the contract method 0xc2eed0de.
//
// Solidity: function releaseReserve(uint256 _id, uint128 _value) returns()
func (_ERC1155 *ERC1155Session) ReleaseReserve(_id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseReserve(&_ERC1155.TransactOpts, _id, _value)
}

// ReleaseReserve is a paid mutator transaction binding the contract method 0xc2eed0de.
//
// Solidity: function releaseReserve(uint256 _id, uint128 _value) returns()
func (_ERC1155 *ERC1155TransactorSession) ReleaseReserve(_id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.ReleaseReserve(&_ERC1155.TransactOpts, _id, _value)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _nextContract) returns()
func (_ERC1155 *ERC1155Transactor) Retire(opts *bind.TransactOpts, _nextContract common.Address) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "retire", _nextContract)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _nextContract) returns()
func (_ERC1155 *ERC1155Session) Retire(_nextContract common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.Retire(&_ERC1155.TransactOpts, _nextContract)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _nextContract) returns()
func (_ERC1155 *ERC1155TransactorSession) Retire(_nextContract common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.Retire(&_ERC1155.TransactOpts, _nextContract)
}

// SafeBatchTransfer is a paid mutator transaction binding the contract method 0x368b2842.
//
// Solidity: function safeBatchTransfer(address _to, uint256[] _ids, uint256[] _values, bytes _data) returns()
func (_ERC1155 *ERC1155Transactor) SafeBatchTransfer(opts *bind.TransactOpts, _to common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeBatchTransfer", _to, _ids, _values, _data)
}

// SafeBatchTransfer is a paid mutator transaction binding the contract method 0x368b2842.
//
// Solidity: function safeBatchTransfer(address _to, uint256[] _ids, uint256[] _values, bytes _data) returns()
func (_ERC1155 *ERC1155Session) SafeBatchTransfer(_to common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransfer(&_ERC1155.TransactOpts, _to, _ids, _values, _data)
}

// SafeBatchTransfer is a paid mutator transaction binding the contract method 0x368b2842.
//
// Solidity: function safeBatchTransfer(address _to, uint256[] _ids, uint256[] _values, bytes _data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeBatchTransfer(_to common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransfer(&_ERC1155.TransactOpts, _to, _ids, _values, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _values, bytes _data) returns()
func (_ERC1155 *ERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeBatchTransferFrom", _from, _to, _ids, _values, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _values, bytes _data) returns()
func (_ERC1155 *ERC1155Session) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, _from, _to, _ids, _values, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _values, bytes _data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, _from, _to, _ids, _values, _data)
}

// SafeMulticastTransfer is a paid mutator transaction binding the contract method 0xdbd3c372.
//
// Solidity: function safeMulticastTransfer(address[] _to, uint256[] _ids, uint256[] _values, bytes _data) returns()
func (_ERC1155 *ERC1155Transactor) SafeMulticastTransfer(opts *bind.TransactOpts, _to []common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeMulticastTransfer", _to, _ids, _values, _data)
}

// SafeMulticastTransfer is a paid mutator transaction binding the contract method 0xdbd3c372.
//
// Solidity: function safeMulticastTransfer(address[] _to, uint256[] _ids, uint256[] _values, bytes _data) returns()
func (_ERC1155 *ERC1155Session) SafeMulticastTransfer(_to []common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeMulticastTransfer(&_ERC1155.TransactOpts, _to, _ids, _values, _data)
}

// SafeMulticastTransfer is a paid mutator transaction binding the contract method 0xdbd3c372.
//
// Solidity: function safeMulticastTransfer(address[] _to, uint256[] _ids, uint256[] _values, bytes _data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeMulticastTransfer(_to []common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeMulticastTransfer(&_ERC1155.TransactOpts, _to, _ids, _values, _data)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xae28b68c.
//
// Solidity: function safeTransfer(address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_ERC1155 *ERC1155Transactor) SafeTransfer(opts *bind.TransactOpts, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeTransfer", _to, _id, _value, _data)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xae28b68c.
//
// Solidity: function safeTransfer(address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_ERC1155 *ERC1155Session) SafeTransfer(_to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransfer(&_ERC1155.TransactOpts, _to, _id, _value, _data)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xae28b68c.
//
// Solidity: function safeTransfer(address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeTransfer(_to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransfer(&_ERC1155.TransactOpts, _to, _id, _value, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_ERC1155 *ERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeTransferFrom", _from, _to, _id, _value, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_ERC1155 *ERC1155Session) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, _from, _to, _id, _value, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, _from, _to, _id, _value, _data)
}

// SetApproval is a paid mutator transaction binding the contract method 0xc35774a1.
//
// Solidity: function setApproval(address _operator, uint256[] _ids, bool _approved) returns()
func (_ERC1155 *ERC1155Transactor) SetApproval(opts *bind.TransactOpts, _operator common.Address, _ids []*big.Int, _approved bool) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setApproval", _operator, _ids, _approved)
}

// SetApproval is a paid mutator transaction binding the contract method 0xc35774a1.
//
// Solidity: function setApproval(address _operator, uint256[] _ids, bool _approved) returns()
func (_ERC1155 *ERC1155Session) SetApproval(_operator common.Address, _ids []*big.Int, _approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApproval(&_ERC1155.TransactOpts, _operator, _ids, _approved)
}

// SetApproval is a paid mutator transaction binding the contract method 0xc35774a1.
//
// Solidity: function setApproval(address _operator, uint256[] _ids, bool _approved) returns()
func (_ERC1155 *ERC1155TransactorSession) SetApproval(_operator common.Address, _ids []*big.Int, _approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApproval(&_ERC1155.TransactOpts, _operator, _ids, _approved)
}

// SetApprovalAdapter is a paid mutator transaction binding the contract method 0x2f0aeda4.
//
// Solidity: function setApprovalAdapter(address _operator, uint256 _id, bool _approved, address _msgSender) returns()
func (_ERC1155 *ERC1155Transactor) SetApprovalAdapter(opts *bind.TransactOpts, _operator common.Address, _id *big.Int, _approved bool, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setApprovalAdapter", _operator, _id, _approved, _msgSender)
}

// SetApprovalAdapter is a paid mutator transaction binding the contract method 0x2f0aeda4.
//
// Solidity: function setApprovalAdapter(address _operator, uint256 _id, bool _approved, address _msgSender) returns()
func (_ERC1155 *ERC1155Session) SetApprovalAdapter(_operator common.Address, _id *big.Int, _approved bool, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalAdapter(&_ERC1155.TransactOpts, _operator, _id, _approved, _msgSender)
}

// SetApprovalAdapter is a paid mutator transaction binding the contract method 0x2f0aeda4.
//
// Solidity: function setApprovalAdapter(address _operator, uint256 _id, bool _approved, address _msgSender) returns()
func (_ERC1155 *ERC1155TransactorSession) SetApprovalAdapter(_operator common.Address, _id *big.Int, _approved bool, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalAdapter(&_ERC1155.TransactOpts, _operator, _id, _approved, _msgSender)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC1155 *ERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC1155 *ERC1155Session) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC1155 *ERC1155TransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, _operator, _approved)
}

// SetMeltFee is a paid mutator transaction binding the contract method 0xa6566f8d.
//
// Solidity: function setMeltFee(uint256 _id, uint16 _fee) returns()
func (_ERC1155 *ERC1155Transactor) SetMeltFee(opts *bind.TransactOpts, _id *big.Int, _fee uint16) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setMeltFee", _id, _fee)
}

// SetMeltFee is a paid mutator transaction binding the contract method 0xa6566f8d.
//
// Solidity: function setMeltFee(uint256 _id, uint16 _fee) returns()
func (_ERC1155 *ERC1155Session) SetMeltFee(_id *big.Int, _fee uint16) (*types.Transaction, error) {
	return _ERC1155.Contract.SetMeltFee(&_ERC1155.TransactOpts, _id, _fee)
}

// SetMeltFee is a paid mutator transaction binding the contract method 0xa6566f8d.
//
// Solidity: function setMeltFee(uint256 _id, uint16 _fee) returns()
func (_ERC1155 *ERC1155TransactorSession) SetMeltFee(_id *big.Int, _fee uint16) (*types.Transaction, error) {
	return _ERC1155.Contract.SetMeltFee(&_ERC1155.TransactOpts, _id, _fee)
}

// SetTransferFee is a paid mutator transaction binding the contract method 0x934930a1.
//
// Solidity: function setTransferFee(uint256 _id, uint256 _fee) returns()
func (_ERC1155 *ERC1155Transactor) SetTransferFee(opts *bind.TransactOpts, _id *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setTransferFee", _id, _fee)
}

// SetTransferFee is a paid mutator transaction binding the contract method 0x934930a1.
//
// Solidity: function setTransferFee(uint256 _id, uint256 _fee) returns()
func (_ERC1155 *ERC1155Session) SetTransferFee(_id *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.SetTransferFee(&_ERC1155.TransactOpts, _id, _fee)
}

// SetTransferFee is a paid mutator transaction binding the contract method 0x934930a1.
//
// Solidity: function setTransferFee(uint256 _id, uint256 _fee) returns()
func (_ERC1155 *ERC1155TransactorSession) SetTransferFee(_id *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.SetTransferFee(&_ERC1155.TransactOpts, _id, _fee)
}

// SetTransferable is a paid mutator transaction binding the contract method 0x34e07ff3.
//
// Solidity: function setTransferable(uint256 _id, uint8 _transferable) returns()
func (_ERC1155 *ERC1155Transactor) SetTransferable(opts *bind.TransactOpts, _id *big.Int, _transferable uint8) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setTransferable", _id, _transferable)
}

// SetTransferable is a paid mutator transaction binding the contract method 0x34e07ff3.
//
// Solidity: function setTransferable(uint256 _id, uint8 _transferable) returns()
func (_ERC1155 *ERC1155Session) SetTransferable(_id *big.Int, _transferable uint8) (*types.Transaction, error) {
	return _ERC1155.Contract.SetTransferable(&_ERC1155.TransactOpts, _id, _transferable)
}

// SetTransferable is a paid mutator transaction binding the contract method 0x34e07ff3.
//
// Solidity: function setTransferable(uint256 _id, uint8 _transferable) returns()
func (_ERC1155 *ERC1155TransactorSession) SetTransferable(_id *big.Int, _transferable uint8) (*types.Transaction, error) {
	return _ERC1155.Contract.SetTransferable(&_ERC1155.TransactOpts, _id, _transferable)
}

// SetURI is a paid mutator transaction binding the contract method 0x862440e2.
//
// Solidity: function setURI(uint256 _id, string _uri) returns()
func (_ERC1155 *ERC1155Transactor) SetURI(opts *bind.TransactOpts, _id *big.Int, _uri string) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setURI", _id, _uri)
}

// SetURI is a paid mutator transaction binding the contract method 0x862440e2.
//
// Solidity: function setURI(uint256 _id, string _uri) returns()
func (_ERC1155 *ERC1155Session) SetURI(_id *big.Int, _uri string) (*types.Transaction, error) {
	return _ERC1155.Contract.SetURI(&_ERC1155.TransactOpts, _id, _uri)
}

// SetURI is a paid mutator transaction binding the contract method 0x862440e2.
//
// Solidity: function setURI(uint256 _id, string _uri) returns()
func (_ERC1155 *ERC1155TransactorSession) SetURI(_id *big.Int, _uri string) (*types.Transaction, error) {
	return _ERC1155.Contract.SetURI(&_ERC1155.TransactOpts, _id, _uri)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0xe6e21c75.
//
// Solidity: function setWhitelisted(uint256 _id, address _account, address _whitelisted, bool _on) returns()
func (_ERC1155 *ERC1155Transactor) SetWhitelisted(opts *bind.TransactOpts, _id *big.Int, _account common.Address, _whitelisted common.Address, _on bool) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setWhitelisted", _id, _account, _whitelisted, _on)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0xe6e21c75.
//
// Solidity: function setWhitelisted(uint256 _id, address _account, address _whitelisted, bool _on) returns()
func (_ERC1155 *ERC1155Session) SetWhitelisted(_id *big.Int, _account common.Address, _whitelisted common.Address, _on bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetWhitelisted(&_ERC1155.TransactOpts, _id, _account, _whitelisted, _on)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0xe6e21c75.
//
// Solidity: function setWhitelisted(uint256 _id, address _account, address _whitelisted, bool _on) returns()
func (_ERC1155 *ERC1155TransactorSession) SetWhitelisted(_id *big.Int, _account common.Address, _whitelisted common.Address, _on bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetWhitelisted(&_ERC1155.TransactOpts, _id, _account, _whitelisted, _on)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address _to, uint256 _id, uint256 _value) returns()
func (_ERC1155 *ERC1155Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "transfer", _to, _id, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address _to, uint256 _id, uint256 _value) returns()
func (_ERC1155 *ERC1155Session) Transfer(_to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Transfer(&_ERC1155.TransactOpts, _to, _id, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address _to, uint256 _id, uint256 _value) returns()
func (_ERC1155 *ERC1155TransactorSession) Transfer(_to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Transfer(&_ERC1155.TransactOpts, _to, _id, _value)
}

// TransferAdapter is a paid mutator transaction binding the contract method 0xff69ecd4.
//
// Solidity: function transferAdapter(address _to, uint256 _id, uint256 _value, address _msgSender) returns()
func (_ERC1155 *ERC1155Transactor) TransferAdapter(opts *bind.TransactOpts, _to common.Address, _id *big.Int, _value *big.Int, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "transferAdapter", _to, _id, _value, _msgSender)
}

// TransferAdapter is a paid mutator transaction binding the contract method 0xff69ecd4.
//
// Solidity: function transferAdapter(address _to, uint256 _id, uint256 _value, address _msgSender) returns()
func (_ERC1155 *ERC1155Session) TransferAdapter(_to common.Address, _id *big.Int, _value *big.Int, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.TransferAdapter(&_ERC1155.TransactOpts, _to, _id, _value, _msgSender)
}

// TransferAdapter is a paid mutator transaction binding the contract method 0xff69ecd4.
//
// Solidity: function transferAdapter(address _to, uint256 _id, uint256 _value, address _msgSender) returns()
func (_ERC1155 *ERC1155TransactorSession) TransferAdapter(_to common.Address, _id *big.Int, _value *big.Int, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.TransferAdapter(&_ERC1155.TransactOpts, _to, _id, _value, _msgSender)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _id, uint256 _value) returns()
func (_ERC1155 *ERC1155Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "transferFrom", _from, _to, _id, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _id, uint256 _value) returns()
func (_ERC1155 *ERC1155Session) TransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.TransferFrom(&_ERC1155.TransactOpts, _from, _to, _id, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _id, uint256 _value) returns()
func (_ERC1155 *ERC1155TransactorSession) TransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.TransferFrom(&_ERC1155.TransactOpts, _from, _to, _id, _value)
}

// TransferFromAdapter is a paid mutator transaction binding the contract method 0x52a55db2.
//
// Solidity: function transferFromAdapter(address _from, address _to, uint256 _id, uint256 _value, address _msgSender) returns()
func (_ERC1155 *ERC1155Transactor) TransferFromAdapter(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "transferFromAdapter", _from, _to, _id, _value, _msgSender)
}

// TransferFromAdapter is a paid mutator transaction binding the contract method 0x52a55db2.
//
// Solidity: function transferFromAdapter(address _from, address _to, uint256 _id, uint256 _value, address _msgSender) returns()
func (_ERC1155 *ERC1155Session) TransferFromAdapter(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.TransferFromAdapter(&_ERC1155.TransactOpts, _from, _to, _id, _value, _msgSender)
}

// TransferFromAdapter is a paid mutator transaction binding the contract method 0x52a55db2.
//
// Solidity: function transferFromAdapter(address _from, address _to, uint256 _id, uint256 _value, address _msgSender) returns()
func (_ERC1155 *ERC1155TransactorSession) TransferFromAdapter(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _msgSender common.Address) (*types.Transaction, error) {
	return _ERC1155.Contract.TransferFromAdapter(&_ERC1155.TransactOpts, _from, _to, _id, _value, _msgSender)
}

// UpdateName is a paid mutator transaction binding the contract method 0x53e76f2c.
//
// Solidity: function updateName(uint256 _id, string _name) returns()
func (_ERC1155 *ERC1155Transactor) UpdateName(opts *bind.TransactOpts, _id *big.Int, _name string) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "updateName", _id, _name)
}

// UpdateName is a paid mutator transaction binding the contract method 0x53e76f2c.
//
// Solidity: function updateName(uint256 _id, string _name) returns()
func (_ERC1155 *ERC1155Session) UpdateName(_id *big.Int, _name string) (*types.Transaction, error) {
	return _ERC1155.Contract.UpdateName(&_ERC1155.TransactOpts, _id, _name)
}

// UpdateName is a paid mutator transaction binding the contract method 0x53e76f2c.
//
// Solidity: function updateName(uint256 _id, string _name) returns()
func (_ERC1155 *ERC1155TransactorSession) UpdateName(_id *big.Int, _name string) (*types.Transaction, error) {
	return _ERC1155.Contract.UpdateName(&_ERC1155.TransactOpts, _id, _name)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ERC1155 *ERC1155Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ERC1155.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ERC1155 *ERC1155Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.Fallback(&_ERC1155.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ERC1155 *ERC1155TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.Fallback(&_ERC1155.TransactOpts, calldata)
}

// ERC1155AcceptAssignmentIterator is returned from FilterAcceptAssignment and is used to iterate over the raw logs and unpacked data for AcceptAssignment events raised by the ERC1155 contract.
type ERC1155AcceptAssignmentIterator struct {
	Event *ERC1155AcceptAssignment // Event containing the contract specifics and raw log

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
func (it *ERC1155AcceptAssignmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155AcceptAssignment)
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
		it.Event = new(ERC1155AcceptAssignment)
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
func (it *ERC1155AcceptAssignmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155AcceptAssignmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155AcceptAssignment represents a AcceptAssignment event raised by the ERC1155 contract.
type ERC1155AcceptAssignment struct {
	Id      *big.Int
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAcceptAssignment is a free log retrieval operation binding the contract event 0x918e15d498938a1120d3b78b391c92ca9921a35fe59080548f8ca665f8348eba.
//
// Solidity: event AcceptAssignment(uint256 indexed _id, address indexed _creator)
func (_ERC1155 *ERC1155Filterer) FilterAcceptAssignment(opts *bind.FilterOpts, _id []*big.Int, _creator []common.Address) (*ERC1155AcceptAssignmentIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "AcceptAssignment", _idRule, _creatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155AcceptAssignmentIterator{contract: _ERC1155.contract, event: "AcceptAssignment", logs: logs, sub: sub}, nil
}

// WatchAcceptAssignment is a free log subscription operation binding the contract event 0x918e15d498938a1120d3b78b391c92ca9921a35fe59080548f8ca665f8348eba.
//
// Solidity: event AcceptAssignment(uint256 indexed _id, address indexed _creator)
func (_ERC1155 *ERC1155Filterer) WatchAcceptAssignment(opts *bind.WatchOpts, sink chan<- *ERC1155AcceptAssignment, _id []*big.Int, _creator []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "AcceptAssignment", _idRule, _creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155AcceptAssignment)
				if err := _ERC1155.contract.UnpackLog(event, "AcceptAssignment", log); err != nil {
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

// ParseAcceptAssignment is a log parse operation binding the contract event 0x918e15d498938a1120d3b78b391c92ca9921a35fe59080548f8ca665f8348eba.
//
// Solidity: event AcceptAssignment(uint256 indexed _id, address indexed _creator)
func (_ERC1155 *ERC1155Filterer) ParseAcceptAssignment(log types.Log) (*ERC1155AcceptAssignment, error) {
	event := new(ERC1155AcceptAssignment)
	if err := _ERC1155.contract.UnpackLog(event, "AcceptAssignment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC1155 contract.
type ERC1155ApprovalIterator struct {
	Event *ERC1155Approval // Event containing the contract specifics and raw log

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
func (it *ERC1155ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Approval)
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
		it.Event = new(ERC1155Approval)
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
func (it *ERC1155ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Approval represents a Approval event raised by the ERC1155 contract.
type ERC1155Approval struct {
	Id      *big.Int
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x69e4aaf23f9318cf40839ac20453d8fbedaac2955eb08a27ae5189cc71925716.
//
// Solidity: event Approval(uint256 indexed _id, address indexed _owner, address indexed _spender, uint256 _value)
func (_ERC1155 *ERC1155Filterer) FilterApproval(opts *bind.FilterOpts, _id []*big.Int, _owner []common.Address, _spender []common.Address) (*ERC1155ApprovalIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Approval", _idRule, _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155ApprovalIterator{contract: _ERC1155.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x69e4aaf23f9318cf40839ac20453d8fbedaac2955eb08a27ae5189cc71925716.
//
// Solidity: event Approval(uint256 indexed _id, address indexed _owner, address indexed _spender, uint256 _value)
func (_ERC1155 *ERC1155Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC1155Approval, _id []*big.Int, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Approval", _idRule, _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Approval)
				if err := _ERC1155.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x69e4aaf23f9318cf40839ac20453d8fbedaac2955eb08a27ae5189cc71925716.
//
// Solidity: event Approval(uint256 indexed _id, address indexed _owner, address indexed _spender, uint256 _value)
func (_ERC1155 *ERC1155Filterer) ParseApproval(log types.Log) (*ERC1155Approval, error) {
	event := new(ERC1155Approval)
	if err := _ERC1155.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155AssignIterator is returned from FilterAssign and is used to iterate over the raw logs and unpacked data for Assign events raised by the ERC1155 contract.
type ERC1155AssignIterator struct {
	Event *ERC1155Assign // Event containing the contract specifics and raw log

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
func (it *ERC1155AssignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Assign)
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
		it.Event = new(ERC1155Assign)
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
func (it *ERC1155AssignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155AssignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Assign represents a Assign event raised by the ERC1155 contract.
type ERC1155Assign struct {
	Id   *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAssign is a free log retrieval operation binding the contract event 0x78eaa4bf5fdbde84109244ab35d7cdb379b3b146d0dd6c94050d86b662ea30ba.
//
// Solidity: event Assign(uint256 indexed _id, address indexed _from, address indexed _to)
func (_ERC1155 *ERC1155Filterer) FilterAssign(opts *bind.FilterOpts, _id []*big.Int, _from []common.Address, _to []common.Address) (*ERC1155AssignIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Assign", _idRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155AssignIterator{contract: _ERC1155.contract, event: "Assign", logs: logs, sub: sub}, nil
}

// WatchAssign is a free log subscription operation binding the contract event 0x78eaa4bf5fdbde84109244ab35d7cdb379b3b146d0dd6c94050d86b662ea30ba.
//
// Solidity: event Assign(uint256 indexed _id, address indexed _from, address indexed _to)
func (_ERC1155 *ERC1155Filterer) WatchAssign(opts *bind.WatchOpts, sink chan<- *ERC1155Assign, _id []*big.Int, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Assign", _idRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Assign)
				if err := _ERC1155.contract.UnpackLog(event, "Assign", log); err != nil {
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

// ParseAssign is a log parse operation binding the contract event 0x78eaa4bf5fdbde84109244ab35d7cdb379b3b146d0dd6c94050d86b662ea30ba.
//
// Solidity: event Assign(uint256 indexed _id, address indexed _from, address indexed _to)
func (_ERC1155 *ERC1155Filterer) ParseAssign(log types.Log) (*ERC1155Assign, error) {
	event := new(ERC1155Assign)
	if err := _ERC1155.contract.UnpackLog(event, "Assign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155CancelTradeIterator is returned from FilterCancelTrade and is used to iterate over the raw logs and unpacked data for CancelTrade events raised by the ERC1155 contract.
type ERC1155CancelTradeIterator struct {
	Event *ERC1155CancelTrade // Event containing the contract specifics and raw log

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
func (it *ERC1155CancelTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155CancelTrade)
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
		it.Event = new(ERC1155CancelTrade)
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
func (it *ERC1155CancelTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155CancelTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155CancelTrade represents a CancelTrade event raised by the ERC1155 contract.
type ERC1155CancelTrade struct {
	TradeId               *big.Int
	FirstParty            common.Address
	ReceivedEnjFirstParty *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCancelTrade is a free log retrieval operation binding the contract event 0x1555f98ec6d1cab7127dbc97d5abf49e0bfc5d2ebb95e628f7d3b20ed6ec9b93.
//
// Solidity: event CancelTrade(uint256 indexed _tradeId, address indexed _firstParty, uint256 _receivedEnjFirstParty)
func (_ERC1155 *ERC1155Filterer) FilterCancelTrade(opts *bind.FilterOpts, _tradeId []*big.Int, _firstParty []common.Address) (*ERC1155CancelTradeIterator, error) {

	var _tradeIdRule []interface{}
	for _, _tradeIdItem := range _tradeId {
		_tradeIdRule = append(_tradeIdRule, _tradeIdItem)
	}
	var _firstPartyRule []interface{}
	for _, _firstPartyItem := range _firstParty {
		_firstPartyRule = append(_firstPartyRule, _firstPartyItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "CancelTrade", _tradeIdRule, _firstPartyRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155CancelTradeIterator{contract: _ERC1155.contract, event: "CancelTrade", logs: logs, sub: sub}, nil
}

// WatchCancelTrade is a free log subscription operation binding the contract event 0x1555f98ec6d1cab7127dbc97d5abf49e0bfc5d2ebb95e628f7d3b20ed6ec9b93.
//
// Solidity: event CancelTrade(uint256 indexed _tradeId, address indexed _firstParty, uint256 _receivedEnjFirstParty)
func (_ERC1155 *ERC1155Filterer) WatchCancelTrade(opts *bind.WatchOpts, sink chan<- *ERC1155CancelTrade, _tradeId []*big.Int, _firstParty []common.Address) (event.Subscription, error) {

	var _tradeIdRule []interface{}
	for _, _tradeIdItem := range _tradeId {
		_tradeIdRule = append(_tradeIdRule, _tradeIdItem)
	}
	var _firstPartyRule []interface{}
	for _, _firstPartyItem := range _firstParty {
		_firstPartyRule = append(_firstPartyRule, _firstPartyItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "CancelTrade", _tradeIdRule, _firstPartyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155CancelTrade)
				if err := _ERC1155.contract.UnpackLog(event, "CancelTrade", log); err != nil {
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

// ParseCancelTrade is a log parse operation binding the contract event 0x1555f98ec6d1cab7127dbc97d5abf49e0bfc5d2ebb95e628f7d3b20ed6ec9b93.
//
// Solidity: event CancelTrade(uint256 indexed _tradeId, address indexed _firstParty, uint256 _receivedEnjFirstParty)
func (_ERC1155 *ERC1155Filterer) ParseCancelTrade(log types.Log) (*ERC1155CancelTrade, error) {
	event := new(ERC1155CancelTrade)
	if err := _ERC1155.contract.UnpackLog(event, "CancelTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155CompleteTradeIterator is returned from FilterCompleteTrade and is used to iterate over the raw logs and unpacked data for CompleteTrade events raised by the ERC1155 contract.
type ERC1155CompleteTradeIterator struct {
	Event *ERC1155CompleteTrade // Event containing the contract specifics and raw log

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
func (it *ERC1155CompleteTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155CompleteTrade)
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
		it.Event = new(ERC1155CompleteTrade)
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
func (it *ERC1155CompleteTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155CompleteTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155CompleteTrade represents a CompleteTrade event raised by the ERC1155 contract.
type ERC1155CompleteTrade struct {
	TradeId                *big.Int
	FirstParty             common.Address
	SecondParty            common.Address
	ReceivedEnjFirstParty  *big.Int
	ChangeEnjFirstParty    *big.Int
	ReceivedEnjSecondParty *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterCompleteTrade is a free log retrieval operation binding the contract event 0x288e5885bcfc1351450074e37b883e1d069889e7e883ced625610663afabd0f0.
//
// Solidity: event CompleteTrade(uint256 indexed _tradeId, address indexed _firstParty, address indexed _secondParty, uint256 _receivedEnjFirstParty, uint256 _changeEnjFirstParty, uint256 _receivedEnjSecondParty)
func (_ERC1155 *ERC1155Filterer) FilterCompleteTrade(opts *bind.FilterOpts, _tradeId []*big.Int, _firstParty []common.Address, _secondParty []common.Address) (*ERC1155CompleteTradeIterator, error) {

	var _tradeIdRule []interface{}
	for _, _tradeIdItem := range _tradeId {
		_tradeIdRule = append(_tradeIdRule, _tradeIdItem)
	}
	var _firstPartyRule []interface{}
	for _, _firstPartyItem := range _firstParty {
		_firstPartyRule = append(_firstPartyRule, _firstPartyItem)
	}
	var _secondPartyRule []interface{}
	for _, _secondPartyItem := range _secondParty {
		_secondPartyRule = append(_secondPartyRule, _secondPartyItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "CompleteTrade", _tradeIdRule, _firstPartyRule, _secondPartyRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155CompleteTradeIterator{contract: _ERC1155.contract, event: "CompleteTrade", logs: logs, sub: sub}, nil
}

// WatchCompleteTrade is a free log subscription operation binding the contract event 0x288e5885bcfc1351450074e37b883e1d069889e7e883ced625610663afabd0f0.
//
// Solidity: event CompleteTrade(uint256 indexed _tradeId, address indexed _firstParty, address indexed _secondParty, uint256 _receivedEnjFirstParty, uint256 _changeEnjFirstParty, uint256 _receivedEnjSecondParty)
func (_ERC1155 *ERC1155Filterer) WatchCompleteTrade(opts *bind.WatchOpts, sink chan<- *ERC1155CompleteTrade, _tradeId []*big.Int, _firstParty []common.Address, _secondParty []common.Address) (event.Subscription, error) {

	var _tradeIdRule []interface{}
	for _, _tradeIdItem := range _tradeId {
		_tradeIdRule = append(_tradeIdRule, _tradeIdItem)
	}
	var _firstPartyRule []interface{}
	for _, _firstPartyItem := range _firstParty {
		_firstPartyRule = append(_firstPartyRule, _firstPartyItem)
	}
	var _secondPartyRule []interface{}
	for _, _secondPartyItem := range _secondParty {
		_secondPartyRule = append(_secondPartyRule, _secondPartyItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "CompleteTrade", _tradeIdRule, _firstPartyRule, _secondPartyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155CompleteTrade)
				if err := _ERC1155.contract.UnpackLog(event, "CompleteTrade", log); err != nil {
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

// ParseCompleteTrade is a log parse operation binding the contract event 0x288e5885bcfc1351450074e37b883e1d069889e7e883ced625610663afabd0f0.
//
// Solidity: event CompleteTrade(uint256 indexed _tradeId, address indexed _firstParty, address indexed _secondParty, uint256 _receivedEnjFirstParty, uint256 _changeEnjFirstParty, uint256 _receivedEnjSecondParty)
func (_ERC1155 *ERC1155Filterer) ParseCompleteTrade(log types.Log) (*ERC1155CompleteTrade, error) {
	event := new(ERC1155CompleteTrade)
	if err := _ERC1155.contract.UnpackLog(event, "CompleteTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155CreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the ERC1155 contract.
type ERC1155CreateIterator struct {
	Event *ERC1155Create // Event containing the contract specifics and raw log

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
func (it *ERC1155CreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Create)
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
		it.Event = new(ERC1155Create)
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
func (it *ERC1155CreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155CreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Create represents a Create event raised by the ERC1155 contract.
type ERC1155Create struct {
	Id            *big.Int
	Creator       common.Address
	IsNonFungible bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x250ed6814ddcc5fc06eec40c015c413d3aa7bfc4e1df91ed205e0d71f0a9408f.
//
// Solidity: event Create(uint256 indexed _id, address indexed _creator, bool _isNonFungible)
func (_ERC1155 *ERC1155Filterer) FilterCreate(opts *bind.FilterOpts, _id []*big.Int, _creator []common.Address) (*ERC1155CreateIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Create", _idRule, _creatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155CreateIterator{contract: _ERC1155.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x250ed6814ddcc5fc06eec40c015c413d3aa7bfc4e1df91ed205e0d71f0a9408f.
//
// Solidity: event Create(uint256 indexed _id, address indexed _creator, bool _isNonFungible)
func (_ERC1155 *ERC1155Filterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *ERC1155Create, _id []*big.Int, _creator []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Create", _idRule, _creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Create)
				if err := _ERC1155.contract.UnpackLog(event, "Create", log); err != nil {
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

// ParseCreate is a log parse operation binding the contract event 0x250ed6814ddcc5fc06eec40c015c413d3aa7bfc4e1df91ed205e0d71f0a9408f.
//
// Solidity: event Create(uint256 indexed _id, address indexed _creator, bool _isNonFungible)
func (_ERC1155 *ERC1155Filterer) ParseCreate(log types.Log) (*ERC1155Create, error) {
	event := new(ERC1155Create)
	if err := _ERC1155.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155CreateTradeIterator is returned from FilterCreateTrade and is used to iterate over the raw logs and unpacked data for CreateTrade events raised by the ERC1155 contract.
type ERC1155CreateTradeIterator struct {
	Event *ERC1155CreateTrade // Event containing the contract specifics and raw log

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
func (it *ERC1155CreateTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155CreateTrade)
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
		it.Event = new(ERC1155CreateTrade)
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
func (it *ERC1155CreateTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155CreateTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155CreateTrade represents a CreateTrade event raised by the ERC1155 contract.
type ERC1155CreateTrade struct {
	TradeId               *big.Int
	FirstParty            common.Address
	SecondParty           common.Address
	EscrowedEnjFirstParty *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCreateTrade is a free log retrieval operation binding the contract event 0xec16055b5b6ecb313865a3862d1d169101c9700959b7cae5bfb25c0f319e97de.
//
// Solidity: event CreateTrade(uint256 indexed _tradeId, address indexed _firstParty, address indexed _secondParty, uint256 _escrowedEnjFirstParty)
func (_ERC1155 *ERC1155Filterer) FilterCreateTrade(opts *bind.FilterOpts, _tradeId []*big.Int, _firstParty []common.Address, _secondParty []common.Address) (*ERC1155CreateTradeIterator, error) {

	var _tradeIdRule []interface{}
	for _, _tradeIdItem := range _tradeId {
		_tradeIdRule = append(_tradeIdRule, _tradeIdItem)
	}
	var _firstPartyRule []interface{}
	for _, _firstPartyItem := range _firstParty {
		_firstPartyRule = append(_firstPartyRule, _firstPartyItem)
	}
	var _secondPartyRule []interface{}
	for _, _secondPartyItem := range _secondParty {
		_secondPartyRule = append(_secondPartyRule, _secondPartyItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "CreateTrade", _tradeIdRule, _firstPartyRule, _secondPartyRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155CreateTradeIterator{contract: _ERC1155.contract, event: "CreateTrade", logs: logs, sub: sub}, nil
}

// WatchCreateTrade is a free log subscription operation binding the contract event 0xec16055b5b6ecb313865a3862d1d169101c9700959b7cae5bfb25c0f319e97de.
//
// Solidity: event CreateTrade(uint256 indexed _tradeId, address indexed _firstParty, address indexed _secondParty, uint256 _escrowedEnjFirstParty)
func (_ERC1155 *ERC1155Filterer) WatchCreateTrade(opts *bind.WatchOpts, sink chan<- *ERC1155CreateTrade, _tradeId []*big.Int, _firstParty []common.Address, _secondParty []common.Address) (event.Subscription, error) {

	var _tradeIdRule []interface{}
	for _, _tradeIdItem := range _tradeId {
		_tradeIdRule = append(_tradeIdRule, _tradeIdItem)
	}
	var _firstPartyRule []interface{}
	for _, _firstPartyItem := range _firstParty {
		_firstPartyRule = append(_firstPartyRule, _firstPartyItem)
	}
	var _secondPartyRule []interface{}
	for _, _secondPartyItem := range _secondParty {
		_secondPartyRule = append(_secondPartyRule, _secondPartyItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "CreateTrade", _tradeIdRule, _firstPartyRule, _secondPartyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155CreateTrade)
				if err := _ERC1155.contract.UnpackLog(event, "CreateTrade", log); err != nil {
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

// ParseCreateTrade is a log parse operation binding the contract event 0xec16055b5b6ecb313865a3862d1d169101c9700959b7cae5bfb25c0f319e97de.
//
// Solidity: event CreateTrade(uint256 indexed _tradeId, address indexed _firstParty, address indexed _secondParty, uint256 _escrowedEnjFirstParty)
func (_ERC1155 *ERC1155Filterer) ParseCreateTrade(log types.Log) (*ERC1155CreateTrade, error) {
	event := new(ERC1155CreateTrade)
	if err := _ERC1155.contract.UnpackLog(event, "CreateTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155DeployERCAdapterIterator is returned from FilterDeployERCAdapter and is used to iterate over the raw logs and unpacked data for DeployERCAdapter events raised by the ERC1155 contract.
type ERC1155DeployERCAdapterIterator struct {
	Event *ERC1155DeployERCAdapter // Event containing the contract specifics and raw log

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
func (it *ERC1155DeployERCAdapterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155DeployERCAdapter)
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
		it.Event = new(ERC1155DeployERCAdapter)
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
func (it *ERC1155DeployERCAdapterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155DeployERCAdapterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155DeployERCAdapter represents a DeployERCAdapter event raised by the ERC1155 contract.
type ERC1155DeployERCAdapter struct {
	Id     *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeployERCAdapter is a free log retrieval operation binding the contract event 0xc2dad9e937f7f5f669e13cada0909d2612ace1aed9b91c47d3632af5e1f32fb6.
//
// Solidity: event DeployERCAdapter(uint256 indexed _id, address indexed _sender)
func (_ERC1155 *ERC1155Filterer) FilterDeployERCAdapter(opts *bind.FilterOpts, _id []*big.Int, _sender []common.Address) (*ERC1155DeployERCAdapterIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "DeployERCAdapter", _idRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155DeployERCAdapterIterator{contract: _ERC1155.contract, event: "DeployERCAdapter", logs: logs, sub: sub}, nil
}

// WatchDeployERCAdapter is a free log subscription operation binding the contract event 0xc2dad9e937f7f5f669e13cada0909d2612ace1aed9b91c47d3632af5e1f32fb6.
//
// Solidity: event DeployERCAdapter(uint256 indexed _id, address indexed _sender)
func (_ERC1155 *ERC1155Filterer) WatchDeployERCAdapter(opts *bind.WatchOpts, sink chan<- *ERC1155DeployERCAdapter, _id []*big.Int, _sender []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "DeployERCAdapter", _idRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155DeployERCAdapter)
				if err := _ERC1155.contract.UnpackLog(event, "DeployERCAdapter", log); err != nil {
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

// ParseDeployERCAdapter is a log parse operation binding the contract event 0xc2dad9e937f7f5f669e13cada0909d2612ace1aed9b91c47d3632af5e1f32fb6.
//
// Solidity: event DeployERCAdapter(uint256 indexed _id, address indexed _sender)
func (_ERC1155 *ERC1155Filterer) ParseDeployERCAdapter(log types.Log) (*ERC1155DeployERCAdapter, error) {
	event := new(ERC1155DeployERCAdapter)
	if err := _ERC1155.contract.UnpackLog(event, "DeployERCAdapter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155InitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the ERC1155 contract.
type ERC1155InitializeIterator struct {
	Event *ERC1155Initialize // Event containing the contract specifics and raw log

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
func (it *ERC1155InitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Initialize)
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
		it.Event = new(ERC1155Initialize)
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
func (it *ERC1155InitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155InitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Initialize represents a Initialize event raised by the ERC1155 contract.
type ERC1155Initialize struct {
	Block       *big.Int
	Storage     common.Address
	OldContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0xd49a56a3e5838c9132e8acda6a2d755d8fdc303f9964b901cfbbb3059876d886.
//
// Solidity: event Initialize(uint256 _block, address _storage, address _oldContract)
func (_ERC1155 *ERC1155Filterer) FilterInitialize(opts *bind.FilterOpts) (*ERC1155InitializeIterator, error) {

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &ERC1155InitializeIterator{contract: _ERC1155.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0xd49a56a3e5838c9132e8acda6a2d755d8fdc303f9964b901cfbbb3059876d886.
//
// Solidity: event Initialize(uint256 _block, address _storage, address _oldContract)
func (_ERC1155 *ERC1155Filterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *ERC1155Initialize) (event.Subscription, error) {

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Initialize)
				if err := _ERC1155.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0xd49a56a3e5838c9132e8acda6a2d755d8fdc303f9964b901cfbbb3059876d886.
//
// Solidity: event Initialize(uint256 _block, address _storage, address _oldContract)
func (_ERC1155 *ERC1155Filterer) ParseInitialize(log types.Log) (*ERC1155Initialize, error) {
	event := new(ERC1155Initialize)
	if err := _ERC1155.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155LogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the ERC1155 contract.
type ERC1155LogIterator struct {
	Event *ERC1155Log // Event containing the contract specifics and raw log

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
func (it *ERC1155LogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Log)
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
		it.Event = new(ERC1155Log)
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
func (it *ERC1155LogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155LogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Log represents a Log event raised by the ERC1155 contract.
type ERC1155Log struct {
	Id   *big.Int
	From common.Address
	Data string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x75bcec5ba6a2febb176c494c16fef3f009bf87d476335ab48b702384710481c9.
//
// Solidity: event Log(uint256 indexed _id, address indexed _from, string _data)
func (_ERC1155 *ERC1155Filterer) FilterLog(opts *bind.FilterOpts, _id []*big.Int, _from []common.Address) (*ERC1155LogIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Log", _idRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155LogIterator{contract: _ERC1155.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x75bcec5ba6a2febb176c494c16fef3f009bf87d476335ab48b702384710481c9.
//
// Solidity: event Log(uint256 indexed _id, address indexed _from, string _data)
func (_ERC1155 *ERC1155Filterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ERC1155Log, _id []*big.Int, _from []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Log", _idRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Log)
				if err := _ERC1155.contract.UnpackLog(event, "Log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x75bcec5ba6a2febb176c494c16fef3f009bf87d476335ab48b702384710481c9.
//
// Solidity: event Log(uint256 indexed _id, address indexed _from, string _data)
func (_ERC1155 *ERC1155Filterer) ParseLog(log types.Log) (*ERC1155Log, error) {
	event := new(ERC1155Log)
	if err := _ERC1155.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155MeltIterator is returned from FilterMelt and is used to iterate over the raw logs and unpacked data for Melt events raised by the ERC1155 contract.
type ERC1155MeltIterator struct {
	Event *ERC1155Melt // Event containing the contract specifics and raw log

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
func (it *ERC1155MeltIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Melt)
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
		it.Event = new(ERC1155Melt)
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
func (it *ERC1155MeltIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155MeltIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Melt represents a Melt event raised by the ERC1155 contract.
type ERC1155Melt struct {
	Id    *big.Int
	Owner common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMelt is a free log retrieval operation binding the contract event 0xba6a480970167b03ed2f35b55c48a436cd01efe96abdf846d1a64da47df0e6d9.
//
// Solidity: event Melt(uint256 indexed _id, address indexed _owner, uint256 _value)
func (_ERC1155 *ERC1155Filterer) FilterMelt(opts *bind.FilterOpts, _id []*big.Int, _owner []common.Address) (*ERC1155MeltIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Melt", _idRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155MeltIterator{contract: _ERC1155.contract, event: "Melt", logs: logs, sub: sub}, nil
}

// WatchMelt is a free log subscription operation binding the contract event 0xba6a480970167b03ed2f35b55c48a436cd01efe96abdf846d1a64da47df0e6d9.
//
// Solidity: event Melt(uint256 indexed _id, address indexed _owner, uint256 _value)
func (_ERC1155 *ERC1155Filterer) WatchMelt(opts *bind.WatchOpts, sink chan<- *ERC1155Melt, _id []*big.Int, _owner []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Melt", _idRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Melt)
				if err := _ERC1155.contract.UnpackLog(event, "Melt", log); err != nil {
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

// ParseMelt is a log parse operation binding the contract event 0xba6a480970167b03ed2f35b55c48a436cd01efe96abdf846d1a64da47df0e6d9.
//
// Solidity: event Melt(uint256 indexed _id, address indexed _owner, uint256 _value)
func (_ERC1155 *ERC1155Filterer) ParseMelt(log types.Log) (*ERC1155Melt, error) {
	event := new(ERC1155Melt)
	if err := _ERC1155.contract.UnpackLog(event, "Melt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ERC1155 contract.
type ERC1155MintIterator struct {
	Event *ERC1155Mint // Event containing the contract specifics and raw log

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
func (it *ERC1155MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Mint)
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
		it.Event = new(ERC1155Mint)
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
func (it *ERC1155MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Mint represents a Mint event raised by the ERC1155 contract.
type ERC1155Mint struct {
	Id    *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xcc9c58b575eabd3f6a1ee653e91fcea3ff546867ffc3782a3bbca1f9b6dbb8df.
//
// Solidity: event Mint(uint256 indexed _id, uint256 _value)
func (_ERC1155 *ERC1155Filterer) FilterMint(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155MintIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Mint", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155MintIterator{contract: _ERC1155.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xcc9c58b575eabd3f6a1ee653e91fcea3ff546867ffc3782a3bbca1f9b6dbb8df.
//
// Solidity: event Mint(uint256 indexed _id, uint256 _value)
func (_ERC1155 *ERC1155Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ERC1155Mint, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Mint", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Mint)
				if err := _ERC1155.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xcc9c58b575eabd3f6a1ee653e91fcea3ff546867ffc3782a3bbca1f9b6dbb8df.
//
// Solidity: event Mint(uint256 indexed _id, uint256 _value)
func (_ERC1155 *ERC1155Filterer) ParseMint(log types.Log) (*ERC1155Mint, error) {
	event := new(ERC1155Mint)
	if err := _ERC1155.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155OperatorApprovalIterator is returned from FilterOperatorApproval and is used to iterate over the raw logs and unpacked data for OperatorApproval events raised by the ERC1155 contract.
type ERC1155OperatorApprovalIterator struct {
	Event *ERC1155OperatorApproval // Event containing the contract specifics and raw log

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
func (it *ERC1155OperatorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155OperatorApproval)
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
		it.Event = new(ERC1155OperatorApproval)
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
func (it *ERC1155OperatorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155OperatorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155OperatorApproval represents a OperatorApproval event raised by the ERC1155 contract.
type ERC1155OperatorApproval struct {
	Owner    common.Address
	Operator common.Address
	Id       *big.Int
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorApproval is a free log retrieval operation binding the contract event 0xc55152c6b6552357df350ada4b090cece4431b93e48dafe22a243786e294f08d.
//
// Solidity: event OperatorApproval(address indexed _owner, address indexed _operator, uint256 indexed _id, bool _approved)
func (_ERC1155 *ERC1155Filterer) FilterOperatorApproval(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address, _id []*big.Int) (*ERC1155OperatorApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "OperatorApproval", _ownerRule, _operatorRule, _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155OperatorApprovalIterator{contract: _ERC1155.contract, event: "OperatorApproval", logs: logs, sub: sub}, nil
}

// WatchOperatorApproval is a free log subscription operation binding the contract event 0xc55152c6b6552357df350ada4b090cece4431b93e48dafe22a243786e294f08d.
//
// Solidity: event OperatorApproval(address indexed _owner, address indexed _operator, uint256 indexed _id, bool _approved)
func (_ERC1155 *ERC1155Filterer) WatchOperatorApproval(opts *bind.WatchOpts, sink chan<- *ERC1155OperatorApproval, _owner []common.Address, _operator []common.Address, _id []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "OperatorApproval", _ownerRule, _operatorRule, _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155OperatorApproval)
				if err := _ERC1155.contract.UnpackLog(event, "OperatorApproval", log); err != nil {
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

// ParseOperatorApproval is a log parse operation binding the contract event 0xc55152c6b6552357df350ada4b090cece4431b93e48dafe22a243786e294f08d.
//
// Solidity: event OperatorApproval(address indexed _owner, address indexed _operator, uint256 indexed _id, bool _approved)
func (_ERC1155 *ERC1155Filterer) ParseOperatorApproval(log types.Log) (*ERC1155OperatorApproval, error) {
	event := new(ERC1155OperatorApproval)
	if err := _ERC1155.contract.UnpackLog(event, "OperatorApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155RetireIterator is returned from FilterRetire and is used to iterate over the raw logs and unpacked data for Retire events raised by the ERC1155 contract.
type ERC1155RetireIterator struct {
	Event *ERC1155Retire // Event containing the contract specifics and raw log

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
func (it *ERC1155RetireIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Retire)
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
		it.Event = new(ERC1155Retire)
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
func (it *ERC1155RetireIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155RetireIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Retire represents a Retire event raised by the ERC1155 contract.
type ERC1155Retire struct {
	Block        *big.Int
	NextContract common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRetire is a free log retrieval operation binding the contract event 0x63c9d69c60481ca0343e92ab401e8607082997131c6f61eafefa7bbae719f46b.
//
// Solidity: event Retire(uint256 _block, address _nextContract)
func (_ERC1155 *ERC1155Filterer) FilterRetire(opts *bind.FilterOpts) (*ERC1155RetireIterator, error) {

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Retire")
	if err != nil {
		return nil, err
	}
	return &ERC1155RetireIterator{contract: _ERC1155.contract, event: "Retire", logs: logs, sub: sub}, nil
}

// WatchRetire is a free log subscription operation binding the contract event 0x63c9d69c60481ca0343e92ab401e8607082997131c6f61eafefa7bbae719f46b.
//
// Solidity: event Retire(uint256 _block, address _nextContract)
func (_ERC1155 *ERC1155Filterer) WatchRetire(opts *bind.WatchOpts, sink chan<- *ERC1155Retire) (event.Subscription, error) {

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Retire")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Retire)
				if err := _ERC1155.contract.UnpackLog(event, "Retire", log); err != nil {
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

// ParseRetire is a log parse operation binding the contract event 0x63c9d69c60481ca0343e92ab401e8607082997131c6f61eafefa7bbae719f46b.
//
// Solidity: event Retire(uint256 _block, address _nextContract)
func (_ERC1155 *ERC1155Filterer) ParseRetire(log types.Log) (*ERC1155Retire, error) {
	event := new(ERC1155Retire)
	if err := _ERC1155.contract.UnpackLog(event, "Retire", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155SetURIIterator is returned from FilterSetURI and is used to iterate over the raw logs and unpacked data for SetURI events raised by the ERC1155 contract.
type ERC1155SetURIIterator struct {
	Event *ERC1155SetURI // Event containing the contract specifics and raw log

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
func (it *ERC1155SetURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155SetURI)
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
		it.Event = new(ERC1155SetURI)
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
func (it *ERC1155SetURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155SetURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155SetURI represents a SetURI event raised by the ERC1155 contract.
type ERC1155SetURI struct {
	Id  *big.Int
	Uri string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetURI is a free log retrieval operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed _id, string _uri)
func (_ERC1155 *ERC1155Filterer) FilterSetURI(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155SetURIIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "SetURI", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155SetURIIterator{contract: _ERC1155.contract, event: "SetURI", logs: logs, sub: sub}, nil
}

// WatchSetURI is a free log subscription operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed _id, string _uri)
func (_ERC1155 *ERC1155Filterer) WatchSetURI(opts *bind.WatchOpts, sink chan<- *ERC1155SetURI, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "SetURI", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155SetURI)
				if err := _ERC1155.contract.UnpackLog(event, "SetURI", log); err != nil {
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

// ParseSetURI is a log parse operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed _id, string _uri)
func (_ERC1155 *ERC1155Filterer) ParseSetURI(log types.Log) (*ERC1155SetURI, error) {
	event := new(ERC1155SetURI)
	if err := _ERC1155.contract.UnpackLog(event, "SetURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC1155 contract.
type ERC1155TransferIterator struct {
	Event *ERC1155Transfer // Event containing the contract specifics and raw log

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
func (it *ERC1155TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Transfer)
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
		it.Event = new(ERC1155Transfer)
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
func (it *ERC1155TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Transfer represents a Transfer event raised by the ERC1155 contract.
type ERC1155Transfer struct {
	Id    *big.Int
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xf2dbd98d79f00f7aff338b824931d607bfcc63d47307162470f25a055102d3b0.
//
// Solidity: event Transfer(uint256 indexed _id, address indexed _from, address indexed _to, uint256 _value)
func (_ERC1155 *ERC1155Filterer) FilterTransfer(opts *bind.FilterOpts, _id []*big.Int, _from []common.Address, _to []common.Address) (*ERC1155TransferIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Transfer", _idRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferIterator{contract: _ERC1155.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xf2dbd98d79f00f7aff338b824931d607bfcc63d47307162470f25a055102d3b0.
//
// Solidity: event Transfer(uint256 indexed _id, address indexed _from, address indexed _to, uint256 _value)
func (_ERC1155 *ERC1155Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC1155Transfer, _id []*big.Int, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Transfer", _idRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Transfer)
				if err := _ERC1155.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xf2dbd98d79f00f7aff338b824931d607bfcc63d47307162470f25a055102d3b0.
//
// Solidity: event Transfer(uint256 indexed _id, address indexed _from, address indexed _to, uint256 _value)
func (_ERC1155 *ERC1155Filterer) ParseTransfer(log types.Log) (*ERC1155Transfer, error) {
	event := new(ERC1155Transfer)
	if err := _ERC1155.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferFeeIterator is returned from FilterTransferFee and is used to iterate over the raw logs and unpacked data for TransferFee events raised by the ERC1155 contract.
type ERC1155TransferFeeIterator struct {
	Event *ERC1155TransferFee // Event containing the contract specifics and raw log

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
func (it *ERC1155TransferFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferFee)
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
		it.Event = new(ERC1155TransferFee)
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
func (it *ERC1155TransferFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferFee represents a TransferFee event raised by the ERC1155 contract.
type ERC1155TransferFee struct {
	Id       *big.Int
	Sender   common.Address
	FeeId    *big.Int
	FeeValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferFee is a free log retrieval operation binding the contract event 0x298f4be3ff40398fb58555008f42eb86e257a96bba5c0f3c32814fa869535fc2.
//
// Solidity: event TransferFee(uint256 indexed _id, address indexed _sender, uint256 indexed _feeId, uint256 _feeValue)
func (_ERC1155 *ERC1155Filterer) FilterTransferFee(opts *bind.FilterOpts, _id []*big.Int, _sender []common.Address, _feeId []*big.Int) (*ERC1155TransferFeeIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _feeIdRule []interface{}
	for _, _feeIdItem := range _feeId {
		_feeIdRule = append(_feeIdRule, _feeIdItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferFee", _idRule, _senderRule, _feeIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferFeeIterator{contract: _ERC1155.contract, event: "TransferFee", logs: logs, sub: sub}, nil
}

// WatchTransferFee is a free log subscription operation binding the contract event 0x298f4be3ff40398fb58555008f42eb86e257a96bba5c0f3c32814fa869535fc2.
//
// Solidity: event TransferFee(uint256 indexed _id, address indexed _sender, uint256 indexed _feeId, uint256 _feeValue)
func (_ERC1155 *ERC1155Filterer) WatchTransferFee(opts *bind.WatchOpts, sink chan<- *ERC1155TransferFee, _id []*big.Int, _sender []common.Address, _feeId []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _feeIdRule []interface{}
	for _, _feeIdItem := range _feeId {
		_feeIdRule = append(_feeIdRule, _feeIdItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferFee", _idRule, _senderRule, _feeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferFee)
				if err := _ERC1155.contract.UnpackLog(event, "TransferFee", log); err != nil {
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

// ParseTransferFee is a log parse operation binding the contract event 0x298f4be3ff40398fb58555008f42eb86e257a96bba5c0f3c32814fa869535fc2.
//
// Solidity: event TransferFee(uint256 indexed _id, address indexed _sender, uint256 indexed _feeId, uint256 _feeValue)
func (_ERC1155 *ERC1155Filterer) ParseTransferFee(log types.Log) (*ERC1155TransferFee, error) {
	event := new(ERC1155TransferFee)
	if err := _ERC1155.contract.UnpackLog(event, "TransferFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155UpdateDecimalsIterator is returned from FilterUpdateDecimals and is used to iterate over the raw logs and unpacked data for UpdateDecimals events raised by the ERC1155 contract.
type ERC1155UpdateDecimalsIterator struct {
	Event *ERC1155UpdateDecimals // Event containing the contract specifics and raw log

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
func (it *ERC1155UpdateDecimalsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155UpdateDecimals)
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
		it.Event = new(ERC1155UpdateDecimals)
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
func (it *ERC1155UpdateDecimalsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155UpdateDecimalsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155UpdateDecimals represents a UpdateDecimals event raised by the ERC1155 contract.
type ERC1155UpdateDecimals struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateDecimals is a free log retrieval operation binding the contract event 0xc6e60c4f2dca4443a6098cb6bb0aebd338eb8fb256652d8f5f83cd577d6faa5f.
//
// Solidity: event UpdateDecimals(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) FilterUpdateDecimals(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155UpdateDecimalsIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "UpdateDecimals", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155UpdateDecimalsIterator{contract: _ERC1155.contract, event: "UpdateDecimals", logs: logs, sub: sub}, nil
}

// WatchUpdateDecimals is a free log subscription operation binding the contract event 0xc6e60c4f2dca4443a6098cb6bb0aebd338eb8fb256652d8f5f83cd577d6faa5f.
//
// Solidity: event UpdateDecimals(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) WatchUpdateDecimals(opts *bind.WatchOpts, sink chan<- *ERC1155UpdateDecimals, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "UpdateDecimals", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155UpdateDecimals)
				if err := _ERC1155.contract.UnpackLog(event, "UpdateDecimals", log); err != nil {
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

// ParseUpdateDecimals is a log parse operation binding the contract event 0xc6e60c4f2dca4443a6098cb6bb0aebd338eb8fb256652d8f5f83cd577d6faa5f.
//
// Solidity: event UpdateDecimals(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) ParseUpdateDecimals(log types.Log) (*ERC1155UpdateDecimals, error) {
	event := new(ERC1155UpdateDecimals)
	if err := _ERC1155.contract.UnpackLog(event, "UpdateDecimals", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155UpdateMaxMeltFeeIterator is returned from FilterUpdateMaxMeltFee and is used to iterate over the raw logs and unpacked data for UpdateMaxMeltFee events raised by the ERC1155 contract.
type ERC1155UpdateMaxMeltFeeIterator struct {
	Event *ERC1155UpdateMaxMeltFee // Event containing the contract specifics and raw log

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
func (it *ERC1155UpdateMaxMeltFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155UpdateMaxMeltFee)
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
		it.Event = new(ERC1155UpdateMaxMeltFee)
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
func (it *ERC1155UpdateMaxMeltFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155UpdateMaxMeltFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155UpdateMaxMeltFee represents a UpdateMaxMeltFee event raised by the ERC1155 contract.
type ERC1155UpdateMaxMeltFee struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxMeltFee is a free log retrieval operation binding the contract event 0x48b634f9512e480422be213cc551f1b7be53d86664006a468ca2fe7782408438.
//
// Solidity: event UpdateMaxMeltFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) FilterUpdateMaxMeltFee(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155UpdateMaxMeltFeeIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "UpdateMaxMeltFee", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155UpdateMaxMeltFeeIterator{contract: _ERC1155.contract, event: "UpdateMaxMeltFee", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxMeltFee is a free log subscription operation binding the contract event 0x48b634f9512e480422be213cc551f1b7be53d86664006a468ca2fe7782408438.
//
// Solidity: event UpdateMaxMeltFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) WatchUpdateMaxMeltFee(opts *bind.WatchOpts, sink chan<- *ERC1155UpdateMaxMeltFee, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "UpdateMaxMeltFee", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155UpdateMaxMeltFee)
				if err := _ERC1155.contract.UnpackLog(event, "UpdateMaxMeltFee", log); err != nil {
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

// ParseUpdateMaxMeltFee is a log parse operation binding the contract event 0x48b634f9512e480422be213cc551f1b7be53d86664006a468ca2fe7782408438.
//
// Solidity: event UpdateMaxMeltFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) ParseUpdateMaxMeltFee(log types.Log) (*ERC1155UpdateMaxMeltFee, error) {
	event := new(ERC1155UpdateMaxMeltFee)
	if err := _ERC1155.contract.UnpackLog(event, "UpdateMaxMeltFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155UpdateMaxTransferFeeIterator is returned from FilterUpdateMaxTransferFee and is used to iterate over the raw logs and unpacked data for UpdateMaxTransferFee events raised by the ERC1155 contract.
type ERC1155UpdateMaxTransferFeeIterator struct {
	Event *ERC1155UpdateMaxTransferFee // Event containing the contract specifics and raw log

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
func (it *ERC1155UpdateMaxTransferFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155UpdateMaxTransferFee)
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
		it.Event = new(ERC1155UpdateMaxTransferFee)
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
func (it *ERC1155UpdateMaxTransferFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155UpdateMaxTransferFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155UpdateMaxTransferFee represents a UpdateMaxTransferFee event raised by the ERC1155 contract.
type ERC1155UpdateMaxTransferFee struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxTransferFee is a free log retrieval operation binding the contract event 0xcdb0342f044c947e0fab52d7ffb8ac8f24c9989d29d2f8c6bb0fd6a94470a2e0.
//
// Solidity: event UpdateMaxTransferFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) FilterUpdateMaxTransferFee(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155UpdateMaxTransferFeeIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "UpdateMaxTransferFee", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155UpdateMaxTransferFeeIterator{contract: _ERC1155.contract, event: "UpdateMaxTransferFee", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxTransferFee is a free log subscription operation binding the contract event 0xcdb0342f044c947e0fab52d7ffb8ac8f24c9989d29d2f8c6bb0fd6a94470a2e0.
//
// Solidity: event UpdateMaxTransferFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) WatchUpdateMaxTransferFee(opts *bind.WatchOpts, sink chan<- *ERC1155UpdateMaxTransferFee, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "UpdateMaxTransferFee", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155UpdateMaxTransferFee)
				if err := _ERC1155.contract.UnpackLog(event, "UpdateMaxTransferFee", log); err != nil {
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

// ParseUpdateMaxTransferFee is a log parse operation binding the contract event 0xcdb0342f044c947e0fab52d7ffb8ac8f24c9989d29d2f8c6bb0fd6a94470a2e0.
//
// Solidity: event UpdateMaxTransferFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) ParseUpdateMaxTransferFee(log types.Log) (*ERC1155UpdateMaxTransferFee, error) {
	event := new(ERC1155UpdateMaxTransferFee)
	if err := _ERC1155.contract.UnpackLog(event, "UpdateMaxTransferFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155UpdateMeltFeeIterator is returned from FilterUpdateMeltFee and is used to iterate over the raw logs and unpacked data for UpdateMeltFee events raised by the ERC1155 contract.
type ERC1155UpdateMeltFeeIterator struct {
	Event *ERC1155UpdateMeltFee // Event containing the contract specifics and raw log

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
func (it *ERC1155UpdateMeltFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155UpdateMeltFee)
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
		it.Event = new(ERC1155UpdateMeltFee)
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
func (it *ERC1155UpdateMeltFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155UpdateMeltFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155UpdateMeltFee represents a UpdateMeltFee event raised by the ERC1155 contract.
type ERC1155UpdateMeltFee struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateMeltFee is a free log retrieval operation binding the contract event 0x7f0cd16fa546c980060c029fd05aea1347eb18597b335cfc714fa9e401cb0288.
//
// Solidity: event UpdateMeltFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) FilterUpdateMeltFee(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155UpdateMeltFeeIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "UpdateMeltFee", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155UpdateMeltFeeIterator{contract: _ERC1155.contract, event: "UpdateMeltFee", logs: logs, sub: sub}, nil
}

// WatchUpdateMeltFee is a free log subscription operation binding the contract event 0x7f0cd16fa546c980060c029fd05aea1347eb18597b335cfc714fa9e401cb0288.
//
// Solidity: event UpdateMeltFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) WatchUpdateMeltFee(opts *bind.WatchOpts, sink chan<- *ERC1155UpdateMeltFee, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "UpdateMeltFee", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155UpdateMeltFee)
				if err := _ERC1155.contract.UnpackLog(event, "UpdateMeltFee", log); err != nil {
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

// ParseUpdateMeltFee is a log parse operation binding the contract event 0x7f0cd16fa546c980060c029fd05aea1347eb18597b335cfc714fa9e401cb0288.
//
// Solidity: event UpdateMeltFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) ParseUpdateMeltFee(log types.Log) (*ERC1155UpdateMeltFee, error) {
	event := new(ERC1155UpdateMeltFee)
	if err := _ERC1155.contract.UnpackLog(event, "UpdateMeltFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155UpdateNameIterator is returned from FilterUpdateName and is used to iterate over the raw logs and unpacked data for UpdateName events raised by the ERC1155 contract.
type ERC1155UpdateNameIterator struct {
	Event *ERC1155UpdateName // Event containing the contract specifics and raw log

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
func (it *ERC1155UpdateNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155UpdateName)
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
		it.Event = new(ERC1155UpdateName)
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
func (it *ERC1155UpdateNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155UpdateNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155UpdateName represents a UpdateName event raised by the ERC1155 contract.
type ERC1155UpdateName struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateName is a free log retrieval operation binding the contract event 0x28bce0e23786df7a86b305fe801506dbf59150e2f634d23d4b6d702f99e60b87.
//
// Solidity: event UpdateName(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) FilterUpdateName(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155UpdateNameIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "UpdateName", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155UpdateNameIterator{contract: _ERC1155.contract, event: "UpdateName", logs: logs, sub: sub}, nil
}

// WatchUpdateName is a free log subscription operation binding the contract event 0x28bce0e23786df7a86b305fe801506dbf59150e2f634d23d4b6d702f99e60b87.
//
// Solidity: event UpdateName(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) WatchUpdateName(opts *bind.WatchOpts, sink chan<- *ERC1155UpdateName, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "UpdateName", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155UpdateName)
				if err := _ERC1155.contract.UnpackLog(event, "UpdateName", log); err != nil {
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

// ParseUpdateName is a log parse operation binding the contract event 0x28bce0e23786df7a86b305fe801506dbf59150e2f634d23d4b6d702f99e60b87.
//
// Solidity: event UpdateName(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) ParseUpdateName(log types.Log) (*ERC1155UpdateName, error) {
	event := new(ERC1155UpdateName)
	if err := _ERC1155.contract.UnpackLog(event, "UpdateName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155UpdateSymbolIterator is returned from FilterUpdateSymbol and is used to iterate over the raw logs and unpacked data for UpdateSymbol events raised by the ERC1155 contract.
type ERC1155UpdateSymbolIterator struct {
	Event *ERC1155UpdateSymbol // Event containing the contract specifics and raw log

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
func (it *ERC1155UpdateSymbolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155UpdateSymbol)
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
		it.Event = new(ERC1155UpdateSymbol)
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
func (it *ERC1155UpdateSymbolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155UpdateSymbolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155UpdateSymbol represents a UpdateSymbol event raised by the ERC1155 contract.
type ERC1155UpdateSymbol struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateSymbol is a free log retrieval operation binding the contract event 0x1d62d100a8eead9f37a04cf683c80f4f788305d58ef0e644569e89216f055be7.
//
// Solidity: event UpdateSymbol(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) FilterUpdateSymbol(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155UpdateSymbolIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "UpdateSymbol", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155UpdateSymbolIterator{contract: _ERC1155.contract, event: "UpdateSymbol", logs: logs, sub: sub}, nil
}

// WatchUpdateSymbol is a free log subscription operation binding the contract event 0x1d62d100a8eead9f37a04cf683c80f4f788305d58ef0e644569e89216f055be7.
//
// Solidity: event UpdateSymbol(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) WatchUpdateSymbol(opts *bind.WatchOpts, sink chan<- *ERC1155UpdateSymbol, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "UpdateSymbol", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155UpdateSymbol)
				if err := _ERC1155.contract.UnpackLog(event, "UpdateSymbol", log); err != nil {
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

// ParseUpdateSymbol is a log parse operation binding the contract event 0x1d62d100a8eead9f37a04cf683c80f4f788305d58ef0e644569e89216f055be7.
//
// Solidity: event UpdateSymbol(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) ParseUpdateSymbol(log types.Log) (*ERC1155UpdateSymbol, error) {
	event := new(ERC1155UpdateSymbol)
	if err := _ERC1155.contract.UnpackLog(event, "UpdateSymbol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155UpdateTransferFeeIterator is returned from FilterUpdateTransferFee and is used to iterate over the raw logs and unpacked data for UpdateTransferFee events raised by the ERC1155 contract.
type ERC1155UpdateTransferFeeIterator struct {
	Event *ERC1155UpdateTransferFee // Event containing the contract specifics and raw log

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
func (it *ERC1155UpdateTransferFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155UpdateTransferFee)
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
		it.Event = new(ERC1155UpdateTransferFee)
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
func (it *ERC1155UpdateTransferFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155UpdateTransferFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155UpdateTransferFee represents a UpdateTransferFee event raised by the ERC1155 contract.
type ERC1155UpdateTransferFee struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateTransferFee is a free log retrieval operation binding the contract event 0x6fbeac6c79c640ca4f5f47271bd7a36c7eb83076224c65a3b3378c8844720343.
//
// Solidity: event UpdateTransferFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) FilterUpdateTransferFee(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155UpdateTransferFeeIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "UpdateTransferFee", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155UpdateTransferFeeIterator{contract: _ERC1155.contract, event: "UpdateTransferFee", logs: logs, sub: sub}, nil
}

// WatchUpdateTransferFee is a free log subscription operation binding the contract event 0x6fbeac6c79c640ca4f5f47271bd7a36c7eb83076224c65a3b3378c8844720343.
//
// Solidity: event UpdateTransferFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) WatchUpdateTransferFee(opts *bind.WatchOpts, sink chan<- *ERC1155UpdateTransferFee, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "UpdateTransferFee", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155UpdateTransferFee)
				if err := _ERC1155.contract.UnpackLog(event, "UpdateTransferFee", log); err != nil {
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

// ParseUpdateTransferFee is a log parse operation binding the contract event 0x6fbeac6c79c640ca4f5f47271bd7a36c7eb83076224c65a3b3378c8844720343.
//
// Solidity: event UpdateTransferFee(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) ParseUpdateTransferFee(log types.Log) (*ERC1155UpdateTransferFee, error) {
	event := new(ERC1155UpdateTransferFee)
	if err := _ERC1155.contract.UnpackLog(event, "UpdateTransferFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155UpdateTransferableIterator is returned from FilterUpdateTransferable and is used to iterate over the raw logs and unpacked data for UpdateTransferable events raised by the ERC1155 contract.
type ERC1155UpdateTransferableIterator struct {
	Event *ERC1155UpdateTransferable // Event containing the contract specifics and raw log

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
func (it *ERC1155UpdateTransferableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155UpdateTransferable)
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
		it.Event = new(ERC1155UpdateTransferable)
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
func (it *ERC1155UpdateTransferableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155UpdateTransferableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155UpdateTransferable represents a UpdateTransferable event raised by the ERC1155 contract.
type ERC1155UpdateTransferable struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateTransferable is a free log retrieval operation binding the contract event 0x32f37a7bc2ffee1189c3c38246fca3004a880fedc7c732a79a1bc2bb27faa3db.
//
// Solidity: event UpdateTransferable(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) FilterUpdateTransferable(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155UpdateTransferableIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "UpdateTransferable", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155UpdateTransferableIterator{contract: _ERC1155.contract, event: "UpdateTransferable", logs: logs, sub: sub}, nil
}

// WatchUpdateTransferable is a free log subscription operation binding the contract event 0x32f37a7bc2ffee1189c3c38246fca3004a880fedc7c732a79a1bc2bb27faa3db.
//
// Solidity: event UpdateTransferable(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) WatchUpdateTransferable(opts *bind.WatchOpts, sink chan<- *ERC1155UpdateTransferable, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "UpdateTransferable", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155UpdateTransferable)
				if err := _ERC1155.contract.UnpackLog(event, "UpdateTransferable", log); err != nil {
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

// ParseUpdateTransferable is a log parse operation binding the contract event 0x32f37a7bc2ffee1189c3c38246fca3004a880fedc7c732a79a1bc2bb27faa3db.
//
// Solidity: event UpdateTransferable(uint256 indexed _id)
func (_ERC1155 *ERC1155Filterer) ParseUpdateTransferable(log types.Log) (*ERC1155UpdateTransferable, error) {
	event := new(ERC1155UpdateTransferable)
	if err := _ERC1155.contract.UnpackLog(event, "UpdateTransferable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155WhitelistIterator is returned from FilterWhitelist and is used to iterate over the raw logs and unpacked data for Whitelist events raised by the ERC1155 contract.
type ERC1155WhitelistIterator struct {
	Event *ERC1155Whitelist // Event containing the contract specifics and raw log

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
func (it *ERC1155WhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Whitelist)
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
		it.Event = new(ERC1155Whitelist)
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
func (it *ERC1155WhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155WhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Whitelist represents a Whitelist event raised by the ERC1155 contract.
type ERC1155Whitelist struct {
	Id          *big.Int
	Account     common.Address
	Whitelisted common.Address
	On          bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWhitelist is a free log retrieval operation binding the contract event 0xc4d4c7ee56ebaa174aef190c7829f5e548c91dea7fd3304e23ff1d49a8f35b9b.
//
// Solidity: event Whitelist(uint256 indexed _id, address indexed _account, address _whitelisted, bool _on)
func (_ERC1155 *ERC1155Filterer) FilterWhitelist(opts *bind.FilterOpts, _id []*big.Int, _account []common.Address) (*ERC1155WhitelistIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Whitelist", _idRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155WhitelistIterator{contract: _ERC1155.contract, event: "Whitelist", logs: logs, sub: sub}, nil
}

// WatchWhitelist is a free log subscription operation binding the contract event 0xc4d4c7ee56ebaa174aef190c7829f5e548c91dea7fd3304e23ff1d49a8f35b9b.
//
// Solidity: event Whitelist(uint256 indexed _id, address indexed _account, address _whitelisted, bool _on)
func (_ERC1155 *ERC1155Filterer) WatchWhitelist(opts *bind.WatchOpts, sink chan<- *ERC1155Whitelist, _id []*big.Int, _account []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Whitelist", _idRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Whitelist)
				if err := _ERC1155.contract.UnpackLog(event, "Whitelist", log); err != nil {
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

// ParseWhitelist is a log parse operation binding the contract event 0xc4d4c7ee56ebaa174aef190c7829f5e548c91dea7fd3304e23ff1d49a8f35b9b.
//
// Solidity: event Whitelist(uint256 indexed _id, address indexed _account, address _whitelisted, bool _on)
func (_ERC1155 *ERC1155Filterer) ParseWhitelist(log types.Log) (*ERC1155Whitelist, error) {
	event := new(ERC1155Whitelist)
	if err := _ERC1155.contract.UnpackLog(event, "Whitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
