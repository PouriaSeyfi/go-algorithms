package mergesort

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {

	result := make([]int, 0)

	for len(left) > 0 && len(right) > 0  {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result
}
