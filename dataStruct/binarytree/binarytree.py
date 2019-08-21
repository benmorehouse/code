# Binary Search tree that will give you a runtime for search of o(log n) and worst case scenario a o(n) runtime for search and insertion

class node:
    def __init__(self, data):
        self.data = data 
        self.left = None # lesser nodes go left
        self.right = None # greater nodes go right 

class binary_search_tree:
    def __init__(self):
        self.root = None

    def insert(self, data):
        newNode = node(data)
        if not self.root: # this means that the list is empty
            self.root = newNode
        else:
            self.insert_node(data, self.root)

    def insert_node(self, data, entered_node): # the entered_node is what we are comparing the data to 
        newNode = node(data)
        if data < entered_node.data:
            if entered_node.left:
                self.insert_node(data,entered_node.left)
            else:
                entered_node.left = newNode
        else:
            if entered_node.right:
                self.insert_node(data,entered_node.right)
            else:
                entered_node.right = newNode

    def get_min_value(self):
        if self.root == None:
            return "The tree is empty: no min"
        else:
            return self.get_min(self.root)

    def get_min(self,entered_node):
        if entered_node.left == None:
            return entered_node.data
        else:
            return self.get_min(entered_node.left)

    def get_max_value(self): # this will be what the user passes
        if self.root == None:
            return "The tree is empty: no min"
        else:
            return self.get_max(self.root)
    
    def get_max(self, entered_node):
        if entered_node.right == None:
            return entered_node.data
        else:
            return self.get_max(entered_node.right)
   
    def remove_node(self, data, entered_node): # this will return what ever is gonna be the next root node 
        # This is broken somehow and i think that there is something with how we have things alligned on the left and the right
        if not entered_node:
            return entered_node # if what we get for some reason is null than we can just let this go 

        if data < entered_node.data:
            print("data is less than the entered_node which is",entered_node.data)
            return self.remove_node(data, entered_node.left)

        elif data > entered_node.data:
            return self.remove_node(data, entered_node.right)

        else:
            if not entered_node.left and not entered_node.right:
                print("Removing a leaf node...");
                del entered_node
                return None
                    
            if not entered_node.left:  
                print("Removing a node with single right child...");
                tempNode = entered_node.right
                del entered_node
                return tempNode
            elif not entered_node.right:   # node instead of self
                print("Removing a node with single left child...");
                tempNode = entered_node.left
                del node
                return tempNode
            
        
        print("Removing node with two children....");
 
    def delete(self, data):
        if self.root:
            self.root = self.remove_node(data, self.root)
       
    def get_predecessor(self,entered_node): # this comes through and finds the greatest value that is after the node that we delete
        if entered_node.right: 
            return self.get_predecessor(entered_node.right)
        return entered_node
    
    def traverse(self):
        if self.root:
            self.traverse_in_order(self.root)
        else:
            print("was an empty tree")

    def traverse_in_order(self, current_node):
        if current_node.left:
            self.traverse_in_order(current_node.left)
        
        print (current_node.data)

        if current_node.right:
            self.traverse_in_order(current_node.right)
    
bst = binary_search_tree()
bst.insert(10);
bst.insert(13);
bst.insert(5);
bst.insert(14);
bst.traverse()
bst.delete(5)
bst.traverse()
# the issue right now is the self.root is fucked up somehow


