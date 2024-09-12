package indices

import "sort"

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	var l int
	r := n / 2
	var ret int
loop:
	for l <= r {
		mid := (l + r) / 2
		for i := 0; i < mid; i++ {
			if nums[i]*2 > nums[n-mid+i] {
				r = mid - 1
				continue loop
			}
		}
		if mid > ret {
			ret = mid
		}
		l = mid + 1
	}
	return ret * 2
}
