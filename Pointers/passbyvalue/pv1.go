package main

import "fmt"

func isi(s int)int{
	return s
}
// Example for pass by value
func main() {
	s := 25
	fmt.Println(s)
	fmt.Println(isi(50))
}