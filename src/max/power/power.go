package power

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	var lo int64 = 1e9
	var hi int64 = int64(k)
	cnt := make([]int64, n+1)
	for i, station := range stations {
		left := max(0, i-r)
		right := min(n, i+r+1)
		v := int64(station)
		cnt[left] += v
		cnt[right] -= v
		if v < lo {
			lo = v
		}
		hi += v
	}
	var out int64
	diff := make([]int64, n+1)
loop:
	for lo <= hi {
		copy(diff, cnt)
		mid := (lo + hi) / 2
		var sum int64
		remaining := int64(k)
		for i := 0; i < n; i++ {
			sum += diff[i]
			if sum < mid {
				val := mid - sum
				if remaining < val {
					hi = mid - 1
					continue loop
				}
				remaining -= val
				sum = mid
				end := min(n, i+r*2+1)
				diff[end] -= val
			}
		}
		if mid > out {
			out = mid
		}
		lo = mid + 1
	}
	return out
}
