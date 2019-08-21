package main

import(
	"fmt"
)

func main(){
	binary := tree{
		root: nil,
	}
	fmt.Print()
	binary.insert(2)
	binary.insert(3)
	binary.insert(1)
	fmt.Println("root is:", binary.return_root())
	binary.insert(22)
	binary.insert(12)
	binary.insert(4)
	binary.traverse()
	binary.delete_node(22)
	binary.traverse()
}
