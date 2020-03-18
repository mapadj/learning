package main

import (
	"fmt"
	"time"
)

func sqrWorker(tasks <-chan int, result chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond) //simulate blocking task
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		result <- num * num
	}
}

func main() {
	fmt.Println("[main] main() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2
	}

	fmt.Println("[main] Wrote 5 Tasks")

	// closing Tasks
	close(tasks)

	// receiving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")
}

/*

In the above program, as we know select will block the main goroutine,
the scheduler will schedule another available goroutine which is service.
But after that, it will die and the schedule has to schedule another available
goroutine, but since main routine is blocked and no other goroutines are available,
resulting in a deadlock.

*/
