package main

import(
	"fmt"
)

type node struct{
	data int
	left *node
	right *node
	parent *node
}

type tree struct{
	root *node
}

func make_node(input int)(*node){
	new_node := node{
		data: input,
		left: nil,
		right: nil,
		parent: nil,
	}
	return &new_node
}

func (bst *tree) insert(data int){
	if bst.root == nil{
		bst.root = make_node(data)
		return
	}else{
		bst.insert_node(bst.root,data)
	}
}


func (bst *tree) insert_node(entered_node *node, data int) {
	// pass in a piece of data to the bst. Return error if there is something that happens with it
	if data < entered_node.data{
		if entered_node.left == nil{
			entered_node.left = make_node(data)
			entered_node.left.parent = entered_node
		}else{
			bst.insert_node(entered_node.left, data)
		}
	}else if data >= entered_node.data{
		if entered_node.right == nil{
			entered_node.right = make_node(data)
			entered_node.right.parent = entered_node
		}else{
			bst.insert_node(entered_node.right, data)
		}
	}
}

func (bst tree) return_root()(int){
	return bst.root.data
}

func (bst tree) traverse(){
	if bst.root == nil{
		return
	}

	bst.traverse_in_order(bst.root)
}

func (bst tree) traverse_in_order(entered_node *node){
	if entered_node.left != nil{
		bst.traverse_in_order(entered_node.left)
	}

	fmt.Println(entered_node.data)

	if entered_node.right != nil{
		bst.traverse_in_order(entered_node.right)
	}
}

func (bst tree) delete_node(data int){
	if bst.root == nil{
		return
	}
	bst.delete_(bst.root,data)
}

func (bst *tree) delete_ (entered_node *node, data int){
	if data < entered_node.data {
		if entered_node.left == nil && entered_node.right == nil{
			return
		}else{
			bst.delete_(entered_node.left, data)
		}
	}else if data > entered_node.data{
		if entered_node.left == nil && entered_node.right == nil{
			return
		}else{
			bst.delete_(entered_node.left, data)
		}
	}else{
		if entered_node.left == nil && entered_node.right == nil{ // this is a leaf node at the end of a tree
			entered_node = nil
		}else if entered_node.left == nil && entered_node.right != nil{ // has right subtree
			entered_node.parent.right = entered_node.right
			entered_node.right.parent = entered_node.parent
			entered_node = nil
		}else if entered_node.left != nil && entered_node.right == nil{
			entered_node.parent.left = entered_node.left
			entered_node.left.parent = entered_node.parent
			entered_node = nil
		}else{ // has both left and right
			next_node := bst.find_next_node(entered_node.right) // finds predessor from right subtree
			temp := entered_node.data
			entered_node.data = next_node.data
			next_node.data = temp
			bst.delete_(entered_node, entered_node.data)
		}
	}
}


func (bst *tree) find_next_node(entered_node *node)(*node){
	// given certain node, finds left most node
	if entered_node.left == nil{
		return entered_node
	}else{
		return bst.find_next_node(entered_node.left)
	}
}

