// Source: http://bit.ly/learning-go-00004package main
// Author: Vladimir Vivien
package main

// Sample Pointers:
var valPtr *float32
var countPtr *int
type person struct{name string, age int}
var prsn *person
var matrix *[1024]int
var row []*int64

func main(){
	//Each pointer type is unique:
	var intPtr *int
	var int32Ptr *int32
	intPtr = int32Ptr
	// $> cannot use int32Ptr (type *int32) as type *int in assignment
}
