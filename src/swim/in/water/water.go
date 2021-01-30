package water

func swimInWater(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = grid[0][0]
	bfs := [][]int{{0, 0}}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for len(bfs) != 0 {
		cur := bfs[len(bfs)-1]
		prevDp := dp[cur[0]][cur[1]]
		bfs = bfs[0 : len(bfs)-1]
		for _, dir := range dirs {
			nextI := cur[0] + dir[0]
			nextJ := cur[1] + dir[1]
			if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n {
				nextVal := grid[nextI][nextJ]
				nextDp := prevDp
				if nextVal > nextDp {
					nextDp = nextVal
				}
				if dp[nextI][nextJ] == -1 || dp[nextI][nextJ] > nextDp {
					dp[nextI][nextJ] = nextDp
					bfs = append(bfs, []int{nextI, nextJ})
				}
			}
		}
	}
	return dp[m-1][n-1]
}
