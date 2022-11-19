package gloomberg

import (
	"math"
	"math/big"

	"github.com/benleb/gloomberg/internal/chainwatcher"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/nodes"
	ossw "github.com/benleb/gloomberg/internal/osstreamwatcher"
	"github.com/benleb/gloomberg/internal/web"
	"github.com/ethereum/go-ethereum/common"
)

type Gloomberg struct {
	GasPrice       int
	ChainWatcher   *chainwatcher.ChainWatcher
	GloomWeb       *web.GloomWeb
	Nodes          *nodes.Nodes
	StreamWatcher  *ossw.OSStreamWatcher
	Watcher        *models.Watcher
	WebEventStream *web.EventStream

	CollectionDB *collections.CollectionDB
	OwnWallets   *wallet.Wallets

	// WatchUsers *models.WatcherUsers

	// WatchGroups *models.WatchGroups
	// WatchRules  *models.WatchRules

	OutputQueues map[string]chan *collections.Event
	QueueSlugs   chan common.Address
	BasicMIWs    map[common.Address]int
}

func (gb *Gloomberg) GetGasPrice() int {
	if gasInfo, err := gb.Nodes.GetRandomLocalNode().GetCurrentGasInfo(); err == nil && gasInfo != nil {
		// gas price
		if gasInfo.GasPriceWei.Cmp(big.NewInt(0)) > 0 {
			gasPriceGwei, _ := nodes.WeiToGwei(gasInfo.GasPriceWei).Float64()
			gasPrice := int(math.Round(gasPriceGwei))
			gb.GasPrice = gasPrice
			// gb.WebEventStream.GasPrice = &gb.GasPrice
			// gbl.Log.Infof("set gas price gb.GasPrice: %v | gb.WebEventStream.GasPrice: %v", gb.GasPrice, gb.WebEventStream.GasPrice)

			return gb.GasPrice
		}
	}

	return 0
}
