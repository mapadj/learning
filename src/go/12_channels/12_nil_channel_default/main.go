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

func service(c chan string) {
	c <- "response"
}

func main() {
	fmt.Println("main() started")

	var chan1 chan string

	go service(chan1)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res)
	default:
		fmt.Println("No data")

	}

	fmt.Println("main() stopped", time.Since(start))

}

/*

Above program not-only ignores the case block but executes the default statement
immediately. Hence scheduler does not get time to schedule service goroutine.

But this is really bad design. You should always check a channel for nil value.

*/
