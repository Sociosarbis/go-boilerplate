package obstacles

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	bfs := [][2]int{{0, 0}}
	for len(bfs) != 0 {
		l := len(bfs)
		for i := 0; i < l; i++ {
			y, x := bfs[i][0], bfs[i][1]
			if y+1 < m {
				if obstacleGrid[y+1][x] == 0 {
					dp[y+1][x] += dp[y][x]
					if dp[y+1][x] == dp[y][x] {
						bfs = append(bfs, [2]int{y + 1, x})
					}
				}
			}
			if x+1 < n {
				if obstacleGrid[y][x+1] == 0 {
					dp[y][x+1] += dp[y][x]
					if dp[y][x+1] == dp[y][x] {
						bfs = append(bfs, [2]int{y, x + 1})
					}
				}
			}
		}
		bfs = bfs[l:]
	}
	return dp[m-1][n-1]
}
