package ticker

import (
	"encoding/json"
	"os"

	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
)

type Wallets struct {
	Addresses []*Wallet `json:"addresses"`
}

type Wallet struct {
	Address common.Address `json:"address"`
	Ens     string         `json:"ens"`
	Types   []HolderTypes
	Score   int32 `json:"score"`
}

type GetOwnersForCollectionResponse struct {
	OwnerAddresses []string `json:"ownerAddresses"`
}

func (s *Wallet) Contains(e HolderTypes) bool {
	for _, a := range s.Types {
		if a == e {
			return true
		}
	}

	return false
}

type HolderTypes int64

const (
	BAYC HolderTypes = iota
	MAYC
	CryptoPunks
	RLD
	DOODLES
	PUDGYPENGUINS
	MOONBIRDS
	CloneX
	Goblintown
	Azuki
	CYBERKONGZ
	Captainz
	DeGods
)

func ReadWalletsFromJSON(filePath string) *GetOwnersForCollectionResponse {
	// read json file
	file, err := os.Open(filePath)
	if err != nil {
		gbl.Log.Error(err)
	}
	defer file.Close()

	// decode json
	var blueChipWallets *GetOwnersForCollectionResponse

	err = json.NewDecoder(file).Decode(&blueChipWallets)
	if err != nil {
		gbl.Log.Error(err)
	}

	return blueChipWallets
}
