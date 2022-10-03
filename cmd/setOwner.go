package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/benleb/gloomberg/internal/analytics"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// var collectionSets = map[string]*analytics.CollectorSet{
// 	"OSFRLD": {
// 		Name: "Red Lite Collection",
// 		Items: []*analytics.SetItem{
// 			{
// 				CollectionName:  "OSF Editions",
// 				ContractAddress: common.HexToAddress("0xc23a563a26afff06e945ace77173e1568f288ce5"),
// 				FirstToken:      1,
// 				LastToken:       6,
// 				ExcludeTokens:   []int{2},
// 			},
// 			{
// 				CollectionName:  "Red Lite District",
// 				ContractAddress: common.HexToAddress("0x513cd71defc801b9c1aa763db47b5df223da77a2"),
// 				FirstToken:      -1,
// 				LastToken:       -1,
// 			},
// 		},
// 	},
// 	"OSF7Sins": {
// 		Name: "OSF's 7 Deadly Sins",
// 		Items: []*analytics.SetItem{
// 			{
// 				CollectionName:  "OSF's 7 Deadly Sins",
// 				ContractAddress: common.HexToAddress("0x8297d8e55c27aa6ce2d8a65b1fa3debb02410efc"),
// 				FirstToken:      2,
// 				LastToken:       4,
// 			},
// 		},
// 	},
// 	"OSF7SinsKey": {
// 		Name: "OSF's 7 Deadly Sins",
// 		Items: []*analytics.SetItem{
// 			{
// 				CollectionName:  "OSF's 7 Deadly Sins",
// 				ContractAddress: common.HexToAddress("0x8297d8e55c27aa6ce2d8a65b1fa3debb02410efc"),
// 				FirstToken:      3,
// 				LastToken:       3,
// 			},
// 		},
// 	},
// 	"OSFEditions": {
// 		Name: "OSF Editions",
// 		Items: []*analytics.SetItem{
// 			{
// 				CollectionName:  "OSF Editions",
// 				ContractAddress: common.HexToAddress("0xc23a563a26afff06e945ace77173e1568f288ce5"),
// 				FirstToken:      -1,
// 				LastToken:       -1,
// 			},
// 		},
// 	},
// }

// setOwnerCmd represents the setOwner command.
var setOwnerCmd = &cobra.Command{
	Use:   "setOwner",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: setOwner,
}

