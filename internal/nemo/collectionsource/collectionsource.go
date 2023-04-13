package collectionsource

type CollectionSource int64

const (
	FromConfiguration CollectionSource = iota
	FromWallet
	FromStream
)

func (cs *CollectionSource) MarshalJSON() ([]byte, error) {
	return []byte(`"` + cs.String() + `"`), nil
}

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

func (cs *CollectionSource) String() string {
	return map[CollectionSource]string{
		FromConfiguration: "configuration",
		FromWallet:        "wallet",
		FromStream:        "stream",
	}[*cs]
}
