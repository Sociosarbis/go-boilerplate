package tribonacci

func tribonacci(n int) int {
	dp := []int{0, 1, 1}
	for i := 3; i <= n; i += 1 {
		dp[i%3] = dp[0] + dp[1] + dp[2]
	}
	return dp[n%3]
}
