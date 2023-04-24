package main

import "fmt"

func sum(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	sum, diff := sum(9, 10)
	fmt.Println(sum,diff)

}