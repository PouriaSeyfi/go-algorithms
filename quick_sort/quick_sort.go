package quicksort

func QuickSort(arr []int) []int {
	// QS([2,1,7,9,3]) :
	// p = 3
	// l = QS([2,1]) = [1,2]
	// r = QS([7,9]) = [7,9]
	// return [1,2,3,7,9]
	//-------------------
	// QS([2,1]) :
	// p = 1
	// l = QS([2]) = [2]
	// r = []
	// return [1,2]
	//--------------------
	// QS([7,9]) :
	// p = 9
	// l = QS([7]) = [7]
	// l = []
	// return [7,9]

	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	var left, right []int

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}
