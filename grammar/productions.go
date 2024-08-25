package grammar

import (
	"github.com/Carter907/phrase-structure-grammar/logic/implication"
	"github.com/Carter907/phrase-structure-grammar/symbol"
)

type Productions struct {
	implications []implication.Implication
}

func NewProductions() Productions {

	return Productions{
		implications: []implication.Implication{
			{
				Antecedent: "sentence",
				Consequent: []symbol.Symbol{"noun phrase", "verb phrase"},
			},
			{
				Antecedent: "noun phrase",
				Consequent: []symbol.Symbol{"article", "adjective", "noun"},
			},
			{
				Antecedent: "adjective",
				Consequent: []symbol.Symbol{"beastly"},
			},
			{
				Antecedent: "article",
				Consequent: []symbol.Symbol{"the"},
			},
			{
				Antecedent: "noun",
				Consequent: []symbol.Symbol{"bear"},
			},
			{
				Antecedent: "verb phrase",
				Consequent: []symbol.Symbol{"verb", "adverb"},
			},
			{
				Antecedent: "verb",
				Consequent: []symbol.Symbol{"eats"},
			},
			{
				Antecedent: "adverb",
				Consequent: []symbol.Symbol{"impatiently"},
			},
		},
	}
}

func FromImplicationSet(implications []implication.Implication) Productions {

	return Productions{implications}
}
