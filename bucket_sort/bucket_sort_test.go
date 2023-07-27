package bucketsort

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"case one", []int{2, 1, 0}, []int{0, 1, 2}},
		{"case two", []int{0, 3, 2, 1, 1, 3, 1, 0}, []int{0, 0, 1, 1, 1, 2, 3, 3}},
		{"case three", []int{1, 0, 0, 1}, []int{0, 0, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := BucketSort(tt.input)

			// Check if actual output matches expected output
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("BucketSort(%v) = %v, expected %v", tt.input, actual, tt.expected)
			}
		})
	}
}
