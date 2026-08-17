package v

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	prefixSum := make([]int, n+1)
	for i, value := range stoneValue {
		prefixSum[i+1] = prefixSum[i] + value
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 2; i <= n; i++ {
		for j := 0; j+i <= n; j++ {
			var max int
			for k := i - 2; k >= 0; k-- {
				a, b := prefixSum[j+k+1]-prefixSum[j], prefixSum[j+i]-prefixSum[j+k+1]
				var temp int
				if a > b {
					temp = b + dp[j+k+1][j+i-1]
				} else {
					temp = a + dp[j][j+k]
					if a == b && b+dp[j+k+1][j+i-1] > temp {
						temp = b + dp[j+k+1][j+i-1]
					}
				}
				if temp > max {
					max = temp
				}
			}
			dp[j][j+i-1] = max
		}
	}
	return dp[0][n-1]
}
