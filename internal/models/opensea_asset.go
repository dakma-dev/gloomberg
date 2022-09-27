package models

type AssetsResponse struct {
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Assets   []Asset     `json:"assets"`
}

type User struct {
	Username string `json:"username"`
}

type Profile struct {
	User          User   `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}

type Fees struct {
	SellerFees  interface{} `json:"seller_fees"`
	OpenseaFees interface{} `json:"opensea_fees"`
}

type LastSale struct {
	Asset struct {
		Decimals interface{} `json:"decimals"`
		TokenID  string      `json:"token_id"`
	} `json:"asset"`
	AssetBundle    interface{}  `json:"asset_bundle"`
	EventType      string       `json:"event_type"`
	EventTimestamp string       `json:"event_timestamp"`
	AuctionType    interface{}  `json:"auction_type"`
	TotalPrice     string       `json:"total_price"`
	PaymentToken   PaymentToken `json:"payment_token"`
	Transaction    interface{}  `json:"transaction"`
	CreatedDate    string       `json:"created_date"`
	Quantity       string       `json:"quantity"`
}

type Asset struct {
	ID                      int           `json:"id"`
	AnimationOriginalURL    string        `json:"animation_original_url"`
	AnimationURL            string        `json:"animation_url"`
	AssetContract           AssetContract `json:"asset_contract"`
	BackgroundColor         interface{}   `json:"background_color"`
	Collection              Collection    `json:"collection"`
	Creator                 Profile       `json:"creator"`
	Decimals                interface{}   `json:"decimals"`
	Description             string        `json:"description"`
	ExternalLink            string        `json:"external_link"`
	ImageOriginalURL        string        `json:"image_original_url"`
	ImagePreviewURL         string        `json:"image_preview_url"`
	ImageThumbnailURL       string        `json:"image_thumbnail_url"`
	ImageURL                string        `json:"image_url"`
	IsNsfw                  bool          `json:"is_nsfw"`
	IsPresale               bool          `json:"is_presale"`
	LastSale                LastSale      `json:"last_sale"`
	ListingDate             interface{}   `json:"listing_date"`
	Name                    string        `json:"name"`
	NumSales                int           `json:"num_sales"`
	Owner                   Profile       `json:"owner"`
	Permalink               string        `json:"permalink"`
	RarityData              interface{}   `json:"rarity_data"`
	SeaportSellOrders       interface{}   `json:"seaport_sell_orders"`
	SupportsWyvern          bool          `json:"supports_wyvern"`
	TokenID                 string        `json:"token_id"`
	TokenMetadata           string        `json:"token_metadata"`
	TopBid                  interface{}   `json:"top_bid"`
	Traits                  []Trait       `json:"traits"`
	TransferFee             interface{}   `json:"transfer_fee"`
	TransferFeePaymentToken interface{}   `json:"transfer_fee_payment_token"`
}

type Collection struct {
	BannerImageURL          string      `json:"banner_image_url"`
	ChatURL                 interface{} `json:"chat_url"`
	CreatedDate             string      `json:"created_date"`
	DefaultToFiat           bool        `json:"default_to_fiat"`
	Description             string      `json:"description"`
	DevBuyerFeeBasisPoints  string      `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints string      `json:"dev_seller_fee_basis_points"`
	DiscordURL              interface{} `json:"discord_url"`
	DisplayData             struct {
		CardDisplayStyle string `json:"card_display_style"`
	} `json:"display_data"`
	ExternalURL                 interface{}     `json:"external_url"`
	Featured                    bool            `json:"featured"`
	FeaturedImageURL            string          `json:"featured_image_url"`
	Fees                        Fees            `json:"fees"`
	Hidden                      bool            `json:"hidden"`
	ImageURL                    string          `json:"image_url"`
	InstagramUsername           interface{}     `json:"instagram_username"`
	IsNsfw                      bool            `json:"is_nsfw"`
	IsRarityEnabled             bool            `json:"is_rarity_enabled"`
	IsSubjectToWhitelist        bool            `json:"is_subject_to_whitelist"`
	LargeImageURL               string          `json:"large_image_url"`
	MediumUsername              interface{}     `json:"medium_username"`
	Name                        string          `json:"name"`
	OnlyProxiedTransfers        bool            `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string          `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string          `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               interface{}     `json:"payout_address"`
	RequireEmail                bool            `json:"require_email"`
	SafelistRequestStatus       string          `json:"safelist_request_status"`
	ShortDescription            interface{}     `json:"short_description"`
	Slug                        string          `json:"slug"`
	Stats                       CollectionStats `json:"stats"`
	TelegramURL                 interface{}     `json:"telegram_url"`
	TwitterUsername             interface{}     `json:"twitter_username"`
	WikiURL                     interface{}     `json:"wiki_url"`
}
