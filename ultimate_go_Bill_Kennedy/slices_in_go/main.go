// Illustrates the
// See: https://www.ardanlabs.com/blog/2013/09/iterating-over-slices-in-go.html?utm_source=Ardan+Labs+Newsletter&utm_campaign=eee13ab87a-ARDAN_LABS_INSIDER_JAN2020&utm_medium=email&utm_term=0_1a6a20b712-eee13ab87a-193662193&mc_cid=eee13ab87a&mc_eid=4e8f7bf7fb
package main

import (
	"fmt"
)

type Dog struct {
	Name string
	Age  int
}

func main() {

	// create two Dog values with composite literals
	// and display the object's address
	jackie := Dog{
		Name: "Jackie",
		Age:  19,
	}

	fmt.Printf("Jackie is stored at: %p\n", &jackie)

	sammy := Dog{
		Name: "Sammy",
		Age:  10,
	}

	fmt.Printf("Sammy is stored at: %p\n", &sammy)

	// create a slice containing the two dogs
	dogs := []Dog{jackie, sammy}

	fmt.Println("==================")

	// iterate over COPIES of the two dogs
	for _, dog := range dogs {
		fmt.Printf("Name: %s, Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n", &dog)
		fmt.Println("-------")
	}

	fmt.Println("\n==================")

	// HOW WE CAN GET INTO TROUBLE...
	// ...by thinking the address of the "dog" variable inside the
	// loop can be used as a pointer to the Dog objects in the
	// slice:
	mySliceOfDogs := []*Dog{}
	for _, dog := range dogs {
		mySliceOfDogs = append(mySliceOfDogs, &dog)
	}

	// We might *think* we have a new slice containing a pointer
	// to jackie and a pointer to sammy, but we don't:
	for _, dog := range mySliceOfDogs {
		fmt.Printf("Name: %s, Age: %d\n", dog.Name, dog.Age)
	}

	fmt.Println("\n==================")

	// Alternative: Using slice of pointers

	dogPtrSlice := []*Dog{&jackie, &sammy}

	for _, dog := range dogPtrSlice {
		fmt.Printf("Name: %s, Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n", dog)
		fmt.Println("-------")
	}

	// NOTE: When the slice is a collection of Dog values or a collection
	// of pointers to Dog values, the range loop is the same. Go handles
	// access to the Dog value regardless of whether we are using a pointer
	// or not. This is awesome but can sometimes lead to a bit of confusion.
}
