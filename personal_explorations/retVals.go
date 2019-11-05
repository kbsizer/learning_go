package main

import "fmt"

// Return values in Go:
//    	- Functions and Methods may return zero, one or N values
//		- Function and Method invocations must either accept all return values or ignore all return values
//		- An underscore may be used to accept-but-throw-away a return value
//
// Examples follow.
func main() {
	fmt.Println("retaining no values from function that returns one value")
	oneRetVal()

	fmt.Println("retaining one value from function that returns one value")
	retVal1 := oneRetVal()
	fmt.Println("oneRetVal() returned", retVal1)

	fmt.Println("retaining no values from function that returns two value")
	twoRetVals()

	fmt.Println("retaining 1st value from function that returns two values")
	retVal2, _ := twoRetVals()
	fmt.Println("twoRetVals() returned", retVal2)

	fmt.Println("retaining both values from function that returns two values")
	retVal2_1, retVal2_2 := twoRetVals()
	fmt.Println("twoRetVals() returned", retVal2_1, retVal2_2)

	fmt.Println("retaining 2nd value from function that returns two values")
	_, retVal2_2 = twoRetVals()
	fmt.Println("twoRetVals() returned", retVal2_2, "as second return value")

	fmt.Println("retaining 1st value from function that returns three values")
	retVal3, _, _ := threeRetVals()
	fmt.Println("threeRetVals() returned", retVal3)

	// RANGE with UNICODE CHARACTERS
	msg := "abc日本語123"
	for i, ch := range msg {
		fmt.Printf("Character %#U starts at byte position %d\n", ch, i)
	}
}

func oneRetVal() string {
	return "1foo"
}

func twoRetVals() (string, string) {
	return "2foo", "2bar"
}

func threeRetVals() (string, string, string) {
	return "3foo", "3bar", "3baz"
}
