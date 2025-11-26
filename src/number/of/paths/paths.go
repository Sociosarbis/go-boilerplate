package paths

const mask int = 1e9 + 7

func numberOfPaths(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][][]int, 2)
	tmpl := make([]int, k)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k)
		}
	}
	dp[0][0][grid[0][0]%k]++
	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			if dp[0][i-1][j] != 0 {
				dp[0][i][(j+grid[0][i])%k]++
				break
			}
		}
	}
	for i := 1; i < m; i++ {
		index := i % 2
		prevIndex := 1 - index
		copy(dp[index][0], tmpl)
		for j := 0; j < k; j++ {
			if dp[prevIndex][0][j] != 0 {
				dp[index][0][(j+grid[i][0])%k]++
				break
			}
		}
		for j := 1; j < n; j++ {
			copy(dp[index][j], tmpl)
			for ii := 0; ii < k; ii++ {
				if dp[prevIndex][j][ii] != 0 {
					dp[index][j][(ii+grid[i][j])%k] = (dp[index][j][(ii+grid[i][j])%k] + dp[prevIndex][j][ii]) % mask
				}
				if dp[index][j-1][ii] != 0 {
					dp[index][j][(ii+grid[i][j])%k] = (dp[index][j][(ii+grid[i][j])%k] + dp[index][j-1][ii]) % mask
				}
			}
		}
	}
	return dp[(m-1)%2][n-1][0]
}
