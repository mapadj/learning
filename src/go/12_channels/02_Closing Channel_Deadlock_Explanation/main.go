package main

import "fmt"

// A deadlock example

func greet(c chan string) {
	<-c
	<-c
}

func main() {
	fmt.Println("main() started")

	c := make(chan string)

	go greet(c)

	c <- "John"

	close(c)

	// Deadlock here!

	c <- "Mike"
	fmt.Println("main() stopped")

}
