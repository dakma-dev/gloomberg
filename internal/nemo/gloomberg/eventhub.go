package gloomberg

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	chawagoModels "github.com/benleb/gloomberg/internal/chawago/models"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/spf13/viper"
)

var (
	TerminalPrinterQueue = make(chan string, viper.GetInt("gloomberg.eventhub.inQueuesSize"))

	counterItemListed              int64
	counterItemReceivedBid         int64
	counterItemMetadataUpdated     int64
	counterCollectionOffer         int64
	counterTxWithLogs              int64
	counterTokenTransactions       int64
	counterParsedEvents            int64
	counterRecentOwnEvents         int64
	counterSeawatcherMgmt          int64
	counterSeawatcherSubscriptions int64
	// counterPrintToTerminal      int64.
	counterTerminalPrinterQueue int64
	counterNewBlock             int64
)

// eventHub is a central hub for all events.
type eventHub struct {
	In  eventChannelsIn
	out eventChannelsOut

	counters map[string]*int64

	// info
	CurrentBlock uint64
}

type eventChannelsIn struct {
	ItemListed          chan *models.ItemListed
	ItemReceivedBid     chan *models.ItemReceivedBid
	ItemMetadataUpdated chan *models.ItemMetadataUpdated

	CollectionOffer chan *models.CollectionOffer

	TxWithLogs        chan *chawagoModels.TxWithLogs
	TokenTransactions chan *totra.TokenTransaction

	ParsedEvents    chan *degendb.PreformattedEvent
	RecentOwnEvents chan []*degendb.PreformattedEvent

	SeawatcherMgmt          chan *models.MgmtEvent
	SeawatcherSubscriptions chan *models.SubscriptionEvent

	// PrintToTerminal chan string
	NewBlock chan uint64
}

type eventChannelsOut struct {
	ItemListed          mapset.Set[chan *models.ItemListed]
	ItemReceivedBid     mapset.Set[chan *models.ItemReceivedBid]
	ItemMetadataUpdated mapset.Set[chan *models.ItemMetadataUpdated]
	CollectionOffer     mapset.Set[chan *models.CollectionOffer]

	TxWithLogs        []chan *chawagoModels.TxWithLogs
	TokenTransactions []chan *totra.TokenTransaction

	ParsedEvents    []chan *degendb.PreformattedEvent
	RecentOwnEvents []chan []*degendb.PreformattedEvent

	SeawatcherMgmt          []chan *models.MgmtEvent
	SeawatcherSubscriptions []chan *models.SubscriptionEvent

	// PrintToTerminal []chan string
	NewBlock []chan uint64
}

