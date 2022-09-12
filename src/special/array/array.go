package array

import (
	"sort"
)

func specialArray(nums []int) int {
	sort.Ints(nums)
	ret := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= ret {
			if (i > 0 && nums[i-1] < ret) || i == 0 {
				return ret
			}
		} else {
			break
		}
		ret++
	}
	return -1
}
