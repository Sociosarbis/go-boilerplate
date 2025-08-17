package game

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1
	}
	var out float64
	dp := make([]float64, k+1)
	dp[0] = 1
	var denom float64 = 1 / float64(maxPts)
	for i := 0; i < k; i++ {
		for j := 1; j <= maxPts; j++ {
			index := i + j
			if index <= k {
				dp[index] += dp[i] * denom
			}
			if index >= k {
				max := i + maxPts
				if max > n {
					max = n
				}
				out += (dp[i] * denom) * float64(max-k+1)
				break
			}
		}
	}
	return out
}
