// Source: https://bit.ly/go-learning-00005
package main

import (
	"fmt"
	"io"
)
// io.Pipe creates a synchronous pipe between io.Reader and io.Writer.

// When some data is written to the io.Writer it is available to 
// be read by io.Reader instantly.

// func Pipe() (*PipeReader, *PipeWriter)



// Pipe returns pointer *io.PipeReader and *io.PipeWriter ObjectPointer

// we read from the source, and we write to the target.
source, destination := io.Pipe()

// destination.Read()

// we can read as often as there will be data

// current goroutine blocks when no data comes

// go schedules another go routine that may write data
// destination.Write(data) 

// each Write blocks the current go routine.



// Each Read() should get all available data, but must not

// we can make multiple calls to Read() until all data is extracted

// thats when current goroutine blocks to schedule other routines
// with a possible write to get additional data

// destination.Close() closes write operations on destination

// read operations on closed destination will return 0 bytes read 
// and io.EOF error

// if source is closed for any read using source.Close(), then any
// destination Write will return io.ErrClosedPipe error



func main() {

	// create a pipe
	src, dst := io.Pipe()

	// start goroutine that writes data to `dst`
	go func() {
		dst.Write([]byte("DATA_1")) // write and block
		dst.Write([]byte("DATA_2")) // write and block
		dst.Close() // indicate EOF
	}()
	
	// data transfer packet
	packet := make( []byte, 6 )
	
	// read from `src`
	bytesRead1, err1 := src.Read(packet)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead1, packet, err1)

	// read from `src`
	bytesRead2, err2 := src.Read(packet)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead2, packet, err2)

	// read from `src`
	bytesRead3, err3 := src.Read(packet)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead3, packet, err3)

}