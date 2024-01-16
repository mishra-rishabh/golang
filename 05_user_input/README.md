# Taking User Input

There are two main ways to take user input in Go:
## 1. Using The bufio Package

The `bufio` package provides a function called `NewReader()` that can be used to create a reader object that reads from the user's input. The `NewReader()` function takes a value of type `io.Reader` as an argument, which can be any type that implements the `io.Reader` interface. <br/>
**example:** refer *user_input.go*


## 2. Using The fmt Package

The `fmt` package provides a function called `Scan()` that can be used to read input from the user. The `Scan()` function takes a variable number of arguments, and it will read input from the user until it encounters a newline character. <br/>
**example:** refer *user_input.go*

**NOTE:** Which method you use to take user input in Go depends on your specific needs. If you need to read a **single line of input**, then using `fmt.Scan()` function is the simplest and most straightforward way to do it. However, if you need to **read multiple lines of input**, or if you need to process the input in a more complex way, then using the `bufio` package is a better option.

## Official Link For Golang Packages

https://pkg.go.dev/