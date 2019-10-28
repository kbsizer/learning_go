// Demonstrates use of Go Routines, GOMAXPROCS, WaitGroups,
// critical sections, atomic operations, synchronization blocks, etc.
// by spawning multiple worker threads computing Pi and other silliness.
package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

const maxOsThreads = 2
const maxWorkers = 5
const maxIterations = 100000

func main() {
	// setup
	// GOMAXPROCS limits the number of operating system threads that can execute user-level Go code simultaneously
	runtime.GOMAXPROCS(maxOsThreads)

	fmt.Println("#########################################################")
	fmt.Println("#### PART 1: Basic Thread Synchronization")
	fmt.Println("#########################################################")

	// initialize wait group semaphore
	var wg sync.WaitGroup
	wg.Add(maxWorkers)
	// spawn mulitple worker threads
	for w := 1; w <= maxWorkers; w++ {
		name := fmt.Sprintf("%s#%d", strings.Repeat("        ", w), w)
		go func() {
			computePi(name, maxIterations)
			wg.Done()
		}()
	}

	// wait here for all workers to finish
	wg.Wait()

	fmt.Println("\n\n#########################################################")
	fmt.Println("#### PART 2: Synchronization using the atomic package")
	fmt.Println("####")
	fmt.Println("#### There are", maxWorkers, "worker threads, each making 2")
	fmt.Println("#### iterations, therefore, we expect the final count")
	fmt.Println("#### to be:", maxWorkers*2)
	fmt.Println("#########################################################")

	// NOTE: Atomic instructions are faster than mutexes,
	// which are faster than channels

	fmt.Print("\n\nVERSION 1\n\n")

	// Version 1: no atomic access; very broken
	var counter int
	wg.Add(maxWorkers) // reset wait group
	for i := 0; i < maxWorkers; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				value := counter
				value++
				fmt.Println(value)
				counter = value
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Print("\n\nVERSION 2\n\n")

	// Version 2: Using Atomic Operations --> see sync/atomic package
	var counter64 int64
	wg.Add(maxWorkers) // reset wait group
	for i := 0; i < maxWorkers; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				value := atomic.AddInt64(&counter64, 1)
				fmt.Println(value)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Print("\n\nVERSION 3\n\n")

	// Version 3: Using Mutexes --> see sync/atomic package
	mutex := &sync.Mutex{}
	counter = 0        // reset counter
	wg.Add(maxWorkers) // reset wait group
	for i := 0; i < maxWorkers; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				mutex.Lock() // BEGIN Critical Section
				{
					value := counter
					value++
					fmt.Println(value)
					counter = value
				}
				mutex.Unlock() // END Critical Section
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Print("\n\nVERSION 4\n\n")

	// Version 4: Using Read-Write Mutexes
	rfMutex := &sync.RWMutex{}
	counter = 0            // reset counter
	wg.Add(2 * maxWorkers) // initialize wait group
	// a group of readers
	for i := 0; i < 2*maxWorkers; i++ {
		go func(iAmAWriter bool) {
			for count := 0; count < 2; count++ {
				if iAmAWriter {
					rfMutex.Lock() // BEGIN Critical Section
					{
						value := counter
						value++
						fmt.Println("WRITER found", value)
						counter = value
					}
					rfMutex.Unlock() // END Critical Section
				} else {
					rfMutex.RLock() // BEGIN Critical Section
					{
						value := counter
						fmt.Println("READER found", value)
					}
					rfMutex.RUnlock() // END Critical Section
				}
			}
			wg.Done()
		}(i%2 == 0)
	}
	wg.Wait()
}

// computePi computes the value of Pi using:
//     Pi/4 = 1 - 1/3 + 1/5 - 1/7 + 1/9 - ...
// See: http://mathworld.wolfram.com/PiFormulas.html
func computePi(name string, iterations int) {
	fmt.Println(name, "START")
	var piOver4 float64
	for k := 1; k < iterations; k++ {
		term := 1.0 / (2.0*float64(k) - 1.0)
		if k%2 == 0 {
			term = -term
		}
		piOver4 += term
		// show progress
		if k%1000 == 0 {
			fmt.Printf("%v: pi[%v]=%v\n", name, k, piOver4*4.0)
			// runtime.Gosched() //XXX
		}
	}
	fmt.Println(name, "DONE. Final answer:", piOver4*4.0)
}
