# Decision Making In Go

* Decision making structures require that the programmer specify one or more conditions to be evaluated or tested by the program, along with a statement or statements to be executed if the condition is determined to be true, and optionally, other statements to be executed if the condition is determined to be false.

* Go programming language provides the following types of decision making statements.

    ## 1. `if` Statement

    * An `if` statement consists of a boolean expression followed by one or more statements.

        **Syntax:**
        ```
        if(boolean_expression) {
            /* statement(s) will execute if the boolean expression is true */
        }
        ```

        **example:**
        ```go
        var a int = 10
    
        if a < 20 {
            fmt.Printf("a is less than 20\n" )
        }
        ```

    
    ## 2. `if...else` Statement

    * An `if` statement can be followed by an optional `else` statement, which executes when the boolean expression is `false`.

        **Syntax:**
        ```
        if(boolean_expression) {
            /* statement(s) will execute if the boolean expression is true */
        } else {
            /* statement(s) will execute if the boolean expression is false */
        }
        ```

        **example:**
        ```go
        var a int = 100;
 
        if( a < 20 ) {
            fmt.Printf("a is less than 20\n" );
        } else {
            fmt.Printf("a is not less than 20\n" );
        }

        /*
            Other way is to declare and initialize a variable
            and at the same time check it for some condition.
        */

        if age := 17; age < 18 {
            fmt.Println("Not eligible to vote!")
        } else {
            fmt.Println("Eligible to vote!")
        }
        ```

    
    ## 3. `if...else if...else` Statement

    * An `if` statement can be followed by an optional `else if...else` statement, which is very useful to test various conditions using single `if...else if` statement.
    * When using `if` , `else if` , `else` statements there are few points to keep in mind −
        - An if can have zero or one else's and it must come after any else if's.
        - An if can have zero to many else if's and they must come before the else.
        - Once an else if succeeds, none of the remaining else if's or else's will be tested.
    
        **Syntax:**
        ```
        if(boolean_expression 1) {
            /* Executes when the boolean expression 1 is true */
        } else if( boolean_expression 2) {
            /* Executes when the boolean expression 2 is true */
        } else if( boolean_expression 3) {
            /* Executes when the boolean expression 3 is true */
        } else {
            /* executes when the none of the above condition is true */
        }
        ```

        **example:**
        ```go
        number := 5

        if number < 0 {
            fmt.Println("Negative")
        } else if number > 0 {
            fmt.Println("Positive")
        } else {
            fmt.Println("Zero")
        }
        ```

    
    ## 4. Nested `if` Statements

    * It is always legal in Go programming to **nest** `if-else` statements, which means you can use one `if` or `else if` statement inside another `if` or `else if` statement(s).
    * You can nest `else if...else` in the similar way as you have nested `if` statement.

        **Syntax:**
        ```
        if( boolean_expression 1) {
            /* Executes when the boolean expression 1 is true */
            if(boolean_expression 2) {
                /* Executes when the boolean expression 2 is true */
            }
        }
        ```

        **example:**
        ```go
        var a int = 100
        var b int = 200
        
        if( a == 100 ) {
            if( b == 200 )  {
                fmt.Printf("Value of a is 100 and b is 200\n" );
            }
        }
        ```


    ## 5. `switch` Statement

    * Switch statements express conditionals across many branches.
    * A `switch` statement allows a variable to be tested for equality against a list of values.
    * Each value is called a case, and the variable being switched on is checked for each **switch case**.

        **example:**
        ```go
        rand.NewSource(time.Now().UnixNano())
        diceNumber := rand.Intn(6) + 1

        fmt.Printf("Value of dice is: %v\n", diceNumber)

        switch diceNumber {
        case 1:
            fmt.Println("Dice value is 1 and you can open.")
        case 2:
            fmt.Println("You can move 2 spots.")
        case 3:
            fmt.Println("You can move 3 spots.")
        case 4:
            fmt.Println("You can move 4 spots.")
            // fallthrough // "fallthrough" statement will also execute the next case. Uncomment to check the output
        case 5:
            fmt.Println("You can move 5 spots.")
            // fallthrough
        case 6:
            fmt.Println("You can move 6 spots and roll a dice again.")
        default:
            fmt.Println("No such number found on a dice!")
        }
        ```
    * You can use **commas** to separate multiple expressions in the same case statement. We use the optional `default` case in this example as well.

        **example:**
        ```go
        /* Multiple expressions in the same case statement. */
        switch time.Now().Weekday() {
        case time.Saturday, time.Sunday:
            fmt.Println("Weekend!!")
        default:
            fmt.Println("Weekday")
        }
        ```


    ## 6. `select` Statement

    * A `select` statement is similar to `switch` statement with difference that case statements refers to `channel` communications.
    * We will look into `select` statement in upcoming chapters

**NOTE:**
* There is no ***ternary if*** in Go, so you’ll need to use a full if statement even for basic conditions.
* You don’t need parentheses around conditions in Go, but that the braces are required.