package array

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	ret := 1
	start := nums[0]
	for i := 1; i < n; i++ {
		if nums[i]-start > k {
			start = nums[i]
			ret++
		}
	}
	return ret
}
