# Methods In Go

* Go supports special types of functions called **methods** defined on struct types.
* In method declaration syntax, a **"receiver"** is present to represent the container of the function.
* This receiver can be used to call a function using **"."** operator.

    **example:**
    ```go
    type Rectangle struct {
        width  int
        length int
    }

    func (rect Rectangle) area() int {
        return (rect.length * rect.width)
    }

    /* creating struct */
    rect := Rectangle{length: 5, width: 9}

    /* calling a fuction */
    fmt.Println("Area: ", rect.area())
    ```