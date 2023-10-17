package main

import "fmt"

func main() {
	fmt.Println("Pointer tutprial")

	// decalring a pointer of tupe int
	var ptr *int
	fmt.Println("value of ptr: ", ptr)

	someNum := 6
	var somePtr = &someNum // storing the address of someNum variable

	fmt.Println("somePtr: ", somePtr)
	fmt.Println("value stored in somePtr: ", *somePtr)

	*somePtr += 3
	fmt.Println("new value of someNum: ", someNum)
}
