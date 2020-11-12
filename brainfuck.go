package main

import (
	"flag"
	"os"
	"io/ioutil"
)

var (
	optimizationEnabled bool
)

func init() {
	flag.BoolVar(&optimizationEnabled, "o", false, "Optimizes the code before the execution (SPERIMENTAL!)")
	flag.Parse()
}

func main() {
	// Load the program as an array of characters
	instructions, _ := ioutil.ReadFile(os.Args[len(os.Args) - 1]) // Get the last available arguments that must be the file name that needs to be run

	// Optimizations
	if optimizationEnabled {
		instructions = optimizer(instructions)
	}

	bracketOffsetMap := preloadBrackets(instructions) // Preloads the offsets in the code for each parentheses

	// Execution loop
	interpretate(instructions, bracketOffsetMap)
}