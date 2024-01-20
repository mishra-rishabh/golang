package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	/* 1. Single condition loop */
	i := 1

	for i <= 5 {
		fmt.Printf("Value of i: %v\n", i)
		i++
	}

	fmt.Println()

	/* 2. A classic initial/condition/after for loop. */
	countries := []string{"India", "USA", "Japan", "Sri_Lanka", "Australia", "Afghanistan", "England"}

	for j := 0; j < len(countries); j++ {
		fmt.Printf("Country: %v\n", countries[j])
	}

	fmt.Println()

	/* break statement */
	for k := 0; k <= 50; k++ {
		if k == 5 {
			break
		}

		fmt.Println("Value of k: ", k)
	}

	fmt.Println()

	/* continue statement */
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println("Value of n: ", n)
	}

	fmt.Println()

	/* goto statement */

	for x := 0; x < 5; x++ {
		if x == 3 {
			goto myLabel
		}

		fmt.Println("Value of x: ", x)
	}

myLabel:
	fmt.Println("This is my label.")

	fmt.Println()

	/* 3. for...range loop */
	cities := []string{"Fatehpur", "Kanpur", "Surat", "Noida", "Gurugram"}

	for index, element := range cities {
		fmt.Printf("index: %v , element: %v\n", index, element)
	}

	fmt.Println()

	/* this will give only element and not the index */
	for _, element := range cities {
		fmt.Printf("only element: %v\n", element)
	}

	fmt.Println()

	/* iterating over maps */
	numSquares := map[int]int{2: 4, 3: 9, 4: 16, 5: 25, 6: 36}

	for key, value := range numSquares {
		fmt.Printf("key-num: %v, value-square: %v\n", key, value)
	}

	fmt.Println()

	/* print only keys */
	for key := range numSquares {
		fmt.Printf("only key-num: %v\n", key)
	}
}
