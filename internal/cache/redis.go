package cache

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

const (
	keywordContractName      string = "contractName"
	keywordENS               string = "ensDomain"
	keywordOSSlug            string = "osslug"
	keywordBlurSlug          string = "blurslug"
	keywordFloorPrice        string = "floor"
	keywordSalira            string = "salira"
	keywordNotificationsLock string = "notification"
	keyDelimiter             string = ":"
)

func keyContract(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordContractName)
}

func keyENS(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordENS)
}

func keyNotificationsLock(txID common.Hash) string {
	return fmt.Sprint(txID.Hex(), keyDelimiter, keywordNotificationsLock)
}

func keyOSSlug(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordOSSlug)
}

func keyBlurSlug(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordBlurSlug)
}

func keyFloorPrice(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordFloorPrice)
}

func keySalira(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordSalira)
}
