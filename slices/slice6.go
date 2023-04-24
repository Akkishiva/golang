package main

import "fmt"

func main() {
        arr := [5]int{10, 20, 90, 70, 60}
        slice := append(arr[:2], arr[3:]...)//appending one slice to another slice
        fmt.Println(slice)
}