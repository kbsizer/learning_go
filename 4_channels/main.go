package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("===== PART 1: worker waiting for a task =====")
	waitForTask()
	fmt.Println("===== PART 2: manager waiting on worker =====")
	waitForResult()
	fmt.Println("===== PART 3: worker sends signal without data =====")
	waitForFinished()
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
