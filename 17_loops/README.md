# Loops in Golang

A loop statement allows us to execute a statement or group of statements multiple times.


## For Loop

* A `for` loop is a repetition control structure.
* It allows you to write a loop that needs to execute a specific number of times.
* `for` is Go’s only looping construct.

    ### 1. For Loop With Single Condition.

    ```go
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i++
    }
    ```

    ### 2. For Loop With Initialization, Condition And Update

    ```go
    for j := 7; j <= 9; j++ {
		fmt.Printf("Value of j: %v\n", j)
	}
    ```


## Loop Control Statements

* Loop control statements change an execution from its normal sequence.
* When an execution leaves its scope, all automatic objects that were created in that scope are destroyed.

    ### 1. `break` Statement

    * When a break statement is encountered inside a loop, the loop is immediately terminated and the program control resumes at the next statement following the loop.
    * Go `break` statement can be used to break a `for` loop or to break a `switch` statement
    * If you are using nested loops, the break statement will stop the execution of the innermost loop and start executing the next line of code after the block.

        **example:**
        ```go
        var i int = 0
        
        for i < 50 {
            if i == 5 {
                break
            }
            fmt.Println(i)
            i = i + 1
        }
        ```

    ### 2. `continue` Statement

    * Go continue statement is used to skip the other statements in the loop for current iteration, and continue with the next iterations of the loop.

        **example:**
        ```go
        for n := 0; n <= 5; n++ {
            if n%2 == 0 {
                continue
            }

            fmt.Println(n)
        }
        ```

    ### 3. `goto` Statement

    * Goto statement is used to go or jump to a labeled statement in the same function.

        **example:**
        ```go
        for x := 0; x < 5; x++ {
            if x == 3 {
                goto myLabel
            }

		    fmt.Println("Value of x: ", x)
	    }

        myLabel:
            fmt.Println("This is my label.")
        ```

        **Note:** Use of goto statement is highly discouraged in any programming language because it becomes difficult to trace the control flow of a program, making the program difficult to understand and hard to modify. Any program that uses a goto can be rewritten using some other construct.


## For Range Loop

* Go language does not have a builtin `forEach` kind of statement. But the nearest kind of syntax could be achieved using for loop with range.
* The `range` keyword is used in `for` loop to iterate over items of an **array**, **slice**, **channel** or **map**.
* During each iteration, we get access to the index of the element, and the element itself.
* With array and slices, it returns the **index** of the item as **integer**. With maps, it returns the **key** of the next **key-value pair**.

    **example:**
    ```go
    x := []int{10, 20, 30, 40, 50}
    
    for index, element := range x {
        fmt.Println(index, " - ", element)
    }
    ```

* If we do not need index while looping, we may provide underscore symbol instead of index.

    **example:**
    ```go
    x := [5]int{11, 22, 33, 44, 55}
 
    for _, element := range x {
        fmt.Println(element)
    }
    ```


**NOTE:** **pre-increment** and **pre-decrement** are not supported in Go, e.g. `--i`, `++i`