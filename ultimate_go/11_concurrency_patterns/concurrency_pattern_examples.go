// Introduction to concurrrency patterns and Go's context package.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// TraceID represents the trace ID.
type TraceID string

// TraceIDKey is the type of value to use for the key.  The key is
// type specific and only values of the same type will match.
type TraceIDKey int

// Examples in this file covers the following sections of Ultimate Go Programming:
//    11.1 Context (Part 1)
//    11.1 Context (Part 2)
func main() {
	fmt.Printf("\n\n----------------- example 1 -----------------\n\n")
	example1()

	fmt.Printf("\n\n----------------- example 4 -----------------\n\n")
	example4()

	fmt.Printf("\n\n----------------- example 5 -----------------\n\n")
	example5()

	fmt.Printf("\n\n----------------- failureDetectionDemo -----------------\n\n")
	failureDetectionDemo()
}

// example1 illustrates:
//  * storing values into and retrieving them from a context
//	* the "if..ok" idiom for retrieving and testing for success
func example1() {
	// create a traceID for this request
	traceID := TraceID("12345-67890-gamma-delta")

	// declare a key with the value of zero of type userKey
	const traceIDKey TraceIDKey = 0

	// store the traceID valeu inside the CONTEXT with a value of
	// zero for the key type
	emptyDummyContext := context.Background() // See also: context.TODO()
	ctx := context.WithValue(emptyDummyContext, traceIDKey, traceID)

	// Correct way to retrieve trace ID value from the context value bag
	if uuid, ok := ctx.Value(traceIDKey).(TraceID); !ok {
		log.Fatal("Failed while attempting to retrieve value from context")
	} else {
		log.Printf("Retrieved traceID '%s'\n", uuid)
	}

	// what happens if we try to retrieve that traceID value
	// from the context value bag NOT using the proper key type
	if _, ok := ctx.Value(0).(TraceID); !ok {
		log.Println("You can't do this! --> ctx.Value(0) returned ok = FALSE")
	}
}

type data struct{ string }

// Illustrates:
//	* creating a context with timeout
//	* using defer to ensure cancel() is called
//  * using select wait on multiple channels
func example4() {
	// specify timeout duration
	duration := 150 * time.Millisecond

	for i := 1; i <= 10; i++ {

		// create a context that is both manually cancellable and will
		// signal a cancel at the specified duration
		// NOTE: second parameter is a reference to a handler function which we MUST call
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		defer cancel() // ProTip: place a "defer" immediately after "WithTimeout"

		// create a channel to receive a signal indicating that work is done
		ch := make(chan data, 1)

		// ask a goroutine to do some work for us
		go func() {
			// simulate random amount of work (0~500ms)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			// report that work is done
			ch <- data{"done"}
		}()

		// wait here for work to finish or timeout
		select {
		case d := <-ch:
			log.Printf("worker completed work and sent '%s'\n", d)
		case <-ctx.Done():
			log.Printf("Received DONE signal from our CONTEXT; work cancelled.\n")
		}
	}
	time.Sleep(100 * time.Millisecond)
}

// example5 demonstrates making an HTTP request for information from an external source with
// a bounded wait time.
// Utiltizes:
//	* Context package
// 	* Net and HTTP packages
//	* Creating HTTP requests with timeout
//	* Streaming data using io.Copy
func example5() {
	// create an HTTP request (note: we are DEFINING, not ISSUING the request)
	// req, err := http.NewRequest("GET", "https://www.ardenlabs.com/blog/post/index.xml", nil)
	req, err := http.NewRequest("GET", "https://blog.golang.org/index", nil)
	if err != nil {
		log.Fatal(err)
	}

	for delay := 25; delay < 200; delay += 25 {
		log.Printf("\n\n~~~~~ Setting max wait time to %d ms ~~~~~\n\n", delay)

		// create a context with a timeout of 50 ms
		// (note: Unlike previous examples, we have a real context from the request)
		ctx, cancel := context.WithTimeout(req.Context(), time.Duration(delay)*time.Millisecond)
		defer cancel()

		// declare a new transport and client for the call
		tr := http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
		client := http.Client{
			Transport: &tr,
		}

		// make the web call in a separate goroutine so it can be cancelled
		ch := make(chan error, 1) // create buffered channel of size 1
		go func() {
			log.Println("Starting request")
			// make the web call and return any error
			resp, err := client.Do(req)
			if err != nil {
				ch <- err
				return
			}
			// ensure we close on return
			defer resp.Body.Close()
			// write the response (first 64 bytes) to STDOUT
			io.CopyN(os.Stdout, resp.Body, 64)
			fmt.Println()
			ch <- nil // put nil into error channel to signal successful completion
		}()

		// Wait for EITHER request success OR request failed OR timeout
		select {
		case err := <-ch:
			if err != nil {
				log.Println("CASE 2: request failed.")
				log.Fatal(err)
			}
			log.Println("CASE 1: Request completed successfully and written to STDOUT.")
		case <-ctx.Done():
			log.Println("CASE 3: timed-out while waiting.  Cancelling work.")
			tr.CancelRequest(req) // send cancel
			log.Println(<-ch)
		}
	}
}

//
// ----------- begin Failure Detection Demo code ------------
//

// device allows us to mock a device we use to persist log events
type device struct {
	problem bool
}

// Write is a method on device; implements the io.Writer interface.
func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		// simulate disk I/O problems
		time.Sleep(time.Second)
	}
	fmt.Println("p:", string(p))
	return len(p), nil
}

// 11.2 Failure Detection
func failureDetectionDemo() {
	// number of goroutines writing to log
	const grs = 10
	// create a loggervalue with a buffer of capacity
	// equal to number of goroutines that will be logging
	var d device
	l := log.New(&d, "prefix", 0)
	// generate goroutines
	for i := 1; i <= grs; i++ {
		go func(id int) {
			for {
				l.Println(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// we want to control the simulated disk blocking. Capture
	// interrupt signals to toggle device issues. Use ctrl-z
	// to kill the program.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

}
