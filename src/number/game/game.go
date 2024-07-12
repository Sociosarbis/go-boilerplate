package game

import "sort"

func numberGame(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}
