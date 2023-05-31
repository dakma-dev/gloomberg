package osmodels

import (
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/ethereum/go-ethereum/common"
)

type EventType string

const (
	ItemListed          EventType = "item_listed"
	ItemSold            EventType = "item_sold"
	ItemReceivedBid     EventType = "item_received_bid"
	ItemReceivedOffer   EventType = "item_received_offer"
	ItemMetadataUpdated EventType = "item_metadata_updated"

	CollectionOffer EventType = "collection_offer"

	// ItemCancelled       EventType = "item_cancelled".
	// ItemTransferred     EventType = "item_transferred".

	StreamAPIEndpoint string = "wss://stream.openseabeta.com/socket"
)

var TxType = map[EventType]totra.TxType{
	ItemListed:      totra.Listing,
	ItemSold:        totra.Sale,
	CollectionOffer: totra.CollectionOffer,
	ItemReceivedBid: totra.ItemBid,
}

type ItemEvent interface {
	StreamEventType() EventType
}

type BaseStreamMessage struct {
	StreamEvent EventType `json:"event_type" mapstructure:"event_type"`
	SentAt      string    `json:"sent_at" mapstructure:"sent_at"`
}

func (m *BaseStreamMessage) StreamEventType() EventType {
	return m.StreamEvent
}

type BaseItemMetadataType struct {
	Name         string `json:"name" mapstructure:"name"`
	ImageURL     string `json:"image_url" mapstructure:"image_url"`
	AnimationURL string `json:"animation_url" mapstructure:"animation_url"`
	MetadataURL  string `json:"metadata_url" mapstructure:"metadata_url"`
}

type BaseItemType struct {
	NftID     string               `json:"nft_id" mapstructure:"nft_id"`
	Permalink string               `json:"permalink" mapstructure:"permalink"`
	Metadata  BaseItemMetadataType `json:"metadata" mapstructure:"metadata"`
	Chain     Chain                `json:"chain" mapstructure:"chain"`
}

type PayloadItemAndColl struct {
	Item       BaseItemType   `json:"item" mapstructure:"item"`
	Collection CollectionSlug `json:"collection" mapstructure:"collection"`
}

type CollectionSlug struct {
	CollectionName string `json:"collection_name" mapstructure:"collection_name"`
	CollectionSlug string `json:"collection_slug" mapstructure:"collection_slug"`
}

type Chain struct {
	Name string `json:"name" mapstructure:"name"`
}

// type ItemEvent struct {
// 	BaseStreamMessage `json:"base_stream_message" mapstructure:",squash"`
// 	Payload           ItemEventPayload `json:"payload" mapstructure:"payload"`
// }

// func (e ItemEvent) GetNftID() []string {
// 	return strings.Split(e.Payload.Item.NftID, "/")
// }

// func (e ItemEvent) ContractAddress() common.Address {
// 	return common.HexToAddress(e.GetNftID()[1])
// }

type ItemEventPayload struct {
	PayloadItemAndColl `json:"payload_item_and_coll" mapstructure:",squash"`
}

type ItemListedEvent struct {
	BaseStreamMessage `json:"base_stream_message" mapstructure:",squash"`
	// Payload           ItemEventPayload `json:"payload" mapstructure:"payload"`
	Payload ItemListedEventPayload `json:"payload" mapstructure:"payload"`
}

func (e *ItemListedEvent) GetNftID() []string {
	return strings.Split(e.Payload.Item.NftID, "/")
}

func (e *ItemListedEvent) ContractAddress() common.Address {
	return common.HexToAddress(e.GetNftID()[1])
}

type ItemListedEventPayload struct {
	PayloadItemAndColl `json:"payload_item_and_coll" mapstructure:",squash"`
	Quantity           int          `json:"quantity" mapstructure:"quantity"`
	ListingType        string       `json:"listing_type" mapstructure:"listing_type"`
	ListingDate        string       `json:"listing_date" mapstructure:"listing_date"`
	ExpirationDate     string       `json:"expiration_date" mapstructure:"expiration_date"`
	Maker              Account      `json:"maker" mapstructure:"maker"`
	Taker              Account      `json:"taker" mapstructure:"taker"`
	BasePrice          string       `json:"base_price" mapstructure:"base_price"`
	PaymentToken       PaymentToken `json:"payment_token" mapstructure:"payment_token"`
	IsPrivate          bool         `json:"is_private" mapstructure:"is_private"`
	EventTimestamp     string       `json:"event_timestamp" mapstructure:"event_timestamp"`
}

