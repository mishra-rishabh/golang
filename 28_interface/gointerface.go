package main

import "fmt"

type Animal interface {
	speak()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) speak() {
	fmt.Println(d.Name, "says Woof!")
}

func (c Cat) speak() {
	fmt.Println(c.Name, "says Meow!")
}

func main() {
	fmt.Println("Interface in Go")

	/* variables of type interface */
	var animalDog Animal
	var animalCat Animal

	/* initializing structs */
	dog := Dog{Name: "Spike"}
	cat := Cat{Name: "Tom"}

	/* assigning structs to interface variable, i.e., implicit implementation */
	animalDog = dog
	animalDog.speak()

	animalCat = cat
	animalCat.speak()
}
