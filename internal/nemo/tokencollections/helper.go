package tokencollections

import (
	"context"
	"math/big"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/ethereum/go-ethereum/common"
)

func GetCollection(gb *gloomberg.Gloomberg, contractAddress common.Address, tokenID int64) *collections.Collection {
	// collection information
	gb.CollectionDB.RWMu.RLock()
	collection := gb.CollectionDB.Collections[contractAddress]
	gb.CollectionDB.RWMu.RUnlock()

	if collection == nil && gb.ProviderPool != nil {
		name := ""

		// if tokenName, err := gb.Nodes.GetERC1155TokenName(contractAddress, big.NewInt(tokenID)); err == nil && tokenName != "" {
		if tokenName, err := gb.ProviderPool.ERC1155TokenName(context.Background(), contractAddress, big.NewInt(tokenID)); err == nil && tokenName != "" {
			name = tokenName
			gbl.Log.Debugf("found token name: %s | %s", name, contractAddress.String())
		} else if err != nil {
			gbl.Log.Debugf("failed to get collection name: %s", err)
		}

		collection = collections.NewCollection(contractAddress, name, gb.ProviderPool, collections.FromStream, gb.Rueidi)

		if collection != nil {
			gb.CollectionDB.RWMu.Lock()
			gb.CollectionDB.Collections[contractAddress] = collection
			gb.CollectionDB.RWMu.Unlock()
		} else {
			gbl.Log.Warnf("❗️ collection not found: %s", contractAddress.String())

			return nil
		}
	}

	return collection
}
