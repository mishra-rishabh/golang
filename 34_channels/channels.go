package main

import "fmt"

func channelData(num chan int, msg chan string) {
	/* Send data to channel */
	num <- 15
	msg <- "Learning channels in golang"
}

func main() {
	fmt.Println("Channels in golang")

	/* Create channel of integer and string type */
	number := make(chan int)
	message := make(chan string)

	/* Function call with go routine */
	go channelData(number, message)

	/* Retrieve channel data */
	fmt.Println("Channel data number: ", <-number)
	fmt.Println("Channel data message: ", <-message)
}
