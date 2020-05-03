package logger

import (
	"fmt"
	"io"
	"sync"
)

// Logger is ...
type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

// New creates a new Logger (factory function)
func New(w io.Writer, capacity int) *Logger {
	l := Logger{
		ch: make(chan string, capacity),
	}

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		for v := range l.ch {
			fmt.Fprintln(w, v)
		}
	}()

	return &l
}

// Close closes the Logger
func (l *Logger) Close() {
	close(l.ch) // close the channel
	l.wg.Wait() // wait here until our logger's goroutine has shutdown nicely
}

// Println writes out a log message
func (l *Logger) Println(v string) {
	select {
	case l.ch <- v:
		//
	default:
		fmt.Println("Too much resource contention; message dropped")
	}
}
