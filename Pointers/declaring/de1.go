package main

import "fmt"

func main() {
	// Declaring the pointer

	i:=10
	var x *int = &i //Initalising variable to pointer
    var k = &i //Second way of initalising pointer to the variable
	//short hand declaration
	s:="hello"
	ptr_s:=&s
	var y *string
	fmt.Println(x,y,k,ptr_s)
}