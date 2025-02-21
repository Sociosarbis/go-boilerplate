package tiles

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	n := len(floor)
	prefixSum := make([]int, n+1)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		if floor[i] == '1' {
			prefixSum[i+1] = prefixSum[i] + 1
		} else {
			prefixSum[i+1] = prefixSum[i]
		}
		dp[i] = make([]int, numCarpets+1)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][0] = prefixSum[n] - prefixSum[i]
	}
	for j := 1; j <= numCarpets; j++ {
		for i := n - carpetLen - 1; i >= 0; i-- {
			a := dp[i+carpetLen][j-1]
			b := dp[i+1][j]
			if floor[i] == '1' {
				b++
			}
			if a < b {
				dp[i][j] = a
			} else {
				dp[i][j] = b
			}
		}
	}
	return dp[0][numCarpets]
}
