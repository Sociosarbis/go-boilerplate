package removal

import "sort"

func minRemoval(nums []int, k int) int {
	n := len(nums)
	var count int
	r := n - 1
	sort.Ints(nums)
	for i := n - 2; i >= 0; i-- {
		for nums[r] > nums[i] && int64(nums[r]) > int64(k)*int64(nums[i]) {
			r--
		}
		if r-i > count {
			count = r - i
		}
	}
	return n - (count + 1)
}
