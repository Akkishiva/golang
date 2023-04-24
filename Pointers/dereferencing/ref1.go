package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(s)
	ptr_s := &s
	fmt.Println(*ptr_s)
	fmt.Println(ptr_s)
	*ptr_s = "win win baby"
	fmt.Println(*ptr_s)
}