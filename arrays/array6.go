package main

import "fmt"

func main() {
        arr := []string{"a", "b", "c"}
        for index, element := range arr {  //range will iterate all over the array
                fmt.Println(index, "->", element)
        }
}