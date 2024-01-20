package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://gobyexample.com/"

func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Handling web request in go")

	response, err := http.Get(url)

	checkForError(err)

	fmt.Printf("Type of response: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	dataBytes, err := io.ReadAll(response.Body)

	checkForError(err)

	content := string(dataBytes)

	fmt.Printf("Response: %v", content)
}
