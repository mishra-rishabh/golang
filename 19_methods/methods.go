package main

import "fmt"

type Rectangle struct {
	width  int
	length int
}

func (rect Rectangle) area() int {
	return (rect.length * rect.width)
}

func (rect Rectangle) perimeter() int {
	return (2 * (rect.length + rect.width))
}

func main() {
	fmt.Println("Methods in go")

	rect := Rectangle{length: 5, width: 9}

	fmt.Println("Area: ", rect.area())
	fmt.Println("Perimeter: ", rect.perimeter())
}
