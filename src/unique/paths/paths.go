package paths

const mod int = 1e9 + 7

func resolveNext(grid [][]int, fromTop bool, m, n, i, j int) []int {
	if i+1 > m || j+1 > n {
		return nil
	}
	if grid[i][j] == 0 {
		return []int{i, j}
	}
	if fromTop {
		return resolveNext(grid, false, m, n, i, j+1)
	}
	return resolveNext(grid, true, m, n, i+1, j)
}

func safeAdd(value *int, c int) {
	*value = (*value + c) % mod
}

func uniquePaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i, row := range dp {
		for j, c := range row {
			if grid[i][j] == 1 {
				continue
			}
			next := resolveNext(grid, false, m, n, i, j+1)
			if next != nil {
				safeAdd(&dp[next[0]][next[1]], c)
			}
			next = resolveNext(grid, true, m, n, i+1, j)
			if next != nil {
				safeAdd(&dp[next[0]][next[1]], c)
			}
		}
	}
	return dp[m-1][n-1]
}
