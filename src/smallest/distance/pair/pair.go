package pair

import (
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	l := nums[1] - nums[0]
	r := nums[n-1] - nums[0]
	for i := 2; i < n; i++ {
		temp := nums[i] - nums[i-1]
		if temp < l {
			l = temp
		}
	}
	ret := 0
	for l <= r {
		mid := (l + r) >> 1
		count := 0
		for i := n - 1; i > 0; i-- {
			target := nums[i] - mid
			ll := 0
			rr := i - 1
			for ll <= rr {
				mm := (ll + rr) >> 1
				if nums[mm] >= target {
					if mm > ll && nums[mm-1] >= target {
						rr = mm - 1
					} else {
						ll = mm
						break
					}
				} else {
					ll = mm + 1
				}
			}
			if ll < i {
				count += i - ll
			}
		}
		if count >= k {
			ret = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ret
}