type ItemReceivedOfferEvent struct {
	BaseStreamMessage `json:"base_stream_message" mapstructure:",squash"`
	Payload           ItemReceivedOfferEventPayload `json:"payload" mapstructure:"payload"`
}

type ItemReceivedOfferEventPayload struct {
	PayloadItemAndColl `json:"payload_item_and_coll" mapstructure:",squash"`
	Quantity           int          `json:"quantity" mapstructure:"quantity"`
	CreatedDate        string       `json:"created_date" mapstructure:"created_date"`
	ExpirationDate     string       `json:"expiration_date" mapstructure:"expiration_date"`
	Maker              Account      `json:"maker" mapstructure:"maker"`
	Taker              Account      `json:"taker" mapstructure:"taker"`
	BasePrice          string       `json:"base_price" mapstructure:"base_price"`
	PaymentToken       PaymentToken `json:"payment_token" mapstructure:"payment_token"`
	EventTimestamp     string       `json:"event_timestamp" mapstructure:"event_timestamp"`
}

type CollectionOfferEvent struct {
	BaseStreamMessage `json:"base_stream_message" mapstructure:",squash"`
	Payload           CollectionOfferPayload `json:"payload" mapstructure:"payload"`
}

func (co CollectionOfferEvent) NftID() []string {
	return nil
}

func (co CollectionOfferEvent) ContractAddress() common.Address {
	return common.HexToAddress(co.Payload.AssetContractCriteria.Address)
}

type CollectionOfferPayload struct {
	AssetContractCriteria struct {
		Address string `json:"address"`
	} `json:"asset_contract_criteria" mapstructure:"asset_contract_criteria"`
	BasePrice  string `json:"base_price" mapstructure:"base_price"`
	Collection struct {
		Slug string `json:"slug"`
	} `json:"collection"`
	CollectionCriteria struct {
		Slug string `json:"slug"`
	} `json:"collection_criteria"`
	CreatedDate    string              `json:"created_date"`
	EventTimestamp time.Time           `json:"event_timestamp"`
	ExpirationDate string              `json:"expiration_date"`
	Maker          Account             `json:"maker"`
	OrderHash      string              `json:"order_hash"`
	PaymentToken   PaymentToken        `json:"payment_token" mapstructure:"payment_token"`
	ProtocolData   SeaportProtocolData `json:"protocol_data" mapstructure:"protocol_data"`
	Quantity       int                 `json:"quantity"`
	Taker          any                 `json:"taker"`
}

type Account struct {
	Address string `json:"address" mapstructure:"address"`
	User    string `json:"user" mapstructure:"user"`
}

type PaymentToken struct {
	ID       int    `json:"id" mapstructure:"id"`
	Symbol   string `json:"symbol" mapstructure:"symbol"`
	Address  string `json:"address" mapstructure:"address"`
	ImageURL string `json:"image_url" mapstructure:"image_url"`
	Name     string `json:"name" mapstructure:"name"`
	Decimals int    `json:"decimals" mapstructure:"decimals"`
	EthPrice string `json:"eth_price" mapstructure:"eth_price"`
	UsdPrice string `json:"usd_price" mapstructure:"usd_price"`
}

// bid

type ItemReceivedBidEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemReceivedBidEventPayload `mapstructure:"payload"`
}

