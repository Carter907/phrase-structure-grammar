package grammar

import "github.com/Carter907/phrase-structure-grammar/symbol"

// words that cannot be replaced
// {a, the, rabbit, mathematician, hops, eats, quickly, wildly}
type Terminals struct {
	Symbols []symbol.Symbol
}
