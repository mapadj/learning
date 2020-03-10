
// Source: http://bit.ly/learning-go-00005
// Author: Uday Hiwarale

package main

import (
	"fmt"
	"strings"
	"io"
)

func main() {
	
	// create data source
	src := strings.NewReader("Hello Amazing World!") // 20 characters

	// create buffer of length 14
	buf := make([]byte, 14)

	// call 1: read from `src`
	bytesRead1, err1 := io.ReadFull( src, buf )
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead1, buf[:bytesRead1], err1)

	// call 2: read from `src`
	bytesRead2, err2 := io.ReadFull( src, buf )
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead2, buf[:bytesRead2], err2)

	// call 3: read from `src`
	bytesRead3, err3 := io.ReadFull( src, buf )
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead3, buf[:bytesRead3], err3)

}