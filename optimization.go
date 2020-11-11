package main

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