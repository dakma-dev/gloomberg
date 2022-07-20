package internal

import (
	"fmt"
	"strings"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/ethereum/go-ethereum/common"
)

func FormatTokenInfo(tokenID uint64, collection *collections.GbCollection, isMint bool, color bool) string {
	var (
		collectionName = collection.Name
		prefix         = "#"
		id             = fmt.Sprint(tokenID)

		tokenInfo string
	)

	// shorten some names
	collectionName = strings.ReplaceAll(collectionName, "Psychedelics Anonymous", "PA")
	collectionName = strings.ReplaceAll(collectionName, "Open Edition", "OE")
	collectionName = strings.ReplaceAll(collectionName, " Collection", "")

	// remove useless IDs for ens domains
	if collection.ContractAddress == common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85") {
		collectionName = "ENS: Ethereum Name Service"
		prefix = ""
		id = ""
	}

	if color {
		collectionName = collection.Style().Faint(isMint).Render(collectionName)
		id = collection.Style().Faint(isMint).Render(fmt.Sprint(id))
		prefix = collection.StyleSecondary().Faint(isMint).Render(prefix)
	}

	tokenInfo = fmt.Sprintf("%s %s%s", collectionName, prefix, id)

	return tokenInfo
}
