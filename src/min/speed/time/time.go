package time

func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if hour <= float64(n-1) {
		return -1
	}
	r := dist[0]
	for i := 1; i < n; i++ {
		if r < dist[i] {
			r = dist[i]
		}
	}
	if r < dist[n-1]*100 {
		r = dist[n-1] * 100
	}
	l := 1
	ret := r
	for l <= r {
		mid := (l + r) / 2
		var temp int
		for i := n - 2; i >= 0; i-- {
			temp += dist[i] / mid
			if dist[i]%mid != 0 {
				temp++
			}
		}
		remain := float64(dist[n-1]) / float64(mid)
		if float64(temp)+remain <= hour {
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
