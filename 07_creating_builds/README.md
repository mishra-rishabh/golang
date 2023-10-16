## Creating Builds
To find the OS name type the below command and search for `GOOS`.
```
go env
```
You will find your OS name-
* `GOOS=windows` for windows OS
* `GOOS=linux` for linux OS
* `GOOS=darwin` for mac OS

To create a build based on the OS, go the project directory and type below command-
```
GOOS=windows go build
```

**NOTE:** Use GOOS=<os_name> if you are making a build for another OS from the other one.<br/>
**example:** If you are using ***windows*** but want to take a build for mac OS then use `GOOS=darwin go build` and the same goes for other OS also.