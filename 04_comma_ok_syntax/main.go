package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	/*
		Whatever the error will be will come into err variable.
		Here, we haven't used err variable because it will give error if we don't use it in the code.
		So, we have used _ (underscore) instead.
	*/
	// input, err := reader.ReadString('\n')
	input, _ := reader.ReadString('\n')

	fmt.Println("Enter you name:")
	fmt.Print("\nHello ", input)
}
