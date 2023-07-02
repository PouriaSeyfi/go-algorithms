package selectionsort

func SelectionSort(arr []int) {
    n := len(arr)

    for i := 0; i < n-1; i++ {
        minIndex := i

        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
		//In terms of performance, using the parallel assignment statement for swapping is generally more efficient because it eliminates the overhead of creating and managing a temporary variable.
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}
