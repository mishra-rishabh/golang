package main

import "fmt"

func main() {
	fmt.Printf("Maps in Go\n\n")

	studentsAndRollNum := make(map[int]string)
	fmt.Println(studentsAndRollNum) // empty map

	studentsAndRollNum[1] = "Naveen"
	studentsAndRollNum[23] = "Akshay"
	studentsAndRollNum[56] = "Nikhil"
	studentsAndRollNum[45] = "Parth"

	// print whole map
	fmt.Println("\nStudents and roll num:", studentsAndRollNum)

	// print value for a specific key
	fmt.Println("Student for roll num 45:", studentsAndRollNum[45])

	// declare and define map in a single line
	countryCapital := map[string]string{"India": "New_Delhi", "Japan": "Tokyo"}
	fmt.Println("\nCountry and capital:", countryCapital)
	fmt.Println("countryCapital map length:", len(countryCapital))

	// delete an entry from a map => we will take studentsAndRollNum map
	delete(studentsAndRollNum, 56)
	fmt.Println("\nStudents and roll num after deleting an entry:", studentsAndRollNum)

	// delete a whole map
	clear(countryCapital)
	fmt.Printf("\ncountryCapital map after clearing %v:\n\n", countryCapital)

	// loop through a map to print it
	for key, val := range studentsAndRollNum {
		fmt.Printf("key: %v, value: %v\n", key, val)
	}

	/* We will look into loops in upcoming chapters */

	// checking whether the value is present for a given key or not
	_, present := studentsAndRollNum[67]
	fmt.Printf("\n%v", present)
}
