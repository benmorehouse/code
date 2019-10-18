package main

import(
	"fmt"
	"io/ioutil"
)

func main(){
	returnVal, err := ioutil.ReadFile("/Users/benmorehouse/repositories/BMA/example.txt")
	if err != nil{
		fmt.Println("error")
	}else{
		fmt.Println("ioutile returned:",string(returnVal))
	}

}
