package main

import (
	"log"
	"os"

	_ "sample/matchers"
	"sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var input string
	input = "网红"
	//input = "president"
	search.Run(input)
}
