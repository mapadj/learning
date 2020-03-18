package main

import "fmt"

/*

Generator Pattern

*/
var i int // i == 0

func fib(length int) <-chan int {

	// make buffered channel

	c := make(chan int, length)

	// run generation concurrently
	go func() {
		for i, j := 0, 1; i < length; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()

	// return channel
	return c
}

func main() {

	// read 10 fibonacci numbers from channel returned by 'fib' function

	for fn := range fib(10) {
		fmt.Println("Current fibonacci number is", fn)
	}
}
