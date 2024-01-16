package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is a type conversion tutorial")
	fmt.Println("Enter any number between 1 to 10")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Print("You have entered: ", input)
	fmt.Printf("Type of given input is %T\n", input)

	// below line will give an error, because we are also taking a new line ('\n') as an input so, we have to remove it
	// incrementNum, err := strconv.ParseFloat(input, 64)

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	incrementedNum := num + 1

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number incremented by 1:", incrementedNum)
		fmt.Printf("Type is %T\n", incrementedNum)
	}

	// Convert a boolean value to a string.
	var booleanValue bool = true
	var stringValue string = strconv.FormatBool(booleanValue)
	fmt.Printf("Type is %T\n", booleanValue)
	fmt.Printf("Type is %T\n", stringValue)

	// implicit type conversion
	var floatNum float64 = 34.5
	// var intNum int = 6
	var sum float64 = floatNum + 6 // here, 6 will be converted to float automatically by compiler
	fmt.Println("Sum", sum)
	fmt.Printf("Type is %T\n", sum)

}
