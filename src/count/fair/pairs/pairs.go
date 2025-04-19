package pairs

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	n := len(nums)
	var ret int64
	l := 1
	r := n - 1
	for i := 0; i < r; i++ {
		if l <= i {
			l = i + 1
		}
		for nums[l]+nums[i] < lower && l+1 <= r {
			l++
		}
		for l-1 > i && nums[l-1]+nums[i] >= lower {
			l--
		}
		for nums[r]+nums[i] > upper && r > l {
			r--
		}

		if nums[l]+nums[i] >= lower {
			if nums[r]+nums[i] <= upper {
				ret += int64(r - l + 1)
			}
		}
	}
	return ret
}
