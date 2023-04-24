package main

import "fmt"

func modify(i1 *int) {
	 *i1=25

}

func main() {
	i := 10
	fmt.Println(i)
	modify(&i)
	fmt.Println(i)

}