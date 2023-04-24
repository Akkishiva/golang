package main

import "fmt"

func main() {
        var x, y float64 = 27.9, 7.0
		var v,b int = 25,5
        x -= y
        fmt.Println(x)
        x += y
        fmt.Println(x)
		x /= y
        fmt.Println(x)
		v*=b
		fmt.Println(v)
        v %= b
        fmt.Println(v)
		
}