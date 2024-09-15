package points

import "sort"

func numberOfPoints(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	n := len(nums)
	var ret int
	var prev int
	for i := 0; i < n; i++ {
		if nums[i][0] > prev {
			ret += nums[i][1] - nums[i][0] + 1
			prev = nums[i][1]
		} else {
			if nums[i][1] > prev {
				ret += nums[i][1] - prev
				prev = nums[i][1]
			}
		}
	}
	return ret
}
