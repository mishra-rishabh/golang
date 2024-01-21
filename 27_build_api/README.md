# Building API In Golang

## Install Gorilla Mux Dependency

* Package **gorilla/mux** implements a request router and dispatcher for matching incoming requests to their respective handler.
* The name mux stands for "HTTP request multiplexer".
    ```
    go get -u github.com/gorilla/mux
    ```


## Reader And Writer
* **Reader** is where you get the value from whoever is sending the request.
* **Writer** is where you write the response for that.


## Create A Build
```
go build .
```


## Gorilla Mux Documentation

https://github.com/gorilla/mux