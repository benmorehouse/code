package main

import(
	"fmt"
	"net/http"
	"os"
	"bufio"
	"flag"
	"strings"
)

func httpCheck(file string)(error){
	_ ,err := http.Get("http://"+file+".com")
	return err
}

func main(){
	fmt.Println("                  ***Welcome to Autofill***\n")
	fmt.Println("/***********************************************************************/\n")
	fmt.Println("This is a tool that will be used to parse through files and change the \nnames of the file that we are working on to a new file which will be a new scraper/etc.\n")
	fmt.Println("/***********************************************************************/\n")

	fmt.Print("Enter in your website (EX: buzzfeed) : ")
	var website string
	fmt.Scan(&website)
	err := httpCheck(website)

	if err != nil{
		fmt.Println("Error: Not a valid website")
		os.Exit(1)
	}

	website = strings.ToLower(website)
	fmt.Print("Now enter in the new website scraper : ")
	var newWebsite string
	fmt.Scan(&newWebsite)
	err = httpCheck(newWebsite)

	if err != nil{
		fmt.Println("Error: Not a valid website")
		os.Exit(1)
	}

	//origNewWebsite := newWebsite // used later on in the process
	newWebsite = strings.ToLower(newWebsite)

	//data.txt is what we will use for now, and we will import to new data

	fptr := flag.String("filepath","data.txt","") // the pointer needed to do os.Open()
	flag.Parse()
	f, err:= os.Open(*fptr) // file is the returned value of os.open. It is what we read through 
	if err != nil{
		fmt.Println("error: file not able to be opened")
	}

	scan := bufio.NewScanner(f)
	//channelCount:=0
	for scan.Scan(){ // we read in the string line by line 
		// use the field function to scan through each element. Then change the ones that match
		// will convert the string to all lower
		// only if the first letter is what we want
		temp := strings.Fields(scan.Text())
		fmt.Println(temp)
	}

}

