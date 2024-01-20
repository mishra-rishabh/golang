# Closures In Golang

* Go supports **anonymous** functions, which can form **closures**.
* Anonymous functions are useful when you want to define a function **inline** without having to name it.

    **example:**
    ```go
    /*
        This function intSequence returns another function,
        which we define anonymously in the body of intSequence.
        The returned function closes over the variable i
        to form a closure.
    */
    func intSequence() func() int {
        i := 0

        return func() int {
            i++

            return i
        }
    }

    /*
        We call intSequence, assigning the result (a function) to nextInt.
        This function value captures its own i value, which will be
        updated each time we call nextInt.
    */
    nextInt := intSequence()

    fmt.Println(nextInt())
    ```