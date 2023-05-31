package gloomberg

import (
	chawagoModels "github.com/benleb/gloomberg/internal/chawago/models"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/charmbracelet/log"
)

// eventHub is a central hub for all events.
type eventHub struct {
	In struct {
		ItemListedEvents chan *osmodels.ItemListedEvent
		TxWithLogs       chan *chawagoModels.TxWithLogs
		PrintToTerminal  chan string

		ItemEvents chan *osmodels.ItemEvent
	}
	out struct {
		ItemListedEvents []chan *osmodels.ItemListedEvent
		TxWithLogs       []chan *chawagoModels.TxWithLogs
		PrintToTerminal  []chan string

		ItemEvents []chan *osmodels.ItemEvent
	}
}

func newEventHub() *eventHub {
	eh := eventHub{
		In: struct {
			ItemListedEvents chan *osmodels.ItemListedEvent
			TxWithLogs       chan *chawagoModels.TxWithLogs
			PrintToTerminal  chan string

			ItemEvents chan *osmodels.ItemEvent
		}{
			ItemListedEvents: make(chan *osmodels.ItemListedEvent, 1024),
			TxWithLogs:       make(chan *chawagoModels.TxWithLogs, 1024),
			PrintToTerminal:  make(chan string, 1024),

			ItemEvents: make(chan *osmodels.ItemEvent, 1024),
		},
		out: struct {
			ItemListedEvents []chan *osmodels.ItemListedEvent
			TxWithLogs       []chan *chawagoModels.TxWithLogs
			PrintToTerminal  []chan string

			ItemEvents []chan *osmodels.ItemEvent
		}{
			ItemListedEvents: make([]chan *osmodels.ItemListedEvent, 0),
			TxWithLogs:       make([]chan *chawagoModels.TxWithLogs, 0),
			PrintToTerminal:  make([]chan string, 0),

			ItemEvents: make([]chan *osmodels.ItemEvent, 0),
		},
	}

	go eh.worker()
	go eh.worker()
	go eh.worker()

	return &eh
}

func (eh *eventHub) SubscribeItemEvents() chan *osmodels.ItemEvent {
	outChannel := make(chan *osmodels.ItemEvent, 1024)
	eh.out.ItemEvents = append(eh.out.ItemEvents, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeItemListed() chan *osmodels.ItemListedEvent {
	outChannel := make(chan *osmodels.ItemListedEvent, 1024)
	eh.out.ItemListedEvents = append(eh.out.ItemListedEvents, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribeTxWithLogs() chan *chawagoModels.TxWithLogs {
	outChannel := make(chan *chawagoModels.TxWithLogs, 1024)
	eh.out.TxWithLogs = append(eh.out.TxWithLogs, outChannel)

	return outChannel
}

func (eh *eventHub) SubscribePrintToTerminal() chan string {
	outChannel := make(chan string, 1024)
	eh.out.PrintToTerminal = append(eh.out.PrintToTerminal, outChannel)

	return outChannel
}

func (eh *eventHub) worker() {
	for {
		select {
		case event := <-eh.In.ItemListedEvents:
			log.Debugf("ItemListedEvents event | pushing to %d receivers", len(eh.out.ItemListedEvents))

			for _, ch := range eh.out.ItemListedEvents {
				ch <- event
			}
		case event := <-eh.In.PrintToTerminal:
			log.Debugf("PrintToTerminal event | pushing to %d receivers", len(eh.out.PrintToTerminal))

			for _, ch := range eh.out.PrintToTerminal {
				ch <- event
			}
		case event := <-eh.In.TxWithLogs:
			log.Debugf("TxWithLogs event | pushing to %d receivers", len(eh.out.TxWithLogs))

			for _, ch := range eh.out.TxWithLogs {
				ch <- event
			}
		}
	}
}
