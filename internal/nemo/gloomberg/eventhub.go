package gloomberg

import (
	"time"

	chawagoModels "github.com/benleb/gloomberg/internal/chawago/models"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// eventHub is a central hub for all events.
type eventHub struct {
	In  eventChannelsIn
	out eventChannelsOut

	// info
	CurrentBlock uint64
}

type eventChannelsIn struct {
	ItemListed      chan *models.ItemListed
	ItemReceivedBid chan *models.ItemReceivedBid

	CollectionOffer chan *models.CollectionOffer

	TxWithLogs        chan *chawagoModels.TxWithLogs
	TokenTransactions chan *totra.TokenTransaction

	ParsedEvents chan *degendb.ParsedEvent

	SeawatcherMgmt chan *models.MgmtEvent

	PrintToTerminal chan string
	NewBlock        chan uint64
}

type eventChannelsOut struct {
	ItemListed      []chan *models.ItemListed
	ItemReceivedBid []chan *models.ItemReceivedBid

	CollectionOffer []chan *models.CollectionOffer

	TxWithLogs        []chan *chawagoModels.TxWithLogs
	TokenTransactions []chan *totra.TokenTransaction

	ParsedEvents []chan *degendb.ParsedEvent

	SeawatcherMgmt []chan *models.MgmtEvent

	PrintToTerminal []chan string
	NewBlock        []chan uint64
}

func newEventHub() *eventHub {
	eh := eventHub{
		CurrentBlock: 0,

		In: eventChannelsIn{
			ItemListed:      make(chan *models.ItemListed, 1024),
			ItemReceivedBid: make(chan *models.ItemReceivedBid, 1024),

			CollectionOffer: make(chan *models.CollectionOffer, 1024),

			TxWithLogs:        make(chan *chawagoModels.TxWithLogs, 1024),
			TokenTransactions: make(chan *totra.TokenTransaction, 1024),

			ParsedEvents: make(chan *degendb.ParsedEvent, 1024),

			SeawatcherMgmt: make(chan *models.MgmtEvent, 1024),

			PrintToTerminal: make(chan string, 1024),
			NewBlock:        make(chan uint64, 1024),
		},

		out: eventChannelsOut{
			ItemListed:      make([]chan *models.ItemListed, 0),
			ItemReceivedBid: make([]chan *models.ItemReceivedBid, 0),

			CollectionOffer: make([]chan *models.CollectionOffer, 0),

			TxWithLogs:        make([]chan *chawagoModels.TxWithLogs, 0),
			TokenTransactions: make([]chan *totra.TokenTransaction, 0),

			ParsedEvents: make([]chan *degendb.ParsedEvent, 0),

			SeawatcherMgmt: make([]chan *models.MgmtEvent, 0),

			PrintToTerminal: make([]chan string, 0),
			NewBlock:        make([]chan uint64, 0),
		},
	}

	for i := 0; i < viper.GetInt("gloomberg.numEventHubHandlers"); i++ {
		go eh.worker(i)
	}

	// run goroutine that periodically checks all In channel sizes
	go func() {
		for {
			// sum up all In channel sizes
			sum := 0
			sum += len(eh.In.ItemListed)
			sum += len(eh.In.ItemReceivedBid)
			sum += len(eh.In.CollectionOffer)
			sum += len(eh.In.TxWithLogs)
			sum += len(eh.In.TokenTransactions)
			sum += len(eh.In.ParsedEvents)
			sum += len(eh.In.SeawatcherMgmt)
			sum += len(eh.In.PrintToTerminal)
			sum += len(eh.In.NewBlock)

			if sum > 0 {
				log.Printf(
					"eventHub | IListed: %d, IReceivedBid: %d, COffer: %d, TxWithLogs: %d, TTransactions: %d, ParsedEvents: %d, SeawatcherMgmt: %d, PrintToTerminal: %d, NewBlock: %d",
					len(eh.In.ItemListed),
					len(eh.In.ItemReceivedBid),
					len(eh.In.CollectionOffer),
					len(eh.In.TxWithLogs),
					len(eh.In.TokenTransactions),
					len(eh.In.ParsedEvents),
					len(eh.In.SeawatcherMgmt),
					len(eh.In.PrintToTerminal),
					len(eh.In.NewBlock),
				)
			}

			time.Sleep(13 * time.Second)
		}
	}()

	return &eh
}

func (eh *eventHub) SubscribeItemListed() chan *models.ItemListed {
	outChannel := make(chan *models.ItemListed, 1024)
	eh.out.ItemListed = append(eh.out.ItemListed, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeItemReceivedBid() chan *models.ItemReceivedBid {
	outChannel := make(chan *models.ItemReceivedBid, 1024)
	eh.out.ItemReceivedBid = append(eh.out.ItemReceivedBid, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeCollectionOffer() chan *models.CollectionOffer {
	outChannel := make(chan *models.CollectionOffer, 1024)
	eh.out.CollectionOffer = append(eh.out.CollectionOffer, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeTxWithLogs() chan *chawagoModels.TxWithLogs {
	outChannel := make(chan *chawagoModels.TxWithLogs, 1024)
	eh.out.TxWithLogs = append(eh.out.TxWithLogs, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeTokenTransactions() chan *totra.TokenTransaction {
	outChannel := make(chan *totra.TokenTransaction, 1024)
	eh.out.TokenTransactions = append(eh.out.TokenTransactions, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeParsedEvents() chan *degendb.ParsedEvent {
	outChannel := make(chan *degendb.ParsedEvent, 1024)
	eh.out.ParsedEvents = append(eh.out.ParsedEvents, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeSeawatcherMgmt() chan *models.MgmtEvent {
	outChannel := make(chan *models.MgmtEvent, 1024)
	eh.out.SeawatcherMgmt = append(eh.out.SeawatcherMgmt, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribePrintToTerminal() chan string {
	outChannel := make(chan string, 1024)
	eh.out.PrintToTerminal = append(eh.out.PrintToTerminal, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribNewBlocks() chan uint64 {
	outChannel := make(chan uint64, 1024)
	eh.out.NewBlock = append(eh.out.NewBlock, outChannel)

	return outChannel
}

func (eh *eventHub) worker(workerID int) {
	for {
		select {
		case event := <-eh.In.TxWithLogs:
			log.Debugf("TxWithLogs event | %d |  pushing to %d receivers", workerID, len(eh.out.TxWithLogs))

			for _, ch := range eh.out.TxWithLogs {
				ch <- event
			}
		case event := <-eh.In.TokenTransactions:
			log.Debugf("TokenTransactions event | %d | pushing to %d receivers", workerID, len(eh.out.TokenTransactions))

			for _, ch := range eh.out.TokenTransactions {
				ch <- event
			}
		case event := <-eh.In.ItemListed:
			log.Debugf("ItemListedEvents event | %d | pushing to %d receivers", workerID, len(eh.out.ItemListed))

			for _, ch := range eh.out.ItemListed {
				ch <- event
			}
		case event := <-eh.In.ItemReceivedBid:
			log.Debugf("ItemReceivedBid event | %d | pushing to %d receivers", workerID, len(eh.out.ItemReceivedBid))

			for _, ch := range eh.out.ItemReceivedBid {
				ch <- event
			}
		case event := <-eh.In.CollectionOffer:
			log.Debugf("CollectionOffer event | %d | pushing to %d receivers", workerID, len(eh.out.CollectionOffer))

			for _, ch := range eh.out.CollectionOffer {
				ch <- event
			}
		case event := <-eh.In.ParsedEvents:
			log.Debugf("ParsedEvents event | %d | pushing to %d receivers", workerID, len(eh.out.ParsedEvents))

			for _, ch := range eh.out.ParsedEvents {
				ch <- event
			}
		case event := <-eh.In.SeawatcherMgmt:
			log.Debugf("SeawatcherMgmt event | %d | pushing to %d receivers", workerID, len(eh.out.SeawatcherMgmt))

			for _, ch := range eh.out.SeawatcherMgmt {
				ch <- event
			}
		case event := <-eh.In.PrintToTerminal:
			log.Debugf("PrintToTerminal event | %d | pushing to %d receivers", workerID, len(eh.out.PrintToTerminal))

			for _, ch := range eh.out.PrintToTerminal {
				ch <- event
			}
		case event := <-eh.In.NewBlock:
			log.Debugf("CurrentBlock event | %d | pushing to %d receivers", workerID, len(eh.out.NewBlock))

			for _, ch := range eh.out.NewBlock {
				ch <- event
			}
		}
	}
}
