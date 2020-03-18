package main

import (
	"fmt"
	"sync"
)

/*

This example solves previous race condition through a sync.Mutex

*/
var i int // i == 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	i = i + 1
	m.Unlock() // release lock
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	wg.Wait()

	fmt.Println("value of i after 1000 operations", i)

}

/*

You can test for race condition in Go using race flag while running a program
like Go run -race program.Go. Read more about race detector:

https://blog.golang.org/race-detector

*/
