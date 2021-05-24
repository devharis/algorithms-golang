package sorting

import (
	"fmt"
)

func QuickSort() {
	fmt.Println("QuickSort")

	arr := []int{1, 20, 5, 10, 3, 6}
	n := len(arr)

	arr = qsort(arr, 0, n-1)

	fmt.Println(arr)
}

func qsort(arr []int, low int, high int) []int {
	if low < high {
		pi := partition(arr, low, high)

		qsort(arr, low, pi-1)
		qsort(arr, pi+1, high)
	}
	return arr
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, high)
	return i + 1
}

func swap(arr []int, i int, j int) []int {
	arr[i], arr[j] = arr[j], arr[i]
	return arr
}
