package main

type node struct{
	data int
	left *node
	right *node
}

func make_node(input int)(*node){
	new_node := node{
		data: input,
		left: nil,
		right: nil,
	}
	return &new_node
}

