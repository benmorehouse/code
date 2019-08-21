#ifndef DELIST_H
#define DELIST_H

struct node{
	int size;
	*node prev;
	*node next;
};

class linked_list {
	//Return the current size	
	private:
	int return_size();
	
	// Returns the head and tail of whatever we are looking for 
	*Node return_head();
	*Node return_tail(*Node); 

	void traverse_head_to_tail(*Node);
	void traverse_tail_to_head(*Node);
}
