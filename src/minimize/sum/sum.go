package sum

import "sort"

func minimizeSum(nums []int) int {
	n := len(nums)
	if n < 4 {
		return 0
	}
	sort.Ints(nums)
	out := nums[n-2] - nums[1]
	if nums[n-1]-nums[2] < out {
		out = nums[n-1] - nums[2]
	}
	if nums[n-3]-nums[0] < out {
		out = nums[n-3] - nums[0]
	}
	return out
}
