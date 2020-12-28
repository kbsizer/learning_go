package module01

import "fmt"

// var digits = [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	decValue := 0
	lastNdx := len(value) - 1
	for ndx, rune := range value {
		digitValue := valueOf(rune)
		decValue += digitValue
		// would be slick to get rid of this special case "if"
		// maybe remove the last character or use a regular for loop
		// and stop before the last character?
		if ndx != lastNdx {
			decValue *= base
		}
	}
	return decValue
}

func valueOf(r rune) int {
	for ndx, rune := range digits {
		if r == rune {
			return ndx
		}
	}
	panic(fmt.Sprintf("Unexpected digit: %v", r))
}
