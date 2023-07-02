package main

import (
	"fmt"

	binarysearch "github.com/PouriaSeyfi/go-algorithms/binary_search"
)


func main(){

	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23
	
	index := binarysearch.BinarySearch(arr, target)
	
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}

}