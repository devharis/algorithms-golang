package sorting

import "fmt"

// Bubble Sort sorts values in pairs of two, adjacent elements.
// Note: This implementation is stable, as in it preserves the order of duplicate values.
//
//
// Time complexity O(n^2)
// Space complexity O(1)

func BubbleSort() {
	arr := [10]int{1, 20, 5, 15, 19, 3, 8, 17, 6, 12}
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
