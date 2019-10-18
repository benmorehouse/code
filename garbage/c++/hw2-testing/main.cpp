#include <iostream> 
#include <sstream>
#include <fstream> 
#include <string>
#include <vector>

/* 
 * Intial Check is to make sure that the intial word in the 
 * command is something that we have
 */
 
int main(){
	std::ofstream mystream;
	mystream.open("newFile.txt", std::ofstream::app);
	if (mystream.fail()){
		std::cout<<"This file failed to open"<<std::endl;
	}
	mystream << "Whats up bitches"<<std::endl;
}


