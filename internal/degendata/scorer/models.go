package scorer

import (
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	ContractAddress common.Address
	TokenID         int64
	Metadata        []TokenAttribute
}

type TokenAttribute struct {
	Name  string
	Value interface{}
}

type CollectionAttribute struct {
	Attribute   TokenAttribute
	TotalTokens int
}

// Collection.
type Collection struct {
	Name                      string
	Tokens                    []*Token
	AttributesFrequencyCounts map[string]map[string]int
}

func NewCollection(name string, tokens []*Token) *Collection {
	collection := &Collection{
		Name:   name,
		Tokens: tokens,
	}

	collection.AttributesFrequencyCounts = collection.ExtractCollectionAttributeFrequencyCounts()

	// "trait countify"
	for _, token := range collection.Tokens {
		token.Metadata = append(token.Metadata, TokenAttribute{
			Name:  "meta_trait:trait_count",
			Value: len(token.Metadata),
		})
	}

	return collection
}

func (collection *Collection) ExtractNullAttributes() map[string]*CollectionAttribute {
	nullAttributes := make(map[string]*CollectionAttribute)
	for _, token := range collection.Tokens {
		for _, attribute := range token.Metadata {
			if attribute.Value == nil {
				if nullAttributes[attribute.Name] == nil {
					nullAttributes[attribute.Name] = &CollectionAttribute{
						Attribute:   attribute,
						TotalTokens: 0,
					}
				}

				nullAttributes[attribute.Name].TotalTokens++
			}
		}
	}

	return nullAttributes
}

func (collection *Collection) ExtractCollectionAttributes() map[string][]CollectionAttribute {
	collectionTraits := make(map[string][]CollectionAttribute, 0)

	for name, attributeValues := range collection.AttributesFrequencyCounts {
		for attributeValue, count := range attributeValues {
			collectionTraits[name] = append(collectionTraits[name], CollectionAttribute{
				Attribute: TokenAttribute{
					Name:  name,
					Value: attributeValue,
				},
				TotalTokens: count,
			})
		}
	}

	return collectionTraits
}

func (collection *Collection) ExtractCollectionAttributeFrequencyCounts() map[string]map[string]int {
	collectionAttributes := make(map[string]map[string]int)
	for _, token := range collection.Tokens {
		for _, attribute := range token.Metadata {
			if attribute.Value == nil {
				continue
			}
			if collectionAttributes[attribute.Name] == nil {
				collectionAttributes[attribute.Name] = make(map[string]int)
			}

			attributeValue, ok := attribute.Value.(string)
			if !ok {
				continue
			}

			collectionAttributes[attribute.Name][attributeValue]++
		}
	}

	return collectionAttributes
}

func (collection *Collection) ExtractCollectionAttributeEntropy() map[string]float64 {
	collectionAttributes := collection.ExtractCollectionAttributeFrequencyCounts()
	collectionAttributesEntropy := make(map[string]float64)
	for key := range collectionAttributes {
		collectionAttributesEntropy[key] = 0.0 // entropy(value)
	}

	return collectionAttributesEntropy
}

func (collection *Collection) DeriveNormalizedAttributesFrequencyCounts() map[string]map[string]int {
	attrFrequencyCounts := make(map[string]map[string]int, 0)

	// for token in self._tokens:
	//     for (
	//         attr_name,
	//         str_attr,
	//     ) in token.metadata.string_attributes.items():
	//         normalized_name = normalize_attribute_string(attr_name)
	//         if str_attr.value not in attrs_freq_counts[attr_name]:
	//             attrs_freq_counts[normalized_name][str_attr.value] = 1
	//         else:
	//             attrs_freq_counts[normalized_name][str_attr.value] += 1

	for _, token := range collection.Tokens {
		for _, attribute := range token.Metadata {
			if attribute.Value == nil {
				continue
			}

			value, ok := attribute.Value.(string)
			if !ok {
				continue
			}

			if attrFrequencyCounts[attribute.Name] == nil {
				attrFrequencyCounts[attribute.Name] = make(map[string]int)
			}

			attrFrequencyCounts[attribute.Name][value]++
		}
	}

	return attrFrequencyCounts
}

func (collection *Collection) TotalTokensWithAttribute(attribute TokenAttribute) int {
	attributeValue, ok := attribute.Value.(string)
	if !ok {
		return 0
	}

	return collection.AttributesFrequencyCounts[attribute.Name][attributeValue]
}
