package partition

import "sort"

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ret := nums[1] - nums[0]
	for i := 2; i < n; i++ {
		temp := nums[i] - nums[i-1]
		if temp < ret {
			ret = temp
		}
	}
	return ret
}
