package servings

func soupServings(n int) float64 {
	if n%25 != 0 {
		n = n/25*25 + 25
	}
	if n >= 5000 {
		return 1
	}
	// 最小单位是25
	units := n / 25
	dp := make([][]float64, units+1)
	for i := range dp {
		dp[i] = make([]float64, units+1)
		dp[i][0] = 0
		dp[0][i] = 1
	}
	dp[0][0] = 0.5
	for i := 1; i <= units; i++ {
		for j := 1; j <= units; j++ {
			for k := 0; k < 4; k++ {
				var ii int
				var jj int
				if i > 4-k {
					ii = i - 4 + k
				}
				if j > k {
					jj = j - k
				}
				// 当前的值可从之前的计算得来
				dp[i][j] += dp[ii][jj]
			}
			dp[i][j] *= 0.25
		}
	}
	return dp[units][units]
}
