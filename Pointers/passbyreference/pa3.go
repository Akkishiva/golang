package main

import "fmt"

func modify(a map[string]int) {
	a["g"] = 4
}

func main() {
	s := make(map[string]int)
	s["a"] = 1
	s["b"] = 2
	s["c"] = 3
	fmt.Println(s)
	modify(s)
	fmt.Println(s)
}