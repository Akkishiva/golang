package main

import "fmt"

func main() {

	name:="shiva"
	age:=25
	gender:="Male"

	fmt.Print("My name is ", name," my age is ", age," my gender is ", gender,"\n") //easy typed
	fmt.Println(name,age,gender) //newline
	fmt.Printf("My name is %v any age is %d and gender is %v",name,age,gender) //old method format specifier

}