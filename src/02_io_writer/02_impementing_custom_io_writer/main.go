package main

import "fmt"

type chanWriter struct {
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{make(chan byte, 1024)}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}
package main

import "fmt"

type chanWriter struct {
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{make(chan byte, 1024)}
}

// here we have our output Function which returns a stream of bytes
func (w *chanWriter) Chan() <-chan byte{
	return w.ch
}


// The Close Message closes the Channels and prevents deadlocks!
func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}

func main() {
	writer := newChanWriter()
	go func() {
		//here we start a new go routine and after we our work we close the channel with the defer message and we are safe!!!!
		defer writer.Close()
		writer.Write([]byte("Sesam, "))
		writer.Write([]byte("stream me like a bitch!"))
	}()

	for c := range writer.Chan() {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}


