package averages

import "sort"

func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	n := len(nums)
	half := n / 2
	ret := float64(nums[0]+nums[n-1]) / 2
	for i := 1; i < half; i++ {
		temp := float64(nums[i]+nums[n-1-i]) / 2
		if temp < ret {
			ret = temp
		}
	}
	return ret
}
