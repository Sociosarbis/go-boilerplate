package value

import (
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var l int
	n := len(arr2)
	var out int
	for i, num := range arr1 {
		for l < n && arr2[l] < num-d {
			l++
		}
		if l == n {
			out += len(arr1) - i
			return out
		}
		if arr2[l] > num+d {
			out++
		}
	}
	return out
}
