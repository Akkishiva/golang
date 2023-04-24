package main

import "fmt"

// Multidimensional arrays 2-d array
func main() {
        arr := [5][2]string{{"a"}, {"b"}, {"c"}}
        fmt.Println(arr[0][0])
        fmt.Println(arr[1][1])
        fmt.Println(arr[2][0])
}