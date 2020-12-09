package solutions

import "fmt"

func Puzzle1(puzzleInputFile string) {
	countValid := 0
	dataSets := readInputFile(puzzleInputFile)
	for _, dataSet := range dataSets {
		if validate(dataSet) {
			countValid++
		}
	}

	fmt.Printf("puzzle 1 -- count valid passports: %d\n", countValid)
}

func validate(dataSet map[string]string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, requiredField := range requiredFields {
		if _, exists := dataSet[requiredField]; !exists {
			return false
		}
	}

	return true
}
