package rectangle

// 记录当前位置向左和向上能到达的位置
type leftTop struct {
	l int
	t int
}

func maximalRectangle(matrix [][]byte) int {
	ret := 0
	dp := make([][]leftTop, len(matrix))
	for i := range dp {
		dp[i] = make([]leftTop, len(matrix[i]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 49 {
				if i > 0 && matrix[i-1][j] == 49 {
					dp[i][j].t = dp[i-1][j].t
				} else {
					dp[i][j].t = i
				}

				if j > 0 && matrix[i][j-1] == 49 {
					dp[i][j].l = dp[i][j-1].l
				} else {
					dp[i][j].l = j
				}

				m, n := dp[i][j].l, i
				for n >= dp[i][j].t && matrix[n][j] == 49 {
					if dp[n][j].l > m {
						m = dp[n][j].l
					}
					tmp := (j - m + 1) * (i - n + 1)
					if tmp > ret {
						ret = tmp
					}
					n--
				}
			}
		}
	}
	return ret
}
