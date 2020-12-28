package module01

// GCD returns the greatest common denominator using
// Euclidean algorithm
// Step 1: If B == 0, return A
// Step 2: A becomes B, and B becomes the remainder of dividing A by B
func GCD(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
