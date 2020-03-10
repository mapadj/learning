// Source: http://bit.ly/learning-go-00004package main
// Author: Vladimir Vivien
package main

// What is a composite type literal expression?

// These are three composite type literal expressions:
type person struct{name string, age int}
prsn := &person{"Prince", 57}
pair := &[2]string{"left-sock", "right-sock"}

// And you can apply ampersand on them and return their adress.

func main(){
	fmt.Println(prsn)
	fmt.Println(pair)
	fmt.Println("penis")	
}
