package subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
		for j := 0; j < len(text2); j++ {
			dp[i][j] = 0
		}
	}

	for i := 0; i < len(text1); i++ {
		c1 := text1[i]
		for j := 0; j < len(text2); j++ {
			if c1 == text2[j] {
				if i != 0 && j != 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				if i != 0 && (j == 0 || dp[i-1][j] >= dp[i][j-1]) {
					dp[i][j] = dp[i-1][j]
				} else if j != 0 && (i == 0 || dp[i-1][j] < dp[i][j-1]) {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[len(text1)-1][len(text2)-1]
}
