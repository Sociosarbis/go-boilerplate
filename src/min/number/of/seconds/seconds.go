package seconds

import "math"

func findX(workerTime int64, totalTime int64) int64 {
	target := totalTime * 2
	root := float64(target) / float64(workerTime)
	num := int64(math.Sqrt(root)) + 1
	for i := num + 1; i > 0; i-- {
		if (i+1)*i*workerTime <= target {
			return i
		}
	}
	return 0
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	out := int64(1+mountainHeight) * int64(mountainHeight) * int64(workerTimes[0]) / 2
	var l int64 = 1
	r := out
	for l <= r {
		mid := l + (r-l)/2
		var temp int64
		for _, workerTime := range workerTimes {
			temp += findX(int64(workerTime), mid)
		}
		if temp >= int64(mountainHeight) {
			out = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return out
}
