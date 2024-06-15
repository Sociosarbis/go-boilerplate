package beauty

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)

	n := len(nums)
	start := 0
	ret := 1
	for i := 1; i < n; i++ {
		l := start
		r := i
		min := nums[i] - k
		for l <= r {
			mid := (l + r) / 2
			if nums[mid]+k < min {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if i-l+1 > ret {
			ret = i - l + 1
		}
		start = l
	}
	return ret
}
