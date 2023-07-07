package quicksort

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Choose the last element as the pivot
	pivot := arr[len(arr)-1]
	left := 0
	right := len(arr) - 2

	// Partitioning step
	for left <= right {

		// Find an element on the left side that is greater than the pivot
		if arr[left] > pivot && arr[right] < pivot {
			arr[left], arr[right] = arr[right], arr[left]
		}

		// Find an element on the right side that is smaller than the pivot
		if arr[left] <= pivot {
			left++
		}

		// Swap the elements if necessary
		if arr[right] >= pivot {
			right--
		}
	}

	// Place the pivot in its correct position
	arr[left], arr[len(arr)-1] = arr[len(arr)-1], arr[left]

	// Recursive calls on the sub-arrays
	QuickSort(arr[:left])   // Sort elements smaller than the pivot
	QuickSort(arr[left+1:]) // Sort elements greater than the pivot
}
