#include <string>
#include <iostream>

class shape{
	public:
		shape(std::string name, int sides){
			_name = name;
			_sides = sides;
		}
		~shape(){};
		virtual void printme(){
			std::cout<<"this shouldnt print"<<std::endl; 
		}
	protected:
		std::string _name;
		int _sides;
};

class triangle : protected shape{
	public: 
		triangle(int _size):shape("triangle",3){
			this->size = _size;	
		}
		~triangle(){};
		void printme(){
			std::cout<<"triangle"<<std::endl;	
		}
	private:
		int size;
};

class square : protected shape{
	public:
		square(int _size):shape("square",4){
			this->size = _size;
			this->_sides = 4;
			this->_name = "square";		
		}
		~square(){};
		void printme(){
			std::cout<<"square"<<std::endl;	
		}
		
	private:
		int size;
};

int main(){
	triangle *tr1 = new triangle(4);
	square *sqr1 = new square(4);
	tr1->printme();
	sqr1->printme();
	delete tr1;
	delete sqr1;
}

/*
#include <iostream>
using namespace std;

class base{
	public: 
		virtual void print () 
		{ cout<< "print base class" <<endl; } 
   
		void show () 
		{ cout<< "show base class" <<endl; } 
}; 
   
class derived:public base{
	public: 
		void print () //print () is already virtual function in derived class, we could also declared as virtual void print () explicitly 
		{ cout<< "print derived class" <<endl; } 
   
		void show () 
		{ cout<< "show derived class" <<endl; } 
}; 
  
//main function 
int main()  
{ 
	base *bptr; 
	derived d; 
	bptr = &d; 
	//virtual function, binded at runtime (Runtime polymorphism) 
	bptr->print();  
	// Non-virtual function, binded at compile time 
	bptr->show();  
	return 0; 
}  
*/

