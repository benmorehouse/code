package main

import (
	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"
	"fmt"
	"log"
)

var bucketName = []byte("Lists")
// This needs to be fixed 
var readList = &cobra.Command{ // just displays the list 
	Use: "Read",
	Short:"Reads what is on the current list",
	Run: func(cmd *cobra.Command, args []string){
		db, err := bolt.Open("mainDatabase.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(bucketName) // here what we will do is read in whats in temp and change to byte
			if bucket == nil { // this will be nil if there is nothing at all in the bucket!
				fmt.Println("Empty")
			}
			val := bucket.Get(strToByteSlice("second")) //everything is mapped to a key. This returns the map of that key which is data
			fmt.Println(string(val))
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

var writeList = &cobra.Command{ // appends to the end of the bucket
// store some data that is written by the user
	Use: "Write",
	Short:"Write to the current list",
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
			for checkKey == nil{
				fmt.Println("Not a valid list, please enter in an existing key")
				var temp string
				fmt.Scan(&temp)
				checkKey = bucket.Get(strToByteSlice(temp))
			} // at this point now we know that checkKey exists
			desKey = checkKey
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

