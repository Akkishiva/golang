package main

import "fmt"

func main() {
	var a,b int = 5,10
	var c,d string = "foo","bar"
    

	fmt.Println(a+b)
	fmt.Println(c+d)
	fmt.Println(a-b)
	fmt.Println(a/b)
	fmt.Println(a%b)
	a--
	fmt.Println(a)
	b++
	fmt.Println(b)

}