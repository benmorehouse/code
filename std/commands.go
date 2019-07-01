package main

import (
	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"
	"fmt"
	"log"
)

var testSlice = []byte("Hello world\nthis is a testSlice")
var key = []byte("Output")

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
			bucket := tx.Bucket(testSlice)
			if bucket == nil {
				log.Fatal(err)
			}
			val := bucket.Get(key) //everything is mapped to a key. This returns the map of that key which is data
			fmt.Println(string(val))
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

var writeList = &cobra.Command{ // appends to the end of the bucket
// store some data
	Use: "Write",
	Short:"Write to the current list",
	Run: func(cmd *cobra.Command, args []string){
		db, err := bolt.Open("mainDatabase.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(testSlice)
			if err != nil {
			    return err
			}

			err = bucket.Put(key,testSlice)
			if err != nil{
				return err
			}
			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}