func newEventHub() *eventHub {
	eh := eventHub{
		CurrentBlock: 0,

		counters: map[string]*int64{
			"ItemListed":             &counterItemListed,
			"ItemReceivedBid":        &counterItemReceivedBid,
			"ItemMetadataUpdated":    &counterItemMetadataUpdated,
			"CollectionOffer":        &counterCollectionOffer,
			"TxWithLogs":             &counterTxWithLogs,
			"TokenTransactions":      &counterTokenTransactions,
			"ParsedEvents":           &counterParsedEvents,
			"RecentOwnEvents":        &counterRecentOwnEvents,
			"SeawatcherMgmt":         &counterSeawatcherMgmt,
			"SewatcherSubscriptions": &counterSeawatcherSubscriptions, // TODO: add counter
			// "PrintToTerminal":     &counterPrintToTerminal,
			"TerminalPrinterQueue": &counterTerminalPrinterQueue,
			"NewBlock":             &counterNewBlock,
		},

		In: eventChannelsIn{
			ItemListed:          make(chan *models.ItemListed, viper.GetInt("gloomberg.eventhub.inQueuesSize")),
			ItemReceivedBid:     make(chan *models.ItemReceivedBid, viper.GetInt("gloomberg.eventhub.inQueuesSize")),
			ItemMetadataUpdated: make(chan *models.ItemMetadataUpdated, viper.GetInt("gloomberg.eventhub.inQueuesSize")),
			CollectionOffer:     make(chan *models.CollectionOffer, viper.GetInt("gloomberg.eventhub.inQueuesSize")),

			TxWithLogs:        make(chan *chawagoModels.TxWithLogs, viper.GetInt("gloomberg.eventhub.inQueuesSize")),
			TokenTransactions: make(chan *totra.TokenTransaction, viper.GetInt("gloomberg.eventhub.inQueuesSize")),

			ParsedEvents:    make(chan *degendb.PreformattedEvent, viper.GetInt("gloomberg.eventhub.inQueuesSize")),
			RecentOwnEvents: make(chan []*degendb.PreformattedEvent, viper.GetInt("gloomberg.eventhub.inQueuesSize")),

			SeawatcherMgmt:          make(chan *models.MgmtEvent, viper.GetInt("gloomberg.eventhub.inQueuesSize")),
			SeawatcherSubscriptions: make(chan *models.SubscriptionEvent, viper.GetInt("gloomberg.eventhub.inQueuesSize")),

			// PrintToTerminal: make(chan string, viper.GetInt("gloomberg.eventhub.inQueuesSize")),
			NewBlock: make(chan uint64, viper.GetInt("gloomberg.eventhub.inQueuesSize")),
		},

		out: eventChannelsOut{
			ItemListed:          mapset.NewSet[chan *models.ItemListed](),
			ItemReceivedBid:     mapset.NewSet[chan *models.ItemReceivedBid](),
			ItemMetadataUpdated: mapset.NewSet[chan *models.ItemMetadataUpdated](),
			CollectionOffer:     mapset.NewSet[chan *models.CollectionOffer](),

			TxWithLogs:        make([]chan *chawagoModels.TxWithLogs, 0),
			TokenTransactions: make([]chan *totra.TokenTransaction, 0),

			ParsedEvents:    make([]chan *degendb.PreformattedEvent, 0),
			RecentOwnEvents: make([]chan []*degendb.PreformattedEvent, 0),

			SeawatcherMgmt:          make([]chan *models.MgmtEvent, 0),
			SeawatcherSubscriptions: make([]chan *models.SubscriptionEvent, 0),

			// PrintToTerminal: make([]chan string, 0),
			NewBlock: make([]chan uint64, 0),
		},
	}

	for i := 0; i < viper.GetInt("gloomberg.eventhub.numHandler"); i++ {
		go eh.worker(i)
	}

	// run goroutine that periodically checks all In channel sizes
	go func() {
		for {
			// // sum up all In channel sizes
			// sum := 0
			// sum += len(eh.In.ItemListed)
			// sum += len(eh.In.ItemReceivedBid)
			// sum += len(eh.In.CollectionOffer)
			// sum += len(eh.In.TxWithLogs)
			// sum += len(eh.In.TokenTransactions)
			// sum += len(eh.In.ParsedEvents)
			// sum += len(eh.In.RecentOwnEvents)
			// sum += len(eh.In.SeawatcherMgmt)
			// sum += len(eh.In.PrintToTerminal)
			// sum += len(eh.In.NewBlock)

			// pretty.Println(eh.counters)
			chans := map[string]int{
				"ItemListed":              len(eh.In.ItemListed),
				"ItemReceivedBid":         len(eh.In.ItemReceivedBid),
				"ItemMetadataUpdated":     len(eh.In.ItemMetadataUpdated),
				"CollectionOffer":         len(eh.In.CollectionOffer),
				"TxWithLogs":              len(eh.In.TxWithLogs),
				"TokenTransactions":       len(eh.In.TokenTransactions),
				"ParsedEvents":            len(eh.In.ParsedEvents),
				"RecentOwnEvents":         len(eh.In.RecentOwnEvents),
				"SeawatcherMgmt":          len(eh.In.SeawatcherMgmt),
				"SeawatcherSubscriptions": len(eh.In.SeawatcherSubscriptions),
				// "PrintToTerminal":     len(eh.In.PrintToTerminal),
				"TerminalPrinterQueue": len(TerminalPrinterQueue),
				"NewBlock":             len(eh.In.NewBlock),
			}

			outChans := map[string]int{
				// "outItemListed":          len(eh.out.ItemListed),
				"outItemListed":              eh.out.ItemListed.Cardinality(),
				"outItemReceivedBid":         eh.out.ItemReceivedBid.Cardinality(),
				"outItemMetadataUpdated":     eh.out.ItemMetadataUpdated.Cardinality(),
				"outCollectionOffer":         eh.out.CollectionOffer.Cardinality(),
				"outTxWithLogs":              len(eh.out.TxWithLogs),
				"outTokenTransactions":       len(eh.out.TokenTransactions),
				"outParsedEvents":            len(eh.out.ParsedEvents),
				"outRecentOwnEvents":         len(eh.out.RecentOwnEvents),
				"outSeawatcherMgmt":          len(eh.out.SeawatcherMgmt),
				"outSeawatcherSubscriptions": len(eh.out.SeawatcherSubscriptions),
				// "outPrintToTerminal":     len(eh.out.PrintToTerminal),
				"outNewBlock": len(eh.out.NewBlock),
			}

			inWarnings := strings.Builder{}
			outWarnings := strings.Builder{}

			for name, inQueue := range chans {
				if inQueue > 3 {
					inWarnings.WriteString(fmt.Sprintf(" | %s inQueue queue filling up: %d", name, inQueue))
				}
			}

			for name, outQueues := range outChans {
				if outQueues > 3 {
					outWarnings.WriteString(fmt.Sprintf(" | %s: %d", name, outQueues))
				}
			}

			if outWarnings.Len() > 0 {
				saved := outWarnings.String()
				outWarnings.Reset()
				inWarnings.WriteString("subscriptions: " + saved)
			}

			if inWarnings.Len() > 0 && outWarnings.Len() > 0 {
				inWarnings.WriteString("\n")
			}

			if inWarnings.Len()+outWarnings.Len() > 0 {
				log.Printf("🤬 eventHub" + inWarnings.String() + outWarnings.String())
			}

			time.Sleep(3 * time.Second)
		}
	}()

	return &eh
}

