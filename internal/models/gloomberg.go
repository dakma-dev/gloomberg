package models

import (
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/ethereum/go-ethereum/core/types"
)

type WorkerChannels struct {
	NewLogs chan *types.Log
	// NewListings chan []models.AssetEvent

	EventsToFormat chan *collections.Event

	OutputWs    chan *collections.Event
	OutputLines chan string
}
