package main

import "fmt"

func marks(S string, marks ...int) int {
	fmt.Println("Hey ",S,"Total marks is -")
	i:=0
	for _,value:=range marks{
		i+=value
	}
	return i

}

func main() {

	fmt.Println(marks("shiva",50,60,70))
}