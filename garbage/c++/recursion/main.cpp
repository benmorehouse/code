class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;

    Node() {}

    Node(int _val, Node* _prev, Node* _next, Node* _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
    }
};

class Solution {
public:
    Node* flatten(Node* head) {
        if (!head){return getStart(head->prev);} 
        if (head->child){
            head->child->getEnd(head->child)->next = head->next;
            head->child->prev = head;
            head->next = head->next;
        }
        
        return flatten(head->next);
    }
    
    Node* getEnd(Node* head){
        if (!head->next){
            return head;
        }else{
            return this->getEnd(head->next);
        }
    }
    
    Node *getStart(Node* tail){
        if(tail->prev == NULL){
            return tail;
        }else{
            return getStart(tail->prev);
        }
    }
};
