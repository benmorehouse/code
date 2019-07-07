package main

import (
	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"
	"fmt"
	"log"
)

var testSlice = []byte("Hello world\nthis is a testSlice")
var key = []byte("Output")
var BucketName = []byte("Work")

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
			bucket := tx.Bucket(BucketName) // here what we will do is read in whats in temp and change to byte
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
		var chosenBucket string
		if bucketExists(strToByteSlice(args[0])) == false && args[0] == "new"{ // if user argument is not a bucket
			fmt.Println("Enter in a list to work with or for new list, enter new")
			fmt.Scan(&chosenBucket)
			for bucketExists(strToByteSlice(chosenBucket)){
				fmt.Println("Enter in a list to work with or for new list, enter new")
				fmt.Scan(&chosenBucket)
			}
		}

		db, err := bolt.Open("mainDatabase.db", 0600, nil)
		fmt.Println(args)

		if err != nil {
			log.Fatal(err)
		}
		defer db.Close() // will close at end of function run
		err = db.Update(func(tx *bolt.Tx) error{
			bucket, err := tx.CreateBucketIfNotExists(strToByteSlice(chosenBucket)) // tx is used to create and interact with buckets
			if err != nil {
				log.Println("blank bucket name or too long of a name")
			}
			err = bucket.Put(strToByteSlice("second"),getInput()) // getinput opens the file
			if err != nil{
				return err
			}
			return nil
		})
		if err != nil {
			log.Fatal("error in write command:",err) // this will 
		}
	},
}

