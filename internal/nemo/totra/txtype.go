package totra

type TxType int64

const (
	Unknown TxType = iota
	Sale
	Mint
	Transfer
	Listing
	Purchase
	Burn
	ReBurn
	Airdrop
	Loan
	CollectionOffer
	ItemBid
	OwnBid
)

func (et TxType) String() string {
	return map[TxType]string{
		Sale:            "sale",
		Mint:            "mint",
		Transfer:        "transfer",
		Listing:         "listing",
		Purchase:        "purchase",
		Burn:            "burn",
		ReBurn:          "reBurn",
		Airdrop:         "airdrop",
		Loan:            "loan",
		CollectionOffer: "collectionOffer",
		ItemBid:         "itemBid",
		OwnBid:          "ownBid",
		Unknown:         "unknown",
	}[et]
}

func (et TxType) Icon() string {
	switch et {
	case Sale:
		return "💰"
	case Mint:
		return "Ⓜ️"
	case Transfer:
		return "📦"
	case Listing:
		return "📢"
	case Purchase:
		return "🛒"
	case Burn:
		return "🔥"
	case ReBurn:
		return "♻️"
	case Airdrop:
		return "🎁"
	case Loan:
		return "💸"
	case CollectionOffer:
		return "🧊"
	case ItemBid:
		return "💦"
	case OwnBid:
		return "🤑"
	case Unknown:
		return "🔬"
	default:
		return "⁉️"
	}
}

func (et TxType) ActionName() string {
	switch et {
	case Sale:
		return "sold"
	case Mint:
		return "minted"
	case Transfer:
		return "transferred"
	case Listing:
		return "listed"
	case Purchase:
		return "purchased"
	case Burn:
		return "burned"
	case ReBurn:
		return "re-burned"
	case Airdrop:
		return "got airdropped"
	case Loan:
		return "(un?)loaned"
	case CollectionOffer:
		return "offered"
	case ItemBid:
		return "got bid"
	case OwnBid:
		return "bid on"
	case Unknown:
		return "did something"
	}

	return "⁉️"
}
