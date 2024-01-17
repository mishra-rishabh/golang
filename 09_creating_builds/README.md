# Creating Builds

## Find The Operating System Name

To find the OS name type the below command and search for `GOOS`.
```
go env
```

You will find your OS name-
* `GOOS=windows` for windows OS
* `GOOS=linux` for linux OS
* `GOOS=darwin` for mac OS


## Create A Build Based On Your Current Operating System

Go to your project directory and type below command-

```
go build
```


## Create A Build For Other OS

If you are using ***windows*** but want to take a build for ***mac OS*** then, go to your project directory and type below command-

`GOOS=<os_name> go build`

**Example:**
```
GOOS=darwin go build

GOOS=linux go build
```