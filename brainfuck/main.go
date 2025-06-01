// This solution is created by AI. just to study it. 

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	program := os.Args[1]
	memory := make([]int, 30000)
	ptr := 0
	pc := 0
	bracketPairs := make(map[int]int)

	// First pass: precompute all bracket matches
	tempStack := make([]int, 0)
	for i, c := range program {
		switch c {
		case '[':
			tempStack = append(tempStack, i)
		case ']':
			if len(tempStack) == 0 {
				return // Unmatched ]
			}
			start := tempStack[len(tempStack)-1]
			tempStack = tempStack[:len(tempStack)-1]
			bracketPairs[start] = i
			bracketPairs[i] = start
		}
	}

	// Second pass: execution
	for pc < len(program) {
		cmd := program[pc]
		switch cmd {
		case '>':
			ptr++
			if ptr >= len(memory) {
				ptr = 0
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = len(memory) - 1
			}
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			z01.PrintRune(rune(memory[ptr]))
		case '[':
			if memory[ptr] == 0 {
				pc = bracketPairs[pc]
			}
		case ']':
			if memory[ptr] != 0 {
				pc = bracketPairs[pc]
			}
		}
		// All other characters are ignored
		pc++
	}
}