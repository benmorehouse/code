package main

import(
	"io/ioutil"
	"os"
	"bufio"
	"github.com/boltdb/bolt"
)

type cron struct{ // these conditions will be written to a file and held
	filename string
	condition int // 0 if initial state with no comments, 1 if it is initialized to something and non-binary if it is broken
	condition_name string
}

//This will reset cron to the initial stance that it was supposed to be before... call this many a times!

func reset_cron_conditions()error{
	db, err := bolt.Open("mainDatabase.db", 0600, nil)
	defer db.Close()
	reset_val := "cron.txt\n0\nINITIAL_CONDITION\n"
	db.Update(func(tx *bolt.Tx) error{
		cron_conditions := "cron_conditions"
		bucket, err := tx.CreateBucketIfNotExists([]byte(cron_conditions)) // tx is used to create and interact with buckets 
		if err != nil{
			log.Fatal("Broken function: reset_cron ... error:",err)
		}
		key := "cron_conditions"
		err = bucket.Put([]byte(key), []byte(reset_val)) // this will put the cron conditions back to normal
		if err != nil{
			log.Fatal("Broken function: reset_cron ... error:", err)
		}
		// there is a bucket with m
	}()
	return err
}

func (scheduler cron) update_cron_conditions(reset_val string)error{ // this will be in the form above
	db, err := bolt.Open("mainDatabase.db", 0600, nil)
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error{
		cron_conditions := "cron_conditions"
		bucket, err := tx.CreateBucketIfNotExists([]byte(cron_conditions)) // tx is used to create and interact with buckets 
		if err != nil{
			log.Fatal("Broken function: update_cron ... error:",err)
		}
		key := "cron_conditions"
		err = bucket.Put([]byte(key), []byte(reset_val)) // this will put the cron conditions back to normal
		if err != nil{
			log.Fatal("Broken function: update_cron ... error:", err)
		}
	}()
	return err
}

func (scheduler cron) parse_cron()(error){
	// this will go in and make the neccessary edits to cron like we wish 
	// pass in the file flag which we will use as t
	file ,err = os.Open("cron.txt")
	db, err := bolt.Open("mainDatabase.db", 0600, nil)
	if err != nil{
		log.Fatal("There is an error with parse_cron()")
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error{
		var cron_scan = [4]bufio.NewScanner(file) // an array of scanners for each of the goroutines
		var cron_temp = [4]string
		for i := range cron_temp{
			cron_temp[i] = ""
		}
		newscraper := scheduler.condition_name //updating tool... this is what we use as newscraper
		for i := range cron_channel{
			cron_channel[i] := make(chan string) // this will make blocking channels
		}

	// this is the first scanner. It scans the initial lines up to and including all of the imports
		go func(){
			for i:=0;i<13;i++{
				cron_scan[0].Scan()
				cron_temp[0] += cron_scan[0].Text() + "\n"
			} // will scan to line 14
			// then add in line fourteen the /*
			cron_temp[0] += "/* \n"
			cron_scan[0].Scan()
			for cron_scan[0].Text() != ""{ // this will scan until the whitespace 
				cron_scan[0].scan()
				cron_temp[0] += cron_scan[0].Text() + "\n"
			}
			cron_temp[0] +="*/ \n"
			cron_temp[0] += "github.com/overlooked-incorporated/backend/scraper/" + scheduler.condition_name + "\n"
			for i:=0; i<2; i++{
				cron_scan[0].scan()
				cron_temp[0] += cron_scan[0].Text() + "\n"
			}

			first_bucket, err := tx.CreateBucketIfNotExists([]byte("first")) //first is the first of three goroutines

			if err != nil{
				log.Fatal("Error at the first goroutine of parse_cron:",err)
			}

			key := "first"
			err = first_bucket.Put([]byte(first),[]byte(cron_temp[0]))

			if err != nil{
				log.Fatal("Error at the first goroutine of parse_cron:",err)
			}
			cron_channel[0] <- "first goroutine finished"
		}()
	// second goroutine
		go func(){
			//space count will scan the buffer all the way down to where we start second goroutine
			space_count := 0
			for space_count != 5{
				cron_scan[1].Scan()
				if cron_scan[1].Text() == ""{
					space_count++
				}else{
					continue
				}
			}

			for i:=0;i<21;i++{
				cron_scan[1].Scan()
				cron_temp[1] += cron_scan[1].Text() + "\n"
			}
			cron_scan[1].Scan()
			cron_temp[1] += cron_scan[1].Text() + "\n" + "/*\n"
			for cron_scan[1].Text() != ""{
				cron_temp[1] += cron_scan[1].Text()
				cron_scan[1].Scan()
			}

			cron_temp[1] += "*/\n"

			temp := scheduler.condition_name + " := " + scheduler.condition_name
			temp += ".New(logger, httpClient, db, fakebox, retrier)\n"

			cron_temp[1] += temp

			second_bucket, err := tx.CreateBucketIfNotExists([]byte("second")) //first is the first of three goroutines

			if err != nil{
				log.Fatal("Error at the second goroutine of parse_cron:",err)
			}

			key := "second"
			err = second_bucket.Put([]byte(second),[]byte(cron_temp[1]))

			if err != nil{
				log.Fatal("Error at the second goroutine of parse_cron:",err)
			}

			cron_channel[1] <- "second goroutine finished"
		}()

		go func(){
			space_count := 0
			for space_count != 9{
				cron_scan[2].Scan()
				if cron_scan[2].Text() == ""{
					space_count++
				}else{
					continue
				}
			}
			cron_scan[2].Scan()

			for cron_scan[2].Text() != ""{
				cron_temp[2] += cron_scan[2].Text() + "\n"
				cron_scan[2].Scan()
			}

			cron_temp[2] += "\t\t" + scheduler.condition_name + ",\n"

			third_bucket, err := tx.CreateBucketIfNotExists([]byte("third")) //first is the first of three goroutines

			if err != nil{
				log.Fatal("Error at the third goroutine of parse_cron:",err)
			}
			key := "third"
			err = third_bucket.Put([]byte(key),[]byte(cron_temp[2]))

			if err != nil{
				log.Fatal("Error at the third goroutine of parse_cron:",err)
			}

			cron_channel[2] <- "third goroutine finished"
		}()

		go func(){
			space_count := 0
			for space_count != 10{
				cron_scan[3].Scan()
				if cron_scan[3].Text() == ""{
					space_count++
				}else{
					continue
				}
			}
			for cron_scan[3].Scan(){
				cron_temp[3] += cron_scan[3].Text() + "\n"
				cron_scan[3].Scan()
			}

			fourth_bucket, err := tx.CreateBucketIfNotExists([]byte("fourth")) //first is the first of three goroutines

			if err != nil{
				log.Fatal("Error at the fourth goroutine of parse_cron:",err)
			}
			key := "fourth"
			err = fourth_bucket.Put([]byte(key),[]byte(cron_temp[3]))

			if err != nil{
				log.Fatal("Error at the fourth goroutine of parse_cron:",err)
			}

			cron_channel[3] <- "fourth goroutine finished"
		}()

		// wait for all the goroutines to return and then 
		new_cron , err := tx.CreateBucketIfNotExists([]byte("final"))
		<-cron_channel[0]
		<-cron_channel[1]
		<-cron_channel[2]
		<-cron_channel[3]
		var temp string
		for i := range cron_temp{
			temp += cron_temp[i]
		}
		key := "final"
		final_err := new_cron.Put([]byte(key),[]byte(temp))
		return final_err
	}()
}


func (scheduler cron) reset_cron() error{ // this will be used to uncomment cron when it runs
}




