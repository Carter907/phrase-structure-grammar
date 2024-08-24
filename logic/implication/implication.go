package implication

import (
	"github.com/Carter907/phrase-structure-grammar/symbol"
)

type Implication struct {
	antecedent symbol.Symbol
	cosequent  symbol.Symbol
}
