package main

import (
	"fmt"
	"time"
)

func greet(str string) {
	for i := 0; i < 4; i++ {
		time.Sleep(1 * time.Second) // Comment and uncomment this line to see the changes
		fmt.Println(str)
	}
}

func main() {
	fmt.Println("Goroutines in go")

	go greet("hello")
	greet("dude")

}
