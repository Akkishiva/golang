package main

import "fmt"

func main() {
	// s := func(x int, y int) int {
	// 	return x * y
	// }
	// fmt.Println(s(4,5))

	s:=func(x int,y int)int{
		return x*y
	}(25,45)
	fmt.Println(s)
}

