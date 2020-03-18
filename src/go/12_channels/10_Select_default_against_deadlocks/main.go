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
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	select {
	case res := <-chan1:
		fmt.Println("Response from channel 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from channel 2", res, time.Since(start))
	default:
		fmt.Println("No Data Available", time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))

}

/*

Since with default, select is non-blocking, the scheduler does not get a call from
main goroutine to schedule available goroutines. But we can do that manually by
calling time.Sleep. This way, all goroutines will execute and die, returning control
to main goroutine which will wake up after some time. When main goroutine wakes up,
channels will have values immediately available.

Result:

Either Message from Channel 1 or 2, but never default, since there will be a message.

*/
