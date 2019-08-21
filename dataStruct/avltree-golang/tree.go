package main

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




