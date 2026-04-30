package score

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathScore(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	k = min(k, m+n)
	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}
	dp[0][0][0] = 1
	for i := 0; i < m; i++ {
		index := i % 2
		prevIndex := 1 - index
		var j int
		if i == 0 {
			j = 1
		}
		for ; j < n; j++ {
			var cost, score int
			if grid[i][j] == 1 {
				cost, score = 1, 1
			} else if grid[i][j] == 2 {
				cost, score = 1, 2
			}
			dp[index][j][0] = 0
			for l := k - cost; l >= 0; l-- {
				dp[index][j][l+cost] = 0
				if i > 0 && dp[prevIndex][j][l] != 0 {
					dp[index][j][l+cost] = dp[prevIndex][j][l] + score
				}
				if j > 0 && dp[index][j-1][l] != 0 {
					dp[index][j][l+cost] = max(dp[index][j][l+cost], dp[index][j-1][l]+score)
				}
			}
		}
	}
	var out int
	index := (m - 1) % 2
	for i := 0; i <= k; i++ {
		out = max(out, dp[index][n-1][i])
	}
	return out - 1
}
