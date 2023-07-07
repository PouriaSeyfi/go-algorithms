package binarysearch


func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func BinarySearchInDescendingArray(arr []int, target int) int {

	left := 0
	right := len(arr) - 1

	for left <= right {
		m := (left + right)/2
		if arr[m] == target {
			return m
		} else if arr[m] < target {
			right = m - 1
		} else {
			left = m + 1
		}
	}

	return -1
}
