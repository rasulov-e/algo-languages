package main

import (
	"testing"
)

func BenchmarkQuicksort(b *testing.B) {
	b.StopTimer()
	length := 1_000_000
	arr, _ := randomSlice(length)
	b.StartTimer()
	quickSort(arr, 0, length-1)
}

func TestQuicksort(t *testing.T) {
	arr := []int{10, 9, 2, 3, 3}
	sortedArr := []int{2, 3, 3, 9, 10}

	result := quickSort(arr, 0, len(arr)-1)

	for i := 0; i < len(arr); i++ {
		if result[i] != sortedArr[i] {
			t.Error("sorting is wrong")
		}
	}
}
