package main

import(
	"fmt"
)

type node struct{
	data int
	leftChild *node
	rightChild *node
}

type tree struct{
	root *node
}

func (bst tree) insert(data int){
	if bst.root == nil{
		bst.root 
