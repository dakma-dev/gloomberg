package degendb

type SlugSubscriptions []SlugSubscription

type SlugSubscription struct {
	Slug   string      `json:"slug"   mapstructure:"slug"`
	Events []EventType `json:"events" mapstructure:"events"`
}

func (ss SlugSubscription) ToStringSlice() []string {
	eventTypes := make([]string, 0, len(ss.Events))

	for _, s := range ss.Events {
		eventTypes = append(eventTypes, s.String())
	}

	return eventTypes
}
