package main

import (
	"fmt"
	"time"
)

/*

Like for{} empty loop, an empty select{} syntax is also valid but there is a gotcha.
As we know, select statement is blocked until one of the cases unblocks, and since
there are no case statements available to unblock it, the main goroutine will block
forever resulting in a deadlock.

*/

var start time.Time

func init() {
	start = time.Now()
}

func service() {
	fmt.Println("Hello from Service", time.Since(start))
}

func main() {
	fmt.Println("main() started", time.Since(start))

	go service()

	select {}

	fmt.Println("main() stopped", time.Since(start))
}

/*

In the above program, as we know select will block the main goroutine,
the scheduler will schedule another available goroutine which is service.
But after that, it will die and the schedule has to schedule another available
goroutine, but since main routine is blocked and no other goroutines are available,
resulting in a deadlock.

*/