type ItemReceivedBidEventPayload struct {
	PayloadItemAndColl `mapstructure:",squash"`
	Quantity           int          `mapstructure:"quantity"`
	CreatedDate        string       `mapstructure:"created_date"`
	ExpirationDate     string       `mapstructure:"expiration_date"`
	Maker              Account      `mapstructure:"maker"`
	Taker              Account      `mapstructure:"taker"`
	BasePrice          string       `mapstructure:"base_price"`
	PaymentToken       PaymentToken `mapstructure:"payment_token"`
	EventTimestamp     string       `mapstructure:"event_timestamp"`
	Collection         struct {
		Slug string `json:"slug"`
	} `json:"collection"`
}

// sale

type ItemSoldEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemSoldEventPayload `mapstructure:"payload"`
}

type ItemSoldEventPayload struct {
	PayloadItemAndColl `mapstructure:",squash"`
	ListingType        string       `mapstructure:"listing_type"`
	ClosingDate        string       `mapstructure:"closing_date"`
	Transaction        Transaction  `mapstructure:"transaction"`
	Maker              Account      `mapstructure:"maker"`
	Taker              Account      `mapstructure:"taker"`
	SalePrice          string       `mapstructure:"sale_price"`
	PaymentToken       PaymentToken `mapstructure:"payment_token"`
	IsPrivate          bool         `mapstructure:"is_private"`
	EventTimestamp     string       `mapstructure:"event_timestamp"`
}

// transfer

type ItemTransferredEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemTransferredEventPayload `mapstructure:"payload"`
}
type ItemTransferredEventPayload struct {
	PayloadItemAndColl `mapstructure:",squash"`
	FromAccount        Account     `mapstructure:"from_account"`
	Quantity           int         `mapstructure:"quantity"`
	ToAccount          Account     `mapstructure:"to_account"`
	Transaction        Transaction `mapstructure:"transaction"`
	EventTimestamp     string      `mapstructure:"event_timestamp"`
}

// cancel

type ItemCancelledEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemCancelledEventPayload `mapstructure:"payload"`
}
type ItemCancelledEventPayload struct {
	PayloadItemAndColl `mapstructure:",squash"`
	Quantity           int          `mapstructure:"quantity"`
	ListingType        string       `mapstructure:"listing_type"`
	Transaction        Transaction  `mapstructure:"transaction"`
	PaymentToken       PaymentToken `mapstructure:"payment_token"`
	EventTimestamp     string       `mapstructure:"event_timestamp"`
}

// metadata update

type ItemMetadataUpdateEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemMetadataUpdatePayload `mapstructure:"payload"`
}

type ItemMetadataUpdatePayload struct {
	PayloadItemAndColl   `mapstructure:",squash"`
	BaseItemMetadataType `mapstructure:",squash"`
	Description          string  `mapstructure:"description"`
	BackgroundColor      string  `mapstructure:"background_color"`
	Traits               []Trait `mapstructure:"traits"`
}

type Transaction struct {
	Hash      string `mapstructure:"hash"`
	Timestamp string `mapstructure:"timestamp"`
}

type Trait struct {
	TraitType   interface{} `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType interface{} `json:"display_type"`
	MaxValue    interface{} `json:"max_value"`
	TraitCount  interface{} `json:"trait_count"`
	Order       interface{} `json:"order"`
}

type AssetContract struct {
	Collection                  OSCollection `json:"collection"`
	Address                     string       `json:"address"`
	AssetContractType           string       `json:"asset_contract_type"`
	CreatedDate                 string       `json:"created_date"`
	Name                        string       `json:"name"`
	NftVersion                  string       `json:"nft_version"`
	OpenseaVersion              any          `json:"opensea_version"`
	Owner                       int          `json:"owner"`
	SchemaName                  string       `json:"schema_name"`
	Symbol                      string       `json:"symbol"`
	TotalSupply                 any          `json:"total_supply"`
	Description                 string       `json:"description"`
	ExternalLink                string       `json:"external_link"`
	ImageURL                    string       `json:"image_url"`
	DefaultToFiat               bool         `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int          `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int          `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool         `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  interface{}  `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints interface{}  `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int          `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int          `json:"seller_fee_basis_points"`
	PayoutAddress               string       `json:"payout_address"`
}

