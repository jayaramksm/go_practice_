package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	hello := User{"hello", "hello@go.dev", true, 16}
	fmt.Println(hello)
	fmt.Printf("hello details are: %+v\n", hello)
	fmt.Printf("Name is %v and email is %v.", hello.Name, hello.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
