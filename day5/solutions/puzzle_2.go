package solutions

import "fmt"

func Puzzle2(puzzleInputFile string) {
	seatDescriptions := readInputFile(puzzleInputFile)
	highestSeatNumber := 0
	for _, seatDescription := range seatDescriptions {
		rowDescription := rowDescriptionFromSeatDescription(seatDescription)
		placeDescription := placeDescriptionFromSeatDescription(seatDescription)

		row := calcRow(rowDescription)
		place := calcPlace(placeDescription)

		seatNumber := calcSeatNumber(row, place)

	}

	fmt.Printf("puzzle 1 -- highest seat number: %d", highestSeatNumber)
}

