package amount

const min int = -250000000

func reset(dp [][3]int, i int) {
	dp[i][0], dp[i][1], dp[i][2] = min, min, min
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	dp := [2][][3]int{
		make([][3]int, n),
		make([][3]int, n),
	}
	reset(dp[0], 0)
	if coins[0][0] >= 0 {
		dp[0][0][0] = coins[0][0]
	} else {
		dp[0][0][0] = coins[0][0]
		dp[0][0][1] = 0
	}
	for i := 1; i < n; i++ {
		reset(dp[0], i)
		for j := 0; j < 3; j++ {
			dp[0][i][j] = dp[0][i-1][j] + coins[0][i]
		}
		if coins[0][i] < 0 {
			for j := 1; j < 3; j++ {
				dp[0][i][j] = max(dp[0][i][j], dp[0][i-1][j-1])
			}
		}
	}

	for i := 1; i < m; i++ {
		index := i & 1
		reset(dp[index], 0)
		for j := 0; j < 3; j++ {
			dp[index][0][j] = dp[1-index][0][j] + coins[i][0]
		}
		if coins[i][0] < 0 {
			for j := 1; j < 3; j++ {
				dp[index][0][j] = max(dp[index][0][j], dp[1-index][0][j-1])
			}
		}
		for j := 1; j < n; j++ {
			reset(dp[index], j)
			for k := 0; k < 3; k++ {
				dp[index][j][k] = max(dp[index][j-1][k], dp[1-index][j][k]) + coins[i][j]
			}
			if coins[i][j] < 0 {
				for k := 1; k < 3; k++ {
					dp[index][j][k] = max(dp[index][j][k], max(dp[index][j-1][k-1], dp[1-index][j][k-1]))
				}
			}
		}
	}
	index := (m - 1) & 1
	return max(max(dp[index][n-1][0], dp[index][n-1][1]), dp[index][n-1][2])
}
