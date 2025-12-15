package periods

func getDescentPeriods(prices []int) int64 {
	var start int
	n := len(prices)
	var out int64
	for i := 1; i < n; i++ {
		if prices[i]+1 != prices[i-1] {
			out += int64(i-start) * int64(i-start+1) / 2
			start = i
		}
	}
	out += int64(n-start) * int64(n-start+1) / 2
	return out
}
