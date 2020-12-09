package solutions

import (
	"fmt"
	"regexp"
	"strconv"
)

var fourDigitsRegex = regexp.MustCompile(`^\d{4}$`)
var eyeColorRegex = regexp.MustCompile(`^#[0-9 a-f]{6}$`)
var passportIDRegex = regexp.MustCompile(`^\d{9}$`)

var validEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

const (
	BirthYear      = "byr"
	IssueYear      = "iyr"
	ExpirationYear = "eyr"
	Height         = "hgt"
	HairColor      = "hcl"
	EyeColor       = "ecl"
	PassportID     = "pid"
)

func Puzzle2(puzzleInputFile string) {
	dataSets := readInputFile(puzzleInputFile)

	countValidPassports := 0
	for _, dataSet := range dataSets {
		if isValidDataset(dataSet) {
			countValidPassports++
		}
	}

	fmt.Printf("puzzle 2 -- count valid passorts: %d\n", countValidPassports)
}

func isValidDataset(data map[string]string) bool {
	return hasValidBirthYear(data) &&
		hasValidIssueYear(data) &&
		hasValidExpirationYear(data) &&
		hasValidHeight(data) &&
		hasValidHairColor(data) &&
		hasValidEyeColor(data) &&
		hasValidPassportID(data)
}

func hasValidBirthYear(data map[string]string) bool {
	if year, ok := data[BirthYear]; ok {
		return isFourDigitInRange(year, 1920, 2002)
	}
	return false
}

func hasValidIssueYear(data map[string]string) bool {
	if year, ok := data[IssueYear]; ok {
		return isFourDigitInRange(year, 2010, 2020)
	}
	return false
}

func hasValidExpirationYear(data map[string]string) bool {
	if year, ok := data[ExpirationYear]; ok {
		return isFourDigitInRange(year, 2020, 2030)
	}
	return false
}

func hasValidHeight(data map[string]string) bool {
	if heightWithUnit, ok := data[Height]; ok {
		if len(heightWithUnit) > 2 {
			unit := heightWithUnit[len(heightWithUnit)-2:]
			heightStr := heightWithUnit[:len(heightWithUnit)-2]
			height, err := strconv.Atoi(heightStr)
			if err != nil {
				return false
			}
			switch unit {
			case "in":
				return isInRange(height, 59, 76)
			case "cm":
				return isInRange(height, 150, 193)
			default:
				return false
			}
		}
	}
	return false
}

func hasValidHairColor(data map[string]string) bool {
	if hairColor, ok := data[HairColor]; ok {
		return eyeColorRegex.MatchString(hairColor)
	}
	return false
}

func hasValidEyeColor(data map[string]string) bool {
	if eyeColor, ok := data[EyeColor]; ok {
		for _, e := range validEyeColors {
			if eyeColor == e {
				return true
			}
		}
	}
	return false
}

func hasValidPassportID(data map[string]string) bool {
	if passportID, ok := data[PassportID]; ok {
		return passportIDRegex.MatchString(passportID)
	}
	return false
}

func isFourDigitInRange(suspect string, min, max int) bool {
	if fourDigitsRegex.MatchString(suspect) {
		y, _ := strconv.Atoi(suspect)
		return isInRange(y, min, max)
	}
	return false
}

func isInRange(suspect, min, max int) bool {
	return suspect >= min && suspect <= max
}
