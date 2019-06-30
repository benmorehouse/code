package main

import(
	"fmt"
)
/*
type list struct{
	// Use the get function with this to be able to go in and get the data 
	key []byte
	content []byte
}*/
func showLists(lists []list){ // an array of lists 
	// shows all of the lists that we have thus far
	fmt.Println("Available Lists\n____________________________________")
	for i:=0;i<len(lists);i++{
		fmt.Println(string(lists[i].key))
	}
}
