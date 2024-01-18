package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Slices in Golang\n\n")

	/* Defining a slice */
	var animes = []string{"Naruto", "Death_Note", "Bleach"}

	fmt.Printf("Type of animes: %T:\n", animes)
	fmt.Println("animes list: ", animes)

	/* adding items to slice */
	animes = append(animes, "Black_Clover", "AOT")
	fmt.Println("animes list after appending: ", animes)

	/*
		Subslicing the slices-
		For this we can use the 'colon syntax'. We can provide the
		start_index (included) and end_index (excluded).
		The below statement will return a new slice with elements from
		index 1 to index (3-1), i.e., 2
	*/
	// animes = append(animes[1:3])

	/*
		This will subslice the actual slice from index 1 to the end.
	*/
	// animes = append(animes[1:])

	/*
		The below statement will return a new slice with elements from
		index 1 to index (3-1), i.e., 2 and add a new elemet "Erased" after slicing.
	*/
	animes = append(animes[1:3], "Erased")

	fmt.Printf("animes list after slicing: %v\n\n", animes)

	/*
		make() function to create slice
	*/
	numbers := make([]int, 5) // slice of length 5

	numbers[0] = 9
	numbers[1] = 8
	numbers[2] = 7
	numbers[3] = 6
	numbers[4] = 5

	// this will give an error: index out of bound range, because we have allocated the size of slice as 5
	// numbers[5] = 4

	// this will work because append will reallocate the memory for the slice so that all the values would be accomodated
	numbers = append(numbers, 4, 3, 2, 1)

	// sort the slice of integers
	sort.Ints(numbers)
	// check if the slice of integers is sorted or not
	fmt.Println("Numbers slice: ", numbers)
	fmt.Println(sort.IntsAreSorted(numbers))

	// slice of length 3 and capacity 4
	scores := make([]int, 3, 4)

	fmt.Printf("\nFor scores slice => length: %d, capacity: %d, slice: %v \n", len(scores), cap(scores), scores)
	scores[0] = 89
	scores[1] = 91
	scores[2] = 97
	// scores[3] = 100 // this will give index out of bound range error, because we have defined its length as 3

	scores = append(scores, 100)
	// scores = append(scores, 82, 87, 73, 68)

	fmt.Printf("For scores slice after adding elements => length: %d, capacity: %d, slice: %v \n\n", len(scores), cap(scores), scores)

	copyScores := make([]int, len(scores), cap(scores))

	copy(copyScores, scores) // copying from scores to copyScores

	fmt.Printf("For copyScores slice after copying => length: %d, capacity: %d, slice: %v \n\n", len(copyScores), cap(copyScores), copyScores)

	/*
		Comment and Uncomment the statements to check the outputs
	*/
}
