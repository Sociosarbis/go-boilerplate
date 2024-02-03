package vii

func stoneGameVII(stones []int) int {
	n := len(stones)
	dp := make([][][2]int, n)
	prefix_sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix_sums[i+1] = prefix_sums[i] + stones[i]
		dp[i] = make([][2]int, n)
	}
	return dfs(stones, 0, n-1, 0, dp, prefix_sums) - dp[0][n-1][1]
}

func dfs(stones []int, l, r, turn int, dp [][][2]int, prefix_sums []int) int {
	if l >= r {
		return 0
	}
	if dp[l][r][turn] != 0 {
		return dp[l][r][turn]
	}
	option1 := dfs(stones, l+1, r, 1-turn, dp, prefix_sums)
	gain1 := dp[l+1][r][turn] + prefix_sums[r+1] - prefix_sums[l+1]
	option2 := dfs(stones, l, r-1, 1-turn, dp, prefix_sums)
	gain2 := dp[l][r-1][turn] + prefix_sums[r] - prefix_sums[l]
	if gain1-option1 > gain2-option2 {
		dp[l][r][turn] = gain1
		dp[l][r][1-turn] = option1
	} else {
		dp[l][r][turn] = gain2
		dp[l][r][1-turn] = option2
	}
	return dp[l][r][turn]
}
