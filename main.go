package main

import (
	"fmt"

	binarysearch "github.com/PouriaSeyfi/go-algorithms/binary_search"
	quicksort "github.com/PouriaSeyfi/go-algorithms/quick_sort"
	selectionsort "github.com/PouriaSeyfi/go-algorithms/selection_sort"
)

func main() {

	// Binary Search
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	index := binarysearch.BinarySearch(arr, target)
	fmt.Println(index) // Output: 5


	// Selection Sort
	unsortedArr1 := []int{64, 25, 12, 22, 11}
	selectionsort.SelectionSort(unsortedArr1)
	fmt.Println(unsortedArr1) // Output: [11 12 22 25 64]


	// Quick Sort
	unsortedArr2 := []int{75, 22, 11, 17, 19}
	quicksort.QuickSort(unsortedArr2)
	fmt.Println(unsortedArr2) // Output: [12 17 19 22 75]
}
