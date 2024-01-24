# WaitGroup In Golang

Previously we saw about goroutines. Here, we will discuss problems with it and how to handle the same.

To wait for multiple goroutines to finish, we can use a wait group.


## Problems With Goroutines In And Ways To Solve Them

* When you run the below program, using `go run` command in command prompt or terminal. You may not see the strings passed to `greet()` function in the goroutines printed to the console. It is because, when you run the program, the other goroutines as well run on the same address space and access the same shared memory and resources.
* If the **main** goroutine or other goroutine is accessing console output to print the message, and if at the same time other goroutine checks for a console output device, it might find the console output to be in use. Hence, the string will not be displayed.

    **example:**
    ```go
    package main

    import (
        "fmt"
        "time"
    )

    func greet(str string) {
        for i := 0; i < 4; i++ {
            fmt.Println(str)
        }
    }

    func main() {
        fmt.Println("Goroutines in go")

        go greet("hello")
        greet("dude")
    }
    ```

* The problem can be handled in two ways:
    1. Using `time` package to add delay in main functions.
    2. Implementing a waitgroup in golang.


## Using Time Package

* By introducing a **delay** in the program, it gives goroutines time to complete the execution.

    **example:**
    ```go
    // modified greet()
    func greet(str string) {
        for i := 0; i < 4; i++ {
            time.Sleep(1 * time.Second)
            fmt.Println(str)
        }
    }
    ```

* The execution of goroutines can be random. Using the `time.Sleep` program is waiting for **1 second**. This gives enough time to complete the other goroutines. This workaround is not the ideal way to handle go-routines. What if other goroutines take more time than **1 second**. Then in that scenario it will end its execution at the end of the main goroutine.
* To handle this better we have **Waitgroup**.


## What Are The Waitgroups In Golang?

* Golang Waitgroup allows you to block a specific code block to allow a set of goroutines to complete execution.
* An example would be to block the main function until the goroutines are completed and then unblocks the group.
* Waitgroup is available as a standard package and can be imported from the `sync` package.

## Methods Of Waitgroups In Golang

* **Add:** The Waitgroup acts as a counter holding the number of functions or go routines to wait for. When the counter becomes **0** the Waitgroup releases the goroutines. 
* **Wait:** The wait method blocks the execution of the application until the Waitgroup counter becomes 0.
* **Done:** Decreases the Waitgroup counter by a value of 1. Each of the goroutine runs and calls `Done()` when finished.

    **example:**
    ```go
    package main

    import (
        "fmt"
        "net/http"
        "sync"
    )

    var wg sync.WaitGroup

    func getStatusCode(endpoint string) {
        defer wg.Done() // signal that task in done

        resp, err := http.Get(endpoint)

        if err != nil {
            fmt.Println("error occurred")
        }

        fmt.Printf("%d status code for %s\n", resp.StatusCode, endpoint)
    }

    func main() {
        fmt.Println("waitgroup in go")

        websiteList := []string{
            "https://google.com",
            "https://youtube.com",
            "https://go.dev",
            "https://github.com",
        }

        for _, web := range websiteList {
            go getStatusCode(web)
            wg.Add(1) // add goroutine to wait for
        }

        wg.Wait() // always comes at the end of the main method
    }
    ```

* Here, `var wg sync.WaitGroup` creates a new WaitGroup while `wg.Add(1)` informs WaitGroup that it must wait for one goroutine on each iteration of a loop.
* This is followed by `defer wg.Done()` alerting the WaitGroup when a goroutine completes.
* `wg.Wait()` then blocks the execution until the goroutine's execution completes.
* The whole process is like **adding** to a counter in `wg.Add()`, **subtracting** from the counter in `wg.Done()`, and waiting for the counter to hit **0** in `wg.Wait()`.