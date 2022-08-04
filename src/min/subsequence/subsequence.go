package subsequence

import (
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	total := 0
	for _, num := range nums {
		total += num
	}

	temp := 0

	for i := len(nums) - 1; i >= 0; i-- {
		temp += nums[i]
		if temp<<1 > total {
			l := i
			r := len(nums) - 1
			for l < r {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			}
			return nums[i:]
		}
	}
	return nums
}
