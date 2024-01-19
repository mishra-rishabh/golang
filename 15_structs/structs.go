package main

import "fmt"

// defining struct
type Voters struct {
	Name             string
	Age              uint
	isEligibleToVote bool
}

func main() {
	fmt.Println("Structs in Golang")

	/* inserting values in struct */
	newVoter1 := Voters{Name: "Jack", Age: 19, isEligibleToVote: true}
	fmt.Println(newVoter1)

	// or

	newVoter2 := Voters{"Naveen", 17, false}
	fmt.Printf("Voter Details: %+v\n", newVoter2) // we have to use +v as format verb for structs

	// or

	var newVoter3 Voters
	newVoter3.Name = "Parth"
	newVoter3.Age = 20
	newVoter3.isEligibleToVote = true

	// accessing specific value
	fmt.Printf("Voter Name: %+v\n", newVoter3.Name)
}
