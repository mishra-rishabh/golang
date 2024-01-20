package main

import (
	"fmt"
)

/* Function with no return-type */
func greet() {
	fmt.Println("Vande Mataram!")
}

/* Function returning a single value */
func add(a int, b int) int {

	add := a + b

	return add
}

/* Function returning multiple values */
func sumProduct(a int, b int) (int, int) {
	var add int
	var product int

	add = a + b
	product = a * b

	return add, product
}

/* Variadic function: function with variable number of arguments */
func addVariadicFunc(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println("Function in golang")

	greet()

	sum := add(5, 6)
	fmt.Println("Sum is: ", sum)

	add, multiply := sumProduct(5, 7)
	fmt.Printf("sum: %v, product: %v", add, multiply)

	fmt.Println()

	/* same function with different number of arguments */
	fmt.Println("sum of 2 numbers: ", addVariadicFunc(2, 3))
	fmt.Println("sum of 3 numbers: ", addVariadicFunc(2, 3, 5))
	fmt.Println("sum of 4 numbers: ", addVariadicFunc(2, 3, 5, 9))

	/* same variadic function can be used to add elements of a slice */
	integerSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("sum of slice elements: ", addVariadicFunc(integerSlice...))

}
