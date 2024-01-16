# The Comma OK Syntax

* The **comma ok syntax** or **error ok syntax** in Go is used to check for errors and handle them accordingly.
* It is a common idiom that is used in many different situations, such as reading from a map, opening a file, and decoding JSON.
* The comma ok syntax is used to check for errors when returning a value from a function.
* It works by returning two values: **the value itself** and a **boolean value** indicating whether there was an error.
* In Go we don't have **try/catch**, so if something goes wrong there's nobody to catch that, because Go is designed in such a way that it treats a problem or error as a variable. Whatever the reader actually reads, we want to store it in a variable and this is where the comma ok syntax come in.

**NOTE:** We have already used this syntax in *user_input* section.

