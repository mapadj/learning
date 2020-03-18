package main

import (
	"fmt"
	"sync"
	"time"
)

/*

Let’s imagine a condition where you need to know if all goroutines finished their
job. This is somewhat opposite to select where you needed only one condition to be
true, but here you need all conditions to be true in order to unblock the main
goroutine. Here the condition is successful channel operation.

WaitGroup is a struct with a counter value which tracks how many goroutines were
spawned and how many have completed their job. This counter when reaches zero,
means all goroutines have done their job.

*/

var start time.Time

func init() {
	start = time.Now()
}

func service(wg *sync.WaitGroup, instance int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done() // decrement counter
}

func main() {
	fmt.Println("main() started", time.Since(start))

	var wg sync.WaitGroup // create waitgroup (empty struct)

	for i := 1; i < 3; i++ {
		wg.Add(1) // increment counter
		go service(&wg, i)
	}

	wg.Wait() // blocks here
	fmt.Println("main() stopped", time.Since(start))
}

/*

In the above program, as we know select will block the main goroutine,
the scheduler will schedule another available goroutine which is service.
But after that, it will die and the schedule has to schedule another available
goroutine, but since main routine is blocked and no other goroutines are available,
resulting in a deadlock.

*/
