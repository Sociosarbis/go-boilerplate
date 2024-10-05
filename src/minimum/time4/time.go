package time4

func minimumTime(time []int, totalTrips int) int64 {
	n := len(time)
	minTime := time[0]
	maxTime := time[0]
	for i := 1; i < n; i++ {
		if time[i] > maxTime {
			maxTime = time[i]
		}
		if time[i] < minTime {
			minTime = time[i]
		}
	}
	l := int64(totalTrips/n) * int64(minTime)
	r := int64(totalTrips/n+1) * int64(maxTime)
	ret := r
	for l <= r {
		mid := (l + r) / 2
		var temp int
		for i := 0; i < n; i++ {
			temp += int(mid / int64(time[i]))
		}
		if temp >= totalTrips {
			if mid < ret {
				ret = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ret
}
