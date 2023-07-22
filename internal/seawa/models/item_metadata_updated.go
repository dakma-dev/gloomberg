package models

import "time"

type ItemMetadataUpdated struct {
	Event   string                     `json:"event_type" mapstructure:"event_type"`
	SentAt  time.Time                  `json:"sent_at"    mapstructure:"sent_at"`
	Payload itemMetadataUpdatedPayload `json:"payload"    mapstructure:"payload"`

	Other map[string]interface{} `mapstructure:",remain"`
}

type itemMetadataUpdatedPayload struct {
	Item               `json:"item"       mapstructure:"item"`
	CollectionCriteria `json:"collection" mapstructure:"collection"`

	CreatedDate time.Time `json:"created_date" mapstructure:"created_date"`

	Other map[string]interface{} `mapstructure:",remain"`
}
