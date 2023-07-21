package insertionsort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty array", []int{}, []int{}},
		{"one element", []int{5}, []int{5}},
		{"two elements", []int{2, 1}, []int{1, 2}},
		{"already sorted", []int{1, 2, 3}, []int{1, 2, 3}},
		{"reverse order", []int{3, 2, 1}, []int{1, 2, 3}},
	}

	// Iterate through test cases
	for _, tt := range tests {
		// Call insertionSort function
		t.Run(tt.name, func(t *testing.T) {
			actual := InsertionSort(tt.input)

			// Check if actual output matches expected output
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("InsertionSort(%v) = %v, expected %v", tt.input, actual, tt.expected)
			}
		})
	}
}