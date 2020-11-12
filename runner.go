package main

import (
	"time"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func compressedCount(instructions []byte) (uint8, int) { // instruction must be a slice starting with the first significative character
	offset := 0

	for ;; {
		if !unicode.IsDigit(rune(instructions[offset])) {	// Count all the characters that are actually digits and stop the first non-digit it finds
			break
		}
		offset++ // Increase the offset as it goes through all the numbers
		if offset >= len(instructions) { break }
	}

	characterCount, err := strconv.Atoi(string(instructions[:offset])) // Calculate the amount of characters that have been replaced with numbers
	if err != nil {
		log.Fatal(err)	// Throw error on parse error
	}

	return uint8(characterCount), offset	// Return the count and the next instruction offset
}

// Main code execution loop
func interpretate(instructions []byte, bracketOffsetMap map[int]int) {
	// Initializing the fuckMemory
	memPtr := int64(0)          // Empty memory pointer, initialized as uint64, starts from 0
	mem := make([]uint8, 30000) // 30000 is the cell limit of the original brainfuck compiler

	// Initializing the instruction offset
	instructionOffset := 0

	for {
		// Switch through all the possible instructions
		switch instructions[instructionOffset] { // NOTE: Switching through single runes of the whole file
		case '.': // Print the rune of the memory cell where the memory pointer is pointing right now
			fmt.Print(string(mem[memPtr]))
			break
		case ',':
			b := make([]byte, 1)
			os.Stdin.Read(b)
			mem[memPtr] = b[0]
			break
		case '+': // Add to the current pointer value
			if optimizationEnabled {
				charCount, offset := compressedCount(instructions[instructionOffset+1:])
				instructionOffset += offset
				mem[memPtr] += charCount // From the map select the current element and add to itself
			} else { mem[memPtr]++ }
			break
		case '-': // Remove from the current pointer value
		if optimizationEnabled {
			charCount, offset := compressedCount(instructions[instructionOffset+1:])
			instructionOffset += offset
			mem[memPtr] -= charCount // From the map select the current element and add to itself
		} else { mem[memPtr]-- }
			break
		case '>': // Increment mem pointer
			memPtr++ // Increments the memory pointer by one, switching to the next cell
			break
		case '<': // Decrement mem pointer
			memPtr-- // Decrements the memory pointer by one, switching to the previous cell
			break
		case '[': // Jump to the address after the corrisponding closing bracket if the value is zero
			if mem[memPtr] != 0 {
				break // If the value is zero continue executing normally
			}
			instructionOffset = bracketOffsetMap[instructionOffset]
			break
		case ']': // Jump to the address after the corrisponding opening bracket if the value is not zero
			if mem[memPtr] == 0 {
				break // If the value is zero then continue executing normally
			}
			instructionOffset = bracketOffsetMap[instructionOffset]
			break
		case '=':	// Visualizes the chunk of memory around the memory pointer
			if !enhancedDebugging { break } // If enhanced debug characters are disabled then just do nothing
			printMemory(mem, memPtr, false)	// Prints the memory
			break
		case '#':	// Visualizes the WHOLE memory
			if !enhancedDebugging { break }
			printMemory(mem, memPtr, true)	// Prints the WHOLE memory
			break
		}

		if memVisualizer {	// If the memory visualizer option is turned on, print the memory after each instruction
			printMemory(mem, memPtr, false)
		}

		if instructionDelay != 0 {
			time.Sleep(instructionDelay)
		}

		instructionOffset++                         // Increase the instruction pointer after each instruction
		if instructionOffset >= len(instructions) { // Check if the instruction pointer is exceeding the number of available instructions
			break
		}
	}
}

func printMemory(mem []uint8, memPtr int64, whole bool) {
	if whole { 
		fmt.Println(mem) 
	}	else {
		for i := int64(-4); i < 5; i++ {
			if memPtr + i < 0 { fmt.Print(sprintCell(0)); continue }
			fmt.Print(sprintCell(mem[memPtr + i]))
		}
		fmt.Print("\n\n")
	}
}

func sprintCell(val uint8) string{
	return fmt.Sprintf(" [%v] ", val)
}