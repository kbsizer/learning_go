package main

import "fmt"

func main() {
	fmt.Println("retaining no values from function that returns one value")
	oneRetVal()

	fmt.Println("retaining one value from function that returns one value")
	s1 := oneRetVal()
	fmt.Println("oneRetVal() returned", s1)

	fmt.Println("retaining 1st value from function that returns two values")
	s2, _ := twoRetVals()
	fmt.Println("twoRetVals() returned", s2)

	fmt.Println("retaining both values from function that returns two values")
	s2_1, s2_2 := twoRetVals()
	fmt.Println("twoRetVals() returned", s2_1, s2_2)

	fmt.Println("retaining 2nd value from function that returns two values")
	_, s2_2 = twoRetVals()
	fmt.Println("twoRetVals() returned", s2_2, "as second return value")

	fmt.Println("retaining 1st value from function that returns three values")
	s3, _, _ := threeRetVals()
	fmt.Println("threeRetVals() returned", s3)

	// RANGE

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
