package main

import (
	"fmt"
	"time"
)

/*

Above program is not very useful since only default case is getting executed.

But sometimes, what we want is that any available services should respond in a
desirable time, if it doesn’t, then default case should get executed.

This can be done using a case with a channel operation that unblocks after defined
time. This channel operation is provided by time package’s After function.
Let’s see an example.

*/

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "Hello from service 1"

	fmt.Println("service1 stopped", time.Since(start))

}

func service2(c chan string) {
	time.Sleep(time.Second * 3)
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
	case <-time.After(1 * time.Second):
		fmt.Println("No Response received", time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))

}

/*

In the above program, <-time.After(2 * time.Second) unblocks after 2 seconds
returning time at which it was unblocked, but here, we are not interested in
its return value. Since it also acts like a goroutine, we have 3 goroutines
out of which this one unblocks first. Hence, the case corresponding to that
goroutine operation gets executed.

This is useful because you don’t want to wait too long for a response from
available services, where the user has to wait a long time before getting
anything from the service. If we add 10 * time.Second in the above example,
the response from service1will be printed, I guess that’s obvious now.

*/
