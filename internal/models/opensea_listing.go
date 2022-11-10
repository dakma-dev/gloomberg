package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type OpenSeaListingsResponse struct {
	Next     interface{}    `json:"next"`
	Previous interface{}    `json:"previous"`
	Orders   []SeaportOrder `json:"orders"`
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
	Offerer                         common.Address             `json:"offerer"`
	Offer                           []SeaportOffer             `json:"offer"`
	Consideration                   []SeaportConsiderationItem `json:"consideration"`
	StartTime                       string                     `json:"startTime"`
	EndTime                         string                     `json:"endTime"`
	OrderType                       int                        `json:"orderType"`
	Zone                            common.Address             `json:"zone"`
	ZoneHash                        common.Hash                `json:"zoneHash"`
	Salt                            string                     `json:"salt"`
	ConduitKey                      common.Hash                `json:"conduitKey"`
	TotalOriginalConsiderationItems int                        `json:"totalOriginalConsiderationItems"`
	Counter                         int                        `json:"counter"`
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
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
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
	OpenseaBuyerFeeBasisPoints  string             `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string             `json:"opensea_seller_fee_basis_points"`
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
