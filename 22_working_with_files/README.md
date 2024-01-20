# Working With Files In Golang

Reading and writing files are basic tasks needed for many Go programs.

## Reading A File

```go
func readFile(fileName string) {
	// file is read in bytes format and not in string format
	dataBytes, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data in bytes format: ", dataBytes)
	fmt.Println("Data in string format: ", string(dataBytes))
}


// calling a function
readFile("./myFile.txt")
```


## Creating A File

```go
func createFile(fileContent string) int {
	file, err := os.Create("./myFile.txt")

	checkForError(err)

	// WriteString() returns the length of the content
	length, err := io.WriteString(file, fileContent)

	checkForError(err)

	defer file.Close()

	return length
}

// calling a function
fileContent := "This is a golang tutorial for files in golang"

length := createFile(fileContent)
```