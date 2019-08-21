package main

import(
	"fmt"
)

func main(){
	integer := 0
	return_val(&integer)
	fmt.Println(integer)
}


func return_val(mystuff *int){
	*mystuff++
}

