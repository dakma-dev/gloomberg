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
)

func (et TxType) String() string {
	return map[TxType]string{
		Sale: "Sale", Mint: "Mint", Transfer: "Transfer", Listing: "Listing", Purchase: "Purchase", Burn: "Burn", ReBurn: "ReBurn", Airdrop: "Airdrop", Unknown: "Unknown", Loan: "Loan",
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
		return "👋"
	case ItemBid:
		return "🖐️"
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
	case Loan:
		return "(un?)loaned"
	case Airdrop:
		return "got airdropped"
	case Unknown:
		return "did something"
	}

	return "⁉️"
}
