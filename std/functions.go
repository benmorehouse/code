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

	err := db.update(func(tx *bolt.tx) error{
		bucket := tx.bucket(bucketname)
		if bucket == nil{
			log.fatal("show_lists couldnt open the bucket")
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


