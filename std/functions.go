package main

import(
	"fmt"
	"os"
	"os/exec"
	"log"
	"github.com/boltdb/bolt"
	"strings"
)

func openFile(){
	cmd := exec.Command("vim","-o","buffer.md")
	// the hope is that buffer.txt will have everything loaded in from the bucket
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func show_lists(db *bolt.DB) error {
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
		for _ , val := range final_list {
			fmt.Print("- ")
			fmt.Println(val)
		}
		fmt.Println("______________________________\n")
		return nil
	})

	return err
}

func rc_content_manip(input, new_list  string)(string){
// Need to come here, create a system that takes in content, adds marks wherever there are \n, then does fields and joins
	var marker [0]int // used to keep place of where there ~
	for i , val := range input{
		if val == "~"{
			val = "~"
			marker = append(marker,i + 2)
		}
	}
	temp_content := strings.Field(input)
	temp_content[1] = new_list + "\n\n"
	input = strings.Join(temp_content, " ")
	for _ , val := range marker{
		if input[val]!="~"{
			continue
		}else{
			input[val] = "\n"
		}
	}
	return input
}












