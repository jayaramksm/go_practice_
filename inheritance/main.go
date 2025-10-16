package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
	Animal
	age   int
	types string
}

func main() {
	a := Animal{"Buddy"}

	d := Dog{a, 27, "animal"}

	d.Speak() // Inherited method

	fmt.Println(a, d.age, d.types, d.Name, "makes a sound")
}
