package string

func uniqueLetterString(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, 28)
	}
	ret := 0
	for j := 0; j < len(s); j++ {
		for i := 1; j+i <= len(s); i++ {
			if dp[j][27] == 26 && dp[j][26] == 0 {
				break
			}
			c := s[j+i-1] - 65
			dp[j][c]++
			if dp[j][c] == 2 {
				dp[j][26]--
			} else if dp[j][c] == 1 {
				dp[j][26]++
				dp[j][27]++
			}
			ret += dp[j][26]
		}
	}
	return ret
}

func uniqueLetterStringBest(s string) int {
	last := [26][2]int{}
	for i := range last {
		last[i] = [2]int{-1, -1}
	}
	ans := 0
	for i, c := range s {
		c -= 'A'
		// 某个字符成为唯一字符的区间在前后两个相同字符之间
		ans += (len(s) - i) * (i - last[c][1]*2 + last[c][0])
		last[c][0], last[c][1] = last[c][1], i
	}
	return ans
}
