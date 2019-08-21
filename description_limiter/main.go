package main

import(
	"fmt"
	"strings"
)

func main(){
	description := "Needles officials say they plan to work with state representatives on legislation that would exempt the city — and possibly other border towns — from rules on purchasing ammunition out of California and reciprocity with other states over concealed weapons permits."
	trim_description(&description)
	fmt.Println(description)
	fmt.Println(len(description))
}

func trim_description(description *string){
	length := len(*description)
	if length < 255{
		return
	}

	temp := []byte(*description)

	i := 255
	for temp[i] != ' '{
		i--
	}

	temp_letter := temp[i]
	temp[i] = ' '
	temp[i+1] = '|'
	temp[i+2] = ' '

	*description = string(temp)
	temp_field := strings.Fields(*description)

	i = 0
	for temp_field[i] != "|"{
		i++
	}
	temp_field[i] = string(temp_letter)

	*description = strings.Join(temp_field[:i], " ")
}
