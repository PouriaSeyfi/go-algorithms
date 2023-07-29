package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"case one", []int{2, 1, 7, 9, 3}, []int{1, 2, 3, 7, 9}},
		{"case two", []int{0, 3, 2, 1, 1, 3, 1, 0}, []int{0, 0, 1, 1, 1, 2, 3, 3}},
		{"case three", []int{9, 8, 7, 6}, []int{6, 7, 8, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := QuickSort(tt.input)

			// Check if actual output matches expected output
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("BucketSort(%v) = %v, expected %v", tt.input, actual, tt.expected)
			}
		})
	}

	arr := []int{2, 1, 7, 9, 3}
	expected := []int{1, 2, 3, 7, 9}

	arr = QuickSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, but got %v", expected, arr)
	}
}
