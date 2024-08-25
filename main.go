package main

import (
	"fmt"

	"github.com/Carter907/phrase-structure-grammar/grammar"
)

func main() {

	gram := grammar.NewGrammar()
	test := "the beastly bear eats impatiently"
	res := gram.IsInLanguage(test)
	fmt.Println(res)
}