func init() {
	rootCmd.AddCommand(setOwnerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setOwnerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setOwnerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	setOwnerCmd.Flags().StringVar(&apiKeyMoralis, "moralis", "", "Moralis API Key")
	_ = viper.BindPFlag("api_keys.moralis", setOwnerCmd.Flags().Lookup("moralis"))
}

func setOwner(_ *cobra.Command, args []string) {
	if !viper.IsSet("api_keys.moralis") {
		log.Fatal("api_keys.moralis not set")
	}

	// collectionSet := analytics.CollectionSets[args[0]]
	// if collectionSet == nil {
	// 	log.Fatal("set not found")
	// }

	// multiCollectionSet := analytics.MultiCollectionSets[args[0]]
	// if multiCollectionSet == nil {
	// 	log.Fatal("multiCollectionSet not found")
	// }

	// print header
	header := style.GetHeader(Version)
	fmt.Println(header)
	gbl.Log.Info(header)

	// // print set items
	// for _, item := range set.Items {
	// 	fmt.Println(style.DarkWhiteStyle.Copy().PaddingLeft(4).Render(item.CollectionName))
	// }

	// create clean output dir
	path := "dist/www"
	if err := os.RemoveAll(path); err != nil {
		log.Println(err)
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Println(err)
	}

	ownersPerSet := make(map[string]int, 0)
	ownersPerMuCoSet := make(map[string]int, 0)

	for _, artist := range analytics.Artists {
		for _, multiCollectionSet := range artist.MultiCollectionSets {
			// print set info
			fmt.Println(style.BoldStyle.Copy().Padding(1, 0, 1, 2).Render(multiCollectionSet.Name))

			uniqueMultiSetOwners := make([]common.Address, 0)

			multiSetOwners := make(map[common.Address][]analytics.CollectionSet, 0)

			for _, collectionSet := range multiCollectionSet.CollectionSets {
				setOwners, err := analytics.FetchSetOwnersFor(viper.GetString("api_keys.moralis"), &collectionSet)
				if err != nil {
					gbl.Log.Errorf("error fetching set owners: %v\n", err)
				}

				for _, owner := range setOwners {
					multiSetOwners[owner] = append(multiSetOwners[owner], collectionSet)
				}

				fmt.Println(lipgloss.NewStyle().PaddingLeft(4).Render(fmt.Sprintf("%s owners: %d", collectionSet.Name, len(setOwners))))

				ownersPerSet[collectionSet.ID] = len(setOwners)

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

			for ownerAddress, sets := range multiSetOwners {
				if len(multiCollectionSet.CollectionSets) == len(sets) {
					uniqueMultiSetOwners = append(uniqueMultiSetOwners, ownerAddress)
				}
			}

			ownersPerMuCoSet[multiCollectionSet.ID] = len(uniqueMultiSetOwners)

			f, err := os.Create((path + "/" + multiCollectionSet.ID + ".txt"))
			if err != nil {
				gbl.Log.Errorf("error creating file: %v\n", err)
			}

			for _, owner := range uniqueMultiSetOwners {
				_, err := f.WriteString(owner.Hex() + "\n")
				if err != nil {
					gbl.Log.Errorf("error writing to file: %v\n", err)
				}
			}

			fmt.Println(style.BoldStyle.Copy().Padding(1, 0, 1, 2).Render(fmt.Sprintf("set owners: %d\n", len(uniqueMultiSetOwners))))
		}

		artistFile, err := os.Create((path + "/" + artist.ID + ".html"))
		if err != nil {
			gbl.Log.Errorf("error creating file: %v\n", err)
		}

		artistFile.WriteString("<html><head><title>" + artist.Name + " Sets</title></head><body>")

		for _, muCoSet := range artist.MultiCollectionSets {
			artistFile.WriteString("<h3>" + muCoSet.Name + ": <a href=\"" + muCoSet.ID + ".txt\">" + strconv.Itoa(ownersPerMuCoSet[muCoSet.ID]) + " owners</a></h3>")
			artistFile.WriteString("<ul>")

			for _, set := range muCoSet.CollectionSets {
				artistFile.WriteString("<li>" + set.Name + ": <a href=\"" + set.ID + ".txt\">" + strconv.Itoa(ownersPerSet[set.ID]) + " owners</a></li>")
			}

			artistFile.WriteString("</ul>")
		}

		artistFile.WriteString(fmt.Sprint("<p>last update: ", time.Now().Format("2006-01-02 15:04:05"), "</p>"))
		artistFile.WriteString("</body></html>")
	}

	// write index html
	indexFile, err := os.Create((path + "/index.html"))
	if err != nil {
		gbl.Log.Errorf("error creating file: %v\n", err)
	}

	indexFile.WriteString("<html><head><title>Artists</title></head><body>")

	for _, artist := range analytics.Artists {
		indexFile.WriteString("<h3><a href=\"" + artist.ID + ".html\">" + artist.Name + "</a></h3>")
	}

	indexFile.WriteString(fmt.Sprint("<p>last update: ", time.Now().Format("2006-01-02 15:04:05"), "</p>"))
	indexFile.WriteString("</body></html>")
}

// func getSetOwners(set *analytics.CollectorSet) {
// 	itemOwners := make(map[int][]string, 0)

// 	for idx, item := range set.Items {
// 		item.ItemID = idx
// 		itemOwners[idx] = getItemOwners(item)
// 	}

// 	completeSetOwner := make([]string, 0)

// 	if len(itemOwners) > 1 {
// 		for _, owner := range intersect.Hash(itemOwners[0], itemOwners[1]) {
// 			completeSetOwner = append(completeSetOwner, owner.(string))
// 		}
// 	} else {
// 		completeSetOwner = itemOwners[0]
// 	}

// 	fmt.Println(style.BoldStyle.Copy().PaddingTop(1).Render(fmt.Sprintf("set owner: %4d", len(completeSetOwner))))

// 	fmt.Println(lipgloss.NewStyle().Padding(1, 0, 1, 0).Render(strings.Join(completeSetOwner, "\n")))
// }

// func getItemOwners(item *analytics.SetItem) []string {
// 	// tokenOwners := make(map[int]map[string]bool, 0)
// 	tOwners := make(map[string][]int, 0)
// 	tokenNames := make(map[int]string, 0)

// 	for tokenID := item.FirstToken; tokenID <= item.LastToken; tokenID++ {
// 		if slices.Contains(item.ExcludeTokens, tokenID) {
// 			continue
// 		}

// 		// owners := make([]*analytics.Owner, 0)
// 		tokenNames[tokenID] = ""

// 		responseOwners, err := analytics.FetchOwnersFor(viper.GetString("api_keys.moralis"), item.ContractAddress, tokenID) // getOwnerPage(client, viper.GetString("api_keys.moralis"), item.ContractAddress, tokenID, cursor)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		for ownerAddress := range responseOwners {
// 			// owners = append(owners, owner)
// 			tOwners[ownerAddress] = append(tOwners[ownerAddress], tokenID)
// 		}

// 		// for _, owner := range owners {
// 		// 	if tokenOwners[tokenID] == nil {
// 		// 		tokenOwners[tokenID] = make(map[string]bool, 0)
// 		// 	}

// 		// 	tokenOwners[tokenID][owner.OwnerOf] = true
// 		// }

// 		time.Sleep(time.Second * 1)
// 	}

// 	padstyle := lipgloss.NewStyle().PaddingLeft(6)

// 	for tokenID := item.FirstToken; tokenID <= item.LastToken; tokenID++ {
// 		for ownerAddresses := range tOwners {
// 			fmt.Println(padstyle.Render(fmt.Sprintf("#%d | %4d  - %s", tokenID, len(ownerAddresses), tokenNames[tokenID])))
// 		}
// 	}

// 	completeItemOwner := make([]string, 0)

// 	// var setOwner []interface{}

// 	// for tokenID := item.FirstToken; tokenID <= item.LastToken; tokenID++ {
// 	// 	owners := make(map[string]bool, 0)

// 	// 	if tokenID == item.FirstToken && len(setOwner) == 0 {
// 	// 		owners = tokenOwners[tokenID]
// 	// 	}

// 	// 	for _, owner := range setOwner {
// 	// 		owners[owner.(string)] = true
// 	// 	}

// 	// 	setOwner = intersect.Hash(keys(owners), keys(tokenOwners[tokenID]))

// 	// 	if tokenID == item.LastToken {
// 	// 		for _, owner := range setOwner {
// 	// 			completeItemOwner = append(completeItemOwner, owner.(string))
// 	// 		}
// 	// 	}
// 	// }
// 	// fmt.Printf("completeSetOwner: %d\n\n", len(completeItemOwner))

// 	return completeItemOwner
// }

// func keys(m map[string]bool) []string {
// 	keys := make([]string, 0)
// 	for key := range m {
// 		keys = append(keys, key)
// 	}

// 	return keys
// }
