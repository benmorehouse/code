package main

import(
	"fmt"
	"os"
	"os/exec"
	"log"
	"bufio"

	"github.com/boltdb/bolt"
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

func bucketExists(input []byte)(bool){
	// is going to go through bucket text file, and given user input, will tell if it exists
	file , err := os.Open("bucketLust.txt")
	if err != nil{
		log.Fatal("Error with opening bucketLust.txt...",err)
	}
	defer file.Close() // closes file when function is ready to end (excluding os.Exit(1))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == string(input){
			return true // this bucket exists
		}
	}
	return false // this bucket was never found
}

func showKeys(){
	file , err := os.Open("bucketLust.txt")
	if err != nil{
		log.Fatal("Error with opening bucketLust.txt...",err)
	}
	scanner := bufio.NewScanner(file)
	fmt.Println("Available lists\n_________________________________________")
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func deleteKey(input []byte,curBucket *Bucket)error{
	db, err := bolt.Open("mainDatabase.db", 0600, nil)
	defer db.Close()
	if err!=nil{
		log.Println("deleteKey open database failed with error message:",err)
	}
	err = .Delete(input)
	if err == nil{ // this means that if the input was not in the db, then it returned something
		Println("error: list doesnt exist")
		return err
	}
	file , err := os.Open("bucketLust.txt")
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
	file , err = os.Create("bucketLust.txt")
	if err != nil{
		log.Println("This file already exists")
	}
	for _, place := range lists{
		fmt.Fprintln(file,place) // will write in the lists now but without the one that we wanted 
	}

	err = file.Close()

}

func addKey(input []byte)error{
	db , err := bolt.Open("mainDatabase.db",0600,nil)
	defer db.Close()
	if err!=nil{
		log.Println("deleteKey open database failed with error message:",err)
	}
}
