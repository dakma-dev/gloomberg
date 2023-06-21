package degendb

type EventType int

const (
	Sale EventType = iota
	Listing
	AcceptedOffer
	AcceptedCollectionOffer
)

func (et EventType) String() string {
	var etName string
	if etName = map[EventType]string{
		Sale:                    "Sale",
		Listing:                 "Listing",
		AcceptedOffer:           "AcceptedOffer",
		AcceptedCollectionOffer: "AcceptedCollectionOffer",
	}[et]; etName == "" {
		etName = "Unknown"
	}

	return etName
}
