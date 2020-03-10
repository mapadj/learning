// Source: http://bit.ly/learning-go-00005
// Author: Uday Hiwarale

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	// create a main data source
	mainSrc := strings.NewReader("Hello Amazing World!") // 20 characters

	// create data source from `mainSrc` with cap of `10` bytes
	src := io.LimitReader(mainSrc, 10)

	// create a packet
	p := make([]byte, 3) // slice of length `3`

	// read `src` until an error is returned
	for {

		// read `p` bytes from `src`
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		// handle error
		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occured!", err)
			break
		}
	}

}
