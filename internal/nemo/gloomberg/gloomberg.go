package gloomberg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/nemo/watch"
	"github.com/benleb/gloomberg/internal/seawa"
	"github.com/benleb/gloomberg/internal/stats"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
)

type Gloomberg struct {
	// Nodes        *nodes.Nodes
	ProviderPool *provider.Pool
	Watcher      *watch.Watcher

	CollectionDB *collections.CollectionDB
	OwnWallets   *wallet.Wallets
	Stats        *stats.Stats

	Rdb *redis.Client

	QueueSlugs chan common.Address
}

func (gb *Gloomberg) SendSlugsToServer() {
	// to enable multiple users to use the central gloomberg instance for events from opensea,
	// we first send the slugs of 'our' collections to the events-subscriptions channel.
	// the central gloomberg instance then creates a subscription on the opensea
	// api and publishes upcoming incoming events to the pubsub channel
	// marshal event to json
	slugs := gb.CollectionDB.OpenseaSlugs()
	if len(slugs) == 0 {
		gbl.Log.Warn("❌ no slugs to send to gloomberg server")

		return
	}

	mgmtEvent := &seawa.MgmtEvent{Action: seawa.Subscribe, Slugs: slugs}

	jsonMgmtEvent, err := json.Marshal(mgmtEvent)
	if err != nil {
		gbl.Log.Error("❌ marshal failed for outgoing list of collection slugs: %s | %v", err, gb.CollectionDB.OpenseaSlugs())

		return
	}

	if err := gb.Rdb.Publish(context.Background(), internal.TopicSeaWatcherMgmt, jsonMgmtEvent).Err(); err != nil {
		gbl.Log.Warnf("error publishing event to redis: %s", err.Error())
	} else {
		gbl.Log.Infof("📢 sent %s collection slugs to %s", style.BoldStyle.Render(fmt.Sprint(len(gb.CollectionDB.OpenseaSlugs()))), style.BoldStyle.Render(internal.TopicSeaWatcherMgmt))
	}
}
