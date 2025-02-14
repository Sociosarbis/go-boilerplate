package distance

import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	if n == 2 {
		return position[n-1] - position[0]
	}
	l := 1
	r := (position[n-1] - position[0]) / (m - 1)
	ret := 1
	for l <= r {
		mid := (l + r) >> 1
		p := position[0]
		count := 1
		for i := 1; i < n; i++ {
			if p+mid <= position[i] {
				count++
				p = position[i]
			}
		}
		if count >= m {
			if mid > ret {
				ret = mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ret
}
