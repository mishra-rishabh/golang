package main

import "fmt"

func main() {
	fmt.Println("Array tutorial")

	var students [3]string

	fmt.Println("Student list initial:", students)                // will print an empty array with 3 spaces
	fmt.Println("Length of Student list initial:", len(students)) // will print length = 3

	students[0] = "Naveen"
	students[1] = "Akshay"
	students[2] = "Parth"

	fmt.Println("Student list:", students)
	fmt.Println("Length of Student list:", len(students))

	animes := [6]string{"Naruto", "Death_Note", "Erased", "Dragon_Ball_Z", "Black_Clover", "One_Punch_Man"}

	fmt.Println("Anime list:", animes)
	fmt.Println("Length of Anime list:", len(animes))

}
