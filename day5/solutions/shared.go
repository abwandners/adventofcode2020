package solutions


func calcRow(rowDescription string) int {
	lowerRow := 0
	upperRow := 127

	return partition(rowDescription, lowerRow, upperRow)
}

func calcPlace(placeDescription string) int {
	lowerSeat := 0
	upperSeat := 7

	return partition(placeDescription, lowerSeat, upperSeat)
}

func partition(designators string, low, high int) int {
	for _, designator := range designators {
		if designator == lower || designator == left {
			high = (high-low)/2 + low
		} else {
			low = (high-low)/2 + low + 1
		}
	}
	return low
}

func rowDescriptionFromSeatDescription(seatDescription string) string {
	return seatDescription[:7]
}

func placeDescriptionFromSeatDescription(seatDescription string) string {
	return seatDescription[7:]
}

func calcSeatNumber(row, place int) int {
	return row * 8 + place
}
