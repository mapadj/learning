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

	time.Sleep(time.Second)

	select {
	case res := <-chan1:
		fmt.Println("Response from channel 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from channel 2", res, time.Since(start))
	default:
		fmt.Println("No Response received", time.Since(start))
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
