package degendata

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func LoadOpenseaRanks(gb *gloomberg.Gloomberg) error {
	ddPathRanks := path.Join(viper.GetString("degendata.path"), "ranks")
	log.Debugf("loading opensea ranks from %s", ddPathRanks)

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

		address, ok := gb.CollectionDB.OpenSeaSlugsAndAddresses()[slug]
		if !ok {
			if addr, err := gb.Rueidi.GetAddressForOSSlug(context.Background(), slug); err == nil && addr != "" {
				address = common.HexToAddress(addr)

				gloomberg.PrDModf("ddb", "found address %s for slug %s in cache", style.AlmostWhiteStyle.Render(addr), style.AlmostWhiteStyle.Render(slug))
			} else if collectionResponse := opensea.GetCollection(slug); collectionResponse != nil {
				// don't fuck opensea
				time.Sleep(time.Millisecond * 337)

				gloomberg.PrDModf("ddb", "fetched address %s for slug %s from opensea", style.AlmostWhiteStyle.Render(addr), style.AlmostWhiteStyle.Render(slug))

				if len(collectionResponse.Collection.PrimaryAssetContracts) > 0 {
					address = common.HexToAddress(collectionResponse.Collection.PrimaryAssetContracts[0].Address)
				} else {
					log.Warnf("failed to get address for %s from opensea", style.AlmostWhiteStyle.Render(slug))

					continue
				}
			} else {
				log.Warnf("failed to get collection data for %s from opensea", style.AlmostWhiteStyle.Render(slug))

				continue
			}
		} else {
			gloomberg.PrDModf("ddb", fmt.Sprintf("address %s for slug %s from our collectionDB", style.AlmostWhiteStyle.Render(address.Hex()), style.AlmostWhiteStyle.Render(slug)))
		}

		// cache
		if slug != "" && address != (common.Address{}) {
			_ = gb.Rueidi.StoreAddressForOSSlug(context.Background(), slug, address)
			_ = gb.Rueidi.StoreOSSlugForAddress(context.Background(), address, slug)

			gloomberg.PrDModf("ddb", "stored address %s for slug %s in cache", style.AlmostWhiteStyle.Render(address.Hex()), style.AlmostWhiteStyle.Render(slug))
		}

		ranksOpensea := make(degendb.OpenSeaRanks)
		ranksFile, err := os.Open(filePath)
		if err != nil {
			log.Errorf("failed to open file: %s", err)

			continue
		}

		ranksBytes, err := io.ReadAll(ranksFile)
		if err != nil {
			log.Errorf("failed to read content from file: %s", err)

			continue
		}
		ranksFile.Close()

		// parse json to structs
		err = json.Unmarshal(ranksBytes, &ranksOpensea)
		if err != nil {
			log.Printf("failed to decode ranks from json file: %s", err)

			continue
		}

		// validate
		for tokenID, rank := range ranksOpensea {
			if rank.Rank <= 0 {
				gbl.Log.Debugf("%s | rank is <=0 for %s", style.AlmostWhiteStyle.Render(slug), style.AlmostWhiteStyle.Render(strconv.FormatInt(tokenID, 10)))
				gloomberg.PrDModf("ddb", "%s | rank is <=0 for %s", style.AlmostWhiteStyle.Render(slug), style.AlmostWhiteStyle.Render(strconv.FormatInt(tokenID, 10)))

				continue
			}

			if rank.Score <= 0 {
				gbl.Log.Debugf("%s | score is <=0 for %s", style.AlmostWhiteStyle.Render(slug), style.AlmostWhiteStyle.Render(strconv.FormatInt(tokenID, 10)))
				gloomberg.PrDModf("ddb", "%s | score is <=0 for %s", style.AlmostWhiteStyle.Render(slug), style.AlmostWhiteStyle.Render(strconv.FormatInt(tokenID, 10)))

				continue
			}
		}

		gloomberg.PrDModf("ddb", "added %s ranks for %s", style.AlmostWhiteStyle.Render(strconv.Itoa(len(ranksOpensea))), style.AlmostWhiteStyle.Render(slug))

		gb.Ranks[address] = ranksOpensea
		totalRanks += len(ranksOpensea)
	}

	gloomberg.PrMod("ddb", fmt.Sprintf("%s collections with %s ranks in total (opensea)", style.AlmostWhiteStyle.Render(strconv.Itoa(len(gb.Ranks))), style.AlmostWhiteStyle.Render(strconv.Itoa(totalRanks))))

	return nil
}
