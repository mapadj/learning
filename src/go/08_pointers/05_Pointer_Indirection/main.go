// Source: http://bit.ly/learning-go-00004package main
// Author: Vladimir Vivien

package main

import "fmt"

func main() {
	a := 71
	print(&a)
}

func print(val *int) {
	// Pointer is simply an adress aof an actual value in memory
	fmt.Println(val)
	//To access it simply apply the death star to it *
	fmt.Println(*val * 2000000000)

	fmt.Println("PENIS")
}
