package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"net/http"
)

func httpCheck(file string){
	resp,_ := http.Get("http://"+file+".com")
	fmt.Println("http://"+file+".com")
}

func main(){
	fmt.Println("***Welcome to Autofill***")
	fmt.Println("This is a tool that will be used to parse through files and change the names of the file that we are working on to a new file which will be a new scraper/etc.")
	var startFile string
	fmt.Scan(&x)
	httpCheck("buzzfeed")
}
	// this will read in a single string of what they enter
	// Now we check it by pushing an httprequest to see if it is a valid address
