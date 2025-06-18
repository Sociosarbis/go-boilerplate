package array

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ret := make([][]int, 0, n/3)
	for i := 0; i < n; i += 3 {
		if nums[i+2]-nums[i] <= k {
			ret = append(ret, nums[i:i+3])
		} else {
			return [][]int{}
		}
	}
	return ret
}
