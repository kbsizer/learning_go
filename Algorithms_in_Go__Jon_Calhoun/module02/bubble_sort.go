package module02

import (
	"sort"
)

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for {
		swapped := false
		for i := 1; i < len(list); i++ {
			if list[i-1] > list[i] {
				list[i], list[i-1] = list[i-1], list[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortString is a bubble sort for string slices.
//
// In GOLANG, the string is an immutable chain of arbitrary bytes encoded with UTF-8 encoding.
// You are allowed to compare strings with each other using two different ways:
//
// 1. Using comparison operators: Go strings support comparison operators, i.e, ==, !=, >=, <=, <, >.
//
// 2. Using Compare() method: You can also compare two strings using the built-in function Compare()
// provided by the strings package. This function returns an integer value after comparing two strings
// lexicographically. The return values are:
//    strings.Compare("A","A") == 0
//    strings.Compare("A","B") == -1
//    strings.Compare("B","A") == 1
// Note: Compare is included only for symmetry with package bytes. It is usually clearer and always faster
// to use the built-in string comparison operators ==, <, >, and so on.
func BubbleSortString(list []string) {
	for {
		swapped := false
		for i := 1; i < len(list); i++ {
			if list[i-1] > list[i] {
				list[i], list[i-1] = list[i-1], list[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
//
//    type Interface interface {
//       Len() int   // Len is the number of elements in the collection.
//       Less(i, j int) bool   // Less reports whether the element with index i should sort before the element with index j.
//       Swap(i, j int)   // Swap swaps the elements with indexes i and j.
//    }
func BubbleSort(list sort.Interface) {
	//fmt.Println("entering")
	for {
		swapped := false
		for i := 1; i < list.Len(); i++ {
			if list.Less(i, i-1) {
				//fmt.Println("Swapping", i, "and", i-1)
				list.Swap(i, i-1)
				swapped = true
			}
		}
		//fmt.Println("bottom of loop, swapped=", swapped)
		if !swapped {
			break
		}
	}
}
