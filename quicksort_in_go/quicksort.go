package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	intArray := make([]int, len(argsWithoutProg))
	for i, arg := range argsWithoutProg {
		element, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("ERROR: Unable to parse %s as an integer: %v\n", arg, err)
			os.Exit(2)
		}
		intArray[i] = element
	}

	fmt.Printf("Sorting %v\n...\n", intArray)
	sortedArray := QuickSort(intArray)
	fmt.Printf("result: %v\n", sortedArray)
}

// QuickSort demonstrates Go implementation of the recursive sort algorithm.
func QuickSort(arr []int) []int {
	newArr := make([]int, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	sort(newArr, 0, len(arr)-1)

	return newArr
}

func sort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	sort(arr, start, splitIndex-1)
	sort(arr, splitIndex+1, end)
}
