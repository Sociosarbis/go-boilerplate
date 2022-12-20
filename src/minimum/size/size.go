package size

import "sort"

func minimumSize(nums []int, maxOperations int) int {
	sort.Ints(nums)
	n := len(nums)
	l := 1
	r := nums[n-1]
	for l <= r {
		mid := (l + r) / 2
		ll := 0
		rr := n - 1
		for ll <= rr {
			mm := (ll + rr) / 2
			if nums[mm] > mid {
				if mm > 0 && nums[mm-1] > mid {
					rr = mm - 1
				} else {
					ll = mm
					break
				}
			} else {
				ll = mm + 1
			}
		}
		temp := 0
		for i := ll; i < n; i++ {
			if nums[i]%mid == 0 {
				temp += nums[i]/mid - 1
			} else {
				temp += nums[i] / mid
			}
		}
		if temp > maxOperations {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
