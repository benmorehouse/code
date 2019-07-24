package main

import(
	"fmt"
	"os"
	"os/exec"
	"log"
	"github.com/boltdb/bolt"
	"strings"
)

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
	cmd := exec.Command("vim","-o","buffer.md")
	// the hope is that buffer.txt will have everything loaded in from the bucket
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func show_lists(db *bolt.DB) error { // this will show the list of all the lists that are in that bucket
	if db == nil { // this means if DB is nil
		log.Fatal("show_lists was given a null database")
	}

	err := db.Update(func(tx *bolt.Tx) error{
		bucket := tx.Bucket(bucketName)
		if bucket == nil{
			log.Fatal("show_lists couldnt open the bucket")
		}

		get_list := bucket.Get([]byte("show_lists"))

		if get_list == nil{
			fmt.Println("No lists yet")
			return nil
		}
		temp_list := string(get_list)

		final_list := strings.Fields(temp_list)

		fmt.Println("\tAVAILABLE LISTS\n______________________________\n")
		for _ , val := range final_list{
			fmt.Print("- ")
			fmt.Println(val)
		}
		fmt.Println("______________________________\n")
		return nil
	})
	return err
}

