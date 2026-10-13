package sum

import "sort"

func maxAlternatingSum(nums []int) int64 {
	for i, num := range nums {
		if num < 0 {
			nums[i] = -num
		}
	}
	sort.Ints(nums)
	var l int
	r := len(nums) - 1
	var out int64
	for l < r {
		out += int64(nums[r])*int64(nums[r]) - int64(nums[l])*int64(nums[l])
		l++
		r--
	}
	if l == r {
		out += int64(nums[r]) * int64(nums[r])
	}
	return out
}
