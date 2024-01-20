package main

import (
	"fmt"
	"io"
	"os"
)

func checkForError(err error) {
	if err != nil {
		panic(err) // we will look more into 'panic' in upcoming chapters
	}
}

func createFile(fileContent string) int {
	file, err := os.Create("./myFile.txt")

	checkForError(err)

	// WriteString() returns the length of the content
	length, err := io.WriteString(file, fileContent)

	checkForError(err)

	defer file.Close()

	return length
}

func readFile(fileName string) {
	// file is read in bytes format and not in string format
	dataBytes, err := os.ReadFile(fileName)

	checkForError(err)

	fmt.Println("Data in bytes format: ", dataBytes)
	fmt.Println("Data in string format: ", string(dataBytes))
}

func main() {
	fmt.Println("Working with files in go")

	fileContent := "This is a golang tutorial for files in golang"

	length := createFile(fileContent)

	fmt.Println("Length of file content: ", length)

	readFile("./myFile.txt")
}
