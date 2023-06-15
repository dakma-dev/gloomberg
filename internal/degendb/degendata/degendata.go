package degendata

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func LoadOpenseaRanks(gb *gloomberg.Gloomberg) error {
	ddPathRanks := path.Join(viper.GetString("degendata.path"), "ranks")

	ranksFiles, err := os.ReadDir(ddPathRanks)
	if err != nil {
		log.Errorf("error reading metadata directory %s: %s", ddPathRanks, err)
	}

	log.Debugf("found %d rank files: %+v", len(ranksFiles), ranksFiles)

	totalRanks := 0

	for _, rFile := range ranksFiles {
		log.Debugf("%+v", rFile)

		slug := strings.TrimSuffix(rFile.Name(), "_opensea.json")
		filePath := path.Join(ddPathRanks, rFile.Name())

		address, ok := gb.CollectionDB.OpenseaSlugsAndAddresses()[slug]
		if !ok {
			collectionResponse := opensea.GetCollection(slug)

			if collectionResponse == nil {
				log.Warnf("failed to get collection data for %s from opensea", style.AlmostWhiteStyle.Render(slug))

				continue
			} else if len(collectionResponse.Collection.PrimaryAssetContracts) > 0 {
				address = common.HexToAddress(collectionResponse.Collection.PrimaryAssetContracts[0].Address)
			}

			// don't fuck opensea
			time.Sleep(time.Millisecond * 337)
		}

		ranksOpensea := make(degendb.OpenSeaRanks, 0)
		ranksFile, err := os.Open(filePath)
		if err != nil {
			log.Errorf("failed to open file: %s", err)

			continue
		}
		defer ranksFile.Close()

		ranksBytes, err := io.ReadAll(ranksFile)
		if err != nil {
			log.Errorf("failed to read content from file: %s", err)

			continue
		}

		// parse json to structs
		err = json.Unmarshal(ranksBytes, &ranksOpensea)
		if err != nil {
			log.Printf("failed to decode ranks from json file: %s", err)

			continue
		}

		// gb.PrMod("ddb", fmt.Sprintf("added %s ranks for %s", style.AlmostWhiteStyle.Render(fmt.Sprint(len(ranksOpensea))), style.AlmostWhiteStyle.Render(slug)))

		gb.Ranks[address] = ranksOpensea
		totalRanks += len(ranksOpensea)
	}

	gb.PrMod("ddb", fmt.Sprintf("%s collections with %s ranks in total (opensea)", style.AlmostWhiteStyle.Render(fmt.Sprint(len(gb.Ranks))), style.AlmostWhiteStyle.Render(fmt.Sprint(totalRanks))))

	return nil
}
