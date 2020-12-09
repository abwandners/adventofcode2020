package solutions

import (
	"fmt"
)

const (
	lower = 'F'
	left  = 'L'
)

func Puzzle1(puzzleInputFile string) {
	seatDescriptions := readInputFile(puzzleInputFile)
	highestSeatNumber := 0
	for _, seatDescription := range seatDescriptions {
		rowDescription := rowDescriptionFromSeatDescription(seatDescription)
		placeDescription := placeDescriptionFromSeatDescription(seatDescription)

		row := calcRow(rowDescription)
		place := calcPlace(placeDescription)

		seatNumber := calcSeatNumber(row, place)
		if seatNumber > highestSeatNumber {
			highestSeatNumber = seatNumber
		}
	}

	fmt.Printf("puzzle 1 -- highest seat number: %d", highestSeatNumber)
}
