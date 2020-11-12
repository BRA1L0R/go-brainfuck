package main

import (
	"log"
	"flag"
	"os"
	"io/ioutil"
)

var (
	optimizationEnabled bool
)

// Get the flags before program execution
func init() {
	flag.BoolVar(&optimizationEnabled, "o", false, "Optimizes the code before the execution (SPERIMENTAL!)")
	flag.Parse() // Parse the flag and assign the variables
}

func main() {
	// Load the program as an array of characters
	instructions, err := ioutil.ReadFile(os.Args[len(os.Args) - 1]) // Get the last available arguments that must be the file name that needs to be run
	if err != nil {
		log.Fatal(err)
	}

	// Optimizations
	if optimizationEnabled {
		instructions = optimizer(instructions)
	}

	// Preloads the offsets in the code for each parentheses
	bracketOffsetMap := preloadBrackets(instructions)

	// Execution loop
	interpretate(instructions, bracketOffsetMap)
}