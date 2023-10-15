## Type Conversion
Type conversion in Go is the process of converting one data type to another. There are two types of type conversion in Go:

1. **Explicit type conversion:** Explicit type conversion is the process of manually converting one data type to another. It is done by using the type name followed by the variable name in parentheses.<br/>
**example:**
    ```
    var booleanValue bool = true
    var stringValue string = strconv.FormatBool(booleanValue)
    ```
2. **Implicit type conversion:** Implicit type conversion is the process of automatically converting one data type to another. It is done by the compiler when it is necessary to do so. For example, if you add an integer and a float, the compiler will automatically convert the integer to a float so that the addition can be performed.<br/>
**example:**
    ```
    var sum float64 = floatNum + 6
    ```

**NOTE:** Implicit type conversion is only possible when the two data types are compatible. For example, you cannot implicitly convert an integer to a string, because strings are not compatible with integers.