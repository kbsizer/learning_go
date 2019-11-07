// Lesson 9.6 Interface-Based Race Conditions
package main

import (
	"fmt"
	"log"
	"sync"
)

//--------------- Speaker Interface ---------------

// Speaker defines speaking behavior
type Speaker interface {
	Speak() bool
}

//--------------- Ben ---------------

// Ben is a person who can speak
type Ben struct {
	name string
}

// Speak allows Ben to say hello. Returns false if the
// method is called through the interface value after a partial write.
func (b *Ben) Speak() bool {
	if b.name != "Ben" {
		fmt.Println("Ben says, \"Hello, my name is", b.name, "\"")
		return false
	}
	return true
}

//--------------- Jerry ---------------

// Jerry is a person who can speak
type Jerry struct {
	name string
}

// Speak allows Jerry to say hello. Returns false if the
// method is called through the interface value after a partial write.
func (j *Jerry) Speak() bool {
	if j.name != "Jerry" {
		fmt.Println("Jerry says, \"Hello, my name is", j.name, "\"")
		return false
	}
	return true
}

//--------------- Execution begins here ---------------

func main() {
	ben := Ben{"Ben"}
	fmt.Printf("Initializing %+v\n", ben)
	jerry := Jerry{"Jerry"}
	fmt.Printf("Initializing %+v\n", jerry)

	person := Speaker(&ben)

	// Simulate a series of requests for Ben
	go func() {
		for i := 1; i <= 10000; i++ {
			if i%100 == 0 {
				fmt.Print("B")
			}
			person = &ben
			if !person.Speak() {
				log.Fatal("Speak() returned false while voting for Ben")
			}
		}
	}()

	// Simulate a series of requests for Jerry
	go func() {
		for i := 1; i <= 10000; i++ {
			if i%100 == 0 {
				fmt.Print("J")
			}
			person = &jerry
			if !person.Speak() {
				log.Fatal("Speak() returned false while voting for Jerry")
			}
		}
	}()

	// wait here forever... or until a fatal error occurs in the code above  ;-)
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
