package time

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	m, n := len(landStartTime), len(waterStartTime)
	a := 3000
	for i := 0; i < m; i++ {
		if landStartTime[i]+landDuration[i] < a {
			a = landStartTime[i] + landDuration[i]
		}
	}
	out := 3000
	for i := 0; i < n; i++ {
		temp := max(waterStartTime[i], a) + waterDuration[i]
		if temp < out {
			out = temp
		}
	}
	a = 3000
	for i := 0; i < n; i++ {
		if waterStartTime[i]+waterDuration[i] < a {
			a = waterStartTime[i] + waterDuration[i]
		}
	}
	for i := 0; i < m; i++ {
		temp := max(landStartTime[i], a) + landDuration[i]
		if temp < out {
			out = temp
		}
	}
	return out
}
