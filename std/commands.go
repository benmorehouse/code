package main

import (
	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

var bucketName = []byte("Lists")
// This needs to be fixed 

var writeList = &cobra.Command{ // appends to the end of the bucket
// store some data that is written by the user
	Use: "open",
	Short:"Open the current list",
	// Args: this is how you can add in arguments into your cobra commands. You should be able to pass in
	Run: func(cmd *cobra.Command, args []string){ // args is gonna be what we pass through 
		// open tmp, let user input, then read line for line and add into bucket
		db, err := bolt.Open("mainDatabase.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close() // will close at end of function run

		var temp string
		if len(args) != 1{
			show_lists(db)
			fmt.Scan(&temp)
		}else{
			temp = args[0]
		}
		chosen_list_key := []byte(temp)

		err = db.Update(func(tx *bolt.Tx) error{
			bucket := tx.Bucket(bucketName)
			if bucket == nil {
				log.Fatal("Not able to open the bucket with all the lists:")
			}

			content := bucket.Get([]byte(chosen_list_key)) // this will return what is in the list
			if content == nil{ // this means that chose_list_key is not a key within the lists bucket
				for content == nil{
					fmt.Println("Not a valid list, please enter in an existing list")
					show_lists(db)
					fmt.Scan(&chosen_list_key)
					content = bucket.Get([]byte(chosen_list_key))
				} // at this point now we know that checkKey exists
			}

			// loop through the file and put it into one big ass string. Then push that string to the bucket
			//first we need to write whats in the key to the file
			// then we let the user manipulate
			// then we .put it back in
			file , err := os.Create("buffer.md")
			if err != nil{
				log.Println("Error opening file in writelist:",err)
			}
			_ , err = file.Write(content)

			if err != nil{
				log.Println("Error writing file in writelist: ", err)
			}

			openFile() // this will open the file and let the user input 

			content, err = ioutil.ReadFile("buffer.md")

			err = bucket.Put(chosen_list_key,content)

			for err != nil{
				log.Println("error in write command on line 77",err)
			}
			return nil
		})
		if err != nil {
			log.Fatal("error in write command on line 82:",err) // this will return if the database isnt open?
		}
	},
}

var createList = &cobra.Command{
	Use: "Create",
	Short: "create a list",
	Example: "./std create work",
	Run: func(cmd *cobra.Command, args []string){ // args is gonna be what we pass through 
		db, err := bolt.Open("mainDatabase.db", 0600, nil)
		if err != nil{
			log.Println("Error opening database at createlist command:",err)
		}
		defer db.Close()

		var temp string
		if len(args) != 1{ // this means that they didnt enter in any sort of list to add into the bucket
			fmt.Println("What is your new list called")
			fmt.Scan(&temp)
		}else{
			temp = args[0]
		}

		chosen_list_key := []byte(temp) // this is the new list name within the bucket lists

		err = db.Update(func(tx *bolt.Tx) error{ // error happening here 
			bucket , err := tx.CreateBucketIfNotExists(bucketName) // this is going into the database and accessing one bucket
			if err != nil {
				log.Println("blank bucket name or too long of a name")
			}

			checkKey := bucket.Get(chosen_list_key) //  if the list already exists then it will come up 
			// this will return nil if this doesnt exist which is what we want 
			for checkKey != nil{ // this means that they entered something that already exists
				fmt.Println("list already exists:")
				show_lists(db)
				fmt.Scan(&temp)
				checkKey = bucket.Get([]byte(temp))
				chosen_list_key =[]byte(temp)
			} // at this point now we know that checkKey exists

			err = bucket.Put(chosen_list_key,[]byte(""))// creates the new bucket with nothing in it 

			if err != nil{
				log.Println("Unable to add new list in bucket in write command")
				fmt.Println("chosen_list_key: ",chosen_list_key)
				fmt.Println("value of newList:",bucket.Get(chosen_list_key))
			}

			show_list_temp := bucket.Get([]byte("show_lists"))

			if show_list_temp == nil{
				// this means that show_list has yet to be created within the database
				bucket.Put([]byte("show_lists"),[]byte(chosen_list_key))
			}else{
				bucket.Put([]byte("show_lists"),[]byte(string(chosen_list_key) + string(show_list_temp)))
			}

			return nil
		})

		if err != nil {
			log.Fatal("error in create command:",err) // this will 
		}
	},
}

