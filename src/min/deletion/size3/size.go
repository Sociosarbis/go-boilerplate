package size3

func minDeletionSize(strs []string) int {
	m := len(strs)
	n := len(strs[0])
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		max := i
	loop:
		for j := 0; j < i; j++ {
			for k := 0; k < m; k++ {
				if strs[k][i] >= strs[k][j] {
					continue
				} else {
					continue loop
				}
			}
			if dp[j]+i-j-1 < max {
				max = dp[j] + i - j - 1
			}
		}
		dp[i] = max
	}
	out := n - 1
	for i := 1; i < n; i++ {
		if dp[i]+n-i-1 < out {
			out = dp[i] + n - i - 1
		}
	}
	return out
}
