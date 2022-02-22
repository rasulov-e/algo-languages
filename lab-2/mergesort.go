package main

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	logger.Println("MergeSort: separated array into two parts")
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	logger.Println("MergeSort: started merging arrays")
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	logger.Println("MergeSort: started sorting")
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	logger.Println("MergeSort: merged sorted arrays")
	return final
}
