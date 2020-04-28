package module01

import "strings"

var digits = [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
// Implementation notes:
//   - Builder is much more efficient than concatenating (immutable) strings
//   - Builder doesn't have "prepend" method, so we build the result backward (least significant digit to most
//     significant digit), then reverse at the end
func DecToBase(dec, base int) string {
	var sb strings.Builder
	for dec > 0 {
		remainder := dec % base
		sb.WriteRune(digits[remainder])
		dec = dec / base
	}
	return reverse(sb.String())
}

// reverse returns a rune-by-rune reversal of the given string
func reverse(s string) string {
	// convers string to slice of runes
	runes := []rune(s)
	// initialize i to the start, j to the end, then swap runes
	// until the indices meet
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
