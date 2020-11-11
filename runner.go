package main

import (
	"fmt"
	"os"
)

func interpretate(instructions []byte, bracketOffsetMap map[int]int) {
	// Initializing the fuckMemory
	memPtr := int64(0)						// Empty memory pointer, initialized as uint64, starts from 0
	mem := make(map[int64]uint8) // Creating the map of the memory

	// Initializing the fuckPointer
	instructionPointer := 0

	for ;; {		
		// Switch through all the possible instructions
		switch instructions[instructionPointer] { // NOTE: Switching through single runes of the whole file
		case '.': // Print the rune of the memory cell where the memory pointer is pointing right now
			fmt.Print(string(mem[memPtr]))
			break
		case ',':
			b := make([]byte, 1)
			os.Stdin.Read(b)

			mem[memPtr] = b[0]
			break
		case '+': 			// Add to the current pointer value
			mem[memPtr]++	// From the map select the current element and add to itself
			break
		case '-': 			// Remove from the current pointer value
			mem[memPtr]--	// Subtract from itself
			break
		case '>': 	// Increment mem pointer
			memPtr++ 	// Increments the memory pointer by one, switching to the next cell
			break
		case '<': 	// Decrement mem pointer
			memPtr-- 	// Decrements the memory pointer by one, switching to the previous cell
			break
		case '[': 	// Jump to the address after the corrisponding closing bracket if the value is zero
			if mem[memPtr] != 0 { 
				break		// If the value is zero continue executing normally
			}
			instructionPointer = bracketOffsetMap[instructionPointer]
			break
		case ']': // Jump to the address after the corrisponding opening bracket if the value is not zero
			if mem[memPtr] == 0 {
				break // If the value is zero then continue executing normally 
			}
			instructionPointer = bracketOffsetMap[instructionPointer]
			break
		}

		instructionPointer ++ // Increase the instruction pointer after each instruction
		if instructionPointer >= len(instructions) { // Check if the instruction pointer is exceeding the number of available instructions
			break
		}
	}
}