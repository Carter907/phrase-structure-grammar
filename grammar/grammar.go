package grammar

import (
	"github.com/Carter907/phrase-structure-grammar/logic/implication"
	"github.com/Carter907/phrase-structure-grammar/symbol"
)

type PSGrammar struct {
	V Vocabulary
	S symbol.Symbol
	P Production
}

// words that cannot be replaced
// {a, the, rabbit, mathematician, hops, eats, quickly, wildly}
type Terminals struct {
	Symbols []symbol.Symbol
}

// symbols resembling labels for your syntax
// can be replaced
// {sentence, noun phrase, verb phrase, adjective, article, noun, verb, adverb}
type NonTerminals struct {
	Symbols []symbol.Symbol
}

type Vocabulary struct {
	Terminals
	NonTerminals
}

type Production struct {
	implications []implication.Implication
}
