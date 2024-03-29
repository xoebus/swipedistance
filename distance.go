package swipedistance

import (
	"math"
	"strings"
)

type coordinate struct {
	X int
	Y int
}

var coords = map[rune]coordinate{
	'Q': {42, 52},
	'W': {122, 52},
	'E': {201, 52},
	'R': {277, 52},
	'T': {355, 52},
	'Y': {443, 52},
	'U': {512, 52},
	'I': {590, 52},
	'O': {668, 52},
	'P': {746, 52},
	'A': {82, 159},
	'S': {160, 159},
	'D': {239, 159},
	'F': {316, 159},
	'G': {395, 159},
	'H': {474, 159},
	'J': {552, 159},
	'K': {629, 159},
	'L': {729, 159},
	'Z': {227, 266},
	'X': {240, 266},
	'C': {317, 266},
	'V': {396, 266},
	'B': {473, 266},
	'N': {551, 266},
	'M': {629, 266},
}

func (c coordinate) distanceTo(other coordinate) float64 {
	xd := float64(other.X - c.X)
	yd := float64(other.Y - c.Y)

	xd2 := math.Pow(xd, 2)
	yd2 := math.Pow(yd, 2)

	return math.Sqrt(xd2 + yd2)
}

// Sum returns the swipe distance required to type the word. The result has no
// units associated with it. Results are not guaranteed to be comparable
// between different versions of the package. If there is a character in the
// word which cannot be typed then the second argument will be false.
func Sum(word string) (float64, bool) {
	if len(word) <= 1 {
		return 0, true
	}

	word = strings.ToUpper(word)

	var previous coordinate
	var sum float64
	for i, c := range word {
		next, found := coords[c]
		if !found {
			return 0, false
		}

		if i != 0 {
			sum += previous.distanceTo(next)
		}

		previous = next
	}

	return sum, true
}