type OSCollection struct {
	BannerImageURL          string `json:"banner_image_url"`
	ChatURL                 any    `json:"chat_url"`
	CreatedDate             string `json:"created_date"`
	DefaultToFiat           bool   `json:"default_to_fiat"`
	Description             string `json:"description"`
	DevBuyerFeeBasisPoints  string `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints string `json:"dev_seller_fee_basis_points"`
	DiscordURL              string `json:"discord_url"`
	DisplayData             struct {
		CardDisplayStyle string `json:"card_display_style"`
	} `json:"display_data"`
	ExternalURL                 string      `json:"external_url"`
	Featured                    bool        `json:"featured"`
	FeaturedImageURL            any         `json:"featured_image_url"`
	Hidden                      bool        `json:"hidden"`
	SafelistRequestStatus       string      `json:"safelist_request_status"`
	ImageURL                    string      `json:"image_url"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
	LargeImageURL               any         `json:"large_image_url"`
	MediumUsername              any         `json:"medium_username"`
	Name                        string      `json:"name"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  interface{} `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints interface{} `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string      `json:"payout_address"`
	RequireEmail                bool        `json:"require_email"`
	ShortDescription            any         `json:"short_description"`
	Slug                        string      `json:"slug"`
	TelegramURL                 any         `json:"telegram_url"`
	TwitterUsername             any         `json:"twitter_username"`
	InstagramUsername           any         `json:"instagram_username"`
	WikiURL                     any         `json:"wiki_url"`
}

type OpenSeaListingsResponse struct {
	Next     interface{}    `json:"next"`
	Previous interface{}    `json:"previous"`
	Orders   []SeaportOrder `json:"orders"`
}

type OpenSeaAssetsResponse struct {
	Assets   []SeaportAsset `json:"assets"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
}

type SeaportUser struct {
	Username string `json:"username"`
}

type SeaportAccount struct {
	User          int    `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}

type SeaportOwner struct {
	User          SeaportUser `json:"user"`
	ProfileImgURL string      `json:"profile_img_url"`
	Address       string      `json:"address"`
	Config        string      `json:"config"`
}

type SeaportMakerFees struct {
	Account     SeaportAccount `json:"account"`
	BasisPoints string         `json:"basis_points"`
}

type SeaportConsiderationItem struct {
	ItemType             int            `json:"itemType"`
	Token                string         `json:"token"`
	IdentifierOrCriteria string         `json:"identifierOrCriteria"`
	StartAmount          string         `json:"startAmount"`
	EndAmount            string         `json:"endAmount"`
	Recipient            common.Address `json:"recipient"`
}

type SeaportOffer struct {
	ItemType             int            `json:"itemType"`
	Token                common.Address `json:"token"`
	IdentifierOrCriteria string         `json:"identifierOrCriteria"`
	StartAmount          string         `json:"startAmount"`
	EndAmount            string         `json:"endAmount"`
}

type SeaportParameters struct {
	Offerer       common.Address             `json:"offerer"`
	Offer         []SeaportOffer             `json:"offer"`
	Consideration []SeaportConsiderationItem `json:"consideration"`
	StartTime     string                     `json:"startTime"`
	EndTime       string                     `json:"endTime"`
	OrderType     int                        `json:"orderType"`
	// Zone                            common.Address             `json:"zone"`
	// ZoneHash common.Hash `json:"zoneHash"`
	Salt string `json:"salt"`
	// ConduitKey                      common.Hash                `json:"conduitKey"`
	TotalOriginalConsiderationItems int         `json:"totalOriginalConsiderationItems"`
	Counter                         interface{} `json:"counter"`
}

type SeaportDisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}

type SeaportProtocolData struct {
	Parameters SeaportParameters `json:"parameters"`
	Signature  string            `json:"signature"`
}

type SeaportFees struct {
	SellerFees  map[string]int `json:"seller_fees"`
	OpenseaFees map[string]int `json:"opensea_fees"`
}

type SeaportAssetContract struct {
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  interface{} `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       interface{} `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                interface{} `json:"external_link"`
	ImageURL                    interface{} `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  interface{} `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints interface{} `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               interface{} `json:"payout_address"`
}

