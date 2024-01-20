package main

import "fmt"

func deferDemoFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func main() {
	fmt.Println("Defer in go")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	fmt.Println("Hello")

	deferDemoFunc()

	fmt.Println() // just to add space for readability

	/*
		In defer statements, first deferDemoFunc() will be called
		then defer fmt.Println("Three") will be called then
		defer fmt.Println("Tow") and then defer fmt.Println("One")
		will be called

		In deferDemoFunc(), value of 'i' will be printed in reverse order,
		because 'defer' puts the statement at last of all the other statements
		with which it is used.

		As soon as you see the defer keyword with the statement,
		it means the execution will happen a little bit later.
	*/
}
