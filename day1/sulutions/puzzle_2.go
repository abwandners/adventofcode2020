package sulutions

import (
	"fmt"
)

func Puzzle2(puzzleInputFile string, expectedNumber int) {
	puzzleInput := readPuzzleInput(puzzleInputFile)
	number1, number2, number3 := findMatchingOperatorTriple(puzzleInput, expectedNumber)

	fmt.Printf("puzzle 2 output -- number1: %d; number2: %d; number3: %d; restult: %d\n", number1, number2, number3, number1*number2*number3)
}

func findMatchingOperatorTriple(input []int, search int) (int, int, int) {
	for _, op1 := range input {
		for _, op2 := range input {
			for _, op3 := range input {
				sum := op1 + op2 + op3
				if sum == search {
					return op1, op2, op3
				}
			}
		}
	}
	panic("combination not found")
}
