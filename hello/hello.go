package main

import (
	"fmt"
	"math"
)

const helloText = "Hello, world"

// Timezone illustrates the use of "iota".
type Timezone int

const (
	// EST is Eastern Standard Time (GMT-5)
	EST Timezone = -(5 + iota) // iota: 0, EST: -5
	// CST is Central Standard Time (GMT-6)
	CST // iota: 1, CST: -6
	// MST is Mountain Standard Time (GMT-7)
	MST // iota: 2, MST: -7
	// PST is Pacific Standard Time (GMT-8)
	PST // iota: 3, MST: -8
)

func main() {
	fmt.Printf("%s\n", GetHello())
	fmt.Printf("The hypotenuse of a right triangle with legs 10 feet long is %f feet.\n", Hypotenuse(10, 10))
	fmt.Printf("The average of 75, 82 and 98 is %f.\n", Average(75, 82, 98))
	fmt.Printf("Timezones (%T)\n\t%v\n\t%v\n\t%v\n", EST, EST, CST, PST)
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
