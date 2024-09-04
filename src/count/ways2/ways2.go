package ways2

import "sort"

func countWays(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	var ret int
	if nums[0] > 0 {
		ret = 1
	}
	for i, num := range nums {
		if i+1 < n && num == nums[i+1] {
			continue
		}
		if num < i+1 && (i+1 >= n || nums[i+1] > i+1) {
			ret++
		}
	}
	return ret
}
