package main

import "fmt"

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
	fmt.Printf("defaultForFloat: %f \n", defaultForFloat)
}
