package solutions

import "fmt"

func Puzzle2(puzzleInputFile string) {
	inputs := readInputFile(puzzleInputFile)
	countValid := 0
	for _, in := range inputs {
		if validatePuzzle2(in) {
			countValid++
		}
	}

	fmt.Printf("puzzle 2 -- count valid passwords: %d\n", countValid)
}

func validatePuzzle2(in input) bool {
	posMinMatch := len(in.password) >= in.policy.minCount && in.password[in.policy.minCount] == in.policy.char
	posMaxMatch := len(in.password) >= in.policy.maxCount && in.password[in.policy.maxCount] == in.policy.char

	/*
		posMin | posMax | result
		false  | false  | false
		true   | false  | true
		false  | true   | true
		true   | true   | false
	=> XOR
	*/

	return posMinMatch != posMaxMatch
}
