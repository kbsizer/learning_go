package main

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

// Adding methods to the Stack struct

// Push adds a string to the end(top) of the stack
//
// Note: method declaration syntax is:
//       func (refName [*]ParentStruct) methodName(args) [returnType(s)]
//
// Note: append() and len() are frequently-used built-in functions
//       for working with slices.
func (stack *Stack) Push(s string) {
	stack.data = append(stack.data, s)
}

// Pop removes a string from the end(top) of the stack
//
// Note: no parens are needed around the return type types if there is only one
func (stack *Stack) Pop() string {
	topNdx := len(stack.data) - 1
	if topNdx > 0 {
		panic("Attempt to pop from empty stack")
	}
	topElement := stack.data[topNdx] // get last(top) element
	stack.data[topNdx] = ""          // Java equivalent to setting to null
	stack.data = stack.data[:topNdx] // shrink slice
	return topElement
}
