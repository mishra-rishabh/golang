package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("If-Else Statements in Golang")

	number := 5

	if number < 0 {
		fmt.Println("Negative")
	} else if number > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Zero")
	}

	/*
		Other way is to declare and initialize a variable
		and at the same time check it for some condition.
	*/

	if age := 17; age < 18 {
		fmt.Println("Not eligible to vote!")
	} else {
		fmt.Println("Eligible to vote!")
	}

	fmt.Printf("\nSwitch and case in golang\n\n")

	rand.NewSource(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Printf("Value of dice is: %v\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open.")
	case 2:
		fmt.Println("You can move 2 spots.")
	case 3:
		fmt.Println("You can move 3 spots.")
	case 4:
		fmt.Println("You can move 4 spots.")
		// fallthrough // "fallthrough" statement will also execute the next case. Uncomment to check the output
	case 5:
		fmt.Println("You can move 5 spots.")
		// fallthrough
	case 6:
		fmt.Println("You can move 6 spots and roll a dice again.")
	default:
		fmt.Println("No such number found on a dice!")
	}

	fmt.Println()
	/* Multiple expressions in the same case statement. */
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!!")
	default:
		fmt.Println("Weekday")
	}
}
