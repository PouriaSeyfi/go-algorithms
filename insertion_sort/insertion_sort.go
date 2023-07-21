package insertionsort

func InsertionSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j = j - 1
		}
	}
	return arr
}
