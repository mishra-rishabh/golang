package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getRequest() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)

	checkError(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count: ", byteCount)
	fmt.Println("content: ", responseString.String()) // better way than below one

	// fmt.Println("content: ", string(content))
}

func postRequest() {
	const myUrl = "http://localhost:3000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"language": "golang",
			"topic": "handling web requests"
		}
	`)

	resp, err := http.Post(myUrl, "application/json", requestBody)

	checkError(err)

	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)

	fmt.Println("content: ", string(content))
}

func postFormOperation() {
	const myUrl = "http://localhost:3000/postform"

	// fake json payload
	data := url.Values{}
	data.Add("firstName: ", "Rishabh")
	data.Add("lastName: ", "Mishra")
	data.Add("Age: ", "28")

	resp, err := http.PostForm(myUrl, data)

	checkError(err)

	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)

	fmt.Println("form data: ", string(content))
}

func main() {
	fmt.Println("Making web requests in go")

	// getRequest()
	// postRequest()
	postFormOperation()
}
