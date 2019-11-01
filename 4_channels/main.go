package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	fanout()
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

//
// of paper before they start
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

// manager parcels out work to 20 employees and wait for each to complete
func fanout() {
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
