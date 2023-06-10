package models

import (
	"errors"
	"math/big"
	"reflect"
	"strings"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
)

type NftID []string

func NewNftID(chain string, contractAddress common.Address, tokenID *big.Int) NftID {
	return NftID{chain, contractAddress.Hex(), tokenID.String()}
}

func ParseNftID(combinedNftID string) NftID {
	nftID := strings.Split(combinedNftID, "/")

	if len(nftID) != 3 {
		gbl.Log.Error("Invalid NFT ID: %s", combinedNftID)

		return nil
	}

	tokenID, ok := new(big.Int).SetString(nftID[2], 10)
	if !ok {
		gbl.Log.Error("Invalid NFT ID: %s", combinedNftID)

		return nil
	}

	return NewNftID(nftID[0], common.HexToAddress(nftID[1]), tokenID)
}

func (n NftID) Chain() string {
	return n[0]
}

func (n NftID) ContractAddress() common.Address {
	return common.HexToAddress(n[1])
}

func (n NftID) TokenID() *big.Int {
	tokenID, ok := new(big.Int).SetString(n[2], 10)
	if !ok {
		gbl.Log.Error("Invalid NFT ID: %s", n)

		return nil
	}

	return tokenID
}

func (n NftID) TID() string {
	return strings.Join([]string{n.ContractAddress().Hex(), n.TokenID().String()}, "/")
}

func (n NftID) LinkOS() string {
	return style.TerminalLink(n.TID(), "https://opensea.io/assets/"+n.TID())
}

func (n NftID) Equal(other NftID) bool {
	return n.Chain() == other.Chain() &&
		n.ContractAddress() == other.ContractAddress() &&
		n.TokenID().Cmp(other.TokenID()) == 0
}

func (n NftID) String() string {
	return n.TID()
}

// StringToNftIDHookFunc is a mapstructure hook function that converts a string to a NftID.
func StringToNftIDHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(NftID{}) {
			return data, nil
		}

		idString, ok := data.(string)
		if !ok {
			return nil, errors.New("casting NftID to string failed")
		}

		// convert it by parsing
		return ParseNftID(idString), nil
	}
}
