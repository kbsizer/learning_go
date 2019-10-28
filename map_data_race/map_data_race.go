// Demonstrates:
//      1) Maps in Go are NOT inherently thread-safe.  Even when threads are never touching the
//         same key-value pair
//      2) Go's built-in data race detection code for map access
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var scoreMap = make(map[string]int)

func main() {
	// setup
	// GOMAXPROCS limits the number of operating system threads that can execute user-level Go code simultaneously
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i := 0; i < 1000; i++ {
			if i%100 == 0 {
				fmt.Println("\tscoreMap[A]=", scoreMap["A"])
			}
			scoreMap["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			if i%100 == 0 {
				fmt.Println("\t\t\tscoreMap[B]=", scoreMap["B"])
			}
			scoreMap["B"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			if i%100 == 0 {
				fmt.Println("\t\t\tscoreMap[C]=", scoreMap["C"])
			}
			scoreMap["C"]++
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("scoreMap[A]=", scoreMap["A"])
	fmt.Println("scoreMap[B]=", scoreMap["B"])
	fmt.Println("scoreMap[C]=", scoreMap["C"])
}
