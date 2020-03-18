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
	}

	fmt.Println("main() stopped", time.Since(start))

}

/*

From the above result, we can see that select (no cases) means that select
statement is virtually empty because cases with nil channel are ignored.
But as empty select{} statement blocks the main goroutine and service goroutine
is scheduled in its place, channel operation on nil channels throws chan send
(nil chan) error. To avoid this, we use default case.

*/
