package ii

func distinctSubseqII(s string) int {
	M := int(1e9 + 7)
	// 记录以某个字符为结尾的数量
	dp := make([]int, 26)
	temp := 0
	for _, c := range s {
		// 空串(1) + 之前积累的子序列数量
		nextTemp := (temp*2 + 1 - dp[c-'a']) % M
		if nextTemp < 0 {
			nextTemp = nextTemp + M
		}
		nextVal := nextTemp - temp + dp[c-'a']
		if nextVal < 0 {
			nextVal += M
		}
		dp[c-'a'] = nextVal
		temp = nextTemp
	}
	return temp
}
