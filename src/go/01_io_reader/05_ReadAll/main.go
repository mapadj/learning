// Source: http://bit.ly/learning-go-00005
// Author: Uday Hiwarale

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	// create data source
	src := strings.NewReader("Hello Amazing World!")

	// read all data from `src`
	data, _ := ioutil.ReadAll(src)

	// print `data`
	fmt.Printf("Read data of length %d : %s\n", len(data), data)

}
