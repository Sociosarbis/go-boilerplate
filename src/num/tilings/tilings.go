package tilings

func numTilings(n int) int {
	dp := make([]int, n)
	sum := 1
	var mask int = 1e9 + 7
	for i := 0; i < n; i++ {
		if i > 2 {
			sum = (sum + dp[i-3]) % mask
		}
		for j := 0; j < 3; j++ {
			if i >= j {
				prev := 1
				if i > j {
					prev = dp[i-j-1]
				}
				if j == 0 || j == 1 {
					dp[i] += prev
				} else {
					dp[i] += (sum * 2) % mask
				}
				dp[i] %= mask
			}
		}
	}
	return dp[n-1]
}
