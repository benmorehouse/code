package main

import (
	"fmt"
	"os"
)

func main(){
	before,_:=os.Getwd()
	before += "/demo.txt"
	after,_:=os.Getwd()
	after += "/demo.go"
	fmt.Println(before)
	fmt.Println(after)

}

