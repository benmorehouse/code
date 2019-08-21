# Notes on the binary tree

	The best thing about the binary tree is that unlike the sorted array or the linked list, you are able to get a 
	O (log n) runtime for everything including searching, inserting and deleting

	essentially you build a tree much like what you have seen in 109 and you make it where there is a specific path to 
	each element that you could get your hands on and if there is more than one way to get access to it then you have not
	made a tree

	A leaf node is essentially a node that has no children nodes on it

	binary search trees are trees but they only are allowed to have two leaf nodes

	When you organize the tree, the leaf to the left of the root node is always less than the root and everyting to the 
	right is greater. So essentially you do a binary sort and then you create the tree based off the final data

	The reason that this is so fast is that everytime that you are doing something for the tree, the placement and
	figuring out where to put it is already partially figured out	


