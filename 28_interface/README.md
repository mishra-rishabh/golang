# Interface In Golang

* Go language interfaces are different from other languages.
* In Go language, the interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface.
* But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods the interface requires.
* In other words, the interface is a collection of methods as well as it is a custom type.
* The `struct` data type implements these interfaces to have method definitions for the method signature of the interfaces.


## How To Create An Interface?
**Syntax:**
```
type interface_name interface{

// Method signatures

}
```

**example:**
```go
type myInterface interface{

    // Methods
    fun1() int
    fun2() float64
}
```

**NOTE:** In Go language, it is necessary to implement all the methods declared in the interface for implementing an interface.


## Use Of Interface

You can use interface when in methods or functions you want to pass different types of argument in them just like `Println()` function. Or you can also use interface when multiple types implement same interface.