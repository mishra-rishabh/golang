Variables are containers for storing data values.

### Go Variable Types
Variable type should be known in advance.<br/>
In Go, there are different types of variables, for example:
* **String:** `string`, stores text, such as "Vande Mataram". String values are surrounded by double quotes.
* **Boolean:** `bool`, stores values with two states: true or false.
* **Integer:** `int8, int16, int32, int64, uint8, uint16, uint32, uint64, uintptr`, stores integers (whole numbers), such as 123 or -123.
* **Floating:** `float32, float64`, stores floating point numbers, with decimals, such as 19.99 or -19.99.
* **Complex:** `complex64, complex128`, both these data types take 2 floating point literals, one for the real part and one for the imaginary part. complex64 takes 2 float32 literals, whereas complex128 takes 2 float64 literals.

**NOTE:**
* `byte` is an alias for `uint8`.
* `rune` is an alias for `int32` and it represents a Unicode characters, which is a broader set than ASCII characters. These Unicode characters are encoded in the UTF-8 format.

Apart from the above types there are some advance types-
* **Arrays**
* **Slices**
* **Maps**
* **Structs**
* **Pointers**

NOTE: In Go almost everything is a **type**, functions, channels, mutex, etc, all are considered as types internally.

### Declaring a Variable
* Variables are declared using `var` keyword.<br/>
**example:** `var someNumber int = 123` <br/>
* The value of an initialized variable with no assignment will be its default value.

### Declaring a Variable (Other Way)
#### 1. Implicit Type:
* When variable is declared but type is not mentioned then **implicit type** comes into picture.
* It means, Go will automatically decide the type of the variable based on the value.<br/>
**example:** `var myName = "Rishabh Mishra"` here, Go will decide the type of `myName` variable based on the value it holds, i.e., it will treat this variable as string type.

#### 2. No `var` Style (Walrus Operator):
* We can totally ignore the keyword `var` and still can declare the variables.
* We can make use of the **Walrus Operator (`:=`)** to do the same.<br/>
**example:** `totalAmount := 60000`
* The walrus operator can only be used inside of functions, and the variable must be assigned a value immediately. It cannot be used to declare global variables, and it cannot be used to declare multiple variables in a single statement.

### Public and Private Variables
* In Go, variables are considered public if they start with an uppercase letter. Variables that start with a lowercase letter are considered private.
* This means that public variables can be accessed from anywhere in the program, while private variables can only be accessed from within the package in which they are declared.<br/>
**example:**
    ```
    var Name string = "Rishabh"   // this is a valid public variable
    var name string = "Rishabh"   // this is an invalid public variable because it is a private variable
    ```

### Constants
* Constants in Go are immutable values that cannot be changed once they are declared. They can be of any type, including numeric types, strings, and booleans.
* To declare a constant, we use the `const` keyword followed by the name of the constant and its value.<br/>
**example:**
    ```
    const Pi = 3.14159
    const Name = "Rishabh"
    const Active = true
    ```

[Reference Link for Numeric Types](https://go.dev/ref/spec#Numeric_types)