package move2

import (
	"sort"
)

// 中位数是最短距离
func minMoves2(nums []int) int {
	sort.Ints(nums)

	mid := len(nums) / 2
	medium := nums[len(nums)/2]

	i := 0

	ret := 0
	for ; i < mid; i++ {
		ret += medium - nums[i]
	}

	for ; i < len(nums); i++ {
		ret += nums[i] - medium
	}
	return ret
}
