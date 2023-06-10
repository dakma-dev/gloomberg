package models

import (
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
)

type MgmtAction int64

const (
	// client actions.
	Subscribe MgmtAction = iota
	Unsubscribe

	// server actions.

	// SendSlugs is used to request the slugs/events to subscribe to from the clients.
	SendSlugs
)

func (a MgmtAction) String() string {
	switch a {
	case Subscribe:
		return "Subscribe"
	case Unsubscribe:
		return "Unsubscribe"
	case SendSlugs:
		return "SendSlugs"
	default:
		return "unknown"
	}
}

type MgmtEvent struct {
	Action MgmtAction           `json:"action"`
	Events []osmodels.EventType `json:"events"`
	Slugs  []string             `json:"slugs"`
}
