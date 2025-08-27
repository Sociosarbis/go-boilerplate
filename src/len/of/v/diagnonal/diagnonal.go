package diagnonal

var dirs = [4][2]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

func dfs(grid [][]int, dp [][][4][2]int, i, j, d, idx int) int {
	if dp[i][j][d][idx] != 0 {
		return dp[i][j][d][idx]
	}
	num := grid[i][j]
	var temp int
	for k, dir := range dirs {
		if (idx == 0 && (d+1)%4 == k) || d == k {
			if i+dir[0] >= 0 && i+dir[0] < len(grid) && j+dir[1] >= 0 && j+dir[1] < len(grid[0]) {
				if 2-num == grid[i+dir[0]][j+dir[1]] {
					var value int
					if k == d {
						value = dfs(grid, dp, i+dir[0], j+dir[1], k, idx)
					} else if idx == 0 {
						value = dfs(grid, dp, i+dir[0], j+dir[1], k, 1)
					}
					if value > temp {
						temp = value
					}
				}
			}
		}
	}
	dp[i][j][d][idx] = temp + 1
	return temp + 1
}

func lenOfVDiagonal(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][4][2]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][4][2]int, n)
	}
	var out int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				var temp int
				for k, dir := range dirs {
					if i+dir[0] >= 0 && i+dir[0] < m && j+dir[1] >= 0 && j+dir[1] < n {
						if grid[i+dir[0]][j+dir[1]] == 2 {
							value := dfs(grid, dp, i+dir[0], j+dir[1], k, 0)
							if value > temp {
								temp = value
							}
						}
					}
				}
				if temp+1 > out {
					out = temp + 1
				}
			}
		}
	}
	return out
}
