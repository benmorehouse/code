package main

import(
	"fmt"
	"bufio"
	"os"
)
// this will prompt the user for the input
func messageWriter(name string)string{
	fmt.Println("Please enter your message \n To enter name, type '~'\n ~")
	writer := bufio.NewReader(os.Stdin)
	output,_ := writer.ReadString('\n')
	output +="\b" // this gets rid of newline
	// this will scan through and add the name in 
	for i:=0;i<len(output);i++{
		if output[i]=='~'{
			output +="\b "
			output +=name
		}
	}

	fmt.Println("You have entered:\n\n",output,"\n Is this what you want? [y/n]")
	var x string
	fmt.Scan(&x)
	if x=="n" || x=="N"{
		return messageWriter(name)
		// recursively return what they do when they go through the entire process again 
	}else{
		return output
	}
}
// this will return the terminal writer we need 
func terminalWriter(message, number string)string{
//osascript sendMessage.applescript 1235551234 "Hello there!"
	var output string 
	output += "osascript sendMessage.applescipt "
	output += number + " "
	output += string('"')
	output += message
	output += string('"')
	return output

}



