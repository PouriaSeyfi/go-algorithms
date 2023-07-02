package main

import (
	"fmt"

	binarysearch "github.com/PouriaSeyfi/go-algorithms/binary_search"
	selectionsort "github.com/PouriaSeyfi/go-algorithms/selection_sort"
)

func main() {

	// Binary Search
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	index := binarysearch.BinarySearch(arr, target)
	fmt.Println(index) // Output: 5


	// Selection Sort
	unsortedArr := []int{64, 25, 12, 22, 11}
	selectionsort.SelectionSort(unsortedArr)
	fmt.Println(unsortedArr) // Output: [11 12 22 25 64]

}
