// Source: http://bit.ly/learning-go-00004package main
// Author: Vladimir Vivien
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadFile("./planets.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", bytes)
}
