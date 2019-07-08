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
		var temp string
		if len(args) < 1{
			fmt.Println("Which list?")
			fmt.Scan(&temp)
		}
		desKey := strToByteSlice(temp)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close() // will close at end of function run
		err = db.Update(func(tx *bolt.Tx) error{
			bucket, err := tx.CreateBucketIfNotExists(bucketName) // tx is used to create and interact with buckets
			if err != nil {
				log.Println("blank bucket name or too long of a name")
			}
			checkKey := bucket.Get(desKey)
			if checkKey == nil{
				for checkKey==nil{
					fmt.Println("Not a valid list, please enter in an existing key")
					var temp string
					fmt.Scan(&temp)
					checkKey = bucket.Get(strToByteSlice(temp))
				} // at this point now we know that checkKey exists
			}
			desKey = checkKey

			// loop through the file and put it into one big ass string. Then push that string to the bucket
			//first we need to write whats in the key to the file
			// then we let the user manipulate
			// then we .put it back in
			content := bucket.Get(desKey) // now content is what we have 
			file , err := os.Create("buffer.txt")
			if err != nil{
				log.Println("Error:",err)
			}
			_ , err = file.Write(content)

			if err != nil{
				log.Println("Error writing file: ", err)
			}

			openFile() // this will open the file and let the user input 

			content, err = ioutil.ReadFile("buffer.txt")
			fmt.Println(string(content))

			err = bucket.Put(desKey,content)
			for err != nil{
				log.Println("error in write command",err)
			}
			return nil
		})
		if err != nil {
			log.Fatal("error in write command:",err) // this will 
		}
	},
}

var createList = &cobra.Command{
	Use: "Create",
	Short: "create a list",
	Example: "./std create work",
	Run: func(cmd *cobra.Command, args []string){ // args is gonna be what we pass through 
		db, err := bolt.Open("mainDatabase.db", 0600, nil)
		desKey := strToByteSlice(args[0])
		if err != nil{
			log.Println("Error at createlist command:",err)
		}
		defer db.Close()
	newKey := strToByteSlice(args[0])
	err = db.Update(func(tx *bolt.Tx) error{
		bucket, err := tx.CreateBucketIfNotExists(bucketName) // tx is used to create and interact with buckets
		if err != nil {
			log.Println("blank bucket name or too long of a name")
		}
		checkKey := bucket.Get(desKey)
		for checkKey != nil{
			fmt.Println("list already exists")
			var temp string
			fmt.Scan(&temp)
			checkKey = bucket.Get(strToByteSlice(temp))
		} // at this point now we know that checkKey exists
		newKey = checkKey
		err = bucket.Put(desKey,getInput()) // getinput opens the file
		for err != nil{
			log.Println("error in write command",err)
		}
		return nil
	})
		if err != nil {
			log.Fatal("error in write command:",err) // this will 
		}
	},
}

