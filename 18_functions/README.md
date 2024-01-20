# Functions In Go

* A function is a group of statements that together perform a task.
* Every Go program has at least one function, which is `main()`.
* A function **declaration** tells the compiler about a function name, return type, and parameters.
* A function **definition** provides the actual body of the function.
* Functions are also known as **method**, **sub-routine**, or **procedure**.

    **Syntax:**
    ```
    func function_name( [parameter_list] ) [return_types] {
        statement(s)
    }
    ```
    **where,**
    * **func** is the keyword used to define a function.
    * **function_name** is the name by which we can call the function.
    * **parameter_list** is optional and these are the arguments that you provide to our function to transform them or consider them as input.
    * **return_types** is the list of data types of the multiple values that our function can return.


## Function With No Return Type And Returning Nothing
```go
func greet() {
	fmt.Println("Vande Mataram!")
}
``` 


## Function Returning A Single Value
```go
func add(a int, b int) int {

	add := a + b

	return add
}
```


## Function Returning Multiple Values
```go
func sumProduct(a int, b int) (int, int) {
	var add int
	var product int

	add = a + b
	product = a * b

	return add, product
}
```


## Variadic Functions

Variadic functions can be called with any number of trailing arguments. For example, `fmt.Println()` is a common variadic function.

```go
func addVariadicFunc(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
```