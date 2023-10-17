## Arrays
Arrays in Go are fixed-length sequences of elements of the same type. They are declared using the `[]` syntax, followed by the type of the elements in the array.<br/>
**example:**
```
// the following code declares an array of 10 integers
var a [10]int
```

Arrays can be initialized with values when they are declared, or they can be initialized later using the `=` operator.<br/>
**example:**
```
// the following code initializes the array a with the values 1, 2, 3, etc.
var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

**NOTE:** Array is very less used in Go. In majority of the other programming languages, array is one of the most used data structure but Go doesn't allow us to use too much of the array. Under the hood there are other datatypes which actually utilizes array.