# This is first in last out method of insertion and deletions

class stack:

    def __init__(self):
        self.Stack = []

    def isEmpty(self):
        return self.Stack == [] # this is how you can tell if there is something that is empty

    def push(self, data):
        self.Stack.append(data)

    def pop(self):
        data = self.Stack[-1]
        del self.Stack[-1]
        return data

    def peek(self):
        return self.Stack[-1]

    def size(self):
        return len(self.Stack)
    
    def show(self):
        i = 0

        while i!= len(self.Stack):
            print(self.Stack[i])
            i = i+1

myshit = stack()
myshit.push("First entry")
myshit.push("second entry")
myshit.push("third entry")
myshit.show()
print("popped",myshit.pop())
print("popped",myshit.pop())
myshit.show()
