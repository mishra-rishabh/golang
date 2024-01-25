# Mutex In Golang

* A Mutex (Mutual Exclusion) is a method used as a locking mechanism to ensure that only one Goroutine is accessing the critical section of code at any point of time.
* This is done to prevent race conditions from happening.
* It is used to provide synchronization in concurrent programs.


## Introduction To Mutex

* Golang provides us with a very powerful tool like Goroutines that is used to do **multi-Threading**.
* When two or more Goroutines are dealing with the same variable in a program, which is manipulated frequently, it causes a condition called **"Racing"**.
* Racing happens when two or more Goroutines try to access and update the value of a shared resource/data concurrently and this updation causes the value to be some random value making our application behavior erroneous, undesirable and unpredictable.
* To avoid such issues to happen, we can use **locks**.
* Mutex in Golang is used to do locking.
* This is available as a standard package called `sync`.
* A Mutex locks a shared resource and releases the lock once the updation of the shared resource is completed. This avoids the race condition.
* In other words, mutexes prevent race conditions by allowing only one goroutine to acquire the lock and access the shared resource, while other goroutines wait until the lock is released.


## Using A Mutex To Lock/Unlock Data

* In Go, the `sync` package provides the Mutex type, which includes two main methods:
    - `Lock()`
    - `Unlock()`
* Any code present between a call to Lock and Unlock will be executed by only one Goroutine.
* If one Goroutine already has the lock and if a new Goroutine is trying to get the lock, then the new Goroutine will be stopped until the mutex is unlocked.

### Program With Race Condition

```go
// Program with race condition 
package main 
import ( 
	"fmt"
	"sync"
) 

var counter = 0 

// This is the function we’ll run in every 
// goroutine. Note that a WaitGroup must 
// be passed to functions by pointer. 
func worker(wg *sync.WaitGroup) { 
	counter += 1 

	// On return, notify the 
	// WaitGroup that we’re done. 
	wg.Done() 
} 
func main() { 

	// This WaitGroup is used to wait for 
	// all the goroutines launched here to finish. 
	var wg sync.WaitGroup 

	// Launch several goroutines and increment 
	// the WaitGroup counter for each 
	for i := 0; i < 1000; i++ { 
		wg.Add(1)		 
		go worker(&wg) 
	} 
	// Block until the WaitGroup counter 
	// goes back to 0; all the workers 
	// notified they’re done. 
	wg.Wait() 
	fmt.Println("Value of x", counter) 
}
```

In the program above, the `worker()` function increments the value of `counter` by **1** and then calls `Done()` on the WaitGroup to inform its completion. The worker function is called **1000** times. Each of these Goroutines run simultaneously and race condition occurs when trying to increment `counter` as multiple Goroutines try to access the value of `counter` concurrently. Running the same program multiple times gives different outputs each time because of the race condition.

### Solving The Above Problem Using Mutex

```go
// Program with race condition fixed by mutex 
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
	// Lock() the mutex to ensure 
	// exclusive access to the state, 
	// increment the value, 
	// Unlock() the mutex 
	m.Lock() 
	counter += 1 
	m.Unlock() 

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
```

In the program above, variable `m` of type **Mutex** is created. The `worker()` function is changed so that the code increments `counter` between `m.Lock()` and `m.Unlock()`. Now only one Goroutine is allowed to execute this piece of code at any point of time and thus race condition is dealt with.

The address of mutex has to be passed while calling `worker()`. If the mutex is passed by value instead then each Goroutine will have its own copy of the mutex and the race condition will still persist.