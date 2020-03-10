// Source: http://bit.ly/learning-go-00004package main
// Author: Vladimir Vivien
package main

// Sample Pointers:
val := flaot32(5.5)
var valPtr *float32 = &val

score := 79
scorePtr := &score

func printId(id *string){
	//---
}


func main(){
	uid := "abcd-eff-33cc-5534"
	printId(&uid)	
}
