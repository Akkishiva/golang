package main

import "fmt"
func main() {

	var i,j int = 10,20

	fmt.Println((i<100)&&(j<200)) //AND
	fmt.Println((i<100||j<10))   //OR
	fmt.Println(!(i<j))          //Unary
	
}