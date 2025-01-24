package coins

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumCoins(prices []int) int {
	n := len(prices)

	dp := make([][2]int, n+1)
	dp[1][1] = 1e8
	for i, price := range prices {
		max := i + 1
		if i+max >= n {
			max = n - i - 1
		}
		dp[i+1][0] = min(dp[i][1], dp[i+1][1]) + price
		v := dp[i+1][0]
		for j := 1; j <= max; j++ {
			if dp[i+1+j][1] == 0 || dp[i+1+j][1] > v {
				dp[i+1+j][1] = v
			}
		}
	}
	return min(dp[n][0], dp[n][1])
}
