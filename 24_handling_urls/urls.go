package main

import (
	"fmt"
	"net/url"
)

const myUrl = "postgres://user:pass@host.com:5432/path?k=v#f"

func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Handling URL in go")

	fmt.Println("URL:", myUrl)

	/* parsing */
	result, _ := url.Parse(myUrl) // Here, we have ignored the error. It's better to handle the same.

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Port:", result.Port())
	fmt.Println("Path:", result.Path)
	fmt.Println("Fragment:", result.Fragment)
	fmt.Println("Raw Query:", result.RawQuery)

	/*
		Parse query params into a map.

		The parsed query param maps are from strings to slices of strings,
		so index into [0] if you only want the first value.

	*/
	queryParams := result.Query()

	fmt.Printf("Type of queryParams: %T\n", queryParams)
	fmt.Println(queryParams["k"])
	fmt.Println(queryParams["k"][0])

	/* print params using loop */
	for _, val := range queryParams {
		fmt.Println("Param: ", val)
	}

	/* making url from components */
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "abcd.dev",
		Path:     "/path",
		RawQuery: "user=goku",
	}

	fullUrl := partsOfUrl.String()
	fmt.Println("full url: ", fullUrl)
}
