package main

import (
	"flag"
	"fmt"
	"os"

	"swab.dev/kjbc/cons"
	"swab.dev/kjbc/dicts"
)

var (
	lang              string
	usage, showVerses bool
)

func init() {
	flag.StringVar(&lang, "l", "greek", "The target language to search within. Options are \"hebrew\" or \"greek\".")
	flag.BoolVar(&showVerses, "v", false, "Display the verses that correspond with the concordance entry number.")
	flag.BoolVar(&usage, "u", false, "Display the usage information.")
}

func getEntryNumber(term string) {
	data, err := cons.Find(lang, term)
	if err != nil {
		fmt.Println(err)
	}

	// Output word, entry, and definition when only one entry is listed
	if len(data) == 1 {
		for _, v := range data {
			fmt.Printf("One entry found for \"%s\"...\n", term)
			fmt.Printf("Entry Number: %s\n", v.Number)
			getDefinition(v.Number)
		}
	}

	fmt.Printf("\"%s\" contains more than one entry number.\n", term)
	fmt.Println("Use -v to see associated verses for each entry number.")
	for _, v := range data {
		fmt.Printf("\nEntry Number: %s\n", v.Number)

		if showVerses {
			fmt.Printf("In Verses: ")
			for j, w := range v.Verses {
				if j == len(v.Verses)-1 {
					fmt.Printf("%s\n", w)
				} else {
					fmt.Printf("%s, ", w)
				}
			}
		}
	}
}

func getDefinition(term string) {
	output, err := dicts.Get(lang, term)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}

func main() {
	flag.Parse()
	if usage || len(os.Args) <= 2 {
		word := "kjbc [options] word {word} - to find the entry number for an English word."
		oneWord := "  If only one entry exists for a word, kjbc will return the entry definition as well."
		getEntry := "kjbc [options] entry {number} - to return the information for a given entry number."
		moreHelp := "Run 'kjbc -help' for options."

		msg := fmt.Sprintf("%s\n%s\n\n%s\n\n%s", word, oneWord, getEntry, moreHelp)
		fmt.Println(msg)
		return
	}

	comm := os.Args[len(os.Args)-2]
	term := os.Args[len(os.Args)-1]

	if comm == "word" {
		getEntryNumber(term)
	}

	if comm == "entry" {
		getDefinition(term)
	}
}
