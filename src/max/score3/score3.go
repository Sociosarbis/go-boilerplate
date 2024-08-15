package score3

func maxScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var ret int = -1e5
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			var temp int
			if i+1 != m {
				if dp[i+1][j] > temp {
					temp = dp[i+1][j]
				}
			}
			if j+1 != n {
				if dp[i][j+1] > temp {
					temp = dp[i][j+1]
				}
			}
			if temp != 0 {
				if temp-grid[i][j] > ret {
					ret = temp - grid[i][j]
				}
			}
			if grid[i][j] > temp {
				dp[i][j] = grid[i][j]
			} else {
				dp[i][j] = temp
			}
		}
	}
	return ret
}
