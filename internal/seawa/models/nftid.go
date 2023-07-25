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

// NewNftID returns a new NftID.
func NewNftID(chain string, contractAddress common.Address, tokenID *big.Int) *NftID {
	return &NftID{chain, contractAddress.Hex(), tokenID.String()}
}

// ParseNftID parses a NftID from a string in the format "chain/contractAddress/tokenID".
func ParseNftID(combinedNftID string) *NftID {
	nftID := strings.Split(combinedNftID, "/")

	chain := "ethereum"

	var collection common.Address

	var tokenID *big.Int

	var rawTokenID string

	switch len(nftID) {
	case 3:
		chain = nftID[0]

		collection = common.HexToAddress(nftID[1])

		rawTokenID = nftID[2]

	case 2:
		collection = common.HexToAddress(nftID[0])

		rawTokenID = nftID[1]

	default:
		gbl.Log.Errorf("Invalid NFT ID: %s", combinedNftID)

		empty := []string{"", "", ""}

		return (*NftID)(&empty)
	}

	var ok bool
	tokenID, ok = new(big.Int).SetString(rawTokenID, 10)
	if !ok {
		gbl.Log.Errorf("Invalid NFT ID - error parsing tokenID: %s", combinedNftID)

		empty := []string{"", "", ""}

		return (*NftID)(&empty)
	}

	return NewNftID(chain, collection, tokenID)
}

// Chain returns the chain of the token.
func (n *NftID) Chain() string {
	return (*n)[0]
}

// ContractAddress returns the contract address of the token.
func (n *NftID) ContractAddress() common.Address {
	if len(*n) < 2 {
		gbl.Log.Error("Invalid NFT ID: %s", n)

		return common.Address{}
	}

	return common.HexToAddress((*n)[1])
}

// TokenID returns the tokens ID.
func (n *NftID) TokenID() *big.Int {
	tokenID, ok := new(big.Int).SetString((*n)[2], 10)
	if !ok {
		gbl.Log.Error("Invalid NFT ID: %s", n)

		return nil
	}

	return tokenID
}

// TID returns a shorter variant of the NftID, lacking the chain.
func (n *NftID) TID() string {
	return strings.Join([]string{n.ContractAddress().Hex(), n.TokenID().String()}, "/")
}

// LinkOS returns a link to the token on OpenSea.
func (n *NftID) LinkOS() string {
	return style.TerminalLink(n.TID(), "https://opensea.io/assets/"+n.TID())
}

// Equal returns true if the NftID is equal to the other NftID.
func (n *NftID) Equal(other NftID) bool {
	return n.Chain() == other.Chain() &&
		n.ContractAddress() == other.ContractAddress() &&
		n.TokenID().Cmp(other.TokenID()) == 0
}

func (n *NftID) String() string {
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
