package main

import (
	"fmt"
	"sync"
)

/*

FAN IN & FAN OUT

fan in:
multiplexing strategy where the inputs of several channels are compined to produce an output channel.

fan out:
demultiplexing strategy where a single channel is split into multiple channels

*/
func getInputChan() <-chan int {
	// make return channel
	input := make(chan int, 100)

	// sample numbers
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// run goroutine
	go func() {
		for num := range numbers {
			input <- num
		}

		// close channel once all numbers are sent to channel
		close(input)
	}()

	return input
}

// returns a channel which returns square of numbers
func getSquareChan(input <-chan int) <-chan int {
	//make return channel
	output := make(chan int, 100)

	//run goroutine
	go func() {
		//push squares until input channel closes
		for num := range input {
			output <- num * num
		}

		//close output channel once for loop finishes
		close(output)
	}()

	return output
}

// returns a merged channel of outputsChan channels
// this produce fan-in channel
// this is veriadic function

func merge(outputsChan ...<-chan int) <-chan int {
	// Create a WaitGroup
	var wg sync.WaitGroup

	// make return channel
	merged := make(chan int, 100)

	// increase counter to number of channels 'len(outputChan)'
	// as we will spawn number of goroutines equal to number if channels received

	wg.Add(len(outputsChan))

	// function that accept a channel(which sends square numbers)
	// to push numbers to merged channel
	output := func(sc <-chan int) {
		// run until channel (square numbers sender) closes
		for sqr := range sc {
			merged <- sqr
		}
		// once Channel (square numbers sender) closes,
		// call Done on WaitGroup to devrement counter
		wg.Done()
	}

	// run above output function as goroutine n numbers of times
	// where n is equal to number of channels received as argument the function
	// here we are using for range loop on outputchan hence no need to manually
	// tell n

	for _, optChan := range outputsChan {
		go output(optChan)
	}

	// run goroutine to close merged channel once done
	go func() {
		//wait until WaitGrpoup finishes
		wg.Wait()
		close(merged)

	}()

	return merged
}

func main() {
	// step 1:
	// get input numbers channel
	// by calling getInputChan function, it runs a goroutine which sends number
	// to returned channel

	chanInputNums := getInputChan()

	// step 2:
	// `fan-out` square operations to multiple goroutines
	// this can be done by calling `getSquareChan` function multiple times where
	// individual function call returns a channel which sends square of numbers
	// provided `chanInputNums` channel
	// `getSquareChan` function runs goroutines internally where squaring operation
	// is ran concurrently

	chanOptSqr1 := getSquareChan(chanInputNums)
	chanOptSqr2 := getSquareChan(chanInputNums)

	// step 3:
	// fan-in (combine) `chanOptSqr1` and `chanOptSqr2` output to merged channel.
	// this is achieved by calling `merge` function which takes multiple channels
	// as arguments and using `WaitGroup` and multiple goroutines to receive square
	// number, we can send square numbers to `merged` channel and close it

	chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)

	// step 4: let's sum all the squares from 0 to 9 whcih should beabout285
	// this is done by using for range loop on chanMergedSqr

	sqrSum := 0

	// run until chanMergedSqr or merged channel closes
	//that happens in merge function when all goroutines pusheing to merged
	// channel finishes (check line 96 and 100)

	for num := range chanMergedSqr {
		sqrSum += num
	}

	// step 5: print sum when above 'for loop' is done executing which is after
	// 'chanMergedSqr' channel closes

	fmt.Println("Sum of squares between 0-9 is", sqrSum)
}
