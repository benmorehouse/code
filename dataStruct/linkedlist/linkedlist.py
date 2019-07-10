class node(object): # this is how you create a class in python
    def __init__(Node, data): # this is somewhat like a constructor for python 
        # we are constructing self here. We make a variable self.data which will be data and then self.next
        Node.data = data
        Node.Next = None

class Linkedlist():
    
    def __init__(List): # constructing a list called self
        List.head = None
        List.size = 0 
    
    def isEmpty(List):
        if not List.head:
            return True

        else:
            return False


    def insertStart(List, data):
        newNode = node(data)
        List.size = List.size + 1

        if not List.head: # meaning if self.head is null then ....
            List.head = newNode
        
        else:
            newNode.Next = List.head
            List.head = newNode
            List.size += 1
    
    def size(List):
        return List.size

    def findEnd(List):
        Node = List.head

        if not Node:
            return
        
        else:
            while Node.Next != None:
                Node = Node.Next
        
        return Node

    def insertEnd(List, data):
        newNode = node(data)
        endNode = List.findEnd
       
        if not endNode:
            return 
        else:
            endNode.Next = newNode
            List.size += 1

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
            prevNode = None
            curNode = List.head
            if curNode.Next == None: #this list only has one element
                List.head = None
                List.size -= 1
            
            List.head = curNode.Next
                   


workList = Linkedlist()
workList.insertStart("hello world")
workList.insertEnd("This is my program")
workList.traverse()
