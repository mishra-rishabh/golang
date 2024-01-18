# Slices

* Slices are much more powerful and mostly used datatype in the entire golang.
* Slices are actually under the hood arrays, which also means, Go Slice is an abstraction over Go Array.
* Go Array allows you to define variables that can hold several data items of the same kind but it does not provide any inbuilt method to increase its size dynamically or get a sub-array of its own.
* Slices overcome this limitation. It provides many utility functions required on Array and is widely used in Go programming.


## Defining A Slice

To define a slice, you can declare it as an array without specifying its size. Alternatively, you can use `make` function to create a slice.

**example:**
```go
var numbers []int /* a slice of unspecified size */
/* numbers == []int{0, 0, 0, 0, 0}*/
numbers = make([]int, 5) /* a slice of length 5*/
```


## Subslicing

Slice allows **lower-bound** and **upper-bound** to be specified to get the subslice of it using `[lower-bound:upper-bound]` syntax.

**example:**
```go
/* create a slice */
numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}   
printSlice(numbers)

/* print the sub slice starting from index 1(included) to index 4(excluded)*/
fmt.Println("numbers[1:4] ==", numbers[1:4])

/* missing lower bound implies start-index as 0*/
fmt.Println("numbers[:3] ==", numbers[:3])

/* missing upper bound implies end-index as len(s)*/
fmt.Println("numbers[4:] ==", numbers[4:])
```


## `len()` And `cap()` Functions

* A slice is an abstraction over array.
* It actually uses arrays as an underlying structure.

    ### `len()`

    * The `len()` function returns the number of elements presents in the slice.
    * In other words, length is total no of elements the slice is having and can be used for looping through the elements we stored in slice. Also when we print the slice all elements till length gets printed.

    
    ### `cap()`

    * The `cap()` function tells you the capacity of the underlying array.
    * In other words, capacity is total no elements in underlying array, when you append more elements the length increases till capacity. After that any further append to slice causes the capacity to increase automatically(apprx double) and length by no of elements appended.

**example:**
```go
// slice of length 3 and capacity 4
scores := make([]int, 3, 4)

scores[0] = 89
scores[1] = 91
scores[2] = 97

scores = append(scores, 100)
fmt.Printf("For scores slice after adding elements => length: %d, capacity: %d, slice: %v", len(scores), cap(scores), scores)
/* 
    O/P: For scores slice after adding elements => length: 4, capacity: 4, slice: [89 91 97 100]
*/

scores = append(scores, 75) // this will increase the capacity, because we are adding an extra element
fmt.Printf("For scores slice after adding more element => length: %d, capacity: %d, slice: %v", len(scores), cap(scores), scores)

/* 
    O/P: For scores slice after adding another elements => length: 5, capacity: 8, slice: [89 91 97 100 75]
*/
```


## Nil slice
* If a slice is declared with no inputs, then by default, it is initialized as `nil`.
* Its length and capacity are zero.

**example:**
```go
var someSlice []int // slice with no length and capacity
```


## append() And copy() Functions

* Using these functions we can increase the capacity and copy the contents of one slice to another. 

    ### `append()`
    
    One can increase the capacity of a slice using the `append()` function.

    **example:**
    ```go
    // slice of length 2 and capacity 3
    scores := make([]int, 2, 3)

    scores[0] = 89
    scores[1] = 91
    // scores[2] = 97  // this will give an error, because the length is 2

    scores = append(scores, 100, 97, 78)    // this will increase the capacity of the array and length as well
    ```

    
    ### `copy()`

    Using `copy()` function, the contents of a source slice are copied to a destination slice.

    **example:**
    ```go
    copyScores := make([]int, len(scores), cap(scores))

	copy(copyScores, scores) // copying from scores to copyScores
    ```


## Removing A Value From A Slice

We can also use `append()` function to remove an element from a slice.

**example:**

```go
var codingLangs = []string{"go", "rust", "javascript", "solidity", "c++"}
/*
    We know that append actually reallocates the memory accordingly.
    So, if we want to remove 'javascript' (index 2) from the above slice,
    then we have to append the elements from 'go' (index 0) to 'rust' (index 1)
    and from 'solidity' (index 3) till end of the slice

    NOTE: Remember that, in append() the upper-bound is always excluded
*/

codingLangs = append(codingLangs[:2], codingLangs[3:]...)
/*
    O/P: Coding languages after removing an element: [go rust solidity c++]
*/
```
