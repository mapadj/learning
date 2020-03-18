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

func main() {
	defer fmt.Println("main reached defer", time.Since(start))
	fmt.Println("main() started", time.Since(start))

	// Start Channels
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	// Prefill Buffers
	chan1 <- "Value1"
	chan1 <- "Value2"
	chan2 <- "Value1"
	chan2 <- "Value2"

	// Select only the first incoming from either channel then exit
	select {
	case res := <-chan1:
		fmt.Println("Response from channel 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from channel 2", res, time.Since(start))
	}

	// End program
	fmt.Println("main() reached end", time.Since(start))
}

/*

Explanation:

Since cases are non blocking and buffers filled you will result in a Value 1 of a random chosen channel

*/
