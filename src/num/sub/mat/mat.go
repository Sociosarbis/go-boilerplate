package mat

func numSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
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
		loop:
			for k := i; k < m; k++ {
				for l := j; l < n; l++ {
					temp := dp[k][l]
					if i > 0 {
						temp -= dp[i-1][l]
						if j > 0 {
							temp -= dp[k][j-1] - dp[i-1][j-1]
						}
					} else if j > 0 {
						temp -= dp[k][j-1]
					}
					if temp == (k-i+1)*(l-j+1) {
						out++
					} else {
						break loop
					}
				}
			}
		}
	}
	return out
}
