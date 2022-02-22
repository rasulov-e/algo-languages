package main

import (
	"testing"
)

func BenchmarkMergesort(b *testing.B) {
	b.StopTimer()
	length := 1_000_000
	arr, _ := randomSlice(length)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.Run("mergeSort", func(b *testing.B) {
			mergeSort(arr)
		})
	}
}

func TestMergesort(t *testing.T) {
	arr := []int{10, 9, 2, 3, 3}
	sortedArr := []int{2, 3, 3, 9, 10}

	result := mergeSort(arr)

	for i := 0; i < len(arr); i++ {
		if result[i] != sortedArr[i] {
			t.Error("sorting is wrong")
		}
	}
}
