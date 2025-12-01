package time

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxRunTime(n int, batteries []int) int64 {
	var sum int64
	for _, battery := range batteries {
		sum += int64(battery)
	}
	var l int64
	r := sum / int64(n)
	var out int64
	for l <= r {
		mid := (l + r) / 2
		var total int64
		for _, battery := range batteries {
			total += min(mid, int64(battery))
		}
		if total >= mid*int64(n) {
			if mid > out {
				out = mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return out
}
