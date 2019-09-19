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

	// Illustrate value vs reference semantics of for..range
	// In the first case, we make a copy of the array; in the
	// second case we do not
	sArray := [5]string{"Able", "Baker", "Charlie", "Dog", "Eagle"}

	fmt.Printf("Before loop #1, sArray[2] = %s\n", sArray[2])

	for i, s := range sArray {
		if i == 2 {
			sArray[i] = "CHANGED!"
			fmt.Printf("LOOP 1: sArray[2] = %v\n", s)
		}
	}

	fmt.Printf("After loop #1, sArray[2] = %s\n", sArray[2])

	sArray = [5]string{"Able", "Baker", "Charlie", "Dog", "Eagle"}

	for i := range sArray {
		if i == 2 {
			sArray[i] = "CHANGED!"
			fmt.Printf("LOOP 2: sArray[2] = %v\n", sArray[2])
		}
	}

	fmt.Printf("After loop #2, sArray[2] = %s\n", sArray[2])

}

// GetHello returns the content of helloText.
func GetHello() string {
	return helloText
}

// Hypotenuse computes the length of the hypotenuse of a
// right triangle having sides of the given lengths.
func Hypotenuse(a, b float64) float64 {
	//TEMP
	if a == b {
		fmt.Println("Inputs EQUAL")
	}
	//TEMP
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
