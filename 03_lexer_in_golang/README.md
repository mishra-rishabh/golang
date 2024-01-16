## Lexer In Golang

### Code With Semicolon
```go
package main

import "fmt"

func main() {
	fmt.Println("Vande Mataram Bhai!!");    // with ;
}
```

### Code Without Semicolon
```go
package main

import "fmt"

func main() {
	fmt.Println("Vande Mataram Bhai!!")     // without ;
}
```

* There are some places where you can put semicolon because you're allowed to do so but there are places you shouldn't be putting up semicolons. 
* But, one thing to notice is that, since we've installed golang intellisense (go vscode extension) by the golang team itself, it automatically removes the semicolon from the code but there are places where it is not automatically going to remove it like for loop, a special syntax of if-loop.
* So, what is going on and how is it removing the semicolon and still the code is working fine. The actual answer is, we should be putting the semicolon, but eventually the golang team realized there are some places where you shouldn't be putting up so that lexer can come and put up this.

## Read More
https://go.dev/ref/spec#Semicolons