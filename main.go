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
	flag.StringVar(&lang, "l", "greek", "The target language to search within. Options are Hebrew or Greek.")
	flag.BoolVar(&showVerses, "v", false, "Display the verses that correspond with the concordance entry number.")
	flag.BoolVar(&usage, "u", false, "Display the usage information.")
}

func main() {
	flag.Parse()
	if usage || len(os.Args) <= 2 {
		word := "kjbc [-l] word {word} - to find the entry number for an English word"
		getEntry := "kjbc [-l] entry {number} - to return the information for a given entry number."
		moreHelp := "Run 'kjbc -help' for options."

		msg := fmt.Sprintf("%s\n%s\n\n%s", word, getEntry, moreHelp)
		fmt.Println(msg)
		return
	}

	comm := os.Args[len(os.Args)-2]
	term := os.Args[len(os.Args)-1]

	if comm == "word" {
		data, err := cons.Find(lang, term)
		if err != nil {
			fmt.Println(err)
		}

		for i, v := range *data {
			if i > 0 {
				fmt.Printf("\nEntry Number: %s\n", v.Number)
			} else {
				fmt.Printf("Entry Number: %s\n", v.Number)
			}

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

	if comm == "entry" {
		output, err := dicts.Get(lang, term)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)
	}
}
