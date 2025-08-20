package squares

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = 1
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
				if j > 0 {
					dp[i][j] += dp[i][j-1] - dp[i-1][j-1]
				}
			} else if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	var out int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				out++
				for k := 1; i+k < m && j+k < n; k++ {
					temp := dp[i+k][j+k]
					if i > 0 {
						temp -= dp[i-1][j+k]
						if j > 0 {
							temp -= dp[i+k][j-1] - dp[i-1][j-1]
						}
					} else if j > 0 {
						temp -= dp[i+k][j-1]
					}
					if temp == (k+1)*(k+1) {
						out++
					} else {
						break
					}
				}
			}
		}
	}
	return out
}
