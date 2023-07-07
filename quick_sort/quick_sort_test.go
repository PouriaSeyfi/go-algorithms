package quicksort

import (
    "reflect"
    "testing"
)

func TestQuickSort(t *testing.T) {
    arr := []int{64, 25, 12, 22, 11}
    expected := []int{11, 12, 22, 25, 64}

    arr = QuickSort(arr)

    if !reflect.DeepEqual(arr, expected) {
        t.Errorf("Expected %v, but got %v", expected, arr)
    }
}
