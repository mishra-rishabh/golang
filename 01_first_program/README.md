# First Program in Go

## Creating A New Module in Go
Once Go is installed, go to the folder where you want to save your file and run the below command-
```
go mod init <your_module_name>

// example
go mod init first_program
```

The above command creates a new module, initializing the `go.mod` file that describes the module. At the start, it will only add the module path and go version in `go.mod` file.

## What Is `go.mod` File?
* The `go.mod` file lists the specific versions of the dependencies that your project uses.
* The `go.mod` file is the root of dependency management in GoLang.
* All the modules which are needed or to be used in the project are maintained in `go.mod` file.

## Runnig Your Go File
```
go run <file_name>

// example
go run main.go
```

The above command will run the file but won't build it so we are not going to get any executables which we can execute.

## Observation
Notice one thing in the code when you imported the package `fmt` and let's say you haven't used it and save the file, it automatically disappears similarly when you make use of some of the package and it is not imported, go will automatically import it for you when you save it. It is because, go automatically cleans and uncleans the packages which you need or probably you don't.