package good

import "sort"

func isGood(nums []int) bool {
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] != n-1 {
		return false
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] != i+1 {
			return false
		}
	}
	return true
}
