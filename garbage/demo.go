package main

import (
//	"os/exec"
	"bufio"
	"os"
	"log"
	"fmt"
)

var testSlice = []byte("Hello world\nthis is a testSlice")

func main(){
	file , err := os.Open("temp.txt")
	if err != nil{
		log.Fatal("Error with opening bucketLust.txt...",err)
	}
	scanner := bufio.NewScanner(file)
	var lists []string
	for scanner.Scan(){
		if scanner.Text() != "delete"{
			lists = append(lists,scanner.Text())
		}
	}
	file , err = os.Create("temp.txt")
	if err != nil{
		log.Println("This file already exists")
	}
	for _, place := range lists{
		fmt.Fprintln(file,place)
	}


	err = file.Close()

}





				// use this to go into vim and just edit around the line 9 command 
