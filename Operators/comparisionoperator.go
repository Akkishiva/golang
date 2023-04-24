package main

import "fmt"

func main() {
	var a int = 12
	var b int = 12
	fmt.Println(a <= b)

	a = 20
	fmt.Println(a <= b)

	b = 100
	fmt.Println(a <= b)

	c := 0
	fmt.Println(a <= c)

}