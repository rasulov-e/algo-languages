package main

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		logger.Println("QuickSort: quicksort() called")
		arr = quickSort(arr, low, p-1)
		logger.Println("QuickSort: quicksort() called")
		arr = quickSort(arr, p+1, high)
	}
	logger.Println("QuickSort: sorting is done")
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	logger.Println("QuickSort: partitioning began")
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	logger.Println("QuickSort: partitioning is over")
	return arr, i
}
