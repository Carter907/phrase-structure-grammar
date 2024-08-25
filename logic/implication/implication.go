package implication

import (
	"fmt"

	"github.com/Carter907/phrase-structure-grammar/symbol"
)

type Implication struct {
	Antecedent symbol.Symbol
	Consequent []symbol.Symbol
}

func (i Implication) String() string {

	return fmt.Sprintf("(%s -> %s)", i.Antecedent, i.Consequent)
}
