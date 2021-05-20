package sorting

import "fmt"

// Insertion Sort sorts values by inserting each value in its sorted order.
// Note: This implementation is stable, as in it preserves the order of duplicate values.
//
//
// Time complexity O(n^2)
// Space complexity O(1)

func InsertionSort() {
	arr := [3]int{1, 20, 5}
	n := len(arr)

	for i := 1; i < n; i++ {
		current := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = current
	}

	fmt.Println(arr)
}
