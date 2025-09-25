package total

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		index := i & 1
		prevIndex := 1 - index
		for j := 0; j <= i; j++ {
			a := dp[prevIndex][j]
			if j > 0 && dp[prevIndex][j-1] < a {
				a = dp[prevIndex][j-1]
			}
			dp[index][j] = a + triangle[i][j]
		}
	}
	var out int = 2e6
	index := (n - 1) & 1
	for i := 0; i < n; i++ {
		if dp[index][i] < out {
			out = dp[index][i]
		}
	}
	return out
}
