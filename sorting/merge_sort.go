package sorting

import "fmt"

// Merge Sort sorts values by recursively dividing array in to two sub arrays until there is only one value per sub array.
// Then performs a merge which compares values on the way up, comparing and merging until a final array of sorted values.
// Note: This implementation is stable, as in it preserves the order of duplicate values.
//
//
// Time complexity O(nlogn)
// Space complexity O(n)

func MergeSort() {
	arr := []int{1, 20, 5, 10, 3, 6}

	arr = sort(arr)
	fmt.Println(arr)
}

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	m := len(arr) / 2
	l := sort(arr[:m])
	r := sort(arr[m:])

	return merge(l, r)
}

func merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}
