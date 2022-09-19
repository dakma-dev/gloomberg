package models

type CollectionMetadata struct {
	ContractName string `json:"contractName"`
	Symbol       string `json:"symbol"`
	TotalSupply  uint64 `json:"total_supply"`
	TokenURI     string `json:"token_uri"`
}

type MetadataERC721 struct {
	Image string `json:"image"`
}
