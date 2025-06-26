package subsequence

func longestSubsequence(s string, k int) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = -1
	}
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] != -1 {
				value := (dp[j] << 1) | int(s[i]-'0')
				if value <= k && (dp[j+1] == -1 || value < dp[j+1]) {
					dp[j+1] = value
				}
			}
		}
		if dp[n] != -1 {
			return n
		}
	}
	for i := n; i >= 0; i-- {
		if dp[i] != -1 {
			return i
		}
	}
	return 0
}
