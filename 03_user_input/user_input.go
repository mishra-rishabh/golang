package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		walrus operator is used, because when using these packages for user input
		we don't know what is coming up as an input, i.e., what type of input it's
		gonna be (int, string, float)
	*/
	// 1st way
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter you name:")

	// comma ok syntax => we will look it in next chapter
	input, _ := reader.ReadString('\n') // read until new line is found

	var age uint

	fmt.Println("Enter your age:")
	// 2nd way
	fmt.Scanln(&age)

	fmt.Print("\nHello ", input)
	fmt.Printf("You have entered %v as your age", age)
}
