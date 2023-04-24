package main

import "fmt"

func main() {
	s := map[string]int{}

	s["shiva"] = 1
	s["Rama"]=2
	value,found:=s["shiva"]
	fmt.Println(value,found)
	value,found=s["v"] //To check if element is present or not
	fmt.Println(value,found)
	delete(s,"Rama")//delete map
	fmt.Println(s)
	for key,value := range s{  //print key value pair using for loop
		fmt.Println(key,"<--->",value)
	}

	fmt.Println(s)
}