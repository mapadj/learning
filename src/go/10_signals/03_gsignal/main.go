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
	signal.Notify(s1, syscall.SIGWINCH)
	signal.Ignore(syscall.SIGINT)

	go func() {
		<-s1
		fmt.Println("The terminal has been resized.")
		signal.Stop(s1)

		//This will block forever
		<-s1
		done <- true
	}()

	<-done
}

// this program cannot be interrupted with ctrl + c and will never stopped- lol
