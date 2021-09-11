package integers

func findIntegers(n int) int {
	dp := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 1,
	}
	if n > 3 {
		i := 3
		for 1<<i <= n {
			mask := 1<<i - 1
			dp[mask] = dp[mask>>1] + 1<<(i-2) + dp[mask>>2]
			rest := n & mask
			if _, err := dp[rest]; !err {
				if rest&(1<<(i-2)) != 0 {
					dp[rest] = dp[1<<(i-1)-1] + rest - (3 << (i - 2)) + 1 + dp[1<<(i-2)-1]
				} else {
					dp[rest] = dp[1<<(i-1)-1] + dp[rest-(1<<(i-1))]
				}
			}
			i += 1
		}
		if n&(1<<(i-2)) != 0 {
			dp[n] = dp[1<<(i-1)-1] + n - (3 << (i - 2)) + 1 + dp[1<<(i-2)-1]
		} else {
			dp[n] = dp[1<<(i-1)-1] + dp[n-(1<<(i-1))]
		}
	}
	return n + 1 - dp[n]
}
