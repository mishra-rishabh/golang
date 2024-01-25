package main

import (
	"fmt"
	"sync"
)

var counter = 0

// This is the function we’ll run in every
// goroutine. Note that a WaitGroup must
// be passed to functions by pointer.
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	// either use this way
	/* m.Lock()
	counter += 1
	m.Unlock() */

	// or use this way (Good Practice)
	m.Lock()
	defer m.Unlock()
	counter += 1

	// On return, notify the
	// WaitGroup that we’re done.
	wg.Done()
}
func main() {

	// This WaitGroup is used to wait for
	// all the goroutines launched here to finish.
	var wg sync.WaitGroup

	// This mutex will synchronize access to state.
	var m sync.Mutex

	// Launch several goroutines and increment
	// the WaitGroup counter for each
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}
	// Block until the WaitGroup counter
	// goes back to 0; all the workers
	// notified they’re done.
	wg.Wait()

	fmt.Println("Value of counter: ", counter)
}
