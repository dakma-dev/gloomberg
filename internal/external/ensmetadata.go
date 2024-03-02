package external

type ENSMetadataAttribute struct {
	TraitType   string      `json:"trait_type"`
	DisplayType string      `json:"display_type"`
	Value       interface{} `json:"value"`
}

// ENSMetadata defines model for ENSMetadata.
type ENSMetadata struct {
	Attributes      []ENSMetadataAttribute `json:"attributes"`
	BackgroundImage string                 `json:"background_image"`
	Description     string                 `json:"description"`
	ImageURL        string                 `json:"image_url"`
	Name            string                 `json:"name"`
	NameLength      int                    `json:"name_length"`
	SegmentLength   int                    `json:"segment_length"`
	URL             string                 `json:"url"`
	Version         int                    `json:"version"`
}
