package main

import(
	"fmt"
	"bufio"
	"os"
)

type data struct{
	content string
	rank int // may end up scratching this but also may not
}

type node struct{
	next *node
	prev *node
	node_data *data
}

type list struct{
	head *node
	tail *node
	_size int
}

func createnode (temp string)(newnode *node){
	var stuff *node
	stuff.next=nil
	stuff.prev = nil
	stuff.node_data.content = temp
	stuff.node_data.rank = 0
	return newnode
}
//---------------------------------------------
//--------------------------------------------_
//this is all the stuff for node and data
// there should be something that will create the node struct with the data that we have
//this is all the stuff for list struct 
func (todo list) tasks_remaining(){
	fmt.Println("There are",todo._size,"things left on the list")
}

func (todo list) push_task(input *node){ // in the main function we will have another function which will 
	// input is gonna be what the user inputs for the string that they are trying to add to the list
	if todo.head==nil && todo.tail==nil{
		// this is just an empty list then
		todo.head = input
		todo.tail = input
		todo.assign_rank()

	}else{
		// at this point we are trying to see where we want to put this task into the task manager
		fmt.Println("How important is this task?")
		// this is now gonna read in what they are
		reader := bufio.NewReader(os.Stdin)
		text,_ := reader.ReadString('\n')
		// similar to traverse but we prompt the user in each loop 
		// should you return anything
		var ptr *node = todo.head
		for ptr!=nil{
			fmt.Println(ptr.node_data.content)
			fmt.Println("Is it more important than this?[y/n]")
			text,_=reader.ReadString('\n') // read input
			if text=="y"{
				if ptr == todo.head{
					input.next = ptr
					ptr.prev = input
					todo.head = ptr
					todo.assign_rank()
				}else if ptr == todo.tail{
					ptr.next = input
					input.prev = ptr
					todo.tail = input
					todo.assign_rank()
				}else{
					input.next = ptr
					ptr.prev.next = ptr
					ptr.prev = input
					todo.assign_rank()
				}
				return  // break out of this loop once we have inserted what we want to have happen
				//insert the node ahead of the displayed node that we looked at 
			}else{
				fmt.Print("\n\n\n")
				ptr = ptr.next
			}
		}
	}
	todo._size ++
}

func (todo list) traverse(){// this will be used to display the entire list at start
	// should you return anything
	todo.assign_rank()
	if todo._size == 0{
		return
	}else{
		var ptr *node = todo.head
		for ptr!=nil{
			fmt.Println(ptr.node_data.rank,"\b.",ptr.node_data.content)
		}
	}
}

func (todo list) remove(){
	if todo._size == 0{
		return
	}else if todo._size == 1{
		todo.head = nil
		todo.tail = nil
	}else{
		todo.traverse() // this will display all of the nodes
		var num int = 0
		fmt.Println("Which item is complete?")
		fmt.Scan(&num) // they will enter in which node they are trying to remove
		if num > todo._size || num < 1{
			fmt.Println("That's not a task")
			return
		}else{
			// at this point you need to go in and delete which ever one is part of the list
			// then you need to renumber the lis 
			num --
			var ptr *node = todo.head // this is the node that we are gonna scan through with
			for i:=0;i<num;i++{
				ptr = ptr.next
			}
			// at this point ptr is pointing somewhere and we rearrange the pointers
			if ptr == todo.head{
				todo.head = todo.head.next
			}else if ptr == todo.tail{
				todo.tail = todo.tail.prev
			}else{
				ptr.next.prev = ptr.prev
				ptr.prev.next = ptr.next
				ptr.prev = nil
				ptr.next = nil
				todo.assign_rank()
			}
			todo._size --
		}
	}
}

func (todo list) assign_rank(){
	var ptr *node = todo.head
	for i:=0;i<todo._size;i++{
		ptr.node_data.rank = i+1
		ptr=ptr.next
	}
}
