package path

const mask int64 = 1e9 + 7

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var hasZero bool
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				hasZero = true
				break
			}
		}
	}
	dp := make([][][2]int64, 2)
	for i := range dp {
		dp[i] = make([][2]int64, n)
	}
	if grid[0][0] >= 0 {
		dp[0][0][0] = int64(grid[0][0])
	} else {
		dp[0][0][1] = int64(grid[0][0])
	}
	for i := 1; i < n; i++ {
		nums := []int64{int64(grid[0][i]) * dp[0][i-1][0], int64(grid[0][i]) * dp[0][i-1][1]}
		for _, num := range nums {
			if num >= 0 {
				if num > dp[0][i][0] {
					dp[0][i][0] = num
				}
			} else {
				if -num > -dp[0][i][1] {
					dp[0][i][1] = num
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		index := i & 1
		prevIndex := 1 - index
		dp[index][0][0], dp[index][0][1] = 0, 0
		nums := []int64{
			int64(grid[i][0]) * dp[prevIndex][0][0],
			int64(grid[i][0]) * dp[prevIndex][0][1],
		}
		for _, num := range nums {
			if num >= 0 {
				if num > dp[index][0][0] {
					dp[index][0][0] = num
				}
			} else {
				if -num > -dp[index][0][1] {
					dp[index][0][1] = num
				}
			}
		}
		for j := 1; j < n; j++ {
			nums = []int64{
				int64(grid[i][j]) * dp[prevIndex][j][0],
				int64(grid[i][j]) * dp[prevIndex][j][1],
				int64(grid[i][j]) * dp[index][j-1][0],
				int64(grid[i][j]) * dp[index][j-1][1],
			}
			dp[index][j][0], dp[index][j][1] = 0, 0
			for _, num := range nums {
				if num >= 0 {
					if num > dp[index][j][0] {
						dp[index][j][0] = num
					}
				} else {
					if -num > -dp[index][j][1] {
						dp[index][j][1] = num
					}
				}
			}
		}
	}
	index := (m - 1) & 1
	var out int64
	if dp[index][n-1][0] > out {
		out = dp[index][n-1][0]
	}
	if out == 0 && !hasZero {
		return -1
	}
	return int(out % mask)
}
