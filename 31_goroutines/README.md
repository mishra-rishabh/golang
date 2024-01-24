# Goroutines In Golang

* Go language provides a special feature known as a **Goroutines**.
* A goroutine is a lightweight **thread** managed by the Go runtime.
* A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program.
* In other words, every concurrently executing activity in Go language is known as a Goroutines.
* The cost of creating Goroutines is very small as compared to the thread.
* Every program contains at least a single Goroutine and that Goroutine is known as the **main Goroutine**.
* All the Goroutines are working under the main Goroutines if the main Goroutine terminated, then all the goroutine present in the program also terminated.
* Goroutine always works in the background.


## Creating A Goroutine

* To start a goroutine, use keyword `go` before making the function call.

    **Syntax:**
    ```
    func name() {
        // statements
    }

    // using go keyword as the prefix
    go name()
    ```
* When the above statement is executed, the statement itself gets executed in the current goroutine, but `go name()` gets executed in a new goroutine.

    **example:**
    ```go
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

* When you run this program, using `go run` command in command prompt or terminal. You may not see the strings passed to `greet()` function in the goroutines printed to the console. It is because, when you run the program, the other goroutines as well run on the same address space and access the same shared memory and resources.
* If the **main** goroutine or other goroutine is accessing console output to print the message, and if at the same time other goroutine checks for a console output device, it might find the console output to be in use. Hence, the string will not be displayed.
* One can use `time.Sleep()` to make it work.

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
* We added the `Sleep()` method in our program which makes the main Goroutine sleeps for **1 second** in between 1 second the new Goroutine executes, displays **"dude"** on the screen, and then terminate after 1 second main Goroutine re-schedule and perform its operation. This process continues until the value of the `i < 4` after that the main Goroutine terminates.
* Here, both Goroutine and the normal function work concurrently.


## Advantages Of Goroutines

* Goroutines are cheaper than threads.
* Goroutine are stored in the stack and the size of the stack can grow and shrink according to the requirement of the program. But in threads, the size of the stack is fixed.
* Goroutines can communicate using the channel and these channels are specially designed to prevent race conditions when accessing shared memory using Goroutines.
* Suppose a program has one thread, and that thread has many Goroutines associated with it. If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread. All these details are hidden from the programmers.
