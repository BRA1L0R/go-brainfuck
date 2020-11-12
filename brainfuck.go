package main

import (
	"log"
	"flag"
	"os"
	"io/ioutil"
	"time"
)

var (
	optimizationEnabled bool
	enhancedDebugging bool
	memVisualizer bool
	instructionDelay time.Duration
)

// Get the flags before program execution
func init() {
	flag.BoolVar(&optimizationEnabled, "o", false, "Optimizes the code before the execution (SPERIMENTAL!)")
	flag.BoolVar(&enhancedDebugging, "d", false, "Debug your BrainFuck code with special characters")
	flag.BoolVar(&memVisualizer, "v", false, "Visualizes the memory after each instruction")
	flag.DurationVar(&instructionDelay, "t", 0, "Delays t milliseconds after each instruction")
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