type SeaportCollection struct {
	BannerImageURL              string             `json:"banner_image_url"`
	ChatURL                     interface{}        `json:"chat_url"`
	CreatedDate                 time.Time          `json:"created_date"`
	DefaultToFiat               bool               `json:"default_to_fiat"`
	Description                 string             `json:"description"`
	DevBuyerFeeBasisPoints      string             `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string             `json:"dev_seller_fee_basis_points"`
	DiscordURL                  string             `json:"discord_url"`
	DisplayData                 SeaportDisplayData `json:"display_data"`
	ExternalURL                 string             `json:"external_url"`
	Featured                    bool               `json:"featured"`
	FeaturedImageURL            string             `json:"featured_image_url"`
	Hidden                      bool               `json:"hidden"`
	SafelistRequestStatus       string             `json:"safelist_request_status"`
	ImageURL                    string             `json:"image_url"`
	IsSubjectToWhitelist        bool               `json:"is_subject_to_whitelist"`
	LargeImageURL               string             `json:"large_image_url"`
	MediumUsername              interface{}        `json:"medium_username"`
	Name                        string             `json:"name"`
	OnlyProxiedTransfers        bool               `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  interface{}        `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints interface{}        `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string             `json:"payout_address"`
	RequireEmail                bool               `json:"require_email"`
	ShortDescription            interface{}        `json:"short_description"`
	Slug                        string             `json:"slug"`
	TelegramURL                 interface{}        `json:"telegram_url"`
	TwitterUsername             interface{}        `json:"twitter_username"`
	InstagramUsername           interface{}        `json:"instagram_username"`
	WikiURL                     interface{}        `json:"wiki_url"`
	IsNsfw                      bool               `json:"is_nsfw"`
	Fees                        SeaportFees        `json:"fees"`
	IsRarityEnabled             bool               `json:"is_rarity_enabled"`
}

type SeaportAssetBundle struct {
	Assets            []SeaportAsset       `json:"assets"`
	Maker             interface{}          `json:"maker"`
	Slug              interface{}          `json:"slug"`
	Name              string               `json:"name"`
	Description       interface{}          `json:"description"`
	ExternalLink      interface{}          `json:"external_link"`
	AssetContract     SeaportAssetContract `json:"asset_contract"`
	Permalink         string               `json:"permalink"`
	SeaportSellOrders interface{}          `json:"seaport_sell_orders"`
}

type SeaportAsset struct {
	ID                   int                  `json:"id"`
	NumSales             int                  `json:"num_sales"`
	BackgroundColor      interface{}          `json:"background_color"`
	ImageURL             string               `json:"image_url"`
	ImagePreviewURL      string               `json:"image_preview_url"`
	ImageThumbnailURL    string               `json:"image_thumbnail_url"`
	ImageOriginalURL     string               `json:"image_original_url"`
	AnimationURL         interface{}          `json:"animation_url"`
	AnimationOriginalURL interface{}          `json:"animation_original_url"`
	Name                 string               `json:"name"`
	Description          string               `json:"description"`
	ExternalLink         interface{}          `json:"external_link"`
	AssetContract        SeaportAssetContract `json:"asset_contract"`
	Permalink            string               `json:"permalink"`
	Collection           SeaportCollection    `json:"collection"`
	Decimals             interface{}          `json:"decimals"`
	TokenMetadata        string               `json:"token_metadata"`
	IsNsfw               bool                 `json:"is_nsfw"`
	Owner                SeaportOwner         `json:"owner"`
	TokenID              string               `json:"token_id"`
}

type OpenSeaAsset struct {
	Owner          string   `json:"owner"`
	TokenIDs       []string `json:"token_ids"`
	CollectionSlug string   `json:"collection_slug"`
}

