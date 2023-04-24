//fun ,func_name>(<parama1>,<parama2>,<parama3> ...type)<return_type>
//func sum(num ...int)int

package main

import "fmt"

func addnum(num ...int)int{
	i:=0
	for _,value := range num{
		i+=value
	}
	return i
}

func main(){
fmt.Println(addnum())
fmt.Println(addnum(10))
fmt.Println(addnum(10,20))
fmt.Println(addnum(10,20,30,40,50))

}