func (eh *eventHub) SubscribeItemListed() chan *models.ItemListed {
	outChannel := make(chan *models.ItemListed, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.ItemListed.Add(outChannel)

	return outChannel
}

func (eh *eventHub) UnsubscribeItemListed(itemListedChan chan *models.ItemListed) {
	eh.out.ItemListed.Remove(itemListedChan)
}

func (eh *eventHub) SubscribeItemReceivedBid() chan *models.ItemReceivedBid {
	outChannel := make(chan *models.ItemReceivedBid, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.ItemReceivedBid.Add(outChannel)

	return outChannel
}

func (eh *eventHub) UnsubscribeItemReceivedBid(itemReceivedBidChan chan *models.ItemReceivedBid) {
	eh.out.ItemReceivedBid.Remove(itemReceivedBidChan)
}

func (eh *eventHub) SubscribeItemMetadataUpdated() chan *models.ItemMetadataUpdated {
	outChannel := make(chan *models.ItemMetadataUpdated, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.ItemMetadataUpdated.Add(outChannel)

	return outChannel
}

func (eh *eventHub) UnsubscribeItemMetadataUpdated(itemMetadataUpdatedChan chan *models.ItemMetadataUpdated) {
	eh.out.ItemMetadataUpdated.Remove(itemMetadataUpdatedChan)
}

func (eh *eventHub) SubscribeCollectionOffer() chan *models.CollectionOffer {
	outChannel := make(chan *models.CollectionOffer, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.CollectionOffer.Add(outChannel)

	return outChannel
}

func (eh *eventHub) UnsubscribeCollectionOffer(collectionOfferChan chan *models.CollectionOffer) {
	eh.out.CollectionOffer.Remove(collectionOfferChan)
}

func (eh *eventHub) SubscribeTxWithLogs() chan *chawagoModels.TxWithLogs {
	outChannel := make(chan *chawagoModels.TxWithLogs, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
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

func (eh *eventHub) SubscribeSeawatcherMgmt() chan *models.MgmtEvent {
	outChannel := make(chan *models.MgmtEvent, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.SeawatcherMgmt = append(eh.out.SeawatcherMgmt, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeSeawatcherSubscriptions() chan *models.SubscriptionEvent {
	outChannel := make(chan *models.SubscriptionEvent, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
	eh.out.SeawatcherSubscriptions = append(eh.out.SeawatcherSubscriptions, outChannel)

	return outChannel
}

// func (eh *eventHub) SubscribePrintToTerminal() chan string {
// 	outChannel := make(chan string, viper.GetInt("gloomberg.eventhub.outQueuesSize"))
// 	eh.out.PrintToTerminal = append(eh.out.PrintToTerminal, outChannel)

// 	return outChannel
// }

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
			gbl.Log.Debugf("ParsedEvents event | %d | pushing to %d receivers", workerID, len(eh.out.ParsedEvents))

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

			atomic.AddInt64(eh.counters["SeawatcherSubscriptions"], 1)

			for _, ch := range eh.out.SeawatcherSubscriptions {
				ch <- event
			}
		// case event := <-eh.In.PrintToTerminal:
		// 	log.Debugf("PrintToTerminal event | %d | pushing to %d receivers", workerID, len(eh.out.PrintToTerminal))

		// 	atomic.AddInt64(eh.counters["PrintToTerminal"], 1)

		// 	for _, ch := range eh.out.PrintToTerminal {
		// 		ch <- event
		// 	}
		case event := <-eh.In.NewBlock:
			log.Debugf("CurrentBlock event | %d | pushing to %d receivers", workerID, len(eh.out.NewBlock))

			atomic.AddInt64(eh.counters["NewBlock"], 1)

			for _, ch := range eh.out.NewBlock {
				ch <- event
			}
		}
	}
}
