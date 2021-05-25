package main

import "flag"

var lang string

func init() {
	flag.StringVar(&lang, "language", "greek", "The target language to search within. Options are Hebrew or Greek.")
	flag.StringVar(&lang, "l", "greek", "Shorthand for language.")
}

func main() {
	flag.Parse()
}
