#include <iostream>
#include "delist.h"

DEList::DEList(){
	head = NULL;
	tail = NULL;
	_size = 0;
  // Destructor to clean-up memory of current list
}

DEList::~DEList(){
	for(unsigned int i=0;i<_size;i++){
        if(head!=tail){
		    head = head->next;
            delete head->prev;
        }
        else{
            delete head;
            break; 
        }
  // returns a Boolean 'true' if the list is empty
}

bool DEList::empty(){
	if((head==NULL)&&(tail==NULL)){
		return true;
	}
	return false;
  // returns number of items in the list
}

int DEList::size(){
	if(_size>0){
		return _size; 
	}
	return -1;
  // returns front item or -1 if empty list
}

int DEList::front(){
	if(empty()){
		return -1;
	}
	else{
		return head->val;
	}
  // returns back item or -1 if empty list
}

int DEList::back(){
	if(empty()){
		return -1;
	}
	else{
		return tail->val;
	}
  // Adds a new integer to the front of the list
}

void DEList::push_front(int new_val){
	Item* newNode = createItem(new_val);

	if(head == NULL){
		head = newNode;
        tail = newNode; 
	}
	else{
		head->prev = newNode;
		newNode->next = head;
		head = newNode;
	}
	_size++;
  // Adds a new integer to the back of the list
}

void DEList::push_back(int new_val){
	Item* newNode = createItem(new_val);
	if(tail == NULL){
		tail = newNode;
        head = newNode; 
	}
	else{
		tail->next = newNode;
		newNode->prev = tail; 
		tail = newNode;
	}
	_size++;
	// Removes the front item of the list (if it exists)
}

void DEList::pop_front(){
	if(empty()){
		return;
	}
	else{
		Item* temp = head;
		temp = head;
		head = temp->next;
        if(head!=NULL){
            head->prev = NULL; 
        }
        else{
            tail = NULL;
        }
		_size--;
		delete temp;
	}
}

void DEList::pop_back(){
	if(empty()){
		return;
	}
	else{
		Item* temp = tail;
		temp = tail; 
		tail = tail->prev;
        if(tail!=NULL){
            tail->next = NULL;
        }
        else{
            head = NULL;
        }
		_size--;
		delete temp;
	}
}

Item* DEList::createItem(int x){
	Item* item = new Item;
	item->next = NULL;
	item->prev = NULL;
	item->val=x;
	return item;
}
