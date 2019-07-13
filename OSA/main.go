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

	fptr := flag.String("","cron_conditions.txt","") // the file pointer needed to do os.Open(). Leaving the name blank because will never have to edi
	new_scraper := flag.String("ns", "", "user will input this to tell what the name of the new scraper is")
	flag.Parse()
	if *new_scraper == ""{
		log.Fatal("ERROR: NO SCRAPER ENTERED")
	}

	file, err := os.Open(*fptr) // file is the returned value of os.open. It is what we read through 
	if err != nil{
		fmt.Println("error: cron conditions not able to be opened")
	}
	// this is to check that cron_condition is starting out default
	scan := bufio.NewScanner(file)
	var scheduler cron
	scan.Scan()
	scheduler.filename = scan.Text()
	scan.Scan()
	scheduler.condition , _ = strconv.Atoi(scan.Text())
	scan.Scan()
	scheduler.condition_name = scan.Text()

	if scheduler.condition != 0 || scheduler.condition_name != "INITIAL_CONDITION"{
		reset_cron()
		log.Fatal(scheduler.filename," is not in an initial state. Restart program")
	}
	// at this point we are ready to go and edit the file

	*new_scraper = strings.TrimSpace(*new_scraper)
	scheduler.condition_name = *new_scraper
	fmt.Println("new_scraper:", *new_scraper)

	file ,err = os.Open("cron.txt")
	scan = bufio.NewScanner(file)


}
