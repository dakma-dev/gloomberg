package opensea

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/notify"
	"github.com/benleb/gloomberg/internal/seawa"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func StartEventHandler(gb *gloomberg.Gloomberg, eventChannel chan map[string]interface{}, seaWatcher *seawa.SeaWatcher) {
	go func() {
		for itemEvent := range eventChannel {
			itemEventType, ok := itemEvent["event_type"].(string)
			if !ok {
				log.Warnf("⚓️ 🤷‍♀️ unknown event type: %s", itemEvent["event_type"])

				continue
			}

			switch osmodels.EventType(itemEventType) {
			case osmodels.ItemListed:
				log.Debugf("⚓️ received %s: %+v", itemEventType, itemEvent)

				var itemListedEvent osmodels.ItemListedEvent

				err := mapstructure.Decode(itemEvent, &itemListedEvent)
				if err != nil {
					log.Info("⚓️❌ decoding incoming opensea stream api event failed:", err)

					continue
				}

				// lawless sniper
				// nftID is a identification string in the format <chain>/<contract>/<tokenID>
				nftID := strings.Split(itemListedEvent.Payload.Item.NftID, "/")
				if len(nftID) != 3 {
					log.Infof("⚓️ 🤷‍♀️ error parsing nftID: %s", itemListedEvent.Payload.Item.NftID)

					continue
				}

				priceWeiRaw, _, err := big.ParseFloat(itemListedEvent.Payload.BasePrice, 10, 64, big.ToNearestEven)
				if err != nil {
					log.Infof("⚓️❌ error parsing price: %s", err.Error())

					continue
				}

				priceWei, _ := priceWeiRaw.Int(nil)
				pricePerTokenGwei := priceWei.Div(priceWei, big.NewInt(int64(itemListedEvent.Payload.Quantity)))
				offerPricePerTokenEther := price.NewPrice(pricePerTokenGwei).Ether()

				if nftID[1] == "0xb119ec7ee48928a94789ed0842309faf34f0c790" {
					name := itemListedEvent.Payload.Item.Metadata.Name
					log.Infof("lawless listing: %s", name)
					// if name contains "-qf"
					if strings.Contains(name, "-qf") {
						// get opensea url
						openseaURL := fmt.Sprintf("https://opensea.io/assets/ethereum/%s/%s", nftID[1], nftID[2])
						// send telegram message
						notify.SendMessageViaTelegram(fmt.Sprintf("lawless listing: %s \n price: %s  url: %s", name, fmt.Sprintf("%5.3f", offerPricePerTokenEther), openseaURL), viper.GetInt64("notifications.manifold.dakma"), "", 0, nil)
					}
				}

			case osmodels.ItemReceivedOffer:
				log.Debugf("⚓️ received OFFER %s: %+v", itemEventType, itemEvent)

			case osmodels.ItemReceivedBid:
				log.Debugf("⚓️ received BID %s: %+v", itemEventType, itemEvent)
				eventType := osmodels.TxType[osmodels.EventType(itemEventType)]

				itemReceivedBidEvent, err := seaWatcher.DecodeItemReceivedBidEvent(itemEvent)
				if err != nil {
					break
				}

				priceWeiRaw, _, err := big.ParseFloat(itemReceivedBidEvent.Payload.BasePrice, 10, 64, big.ToNearestEven)
				if err != nil {
					log.Infof("⚓️❌ error parsing price: %s", err.Error())

					continue
				}

				priceWei, _ := priceWeiRaw.Int(nil)
				pricePerTokenGwei := priceWei.Div(priceWei, big.NewInt(int64(itemReceivedBidEvent.Payload.Quantity)))
				offerPricePerTokenEther := price.NewPrice(pricePerTokenGwei).Ether()
				paymentTokenSymbol := itemReceivedBidEvent.Payload.PaymentToken.Symbol
				collectionSlug := itemReceivedBidEvent.Payload.Collection.Slug

				// nftID is a identification string in the format <chain>/<contract>/<tokenID>
				nftID := strings.Split(itemReceivedBidEvent.Payload.Item.NftID, "/")
				if len(nftID) != 3 {
					log.Infof("⚓️ 🤷‍♀️ error parsing nftID: %s", itemReceivedBidEvent.Payload.Item.NftID)

					continue
				}

				// check bid against own nfts
				tokenID := nftID[2]

				contractAddress := common.HexToAddress(nftID[1])

				if gb.OwnWallets.ContainsToken(contractAddress, tokenID) {
					log.Infof("⚓️🔸 %s |  %s %s for %s #%s", eventType.Icon(), style.TrendRedStyle.Render(fmt.Sprintf("%5.3f", offerPricePerTokenEther)), paymentTokenSymbol, style.BoldStyle.Render(collectionSlug), nftID[2])
					log.Infof("⚓️ 🤑 own token received bid: %s", itemReceivedBidEvent.Payload.Item.NftID)

					continue
				}

			case osmodels.ItemSold:
				log.Debugf("⚓️ received %s: %+v", itemEventType, itemEvent)
			case osmodels.CollectionOffer:
				collectionOfferEvent, err := seaWatcher.DecodeCollectionOfferEvent(itemEvent)
				if err != nil {
					break
				}

				collectionSlug := collectionOfferEvent.Payload.Collection.Slug

				collectionAddress := common.HexToAddress(collectionOfferEvent.Payload.AssetContractCriteria.Address)
				collection, ok := gb.CollectionDB.Collections[collectionAddress]
				if !ok {
					log.Debugf("⚓️ 🤷‍♀️ collection not found: %s", collectionAddress.String())

					continue
				}

				priceWeiRaw, _, err := big.ParseFloat(collectionOfferEvent.Payload.BasePrice, 10, 64, big.ToNearestEven)
				if err != nil {
					log.Infof("⚓️❌ error parsing price: %s", err.Error())

					continue
				}

				priceWei, _ := priceWeiRaw.Int(nil)

				quantity := collectionOfferEvent.Payload.Quantity
				pricePerTokenGwei := priceWei.Div(priceWei, big.NewInt(int64(quantity)))

				offerPricePerTokenEther := price.NewPrice(pricePerTokenGwei).Ether()
				paymentTokenSymbol := collectionOfferEvent.Payload.PaymentToken.Symbol
				eventType := osmodels.TxType[osmodels.EventType(itemEventType)]

				if offerPricePerTokenEther > collection.HighestCollectionOffer {
					collection.HighestCollectionOffer = offerPricePerTokenEther
					log.Debugf("⚓️🔸 %s | %dx %s %s for %s", eventType.Icon(), quantity, style.TrendRedStyle.Render(fmt.Sprintf("%5.3f", offerPricePerTokenEther)), paymentTokenSymbol, style.BoldStyle.Render(collectionSlug))
				}

				if collection.PreviousFloorPrice != 0 {
					if offerPricePerTokenEther > collection.PreviousFloorPrice {
						log.Infof("⚓️‼ ❗❗❗❗ OFFER: price per token %f is higher than floor price %d", offerPricePerTokenEther, big.NewInt(int64(collection.PreviousFloorPrice)))

						break
					}

					break
				}

				gbl.Log.Info("Requesting floor price...")

				floorPriceAlchemyData := external.GetFloorPriceFromAlchemy(collectionOfferEvent.Payload.AssetContractCriteria.Address)

				if floorPriceAlchemyData == nil {
					log.Infof("⚓️❌ error fetching floor price from alchemy for %s", collectionSlug)

					break
				}
				// log.Infof("%s Floor Price: %f (alchemy)", collectionSlug, floorPriceAlchemyData.Opensea.FloorPrice)

				collectionStats := GetCollectionStats(collectionSlug)

				if collectionStats == nil {
					log.Infof("⚓️❌ error fetching collection stats for %s", collectionSlug)

					break
				}

				if floorPriceAlchemyData.Opensea.FloorPrice != floorPriceAlchemyData.Opensea.FloorPrice {
					log.Infof("⚓️❌ floor price mismatch for %s", collectionSlug)
				}

				gbl.Log.Infof("%s Floor Price (OS): %f", collectionSlug, collectionStats.FloorPrice)
				collection.PreviousFloorPrice = floorPriceAlchemyData.Opensea.FloorPrice

				if offerPricePerTokenEther > collection.PreviousFloorPrice {
					log.Infof("⚓️‼ ❗❗❗❗ OFFER: price per token %f is higher than floor price %d", offerPricePerTokenEther, big.NewInt(int64(collection.PreviousFloorPrice)))

					break
				}
			}
		}
	}()
}
