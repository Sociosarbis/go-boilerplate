package cost9

func minCost(grid [][]int) int {
	dp := [2][][1024]bool{}
	n := len(grid[0])
	for i := 0; i < 2; i++ {
		for j := 0; j < n; j++ {
			dp[i] = make([][1024]bool, n)
		}
	}
	temp := grid[0][0]
	dp[0][0][temp] = true
	for i := 1; i < n; i++ {
		temp ^= grid[0][i]
		dp[0][i][temp] = true
	}
	m := len(grid)
	for i := 1; i < m; i++ {
		index := i & 1
		prevIndex := 1 - index
		value := grid[i][0]
		for j := 0; j < 1024; j++ {
			dp[index][0][j] = dp[prevIndex][0][j^value]
		}
		for j := 1; j < n; j++ {
			value = grid[i][j]
			for k := 0; k < 1024; k++ {
				dp[index][j][k] = dp[prevIndex][j][k^value] || dp[index][j-1][k^value]
			}
		}
	}
	index := (m - 1) & 1
	for i := 0; i < 1024; i++ {
		if dp[index][n-1][i] {
			return i
		}
	}
	return 1023
}
