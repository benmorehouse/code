package main

import(
	"fmt"
	"os"
	"strings"
	"log"
	"errors"
)

func main(){
	mystring := []byte("# backlog \n Hello world this is a string that I have \n i hope that you like it because i really do\n new line \n hit mq when you are finished")
	err, content := backlog_content_manip(mystring)
	if err == nil{
		fmt.Println(string(content))
	}else{
		fmt.Println(err)
	}
}

func changeToDirectory(path string){
	err := os.Chdir("/Users/benmorehouse")
	pathField := strings.Split(path,string('/'))
	for i:=3;i<len(pathField);i++{
		err := os.Chdir(pathField[i])
		if err != nil{
			log.Fatal("This is not a directory")
		}
	}
	fmt.Println(err)
}


func backlog_content_manip(content []byte)(error,[]byte){
	input := string(content)
	if input == ""{
		return errors.New("Content is completely empty"),content
	}
	temp_input := strings.Split(input,string('\n'))
	for i:=1;i<len(temp_input)-1;i++{
		temp := " - " + temp_input[i]
		temp_input[i] = temp
	}
	return nil,[]byte(strings.Join(temp_input, "\n"))
}

