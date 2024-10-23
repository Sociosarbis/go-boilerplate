package pairs2

func countCompleteDayPairs(hours []int) int64 {
	counter := [24]int64{}
	for _, hour := range hours {
		counter[hour%24]++
	}

	ret := counter[0]*(counter[0]-1)/2 + counter[12]*(counter[12]-1)/2
	for i := 1; i < 12; i++ {
		ret += counter[i] * counter[24-i]
	}
	return ret
}
