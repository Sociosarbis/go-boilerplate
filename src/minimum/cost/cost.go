package cost

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	dp := make([][][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([][]int, m)
			for k := 0; k < m; k++ {
				dp[i][j][k] = make([]int, n)
			}
		}
	}

	var end int
	if m > n {
		end = m
	} else {
		end = n
	}

	for k := 1; k < end; k++ {
		for l := 0; l < k; l++ {
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					if i+k < m && j+l < n {
						var value int
						// 从ii上侧切
						for ii := i + 1; ii <= i+k; ii++ {
							temp := dp[i][j][ii-1][j+l] + dp[ii][j][i+k][j+l] + horizontalCut[ii-1]
							if value == 0 || temp < value {
								value = temp
							}
						}
						// 从jj左侧切
						for jj := j + 1; jj <= j+l; jj++ {
							temp := dp[i][j][i+k][jj-1] + dp[i][jj][i+k][j+l] + verticalCut[jj-1]
							if value == 0 || temp < value {
								value = temp
							}
						}
						dp[i][j][i+k][j+l] = value
					}
					if j+k < n && i+l < m {
						var value int
						// 从jj左侧切
						for jj := j + 1; jj <= j+k; jj++ {
							temp := dp[i][j][i+l][jj-1] + dp[i][jj][i+l][j+k] + verticalCut[jj-1]
							if value == 0 || temp < value {
								value = temp
							}
						}
						// 从ii上侧切
						for ii := i + 1; ii <= i+l; ii++ {
							temp := dp[i][j][ii-1][j+k] + dp[ii][j][i+l][j+k] + horizontalCut[ii-1]
							if value == 0 || temp < value {
								value = temp
							}
						}
						dp[i][j][i+l][j+k] = value
					}
				}
			}
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i+k < m && j+k < n {
					var value int
					// 从jj左侧切
					for jj := j + 1; jj <= j+k; jj++ {
						temp := dp[i][j][i+k][jj-1] + dp[i][jj][i+k][j+k] + verticalCut[jj-1]
						if value == 0 || temp < value {
							value = temp
						}
					}
					// 从ii上侧切
					for ii := i + 1; ii <= i+k; ii++ {
						temp := dp[i][j][ii-1][j+k] + dp[ii][j][i+k][j+k] + horizontalCut[ii-1]
						if value == 0 || temp < value {
							value = temp
						}
					}
					dp[i][j][i+k][j+k] = value
				}
			}
		}

	}
	return dp[0][0][m-1][n-1]
}
