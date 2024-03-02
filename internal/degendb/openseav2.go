package degendb

type NFTsResponse struct {
	NFTs []NFT `json:"nfts"`
}

type NFT struct {
	Identifier    string `json:"identifier"`
	Collection    string `json:"collection"`
	Contract      string `json:"contract"`
	TokenStandard string `json:"token_standard"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ImageURL      string `json:"image_url"`
	AnimationURL  string `json:"animation_url"`
	MetadataURL   string `json:"metadata_url"`
	OpenseaURL    string `json:"opensea_url"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	IsDisabled    bool   `json:"is_disabled"`
	IsNsfw        bool   `json:"is_nsfw"`
}

type NFTDetailed struct {
	NFT
	IsSuspicious bool   `json:"is_suspicious"`
	Creator      string `json:"creator"`
	Traits       []struct {
		TraitType   string      `json:"trait_type"`
		DisplayType interface{} `json:"display_type"`
		MaxValue    interface{} `json:"max_value"`
		TraitCount  int         `json:"trait_count"`
		Order       interface{} `json:"order"`
		Value       string      `json:"value"`
	} `json:"traits"`
	Owners []struct {
		Address  string `json:"address"`
		Quantity int    `json:"quantity"`
	} `json:"owners"`
	Rarity struct {
		StrategyID      interface{} `json:"strategy_id"`
		StrategyVersion interface{} `json:"strategy_version"`
		Rank            int         `json:"rank"`
		Score           interface{} `json:"score"`
		CalculatedAt    string      `json:"calculated_at"`
		MaxRank         interface{} `json:"max_rank"`
		TokensScored    int         `json:"tokens_scored"`
		RankingFeatures interface{} `json:"ranking_features"`
	} `json:"rarity"`
}

type ContractInfo struct {
	Address          string `json:"address"`
	Chain            string `json:"chain"`
	Collection       string `json:"collection"`
	ContractStandard string `json:"contract_standard"`
	Name             string `json:"name"`
	Supply           int    `json:"supply"`
}
