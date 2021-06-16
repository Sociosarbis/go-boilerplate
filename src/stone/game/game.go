package game

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	acc := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	acc[0] = piles[0]
	for i := 1; i < n; i++ {
		acc[i] = acc[i-1] + piles[i]
	}
	for i := 0; i < n; i++ {
		dp[i][i] = piles[i]
		for j := i - 1; j >= 0; j-- {
			opt1 := acc[i-1] - dp[j][i-1] + piles[i]
			if j > 0 {
				opt1 -= acc[j-1]
			}
			opt2 := acc[i] - acc[j] - dp[j+1][i] + piles[j]
			if opt1 > opt2 {
				dp[j][i] = opt1
			} else {
				dp[j][i] = opt2
			}
		}
	}
	return dp[0][n-1] > acc[n-1]-dp[0][n-1]
}
