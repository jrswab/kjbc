package dicts

import (
	"encoding/json"
	"fmt"

	lex "github.com/jrswab/kjbc/lexicons"
)

type Dictionary map[string]Entry

type Entry struct {
	Word       string `json:"word,omitempty"`
	Pronounce  string `json:"pronounce,omitempty"`
	Definition string `json:"definition,omitempty"`
}

func marshalGreek() Dictionary {
	greekDict := make(Dictionary)
	json.Unmarshal([]byte(lex.GreekDict), &greekDict)
	return greekDict
}

func marshalHebrew() Dictionary {
	hebrewDict := make(Dictionary)
	json.Unmarshal([]byte(lex.HebrewDict), &hebrewDict)
	return hebrewDict
}

func Get(lang, num string) (Entry, error) {
	var dict Dictionary
	if lang == "greek" {
		dict = marshalGreek()
	}

	if lang == "hebrew" {
		dict = marshalHebrew()
	}

	if _, ok := dict[num]; !ok {
		return dict[num], fmt.Errorf("the entry number %s does not exist in the %s dictionary", num, lang)
	}
	return dict[num], nil
}
