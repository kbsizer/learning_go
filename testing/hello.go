package main

import (
	"fmt"
	"math"
)

const helloText = "Hello, world"

func main() {
	fmt.Printf("%s\n", GetHello())
	fmt.Printf("The hypotenuse of a right triangle with legs 10 feet long is %f feet.\n", Hypotenuse(10, 10))
	fmt.Printf("The average of 75, 82 and 98 is %f.\n", Average(75, 82, 98))
}

// GetHello returns the content of helloText.
func GetHello() string {
	return helloText
}

// Hypotenuse computes the length of the hypotenuse of a
// right triangle having sides of the given lengths.
func Hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

// Average returns the average of a list of numbers.
func Average(values ...float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

func newPrivateFuncThatIsNotTested() {
	fmt.Println("No one called me from a test.")
}
