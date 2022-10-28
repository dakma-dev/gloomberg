package chainwatcher

import (
	"time"

	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
)

func SlugWorker(slugTicker *time.Ticker, slugQueue *chan common.Address) {
	for address := range *slugQueue {
		gbl.Log.Infof("fetching opensea slug for: %s", address.Hex())

		if osslug := opensea.GetCollectionSlug(address); osslug != "" {
			cache.StoreOSSlug(address, osslug)
			gbl.Log.Infof("caching opensea slug for: %s\n", address.Hex())
		} else {
			gbl.Log.Warnf("❌ fetching opensea slug for collection %s failed: no slug found", address.Hex())
			return
		}

		gbl.Log.Infof("fetching blur slug for: %s", address.Hex())

		if blurslug, err := external.GetBlurSlugByName(address); err == nil && blurslug != "" {
			cache.StoreBlurSlug(address, blurslug)
			gbl.Log.Infof("caching blur slug for: %s\n", address.Hex())
		} else {
			gbl.Log.Warnf("❌ fetching blur slug for collection %s failed: no slug found", address.Hex())
			return
		}

		<-slugTicker.C
	}
}
