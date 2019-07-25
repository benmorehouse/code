package main

import(
	"fmt"
)

type string_heap struct{
	heap_size int
	heap_array []string
}

func main(){
	var array = []string{"work","play","hard","easy"}
	min_heap := string_heap {
		heap_size: len(array),
		heap_array: array,
	}
	min_heap.display_heap()
}

func (heap string_heap) display_heap(){
	for i , val := range heap.heap_array{
	}
}
