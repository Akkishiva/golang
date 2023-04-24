package main

import "fmt"

func main() {
 
	var s string
	var i int
	var b bool
	 var c float32 = 25.45 

	fmt.Println("Enter the values")
	fmt.Scanf("%s %d %b %0.1f",&s,&i,&b)
	fmt.Printf("%v is of type %T\n",s,s)
	fmt.Printf("%v is of type %T\n",i,i)
	fmt.Printf("%v is of type %T\n",b,b)
	fmt.Printf("%v is of type %T\n",c,c)
	
}