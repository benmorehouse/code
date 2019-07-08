package main

import(
	"fmt"
	"os"
	"os/exec"
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

func doThis(command string){
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func strToByteSlice(input string)([]byte){
	var output []byte
	for i:=0;i<len(input);i++{
		output = append(output,input[i])
	}
	return output
}

func getInput()([]byte){
	// will get user input, then return the user input to be used in db
	// for now just gonna test with simple buffered scanner - will move to command line temp file afterwards
	var input string = "test string"
	return strToByteSlice(input)

}

func openFile(){
	cmd := exec.Command("vim","-o","buffer.txt")
	// the hope is that buffer.txt will have everything loaded in from the bucket
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
