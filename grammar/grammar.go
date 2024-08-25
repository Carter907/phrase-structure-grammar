package grammar

import (
	"fmt"

	"github.com/Carter907/phrase-structure-grammar/symbol"
)

type PSGrammar struct {
	V Vocabulary
	S symbol.Symbol
	P Productions
}

func NewGrammar() PSGrammar {

	return PSGrammar{
		V: NewVocab(),
		S: "sentence",
		P: NewProductions(),
	}
}

func (grammar *PSGrammar) IsInLanguage(src string) bool {

	for _, prod := range grammar.P.implications {

		fmt.Println(prod)
	}

	return false
}
