package bucketsort


func BucketSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	// Determine the range of input values
	minVal := arr[0]
	maxVal := arr[0]
	for _, val := range arr {
		if val < minVal {
			minVal = val
		} else if val > maxVal {
			maxVal = val
		}
	}

	n := len(arr)
	countsMap := make([]int, maxVal+1)

	for i := 0; i < n; i++ {
		countsMap[arr[i]] = countsMap[arr[i]] + 1
	}

	outputArr := make([]int, n)
	j := 0
	for index, value := range countsMap {
		for i := 0; i < value; i++ {
			outputArr[j] = index
			j++
		}
	}
	return outputArr
}
