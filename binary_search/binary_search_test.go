package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
    arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

    // Test case: Target element exists in the array
    target := 23
    expected := 5
    result := BinarySearch(arr, target)
    if result != expected {
        t.Errorf("Expected index %d, but got index %d", expected, result)
    }

    // Test case: Target element does not exist in the array
    target = 50
    expected = -1
    result = BinarySearch(arr, target)
    if result != expected {
        t.Errorf("Expected index %d, but got index %d", expected, result)
    }

    // Test case: Target element is the first element in the array
    target = 2
    expected = 0
    result = BinarySearch(arr, target)
    if result != expected {
        t.Errorf("Expected index %d, but got index %d", expected, result)
    }

    // Test case: Target element is the last element in the array
    target = 91
    expected = 9
    result = BinarySearch(arr, target)
    if result != expected {
        t.Errorf("Expected index %d, but got index %d", expected, result)
    }
}
