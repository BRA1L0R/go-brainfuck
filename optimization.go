package main

import "fmt"

func compressOperation(instructions []byte, instructionOffset int) []byte {
	compressableChar := instructions[instructionOffset]	// Get the char where the whole operation started, to get a reference on what to compress
	count := 0	// Initialize the counter that will tell us later on how many chars need to be compressed

	for ;; {
		if instructionOffset + count >= len(instructions) { break }	// If the count exceedes the number of available operations then break
		if instructions[instructionOffset + count] != compressableChar { break }	// If the next operation is the same char as before then break
		count ++	// Continue with the count
	}

	beforeSlice := instructions[:instructionOffset]					// Get a reference on what there's before the repeatable sequence
	afterSlice := instructions[instructionOffset + count:]	// Get a reference on what there's after the repeatable sequence

	newInstructions := make([]byte, 0)
	newInstructions = append(newInstructions, beforeSlice...) // Compose the repeatable sequence omitting the repeated characters
	newInstructions = append(newInstructions, compressableChar)
	newInstructions = append(newInstructions, []byte(fmt.Sprint(count))...)
	newInstructions = append(newInstructions, afterSlice...)

	// Return the instructions, but optimized
	return newInstructions
}

// Takes the raw instructions in and spits out an optimized version of the code
func optimizer(instructions []byte) []byte {
	instructionOffset := 0
	for ;; {
		switch instructions[instructionOffset] {
		case '-':
			instructions = compressOperation(instructions, instructionOffset) // Example input: ++++ Example output: +4 (me of the future, don't mess this up)
			break
		case '+':
			instructions = compressOperation(instructions, instructionOffset)
			break
		}
		
		instructionOffset ++;
		if instructionOffset >= len(instructions) { break }
	}

	return instructions
}

func calculateBracketOffset(instructions []byte, instructionPointer int) int {
	depthLevel := 0 	// Var used to store the depth level of the brackets, used for correct bracket matching

	currentBracket := instructions[instructionPointer]
	var oppositeBracket byte
	if currentBracket == '[' { oppositeBracket = ']' } else { oppositeBracket = '[' }

	for ;; {
		// Offset the position of the pointer depending on the type of bracket (closing or opening)
		if currentBracket == '[' { instructionPointer++ } else if currentBracket == ']' { instructionPointer-- }
		if currentBracket == instructions[instructionPointer] { depthLevel++ }	// If the bracket found is the same as the starting bracket increase the depth level
		if oppositeBracket == instructions[instructionPointer] { 
			if depthLevel != 0 { depthLevel-- } else { break }	// If no depth level is left, then return the value of the current pointer
		}
	}

	return instructionPointer
}

// Preloading brackets offsetsw and point directions helps with code efficiency 
// and reduces the single operation time
func preloadBrackets(instructions []byte) map[int]int {
	offsetMap := make(map[int]int)	// Make a map of the various possible point locations of the brackets
	for originOffset, instruction := range instructions {	// Cicle through all the characters in the program looking for brackets
		if instruction == '[' || instruction == ']' { 
			offsetMap[originOffset] = calculateBracketOffset(instructions, originOffset) 	// Calculate the possible offset taking care of bracket depth
		}
	}

	return offsetMap
}