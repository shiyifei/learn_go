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
	input = "童年"
	//input = "president"
	search.Run(input)
}
