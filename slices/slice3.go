package main

import "fmt"

func main() {
	i := []int{10, 20, 90, 70, 60}
	slice := i[:3]
	slice[2] = 900
	fmt.Println(i)
	fmt.Println(slice)

}