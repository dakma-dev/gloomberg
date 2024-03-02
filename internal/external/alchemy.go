package external

type GetFloorPriceAlchemyResponse struct {
	Opensea   FloorPriceAlchemyData `json:"openSea"`
	Looksrare FloorPriceAlchemyData `json:"looksRare"`
}

type FloorPriceAlchemyData struct {
	FloorPrice    float64 `json:"floorPrice"`
	PriceCurrency string  `json:"priceCurrency"`
	CollectionURL string  `json:"collectionUrl"`
	RetrievedAt   string  `json:"retrievedAt"`
	Error         string  `json:"error"`
}
