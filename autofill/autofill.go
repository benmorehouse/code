package main

import(
	"fmt"
	"net/http"
	"os"
	"bufio"
	"flag"
	"strings"
	"log"
)

// This is for you to check and make sure the website still functions
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
	var input string
	fmt.Scan(&input)
	err := httpCheck(input)

	if err != nil{
		fmt.Println("Error: Not a valid website")
		os.Exit(1)
	}

	currentWebsite := website{
		original: input,
		upper: strings.ToUpper(input),
		lower: strings.ToLower(input),
	}

	fmt.Print("Now enter in the new website scraper : ")
	fmt.Scan(&input)
	err = httpCheck(input)

	f, err := os.Create(input + ".txt") // this is the file writer

	if err !=nil{
		fmt.Println("err:",err)
		os.Exit(1)
	}

	newWebsite := website{
		original: input,
	}
	fmt.Println(newWebsite)
	//data.txt is what we will use for now, and we will import to new data
	// also gotta rename the file here from a go file to txt file before we open 

	fptr := flag.String("filepath",currentWebsite.original+".txt","") // the pointer needed to do os.Open()
	flag.Parse()
	file, err := os.Open(*fptr) // file is the returned value of os.open. It is what we read through 
	if err != nil{
		fmt.Println("error: file not able to be opened")
	}

	scan := bufio.NewScanner(file)
	for scan.Scan(){ // we read in the string line by line  // could scan through and instead just use the ones that we found 
		temp := scan.Text() // this gives us each line
		fmt.Println(currentWebsite.all_occurances)
		/*
		newstrings := currentWebsite.all_occurances
		fmt.Println(len(newstrings))
		for i:=0;i<len([]newstrings);i++{
			temp = strings.Replace(temp,newstrings[i],newWebsite.original,-1)
		}
		*/
		temp+=string("\n")
		// now we add old into new
		if _, err := f.Write([]byte(temp)); err != nil {
			log.Fatal(err)
		}
	}
		// essentially renames the file rename(old, new)
		//use os.Getwd to get the rooted path
}