type SeaportOrder struct {
	CreatedDate      string              `json:"created_date"`
	ClosingDate      string              `json:"closing_date"`
	ListingTime      int                 `json:"listing_time"`
	ExpirationTime   int                 `json:"expiration_time"`
	OrderHash        string              `json:"order_hash"`
	ProtocolData     SeaportProtocolData `json:"protocol_data"`
	ProtocolAddress  string              `json:"protocol_address"`
	Maker            SeaportAccount      `json:"maker"`
	Taker            interface{}         `json:"taker"`
	CurrentPrice     string              `json:"current_price"`
	MakerFees        []SeaportMakerFees  `json:"maker_fees"`
	TakerFees        []interface{}       `json:"taker_fees"`
	Side             string              `json:"side"`
	OrderType        string              `json:"order_type"`
	Cancelled        bool                `json:"cancelled"`
	Finalized        bool                `json:"finalized"`
	MarkedInvalid    bool                `json:"marked_invalid"`
	ClientSignature  string              `json:"client_signature"`
	RelayID          string              `json:"relay_id"`
	CriteriaProof    interface{}         `json:"criteria_proof"`
	MakerAssetBundle SeaportAssetBundle  `json:"maker_asset_bundle"`
	TakerAssetBundle SeaportAssetBundle  `json:"taker_asset_bundle"`
}

type AssetCollection struct {
	Editors                 []string        `json:"editors"`
	PaymentTokens           []PaymentToken  `json:"payment_tokens"`
	PrimaryAssetContracts   []AssetContract `json:"primary_asset_contracts"`
	Traits                  struct{}        `json:"traits"`
	Stats                   CollectionStats `json:"stats"`
	BannerImageURL          string          `json:"banner_image_url"`
	ChatURL                 any             `json:"chat_url"`
	CreatedDate             string          `json:"created_date"`
	DefaultToFiat           bool            `json:"default_to_fiat"`
	Description             string          `json:"description"`
	DevBuyerFeeBasisPoints  string          `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints string          `json:"dev_seller_fee_basis_points"`
	DiscordURL              string          `json:"discord_url"`
	DisplayData             struct {
		CardDisplayStyle string `json:"card_display_style"`
	} `json:"display_data"`
	ExternalURL                 string      `json:"external_url"`
	Featured                    bool        `json:"featured"`
	FeaturedImageURL            any         `json:"featured_image_url"`
	Hidden                      bool        `json:"hidden"`
	SafelistRequestStatus       string      `json:"safelist_request_status"`
	ImageURL                    string      `json:"image_url"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
	LargeImageURL               any         `json:"large_image_url"`
	MediumUsername              any         `json:"medium_username"`
	Name                        string      `json:"name"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  interface{} `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints interface{} `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               any         `json:"payout_address"`
	RequireEmail                bool        `json:"require_email"`
	ShortDescription            any         `json:"short_description"`
	Slug                        string      `json:"slug"`
	TelegramURL                 any         `json:"telegram_url"`
	TwitterUsername             any         `json:"twitter_username"`
	InstagramUsername           any         `json:"instagram_username"`
	WikiURL                     any         `json:"wiki_url"`
	OwnedAssetCount             int         `json:"owned_asset_count"`
}

type CollectionStatsResponse struct {
	Stats *CollectionStats `json:"stats"`
}

type CollectionStats struct {
	OneDayVolume          float64 `json:"one_day_volume"`
	OneDayChange          float64 `json:"one_day_change"`
	OneDaySales           float64 `json:"one_day_sales"`
	OneDayAveragePrice    float64 `json:"one_day_average_price"`
	SevenDayVolume        float64 `json:"seven_day_volume"`
	SevenDayChange        float64 `json:"seven_day_change"`
	SevenDaySales         float64 `json:"seven_day_sales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume"`
	ThirtyDayChange       float64 `json:"thirty_day_change"`
	ThirtyDaySales        float64 `json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
	TotalVolume           float64 `json:"total_volume"`
	TotalSales            float64 `json:"total_sales"`
	TotalSupply           float64 `json:"total_supply"`
	Count                 float64 `json:"count"`
	NumOwners             float64 `json:"num_owners"`
	AveragePrice          float64 `json:"average_price"`
	NumReports            float64 `json:"num_reports"`
	MarketCap             float64 `json:"market_cap"`
	FloorPrice            float64 `json:"floor_price"`
}
