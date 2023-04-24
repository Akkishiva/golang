package main

import (
	"fmt"
	"reflect"
)
func main() {
	var s string
	var i int
	var b bool
	var c float64 =35.5

	fmt.Println("Enter the values")
	fmt.Scanf("%s %d %b %f",&s,&i,&b)
	fmt.Printf("%v \n",reflect.TypeOf(s))
	fmt.Printf("%v is of type %v\n",i,reflect.TypeOf(i))
	fmt.Printf("%v \n",reflect.TypeOf(b))
	fmt.Printf("%v is of type %v\n",c,reflect.TypeOf(c))
}