package main

import "fmt"

func sum(a int, b int) (sum int,diff int) {//named return values
	sum = a + b
	diff = a - b
	return 
}

func main() {
	sum, diff := sum(9, 10)
	fmt.Println(sum,diff)

}