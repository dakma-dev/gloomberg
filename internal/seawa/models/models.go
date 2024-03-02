package models

import (
	"time"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/mitchellh/mapstructure"
)

type MgmtAction int64

const (
	// Subscribe client actions.
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
	Action MgmtAction                  `json:"action"`
	Events []osmodels.OpenSeaEventType `json:"events"`
	Slugs  []string                    `json:"slugs"`
}

type SubscriptionEvent struct {
	Action      MgmtAction                `json:"action"      mapstructure:"action"`
	Collections degendb.SlugSubscriptions `json:"collections" mapstructure:"collections"`
}

func GetEventDecoderConfig() mapstructure.DecoderConfig {
	return mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			hooks.StringToAddressHookFunc(),
			hooks.StringToHashHookFunc(),
			hooks.StringToBigIntHookFunc(),
			StringToNftIDHookFunc(),
			mapstructure.OrComposeDecodeHookFunc(
				hooks.StringToUnixTimeHookFunc(),
				mapstructure.StringToTimeHookFunc(time.RFC3339),
			),
		),
	}
}
