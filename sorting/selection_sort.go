package sorting

import "fmt"

// Selection Sort sorts values by tracking index of min value and comparing it with rest of the values in given array.
// If a value is lower than current min, it will swap these values until it reaches EOA.
// Note: This implementation is not stable, as in it won't preserve the order of duplicate values due to swap.
//
//
// Time complexity 	O(n^2)
// Space complexity O(1)

func SelectionSort() {
	arr := [6]int{5, 10, 1, 3, 20, 12}
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

	fmt.Println(arr)
}
