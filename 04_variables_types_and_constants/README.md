# Variable, Types And Constants in Golang

Variables are containers for storing data values.


## Datatypes

* In the Go programming language, data types refer to an extensive system used for declaring variables or functions of different types.
* The type of a variable determines how much space it occupies in storage and how the bit pattern stored is interpreted.
* Variable type should be known in advance.


## Go Variable Types

In Go, there are different types of variables, for example:
1. **String:**
    - `string`, stores text, such as "Vande Mataram".
    - String values are surrounded by double quotes ("").
    - Strings are immutable types that is once created, it is not possible to change the contents of a string.

2. **Boolean:** `bool`, stores values with two states: `true` or `false`.

3. **Numeric Types:** They are arithmetic types and they represents-
    * **Integer types:** `int8,...,int64, uint8,...,uint64, uintptr`, stores integers (whole numbers), such as 123 or -123.

        | Sr.No. | Type   | Description                                       |
        |--------|--------|---------------------------------------------------|
        | 1      | uint8  | Unsigned 8-bit integers (0 to 255)               |
        | 2      | uint16 | Unsigned 16-bit integers (0 to 65535)            |
        | 3      | uint32 | Unsigned 32-bit integers (0 to 4294967295)       |
        | 4      | uint64 | Unsigned 64-bit integers (0 to 18446744073709551615) |
        | 5      | int8   | Signed 8-bit integers (-128 to 127)              |
        | 6      | int16  | Signed 16-bit integers (-32768 to 32767)         |
        | 7      | int32  | Signed 32-bit integers (-2147483648 to 2147483647) |
        | 8      | int64  | Signed 64-bit integers (-9223372036854775808 to 9223372036854775807) |
    
    - **Floating point:** `float32, float64`, stores floating point numbers, with decimals, such as 19.99 or -19.99.
    - **Complex:** `complex64, complex128`, both these data types take 2 floating point literals, one for the real part and one for the imaginary part. complex64 takes 2 float32 literals, whereas complex128 takes 2 float64 literals.

        **NOTE:**
        * `byte` is an alias for `uint8`.
        * `rune` is an alias for `int32` and it represents a Unicode characters, which is a broader set than ASCII characters. These Unicode characters are encoded in the UTF-8 format.
        * `uintptr` is an unsigned integer to store the uninterpreted bits of a pointer value.

4. **Derived types:** There are some advance types-
    - Array
    - Slice
    - Map
    - Struct
    - Pointer
    - Union
    - Function
    - Interface
    - Channel

**NOTE:** In Go almost everything is a **type**, functions, channels, mutex, etc, all are considered as types internally.


## Declaring a Variable

* Variables are declared using `var` keyword. <br/>
    **example:** `var someNumber int = 123`
* The value of an initialized variable with no assignment will be its default value.


## Declaring a Variable (Other Way)

### 1. Implicit Type:
* When variable is declared but type is not mentioned then **implicit type** comes into picture.
* It means, Go will automatically decide the type of the variable based on the value.<br/>
    **example:** `var myName = "Rishabh Mishra"` here, Go will decide the type of `myName` variable based on the value it holds, i.e., it will treat this variable as string type.

### 2. No `var` Style (Walrus Operator):
* We can totally ignore the keyword `var` and still can declare the variables.
* We can make use of the **Walrus Operator (`:=`)** to do the same. <br/>
    **example:** `totalAmount := 60000`
* The walrus operator can only be used inside the functions, and the variable must be assigned a value immediately. It cannot be used to declare global variables, and it cannot be used to declare multiple variables in a single statement.


## Public and Private Variables

* In Go, variables are considered public if they start with an uppercase letter. Variables that start with a lowercase letter are considered private.
* This means that public variables can be accessed from anywhere in the program, while private variables can only be accessed from within the package in which they are declared. <br/>
    **example:**
    ```go
    var Name string = "Rishabh"   // this is a valid public variable
    var name string = "Rishabh"   // this is an invalid public variable because it is a private variable
    ```


## Constants

* Constants in Go are immutable values that cannot be changed once they are declared. They can be of any type, including numeric types, strings, and booleans.
* To declare a constant, we use the `const` keyword followed by the name of the constant and its value. <br/>
    **example:**
    ```go
    const Pi = 3.14159
    const Name = "Rishabh"
    const Active = true
    ```

[Reference Link for Types in Go](https://go.dev/ref/spec#Types)