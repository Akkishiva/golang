package main

import "fmt"

func main() {
	s:=[5]string{"shiva","vishnu","krishna","rama","hanuma"}
	for index,values := range s{
		fmt.Println(index,"<-->",values)
	}
	//default size calculator ... 
	j:=[...]int{5,10,15,20,25,30}
	for i:=0;i<len(j);i++{
		fmt.Println(j[i])
	}
	// Multidimensonal array
	k:=[3][3]int{{0,123,4},{2,5,6},{4,5,6}}
	fmt.Println(k[0][2])

}