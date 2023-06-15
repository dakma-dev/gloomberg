package scorer

import (
	"errors"
	"math"
	"sort"

	"github.com/charmbracelet/log"
	"gonum.org/v1/gonum/floats"
)

type InformationContentScorer struct{}

func NewInformationContentScorer() *InformationContentScorer {
	return &InformationContentScorer{}
}

func (scorer *InformationContentScorer) ScoreCollection(collection *Collection) ([]float64, error) {
	scores := make([]float64, 0)

	collectionNullAttributes := collection.ExtractNullAttributes()

	collectionAttributes := collection.ExtractCollectionAttributes()
	log.Printf("collectionAttributes: %+v", collectionAttributes)

	// collection_entropy = self._get_collection_entropy(
	//     collection=collection,
	//     collection_attributes=collection_attributes,
	//     collection_null_attributes=collection_null_attributes,
	// )

	collectionEntropy := scorer.GetCollectionEntropy(collection, collectionAttributes, collectionNullAttributes)

	if collectionEntropy == 0.0 {
		log.Error("collectionEntropy is 0.0")

		return scores, errors.New("collectionEntropy is 0.0")
	}

	for _, token := range collection.Tokens {
		tokenScore, err := scorer.ScoreToken(collection, token, collectionNullAttributes, collectionEntropy)
		if err != nil {
			log.Errorf("Error scoring token: %+v", err)

			return scores, errors.New("error scoring token")
		}

		scores = append(scores, tokenScore)
	}

	return scores, nil
}

func (scorer *InformationContentScorer) ScoreToken(collection *Collection, token *Token, collectionNullAttributes map[string]*CollectionAttribute, collectionEntropy float64) (float64, error) {
	icTokenScore, err := scorer.GetICScore(collection, token, collectionNullAttributes)
	if err != nil {
		log.Errorf("Error getting IC score: %+v", err)

		return 0.0, errors.New("error getting IC score")
	}

	normalizedTokenScore := icTokenScore / collectionEntropy

	return normalizedTokenScore, nil
}

func (scorer *InformationContentScorer) GetICScore(collection *Collection, token *Token, collectionNullAttributes map[string]*CollectionAttribute) (float64, error) {
	attributeScores, _ := scorer.GetTokenAttributesScoresAndWeights(collection, token, false, collectionNullAttributes)

	log.Print("")
	log.Printf("attributeScores: %+v", attributeScores)

	reciprocals := make([]float64, 0)
	for _, score := range attributeScores {
		reciprocals = append(reciprocals, 1.0/score)
	}

	log.Printf("reciprocals: %+v", reciprocals)

	naturalLogs := make([]float64, 0)
	for _, reciprocal := range reciprocals {
		naturalLogs = append(naturalLogs, math.Log2(reciprocal))
	}

	log.Printf("naturalLogs: %+v", naturalLogs)

	sum := 0.0
	for _, naturalLog := range naturalLogs {
		sum += naturalLog
	}

	return -sum, nil
}

func (scorer *InformationContentScorer) GetTokenAttributesScoresAndWeights(collection *Collection, token *Token, _ bool, collectionNullAttributes map[string]*CollectionAttribute) ([]float64, []float64) {
	combinedAttributes := make(map[string]*CollectionAttribute, 0)

	if collectionNullAttributes == nil {
		collectionNullAttributes = collection.ExtractNullAttributes()
	}

	log.Print("")
	log.Printf("collectionNullAttributes: %+v", collectionNullAttributes)

	// collection attributes
	collectionAttributes := make(map[string]*CollectionAttribute, 0)
	for _, attribute := range token.Metadata {
		if attribute.Value == nil {
			continue
		}

		_, ok := attribute.Value.(string)
		if !ok {
			continue
		}

		collectionAttributes[attribute.Name] = &CollectionAttribute{
			Attribute:   attribute,
			TotalTokens: collection.TotalTokensWithAttribute(attribute),
		}
	}

	log.Printf("collectionAttributes: %+v", collectionAttributes)

	// combine the attributes from the token and the collection
	for key, value := range collectionNullAttributes {
		combinedAttributes[key] = value
	}

	for key, value := range collectionAttributes {
		combinedAttributes[key] = value
	}

	log.Printf("combinedAttributes: %#v", combinedAttributes)

	keys := make([]string, 0)
	for key := range combinedAttributes {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	attrWeights := make([]float64, len(keys))
	for i := range keys {
		attrWeights[i] = 1.0
	}

	sortedCombinedAttributes := make([]*CollectionAttribute, 0)
	for _, key := range keys {
		sortedCombinedAttributes = append(sortedCombinedAttributes, combinedAttributes[key])
	}

	totalSupply := len(collection.Tokens)

	scores := make([]float64, 0)
	for _, attr := range sortedCombinedAttributes {
		scores = append(scores, float64(totalSupply)/float64(attr.TotalTokens))
	}

	return scores, attrWeights
}

func (scorer *InformationContentScorer) GetCollectionEntropy(collection *Collection, collectionAttributes map[string][]CollectionAttribute, collectionNullAttributes map[string]*CollectionAttribute) float64 {
	attributes := collection.ExtractCollectionAttributes()
	log.Printf("e attributes: %+v", len(attributes))
	nullAttributes := collection.ExtractNullAttributes()
	log.Printf("e null_attributes: %+v", len(nullAttributes))
	log.Printf("e collectionNullAttributes: %+v", len(collectionNullAttributes))

	collectionProbabilities := make([]float64, 0)

	for _, attribute := range collectionAttributes {
		for _, attr := range attribute {
			log.Printf("attr: %+v", attr)

			collectionProbabilities = append(collectionProbabilities, float64(attr.TotalTokens)/float64(len(collection.Tokens)))
		}
	}

	log2edCollectionProbabilities := make([]float64, 0)
	for _, probability := range collectionProbabilities {
		log2edCollectionProbabilities = append(log2edCollectionProbabilities, math.Log2(probability))
	}

	dotProd := floats.Dot(collectionProbabilities, log2edCollectionProbabilities)

	log.Printf("dotProd: %+v", dotProd)

	return dotProd
}
