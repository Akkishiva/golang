package main

import (
	"fmt"
	"strconv"
)		 



func main() {
	const i1 int = 56
	const i2=26
	const i3="shiva"
	var f float32  =25.5
	var i int = int(f)
	var s string=strconv.Itoa(i)
	var s1 string = "200abc"
	i,err :=strconv.Atoi(s1)
	fmt.Printf("%v is %T\n",i,i)
	fmt.Printf("%v is %T\n",err,err)
	fmt.Println(f,i)
	fmt.Printf("%q\n",s)
	fmt.Println(i1,i2,i3)




}