package main

import "fmt"

func main() {
	s := make([]int, 5, 11)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}