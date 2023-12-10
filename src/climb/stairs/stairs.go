package stairs

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] += dp[i-1]
		if i > 1 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
