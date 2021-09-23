package main

import "fmt"

// main function is the entrypoint for a standalone application.
//
// Note: "func" keyword introduces both functions and methods.
func main() {
	fmt.Println("Hello, World!")
	s := Stack{}
	fmt.Printf(" Type of s is %T\n", s)
	fmt.Printf("Value of s is %v\n", s)
}

// Stack is a Go structure.  The closest thing we get to a class.
//
// Define a new struct using the keyword "type" (roughly analogous
// to "class" in Java).
//
// Note: no semicolons
//
// Note: struct or variable name PRECEDES its definition or type
type Stack struct {
	data []string
}
