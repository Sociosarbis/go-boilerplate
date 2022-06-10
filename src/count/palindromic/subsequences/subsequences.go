package subsequences

import "fmt"

func countPalindromicSubsequences(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	fmt.Println(1000000000)
	for l := 1; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			if l == 1 {
				dp[i][j] = 1
			} else if l == 2 {
				dp[i][j] = 2
			} else {
				if s[i] == s[j] {
					left := i + 1
					right := j - 1
					for left <= right && s[left] != s[i] {
						left++
					}
					for left <= right && s[right] != s[i] {
						right--
					}
					if left > right {
						dp[i][j] = dp[i+1][j-1]*2 + 2
					} else if left == right {
						dp[i][j] = dp[i+1][j-1]*2 + 1
					} else {
						dp[i][j] = dp[i+1][j-1]*2 - dp[left+1][right-1]
					}
				} else {
					dp[i][j] = dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1]
				}
				if dp[i][j] < 0 {
					dp[i][j] += 1000000007
				} else {
					dp[i][j] %= 1000000007
				}
			}
		}
	}
	return dp[0][n-1]
}
