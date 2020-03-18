package main

import "fmt"

// The Transmitter
func squares(c chan int) {
	// Produce data and transmit
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	// close channel
	close(c)

	// end of goroutine
}

func main() {
	fmt.Println("main() started")

	// Create Channel
	c := make(chan int)

	// Start Transmitter Channel
	go squares(c)

	// Continue as Receiving Channel with Loop range C
	for val := range c {
		fmt.Println(val)
	}

	// End Program
	fmt.Println("main() stopped")
}
