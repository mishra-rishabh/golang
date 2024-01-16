# Gopath And Reading Go Docs In Command-Line

## Help Command
* `go help`: It will list all the commands and topics.

    ![Alt text](<Screenshot 2024-01-16 120940.png>)
* `go help <command>` for more information about a command.

## GOPATH Environment Variable
The GOPATH environment variable lists places to look for Go code.
* On Unix, the value is a colon-separated string.
* On Windows, the value is a semicolon-separated string.
* On Plan 9, the value is a list.
* **Command:** `go env GOPATH`

    **NOTE:** The above command is case-sensitive. `go env gopath` won't work. It should be in uppercase.

* If the environment variable is unset, GOPATH defaults to a subdirectory named **go** in the user's home directory **($HOME/go on Unix, %USERPROFILE%\go on Windows)**, unless that directory holds a Go distribution.
* Run `go env GOPATH` to see the current GOPATH.