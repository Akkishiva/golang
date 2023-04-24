package main

import "fmt"

// arr := [5]string {}
// var arr [5]string = [5]string {}
// -length capacity
func main() {
	s := [3]int{36, 45, 25}
	fmt.Println(s) //Print array
	fmt.Println(s[0:2]) //Indexing
	arr := [...]string{"1","2","3","4","5"}//default size calculated by ...
	fmt.Println(arr)
}