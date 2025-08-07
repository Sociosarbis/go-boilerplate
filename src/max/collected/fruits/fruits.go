package fruits

func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	pos := [2]int{}
	for pos[0] < n {
		dp[n-1][n-1] += fruits[pos[0]][pos[1]]
		pos[0]++
		pos[1]++
	}
	dp[n-1][0] = fruits[n-1][0]
	for i := 1; i < n; i++ {
		start := n - 1 - i
		if start < i {
			start = i
		}
		for j := start; j < n; j++ {
			if j != i {
				dp[j][i] = fruits[j][i]
			}
			var temp int
			if j > 0 {
				temp = dp[j-1][i-1]
			}
			if dp[j][i-1] > temp {
				temp = dp[j][i-1]
			}
			if j+1 < n && dp[j+1][i-1] > temp {
				temp = dp[j+1][i-1]
			}
			dp[j][i] += temp
		}
	}
	dp[0][n-1] = fruits[0][n-1]
	for i := 1; i < n; i++ {
		start := n - 1 - i
		if start < i {
			start = i
		}
		for j := start; j < n; j++ {
			if j != i {
				dp[i][j] = fruits[i][j]
			} else if j+1 != n {
				dp[i][j] = 0
			}
			var temp int
			if j > 0 {
				temp = dp[i-1][j-1]
			}
			if dp[i-1][j] > temp {
				temp = dp[i-1][j]
			}
			if j+1 < n && dp[i-1][j+1] > temp {
				temp = dp[i-1][j+1]
			}
			dp[i][j] += temp
		}
	}
	return dp[n-1][n-1]
}
