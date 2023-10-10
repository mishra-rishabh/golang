first download and install go from the official website.<br/>
check if go is installed or not by running the below command in terminal-
```
go version
```

once installed go to the folder where you want to save your file and run the below command-
```
go mod init <your_module_name>
// example
go mod init first_program
```

now run your go file using the below command-
```
go run <file_name>
// example
go run main.go
```

Notice one thing in the code when you imported the package "fmt" and let's say you haven't used it and save the file, it automatically disappears similarly when you make use of some of the package and it is not imported, go will automatically import it for you when you save it. It is because, go automatically cleans and uncleans the packages which you need or probably you don't.