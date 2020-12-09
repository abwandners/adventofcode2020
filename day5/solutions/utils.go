package solutions

import (
	"io/ioutil"
	"strings"
)

func readInputFile(fileName string) []string {
	rawInput, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(rawInput), "\n")
}
