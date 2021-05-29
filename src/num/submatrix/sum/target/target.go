package target

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	ret := 0
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res := matrix[i][j]
			if i > 0 {
				res += dp[i-1][j]
			}
			if j > 0 {
				res += dp[i][j-1]
			}
			if i > 0 && j > 0 {
				res -= dp[i-1][j-1]
			}
			dp[i][j] = res
			for k := i; k >= 0; k-- {
				for l := j; l >= 0; l-- {
					res := dp[i][j]
					if k < i {
						res -= dp[k][j]
					}
					if l < j {
						res -= dp[i][l]
					}
					if k < i && l < j {
						res += dp[k][l]
					}
					if res == target {
						ret++
					}
				}
			}
		}
	}

	return ret
}
