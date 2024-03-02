package gloomberg

import (
	"sync/atomic"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	seawaModels "github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

var (
	TerminalPrinterQueue = make(chan string, viper.GetInt("gloomberg.eventhub.inQueuesSize"))

	_ int64
	_ int64
	_ int64
	_ int64
	_ int64
	_ int64
	_ int64
	_ int64
	_ int64
	_ int64
	// counterPrintToTerminal      int64.
	_ int64
	_ int64
)

// eventHub is a central hub for all events.
type eventHub struct {
	In  eventChannelsIn
	out eventChannelsOut

	counters map[string]*int64

	// info
	CurrentBlock uint64 // old
}

type eventChannelsIn struct {
	ItemListed          chan *seawaModels.ItemListed
	ItemReceivedBid     chan *seawaModels.ItemReceivedBid
	ItemMetadataUpdated chan *seawaModels.ItemMetadataUpdated

	CollectionOffer chan *seawaModels.CollectionOffer

	TxWithLogs        chan *models.TxWithLogs
	TokenTransactions chan *totra.TokenTransaction

	ParsedEvents    chan *degendb.PreformattedEvent
	RecentOwnEvents chan []*degendb.PreformattedEvent

	SeawatcherMgmt          chan *seawaModels.MgmtEvent
	SeawatcherSubscriptions chan *seawaModels.SubscriptionEvent

	// PrintToTerminal chan string
	NewBlock chan uint64
}

type eventChannelsOut struct {
	ItemListed          mapset.Set[chan *seawaModels.ItemListed]
	ItemReceivedBid     mapset.Set[chan *seawaModels.ItemReceivedBid]
	ItemMetadataUpdated mapset.Set[chan *seawaModels.ItemMetadataUpdated]
	CollectionOffer     mapset.Set[chan *seawaModels.CollectionOffer]

	TxWithLogs        []chan *models.TxWithLogs
	TokenTransactions []chan *totra.TokenTransaction

	ParsedEvents    []chan *degendb.PreformattedEvent
	RecentOwnEvents []chan []*degendb.PreformattedEvent

	SeawatcherMgmt          []chan *seawaModels.MgmtEvent
	SeawatcherSubscriptions []chan *seawaModels.SubscriptionEvent

	// PrintToTerminal []chan string
	NewBlock []chan uint64
}

func (eh *eventHub) SubscribeItemListed() chan *seawaModels.ItemListed {
	outChannel := make(chan *seawaModels.ItemListed, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.ItemListed.Add(outChannel)

	return outChannel
}

func (eh *eventHub) UnsubscribeItemListed(itemListedChan chan *seawaModels.ItemListed) {
	eh.out.ItemListed.Remove(itemListedChan)
}

func (eh *eventHub) SubscribeItemReceivedBid() chan *seawaModels.ItemReceivedBid {
	outChannel := make(chan *seawaModels.ItemReceivedBid, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.ItemReceivedBid.Add(outChannel)

	return outChannel
}

func (eh *eventHub) UnsubscribeItemReceivedBid(itemReceivedBidChan chan *seawaModels.ItemReceivedBid) {
	eh.out.ItemReceivedBid.Remove(itemReceivedBidChan)
}

func (eh *eventHub) SubscribeItemMetadataUpdated() chan *seawaModels.ItemMetadataUpdated {
	outChannel := make(chan *seawaModels.ItemMetadataUpdated, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.ItemMetadataUpdated.Add(outChannel)

	return outChannel
}

func (eh *eventHub) UnsubscribeItemMetadataUpdated(itemMetadataUpdatedChan chan *seawaModels.ItemMetadataUpdated) {
	eh.out.ItemMetadataUpdated.Remove(itemMetadataUpdatedChan)
}

func (eh *eventHub) SubscribeCollectionOffer() chan *seawaModels.CollectionOffer {
	outChannel := make(chan *seawaModels.CollectionOffer, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.CollectionOffer.Add(outChannel)

	return outChannel
}

func (eh *eventHub) UnsubscribeCollectionOffer(collectionOfferChan chan *seawaModels.CollectionOffer) {
	eh.out.CollectionOffer.Remove(collectionOfferChan)
}

func (eh *eventHub) SubscribeTxWithLogs() chan *models.TxWithLogs {
	outChannel := make(chan *models.TxWithLogs, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.TxWithLogs = append(eh.out.TxWithLogs, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeTokenTransactions() chan *totra.TokenTransaction {
	outChannel := make(chan *totra.TokenTransaction, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.TokenTransactions = append(eh.out.TokenTransactions, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeParsedEvents() chan *degendb.PreformattedEvent {
	outChannel := make(chan *degendb.PreformattedEvent, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.ParsedEvents = append(eh.out.ParsedEvents, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeRecentOwnEvents() chan []*degendb.PreformattedEvent {
	outChannel := make(chan []*degendb.PreformattedEvent, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.RecentOwnEvents = append(eh.out.RecentOwnEvents, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeSeawatcherMgmt() chan *seawaModels.MgmtEvent {
	outChannel := make(chan *seawaModels.MgmtEvent, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.SeawatcherMgmt = append(eh.out.SeawatcherMgmt, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeSeawatcherSubscriptions() chan *seawaModels.SubscriptionEvent {
	outChannel := make(chan *seawaModels.SubscriptionEvent, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.SeawatcherSubscriptions = append(eh.out.SeawatcherSubscriptions, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribNewBlocks() chan uint64 {
	outChannel := make(chan uint64, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.NewBlock = append(eh.out.NewBlock, outChannel)

	return outChannel
}

func (eh *eventHub) worker(workerID int) {
	for {
		select {
		case event := <-eh.In.TxWithLogs:
			log.Debugf("workerID: %d | len(eh.out.TxWithLogs): %d", workerID, len(eh.out.TxWithLogs))

			atomic.AddInt64(eh.counters["TxWithLogs"], 1)

			for _, ch := range eh.out.TxWithLogs {
				ch <- event
			}
		case event := <-eh.In.TokenTransactions:
			log.Debugf("TokenTransactions event | %d | pushing to %d receivers", workerID, len(eh.out.TokenTransactions))

			atomic.AddInt64(eh.counters["TokenTransactions"], 1)

			for _, ch := range eh.out.TokenTransactions {
				ch <- event
			}
		case event := <-eh.In.ItemListed:
			// log.Debugf("ItemListedEvents event | %d | pushing to %d receivers", workerID, len(eh.out.ItemListed))
			log.Debugf("ItemListedEvents event | %d | pushing to %d receivers", workerID, eh.out.ItemListed.Cardinality())

			atomic.AddInt64(eh.counters["ItemListed"], 1)

			for _, ch := range eh.out.ItemListed.ToSlice() {
				ch <- event
			}
		case event := <-eh.In.ItemReceivedBid:
			log.Debugf("ItemReceivedBid event | %d | pushing to %d receivers", workerID, eh.out.ItemReceivedBid.Cardinality())

			atomic.AddInt64(eh.counters["ItemReceivedBid"], 1)

			for _, ch := range eh.out.ItemReceivedBid.ToSlice() {
				ch <- event
			}
		case event := <-eh.In.ItemMetadataUpdated:
			log.Debugf("ItemMetadataUpdated event | %d | pushing to %d receivers", workerID, eh.out.ItemMetadataUpdated.Cardinality())

			atomic.AddInt64(eh.counters["ItemMetadataUpdated"], 1)

			for _, ch := range eh.out.ItemMetadataUpdated.ToSlice() {
				ch <- event
			}
		case event := <-eh.In.CollectionOffer:
			log.Debugf("CollectionOffer event | %d | pushing to %d receivers", workerID, eh.out.CollectionOffer.Cardinality())

			atomic.AddInt64(eh.counters["CollectionOffer"], 1)

			for _, ch := range eh.out.CollectionOffer.ToSlice() {
				ch <- event
			}
		case event := <-eh.In.ParsedEvents:
			log.Debugf("ParsedEvents event | %d | pushing to %d receivers", workerID, len(eh.out.ParsedEvents))

			atomic.AddInt64(eh.counters["ParsedEvents"], 1)

			for _, outChannel := range eh.out.ParsedEvents {
				outChannel <- event
			}
		case event := <-eh.In.RecentOwnEvents:
			log.Debugf("RecentOwnEvents event | %d | pushing to %d receivers", workerID, len(eh.out.RecentOwnEvents))

			atomic.AddInt64(eh.counters["RecentOwnEvents"], 1)

			for _, ch := range eh.out.RecentOwnEvents {
				ch <- event
			}
		case event := <-eh.In.SeawatcherMgmt:
			log.Debugf("SeawatcherMgmt event | %d | pushing to %d receivers", workerID, len(eh.out.SeawatcherMgmt))

			atomic.AddInt64(eh.counters["SeawatcherMgmt"], 1)

			for _, ch := range eh.out.SeawatcherMgmt {
				ch <- event
			}
		case event := <-eh.In.SeawatcherSubscriptions:
			log.Debugf("SeawatcherSubscriptions event | %d | pushing to %d receivers", workerID, len(eh.out.SeawatcherSubscriptions))

			pretty.Println(eh.counters)

			atomic.AddInt64(eh.counters["SeawatcherSubscriptions"], 1)

			for _, ch := range eh.out.SeawatcherSubscriptions {
				ch <- event
			}

		case event := <-eh.In.NewBlock:
			log.Debugf("CurrentBlock event | %d | pushing to %d receivers", workerID, len(eh.out.NewBlock))

			atomic.AddInt64(eh.counters["NewBlock"], 1)

			for _, ch := range eh.out.NewBlock {
				ch <- event
			}
		}
	}
}
