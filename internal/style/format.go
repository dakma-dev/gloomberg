package style

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// FormatTokenInfo func FormatTokenInfo(tokenID *big.Int, collection *collections.GbCollection, faint bool, color bool) string {.
func FormatTokenInfo(tokenID *big.Int, collectionName string, primaryStyle lipgloss.Style, secondaryStyle lipgloss.Style, faint bool, color bool) string {
	var (
		prefix = "#"
		id     = fmt.Sprint(tokenID)

		tokenInfo string
	)

	// shorten some names
	collectionName = strings.ReplaceAll(collectionName, "Psychedelics Anonymous", "PA")
	collectionName = strings.ReplaceAll(collectionName, "Open Edition", "OE")
	collectionName = strings.ReplaceAll(collectionName, "Genesis Edition", "Genesis")
	collectionName = strings.ReplaceAll(collectionName, "Golid and Deca", "G&D")
	collectionName = strings.ReplaceAll(collectionName, "[ Ledger ] Market Pass", "Ledger Market Pass")
	collectionName = strings.ReplaceAll(collectionName, "PREMINT Collector Pass - OFFICIAL", "PREMINT Collector Pass")
	collectionName = strings.ReplaceAll(collectionName, " - thestoics.art", "")

	collectionName = strings.ReplaceAll(collectionName, " Collection", "")

	if color {
		collectionName = primaryStyle.Faint(faint).Render(collectionName)
		id = primaryStyle.Faint(faint).Render(id)
		prefix = secondaryStyle.Faint(faint).Render(prefix)
	}

	// convert tokenID to int for more readable comparison
	// if (tokenID.Int64() > 0 && tokenID.Int64() < 999_999) || collectionName == "" {
	if (tokenID.Cmp(big.NewInt(0)) > 0 && tokenID.Cmp(big.NewInt(999_999)) < 0) || collectionName == "" {
		tokenInfo = fmt.Sprintf("%s %s%s", collectionName, prefix, id)
	} else {
		tokenInfo = collectionName
	}

	return tokenInfo
}
