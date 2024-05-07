package pickup

func cherryPickup(grid [][]int) int {
	n := len(grid[0])
	m := len(grid)
	dp := [2][][]int{}
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	dp[0][0][n-1] = grid[0][0] + grid[0][n-1]
	for i := 1; i < m; i++ {
		var s int
		if i > n {
			s = n
		} else {
			s = i
		}
		for j := 0; j < s; j++ {
			index := i & 1
			for k := n - s; k < n; k++ {
				prev_val := dp[1-index][j][k]
				for x := -1; x < 2; x++ {
					next_i := j + x
					if next_i >= 0 && next_i < n {
						for y := -1; y < 2; y++ {
							next_j := k + y
							if next_j >= 0 && next_j < n {
								val := prev_val + grid[i][next_i]
								if next_j != next_i {
									val += grid[i][next_j]
								}
								if val > dp[index][next_i][next_j] {
									dp[index][next_i][next_j] = val
								}
							}
						}
					}
				}
			}
		}
	}
	var ret int
	for _, row := range dp[(m-1)&1] {
		for _, v := range row {
			if v > ret {
				ret = v
			}
		}
	}
	return ret
}
