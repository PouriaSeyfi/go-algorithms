package main

import (
	"fmt"

	binarysearch "github.com/PouriaSeyfi/go-algorithms/binary_search"
	quicksort "github.com/PouriaSeyfi/go-algorithms/quick_sort"
	selectionsort "github.com/PouriaSeyfi/go-algorithms/selection_sort"
	bfs "github.com/PouriaSeyfi/go-algorithms/breath_first_search"
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
