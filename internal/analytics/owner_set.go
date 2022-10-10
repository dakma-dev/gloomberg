package analytics

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CollectionSet struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	ContractAddress common.Address `json:"contractAddress"`
	TokenIDs        []int          `json:"token_ids"`
	Any             bool           `json:"any"`
}

type MultiCollectionSet struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	CollectionSets []CollectionSet `json:"ollectionsets"`
}

type Artist struct {
	ID                  string               `json:"id"`
	Name                string               `json:"name"`
	CollectionSets      []CollectionSet      `json:"collectionsets"`
	MultiCollectionSets []MultiCollectionSet `json:"multicollectionsets"`
}

func GetSetOwner(_ *cobra.Command, _ []string) {
	// create clean output dir
	path := "dist/www"
	if err := os.RemoveAll(path); err != nil {
		log.Println(err)
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Println(err)
	}

	tokenOwners := make(map[string][]common.Address, 0)

	// fetch all owners
	fmt.Println()
	fmt.Println(" " + style.BoldStyle.Render("fetching owners for collection sets..."))
	fmt.Println()

	for _, artist := range Artists {
		for _, collectionSet := range artist.CollectionSets {
			spinnerMsg := "fetching owners for " + style.BoldStyle.Render(artist.Name+" · "+collectionSet.Name)
			spinner := style.GetSpinner(spinnerMsg + "...")
			_ = spinner.Start()

			tokenOwners[collectionSet.ID] = make([]common.Address, 0)

			owners, err := FetchSetOwnersFor(viper.GetString("api_keys.moralis"), &collectionSet)
			if err != nil {
				gbl.Log.Errorf("❌ error fetching owners: %+v", err.Error())
				fmt.Printf("❌ error fetching owners: %+v\n", err.Error())
			}

			// fmt.Printf("%s owners: %v\n", artist.Name, len(owners))

			tokenOwners[collectionSet.ID] = owners

			// fmt.Printf("tokenOwners[collectionSet.ID]: %v\n", len(tokenOwners[collectionSet.ID]))

			// build spinner stop msg with all wallet names
			spinner.StopMessage(spinnerMsg + fmt.Sprintf(": %s", style.BoldStyle.Render(fmt.Sprint(len(tokenOwners[collectionSet.ID])))))

			_ = spinner.Stop()
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(" ✅ fetching owners of collection sets done, lets have a look... 👀")
	fmt.Println()
	fmt.Println()

	for _, artist := range Artists {
		for _, multiCollectionSet := range artist.MultiCollectionSets {
			// print set info
			fmt.Println(style.BoldStyle.Copy().Padding(1, 0, 1, 2).Render(multiCollectionSet.Name))

			multiSetOwners := make(map[common.Address][]CollectionSet, 0)

			for _, collectionSet := range multiCollectionSet.CollectionSets {
				setOwners := tokenOwners[collectionSet.ID]

				for _, owner := range setOwners {
					multiSetOwners[owner] = append(multiSetOwners[owner], collectionSet)
				}

				fmt.Println(lipgloss.NewStyle().PaddingLeft(4).Render(fmt.Sprintf("%s owners: %d", collectionSet.Name, len(setOwners))))

				f, err := os.Create(path + "/" + collectionSet.ID + ".txt")
				if err != nil {
					gbl.Log.Errorf("error creating file: %v\n", err)
				}

				for _, owner := range setOwners {
					_, err := f.WriteString(owner.Hex() + "\n")
					if err != nil {
						gbl.Log.Errorf("error writing to file: %v\n", err)
					}
				}
			}

			// fmt.Printf("multiSetOwners: %d\n", len(multiSetOwners))

			tokenOwners[multiCollectionSet.ID] = make([]common.Address, 0)

			for ownerAddress, sets := range multiSetOwners {
				if len(multiCollectionSet.CollectionSets) == len(sets) {
					tokenOwners[multiCollectionSet.ID] = append(tokenOwners[multiCollectionSet.ID], ownerAddress)
				}
				//  else {
				// 	fmt.Printf("owner %s does not own all sets: %v\n", ownerAddress.Hex(), sets)
				// }
			}

			f, err := os.Create(path + "/" + multiCollectionSet.ID + ".txt")
			if err != nil {
				gbl.Log.Errorf("error creating file: %v\n", err)
			}

			for _, owner := range tokenOwners[multiCollectionSet.ID] {
				_, err := f.WriteString(owner.Hex() + "\n")
				if err != nil {
					gbl.Log.Errorf("error writing to file: %v\n", err)
				}
			}

			fmt.Println(style.BoldStyle.Copy().Padding(1, 0, 1, 4).Render(fmt.Sprintf("set owners: %d\n", len(tokenOwners[multiCollectionSet.ID]))))
		}

		artistFile, err := os.Create(path + "/" + artist.ID + ".html")
		if err != nil {
			gbl.Log.Errorf("error creating file: %v\n", err)
		}

		artistFile.WriteString("<html><head><title>" + artist.Name + " Sets</title></head><body>")

		for _, muCoSet := range artist.MultiCollectionSets {
			artistFile.WriteString("<h3>" + muCoSet.Name + ": <a href=\"" + muCoSet.ID + ".txt\">" + strconv.Itoa(len(tokenOwners[muCoSet.ID])) + " owners</a></h3>")
			artistFile.WriteString("<ul>")

			for _, set := range muCoSet.CollectionSets {
				artistFile.WriteString("<li>" + set.Name + ": <a href=\"" + set.ID + ".txt\">" + strconv.Itoa(len(tokenOwners[set.ID])) + " owners</a></li>")
			}

			artistFile.WriteString("</ul>")
		}

		artistFile.WriteString(fmt.Sprint("<p>last update: ", time.Now().Format("2006-01-02 15:04:05"), "</p>"))
		artistFile.WriteString("</body></html>")
	}

	// write index html
	indexFile, err := os.Create(path + "/index.html")
	if err != nil {
		gbl.Log.Errorf("error creating file: %v\n", err)
	}

	indexFile.WriteString("<html lang=\"en\"><head><title>Artists</title></head><body>")

	for _, artist := range Artists {
		indexFile.WriteString("<h3><a href=\"" + artist.ID + ".html\">" + artist.Name + "</a></h3>")
	}

	indexFile.WriteString(fmt.Sprint("<p>last update: ", time.Now().Format("2006-01-02 15:04:05"), "</p>"))
	indexFile.WriteString("</body></html>")

	fmt.Println()
}

var CollectionSets = map[string]*CollectionSet{
	"Editions": {
		ID:              "Editions",
		Name:            "Editions",
		ContractAddress: common.HexToAddress("0xc23a563a26afff06e945ace77173e1568f288ce5"),
		TokenIDs:        []int{1, 3, 4}, //, 5, 6, 7, 8, 11},
	},
	"RedLiteDistrict": {
		ID:              "RedLiteDistrict",
		Name:            "Red Lite District",
		ContractAddress: common.HexToAddress("0x513cd71defc801b9c1aa763db47b5df223da77a2"),
		TokenIDs:        []int{-1},
	},
	"7DeadlySins": {
		ID:              "7DeadlySins",
		Name:            "7 Deadly Sins",
		ContractAddress: common.HexToAddress("0x8297d8e55c27aa6ce2d8a65b1fa3debb02410efc"),
		TokenIDs:        []int{2, 3, 4}, //, 5, 6, 7, 8, 9},
	},
	"DistilleryPP": {
		ID:              "DistilleryPP",
		Name:            "Distillery - The Private Party",
		ContractAddress: common.HexToAddress("0x26c7de7d475aad40cf8211c0e9ad8469aa4e6878"),
		TokenIDs:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35},
		Any:             true,
	},
	// "DistillerySL": {
	// 	ID:              "DistillerySL",
	// 	Name:            "Distillery - sangsom & lemonade",
	// 	ContractAddress: common.HexToAddress("0x26c7de7d475aad40cf8211c0e9ad8469aa4e6878"),
	// 	TokenIDs:        []int{36, 37, 38, 39, 40},
	// },
}

var MultiCollectionSets = map[string]*MultiCollectionSet{
	"OSFRLDEditions": {
		ID:             "OSFRLDEditions",
		Name:           "RLD & Editions",
		CollectionSets: []CollectionSet{*CollectionSets["Editions"], *CollectionSets["RedLiteDistrict"]},
	},
	"OSFRLDEditionsSins": {
		ID:             "OSFRLDEditionsSins",
		Name:           "RLD & Editions & 7 Sins",
		CollectionSets: []CollectionSet{*CollectionSets["Editions"], *CollectionSets["RedLiteDistrict"], *CollectionSets["7DeadlySins"]},
	},
	"OSFRLDEditionsSinsPP": {
		ID:             "OSFRLDEditionsSinsPP",
		Name:           "RLD & Editions & 7 Sins & Private Party",
		CollectionSets: []CollectionSet{*CollectionSets["Editions"], *CollectionSets["RedLiteDistrict"], *CollectionSets["7DeadlySins"], *CollectionSets["DistilleryPP"]},
	},
}

var Artists = map[string]*Artist{
	"OSF": {
		ID:                  "OSF",
		Name:                "OSF",
		CollectionSets:      []CollectionSet{*CollectionSets["Editions"], *CollectionSets["RedLiteDistrict"], *CollectionSets["7DeadlySins"], *CollectionSets["DistilleryPP"]},
		MultiCollectionSets: []MultiCollectionSet{*MultiCollectionSets["OSFRLDEditions"], *MultiCollectionSets["OSFRLDEditionsSins"], *MultiCollectionSets["OSFRLDEditionsSinsPP"]},
	},
}
