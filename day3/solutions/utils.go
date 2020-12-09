package solutions

import (
	"io/ioutil"
	"strings"
)

const tree = true
const free = false

const treeChar = "#"

// true = tree; false = free
func readInputFile(fileName string) [][]bool {
	rawPuzzleInput, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	inputLines := strings.Split(string(rawPuzzleInput), "\n")

	if len(inputLines) < 1 {
		panic("puzzle input is empty")
	}

	inputLen := len(inputLines)
	input := make([][]bool, inputLen)

	for lineIndex, line := range inputLines {
		lineSlice := make([]bool, len(line))
		for charIndex, char := range inputLines[lineIndex] {
			lineSlice[charIndex] = string(char) == treeChar
		}
		input[lineIndex] = lineSlice
	}

	return input
}
