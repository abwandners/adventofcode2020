package solutions

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type input struct {
	policy policy
	password string
}

type policy struct {
	minCount int
	maxCount int
	char     byte
}

func readInputFile(fileName string) []input {
	rawPuzzleInput, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	inputLines := strings.Split(string(rawPuzzleInput), "\n")

	inputs := make([]input, len(inputLines))
	for i, line := range inputLines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			panic("input does not consist of two parts separated by ':'")
		}
		policyRaw := parts[0]
		password := parts[1]

		p := parsePolicy(policyRaw)
		inputs[i] = input{
			policy:   p,
			password: password,
		}
	}

	return inputs
}

func parsePolicy(policyRaw string) policy {
	parts := strings.Split(policyRaw, " ")
	if len(parts) != 2 {
		panic("policy string does not consist of two parts separated by ' '")
	}
	minMaxRaw := parts[0]
	policyChar := parts[1]

	minMaxParts := strings.Split(minMaxRaw, "-")
	if len(minMaxParts) != 2 {
		panic("min max part of policy string does not consist of two parts separated by '-'")
	}

	min, err := strconv.Atoi(minMaxParts[0])
	if err != nil {
		panic(fmt.Sprintf("unable to convert min part of value %s to int", minMaxParts[0]))
	}

	max, err := strconv.Atoi(minMaxParts[1])
	if err != nil {
		panic(fmt.Sprintf("unable to convert max part of value %s to int", minMaxParts[1]))
	}

	return policy{
		minCount: min,
		maxCount: max,
		char:     []byte(policyChar)[0],
	}
}
