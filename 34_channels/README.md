# Channels In Golang

* Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
* Channels in Go act as a medium for goroutines to communicate with each other and this communication is lock-free.
* In other words, it is an abstract type using which we can send and receive values.
* The communication is bidirectional by default, meaning that you can send and receive values from the same channel.


## Create Channel In Go

* In Go, we use the `make()` function to create a channel.

    **Syntax:**
    ```
    channel_name:= make(chan Type)
    ```

    **example:**
    ```go
    myChannel := make(chan int)
    ```
* Here,
    - `myChannel`: is name of the channel.
    - `(chan int)`: indicates that the channel is of integer type.


    ### Channel Example
    
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Channels in golang")

        /* Create channel of integer type */
        number := make(chan int)

        /* Access type and value of channel */
        fmt.Printf("Channel type: %T", number)
        fmt.Printf("\nChannel value: %v", number)
    }

    /*
        Output:
        Channel type: chan int
        Channel value: 0xc0000280c0
    */
    ```

    In the above example, the channel is of integer type (specified by `chan int`), we get the same output and the value of a channel is a memory address, which acts as a medium through which goroutines send and receive data to communicate.


## Golang Channel Operations

* Once we create a channel, we can send and receive data between different goroutines through the channel.

    ### 1. Send Data To The Channel

    **Syntax:**
    ```
    channelName <- data
    ```

    Here the data after the `<-` operator is sent to `channelName`.

    **example:**
    ```go
    // send integer data to channel
    number <- 15

    // send string data
    message <- "Learning Go Channel"
    ```


    ### 2. Receive Data From The Channel

    **Syntax:**
    ```
    <- channelName
    ```

    This accesses the data from `channelName`.

    **example:**
    ```go
    // receive data 15
    <- number

    // receive data "Learning Go Channel"
    <- message
    ```

    ### Channel Operations Example
    ```go
    package main

    import "fmt"

    func channelData(num chan int, msg chan string) {
        /* Send data to channel */
        num <- 15
        msg <- "Learning channels in golang"
    }

    func main() {
        fmt.Println("Channels in golang")

        /* Create channel of integer and string type */
        number := make(chan int)
        message := make(chan string)

        /* Function call with go routine */
        go channelData(number, message)

        /* Retrieve channel data */
        fmt.Println("Channel data number: ", <-number)
        fmt.Println("Channel data message: ", <-message)
    }
    ```


## Blocking Nature Of Channel

* In Go, the channel automatically blocks the send and receive operations depending on the status of goroutines.

    ### 1. When a goroutine sends data into a channel, the operation is blocked until the data is received by another goroutine.

    **example:**
    ```go
    package main
    import "fmt"

    func main() {
        // create channel
        ch := make(chan string)
        
        // function call with goroutine
        go sendData(ch)

        // receive channel data
        fmt.Println(<-ch)
    }

    func sendData(ch chan string) {
        // data sent to the channel
        ch <- "Received. Send Operation Successful"
        fmt.Println("No receiver! Send Operation Blocked")
    }

    /*
        Output:
        No receiver! Send Operation Blocked
        Received. Send Operation Successful
    */
    ```

    * In the above example, we have created the `sendData()` goroutine to send data to the channel. The goroutine sends the string data to the channel.

    * If the channel is not ready to receive the message, it prints ***"No receiver! Send Operation Blocked"***.

    * Inside the `main()` function, we are calling `sendData()` before receiving data from the channel. That's why the first ***"No receiver..."*** is printed.

    * And, when the channel is ready to receive data, the data sent by goroutine is printed.


    ### 2. When a goroutine receives data from a channel, the operation is blocked until another goroutine sends the data to the channel.

    **example:**
    ```go
    package main
    import "fmt"

    func main() {
        // create channel
        ch := make(chan string)

        // function call with goroutine
        go receiveData(ch)

        fmt.Println("No data. Receive Operation Blocked")

        // send data to the channel 
        ch <- "Data Received. Receive Operation Successful" 
    }

    func receiveData(ch chan string) {
        // receive data from the channel
        fmt.Println(<-ch)
    }

    /*
        Output:
        No data. Receive Operation Blocked
        Data Received. Receive Operation Successful
    */
    ```

    * In the above example, we have created the `receiveData()` goroutine to receive data from the channel. The goroutine receives the string data from the channel.

    * If the channel hasn't yet sent the data, it prints ***"No data. Receive Operation Blocked"***.

    * Inside the `main()` function, we are calling `receiveData()` before sending data to the channel. That's why the first ***"No data..."*** is printed.

    * And, when the channel sends data, the data received by goroutine is printed.


## Closing A Channel

* You can also close a channel with the help of `close()` function. This is an in-built function and sets a flag which indicates that no more value will send to this channel.

    **example:**
    ```go
    ch := make(chan int)

    close(ch)

    ele, ok := <-ch

    fmt.Println("Element: ", ele)
    fmt.Println("OK: ", ok)
    ```

    Here, if the value of `ok` is `true` which means the channel is open so, read operations can be performed. And if the value of `ok` is `false` which means the channel is closed so, read operations can't be performed.


## More About Channels

https://go101.org/article/channel.html

