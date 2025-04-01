package points

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n)
	var base int64
	var ret int64
	for i, q := range questions {
		points, brainpower := q[0], q[1]
		if base < dp[i] {
			base = dp[i]
		}
		temp := base + int64(points)
		next := i + brainpower + 1
		if next < n {
			if dp[next] < temp {
				dp[next] = temp
			}
		}
		if temp > ret {
			ret = temp
		}
	}
	return ret
}
