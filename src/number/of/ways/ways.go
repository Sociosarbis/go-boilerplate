package ways

const mask int = 1e9 + 7

func pow(num, x int) int {
	out := num
	for i := 2; i <= x; i++ {
		out *= num
	}
	return out
}

func numberOfWays(n int, x int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		value := pow(i, x)
		end := n - value
		if end < 0 {
			break
		}
		for j := end; j >= 0; j-- {
			dp[j+value] = (dp[j+value] + dp[j]) % mask
		}
	}
	return dp[n]
}
