package module02

import (
	"fmt"
	"testing"
)

func TestCallSemantics(t *testing.T) {
}

func TestSliceCopy(t *testing.T) {
	// Creating slices
	slc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	var slc2 []int
	slc3 := make([]int, 5)
	slc4 := []int{78, 50, 67, 77}

	// Before copying
	fmt.Println("Before any copy operations...")
	fmt.Println("\tSlice1:", slc1)
	fmt.Println("\tSlice2:", slc2)
	fmt.Println("\tSlice3:", slc3)
	fmt.Println("\tSlice4:", slc4)

	// Copying the slices
	copy1 := copy(slc2, slc1)
	fmt.Println("After executing copy(slice2, slice1)...")
	fmt.Println("\tSlice1:", slc1)
	fmt.Println("\tSlice2:", slc2)
	fmt.Println("\tcopy(slice2,slice1) returned:", copy1, "   (number of elements copied)")

	copy2 := copy(slc3, slc1)
	fmt.Println("After executing copy(slice3, slice1)...")
	fmt.Println("\tSlice3:", slc3)
	fmt.Println("\tTotal number of elements copied:", copy2)

	copy3 := copy(slc4, slc1)
	fmt.Println("\nSlice:", slc4)
	fmt.Println("Total number of elements copied:", copy3)
}

// For more info on variadic functions, see: https://golang.org/ref/spec#Passing_arguments_to_..._parameters
func TestVariadicFunctions(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println("a = ", a)
	// this is legal
	result := append(a, 4, 5, 6)
	fmt.Println("result = ", result)
	// and so is this
	b := []int{4, 5, 6}
	result = append(a, b...)
	fmt.Println("result = ", result)
	// but these are NOT
	// result = append(a, b)              // only first param is a slice, others must be elements
	// result = append(a, 7, b...)        // can only use triple-dot if there are no other elements
	// result = append(a, b..., 7 )       // can only use triple-dot if there are no other elements
}

func TestPrepend(t *testing.T) {
	a := []int{3, 4, 5}
	fmt.Println("a = ", a)
	// prepend a 2 at head of slice
	a = prepend(a, 2)
	fmt.Println("a = ", a)
}

// prepend illustrates (1) an idiomatic and efficient way of inserting a value at the start of a slice, and (2)
// the good practice of using small functions to isolate and clarify low-level implementation details
func prepend(a []int, b int) []int {
	a = append(a, 0)
	copy(a[1:], a)
	a[0] = b
	return a
}
