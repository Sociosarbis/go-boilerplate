package distance4

func pos(a byte) (int, int) {
	return int(a-'A') / 6, int(a-'A') % 6
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func distance(a, b byte) int {
	i1, j1 := pos(a)
	i2, j2 := pos(b)
	return abs(i1-i2) + abs(j1-j2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDistance(word string) int {
	n := len(word)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = 2700
		}
	}

	dp[0][0] = 0
	dp[1][0] = 0
	dp[0][1] = 0

	for i := 2; i <= n; i++ {
		distanceFromPrev := distance(word[i-1], word[i-2])
		// 如果手指1要放到i
		// 如果手指1在上一个字
		for j := 0; j <= i-2; j++ {
			dp[i][j] = dp[i-1][j] + distanceFromPrev
		}
		// 如果手指2在上一个字
		dp[i][i-1] = min(dp[i][i-1], dp[0][i-1])
		for j := 1; j <= i-2; j++ {
			dp[i][i-1] = min(dp[i][i-1], dp[j][i-1]+distance(word[i-1], word[j-1]))
		}
		// 如果手指2要放到i
		// 如果手指2在上一个字
		for j := 0; j <= i-2; j++ {
			dp[j][i] = dp[j][i-1] + distanceFromPrev
		}
		// 如果手指1在上一个字
		dp[i-1][i] = min(dp[i-1][i], dp[i-1][0])
		for j := 1; j <= i-2; j++ {
			dp[i-1][i] = min(dp[i-1][i], dp[i-1][j]+distance(word[i-1], word[j-1]))
		}
	}
	out := 2700
	for i := n - 1; i >= 0; i-- {
		out = min(min(out, dp[n][i]), dp[i][n])
	}
	return out
}
