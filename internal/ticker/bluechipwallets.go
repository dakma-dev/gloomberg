package ticker

import (
	"encoding/json"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"os"
)

type Wallets struct {
	Addresses []*Wallet `json:"addresses"`
}

type Wallet struct {
	Address common.Address `json:"address"`
	Ens     string         `json:"ens"`
	Holder  []HolderTypes
	Score   int32 `json:"score"`
}

type getOwnersForCollection struct {
	OwnerAddresses []string `json:"ownerAddresses"`
}

func (s *Wallet) Contains(e HolderTypes) bool {
	for _, a := range s.Holder {
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
	PUDGY_PENGUINS
	MOONBIRDS
	CloneX
	Goblintown
	Azuki
	CYBERKONGZ
	Captainz
)

func ReadWalletsFromJSON(filePath string) *getOwnersForCollection {
	// read json file
	file, err := os.Open(filePath)
	if err != nil {
		gbl.Log.Error(err)
	}
	defer file.Close()

	// decode json
	var blueChipWallets *getOwnersForCollection

	err = json.NewDecoder(file).Decode(&blueChipWallets)
	if err != nil {
		gbl.Log.Error(err)
	}
	return blueChipWallets
}
