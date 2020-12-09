package sulutions

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readPuzzleInput(fileName string) []int {
	puzzleRawInput, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	puzzleInputStrings := strings.Split(string(puzzleRawInput), "\n")

	puzzleInput := make([]int, len(puzzleInputStrings))
	for i, str := range puzzleInputStrings {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		puzzleInput[i] = number
	}

	return puzzleInput
}
