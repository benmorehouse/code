# This is an implementation of a basic linked list
# This is more for me to get comfy with Python and less for the development of a linked list

class node(object): # this is how you create a class in python
    def __init__(Node, data): # this is somewhat like a constructor for python 
        # we are constructing self here. We make a variable self.data which will be data and then self.next
        Node.data = data
        Node.Next = None

class Linkedlist():
    
    def __init__(List): # constructing a list called self
        List.head = None
        List.current_size = 0 
    
    def isEmpty(List):
        if not List.head:
            return True

        else:
            return False


    def insertStart(List, data):
        newNode = node(data)
        List.current_size += 1

        if not List.head: # meaning if self.head is null then ....
            List.head = newNode
        
        else:
            newNode.Next = List.head
            List.head = newNode
            List.current_size += 1
    
    def insertEnd(List, data):
        newNode = node(data)
        if List.isEmpty == True:
            return
        else:
            current_node = List.head
            while current_node.Next != None:
                current_node = current_node.Next

            current_node.Next = newNode
            List.current_size += 1

    def traverse(List): # This works
        if List.isEmpty == True:
            return "List is empty"
        else:
            curNode = List.head
            while curNode != None:
                print(curNode.data)
                curNode = curNode.Next

    def findNode(List, data):
        if List.isEmpty == True:
            return "List is empty"
        else:
            curNode = List.head
            while curNode != None:
                if curNode.data == data:
                    return curNode
            
            return None
    def deleteTop(List):
        if List.isEmpty == True:
            return "List is empty"
        else:
            curNode = List.head
            if curNode.Next == None: #this list only has one element
                List.head = None
                List.size -= 1
            
            List.head = curNode.Next
     
    def get_size(List):
        return List.current_size

              


workList = Linkedlist()
workList.insertStart("This is something that i have made")
workList.insertEnd("hello world")
workList.traverse()
workList.deleteTop()
workList.traverse()
