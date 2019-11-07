// Ultimate Go Programming (Bill Kennedy).
// Section 10: Channels
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main executes a series of short demonstrations of handy patterns and
// idioms assocuated with channels in Go.
func main() {
	fmt.Println("\n\n===== Section 10.2.1: worker waiting for a task =====")
	waitForTask()
	fmt.Println("\n\n===== Section 10.2.2: manager waiting on worker =====")
	waitForResult()
	fmt.Println("\n\n===== Section 10.2.3: worker sends signal without data =====")
	waitForFinished()
	fmt.Println("\n\n===== Section 10.3: pooling pattern =====")
	pooling()
	fmt.Println("\n\n===== Section 10.4.1: fanout, part 1 =====")
	fanout1()
	fmt.Println("\n\n===== Section 10.4.2: fanout, part 2 =====")
	fanout2()
	fmt.Println("\n\n===== Section 10.5: drop pattern =====")
	drop()
	fmt.Println("\n\n===== Section 10.6: cancellation pattern =====")
	cancel()
	fmt.Println("\n\n===== END =====")
}

// waitForTask: worker waits until handed a task
func waitForTask() {
	// create UNBUFFERED channel for messages of type string
	ch := make(chan string)

	go func() {
		// wait here until manager gives us a task
		// (because channel is unbuffered/blocking)
		p := <-ch // channel receive unary operator
		fmt.Println("worker: received signal:", p)
	}()

	// wait here for a bit
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	// send message
	ch <- "paper"
	fmt.Println("manager: sent signal")

	time.Sleep(time.Second)
	fmt.Println("------------ done ---------")
}

// waitForResult: after handing off work, manager waits for result from worker
func waitForResult() {
	ch := make(chan string) // unbuffered/blocking channel

	go func() {
		// simulate working
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "final answer"
		fmt.Println("worker: message with result sent")
	}()

	p := <-ch
	fmt.Println("manager: received message. Content:", p)

	time.Sleep(time.Second)
	fmt.Println("------------ done ---------")
}

// waitForFinished: sends a signal without data.
// Especially useful for cancel signals and 1-to-many fanout scenarios.
func waitForFinished() {
	ch := make(chan struct{}) // blocking channel with no data

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		close(ch)
		fmt.Println("worker: send signal by closing channel")
	}()

	_, wd := <-ch // want for slgnal
	fmt.Println("manager: signal received: WithDataFlag =", wd)

	time.Sleep(time.Second)
	fmt.Println("------------ done ---------")
}

// pooling creates an unbuffered channel and ranges over it to parcel out
// work to a pool of employees (workers)
//
// 		* range idiom for reading from a channel
func pooling() {
	ch := make(chan string) // unbuffered (blocking) channel with string data
	const employees = 2
	for e := 1; e <= employees; e++ {
		go func(emp int) {
			// NOTE: We are ranging over a channel. This is essentially a channel RECEIVE
			// (on a blocking channel).  Thus, we wait here until a signal arrives.
			// The loop terminates when the channel's state changes from OPEN to CLOSED.
			for p := range ch {
				fmt.Printf("employee %d: received signal '%s'\n", emp, p)
			}
			fmt.Printf("employee %d: received shutdown signal\n", emp)
		}(e)
	}

	// Simulated work loop (pass 10 work orders to our pool of employees)
	const work = 10
	for w := 0; w < work; w++ {
		signal := fmt.Sprintf("work order #%d", w)
		ch <- signal
		fmt.Printf("manager: sent signal '%s'\n", signal)
	}

	close(ch)
	fmt.Println("manager: sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("------------ done ---------")
}

// fanout1 demonstrates the fanout pattern.  In this example, a manager parcels out work to 20 employees
// and waits for each to complete.
func fanout1() {
	employees := 20
	ch := make(chan string, employees) // create a channel with a buffer size of 20 (fanout)

	for e := 1; e <= employees; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			signal := fmt.Sprintf("paper #%d", emp)
			ch <- signal
			fmt.Printf("employee #%d: sent '%s'\n", emp, signal)
		}(e)
	}

	// manager
	for employees > 0 {
		p := <-ch
		fmt.Println(p)
		employees--
		fmt.Printf("manager: received signal. Waiting on %d more.\n", employees)
	}

	time.Sleep(time.Second)
	fmt.Println("------------ done ---------")
}

// fanout2 demonstrates the fanout semaphore pattern. It limits number of go routines which
// can run at a given time.  In this example, a manager parcels out work to 20 employees
// and waits for each to complete
func fanout2() {
	employees := 20
	ch := make(chan string, employees)
	const cap = 5
	sem := make(chan bool, cap) // our semaphore is a boolean channel of size 5

	for e := 1; e <= employees; e++ {
		go func(emp int) {
			sem <- true // signal "I'm going active" (block here if semaphore queue is full)
			{           // BEGIN BLOCK OF WORK
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				msg := fmt.Sprintf("paper from employee #%d", emp)
				ch <- msg
				fmt.Printf("employee: Finished work.  Sending '%s'\n", msg)
			} // END BLOCK OF WORK
			<-sem // signal "I'm done" (decrement active thread semaphore)
		}(e)
	}

	for employees > 0 {
		p := <-ch
		fmt.Printf("manager: received '%s'\n", p)
		employees--
	}
}

// drop uses the drop pattern to manage back pressure when things go bad.  Classic example is a DNS
// server under a denial-of-service attack: Server is being flooded with so many requests
// that it drowns.
//
// 		* range idiom for reading from a channel
//		* select idiom for writing to a buffered channel without blocking
func drop() {
	// create a buffered channel with capacity of 5
	const capacity = 5
	ch := make(chan string, capacity)

	// worker which receives signals from the channel
	go func() {
		for p := range ch {
			fmt.Printf("worker: received signal '%s'\n", p)
		}
		fmt.Println("worker: shutdown signal received")
	}()

	// manager send out out work (signals)
	const work = 20
	for w := 1; w <= work; w++ {
		select { // <== allows us to attempt operations without blocking
		case ch <- "paper":
			fmt.Printf("manager: signal #%d successfully sent\n", w)
		default:
			fmt.Printf("manager: signal #%d dropped\n", w)
		}
	}

	close(ch)
	fmt.Println("manager: sent shutdown signal")

	time.Sleep(time.Second)
}

// cancel uses the cancel pattern to abandon a piece of work we are not willing to
// wait on.  In this example, manager has a hard deadline and, if the worker doesn't
// finish in time, he wants to move on.
//
//		* Uses select idiom to block but only for a predetermined amount of time.
//		* Uses a timer channel to generate an event at a specific time.
func cancel() {
	ch := make(chan string, 1)

	for i := 1; i <= 10; i++ {
		fmt.Println("Trial ", i)
		// worker
		go func() {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch <- "paperwork"
			fmt.Println("employee: sent paperwork")
		}()

		// create a timer channel which will ping us 100ms in the future
		tc := time.After(100 * time.Millisecond)

		select {
		case p := <-ch:
			fmt.Println("manager: received signal from worker:", p)
		case t := <-tc:
			fmt.Println("manager: timeout.  Done waiting. (msg:", t, ")")
		}
	}

	time.Sleep(time.Second)
}
