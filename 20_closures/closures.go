package main

import "fmt"

func intSequence() func() int {
	i := 0

	return func() int {
		i++

		return i
	}
}

func main() {
	fmt.Println("Closures in Go")

	/*
		nextInt is now a function with i as 0
		We call intSequence(), assigning the result (a function) to nextInt.
	*/
	nextInt := intSequence()

	/* invoke nextInt to increase i by 1 and return the same */
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	/* nextInt2 is now a function with i as 0 */
	nextInt2 := intSequence()

	/* invoke nextInt to increase i by 1 and return the same */
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())

}
