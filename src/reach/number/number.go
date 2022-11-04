package number

import "math"

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	l := 1
	r := int(math.Sqrt(float64(target * 2)))
	for l <= r {
		mid := (l + r) / 2
		sum := (mid + 1) * mid / 2
		if sum < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	for {
		diff := (l+1)*l/2 - target
		if diff%2 == 0 {
			return l
		}
		l++
	}
}
