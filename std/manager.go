package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"flag"
// see what is wrong with this 	"golang-list"
)

func manager(todo list, datafile string){
	// read in line by line and push that onto the list that we are working on
	fptr := flag.String("File Scanner",datafile,"file scanner reads what is in the file")
	flag.Parse()
	
	f,err := os.Open(*fptr)
	if err == nil{
		return
	}
	s := bufio.NewScanner(f)
	var temp string
	for s.Scan(){
		temp = s.Text()
		todo.push_task(createnode(temp))
	}
	// now all the stuff in the file is loaded into the list
	todo.tasks_remaining()
	todo.traverse()
	fmt.Println("[add/remove/quit]")
	reader := bufio.NewReader(os.Stdin)
	text,_ := reader.ReadString('\n')
	text = strings.TrimSuffix(text,"\n")
	for text != "quit"{
		if text == "add"{
			fmt.Print("Add Task:")
			text,_ = reader.ReadString('\n')
			text = strings.TrimSuffix(text,"\n")
			todo.push_task(createnode(temp))
		}else if text == "remove"{
			todo.remove()
		}else{
			fmt.Println("NOT A COMMAND")
		}
		text,_ = reader.ReadString('\n')
		text = strings.TrimSuffix(text,"\n")
	}
}

func main(){
	fmt.Println("Tasks?[y/n]")
	reader := bufio.NewReader(os.Stdin)
	text,_ := reader.ReadString('\n')
	text = strings.TrimSuffix(text,"\n")
	if text == "n"{
		return
	}else if text != "y"{
		return
	}

	fmt.Println("\n------------------------------------------------------------------")
	var main list
	var grocery list
	fmt.Println("Which list?")
	fmt.Println("1. Main")
	fmt.Println("2. Grocery")
	var num_input int
	fmt.Scan(&num_input)
	for num_input!=2 && num_input!=1{
		fmt.Println("List Not Available")
		fmt.Scan(&num_input)
	}
	if num_input==1 {
		// go through the process for the main list
		manager(main,"data.txt")
	}else{
		// go through the process for the grocery list
		manager(grocery, "grocery.txt")
	}
}

