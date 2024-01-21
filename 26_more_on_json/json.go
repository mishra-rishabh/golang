package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	/* The below members will appear as it is when json is printed */
	// Name     string
	// Price    int
	// Platform string
	// Password string
	// Tags     []string

	/* The below members will appear as what is there inside backtics `` */
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this will not show the password field
	Tags     []string `json:"tags,omitempty"` // this will keep this field but will remove if it is empty
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func EncodeJson() {
	myCourses := []course{
		{"React", 299, "Udemy", "abcd123", []string{"web-dev", "js"}},
		// {"Golang", 599, "Coursera", "xyz123", []string{"backend", "google"}},
		{"Golang", 599, "Coursera", "xyz123", nil}, // just to have variety of data
		{"Solidity", 199, "Udemy", "pqrs123", []string{"blockchain", "ethereum"}},
	}

	/* packaging the above data as JSON data */
	// finalJson, err := json.Marshal(myCourses)
	// finalJson, err := json.MarshalIndent(myCourses, "something", "\t")
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	/* comment and uncomment the above statements to check different outputs */

	checkError(err)

	fmt.Printf("finalJson: %s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React",
		"price": 299,
		"website": "Udemy",
		"tags": [
				"web-dev",
				"js"
		]
	}
	`)

	var courseData course

	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		fmt.Println("Valid JSON")

		json.Unmarshal(jsonDataFromWeb, &courseData)

		fmt.Printf("Unmarshalled json data: %#v\n", courseData)
	} else {
		fmt.Println("Invalid JSON")
	}

	/* Some cases where you just want to add data to key value */
	var myJsonData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myJsonData)

	fmt.Printf("my json data: %#v\n", myJsonData)

	/* we can also loop through myJsonData and that is the advantage of creating a map */
	for k, v := range myJsonData {
		fmt.Printf("key: %v, value: %v, type: %T\n", k, v, v)
	}
}

func main() {
	fmt.Println("Creating JSON data in golang")

	// EncodeJson()
	DecodeJson()
}
