// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blur

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

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Sell Input
	Buy  Input
}

// Fee is an auto generated low-level Go binding around an user-defined struct.
type Fee struct {
	Rate      uint16
	Recipient common.Address
}

// Input is an auto generated low-level Go binding around an user-defined struct.
type Input struct {
	Order Order
	V     uint8
	R                [32]byte
	S                [32]byte
	ExtraSignature   []byte
	SignatureVersion uint8
	BlockNumber      *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader         common.Address
	Side           uint8
	MatchingPolicy common.Address
	Collection     common.Address
	TokenId        *big.Int
	Amount         *big.Int
	PaymentToken   common.Address
	Price          *big.Int
	ListingTime    *big.Int
	ExpirationTime *big.Int
	Fees           []Fee
	Salt           *big.Int
	ExtraParams    []byte
}

// BlurMetaData contains all meta data concerning the Blur contract.
var BlurMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockRange\",\"type\":\"uint256\"}],\"name\":\"NewBlockRange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIExecutionDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"}],\"name\":\"NewExecutionDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"name\":\"NewFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"NewFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"NewGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPolicyManager\",\"name\":\"policyManager\",\"type\":\"address\"}],\"name\":\"NewPolicyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Opened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"sell\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"buy\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"buyHash\",\"type\":\"bytes32\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INVERSE_BASIS_POINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"_execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"name\":\"bulkExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelledOrFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionDelegate\",\"outputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"},{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInternal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"policyManager\",\"outputs\":[{\"internalType\":\"contractIPolicyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"setBlockRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"}],\"name\":\"setExecutionDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"}],\"name\":\"setPolicyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BlurABI is the input ABI used to generate the binding from.
// Deprecated: Use BlurMetaData.ABI instead.
var BlurABI = BlurMetaData.ABI

// Blur is an auto generated Go binding around an Ethereum contract.
type Blur struct {
	BlurCaller     // Read-only binding to the contract
	BlurTransactor // Write-only binding to the contract
	BlurFilterer   // Log filterer for contract events
}

// BlurCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlurCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlurTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlurFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlurSession struct {
	Contract     *Blur             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlurCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlurCallerSession struct {
	Contract *BlurCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BlurTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlurTransactorSession struct {
	Contract     *BlurTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlurRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlurRaw struct {
	Contract *Blur // Generic contract binding to access the raw methods on
}

// BlurCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlurCallerRaw struct {
	Contract *BlurCaller // Generic read-only contract binding to access the raw methods on
}

// BlurTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlurTransactorRaw struct {
	Contract *BlurTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlur creates a new instance of Blur, bound to a specific deployed contract.
func NewBlur(address common.Address, backend bind.ContractBackend) (*Blur, error) {
	contract, err := bindBlur(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blur{BlurCaller: BlurCaller{contract: contract}, BlurTransactor: BlurTransactor{contract: contract}, BlurFilterer: BlurFilterer{contract: contract}}, nil
}

// NewBlurCaller creates a new read-only instance of Blur, bound to a specific deployed contract.
func NewBlurCaller(address common.Address, caller bind.ContractCaller) (*BlurCaller, error) {
	contract, err := bindBlur(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlurCaller{contract: contract}, nil
}

// NewBlurTransactor creates a new write-only instance of Blur, bound to a specific deployed contract.
func NewBlurTransactor(address common.Address, transactor bind.ContractTransactor) (*BlurTransactor, error) {
	contract, err := bindBlur(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlurTransactor{contract: contract}, nil
}

// NewBlurFilterer creates a new log filterer instance of Blur, bound to a specific deployed contract.
func NewBlurFilterer(address common.Address, filterer bind.ContractFilterer) (*BlurFilterer, error) {
	contract, err := bindBlur(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlurFilterer{contract: contract}, nil
}

// bindBlur binds a generic wrapper to an already deployed contract.
func bindBlur(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlurABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blur *BlurRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blur.Contract.BlurCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blur *BlurRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blur.Contract.BlurTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blur *BlurRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blur.Contract.BlurTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blur *BlurCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blur.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blur *BlurTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blur.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blur *BlurTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blur.Contract.contract.Transact(opts, method, params...)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_Blur *BlurCaller) FEETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "FEE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_Blur *BlurSession) FEETYPEHASH() ([32]byte, error) {
	return _Blur.Contract.FEETYPEHASH(&_Blur.CallOpts)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_Blur *BlurCallerSession) FEETYPEHASH() ([32]byte, error) {
	return _Blur.Contract.FEETYPEHASH(&_Blur.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Blur *BlurCaller) INVERSEBASISPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "INVERSE_BASIS_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Blur *BlurSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Blur.Contract.INVERSEBASISPOINT(&_Blur.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Blur *BlurCallerSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Blur.Contract.INVERSEBASISPOINT(&_Blur.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Blur *BlurCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Blur *BlurSession) NAME() (string, error) {
	return _Blur.Contract.NAME(&_Blur.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Blur *BlurCallerSession) NAME() (string, error) {
	return _Blur.Contract.NAME(&_Blur.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_Blur *BlurCaller) ORACLEORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "ORACLE_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_Blur *BlurSession) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _Blur.Contract.ORACLEORDERTYPEHASH(&_Blur.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_Blur *BlurCallerSession) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _Blur.Contract.ORACLEORDERTYPEHASH(&_Blur.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Blur *BlurCaller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Blur *BlurSession) ORDERTYPEHASH() ([32]byte, error) {
	return _Blur.Contract.ORDERTYPEHASH(&_Blur.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Blur *BlurCallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _Blur.Contract.ORDERTYPEHASH(&_Blur.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Blur *BlurCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Blur *BlurSession) POOL() (common.Address, error) {
	return _Blur.Contract.POOL(&_Blur.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Blur *BlurCallerSession) POOL() (common.Address, error) {
	return _Blur.Contract.POOL(&_Blur.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_Blur *BlurCaller) ROOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "ROOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_Blur *BlurSession) ROOTTYPEHASH() ([32]byte, error) {
	return _Blur.Contract.ROOTTYPEHASH(&_Blur.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_Blur *BlurCallerSession) ROOTTYPEHASH() ([32]byte, error) {
	return _Blur.Contract.ROOTTYPEHASH(&_Blur.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Blur *BlurCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Blur *BlurSession) VERSION() (string, error) {
	return _Blur.Contract.VERSION(&_Blur.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Blur *BlurCallerSession) VERSION() (string, error) {
	return _Blur.Contract.VERSION(&_Blur.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Blur *BlurCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Blur *BlurSession) WETH() (common.Address, error) {
	return _Blur.Contract.WETH(&_Blur.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Blur *BlurCallerSession) WETH() (common.Address, error) {
	return _Blur.Contract.WETH(&_Blur.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_Blur *BlurCaller) BlockRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "blockRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_Blur *BlurSession) BlockRange() (*big.Int, error) {
	return _Blur.Contract.BlockRange(&_Blur.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_Blur *BlurCallerSession) BlockRange() (*big.Int, error) {
	return _Blur.Contract.BlockRange(&_Blur.CallOpts)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_Blur *BlurCaller) CancelledOrFilled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "cancelledOrFilled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_Blur *BlurSession) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _Blur.Contract.CancelledOrFilled(&_Blur.CallOpts, arg0)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_Blur *BlurCallerSession) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _Blur.Contract.CancelledOrFilled(&_Blur.CallOpts, arg0)
}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_Blur *BlurCaller) ExecutionDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "executionDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_Blur *BlurSession) ExecutionDelegate() (common.Address, error) {
	return _Blur.Contract.ExecutionDelegate(&_Blur.CallOpts)
}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_Blur *BlurCallerSession) ExecutionDelegate() (common.Address, error) {
	return _Blur.Contract.ExecutionDelegate(&_Blur.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Blur *BlurCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Blur *BlurSession) FeeRate() (*big.Int, error) {
	return _Blur.Contract.FeeRate(&_Blur.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Blur *BlurCallerSession) FeeRate() (*big.Int, error) {
	return _Blur.Contract.FeeRate(&_Blur.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Blur *BlurCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Blur *BlurSession) FeeRecipient() (common.Address, error) {
	return _Blur.Contract.FeeRecipient(&_Blur.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Blur *BlurCallerSession) FeeRecipient() (common.Address, error) {
	return _Blur.Contract.FeeRecipient(&_Blur.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Blur *BlurCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Blur *BlurSession) Governor() (common.Address, error) {
	return _Blur.Contract.Governor(&_Blur.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Blur *BlurCallerSession) Governor() (common.Address, error) {
	return _Blur.Contract.Governor(&_Blur.CallOpts)
}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_Blur *BlurCaller) IsInternal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "isInternal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_Blur *BlurSession) IsInternal() (bool, error) {
	return _Blur.Contract.IsInternal(&_Blur.CallOpts)
}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_Blur *BlurCallerSession) IsInternal() (bool, error) {
	return _Blur.Contract.IsInternal(&_Blur.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_Blur *BlurCaller) IsOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_Blur *BlurSession) IsOpen() (*big.Int, error) {
	return _Blur.Contract.IsOpen(&_Blur.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_Blur *BlurCallerSession) IsOpen() (*big.Int, error) {
	return _Blur.Contract.IsOpen(&_Blur.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Blur *BlurCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Blur *BlurSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Blur.Contract.Nonces(&_Blur.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Blur *BlurCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Blur.Contract.Nonces(&_Blur.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Blur *BlurCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Blur *BlurSession) Oracle() (common.Address, error) {
	return _Blur.Contract.Oracle(&_Blur.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Blur *BlurCallerSession) Oracle() (common.Address, error) {
	return _Blur.Contract.Oracle(&_Blur.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blur *BlurCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blur *BlurSession) Owner() (common.Address, error) {
	return _Blur.Contract.Owner(&_Blur.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blur *BlurCallerSession) Owner() (common.Address, error) {
	return _Blur.Contract.Owner(&_Blur.CallOpts)
}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_Blur *BlurCaller) PolicyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "policyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_Blur *BlurSession) PolicyManager() (common.Address, error) {
	return _Blur.Contract.PolicyManager(&_Blur.CallOpts)
}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_Blur *BlurCallerSession) PolicyManager() (common.Address, error) {
	return _Blur.Contract.PolicyManager(&_Blur.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Blur *BlurCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Blur *BlurSession) ProxiableUUID() ([32]byte, error) {
	return _Blur.Contract.ProxiableUUID(&_Blur.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Blur *BlurCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Blur.Contract.ProxiableUUID(&_Blur.CallOpts)
}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_Blur *BlurCaller) RemainingETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "remainingETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_Blur *BlurSession) RemainingETH() (*big.Int, error) {
	return _Blur.Contract.RemainingETH(&_Blur.CallOpts)
}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_Blur *BlurCallerSession) RemainingETH() (*big.Int, error) {
	return _Blur.Contract.RemainingETH(&_Blur.CallOpts)
}

// InternalExecute is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blur *BlurTransactor) InternalExecute(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "_execute", sell, buy)
}

// InternalExecute is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blur *BlurSession) InternalExecute(sell Input, buy Input) (*types.Transaction, error) {
	return _Blur.Contract.InternalExecute(&_Blur.TransactOpts, sell, buy)
}

// InternalExecute is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blur *BlurTransactorSession) InternalExecute(sell Input, buy Input) (*types.Transaction, error) {
	return _Blur.Contract.InternalExecute(&_Blur.TransactOpts, sell, buy)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_Blur *BlurTransactor) BulkExecute(opts *bind.TransactOpts, executions []Execution) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "bulkExecute", executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_Blur *BlurSession) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _Blur.Contract.BulkExecute(&_Blur.TransactOpts, executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_Blur *BlurTransactorSession) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _Blur.Contract.BulkExecute(&_Blur.TransactOpts, executions)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_Blur *BlurTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_Blur *BlurSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _Blur.Contract.CancelOrder(&_Blur.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_Blur *BlurTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _Blur.Contract.CancelOrder(&_Blur.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_Blur *BlurTransactor) CancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_Blur *BlurSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _Blur.Contract.CancelOrders(&_Blur.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_Blur *BlurTransactorSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _Blur.Contract.CancelOrders(&_Blur.TransactOpts, orders)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Blur *BlurTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Blur *BlurSession) Close() (*types.Transaction, error) {
	return _Blur.Contract.Close(&_Blur.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Blur *BlurTransactorSession) Close() (*types.Transaction, error) {
	return _Blur.Contract.Close(&_Blur.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blur *BlurTransactor) Execute(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "execute", sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blur *BlurSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _Blur.Contract.Execute(&_Blur.TransactOpts, sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blur *BlurTransactorSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _Blur.Contract.Execute(&_Blur.TransactOpts, sell, buy)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Blur *BlurTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Blur *BlurSession) IncrementNonce() (*types.Transaction, error) {
	return _Blur.Contract.IncrementNonce(&_Blur.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Blur *BlurTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _Blur.Contract.IncrementNonce(&_Blur.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_Blur *BlurTransactor) Initialize(opts *bind.TransactOpts, _executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "initialize", _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_Blur *BlurSession) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _Blur.Contract.Initialize(&_Blur.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_Blur *BlurTransactorSession) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _Blur.Contract.Initialize(&_Blur.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Blur *BlurTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Blur *BlurSession) Open() (*types.Transaction, error) {
	return _Blur.Contract.Open(&_Blur.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Blur *BlurTransactorSession) Open() (*types.Transaction, error) {
	return _Blur.Contract.Open(&_Blur.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blur *BlurTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blur *BlurSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blur.Contract.RenounceOwnership(&_Blur.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blur *BlurTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blur.Contract.RenounceOwnership(&_Blur.TransactOpts)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_Blur *BlurTransactor) SetBlockRange(opts *bind.TransactOpts, _blockRange *big.Int) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "setBlockRange", _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_Blur *BlurSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _Blur.Contract.SetBlockRange(&_Blur.TransactOpts, _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_Blur *BlurTransactorSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _Blur.Contract.SetBlockRange(&_Blur.TransactOpts, _blockRange)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_Blur *BlurTransactor) SetExecutionDelegate(opts *bind.TransactOpts, _executionDelegate common.Address) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "setExecutionDelegate", _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_Blur *BlurSession) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetExecutionDelegate(&_Blur.TransactOpts, _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_Blur *BlurTransactorSession) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetExecutionDelegate(&_Blur.TransactOpts, _executionDelegate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_Blur *BlurTransactor) SetFeeRate(opts *bind.TransactOpts, _feeRate *big.Int) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "setFeeRate", _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_Blur *BlurSession) SetFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _Blur.Contract.SetFeeRate(&_Blur.TransactOpts, _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_Blur *BlurTransactorSession) SetFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _Blur.Contract.SetFeeRate(&_Blur.TransactOpts, _feeRate)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Blur *BlurTransactor) SetFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "setFeeRecipient", _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Blur *BlurSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetFeeRecipient(&_Blur.TransactOpts, _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Blur *BlurTransactorSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetFeeRecipient(&_Blur.TransactOpts, _feeRecipient)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_Blur *BlurTransactor) SetGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "setGovernor", _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_Blur *BlurSession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetGovernor(&_Blur.TransactOpts, _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_Blur *BlurTransactorSession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetGovernor(&_Blur.TransactOpts, _governor)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Blur *BlurTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Blur *BlurSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetOracle(&_Blur.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Blur *BlurTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetOracle(&_Blur.TransactOpts, _oracle)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_Blur *BlurTransactor) SetPolicyManager(opts *bind.TransactOpts, _policyManager common.Address) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "setPolicyManager", _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_Blur *BlurSession) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetPolicyManager(&_Blur.TransactOpts, _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_Blur *BlurTransactorSession) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _Blur.Contract.SetPolicyManager(&_Blur.TransactOpts, _policyManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blur *BlurTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blur *BlurSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blur.Contract.TransferOwnership(&_Blur.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blur *BlurTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blur.Contract.TransferOwnership(&_Blur.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Blur *BlurTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Blur *BlurSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Blur.Contract.UpgradeTo(&_Blur.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Blur *BlurTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Blur.Contract.UpgradeTo(&_Blur.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Blur *BlurTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Blur *BlurSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Blur.Contract.UpgradeToAndCall(&_Blur.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Blur *BlurTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Blur.Contract.UpgradeToAndCall(&_Blur.TransactOpts, newImplementation, data)
}

// BlurAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Blur contract.
type BlurAdminChangedIterator struct {
	Event *BlurAdminChanged // Event containing the contract specifics and raw log

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
func (it *BlurAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurAdminChanged)
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
		it.Event = new(BlurAdminChanged)
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
func (it *BlurAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurAdminChanged represents a AdminChanged event raised by the Blur contract.
type BlurAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Blur *BlurFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BlurAdminChangedIterator, error) {

	logs, sub, err := _Blur.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BlurAdminChangedIterator{contract: _Blur.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Blur *BlurFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BlurAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Blur.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurAdminChanged)
				if err := _Blur.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Blur *BlurFilterer) ParseAdminChanged(log types.Log) (*BlurAdminChanged, error) {
	event := new(BlurAdminChanged)
	if err := _Blur.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Blur contract.
type BlurBeaconUpgradedIterator struct {
	Event *BlurBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BlurBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurBeaconUpgraded)
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
		it.Event = new(BlurBeaconUpgraded)
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
func (it *BlurBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurBeaconUpgraded represents a BeaconUpgraded event raised by the Blur contract.
type BlurBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Blur *BlurFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BlurBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Blur.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BlurBeaconUpgradedIterator{contract: _Blur.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Blur *BlurFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BlurBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Blur.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurBeaconUpgraded)
				if err := _Blur.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Blur *BlurFilterer) ParseBeaconUpgraded(log types.Log) (*BlurBeaconUpgraded, error) {
	event := new(BlurBeaconUpgraded)
	if err := _Blur.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the Blur contract.
type BlurClosedIterator struct {
	Event *BlurClosed // Event containing the contract specifics and raw log

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
func (it *BlurClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurClosed)
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
		it.Event = new(BlurClosed)
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
func (it *BlurClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurClosed represents a Closed event raised by the Blur contract.
type BlurClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_Blur *BlurFilterer) FilterClosed(opts *bind.FilterOpts) (*BlurClosedIterator, error) {

	logs, sub, err := _Blur.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &BlurClosedIterator{contract: _Blur.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_Blur *BlurFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *BlurClosed) (event.Subscription, error) {

	logs, sub, err := _Blur.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurClosed)
				if err := _Blur.contract.UnpackLog(event, "Closed", log); err != nil {
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

// ParseClosed is a log parse operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_Blur *BlurFilterer) ParseClosed(log types.Log) (*BlurClosed, error) {
	event := new(BlurClosed)
	if err := _Blur.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Blur contract.
type BlurInitializedIterator struct {
	Event *BlurInitialized // Event containing the contract specifics and raw log

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
func (it *BlurInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurInitialized)
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
		it.Event = new(BlurInitialized)
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
func (it *BlurInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurInitialized represents a Initialized event raised by the Blur contract.
type BlurInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Blur *BlurFilterer) FilterInitialized(opts *bind.FilterOpts) (*BlurInitializedIterator, error) {

	logs, sub, err := _Blur.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BlurInitializedIterator{contract: _Blur.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Blur *BlurFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BlurInitialized) (event.Subscription, error) {

	logs, sub, err := _Blur.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurInitialized)
				if err := _Blur.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Blur *BlurFilterer) ParseInitialized(log types.Log) (*BlurInitialized, error) {
	event := new(BlurInitialized)
	if err := _Blur.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurNewBlockRangeIterator is returned from FilterNewBlockRange and is used to iterate over the raw logs and unpacked data for NewBlockRange events raised by the Blur contract.
type BlurNewBlockRangeIterator struct {
	Event *BlurNewBlockRange // Event containing the contract specifics and raw log

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
func (it *BlurNewBlockRangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurNewBlockRange)
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
		it.Event = new(BlurNewBlockRange)
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
func (it *BlurNewBlockRangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurNewBlockRangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurNewBlockRange represents a NewBlockRange event raised by the Blur contract.
type BlurNewBlockRange struct {
	BlockRange *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewBlockRange is a free log retrieval operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_Blur *BlurFilterer) FilterNewBlockRange(opts *bind.FilterOpts) (*BlurNewBlockRangeIterator, error) {

	logs, sub, err := _Blur.contract.FilterLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return &BlurNewBlockRangeIterator{contract: _Blur.contract, event: "NewBlockRange", logs: logs, sub: sub}, nil
}

// WatchNewBlockRange is a free log subscription operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_Blur *BlurFilterer) WatchNewBlockRange(opts *bind.WatchOpts, sink chan<- *BlurNewBlockRange) (event.Subscription, error) {

	logs, sub, err := _Blur.contract.WatchLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurNewBlockRange)
				if err := _Blur.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
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

// ParseNewBlockRange is a log parse operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_Blur *BlurFilterer) ParseNewBlockRange(log types.Log) (*BlurNewBlockRange, error) {
	event := new(BlurNewBlockRange)
	if err := _Blur.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurNewExecutionDelegateIterator is returned from FilterNewExecutionDelegate and is used to iterate over the raw logs and unpacked data for NewExecutionDelegate events raised by the Blur contract.
type BlurNewExecutionDelegateIterator struct {
	Event *BlurNewExecutionDelegate // Event containing the contract specifics and raw log

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
func (it *BlurNewExecutionDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurNewExecutionDelegate)
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
		it.Event = new(BlurNewExecutionDelegate)
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
func (it *BlurNewExecutionDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurNewExecutionDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurNewExecutionDelegate represents a NewExecutionDelegate event raised by the Blur contract.
type BlurNewExecutionDelegate struct {
	ExecutionDelegate common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionDelegate is a free log retrieval operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_Blur *BlurFilterer) FilterNewExecutionDelegate(opts *bind.FilterOpts, executionDelegate []common.Address) (*BlurNewExecutionDelegateIterator, error) {

	var executionDelegateRule []interface{}
	for _, executionDelegateItem := range executionDelegate {
		executionDelegateRule = append(executionDelegateRule, executionDelegateItem)
	}

	logs, sub, err := _Blur.contract.FilterLogs(opts, "NewExecutionDelegate", executionDelegateRule)
	if err != nil {
		return nil, err
	}
	return &BlurNewExecutionDelegateIterator{contract: _Blur.contract, event: "NewExecutionDelegate", logs: logs, sub: sub}, nil
}

// WatchNewExecutionDelegate is a free log subscription operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_Blur *BlurFilterer) WatchNewExecutionDelegate(opts *bind.WatchOpts, sink chan<- *BlurNewExecutionDelegate, executionDelegate []common.Address) (event.Subscription, error) {

	var executionDelegateRule []interface{}
	for _, executionDelegateItem := range executionDelegate {
		executionDelegateRule = append(executionDelegateRule, executionDelegateItem)
	}

	logs, sub, err := _Blur.contract.WatchLogs(opts, "NewExecutionDelegate", executionDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurNewExecutionDelegate)
				if err := _Blur.contract.UnpackLog(event, "NewExecutionDelegate", log); err != nil {
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

// ParseNewExecutionDelegate is a log parse operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_Blur *BlurFilterer) ParseNewExecutionDelegate(log types.Log) (*BlurNewExecutionDelegate, error) {
	event := new(BlurNewExecutionDelegate)
	if err := _Blur.contract.UnpackLog(event, "NewExecutionDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurNewFeeRateIterator is returned from FilterNewFeeRate and is used to iterate over the raw logs and unpacked data for NewFeeRate events raised by the Blur contract.
type BlurNewFeeRateIterator struct {
	Event *BlurNewFeeRate // Event containing the contract specifics and raw log

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
func (it *BlurNewFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurNewFeeRate)
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
		it.Event = new(BlurNewFeeRate)
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
func (it *BlurNewFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurNewFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurNewFeeRate represents a NewFeeRate event raised by the Blur contract.
type BlurNewFeeRate struct {
	FeeRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFeeRate is a free log retrieval operation binding the contract event 0x788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae.
//
// Solidity: event NewFeeRate(uint256 feeRate)
func (_Blur *BlurFilterer) FilterNewFeeRate(opts *bind.FilterOpts) (*BlurNewFeeRateIterator, error) {

	logs, sub, err := _Blur.contract.FilterLogs(opts, "NewFeeRate")
	if err != nil {
		return nil, err
	}
	return &BlurNewFeeRateIterator{contract: _Blur.contract, event: "NewFeeRate", logs: logs, sub: sub}, nil
}

// WatchNewFeeRate is a free log subscription operation binding the contract event 0x788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae.
//
// Solidity: event NewFeeRate(uint256 feeRate)
func (_Blur *BlurFilterer) WatchNewFeeRate(opts *bind.WatchOpts, sink chan<- *BlurNewFeeRate) (event.Subscription, error) {

	logs, sub, err := _Blur.contract.WatchLogs(opts, "NewFeeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurNewFeeRate)
				if err := _Blur.contract.UnpackLog(event, "NewFeeRate", log); err != nil {
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

// ParseNewFeeRate is a log parse operation binding the contract event 0x788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae.
//
// Solidity: event NewFeeRate(uint256 feeRate)
func (_Blur *BlurFilterer) ParseNewFeeRate(log types.Log) (*BlurNewFeeRate, error) {
	event := new(BlurNewFeeRate)
	if err := _Blur.contract.UnpackLog(event, "NewFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurNewFeeRecipientIterator is returned from FilterNewFeeRecipient and is used to iterate over the raw logs and unpacked data for NewFeeRecipient events raised by the Blur contract.
type BlurNewFeeRecipientIterator struct {
	Event *BlurNewFeeRecipient // Event containing the contract specifics and raw log

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
func (it *BlurNewFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurNewFeeRecipient)
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
		it.Event = new(BlurNewFeeRecipient)
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
func (it *BlurNewFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurNewFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurNewFeeRecipient represents a NewFeeRecipient event raised by the Blur contract.
type BlurNewFeeRecipient struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewFeeRecipient is a free log retrieval operation binding the contract event 0x412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23.
//
// Solidity: event NewFeeRecipient(address feeRecipient)
func (_Blur *BlurFilterer) FilterNewFeeRecipient(opts *bind.FilterOpts) (*BlurNewFeeRecipientIterator, error) {

	logs, sub, err := _Blur.contract.FilterLogs(opts, "NewFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &BlurNewFeeRecipientIterator{contract: _Blur.contract, event: "NewFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewFeeRecipient is a free log subscription operation binding the contract event 0x412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23.
//
// Solidity: event NewFeeRecipient(address feeRecipient)
func (_Blur *BlurFilterer) WatchNewFeeRecipient(opts *bind.WatchOpts, sink chan<- *BlurNewFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _Blur.contract.WatchLogs(opts, "NewFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurNewFeeRecipient)
				if err := _Blur.contract.UnpackLog(event, "NewFeeRecipient", log); err != nil {
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

// ParseNewFeeRecipient is a log parse operation binding the contract event 0x412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23.
//
// Solidity: event NewFeeRecipient(address feeRecipient)
func (_Blur *BlurFilterer) ParseNewFeeRecipient(log types.Log) (*BlurNewFeeRecipient, error) {
	event := new(BlurNewFeeRecipient)
	if err := _Blur.contract.UnpackLog(event, "NewFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurNewGovernorIterator is returned from FilterNewGovernor and is used to iterate over the raw logs and unpacked data for NewGovernor events raised by the Blur contract.
type BlurNewGovernorIterator struct {
	Event *BlurNewGovernor // Event containing the contract specifics and raw log

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
func (it *BlurNewGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurNewGovernor)
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
		it.Event = new(BlurNewGovernor)
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
func (it *BlurNewGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurNewGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurNewGovernor represents a NewGovernor event raised by the Blur contract.
type BlurNewGovernor struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewGovernor is a free log retrieval operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address governor)
func (_Blur *BlurFilterer) FilterNewGovernor(opts *bind.FilterOpts) (*BlurNewGovernorIterator, error) {

	logs, sub, err := _Blur.contract.FilterLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return &BlurNewGovernorIterator{contract: _Blur.contract, event: "NewGovernor", logs: logs, sub: sub}, nil
}

// WatchNewGovernor is a free log subscription operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address governor)
func (_Blur *BlurFilterer) WatchNewGovernor(opts *bind.WatchOpts, sink chan<- *BlurNewGovernor) (event.Subscription, error) {

	logs, sub, err := _Blur.contract.WatchLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurNewGovernor)
				if err := _Blur.contract.UnpackLog(event, "NewGovernor", log); err != nil {
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

// ParseNewGovernor is a log parse operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address governor)
func (_Blur *BlurFilterer) ParseNewGovernor(log types.Log) (*BlurNewGovernor, error) {
	event := new(BlurNewGovernor)
	if err := _Blur.contract.UnpackLog(event, "NewGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurNewOracleIterator is returned from FilterNewOracle and is used to iterate over the raw logs and unpacked data for NewOracle events raised by the Blur contract.
type BlurNewOracleIterator struct {
	Event *BlurNewOracle // Event containing the contract specifics and raw log

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
func (it *BlurNewOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurNewOracle)
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
		it.Event = new(BlurNewOracle)
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
func (it *BlurNewOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurNewOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurNewOracle represents a NewOracle event raised by the Blur contract.
type BlurNewOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewOracle is a free log retrieval operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_Blur *BlurFilterer) FilterNewOracle(opts *bind.FilterOpts, oracle []common.Address) (*BlurNewOracleIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Blur.contract.FilterLogs(opts, "NewOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return &BlurNewOracleIterator{contract: _Blur.contract, event: "NewOracle", logs: logs, sub: sub}, nil
}

// WatchNewOracle is a free log subscription operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_Blur *BlurFilterer) WatchNewOracle(opts *bind.WatchOpts, sink chan<- *BlurNewOracle, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Blur.contract.WatchLogs(opts, "NewOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurNewOracle)
				if err := _Blur.contract.UnpackLog(event, "NewOracle", log); err != nil {
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

// ParseNewOracle is a log parse operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_Blur *BlurFilterer) ParseNewOracle(log types.Log) (*BlurNewOracle, error) {
	event := new(BlurNewOracle)
	if err := _Blur.contract.UnpackLog(event, "NewOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurNewPolicyManagerIterator is returned from FilterNewPolicyManager and is used to iterate over the raw logs and unpacked data for NewPolicyManager events raised by the Blur contract.
type BlurNewPolicyManagerIterator struct {
	Event *BlurNewPolicyManager // Event containing the contract specifics and raw log

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
func (it *BlurNewPolicyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurNewPolicyManager)
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
		it.Event = new(BlurNewPolicyManager)
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
func (it *BlurNewPolicyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurNewPolicyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurNewPolicyManager represents a NewPolicyManager event raised by the Blur contract.
type BlurNewPolicyManager struct {
	PolicyManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewPolicyManager is a free log retrieval operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_Blur *BlurFilterer) FilterNewPolicyManager(opts *bind.FilterOpts, policyManager []common.Address) (*BlurNewPolicyManagerIterator, error) {

	var policyManagerRule []interface{}
	for _, policyManagerItem := range policyManager {
		policyManagerRule = append(policyManagerRule, policyManagerItem)
	}

	logs, sub, err := _Blur.contract.FilterLogs(opts, "NewPolicyManager", policyManagerRule)
	if err != nil {
		return nil, err
	}
	return &BlurNewPolicyManagerIterator{contract: _Blur.contract, event: "NewPolicyManager", logs: logs, sub: sub}, nil
}

// WatchNewPolicyManager is a free log subscription operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_Blur *BlurFilterer) WatchNewPolicyManager(opts *bind.WatchOpts, sink chan<- *BlurNewPolicyManager, policyManager []common.Address) (event.Subscription, error) {

	var policyManagerRule []interface{}
	for _, policyManagerItem := range policyManager {
		policyManagerRule = append(policyManagerRule, policyManagerItem)
	}

	logs, sub, err := _Blur.contract.WatchLogs(opts, "NewPolicyManager", policyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurNewPolicyManager)
				if err := _Blur.contract.UnpackLog(event, "NewPolicyManager", log); err != nil {
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

// ParseNewPolicyManager is a log parse operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_Blur *BlurFilterer) ParseNewPolicyManager(log types.Log) (*BlurNewPolicyManager, error) {
	event := new(BlurNewPolicyManager)
	if err := _Blur.contract.UnpackLog(event, "NewPolicyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurNonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the Blur contract.
type BlurNonceIncrementedIterator struct {
	Event *BlurNonceIncremented // Event containing the contract specifics and raw log

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
func (it *BlurNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurNonceIncremented)
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
		it.Event = new(BlurNonceIncremented)
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
func (it *BlurNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurNonceIncremented represents a NonceIncremented event raised by the Blur contract.
type BlurNonceIncremented struct {
	Trader   common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_Blur *BlurFilterer) FilterNonceIncremented(opts *bind.FilterOpts, trader []common.Address) (*BlurNonceIncrementedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Blur.contract.FilterLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return &BlurNonceIncrementedIterator{contract: _Blur.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_Blur *BlurFilterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *BlurNonceIncremented, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Blur.contract.WatchLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurNonceIncremented)
				if err := _Blur.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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

// ParseNonceIncremented is a log parse operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_Blur *BlurFilterer) ParseNonceIncremented(log types.Log) (*BlurNonceIncremented, error) {
	event := new(BlurNonceIncremented)
	if err := _Blur.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurOpenedIterator is returned from FilterOpened and is used to iterate over the raw logs and unpacked data for Opened events raised by the Blur contract.
type BlurOpenedIterator struct {
	Event *BlurOpened // Event containing the contract specifics and raw log

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
func (it *BlurOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurOpened)
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
		it.Event = new(BlurOpened)
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
func (it *BlurOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurOpened represents a Opened event raised by the Blur contract.
type BlurOpened struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOpened is a free log retrieval operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_Blur *BlurFilterer) FilterOpened(opts *bind.FilterOpts) (*BlurOpenedIterator, error) {

	logs, sub, err := _Blur.contract.FilterLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return &BlurOpenedIterator{contract: _Blur.contract, event: "Opened", logs: logs, sub: sub}, nil
}

// WatchOpened is a free log subscription operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_Blur *BlurFilterer) WatchOpened(opts *bind.WatchOpts, sink chan<- *BlurOpened) (event.Subscription, error) {

	logs, sub, err := _Blur.contract.WatchLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurOpened)
				if err := _Blur.contract.UnpackLog(event, "Opened", log); err != nil {
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

// ParseOpened is a log parse operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_Blur *BlurFilterer) ParseOpened(log types.Log) (*BlurOpened, error) {
	event := new(BlurOpened)
	if err := _Blur.contract.UnpackLog(event, "Opened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Blur contract.
type BlurOrderCancelledIterator struct {
	Event *BlurOrderCancelled // Event containing the contract specifics and raw log

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
func (it *BlurOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurOrderCancelled)
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
		it.Event = new(BlurOrderCancelled)
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
func (it *BlurOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurOrderCancelled represents a OrderCancelled event raised by the Blur contract.
type BlurOrderCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_Blur *BlurFilterer) FilterOrderCancelled(opts *bind.FilterOpts) (*BlurOrderCancelledIterator, error) {

	logs, sub, err := _Blur.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &BlurOrderCancelledIterator{contract: _Blur.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_Blur *BlurFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *BlurOrderCancelled) (event.Subscription, error) {

	logs, sub, err := _Blur.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurOrderCancelled)
				if err := _Blur.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_Blur *BlurFilterer) ParseOrderCancelled(log types.Log) (*BlurOrderCancelled, error) {
	event := new(BlurOrderCancelled)
	if err := _Blur.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the Blur contract.
type BlurOrdersMatchedIterator struct {
	Event *BlurOrdersMatched // Event containing the contract specifics and raw log

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
func (it *BlurOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurOrdersMatched)
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
		it.Event = new(BlurOrdersMatched)
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
func (it *BlurOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurOrdersMatched represents a OrdersMatched event raised by the Blur contract.
type BlurOrdersMatched struct {
	Maker    common.Address
	Taker    common.Address
	Sell     Order
	SellHash [32]byte
	Buy      Order
	BuyHash  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_Blur *BlurFilterer) FilterOrdersMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address) (*BlurOrdersMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Blur.contract.FilterLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &BlurOrdersMatchedIterator{contract: _Blur.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_Blur *BlurFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *BlurOrdersMatched, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Blur.contract.WatchLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurOrdersMatched)
				if err := _Blur.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_Blur *BlurFilterer) ParseOrdersMatched(log types.Log) (*BlurOrdersMatched, error) {
	event := new(BlurOrdersMatched)
	if err := _Blur.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Blur contract.
type BlurOwnershipTransferredIterator struct {
	Event *BlurOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlurOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurOwnershipTransferred)
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
		it.Event = new(BlurOwnershipTransferred)
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
func (it *BlurOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurOwnershipTransferred represents a OwnershipTransferred event raised by the Blur contract.
type BlurOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blur *BlurFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlurOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blur.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlurOwnershipTransferredIterator{contract: _Blur.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blur *BlurFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlurOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blur.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurOwnershipTransferred)
				if err := _Blur.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Blur *BlurFilterer) ParseOwnershipTransferred(log types.Log) (*BlurOwnershipTransferred, error) {
	event := new(BlurOwnershipTransferred)
	if err := _Blur.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Blur contract.
type BlurUpgradedIterator struct {
	Event *BlurUpgraded // Event containing the contract specifics and raw log

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
func (it *BlurUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurUpgraded)
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
		it.Event = new(BlurUpgraded)
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
func (it *BlurUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurUpgraded represents a Upgraded event raised by the Blur contract.
type BlurUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Blur *BlurFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BlurUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Blur.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BlurUpgradedIterator{contract: _Blur.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Blur *BlurFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BlurUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Blur.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurUpgraded)
				if err := _Blur.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Blur *BlurFilterer) ParseUpgraded(log types.Log) (*BlurUpgraded, error) {
	event := new(BlurUpgraded)
	if err := _Blur.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
