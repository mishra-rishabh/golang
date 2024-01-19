# Structs In Golang

* Go’s **structs** are **user-defined** data type. They are typed collections of fields.
* They’re useful for grouping data together to form records.
* Suppose you want to keep track of the books in a library. You might want to track the following attributes of each book-
    - Title
    - Author
    - Subject
    - Book ID


## Creating A Struct
* To define a structure, we must use `type` and `struct` statements.
* The struct statement defines a new data type, with multiple members for our program.
* The type statement binds a name with the type which is `struct` in our case.

    **Syntax:**
    ```
    type struct_variable struct {
    member definition;
    member definition;
    ...
    member definition;
    }
    ```

    **example:**
    ```go
    type Voters struct {
        Name             string
        Age              uint
        isEligibleToVote bool
    }
    ```


## Inserting Values To Struct

* Once a structure type is defined, it can be used to declare variables of that type.
    
    **Syntax:**
    ```
    variable_name := structure_variable_type {value1, value2...valueN}
    ```

    **example:**
    ```go
    /* inserting values */
	newVoter1 := Voters{Name: "Jack", Age: 19, isEligibleToVote: true}

    // or

	newVoter2 := Voters{"Alex", 17, false}

	// or

	var newVoter3 Voters
	newVoter.Name = "Robert"
	newVoter.Age = 20
	newVoter.isEligibleToVote = true
    ```


## Accessing Struct Values
* To access values from a struct we can use the below way-

    **Syntax:**
    ```
    struct_variable.member
    ```
    
    **example:**
    ```go
    fmt.Printf("Voter Name: %v\n", newVoter3.Name)
    ```

**NOTE:** Structs are mutable.
