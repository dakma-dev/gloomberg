package scorer

type Scorer interface {
	// validate_collection(collection *Collection) error
	scoreToken(collection *Collection, token *Token) (float64, error)
	scoreCollection(collection *Collection, token *Token) (float64, error)
}
