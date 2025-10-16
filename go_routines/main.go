package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	ch <- "Hello Jayaram"

	fmt.Println("Hello from goroutine!")
}

func main() {

	ch := make(chan string, 2)

	var wg sync.WaitGroup

	wg.Add(2)

	go sayHello(&wg, ch)

	msg := <-ch

	fmt.Println(msg)

	fmt.Println(msg)

	go sayHello(&wg, ch)

	wg.Wait() // Waits for both goroutines to finish fmt.Println("Main function done.")
}
