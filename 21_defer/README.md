# Defer In Golang

* `defer` is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
* `defer` is often used where e.g. `ensure` and `finally` would be used in other languages.
* As soon as you see the defer keyword with the statement, it means the execution will happen a little bit later.

**example:**
```go
defer fmt.Println("One")
defer fmt.Println("Two")
defer fmt.Println("Three")

fmt.Println("Hello")

/* 
    Output:

    Hello
    Three
    Two
    One
*/
```