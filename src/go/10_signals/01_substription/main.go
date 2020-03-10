// Source: http://bit.ly/learning-go-00006
// Author: Vincent Blanchon

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	done := make(chan bool, 1)

	s1 := make(chan os.Signal, 1)
	signal.Notify(s1, syscall.SIGINT)

	go func() {
		<-s1
		fmt.Println("The Program is going to exit...")
		done <- true
	}()

	s2 := make(chan os.Signal, 1)
	signal.Notify(s2, syscall.SIGWINCH)

	go func() {
		for {
			<-s2
			fmt.Println("The terminal has been resized.")
		}
	}()

	<-done
}
