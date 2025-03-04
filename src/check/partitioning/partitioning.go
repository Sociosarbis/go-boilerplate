package partitioning

func checkPartitioning(s string) bool {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	var end int
	end = n - 2
	for i := 0; i <= end; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}
	for j := 3; j <= n; j++ {
		end = n - j
		for i := 0; i <= end; i++ {
			if s[i] == s[i+j-1] {
				dp[i][i+j-1] = dp[i+1][i+j-2]

			}
		}
	}
	end = n - 2
	for i := 0; i < end; i++ {
		if dp[0][i] {
			for j := i + 2; j < n; j++ {
				if dp[j][n-1] && dp[i+1][j-1] {
					return true
				}
			}
		}
	}
	return false
}
