// Introduction to concurrrency patterns and Go's context package.
package main

import (
	"context"
	"log"
)

// TraceID represents the trace ID.
type TraceID string

// TraceIDKey is the type of value to use for the key.  The key is
// type specific and only values of the same type will match.
type TraceIDKey int

func main() {
	example1()
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
	// See also: context.TODO()
	ctx := context.WithValue(context.Background(), traceIDKey, traceID)

	// retrieve that trace ID value from teh context value bag
	if uuid, ok := ctx.Value(traceIDKey).(TraceID); !ok {
		log.Fatal("Failed while attempting to retrieve value from context")
	}

	log.Printf("Retrieved traceID '%s'\n", uuid)

	// what happens if we try to retrieve that traceID value
	// from the context value bag NOT using the proper key type
	if _, ok := ctx.Value(0).(TraceID); !ok {
		log.Fatal("TraceID not found!")
	}
}

// Illustrates:
//	* creating a context with timeout
//	* using defer
func example4() {
	// specify timeout duration
	duration := 150 * time.Millisecond
	// create a context that is both manually cancellable and will 
	// signal a cancel at the specified duration
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// create a channel to receive a signal indicating that work is done
	ch := make(chan )
}
