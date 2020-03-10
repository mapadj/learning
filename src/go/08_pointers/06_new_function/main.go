// Source: http://bit.ly/learning-go-00004
// Author: Vladimir Vivien

package main

import "fmt"

func main() {
	intPtr := new(int)
	*intPtr = 77
	type person struct{name string, age int}
	prsn := new(person)
	prsn.first = "Prince"
	prsn.age = 57
}

