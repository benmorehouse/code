package main

import(
	"fmt"
	"net/http"
	"os"
	"bufio"
	"flag"
	"time"
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

	if err != nil{
		fmt.Println("Error: Not a valid website")
		os.Exit(1)
	}

	fmt.Print("Now enter in the new website scraper : ")
	var newFile string
	fmt.Scan(&newFile)
	err = httpCheck(newFile)

	if err != nil{
		fmt.Println("Error: Not a valid website")
		os.Exit(1)
	}

	//data.txt is what we will use for now, and we will import to new data
	time.Sleep(3 * time.Second)

	fptr := flag.String("filepath","data.txt","") // the pointer needed to do os.Open()
	flag.Parse()
	f, err:= os.Open(*fptr) // file is the returned value of os.open. It is what we read through 
	if err != nil{
		fmt.Println("error: file not able to be opened")
	}
	scan := bufio.NewScanner(f)
	buzzfeedCount:=0
	for scan.Scan(){ // we read in the string line by line 
		var temp string = scan.Text()
		fmt.Println(temp)
		for i:=0;i<len(temp);i++{ // this will loop through each line character by character
			if temp[i]==file[0]{ //need to trim the end off and also lower the B
				var match bool = true
				place := 0
				for j:=i;j<i+len(file)-1;j++{
					if temp[j]!=file[place]{
						match = false // this will break out of loop and just continue .. wasnt wanted
						fmt.Println(temp[j],"doesnt =",file[place])
						break
					}else{
						place++
					}
				}
				if match == true{
					fmt.Println("We found a buzzfeed")
					buzzfeedCount++
				}

			}
		}
	}
	fmt.Println(buzzfeedCount,"buzzfeeds were found")
}
