package sign

func orderOfLargestPlusSign(n int, mines [][]int) int {
	mat := make([][]bool, n)
	dp := make([][][]int, n)
	for i := range mat {
		mat[i] = make([]bool, n)
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for _, mine := range mines {
		mat[mine[0]][mine[1]] = true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				if !mat[i][j] {
					dp[i][j][0] = dp[i][j-1][0] + 1
				}
			} else {
				if !mat[i][j] {
					dp[i][j][0] = 1
				}
			}
		}
	}
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if i > 0 {
				if !mat[i][j] {
					dp[i][j][1] = dp[i-1][j][1] + 1
				}
			} else {
				if !mat[i][j] {
					dp[i][j][1] = 1
				}
			}
			if dp[i][j][1] < dp[i][j][0] {
				dp[i][j][0] = dp[i][j][1]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			var prev int
			if j+1 < n {
				prev = dp[i][j+1][1]
			}
			if !mat[i][j] {
				dp[i][j][1] = prev + 1
			} else {
				dp[i][j][1] = 0
			}
			if dp[i][j][1] < dp[i][j][0] {
				dp[i][j][0] = dp[i][j][1]
			}
		}
	}
	for j := 0; j < n; j++ {
		for i := n - 1; i >= 0; i-- {
			var prev int
			if i+1 < n {
				prev = dp[i+1][j][1]
			}
			if !mat[i][j] {
				dp[i][j][1] = prev + 1
			} else {
				dp[i][j][1] = 0
			}
			if dp[i][j][1] < dp[i][j][0] {
				dp[i][j][0] = dp[i][j][1]
			}
		}
	}
	ret := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j][0] > ret {
				ret = dp[i][j][0]
			}
		}
	}
	return ret
}
