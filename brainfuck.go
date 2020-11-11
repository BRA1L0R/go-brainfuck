package main

import (
	"flag"
	"io/ioutil"
)

// Declaring global variabled taken from the run flags
var (
	runFile string
	debugActive bool
)

func init() { 
	// Runs before anything, gets all the flags from the command line
	flag.StringVar(&runFile, "file", "", "Input file that needs to get interpreted")
	flag.BoolVar(&debugActive, "debug", false, "Activates debug messages")
	flag.Parse()
}

func calculateBracketOffset(instructions []byte, instructionPointer int) int {
	depthLevel := 0 	// Var used to store the depth level of the brackets, used for correct bracket matching

	currentBracket := instructions[instructionPointer]
	var oppositeBracket byte
	if currentBracket == '[' { oppositeBracket = ']' } else { oppositeBracket = '[' }

	for ;; {
		// Offset the position of the pointer depending on the type of bracket (closing or opening)
		if currentBracket == '[' { instructionPointer++ } else if currentBracket == ']' { instructionPointer-- }
		if currentBracket == instructions[instructionPointer] { depthLevel++ }
		if oppositeBracket == instructions[instructionPointer] { 
			if depthLevel != 0 { depthLevel-- } else { break }
		}
	}

	return instructionPointer
}

func preloadBrackets(instructions []byte) map[int]int {
	offsetMap := make(map[int]int)
	for originOffset, instruction := range instructions {
		if instruction == '[' || instruction == ']' { 
			offsetMap[originOffset] = calculateBracketOffset(instructions, originOffset) 
		}
	}

	return offsetMap
}

func main() {
	instructions, _ := ioutil.ReadFile(runFile) // Loading the fuckinstructions

	// Optimizations
	bracketOffsetMap := preloadBrackets(instructions) // Preloads the offsets in the code for each parentheses

	// Execution loop
	interpretate(instructions, bracketOffsetMap)
}

