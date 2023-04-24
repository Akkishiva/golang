package main

import "fmt"

var universe string = "multiverse" //global variable

func main(){

country:="india"          //local varible
{
	city:="hubli"          //local variable
	fmt.Println(country)
	fmt.Println(city)
	fmt.Println(universe)
}
fmt.Println(country)
fmt.Println(universe)
}