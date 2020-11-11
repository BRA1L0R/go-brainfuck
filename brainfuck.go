package main

import (
	"os"
	"io/ioutil"
)

func main() {
	// Load the program as an array of characters
	instructions, _ := ioutil.ReadFile(os.Args[1]) // Get the first available arguments that must be the file name that needs to be run

	// Optimizations
	bracketOffsetMap := preloadBrackets(instructions) // Preloads the offsets in the code for each parentheses

	// Execution loop
	interpretate(instructions, bracketOffsetMap)
}