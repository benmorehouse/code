package main

import(
	"fmt"
	"strings"
)

func main(){
	input := "#Work \n hello world this is something that I am making \n if you are interested in anything tlike this you should go and check it out"
	input = rc_content_manip(input, "mywork")
	fmt.Println(input)
}

func rc_content_manip(input, new_list string)(string){
// Need to come here, create a system that takes in content, adds marks wherever there are \n, then does fields and joins
	var marker []int // used to keep place of where there ~
	marker_temp := []byte(input)
	for i , val := range input{
		if string(val) == "\n"{
			marker_temp[i] = '~' // right here is not registering it to input, only to value
			marker = append(marker,i + 3)
		}
	}
	input = string(marker_temp)
	temp_content := strings.Fields(input)
	temp_content[0] = new_list + "\n\n"
	input_temp := []byte(strings.Join(temp_content, " "))
	for _ , val := range marker{ // marker is not working well 
		if input_temp[val] != '~'{
			fmt.Println("input_temp is:",string(input_temp[val]))
			fmt.Println("before input_temp is:",string(input_temp[val-1]))
			fmt.Println("test")
			continue
		}else if val > len(input_temp){
			break // prevents seg fault
		}else{
			input_temp[val] = '\n'   // have to change this to a byte slice
		}
	}
	input = string(input_temp)
	return input
}



