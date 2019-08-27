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

func rc_content_manip(input, new_list  string)(string){
// Need to come here, create a system that takes in content, adds marks wherever there are \n, then does fields and joins
	var marker []int // used to keep place of where there ~
	for i , val := range input{
		if string(val) == "\n"{
			val = '~'
			marker = append(marker,i + 2)
		}
	}
	temp_content := strings.Fields(input)
	temp_content[1] = new_list + "\n\n"
	input = strings.Join(temp_content, " ")
	for _ , val := range marker{
		if input[val]!='~'{
			continue
		}else if val > len(input){
			break
		}else{
			input[val] = '\n'
		}
	}
	return input
}

