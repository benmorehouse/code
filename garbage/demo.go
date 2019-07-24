package main

import(
	"fmt"
	"strings"
)

func main(){
	mystring := "Hello world this is a string that i have made"
	temp := strings.Fields(mystring)
	fmt.Println(temp)
	fmt.Println(temp[2:])
}





