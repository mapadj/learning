package main

import (
	"fmt"
	"time"
)

// select has only non blocking matchcases fifo style

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	c <- "Hello from service 1"

	fmt.Println("service1 stopped", time.Since(start))

}

func service2(c chan string) {
	c <- "hello from service 2"

	fmt.Println("service2 stopped", time.Since(start))

}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from channel 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from channel 2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))

}
