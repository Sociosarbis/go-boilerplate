package stairs2

func climbStairs(n int, costs []int) int {
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		for j := 1; j <= 3; j++ {
			if i+j <= n {
				cost := dp[i] + costs[i+j-1] + j*j
				if dp[i+j] == 0 || cost < dp[i+j] {
					dp[i+j] = cost
				}
			}
		}
	}
	return dp[n]
}
