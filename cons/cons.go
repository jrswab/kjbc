package cons

import (
	"encoding/json"
	"fmt"
)

type Concordance struct {
	Entries []Entry `json:"entries,omitempty"`
}

type Entry struct {
	Word string      `json:"word,omitempty"`
	Refs []Reference `json:"reference,omitempty"`
}

type Reference struct {
	Number string   `json:"number,omitempty"`
	Verses []string `json:"verses,omitempty"`
}

func marshalGreek() *Concordance {
	greekCon := new(Concordance)
	json.Unmarshal([]byte(greek), &greekCon)
	return greekCon
}

func marshalHebrew() *Concordance {
	hebrewCon := new(Concordance)
	json.Unmarshal([]byte(hebrew), &hebrewCon)
	return hebrewCon
}

func Find(lang, word string) (*[]Reference, error) {
	var con *Concordance
	if lang == "greek" {
		con = marshalGreek()
	}

	if lang == "hebrew" {
		con = marshalHebrew()
	}

	var refs *[]Reference
	for _, v := range con.Entries {
		if v.Word == word {
			refs = &v.Refs
			break
		}
	}

	if refs == nil {
		return nil, fmt.Errorf("could not find '%s', in the concordance", word)
	}
	return refs, nil
}
