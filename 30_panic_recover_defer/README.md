# Panic, Recover, & Defer In Golang

## Panic

* A `panic` typically means something went unexpectedly wrong.
* Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.


## Recover

* Go makes it possible to `recover` from a `panic`, by using the `recover` built-in function.
* A `recover` can stop a `panic` from aborting the program and let it continue with execution instead.
* An example of where this can be useful: a server wouldn’t want to crash if one of the client connections exhibits a critical error. Instead, the server would want to close that connection and continue serving other clients. In fact, this is what Go’s ***net/http*** does by default for HTTP servers.


## Defer

* `defer` is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
* `defer` is often used where e.g. `ensure` and `finally` would be used in other languages.
* Refer **folder => 21_defer** for more info on `defer`


**example:**
```go
package main

import "fmt"

func runPanic() {
    /* 
        First panic will be called and then defer statement will be called.
        Since handlePanic() is deferred function it will execute after panic statement.
        The moment defered function is seen it will call the handlepanic() and
        recover() will hold the panic created by the runPanic() function and it will
        store it in variable `a` and since a is not nil, the if statement will get executed and
        it will stop the program from terminating and execute further
    */
	defer handlePanic() 
	panic("critical error")
}

func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("Recover: ", a)
	}
}

func main() {
	fmt.Println("Panic, defer and recover in go")

	fmt.Println("before panic!!")
	runPanic()
	/*
		Below statement won't execute because panic will stop the execution.
		To handle this, wee need to use 'revocer'
	*/
	fmt.Println("after panic!!")
}

```