package degendb

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type EventType interface {
	String() string
	ActionName() string
	Icon() string
	OpenseaEventName() string
}

type GBEventType struct {
	name             string
	actionName       string
	icon             string
	openseaEventName string
}

func GetEventType(name string) *GBEventType {
	if eventType, ok := eventTypes[name]; ok {
		return eventType
	}

	return nil
}

func (et *GBEventType) String() string {
	return et.name
}

func (et *GBEventType) ActionName() string {
	return et.actionName
}

func (et *GBEventType) Icon() string {
	return et.icon
}

func (et *GBEventType) OpenseaEventName() string {
	return et.openseaEventName
}

var (
	// event types.
	Unknown                 = &GBEventType{name: "Unknown", actionName: "did something", icon: "❓", openseaEventName: ""}
	Transfer                = &GBEventType{name: "Transfer", actionName: "transferred", icon: "📦", openseaEventName: "item_transferred"}
	Sale                    = &GBEventType{name: "Sale", actionName: "sold", icon: "💰", openseaEventName: "item_sold"}
	Purchase                = &GBEventType{name: "Purchase", actionName: "purchased", icon: "🛍️", openseaEventName: "item_sold"}
	Mint                    = &GBEventType{name: "Mint", actionName: "minted", icon: "Ⓜ️", openseaEventName: ""}
	Airdrop                 = &GBEventType{name: "Airdrop", actionName: "got airdropped", icon: "🎁", openseaEventName: ""}
	Burn                    = &GBEventType{name: "Burn", actionName: "burned", icon: "🔥", openseaEventName: ""}
	BurnRedeem              = &GBEventType{name: "BurnRedeem", actionName: "redeemed burned", icon: "🔥", openseaEventName: ""}
	Loan                    = &GBEventType{name: "Loan", actionName: "loaned", icon: "💸", openseaEventName: ""}
	RepayLoan               = &GBEventType{name: "RepayLoan", actionName: "repaid loan", icon: "💸", openseaEventName: ""}
	Listing                 = &GBEventType{name: "Listing", actionName: "listed", icon: "📢", openseaEventName: "item_listed"}
	Bid                     = &GBEventType{name: "Bid", actionName: "(got) bid", icon: "💦", openseaEventName: "item_received_bid"}
	OwnBid                  = &GBEventType{name: "OwnBid", actionName: "bid", icon: "🤑", openseaEventName: ""}
	AcceptedOffer           = &GBEventType{name: "AcceptedOffer", actionName: "accepted offer", icon: "🤝", openseaEventName: ""}
	CollectionOffer         = &GBEventType{name: "CollectionOffer", actionName: "(got) collection-offered", icon: "☂️", openseaEventName: "collection_offer"} // 🧊
	AcceptedCollectionOffer = &GBEventType{name: "AcceptedCollectionOffer", actionName: "accepted collection offer", icon: "🤝", openseaEventName: ""}
	MetadataUpdated         = &GBEventType{name: "MetadataUpdated", actionName: "metadata updated", icon: "♻️", openseaEventName: "item_metadata_updated"}
	Cancelled               = &GBEventType{name: "Cancelled", actionName: "cancelled", icon: "❌", openseaEventName: "item_cancelled"}

	// event type sets.
	SaleTypes = mapset.NewSet[EventType](Sale, Purchase)

	// map of lowercase_with_underscores openseaEventName to event type.
	eventTypes = map[string]*GBEventType{
		"item_transferred":      Transfer,
		"item_sold":             Sale,
		"item_listed":           Listing,
		"item_received_bid":     Bid,
		"item_metadata_updated": MetadataUpdated,
		"item_cancelled":        Cancelled,
		"collection_offer":      CollectionOffer,
	}
)
