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

NOTE: In go almost everything is a **type**, functions, channels, mutex, etc, all are considered as types internally.