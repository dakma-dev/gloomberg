package manifold

import "github.com/ethereum/go-ethereum/common"

const APIBaseURL = "https://apps.api.manifoldxyz.com/public/instance/"

// DataResponse is the response from the Manifold API endpoint /data?id=1337...
type DataResponse struct {
	ID            int        `json:"id"`
	Creator       Creator    `json:"creator"`
	Slug          string     `json:"slug"`
	PublicData    PublicData `json:"publicData"`
	AppID         int64      `json:"appId"`
	MintPrice     float64    `json:"mintPrice"`
	IsOpenEdition bool       `json:"isOpenEdition"`
}

type Creator struct {
	ID         int         `json:"id"`
	Image      interface{} `json:"image"`
	Name       string      `json:"name"`
	TwitterURL string      `json:"twitterUrl"`

	// Creator wallet address
	Address common.Address `json:"address"`
}

type PublicData struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Image            string `json:"image"`
	Animation        string `json:"animation"`
	MediaIsLandscape bool   `json:"mediaIsLandscape"`

	ClaimType    string         `json:"claimType"`
	ClaimIndex   int            `json:"claimIndex"`
	MerkleTreeID int            `json:"merkleTreeId"`
	Erc20        common.Address `json:"erc20"`
	AudienceID   interface{}    `json:"audienceId"`

	// Chain ID
	Network int `json:"network"`
	// Manifold contract address
	ExtensionAddress common.Address `json:"extensionAddress"`
	// Collection contract address
	CreatorContractAddress common.Address `json:"creatorContractAddress"`
}
