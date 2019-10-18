#include <iostream>
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
	std::cout<<output->next->value<<std::endl;
	std::cout<<output->next->next->value<<std::endl;
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
#endif
