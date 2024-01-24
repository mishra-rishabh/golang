package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getStatusCode(endpoint string) {
	defer wg.Done() // signal that task in done

	resp, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("error occurred")
	}

	fmt.Printf("%d status code for %s\n", resp.StatusCode, endpoint)
}

func main() {
	fmt.Println("waitgroup in go")

	websiteList := []string{
		"https://google.com",
		"https://youtube.com",
		"https://go.dev",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1) // add goroutine to wait for
	}

	wg.Wait() // always comes at the end of the main method
}
