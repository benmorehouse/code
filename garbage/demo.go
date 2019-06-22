package main

import (
	"fmt"
)

func main(){
/*	directory,_:=os.Getwd()
	fmt.Println(directory)
	directory+="\b\b\b"
	directory+="go"
*/
	string1 :="hello"
	string2 :="hellowworld"
	var strings[]string
	strings = append(strings,string1)
	strings = append(strings,string2)
	fmt.Println(len(strings))
}




