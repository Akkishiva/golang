package main

import "fmt"

func main() {
        arr := [6]int{10, 20, 90, 70, 60,600}
        slice := arr[:3]//capacity of slice is same as that of array
        fmt.Println(cap(slice))
		fmt.Println(slice)
        new_slice := append(slice, 100, 200,300,800)//source,element to be appended
        fmt.Println(cap(new_slice))
		fmt.Println(new_slice)
}