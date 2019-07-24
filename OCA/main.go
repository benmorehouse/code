package main

import(
	"fmt"
	"flag"
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("                  ***Welcome to Overlooked Cron Automation***\n")
	fmt.Println("/***********************************************************************/\n")
	fmt.Println("This is a tool that will be used to automate the tedius cron changes")
	fmt.Println("/***********************************************************************/\n")

	new_scraper := flag.String("ns", "", "user will input this to tell what the name of the new scraper is")
	cron_flag := flag.String("cron","","user will input whether to comment or uncomment")
	flag.Parse()
	if *new_scraper == "" && *cron_flag == ""{
		log.Fatal("ERROR: NO SCRAPER OR CONDITION ENTERED")
	}

	var scheduler cron
	db, err := bolt.Open("mainDatabase.db", 0600, nil)
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error{
		bucket_scan := db.Bucket([]byte("cron_conditions"))
		if bucket_scan == nil{
			fmt.Println("ERROR: CRON_CONDITIONS NOT SET IN DATABASE...RESETTING CRON")
			err = reset_cron_conditions()
			bucket_scan = db.Bucket([]byte("cron_conditions"))
		}

		cron_conditions := string(bucket_scan.Get([]byte("cron_conditions")))
		cron_conditions = strings.Field(cron_conditions)
		scheduler.filename = cron_conditions[0]
		scheduler.condition = cron_conditions[1]
		scheduler.condition_name = cron_conditions[2]

		*new_scraper = strings.TrimSpace(*new_scraper)
		*cron_flag = strings.ToLower(*cron_flag)

		if scheduler.condition == 0{
			// this means that it is in default/unedited mode and that the user is 
			if *new_scraper == ""{
				log.Println("ERROR: NO SCRAPER ENTER")
			}else if *cron_flag
				// there is a scraper
				scheduler.condition = 1
				scheduler.condition_name = *new_scraper
				err = scheduler.Parse()// this will parse the cron file using the data that is within scheudler
				if err != nil{
					log.Println("ERROR: PARSE FUNCTION FAILED IN MAIN")
				}
			}
		}else if scheduler.condition == 1{ // cron is already commented out and everything 
			if *new_scraper != ""{
				log.Println("ERROR: CRON IS COMMENTED. NEEDS TO BE UNCOMMENTED BEFORE NEW SCRAPER ENTERED"
			}else if *cron_flag == "comment"{
				log.Println("ERROR: CRON IS COMMENTED. PLEASE UNCOMMENT BEFORE CONTINUING")
			}else{



		// at this point we are ready to go and edit the file
		// now we are gonna go through and edit 
		// make a flag for the option to reset cron
		// will have to use cron_condition based on what is going on

	}
