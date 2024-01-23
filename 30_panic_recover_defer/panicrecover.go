package main

import "fmt"

func runPanic() {
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
