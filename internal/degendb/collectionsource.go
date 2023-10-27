package degendb

// CollectionSource represents the source of a collection.
type CollectionSource int64

const (
	FromConfiguration CollectionSource = iota // FromConfiguration represents a collection explicitly defined in the configuration.
	FromWallet                                // FromWallet represents a collection from a token in one of our wallets.
	FromStream                                // FromStream represents a collection gathered from stream.
)

// MarshalJSON marshals the CollectionSource to JSON.
func (cs *CollectionSource) MarshalJSON() ([]byte, error) {
	return []byte(`"` + cs.String() + `"`), nil
}

// UnmarshalJSON unmarshals the CollectionSource from JSON.
func (cs *CollectionSource) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `"configuration"`:
		*cs = FromConfiguration
	case `"wallet"`:
		*cs = FromWallet
	case `"stream"`:
		*cs = FromStream
	}

	return nil
}

// String returns the string representation of the CollectionSource.
func (cs *CollectionSource) String() string {
	return map[CollectionSource]string{
		FromConfiguration: "configuration",
		FromWallet:        "wallet",
		FromStream:        "stream",
	}[*cs]
}
