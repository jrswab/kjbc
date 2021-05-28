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
		find := "kjbc [-l] find word - to find the entry number for an English word"
		getEntry := "kjbc [-l] get number - to return the information for a given entry number."
		moreHelp := "Run 'kjbc -help' for options."

		msg := fmt.Sprintf("%s\n%s\n\n%s", find, getEntry, moreHelp)
		fmt.Println(msg)
		return
	}

	comm := os.Args[1]
	term := os.Args[2]

	if len(os.Args) == 5 {
		comm = os.Args[3]
		term = os.Args[4]
	}

	if comm == "find" {
		data, err := cons.Find(lang, term)
		if err != nil {
			fmt.Println(err)
		}

		for _, v := range *data {
			fmt.Printf("Entry Number: %s\n", v.Number)
			if showVerses {
				fmt.Printf("In Verses: ")
				for _, w := range v.Verses {
					fmt.Printf("%s, ", w)
				}
			}
		}
	}

	if comm == "get" {
		output, err := dicts.Get(lang, term)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)
	}
}
