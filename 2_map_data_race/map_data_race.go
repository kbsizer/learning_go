// map_data_race demonstrates the following:
//    - Maps in Go are NOT inherently thread-safe.  Even when threads are never touching the
//      same key-value pair
//    - The use of -race (Go's built-in data race detection facility)
//
// Data races are among the most common and hardest to debug types of bugs in concurrent systems.
// To help diagnose such bugs, Go includes a built-in data race detector. To use it, add the -race
// flag to the go command:
//    $ go test -race mypkg    // to test the package
//    $ go run -race mysrc.go  // to run the source file
//    $ go build -race mycmd   // to build the command
//    $ go install -race mypkg // to install the package
//
// See: https://golang.org/doc/articles/race_detector.html
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
