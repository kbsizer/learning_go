package module02

import (
	"sort"
)

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
//
// NOTE: This is an "in place" implementation: Does not need to create a separate array.
func InsertionSortInt(list []int) {
	for ndx := 1; ndx < len(list); ndx++ {
		value := list[ndx]
		insertionPoint := 0
		for insertionPoint < ndx && value >= list[insertionPoint] {
			insertionPoint++
		}
		// found correct insersion point
		if insertionPoint == ndx {
			// degenerate case; correct position for value is the cell it is in
			continue
		}
		// right-shift elements from target cell (insersionPoint) to value's old home (ndx)
		for j := ndx; j > insertionPoint; j-- {
			list[j] = list[j-1]
		}
		list[insertionPoint] = value
	}
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
//
// NOTE: This implementation uses a separate, target array instead of "in place" sort.
func InsertionSortString(list []string) {
	sortedList := make([]string, 0, len(list))
	for _, valueToInsert := range list {
		ndx := len(sortedList)
		for i, sortedElement := range sortedList {
			//fmt.Println("i=", i, ", sortedElement=", sortedElement, ", valueToInsert=", valueToInsert)
			if valueToInsert < sortedElement {
				ndx = i
				break
			}
		}
		sortedList = insertAt(sortedList, valueToInsert, ndx)
	}
	copy(list, sortedList)
}

// insertAt illustrates the use of slices, the append() command, and the
//
// For more information, see: https://golang.org/ref/spec#Appending_and_copying_slices
func insertAt(list []string, value string, location int) []string {
	//fmt.Println("\tinsertAt(", list, ", ", value, ", ", location)
	list = append(list, "")
	copy(list[location+1:], list[location:])
	list[location] = value
	return list
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
}
