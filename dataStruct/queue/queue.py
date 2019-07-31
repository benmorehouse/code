# This is first in first out method of insertion and deletions
class node(object):
        
    def __init__(Node, data): # this needs to create the list of the nodes that we are trying to make
        Node.Data = data
        Node.Next = None

class Queue():
    
    def __init__(Queue):
        Queue.Head = None

    def Is_empty(Queue):
        if Queue.Head == None:
            return True
        else:
            return False

    def Insert(Queue, entered_data):
        newNode = node(entered_data)
        if Queue.Head == None: # this thing is not empty so we know to just add on 
            Queue.Head = newNode
        else:    
            newNode.Next = Queue.Head
            Queue.Head = newNode

    def Delete(Queue):
        if Queue.Is_empty == True:
            return "This queue is empty"
        else:
            prev_node = None
            cur_node = Queue.Head
            while cur_node.Next != None:
                prev_node = cur_node
                cur_node = cur_node.Next
           # Now we are at the end node
            prev_node.Next = None


            
    def Traverse(Queue):
        if Queue.Is_empty == True:
            print("Nothing to show in this queue")
        else:
            current_node = Queue.Head
            while current_node != None:
                print(current_node.Data)
                current_node = current_node.Next



myQueue = Queue()
myQueue.Insert("First entry")
myQueue.Insert("second entry")
myQueue.Insert("third entry")
myQueue.Traverse()
print()
myQueue.Delete()
myQueue.Traverse()



