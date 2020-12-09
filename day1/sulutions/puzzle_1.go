package sulutions

import (
	"fmt"
)

func Puzzle1(puzzleInputFile string, expectedNumber int) {
	puzzleInput := readPuzzleInput(puzzleInputFile)
	number1, number2 := findMatchingOperators(puzzleInput, expectedNumber)

	fmt.Printf("puzzle 1 output -- number1: %d; number2: %d; restult: %d\n", number1, number2, number1 * number2)
}

func findMatchingOperators(input []int, search int) (int, int) {
	for _, op1 := range input {
		for _, op2 := range input {
			sum := op1 + op2
			if sum == search {
				return op1, op2
			}
		}
	}
	panic("combination not found")
}


