package viii

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stoneGameVIII(stones []int) int {
	n := len(stones)
	pre := make([]int, n)
	pre[0] = stones[0]
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] + stones[i]
	}
	// 在[i, u)时选择移除时，a和b的最大差值
	dp := make([]int, n)
	dp[n-1] = pre[n-1]
	for i := n - 2; i > 0; i-- {
		dp[i] = max(dp[i+1], pre[i]-dp[i+1])
	}
	return dp[1]
}
