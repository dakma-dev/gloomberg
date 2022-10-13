package gloomberg

import (
	"github.com/benleb/gloomberg/internal/chainwatcher"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/nodes"
	ossw "github.com/benleb/gloomberg/internal/osstreamwatcher"
	"github.com/ethereum/go-ethereum/common"
)

type Gloomberg struct {
	Nodes         *nodes.Nodes
	ChainWatcher  *chainwatcher.ChainWatcher
	StreamWatcher *ossw.OSStreamWatcher
	CollectionDB  *collections.CollectionDB
	OwnWallets    *wallet.Wallets

	WatchUsers *models.WatcherUsers

	OutputQueues map[string]chan *collections.Event

	BasicMIWs map[common.Address]int

	Role RoleMap
}
