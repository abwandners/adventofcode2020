package solutions

import (
	"fmt"
	"strings"
)

func Puzzle1(puzzleInputFile string) {
	inputs := readInputFile(puzzleInputFile)
	countValid := 0
	for _, in := range inputs {
		if validatePuzzle1(in) {
			countValid++
		}
	}

	fmt.Printf("puzzle 1 -- count valid passwords: %d\n", countValid)
}

func validatePuzzle1(in input) bool {
	charCount := strings.Count(in.password, string(in.policy.char))
	if charCount < in.policy.minCount {
		return false
	}
	if charCount > in.policy.maxCount {
		return false
	}
	return true
}
