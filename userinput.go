package main

import "fmt"

func main() {
	var s string 
	var i int

	fmt.Printf("Enter the name and value:")
	fmt.Scanf("%s %d",&s,&i)
	fmt.Printf("%s %d",s,i)
}