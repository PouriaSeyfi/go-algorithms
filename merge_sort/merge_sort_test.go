package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	// Test case 1
	arr1 := []int{9, 5, 1, 3, 8, 4, 2, 7, 6}
	expected1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sortedArr1 := MergeSort(arr1)
	if !reflect.DeepEqual(sortedArr1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, sortedArr1)
	}

	// Test case 2
	arr2 := []int{5, 4, 3, 2, 1}
	expected2 := []int{1, 2, 3, 4, 5}
	sortedArr2 := MergeSort(arr2)
	if !reflect.DeepEqual(sortedArr2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, sortedArr2)
	}

	// Test case 3
	arr3 := []int{}
	expected3 := []int{}
	sortedArr3 := MergeSort(arr3)
	if !reflect.DeepEqual(sortedArr3, expected3) {
		t.Errorf("Expected %v, but got %v", expected3, sortedArr3)
	}
}