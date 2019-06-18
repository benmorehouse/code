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

	// This is just the starting interface 
	fmt.Print("Enter in your website (EX: buzzfeed) : ")
	var file string
	fmt.Scan(&file)
	err := httpCheck(file)
	file = strings.ToLower(file)

	if err != nil{
		fmt.Println("Error: Not a valid website")
		os.Exit(1)
	}

	fmt.Print("Now enter in the new website scraper : ")
	var newFile string
	fmt.Scan(&newFile)
	err = httpCheck(newFile)
	//origNewFile := newFile // used later on in the process
	newFile = strings.ToLower(newFile)

	if err != nil{
		fmt.Println("Error: Not a valid website")
		os.Exit(1)
	}

	//data.txt is what we will use for now, and we will import to new data

	fptr := flag.String("filepath","data.txt","") // the pointer needed to do os.Open()
	flag.Parse()
	f, err:= os.Open(*fptr) // file is the returned value of os.open. It is what we read through 
	if err != nil{
		fmt.Println("error: file not able to be opened")
	}
	scan := bufio.NewScanner(f)
	channelCount:=0
	for scan.Scan(){ // we read in the string line by line 
		var temp string = scan.Text()
		for i:=0;i<len(temp);i++{ // this will loop through each line character by character
			if temp[i]==file[0]{ //need to trim the end off and also lower the B
				if checkMatch(file, newFile, temp, i){
					channelCount++
					// here is where i need to write the change function
				}
			}

		}
	}
	fmt.Println(channelCount,"buzzfeeds were found")
}

func changTempToOrig(temp string, OrigNewFile string){
// gotta do a lot of experimentation on this. It needs to delete it then add in the new one all where it was supposed to be
	
}

func checkMatch(file string, newFile string, temp string, i int) bool{
	place := 0
	for j:=i;j<i+len(file)-1;j++{
		if temp[j] != file[place]{
			return false
		}else{
			place++
		}
	}
	return true
}
