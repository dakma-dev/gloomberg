package mintcmd

import "time"

// transform this json to a go struct.
type TokenInfoRequest struct {
	TokenID         uint64      `json:"tokenId"`
	ContractAddress string      `json:"contractAddress"`
	MarketAddresses []string    `json:"marketAddresses"`
	UserAddress     string      `json:"userAddress"`
	Fingerprint     interface{} `json:"fingerprint"`
}

type TokenInfoResponse struct {
	Status string    `json:"status"`
	Result TokenInfo `json:"result"`
}

type TokenInfo struct {
	TokenID          int    `json:"tokenId"`
	UniversalTokenID string `json:"universalTokenId"`
	ContractAddress  string `json:"contractAddress"`
	Tags             []struct {
		Name string `json:"name"`
	} `json:"tags"`
	Metadata struct {
		Image        string   `json:"image"`
		ContractName string   `json:"contractName"`
		Symbol       string   `json:"symbol"`
		MetadataURI  string   `json:"metadataUri"`
		Name         string   `json:"name"`
		Tags         []string `json:"tags"`
		Media        Media    `json:"media"`
		CreatedBy    string   `json:"createdBy"`
		Description  string   `json:"description"`
		YearCreated  string   `json:"yearCreated"`
	} `json:"metadata"`
	MetadataMedia Media `json:"metadataMedia"`
	Owner         User  `json:"owner"`
	Creator       User  `json:"creator"`
	NftImage      struct {
		ImageMedium             string `json:"imageMedium"`
		ImageSmall              string `json:"imageSmall"`
		ImageBlurred            string `json:"imageBlurred"`
		ImageVideoSmall         string `json:"imageVideoSmall"`
		ImageVideoMedium        string `json:"imageVideoMedium"`
		ImageW800               string `json:"imageW800"`
		ImageW390               string `json:"imageW390"`
		ImageVideoW800          string `json:"imageVideoW800"`
		ImageVideoW390          string `json:"imageVideoW390"`
		ImageBlurredW800        string `json:"imageBlurredW800"`
		ImageBlurredW390        string `json:"imageBlurredW390"`
		ImageArtworkDetail      string `json:"imageArtworkDetail"`
		ImageArtworkDetailImgix string `json:"imageArtworkDetailImgix"`
	} `json:"nftImage"`
	SeriesAttributes            []interface{} `json:"seriesAttributes"`
	AsyncBluePrint              interface{}   `json:"asyncBluePrint"`
	CurrentBidder               interface{}   `json:"currentBidder"`
	CurrentBid                  interface{}   `json:"currentBid"`
	Auction                     AuctionInfo   `json:"auction"`
	Name                        string        `json:"name"`
	Image                       string        `json:"image"`
	ShowAttributePercentages    bool          `json:"showAttributePercentages"`
	CurrentPrice                interface{}   `json:"currentPrice"`
	HasLiked                    bool          `json:"hasLiked"`
	ViewCount                   int           `json:"viewCount"`
	LikeCount                   int           `json:"likeCount"`
	BidAmount                   interface{}   `json:"bidAmount"`
	BidderAddress               interface{}   `json:"bidderAddress"`
	CreatorAddress              string        `json:"creatorAddress"`
	OwnerAddress                string        `json:"ownerAddress"`
	Config                      interface{}   `json:"config"`
	AsyncBlueprintID            interface{}   `json:"asyncBlueprintId"`
	MetadataURI                 string        `json:"metadataUri"`
	StandardImage               string        `json:"standardImage"`
	ThumbnailImage              string        `json:"thumbnailImage"`
	Description                 string        `json:"description"`
	CurrentBidContractAddress   interface{}   `json:"currentBidContractAddress"`
	CurrentPriceContractAddress interface{}   `json:"currentPriceContractAddress"`
	ContractName                string        `json:"contractName"`
	ContractSymbol              string        `json:"contractSymbol"`
	Media                       Media         `json:"media"`
	IsSpace                     bool          `json:"isSpace"`
	SpaceOperatorAddress        string        `json:"spaceOperatorAddress"`
	SpaceSlug                   string        `json:"spaceSlug"`
	SpaceName                   string        `json:"spaceName"`
	IsSeries                    bool          `json:"isSeries"`
}

type AuctionInfo struct {
	ID                     string       `json:"id"`
	TokenID                int          `json:"tokenId"`
	StartingTime           time.Time    `json:"startingTime"`
	EndingTime             time.Time    `json:"endingTime"`
	AuctionContractAddress string       `json:"auctionContractAddress"`
	LengthOfAuction        int          `json:"lengthOfAuction"`
	AuctionCreatorAddress  string       `json:"auctionCreatorAddress"`
	AuctionState           string       `json:"auctionState"`
	AuctionType            string       `json:"auctionType"`
	StartingBlock          int          `json:"startingBlock"`
	EndingBlock            int          `json:"endingBlock"`
	ReservePrice           int64        `json:"reservePrice"`
	MinimumBid             int64        `json:"minimumBid"`
	CurrentAuctionBids     []AuctionBid `json:"currentAuctionBids"`
}

type AuctionBid struct {
	BidderAddress string `json:"bidderAddress"`
	Amount        int64  `json:"amount"`
	Bidder        User   `json:"bidder"`
}

type User struct {
	Username        string `json:"username"`
	Avatar          string `json:"avatar"`
	Ethaddress      string `json:"ethaddress"`
	EthereumAddress string `json:"ethereumAddress"`
	MinimumBid      int    `json:"minimumBid"`
	Bio             string `json:"bio"`
}

type Media struct {
	URI        string `json:"uri"`
	Size       string `json:"size"`
	MimeType   string `json:"mimeType"`
	Dimensions string `json:"dimensions"`
}
