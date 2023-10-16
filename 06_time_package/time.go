package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is a time package tutorial")

	currentTime := time.Now()
	fmt.Println(currentTime)

	formattedTime := currentTime.Format("02-01-2006 15:04:05 Monday")
	fmt.Println(formattedTime)

	// Now let's create a date
	createdDate := time.Date(2023, time.January, 6, 15, 9, 6, 0, time.UTC)
	fmt.Println(createdDate)

	formatCreatedDate := createdDate.Format("02-01-2006 15:04:05 Monday")
	fmt.Println(formatCreatedDate)
}
