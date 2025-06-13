package max

import "sort"

func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	var l int
	r := nums[n-1] - nums[0]
	out := r
loop:
	for l <= r {
		mid := (l + r) >> 1
		i := 1
		var c int
		for i < n {
			if nums[i]-nums[i-1] <= mid {
				c++
				if c >= p {
					out = mid
					r = mid - 1
					continue loop
				}
				i += 2
			}
		}
		l = mid + 1
	}
	return out
}
