package grammar

type Vocabulary struct {
	Terminals
	NonTerminals
}

func NewVocab() Vocabulary {

	return Vocabulary{}
}
