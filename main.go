package main

import (
	"fmt"

	binarysearch "github.com/PouriaSeyfi/go-algorithms/binary_search"
	bfs "github.com/PouriaSeyfi/go-algorithms/breath_first_search"
	mergesort "github.com/PouriaSeyfi/go-algorithms/merge_sort"
	quicksort "github.com/PouriaSeyfi/go-algorithms/quick_sort"
	selectionsort "github.com/PouriaSeyfi/go-algorithms/selection_sort"
)

func main() {

	fmt.Println("----------- Binary Search ----------")

	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	index := binarysearch.BinarySearch(arr, target)
	fmt.Println(index) // Output: 5

	descArr := []int{7, 6, 3, 2, 1}
	r := binarysearch.BinarySearchInDescendingArray(descArr, 1)

	fmt.Printf("%v\n", r) // Output: 4

	//----------------------------------------------
	fmt.Println("---------- Selection Sort ----------")
	arr1 := []int{64, 25, 12, 22, 11}
	selectionsort.SelectionSort(arr1)
	fmt.Printf("%v\n", arr1) // Output: [11 12 22 25 64]

	//----------------------------------------------
	fmt.Println("---------- Quick Sort ----------")
	arr2 := []int{75, 22, 11, 17, 19}
	arr2 = quicksort.QuickSort(arr2)
	fmt.Printf("%v\n", arr2) // Output: [12 17 19 22 75]

	//----------------------------------------------
	fmt.Println("---------- Merge Sort ----------")
	arr3 := []int{9, 5, 1, 3, 8, 4, 2, 7, 6}
	mergesort.MergeSort(arr3)
	fmt.Printf("%v\n", arr3) // Output: [1 2 3 4 5 6 7 8 9]

	fmt.Println("--------- Breadth-First Search ----------")
	// Breadth-First Search
	// Create a social network graph for testing
	//   Alice -> Bob -> Charlie -> David
	//          |
	//          -> John

	alice := &bfs.Person{Name: "Alice"}
	bob := &bfs.Person{Name: "Bob"}
	charlie := &bfs.Person{Name: "Charlie"}
	david := &bfs.Person{Name: "David"}
	john := &bfs.Person{Name: "John"}

	alice.Friends = []*bfs.Person{bob, john}
	bob.Friends = []*bfs.Person{charlie}
	charlie.Friends = []*bfs.Person{david}

	person := bfs.BreadthFirstSearch(alice, "Charlie")
	fmt.Println(person.Name) // Charlie

}
