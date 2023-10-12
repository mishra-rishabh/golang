package main

import "fmt"

// someGlobalVar := "abcd"		// invalid, because := operator can only be used inside of functions
var someGlobalVar = "abcd" // allowed

// public variable
var LastName string = "Mishra" // valid public variable

func main() {
	var myName string = "Rishabh"
	fmt.Println(myName)
	fmt.Printf("variable is of type: %T \n\n", myName)

	var isPresent bool = true
	fmt.Println(isPresent)
	fmt.Printf("variable is of type: %T \n\n", isPresent)

	var smallUint uint8 = 123
	fmt.Println(smallUint)
	fmt.Printf("variable is of type: %T \n\n", smallUint)

	var smallFloat float32 = 123.98876
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n\n", smallFloat)

	// default values
	var defaultForInt int
	var defaultForString string
	var defaultForBool bool
	var defaultForFloat float64

	fmt.Printf("defaultForInt: %v \n", defaultForInt)
	fmt.Printf("defaultForString: %q \n", defaultForString)
	fmt.Printf("defaultForBool: %v \n", defaultForBool)
	fmt.Printf("defaultForFloat: %f \n\n", defaultForFloat)

	// implicit type
	var slogan = "Inquilab Zindabad!"
	fmt.Println(slogan)

	// no var style
	age := 27
	fmt.Println(age)

	fmt.Printf("\n")

	fmt.Println(someGlobalVar)

	// const Pi = 3.14159 // valid
	const Pi float64 = 3.14159 // also valid
	fmt.Println(Pi)

	fmt.Printf("\n")

	fmt.Printf("Public variable %q", LastName)
}
