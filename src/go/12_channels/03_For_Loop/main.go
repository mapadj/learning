package main

import "fmt"

// A deadlock example

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c) // close channel
}

func main() {
	fmt.Println("main() started")

	c := make(chan int)

	go squares(c)

	// Preiodic block/unblock of main goroutine until channel closes

	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "Loop broke!")
			break
		} else {
			fmt.Println(val, ok)
		}
	}

	fmt.Println("main() stopped")
}
