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

### strconv Pakacge
For conversions we have used `strconv` package.<br/>
* The `strconv` package in Go provides functions for converting between strings and other data types. It is a standard library package that is included in all Go installations.
* The `strconv` package provides a number of functions for converting between strings and different data types, including:
    * **`Atoi()`:** Converts a string to an integer.
    * **`Itoa()`:** Converts an integer to a string.
    * **`ParseBool()`:** Converts a string to a boolean value.
    * **`FormatBool()`:** Converts a boolean value to a string.
    * **`ParseFloat()`:** Converts a string to a floating-point number.
    * **`FormatFloat()`:** Converts a floating-point number to a string.
* The `strconv` package also provides functions for converting between strings and different types of byte slices, runes, and structs.

[Reference Link for strconv Package](https://pkg.go.dev/strconv)
