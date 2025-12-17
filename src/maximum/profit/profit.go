package profit

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximumProfit(prices []int, k int) int64 {
	n := len(prices)
	dp := make([][][3]int64, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][3]int64, k+1)
	}
	var out int64
	dp[0][1][1], dp[0][1][2] = int64(-prices[0]), int64(prices[0])
	for i := 1; i < n; i++ {
		index := i & 1
		prevIndex := 1 - index
		dp[index][1][1] = int64(-prices[i])
		dp[index][1][2] = int64(prices[i])
		end := min(k, (i+1)/2)
		for t := end; t > 0; t-- {
			if t < k {
				dp[index][t+1][1] = dp[prevIndex][t][0] - int64(prices[i])
				dp[index][t+1][2] = dp[prevIndex][t][0] + int64(prices[i])
			}
			dp[index][t][0] = max(dp[index][t][0], dp[prevIndex][t][1]+int64(prices[i]))
			dp[index][t][0] = max(dp[index][t][0], dp[prevIndex][t][2]-int64(prices[i]))
			if dp[index][t][0] > out {
				out = dp[index][t][0]
			}
		}
		for t := 1; t <= end; t++ {
			for s := 0; s < 3; s++ {
				dp[index][t][s] = max(dp[index][t][s], dp[prevIndex][t][s])
			}
		}
	}
	return out
}
