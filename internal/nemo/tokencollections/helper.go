package tokencollections

import (
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
)

func GetCollection(gb *gloomberg.Gloomberg, contractAddress common.Address) *collections.Collection {
	// collection information
	gb.CollectionDB.RWMu.RLock()
	collection := gb.CollectionDB.Collections[contractAddress]
	gb.CollectionDB.RWMu.RUnlock()

	if collection == nil && gb.Node() != nil {
		name := ""

		// reactivate me
		// // if tokenName, err := gb.Nodes.GetERC1155TokenName(contractAddress, big.NewInt(tokenID)); err == nil && tokenName != "" {
		// if tokenName, err := gb.Node().GetERC1155TokenName(context.Background(), contractAddress, big.NewInt(tokenID)); err == nil && tokenName != "" {
		// 	name = tokenName
		// 	log.Debugf("found token name: %s | %s", name, contractAddress.String())
		// } else if err != nil {
		// 	log.Debugf("failed to get collection name: %s", err)
		// }

		collection = collections.NewCollection(contractAddress, name, degendb.FromStream, gb.GetRueidica())

		if collection != nil {
			gb.CollectionDB.RWMu.Lock()
			gb.CollectionDB.Collections[contractAddress] = collection
			gb.CollectionDB.RWMu.Unlock()
		} else {
			log.Warnf("❗️ collection not found: %s", contractAddress.String())

			return nil
		}
	}

	return collection
}
