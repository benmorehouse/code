/*#include <iostream>
#ifndef NODE_H
#define NODE_H

using namespace std;

struct Node {
    int value;
    Node *next;
};

Node* funcA(Node* in);
Node* funcB(Node* in); 


int main(){
	Node five = {
		5, NULL,
	};
	
	Node four = {
		4, &five,
	};

	Node three = {
		3, &four,
	};

	Node two = {
		2, &three,
	};

	Node one = {
		1, &two, 
	};
	
	Node* output = funcA(&one);
	std::cout<<output->value<<std::endl;
}

Node* funcA(Node* in) {
	Node *out = in;
	while (out->next != nullptr) {
	out = out->next;
	}
	funcB(in)->next = NULL;
	return out;
}

Node* funcB(Node* in) {
	if (in->next != nullptr) {
	funcB(in->next)->next = in;
	}
	return in;
}
#include <iostream>
#include <vector>

struct Node{
	int data;
	Node* next;
};

int main(){
	Node *head = new Node; // the first four lines are o(1) runtime each ... not important
	Node *curr = head;
	head->data = 0;
	head->next = nullptr;
	int n = 9;
	
	for (int i = 1; i < n; i++) // this will go through exactly o(n) times
	{
	   curr->next = new Node;
	   curr = curr->next;
	   curr->data = i;
	   curr->next = nullptr;
	} // this just constructs a linked list

	for (int i = 1; i <= n; i++) { // this will go through exactly o(n) times
	   curr = head;
	   while (curr != nullptr) {  // however long the length of the linked list is
	      if (curr->data % i == 0) {
		 for (int j = 0; j < n; j++) { // this will worst case scenario run o(n!) times
		     A[j] ++;
		 }
	      }
	      curr = curr->next;
	   }
	}
*/

#include <iostream>

int read_input(double*, int); // double* because this is simply a pointer to the beginning of the array

int main(){
    const int CAPACITY = 10; // good use of const
    double values[CAPACITY]; // have you looked into dynamic memory allocation yet?
    int input = read_input(values, CAPACITY);

    for (int i = 0; i < CAPACITY; i++){
        std::cout << values[i] << " ";
    }
    std::cout << "\n";
    return 0;
}


int read_input(double values[], int CAPACITY){ // notice here i indent to the same line is where int main() is...
    int current_size = 0;// you need to initialize this to 0 otherwise the compiler wont know where to start the value!!
    int input;
    std::cout << "Input integers" << std::endl;
    while(current_size < CAPACITY){ // btw -> this is an o(n) process. Look up runtime analysis when you get around to it
        std::cin >> input;
        values[current_size] = input; 
        current_size++; // this means current_size = current_size + 1; but you dont say that current_size = 0 so what happens first run through? 
    }
    return current_size;  // what really is the point of this variable being returned? you dont have to return a variable... just simply say function is void
    
}

/*  Few notes i have: make sure you are commenting out your code and describing what you are doing. Make sure specifically 
 *  that you explain each logic block ie "The next 10 lines are used to assign and assess the value of the array"
 *
 *  you need to look into correct function declaration and indenting
 *
 *  also make sure when you are naming functions, you are naming them precisely. For instance i wouldnt have named it read_input. 
 *  yes you are reading something but it could mean reading from a file etc. instead i would declare as getUserInput (camel casing) the name
 */
