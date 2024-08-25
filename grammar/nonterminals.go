package grammar

import "github.com/Carter907/phrase-structure-grammar/symbol"

// symbols resembling labels for your syntax
// can be replaced
// {sentence, noun phrase, verb phrase, adjective, article, noun, verb, adverb}
type NonTerminals struct {
	Symbols []symbol.Symbol
}

func NewNonTerminals() NonTerminals {

	return NonTerminals{
		Symbols: []symbol.Symbol{
			"sentence",
			"noun phrase",
			"verb phrase",
			"adjectie",
			"article",
			"noun",
			"verb",
			"adverb",
		},
	}
}
